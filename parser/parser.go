package parser

import (
	"github.com/monkey-lang/monkey/scanner"
	"github.com/monkey-lang/monkey/token"
)

type Parser struct {
	l *scanner.Scanner

	curToken  token.Token
	peekToken token.Token
}
