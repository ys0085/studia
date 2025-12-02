package com.tp;

import static org.junit.Assert.assertTrue;
import org.junit.Test;

/**
 * Unit test for simple App.
 */
public class AppTest {
    @Test
    public void test1() {
        Library lib = new Library();
        lib.addBook("ABC", "Adam A", 1234);
        lib.addItem(1);
        lib.addCustomer("Aa", "Bb");

        lib.rentItem(1, 1);

        lib.removeCustomer(1);
        
        assertTrue( lib.getItemStock().getFirst().isRented() == false );
    }

    @Test
    public void test2() {
        Library lib = new Library();
        lib.addBook("ABC", "Adam A", 1234);
        
        lib.addItem(1);
        lib.addItem(1);
        lib.addItem(1);
        lib.addItem(1);
        lib.addItem(1);


        lib.removeBook(1);

        assertTrue( lib.getItemStock().isEmpty() == true );
    }
}
