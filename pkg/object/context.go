package object

import (
	"fmt"
)

// Context is a variable container
type Context interface {
	Set(varName string, object Object) error
	Get(varName string) (Object, error)
}

// context is a native implementation of Context
// it holds varmap as a map[string]object.Object
type context struct {
	varmap map[string]Object
}

// Set sets variable in current context
// returns error is variable had been set already
func (c *context) Set(varName string, object Object) error {
	if _, ok := c.varmap[varName]; ok {
		// redefinition!
		return fmt.Errorf("redifinition of variable %s", varName)
	}
	c.varmap[varName] = object
	return nil
}

// Set gets variable from current context
// returns error is variable has not ever been set
func (c *context) Get(varName string) (Object, error) {
	if val, ok := c.varmap[varName]; !ok {
		// redefinition!
		return nil, fmt.Errorf("undefined variable %s", varName)
	} else {
		return val, nil
	}
}

// NewContext is Context constructor
func NewContext() Context {
	return &context{varmap: make(map[string]Object)}
}
