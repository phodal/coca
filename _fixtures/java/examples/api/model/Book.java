package com.phodal.pholedge.book.model;

import lombok.Builder;
import lombok.Getter;

import java.time.Instant;
import static java.time.Instant.now;

@Getter
@Builder
public class Book {
    private String id;
    private String isbn;
    private String name;
    private Instant createdAt;

    public static Book create(String id, String isbn, String name) {
        return Book.builder()
                .id(id)
                .isbn(isbn)
                .name(name)
                .createdAt(now())

                .build();

    }

    public BookRepresentaion toRepresentation() {
        return new BookRepresentaion(id, name);
    }

    public void save(String isbn, String name) {
        this.name = name;
        this.isbn = isbn;
    }
}
