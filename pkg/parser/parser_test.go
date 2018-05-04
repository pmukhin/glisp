package parser

import (
	"reflect"
	"testing"

	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/token"
)

func TestParser_Parse1(t *testing.T) {
	source := `(* (+ 2 5) (– 7 (/ 21 7)))`
	scn := scanner.New(source)
	parser := New(scn)
	program, err := parser.Parse()

	if err != nil {
		t.Error(err)
	}

	expected := []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.FunctionCall{
				Token: token.New(token.ParenOp, 0),
				Callee: &ast.IdentifierExpression{
					Token: token.New(token.Identifier, 1, "*"),
					Value: "*",
				},
				Args: []ast.Expression{
					&ast.FunctionCall{
						Token: token.New(token.ParenOp, 3),
						Callee: &ast.IdentifierExpression{
							Token: token.New(token.Identifier, 4, "+"),
							Value: "+",
						},
						Args: []ast.Expression{
							&ast.IntegerExpression{
								Token: token.New(token.Integer, 6, "2"),
								Value: 2,
							},
							&ast.IntegerExpression{
								Token: token.New(token.Integer, 8, "5"),
								Value: 5,
							},
						},
					},
					&ast.FunctionCall{
						Token: token.New(token.ParenOp, 11),
						Callee: &ast.IdentifierExpression{
							Token: token.New(token.Identifier, 12, "-"),
							Value: "-",
						},
						Args: []ast.Expression{
							&ast.IntegerExpression{
								Token: token.New(token.Integer, 14, "7"),
								Value: 7,
							},
							&ast.FunctionCall{
								Token: token.New(token.ParenOp, 16),
								Callee: &ast.IdentifierExpression{
									Token: token.New(token.Identifier, 17, "/"),
									Value: "/",
								},
								Args: []ast.Expression{
									&ast.IntegerExpression{
										Token: token.New(token.Integer, 19, "21"),
										Value: 21,
									},
									&ast.IntegerExpression{
										Token: token.New(token.Integer, 22, "7"),
										Value: 7,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, program) {
		t.Errorf("expected %#v got %#v", expected, program)
	}
}
