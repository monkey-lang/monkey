package scanner

import (
	"github.com/monkey-lang/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	src := ""
	tests := []token.Token{
		{token.EOF, ""},
	}

	s := New(src)
	for i, ti := range tests {
		tok := s.NextToken()
		if tok.Type != ti.Type {
			t.Fatalf("Tests[%d] - token type is wrong. Expected %q but got %q", i, ti.Type, tok.Type)
		}

		if tok.Literal != ti.Literal {
			t.Fatalf("Tests[%d] - token literal is wrong. Expected %q but got %q", i, ti.Literal, tok.Literal)
		}
	}
}
