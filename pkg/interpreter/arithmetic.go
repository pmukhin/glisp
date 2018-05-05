package interpreter

import (
	"strings"

	"github.com/pmukhin/glisp/pkg/object"
)

func mul(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, makeArgsLenErr("__mul__", 2, len(args))
	}
	firstType := args[0].Type()
	switch firstType {
	// switch arg types...
	case object.TInt:
		intArgs, err := extractIntArgs(args)
		if err != nil {
			return nil, err
		}

		return iMul(intArgs...)
	case object.TFloat:
		floatArgs, err := extractFloatArgs(args)
		if err != nil {
			return nil, err
		}
		return fMul(floatArgs...)
	case object.TString:
		strToRep, ok := args[0].(*object.String)
		if !ok {
			return nil, makeUnexpectedTypeErr("__mul__", 0,
				object.TString, args[0].Type())
		}
		intArgs, err := extractIntArgs(args[1:])
		if err != nil {
			return nil, err
		}
		return sMul(strToRep, intArgs...)
	default:
		return nil, makeFunNotDefErr("__mul__", firstType)
	}
}

func sMul(strToRep *object.String, muls ...int64) (object.Object, error) {
	value := strToRep.Value
	for _, m := range muls {
		value = strings.Repeat(value, int(m))
	}
	return &object.String{Value: value}, nil
}

func fMul(args ...float64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret *= v
	}
	return &object.Float{Value: ret}, nil
}

func iMul(args ...int64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret *= v
	}
	return &object.Int{Value: ret}, nil
}

func div(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, makeArgsLenErr("__div__", 2, len(args))
	}
	firstType := args[0].Type()
	switch firstType {
	// switch arg types...
	case object.TInt:
		intArgs := make([]int64, len(args))
		for i, ar := range args {
			oInt, ok := ar.(*object.Int)
			if !ok {
				return nil, makeUnexpectedTypeErr("__div__", i,
					object.TInt, ar.Type())
			}
			intArgs[i] = oInt.Value
		}
		return iDiv(intArgs...)
	case object.TFloat:
		floatArgs := make([]float64, len(args))
		for i, ar := range args {
			oFloat, ok := ar.(*object.Float)
			if !ok {
				return nil, makeUnexpectedTypeErr("__div__", i,
					object.TFloat, ar.Type())
			}
			floatArgs[i] = oFloat.Value
		}
		return fDiv(floatArgs...)
	default:
		return nil, makeFunNotDefErr("__div__", firstType)
	}
}

func fDiv(args ...float64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret /= v
	}
	return &object.Float{Value: ret}, nil
}

func iDiv(args ...int64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret /= v
	}
	return &object.Int{Value: ret}, nil
}

func sub(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, makeArgsLenErr("__sub__", 2, len(args))
	}
	firstType := args[0].Type()
	switch firstType {
	// switch arg types...
	case object.TInt:
		intArgs := make([]int64, len(args))
		for i, ar := range args {
			oInt, ok := ar.(*object.Int)
			if !ok {
				return nil, makeUnexpectedTypeErr("__sub__", i,
					object.TInt, ar.Type())
			}
			intArgs[i] = oInt.Value
		}
		return iSub(intArgs...)
	case object.TFloat:
		floatArgs := make([]float64, len(args))
		for i, ar := range args {
			oFloat, ok := ar.(*object.Float)
			if !ok {
				return nil, makeUnexpectedTypeErr("__sub__", i,
					object.TFloat, ar.Type())
			}
			floatArgs[i] = oFloat.Value
		}
		return fSub(floatArgs...)
	default:
		return nil, makeFunNotDefErr("__sub__", firstType)
	}
}

func fSub(args ...float64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret -= v
	}
	return &object.Float{Value: ret}, nil
}

func iSub(args ...int64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret -= v
	}
	return &object.Int{Value: ret}, nil
}

func add(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, makeArgsLenErr("__add__", 2, len(args))
	}
	firstType := args[0].Type()
	switch firstType {
	// switch arg types...
	case object.TInt:
		intArgs := make([]int64, len(args))
		for i, ar := range args {
			oInt, ok := ar.(*object.Int)
			if !ok {
				return nil, makeUnexpectedTypeErr("__add__", i,
					object.TInt, ar.Type())
			}
			intArgs[i] = oInt.Value
		}
		return iAdd(intArgs...)
	case object.TFloat:
		floatArgs := make([]float64, len(args))
		for i, ar := range args {
			oFloat, ok := ar.(*object.Float)
			if !ok {
				return nil, makeUnexpectedTypeErr("__add__", i,
					object.TFloat, ar.Type())
			}
			floatArgs[i] = oFloat.Value
		}
		return fAdd(floatArgs...)
	default:
		return nil, makeFunNotDefErr("__add__", firstType)
	}
}

func fAdd(args ...float64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret += v
	}
	return &object.Float{Value: ret}, nil
}

func iAdd(args ...int64) (object.Object, error) {
	ret := args[0]
	for _, v := range args[1:] {
		ret += v
	}
	return &object.Int{Value: ret}, nil
}
