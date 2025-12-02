%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

void yyerror(const char *s);
int yylex(void);

#define GF_MOD 1234577
#define MAX_STACK 1000
#define MAX_OUTPUT 10000

long long stack[MAX_STACK];
int stack_top = -1;

char postfix[MAX_OUTPUT];
int postfix_len = 0;

void push(long long val) {
    if (stack_top < MAX_STACK - 1) {
        stack[++stack_top] = val;
    }
}

long long pop() {
    if (stack_top >= 0) {
        return stack[stack_top--];
    }
    return 0;
}

void add_to_postfix(const char *str) {
    if (postfix_len > 0) {
        postfix[postfix_len++] = ' ';
    }
    strcpy(postfix + postfix_len, str);
    postfix_len += strlen(str);
}

long long mod(long long x) {
    long long result = x % GF_MOD;
    if (result < 0) result += GF_MOD;
    return result;
}

/* Rozszerzony algorytm Euklidesa dla odwrotności modularnej */
long long mod_inverse(long long a) {
    long long m = GF_MOD;
    long long m0 = m, x0 = 0, x1 = 1;
    
    a = mod(a);
    if (a == 0) {
        fprintf(stderr, "Błąd: dzielenie przez zero\n");
        return 0;
    }
    
    while (a > 1) {
        long long q = a / m;
        long long t = m;
        
        m = a % m;
        a = t;
        t = x0;
        
        x0 = x1 - q * x0;
        x1 = t;
    }
    
    if (x1 < 0) x1 += m0;
    
    return x1;
}

long long mod_pow_pos(long long base, long long exp) {
    long long result = 1;
    base = mod(base);
    while (exp > 0) {
        if (exp & 1) result = mod(result * base);
        base = mod(base * base);
        exp >>= 1;
    }
    return result;
}


long long power(long long base, long long exp) {
    if (exp == 0) return 1;
    if (exp < 0) {
        long long inv = mod_inverse(base);
        return mod_pow_pos(inv, -exp);
    }
    return mod_pow_pos(base, exp);
}

int is_number(const char *str) {
    if (str == NULL || *str == '\0') return 0;
    int i = 0;
    if (str[0] == '-') {
        if (str[1] == '\0') return 0; 
        i = 1;
    }
    int digits = 0;
    for (; str[i] != '\0'; i++) {
        if (str[i] < '0' || str[i] > '9') return 0;
        digits = 1;
    }
    return digits;
}

void preprocess_postfix() {
    if (postfix_len == 0) return;

    char copy[MAX_OUTPUT];
    strcpy(copy, postfix);

    /* Collect tokens */
    char *tokens[2048];
    int tokcount = 0;
    char *tok = strtok(copy, " ");
    while (tok != NULL && tokcount < 2048) {
        tokens[tokcount++] = strdup(tok);
        tok = strtok(NULL, " ");
    }

    /* Adjust negative exponents: token before '^' */
    for (int i = 0; i < tokcount; i++) {
        if (strcmp(tokens[i], "^") == 0) {
            if (i >= 1 && tokens[i-1][0] == '-' && is_number(tokens[i-1])) {
                long long exp = atoll(tokens[i-1]);  /* negative signed */
                long long newexp = mod(exp) - 1;     /* mod(-2)-1 */
                newexp = mod(newexp);                /* ensure 0..GF_MOD-1 */
                char buf[64];
                sprintf(buf, "%lld", newexp);
                free(tokens[i-1]);
                tokens[i-1] = strdup(buf);
            }
        }
    }

    for (int i = 0; i < tokcount; i++) {
        if (tokens[i][0] == '-' && is_number(tokens[i])) {
            long long val = atoll(tokens[i]);
            long long modval = mod(val);  /* convert -n -> (GF_MOD - n) in range 0..GF_MOD-1 */
            char buf[64];
            sprintf(buf, "%lld", modval);
            free(tokens[i]);
            tokens[i] = strdup(buf);
        }
    }

    /* Rebuild postfix */
    postfix_len = 0;
    postfix[0] = '\0';
    for (int i = 0; i < tokcount; i++) {
        if (postfix_len > 0) postfix[postfix_len++] = ' ';
        strcpy(postfix + postfix_len, tokens[i]);
        postfix_len += strlen(tokens[i]);
        free(tokens[i]);
    }
}

long long evaluate_postfix() {
    stack_top = -1;
    char postfix_copy[MAX_OUTPUT];
    strcpy(postfix_copy, postfix);
    char *token = strtok(postfix_copy, " ");
    
    while (token != NULL) {
        if (strcmp(token, "+") == 0) {
            long long b = pop();
            long long a = pop();
            push(mod(a + b));
        } else if (strcmp(token, "-") == 0) {
            long long b = pop();
            long long a = pop();
            push(mod(a - b));
        } else if (strcmp(token, "*") == 0) {
            long long b = pop();
            long long a = pop();
            push(mod(a * b));
        } else if (strcmp(token, "/") == 0) {
            long long b = pop();
            long long a = pop();
            long long b_inv = mod_inverse(b);
            push(mod(a * b_inv));
        } else if (strcmp(token, "^") == 0) {
            long long b = pop();
            long long a = pop();
            push(power(a, b));
        } else {
            push(atoll(token));
        }
        token = strtok(NULL, " ");
    }
    
    return pop();
}

%}

%union {
    long long num;
}

%token <num> NUMBER
%token PLUS MINUS MULT DIV POW LPAREN RPAREN NEWLINE


%left PLUS MINUS 
%left MULT DIV

%right POW


%%

input:
    /* pusty */
    | input line
    ;

line:
    NEWLINE
    | expr NEWLINE {
        preprocess_postfix();
        printf("%s\n", postfix);
        long long result = evaluate_postfix();
        printf("Wynik: %lld\n", result);
        postfix[0] = '\0';
        postfix_len = 0;
    }
    ;

expr:
    NUMBER {
        char buf[100];
        sprintf(buf, "%lld", $1);
        add_to_postfix(buf);
    }
    | expr PLUS expr {
        add_to_postfix("+");
    }
    | expr MINUS expr {
        add_to_postfix("-");
    }
    | expr MULT expr {
        add_to_postfix("*");
    }
    | expr DIV expr {
        add_to_postfix("/");
    }
    | expr POW expr {
        add_to_postfix("^");
    }
    | LPAREN expr RPAREN {q
        /* nawiasy kontrolują kolejność w parserze */
    }
    | MINUS NUMBER {
        char buf[100];
        sprintf(buf, "%lld", -$2);
        add_to_postfix(buf);
    }
    ;

%%

void yyerror(const char *s) {
    fprintf(stderr, "Błąd: %s\n", s);
}

int main(void) {
    printf("Kalkulator z notacją postfiksową w GF(1234577)\n");
    printf("Podaj wyrażenia (# dla komentarza, \\ dla kontynuacji):\n");
    return yyparse();
}