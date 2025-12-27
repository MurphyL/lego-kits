package amap

type ParsedAmapResult map[string]any

func (r ParsedAmapResult) Success() bool {
	return r["status"] == "1" || r["status"] == 1
}

func (r ParsedAmapResult) Code() uint {
	return r["infocode"].(uint)
}

func (r ParsedAmapResult) Get(key string) any {
	return r[key]
}
