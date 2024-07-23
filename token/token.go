package token

type TokenType string

// TODO:Usually token contains file name and line so incase there is any error it is easier to represent
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
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

	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var Symbols = map[string]TokenType{
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  MINUS,
	"!":  BANG,
	"*":  ASTERISK,
	"/":  SLASH,
	"<":  LT,
	">":  GT,
	",":  COMMA,
	";":  SEMICOLON,
	"(":  LPAREN,
	")":  RPAREN,
	"{":  LBRACE,
	"}":  RBRACE,
	"==": EQ,
	"!=": NOT_EQ,
}

var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
