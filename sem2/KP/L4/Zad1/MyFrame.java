package Zad1;
import java.awt.*;
import java.awt.event.*;

class MyWindowAdapter extends WindowAdapter {
    public void windowClosing(WindowEvent e) { System.exit(0); }
}
class RunButton extends Button {
    RunButton(MyFrame f){
        super("Run");
        addActionListener(new RunButtonAL(f));
    }
}
class RunButtonAL implements ActionListener{
    MyFrame f;
    RunButtonAL(MyFrame f) { this.f = f; }
    public void actionPerformed(ActionEvent e) {
        f.run();
        f.repaint();
    }
}


class MyFrame extends Frame{
    Label description;
    Label rows[] = new Label[100];
    TextField inputField;
    RunButton b;
    Container rowBox;
    GridBagConstraints gridC;
    GridBagLayout gridbag;
    MyFrame(){
        super("Trojkat Pascala");
        
        setBounds(new Rectangle(500,500));
        addWindowListener(new MyWindowAdapter());
        setFont(new Font(Font.MONOSPACED, Font.PLAIN, 20));


        for(int i = 0; i < 100; i++){
            rows[i] = new Label("");
        } 
        gridC = new GridBagConstraints();
        gridbag = new GridBagLayout();
        setLayout(gridbag);
        description = new Label("Wpisz liczbe:");
        inputField = new TextField();
        b = new RunButton(this);
        rowBox = new Container();
        

        gridC.gridwidth = GridBagConstraints.REMAINDER;
        gridbag.setConstraints(description, gridC);
        add(description);

        gridC.gridwidth = 1;
        gridC.weighty = 1.0;
        gridC.gridheight = 2;

        gridbag.setConstraints(inputField, gridC);
        add(inputField);


        
        gridbag.setConstraints(b, gridC);
        add(b);

        gridC.gridwidth = GridBagConstraints.REMAINDER;
        gridC.gridheight = 1;
        gridC.weighty = 1.0;
        gridbag.setConstraints(rowBox, gridC);
        add(rowBox);
    }
    void run(){
        try {
            int input = Integer.parseInt(inputField.getText());
            this.rowBox.setLayout(new GridLayout(input+1, 1));
            rowBox.removeAll();
            for(int i = 0; i<input+1; i++){
                rows[i].setText("");
                rows[i].setAlignment(Label.CENTER);
                WierszTrojkataPascala wtp = new WierszTrojkataPascala(i);
                for(int j = 0; j < i+1; j++){
                    rows[i].setText(rows[i].getText() + " " + Integer.toString(wtp.wartosc(j)));
                }
                this.rowBox.add(rows[i]);
            }
            
        } 
        catch (InvalidRowNumberException e) {
            System.out.println("podaj liczbe dodatnia");
        } 
        catch (IndexOutOfRangeException e) {
            System.out.println("??");
        }
        catch (NumberFormatException e) {
            System.out.println("nieprawidlowa liczba");
        }
        this.setVisible(false);
        this.setVisible(true);

    }
}
