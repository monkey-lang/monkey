package token

import (
	"fmt"
)

const (
	EOF TokenType = iota
	ILLEGAL
	ASSIGN
	ADD
	RPAREN
	LPAREN
	RBRACE
	LBRACE
	COMMA
	SEMICOLON
)

var stringTypes = [...]string{
	EOF:       "EOF",
	ILLEGAL:   "ILLEGAL",
	ASSIGN:    "ASSIGN",
	ADD:       "ADD",
	RPAREN:    "RPAREN",
	LPAREN:    "LPAREN",
	RBRACE:    "RBRACE",
	LBRACE:    "LBRACE",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
}

/* Token Types */
type TokenType int

func (typ TokenType) String() string {
	return stringTypes[typ]
}

/* Token */
func New(typ TokenType, lit string) Token {
	return Token{typ, lit}
}

type Token struct {
	Type    TokenType
	Literal string
}

func (tok Token) String() string {
	return fmt.Sprintf("(%s, %s)", stringTypes[tok.Type], tok.Literal)
}
