package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pmukhin/glisp/pkg/token"
)

type Type int8

const (
	FunCall     Type = iota
	ProgramExpr
	Expr
	IdentExpr
	StringExpr
	IntExpr
	FloatExpr
	RuneExpr
	ListExpr
)

var type2str = map[Type]string{
	FunCall:     "FunCall",
	ProgramExpr: "ProgramExpr",
	Expr:        "Expr",
	IdentExpr:   "IdentExpr",
	StringExpr:  "StringExpr",
	IntExpr:     "IntExpr",
	FloatExpr:   "FloatExpr",
	RuneExpr:    "RuneExpr",
	ListExpr:    "ListExpr",
}

func (t Type) String() string {
	return type2str[t]
}

type Node interface {
	Pos() int
	Type() Type
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// ExpressionStatement ...
type ExpressionStatement struct {
	Expression Expression
}

// Type ...
func (es ExpressionStatement) Type() Type {
	return Expr
}

// Pos ...
func (es ExpressionStatement) Pos() int {
	return es.Expression.Pos()
}

// String ...
func (es ExpressionStatement) String() string {
	return es.Expression.String()
}

func (ExpressionStatement) statementNode() {}

// FunctionCall is 99% of all language constructions
type FunctionCall struct {
	Token  token.Token
	Callee Expression
	Args   []Expression
}

// Type ...
func (fc FunctionCall) Type() Type {
	return FunCall
}

// String ...
func (fc FunctionCall) String() string {
	fp := "(" + fc.Callee.String()
	for _, c := range fc.Args {
		fp += " " + c.String()
	}
	return fp + ")"
}

// Pos ...
func (fc FunctionCall) Pos() int {
	return fc.Token.Pos
}

func (FunctionCall) expressionNode() {}

type IdentifierExpression struct {
	Token token.Token
	Value string
}

func (id IdentifierExpression) Pos() int { return id.Token.Pos }

func (id IdentifierExpression) Type() Type { return IdentExpr }

func (id IdentifierExpression) String() string {
	return id.Value
}

func (id IdentifierExpression) expressionNode() {}

// IntegerExpression ...
type IntegerExpression struct {
	Token token.Token
	Value int64
}

func (ie IntegerExpression) Pos() int {
	return ie.Token.Pos
}

func (ie IntegerExpression) Type() Type {
	return IntExpr
}

func (ie IntegerExpression) String() string {
	return strconv.FormatInt(ie.Value, 10)
}

func (ie IntegerExpression) expressionNode() {}

// FloatExpression ...
type FloatExpression struct {
	Token token.Token
	Value float64
}

// Pos ...
func (fe FloatExpression) Pos() int {
	return fe.Token.Pos
}

// Type ...
func (fe FloatExpression) Type() Type {
	return FloatExpr
}

// String ...
func (fe FloatExpression) String() string {
	return fmt.Sprintf("%f", fe.Value)
}

// expressionNode ...
func (fe FloatExpression) expressionNode() {}

// StringExpression ...
type StringExpression struct {
	Token token.Token
	Value string
}

func (se StringExpression) Pos() int {
	return se.Token.Pos
}

func (se StringExpression) Type() Type {
	return StringExpr
}

func (se StringExpression) String() string {
	return se.Value
}

func (se StringExpression) expressionNode() {}

// Program ...
type Program struct {
	Statements []Statement
}

// Pos ...
func (Program) Pos() int {
	return 0
}

// Type ...
func (Program) Type() Type {
	return ProgramExpr
}

type RuneExpression struct {
	Token token.Token
	Value rune
}

func (re RuneExpression) Pos() int {
	return re.Token.Pos
}

func (re RuneExpression) Type() Type {
	return RuneExpr
}

func (re RuneExpression) String() string {
	return string(re.Value)
}

func (re RuneExpression) expressionNode() {}

// ListExpression ...
type ListExpression struct {
	Token    token.Token
	Elements []Expression
}

// Pos ...
func (le ListExpression) Pos() int {
	return le.Token.Pos
}

// Type ...
func (le ListExpression) Type() Type {
	return ListExpr
}

// String ...
func (le ListExpression) String() string {
	strList := make([]string, len(le.Elements))
	for i, el := range le.Elements {
		strList[i] = el.String()
	}
	return "'(" + strings.Join(strList, " ") + ")"
}

// expressionNode ...
func (le ListExpression) expressionNode() {}

// String ...
func (p Program) String() string {
	stmts := make([]string, len(p.Statements))
	for i, s := range p.Statements {
		stmts[i] = s.String()
	}
	return strings.Join(stmts, "\n")
}

func (Program) expressionNode() {}
