package cache

import (
	"net/http"
	"sync"
)

// 客户端连接池, 复用连接, 减少内存分配
var (
	DefaultClientPool = sync.Pool{
		New: func() interface{} {
			return &http.Client{}
		},
	}
)
