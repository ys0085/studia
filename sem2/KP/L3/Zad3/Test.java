package Zad3;

public class Test {
    public static void main(String args[]){
        if(args.length == 0){
            System.out.println("podaj figure");
            return;
        }
        if(args.length == 1){
            System.out.println("podaj co najmniej 1 dana");
            return;
        }
        String f = args[0];
        double val[] = new double[args.length-1];
        try {
            for(int i = 0; i < args.length-1; i++){
                val[i] = Double.parseDouble(args[i+1]);
            }
        } catch (Exception e) {
            System.out.println("nieprawidlowe dane");
            return;
        }
        Figury fig = new Figury();
        if(f.compareTo("o") == 0) fig.wypiszJeden(Figury.Jeden.KOLO, val);
        else if(f.compareTo("p") == 0) fig.wypiszJeden(Figury.Jeden.PIECIOKAT, val);
        else if(f.compareTo("s") == 0) fig.wypiszJeden(Figury.Jeden.SZESCIOKAT, val);
        else if(f.compareTo("c") == 0){
            if(val.length == 1) {fig.wypiszJeden(Figury.Jeden.KWADRAT, val); return; }
            if(val.length < 5){ System.out.println("podaj 5 danych"); return; }
            if(val[4] < 0 || val[4] > 180){ System.out.println("nieprawidlowy kat"); return;}
            if(val[4] == 90){
                if(val[0] == val[1] && val[1] == val[2] && val[2] == val[3]) fig.wypiszJeden(Figury.Jeden.KWADRAT, val); 
                else if((val[0] == val[1] && val[2] == val[3])
                || (val[0] == val[2] && val[1] == val[3])
                || (val[0] == val[3] && val[1] == val[2])) fig.wypiszPiec(Figury.Piec.PROSTOKAT, val); 
            }
            else if(val[0] == val[1] && val[1] == val[2] && val[2] == val[3]) fig.wypiszPiec(Figury.Piec.ROMB, val);
            else { System.out.println("nieobslugiwane dane"); return; }
        }
        else{
            System.out.println("podaj prawidlowa figure");
            return;
        }

        
    }
}
