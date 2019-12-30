package study.huhao.demo.domain.models.blog;

import study.huhao.demo.domain.core.Page;
import study.huhao.demo.domain.core.Repository;

import java.util.Optional;
import java.util.UUID;

public interface BlogRepository extends Repository {
    void save(Blog blog);

    Optional<Blog> findById(UUID id);

    boolean existsById(UUID id);

    void deleteById(UUID id);

    Page<Blog> findAllWithPagination(BlogCriteria criteria);
}
