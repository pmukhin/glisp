package interpreter

import (
	"testing"
	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/token"
	"github.com/pmukhin/glisp/pkg/object"
	"reflect"
)

func TestEval_ListAppend(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
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
		},
	}

	res, err := Eval(program)
	if err != nil {
		t.Error(err)
	}

	list := res.(*object.List)
	expectedElements := []object.Object{
		&object.Int{1}, &object.Int{2}, &object.Int{3},
	}

	if !reflect.DeepEqual(list.Elements, expectedElements) {
		t.Errorf("wrong elements in resulting list: %v vs %v", expectedElements, list.Elements)
	}
}

func TestEval_ArithmeticFunction(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
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
		},
	}

	res, err := Eval(program)
	if err != nil {
		t.Error(err)
	}

	iVal := res.(*object.Int)
	if iVal.Value != 10 {
		t.Errorf("expected 10, got %d", iVal.Value)
	}
}
