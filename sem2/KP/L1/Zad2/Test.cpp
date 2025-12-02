#include<iostream>
#include"LiczbyPierwsze.hpp"
int StrToInt(char a[1000]);
int main(int argc, char *argv[]){
    //printf("%d\n", argc);
    if(argc == 1) return 0;
    int m;
    try{
        m = StrToInt(argv[1]);
    }
    catch(int ex){
        printf("%s - nieprawidlowy zakres\n", argv[1]);
        return 0;
    }
    { 
        LiczbyPierwsze *lp = new LiczbyPierwsze(m);
        int index = 0;
        for(int i = 2; i < argc; i++){
            try{
                index = StrToInt(argv[i]);
                
                if(index < 0 || index >= lp->max){
                    
                    printf("%s - spoza zakresu\n", argv[i]);
                }
                else printf("%d - %d\n", index, lp->liczba(index));
                
            }
            catch(int ex){
                printf("%s - nieprawidlowa dana\n", argv[i]);
            }
        }
    	delete lp;
    }
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
