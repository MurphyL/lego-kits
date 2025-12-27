package main

import "testing"

func TestNewStore(t *testing.T) {
	if tdb, err := NewStore("e:/tmp/a.json"); err == nil {
		coll := tdb.Collection("hello")
		q := NewQuery()
		expr := q.Eq("hello", "x")
		coll.Search(expr)
	} else {
		t.Fatal("创建数据库出错：", err)
	}
}
