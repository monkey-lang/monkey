package parser

import (
	"github.com/monkey-lang/monkey/ast"
	"github.com/monkey-lang/monkey/scanner"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	l := scanner.New(input)
	p := New(l)
	program := p.Parse()
	errors := p.Errors()
	if len(errors) > 0 {
		t.Errorf("Parser has %d errors", len(errors))
		for _, err := range errors {
			t.Errorf("Parser error: %q", err)
		}
		t.FailNow()
	}
	if program == nil {
		t.Fatalf("Parse returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got %d", len(program.Statements))
	}
	tests := []struct{ expectedIdentifier string }{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, ti := range tests {
		stmt := program.Statements[i]
		if stmt.TokenLiteral() != "let" {
			t.Errorf("Expected TokenLiteral 'let', got %q", stmt.TokenLiteral())
		}
		letStatement, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Errorf("Expected *ast.LetStatement got=%T", stmt)
		}
		if letStatement.Name.Value != ti.expectedIdentifier {
			t.Errorf("Expected identifier to be '%s' got '%s'", ti.expectedIdentifier, letStatement.Name.Value)
		}
	}
}
