package com.phodal.pholedge.book;

import com.phodal.pholedge.book.model.BookRepresentaion;
import com.phodal.pholedge.book.model.command.CreateBookCommand;
import com.phodal.pholedge.book.model.command.UpdateBookCommand;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.util.List;
import java.util.Map;

import static com.google.common.collect.ImmutableSortedMap.of;

@RestController
@RequestMapping(value = "/books")
public class BookController {
    private final BookService applicationService;

    public BookController(BookService applicationService) {
        this.applicationService = applicationService;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Map<String, String> createBook(@RequestBody @Valid CreateBookCommand command) {
        return of("id", applicationService.createBook(command));
    }

    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public BookRepresentaion updateBook(@PathVariable(name = "id") String id, @RequestBody @Valid UpdateBookCommand command) {
        return applicationService.updateBook(id, command);
    }

    @GetMapping("/")
    public List<BookRepresentaion> getBookList() {
        return applicationService.getBooksLists();
    }

    @GetMapping("/{id}")
    public BookRepresentaion getBookById(@PathVariable(name = "id") String id) {
        return applicationService.getBookById(id);
    }
}
