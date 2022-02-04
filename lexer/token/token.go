package token

const (
	LBRACKET = "["
	RBRACKET = "]"

	STRING = "STRING"

	ILLEGAL = "ILLEGAL" // Token is not defined.
	EOF     = "EOF"

	IDENT = "IDENT" // All our identifiers like a, b, c.
	INT   = "INT"   // Simple numbers like 1, 2, 3.

	// Operators
	ASSIGN          = "="
	PLUS            = "+"
	MINUS           = "-"
	BANG            = "!"
	ASTERISK        = "*"
	SLASH           = "/"
	LT              = "<"
	GT              = ">"
	EQ              = "=="
	NOT_EQ          = "!="
	PLUS_ASSIGN     = "+="
	MINUS_ASSIGN    = "-="
	ASTERISK_ASSIGN = "*="
	SLASH_ASSIGN    = "/="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"while":  WHILE,
}

func LookupIdentfier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENT
}

type TokenType string // Maybe change to byte

type Token struct {
	Type    TokenType
	Literal string // Defines the content of the token.
}
