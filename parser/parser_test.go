package parser

import (
	"github.com/monkey-lang/monkey/ast"
	"github.com/monkey-lang/monkey/scanner"
	"testing"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := scanner.New(input)
	p := New(l)
	checkParseErrors(t, p)
	program := p.Parse()
	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statements, got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}

}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := scanner.New(input)
	p := New(l)
	checkParseErrors(t, p)
	program := p.Parse()
	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statements, got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func TestLetStatement(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	l := scanner.New(input)
	p := New(l)
	checkParseErrors(t, p)
	program := p.Parse()
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
			t.Errorf("Expected *ast.LetStatement got %T", stmt)
		}
		if letStatement.Name.TokenLiteral() != ti.expectedIdentifier {
			t.Errorf("Expected identifier to be '%s' got '%s'", ti.expectedIdentifier, letStatement.TokenLiteral())
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
		return 5;
		return 10;
		return add(15);
	`
	l := scanner.New(input)
	p := New(l)
	checkParseErrors(t, p)
	program := p.Parse()
	if program == nil {
		t.Fatalf("Parse returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got %d", len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected *ast.ReturnStatement got %T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("Expected literal to be 'return' got '%s'", returnStmt.TokenLiteral())
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) > 0 {
		t.Errorf("Parser has %d errors", len(errors))
		for _, err := range errors {
			t.Errorf("Parser error: %q", err)
		}
		t.FailNow()
	}
}
