"""
Real-time driver behavior detection using MediaPipe landmarks.
Implementation following Google's MediaPipe samples pattern.
Reference: https://github.com/google-ai-edge/mediapipe-samples
"""

import os
import cv2
import torch
import numpy as np
import mediapipe as mp
from mediapipe import solutions
from mediapipe.framework.formats import landmark_pb2


# MediaPipe API imports (following official pattern)
BaseOptions = mp.tasks.BaseOptions
FaceLandmarker = mp.tasks.vision.FaceLandmarker
FaceLandmarkerOptions = mp.tasks.vision.FaceLandmarkerOptions
PoseLandmarker = mp.tasks.vision.PoseLandmarker
PoseLandmarkerOptions = mp.tasks.vision.PoseLandmarkerOptions
HandLandmarker = mp.tasks.vision.HandLandmarker
HandLandmarkerOptions = mp.tasks.vision.HandLandmarkerOptions
VisionRunningMode = mp.tasks.vision.RunningMode


class DriverBehaviorClassifier(torch.nn.Module):
    """Simple neural network for driver behavior classification"""
    
    def __init__(self, input_size, num_classes):
        super().__init__()
        self.network = torch.nn.Sequential(
            torch.nn.Linear(input_size, 256),
            torch.nn.ReLU(),
            torch.nn.Dropout(0.3),
            torch.nn.Linear(256, 128),
            torch.nn.ReLU(),
            torch.nn.Dropout(0.3),
            torch.nn.Linear(128, num_classes)
        )
    
    def forward(self, x):
        return self.network(x)


def extract_landmarks_from_frame(frame, timestamp_ms, face_detector, pose_detector, hand_detector):
    """
    Extract landmarks from video frame using VIDEO mode.
    Following Google's MediaPipe samples pattern.
    """
    
    # Convert BGR to RGB (MediaPipe expects RGB)
    image_rgb = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
    
    # Create MediaPipe Image object
    mp_image = mp.Image(image_format=mp.ImageFormat.SRGB, data=image_rgb)
    
    features = []
    
    # Face landmarks - using detect_for_video with timestamp
    try:
        face_result = face_detector.detect_for_video(mp_image, timestamp_ms)
        if face_result.face_landmarks and len(face_result.face_landmarks) > 0:
            for landmark in face_result.face_landmarks[0]:
                features.extend([landmark.x, landmark.y, landmark.z])
        else:
            features.extend([0.0] * (478 * 3))
    except Exception:
        features.extend([0.0] * (478 * 3))
    
    # Pose landmarks - using detect_for_video with timestamp
    try:
        pose_result = pose_detector.detect_for_video(mp_image, timestamp_ms)
        if pose_result.pose_landmarks and len(pose_result.pose_landmarks) > 0:
            for landmark in pose_result.pose_landmarks[0]:
                features.extend([landmark.x, landmark.y, landmark.z])
        else:
            features.extend([0.0] * (33 * 3))
    except Exception:
        features.extend([0.0] * (33 * 3))
    
    # Hand landmarks - using detect_for_video with timestamp
    try:
        hand_result = hand_detector.detect_for_video(mp_image, timestamp_ms)
        if hand_result.hand_landmarks and len(hand_result.hand_landmarks) > 0:
            num_hands = min(len(hand_result.hand_landmarks), 2)
            for i in range(num_hands):
                for landmark in hand_result.hand_landmarks[i]:
                    features.extend([landmark.x, landmark.y, landmark.z])
            for i in range(2 - num_hands):
                features.extend([0.0] * (21 * 3))
        else:
            features.extend([0.0] * (21 * 3 * 2))
    except Exception:
        features.extend([0.0] * (21 * 3 * 2))
    
    return np.array(features, dtype=np.float32)


def main():
    MODEL_PATH = "saved_models/driver_behavior_classifier.pth"
    
    if not os.path.exists(MODEL_PATH):
        print(f"ERROR: Model not found at {MODEL_PATH}")
        print("Please run: python train_simple.py first")
        return
    
    # Load model
    print("Loading model...")
    checkpoint = torch.load(MODEL_PATH, map_location='cpu')
    classes = checkpoint['classes']
    input_size = checkpoint['input_size']
    
    device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    model = DriverBehaviorClassifier(input_size, len(classes)).to(device)
    model.load_state_dict(checkpoint['model_state_dict'])
    model.eval()
    
    print(f"✓ Model loaded (accuracy: {checkpoint['accuracy']:.2f}%)")
    print(f"Classes: {classes}")
    
    # Initialize MediaPipe with VIDEO mode (following official pattern)
    print("\nInitializing MediaPipe with VIDEO mode...")
    
    face_options = FaceLandmarkerOptions(
        base_options=BaseOptions(model_asset_path='face_landmarker.task'),
        running_mode=VisionRunningMode.VIDEO,  # VIDEO mode for webcam
        num_faces=1,
        min_face_detection_confidence=0.5,
        min_face_presence_confidence=0.5,
        min_tracking_confidence=0.5
    )
    
    pose_options = PoseLandmarkerOptions(
        base_options=BaseOptions(model_asset_path='pose_landmarker_lite.task'),
        running_mode=VisionRunningMode.VIDEO,  # VIDEO mode for webcam
        num_poses=1,
        min_pose_detection_confidence=0.5,
        min_pose_presence_confidence=0.5,
        min_tracking_confidence=0.5
    )
    
    hand_options = HandLandmarkerOptions(
        base_options=BaseOptions(model_asset_path='hand_landmarker.task'),
        running_mode=VisionRunningMode.VIDEO,  # VIDEO mode for webcam
        num_hands=2,
        min_hand_detection_confidence=0.5,
        min_hand_presence_confidence=0.5,
        min_tracking_confidence=0.5
    )
    
    face_detector = FaceLandmarker.create_from_options(face_options)
    pose_detector = PoseLandmarker.create_from_options(pose_options)
    hand_detector = HandLandmarker.create_from_options(hand_options)
    
    print("✓ MediaPipe initialized")
    
    # Open webcam
    cap = cv2.VideoCapture(0, cv2.CAP_DSHOW)
    if not cap.isOpened():
        print("ERROR: Cannot open webcam")
        return
    
    print("\nStarting detection...")
    print("Press 'q' to quit")
    
    # Frame counter for timestamp calculation
    frame_count = 0
    
    while True:
        ret, frame = cap.read()
        if not ret:
            break
        
        # Calculate timestamp in milliseconds (required for VIDEO mode)
        timestamp_ms = int(frame_count * 1000 / 30)  # Assuming 30 FPS
        frame_count += 1
        
        try:
            # Extract landmarks using VIDEO mode (with timestamp)
            landmarks = extract_landmarks_from_frame(
                frame, timestamp_ms, face_detector, pose_detector, hand_detector
            )
            
            # Predict behavior
            x = torch.FloatTensor(landmarks).unsqueeze(0).to(device)
            with torch.no_grad():
                outputs = model(x)
                probs = torch.softmax(outputs, dim=1)[0]
                pred_idx = torch.argmax(outputs, dim=1).item()
            
            # Display results
            label = classes[pred_idx]
            confidence = probs[pred_idx].item() * 100
            
            # Main prediction
            cv2.putText(frame, f"{label}: {confidence:.1f}%", (10, 50),
                       cv2.FONT_HERSHEY_SIMPLEX, 1.2, (0, 0, 139), 3, cv2.LINE_AA)
            
            # All probabilities
            y_offset = 100
            for i, class_name in enumerate(classes):
                prob = probs[i].item() * 100
                text = f"{class_name}: {prob:.1f}%"
                color = (0, 0, 139) if i == pred_idx else (0, 0, 100)
                cv2.putText(frame, text, (10, y_offset),
                           cv2.FONT_HERSHEY_SIMPLEX, 0.7, color, 2, cv2.LINE_AA)
                y_offset += 35
            
        except Exception as e:
            cv2.putText(frame, f"Error: {str(e)}", (10, 50),
                       cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 255), 2, cv2.LINE_AA)
        
        cv2.imshow("Driver Behavior Detection - Press q to quit", frame)
        
        if cv2.waitKey(1) & 0xFF == ord('q'):
            break
    
    # Cleanup
    cap.release()
    cv2.destroyAllWindows()
    face_detector.close()
    pose_detector.close()
    hand_detector.close()
    
    print("\n✓ Done!")


if __name__ == "__main__":
    main()
