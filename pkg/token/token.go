package token

// Type is type of a single token
type Type int8

const (
	EOF         Type = iota
	Illegal
	ParenOp
	ParenCl
	BracketOp
	BracketCl
	SingleQuote
	Colon
	Identifier
	Float
	Integer
	Rune
	String
)

var type2name = map[Type]string{
	EOF:         "EOF",
	Illegal:     "Illegal",
	ParenOp:     "ParenOp<(>",
	ParenCl:     "ParenCl<)>",
	BracketOp:   "BracketOp<[>",
	BracketCl:   "BracketCl<]>",
	SingleQuote: "SingleQuote<'>",
	Colon:       "Colon<:>",
	Identifier:  "Identifier",
	Float:       "Float",
	Integer:     "Integer",
	Rune:        "Rune",
	String:      "String",
}

func (t Type) String() string {
	return type2name[t]
}

var defaultLiteral = map[Type]string{
	ParenOp:     "(",
	ParenCl:     ")",
	BracketOp:   "[",
	BracketCl:   "]",
	Colon:       ":",
	SingleQuote: "'",
}

// Token represents a single token both terminals and non-terminals
type Token struct {
	Type    Type
	Literal string
	Pos     int
}

// New constructs a token
func New(typ Type, pos int, lit ...string) Token {
	defOrLit, ok := defaultLiteral[typ]
	if len(lit) == 0 && !ok {
		panic("non-passing a literal for a token with no default literal")
	}
	if !ok {
		defOrLit = lit[0]
	}

	return Token{
		Type:    typ,
		Pos:     pos,
		Literal: defOrLit,
	}
}
