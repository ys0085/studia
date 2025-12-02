import numpy as np
from tensorflow.keras.datasets import mnist

def load_emnist():
    (x_train, y_train), (_, _) = mnist.load_data()

    x_train = x_train.astype(np.float32) / 255.0

    x_train = x_train.reshape(x_train.shape[0], -1)

    return x_train, y_train