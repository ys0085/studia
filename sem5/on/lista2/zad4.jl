# Krzysztof Kleszcz, 279728
using Polynomials
using Printf

coeffs_P = [1, -210.0, 20615.0,-1256850.0, 53327946.0,-1672280820.0, 40171771630.0, -756111184500.0,
11310276995381.0, -135585182899530.0, 1307535010540395.0, -10142299865511450.0, 63030812099294896.0, -311333643161390640.0, 1206647803780373360.0, -3599979517947607200.0, 8037811822645051776.0, -12870931245150988800.0, 13803759753640704000.0, -8752948036761600000.0, 2432902008176640000.0]

coeffs_P = reverse(coeffs_P) 
P = Polynomial(coeffs_P)

println("\n(a) Analiza wielomianu P w postaci naturalnej")
println("-" ^ 80)

println("Obliczanie pierwiastków wielomianu P...")
roots_P = roots(P)

roots_P_sorted = sort(roots_P, by=real)

println("\nObliczone pierwiastki wielomianu P:")
println(rpad("k", 5), rpad("Re(z_k)", 20), rpad("Im(z_k)", 20), 
        rpad("|P(z_k)|", 15), rpad("|p(z_k)|", 15), rpad("|z_k - k|", 15))
println("-" ^ 100)

p(x) = prod(x - k for k in 1:20)


for (idx, z) in enumerate(roots_P_sorted)
    k = idx 

    val_P = abs(P(z))
    val_p = abs(p(z))
    error = abs(z - k)
    
    println(rpad(k, 5), 
            rpad(round(real(z), digits=10), 20),
            rpad(round(imag(z), digits=10), 20),
            rpad(@sprintf("%.2e", val_P), 15),
            rpad(@sprintf("%.2e", val_p), 15),
            rpad(@sprintf("%.2e", error), 15))
end


println("\n" * "=" ^ 80)
println("(b) Eksperyment Wilkinsona - zaburzenie współczynnika")
println("=" ^ 80)

perturbation = -2.0^(-23)
coeffs_P_perturbed = copy(coeffs_P)
coeffs_P_perturbed[20] = -210 + perturbation

println("\nZmieniamy współczynnik przy x^19:")
println("Oryginalny: -210")
println("Zmieniony:  -210 - 2^(-23) = ", -210 + perturbation)
println("Zaburzenie: 2^(-23) ≈ ", abs(perturbation))
println("Względne zaburzenie: ", abs(perturbation/210))


P_perturbed = Polynomial(coeffs_P_perturbed)

println("\nObliczanie pierwiastków zaburzonego wielomianu...")
roots_P_perturbed = roots(P_perturbed)
roots_P_perturbed_sorted = sort(roots_P_perturbed, by=real)

println("\nPorównanie pierwiastków:")
println(rpad("k", 5), rpad("Re(z_k) - oryginalny", 25), 
        rpad("Re(z_k) - zaburzony", 25), rpad("Im(z_k) - zaburzony", 25),
        rpad("Różnica", 15))
println("-" ^ 105)

for (idx, (z_orig, z_pert)) in enumerate(zip(roots_P_sorted, roots_P_perturbed_sorted))
    k = idx
    diff = abs(z_pert - z_orig)
    
    println(rpad(k, 5),
            rpad(@sprintf("%.6f", real(z_orig)), 25),
            rpad(@sprintf("%.6f", real(z_pert)), 25),
            rpad(@sprintf("%.6f", imag(z_pert)), 25),
            rpad(@sprintf("%.2e", diff), 15))
end