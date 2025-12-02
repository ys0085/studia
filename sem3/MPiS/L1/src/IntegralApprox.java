
import org.apache.commons.math3.random.MersenneTwister;





public class IntegralApprox {

    public Func f;
    private double max_fx;
    private double a;
    private double b;
    private MersenneTwister gen;
    private double approx_max_value(){
        double m_fx = f.calculate(a);
        double delta = 0.05;
    
        for(double x = a; x <= b; x += delta){
            double fx = f.calculate(x);
            if(fx > m_fx) m_fx = fx;
        }
    
    
        return m_fx;
    
    }
    
    public double get_random(double min, double max){
        double r = gen.nextDouble();
        
        //Applying correct constraints on r

        double diff = max - min;
        r = r * diff; // <- r is now in [0, diff]
        r = r + min; // <- r is now in [min, max]
    
        return r;
    }
    
    
    public double approximate_integral(int n){
        
        double C = 0;
        for(int i = 0; i < n; i++){
            double x = get_random(a, b);
            double y = get_random(0, max_fx);
            if(y <= f.calculate(x)) C++;
        }
    
        return (double) (b - a) * max_fx * (double) C / n;
    }
    
    public void runTests(int k){
        double results[][] = new double[k][100];

        for(int i = 0; i < k; i++){
            for(int j = 0; j < 100; j++){
                results[i][j] = approximate_integral((j + 1) * 50);
            }
        }
    
    
        for(int i = 0; i < 100; i++){
            for(int j = 0; j < k; j++){
                System.out.print(results[j][i] + " ");
            }
            System.out.print("\n");
            
        }
        
    }
    IntegralApprox(double val_a, double val_b, Func fn){
        this(val_a, val_b, 0, fn);
    }
    IntegralApprox(double val_a, double val_b, int seed, Func fn){
        gen = new MersenneTwister(seed);
        this.f = fn;
        this.a = val_a;
        this.b = val_b;
        this.max_fx = approx_max_value();
    }
    
}
