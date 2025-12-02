package Zad1;
interface Wartosci{
    public double area(double r);
    public double length(double r);
    public double area(double s1, double s2, double s3, double s4, double angle);
    public double length(double s1, double s2, double s3, double s4, double angle);
}
abstract class Figura implements Wartosci{
    public double area(double s1, double s2, double s3, double s4, double angle){ return -1; }
    public double length(double s1, double s2, double s3, double s4, double angle){ return -1; }
    public String name;
}
class Kolo extends Figura{
    public double area(){return 0;};
    public double length(){return 0;};
    public double area(double radius){
        return radius*radius*Math.PI;
    }
    public double length(double radius){
        return 2*radius*Math.PI;
    }
    Kolo(){
        this.name = "Kolo";
    }
}
class Pieciokat extends Figura{
    public double area(){return 0;};
    public double length(){return 0;};

    public double area(double side){
        return side*side*1.7204774; //trust
    }
    public double length(double side){
        return 5*side;
    }
    Pieciokat(){
        this.name = "Pieciokat foremny";
    
    }
}
class Szesciokat extends Figura{
    public double area(){return 0;};
    public double length(){return 0;};
    public double area(double side){
        return side*side*2.5980762; //trust
    }
    public double length(double side){
        return 6*side;
    }
    Szesciokat(){
        this.name = "Szesciokat foremny";
    }
}
abstract class Czworokat extends Figura{
    public double length(double r){ return -1; }
    public double area(double r){return -1; }
    public double length(double s1, double s2, double s3, double s4, double angle){
        return s1+s2+s3+s4;
    }
}
class Kwadrat extends Czworokat{
    
    public double area(double s1, double s2, double s3, double s4, double angle){
        return s1*s1;
    }
    Kwadrat(){
        this.name = "Kwadrat";
    }
}
class Prostokat extends Czworokat{
    public double area(double s1, double s2, double s3, double s4, double angle){
        if(s1 == s2) return s1*s3;
        else return s1*s2;
    }
    Prostokat(){
        this.name = "Prostokat";
    }
}
class Romb extends Czworokat{
    public double area(double s1, double s2, double s3, double s4, double angle){
        return s1*s1*Math.sin(angle*Math.PI/180);
    }
    Romb(){
        this.name = "Romb";
    }
}
