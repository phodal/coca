package com.phodal.pholedge.book;

import com.phodal.pholedge.book.model.Book;
import com.phodal.pholedge.book.model.BookRepresentaion;
import org.apache.ibatis.annotations.Mapper;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
@Mapper
public interface BookMapper {

    void doSave(Book book);

    List<BookRepresentaion> list();

    Book byId(String id);
}
