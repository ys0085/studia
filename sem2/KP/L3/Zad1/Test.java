package Zad1;

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
        Figura fig = new Kwadrat(); 
        if(f.compareTo("o") == 0) fig = new Kolo();
        else if(f.compareTo("p") == 0) fig = new Pieciokat();
        else if(f.compareTo("s") == 0) fig = new Szesciokat();
        else if(f.compareTo("c") == 0){
            if(val.length < 5){ System.out.println("podaj 5 danych"); return; }
            if(val[4] < 0 || val[4] > 180){ System.out.println("nieprawidlowy kat"); return;}
            if(val[4] == 90){
                if(val[0] == val[1] && val[1] == val[2] && val[2] == val[3]){ fig = new Kwadrat(); }
                else if((val[0] == val[1] && val[2] == val[3])
                || (val[0] == val[2] && val[1] == val[3])
                || (val[0] == val[3] && val[1] == val[2])){ fig = new Prostokat(); }
            }
            else if(val[0] == val[1] && val[1] == val[2] && val[2] == val[3]){ fig = new Romb(); }
            else { System.out.println("nieobslugiwane dane"); return; }
        }
        else{
            System.out.println("podaj prawidlowa figure");
            return;
        }
    
        if(f.compareTo("c") == 0){
            System.out.println("Nazwa figury: " + fig.name);
            System.out.println("Pole: " + fig.area(val[0], val[1], val[2], val[3], val[4]));
            System.out.println("Obwod: " + fig.length(val[0], val[1], val[2], val[3], val[4]));
        }
        else{
            System.out.println("Nazwa figury: " + fig.name);
            System.out.println("Pole: " + fig.area(val[0]));
            System.out.println("Obwod: " + fig.length(val[0]));
        }
    }
}
