import java.sql.*;
import java.util.Scanner;

public class ProductsHandler {
    private static final String DB_URL = "jdbc:mariadb://localhost:3306/app_db?user=root&password=12345678";

    
    public void viewAllProducts() {
        try (Connection connection = DriverManager.getConnection(DB_URL);
             Statement statement = connection.createStatement();
             ResultSet resultSet = statement.executeQuery("SELECT * FROM Products")) {

            System.out.println("\n=== Products List ===");
            System.out.printf("%-5s %-20s %-10s %-10s%n", "ID", "Name", "Price", "Stock");
            System.out.println("---------------------------------------------");

            while (resultSet.next()) {
                System.out.printf("%-5d %-20s %-10.2f %-10d%n",
                        resultSet.getInt("id"),
                        resultSet.getString("name"),
                        resultSet.getBigDecimal("price"),
                        resultSet.getInt("stock"));
            }
        } catch(SQLException e){ System.out.println("Error"); }
    }

    
    public void addNewProduct(Scanner scanner) {
        System.out.print("Enter product name: ");
        String name = scanner.nextLine();

        System.out.print("Enter product price: ");
        double price;
        try {
            price = Double.parseDouble(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid price. Please try again.");
            return;
        }

        System.out.print("Enter product stock: ");
        int stock;
        try {
            stock = Integer.parseInt(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid stock. Please try again.");
            return;
        }

        try (Connection connection = DriverManager.getConnection(DB_URL);
             PreparedStatement statement = connection.prepareStatement(
                     "INSERT INTO Products (name, price, stock) VALUES (?, ?, ?)")) {

            statement.setString(1, name);
            statement.setDouble(2, price);
            statement.setInt(3, stock);

            int rowsInserted = statement.executeUpdate();
            if (rowsInserted > 0) {
                System.out.println("Product added successfully!");
            }
        } catch(SQLException e){ System.out.println("Error"); }
    }

    
    public void deleteProduct(Scanner scanner) {
        System.out.print("Enter the ID of the product to delete: ");
        int productId;
        try {
            productId = Integer.parseInt(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return;
        }

        try (Connection connection = DriverManager.getConnection(DB_URL);
             PreparedStatement statement = connection.prepareStatement("DELETE FROM Products WHERE id = ?")) {

            statement.setInt(1, productId);

            int rowsDeleted = statement.executeUpdate();
            if (rowsDeleted > 0) {
                System.out.println("Product deleted successfully!");
            } else {
                System.out.println("Product with ID " + productId + " does not exist.");
            }
        }
        catch(SQLException e){ System.out.println("Error"); }
    }
}
