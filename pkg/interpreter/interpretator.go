package interpreter

import (
	"github.com/pmukhin/glisp/pkg/ast"
	"fmt"
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
	typeToEvaluatorFunc[ast.Expr] = evalExpr
}

// evalExpr ...
func evalExpr(node ast.Node) (object.Object, error) {
	astExprStmt := node.(*ast.ExpressionStatement)
	return Interpret(astExprStmt.Expression)
}

// evalInt ...
func evalInt(node ast.Node) (object.Object, error) {
	astInt := node.(*ast.IntegerExpression)
	return &object.Int{Value: astInt.Value}, nil
}

func Interpret(n ast.Node) (object.Object, error) {
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
		val, err := Interpret(statement)
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
		objArg, err := Interpret(rawArg)
		if err != nil {
			return nil, err
		}
		args[i] = objArg
	}

	return fun(args...)
}
