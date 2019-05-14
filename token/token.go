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
	INT
	// Identifiers
	IDENT
	LET
	FUNCTION
	RETURN
	IF
	ELSE
	TRUE
	FALSE
)

var stringTypes = [...]string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	// Symbols
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
	// Numbers
	INT: "INT",
	// Identifiers
	IDENT:    "IDENT",
	LET:      "LET",
	FUNCTION: "FUNCTION",
	RETURN:   "RETURN",
	IF:       "IF",
	ELSE:     "ELSE",
	TRUE:     "TRUE",
	FALSE:    "FALSE",
}

var identTypes = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
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

/* utils */
func IdentLookup(ident string) TokenType {
	typ, found := identTypes[ident]
	if !found {
		return IDENT
	}
	return typ
}
