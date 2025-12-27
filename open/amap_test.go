package open

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestAmapPOI2(t *testing.T) {
	var app PlatformApp
	if tokens, ok := os.LookupEnv("AMAP_TOKENS"); ok {
		appKey, appSecret, _ := strings.Cut(tokens, "/")
		app = NewPlatformApp(appKey, appSecret)
	} else {
		log.Fatalln("未配置正常的高德开放平台应用参数")
	}
	u, _ := url.Parse("https://restapi.amap.com/v5/place/text")
	q := u.Query()
	q.Add("key", app.AppKey())
	q.Add("keywords", "武汉大学")
	u.RawQuery = q.Encode()
	r := http.Request{URL: u}
	resp, _ := app.DoRequest(&r)
	if resp.StatusCode == 200 {
		bytes, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		log.Println(string(bytes))
	}
}
