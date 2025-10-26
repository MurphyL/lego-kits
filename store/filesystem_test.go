package main

import (
	"testing"
)

func TestInternalFilesystemStore_CreateCollection(t *testing.T) {
	var st Store
	var err error
	if st, err = NewFilesystemStore("E:/repo"); err == nil {
		var ok bool
		var collName = "hello_y"
		if ok, err = st.CreateCollection("hello_y", "a", "b", "c"); ok && err == nil {
			for i := 0; i < 100; i++ {
				if x := st.AddDocument(collName, i*10+1, i*10+2, i*10+3); x {
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

func TestFilesystemStore_ListDocument(t *testing.T) {
	if st, err := NewFilesystemStore("E:/repo"); err == nil {
		st.ListDocument("hello_y")
	}
}
