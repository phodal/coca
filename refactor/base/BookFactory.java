package com.phodal.pholedge.book;

import com.phodal.pholedge.book.model.Book;
import com.phodal.pholedge.core.IdGenerator;
import org.springframework.stereotype.Component;

@Component
public class BookFactory {
    private final IdGenerator idGenerator;

    public BookFactory(IdGenerator idGenerator) {
        this.idGenerator = idGenerator;
    }

    public Book create(String isbn, String name) {
        String bookId = idGenerator.generate();
        return Book.create(bookId, isbn, name);
    }
}
