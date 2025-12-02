import java.io.*;
import java.net.*;
 
/**Glowny program serwera */
public class ServerProgram {
 
    public static void main(String[] args) {
        //Tworzenie obslugiwanych drzew binarnych 
        BinaryTree<String> stringTree = new BinaryTree<String>();
        BinaryTree<Double> doubleTree = new BinaryTree<Double>();
        BinaryTree<Integer> integerTree = new BinaryTree<Integer>();
        int port = 1234;
        try{
            port = Integer.parseInt(args[1]);
        }
        catch(Exception ex){

        }
    
        try (ServerSocket serverSocket = new ServerSocket(port)) { //Otwieranie serwera
 
            System.out.println("port: " + port);
 
            while (true) {
                Socket socket = serverSocket.accept();
                System.out.println("New client connected");
 
                new MultiThread(socket, stringTree, doubleTree, integerTree).start(); // rozpoczynanie watku obslugi sklienta
            }
 
        } catch (IOException ex) {
            System.out.println("Server exception: " + ex.getMessage());
            ex.printStackTrace();
        }
    }
}