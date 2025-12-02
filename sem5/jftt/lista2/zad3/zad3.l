%{
#include <stdio.h>
#include <string.h>

int preserve_doxygen = 0;
int in_doxygen = 0;
%}

%option noyywrap

%x IN_BLOCK_COMMENT
%x IN_LINE_COMMENT

%%

"/**"                   { if (preserve_doxygen) ECHO; BEGIN(IN_BLOCK_COMMENT); in_doxygen = 1; }
"/*!"                   { if (preserve_doxygen) ECHO; BEGIN(IN_BLOCK_COMMENT); in_doxygen = 1; }

"/*"                    { BEGIN(IN_BLOCK_COMMENT); in_doxygen = 0; }

"///"[/!]?             { if (preserve_doxygen) ECHO; BEGIN(IN_LINE_COMMENT); in_doxygen = 1; }
"//!"                   { if (preserve_doxygen) ECHO; BEGIN(IN_LINE_COMMENT); in_doxygen = 1; }

"//"                    { BEGIN(IN_LINE_COMMENT); in_doxygen = 0; }

<IN_BLOCK_COMMENT>"*/"  { if (preserve_doxygen && in_doxygen) 
                            ECHO; 
                          BEGIN(INITIAL); 
                        }
<IN_BLOCK_COMMENT>\n    { if (preserve_doxygen && in_doxygen) 
                            ECHO; 
                          else 
                            putchar('\n'); 
                        }
<IN_BLOCK_COMMENT>.     { if (preserve_doxygen && in_doxygen) 
                            ECHO; 
                        }

<IN_LINE_COMMENT>\n     { putchar('\n'); BEGIN(INITIAL); }
<IN_LINE_COMMENT>.      { if (preserve_doxygen && in_doxygen) 
                            ECHO; 
                        }

\"([^\\\"]|\\.)*\"     { ECHO; }
\'([^\\\']|\\.)*\'     { ECHO; }

.|\n                    { ECHO; }

%%

int main(int argc, char* argv[]) {
    if (argc > 1 && strcmp(argv[1], "--preserve-doxygen") == 0) {
        preserve_doxygen = 1;
    }
    yylex();
    return 0;
}