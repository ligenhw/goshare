package orm

import (
	"fmt"
)

type condValue struct {
	exprs string
	args  []interface{}
	cond  *Condition
	isOr  bool
	isNot bool
	isIn  bool
	isRaw bool
	sql   string
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

// And add expression to condition
func (c Condition) Raw(sql string) *Condition {
	if sql == "" {
		panic(fmt.Errorf("<Condition.Raw> sql cannot empty"))
	}
	c.params = append(c.params, condValue{sql: sql, isRaw: true})
	return &c
}

func (c Condition) In(expr string, args ...interface{}) *Condition {
	if expr == "" || len(args) == 0 {
		panic(fmt.Errorf("<Condition.In> args cannot empty"))
	}
	c.params = append(c.params, condValue{exprs: expr, args: args, isIn: true})
	return &c
}
