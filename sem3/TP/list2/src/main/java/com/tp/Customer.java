package com.tp;

import java.util.ArrayList;
import java.util.List;

/**
 * Customer class - stores info about the customer.
 */
public class Customer {
    private final String firstName;
    private final String lastName;
    private final List<Item> rentedItemList = new ArrayList<>();

    private final int id;

    /**
     * Returns the id of the customer.
     * @return ID
     */
    public final int getID() {
        return id;
    }

    /**
     * Getter full name of the customer.
     * @return customer full name
     */
    public final String getFullName() {
        return firstName + " " + lastName;
    }

    /**
     * Getter of first name of the customer.
     * @return customer first name
     */
    public final String getFirstName() {
        return firstName;
    }

    /**
     * Getter of first name of the customer.
     * @return customer last name
     */
    public final String getLastName() {
        return lastName;
    }

    /**
     * Getter of the customer's rented item list.
     * @return rented item list
     */
    public final List<Item> getRentedItemList() {
        return rentedItemList;
    }

    /**
     * Prints the customer's rented items.
     */
    public final void printRentedItems() {
        System.out.println("====== Rented Books =======");
        System.out.println("IID | BID | Year | Position");
        System.out.println("===========================");
        for (Item i : rentedItemList) {
            System.out.println(String.format("%03d", i.getItemID())
                    + " | " + String.format("%03d", i.getBookID())
                    + " | " + String.format("%04d", i.getReleaseYear())
                    + " | " + i.getAuthor() + " - " + i.getTitle());
        }

    }

    Customer(final String lN, final String fN, final int customerID) {
        this.firstName = fN;
        this.lastName = lN;
        this.id = customerID;
    }
}
