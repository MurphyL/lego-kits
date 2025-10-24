package repo

import (
	"fmt"

	"gorm.io/gorm"
)

type Model interface {
}

type QueryResult[T Model, Q any] struct {
	Params  Q
	Total   int64
	Records []T
}

type PageQuery struct {
	PageNum  int
	PageSize int
}

type SkipQuery struct {
	Key   string
	Value string
	Count int
}

func NewPageQuery[T Model](db *gorm.DB, query *PageQuery, model *T) (*gorm.DB, int64) {
	var total int64
	db.Model(model).Count(&total)
	offset := (query.PageNum - 1) * query.PageSize
	return db.Offset(offset).Limit(query.PageSize), total
}

func NewSkipQuery[T Model](db *gorm.DB, query *SkipQuery, model *T) (*gorm.DB, int64) {
	var total int64
	db.Model(model).Count(&total)
	return db.Where(fmt.Sprintf("%s > ?", query.Key), query.Value).Limit(query.Count), total
}
