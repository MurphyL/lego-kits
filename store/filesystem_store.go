package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

type WithFilesystemStoreOption func(*filesystemStore)

func (s filesystemStore) createCollectionName(name string) string {
	return strings.Join([]string{name, s.fileSuffix}, ".")
}

func (s filesystemStore) CreateCollection(name string, fields ...string) (bool, error) {
	if len(name) == 0 {
		return false, fmt.Errorf("集合名称不能为空")
	}
	if len(fields) == 0 {
		return false, fmt.Errorf("至少需要一个字段")
	}
	collName := s.createCollectionName(name)
	if _, err := s.root.Stat(collName); err == nil && os.IsNotExist(err) {
		return false, fmt.Errorf("已存在同名（%s）集合", collName)
	}
	var err error
	var fh *os.File
	if fh, err = s.root.OpenFile(collName, os.O_WRONLY|os.O_CREATE, 0644); err == nil {
		data, _ := json.Marshal(fields)
		_, err = fh.Write(fmt.Appendln(data))
	}
	if err == nil {
		return true, nil
	} else {
		s.DropCollection(collName)
		return false, err
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
