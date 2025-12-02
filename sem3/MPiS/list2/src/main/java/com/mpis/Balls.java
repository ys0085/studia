package com.mpis;

import org.apache.commons.math3.random.MersenneTwister;

public class Balls {
    final private MersenneTwister gen;

    private int bins[];
    private int bin_count;

    private int ball_counter = 0;

    public int B;
    public int U;
    public int C;
    public int D;

    private void checkB(){
        if (B != 0) return;
        for (int i = 0; i < bin_count; i++) {
            if (bins[i] > 1) {
                B = ball_counter;
                return;
            }
        }
    }

    private void checkU(){
        if (ball_counter != bin_count) return;
        for (int i = 0; i < bin_count; i++) {
            if (bins[i] == 0) U++;
        }
    }

    private void checkC(){
        if (C != 0) return;
        for (int i = 0; i < bin_count; i++) {
            if (bins[i] == 0) return;
        } 
        C = ball_counter;
    }

    private void checkD(){
        if (D != 0) return;
        for (int i = 0; i < bin_count; i++) {
            if (bins[i] < 2) return;
        } 
        D = ball_counter;
    }

    public void test(int n){
        bin_count = n;
        B = 0;
        U = 0;
        C = 0;
        D = 0;
        bins = new int[bin_count];
        for (int i = 0; i < bin_count; i++) {
            bins[i] = 0;
        }
        ball_counter = 0;
        while (B == 0 || C == 0 || U == 0 || D == 0) {
            bins[gen.nextInt(bin_count)]++;
            checkB();
            checkU();
            checkC();
            checkD();
            ball_counter++;
        }
    }

    Balls(int seed) {
        gen = new MersenneTwister(seed);
        B = 0;
        U = 0;
        C = 0;
        D = 0;
    }

    

    

}
