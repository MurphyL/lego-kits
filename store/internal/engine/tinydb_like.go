package engine

import (
	"os"

	"github.com/MurphyL/lego-kits/store/internal/conds"
)

// 参考 TinyDB 实现的一个 JSON 文档数据库 - https://tinydb.readthedocs.io/en/latest/index.html

func NewTinyDB(filepath string) (*TinyDB, error) {
	return &TinyDB{}, nil
}

type Collection interface {
	Insert(doc any) uint64
	Search(*conds.QueryExpr)
}

// TinyDB - tiny, document oriented database
type TinyDB struct {
	fh *os.File
}

type TinyCollection struct {
	name string
}

func (s *TinyDB) Truncate() bool {
	return false
}

func (s *TinyDB) Collection(name string) Collection {
	return &TinyCollection{name: name}
}

func (c *TinyCollection) Insert(doc any) uint64 {
	return 0
}

func (c *TinyCollection) Search(*conds.QueryExpr) {
}
