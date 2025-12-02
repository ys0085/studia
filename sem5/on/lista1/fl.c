#include<float.h>
#include<stdio.h>

int main() {
    printf("%g\n%g\n", FLT_EPSILON, DBL_EPSILON);
    printf("%g\n%g\n", FLT_MIN, DBL_MIN);
    printf("%g\n%g\n", FLT_MAX, DBL_MAX);
}