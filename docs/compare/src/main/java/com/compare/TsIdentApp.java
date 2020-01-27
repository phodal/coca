package com.compare;

import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.tree.ParseTreeWalker;
import pyantlr.PythonLexer;
import pyantlr.PythonParser;
import tsantlr.TypeScriptLexer;
import tsantlr.TypeScriptParser;

public class TsIdentApp {
    static void processString(String inputStr) {
        CharStream stream = CharStreams.fromString(inputStr);;
        TypeScriptLexer lexer = new TypeScriptLexer(stream);
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        TypeScriptParser parser = new TypeScriptParser(tokens);
        TypeScriptParser.ProgramContext tree = parser.program();

        TsIdentListener listener = new TsIdentListener();
        ParseTreeWalker.DEFAULT.walk(listener, tree);
    }
}
