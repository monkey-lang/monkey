package scanner

import (
	"github.com/monkey-lang/monkey/token"
	"strings"
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
	go s.run()
	return s
}

type scanner struct {
	input  string           // the "source code"
	begin  int              // start endition of this item
	end    int              // current position of this item
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
	if s.end >= len(s.input) {
		s.width = 0
		return EOF
	}
	r, s.width = utf8.DecodeRuneInString(s.input[s.end:])
	s.end += s.width
	return r
}

// prev steps back one rune (it undoes what next did). It should only be called once after next
func (s *scanner) prev() {
	s.end -= s.width
}

// peek returns the next rune but does not consume
func (s *scanner) peek() rune {
	r := s.next()
	s.prev()
	return r
}

func (s *scanner) confirm() {
	s.begin = s.end
}

func (s *scanner) emit(typ token.TokenType) {
	s.tokens <- token.New(typ, s.input[s.begin:s.end])
	s.confirm()
}

/* scanners */
type stateFn func(s *scanner) stateFn

func scan(s *scanner) stateFn {
	r := s.peek()
	switch {
	case isEOF(r):
		return scanEOF
	case isSpace(r):
		return scanSpace
	case isSymbol(r):
		return scanSymbol
	case isLetter(r):
		return scanIdent
	case isNumber(r):
		return scanInt
	default:
		return scanIllegal
	}
}

func scanEOF(s *scanner) stateFn {
	s.emit(token.EOF)
	return nil
}

func scanSpace(s *scanner) stateFn {
	for r := s.peek(); isSpace(r); r = s.peek() {
		s.next()
	}
	s.confirm()
	return scan
}

func scanSymbol(s *scanner) stateFn {
	var typ token.TokenType
	switch s.next() {
	case '=':
		if s.peek() == '=' {
			s.next()
			typ = token.EQ
		} else {
			typ = token.ASSIGN
		}
	case '+':
		typ = token.ADD
	case '(':
		typ = token.LPAREN
	case ')':
		typ = token.RPAREN
	case '{':
		typ = token.LBRACE
	case '}':
		typ = token.RBRACE
	case ',':
		typ = token.COMMA
	case ';':
		typ = token.SEMICOLON
	case '!':
		if s.peek() == '=' {
			s.next()
			typ = token.NOT_EQ
		} else {
			typ = token.BANG
		}
	case '-':
		typ = token.MINUS
	case '/':
		typ = token.SLASH
	case '*':
		typ = token.ASTERISK
	case '<':
		typ = token.LT
	case '>':
		typ = token.GT
	default:
		return scanIllegal
	}
	s.emit(typ)
	return scan
}

func scanIdent(s *scanner) stateFn {
	for r := s.peek(); isLetter(r); r = s.peek() {
		s.next()
	}
	s.emit(token.IDENT)
	return scan
}

func scanInt(s *scanner) stateFn {
	for r := s.peek(); isNumber(r); r = s.peek() {
		s.next()
	}
	s.emit(token.INT)
	return scan
}

func scanIllegal(s *scanner) stateFn {
	s.next()
	s.emit(token.ILLEGAL)
	return nil
}

/* helpers */
func isEOF(r rune) bool {
	return r == EOF
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func isSymbol(r rune) bool {
	return strings.ContainsRune("=+(){},;!-/*<>", r)
}

func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}
