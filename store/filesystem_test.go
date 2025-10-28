package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInternalFilesystemStore_CreateCollection(t *testing.T) {
	var st Store
	var err error
	type et struct {
		Name string `lego:"name"`
	}
	x := new(et)
	if st, err = NewFilesystemStore("E:/repo"); err == nil {
		var collName = "hello_y"
		if coll, ex := st.CreateCollection(collName, x); ex == nil {
			for i := 0; i < 100; i++ {
				e := et{Name: fmt.Sprintf("hello #%05d", i)}
				if ok := coll.AddDocument(e); ok {
					t.Log("集合", collName, "的记录新增成功#", i)
				}
			}

		} else {
			t.Errorf("创建集合出错：%s", err)
		}
	} else {
		t.Errorf("创建客户端出错：%s", err)
	}
}

func TestTags(t *testing.T) {
	x := new(struct {
		Name string `lego:"name"`
	})
	pt := reflect.TypeOf(x).Elem()
	for i := 0; i < pt.NumField(); i++ {
		xf := pt.Field(i)
		tag := xf.Tag.Get("lego")
		t.Log(tag)
	}
}
