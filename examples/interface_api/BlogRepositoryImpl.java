package study.huhao.demo.adapters.outbound.persistence.blog;

import org.springframework.stereotype.Repository;
import study.huhao.demo.domain.contexts.blogcontext.blog.Blog;
import study.huhao.demo.domain.contexts.blogcontext.blog.BlogCriteria;
import study.huhao.demo.domain.contexts.blogcontext.blog.BlogRepository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static java.util.stream.Collectors.toList;


@Repository
public class BlogRepositoryImpl implements BlogRepository {

    private final BlogMapper blogMapper;

    @Override
    public List<Blog> findAll(BlogCriteria criteria) {
        return blogMapper.selectAllByCriteria(criteria).stream().map(BlogPO::toDomainModel).collect(toList());
    }

}

