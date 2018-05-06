package parser

import (
	"reflect"
	"testing"

	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/scanner"
	"github.com/pmukhin/glisp/pkg/token"
	"fmt"
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
		tokTrace := ast.Print(program)

		fmt.Println(tokTrace)
	}
}

func TestParser_Parse_MacroDefVar_WithOutComment(t *testing.T) {
	do(t, `(defvar int-list '(1 2))`, []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.DefVarExpression{
				Token: token.New(token.Identifier, 1, "defvar"),
				Name: &ast.IdentifierExpression{
					Token: token.New(token.Identifier, 8, "int-list"),
					Value: "int-list",
				},
				Value: &ast.ListExpression{
					Token: token.New(token.SingleQuote, 17),
					Elements: []ast.Expression{
						&ast.IntegerExpression{
							Token: token.New(token.Integer, 19, "1"),
							Value: 1,
						},
						&ast.IntegerExpression{
							Token: token.New(token.Integer, 21, "2"),
							Value: 2,
						},
					},
				},
				Comment: nil,
			},
		},
	})
}

func TestParser_Parse_MacroDefVar_WithComment(t *testing.T) {
	do(t, `(defvar int-list '(1 2) "a list of ints")`, []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.DefVarExpression{
				Token: token.New(token.Identifier, 1, "defvar"),
				Name: &ast.IdentifierExpression{
					Token: token.New(token.Identifier, 8, "int-list"),
					Value: "int-list",
				},
				Value: &ast.ListExpression{
					Token: token.New(token.SingleQuote, 17),
					Elements: []ast.Expression{
						&ast.IntegerExpression{
							Token: token.New(token.Integer, 19, "1"),
							Value: 1,
						},
						&ast.IntegerExpression{
							Token: token.New(token.Integer, 21, "2"),
							Value: 2,
						},
					},
				},
				Comment: &ast.StringExpression{
					Token: token.New(token.String, 24, "a list of ints"),
					Value: "a list of ints",
				},
			},
		},
	})
}

func TestParser_Parse_VectorOfStrings(t *testing.T) {
	do(t, `["a" "b" "c"]`, []ast.Statement{
		&ast.ExpressionStatement{
			Expression: &ast.VectorExpression{
				Token: token.New(token.BracketOp, 0),
				Elements: []ast.Expression{
					&ast.StringExpression{
						Token: token.New(token.String, 1, "a"),
						Value: "a",
					},
					&ast.StringExpression{
						Token: token.New(token.String, 5, "b"),
						Value: "b",
					},
					&ast.StringExpression{
						Token: token.New(token.String, 9, "c"),
						Value: "c",
					},
				},
			},
		},
	})
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
