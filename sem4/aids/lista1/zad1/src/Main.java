public class Main {
    public static void main(String[] args) throws Exception {
        Stack<Integer> stack = new Stack<>(Integer.class, 50);
        for(int i = 1; i < 53; i++){
            try {
                stack.push(i);
                System.out.println("Push: " + i);
            } catch (StackFullException e) {
                System.out.println("Stack full! at " + i);
            }
            
        }
        for(int i = 1; i < 53; i++){
            try {
                System.out.println("Pop: " + stack.pop());
            } catch (StackEmptyException e) {
                System.out.println("Stack empty!");
            }
        }

        Queue<Integer> queue = new Queue<>(Integer.class, 50);
        for(int i = 1; i < 53; i++){
            try {
                queue.enqueue(i);
                System.out.println("Queued: " + i);
            } catch (QueueFullException e) {
                System.out.println("Queue full! at " + i);
            }
            
        }
        for(int i = 1; i < 53; i++){
            try {
                System.out.println("Pop: " + queue.dequeue());
            } catch (QueueEmptyException e) {
                System.out.println("Queue empty!");
            }
        }
    }
}
