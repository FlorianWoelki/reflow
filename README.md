# reflow

reflow is a custom programming languages that includes a custom built interpreter (and future compiler). This language was designed to learn more about what programming languages are and what interpreters and compilers are.

## Interpreter Packages

- `lexer`: takes source code as input and outputs tokens that represent the source code,
- `lexer/token`: defines the tokens and useful functions to look up tokens,
- `parser`: takes source code (in tokens) as input and produces a data structure which represents the source code,
- `evaluator`: takes the AST and interpret it on the fly via a tree-walking interpreter without preprocessing or extra compilation,
- `object`: value system or object representation and it represents the values it generates when evaluating the AST,
- `repl`: read eval print loop

## How to use

To get started, clone the repository by following this command:

```sh
$ git clone git@github.com:FlorianWoelki/reflow.git
```

After that, you can feel free to directly run the shell in `cmd/shell.go` via:

```sh
go run cmd/shell.go
```

This will open up the reflow shell. From there, you can experiment with the programming language itself.
