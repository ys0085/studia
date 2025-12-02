package Zad1;
public class Test{
    public static void main( String args[] ){
        int m = 1;
        try{
            m = Integer.parseInt(args[0]);
            if(m < 1) {
                System.out.println(args[0] + " - nieprawidlowy zakres");
                return;
            }
        }
        catch(NumberFormatException ex){
            System.out.println(args[0] + " - nieprawidlowy zakres");
            return;
        }
        catch(ArrayIndexOutOfBoundsException ex) {
            System.out.println("podaj zakres");
            return;
        }
        LiczbyPierwsze lp = new LiczbyPierwsze(m);
        for(int i = 1; i < args.length; i++){
            try{
                int n = Integer.parseInt(args[i]);
                if(n < 0 || lp.liczba(n) == 0){
                    System.out.println(n + " - wartosc spoza zakresu tablicy");
                }   
                else System.out.println(n + " - " + lp.liczba(n));
            }
            catch(NumberFormatException ex){
                System.out.println(args[i] + " - nieprawidlowa dana");
            }
        }
    }
}
