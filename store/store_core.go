package main

type Store interface {
	CreateCollection(name string) (Collection, error) // name, fieldNames
	DropCollection(name string) (bool, error)
}

type ReadableCollection interface {
	ForEach(func(v []byte, i uint)) bool
}

type ModifiableCollection interface {
	Append([]byte) bool
}

type Collection interface {
	ReadableCollection
	ModifiableCollection
	Iterator() *CollectionIterator
}

type CollectionIterator struct {
	Collection ReadableCollection
	Offset     uint
	Limit      uint
}
