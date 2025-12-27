package amap

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	var appKey, appSecret string
	if tokens, ok := os.LookupEnv("AMAP_TOKENS"); ok {
		log.Println("高德开放平台应用参数：", tokens)
		appKey, appSecret, _ = strings.Cut(tokens, "/")
	} else {
		log.Fatalln("未配置正常的高德开放平台应用参数")
	}
	log.Println("应用数据：", appKey, appSecret)
}
