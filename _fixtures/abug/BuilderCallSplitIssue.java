package com.phodal.coca.bug;

import static org.assertj.core.api.Assertions.assertThat;

public class BuilderCallSplitIssue {
	void someAssert() {
		assertThat(ctx.containsBean(TRANSITIVE_BEAN_NAME)).isFalse();
	}
}
