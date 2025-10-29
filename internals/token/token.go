package Token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // variable/function names
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	DIVIDE = "/"
	MULT   = "*"
	POW    = "^"

	EQUAL = "=="
	LD    = "<"
	GD    = ">"
	LDE   = "<="
	GDE   = ">="
	NOT   = "!"
	NOTE  = "!="

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACK    = "["
	RBRACK    = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	THEN     = "THEN"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var KeywordsTable = map[string]TokenType{
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"then":   THEN,
	"else":   ELSE,
	"return": RETURN,
}

// user defined identifies
var IdentTable = map[string]TokenType{}

var FunctionTable = map[string]TokenType{
	"print": FUNCTION,
}
