# Krzysztof Kleszcz, 279728
function compute_eps_kahan(::Type{T}) where T
    return T(3) * (T(4)/T(3) - T(1)) - T(1)
end

println(compute_eps_kahan(Float16), " ", eps(Float16))
println(compute_eps_kahan(Float32), " ", eps(Float32))
println(compute_eps_kahan(Float64), " ", eps(Float64))