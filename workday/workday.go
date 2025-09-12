package workday

import (
	"bufio"
	"embed"
	"log"
	"strings"
)

//go:embed vendors
var vendors embed.FS

func init() {
	fh, _ := vendors.Open("vendors/data.csv")
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		seps := strings.SplitN(scanner.Text(), ",", 3)
		log.Println(seps)
	}
}

func Get() {

}

// IsChineseWorkday - 确认格式为（YYYY-MM-DD）日期字符串对应的日期是否为中国大陆的工作日
func IsChineseWorkday(date string) {

}
