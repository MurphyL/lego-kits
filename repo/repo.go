package dao

import (
	"repo/internal/repo_util"
)

type ID interface{ uint64 | string }

type Repo interface {
	Get(dest interface{}, query string, args ...interface{}) error
}

func NewRepo(driverName, dataSourceName string) Repo {
	return repo_util.NewInternalRepo(driverName, dataSourceName)
}
