package interpreter

import (
	"fmt"

	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/object"
)

type evaluatorFunc func(node ast.Node) (object.Object, error)

var typeToEvaluatorFunc map[ast.Type]evaluatorFunc

// register ast handlers
func init() {
	typeToEvaluatorFunc = make(map[ast.Type]evaluatorFunc)

	typeToEvaluatorFunc[ast.ProgramExpr] = evalProgram
	typeToEvaluatorFunc[ast.FunCall] = evalFunctionCall
	typeToEvaluatorFunc[ast.IntExpr] = evalInt
	typeToEvaluatorFunc[ast.FloatExpr] = evalFloat
	typeToEvaluatorFunc[ast.StringExpr] = evalString
	typeToEvaluatorFunc[ast.RuneExpr] = evalRune
	typeToEvaluatorFunc[ast.ListExpr] = evalList
	typeToEvaluatorFunc[ast.Expr] = evalExpr
}

// evalString ...
func evalString(node ast.Node) (object.Object, error) {
	astStrStmt := node.(*ast.StringExpression)
	return &object.String{Value: astStrStmt.Value}, nil
}

// evalList ...
func evalList(node ast.Node) (object.Object, error) {
	listStmt := node.(*ast.ListExpression)
	list := &object.List{Elements: make([]object.Object, 0, 32)}
	for _, astElem := range listStmt.Elements {
		oElem, err := Eval(astElem)
		if err != nil {
			return nil, err
		}
		list.Elements = append(list.Elements, oElem)
	}
	return list, nil
}

// evalRune ...
func evalRune(node ast.Node) (object.Object, error) {
	astRuneStmt := node.(*ast.RuneExpression)
	return &object.Rune{Value: astRuneStmt.Value}, nil
}

// evalExpr ...
func evalExpr(node ast.Node) (object.Object, error) {
	astExprStmt := node.(*ast.ExpressionStatement)
	return Eval(astExprStmt.Expression)
}

// evalInt ...
func evalInt(node ast.Node) (object.Object, error) {
	astInt := node.(*ast.IntegerExpression)
	return &object.Int{Value: astInt.Value}, nil
}

// evalFloat ...
func evalFloat(node ast.Node) (object.Object, error) {
	astFloat := node.(*ast.FloatExpression)
	return &object.Float{Value: astFloat.Value}, nil
}

// Eval ...
func Eval(n ast.Node) (object.Object, error) {
	evaluator, ok := typeToEvaluatorFunc[n.Type()]
	if !ok {
		return nil, fmt.Errorf("can not evaluate %s", n.Type())
	}
	return evaluator(n)
}

// evalProgram ...
func evalProgram(node ast.Node) (object.Object, error) {
	program := node.(*ast.Program)

	var lastVal object.Object = nil
	for _, statement := range program.Statements {
		val, err := Eval(statement)
		if err != nil {
			return nil, err
		}
		lastVal = val
	}

	return lastVal, nil
}

func evalFunctionCall(node ast.Node) (object.Object, error) {
	fc := node.(*ast.FunctionCall)
	fName := fc.Callee.(*ast.IdentifierExpression).Value
	fun, ok := internalFunctionTable[fName]

	if !ok {
		return nil, fmt.Errorf("function `%s` is not defined", fName)
	}

	args := make([]object.Object, len(fc.Args))
	for i, rawArg := range fc.Args {
		objArg, err := Eval(rawArg)
		if err != nil {
			return nil, err
		}
		args[i] = objArg
	}

	return fun(args...)
}
