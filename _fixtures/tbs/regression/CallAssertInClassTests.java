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

package org.springframework.context.event;

import org.junit.jupiter.api.Test;

import org.springframework.context.ApplicationEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.core.ResolvableType;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;

/**
 * @author Stephane Nicoll
 */
@RunWith(PowerMockRunner.class)
public class CallAssertInClassTests extends AbstractApplicationEventListenerTests {

    @Mock
    Connection connection = PowerMockito.mock(Connection.class);

	@Test  // Demonstrates we cant inject that event because the listener has a raw type
	public void genericListenerRawTypeTypeErasure() {
		GenericTestEvent<String> stringEvent = createGenericTestEvent("test");
		ResolvableType eventType = ResolvableType.forType(stringEvent.getClass());
		supportsEventType(true, RawApplicationListener.class, eventType);
	}

	@SuppressWarnings("rawtypes")
	private void supportsEventType(
			boolean match, Class<? extends ApplicationListener> listenerType, ResolvableType eventType) {

		ApplicationListener<?> listener = mock(listenerType);
		GenericApplicationListenerAdapter adapter = new GenericApplicationListenerAdapter(listener);
		assertThat(adapter.supportsEventType(eventType)).as("Wrong match for event '" + eventType + "' on " + listenerType.getClass().getName()).isEqualTo(match);
	}
}
