package parser

import (
	"os"
	"testing"

	"github.com/QuodEstDubitandum/D--/ast"
	"github.com/QuodEstDubitandum/D--/lexer"
)

func TestVarStatements(t *testing.T) {
	file, err := os.Open("../test_files/var_statement.txt")
	if err != nil {
		t.Fatal(err)
	}
	l := lexer.NewLexer(file)

	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		varStatement, ok := stmt.(*ast.VarStatement)
		if !ok {
			t.Errorf("stmt not *ast.VarStatement. got=%T", stmt)
			continue
		}
		if varStatement.TokenLiteral() != "var" {
			t.Errorf("varStatement.TokenLiteral not 'var', got %q", varStatement.TokenLiteral())
		}
	}
}

func TestReturnStatements(t *testing.T) {
	file, err := os.Open("../test_files/return_statement.txt")
	if err != nil {
		t.Fatal(err)
	}
	l := lexer.NewLexer(file)

	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	}

	t.Errorf("parser encountered %d errors", len(p.errors))
	
	for _, msg := range p.errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func TestExpressions(t *testing.T) {
	file, err := os.Open("../test_files/expression_statement.txt")
	if err != nil {
		t.Fatal(err)
	}
	l := lexer.NewLexer(file)

	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 3, len(program.Statements))
	}

	for _, stmt := range program.Statements {
		stmt, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		_ , ok = stmt.Value.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.InfixExpression. got=%T", stmt.Value)
		}
	}
}

