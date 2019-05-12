package scanner

import (
	"github.com/monkey-lang/monkey/token"
)

type scanner struct {
	src string
}

func New(src string) *scanner {
	s := &scanner{src}
	return s
}

func (s *scanner) NextToken() token.Token {
	t := token.New(token.EOF, "")
	return t
}
