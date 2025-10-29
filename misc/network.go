package misc

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

// CacheEntry 缓存条目
type CacheEntry struct {
	IPs       []string
	ExpiresAt time.Time
}

// DNSCache 线程安全的DNS缓存
type DNSCache struct {
	cache sync.Map
	ttl   time.Duration
}

// NewDNSCache 创建DNS缓存实例
func NewDNSCache(ttl time.Duration) *DNSCache {
	return &DNSCache{ttl: ttl}
}

// Lookup 查询域名
func (c *DNSCache) Lookup(ctx context.Context, domain string) ([]string, error) {
	if entry, ok := c.cache.Load(domain); ok {
		if ce, ok := entry.(CacheEntry); ok && time.Now().Before(ce.ExpiresAt) {
			return ce.IPs, nil
		}
	}
	ips, err := net.DefaultResolver.LookupHost(ctx, domain)
	if err != nil {
		return nil, err
	}
	c.cache.Store(domain, CacheEntry{
		IPs:       ips,
		ExpiresAt: time.Now().Add(c.ttl),
	})
	return ips, nil
}

func main() {
	cache := NewDNSCache(10 * time.Minute)
	domain := "example.com"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ips, err := cache.Lookup(ctx, domain)
	if err != nil {
		fmt.Println("查询错误:", err)
		return
	}
	fmt.Printf("%s 的A记录: %v\n", domain, ips)
}
