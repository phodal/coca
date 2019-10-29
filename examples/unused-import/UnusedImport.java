package study.huhao.demo.adapters.restapi.providers;

import study.huhao.demo.domain.core.excpetions.EntityExistedException;

import javax.ws.rs.core.Response;
import javax.ws.rs.ext.ExceptionMapper;
import javax.ws.rs.ext.Provider;
import java.util.Map;

import static javax.ws.rs.core.Response.Status.CONFLICT;
import static javax.ws.rs.core.Response.status;


@Provider
public class EntityExistedExceptionMapper implements ExceptionMapper<EntityExistedException> {
    @Override
    public Response toResponse(EntityExistedException ex) {
        var entity = Map.of("message", ex.getMessage());
        return status(CONFLICT).entity(entity).build();
    }
}
