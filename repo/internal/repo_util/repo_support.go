package repo_util

import (
	"github.com/jmoiron/sqlx"
)

type SqlxRepo struct {
	*sqlx.DB
}

func NewInternalRepo(driverName, dataSourceName string) *SqlxRepo {
	return &SqlxRepo{DB: sqlx.MustOpen(driverName, dataSourceName)}
}
