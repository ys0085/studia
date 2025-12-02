# Krzysztof Kleszcz, 279728
include("func/func.jl")

f1(x) = exp(1-x)-1.0
f2(x) = x*exp(-x)

pf1(x) = -exp(1-x)
pf2(x) = -exp(-x)*(x-1)

delta = 10^(-5)
epsilon = delta

max_iter = 100


println("f1(x) = e^(1-x) - 1")

r, v, it, err = mbisekcji(f1, -5.0, 5.0, delta, epsilon)
println("Metoda bisekcji")
println((r=r, v=v, it=it, err=err))

r, v, it, err = mstycznych(f1, pf1, 0.0, delta, epsilon, max_iter)
println("Metoda Newtona")
println((r=r, v=v, it=it, err=err))

r, v, it, err = msiecznych(f1, -2.0, 2.0, delta, epsilon, max_iter)
println("Metoda siecznych")
println((r=r, v=v, it=it, err=err))


println("f2(x) = x*e^(-x)")

r, v, it, err = mbisekcji(f2, -11.0, 10.0, delta, epsilon)
println("Metoda bisekcji")
println((r=r, v=v, it=it, err=err))

r, v, it, err = mstycznych(f2, pf2, 2.0, delta, epsilon, max_iter)
println("Metoda Newtona")
println((r=r, v=v, it=it, err=err))

r, v, it, err = msiecznych(f2, -2.0, 2.0, delta, epsilon, max_iter)
println("Metoda siecznych")
println((r=r, v=v, it=it, err=err))




println("testy:")
println("Metoda Newtona")
println("x0 ∈ (1, ∞], f1")

println("x0 = 5")
r, v, it, err = mstycznych(f1, pf1, 5.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 20")
r, v, it, err = mstycznych(f1, pf1, 20.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 100")
r, v, it, err = mstycznych(f1, pf1, 100.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 1000")
r, v, it, err = mstycznych(f1, pf1, 1000.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 10000")
r, v, it, err = mstycznych(f1, pf1, 10000.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 1000000")
r, v, it, err = mstycznych(f1, pf1, 1000000.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 ∈ [1, ∞], f2")

println("x0 = 1")
r, v, it, err = mstycznych(f2, pf2, 1.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 5")
r, v, it, err = mstycznych(f2, pf2, 5.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 10")
r, v, it, err = mstycznych(f2, pf2, 10.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 100")
r, v, it, err = mstycznych(f2, pf2, 100.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))

println("x0 = 1000")
r, v, it, err = mstycznych(f2, pf2, 1000.0, delta, epsilon, max_iter)
println((r=r, v=v, it=it, err=err))




