# Driver Behavior Detection System

Real-time driver behavior classification using MediaPipe landmarks and neural networks.

## Overview

This system detects and classifies 5 different driver behaviors:
- **safe_driving** - Normal driving posture
- **texting_phone** - Using phone for texting
- **talking_phone** - Talking on the phone
- **turning** - Turning the steering wheel
- **other_activities** - Other distracting activities

**Current Model Accuracy:** 70.40%

## How It Works

1. **MediaPipe** extracts landmarks from webcam feed:
   - Face landmarks (478 points)
   - Pose landmarks (33 points) 
   - Hand landmarks (42 points)

2. **Neural Network** processes these landmarks to classify behavior

3. **Real-time display** shows:
   - Current behavior prediction
   - Confidence percentages for all classes

## Files

- `train_simple.py` - Train the model on your dataset
- `detect_simple.py` - Run real-time detection on webcam
- `download_models.py` - Download MediaPipe model files
- `saved_models/driver_behavior_classifier.pth` - Trained model
- `face_landmarker.task`, `pose_landmarker_lite.task`, `hand_landmarker.task` - MediaPipe models

## Setup

### 1. Install Dependencies

```bash
pip install mediapipe torch opencv-python numpy scikit-learn tqdm
```

### 2. Download MediaPipe Models

```bash
python download_models.py
```

This downloads the required MediaPipe model files (~15MB total).

### 3. Prepare Your Dataset (Optional - for retraining)

Organize images in this structure:
```
dataset/
├── safe_driving/
│   ├── img001.jpg
│   └── ...
├── texting_phone/
│   └── ...
├── talking_phone/
│   └── ...
├── turning/
│   └── ...
└── other_activities/
    └── ...
```

### 4. Train the Model (Optional)

```bash
python train_simple.py
```

This will:
- Process all images in the dataset
- Extract MediaPipe landmarks
- Train a neural network for 30 epochs
- Save the best model to `saved_models/driver_behavior_classifier.pth`

**Training time:** 1-2 hours for ~7,000 images

### 5. Run Real-Time Detection

```bash
python detect_simple.py
```

Press **'q'** to quit.

## Technical Details

### MediaPipe Integration

Following Google's official MediaPipe samples pattern:
- Uses **VIDEO mode** for optimal webcam performance
- Implements `detect_for_video()` with timestamps
- Enables tracking for reduced latency
- Processes at ~30 FPS

### Neural Network Architecture

```
Input: 1659 features (478×3 + 33×3 + 21×3×2)
    ↓
Linear(1659 → 256) + ReLU + Dropout(0.3)
    ↓
Linear(256 → 128) + ReLU + Dropout(0.3)
    ↓
Linear(128 → 5 classes)
```

### Training Configuration

- **Optimizer:** Adam (lr=0.001)
- **Loss:** CrossEntropyLoss
- **Epochs:** 30
- **Batch Size:** 32
- **Train/Test Split:** 80/20

## Improving Accuracy

To get better results:

1. **Add more training data** (aim for 500+ images per class)
2. **Increase epochs** (edit `EPOCHS = 50` in `train_simple.py`)
3. **Balance your dataset** (equal images per class)
4. **Use better quality images** (good lighting, clear poses)
5. **Add data augmentation** (flips, rotations, brightness changes)

## Troubleshooting

**"Model not found" error:**
- Run `python train_simple.py` first to train the model

**"Cannot open webcam":**
- Check camera permissions
- Try different camera index: Change `cap = cv2.VideoCapture(0)` to `cap = cv2.VideoCapture(1)`

**Low accuracy:**
- Need more diverse training data
- Increase training epochs
- Check if dataset is balanced

**Slow performance:**
- MediaPipe VIDEO mode should run at 30 FPS
- Check CPU usage
- Consider using GPU if available

## Dataset

Current model trained on 7,276 images:
- safe_driving: 1,679 images
- talking_phone: 1,513 images
- texting_phone: 1,561 images
- turning: 1,339 images
- other_activities: 1,184 images

## References

- [MediaPipe Documentation](https://ai.google.dev/edge/mediapipe/solutions/guide)
- [MediaPipe Pose Landmarker](https://ai.google.dev/edge/mediapipe/solutions/vision/pose_landmarker)
- [MediaPipe Face Landmarker](https://ai.google.dev/edge/mediapipe/solutions/vision/face_landmarker)
- [MediaPipe Hand Landmarker](https://ai.google.dev/edge/mediapipe/solutions/vision/hand_landmarker)

## License

Apache 2.0
