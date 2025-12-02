# Krzysztof Kleszcz, 279728

function f(x)
    return sqrt(x^2 + Float64(1.0)) - Float64(1.0)
end

function g(x)
    return x^2 / (sqrt(x^2 + Float64(1.0)) + Float64(1.0))
end

exp = -1
eight = Float64(8)
println("= x =|======== f(x) ========|======== g(x) ========")
while eight^exp > 0.0
    println("8^", Int32(exp), " ", f(eight^exp), " ", g(eight^exp))
    global exp -= 1.0
end