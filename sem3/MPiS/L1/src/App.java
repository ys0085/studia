import java.util.Random;
/*
 * Main class 
 * Run jar with:
 * "java -jar L1.jar arg1 arg2 > out.txt"
 * arg1 - integer in [0, 3] indicating the function to approximate (see Func f[])
 * arg2 - number of tests to run 
 * out.txt - output file 
 * 
 * This project uses the MersenneTwister class from the Apache Commons math3 library, 
 * but the compiled jar should work on its own.
 */
public class App {

    public static void main(String args[]){
        Random rand = new Random();
        int i = Integer.parseInt(args[0]);
        int k = Integer.parseInt(args[1]);
        Func f[] =  {
            (double x) -> Math.pow(x, (double) 1/3),
            (double x) -> Math.sin(x),
            (double x) -> 4 * x * Math.pow(1 - x, 3),
            (double x) -> 2 * Math.sqrt(1 - x * x)
        };

        double a[] = {
            0,
            0,
            0,
            -1
        };

        double b[] = {
            8,
            (Math.PI),
            1,
            1
        };

        IntegralApprox iApp = new IntegralApprox(a[i], b[i], rand.nextInt(), f[i]);

        iApp.runTests(k);
    }
}   
