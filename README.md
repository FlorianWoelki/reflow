# reflow

## Interpreter Packages

- `lexer`: takes source code as input and outputs tokens that represent the source code,
- `lexer/token`: defines the tokens and useful functions to look up tokens,
- `parser`: takes source code (in tokens) as input and produces a data structure which represents the source code,
- `evaluator`: takes the AST and interpret it on the fly via a tree-walking interpreter without preprocessing or extra compilation,
- `object`: value system or object representation and it represents the values it generates when evaluating the AST,
- `repl`: read eval print loop
