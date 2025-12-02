#include<math.h>
#include"WierszTrojkataPascala.hpp"
#include<iostream>
#include<algorithm>
void WierszTrojkataPascala::generate(){
    this->val[0] = 1;
    int r = 1;
    int old_val[1000];
    while(r <= this->n){
        std::copy(this->val, this->val + r, old_val);
        val[r] = 1;
        for(int i = 1; i <= r-1; i++){
            this->val[i] = old_val[i-1] + old_val[i];
        }
        r++;
    }
}
int WierszTrojkataPascala::wartosc(int m){
    if(m < 0 || m > this->n) throw -2;
    else return this->val[m];
}
WierszTrojkataPascala::WierszTrojkataPascala(int n){
    if(n < 0){
        throw -3;
        return;
    }
    this->n = n;
    generate();
}
WierszTrojkataPascala::~WierszTrojkataPascala(){
    printf("dest\n");
    return;
}