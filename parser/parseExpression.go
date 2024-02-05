package parser

import (
	"fmt"
	"strconv"

	"github.com/QuodEstDubitandum/D--/ast"
	"github.com/QuodEstDubitandum/D--/lexer"
)

func (p *Parser) ParseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}
	stmt.Value = p.parseExpression(LOWEST)

	// if p.IsNextToken(lexer.NEWLINE) {
	// 	p.nextToken()
	// }

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.NoPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	leftExp := prefix()
	

	for p.evalToken.Type != lexer.NEWLINE && p.evalToken.Type != lexer.EOF{
		infix := p.infixParseFns[p.evalToken.Type]
		if infix == nil || precedence >= p.nextPrecedence() {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	if !p.NextTokenIsSpace(){
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token: p.currentToken,
		Operator: p.currentToken.Literal,
	}
	
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token: p.currentToken,
		Operator: p.currentToken.Literal,
		Left: left,
	}

	precedence := p.currentPrecedence()
	
	if !p.NextTokenIsSpace(){
		return nil
	}

	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}