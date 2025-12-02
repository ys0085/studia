package Zad3;
import java.awt.*;
import java.awt.event.*;


import javax.swing.BoxLayout;
import javax.swing.JLabel;
import javax.swing.JPanel;

import javax.swing.JScrollPane;


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
    JLabel description;
    TextField inputField;
    RunButton b;
    
    GridLayout all;
    Container top;
    JPanel bottom;
    JLabel output[] = new JLabel[1000];
    MyFrame(){
        super("Trojkat Pascala");
        setBounds(new Rectangle(500,500));
        addWindowListener(new MyWindowAdapter());
        setFont(new Font(Font.MONOSPACED, Font.PLAIN, 20));
        all = new GridLayout(2,1);
        setLayout(all);
        description = new JLabel("Wpisz liczbe:");
        inputField = new TextField();
        b = new RunButton(this);

        top = new Container();
        top.setLayout(new GridLayout(3,1));
        top.add(description);
        top.add(inputField);
        top.add(b);

        bottom = new JPanel();
        JScrollPane sp = new JScrollPane(bottom);
        
        
        bottom.setLayout(new BoxLayout(bottom, BoxLayout.PAGE_AXIS));
        // sp.setLayout(new ScrollPaneLayout());
        // bottom.add(sp);
        

        add(top);
        add(sp);
    }
    void run(){
        try {
            MyProcess process = new MyProcess();
            process.start(inputField.getText());
            bottom.removeAll();
            String line = "";
            int i = 0;
            while((line = process.getOutputLine(i)) != ""){
                output[i] = new JLabel(line);
                bottom.add(output[i]);
                i++;
            }
            
            
        } 
        catch (Exception e) {
            System.out.println(e.getMessage());
        } 
        
        this.setVisible(false);
        this.setVisible(true);

    }
}
