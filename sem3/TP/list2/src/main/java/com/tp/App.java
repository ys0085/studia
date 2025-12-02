package com.tp;

import java.util.Scanner;

/**
 * App class - handling user interface.
 */
final class App {
    private static final String HELP_MESSAGE = (""
            + "==============="
            + "\nCommand list: "
            + "\nadd book"
            + "\nadd item"
            + "\nadd customer"
            + "\n==============="
            + "\nremove book"
            + "\nremove item"
            + "\nremove customer"
            + "\n==============="
            + "\nrent"
            + "\nreturn"
            + "\n==============="
            + "\nprint all"
            + "\nprint stock"
            + "\nprint books"
            + "\nprint customers"
            + "\nprint rented"
            + "\n==============="
            + "\nquit"
            + "\n===============");

    private static final Scanner MAIN_SCANNER = new Scanner(System.in);
    private static final Library MAIN_LIBRARY = new Library();

    private static void add(final CommandSubject subject) {
        switch (subject) {
            case BOOK:
                System.out.println("Book title: ");
                String title = MAIN_SCANNER.nextLine();
                System.out.println("Book author(s): ");
                String authors = MAIN_SCANNER.nextLine();
                System.out.println("Book release year: ");
                int releaseYear;
                try {
                    releaseYear = Integer.parseInt(MAIN_SCANNER.nextLine());
                } catch (NumberFormatException e) {
                    System.out.println("Invalid input");
                    return;
                }
                MAIN_LIBRARY.addBook(title, authors, releaseYear);
                break;

            case ITEM:
                System.out.println("Book ID: ");
                int bookID;
                try {
                    bookID = Integer.parseInt(MAIN_SCANNER.nextLine());
                } catch (NumberFormatException e) {
                    System.out.println("Invalid input");
                    return;
                }
                MAIN_LIBRARY.addItem(bookID);
                break;

            case CUSTOMER:
                System.out.println("First name: ");
                String firstName = MAIN_SCANNER.nextLine();
                System.out.println("Last name: ");
                String lastName = MAIN_SCANNER.nextLine();

                MAIN_LIBRARY.addCustomer(lastName, firstName);
                break;

            default:
                System.out.println("Invalid command.");
                break;
        }
    }

    private static void remove(final CommandSubject subject) {
        switch (subject) {
            case BOOK:
                System.out.println("Book ID: ");
                int bookID;
                try {
                    bookID = Integer.parseInt(MAIN_SCANNER.nextLine());
                } catch (NumberFormatException e) {
                    System.out.println("Invalid input");
                    return;
                }
                MAIN_LIBRARY.removeBook(bookID);
                break;

            case ITEM:
                System.out.println("Item ID: ");
                int itemID;
                try {
                    itemID = Integer.parseInt(MAIN_SCANNER.nextLine());
                } catch (NumberFormatException e) {
                    System.out.println("Invalid input");
                    return;
                }
                MAIN_LIBRARY.removeItem(itemID);
                break;

            case CUSTOMER:
                System.out.println("Customer ID: ");
                int customerID;
                try {
                    customerID = Integer.parseInt(MAIN_SCANNER.nextLine());
                } catch (NumberFormatException e) {
                    System.out.println("Invalid input.");
                    return;
                }
                MAIN_LIBRARY.removeCustomer(customerID);
                break;

            default:
                System.out.println("Invalid command.");
                break;
        }
    }

    private static void rentItem() {
        System.out.println("Customer ID: ");
        int customerID;
        try {
            customerID = Integer.parseInt(MAIN_SCANNER.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input.");
            return;
        }

        System.out.println("Item ID: ");
        int itemID;
        try {
            itemID = Integer.parseInt(MAIN_SCANNER.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input.");
            return;
        }

        MAIN_LIBRARY.rentItem(customerID, itemID);

    }

    private static void returnItem() {

        System.out.println("Item ID: ");
        int itemID;
        try {
            itemID = Integer.parseInt(MAIN_SCANNER.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input.");
            return;
        }

        MAIN_LIBRARY.returnItem(itemID);
    }

    private static void printCustomer() {
        System.out.println("Customer ID: ");
        int customerID;
        try {
            customerID = Integer.parseInt(MAIN_SCANNER.nextLine());
        } catch (NumberFormatException e) {
            System.out.println("Invalid input.");
            return;
        }

        MAIN_LIBRARY.printCustomerRentedItems(customerID);
    }

    /**
     * Main function running the app.
     * @param args
     */
    public static void main(final String[] args) {

        String input = "";
        System.out.println("enter \"help\" for message list");
        MAIN_LIBRARY.fillSampleStock();

        while (!input.equals("exit")) {

            input = MAIN_SCANNER.nextLine().toLowerCase().strip();

            switch (input) {

                case "exit":
                    System.out.println("Exiting the app.");
                    break;

                case "add book":
                    add(CommandSubject.BOOK);
                    break;

                case "add item":
                    add(CommandSubject.ITEM);
                    break;

                case "add customer":
                    add(CommandSubject.CUSTOMER);
                    break;

                case "remove book":
                    remove(CommandSubject.BOOK);
                    break;

                case "remove item":
                    remove(CommandSubject.ITEM);
                    break;

                case "remove customer":
                    remove(CommandSubject.CUSTOMER);
                    break;

                case "rent":
                    rentItem();
                    break;

                case "return":
                    returnItem();
                    break;

                case "print all":
                    MAIN_LIBRARY.printFullStock();
                    break;

                case "print stock":
                    MAIN_LIBRARY.printItemStock();
                    break;

                case "print books":
                    MAIN_LIBRARY.printBookList();
                    break;

                case "print rented":
                    printCustomer();
                    break;

                case "print customers":
                    MAIN_LIBRARY.printCustomerList();
                    break;

                case "help":
                    System.out.println(HELP_MESSAGE);
                    break;

                default:
                    System.out.println("Invalid command.");
                    break;
            }
        }
    }
    private App() {
        //
    }
}
