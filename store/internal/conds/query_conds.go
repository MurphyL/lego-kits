package conds

type QueryBuilder struct {
}

type QueryExpr struct {
	qb *QueryBuilder
}

func (qb *QueryBuilder) Eq(field string, value any) *QueryExpr {
	return &QueryExpr{qb: qb}
}

func (qb *QueryBuilder) Gt(field string, value any) *QueryExpr {
	return &QueryExpr{qb: qb}
}

func (qb *QueryBuilder) Lt(field string, value any) *QueryExpr {
	return &QueryExpr{qb: qb}
}

func (qb *QueryBuilder) Gte(field string, value any) *QueryExpr {
	return &QueryExpr{qb: qb}
}

func (qb *QueryBuilder) Lte(field string, value any) *QueryExpr {
	return &QueryExpr{qb: qb}
}
