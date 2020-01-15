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
        String s = "# argument\n" +
                "#     : test (comp_for | ASSIGN test)?\n" +
                "#     | (POWER | STAR) test\n" +
                "#     ;\n" +
                "\n" +
                "# test\n" +
                "b(x)\n" +
                "\n" +
                "# test comp_for\n" +
                "b(x for x in a)\n" +
                "\n" +
                "# test ASSIGN test\n" +
                "b(x=i)\n" +
                "\n" +
                "# test COMMA test ASSIGN test COMMA test ASSIGN test\n" +
                "b(z, x=i, y=u)\n" +
                "\n" +
                "# POWER test\n" +
                "b(**z)\n" +
                "\n" +
                "# STAR test\n" +
                "b(*z)\n" +
                "\n" +
                "# test COMMA STAR test COMMA test ASSIGN test\n" +
                "b(y, *z, x=i)\n";
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
    @Test
    void testTryStmt() {
        String s = "# try_stmt: TRY COLON suite (except_clause+ else_clause? finaly_clause? | finaly_clause)\n" +
                "\n" +
                "# TRY COLON suite except_clause\n" +
                "try:\n" +
                "    pass\n" +
                "except:\n" +
                "    pass\n" +
                "\n" +
                "# TRY COLON suite except_clause except_clause else_clause\n" +
                "try:\n" +
                "    pass\n" +
                "except Exception as ex:\n" +
                "    pass\n" +
                "except:\n" +
                "    pass\n" +
                "else:\n" +
                "    pass\n" +
                "\n" +
                "# TRY COLON suite except_clause finaly_clause\n" +
                "try:\n" +
                "    pass\n" +
                "except Exception:\n" +
                "    pass\n" +
                "finally:\n" +
                "    pass\n" +
                "\n" +
                "# TRY COLON suite finaly_clause\n" +
                "try:\n" +
                "    pass\n" +
                "finally:\n" +
                "    pass\n";
        PythonIdentApp.processString(s);
    }
}
