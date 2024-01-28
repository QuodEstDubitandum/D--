package parser

import (
	"github.com/QuodEstDubitandum/D--/ast"
	"github.com/QuodEstDubitandum/D--/lexer"
)

type Parser struct {
	l *lexer.Lexer
	currentToken lexer.Token
	evalToken lexer.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
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

	
	