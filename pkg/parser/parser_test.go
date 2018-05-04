package parser

import (
	"reflect"
	"testing"

	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/token"
)

func do(t *testing.T, s string, e []ast.Statement) {
	scn := scanner.New(s)
	parser := New(scn)
	program, err := parser.Parse()

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(e, program.Statements) {
		t.Errorf("expected %#v got %#v", e, program.Statements)
	}
}

func TestParser_Parse_FunctionCall(t *testing.T) {
	do(t, `(* 2 5)`, []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.FunctionCall{
				Token: token.New(token.ParenOp, 0),
				Callee: &ast.IdentifierExpression{
					Token: token.New(token.Identifier, 1, "*"),
					Value: "*",
				},
				Args: []ast.Expression{
					&ast.IntegerExpression{
						Token: token.New(token.Integer, 3, "2"),
						Value: 2,
					},
					&ast.IntegerExpression{
						Token: token.New(token.Integer, 5, "5"),
						Value: 5,
					},
				},
			},
		},
	})
}

func TestParser_Parse_RecFunctionCall(t *testing.T) {
	do(t, `(* 2 (- 5 1))`, []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.FunctionCall{
				Token: token.New(token.ParenOp, 0),
				Callee: &ast.IdentifierExpression{
					Token: token.New(token.Identifier, 1, "*"),
					Value: "*",
				},
				Args: []ast.Expression{
					&ast.IntegerExpression{
						Token: token.New(token.Integer, 3, "2"),
						Value: 2,
					},
					&ast.FunctionCall{
						Token: token.New(token.ParenOp, 5),
						Callee: &ast.IdentifierExpression{
							Token: token.New(token.Identifier, 6, "-"),
							Value: "-",
						},
						Args: []ast.Expression{
							&ast.IntegerExpression{
								Token: token.New(token.Integer, 8, "5"),
								Value: 5,
							},
							&ast.IntegerExpression{
								Token: token.New(token.Integer, 10, "1"),
								Value: 1,
							},
						},
					},
				},
			},
		},
	})
}
