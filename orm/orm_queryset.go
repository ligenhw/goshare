package orm

import "strings"

// real query struct
type QuerySeter struct {
	mi       *modelInfo
	cond     *Condition
	limit    int64
	offset   int64
	orders   []string
	distinct bool
	orm      *orm
}

func getCondSQL(cond *Condition) (where string, args []interface{}) {
	var wheres []string
	for i, p := range cond.params {
		if i > 0 {
			if p.isOr {
				where += "OR "
			} else {
				where += "AND "
			}
		}

		wheres = append(wheres, p.exprs+" = ? ")
		args = append(args, p.args...)
	}

	where = strings.Join(wheres, ", ")

	if where != "" {
		where = "WHERE " + where
	}
	return
}

// create new QuerySeter.
func newQuerySet(orm *orm, mi *modelInfo) *QuerySeter {
	o := new(QuerySeter)
	o.mi = mi
	o.orm = orm
	return o
}

// add condition expression to QuerySeter.
func (o *QuerySeter) Filter(expr string, args ...interface{}) *QuerySeter {
	if o.cond == nil {
		o.cond = NewCondition()
	}
	o.cond = o.cond.And(expr, args...)
	return o
}

func (o *QuerySeter) All(container interface{}, cols ...string) (int64, error) {
	return o.orm.DbBaser.ReadBatch(o.orm.db, o, o.mi, o.cond, container, cols)
}

// add ORDER expression.
// "column" means ASC, "-column" means DESC.
func (o *QuerySeter) OrderBy(exprs ...string) *QuerySeter {
	o.orders = exprs
	return o
}
