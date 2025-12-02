import java.lang.reflect.Array;


public class Stack<T> {
    private final T elements[];
    private int top;
    
    public Stack(Class<T> c, int capacity){
        @SuppressWarnings("unchecked")
        final T e[] = (T[]) Array.newInstance(c, capacity); 
        elements = e;  
        top = -1;
    }
    public void push(T e) throws StackFullException {
        if(top == elements.length - 1) throw new StackFullException();
        elements[++top] = e;
    }

    public T pop() throws StackEmptyException {
        if(top == -1) throw new StackEmptyException();
        return elements[top--];
    }
    public T top() throws StackEmptyException {
        if(top == -1) throw new StackEmptyException();
        else return elements[top];
    }
}
