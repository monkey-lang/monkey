package token

import (
	"fmt"
)

const (
	EOF TokenType = iota
)

var stringTypes = [...]string{
	EOF: "EOF",
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
