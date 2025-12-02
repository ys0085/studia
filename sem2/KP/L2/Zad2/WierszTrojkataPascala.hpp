#pragma once
class WierszTrojkataPascala {
    private: 
        int n;
        int val[1000];
    public:
        void generate();
        int wartosc(int m);
        WierszTrojkataPascala(int n);
        ~WierszTrojkataPascala();
};