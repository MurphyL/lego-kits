package core

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormConfig struct{ Driver, Dsn string }

func ConnectGorm(opts *GormConfig) (*gorm.DB, error) {
	var conn gorm.Dialector
	switch opts.Driver {
	default:
		conn = sqlite.Open(opts.Dsn)
	}
	return gorm.Open(conn, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
	})
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
