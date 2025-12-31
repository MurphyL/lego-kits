package main

import (
	"log"
	"testing"
)

func TestNewStore(t *testing.T) {
	if tdb, err := NewStore("e:/tmp/a.json"); err == nil {
		coll := tdb.Collection("hello")
		coll.Insert(map[string]any{"a": 1, "b": 1})
		ret := coll.Search(func(field string, value any) bool {
			return false
		})
		log.Println("搜索结构：", ret)
	} else {
		t.Fatal("创建数据库出错：", err)
	}
}
