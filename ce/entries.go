package ce

type CustomEntryKind string

const (
	Text   CustomEntryKind = "text"
	Number CustomEntryKind = "number"
	Enum   CustomEntryKind = "enum"
)

// CustomEntry 自定义字段
type CustomEntry[K any] interface {
	Kind() CustomEntryKind
	Value() K
}

type internalTextEntry struct {
	value string
}

func (t internalTextEntry) Kind() CustomEntryKind {
	return Text
}

func (t internalTextEntry) Value() string {
	return t.value
}

type internalNumberEntry struct {
	value int
}

func (t internalNumberEntry) Kind() CustomEntryKind {
	return Number
}

func (t internalNumberEntry) Value() int {
	return t.value
}

type internalEnumEntry struct {
	value int
}

func (t internalEnumEntry) Kind() CustomEntryKind {
	return Number
}

func (t internalEnumEntry) Value() int {
	return t.value
}
