package com.compare;

import pyantlr.PythonParser;
import pyantlr.PythonParserBaseListener;

public class PythonIdentListener extends PythonParserBaseListener {
    @Override
    public void enterSingle_input(PythonParser.Single_inputContext ctx) {
        super.enterSingle_input(ctx);
    }

    @Override
    public void enterClassdef(PythonParser.ClassdefContext ctx) {
        System.out.println(ctx.name().getText());
        super.enterClassdef(ctx);
    }
}
