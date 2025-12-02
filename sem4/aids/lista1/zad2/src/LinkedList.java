
public class LinkedList {
    public Node head = null;
    public Node tail = null;
    public int count;
    LinkedList() {
        head = null;
        count = 0;
    }

    public void insert(int value){
        Node node = new Node(value);

        if(head == null){
            head = node;
        } 
        else{
            tail.next = node;
        }
        tail = node;
        tail.next = head;
        count++;
    }

    public Node search(int v){
        Node n = head;
        for (int i = 0; i < count; i++) {
            if(n.value == v) return n;
            n = n.next;
        }
        return null;
    }

    public void printFull(){
        if(head == null) return;
        Node n = head;
        System.out.print("| ");
        for(int i = 0; i < count; i++){
            n = n.next;
            System.out.print(n.value + " ");
        }
        System.out.println("|");
        
    }

    public static void merge(LinkedList list1, LinkedList list2){
        int total = list1.count + list2.count;
        list1.count = total;
        list2.count = total;
        list1.tail.next = list2.head;
        list2.tail.next = list1.head;
        
    }

    
    
}
