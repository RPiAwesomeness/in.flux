package token

// TokenType defines a separate type that is specifically used by the parser for identifying tokens,
// it's just a string though
type TokenType string

// Token defines a type that the parser uses to track tokens, will be used to build the AST
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foo, bar, x, y, i
	INT   = "INT"   // 1, 23, 42, 991290

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "{"
	RBRACKET = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
