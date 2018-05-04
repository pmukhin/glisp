package ast

import (
	"github.com/pmukhin/glisp/pkg/token"
	"strconv"
)

type Type int8

const (
	FunCall    Type = iota
	Expr
	IdentExpr
	StringExpr
	IntExpr
	FloatExpr
	RuneExpr
)

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