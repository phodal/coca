parser grammar RefactorMethodSignatureParser;

options { tokenVocab=AspectJLexer; }

import JavaParser;

classNameOrInterface
    :	(IDENTIFIER | '*' | '.' | '..')+ ('[' ']')*
    ;
