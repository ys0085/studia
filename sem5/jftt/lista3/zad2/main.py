# ...existing code...
"""
PLY-based calculator: builds postfix with the same grammar/behavior as the Bison/Flex example.
- Modulo field GF(1234577).
- Unary minus only for numbers: "-1--2" -> (-1)-(-2).
- 2^-2 gets rewritten to 2^(mod(-2)-1) before evaluation.
- Negative numeric tokens are canonicalized to their modular representative.
"""

import sys
import ply.lex as lex
import ply.yacc as yacc

GF_MOD = 1234577
MAX_OUTPUT = 10000


# --- Helpers & modular arithmetic ---

def mod(x: int) -> int:
    r = x % GF_MOD
    if r < 0:
        r += GF_MOD
    return r


def egcd(a: int, b: int):
    # extended gcd
    if a == 0:
        return b, 0, 1
    gcd, x1, y1 = egcd(b % a, a)
    x = y1 - (b // a) * x1
    y = x1
    return gcd, x, y


def mod_inverse(a: int):
    a = mod(a)
    if a == 0:
        return None
    gcd, x, y = egcd(a, GF_MOD)
    if gcd != 1:
        return None
    return x % GF_MOD


def mod_pow_pos(base: int, exp: int) -> int:
    result = 1
    base = mod(base)
    while exp > 0:
        if exp & 1:
            result = mod(result * base)
        base = mod(base * base)
        exp >>= 1
    return result


def power(base: int, exp: int) -> int:
    if exp == 0:
        return 1
    if exp < 0:
        inv = mod_inverse(base)
        if inv is None:
            raise ValueError("No modular inverse for dividing by base")
        return mod_pow_pos(inv, -exp)
    return mod_pow_pos(base, exp)


# --- Postfix builder & evaluator ---

postfix = ""
postfix_len = 0


def add_to_postfix(tok: str):
    global postfix, postfix_len
    if postfix_len > 0:
        postfix += " "
        postfix_len += 1
    postfix += tok
    postfix_len += len(tok)


def is_number_token(s: str) -> bool:
    if not s:
        return False
    i = 0
    if s[0] == '-':
        if len(s) == 1:
            return False
        i = 1
    digits = False
    for ch in s[i:]:
        if not ch.isdigit():
            return False
        digits = True
    return digits and (i == 0 or digits)


def preprocess_postfix():
    """
    Modify postfix string:
    1) For each '^', if the token immediately before it is a negative number "-n",
       replace it with (mod(-n) - 1) (modded into 0..GF_MOD-1).
    2) Convert all remaining negative number tokens into their modular representatives mod(n).
    """
    global postfix, postfix_len
    if postfix_len == 0:
        return

    tokens = postfix.split()
    tokcount = len(tokens)

    for i in range(tokcount):
        if tokens[i] == "^" and i >= 1 and is_number_token(tokens[i - 1]) and tokens[i - 1].startswith("-"):
            exp = int(tokens[i - 1])
            newexp = mod(exp) - 1
            newexp = mod(newexp)
            tokens[i - 1] = str(newexp)

    for i in range(tokcount):
        if tokens[i].startswith("-") and is_number_token(tokens[i]):
            val = int(tokens[i])
            tokens[i] = str(mod(val))

    postfix = " ".join(tokens)
    postfix_len = len(postfix)


def evaluate_postfix() -> int:
    stack = []
    if postfix_len == 0:
        return 0
    for token in postfix.split():
        if token == "+":
            b = stack.pop()
            a = stack.pop()
            stack.append(mod(a + b))
        elif token == "-":
            b = stack.pop()
            a = stack.pop()
            stack.append(mod(a - b))
        elif token == "*":
            b = stack.pop()
            a = stack.pop()
            stack.append(mod(a * b))
        elif token == "/":
            b = stack.pop()
            a = stack.pop()
            inv_b = mod_inverse(b)
            if inv_b is None:
                raise ValueError("Division by non-invertible element")
            stack.append(mod(a * inv_b))
        elif token == "^":
            b = stack.pop()
            a = stack.pop()
            stack.append(power(a, b))
        else:
            stack.append(int(token))
    if stack:
        return stack.pop()
    return 0


# --- LEXER (PLY) ---

tokens = (
    "NUMBER",
    "PLUS",
    "MINUS",
    "MULT",
    "DIV",
    "POW",
    "LPAREN",
    "RPAREN",
    "NEWLINE",
)

t_PLUS = r'\+'
t_MULT = r'\*'
t_DIV = r'/'
t_POW = r'\^'
t_LPAREN = r'\('
t_RPAREN = r'\)'


def t_NUMBER(t):
    r'\d+'
    t.value = int(t.value)
    return t


t_MINUS = r'-'


t_ignore = ' \t\r'


def t_comment(t):
    r'\#.*'
    pass


def t_NEWLINE(t):
    r'\n+'
    t.type = 'NEWLINE'
    t.value = '\n'
    return t


def t_error(t):
    print(f"Illegal character '{t.value[0]}'")
    t.lexer.skip(1)


lexer = lex.lex()


# --- PARSER (PLY YACC) ---

precedence = (
    ('left', 'PLUS', 'MINUS'),
    ('left', 'MULT', 'DIV'),
    ('right', 'POW'),
)


def p_input_empty(p):
    'input : '
    pass


def p_input_lines(p):
    'input : input line'
    pass


def p_line_newline(p):
    'line : NEWLINE'
    pass


def p_line_expr(p):
    'line : expr NEWLINE'
    global postfix, postfix_len

    preprocess_postfix()
    print(postfix)
    try:
        result = evaluate_postfix()
        print(f"Wynik: {result}")
    except Exception as e:
        print("Error:", e)
    postfix = ""
    postfix_len = 0


def p_expr_number(p):
    'expr : NUMBER'
    add_to_postfix(str(p[1]))


def p_expr_minus_number(p):
    'expr : MINUS NUMBER'
    add_to_postfix(str(-p[2]))


def p_expr_paren(p):
    'expr : LPAREN expr RPAREN'


def p_expr_binop_plus(p):
    'expr : expr PLUS expr'
    add_to_postfix("+")


def p_expr_binop_minus(p):
    'expr : expr MINUS expr'
    add_to_postfix("-")


def p_expr_binop_mult(p):
    'expr : expr MULT expr'
    add_to_postfix("*")


def p_expr_binop_div(p):
    'expr : expr DIV expr'
    add_to_postfix("/")


def p_expr_binop_pow(p):
    'expr : expr POW expr'
    add_to_postfix("^")


def p_error(p):
    if p:
        print("Syntax error at token", p.type)
    else:
        print("Syntax error at EOF")


parser = yacc.yacc()


# --- Main REPL ---

def main():
    print("PLY postfix calculator in GF(1234577)")
    print("Enter expressions; blank lines or EOF to quit.")
    input_buffer = ""
    for line in sys.stdin:
        if '#' in line:
            line = line.split('#', 1)[0]
        if not line.strip():
            continue
        if line.rstrip().endswith('\\'):
            part = line.rstrip('\n')
            part = part.rstrip()
            part = part[:-1] 
            input_buffer += part + ' '
            continue
        input_buffer += line
        if not input_buffer.endswith("\n"):
            input_buffer += "\n"
        parser.parse(input_buffer, lexer=lexer)


if __name__ == "__main__":
    main()
