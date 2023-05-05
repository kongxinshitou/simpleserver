package cache

import (
	"sync"
	"time"
)

// 定义缓存全局变量
var (
	DefaultCache = NewDefaultTimeoutCache()
)

// 实际上有多个超时缓存, 提高并发度
type TimeoutCache struct {
	num    int
	caches []*Cache
}

// 实际工作的超时缓存
type Cache struct {
	mu      sync.Mutex
	data    map[string]result
	timeout time.Duration
}

// 超时缓存中的value类型
type result struct {
	val      string
	lastTime time.Time
}

func (c *Cache) Set(key string, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = result{val, time.Now()}
}

// 返回true表示需要发送请求
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.data[key]
	if !ok {
		return "", false
	}
	val, lastTime := v.val, v.lastTime
	if time.Since(lastTime) > c.timeout {
		return "", false
	}
	return val, true
}

func NewDefaultTimeoutCache() *TimeoutCache {
	TC := &TimeoutCache{}
	TC.num = 20
	for i := 0; i < 20; i++ {
		TC.caches = append(TC.caches, &Cache{
			data:    make(map[string]result),
			timeout: 10 * time.Second,
		})
	}
	return TC
}

func (TC *TimeoutCache) Key2Index(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])) * 13 % 1000000000
	}
	return hash % TC.num
}

func (TC *TimeoutCache) Set(key, val string) {
	TC.caches[TC.Key2Index(key)].Set(key, val)
}
func (TC *TimeoutCache) Get(key string) (string, bool) {
	return TC.caches[TC.Key2Index(key)].Get(key)
}
