import java.sql.CallableStatement;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.Scanner;

public class OrdersHandler {
    private static final String DB_URL = "jdbc:mariadb://localhost:3306/app_db?user=root&password=12345678";

    
    public void viewAllOrders(){
        try (Connection connection = DriverManager.getConnection(DB_URL);
             Statement statement = connection.createStatement();
             ResultSet resultSet = statement.executeQuery("SELECT * FROM Orders_full")) {

            System.out.println("\n=== Orders List ===");
            System.out.printf("%-3s %-25s %-9s %-20s %-20s %-4s %-8s%n", "ID", "Date", "Status", "Client", "Product", "Count", "Total");
            System.out.println("---------------------------------------------------------------------------------");

            while (resultSet.next()) {
                System.out.printf("%-3s %-25s %-9s %-20s %-20s %-4s %-8s%n",
                        resultSet.getInt("order_id"),
                        resultSet.getString("order_date"),
                        resultSet.getString("status"),
                        resultSet.getString("client"),
                        resultSet.getString("name"),
                        resultSet.getInt("count"),
                        resultSet.getDouble("price")
                        );
            }
        } catch(SQLException e) {e.printStackTrace();}
    }

    public void addNewOrder(Scanner scanner) {
        System.out.print("Enter client id: ");
        int clientId;
        try {
            clientId = Integer.parseInt(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return;
        }

        System.out.print("Enter product id: ");
        int productId;
        try {
            productId = Integer.parseInt(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return;
        }

        System.out.print("Enter product amount: ");
        int count;
        try {
            count = Integer.parseInt(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return;
        }

        addOrder(clientId, productId, count);
    }
    
    private void addOrder(int clientId, int productId, int count){
        String procedureCall = "{CALL AddOrder(?, ?, ?)}";

        try (Connection connection = DriverManager.getConnection(DB_URL);
             CallableStatement callableStatement = connection.prepareCall(procedureCall)) {

            callableStatement.setInt(1, clientId);
            callableStatement.setInt(2, productId);
            callableStatement.setInt(3, count);

            callableStatement.execute();
            System.out.println("Order added successfully.");
        }
        catch(SQLException e){ System.out.println("Error"); }
    }
    

    public void payForOrder(Scanner scanner) {
        System.out.print("Enter order id: ");
        int orderId;
        try {
            orderId = Integer.parseInt(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return;
        }

        System.out.print("Payment type? (Cash/Debit/Credit)");
        String type;
        
        type = scanner.nextLine();

        if(type.toLowerCase().equals("cash")) type = "Cash";
        if(type.toLowerCase().equals("debit")) type = "Debit";
        if(type.toLowerCase().equals("credit")) type = "Credit";


        payOrder(orderId, type);
    }

    private void payOrder(int orderId, String paymentType){
        String procedureCall = "{CALL PayOrder(?, ?)}";

        try (Connection connection = DriverManager.getConnection(DB_URL);
             CallableStatement callableStatement = connection.prepareCall(procedureCall)) {

            callableStatement.setInt(1, orderId);
            callableStatement.setString(2, paymentType);

            callableStatement.execute();
            System.out.println("Order was paid for successfully.");
        }
        catch(SQLException e){ System.out.println("Error"); }
    }
}
