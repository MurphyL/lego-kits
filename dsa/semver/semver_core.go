package semver

import (
	"log"
	"regexp"
)

// 参考文档 - https://semver.org/
var SV_PATTERN = regexp.MustCompile("^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)$")

func Resolve(plainText string) SemVer {
	matches := SV_PATTERN.FindAllString(plainText, -1)
	log.Println(matches)
	return nil
}

type SemVer interface {
	Major() uint8
	Minor() uint8
	Patch() uint8
	Labels() []string
}
