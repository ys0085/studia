package Zad3;
class Figury{
    interface WartosciJedna{
        public double pole(double r);
        public double obwod(double r);
        public String nazwa();
        
    }
    interface WartosciPiec{
        public double pole(double s1, double s2, double s3, double s4, double angle);
        public double obwod(double s1, double s2, double s3, double s4, double angle);
        public String nazwa();
    }
    public void wypiszJeden(Figury.Jeden e, double val[]){
        System.out.println("Nazwa figury: " + e.nazwa());
        System.out.println("Pole: " + e.pole(val[0]));
        System.out.println("Obwod: " + e.obwod(val[0]));
    }
    public void wypiszPiec(Figury.Piec e, double val[]){
        System.out.println("Nazwa figury: " + e.nazwa());
        System.out.println("Pole: " + e.pole(val[0], val[1], val[2], val[3], val[4]));
        System.out.println("Obwod: " + e.obwod(val[0], val[1], val[2], val[3], val[4]));
    }
    public enum Jeden implements WartosciJedna{
        KOLO{
            public double pole(double r){
                return Math.PI*r*r;
            }
            public double obwod(double r){
                return 2*Math.PI*r;
            }
            public String nazwa(){
                return "Koło";
            }
        },
        KWADRAT{
            public double pole(double r){
                return r*r;
            }
            public double obwod(double r){
                return 4*r;
            }
            public String nazwa(){
                return "Kwadrat";
            }
        },
        PIECIOKAT{
            public double pole(double r){
                return r*r*1.7204774;
            }
            public double obwod(double r){
                return 5*r;
            }
            public String nazwa(){
                return "Pieciokąt";
            }
        },
        SZESCIOKAT{
            public double pole(double r){
                return r*r*2.5980762;
            }
            public double obwod(double r){
                return 6*r;
            }
            public String nazwa(){
                return "Sześciokąt";
            }
        }
    }
    enum Piec implements WartosciPiec{
        PROSTOKAT{
            public double pole(double s1, double s2, double s3, double s4, double angle){
                if(s1 == s2) return s1*s3;
                else return s1*s2;
            }
            public double obwod(double s1, double s2, double s3, double s4, double angle){
                return s1+s2+s3+s4;
            }
            public String nazwa(){
                return "Prostokąt";
            }
        },
        ROMB{
            public double pole(double s1, double s2, double s3, double s4, double angle){
                return s1*s1*Math.sin(angle*Math.PI/180);
            }
            public double obwod(double s1, double s2, double s3, double s4, double angle){
                return s1+s2+s3+s4;
            }
            public String nazwa(){
                return "Romb";
            }
        }
    }
    
}

