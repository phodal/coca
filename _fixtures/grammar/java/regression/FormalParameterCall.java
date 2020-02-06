package regression;

import hello.CreateBookCommand;

@Component
public class BookService implements Service {
    @Transactional
    public void getIsbnId(CreateBookCommand command) {
        command.getIsbn();
    }
}
