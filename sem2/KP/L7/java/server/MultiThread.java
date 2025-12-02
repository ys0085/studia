import java.io.*;
import java.net.*;

/**Klasa wątku obslugujacego klienta */
public class MultiThread extends Thread{
    private Socket socket;
    private BinaryTree<String> stringTree;
    private BinaryTree<Double> doubleTree;
    private BinaryTree<Integer> integerTree;
    
    /**
     * Konstruktor
     * @param socket socket połączenia z klientem
     * @param sTree drzewo String
     * @param dTree drzewo Double
     * @param iTree drzewo Integer
     */
    public MultiThread(Socket socket, BinaryTree<String> sTree, BinaryTree<Double> dTree, BinaryTree<Integer> iTree) {
        this.socket = socket;
        stringTree = sTree;
        doubleTree = dTree;
        integerTree = iTree;
        
    }

    private enum TreeMode{
        STRING,
        DOUBLE,
        INTEGER
    }

    private TreeMode mode;
    
    private Node<?> search(int key){
        
        if(mode == TreeMode.STRING)
            return stringTree.treeSearch(stringTree.root, key);
        if(mode == TreeMode.DOUBLE)
            return doubleTree.treeSearch(doubleTree.root, key);
        if(mode == TreeMode.INTEGER)
            return integerTree.treeSearch(integerTree.root, key);

        return null;
        
    }

    private void insert(int key, String data){
        if(mode == TreeMode.STRING)
            stringTree.treeInsert(new Node<String>(key, data));
        if(mode == TreeMode.DOUBLE)
            doubleTree.treeInsert(new Node<Double>(key, Double.parseDouble(data)));
        if(mode == TreeMode.INTEGER)
            integerTree.treeInsert(new Node<Integer>(key, Integer.parseInt(data)));
    }

    private void delete(int key){
        if(mode == TreeMode.STRING)
            stringTree.treeDelete(stringTree.treeSearch(stringTree.root, key));
        if(mode == TreeMode.DOUBLE)
            doubleTree.treeDelete(doubleTree.treeSearch(doubleTree.root, key));
        if(mode == TreeMode.INTEGER)
            integerTree.treeDelete(integerTree.treeSearch(integerTree.root, key));
    }

    private void draw(PrintWriter out){
        if(mode == TreeMode.STRING)
            stringTree.printOnSocket(out, "", stringTree.root, false);
        if(mode == TreeMode.DOUBLE)
            doubleTree.printOnSocket(out, "", doubleTree.root, false);
        if(mode == TreeMode.INTEGER)
            integerTree.printOnSocket(out, "", integerTree.root, false);
        
    }

    public void run() {
        mode = TreeMode.STRING;
        try {
             //Odbieranie od socketa
            InputStream input = socket.getInputStream();
            BufferedReader in = new BufferedReader(new InputStreamReader(input));
    
            //Wysylanie do socketa
            OutputStream output = socket.getOutputStream();
            PrintWriter out = new PrintWriter(output, true);
        
            String line, args[];
            
            do {
                
                //pobieranie danych z klienta
                line = in.readLine();
                args = line.split(" ");
                System.out.println("input: " + line);
                
                if(args[0].equals("search")){ //funkcjonalnosc search
                    int key = Integer.parseInt(args[1]);

                    Node<?> node = search(key);
                    if(node == null){
                        out.println("null");
                    }
                    else{
                        out.println(key);
                        out.println(node.getData());
                    }
                }
                else if(args[0].equals("insert")){ //funkcjonalnosc insert
                    int key = Integer.parseInt(args[1]);
                    String data = in.readLine();
                    
                    Node<?> n = search(key);
                    if(n != null) delete(key);
                    try {
                        insert(key, data);
                    } catch (NumberFormatException e) {
                    }
                    
                    
                }
                else if(args[0].equals("delete")){ //funkcjonalnosc delete
                    int key = Integer.parseInt(args[1]);
                    delete(key);
                }
                else if(args[0].equals("draw")){ //funkcjonalnosc draw
                    out.println("Printing tree...");
                    draw(out);
                    out.println("end");
                }
                else if(args[0].equals("mode")){ //zmiana trybu drzewa
                    if(args[1].equals("s")) mode = TreeMode.STRING;
                    if(args[1].equals("d")) mode = TreeMode.DOUBLE;
                    if(args[1].equals("i")) mode = TreeMode.INTEGER;
                }
                
                
            } while (!line.equals("quit"));
    
            //odlaczanie klienta
            socket.close();
            System.out.println("Client disconnected");
        } 
        catch (IOException ex) {
            System.out.println("Server exception: " + ex.getMessage());
            ex.printStackTrace();
        }
        catch (NumberFormatException ex){
            System.out.println("Wrong data: " + ex.getMessage());
            ex.printStackTrace();
        }
        
    }
}
