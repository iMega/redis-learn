package tests

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

type redisFailoverConfig struct {
	MasterName    string   `json:"master_name"`
	SentinelAddrs []string `json:"sentinel_addrs"`
	DB            int      `json:"db"`
}

func GetConfigValue(key string) (string, error) {
	value := os.Getenv(key)
	return value, nil
}

func GetRedis() *redis.Client {
	if redisClient == nil {
		if err := WaitForDB(); err != nil {
			panic(err)
		}

	}

	return redisClient
}

func WaitForDB() error {
	raw, err := GetConfigValue("REDIS_CONFIG")
	if err != nil {
		return err
	}

	c := &redisFailoverConfig{}
	err = json.Unmarshal([]byte(raw), c)
	if err != nil {
		return err
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: c.SentinelAddrs[0],
		DB:   c.DB,
	})

	maxRetries := 30
	for {
		_, err := redisClient.Ping().Result()
		if err == nil {
			break
		}
		log.Print(".")
		if maxRetries == 0 {
			break
		}
		maxRetries--
		<-time.After(time.Duration(1 * time.Second))
	}

	_, err = redisClient.Ping().Result()

	return err
}
