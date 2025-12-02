#include<iostream>
#include"WierszTrojkataPascala.hpp"
int StrToInt(char a[1000]);
int main(int argc, char *argv[]){
    //printf("%d\n", argc);
    if(argc == 1) return 0;
    WierszTrojkataPascala *wtp;
    int m;
    try{
        m = StrToInt(argv[1]);
        wtp = new WierszTrojkataPascala(m);
    }
    catch(int ex){
        if(ex == -1) std::cout << argv[1] << " - nieprawidlowa dana\n";
        else if (ex == -3) printf("%s - nieprawidlowy zakres\n", argv[1]);
        return 0;
    }
    for(int i = 2; i < argc; i++){
        try{
            printf("%d - %d\n", StrToInt(argv[i]), wtp->wartosc(StrToInt(argv[i])));
        }
        catch(int ex){
            if(ex == -1) printf("%s - nieprawidlowa dana\n", argv[i]);
            else if(ex == -2) printf("%s - wartosc spoza zakresu tablicy\n", argv[i]);
        }
    }

    delete wtp;
    return 0;
}
int StrToInt(char a[1000]){
    int i = 0;
    int num = 0;
    bool neg = false;
    if(a[0] == '-'){ neg = true; i++; }
    while(a[i] != 0){
        if((a[i] < '0' || a[i] > '9')){
            throw -1;
            return 0;
        }
        num = (a[i] - '0')  + (num * 10);
        i++;
    }
    return neg ? -num : num;
}
