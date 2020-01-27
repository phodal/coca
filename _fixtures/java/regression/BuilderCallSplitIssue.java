package com.phodal.coca.bug;

import static org.assertj.core.api.Assertions.assertThat;
import org.springframework.web.context.support.AnnotationConfigWebApplicationContext;

public class BuilderCallSplitIssue {
	void someAssert() {
	    AnnotationConfigApplicationContext ctx = new AnnotationConfigApplicationContext();

		assertThat(ctx.containsBean(TRANSITIVE_BEAN_NAME)).isFalse();
	}
}
