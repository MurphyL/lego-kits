package semver

import "testing"

func TestResolve(t *testing.T) {
	v1 := Resolve("1.2.3")
	if v1.Major() != 1 {
		t.Errorf("解析主版本号出错")
	}
	if v1.Minor() != 2 {
		t.Errorf("解析次版本号出错")
	}
	if v1.Patch() != 3 {
		t.Errorf("解析修订版本号出错")
	}
}
