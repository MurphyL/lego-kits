package dsa

// Pair - K 的类型约束为 comparable，V 的类型约束为 any
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type Record map[string]interface{}

func (record Record) Set(key string, value interface{}) {
	record[key] = value
}

func (record Record) Get(key string) interface{} {
	return record[key]
}

func (record Record) Remove(key string) {
	delete(record, key)
}

func (record Record) Lookup(key string) (interface{}, bool) {
	value, ok := record[key]
	return value, ok
}

type DataFrame struct {
	Columns []string
	Payload [][]interface{}
}

func NewDataFrame(payload [][]interface{}, columns []string) DataFrame {
	return DataFrame{Columns: columns, Payload: payload}
}

func NewPair[T interface{}](key string, value T) Pair[string, T] {
	return Pair[string, T]{Key: key, Value: value}
}

type Dataset struct {
	Name  string
	Pairs []Pair[string, any]
}
