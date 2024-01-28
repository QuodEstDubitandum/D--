package parser

import (
	"fmt"
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

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	fmt.Println(program)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
		len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	
	for i, tt := range tests {
		stmt := program.Statements[i]
		fmt.Println(stmt)
		if !testVarStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got=%q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got=%T", s)
		return false
	}

	if varStmt.IdentifierName != name {
		t.Errorf("varStmt.Name.Value not '%s'. got=%s", name, varStmt.IdentifierName)
		return false
	}

	return true
}