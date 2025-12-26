package semver

import (
	"fmt"
	"testing"
)

func TestResolveVersions(t *testing.T) {
	versions := []string{"1.2.3", "v1.2.3", "v1.2.3-beta", "v1.2.3-beta-v1"}
	for _, version := range versions {
		t.Run(version, func(t *testing.T) {
			v, ok := Resolve(version)
			if !ok {
				t.Fatal("解析出错：", v)
			}
			if v.Major() != 1 {
				t.Fatal("解析主版本号出错：", v.Major())
			}
			if v.Minor() != 2 {
				t.Fatal("解析次版本号出错：", v.Minor())
			}
			if v.Patch() != 3 {
				t.Fatal("解析修订版本号出错：", v.Patch())
			}
			t.Log("major:", v.Major(), "minor:", v.Minor(), "patch:", v.Patch(), "labels:", v.Labels())
		})
	}
}

func TestCompareVersions(t *testing.T) {
	versions := []string{"1.2.3", "v1.2.3", "v1.2.3-beta", "v1.2.3-beta-v1"}
	a1, _ := Resolve("1.2.3")
	a2, _ := Resolve("1.2.1")
	a3, _ := Resolve("1.2.4")
	for _, version := range versions {
		t.Run(version, func(t *testing.T) {
			b, _ := Resolve(version)
			t.Log(fmt.Sprintf("compare('%s', '%s')", a1.Value(), version), a1.Compare(b))
			t.Log(fmt.Sprintf("compare('%s', '%s')", a2.Value(), version), a2.Compare(b))
			t.Log(fmt.Sprintf("compare('%s', '%s')", a3.Value(), version), a3.Compare(b))
		})
	}
}
