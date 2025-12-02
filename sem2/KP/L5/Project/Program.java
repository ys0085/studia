import java.io.File;
import java.lang.Runtime.Version;

import javafx.application.Application;
import javafx.geometry.Insets;
 
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.control.Alert.AlertType;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.scene.layout.Background;
import javafx.scene.layout.ColumnConstraints;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.Pane;
import javafx.scene.layout.Priority;
import javafx.scene.layout.RowConstraints;
import javafx.scene.paint.Color;
import javafx.scene.shape.Rectangle;
import javafx.scene.shape.Circle;
import javafx.scene.shape.Polygon;
import javafx.scene.shape.Shape;
import javafx.stage.FileChooser;
import javafx.stage.FileChooser.ExtensionFilter;
import javafx.stage.Stage;

/**
 * Glowna klasa programu
 * @version 1.0.1
 * @author Krzysztof Kleszcz
 */
public class Program extends Application {
    /** Wersja programu*/
    private Version version = Version.parse("1.0.1");

    
    /** Funkcja start
     * @param stage glowna scena
     */
    @Override
    public void start(Stage stage) {
        stage.setMinWidth(640);
        stage.setMinHeight(480);
        stage.setTitle("SuperKształty " + version.toString());
        Insets globalPadding = new Insets(5, 5, 5, 5);

        // -- INTERFEJS -- 

        //ustawianie paska kształtów z prawej strony

        GridPane shapesbar = new GridPane(); 

        
        ColumnConstraints buttonColumn = new ColumnConstraints();   //rozmieszczanie siatki
        buttonColumn.setPercentWidth(100);
        
        RowConstraints buttonRows = new RowConstraints();
        buttonRows.setPercentHeight(100.0 / 3);     
        
        shapesbar.setBackground(Background.fill(Color.AQUAMARINE));
        shapesbar.getColumnConstraints().add(buttonColumn);
        shapesbar.getRowConstraints().addAll(buttonRows);
        shapesbar.setVgap(10);
        shapesbar.setPadding(globalPadding);
    


        
        Button b1 = new Button("", new ImageView(new Image("icons/rectangle.png"))); //inicjalizacja przycisków
        b1.setPrefSize(10000, 10000);
        
       
        Button b2 = new Button("", new ImageView(new Image("icons/circle.png")));
        b2.setPrefSize(10000,10000);
        
        
        Button b3 = new Button("", new ImageView(new Image("icons/triangle.png")));
        b3.setPrefSize(10000,10000);
        
        
        GridPane.setConstraints(b1, 0, 0);
        GridPane.setConstraints(b2, 0, 1);
        GridPane.setConstraints(b3, 0, 2);
        

        shapesbar.getChildren().addAll(b1, b2, b3); //dodawanie przyciskow do siatki

        //koniec paska kształtów

        
        //inicjalizacja płótna

        
        Pane canvas = new Pane();
        canvas.setPrefSize(10000, 10000);
        canvas.setStyle("-fx-border-color:#000000; -fx-border-width:1px; -fx-background-color:#ffffff");

        GridPane.setVgrow(canvas, Priority.ALWAYS); //płótno ma mieć priorytet w powiększaniu się, w przypadku dziwnych ustawień okna

        //koniec inicjalizacji płótna
    

        //ustawianie paska narzędzi na dole

        GridPane toolbar = new GridPane();
        
        RowConstraints toolRowConstraints = new RowConstraints(); // rozmieszczanie siatki
        toolRowConstraints.setPercentHeight(100);
        ColumnConstraints toolColumnConstraints = new ColumnConstraints();
        toolColumnConstraints.setPercentWidth((double)100/6);

        toolbar.setBackground(Background.fill(Color.AQUAMARINE));
        toolbar.getColumnConstraints().addAll(toolColumnConstraints);
        toolbar.getRowConstraints().add(toolRowConstraints);
        toolbar.setHgap(10);
        toolbar.setPadding(globalPadding);

        Button t1 = new Button("", new ImageView(new Image("icons/ccw.png"))); // inicjalizacja przycisków 
        t1.setPrefSize(10000, 10000);
        
        Button t2 = new Button("", new ImageView(new Image("icons/cw.png")));
        t2.setPrefSize(10000, 10000);

        ColorPicker t3 = new ColorPicker(); 
        t3.setPrefSize(10000, 10000);
        t3.setValue(Color.BLACK);  //ustawienie default koloru na czarny
        
        Button t4 = new Button("Save");
        t4.setPrefSize(10000, 10000);

        Button t5 = new Button("Load");
        t5.setPrefSize(10000, 10000);

        Button t6 = new Button("Instrukcje");
        t6.setPrefSize(10000, 10000);
        


        GridPane.setConstraints(t1, 0, 0);
        GridPane.setConstraints(t2, 1, 0);
        GridPane.setConstraints(t3, 2, 0);
        GridPane.setConstraints(t4, 3, 0);
        GridPane.setConstraints(t5, 4, 0);
        GridPane.setConstraints(t6, 5, 0);
        
        toolbar.getChildren().addAll(t1, t2, t3, t4, t5, t6); // dodawanie przycisków do siatki

        //koniec paska narzędzi


        // inicjalizacja przycisku info

        Button infoButton = new Button("About"); 
        infoButton.setStyle("-fx-border-color:#000000; -fx-border-width:2px; -fx-background-color:#eeeeee; -fx-border-radius:6px");
        infoButton.setPrefSize(10000, 10000);
        
        // koniec przycisku info


        // siatka bazowa, zajmująca całe okno

        GridPane base = new GridPane();

        ColumnConstraints canvasColumn = new ColumnConstraints();
        canvasColumn.setPercentWidth(90);
        RowConstraints canvasRow = new RowConstraints();
        canvasRow.setPercentHeight(90);     //ustawianie rozmiarow glownej siatki

        base.getColumnConstraints().add(canvasColumn);
        base.getRowConstraints().add(canvasRow);
        
        GridPane.setConstraints(canvas, 0, 0);
        GridPane.setConstraints(toolbar, 0, 1);
        GridPane.setConstraints(shapesbar, 1, 0);
        GridPane.setConstraints(infoButton, 1, 1);

        base.getChildren().addAll(canvas, toolbar, shapesbar, infoButton); //dodawanie elementów do siatki
        
        canvas.toBack(); //ustawienie płótna za paskami narzędzi
        
        //koniec siatki bazowej

        // -- KONIEC INTERFEJSU -- 

        

        // -- DZIAŁANIE PROGRAMU -- 

        Label currentShapeLabel = new Label(); // napis wskazujący na wybrany kształt
        canvas.getChildren().add(currentShapeLabel);

        ShapeHandler sh = new ShapeHandler(canvas); //inicjalizacja klas ShapeHandler i ShapeMaker
        ShapeMaker sm = new ShapeMaker(sh, currentShapeLabel);


        
        //Dzialanie przyciskow tworzenia ksztaltow
        b1.setOnAction(e -> {   //tworzenie Prostokąta
            if(sm.getCurrentShape() == ShapeMaker.Shapes.RECTANGLE){
                sm.setShape(ShapeMaker.Shapes.NONE);
            } 
            else {
                sm.setShape(ShapeMaker.Shapes.RECTANGLE);
                sh.setAsActive(null, false);
                sm.addShapes(canvas);
            }
        });


        b2.setOnAction(e -> {   //tworzenia koła
            if(sm.getCurrentShape() == ShapeMaker.Shapes.CIRCLE){
                sm.setShape(ShapeMaker.Shapes.NONE);
                sm.addShapes(canvas);
                
            }
            else{
                sm.setShape(ShapeMaker.Shapes.CIRCLE);
                sh.setAsActive(null, false);
                sm.addShapes(canvas);
            }
        });

        b3.setOnAction(e -> {   //tworzenie trójkąta
            if(sm.getCurrentShape() == ShapeMaker.Shapes.TRIANGLE){
                sm.setShape(ShapeMaker.Shapes.NONE);
                sm.addShapes(canvas);
                
            }
            else{
                sm.setShape(ShapeMaker.Shapes.TRIANGLE);
                sh.setAsActive(null, false);
                sm.addShapes(canvas);
            }
        });


        //Dzialanie przyciskow narzedzi

        double rotateAmount = 5;
        t1.setOnAction(e -> {   // przyciski obrotu
            sh.rotateShape(sh.getActiveShape(), -rotateAmount);
        });

        t2.setOnAction(e -> {
            sh.rotateShape(sh.getActiveShape(), rotateAmount);
        });
        
        t3.setOnAction(e -> {   // zmiana koloru
            sh.setCurrentColor(t3.getValue());
            sh.colorShape(sh.getActiveShape());
        });

        t4.setOnAction(e -> {   // zapis do pliku
            Shape activeShape = sh.getActiveShape();  
            if(activeShape != null){    //sprawdzanie czy istnieje aktywny kształt
                FileChooser chooser = new FileChooser();    //tworzenie FileChooser
                chooser.setTitle("Zapisz plik");
                chooser.setInitialDirectory(new File("saves"));
                ExtensionFilter shp = new ExtensionFilter("kształt (*.shp)", "*.shp");
                chooser.getExtensionFilters().add(shp);

                File file = chooser.showSaveDialog(stage);  //pokazanie okna eksploratora plikow

                if(file == null) return;

                String name = activeShape.getClass().getName(); //zapis za uzyciem roznych metod saveShape w zaleznosci od rodzaju kształtu
                if(name.contains("Rectangle")){
                    sh.saveShape((Rectangle) sh.getActiveShape(), file);
                }
                else if(name.contains("Circle")){
                    sh.saveShape((Circle) sh.getActiveShape(), file);
                }
                else if(name.contains("Polygon")){
                    sh.saveShape((Polygon) sh.getActiveShape(), file);
                }
            }
        });
        
        t5.setOnAction(e -> {   // odczyt z pliku
            FileChooser chooser = new FileChooser(); //tworzenie FileChooser
            chooser.setTitle("Wybierz plik");
            chooser.setInitialDirectory(new File("saves"));
            ExtensionFilter shp = new ExtensionFilter("kształt (*.shp)", "*.shp");
            chooser.getExtensionFilters().add(shp);

            File file = chooser.showOpenDialog(stage); // pokazanie FileChooser
            
            if(file == null) return;

            Shape loadedShape = sh.loadShape(file);

            if(loadedShape != null){ //łapanie źle załadowanych plików
                sh.init(loadedShape);
            }
            else System.out.println("failed to load shape"); //informacja w konsoli w razie źle odczytanego pliku
            
        });

        Alert instructions = new Alert(AlertType.INFORMATION); //tworzenie alertu z instrukcjami
        instructions.setTitle("Instrukcje");
        instructions.setHeaderText("Instrukcje");
        instructions.setContentText(""
            + "Użyj przycisków po prawej stronie, aby wybrać kształt:\n" 
            + "Prostokąt - zaznacz jego przeciwległe wierzchołki\n" 
            + "Koło - zaznacz środek oraz punkt na krawędzi koła\n"
            + "Trójkąt - zaznacz jego wierzchołki\n\n"
            + "Zaznacz kształt prawym przyciskiem myszy.\n"
            + "Usunięcie kształtu - Delete\n"
            + "Zmiana rozmiaru - Scroll\n"
            + "Obrót kształtu - przyciski obrotu na pasku narzędzi\n"
            + "Zmiana koloru - menu z kolorami na pasku narzędzi\n"
            + "Zapis i odczyt z pliku - przyciski Save i Load\n"
            + "Informacje o programie - przycisk Info"
            );

        t6.setOnAction(e -> { // pokazanie instrukcji
            instructions.show();
        });

        Alert info = new Alert(AlertType.INFORMATION); //tworzenie alertu info
        info.setTitle("Info");
        info.setHeaderText("SuperKształty");
        info.setContentText("" 
        + "Program do edycji podstawowych kształtów\n"
        + "Autor: Krzysztof Kleszcz\n"
        + "Wersja: " + version.toString()
        );
        
        infoButton.setOnAction(e -> { // pokazanie alertu info
            info.show();
        });

        Scene scene = new Scene(base, 640, 480); //tworzenie sceny i pokazanie jej
        stage.setScene(scene);
        stage.show();
    }
    

    
    /** Funkcja glowna programu rozpoczynajaca aplikacje
     * @param args argumenty funkcji main
     */
    public static void main(String[] args) {
        launch();
    }

}