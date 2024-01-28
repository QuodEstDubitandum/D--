package lexer

type TokenType byte

type Token struct {
	Type TokenType
	Literal string
	Err string
}

type Error struct {
	File string
	Line int
	Message string
}

var keywords = map[string]TokenType{
	"f": FUNCTION,
	"const": CONST,
	"var": VAR,
	"for": FOR,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"break": BREAK,
}

const (
	ILLEGAL TokenType = iota
	EOF
	ID	

	// Primitives
	INT
	FLOAT

	// Brackets
	LRBRKT
	RRBKRT
	LCBRKT
	RCBRKT
	LSBRKT
	RSBRKT

	// Operators
	ASSIGN
	GREATER
	SMALLER
	NOT_EQUAL
	EQUAL
	GOR_EQUAL
	SOR_EQUAL
	MAPS_TO 
	PLUS 
	MINUS
	ASTERISK
	PLUS_EQUAL
	MINUS_EQUAL
	DOUBLE_ASTERISK
	SLASH 
	PERCENT
	DOLLAR 
	EXCLAMATION 
	QUESTION 
	AMPERSAND 
	PIPE 
	COMMA 
	SEMICOLON 
	DOT 
	DOUBLEDOT 

	// Keywords
	FUNCTION 
	CONST 
	VAR 
	FOR
	IF
	ELSE
	RETURN
	BREAK

	// SPACES
	SPACE
	TAB
	NEWLINE
)