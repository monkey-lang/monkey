package scanner

import (
	"github.com/monkey-lang/monkey/token"
	"unicode/utf8"
)

const (
	EOF rune = 0
)

func New(input string) *scanner {
	s := &scanner{
		input:  input,
		tokens: make(chan token.Token),
	}
	s.run()
	return s
}

type scanner struct {
	input  string           // the "source code"
	start  int              // start position of this item
	pos    int              // current position of this item
	width  int              // width of last rune read from input.
	tokens chan token.Token // channel of scanned tokens
}

func (s *scanner) NextToken() token.Token {
	return <-s.tokens
}

func (s *scanner) run() {
	for state := scan(s); state != nil; state = state(s) {
	}
	close(s.tokens)
}

// next returns the next rune in the string
func (s *scanner) next() rune {
	var r rune
	if s.pos > len(s.input) {
		s.width = 0
		return EOF
	}
	r, s.width = utf8.DecodeRuneInString(s.input[s.pos:])
	s.pos += s.width
	return r
}

// prev steps back one rune (it undoes what next did). It should only be called once after next
func (s *scanner) prev() {
	s.pos -= s.width
}

// peek returns the next rune but does not consume
func (s *scanner) peek() rune {
	r := s.next()
	s.prev()
	return r
}

func (s *scanner) confirm() {
	s.start = s.pos
}

func (s *scanner) emit(typ token.TokenType) {
	s.tokens <- token.New(typ, s.input[s.start:s.pos])
	s.confirm()
}

/* scanners */
type stateFn func(s *scanner) stateFn

func scan(s *scanner) stateFn {
	r := s.peek()
	switch {
	case r == EOF:
		return scanEOF
	default:
		return nil
	}
}

func scanEOF(s *scanner) stateFn {
	s.emit(token.EOF)
	return nil
}

/* helpers */
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}
