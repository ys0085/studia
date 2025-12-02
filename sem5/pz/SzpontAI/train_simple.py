"""
Train a classifier using MediaPipe landmarks directly from images.
Simplified approach: detect landmarks -> train neural network -> classify behaviors
"""

import os
import cv2
import torch
import numpy as np
import mediapipe as mp
from torch import nn
from torch.optim import Adam
from torch.utils.data import Dataset, DataLoader
from tqdm import tqdm


# Initialize MediaPipe
BaseOptions = mp.tasks.BaseOptions
FaceLandmarker = mp.tasks.vision.FaceLandmarker
FaceLandmarkerOptions = mp.tasks.vision.FaceLandmarkerOptions
PoseLandmarker = mp.tasks.vision.PoseLandmarker
PoseLandmarkerOptions = mp.tasks.vision.PoseLandmarkerOptions
HandLandmarker = mp.tasks.vision.HandLandmarker
HandLandmarkerOptions = mp.tasks.vision.HandLandmarkerOptions
VisionRunningMode = mp.tasks.vision.RunningMode


class DriverBehaviorClassifier(nn.Module):
    """Simple neural network for driver behavior classification"""
    
    def __init__(self, input_size, num_classes):
        super().__init__()
        self.network = nn.Sequential(
            nn.Linear(input_size, 256),
            nn.ReLU(),
            nn.Dropout(0.3),
            nn.Linear(256, 128),
            nn.ReLU(),
            nn.Dropout(0.3),
            nn.Linear(128, num_classes)
        )
    
    def forward(self, x):
        return self.network(x)


def extract_landmarks_from_image(image_path, face_detector, pose_detector, hand_detector):
    """Extract landmarks from a single image using MediaPipe (following Google's pattern)"""
    
    # Read and prepare image
    image = cv2.imread(image_path)
    if image is None:
        return None
    
    # Convert BGR to RGB as MediaPipe expects RGB
    image_rgb = cv2.cvtColor(image, cv2.COLOR_BGR2RGB)
    
    # Create MediaPipe Image object
    mp_image = mp.Image(image_format=mp.ImageFormat.SRGB, data=image_rgb)
    
    features = []
    
    # Face landmarks detection
    try:
        face_result = face_detector.detect(mp_image)
        if face_result.face_landmarks and len(face_result.face_landmarks) > 0:
            # Get first face's landmarks
            for landmark in face_result.face_landmarks[0]:
                features.extend([landmark.x, landmark.y, landmark.z])
        else:
            # No face detected - use zeros
            features.extend([0.0] * (478 * 3))
    except Exception:
        features.extend([0.0] * (478 * 3))
    
    # Pose landmarks detection
    try:
        pose_result = pose_detector.detect(mp_image)
        if pose_result.pose_landmarks and len(pose_result.pose_landmarks) > 0:
            # Get pose landmarks
            for landmark in pose_result.pose_landmarks[0]:
                features.extend([landmark.x, landmark.y, landmark.z])
        else:
            features.extend([0.0] * (33 * 3))
    except Exception:
        features.extend([0.0] * (33 * 3))
    
    # Hand landmarks detection
    try:
        hand_result = hand_detector.detect(mp_image)
        if hand_result.hand_landmarks and len(hand_result.hand_landmarks) > 0:
            # Process up to 2 hands
            num_hands = min(len(hand_result.hand_landmarks), 2)
            for i in range(num_hands):
                for landmark in hand_result.hand_landmarks[i]:
                    features.extend([landmark.x, landmark.y, landmark.z])
            # Pad remaining hands with zeros
            for i in range(2 - num_hands):
                features.extend([0.0] * (21 * 3))
        else:
            # No hands detected
            features.extend([0.0] * (21 * 3 * 2))
    except Exception:
        features.extend([0.0] * (21 * 3 * 2))
    
    return np.array(features, dtype=np.float32)


def load_dataset(dataset_path, face_detector, pose_detector, hand_detector):
    """Load and process entire dataset"""
    
    X = []
    y = []
    classes = []
    
    print("\nLoading dataset...")
    
    for class_idx, class_name in enumerate(sorted(os.listdir(dataset_path))):
        class_path = os.path.join(dataset_path, class_name)
        
        if not os.path.isdir(class_path):
            continue
        
        classes.append(class_name)
        print(f"\nProcessing class: {class_name}")
        
        image_files = [f for f in os.listdir(class_path) 
                      if f.lower().endswith(('.jpg', '.jpeg', '.png', '.bmp'))]
        
        for img_file in tqdm(image_files, desc=f"  {class_name}"):
            img_path = os.path.join(class_path, img_file)
            landmarks = extract_landmarks_from_image(img_path, face_detector, pose_detector, hand_detector)
            
            if landmarks is not None:
                X.append(landmarks)
                y.append(class_idx)
    
    return np.array(X), np.array(y), classes


class LandmarkDataset(Dataset):
    """PyTorch dataset for landmarks"""
    
    def __init__(self, X, y):
        self.X = torch.FloatTensor(X)
        self.y = torch.LongTensor(y)
    
    def __len__(self):
        return len(self.X)
    
    def __getitem__(self, idx):
        return self.X[idx], self.y[idx]


if __name__ == "__main__":
    
    # Configuration
    DATASET_PATH = "Multi-Class Driver Behavior Image Dataset"
    MODEL_SAVE_PATH = "saved_models/driver_behavior_classifier.pth"
    EPOCHS = 200
    BATCH_SIZE = 32
    LEARNING_RATE = 0.0001
    
    # Initialize MediaPipe detectors
    print("Initializing MediaPipe...")
    
    face_options = FaceLandmarkerOptions(
        base_options=BaseOptions(model_asset_path='face_landmarker.task'),
        running_mode=VisionRunningMode.IMAGE,
        num_faces=1
    )
    
    pose_options = PoseLandmarkerOptions(
        base_options=BaseOptions(model_asset_path='pose_landmarker_lite.task'),
        running_mode=VisionRunningMode.IMAGE,
        num_poses=1
    )
    
    hand_options = HandLandmarkerOptions(
        base_options=BaseOptions(model_asset_path='hand_landmarker.task'),
        running_mode=VisionRunningMode.IMAGE,
        num_hands=2
    )
    
    face_detector = FaceLandmarker.create_from_options(face_options)
    pose_detector = PoseLandmarker.create_from_options(pose_options)
    hand_detector = HandLandmarker.create_from_options(hand_options)
    
    print("✓ MediaPipe initialized")
    
    # Load dataset and extract landmarks
    X, y, classes = load_dataset(DATASET_PATH, face_detector, pose_detector, hand_detector)
    
    print(f"\nDataset loaded:")
    print(f"  Samples: {len(X)}")
    print(f"  Classes: {classes}")
    print(f"  Features per sample: {X.shape[1]}")
    
    # Split train/test
    from sklearn.model_selection import train_test_split
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42, stratify=y)
    
    train_dataset = LandmarkDataset(X_train, y_train)
    test_dataset = LandmarkDataset(X_test, y_test)
    
    train_loader = DataLoader(train_dataset, batch_size=BATCH_SIZE, shuffle=True)
    test_loader = DataLoader(test_dataset, batch_size=BATCH_SIZE, shuffle=False)
    
    # Create model
    device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    print(f"\nUsing device: {device}")
    
    model = DriverBehaviorClassifier(X.shape[1], len(classes)).to(device)
    criterion = nn.CrossEntropyLoss()
    optimizer = Adam(model.parameters(), lr=LEARNING_RATE)

    # Train
    print("\nTraining...")
    best_acc = 0.0
    
    for epoch in range(EPOCHS):
        model.train()
        train_loss = 0.0
        train_correct = 0
        train_total = 0
        
        for batch_x, batch_y in train_loader:
            batch_x, batch_y = batch_x.to(device), batch_y.to(device)
            
            optimizer.zero_grad()
            outputs = model(batch_x)
            loss = criterion(outputs, batch_y)
            loss.backward()
            optimizer.step()
            
            train_loss += loss.item()
            _, predicted = torch.max(outputs, 1)
            train_total += batch_y.size(0)
            train_correct += (predicted == batch_y).sum().item()
        
        train_acc = 100 * train_correct / train_total

        
        # Validate
        model.eval()
        test_correct = 0
        test_total = 0
        
        with torch.no_grad():
            for batch_x, batch_y in test_loader:
                batch_x, batch_y = batch_x.to(device), batch_y.to(device)
                outputs = model(batch_x)
                _, predicted = torch.max(outputs, 1)
                test_total += batch_y.size(0)
                test_correct += (predicted == batch_y).sum().item()
        
        test_acc = 100 * test_correct / test_total
        
        print(f"Epoch [{epoch+1}/{EPOCHS}] Train Acc: {train_acc:.2f}% | Test Acc: {test_acc:.2f}%")
        
        # Save best model
        if test_acc > best_acc:
            best_acc = test_acc
            os.makedirs("saved_models", exist_ok=True)
            
            torch.save({
                'model_state_dict': model.state_dict(),
                'classes': classes,
                'input_size': X.shape[1],
                'accuracy': best_acc
            }, MODEL_SAVE_PATH)
            
            print(f"  ✓ Saved best model (accuracy: {best_acc:.2f}%)")
    
    print(f"\n✓ Training complete! Best accuracy: {best_acc:.2f}%")
    print(f"Model saved to: {MODEL_SAVE_PATH}")
    
    # Cleanup
    face_detector.close()
    pose_detector.close()
    hand_detector.close()
