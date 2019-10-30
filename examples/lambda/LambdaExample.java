package study.huhao.demo.infrastructure.persistence.blog;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import study.huhao.demo.domain.core.Page;
import study.huhao.demo.domain.models.blog.Blog;
import study.huhao.demo.domain.models.blog.BlogCriteria;
import study.huhao.demo.domain.models.blog.BlogRepository;

import java.util.Optional;
import java.util.UUID;

import static study.huhao.demo.infrastructure.persistence.utils.PaginationUtil.buildPageRequest;

@Component
public class BlogRepositoryImpl implements BlogRepository {
    @Override
    public Page<Blog> findAllWithPagination(BlogCriteria criteria) {
        var pagedBlog = blogJpaRepository.findAll(buildPageRequest(criteria))
                .map(BlogPO::toDomainModel);

        return "";
    }
}
