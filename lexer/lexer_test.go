package lexer

import (
	"fmt"
	"os"
	"testing"
)

func TestConstStatement(t *testing.T){
	file, err := os.Open("../test_files/const_statement.txt")
	if err != nil {
		t.Fatal(err)
	}

	l := NewLexer(file)

	tests := []Token{
		{CONST, "const", ""},
		{SPACE, " ", ""},
		{ID, "five", ""},
		{SPACE, " ", ""},
		{ASSIGN, "=", ""},
		{SPACE, " ", ""},
		{INT, "5", ""},
		{EOF, "", ""},
	}
	

	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d", i, test.Type, tok.Type)
		}
		if tok.Literal != test.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",i, test.Literal, tok.Literal)
		}
	}
}

func TestVarStatements(t *testing.T){
	file, err := os.Open("../test_files/var_statement.txt")
	if err != nil {
		t.Fatal(err)
	}

	l := NewLexer(file)

	tests := []Token{
		{VAR, "var", ""},
		{SPACE, " ", ""},
		{ID, "x", ""},
		{SPACE, " ", ""},
		{ASSIGN, "=", ""},
		{SPACE, " ", ""},
		{INT, "5", ""},
		{NEWLINE, "\n", ""},
		{VAR, "var", ""},
		{SPACE, " ", ""},
		{ID, "y", ""},
		{SPACE, " ", ""},
		{ASSIGN, "=", ""},
		{SPACE, " ", ""},
		{INT, "10", ""},
		{NEWLINE, "\n", ""},
		{VAR, "var", ""},
		{SPACE, " ", ""},
		{ID, "foobar", ""},
		{SPACE, " ", ""},
		{ASSIGN, "=", ""},
		{SPACE, " ", ""},
		{INT, "69420", ""},
		{EOF, "", ""},
	}

	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.Type {
			fmt.Println(tok.Type, tok.Literal)
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d", i, test.Type, tok.Type)
		}
		if tok.Literal != test.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",i, test.Literal, tok.Literal)
		}
	}
}