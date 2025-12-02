import java.util.Scanner;

public class ConsoleApp {
    public static String pwd;
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        ClientsHandler clientsHandler = new ClientsHandler();
        ProductsHandler productsHandler = new ProductsHandler();
        OrdersHandler ordersHandler = new OrdersHandler();
        
        while (true) {
            System.out.println("\n======= System =======");
            System.out.println("1. Manage Clients");
            System.out.println("2. Manage Products");
            System.out.println("3. Manage Orders");
            System.out.println("4. Exit");
            System.out.print("Choose an option: ");

            int mainChoice;
            try {
                mainChoice = Integer.parseInt(scanner.nextLine());
            } catch (NumberFormatException e) {
                System.out.println("Invalid input. Please enter a number.");
                continue;
            }

            switch (mainChoice) {
                case 1 -> manageClients(scanner, clientsHandler);
                case 2 -> manageProducts(scanner, productsHandler);
                case 3 -> manageOrders(scanner, ordersHandler);
                case 4 -> {
                    System.out.println("Exiting the application. Goodbye!");
                    return;
                }
                default -> System.out.println("Invalid option. Please try again.");
            }
        }
    }

    private static void manageClients(Scanner scanner, ClientsHandler clientsHandler) {
        while (true) {
            System.out.println("\n=== Clients Management ===");
            System.out.println("1. View All Clients");
            System.out.println("2. Add New Client");
            System.out.println("3. Delete Client");
            System.out.println("4. Back to Main Menu");
            System.out.print("Choose an option: ");

            int choice;
            try {
                choice = Integer.parseInt(scanner.nextLine());
            } catch (NumberFormatException e) {
                System.out.println("Invalid input. Please enter a number.");
                continue;
            }

            switch (choice) {
                case 1 -> clientsHandler.viewAllClients();
                case 2 -> clientsHandler.addNewClient(scanner);
                case 3 -> clientsHandler.deleteClient(scanner);
                case 4 -> {
                    return;
                }
                default -> System.out.println("Invalid option. Please try again.");
            }
        }
    }

    private static void manageProducts(Scanner scanner, ProductsHandler productsHandler) {
        while (true) {
            System.out.println("\n=== Products Management ===");
            System.out.println("1. View All Products");
            System.out.println("2. Add New Product");
            System.out.println("3. Delete Product");
            System.out.println("4. Back to Main Menu");
            System.out.print("Choose an option: ");

            int choice;
            try {
                choice = Integer.parseInt(scanner.nextLine());
            } catch (NumberFormatException e) {
                System.out.println("Invalid input. Please enter a number.");
                continue;
            }

            switch (choice) {
                case 1 -> productsHandler.viewAllProducts();
                case 2 -> productsHandler.addNewProduct(scanner);
                case 3 -> productsHandler.deleteProduct(scanner);
                case 4 -> {
                    return;
                }
                default -> System.out.println("Invalid option. Please try again.");
            }
        }
    }
    private static void manageOrders(Scanner scanner, OrdersHandler ordersHandler){
        while (true) {
            System.out.println("\n=== Products Management ===");
            System.out.println("1. View All Orders");
            System.out.println("2. Add New Order");
            System.out.println("3. Pay for Order");
            System.out.println("4. Back to Main Menu");
            System.out.print("Choose an option: ");

            int choice;
            try {
                choice = Integer.parseInt(scanner.nextLine());
            } catch (NumberFormatException e) {
                System.out.println("Invalid input. Please enter a number.");
                continue;
            }

            switch (choice) {
                case 1 -> ordersHandler.viewAllOrders();
                case 2 -> ordersHandler.addNewOrder(scanner);
                case 3 -> ordersHandler.payForOrder(scanner);
                case 4 -> {
                    return;
                }
                default -> System.out.println("Invalid option. Please try again.");
            }
        }
    }
}