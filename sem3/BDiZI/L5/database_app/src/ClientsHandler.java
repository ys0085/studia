import java.sql.*;
import java.util.Scanner;

public class ClientsHandler {
    private static final String DB_URL = "jdbc:mariadb://localhost:3306/app_db?user=root&password=12345678";

    public void viewAllClients() {
        try (Connection connection = DriverManager.getConnection(DB_URL);
             Statement statement = connection.createStatement();
             ResultSet resultSet = statement.executeQuery("SELECT * FROM Clients")) {

            System.out.println("\n=== Clients List ===");
            System.out.printf("%-5s %-20s %-20s %-30s%n", "ID", "First Name", "Last Name", "Email");
            System.out.println("----------------------------------------------------------------------");

            while (resultSet.next()) {
                System.out.printf("%-5d %-20s %-20s %-30s%n",
                        resultSet.getInt("id"),
                        resultSet.getString("first_name"),
                        resultSet.getString("last_name"),
                        resultSet.getString("email"));
            }
        } catch(SQLException e){ System.out.println("Error"); }
    }

    
    public void addNewClient(Scanner scanner) {
        System.out.print("Enter first name: ");
        String firstName = scanner.nextLine();

        System.out.print("Enter last name: ");
        String lastName = scanner.nextLine();

        System.out.print("Enter email: ");
        String email = scanner.nextLine();

        try (Connection connection = DriverManager.getConnection(DB_URL);
             PreparedStatement statement = connection.prepareStatement(
                     "INSERT INTO Clients (first_name, last_name, email) VALUES (?, ?, ?)")) {

            statement.setString(1, firstName);
            statement.setString(2, lastName);
            statement.setString(3, email);

            int rowsInserted = statement.executeUpdate();
            if (rowsInserted > 0) {
                System.out.println("Client added successfully!");
            }
        } catch(SQLException e){ System.out.println("Error"); }
    }

    public void deleteClient(Scanner scanner) {
        System.out.print("Enter the ID of the client to delete: ");
        int clientId;
        try {
            clientId = Integer.parseInt(scanner.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input. Please enter a valid number.");
            return;
        }

        try (Connection connection = DriverManager.getConnection(DB_URL);
             PreparedStatement statement = connection.prepareStatement("DELETE FROM Clients WHERE id = ?")) {

            statement.setInt(1, clientId);

            int rowsDeleted = statement.executeUpdate();
            if (rowsDeleted > 0) {
                System.out.println("Client deleted successfully!");
            } else {
                System.out.println("Client with ID " + clientId + " does not exist.");
            }
        } catch(SQLException e){ System.out.println("Error"); }
    }
}
