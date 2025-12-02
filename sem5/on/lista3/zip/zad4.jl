# Krzysztof Kleszcz, 279728

include("func/func.jl")

f(x) = sin(x) - (x/2)^2
pf(x) = cos(x) - x/2

delta = 0.5 * 10^(-5)
epsilon = delta
max_iter = 100

a = 1.5
b = 2.0

r, v, it, err = mbisekcji(f, a, b, delta, epsilon)
println("Metoda bisekcji")
println((r=r, v=v, it=it, err=err))

x0 = 1.5

r, v, it, err = mstycznych(f, pf, x0, delta, epsilon, max_iter)
println("Metoda Newtona")
println((r=r, v=v, it=it, err=err))

x1 = 1.0
x2 = 2.0

r, v, it, err = msiecznych(f, x1, x2, delta, epsilon, max_iter)
println("Metoda siecznych")
println((r=r, v=v, it=it, err=err))

