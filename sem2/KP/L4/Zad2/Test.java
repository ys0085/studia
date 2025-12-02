
import javafx.application.Application;
import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.layout.ColumnConstraints;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.RowConstraints;
import javafx.stage.Stage;

public class Test extends Application {

    @Override
    public void start(Stage stage) {
        Label description = new Label("Wpisz liczbe: ");
        TextField input = new TextField();
        Button b = new Button("Run");
        Label numbers[][] = new Label[100][100];
        int MAX_COL_WIDTH = 30;
        

        GridPane text = new GridPane();

        GridPane.setConstraints(description, 1, 1);
        GridPane.setConstraints(input, 2, 1);
        GridPane.setConstraints(b, 3, 1);

        ColumnConstraints col0 = new ColumnConstraints(10);
        ColumnConstraints col1 = new ColumnConstraints(150);
        ColumnConstraints col2 = new ColumnConstraints(100);
        ColumnConstraints col3 = new ColumnConstraints(100);

        RowConstraints row0 = new RowConstraints(10);
        RowConstraints row1 = new RowConstraints(40);
        
        text.getChildren().addAll(description, input, b);
        text.getColumnConstraints().addAll(col0, col1, col2, col3);
        text.getRowConstraints().addAll(row0, row1);
        
        GridPane triangle = new GridPane();

        triangle.getColumnConstraints().add(new ColumnConstraints(20));

        GridPane all = new GridPane();

        GridPane.setConstraints(text, 0, 0);
        GridPane.setConstraints(triangle, 0, 1);

        all.getChildren().addAll(text, triangle);

        Scene scene = new Scene(all, 640, 480);
        b.setOnAction(new EventHandler<ActionEvent>() {
            @Override
            public void handle(ActionEvent e){
                try {
                    triangle.getChildren().clear();
                    int num = Integer.parseInt(input.getText());
                    
                    for(int i = 0; i<num; i++){
                        WierszTrojkataPascala wtp = new WierszTrojkataPascala(i);
                        for(int j = 0; j < i+1; j++){
                            numbers[i][j] = new Label(Integer.toString(wtp.wartosc(j)));
                            numbers[i][j].minWidth(40);
                            GridPane.setConstraints(numbers[i][j], num - i + 2*j, i);
                            triangle.getChildren().add(numbers[i][j]);
                        }
                        triangle.getColumnConstraints().addAll(new ColumnConstraints(MAX_COL_WIDTH), new ColumnConstraints(MAX_COL_WIDTH));

                    }
                    
                } 
                catch (InvalidRowNumberException ex) {
                    System.out.println("podaj liczbe dodatnia");
                } 
                catch (IndexOutOfRangeException ex) {
                    System.out.println("??");
                }
                catch (NumberFormatException ex) {
                    System.out.println("nieprawidlowa liczba");
                }
            }
        });
        stage.setScene(scene);
        stage.show();
    }

    public static void main(String[] args) {
        launch();
    }

}