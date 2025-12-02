# Krzysztof Kleszcz, 279728

using Printf

function iterate_map(x0, c, N=40)
    x = x0
    seq = Vector{Float64}(undef, N+1)
    seq[1] = x
    for i in 1:N
        x = x^2 + c
        seq[i+1] = x
    end
    return seq
end

cases = [
    (-2.0, 1.0),
    (-2.0, 2.0),
    (-2.0, 1.99999999999999),
    (-1.0, 1.0),
    (-1.0, -1.0),
    (-1.0, 0.75),
    (-1.0, 0.25)
]

for (c, x0) in cases
    seq = iterate_map(x0, c)
    @printf("\nCase c=%g, x0=%1.16g\n", c, x0)
    for i in 1:41
        @printf("%3d: %16.25g\n", i-1, seq[i])
    end
end