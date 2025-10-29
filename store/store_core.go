package main

type Store interface {
	CreateCollection(name string) (Collection, error) // name, fieldNames
	DropCollection(name string) (bool, error)
}

type Collection interface {
	ForEach(func(v []byte, i uint)) bool
	ForUpdate(func(v []byte, i uint) bool) error
	Append([]byte) bool
}

type CollectionIterator struct {
	Offset uint
	Limit  uint
}
