public class Program{
    public static void main(String[] args){
        int len = args.length;
        int[] n = new int[len];
        int[] d = new int[len];
        for(int i = 0; i < len; i++){
            try{
                n[i] = Integer.parseInt(args[i]);
                d[i] = div(Math.abs(n[i]));
            	System.out.println(n[i] + " " + d[i]);
            }
            catch (NumberFormatException ex){
                System.out.println(args[i] + " nie jest liczba naturalna");
            }
            
            
        }
    }
    public static int div(int n){ //zakladam ze najwiekszy dzielnik wykluczjac liczbe n
        int i = n-1;
        for(; i > 0; i--){
            if(n % i == 0) break;
        }
        return n != 1 ? i : 1;
    }
}
