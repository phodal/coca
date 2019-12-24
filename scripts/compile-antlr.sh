#!/usr/bin/env bash

cd core/languages/g4

antlr -Dlanguage=Go -listener JavaLexer.g4 -o ../java
antlr -Dlanguage=Go -listener JavaParser.g4 -o ../java

antlr -Dlanguage=Go -listener Sql.g4 -o ../sql
