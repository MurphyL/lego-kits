package data_dict

type Loader func(group, label string) []Entry

func New(loader Loader) DataDict {
	return &internalDataDict{loader: loader}
}

type DataDict interface {
	Group(name string) []Entry
	Entry(group, name string) Entry
	Value(group, name string) string
}

type Entry interface {
	Group() string
	Label() string
	Value() string
	Intro() string
}

type internalDataDict struct {
	loader Loader
}

func (dd *internalDataDict) Group(name string) []Entry {
	return dd.loader(name, "")
}

func (dd *internalDataDict) Entry(group, label string) Entry {
	ret := dd.loader(group, label)
	if nil == ret || len(ret) == 0 {
		return nil
	} else {
		return ret[0]
	}
}

func (dd *internalDataDict) Value(group, label string) string {
	ret := dd.Entry(group, label)
	if nil == ret {
		return ""
	} else {
		return ret.Value()
	}
}
