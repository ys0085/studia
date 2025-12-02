import java.util.Random;

import javafx.application.Application;
import javafx.scene.Scene;
import javafx.scene.layout.ColumnConstraints;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.RowConstraints;
import javafx.stage.Stage;


/**
 * Glowna klasa programu
 * @version 1.0.1
 * @author Krzysztof Kleszcz
 */
public class Program extends Application{
    /**Podstawowa wielkosc pola */
    public static double FIELD_SIZE = 150.0;
    /**Glowny generator losowy programu */
    public static Random RAND = new Random();
    /**Okres dzialania wątków */
    public static double TIME;
    /**Prawdopodobienstwo akcji wątku */
    public static double PROB;
    /**Wymiary pola programu */
    public static int m, n;

    /**Tablica głównych wątków programu */
    public static ColorField fields[];
    
    /** Funkcja start
     * @param stage glowna scena
     */
    public void start(Stage stage){

        //tworzenie siatki pól
        GridPane mainGrid = new GridPane();
        mainGrid.setHgap(0);
        mainGrid.setVgap(0);
        if(FIELD_SIZE * Math.max(m, n) > 900) FIELD_SIZE = (double) 640 / Math.max(m, n);

        RowConstraints rows = new RowConstraints(FIELD_SIZE);
        ColumnConstraints columns = new ColumnConstraints(FIELD_SIZE);
        
        mainGrid.getRowConstraints().add(rows);
        mainGrid.getColumnConstraints().add(columns);

        //inicjalizacja wątków

        fields = new ColorField[n*m];

        
        
        for(int i = 0; i < m; i++){
            for(int j = 0; j < n; j++){
                int index = i * n + j;
                fields[index] = new ColorField(index);
                GridPane.setConstraints(fields[index].getPane(), i, j);
                mainGrid.getChildren().add(fields[index].getPane());
            }
        }

        //ustalanie sąsiadów każdego wątku

        
        for(int i = 0; i < m * n; i++){
            int col = i / n;
            int row = i % n;
            int right, left, down, up;
            
            right = i + n;
            left = i - n;
            
            down = i + 1;
            up = i - 1;

            if(col == m - 1){
                right = row; 
                left = i - n;
            }
            if(col == 0){
                right = i + n;
                left = row + (m - 1) * n;
            }
            if(m == 1){
                right = i;
                left = i;
            }
            

            if(row == n - 1){
                down = col * n;
                up = i - 1;
            }
            if(row == 0){
                down = i + 1;
                up = i + n - 1;
            }
            if(n == 1){
                up = i;
                down = i;
            }
            
            fields[i].setNeighbors(fields[right], fields[left], fields[down], fields[up]);
        }
        
        //rozpoczynanie wątków przy pokazaniu sceny
        stage.setOnShowing(e -> {
            for(int i = 0; i < m * n; i++){
                fields[i].start();
                while(!fields[i].isRunning()){
                    try {
                        Thread.sleep(1);
                    } catch (InterruptedException e1) {
                        // TODO Auto-generated catch block
                        e1.printStackTrace();
                    }
                }
            }
        });

        //zamykanie wątków przy zamknieciu sceny
        stage.setOnCloseRequest(e -> {
            for(int i = 0; i < m * n; i++){
                fields[i].stopT();
            }
            System.exit(0);
        });

        Scene mainScene = new Scene(mainGrid, m * FIELD_SIZE, n * FIELD_SIZE); //tworzenie sceny i pokazanie jej
        
        stage.setResizable(false);
        stage.setScene(mainScene);
        stage.show();

    }
    
    /** Funkcja glowna programu rozpoczynajaca aplikacje
     * @param args argumenty funkcji main
     */
    public static void main(String[] args){
        
            try { //przetwarzanie podanych argumentow
                m = Integer.parseInt(args[0]);
                if(m < 1) throw new Exception();
                n = Integer.parseInt(args[1]);
                if(n < 1) throw new Exception();
                TIME = Integer.parseInt(args[2]);
                if(TIME <= 0) throw new Exception();
                PROB = Double.parseDouble(args[3]);
                if(PROB < 0 || PROB > 1) throw new Exception();

            } 
            catch (Exception e) {
                System.out.println("Podaj prawidlowe parametry: m, n, k, p" + 
                "\n m > 0, n > 0 - wymiary pola" +
                "\n k > 0 - czas w milisekundach" + 
                "\n 0 < p < 1 - prawdopodobienstwo");

                System.exit(0);
            }

        launch();
    
    }
}