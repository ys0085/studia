#include<iostream>
#include"BinaryTree.cpp"

int main(){
    BinaryTree *t = new BinaryTree();
    
    BinaryTree *sTree = new BinaryTree();
    BinaryTree *dTree = new BinaryTree();
    BinaryTree *iTree = new BinaryTree();
    
    
    bool q = false;
    t = sTree;
    int mode = 0;
    while(!q){

        std::cout << "\n\nKomendy:\n"
            << "search, insert, delete, draw, mode, quit" << std::endl;
    

        std::string input = "";
        std::cin.sync();
        std::getline(std::cin, input);

        

        try{
            if(input == "search" || input == "s"){

                std::string s;
                std::cout << "Podaj klucz: ";
                std::cin.sync();
                std::getline(std::cin, s);
                int key = std::stoi(s);
                std::cout << std::endl;

                Node *n = t->treeSearch(t->root, key);

                if(n != nullptr){
                    std::cout << "Klucz "<< s << ": " << std::endl;
                    if(mode == 0) std::cout << n->sData << std::endl;
                    if(mode == 1) std::cout << n->dData << std::endl;
                    if(mode == 2) std::cout << n->iData << std::endl;
                }
                    
                else 
                    std::cout << "Wezel o tym kluczu nie istnieje" << std::endl;

            }
            else if(input == "insert" || input == "i"){
                std::string s;

                std::cout << "Podaj klucz: ";
                std::cin.sync();
                std::getline(std::cin, s);
                int key = std::stoi(s);
                std::cout << std::endl;

                std::string sData = "";

                if(mode == 0){
                    std::cout << "Podaj wartosc typu string: ";
                    std::cin.sync();
                    std::getline(std::cin, s);
                    sData = s;
                }
                
                double dData = 0.0;
                if(mode == 1){
                    std::cout << "Podaj wartosc typu double: ";
                    std::cin.sync();
                    std::getline(std::cin, s);
                    dData = std::stod(s);
                }
                
                int iData = 0;
                
                if(mode == 2){
                    std::cout << "Podaj wartosc typu int: ";
                    std::cin.sync();
                    std::getline(std::cin, s);
                    iData = std::stoi(s);
                }

                
                
                
                Node* node = new Node(key, sData, dData, iData);

                Node *n = t->treeSearch(t->root, key);

                if(n != nullptr){
                    std::cout << "Wezel o kluczu "<< s << " juz istnieje. Nadpisac? (y/n) : ";
                    std::cin.sync();
                    std::getline(std::cin, s);
                    if(s == "y" || s == "Y") {
                        t->treeDelete(n);
                        
                        t->treeInsert(node);
                    }
                }
                else t->treeInsert(node);


            }
            else if(input == "delete" || input == "x"){
                std::string s;

                std::cout << "Podaj klucz: ";
                std::cin.sync();
                std::getline(std::cin, s);
                int key = std::stoi(s);
                std::cout << std::endl;
                t->treeDelete(t->treeSearch(t->root, key));

            }
            else if(input == "draw" || input == "d"){
                t->print("", t->root);
            }
            else if(input == "quit" || input == "q"){
                delete t;
                return 0;
            }
            else if(input == "mode" || input == "m"){
                std::string s;

                std::cout << "Podaj tryb (string/double/int): ";
                std::cin.sync();
                std::getline(std::cin, s);

                if(s == "string"){
                    mode = 0;
                    t = sTree;
                } 
                else if(s == "double"){
                    mode = 1;
                    t = dTree;
                } 
                else if(s == "int") {
                    mode = 2;
                    t = iTree;
                }
                else std::cout << "Nieprawidlowy tryb";
            }
            else {
                std::cout << "Podaj prawidlowa opcje" << std::endl;
            }
        }
        catch(std::exception e){
            std::cout << "Podaj prawidlowa wartosc" << std::endl;
        }
        
    }
    

}