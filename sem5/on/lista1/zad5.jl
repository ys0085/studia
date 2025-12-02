# Krzysztof Kleszcz, 279728

using LinearAlgebra

function dot_1(x, y)
    T = typeof(x[1])
    S = T(0)
    for i in 1:length(x)
        S += x[i]*y[i]
    end
    return S
end

function dot_2(x, y)
    T = typeof(x[1])
    S = T(0)
    for i in 1:length(x)
        S += x[(length(x)+1) - i]*y[(length(x)+1) - i]
    end
    return S
end

function dot_3(x, y)
    T = typeof(x[1])
    pos = zeros(T, 0)
    neg = zeros(T, 0)
    for i in 1:length(x)
        m = x[i]*y[i]
        if m >= 0
            append!(pos, m)
        else 
            append!(neg, m)
        end
    end
    sort!(pos, rev=true)
    sort!(neg)
    pos_sum = T(0)
    neg_sum = T(0)
    S = T(0)
    for n in pos
        pos_sum += n
    end
    for n in neg
        neg_sum += n
    end
    S = pos_sum + neg_sum
end

function dot_4(x, y)
    T = typeof(x[1])
    pos = zeros(T, 0)
    neg = zeros(T, 0)
    for i in 1:length(x)
        m = x[i]*y[i]
        if m >= 0
            append!(pos, m)
        else 
            append!(neg, m)
        end
    end
    sort!(pos)
    sort!(neg, rev=true)
    pos_sum = T(0)
    neg_sum = T(0)
    S = T(0)
    for n in pos
        pos_sum += n
    end
    for n in neg
        neg_sum += n
    end
    S = pos_sum + neg_sum
    return S
end


x = [2.718281828, -3.141592654, 1.414213562, 0.5772156649, 0.3010299957]
y = [1486.2497, 878366.9879, -22.37492, 4773714.647, 0.000185049]
correct = BigFloat(-1.00657107000000 * 10^(-11))

println("correct: ", correct)

xv = Float32.(x)
yv = Float32.(y)

println("====== Float32 ======")

d = dot_1(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_2(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_3(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_4(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct))) 


xv = Float64.(x)
yv = Float64.(y)


println("====== Float64 ======")

d = dot_1(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_2(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_3(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_4(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct))) 


xv = BigFloat.(x)
yv = BigFloat.(y)


println("====== BigFloat ======")

d = dot_1(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_2(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_3(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct)))

d = dot_4(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct))) 

d = dot(xv, yv)
println(d, " delta: ", round(abs(d-correct)/abs(correct))) 