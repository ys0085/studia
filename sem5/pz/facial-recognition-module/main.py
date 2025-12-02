import os
import cv2
import torch
import torchvision.transforms as transforms
from PIL import Image
from model import FaceNet  # imports only the class; model.py won't retrain because of __main__ guard

# Load checkpoint
save_dir = os.path.join(os.path.dirname(__file__), "saved_models")
checkpoint_path = os.path.join(save_dir, "facenet.pth")
if not os.path.exists(checkpoint_path):
    raise FileNotFoundError(f"Saved model not found at: {checkpoint_path}")

checkpoint = torch.load(checkpoint_path, map_location="cpu")
classes = checkpoint["classes"]
IMAGE_SIZE = checkpoint.get("image_size", 64)
NUM_FEATURES = checkpoint.get("num_features", IMAGE_SIZE * IMAGE_SIZE)
hidden_size = checkpoint.get("hidden_size", 256)
num_classes = len(classes)

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

# recreate model and load weights
model = FaceNet(number_of_features=NUM_FEATURES, hidden_size=hidden_size, number_of_classes=num_classes)
model.load_state_dict(checkpoint["state_dict"])
model.to(device)
model.eval()

# transforms — must match training
transform = transforms.Compose([
    transforms.Grayscale(num_output_channels=1),
    transforms.Resize((IMAGE_SIZE, IMAGE_SIZE)),
    transforms.ToTensor(),
    transforms.Normalize(mean=[0.5], std=[0.5])
])

def predict_frame(frame_bgr):
    # convert BGR (cv2) to RGB then to PIL Image
    img = cv2.cvtColor(frame_bgr, cv2.COLOR_BGR2RGB)
    pil = Image.fromarray(img)
    x = transform(pil)  # [1, H, W] with channel=1
    x = x.view(1, -1).to(device)  # flatten
    with torch.no_grad():
        logits = model(x)
        pred = torch.argmax(logits, dim=1).item()
    return classes[pred], pred

def run_webcam(device_index=0):
    cap = cv2.VideoCapture(device_index, cv2.CAP_DSHOW)  # use DirectShow on Windows
    if not cap.isOpened():
        print("Cannot open webcam")
        return

    print("Press 'q' to quit.")
    while True:
        ret, frame = cap.read()
        if not ret:
            break

        try:
            label, idx = predict_frame(frame)
            display_text = f"{label} ({idx})"
        except Exception as e:
            display_text = "Error"

        # overlay text
        cv2.putText(frame, display_text, (10, 30),
                    cv2.FONT_HERSHEY_SIMPLEX, 1.0, (0, 255, 0), 2, cv2.LINE_AA)

        cv2.imshow("Webcam - Press q to quit", frame)
        if cv2.waitKey(1) & 0xFF == ord('q'):
            break

    cap.release()
    cv2.destroyAllWindows()

if __name__ == "__main__":
    run_webcam(0)