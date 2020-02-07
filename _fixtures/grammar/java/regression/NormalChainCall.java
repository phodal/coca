package hello.outer;

import org.springframework.web.util.UriComponents;
import org.springframework.web.util.UriComponentsBuilder;

public class PublishedBlogResource {
    public ResponseEntity post(UriComponentsBuilder uriComponentsBuilder) {
        UriComponents uriComponents = uriComponentsBuilder.path("/published-blog/{id}").buildAndExpand("1234");
//         return uriComponents.toUri();
    }
}
