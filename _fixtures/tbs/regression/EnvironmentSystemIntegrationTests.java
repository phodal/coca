/*
 * Copyright 2002-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.springframework.core.env;

import java.io.File;
import java.io.IOException;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import org.springframework.beans.factory.support.BeanDefinitionRegistry;
import org.springframework.beans.factory.support.DefaultListableBeanFactory;
import org.springframework.beans.factory.xml.XmlBeanDefinitionReader;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.EnvironmentAware;
import org.springframework.context.annotation.AnnotatedBeanDefinitionReader;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ClassPathBeanDefinitionScanner;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.context.annotation.Profile;
import org.springframework.context.support.ClassPathXmlApplicationContext;
import org.springframework.context.support.FileSystemXmlApplicationContext;
import org.springframework.context.support.GenericApplicationContext;
import org.springframework.context.support.GenericXmlApplicationContext;
import org.springframework.context.support.StaticApplicationContext;
import org.springframework.core.io.ClassPathResource;
import org.springframework.jca.context.ResourceAdapterApplicationContext;
import org.springframework.jca.support.SimpleBootstrapContext;
import org.springframework.jca.work.SimpleTaskWorkManager;
import org.springframework.mock.env.MockEnvironment;
import org.springframework.mock.env.MockPropertySource;
import org.springframework.mock.web.MockServletConfig;
import org.springframework.mock.web.MockServletContext;
import org.springframework.util.FileCopyUtils;
import org.springframework.web.context.WebApplicationContext;
import org.springframework.web.context.support.AbstractRefreshableWebApplicationContext;
import org.springframework.web.context.support.AnnotationConfigWebApplicationContext;
import org.springframework.web.context.support.GenericWebApplicationContext;
import org.springframework.web.context.support.StandardServletEnvironment;
import org.springframework.web.context.support.StaticWebApplicationContext;
import org.springframework.web.context.support.XmlWebApplicationContext;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatExceptionOfType;
import static org.springframework.beans.factory.support.BeanDefinitionBuilder.rootBeanDefinition;
import static org.springframework.context.ConfigurableApplicationContext.ENVIRONMENT_BEAN_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.DERIVED_DEV_BEAN_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.DERIVED_DEV_ENV_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.DEV_BEAN_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.DEV_ENV_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.ENVIRONMENT_AWARE_BEAN_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.PROD_BEAN_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.PROD_ENV_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.TRANSITIVE_BEAN_NAME;
import static org.springframework.core.env.EnvironmentSystemIntegrationTests.Constants.XML_PATH;

/**
 * System integration tests for container support of the {@link Environment} API.
 *
 * <p>
 * Tests all existing BeanFactory and ApplicationContext implementations to ensure that:
 * <ul>
 * <li>a standard environment object is always present
 * <li>a custom environment object can be set and retrieved against the factory/context
 * <li>the {@link EnvironmentAware} interface is respected
 * <li>the environment object is registered with the container as a singleton bean (if an
 * ApplicationContext)
 * <li>bean definition files (if any, and whether XML or @Configuration) are registered
 * conditionally based on environment metadata
 * </ul>
 *
 * @author Chris Beams
 * @author Sam Brannen
 * @see org.springframework.context.support.EnvironmentIntegrationTests
 */
@SuppressWarnings("resource")
public class EnvironmentSystemIntegrationTests {

	private final ConfigurableEnvironment prodEnv = new StandardEnvironment();

	private final ConfigurableEnvironment devEnv = new StandardEnvironment();

	private final ConfigurableEnvironment prodWebEnv = new StandardServletEnvironment();

	@BeforeEach
	void setUp() {
		prodEnv.setActiveProfiles(PROD_ENV_NAME);
		devEnv.setActiveProfiles(DEV_ENV_NAME);
		prodWebEnv.setActiveProfiles(PROD_ENV_NAME);
	}

	@Test
	void mostSpecificDerivedClassDrivesEnvironment_withDevEnvAndDerivedDevConfigClass() {
		AnnotationConfigApplicationContext ctx = new AnnotationConfigApplicationContext();
		ctx.setEnvironment(devEnv);
		ctx.register(DerivedDevConfig.class);
		ctx.refresh();

		assertThat(ctx.containsBean(DEV_BEAN_NAME)).isFalse();
		assertThat(ctx.containsBean(DERIVED_DEV_BEAN_NAME)).isFalse();
		assertThat(ctx.containsBean(TRANSITIVE_BEAN_NAME)).isFalse();
	}

	@Test
	void annotationConfigApplicationContext_withProfileExpressionMatchOr() {
		testProfileExpression(true, "p3");
	}

	@Test
	void annotationConfigApplicationContext_withProfileExpressionMatchAnd() {
		testProfileExpression(true, "p1", "p2");
	}

	@Test
	void annotationConfigApplicationContext_withProfileExpressionNoMatchAnd() {
		testProfileExpression(false, "p1");
	}

	@Test
	void annotationConfigApplicationContext_withProfileExpressionNoMatchNone() {
		testProfileExpression(false, "p4");
	}

	private void testProfileExpression(boolean expected, String... activeProfiles) {
		AnnotationConfigApplicationContext ctx = new AnnotationConfigApplicationContext();
		StandardEnvironment environment = new StandardEnvironment();
		environment.setActiveProfiles(activeProfiles);
		ctx.setEnvironment(environment);
		ctx.register(ProfileExpressionConfig.class);
		ctx.refresh();
		assertThat(ctx.containsBean("expressionBean")).isEqualTo(expected);
	}
}
