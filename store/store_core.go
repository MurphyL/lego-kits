package main

type Store interface {
	CreateCollection(name string, fields ...string) (bool, error) // name, fieldNames
	DropCollection(name string) (bool, error)
	AddDocument(collName string, vals ...any) bool
	ListDocument(collName string) (bool, [][]byte)
}
