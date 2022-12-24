package plugins_test

import (
	"testing"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/plugins"
)

func TestNewRedisCacheConnection(t *testing.T) {
	// create redis connection
	redisUrl := "redis://:@localhost:6379/0"
	rdb, rContext := plugins.NewRedisCacheConnection(redisUrl)

	// check redis connection
	pong, err := rdb.Ping(rContext).Result()
	if err != nil {
		t.Errorf("NewRedisCacheConnection() error = %v", err)
		return
	}

	if pong != "PONG" {
		t.Errorf("NewRedisCacheConnection() pong = %v, want %v", pong, "PONG")
		return
	}

	// close redis connection
	err = rdb.Close()
	if err != nil {
		t.Errorf("NewRedisCacheConnection() error = %v", err)
		return
	}
}
