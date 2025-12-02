import javafx.scene.control.Label;
import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;
import javafx.scene.shape.Rectangle;


/**
 * Klasa - główny wątek uzywany w programie
 */
public class ColorField extends Thread {
    private Pane pane;
    private Rectangle rect; 
    private Color color;
    private int index;

    /**
     * funkcja zwracająca aktualny kolor pola, lub kolor #000000 jesli wątek jest wstrzymany
     * @return kolor pola lub kolor #000000
     */
    public Color getColor(){
        return isSuspended() ? 
        new Color(0.0, 0.0, 0.0, 0.0) : // zwracam kolor #000000 aby nie ingerowac w funkcje wybierania średniej z kolorów
        color;
    }

    /**
     * funkcja zmiany koloru na inny
     * @param newColor nowy kolor
     */
    public void setColor(Color newColor){
        color = newColor;
        rect.setFill(color);;
    }
    
    /**
     * funkcja zwracająca widoczny element wątku
     * @return Pane wątku
     */
    public Pane getPane(){
        return pane;
    }
    
    private ColorField f1;
    private ColorField f2;
    private ColorField f3;
    private ColorField f4;

    /**
     * funkcja ustawiania sąsiadów danego wątku
     * @param f1 sasiad 1
     * @param f2 sasiad 2
     * @param f3 sasiad 3
     * @param f4 sasiad 4
     */
    public void setNeighbors(ColorField f1, ColorField f2, ColorField f3, ColorField f4){
        this.f1 = f1;
        this.f2 = f2;
        this.f3 = f3;
        this.f4 = f4;
    }

    synchronized private void setRandomColor(){
        setColor(new Color(Program.RAND.nextDouble(), Program.RAND.nextDouble(), Program.RAND.nextDouble(), 1.0));
    }
    synchronized private void setAverageColor(){
        int v = 4;
        if(f1.threadSuspended) v--;
        if(f2.threadSuspended) v--;
        if(f3.threadSuspended) v--;
        if(f4.threadSuspended) v--;
        if(v == 0) return;
        double newRed = f1.getColor().getRed() + f2.getColor().getRed() + f3.getColor().getRed() + f4.getColor().getRed();
        double newGreen = f1.getColor().getGreen() + f2.getColor().getGreen() + f3.getColor().getGreen() + f4.getColor().getGreen();
        double newBlue = f1.getColor().getBlue() + f2.getColor().getBlue() + f3.getColor().getBlue() + f4.getColor().getBlue();
        
        Color newColor = new Color(newRed / v, newGreen / v, newBlue / v, 1.0);
        setColor(newColor);
    }
    
    private boolean threadSuspended;
    private boolean threadRunning;

    /**
     * funkcja zwracająca czy wątek jest wstrzymany
     * @return stan wstrzymania wątku
     */
    public boolean isSuspended() {
        return threadSuspended;
    }

    /**
     * funkcja zwracająca czy wątek jest rozpoczęty
     * @return stan wątku
     */
    public boolean isRunning(){
        return threadRunning;
    }
    

    private void suspendT(){
        threadSuspended = true;
    }

    private void resumeT(){
        threadSuspended = false;
        synchronized(this){
            notify();
        }
    }

    private void toggleSuspend(){
        if(threadSuspended) resumeT();
        else suspendT();
    }


    /**
     * funkcja zatrzymująca wątek
     */
    public void stopT(){
        threadRunning = false;
    }

    /**
     * funkcja rozpoczynająca wątek
     */
    synchronized public void run(){
        threadRunning = true;
        threadSuspended = false;
        //System.out.println("Start: " + getName());
        try{
            sleep(500);
            while(threadRunning){
                synchronized(this){
                    
                
                    sleep((long) ((Program.RAND.nextDouble() + 0.5) * Program.TIME));
            
                    while (threadSuspended)
                        wait();
                

                
                    System.out.println("Start: " + getName());
                    
                    if(Program.RAND.nextDouble() < Program.PROB){
                        setRandomColor();
                    }
                    else{
                        setAverageColor();
                    }
                    System.out.println("End: " + getName());

                    
                }
                
            }
            
        }
        catch(InterruptedException e){

        }
        //System.out.println("End: " + getName());
    }

    private void suspendOnMouseClick(){
        pane.setOnMouseClicked(e -> {
            toggleSuspend();
        });
    }
    
    /**
     * Konstruktor wątku 
     * @param index liczba porządkowa wątku
     */
    ColorField(int index){
        this.index = index;
        pane = new Pane();
        pane.setMaxWidth(Program.FIELD_SIZE);
        pane.setMinWidth(Program.FIELD_SIZE);
        pane.setMaxHeight(Program.FIELD_SIZE);
        pane.setMinHeight(Program.FIELD_SIZE);
        rect = new Rectangle(Program.FIELD_SIZE, Program.FIELD_SIZE);
        pane.getChildren().add(rect);
        setRandomColor();
        setName("Field " + this.index);
        //pane.getChildren().add(new Label(index + ""));
        suspendOnMouseClick();
        setPriority(MIN_PRIORITY);
        threadRunning = false;
        threadSuspended = false;
    }
}
