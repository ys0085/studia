"""
Download MediaPipe model files needed for landmark extraction.
"""
import urllib.request
import os

print("Downloading MediaPipe model files...")
print("=" * 60)

models = [
    {
        "name": "face_landmarker.task",
        "url": "https://storage.googleapis.com/mediapipe-models/face_landmarker/face_landmarker/float16/1/face_landmarker.task",
        "size": "~5 MB"
    },
    {
        "name": "pose_landmarker_lite.task",
        "url": "https://storage.googleapis.com/mediapipe-models/pose_landmarker/pose_landmarker_lite/float16/1/pose_landmarker_lite.task",
        "size": "~5 MB"
    },
    {
        "name": "hand_landmarker.task",
        "url": "https://storage.googleapis.com/mediapipe-models/hand_landmarker/hand_landmarker/float16/1/hand_landmarker.task",
        "size": "~5 MB"
    }
]

for model in models:
    filepath = model["name"]
    
    if os.path.exists(filepath):
        print(f"✓ {model['name']} already exists")
        continue
    
    print(f"\nDownloading {model['name']} ({model['size']})...")
    try:
        urllib.request.urlretrieve(model["url"], filepath)
        print(f"✓ Downloaded {model['name']}")
    except Exception as e:
        print(f"❌ Error downloading {model['name']}: {e}")
        print(f"   URL: {model['url']}")

print("\n" + "=" * 60)
print("✓ All models downloaded!")
print("\nYou can now run:")
print("  python extract_landmarks.py")
