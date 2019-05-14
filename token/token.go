package token

import (
	"fmt"
)

const (
	EOF TokenType = iota
	ILLEGAL
	ASSIGN
	ADD
	MINUS
	ASTERISK
	BANG
	RPAREN
	LPAREN
	RBRACE
	LBRACE
	LT
	GT
	EQ
	NOT_EQ
	SLASH
	COMMA
	SEMICOLON
	IDENT
	INT
)

var stringTypes = [...]string{
	EOF:       "EOF",
	ILLEGAL:   "ILLEGAL",
	ASSIGN:    "ASSIGN",
	ADD:       "ADD",
	MINUS:     "MINUS",
	ASTERISK:  "ASTERISK",
	BANG:      "BANG",
	RPAREN:    "RPAREN",
	LPAREN:    "LPAREN",
	RBRACE:    "RBRACE",
	LBRACE:    "LBRACE",
	LT:        "LT",
	GT:        "GT",
	EQ:        "EQ",
	NOT_EQ:    "NOT_EQ",
	SLASH:     "SLASH",
	COMMA:     "COMMA",
	SEMICOLON: "SEMICOLON",
	IDENT:     "IDENT",
	INT:       "INT",
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
