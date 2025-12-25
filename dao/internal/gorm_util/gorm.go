package gorm_util

import (
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
