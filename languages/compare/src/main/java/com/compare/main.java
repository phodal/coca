package com.compare;

import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.tree.ParseTreeWalker;
import pyantlr.PythonLexer;
import pyantlr.PythonParser;
import org.antlr.v4.runtime.CommonTokenStream;

public class main {
    public static void main(String[] args) {
        CharStream input = CharStreams.fromString("class foo:\n    pass\n");;

        PythonLexer lexer = new PythonLexer(input);
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        PythonParser parser = new PythonParser(tokens);
        PythonParser.RootContext tree = parser.root(); // see the grammar ->

        PythonIdentListener pythonIdentListener = new PythonIdentListener();
        ParseTreeWalker.DEFAULT.walk(pythonIdentListener, tree);
    }
}
