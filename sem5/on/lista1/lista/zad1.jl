# Krzysztof Kleszcz, 279728

function compute_eps(::Type{T}) where T
    one = T(1)
    epsilon = T(1)
    while one + epsilon/2 > one
        epsilon = epsilon / 2
    end
    return epsilon
end

function compute_min(::Type{T}) where T
    num = T(1)
    while num/2 > 0
        num /= 2
    end
    return num
end

function compute_max(::Type{T}) where T
    num = T(1)
    while !isinf(num * T(2))
        num *= T(2)
    end

    step = num
    while true
        if !isinf(num + step) && num + step != num
            num += step
        else
            step /= 2
            if step == 0
                break
            end
        end
    end
    return num
end



c_eps16 = compute_eps(Float16)
b_eps16 = eps(Float16)
c_eps32 = compute_eps(Float32)
b_eps32 = eps(Float32)
c_eps64 = compute_eps(Float64)
b_eps64 = eps(Float64)

println(repeat("=", 80))
println("      My Float16_eps: ", c_eps16, " ", bitstring(c_eps16))
println("Built-in Float16_eps: ", b_eps16, " ", bitstring(b_eps16))
println("      My Float32_eps: ", c_eps32, " ", bitstring(c_eps32))
println("Built-in Float32_eps: ", b_eps32, " ", bitstring(b_eps32))
println("      My Float64_eps: ", c_eps64, " ", bitstring(c_eps64))
println("Built-in Float64_eps: ", b_eps64, " ", bitstring(b_eps64))

c_min16 = compute_min(Float16)
b_min16 = nextfloat(Float16(0.0))
c_min32 = compute_min(Float32)
b_min32 = nextfloat(Float32(0.0))
c_min64 = compute_min(Float64)
b_min64 = nextfloat(Float64(0.0))

println(repeat("=", 80))
println("      My Float16_min: ", c_min16, " ", bitstring(c_min16))
println("Built-in Float16_min: ", b_min16, " ", bitstring(b_min16))
println("      My Float32_min: ", c_min32, " ", bitstring(c_min32))
println("Built-in Float32_min: ", b_min32, " ", bitstring(b_min32))
println("      My Float64_min: ", c_min64, " ", bitstring(c_min64))
println("Built-in Float64_min: ", b_min64, " ", bitstring(b_min64))

c_max16 = compute_max(Float16)
b_max16 = floatmax(Float16)
c_max32 = compute_max(Float32)
b_max32 = floatmax(Float32)
c_max64 = compute_max(Float64)
b_max64 = floatmax(Float64)

println(repeat("=", 80))
println("      My Float16_max: ", c_max16, " ", bitstring(c_max16))
println("Built-in Float16_max: ", b_max16, " ", bitstring(b_max16))
println("      My Float32_max: ", c_max32, " ", bitstring(c_max32))
println("Built-in Float32_max: ", b_max32, " ", bitstring(b_max32))
println("      My Float64_max: ", c_max64, " ", bitstring(c_max64))
println("Built-in Float64_max: ", b_max64, " ", bitstring(b_max64))