import matplotlib.pyplot as plt
import tensorflow as tf
import numpy as np
from sklearn.metrics import precision_score, recall_score, accuracy_score, classification_report

def main():
    mnist = tf.keras.datasets.mnist
    (x_train, y_train), (x_test, y_test) = mnist.load_data()

    
    x_train, x_test = x_train / 255.0, x_test / 255.0   # Normalize pixel values (0-255 → 0-1)


    model = tf.keras.models.Sequential([
        tf.keras.layers.Flatten(input_shape=(28, 28)),  # Flatten 28x28 images
        tf.keras.layers.Dense(128, activation='relu'),  # Hidden layer with 128 neurons
        tf.keras.layers.Dropout(0.2),                   # Dropout for regularization
        tf.keras.layers.Dense(10, activation='softmax') # Output layer (10 classes)
    ])

    model.compile(optimizer='adam', loss='sparse_categorical_crossentropy', metrics=['accuracy'])

    model.fit(x_train, y_train, epochs=5)
    

    y_pred = model.predict(x_test)
    pred_classes = np.argmax(y_pred, axis=1) 

    acc = accuracy_score(y_test, pred_classes)
    precision = precision_score(y_test, pred_classes, average="macro")
    recall = recall_score(y_test, pred_classes, average="macro")

    print(f'\nTest accuracy: {acc:.4f}')
    print(f'\nTest recall: {recall:.4f}')
    print(f'\nTest precision: {precision:.4f}')

    plt.imshow(x_test[0], cmap="gray")
    plt.title(f'Predicted: {pred_classes[0]}')
    plt.show()

    model.save("mnist_digits.keras")

    print(classification_report(y_test, pred_classes))




if __name__ == "__main__":
    main()

