package parser

import (
	"github.com/QuodEstDubitandum/D--/ast"
	"github.com/QuodEstDubitandum/D--/lexer"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case lexer.VAR:
		return p.parseVarStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	case lexer.NEWLINE:
		p.nextToken()
		return p.parseStatement()
	default:
		return p.ParseExpressionStatement()
	}
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.currentToken}

	if !p.IsNextToken(lexer.SPACE) {
		return nil
	}

	if !p.IsNextToken(lexer.ID) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal} 

	if !p.IsNextToken(lexer.SPACE) {
		return nil
	}

	if !p.IsNextToken(lexer.ASSIGN) {
		return nil
	}

	if !p.IsNextToken(lexer.SPACE) {
		return nil
	}

	// TODO: also parse other primary types
	if !p.IsNextToken(lexer.INT) {
		return nil
	}


	for p.evalToken.Type != lexer.NEWLINE && p.evalToken.Type != lexer.EOF {
		if p.evalToken.Type == lexer.SPACE || p.evalToken.Type == lexer.TAB {
			p.nextToken()
		} else {
			p.NextTokenError(lexer.NEWLINE, p.evalToken.Type)
			return nil
		}
	}
	p.nextToken()

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	if !p.IsNextToken(lexer.SPACE) {
		return nil
	}

	// TODO: also parse other primary types
	if !p.IsNextToken(lexer.INT) {
		return nil
	}

	for p.evalToken.Type != lexer.NEWLINE && p.evalToken.Type != lexer.EOF {
		if p.evalToken.Type == lexer.SPACE || p.evalToken.Type == lexer.TAB {
			p.nextToken()
		} else {
			p.NextTokenError(lexer.NEWLINE, p.evalToken.Type)
			return nil
		}
	}
	p.nextToken()

	return stmt
}