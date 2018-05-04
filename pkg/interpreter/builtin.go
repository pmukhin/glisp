package interpreter

import (
	"fmt"
	"strings"

	"github.com/pmukhin/glisp/pkg/object"
)

type internalFunc func(args ...object.Object) (object.Object, error)

var internalFunctionTable = map[string]internalFunc{
	"+":     add,
	"-":     sub,
	"/":     div,
	"*":     mul,
	"print": glispPrint,
}

func makeArgErr(funName string, expected int, given int) error {
	return fmt.Errorf("%s expects at least %d args, %d given", funName, expected, given)
}

func makeFunNotDefErr(funName string, oType object.Type) error {
	return fmt.Errorf("%s is not defined for type %d", funName, oType)
}

func makeUnexpectedTypeErr(funName string, pos int, oTypeExp, oTypeGiven object.Type) error {
	return fmt.Errorf("%s expects positional argument #%d to be of type %s, %s given",
		funName, pos, oTypeExp, oTypeGiven)
}

func extractIntArgs(args []object.Object) ([]int64, error) {
	intArgs := make([]int64, len(args))
	for i, ar := range args {
		oInt, ok := ar.(*object.Int)
		if !ok {
			return nil, makeUnexpectedTypeErr("__mul__", i,
				object.TInt, ar.Type())
		}
		intArgs[i] = oInt.Value
	}
	return intArgs, nil
}

func extractFloatArgs(args []object.Object) ([]float64, error) {
	floatArgs := make([]float64, len(args))
	for i, ar := range args {
		oFloat, ok := ar.(*object.Float)
		if !ok {
			return nil, makeUnexpectedTypeErr("__mul__", i,
				object.TFloat, ar.Type())
		}
		floatArgs[i] = oFloat.Value
	}
	return floatArgs, nil
}

func glispPrint(args ...object.Object) (object.Object, error) {
	strList := make([]string, len(args))
	for i, v := range args {
		strList[i] = v.String()
	}

	fmt.Printf("%s\n", strings.Join(strList, " "))

	return nil, nil
}
