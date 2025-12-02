%{
#include <stdio.h>
#include <stdlib.h>
#include "gf.tab.h"
%}

%option noyywrap

%%

[0-9]+          { yylval.num = atoll(yytext); return NUMBER; }
"+"             { return PLUS; }
"-"             { return MINUS; }
"*"             { return MULT; }
"/"             { return DIV; }
"^"             { return POW; }
"("             { return LPAREN; }
")"             { return RPAREN; }
"#"[^\n]*       { /* komentarz - ignoruj */ }
"\\\n"          { /* kontynuacja linii - ignoruj */ }
[ \t]+          { /* białe znaki - ignoruj */ }
\n              { return NEWLINE; }
.               { /* inne znaki - ignoruj */ }

%%