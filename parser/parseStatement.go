package parser

import (
	"github.com/QuodEstDubitandum/D--/ast"
	"github.com/QuodEstDubitandum/D--/lexer"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case lexer.VAR:
		return p.parseVarStatement()
	default:
		return nil
	}
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.currentToken}

	if !p.isNextToken(lexer.SPACE) {
		return nil
	}

	if !p.isNextToken(lexer.ID) {
		return nil
	}

	stmt.IdentifierName = p.currentToken.Literal

	if !p.isNextToken(lexer.SPACE) {
		return nil
	}

	if !p.isNextToken(lexer.ASSIGN) {
		return nil
	}

	if !p.isNextToken(lexer.SPACE) {
		return nil
	}

	if !p.isNextToken(lexer.INT) {
		return nil
	}

	for p.evalToken.Type != lexer.NEWLINE && p.evalToken.Type != lexer.EOF {
		if p.evalToken.Type == lexer.SPACE || p.evalToken.Type == lexer.TAB {
			p.nextToken()
		} else {
			return nil
		}
	}

	return stmt
}

func (p *Parser) isNextToken(tokenType lexer.TokenType) bool {
	if p.evalToken.Type != tokenType {
		return false
	}
	p.nextToken()
	return true
}