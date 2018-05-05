package interpreter

import (
	"github.com/pmukhin/glisp/pkg/object"
)

func glispAppend(args ...object.Object) (object.Object, error) {
	if len(args) < 2 {
		return nil, makeArgsLenErr("append", 2, len(args))
	}
	if args[0].Type() != object.TList {
		return nil, makeUnexpectedTypeErr("append", 0,
			object.TList, args[0].Type())
	}
	list := args[0].(*object.List)
	newList := &object.List{Elements: list.Elements}

	for _, argument := range args[1:] {
		newList.Elements = append(newList.Elements, argument)
	}

	return newList, nil
}
