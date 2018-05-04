package object

import "fmt"

type Type int8

const (
	TInt      Type = iota
	TFunction
	TString
	TRune
	TFloat
)

var type2str = map[Type]string{
	TInt:      "TInt",
	TFunction: "TFunction",
	TString:   "TString",
	TRune:     "TRune",
	TFloat:    "TFloat",
}

func (t Type) String() string {
	return type2str[t]
}

type Object interface {
	String() string
	Type() Type
}

type Float struct {
	Value float64
}

func (f Float) String() string {
	return fmt.Sprintf("%f", f.Value)
}

func (f Float) Type() Type {
	return TFloat
}

type Int struct {
	Value int64
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.Value)
}

func (Int) Type() Type {
	return TInt
}

// String ...
type String struct {
	Value string
}

func (s String) String() string {
	return fmt.Sprintf("%s", s.Value)
}

func (String) Type() Type {
	return TString
}

// Rune ...
type Rune struct {
	Value rune
}

func (r Rune) String() string {
	return fmt.Sprintf("%s", string(r.Value))
}

func (Rune) Type() Type {
	return TRune
}
