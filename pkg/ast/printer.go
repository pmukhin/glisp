package ast

import (
	"strings"
	"fmt"
)

var astType2printer map[Type]func(node Node) string

func init() {
	astType2printer = map[Type]func(node Node) string{
		ProgramExpr: printProgram,
		Expr:        printExpr,
		IntExpr:     printInt,
		StringExpr:  printStr,
		IdentExpr:   printIdent,
		FloatExpr:   printFloat,
		DefVarExpr:  printDefVar,
		ListExpr:    printList,
	}
}

func printDefVar(node Node) string {
	defVar := node.(*DefVarExpression)
	return fmt.Sprintf("<ast.DefVarExpr> pos: %d value: %s comment: %s", defVar.Pos(),
		Print(defVar.Value), defVar.Comment.String())
}

func printList(node Node) string {
	listExpr := node.(*ListExpression)
	values := make([]string, 0, len(listExpr.Elements))
	for _, el := range listExpr.Elements {
		values = append(values, Print(el))
	}

	return fmt.Sprintf("<ast.ListExpr pos: %d value: [%s]>", listExpr.Pos(),
		strings.Join(values, ", "))
}

func printIdent(node Node) string {
	astStr := node.(*IdentifierExpression)
	return fmt.Sprintf("<ast.IdentExpr pos: %d value: %s>", astStr.Pos(), astStr.Value)
}

func printStr(node Node) string {
	astStr := node.(*StringExpression)
	return fmt.Sprintf("<ast.StringExpr pos: %d value: %s>", astStr.Pos(), astStr.Value)
}

func printFloat(node Node) string {
	astFloat := node.(*FloatExpression)
	return fmt.Sprintf("<ast.FloatExpr pos: %d value: %f>", astFloat.Pos(), astFloat.Value)
}

func printInt(node Node) string {
	astInt := node.(*IntegerExpression)
	return fmt.Sprintf("<ast.IntExpr pos: %d value: %d>", astInt.Pos(), astInt.Value)
}

func printExpr(node Node) string {
	exprSt := node.(*ExpressionStatement)
	return Print(exprSt.Expression)
}

func printProgram(node Node) string {
	program := node.(*Program)
	strStmts := make([]string, 0, len(program.Statements))

	for _, st := range program.Statements {
		strStmts = append(strStmts, Print(st))
	}

	return strings.Join(strStmts, "\n")
}

func Print(node Node) string {
	printer, ok := astType2printer[node.Type()]
	if !ok {
		panic("there's no printer for " + node.Type().String())
	}
	return printer(node)
}
