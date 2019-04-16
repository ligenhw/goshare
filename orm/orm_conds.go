package orm

import (
	"fmt"
)

type condValue struct {
	exprs  string
	args   []interface{}
	cond   *Condition
	isOr   bool
	isNot  bool
	isCond bool
	isRaw  bool
	sql    string
}

// Condition struct.
// work for WHERE conditions.
type Condition struct {
	params []condValue
}

// NewCondition return new condition struct
func NewCondition() *Condition {
	c := &Condition{}
	return c
}

// And add expression to condition
func (c Condition) And(expr string, args ...interface{}) *Condition {
	if expr == "" || len(args) == 0 {
		panic(fmt.Errorf("<Condition.And> args cannot empty"))
	}
	c.params = append(c.params, condValue{exprs: expr, args: args})
	return &c
}
