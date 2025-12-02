package com.tp;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

/**
 * Klasa "zlepiająca wszystko w całość". Pełni rolę "Eksperta" w metodach GRASP,
 * oraz rolę "Kreatora" dla klas Customer, Item, Book.
 */
public class Library {
    private boolean printActionFeedback = true;
    private boolean printErrorFeedback = true;

    private final int bookStartingID = 1;
    private final int itemStartingID = 1;
    private final int customerStartingID = 1;

    private final List<Item> itemStock = new ArrayList<>();
    private final List<Book> bookList = new ArrayList<>();
    private final List<Customer> customerList = new ArrayList<>();

    private final Comparator<Book> bookIDComparator =
    (Book b1, Book b2) -> {
        if (b1.getBookID() == b2.getBookID()) {
            return 0;
        } else if (b1.getBookID() > b2.getBookID()) {
            return 1;
        } else {
            return -1;
        }
    };
    private final Comparator<Item> itemIDComparator =
    (Item i1, Item i2) -> {
        if (i1.getItemID() == i2.getItemID()) {
            return 0;
        } else if (i1.getItemID() > i2.getItemID()) {
            return 1;
        } else {
            return -1;
        }
    };
    private final Comparator<Customer> customerIDComparator =
    (Customer c1, Customer c2) -> {
        if (c1.getID() == c2.getID()) {
            return 0;
        } else if (c1.getID() > c2.getID()) {
            return 1;
        } else {
            return -1;
        }
    };

    /**
     * Defines if the program should print out feedback on successful action.
     * @param b true / false
     */
    public final void doPrintActionFeedback(final boolean b) {
        printActionFeedback = b;
    }

    /**
     * Fills the library with sample stock.
     */
    public final void fillSampleStock() {
        doPrintActionFeedback(false);
        addBook("Lalka", "Bolesław Prus", 1889);
        addBook("Lalka 2", "Kolesław Prus", 1999);
        addBook("Lalka 3", "Bolesław Obrus", 1909);

        addItem(1);
        addItem(1);
        addItem(1);

        addItem(2);
        addItem(2);
        addItem(2);

        addItem(3);
        addItem(3);
        addItem(3);

        doPrintActionFeedback(true);
    }

    /**
     * Getter for the library's item stock.
     * @return item stock
     */
    public final List<Item> getItemStock() {
        return itemStock;
    }

    /**
     * Getter of the library's book list.
     * @return book list
     */
    public final List<Book> getBookList() {
        return bookList;
    }


    private Book findBook(final int bookID) {
        Book book = null;
        for (Book b : bookList) {
            if (b.getBookID() == bookID) {
                book = b;
            }
        }
        return book;
    }

    private Item findItem(final int itemID) {
        Item item = null;
        for (Item i : itemStock) {
            if (i.getItemID() == itemID) {
                item = i;
            }
        }
        return item;
    }

    private Customer findCustomer(final int customerID) {
        Customer customer = null;
        for (Customer c : customerList) {
            if (c.getID() == customerID) {
                customer = c;
            }
        }
        return customer;
    }

    /**
     * Adds a book position to the library.
     * @param title Title
     * @param author Author
     * @param releaseYear Release year
     */
    public final void addBook(final String title,
                            final String author,
                            final int releaseYear) {
        for (Book b : bookList) {
            if (b.getTitle().equals(title)
                    && b.getAuthor().equals(author)
                    && b.getReleaseYear() == releaseYear) {
                if (printErrorFeedback) {
                    System.out.println("This book is already added.");
                }
                return;
            }
        }

        int freeID = bookStartingID;
        for (Book b : bookList) {
            if (b.getBookID() == freeID) {
                freeID++;
            }
        }
        bookList.add(new Book(title, author, releaseYear, freeID));
        bookList.sort(bookIDComparator);

        if (printActionFeedback) {
            System.out.println("Book \"" + title
                             + "\" added with Book ID " + freeID + ".");
        }

    }

    /**
     * Adds a copy of an existing book position.
     * @param bookID Book ID
     */
    public final void addItem(final int bookID) {
        Book book = findBook(bookID);
        if (book == null) {
            if (printErrorFeedback) {
                System.out.println("Wrong book ID!");
            }
            return;
        }
        int freeID = itemStartingID;
        for (Item i : itemStock) {
            if (i.getItemID() == freeID) {
                freeID++;
            }
        }
        itemStock.add(new Item(book, freeID));
        itemStock.sort(itemIDComparator);

        if (printActionFeedback) {
            System.out.println("Added a copy of \"" + book.getTitle()
                                + "\" with Item ID " + freeID + ".");
        }
    }

    /**
     * Adds a customer to the library.
     * @param lN Last name
     * @param fN First name
     */
    public final void addCustomer(final String lN, final String fN) {
        int freeID = customerStartingID;
        for (Customer c : customerList) {
            if (c.getID() == freeID) {
                freeID++;
            }
        }

        Customer customer = new Customer(lN, fN, freeID);
        customerList.add(customer);
        customerList.sort(customerIDComparator);
        if (printActionFeedback) {
            System.out.println("Customer " + customer.getFullName()
                            + " added with Customer ID " + freeID + ".");
        }
    }

    /**
     * Removes a book from the library along with all of its copies.
     * @param bookID Book ID
     */
    public final void removeBook(final int bookID) {
        Book removedBook = findBook(bookID);
        if (removedBook == null) {
            if (printErrorFeedback) {
                System.out.println("No book of such ID exists.");
            }
            return;
        }

        List<Item> removedItems = new ArrayList<>();
        for (Item i : itemStock) {
            if (i.getBookID() == bookID) {
                removedItems.add(i);
            }
        }

        int count = removedItems.size();

        for (Item i : removedItems) {
            itemStock.remove(i);
        }

        String title = removedBook.getTitle();
        bookList.remove(removedBook);

        if (printActionFeedback) {
            System.out.println("Removed "
                    + (count == 1 ? "a copy " : (count + " copies "))
                    + "of \"" + title + "\".");
        }

    }

    /**
     * Removes a copy of a book from the library.
     * @param itemID Item ID
     */
    public final void removeItem(final int itemID) {
        Item removedItem = findItem(itemID);
        if (removedItem == null) {
            if (printErrorFeedback) {
                System.out.println("No item of such ID exists.");
            }
            return;
        }

        String title = removedItem.getTitle();
        itemStock.remove(removedItem);

        if (printActionFeedback) {
            System.out.println("Removed a copy of \"" + title + "\".");
        }
    }

    /**
     * Removes a customer from the library and returnes all their rented books.
     * @param customerID Customer ID
     */
    public final void removeCustomer(final int customerID) {
        Customer removedCustomer = findCustomer(customerID);
        if (removedCustomer == null) {
            if (printErrorFeedback) {
                System.out.println("No customer of such ID exists.");
            }
            return;
        }

        List<Item> returnedItems = new ArrayList<>();

        for (Item i : removedCustomer.getRentedItemList()) {
            returnedItems.add(i);
        }

        for (Item i : returnedItems) {
            returnItem(i.getItemID());
        }

        String fullName = removedCustomer.getFullName();

        customerList.remove(removedCustomer);

        if (printActionFeedback) {
            System.out.println("Removed customer " + fullName + ".");
        }

    }

    /**
     * Prints the full stock of the library.
     */
    public final void printFullStock() {
        printItemStock();
        System.out.println("");
        printBookList();
        System.out.println("");
        printCustomerList();
    }

    /**
     * Prints all book copies in the library.
     */
    public final void printItemStock() {
        System.out.println("================== Book stock ==================");
        System.out.println("IID | Availability | CID | BID | Year | Position");
        System.out.println("================================================");
        for (Item i : itemStock) {
            int customerID = i.getCustomerID();
            System.out.println(String.format("%03d", i.getItemID())
                    + " | " + (i.isRented() ? "Rented      " : "Available   ")
                    + " | " + (customerID == -1
                                ? "---"
                                : String.format("%03d", customerID))
                    + " | " + String.format("%03d", i.getBookID())
                    + " | " + String.format("%04d", i.getReleaseYear())
                    + " | " + i.getAuthor() + " - " + i.getTitle());
        }
        System.out.println("================================================");
    }

    /**
     * Prints all the book positions in the library.
     */
    public final void printBookList() {
        System.out.println("=========== Position list ============");
        System.out.println("BID | Total | Available | Year | Title");
        System.out.println("======================================");
        for (Book b : bookList) {
            int totalCount = countTotalCopies(b);
            int rentedCount = countRentedCopies(b);
            System.out.println(String.format("%03d", b.getBookID())
                    + " | " + String.format("%-5d", totalCount)
                    + " | " + String.format("%-9d", totalCount - rentedCount)
                    + " | " + String.format("%04d", b.getReleaseYear())
                    + " | " + b.getAuthor() + " - " + b.getTitle());
        }
        System.out.println("======================================");
    }

    /**
     * Prints all customers.
     */
    public final void printCustomerList() {
        System.out.println("==== Customer list =====");
        System.out.println("CID | Rented | Full name");
        System.out.println("========================");
        for (Customer c : customerList) {
            int rentedCount = c.getRentedItemList().size();
            System.out.println(String.format("%03d", c.getID())
                    + " | " + String.format("%-6d", rentedCount)
                    + " | " + c.getFullName());
        }
        System.out.println("========================");
    }

    /**
     * Rents an item to a customer.
     * @param customerID Customer ID
     * @param itemID Item ID
     */
    public final void rentItem(final int customerID, final int itemID) {
        Customer customer = findCustomer(customerID);
        Item item = findItem(itemID);
        if (customer == null) {
            if (printErrorFeedback) {
                System.out.println("There is no customer of such ID.");
            }
            return;
        }

        if (item == null) {
            if (printErrorFeedback) {
                System.out.println("There is no item of such ID.");
            }
            return;
        }

        if (customer.getRentedItemList().contains(item)) {
            if (printErrorFeedback) {
                System.out.println(
                    "This customer is already renting this item.");
            }
            return;
        }

        if (item.isRented()) {
            if (printErrorFeedback) {
                System.err.println(
                    "This item is already rented to another customer.");
            }
            return;
        }

        customer.getRentedItemList().add(item);
        item.itemRent(customerID);

        customer.getRentedItemList().sort(itemIDComparator);
        if (printActionFeedback) {
            System.out.println("Rented item \""
                            + item.getTitle() + "\" to customer "
                            + customer.getFullName() + ".");
        }

    }

    /**
     * Returns an item to the library.
     * @param itemID Item ID
     */
    public final void returnItem(final int itemID) {
        Item item = findItem(itemID);

        if (item == null) {
            if (printErrorFeedback) {
                System.out.println("There is no item of such ID.");
            }
            return;
        }

        Customer customer = findCustomer(item.getCustomerID());
        if (customer == null) {
            if (printErrorFeedback) {
                System.out.println("No customer is renting this item.");
            }
            return;
        }

        customer.getRentedItemList().remove(item);
        item.itemReturn();

        if (printActionFeedback) {
            System.out.println("Item \"" + item.getTitle()
                                + "\" returned to stock.");
        }

    }

    /**
     * Prints all the rented items of a customer.
     * @param customerID Customer ID
     */
    public final void printCustomerRentedItems(final int customerID) {
        Customer customer = findCustomer(customerID);
        List<Item> rentedItemList = customer.getRentedItemList();
        System.out.println("Customer name: " + customer.getFullName());
        System.out.println("Customer ID: " + customerID);
        System.out.println("====== Rented Books =======");
        System.out.println("IID | BID | Year | Position");
        System.out.println("===========================");
        for (Item i : rentedItemList) {
            System.out.println(String.format("%03d", i.getItemID())
                    + " | " + String.format("%03d", i.getBookID())
                    + " | " + String.format("%04d", i.getReleaseYear())
                    + " | " + i.getAuthor() + " - " + i.getTitle());
        }
        System.out.println("===========================");
    }

    private int countTotalCopies(final Book book) {
        int count = 0;
        for (Item i : itemStock) {
            if (i.getBook().equals(book)) {
                count++;
            }
        }
        return count;
    }

    private int countRentedCopies(final Book book) {
        int count = 0;
        for (Item i : itemStock) {
            if (i.getBook().equals(book)
                    && i.isRented()) {
                count++;
            }
        }
        return count;
    }

    Library() {

    }
}
