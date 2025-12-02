
public class BiLinkedList {
    public Node head = null;
    public Node tail = null;
    public int count;
    BiLinkedList() {
        count = 0;
    }

    public void insert(int value){
        Node node = new Node(value);

        if(head == null){
            head = node;
            tail = node;
        } 
        else{
            tail.next = node;
        }
        node.prev = tail;
        node.next = head;
        tail = node;
        head.prev = tail;
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
            n = n.prev;
            System.out.print(n.value + " ");
        }
        System.out.println("|");
        
    }

    public static void merge(BiLinkedList list1, BiLinkedList list2){
        int total = list1.count + list2.count;
        list1.count = total;
        list2.count = total;

        list1.tail.next = list2.head;
        list2.head.prev = list1.tail;

        list2.tail.next = list1.head;
        list1.head.prev = list2.tail;
        
    }

    
    
}
