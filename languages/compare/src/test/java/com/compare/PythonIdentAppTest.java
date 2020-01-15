package com.compare;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class PythonIdentAppTest {

    @Test
    void processString() {
        String s = "class foo:\n    pass\n";
        PythonIdentApp.processString(s);
    }

    @Test
    void testStar() {
        String s = "b(**z)\nb(*z)";
        PythonIdentApp.processString(s);
    }

    @Test
    void testPy2() {
        String s = "print 'a'";
        PythonIdentApp.processString(s);
    }

    @Test
    void testPy3() {
        String s = "print('a')";
        PythonIdentApp.processString(s);
    }
}
