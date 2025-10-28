package misc

type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)
