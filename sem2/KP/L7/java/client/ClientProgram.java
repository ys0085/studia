import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import java.net.UnknownHostException;
/**
 * Glowny program klienta
 */
public class ClientProgram {
    public static void main(String args[]){
        
        try {
            boolean quit = false;
            
            Socket socket = new Socket("localhost", 1234); 
            // Wysylanie do serwera
            PrintWriter out = new PrintWriter(socket.getOutputStream(), true);
            // Odbieranie z serwera
            BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
            
            //Czytanie z konsoli
            BufferedReader bfn = new BufferedReader(new InputStreamReader(System.in));
            
            int mode = 0;
            String line;
            do {
                
                System.out.println("search, insert, delete, draw, mode, quit");
                
                line = bfn.readLine();

                try {
                    if(line.equalsIgnoreCase("search")){ //funkcjonalnosc search
                        System.out.println("Podaj klucz: ");
                        line = bfn.readLine();
                        int key = Integer.parseInt(line);
                        
                        out.println("search " + key);
                        String s = in.readLine();
                        if(s.equals("null")){
                            System.out.println("Wezel o kluczu " + key + " nie istnieje");
                            
                        }
                        else {
                            System.out.println("Wezel o kluczu " + key + ":");
                            System.out.println("Dane: " + in.readLine());
                            
                        }
                    }
                    else if(line.equalsIgnoreCase("insert")){ //funkcjonalnosc insert
                        System.out.println("Podaj klucz: ");
                        line = bfn.readLine();
                        int key = Integer.parseInt(line);
        
                        System.out.println("Podaj wartosc: ");
                        line = bfn.readLine();
                        String data = line;
        

                        if(mode == 0){
                            out.println("insert " + key);
                            out.println(data);
                        }
                        if(mode == 1){
                            double d = Double.parseDouble(data);
                            out.println("insert " + key);
                            out.println(d);
                        }
                        if(mode == 2){
                            int i = Integer.parseInt(data);
                            out.println("insert " + key);
                            out.println(i);
                        }
                        
                        
                        
        
                        
                    }
                    else if(line.equalsIgnoreCase("delete")){ //funkcjonalnosc delete
                        System.out.println("Podaj klucz: ");
                        line = bfn.readLine();
                        int key = Integer.parseInt(line);
                        
                        out.println("delete " + key);
                    }
                    else if(line.equalsIgnoreCase("draw")){ //funkcjonalnosc draw
                        out.println("draw");
                        String s;
                        while(!(s = in.readLine()).equals("end")){
                            System.out.println(s);
                        }
                    }
                    else if(line.equalsIgnoreCase("quit")){ //wyjscie z programu
                        out.println("quit");
                        quit = true;
                    }
                    else if(line.equalsIgnoreCase("mode")){ //zmiana trybu
                        System.out.println("Podaj tryb (string/double/int): ");
                        line = bfn.readLine();
                        if(line.equalsIgnoreCase("string")){
                            out.println("mode s");
                            mode = 0;
                        }
                        if(line.equalsIgnoreCase("double")){
                            out.println("mode d");
                            mode = 1;
                        }
                        if(line.equalsIgnoreCase("int") || line.equalsIgnoreCase("integer")){
                            out.println("mode i");
                            mode = 2;
                        }
                    }
                    else {
                        System.out.println("Nieprawidlowa akcja");
                    }
                    
                    line = "";

                } catch (NumberFormatException e) {
                    System.out.println("Nieprawidlowe dane");
                }
        
            } while (!quit);
            
            //rozlaczanie z serwerem
            socket.close();
            System.out.println("Connection closed");

        } 
        catch (UnknownHostException ex) {
            System.out.println("Server not found: " + ex.getMessage());
        }
        catch (IOException ex) {
            System.out.println("I/O exception: " + ex.getMessage());
        }
        

    }
}