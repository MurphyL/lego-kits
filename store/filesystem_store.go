package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func NewFilesystemStore(path string, withOptions ...WithFilesystemStoreOption) (Store, error) {
	if root, err := os.OpenRoot(path); err == nil {
		st := filesystemStore{root: root}
		if len(withOptions) > 0 {
			for _, withOption := range withOptions {
				withOption(&st)
			}
		} else {
			WithFileSuffix("st")(&st)
		}
		return &st, nil
	} else {
		return nil, fmt.Errorf("创建文件存储客户端出错：%s", err.Error())
	}
}

type filesystemStore struct {
	root       *os.Root
	fileSuffix string
}

type filesystemCollection struct {
	collName    string
	storeRoot   *os.Root
	elementType reflect.Type
}

func (c filesystemCollection) AddDocument(v any) bool {
	if fh, err := c.storeRoot.OpenFile(c.collName, os.O_WRONLY, 0644); err == nil {
		if data, em := json.Marshal(v); em == nil {
			_, ew := fh.Write(fmt.Appendln(data))
			return ew == nil
		}
	}
	return false
}

type WithFilesystemStoreOption func(*filesystemStore)

func (s filesystemStore) createCollectionName(name string) string {
	return strings.Join([]string{name, s.fileSuffix}, ".")
}

func (s filesystemStore) CreateCollection(name string, v any) (Collection, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("集合名称不能为空")
	}
	var valueType reflect.Type
	var fields []string
	if nil == v {
		return nil, fmt.Errorf("数值对象不能为空")
	}
	if vt := reflect.TypeOf(v); vt.Kind() == reflect.Ptr {
		valueType = reflect.TypeOf(v).Elem()

	} else {
		valueType = reflect.TypeOf(v)
	}
	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		structTag := structField.Tag.Get("lego")
		if len(structTag) > 0 {
			tags := strings.Split(structTag, ",")[:]
			fields = append(fields, tags[0])
		} else {
			fields = append(fields, structField.Name)
		}
	}
	collName := s.createCollectionName(name)
	if _, err := s.root.Stat(collName); err == nil && os.IsNotExist(err) {
		return nil, fmt.Errorf("已存在同名（%s）集合", collName)
	}
	// TODO 拆分文件
	if _, err := s.root.OpenFile(collName, os.O_WRONLY|os.O_CREATE, 0644); err == nil {
		return &filesystemCollection{collName: collName, elementType: valueType}, err
	} else {
		return nil, err
	}
}

func (s filesystemStore) DropCollection(name string) (bool, error) {
	return false, nil
}

func (s filesystemStore) AddDocument(name string, vals ...any) bool {
	collName := s.createCollectionName(name)
	if fh, err := s.root.OpenFile(collName, os.O_APPEND, 0644); err == nil {
		data, _ := json.Marshal(vals)
		fh.Write(fmt.Appendln(data))
		return true
	} else {
		return false
	}
}

func (s filesystemStore) ListDocument(name string) (bool, [][]byte) {
	collName := s.createCollectionName(name)
	if fh, err := s.root.Open(collName); err == nil {
		scanner := bufio.NewScanner(fh)
		for scanner.Scan() {
			data := scanner.Bytes()
			fmt.Println(string(data))
		}
		return true, nil
	} else {
		return false, nil
	}
}

func WithFileSuffix(suffix string) WithFilesystemStoreOption {
	return func(store *filesystemStore) {
		store.fileSuffix = suffix
	}
}
