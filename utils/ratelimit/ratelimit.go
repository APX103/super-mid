package ratelimit

import (
	"errors"
	"fmt"

	"apx103.com/super-mid/utils/config"

	"github.com/gin-gonic/gin"
	lru "github.com/hashicorp/golang-lru/v2"
	"golang.org/x/time/rate"
)

const (
	ServerLimitType string = "server"
	IPLimitType     string = "ip"
)

type RateLimiter struct {
	limitType          string
	keyFunc            func(*gin.Context) string
	cache              *lru.Cache[string, *rate.Limiter]
	rateLimiterFactory func() *rate.Limiter
}

func NewRateLimiter(conf *config.LimitConfig) (*RateLimiter, error) {
	if conf == nil {
		return nil, errors.New("invalid config")
	}

	if err := conf.Validate(); err != nil {
		return nil, err
	}

	var keyFunc func(*gin.Context) string
	switch conf.LimitType {
	case ServerLimitType:
		keyFunc = func(c *gin.Context) string {
			return ""
		}
	case IPLimitType:
		keyFunc = func(c *gin.Context) string {
			return c.ClientIP()
		}
	default:
		return nil, fmt.Errorf("unknow limit type %s", conf.LimitType)
	}

	c, err := lru.New[string, *rate.Limiter](conf.CacheSize)
	if err != nil {
		return nil, err
	}

	rateLimiterFactory := func() *rate.Limiter {
		return rate.NewLimiter(rate.Limit(conf.QPS), conf.Burst)
	}

	return &RateLimiter{
		limitType:          conf.LimitType,
		keyFunc:            keyFunc,
		cache:              c,
		rateLimiterFactory: rateLimiterFactory,
	}, nil
}

func (rl *RateLimiter) Accept(c *gin.Context) error {
	key := rl.keyFunc(c)
	limiter := rl.get(key)

	if !limiter.Allow() {
		return fmt.Errorf("limit reached on %s for key %v", rl.limitType, key)
	}

	return nil
}

func (rl *RateLimiter) get(key string) *rate.Limiter {
	value, found := rl.cache.Get(key)
	if !found {
		limiter := rl.rateLimiterFactory()
		rl.cache.Add(key, limiter)
		return limiter
	}
	return value
}
