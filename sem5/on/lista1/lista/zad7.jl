# Krzysztof Kleszcz, 279728

function derivative(f, h, x)
    return (f(x + h) - f(x)) / h
end
    
function f(x)
    return sin(x) + cos(3.0*x)
end

realvalue = BigFloat(0.11694228168853805109870219901864576419151062786112547521444933168383436487)


for exp in 1:54
    d = derivative(f, 2.0^(-exp), 1.0)
    println("2^-", exp, ": ", d, " ", abs(realvalue - d))
end