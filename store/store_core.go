package main

type Store interface {
	CreateCollection(name string, v any) (Collection, error) // name, fieldNames
	DropCollection(name string) (bool, error)
}

type Collection interface {
	AddDocument(v any) bool
}
