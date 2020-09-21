package proxy

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	RedisUrl         string
	Port             string
	CacheKeyCapacity *int
	CacheTTL         *time.Duration
	ProxyClientLimit *int
}

func (c Config) getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// NewConfig creates the struct Config by parsing environment
// vars or setting them to defaul values, if applicable
func NewConfig() Config {
	c := Config{}
	c.RedisUrl = c.getEnv("REDIS_URL", "localhost:6379")
	c.Port = ":" + c.getEnv("PORT", "8080")
	ckp := c.getEnv("CACHE_KEY_CAPACITY", "")
	if ckp != "" {
		x, err := strconv.ParseInt(ckp, 10, 64)
		if err != nil {
			log.Fatal(err)
		} else {
			xc := int(x)
			c.CacheKeyCapacity = &xc
		}

	}
	cttl := c.getEnv("CACHE_TTL", "")
	if cttl != "" {
		ct, err := time.ParseDuration(cttl + "s")
		if err != nil {
			log.Fatal(err)
		} else {
			c.CacheTTL = &ct
		}
	}
	pcl := c.getEnv("PROXY_CLIENT_LIMIT", "")
	if pcl != "" {
		l, err := strconv.ParseInt(pcl, 10, 64)
		if err != nil {
			log.Fatal(err)
		} else {
			lc := int(l)
			c.CacheKeyCapacity = &lc
		}
	}
	return c
}
