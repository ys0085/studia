# Krzysztof Kleszcz, 279728

x = Float64(1)
smallest = Float64(2)
loops = 0
ex = 0

while true
    if loops > 10000000
        break
    end
    global ex
    global x = rand(Float64) / 2^ex + Float64(1)
    if x * (1/x) != Float64(1) && x < smallest
        global smallest = x
        global loops = 0
        ex += 1
    end
    loops+=1 
end


println(smallest)


x = Float64(1)
loops = 0
while true
    if loops > 10000000 || x * 1/x != 1
        break
    end
    global x = nextfloat(x)
    global loops += 1
end

println(x)