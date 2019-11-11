package com.phodal.pholedge.book;

import com.phodal.pholedge.book.model.Book;
import com.phodal.pholedge.book.model.BookRepresentaion;
import com.phodal.pholedge.core.domain.Repository;
import com.phodal.pholedge.core.exception.NotFoundException;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class BookRepository implements Repository {
    private final BookMapper bookMapper;

    public BookRepository(BookMapper bookMapper) {
        this.bookMapper = bookMapper;
    }

    public void save(Book book) {
        this.bookMapper.doSave(book);
    }

    public List<BookRepresentaion> list() {
        return this.bookMapper.list();
    }

    public Book byId(String id) {
        try {
            return this.bookMapper.byId(id);
        } catch (EmptyResultDataAccessException e) {
            throw new NotFoundException(id);
        }
    }
}
