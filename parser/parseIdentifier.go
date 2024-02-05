package parser

import "github.com/QuodEstDubitandum/D--/ast"

func (p *Parser) parseIdentifier() ast.Expression {
	expression := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.NextTokenIsSpace(){
		return nil
	}
	
	return expression
}