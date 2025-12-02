# Krzysztof Kleszcz, 279728

x = Float64(1)
next_x = nextfloat(Float64(x))
delta_x = next_x - x

random = rand(Float64) + Float64(1)
next_r = nextfloat(Float64(random))
delta_r = next_r - random

d = Float64(2^-52)

println("Delta k = nextfloat(k) - k")
println("r ∈ [x, 2*x)")


println("=================== x = 1.0 ===================")
println("Delta x: ", bitstring(delta_x), " ", delta_x)
println("Delta r: ", bitstring(delta_r), " ", delta_r)
println("  2^-52: ", bitstring(d), " ", d)



x = Float64(2)
next_x = nextfloat(Float64(x))
delta_x = next_x - x

random = rand(Float64)*2 + Float64(2)
next_r = nextfloat(Float64(random))
delta_r = next_r - random

d = Float64(2^-51)


println("=================== x = 2.0 ===================")
println("Delta x: ", bitstring(delta_x), " ", delta_x)
println("Delta r: ", bitstring(delta_r), " ", delta_r)
println("  2^-51: ", bitstring(d), " ", d)


x = Float64(0.5)
next_x = nextfloat(Float64(x))
delta_x = next_x - x

random = rand(Float64)*0.5 + Float64(0.5)
next_r = nextfloat(Float64(random))
delta_r = next_r - random

d = Float64(2^-53)


println("=================== x = 0.5 ===================")
println("Delta x: ", bitstring(delta_x), " ", delta_x)
println("Delta r: ", bitstring(delta_r), " ", delta_r)
println("  2^-53: ", bitstring(d), " ", d)

