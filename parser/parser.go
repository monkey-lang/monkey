package parser

import (
	"fmt"
	"github.com/monkey-lang/monkey/ast"
	"github.com/monkey-lang/monkey/scanner"
	"github.com/monkey-lang/monkey/token"
)

func New(input *scanner.Scanner) *Parser {
	p := &Parser{input: input, errors: []string{}}

	// Read two tokens, so curToken and peekToken are both set
	p.next()
	p.next()

	return p
}

type Parser struct {
	errors []string
	// TODO: input, peekToken and token can be refactored in its own module
	input     *scanner.Scanner
	peekToken token.Token
	curToken  token.Token // We only need to look one token ahead
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for ; p.curToken.Type != token.EOF; p.next() {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}
	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

/* Helpers */
func (p *Parser) next() {
	p.curToken = p.peekToken
	p.peekToken = p.input.NextToken()
}

/* Parsers */
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	p.next()

	// let [x] = 10;
	if p.curToken.Type != token.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("Expected next token to be '%s' got '%s'", token.IDENT, p.curToken))
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	p.next()

	// let x [=] 10;
	if p.curToken.Type != token.ASSIGN {
		p.errors = append(p.errors, fmt.Sprintf("Expected next token to be '%s' got '%s'", token.ASSIGN, p.curToken))
		return nil
	}
	p.next()

	// let x = [10];
	// TODO: replace this
	for p.curToken.Type != token.SEMICOLON {
		p.next()
	}

	return stmt
}
