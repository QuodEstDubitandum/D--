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

type Identifier struct {
	Token lexer.Token
	Value string
}
	
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ExpressionStatement struct{
	Token lexer.Token
	Value Expression
}

func (vs *ExpressionStatement) statementNode() {}
func (vs *ExpressionStatement) TokenLiteral() string { return vs.Token.Literal }


type VarStatement struct {
	Token lexer.Token
	Name *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }


type ReturnStatement struct {
	Token lexer.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type IntegerLiteral struct {
	Token lexer.Token
	Value int64
}
	
func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

type PrefixExpression struct {
	Token lexer.Token
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

type InfixExpression struct {
	Token lexer.Token
	Operator string
	Left Expression
	Right Expression
}

func (pe *InfixExpression) expressionNode() {}
func (pe *InfixExpression) TokenLiteral() string { return pe.Token.Literal }