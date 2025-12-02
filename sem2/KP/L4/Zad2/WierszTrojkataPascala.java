
class InvalidRowNumberException extends Exception{
}
class IndexOutOfRangeException extends Exception{
}
class WierszTrojkataPascala {
    private int n;
    private int[] val;
    public void generate(){
        val = new int[n+1];
        val[0] = 1;
        int r = 1;
        while(r <= this.n){
            int old_val[] = val.clone();
            val[r] = 1;
            for(int i = 1; i <= r-1; i++){
                val[i] = old_val[i-1] + old_val[i];
            }
            r++;
        }
    }
    public int wartosc(int m) throws IndexOutOfRangeException{
        if(m > this.n || m < 0) throw new IndexOutOfRangeException();
        else return this.val[m];
    }
    WierszTrojkataPascala(int n) throws InvalidRowNumberException{ 
        if(n < 0){
            throw new InvalidRowNumberException();
        }
        this.n = n;
        generate();
    }
}
