package parser

import "github.com/QuodEstDubitandum/D--/lexer"

const (
	_ int = iota
	LOWEST
	EQUALS // ==
	LESSGREATER // > or <
	SUM // +
	PRODUCT // *
	POWER_OF // **
	PREFIX // -X or !X
	CALL // myFunction(X)
)

// map tokenTypes to their precedence
var precedences = map[lexer.TokenType]int{
	lexer.EQUAL: EQUALS,
	lexer.NOT_EQUAL: EQUALS,
	lexer.SMALLER: LESSGREATER,
	lexer.GREATER: LESSGREATER,
	lexer.GOR_EQUAL: LESSGREATER,
	lexer.SOR_EQUAL: LESSGREATER,
	lexer.PLUS: SUM,
	lexer.MINUS: SUM,
	lexer.SLASH: PRODUCT,
	lexer.ASTERISK: PRODUCT,
	lexer.DOUBLE_ASTERISK: POWER_OF,
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) nextPrecedence() int {
	if p, ok := precedences[p.evalToken.Type]; ok {
		return p
	}
	return LOWEST
}