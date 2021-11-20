package limiterx

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type Limiters struct {
	limiters *sync.Map
}

type Limiter struct {
	limiter *rate.Limiter
	lastGet time.Time // 上一次获取token的时间
	key     string
}

var GlobalLimiters = &Limiters{
	limiters: &sync.Map{},
}

var once = sync.Once{}

func NewLimiter(interval time.Duration, burst int, key string) *Limiter {
	once.Do(func() {
		go GlobalLimiters.clearLimiter()
	})
	rateLimit := rate.Every(interval)
	keyLimiter := GlobalLimiters.getLimiter(rateLimit, burst, key)
	return keyLimiter
}

func (limit *Limiter) Allow() bool {
	limit.lastGet = time.Now()
	return limit.limiter.Allow()
}

func (limits *Limiters) getLimiter(rateLimit rate.Limit, burst int, key string) *Limiter {
	limiter, ok := limits.limiters.Load(key)
	if ok {
		return limiter.(*Limiter)
	}
	limit := &Limiter{
		limiter: rate.NewLimiter(rateLimit, burst),
		lastGet: time.Now(),
		key:     key,
	}
	limits.limiters.Store(key, limit)
	return limit
}

// 清除过期的限流器
func (limits *Limiters) clearLimiter() {
	for {
		time.Sleep(1 * time.Minute)
		limits.limiters.Range(func(key, value interface{}) bool {
			// 超过1分钟
			if time.Now().Unix()-value.(*Limiter).lastGet.Unix() > 60 {
				limits.limiters.Delete(key)
			}
			return true
		})
	}
}
