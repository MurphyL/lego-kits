package person

type MetaInfo[K any] struct {
	Id            int64
	Name          string
	Intro         string
	Birthday      string // 生日
	PhoneOrMobile string // 手机或电话
}

type GenericPerson[K any] interface {
	MetaInfo[K]
}
