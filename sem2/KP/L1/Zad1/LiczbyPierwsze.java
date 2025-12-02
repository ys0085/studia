package Zad1;

public class LiczbyPierwsze{
    private int[] tab;
    public int liczba(int m){ return tab[m]; }
    LiczbyPierwsze(int n){
        tab = new int[n];
        int index = 0, i = 2;
        boolean p;
        while(i < n){
            p = true;
            for(int j = 2; j <= Math.sqrt(i); j++){
                if(i%j == 0){
                    p = false;
                    break;
                }
            }
            if(p){
                tab[index] = i;
                index++;
            }
            i++;
        }
    }
}
