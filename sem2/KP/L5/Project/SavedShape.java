import java.io.Serializable;
import javafx.scene.paint.Color;
import javafx.scene.shape.Circle;
import javafx.scene.shape.Polygon;
import javafx.scene.shape.Rectangle;
import javafx.scene.shape.Shape;
/**
 * Klasa rozpisujaca obiekty Shape w celu mozliwosci ich zapisania w pliku
 */
public class SavedShape implements Serializable { //Klasa implementująca Serializable, służąca do rozkodowywania i kodowania obiektów klasy Shape
    private static final long serialVersionUID = 1L;

    /**Rodzaj zapisanego ksztaltu */
    public ShapeMaker.Shapes shapeType = ShapeMaker.Shapes.NONE;

    /**Tablica zawierajaca kolor ksztaltu */
    private double colorHexCode[] = new double[3];
    /**Obrot ksztaltu */
    private double rotation = 0;
    /**Powiekszenie po osi X*/
    private double scaleX = 0;
    /**Powiekszenie po osi Y */
    private double scaleY = 0;

    /**Szerokosc prostokata */
    private double width = 0;
    /**Wysokosc prostokata */
    private double height = 0;

    /**Promien kola */
    private double radius = 0;

    /**Tablica z wspolrzednymi wierzcholkow trojkata */
    private double points[] = {0, 0, 0, 0, 0, 0};
    
    
    /**Konkstruktor dla prostokata 
     * @param s prostokat do zapisania
     */ 
    SavedShape(Rectangle s){ //Rozpisywanie obiektów Shape na parametry
        shapeType = ShapeMaker.Shapes.RECTANGLE;
        Color c = (Color) s.getFill();
        colorHexCode[0] = c.getRed();
        colorHexCode[1] = c.getGreen();
        colorHexCode[2] = c.getBlue();
        scaleX = s.getScaleX();
        scaleY = s.getScaleY();
        rotation = s.getRotate();

        width = s.getWidth();
        height = s.getHeight();
    }
    
    /**Konkstruktor dla kola
     * @param s kolo do zapisania
     */ 
    SavedShape(Circle s){ //dla każdego rodzaju kształtu
        shapeType = ShapeMaker.Shapes.CIRCLE;
        Color c = (Color) s.getFill();
        colorHexCode[0] = c.getRed();
        colorHexCode[1] = c.getGreen();
        colorHexCode[2] = c.getBlue();
        scaleX = s.getScaleX();
        scaleY = s.getScaleY();
        rotation = s.getRotate();

        radius = s.getRadius();
    }

    /**Konkstruktor dla trojkata 
     * @param s trojkat do zapisania
     */ 
    SavedShape(Polygon s){
        shapeType = ShapeMaker.Shapes.TRIANGLE;
        Color c = (Color) s.getFill();
        colorHexCode[0] = c.getRed();
        colorHexCode[1] = c.getGreen();
        colorHexCode[2] = c.getBlue();
        scaleX = s.getScaleX();
        scaleY = s.getScaleY();
        rotation = s.getRotate();

        points[0] = s.getPoints().get(0);
        points[1] = s.getPoints().get(1);
        points[2] = s.getPoints().get(2);
        points[3] = s.getPoints().get(3);
        points[4] = s.getPoints().get(4);
        points[5] = s.getPoints().get(5);
        
    }

    
    
    

    
    /** 
     * Funkcja odczytujaca ksztalt z zapisanych danych 
     * @return odczytany ksztalt
     */
    public Shape readShape(){ //odczytywanie paramterów z powrotem w obiekt Shape
        Shape s = null;
        if(shapeType == ShapeMaker.Shapes.RECTANGLE){
            s = new Rectangle(150, 100, width, height);
        }
        else if(shapeType == ShapeMaker.Shapes.CIRCLE){
            s = new Circle(150, 100, radius);
        }
        else if(shapeType == ShapeMaker.Shapes.TRIANGLE){
            s = new Polygon(points[0], points[1], points[2], points[3], points[4], points[5]);
        }
        
        s.setFill(Color.rgb((int) colorHexCode[0], (int) colorHexCode[1], (int) colorHexCode[2]));
        s.setScaleX(scaleX);
        s.setScaleY(scaleY);
        s.setRotate(rotation);   
        

        return s;
    }

}   
