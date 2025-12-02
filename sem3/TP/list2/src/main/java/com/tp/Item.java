package com.tp;

public class Item extends Book {
    private boolean rented;
    private int renterId;

    private final Book book;

    private final int itemID;

    /**
     * Getter of the item's ID.
     * @return Item ID
     */
    public final int getItemID() {
        return itemID;
    }

    /**
     * Check of whether the item is currently being rented.
     * @return true/false
     */
    public final boolean isRented() {
        return rented;
    }

    /**
     * Getter of the item's book position.
     * @return the book in question
     */
    public final Book getBook() {
        return book;
    }

    /**
     * Getter of the item's renter's ID.
     * @return -1 if the book is not being rented, else renter's ID
     */
    public final int getCustomerID() {
        return renterId;
    }

    /**
     * Rents the item to a specified customer.
     * @param customerId Customer ID
     */
    public final void itemRent(final int customerId) {
        rented = true;
        renterId = customerId;
    }

    /**
     * Returns the item to the library.
     */
    public final void itemReturn() {
        rented = false;
        renterId = -1;
    }

    Item(final Book b, final int id) {
        super(b);
        this.book = b;
        this.itemID = id;

        this.rented = false;
        renterId = -1;
    }
}
