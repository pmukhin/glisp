package parser

import (
	"reflect"
	"testing"

	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/token"
)

// do does the testwork
func do(t *testing.T, s string, e []ast.Statement) {
	scn := scanner.New(s)
	parser := New(scn)
	program, err := parser.Parse()

	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(e, program.Statements) {
		t.Errorf("expected %#v got %#v", e, program.Statements)
	}
}

func TestParser_Parse_ApplyToListOfStrings(t *testing.T) {
	do(t, `(print '("a" "b" "c"))`, []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.FunctionCall{
				Token: token.New(token.ParenOp, 0),
				Callee: &ast.IdentifierExpression{
					Token: token.New(token.Identifier, 1, "print"),
					Value: "print",
				},
				Args: []ast.Expression{
					&ast.ListExpression{
						Token: token.New(token.SingleQuote, 7),
						Elements: []ast.Expression{
							&ast.StringExpression{
								Token: token.New(token.String, 9, "a"),
								Value: "a",
							},
							&ast.StringExpression{
								Token: token.New(token.String, 13, "b"),
								Value: "b",
							},
							&ast.StringExpression{
								Token: token.New(token.String, 17, "c"),
								Value: "c",
							},
						},
					},
				},
			},
		},
	})
}

func TestParser_Parse_List(t *testing.T) {
	do(t, `(append '(1 2) 3)`, []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.FunctionCall{
				Token: token.New(token.ParenOp, 0),
				Callee: &ast.IdentifierExpression{
					Token: token.New(token.Identifier, 1, "append"),
					Value: "append",
				},
				Args: []ast.Expression{
					&ast.ListExpression{
						Token: token.New(token.SingleQuote, 8),
						Elements: []ast.Expression{
							&ast.IntegerExpression{
								Token: token.New(token.Integer, 10, "1"),
								Value: 1,
							},
							&ast.IntegerExpression{
								Token: token.New(token.Integer, 12, "2"),
								Value: 2,
							},
						},
					},
					&ast.IntegerExpression{
						Token: token.New(token.Integer, 15, "3"),
						Value: 3,
					},
				},
			},
		},
	})
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
