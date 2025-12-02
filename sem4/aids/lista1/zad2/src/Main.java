import java.util.Random;

public class Main {
    public static void main(String[] args) throws Exception {
        LinkedList list1 = new LinkedList();
        for(int i = 0; i < 10; i++){
            list1.insert(i + 30);
        }
        LinkedList list2 = new LinkedList();
        for(int i = 0; i < 10; i++){
            list2.insert(i + 40);
        }
        System.out.println("List 1:");
        list1.printFull();

        System.out.println("List 2:");
        list2.printFull();

        System.out.println("Merging...");
        LinkedList.merge(list1, list2);

        System.out.println("List 1:");
        list1.printFull();
        
        System.out.println("List 2:");
        list2.printFull();

        // TEST

        System.out.println("Tests");

        Random rand = new Random();
        int table[] = new int[10000];
        LinkedList testList = new LinkedList();
        for(int i = 0; i < 10000; i++){
            table[i] = rand.nextInt(100000);
            testList.insert(table[i]);
        }

        Node n = testList.head;
        int counter1 = 0;
        for(int i = 0; i < 1000; i++){
            int v = table[rand.nextInt(table.length)];
            for (int j = 0; j < testList.count; j++) {
                if(n.value == v) break;
                counter1++;
                n = n.next;
            }
        }
        double avg1 = counter1 / 1000;

        n = testList.head;
        int counter2 = 0;
        for(int i = 0; i < 1000; i++){
            int v = rand.nextInt(100000);
            for (int j = 0; j < testList.count; j++) {
                if(n.value == v) break;
                counter2++;
                n = n.next;
            }
        }
        double avg2 = counter2 / 1000;

        System.out.println("Cost: " + avg1 + " " + avg2);
    }
}
