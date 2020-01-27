package com.compare;

import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTreeWalker;
import pyantlr.PythonLexer;
import pyantlr.PythonParser;

public class PythonIdentApp {
    static void processString(String inputStr) {
        CharStream stream = CharStreams.fromString(inputStr);;
        PythonLexer lexer = new PythonLexer(stream);
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        PythonParser parser = new PythonParser(tokens);
        PythonParser.RootContext tree = parser.root(); // see the grammar ->

        PythonIdentListener pythonIdentListener = new PythonIdentListener();
        ParseTreeWalker.DEFAULT.walk(pythonIdentListener, tree);
    }
}
