# Krzysztof Kleszcz, 279728

using LinearAlgebra
using Random
using Printf

function hilb(n::Int)
# Function generates the Hilbert matrix  A of size n,
#  A (i, j) = 1 / (i + j - 1)
# Inputs:
#	n: size of matrix A, n>=1
#
#
# Usage: hilb(10)
#
# Pawel Zielinski
        if n < 1
         error("size n should be >= 1")
        end
        return [1 / (i + j - 1) for i in 1:n, j in 1:n]
end

function matcond(n::Int, c::Float64)
# Function generates a random square matrix A of size n with
# a given condition number c.
# Inputs:
#	n: size of matrix A, n>1
#	c: condition of matrix A, c>= 1.0
#
# Usage: matcond(10, 100.0)
#
# Pawel Zielinski
        if n < 2
         error("size n should be > 1")
        end
        if c< 1.0
         error("condition number  c of a matrix  should be >= 1.0")
        end
        (U,S,V)=svd(rand(n,n))
        return U*diagm(0 =>[LinRange(1.0,c,n);])*V'
end

function test_hilbert()
    println("Macierze Hilberta:")
    for n in 2:16
        A = hilb(n)
        x_true = ones(n)
        b = A * x_true
        x1 = A \ b
        x2 = inv(A) * b
        err1 = norm(x1 - x_true) / norm(x_true)
        err2 = norm(x2 - x_true) / norm(x_true)
        @printf("n=%2d cond(A)=%.2e  err1=%.2e  err2=%.2e\n", n, cond(A), err1, err2)
    end
end

function test_matcond()
    println("\nMacierze o zadanym condition number:")
    for n in [5, 10, 20], c in [1, 10, 1e3, 1e7, 1e12, 1e16]
        A = matcond(n, c)
        x_true = ones(n)
        b = A * x_true
        x1 = A \ b
        x2 = inv(A) * b
        err1 = norm(x1 - x_true) / norm(x_true)
        err2 = norm(x2 - x_true) / norm(x_true)
        @printf("n=%2d c=%.0e cond=%.2e err1=%.2e err2=%.2e\n", n, c, cond(A), err1, err2)
    end
end

test_hilbert()
test_matcond()