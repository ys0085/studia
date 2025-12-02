%{
#include <stdio.h>
int started = 0;
int pending_space = 0;
int word_count = 0;
int line_count = 1;
%}
%option noyywrap
%%
[ \t\r]+      { if (started) pending_space = 1; }
\n            { line_count++; if (started) pending_space = 1; }
.             { if (pending_space) { putchar(' '); pending_space = 0; word_count++; } putchar(yytext[0]); started = 1; }
%%
int main(void)
{
    yylex();
    printf("\n%d %d", line_count, word_count + (started ? 1 : 0));
    return 0;
}