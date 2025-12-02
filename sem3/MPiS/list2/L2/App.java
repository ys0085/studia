package com.mpis;
import java.util.Random;

public class App {
    public static void main (String args[]) {
        Random rand = new Random();
        Balls balls = new Balls(rand.nextInt());
        for (int k = 0; k < 50; k++){
            for (int i = 1; i <= 100; i++) {
                balls.test(1000*i);
                System.out.println(balls.B + " " + balls.U + " " + balls.C + " " + balls.D);
            }   
        }
    }

}
