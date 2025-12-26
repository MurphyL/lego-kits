package semver

import (
	"regexp"
	"strconv"
	"strings"
)

// 参考文档 - https://semver.org/
const SV_PATTERN = "(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-[0-9a-zA-Z-]+)*"

type VERSION_STEP uint

const MAJOR_PART, MINOR_PART, PATCH_PART VERSION_STEP = 0, 1, 2

func Resolve(plainText string) (SemVer, bool) {
	matche := regexp.MustCompile(SV_PATTERN).FindString(plainText)
	if len(matche) == 0 {
		return nil, false
	} else {
		sv := internalSerVer{}
		sv.versionStr, sv.additional, _ = strings.Cut(matche, "-")
		parts := strings.SplitN(sv.versionStr, ".", 3)
		sv.majorVersionNum, _ = strconv.ParseUint(parts[MAJOR_PART], 10, 8)
		sv.minorVersionNum, _ = strconv.ParseUint(parts[MINOR_PART], 10, 8)
		sv.patchVersionNum, _ = strconv.ParseUint(parts[PATCH_PART], 10, 8)
		return &sv, true
	}
}

type SemVer interface {
	Major() uint                    // 主版本号
	Minor() uint                    // 次版本号
	Patch() uint                    // 修正版本
	Labels() []string               // 额外的标签：alpha, beta, snapshot
	Value() string                  // 干净的版本号
	Compare(otherSemVer SemVer) int // 对比两个版本
}

type internalSerVer struct {
	versionStr      string
	additional      string
	majorVersionNum uint64
	minorVersionNum uint64
	patchVersionNum uint64
}

func (v *internalSerVer) Compare(otherSemVer SemVer) int {
	return strings.Compare(v.versionStr, otherSemVer.Value())
}

func (v *internalSerVer) Major() uint {
	return uint(v.majorVersionNum)
}

func (v *internalSerVer) Minor() uint {
	return uint(v.minorVersionNum)
}

func (v *internalSerVer) Patch() uint {
	return uint(v.patchVersionNum)
}

func (v *internalSerVer) Value() string {
	return v.versionStr
}

func (v *internalSerVer) Labels() []string {
	if len(v.additional) > 0 {
		return strings.Split(v.additional, "-")
	}
	return []string{}
}
