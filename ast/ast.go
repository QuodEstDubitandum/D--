package ast

import "github.com/QuodEstDubitandum/D--/lexer"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	TokenLiteral() string
	statementNode()
}
	
type Expression interface {
	TokenLiteral() string
	expressionNode()
}

type Program struct {
	Statements []Statement
}
	
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type VarStatement struct {
	Token lexer.Token
	IdentifierName string
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

type ReturnStatement struct {
	Token lexer.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }