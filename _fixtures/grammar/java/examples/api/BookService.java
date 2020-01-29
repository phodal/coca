package com.phodal.pholedge.book;

import com.phodal.pholedge.book.model.Book;
import com.phodal.pholedge.book.model.BookRepresentaion;
import com.phodal.pholedge.book.model.command.CreateBookCommand;
import com.phodal.pholedge.book.model.command.UpdateBookCommand;
import com.phodal.pholedge.core.domain.Service;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import javax.validation.Valid;
import java.util.List;

@Component
public class BookService implements Service {
    private final BookFactory bookFactory;
    private final BookRepository bookRepository;

    public BookService(BookFactory bookFactory, BookRepository bookRepository) {
        this.bookFactory = bookFactory;
        this.bookRepository = bookRepository;
    }

    @Transactional
    public String createBook(CreateBookCommand command) {
        Book book = bookFactory.create(command.getIsbn(), command.getName());
        bookRepository.save(book);
        return book.getId();
    }

    public List<BookRepresentaion> getBooksLists() {
        return bookRepository.list();
    }

    public BookRepresentaion getBookById(String id) {
        Book book = bookRepository.byId(id);
        return book.toRepresentation();
    }

    public BookRepresentaion updateBook(String id, @Valid UpdateBookCommand command) {
        Book book = bookRepository.byId(id);
        book.save(command.getIsbn(), command.getName());
        bookRepository.save(book);
        return book.toRepresentation();
    }
}
