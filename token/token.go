package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // All our identifiers
	INT   = "INT"   // Simple numbers

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COLON     = ":"

	// Keywords
	FUNCTION = "FN"
	LET      = "LET"
)

type TokenType string // Maybe change to byte

type Token struct {
	Type    TokenType
	Literal string
}
