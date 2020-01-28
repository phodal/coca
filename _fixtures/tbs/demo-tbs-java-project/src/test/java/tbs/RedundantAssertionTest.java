package tbs;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class RedundantAssertionTest {
    @Test
    public void testTrue() {
        assertEquals(true, true);
    }
}
