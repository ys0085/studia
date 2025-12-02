package Zad1;

import java.util.Objects;

public class Test{
    public static void main(String args[]){
        try{
            if (Objects.equals(args[0], new String("test"))) {
                int overflow = 1;
                int prev = 0;
                int row = 0;
                WierszTrojkataPascala wtpTest;
                while(overflow > 0){
                    row++;
                    //System.out.println(row);
                    prev = overflow;
                    wtpTest = new WierszTrojkataPascala(row);
                    overflow = wtpTest.wartosc((int) row/2);
                }
                System.out.println(row-1 + " " + prev);



                return;
            }
        }
        catch(InvalidRowNumberException e){  }
        catch(IndexOutOfRangeException e){  }
        catch(ArrayIndexOutOfBoundsException e){ System.out.println("podaj numer wiersza lub wpisz \"test\""); return; }

        WierszTrojkataPascala wtp;
        try{
            wtp = new WierszTrojkataPascala(Integer.parseInt(args[0]));
        }
        catch(NumberFormatException e){
            System.out.println(args[0] + " - nieprawidlowa dana");
            return;
        }
        catch(InvalidRowNumberException e){
            System.out.println(args[0] + " - nieprawidlowy numer wiersza");
            return;
        }
        catch(ArrayIndexOutOfBoundsException ex) {
            System.out.println("podaj numer wiersza lub wpisz \"test\"");
            return;
        }
        for(int i = 1; i < args.length; i++){
            try{
                System.out.println(args[i] + " - " + wtp.wartosc(Integer.parseInt(args[i])));
            }
            catch(NumberFormatException e){
                System.out.println(args[i] + " - nieprawidlowa dana");
            }
            catch(IndexOutOfRangeException e){
                System.out.println(args[i] + " - liczba spoza zakresu");
            }

        }
        

    }
}