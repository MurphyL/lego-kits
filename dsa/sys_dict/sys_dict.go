package sys_dict

import "errors"

type DictType interface {
	// 字典类型编码
	DictCode() string
	// 字典类型名称
	DictName() string
	// 字典项
	DictItems() []DictItem
	// 保存
	SaveType() (bool, error)
}

type DictItem interface {
	// 字典类型编码
	DictCode() string
	// 字典项标签
	ItemLabel() string
	// 字典项文本值
	ItemValue() string
	// 保存
	SaveItem() (bool, error)
}

func SaveDictType(dt DictType) (bool, error) {
	if nil == dt {
		return false, errors.New("参数为空")
	}
	return dt.SaveType()
}

func SaveDictItem(di DictItem) (bool, error) {
	if nil == di {
		return false, errors.New("参数为空")
	}
	return di.SaveItem()
}
