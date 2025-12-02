#include<stdbool.h>
#include<stdio.h>


bool isEven(int num){
    if(num == 0) return true;
    if(num == 1) return false;
    return isEven(num - 2);
}

int main(){
    printf("%d", isEven(100000));
}