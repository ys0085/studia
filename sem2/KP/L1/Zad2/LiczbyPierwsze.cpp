#include<math.h>
#include"LiczbyPierwsze.hpp"
#include<iostream>
int LiczbyPierwsze::liczba(int m){ return this->tab[m]; }
LiczbyPierwsze::LiczbyPierwsze(int n){
    //printf("gen %d\n", n);
    
    int index = 0, i = 2;
    bool prime;
    while(i < n){
        prime = true;
        for(int j = 2; j < i; j++){
            if(i%j == 0){
                prime = false;
                break;
            }
        }
        if(prime){
            this->tab[index] = i;
            index++;
        }
        i++;
    }
    this->max = index;
}
LiczbyPierwsze::~LiczbyPierwsze(){
    printf("dest\n");
    return;
}
