# This is a single line comment
print("Hello World")  # This is an end of line comment

"""
This is a multi-line
comment that should be removed
"""

# Another comment
x = """This is a triple-quoted
string literal that should
be preserved"""

y = '''Another multi-line
string literal that should
stay in the code'''


data = [
    """This is a string
    in a list""",
    '''Another string
    in a list'''
]

# Final comment
print(x, y)