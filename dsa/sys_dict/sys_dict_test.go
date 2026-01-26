package sys_dict

import (
	"errors"
	"log"
	"testing"
)

type testDictType struct {
	code, name, intro string
}

type testDictItem struct {
	code, label, value, intro string
}

func (d *testDictType) DictCode() string {
	return d.code
}

func (d *testDictType) DictName() string {
	return d.name
}

func (d *testDictType) SaveType() (bool, error) {
	return false, errors.New("no impl")
}

func (d *testDictType) DictItems() []DictItem {
	var ret []DictItem = []DictItem{}
	return ret
}

func (d *testDictItem) DictCode() string {
	return d.code
}

func (d *testDictItem) ItemLabel() string {
	return d.label
}

func (d *testDictItem) ItemValue() string {
	return d.value
}

func (d *testDictItem) SaveItem() (bool, error) {
	return false, errors.New("no impl")
}

func TestDictType(t *testing.T) {
	tdt := &testDictType{code: "hello"}
	if ok, err := SaveDictType(tdt); err == nil {
		log.Println("保存完成：", ok, tdt)
	} else {
		log.Println("保存出错：", ok, tdt)
	}
	items := tdt.DictItems()
	log.Println("字典项：", items)
}

func TestDictItem(t *testing.T) {
	tdi := &testDictItem{code: "hello"}
	if ok, err := SaveDictItem(tdi); err == nil {
		log.Println("保存完成：", ok, tdi)
	} else {
		log.Println("保存出错：", ok, tdi)
	}
}
