package repo

import (
	"github.com/jmoiron/sqlx"

	"github.com/MurphyL/lego-kits/repo/internal/repo_util"
)

type ID interface{ uint64 | string }

func NewRepo(driverName, dataSourceName string) repo_util.SqlxRepo {
	return repo_util.SqlxRepo{DB: sqlx.MustOpen(driverName, dataSourceName)}
}
