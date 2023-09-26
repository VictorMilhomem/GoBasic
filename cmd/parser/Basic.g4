grammar Basic;

program
    : line* EOF
    ;

line
   : number statement (CR | EOF)
   | statement (CR | EOF)
   ;

statement
   : 'PRINT' exprlist
   | 'IF' expression relop expression 'THEN'? statement
   | 'GOTO' number
   | 'INPUT' varlist
   | 'LET'? vara '=' expression
   | 'GOSUB' expression
   | 'RETURN'
   | 'CLEAR'
   | 'LIST'
   | 'RUN'
   | 'END'
   ;

exprlist
   : (STRING | expression) (',' (STRING | expression))*
   ;

varlist
   : vara (',' vara)*
   ;

expression
   : ( unary)? term (('+' | '-') term)*
   ;

term
   : factor (mutlop factor)*
   ;

factor
   :
   vara
   | number
   ;

unary: '+' | '-';
mutlop: ('*' | '/');

vara
    : VAR
    | STRING
    ;

number
   : DIGIT +
   ;

relop
   : ('<' ('>' | '=' )?)
   | ('>' ('<' | '=' )?)
   | '='
   ;


STRING
   : '"' ~ ["\r\n]* '"'
   ;


DIGIT
   : '0' .. '9'
   ;


VAR
   : 'A' .. 'Z'
   ;


CR
   : [\r\n]+
   ;


WS
   : [ \t] -> skip
   ;