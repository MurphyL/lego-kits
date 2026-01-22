package engine

import (
	"encoding/json"
	"io"
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
	Search(conds.Query) []any
}

// TinyDB - tiny, document oriented database
type TinyDB struct {
	fh *os.File
}

type TinyCollection struct {
	holder *TinyDB
	name   string
}

func (d *TinyDB) Truncate() bool {
	if err := d.fh.Truncate(0); err != nil {
		log.Panicln("清空数据文件出错：", err)
		return false
	}
	if _, err := d.fh.WriteString("{}"); err != nil {
		log.Panicln("初始化数据文件出错：", err)
		return false
	}
	return true
}

func (d *TinyDB) load() (map[string][]any, error) {
	if data, err := io.ReadAll(d.fh); err == nil {
		ret := map[string][]any{}
		err = json.Unmarshal(data, &ret)
		if err != nil {
			return nil, err
		}
		return ret, nil
	} else {
		return nil, err
	}
}

func (d *TinyDB) dump(dict map[string][]any) error {
	if data, err := json.Marshal(dict); err == nil {
		d.fh.Truncate(0)
		d.fh.Seek(0, io.SeekStart)
		_, err = d.fh.Write(data)
		return err
	} else {
		return err
	}
}

func (d *TinyDB) Collection(name string) Collection {
	return &TinyCollection{holder: d, name: name}
}

func (c *TinyCollection) useCollection(handleColl func(coll []any, db map[string][]any)) {
	if dict, err := c.holder.load(); err == nil {
		handleColl(dict[c.name], dict)
	} else {
		handleColl([]any{}, map[string][]any{})
	}
}

func (c *TinyCollection) Insert(doc any) uint64 {
	log.Println("尝试写入数据记录：", doc)
	c.useCollection(func(coll []any, db map[string][]any) {
		db[c.name] = append(coll, doc)
		if err := c.holder.dump(db); err != nil {
			log.Panicln("数据写入出错：", err)
		}
	})
	return 0
}

func (c *TinyCollection) Search(q conds.Query) []any {
	var ret []any
	c.useCollection(func(coll []any, db map[string][]any) {
		ret = coll
	})
	return ret
}
