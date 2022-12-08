#!/usr/bin/env bash

cd languages/java

antlr -Dlanguage=Go -listener JavaLexer.g4
antlr -Dlanguage=Go -listener JavaParser.g4

cd ../g4

antlr -Dlanguage=Go -listener GroovyLexer.g4 -o ../groovy
antlr -Dlanguage=Go -listener GroovyParser.g4 -o ../groovy

antlr -Dlanguage=Go -listener JavaScriptLexer.g4 -o ../js
antlr -Dlanguage=Go -listener JavaScriptParser.g4 -o ../js


#antlr -Dlanguage=Go -listener GoLexer.g4 -o ../go
#antlr -Dlanguage=Go -listener GoParser.g4 -o ../go

antlr -Dlanguage=Go -listener PythonLexer.g4 -o ../python
antlr -Dlanguage=Go -listener PythonParser.g4 -o ../python

#antlr -Dlanguage=Java -listener PythonLexer.g4 -o ../compare/java
#antlr -Dlanguage=Java -listener PythonParser.g4 -o ../compare/java

#antlr -Dlanguage=Java -listener TypeScriptLexer.g4 -o ../compare/src/main/java/tsantlr
#antlr -Dlanguage=Java -listener TypeScriptParser.g4 -o ../compare/src/main/java/tsantlr

antlr -Dlanguage=Go -listener CommentLexer.g4 -o ../comment
