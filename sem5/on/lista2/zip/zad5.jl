# Krzysztof Kleszcz, 279728

using Printf

function logistic(p0, r, n, T)
    p = convert(T, p0)
    seq = Vector{T}(undef, n+1)
    seq[1] = p
    for i in 1:n
        p = p + r * p * (1 - p)
        seq[i+1] = p
    end
    return seq
end

seq1 = logistic(0.01, 3, 40, Float32)
seq2 = logistic(0.01, 3, 10, Float32)
cut = round(seq2[end], digits=3)
seq2b = logistic(cut, 3, 30, Float32)
seq2 = vcat(seq2, seq2b[2:end])

println("iter  Float32  po obcięciu")
for i in 1:41
    @printf("%3d  %10.6f  %10.6f\n", i-1, seq1[i], seq2[i])
end

seq32 = logistic(0.01, 3, 40, Float32)
seq64 = logistic(0.01, 3, 40, Float64)

println("\nPorównanie Float32 i Float64:")
for i in 1:41
    @printf("iter %2d: Float32=%12.7g Float64=%12.16g diff=%0.3e\n",
            i-1, seq32[i], seq64[i], abs(Float64(seq32[i]) - seq64[i]))
end