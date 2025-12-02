import cv2
import numpy as np
import tensorflow as tf
import matplotlib.pyplot as plt
import sys

def process_image(image_path):

    image = cv2.imread(image_path, cv2.IMREAD_GRAYSCALE)

    # Set image to pure black/white pixels
    _, image = cv2.threshold(image, 185, 255, cv2.THRESH_BINARY) 

    # Resize to 28x28
    image = cv2.resize(image, (28, 28))

    # Invert colors , because MNIST is white on black
    image = 255 - image  

    # Normalize the image (convert pixel values to 0-1 range)
    image = image / 255.0

    # Reshape to fit the model's input shape (batch_size, height, width, channels)
    image = image.reshape(1, 28, 28)  # Model expects (1, 28, 28)

    return image



def main():
    model = tf.keras.models.load_model("mnist_digits.keras")
    image_path = "img/ola_crop/{}_{}.png"
    correct, incorrect = 0, 0
    for digit in range(10):
        for rep in range(4):
            
            image_path_full = image_path.format(digit, rep)
            image = process_image(image_path_full)
            image_original = cv2.imread(image_path_full)
            
            prediction = model.predict(image)
            predicted_digit = np.argmax(prediction)
            perdiction_certainty = np.amax(prediction)

            if predicted_digit == digit:
                correct += 1
                print(f"Correct! Guessed {predicted_digit}.")
            else:
                incorrect += 1
                print(f"Incorrect! Guessed {predicted_digit}, correct answer is {digit}.")

            
            if len(sys.argv) > 1 and sys.argv[1] == "show":
                fig, (ax1, ax2) = plt.subplots(1, 2)

                ax1.imshow(image.squeeze())
                ax2.imshow(image_original.squeeze())

                fig.suptitle(f"Predicted digit: {predicted_digit}, Certainty: {perdiction_certainty}")
                plt.show()
    
    accuracy = correct / (correct + incorrect)
    print("=================")
    print(f"Correct: {correct} out of {correct + incorrect}")
    print(f"Incorrect: {incorrect} out of {correct + incorrect}")
    print(f"Accuracy: {(accuracy * 100):.2f}%")
    print("=================")



if __name__ == "__main__":
    main()