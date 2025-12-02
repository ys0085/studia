import javafx.scene.input.KeyCode;
import javafx.scene.input.MouseButton;
import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;
import javafx.scene.shape.Rectangle;
import javafx.scene.shape.Circle;
import javafx.scene.shape.Polygon;
import javafx.scene.shape.Shape;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;

/**
 * Klasa odpowiedzialna za zmiane stanu ksztaltow
 */

public class ShapeHandler {
    /**Glowne plotno */
    private Pane canvas;
    /**Aktualnie aktywny ksztalt */
    private Shape currentActiveShape;
    /**Aktynwny kolor ksztaltu */
    private Color currentColor = Color.BLACK;
    

    /**
     * Funkcja inicjalizujaca ksztalt i dodajaca go do plotna
     * @param s ksztalt
     */
    public void init(Shape s){ //inicjalizacja ksztaltu
        if(s == null) return;
        s.setFill(currentColor);
        canvas.getChildren().add(s);
        makeActivable(s);
        makeDraggable(s);
        makeDeletable(s);
        
    }

    /**
     * Funkcja umozliwiajaca oznaczenie ksztaltu jako aktywny
     * @param s ksztalt
     */
    private void makeActivable(Shape s){  //dodwanie dzialania prawego przycisku
        s.setOnMouseClicked(e -> {
            if(e.getButton() == MouseButton.SECONDARY){
                if(currentActiveShape == null || !(currentActiveShape.equals(s))){
                    setAsActive(s, true);
                }
                else setAsActive(s, false);
                
            }
        });
    }
    
    
    /** Funkcja oznaczajaca ksztalt jako aktywny
     * @param s ksztalt do (dez)aktywowania
     * @param active okresla, czy ksztalt ma byc aktywowany czy dezaktywowany
     */
    public void setAsActive(Shape s, boolean active){   //ustawianie kształtu jako aktywny
        if(s == null){ 
            if(currentActiveShape != null){
                setAsActive(currentActiveShape, false); 
            }
        }
        else{
            if(active){
                setAsActive(currentActiveShape, false);
                currentActiveShape = s;
                s.toFront();
                s.requestFocus();
                s.setStrokeWidth(3.0 / s.getScaleX());
                currentActiveShape.getStrokeDashArray().clear();     //dodaję do aktywnego kształtu czerwoną obwódkę
                currentActiveShape.getStrokeDashArray().addAll(5.0 / currentActiveShape.getScaleX());
                currentActiveShape.setStrokeDashOffset(3.0 / currentActiveShape.getScaleX());

                s.setStroke(Color.RED);
            }
            else{
                currentActiveShape = null;
                s.setStrokeWidth(0);
                s.setStroke(s.getFill());
                canvas.requestFocus();
            }
        }
    }

   

    private double dragOffsetX = 0;
    private double dragOffsetY = 0;
    
    /** 
     * Funkcja umozliwiajaca przeciaganie ksztaltu po plotnie
     * @param s ksztalt
     */
    private void makeDraggable(Shape s){
        s.setOnMousePressed(e -> {
            dragOffsetX = s.getLayoutX() - e.getSceneX();   //ustalam miejsce złapania za kształt jako offset
            dragOffsetY = s.getLayoutY() - e.getSceneY();
        }); 
        s.setOnMouseDragged(e -> {
            if(s.equals(currentActiveShape)){
                
                s.setLayoutX(e.getSceneX() + dragOffsetX);  //korekta o offset
                s.setLayoutY(e.getSceneY() + dragOffsetY);
            }
        });
    }

    
    
    /** Funkcja umozliwiajaca usuwanie ksztaltu
     * @param s ksztalt
     */
    private void makeDeletable(Shape s){
        s.setOnKeyPressed(e -> {
            if (e.getCode() == KeyCode.DELETE){ //usuwanie kształtu na klawiszu Delete
                canvas.getChildren().remove(s);
            }
        });
    }

    /**
     * Funkcja nadajaca mozliwosc zmiany rozmiaru aktywnego ksztaltu na scrollu
     */
    private void makeSizable(){ 
        canvas.setOnScroll(e -> {
            double scale = 1.1;
            double deltaY = e.getDeltaY();
            if(deltaY < 0){ //jeśli scroll idzie w dół (deltaY < 0), to wspolczynnik zmiany wielkości jest mniejszy od 1
                scale = 0.9;
            }
            if(currentActiveShape != null)
                changeSize(currentActiveShape, scale);
            
        });
    }    
    /**
     * Funkcja zmieniajaca rozmiar ksztaltu
     * @param s ksztalt
     * @param scale skala zmiany rozmiaru
     */
    private void changeSize(Shape s, double scale){
        s.setScaleX(s.getScaleX() * scale);
        s.setScaleY(s.getScaleY() * scale);
        s.setStrokeWidth(3.0 / s.getScaleX());        //aktualizacja czerwonej obwódki zgodnie ze zmianą wielkości kształtu
        s.getStrokeDashArray().clear();
        s.getStrokeDashArray().addAll(5.0 / s.getScaleX(), 5.0 / s.getScaleX()); 
        s.setStrokeDashOffset(3.0 / s.getScaleX());
    }

    /**
     * Funkcja zwracajaca aktywny ksztalt
     * @return aktywny ksztalt
     */
    public Shape getActiveShape(){ //zwraca aktywny kształt
        return currentActiveShape;
    }
    
    /**
     * Funkcja obracajaca ksztalt
     * @param s ksztalt
     * @param degrees ilosc stopni obrotu
     */
    public void rotateShape(Shape s, double degrees){ //obrót kształtem
        if(s != null){
            s.setRotate(s.getRotate() + degrees);
        }
    }

    /**
     * Funkcja zmieniajaca kolor ksztaltu na aktywny kolor {@link currentColor currentColor}
     * @param s ksztalt
     */
    public void colorShape(Shape s){ //zmiana koloru kształtu
        if(s != null){
            s.setFill(currentColor);
        }
    }
    
    /**
     * Funkcja zmieniajaca aktywny kolor {@link currentColor currentColor} na wybrany 
     * @param c nowy kolor
     */
    public void setCurrentColor(Color c){ //ustawienie koloru
        currentColor = c;
    }   

    /**
     * Funkcja zapisujaca prostokat do pliku
     * @param s prostokat do zapisu
     * @param file plik docelowy
     */
    public void saveShape(Rectangle s, File file){ // zmiana obiektu Shape w obiekt SavedShape oraz zapis do pliku 
        try(ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(file))){
            SavedShape ss = new SavedShape(s); //tworzenie klasy SavedShape która daje się zapisać
            oos.writeObject(ss);
        }
        catch(IOException ex){
            System.err.println(ex.getMessage());
        }
    }

    /**
     * Funkcja zapisujaca kolo do pliku
     * @param s kolo do zapisu
     * @param file plik docelowy
     */
    public void saveShape(Circle s, File file){ // to samo dla każdej klasy Rectangle, Circle i Polygon
        try(ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(file))){
            SavedShape ss = new SavedShape(s);
            oos.writeObject(ss);
        }
        catch(IOException ex){
            System.err.println(ex.getMessage());
        }
    }

    /**
     * Funkcja zapisujaca trojkat do pliku
     * @param s trojkat do zapisu
     * @param file plik docelowy
     */
    public void saveShape(Polygon s, File file){
        try(ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(file))){
            SavedShape ss = new SavedShape(s);
            oos.writeObject(ss);
        }
        catch(IOException ex){
            System.err.println(ex.getMessage());
        }
    }

    /**
     * Funkcja ladujaca ksztalt z pliku
     * @param file plik zrodla
     * @return zaladowany ksztalt
     */
    public Shape loadShape(File file){ //odczyt obiektu SavedShape z pliku i zmiana w obiekt Shape
        try (ObjectInputStream ois = new ObjectInputStream(new FileInputStream(file))) {
            SavedShape s = (SavedShape) ois.readObject();
            
            if(s.shapeType == ShapeMaker.Shapes.RECTANGLE){ //zwracanie figury w zależności od odczytanego typu
                return (Rectangle) s.readShape();
            }
            if(s.shapeType == ShapeMaker.Shapes.CIRCLE){
                return (Circle) s.readShape();
            }
            if(s.shapeType == ShapeMaker.Shapes.TRIANGLE){
                return (Polygon) s.readShape();
            }
        }
        catch(IOException ex){
            System.err.println(ex.getMessage());
        }
        catch(ClassNotFoundException ex){
            System.err.println(ex.getLocalizedMessage());
        }
        return null; // metoda zwraca null w razie źle odczytanego pliku
    }

    /**
     * Konstruktor
     * @param c plotno {@link canvas canvas}
     */
    ShapeHandler(Pane c){ //konstruktor
        this.canvas = c;
        this.currentActiveShape = null;
        this.currentColor = Color.BLACK;
        makeSizable();
    }
}
