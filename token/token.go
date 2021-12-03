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

	// Data Types
	INT_TYPE = "INT_TYPE"

	// Keywords
	FUNCTION = "FN"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

var dataTypes = map[string]TokenType{
	"int": INT_TYPE,
}

func LookupIdentfier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	if tok, ok := dataTypes[identifier]; ok {
		return tok
	}

	return IDENT
}

type TokenType string // Maybe change to byte

type Token struct {
	Type    TokenType
	Literal string
}
