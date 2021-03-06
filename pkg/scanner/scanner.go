package scanner

import (
	"unicode"

	"github.com/pmukhin/glisp/pkg/token"
)

func isIdentifier(ch rune) bool {
	return unicode.IsLetter(ch) ||
		ch == '<' ||
		ch == '>' ||
		ch == '=' ||
		ch == '*' ||
		ch == '/' ||
		ch == '+' ||
		ch == '-'
}

type Scanner struct {
	src    []rune
	ch     rune
	offset int
}

func New(source string) *Scanner {
	s := new(Scanner)
	s.src = []rune(source)
	s.ch = -1
	s.offset = -1

	return s
}

func (s *Scanner) nextChar() {
	s.offset++
	if s.offset >= len(s.src) {
		s.ch = -1
	} else {
		s.ch = s.src[s.offset]
	}
}

func (s *Scanner) Next() token.Token {
	s.nextChar()
	s.skipWhitespace()

	tokType := token.Illegal
	switch s.ch {
	case -1:
		return token.New(token.EOF, s.offset, "")
	case '(':
		tokType = token.ParenOp
	case ')':
		tokType = token.ParenCl
	case '[':
		tokType = token.BracketOp
	case ']':
		tokType = token.BracketCl
	case '"':
		return s.scanString()
	case '\'':
		tokType = token.SingleQuote
	case ':':
		tokType = token.Colon
	default:
		switch true {
		case unicode.IsDigit(s.ch):
			return s.scanNumber()
		case isIdentifier(s.ch):
			return s.scanIdentifier()
		default:
			return token.New(token.Illegal, s.offset, string(s.ch))
		}
	}

	return token.New(tokType, s.offset)
}

func (s *Scanner) scanIdentifier() token.Token {
	pos := s.offset // preserve the position
	str := make([]rune, 0, 32)

	for isIdentifier(s.ch) {
		str = append(str, s.ch)
		s.nextChar()
	}
	s.un()

	return token.New(token.Identifier, pos, string(str))
}

func (s *Scanner) skipWhitespace() {
	for s.ch == '\n' || s.ch == ' ' || s.ch == '\r' || s.ch == '\t' {
		s.nextChar()
	}
}

func (s *Scanner) scanString() token.Token {
	pos := s.offset // preserve the position
	s.nextChar()    // eat `"`
	str := make([]rune, 0, 32)

	for s.ch != '"' {
		str = append(str, s.ch)
		s.nextChar()
	}

	return token.New(token.String, pos, string(str))
}

func (s *Scanner) scanNumber() token.Token {
	pos := s.offset // preserve offset
	typ := token.Integer
	str := make([]rune, 0, 8)

	for s.ch == '.' || unicode.IsDigit(s.ch) {
		if s.ch == '.' {
			if typ == token.Float {
				return token.New(token.Illegal, s.offset, string(s.ch))
			}
			typ = token.Float
		}

		str = append(str, s.ch)
		s.nextChar()
	}
	s.un()

	return token.New(typ, pos, string(str))
}

func (s *Scanner) peek() rune {
	s.nextChar()
	defer s.un()

	return s.ch
}

func (s *Scanner) un() {
	s.offset--
	s.ch = s.src[s.offset]
}
