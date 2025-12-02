#include<iostream>
#include<string>
#include<math.h>
//#include"Figura.hpp"
class Figura{
    public:
        std::string name;
        virtual double area(double r)=0;
        virtual double length(double r)=0;
        virtual ~Figura() {};
        
};
class Kolo : public Figura{
    public:
        double area(double s1, double s2, double s3, double s4, double angle){ return -1; }
        double area(double radius){
            return radius*radius*3.1415;
        }
        double length(double radius){
            return 2*radius*3.1415;
        }
    Kolo(){
        this->name = "Kolo";
    }
    ~Kolo(){}
};
class Pieciokat : public Figura{
    public: 
        double area(double s1, double s2, double s3, double s4, double angle){ return -1; }
        double area(double side){
            return side*side*1.7204774; //trust
        }
        double length(double side){
            return 5*side;
        }
    Pieciokat(){
        this->name = "Pieciokat foremny";
    }
    ~Pieciokat(){}
};
class Szesciokat : public Figura{
    public:
        double area(double s1, double s2, double s3, double s4, double angle){ return -1; }
        double area(double side){
            return side*side*2.5980762; //trust
        }
        double length(double side){
            return 6*side;
        }
    Szesciokat(){
        this->name = "Szesciokat foremny";
    }
    ~Szesciokat(){}
};
class Czworokat : public Figura{
    public:
        double area(double r){return -1;}
        double length(double r){return -1;}
        virtual ~Czworokat(){}
        virtual double area(double s1, double s2, double s3, double s4, double angle){ return 0; }
        double length(double s1, double s2, double s3, double s4, double angle){
            return s1+s2+s3+s4;
        }

};
class Kwadrat : public Czworokat{
    public:
        double area(double s1, double s2, double s3, double s4, double angle){
            return s1*s1;
        }
    Kwadrat(){
        this->name = "Kwadrat";
    }
    ~Kwadrat(){}
};
class Prostokat : public Czworokat{
    public:
        double area(double s1, double s2, double s3, double s4, double angle){
            if(s1 == s2) return s1*s3;
            else return s1*s2;
        }
    Prostokat(){
        this->name = "Prostokat";
    }
    ~Prostokat(){}
};
class Romb : public Czworokat{
    public:
        double area(double s1, double s2, double s3, double s4, double angle){
            return s1*s1*sin(angle*3.1415/180);
        }
    Romb(){
        this->name = "Romb";
    }
    ~Romb(){}
};
int main(int argc, char* argv[]){
    
    if(argc == 1){
        std::cout << "podaj figure\n";
        return 0;
    }
    if(argc == 2){
        std::cout << "podaj co najmniej 1 dana\n";
        return 0;
    }
    std::string f = std::string(argv[1]);
    double val[argc-2];
    try {
        for(int i = 0; i < argc-2; i++){
            val[i] = std::stod(std::string(argv[i+2]));
        }
    } catch (std::invalid_argument e) {
        std::cout << "nieprawidlowe dane\n";
        return 0;
    }
    
    if(f.compare("o") == 0) {
        Kolo *fig = new Kolo();
        std::cout << "Nazwa figury: " << fig->name << std::endl;
        std::cout << "Pole: " << fig->area(val[0]) << std::endl;
        std::cout << "Obwod: " << fig->length(val[0]) << std::endl;
        delete fig;
    }
    else if(f.compare("p") == 0){
        Pieciokat *fig = new Pieciokat();
        std::cout << "Nazwa figury: " << fig->name << std::endl;
        std::cout << "Pole: " << fig->area(val[0]) << std::endl;
        std::cout << "Obwod: " << fig->length(val[0]) << std::endl;
        delete fig;
    } 
    else if(f.compare("s") == 0) {
        Szesciokat *fig = new Szesciokat();
        std::cout << "Nazwa figury: " << fig->name << std::endl;
        std::cout << "Pole: " << fig->area(val[0]) << std::endl;
        std::cout << "Obwod: " << fig->length(val[0]) << std::endl;
        delete fig;
    }
    else if(f.compare("c") == 0){
        if(argc < 7){ std::cout << "podaj 5 danych;"; return 0; }
        if(val[4] < 0 || val[4] > 180){ std::cout << "nieprawidlowy kat\n"; return 0; }
        if(val[4] == 90){
            if(val[0] == val[1] && val[1] == val[2] && val[2] == val[3]){ 
                Kwadrat *fig = new Kwadrat(); 
                std::cout << "Nazwa figury: " << fig->name << std::endl;
                std::cout << "Pole: " << fig->area(val[0], val[1], val[2], val[3], val[4]) << std::endl;
                std::cout << "Obwod: " << fig->length(val[0], val[1], val[2], val[3], val[4]) << std::endl;
                delete fig;
            }
            else if((val[0] == val[1] && val[2] == val[3])
            || (val[0] == val[2] && val[1] == val[3])
            || (val[0] == val[3] && val[1] == val[2]))
            { 
                Prostokat *fig = new Prostokat(); 
                std::cout << "Nazwa figury: " << fig->name << std::endl;
                std::cout << "Pole: " << fig->area(val[0], val[1], val[2], val[3], val[4]) << std::endl;
                std::cout << "Obwod: " << fig->length(val[0], val[1], val[2], val[3], val[4]) << std::endl;
                delete fig;
            }
        }
        else if(val[0] == val[1] && val[1] == val[2] && val[2] == val[3]){ 
            Romb *fig = new Romb();
            std::cout << "Nazwa figury: " << fig->name << std::endl;
            std::cout << "Pole: " << fig->area(val[0], val[1], val[2], val[3], val[4]) << std::endl;
            std::cout << "Obwod: " << fig->length(val[0], val[1], val[2], val[3], val[4]) << std::endl;
            delete fig; 
        }
        else { 
            std::cout << "nieobslugiwane dane\n"; 
            return 0; 
        }
    }
    else{
        std::cout << "podaj prawidlowa figure\n";
        return 0;
    }
    return 0;
}
