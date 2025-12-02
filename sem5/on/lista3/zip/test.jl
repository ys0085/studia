# Krzysztof Kleszcz, 279728

include("func/func.jl")


f(x) = x^3 - 2x - 5
pf(x) = 3x^2 - 2


println("--- Test metody bisekcji ---")
r, v, it, err = mbisekcji(f, 1.0, 3.0, 1e-6, 1e-6)
println((r=r, v=v, it=it, err=err))


println("--- Test metody Newtona ---")
r, v, it, err = mstycznych(f, pf, 2.0, 1e-6, 1e-6, 50)
println((r=r, v=v, it=it, err=err))


println("--- Test metody siecznych ---")
r, v, it, err = msiecznych(f, 1.0, 3.0, 1e-6, 1e-6, 50)
println((r=r, v=v, it=it, err=err))




println("--- Test błędów: bisekcja (brak zmiany znaku) ---")
r, v, it, err = mbisekcji(f, -1.0, 0.0, 1e-6, 1e-6)
println((r=r, v=v, it=it, err=err))


println("--- Test błędów: Newton (pochodna bliska zeru) ---")
f2(x) = x^3 + 1.0
pf2(x) = 3x^2
r, v, it, err = mstycznych(f2, pf2, 0.0, 1e-6, 1e-6, 20)
println((r=r, v=v, it=it, err=err))


println("--- Test błędów: metoda siecznych (dzielenie przez zero) ---")
f3(x) = 1.0 
r, v, it, err = msiecznych(f3, 1.0, 2.0, 1e-6, 1e-6, 20)
println((r=r, v=v, it=it, err=err))