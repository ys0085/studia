package Zad3;
import java.io.BufferedReader;
import java.io.InputStreamReader;

public class MyProcess {
    public String output[] = new String[1000];
    public void start(String args){
        try {
            Process p = Runtime.getRuntime().exec("Zad3/a.exe " + args);
            BufferedReader reader = new BufferedReader(new InputStreamReader(p.getInputStream()));
            String line = "";
            int i = 0;
            while((line = reader.readLine()) != null) {
                output[i] = line;
                i++;
            }
            output[i] = "";
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
    public String getOutputLine(int line){
        return output[line];
    } 
}
