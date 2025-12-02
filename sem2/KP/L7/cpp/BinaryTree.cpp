#include<iostream>

class Node{
    public:
        std::string sData;
        double dData;
        int iData;

        int key;    
        Node* right;
        Node* left;
        Node* parent;

        Node(int k, std::string s = "", double d = 0, int i = 0){
            key = k; 
            right = left = parent = nullptr;
            sData = s;
            dData = d;
            iData = i;
        }
        ~Node(){
            delete right;
            delete left;
            delete parent;
        }

};


class BinaryTree{
    public:
        Node* root;

        Node* treeSearch(Node* x, int key){
            if(x == nullptr || key == x->key){
                return x;
            }
            if(key < x->key) return treeSearch(x->left, key);
            else return treeSearch(x->right, key);
        }

        void treeInsert(Node* z){
            if(root == nullptr) {
                root = z;
                return;
            }
            Node* y = nullptr;
            Node* x = root;
            while(x != nullptr){
                y = x;
                if(z->key < x->key) x = x->left;
                else x = x->right;
            }
            z->parent = y;
            if(y == nullptr) root = z;
            else if(z->key < y->key) y->left = z;
            else y->right = z;
        }

        void treeDelete(Node* z){
            Node* x, * y;
            if(z->left == nullptr || z->right == nullptr) y = z;
            else y = treeSuccessor(z);
            if(y->left != nullptr) x = y->left;
            else x = y->right;
            if(x != nullptr) x->parent = y->parent;
            if(y->parent == nullptr) root = x;
            else if(y == (y->parent)->left) (y->parent)->left = x;
            else (y->parent)->right = x;
            if(y != z){
                z->key = y->key;
                z->dData = y->dData;
                z->sData = y->sData;
                z->iData = y->iData;
            }
        }

        Node* treeSuccessor(Node* x){
            Node sx = *x;
            if(x->right != nullptr) return treeMinimum(x->right);
            Node* y = x->parent;
            while(y != nullptr && x == y-> right){
                x = y;
                y = y->parent;
            }
            *x = sx;
            return y;
        }

        Node* treeMinimum(Node* x){
            Node* y = x->left;
            if(y == nullptr) return x;
            while(y != nullptr){
                y = y->left;
            }
            return y;
        }

        void print(const std::string prefix, Node* node, bool isLeft = false){  
            if(node != nullptr){
                std::cout << prefix;

                std::cout << (isLeft ? "|----" : "'----");

                std::cout << node->key << std::endl;

                print(prefix + (isLeft ? "|    " : "     "), node->left, true);
                print(prefix + (isLeft ? "|    " : "     "), node->right, false);
            }
        }

        BinaryTree(){
            root = nullptr;
        }

        ~BinaryTree(){
            delete root;
        }
};


