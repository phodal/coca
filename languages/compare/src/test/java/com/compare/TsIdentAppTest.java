package com.compare;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class TsIdentAppTest {
    @Test
    void shouldHandleField() {
        String str = "console.log('a')";
        TsIdentApp.processString(str);
    }
}
