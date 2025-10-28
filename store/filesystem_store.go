package main

import (
	"bufio"
	"errors"
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

type filesystemCollection struct {
	collName  string
	storeRoot *os.Root
}

func (c filesystemCollection) Iterator() *CollectionIterator {
	return &CollectionIterator{}
}

type WithFilesystemStoreOption func(*filesystemStore)

func (s filesystemStore) createCollectionName(name string) string {
	return strings.Join([]string{name, s.fileSuffix}, ".")
}

func (s filesystemStore) CreateCollection(name string) (Collection, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("集合名称不能为空")
	}
	collName := s.createCollectionName(name)
	if _, err := s.root.Stat(collName); err == nil && os.IsNotExist(err) {
		return nil, fmt.Errorf("已存在同名（%s）集合", collName)
	}
	// TODO 拆分文件
	if _, err := s.root.OpenFile(collName, os.O_WRONLY|os.O_CREATE, 0644); err == nil {
		return &filesystemCollection{collName: collName, storeRoot: s.root}, err
	} else {
		return nil, errors.Join(errors.New("写入数据文件失败"), err)
	}
}

func (s filesystemStore) DropCollection(name string) (bool, error) {
	collName := s.createCollectionName(name)
	err := s.root.Remove(collName)
	return err == nil, errors.Join(errors.New("删除数据文件失败"), err)
}

func WithFileSuffix(suffix string) WithFilesystemStoreOption {
	return func(store *filesystemStore) {
		store.fileSuffix = suffix
	}
}

func (c filesystemCollection) ForEach(handleEach func(v []byte, i uint)) bool {
	if fh, err := c.storeRoot.Open(c.collName); err == nil {
		scanner := bufio.NewScanner(fh)
		var i uint
		for scanner.Scan() {
			handleEach(scanner.Bytes(), i)
			i++
		}
		return true
	} else {
		return false
	}
}

func (c filesystemCollection) Append(v []byte) bool {
	if fh, err := c.storeRoot.OpenFile(c.collName, os.O_APPEND, 0644); err == nil {
		_, ew := fh.Write(fmt.Appendln(v))
		return ew == nil
	}
	return false
}
