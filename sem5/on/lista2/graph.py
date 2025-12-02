

import numpy as np
import matplotlib.pyplot as plt


def f(x):
    return np.exp(x) * np.log(1 + np.exp(-x))


x = np.linspace(-10, 50, 100000)
y = f(x)

# Wykres
plt.figure(figsize=(8,5))
plt.plot(x, y, label=r'$f(x)=e^{x}\ln(1+e^{-x})$', color='blue')
plt.axhline(1, color='red', linestyle='--', label='granica = 1')
plt.title(r'Wykres funkcji $f(x)=e^{x}\ln(1+e^{-x})$')
plt.xlabel('x')
plt.ylabel('f(x)')
plt.grid(True)
plt.legend()
plt.tight_layout()

plt.savefig("wykres_funkcji.png", dpi=150)
plt.show()