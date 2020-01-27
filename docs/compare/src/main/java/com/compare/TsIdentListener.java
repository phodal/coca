package com.compare;

import tsantlr.TypeScriptParser;
import tsantlr.TypeScriptParserBaseListener;

public class TsIdentListener extends TypeScriptParserBaseListener {
    @Override
    public void enterClassDeclaration(TypeScriptParser.ClassDeclarationContext ctx) {
        super.enterClassDeclaration(ctx);
    }
}
