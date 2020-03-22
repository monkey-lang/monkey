package parser

import (
	"fmt"
	"github.com/monkey-lang/monkey/ast"
	"github.com/monkey-lang/monkey/scanner"
	"github.com/monkey-lang/monkey/token"
)

const (
	_ = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

func New(input *scanner.Scanner) *Parser {
	p := &Parser{input: input, errors: []string{}}

	// Read two tokens, so curToken and peekToken are both set
	p.next()
	p.next()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	return p
}

type prefixParseFn func() ast.Expression
type infixParseFn func() ast.Expression

type Parser struct {
	errors []string
	// TODO: input, peekToken and token can be refactored in its own module
	input     *scanner.Scanner
	peekToken token.Token
	curToken  token.Token // We only need to look one token ahead

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
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
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) next() {
	p.curToken = p.peekToken
	p.peekToken = p.input.NextToken()
}

func (p *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}

/* Parsers */
// TODO: assertions with nil are failing, it might should be *ast.Statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// [let] x = 10;
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	// [return] 10;
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.next()

	// return [10];
	// TODO: replace this
	for p.curToken.Type != token.SEMICOLON {
		p.next()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.next()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
