package orm

// real query struct
type QuerySeter struct {
	mi       *modelInfo
	cond     *Condition
	limit    int64
	offset   int64
	distinct bool
	orm      *orm
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
