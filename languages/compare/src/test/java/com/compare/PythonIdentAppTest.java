package com.compare;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class PythonIdentAppTest {

    @Test
    void processString() {
        String s = "class foo:\n    pass\n";
        PythonIdentApp.processString(s);
    }
}
