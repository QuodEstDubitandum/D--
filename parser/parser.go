package parser

import (
	"github.com/QuodEstDubitandum/D--/ast"
	"github.com/QuodEstDubitandum/D--/lexer"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer
	currentToken lexer.Token
	evalToken lexer.Token
	errors []string

	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns map[lexer.TokenType]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}
	p.prefixParseFns = make(map[lexer.TokenType]prefixParseFn)
	p.infixParseFns = make(map[lexer.TokenType]infixParseFn)
	p.registerPrefix(lexer.ID, p.parseIdentifier)
	p.registerPrefix(lexer.INT, p.parseIntegerLiteral)
	p.registerPrefix(lexer.EXCLAMATION, p.parsePrefixExpression)
	p.registerPrefix(lexer.MINUS, p.parsePrefixExpression)
	p.registerInfix(lexer.PLUS, p.parseInfixExpression)
	p.registerInfix(lexer.MINUS, p.parseInfixExpression)
	p.registerInfix(lexer.SLASH, p.parseInfixExpression)
	p.registerInfix(lexer.ASTERISK, p.parseInfixExpression)
	p.registerInfix(lexer.EQUAL, p.parseInfixExpression)
	p.registerInfix(lexer.NOT_EQUAL, p.parseInfixExpression)
	p.registerInfix(lexer.SMALLER, p.parseInfixExpression)
	p.registerInfix(lexer.GREATER, p.parseInfixExpression)
	p.registerInfix(lexer.SOR_EQUAL, p.parseInfixExpression)
	p.registerInfix(lexer.GOR_EQUAL, p.parseInfixExpression)
	p.registerInfix(lexer.DOUBLE_ASTERISK, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.evalToken
	p.evalToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		
		// p.nextToken()
		if !p.advanceUntilNewLine() {
			break
		}
	}
	return program
}


func (p *Parser) IsNextToken(tokenType lexer.TokenType) bool {
	if p.evalToken.Type != tokenType {
		p.NextTokenError(tokenType, p.evalToken.Type)
		return false
	}
	p.nextToken()
	return true
}
	
func (p *Parser) NextTokenIsSpace() bool {
	if p.evalToken.Type == lexer.SPACE {
		p.nextToken()
		return true
	}

	if p.evalToken.Type == lexer.EOF || p.evalToken.Type == lexer.NEWLINE {
		return true
	}

	p.MissingSpaceError(p.evalToken.Type)
	return false
}

func (p *Parser) advanceUntilNewLine() bool {
	for p.evalToken.Type != lexer.NEWLINE && p.evalToken.Type != lexer.EOF {
		if p.evalToken.Type == lexer.SPACE || p.evalToken.Type == lexer.TAB {
			p.nextToken()
		} else {
			p.NextTokenError(lexer.NEWLINE, p.evalToken.Type)
			return false
		}
	}
	p.nextToken()
	return true
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
	
func (p *Parser) registerInfix(tokenType lexer.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}