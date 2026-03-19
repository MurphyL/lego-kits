package amap

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestAmapInterfaces(t *testing.T) {
	var app *Assistant
	if tokens, ok := os.LookupEnv("AMAP_TOKENS"); ok {
		log.Println("高德开放平台应用参数：", tokens)
		appKey, appSecret, _ := strings.Cut(tokens, "/")
		app = NewAmapClient(appKey, appSecret)
	} else {
		log.Fatalln("未配置正常的高德开放平台应用参数")
	}
	t.Run("POI2.0", func(t *testing.T) {
		if ret, ok := app.SearchAround("武汉大学"); ok {
			log.Println("调用高德地图开发平台 POI2.0 接口成功：", ret)
		} else {
			log.Println("调用高德地图开发平台 POI2.0 接口出错：", ok)
		}
	})
	t.Run("ReGEO", func(t *testing.T) {
		if ret, ok := app.ReGEO("武汉大学"); ok {
			log.Println("调用高德地图开发平台地理/逆地理编码接口成功：", ret)
		} else {
			log.Println("调用高德地图开发平台地理/逆地理编码接口出错：", ok)
		}
	})
}
