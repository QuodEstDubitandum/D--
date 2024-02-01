package parser

import (
	"fmt"

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
		p.nextToken()
	}
	return program
}

func (p *Parser) nextTokenError(expectToken lexer.TokenType, gotToken lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", lexer.ASCIIMap[expectToken], lexer.ASCIIMap[gotToken])
	p.errors = append(p.errors, msg)
}

func (p *Parser) IsNextToken(tokenType lexer.TokenType) bool {
	if p.evalToken.Type != tokenType {
		p.nextTokenError(tokenType, p.evalToken.Type)
		return false
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