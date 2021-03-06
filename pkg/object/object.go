package object

import (
	"fmt"
	"strings"
)

type Type int8

const (
	TInt      Type = iota
	TFunction
	TString
	TRune
	TFloat
	TBool
	TList
	TVector
)

var type2str = map[Type]string{
	TInt:      "TInt",
	TFunction: "TFunction",
	TString:   "TString",
	TRune:     "TRune",
	TFloat:    "TFloat",
	TBool:     "TBool",
	TList:     "TList",
	TVector:   "TVector",
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

// Bool ...
type Bool struct {
	Value bool
}

// String ...
func (b Bool) String() string {
	if b.Value {
		return "true"
	}
	return "false"
}

// Type ...
func (b Bool) Type() Type {
	return TBool
}

// List ...
type List struct {
	Elements []Object
}

// String ...
func (l List) String() string {
	strElements := make([]string, len(l.Elements))
	for i, el := range l.Elements {
		strElements[i] = el.String()
	}
	return "'(" + strings.Join(strElements, " ") + ")"
}

// Type ...
func (l List) Type() Type {
	return TList
}

// Vector ...
type Vector struct {
	Elements []Object
}

// String ...
func (v Vector) String() string {
	strElements := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		strElements[i] = el.String()
	}
	return "[" + strings.Join(strElements, " ") + "]"
}

// Type ...
func (Vector) Type() Type {
	return TVector
}
