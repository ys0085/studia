# Krzysztof Kleszcz, 279728

""" 
wykresy:
y = 3x
y = e^x

znalezc miejsce przeciecia, czyli:
e^x = 3x
e^x - 3x = 0
f(x) = e^x - 3x

latwo mozna oszacowac, ze 
f(x<0) > 0
f(0) = 1 > 0
f(1) = e - 3 < 0
f(2) = e^2 - 6 > 0
f(x>2) > 0
"""

include("func/func.jl")

f(x) = exp(x) - 3x

delta = 10^(-4)
epsilon = delta

a = 0.0
b = 1.0
c = 2.0

r, v, it, err = mbisekcji(f, a, b, delta, epsilon)
println("Pierwiastek 1")
println((r=r, v=v, it=it, err=err))

r, v, it, err = mbisekcji(f, b, c, delta, epsilon)
println("Pierwiastek 2")
println((r=r, v=v, it=it, err=err))

