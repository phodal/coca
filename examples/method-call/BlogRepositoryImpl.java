package study.huhao.demo.infrastructure.persistence.blog;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import study.huhao.demo.domain.core.Page;
import study.huhao.demo.domain.models.blog.Blog;
import study.huhao.demo.domain.models.blog.BlogCriteria;
import study.huhao.demo.domain.models.blog.BlogRepository;

import java.util.Optional;
import java.util.UUID;

import static java.util.stream.Collectors.toList;


@Component
public class BlogRepositoryImpl implements BlogRepository {
    @Override
    public void save(Blog blog) {
        var blogPO = BlogPO.of(blog);

        blogMapper.findById(blog.getId().toString()).ifPresentOrElse(
                po -> blogMapper.update(blogPO),
                () -> blogMapper.insert(blogPO)
        );
    }
}
