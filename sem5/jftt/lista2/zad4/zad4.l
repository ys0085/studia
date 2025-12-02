%{
#include <stdio.h>
#include <string.h>
#include <math.h>

struct Stack {
    int position;
    int *data;
};
    
int push(struct Stack* stack, int n) {
    if(stack->position >= 1023) return 0;
    stack->position++;
    stack->data[stack->position] = n;
    return 1;
}

int pop(struct Stack* stack) {
    stack->position--;
    return stack->data[stack->position+1];
}

int d[1024];
struct Stack stack = {-1, d};

void reset(struct Stack* stack) {
    for(int i = 0; i < 1024; i++) stack->data[i] = 0;
    stack->position = -1;
}

int _pow(int x, int y) {
    if(y == 0) return 1;
    if(y == 1) return x;
    int p = _pow(x, y/2);
    if(y % 2) return p*p*x;
    else return p*p;
}

int err_char = 0;
%}

%option noyywrap

%x ERROR
%x UNEXP_ERR
OPERATOR        [\+\-\*\/%\^]
%%
[\-]?[0-9]+    { push(&stack, atoi(yytext));}
[ \t]           {;}
{OPERATOR}      {
                    
                    if(stack.position < 1) {
                        printf("\nError: za mała liczba argumentów");
                        BEGIN(ERROR);
                    } 
                    else {
                        int a = pop(&stack);
                        int b = pop(&stack);
                        int ans;
                        switch(yytext[0]){
                            case '+': ans = b + a; break;
                            case '-': ans = b - a; break;
                            case '*': ans = b * a; break;
                            case '/': if(a == 0) {printf("\nError: dzielenie przez 0"); BEGIN(ERROR);}ans = b / a; break;
                            case '%': ans = b % a; break;
                            case '^': ans = _pow(b, a); break;
                            default: break;
                        }
                        push(&stack, ans);
                    }
                }

<ERROR>\n       {reset(&stack); BEGIN(INITIAL);}
<ERROR>.        {}
\r?\n           {           /* Windows and Linux newlines */
                    
                    if(stack.position != 0) {
                        printf("Error: za mała liczba operatorów\n");
                    }
                    else printf("= %d\n", pop(&stack));
                    reset(&stack);
                }
.               { 
                    err_char = yytext[0];
                    
                    BEGIN(UNEXP_ERR);
                }
<UNEXP_ERR>.    {}
<UNEXP_ERR>\n   {reset(&stack); printf("Error: niespodziewany znak '%c'\n", err_char); BEGIN(INITIAL);}
%%

int main(int argc, char* argv[]) {
    yylex();
    return 0;
}