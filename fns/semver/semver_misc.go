package semver

import (
	"fmt"
	"strings"
)

func Upgrage(v SemVer, step VERSION_STEP) SemVer {
	ret := internalSerVer{
		majorVersionNum: uint64(v.Major()),
		minorVersionNum: uint64(v.Minor()),
		patchVersionNum: uint64(v.Patch()),
		additional:      strings.Join(v.Labels(), "-"),
	}
	switch step {
	case MAJOR_PART:
		ret.majorVersionNum += 1
	case MINOR_PART:
		ret.minorVersionNum += 1
	case PATCH_PART:
		ret.patchVersionNum += 1
	}
	ret.versionStr = fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch())
	return &ret
}

func UpgrageMajor(v SemVer) SemVer {
	return Upgrage(v, MAJOR_PART)
}

func UpgrageMinor(v SemVer) SemVer {
	return Upgrage(v, MAJOR_PART)
}

func UpgragePatch(v SemVer) SemVer {
	return Upgrage(v, PATCH_PART)
}
