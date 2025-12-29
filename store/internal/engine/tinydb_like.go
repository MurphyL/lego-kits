package engine

import (
	"log"
	"os"

	"github.com/MurphyL/lego-kits/store/internal/conds"
)

// 参考 TinyDB 实现的一个 JSON 文档数据库 - https://tinydb.readthedocs.io/en/latest/index.html

func NewTinyDB(filepath string) (*TinyDB, error) {
	if fh, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0666); err == nil {
		return &TinyDB{fh: fh}, nil
	} else {
		return nil, err
	}
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
	if err := s.fh.Truncate(0); err != nil {
		log.Panicln("清空数据文件出错：", err)
		return false
	}
	if _, err := s.fh.WriteString("{}"); err != nil {
		log.Panicln("初始化数据文件出错：", err)
		return false
	}
	return true
}

func (s *TinyDB) Collection(name string) Collection {
	return &TinyCollection{name: name}
}

func (c *TinyCollection) Insert(doc any) uint64 {
	return 0
}

func (c *TinyCollection) Search(*conds.QueryExpr) {
}
