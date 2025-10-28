package main

import (
	"encoding/json"
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
	if st, err = NewFilesystemStore("E:/repo"); err == nil {
		var collName = "hello_y"
		if coll, ex := st.CreateCollection(collName); ex == nil {
			for i := 0; i < 100; i++ {
				d, _ := json.Marshal(et{Name: fmt.Sprintf("hello #%05d", i)})
				if ok := coll.Append(d); ok {
					t.Log("写入集合", collName, "的记录新增成功#", i)
				}
			}

			coll.ForEach(func(v []byte, i uint) {
				e := new(et)
				json.Unmarshal(v, e)
				t.Log("读取集合", collName, "的记录成功#", i, e)
			})
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
