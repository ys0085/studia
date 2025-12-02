import java.lang.reflect.Array;

public class Queue<T> {
    private final T elements[];
    private int top;

    public Queue(Class<T> c, int capacity){
        @SuppressWarnings("unchecked")
        final T e[] = (T[]) Array.newInstance(c, capacity); 
        elements = e;  
        top = -1;
    }
    public void enqueue(T e) throws QueueFullException {
        if(top == elements.length - 1) throw new QueueFullException();
        elements[++top] = e;
    }
    public T dequeue() throws QueueEmptyException {
        T front = getFront();
        for (int i = 1; i < elements.length; i++) {
            elements[i-1] = elements[i];
        }
        top--;
        return front;
    }
    public T getFront() throws QueueEmptyException {
        if(top == -1) throw new QueueEmptyException();
        return elements[0];
    }
}
