import java.io.PrintWriter;
/**
 * Klasa drzewa binarnego
 */
public class BinaryTree<T> {
    /** korzen drzewa
     */
    public Node<T> root;

    /** Konstruktor drzewa binarnego */
    BinaryTree(){
        root = null;
    }
    /**
     * Rekurencyjne yszukiwanie na drzewie
     * @param x poczatkowy wezel
     * @param key klucz szukanego wezla
     * @return szukany wezel
     */
    public Node<T> treeSearch(Node<T> x, int key){
        if(x == null || key == x.key){
            return x;
        }
        if(key < x.key) return treeSearch(x.left, key);
        else return treeSearch(x.right, key);
    }

    /**
     * wstawienie nowego wezla
     * @param z nowy wezel
     */
    public void treeInsert(Node<T> z){
        Node<T> y = null;
        Node<T> x = root;
        while(x != null){
            y = x;
            if(z.key < x.key) x = x.left;
            else x = x.right;
        }
        z.parent = y;
        if(y == null) root = z;
        else if(z.key < y.key) y.left = z;
        else y.right = z;
    }

    /**
     * usuwanie węzla z drzewa
     * @param z węzel o usuniecia
     */
    public void treeDelete(Node<T> z){
        Node<T> x, y;
        if(z.left == null || z.right == null) y = z;
        else y = treeSuccessor(z);
        if(y.left != null) x = y.left;
        else x = y.right;
        if(x != null) x.parent = y.parent;
        if(y.parent == null) root = x;
        else if(y == (y.parent).left) (y.parent).left = x;
        else (y.parent).right = x;
        if(y != z){
            z.key = y.key;
            z.setData(y.getData());
        }
    }
    /**
     * nastepnik
     * @param x poczatkowy wezel
     * @return nastepnik
     */
    public Node<T> treeSuccessor(Node<T> x){
        Node<T> sx = x;
        if(x.right != null) return treeMinimum(x.right);
        Node<T> y = x.parent;
        while(y != null && x == y. right){
            x = y;
            y = y.parent;
        }
        x = sx;
        return y;
    }
    /**
     * funkcja minimum drzewa
     * @param x początkowy wezel
     * @return minimum
     */
    public Node<T> treeMinimum(Node<T> x){
        Node<T> y = x.left;
        if(y == null) return x;
        while(y != null){
            y = y.left;
        }
        return y;
    }

    /**
     * Wypisywanie drzewa na konsoli
     * @param prefix przedrostek drzewa, zwiekszajacy sie rekrencyjnie
     * @param node początkowy węzeł
     * @param isLeft 
     */
    public void print(String prefix, Node<T> node, boolean isLeft){  
        if(node != null){
            System.out.println(prefix + (isLeft ? "|----" : "'----") + node.key);
            print(prefix + (isLeft ? "|    " : "     "), node.left, true);
            print(prefix + (isLeft ? "|    " : "     "), node.right, false);
        }
    }
    /**
     * Wypisywanie drzewa na okreslonym wyjsciu 
     * @param out wyjście 
     * @param prefix    przedrostek drzewa, zwiekszajacy sie rekrencyjnie
     * @param node  początkowy węzeł
     * @param isLeft
     */
    public void printOnSocket(PrintWriter out, String prefix, Node<T> node, boolean isLeft){  
        
        if(node != null){
            out.println(prefix + (isLeft ? "|----" : "'----") + node.data);


            printOnSocket(out, prefix + (isLeft ? "|    " : "     "), node.left, true);
            printOnSocket(out, prefix + (isLeft ? "|    " : "     "), node.right, false);
        }
    }
}