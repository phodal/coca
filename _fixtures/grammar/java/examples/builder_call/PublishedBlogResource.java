package study.huhao.demo.adapters.inbound.rest.resources.publishedblog;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.util.UriComponents;
import org.springframework.web.util.UriComponentsBuilder;
import study.huhao.demo.application.usecases.EditBlogUseCase;
import study.huhao.demo.application.usecases.QueryPublishedBlogUseCase;

import static org.springframework.http.MediaType.APPLICATION_JSON_VALUE;

@RestController
@RequestMapping(value = "/published-blog", produces = APPLICATION_JSON_VALUE)
public class PublishedBlogResource {

    private final EditBlogUseCase editBlogUseCase;
    private final QueryPublishedBlogUseCase queryPublishedBlogUseCase;

    @Autowired
    public PublishedBlogResource(EditBlogUseCase editBlogUseCase, QueryPublishedBlogUseCase queryPublishedBlogUseCase) {
        this.editBlogUseCase = editBlogUseCase;
        this.queryPublishedBlogUseCase = queryPublishedBlogUseCase;
    }

    @PostMapping(consumes = APPLICATION_JSON_VALUE)
    public ResponseEntity<?> post(@RequestBody PublishBlogRequest data, UriComponentsBuilder uriComponentsBuilder) {
        var blog = editBlogUseCase.publish(data.blogId);

        UriComponents uriComponents = uriComponentsBuilder.path("/published-blog/{id}").buildAndExpand(blog.getId());
        return ResponseEntity.created(uriComponents.toUri()).body(PublishedBlogDto.of(blog));
    }
}
