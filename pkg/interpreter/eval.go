package interpreter

import (
	"fmt"

	"github.com/pmukhin/glisp/pkg/ast"
	"github.com/pmukhin/glisp/pkg/object"
)

type evaluatorFunc func(node ast.Node, ctx object.Context) (object.Object, error)

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
	typeToEvaluatorFunc[ast.VectorExpr] = evalVector
	typeToEvaluatorFunc[ast.DefVarExpr] = evalDefVar
	typeToEvaluatorFunc[ast.Expr] = evalExpr
}

// evalDefVar defines a variable in a given context
func evalDefVar(node ast.Node, ctx object.Context) (object.Object, error) {
	defVarExpr := node.(*ast.DefVarExpression)
	value, err := Eval(defVarExpr.Value, ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.Set(defVarExpr.Name.Value, value)
	if err != nil {
		return nil, nil
	}

	return nil, nil
}

// evalString ...
func evalString(node ast.Node, ctx object.Context) (object.Object, error) {
	astStrStmt := node.(*ast.StringExpression)
	return &object.String{Value: astStrStmt.Value}, nil
}

// evalList ...
func evalList(node ast.Node, ctx object.Context) (object.Object, error) {
	listStmt := node.(*ast.ListExpression)
	list := &object.List{Elements: make([]object.Object, 0, 32)}
	for _, astElem := range listStmt.Elements {
		oElem, err := Eval(astElem, ctx)
		if err != nil {
			return nil, err
		}
		list.Elements = append(list.Elements, oElem)
	}
	return list, nil
}

// evalVector ...
func evalVector(node ast.Node, ctx object.Context) (object.Object, error) {
	listStmt := node.(*ast.VectorExpression)
	list := &object.Vector{Elements: make([]object.Object, 0, 32)}

	var fType object.Type = -1
	for _, astElem := range listStmt.Elements {
		oElem, err := Eval(astElem, ctx)
		if err != nil {
			return nil, err
		}
		if fType == -1 {
			fType = oElem.Type()
		} else {
			if fType != oElem.Type() {
				return nil, fmt.Errorf("vectors contain only values "+
					"of the same type: %s is expected, %s given", fType, oElem.Type())
			}
		}

		list.Elements = append(list.Elements, oElem)
	}
	return list, nil
}

// evalRune ...
func evalRune(node ast.Node, ctx object.Context) (object.Object, error) {
	astRuneStmt := node.(*ast.RuneExpression)
	return &object.Rune{Value: astRuneStmt.Value}, nil
}

// evalExpr ...
func evalExpr(node ast.Node, ctx object.Context) (object.Object, error) {
	astExprStmt := node.(*ast.ExpressionStatement)
	return Eval(astExprStmt.Expression, ctx)
}

// evalInt ...
func evalInt(node ast.Node, ctx object.Context) (object.Object, error) {
	astInt := node.(*ast.IntegerExpression)
	return &object.Int{Value: astInt.Value}, nil
}

// evalFloat ...
func evalFloat(node ast.Node, ctx object.Context) (object.Object, error) {
	astFloat := node.(*ast.FloatExpression)
	return &object.Float{Value: astFloat.Value}, nil
}

// Eval ...
func Eval(n ast.Node, ctx object.Context) (object.Object, error) {
	evaluator, ok := typeToEvaluatorFunc[n.Type()]
	if !ok {
		return nil, fmt.Errorf("can not evaluate %s", n.Type())
	}
	return evaluator(n, ctx)
}

// evalProgram ...
func evalProgram(node ast.Node, ctx object.Context) (object.Object, error) {
	program := node.(*ast.Program)

	var lastVal object.Object = nil
	for _, statement := range program.Statements {
		val, err := Eval(statement, ctx)
		if err != nil {
			return nil, err
		}
		lastVal = val
	}

	return lastVal, nil
}

func evalFunctionCall(node ast.Node, ctx object.Context) (object.Object, error) {
	fc := node.(*ast.FunctionCall)
	fName := fc.Callee.(*ast.IdentifierExpression).Value
	fun, ok := internalFunctionTable[fName]

	if !ok {
		return nil, fmt.Errorf("function `%s` is not defined", fName)
	}

	args := make([]object.Object, len(fc.Args))
	for i, rawArg := range fc.Args {
		objArg, err := Eval(rawArg, ctx)
		if err != nil {
			return nil, err
		}
		args[i] = objArg
	}

	return fun(args...)
}
