package token

import (
	"fmt"
)

type TokenType int

func (typ TokenType) String() string {
	return tokens[typ]
}

type TokenLiteral string

type Token struct {
	Type    TokenType
	Literal TokenLiteral
}

func (tok Token) String() string {
	return fmt.Sprintf("(%s, %s)", tokens[tok.Type], tok.Literal)
}

const (
	EOF TokenType = iota
	ASSIGN
	ADD
	LPAREN
	LBRACE
	RPAREN
	RBRACE
	COMMA
	SEMICOLON
)

var tokens = [...]string{
	EOF:       "EOF",
	ASSIGN:    "ASSIGN",
	ADD:       "ADD",
	LPAREN:    "LPAREN",
	LBRACE:    "LBRACE",
	RPAREN:    "RPAREN",
	RBRACE:    "RBRACE",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
}

func New(typ TokenType, lit TokenLiteral) Token {
	return Token{typ, lit}
}
