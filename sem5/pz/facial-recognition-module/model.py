import os
import torch
import torchvision.transforms as transforms
from torchvision.datasets import ImageFolder
from torch.utils.data import DataLoader, random_split
from torch import nn
from torch.optim import Adam


# =======================================================
# 1. FIXED VERSION OF YOUR MODEL
# =======================================================

class FaceNet(nn.Module):
    def __init__(self, number_of_features, hidden_size, number_of_classes):
        super().__init__()

        # Input shape for each image after processing: [features, 2]
        # But we are flattening images → 1D vector, so adjust:
        self.classifier = nn.Sequential(
            nn.Flatten(),
            nn.Linear(number_of_features, hidden_size),
            nn.LeakyReLU(),
            nn.Linear(hidden_size, hidden_size),
            nn.LeakyReLU(),
            nn.Linear(hidden_size, number_of_classes)
            # Softmax removed — CrossEntropyLoss applies it internally
        )

    def forward(self, x):
        return self.classifier(x)

if __name__ == "__main__":
    # =======================================================
    # 2. DATA LOADING + NORMALIZATION (mean=0, variance=1)
    # =======================================================

    path = "C:/Users/Lenovo/.cache/kagglehub/datasets/arafatsahinafridi/multi-class-driver-behavior-image-dataset/versions/1/Multi-Class Driver Behavior Image Dataset"

    # We will:
    # - Convert images to grayscale
    # - Resize to fixed size (e.g., 64×64)
    # - Convert to tensor
    # - Normalize to mean=0, variance=1
    #
    # After flattening, each image will be 64×64 = 4096 features

    IMAGE_SIZE = 64
    NUM_FEATURES = IMAGE_SIZE * IMAGE_SIZE  # 4096

    transform = transforms.Compose([
        transforms.Grayscale(num_output_channels=1),
        transforms.Resize((IMAGE_SIZE, IMAGE_SIZE)),
        transforms.ToTensor(),
        transforms.Normalize(mean=[0.5], std=[0.5])  # approx mean=0, var=1
    ])

    dataset = ImageFolder(root=path, transform=transform)

    num_classes = len(dataset.classes)
    print("Discovered classes:", dataset.classes)


    # =======================================================
    # 3. TRAIN / TEST SPLIT
    # =======================================================

    train_ratio = 0.8
    train_size = int(train_ratio * len(dataset))
    test_size = len(dataset) - train_size

    train_dataset, test_dataset = random_split(dataset, [train_size, test_size])

    train_loader = DataLoader(train_dataset, batch_size=32, shuffle=True)
    test_loader = DataLoader(test_dataset, batch_size=32, shuffle=False)


    # =======================================================
    # 4. CREATE MODEL
    # =======================================================

    model = FaceNet(
        number_of_features=NUM_FEATURES,
        hidden_size=256,
        number_of_classes=num_classes
    )

    criterion = nn.CrossEntropyLoss()
    optimizer = Adam(model.parameters(), lr=0.001)


    # =======================================================
    # 5. TRAINING LOOP
    # =======================================================

    EPOCHS = 10

    print("Starting training...\n")

    for epoch in range(EPOCHS):
        model.train()
        total_loss = 0

        for images, labels in train_loader:
            # Flatten: [batch, 1, 64, 64] → [batch, 4096]
            images = images.view(images.size(0), -1)

            optimizer.zero_grad()
            logits = model(images)
            loss = criterion(logits, labels)
            loss.backward()
            optimizer.step()

            total_loss += loss.item()

        avg_loss = total_loss / len(train_loader)
        print(f"Epoch {epoch + 1}/{EPOCHS} | Loss: {avg_loss:.4f}")


    # =======================================================
    # 6. TEST EVALUATION
    # =======================================================

    model.eval()
    correct = 0
    total = 0

    with torch.no_grad():
        for images, labels in test_loader:
            images = images.view(images.size(0), -1)
            logits = model(images)
            preds = torch.argmax(logits, dim=1)
            correct += (preds == labels).sum().item()
            total += labels.size(0)

    accuracy = correct / total
    print(f"\nTest Accuracy: {accuracy * 100:.2f}%")

    # =======================================================
    # 7. SAVE MODEL + METADATA
    # =======================================================

    save_dir = os.path.join(os.path.dirname(__file__), "saved_models")
    os.makedirs(save_dir, exist_ok=True)
    save_path = os.path.join(save_dir, "facenet.pth")

    torch.save({
        "state_dict": model.state_dict(),
        "classes": dataset.classes,
        "image_size": IMAGE_SIZE,
        "num_features": NUM_FEATURES
    }, save_path)

    print(f"Model and metadata saved to: {save_path}")
