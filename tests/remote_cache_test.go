package tests_test

import (
	"testing"
	"time"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/plugins"
)

func TestNewRedisCacheSet(t *testing.T) {
	// create redis connection
	redisUrl := "redis://:@localhost:6379/0"
	rdb, rContext := plugins.NewRedisCacheConnection(redisUrl)

	// set value
	key := "test-for-redis"
	value := "test-value"
	exp := time.Minute
	err := rdb.Set(rContext, key, value, exp).Err()
	if err != nil {
		t.Errorf("TestNewRedisCacheSet() could not set value to redis. error = %v", err)
		return
	}
}

func TestNewRedisCacheGet(t *testing.T) {
	// create redis connection
	redisUrl := "redis://:@localhost:6379/0"
	rdb, rContext := plugins.NewRedisCacheConnection(redisUrl)

	// set value
	key := "test-for-redis"
	value := "test-value"
	exp := time.Minute
	err := rdb.Set(rContext, key, value, exp).Err()
	if err != nil {
		t.Errorf("TestNewRedisCacheGet() could not set value to redis. error = %v", err)
		return
	}

	// get value
	val, err := rdb.Get(rContext, key).Result()
	if err != nil {
		t.Errorf("TestNewRedisCacheGet() could not get value from redis. error = %v", err)
		return
	}
	if val != value {
		t.Errorf("TestNewRedisCacheGet() could not get value from redis. error = %v", err)
		return
	}

}
