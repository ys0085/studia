import javafx.geometry.Point2D;
import javafx.scene.control.Label;
import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;
import javafx.scene.shape.Circle;
import javafx.scene.shape.Polygon;
import javafx.scene.shape.Rectangle;
import javafx.scene.shape.Shape;

/**
 * Klasa tworzaca ksztalty
 */
public class ShapeMaker {
     /**{@link ShapeHandler ShapeHandler} odpowiedzialny za zarzadzanie dodawanymi ksztaltami*/
    private ShapeHandler sh;
    /**Napis z aktualnie wybranym ksztaltem wyswietlany na plotnie */
    private Label cs;

    /**
     * Klasa enum z nazwami ksztaltow
     */
    public enum Shapes {
        /**Prostokat */
        RECTANGLE,
        /**Kolo */
        CIRCLE,
        /**Trojkat */
        TRIANGLE,
        /**Brak ksztaltu */
        NONE
    }

    /** Tablica z napisami wyswielanymi na plotnie */
    private String[] napisy =  {
        "Rysujesz prostokąt",
        "Rysujesz koło",
        "Rysujesz trójkąt",
        ""
    };

    
    /**Aktualnie wybrany rodzaj ksztaltu */
    private Shapes currentShape;
    
    


    private Point2D firstClick;
    private Circle firstClickIndicator;
    private Point2D secondClick;
    private Circle secondClickIndicator;
    /** 
     * Funkcja dodajaca wybrany ksztalt {@link currentShape currentShape} do plotna
     * @param canvas plotno
     */
    public void addShapes(Pane canvas){
        firstClick = null;
        secondClick = null;
        
        canvas.getChildren().remove(firstClickIndicator);
        canvas.getChildren().remove(secondClickIndicator);
        
        canvas.setOnMousePressed(e -> {
            Shape s;
            if(currentShape == Shapes.RECTANGLE){
                if(firstClick == null){     //Zapisuję pozycję pierwszego punktu
                    canvas.getChildren().remove(firstClickIndicator);
                    firstClick = new Point2D(e.getX(), e.getY());
                    firstClickIndicator = new Circle(firstClick.getX(), firstClick.getY(), 3);
                    firstClickIndicator.setFill(Color.RED);
                    canvas.getChildren().add(firstClickIndicator);
                }
                else{  
                    s = new Rectangle(Math.min(e.getX(), firstClick.getX()), Math.min(e.getY(), firstClick.getY()),Math.abs(firstClick.getX() - e.getX()), Math.abs(firstClick.getY() - e.getY()));
                    firstClick = null; //tworzę kształt i usuwam pierwszy punkt
                    setShape(Shapes.NONE);
                    canvas.getChildren().remove(firstClickIndicator);
                    sh.init(s); //dodawanie funkcjonalności do nowego kształtu
                }
            }
            else if(currentShape == Shapes.CIRCLE){ //analogicznie jak dla kwadratu
                if(firstClick == null){
                    canvas.getChildren().remove(firstClickIndicator);
                    firstClick = new Point2D(e.getX(), e.getY());
                    firstClickIndicator = new Circle(firstClick.getX(), firstClick.getY(), 3);
                    firstClickIndicator.setFill(Color.RED);
                    canvas.getChildren().add(firstClickIndicator);
                }
                else{
                    s = new Circle(firstClick.getX(), firstClick.getY(), Math.hypot(firstClick.getX() - e.getX(), firstClick.getY() - e.getY()));
                    firstClick = null;
                    setShape(Shapes.NONE);
                    canvas.getChildren().remove(firstClickIndicator);
                    sh.init(s);
                    
                }
            }
            else if(currentShape == Shapes.TRIANGLE){ //analogicznie, lecz z dwoma punktami
                if(firstClick == null){
                    canvas.getChildren().remove(firstClickIndicator);
                    firstClick = new Point2D(e.getX(), e.getY());
                    firstClickIndicator = new Circle(firstClick.getX(), firstClick.getY(), 3);
                    firstClickIndicator.setFill(Color.RED);
                    canvas.getChildren().add(firstClickIndicator);
                }
                else if(secondClick == null){
                    canvas.getChildren().remove(secondClickIndicator);
                    secondClick = new Point2D(e.getX(), e.getY());
                    secondClickIndicator = new Circle(secondClick.getX(), secondClick.getY(), 3);
                    secondClickIndicator.setFill(Color.BLUE);
                    canvas.getChildren().add(secondClickIndicator);
                }
                else{
                    s = new Polygon(firstClick.getX(), firstClick.getY(), secondClick.getX(), secondClick.getY(), e.getX(), e.getY());
                    firstClick = null;
                    setShape(Shapes.NONE);
                    canvas.getChildren().remove(firstClickIndicator);
                    canvas.getChildren().remove(secondClickIndicator);
                    sh.init(s);
                }
            }
            else if(currentShape == Shapes.NONE){
                canvas.getChildren().remove(firstClickIndicator);
            }
        });
        
    }

    
    /** Funkcja ustawiajaca podany ksztalt na aktualnie wybrany oraz aktualizujaca napis na plotnie
     * @param s rodzaj ksztaltu
     */
    public void setShape(Shapes s){
        currentShape = s;
        cs.setText(napisy[s.ordinal()]);
    }

    /**
     * Funkcja zwracajaca aktualnie wybrany ksztalt 
     * @return aktualnie wybrany ksztalt
     */
    public Shapes getCurrentShape(){
        return currentShape;
    }

    /**
     * Konstruktor 
     * @param sh {@link ShapeHandler ShapeHandler} odpowiedzialny za zarzadzanie dodawanymi ksztaltami
     * @param cs Napis z aktualnie wybranym ksztaltem wyswietlany na plotnie
     */
    ShapeMaker(ShapeHandler sh, Label cs){
        firstClick = null;
        secondClick = null;
        currentShape = Shapes.NONE;
        this.sh = sh;
        this.cs = cs;
        
    }
}
