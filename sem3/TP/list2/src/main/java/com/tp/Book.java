package com.tp;

public class Book {
    private final String title;
    private final String author;
    private final int releaseYear;

    private final int bookID;

    /**
     * Getter of the book's ID.
     * @return Book ID
     */
    public final int getBookID() {
        return bookID;
    }

    /**
     * Getter of the book's title.
     * @return book title
     */
    public final String getTitle() {
        return title;
    }

    /**
     * Getter of the book's author(s).
     * @return author string
     */
    public final String getAuthor() {
        return author;
    }

    /**
     * Getter of the book's release year.
     * @return release yeara
     */
    public final int getReleaseYear() {
        return releaseYear;
    }

    Book(final String t, final String a, final int rY, final int id) {
        this.title = t;
        this.author = a;
        this.releaseYear = rY;
        this.bookID = id;
    }

    Book(final Book b) {
        this(b.title, b.author, b.releaseYear, b.bookID);
    }
}
