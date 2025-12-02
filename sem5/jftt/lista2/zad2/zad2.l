%{
#include <stdio.h>

%}

%option noyywrap

%x IN_DOCSTRING_SINGLE
%x IN_DOCSTRING_DOUBLE
%x IN_STRING_SINGLE
%x IN_STRING_DOUBLE

%%
    /* Single line comments */
#.*$                    { /* ignore */ }

    /* String literals - must be handled before docstrings */
[=({[\,][:blank:]*\'\'\'   { ECHO; BEGIN(IN_STRING_SINGLE); }
[=({[\,][:blank:]*\"\"\"   { ECHO; BEGIN(IN_STRING_DOUBLE); }

    /* Docstrings */
^[:blank:]*\'\'\'          { BEGIN(IN_DOCSTRING_SINGLE); }
^[:blank:]*\"\"\"          { BEGIN(IN_DOCSTRING_DOUBLE); }

    /* Handle string content */
<IN_STRING_SINGLE>\'\'\'    { ECHO; BEGIN(INITIAL); }
<IN_STRING_SINGLE>.|\n      { ECHO; }
<IN_STRING_DOUBLE>\"\"\"    { ECHO; BEGIN(INITIAL); }
<IN_STRING_DOUBLE>.|\n      { ECHO; }

    /* Handle docstring content */
<IN_DOCSTRING_SINGLE>\'\'\'    { BEGIN(INITIAL); }
<IN_DOCSTRING_SINGLE>.|\n      { /* ignore */ }
<IN_DOCSTRING_DOUBLE>\"\"\"    { BEGIN(INITIAL); }
<IN_DOCSTRING_DOUBLE>.|\n      { /* ignore */ }

    /* Regular strings */
\'[^\'\n\\]*\'         { ECHO; }
\"[^\"\n\\]*\"         { ECHO; }

    /* Everything else */
.|\n                    { ECHO; }

%%


int main() {
    yylex();
    return 0;
}