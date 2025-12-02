# Krzysztof Kleszcz, 279728

function mbisekcji(f, a::Float64, b::Float64, delta::Float64, epsilon::Float64)
    fa = f(a)
    fb = f(b)

    if fa * fb > 0
        return (NaN, NaN, 0, 1)  # brak zmiany znaku
    end

    it = 0
    while (b - a) > delta
        it += 1
        r = (a + b) / 2
        fr = f(r)

        if abs(fr) < epsilon
            return (r, fr, it, 0)
        end

        if fa * fr < 0
            b = r
            fb = fr
        else
            a = r
            fa = fr
        end
    end
    r = (a + b) / 2
    return (r, f(r), it, 0)
end

function mstycznych(f, pf, x0::Float64, delta::Float64, epsilon::Float64, maxit::Int)
    x = x0

    for it in 1:maxit
        fx = f(x)
        if abs(fx) < epsilon
            return (x, fx, it, 0)
        end

        pfx = pf(x)
        if abs(pfx) < eps()
            return (x, fx, it, 2)  # pochodna bliska zero
        end

        x_new = x - fx / pfx

        if abs(x_new - x) < delta
            return (x_new, f(x_new), it, 0)
        end

        x = x_new
    end

    return (x, f(x), maxit, 1)
end


function msiecznych(f, x0::Float64, x1::Float64, delta::Float64, epsilon::Float64, maxit::Int)
    f0 = f(x0)
    f1 = f(x1)

    for it in 1:maxit
        if abs(f1 - f0) < eps()
            return (x1, f1, it, 1)  # brak poprawy, dzielenie przez zero
        end

        x2 = x1 - f1 * (x1 - x0) / (f1 - f0)
        f2 = f(x2)

        if abs(x2 - x1) < delta || abs(f2) < epsilon
            return (x2, f2, it, 0)
        end

        x0, f0 = x1, f1
        x1, f1 = x2, f2
    end

    return (x1, f1, maxit, 1)
end
