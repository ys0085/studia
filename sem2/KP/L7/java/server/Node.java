
public class Node<T>{
    
    T data;

    public int key;

    public Node<T> right;
    public Node<T> left;
    public Node<T> parent;

    /**
     * Konstruktor wezla z danymi
     * @param k klucz wezla
     * @param data dane wezla
     */
    Node(int k, T data){
        key = k;
        right = left = parent = null;
        this.data = data;
    }
    /**
     * Podstawowy konstruktor wezla
     * @param k klucz wezla
     */
    Node(int k){
        key = k;
        right = left = parent = null;
        this.data = null;
    }

    /**
     * Pobieranie danych z wezla
     * @return dane wezla
     */
    public T getData(){
        return data;
    }
    /**
     * ustawianie danych wezla
     * @param data dane wezla
     */
    public void setData(T data){
        this.data = data;
    }
    
}
