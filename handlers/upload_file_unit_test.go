package handlers_test

import (
	"testing"
	"time"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/handlers"
	"github.com/gin-contrib/cache/persistence"
)

func TestCacheSet(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "test"
	value := "test"
	expiration := time.Second

	// call service method
	err := handlers.CacheSet(inAppCache, key, value, expiration)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}
}

func TestCacheGet(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "test"
	value := "test"
	expiration := time.Minute

	// call service method
	err := handlers.CacheSet(inAppCache, key, value, expiration)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}

	// call service method
	var valueCached string
	err = handlers.CacheGet(inAppCache, key, &valueCached)
	if err != nil {
		t.Errorf("CacheGet() error = %v", err)
		return
	}

	if valueCached != value {
		t.Errorf("CacheGet() valueCached = %v, want %v", valueCached, value)
		return
	}
}

func TestCacheGetWithTTLSuccess(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "testttl"
	value := "testttlsuccess"
	exp := time.Second
	err := handlers.CacheSet(inAppCache, key, value, exp)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}

	// call service method
	var valueCached string
	err = handlers.CacheGet(inAppCache, key, &valueCached)
	if err != nil {
		t.Errorf("CacheGet() error = %v", err)
		return
	}

	if valueCached != value {
		t.Errorf("CacheGet() valueCached = %v, want %v", valueCached, value)
		return
	}
}

func TestCacheGetWithTTLExpired(t *testing.T) {
	// create memory cache
	inAppCache := persistence.NewInMemoryStore(time.Minute)

	key := "testttlfail"
	value := "testttlfailed"
	exp := time.Second
	err := handlers.CacheSet(inAppCache, key, value, exp)
	if err != nil {
		t.Errorf("CacheSet() error = %v", err)
		return
	}

	// wait for expiration
	time.Sleep(2 * time.Second)

	// call service method
	var valueCached string
	err = handlers.CacheGet(inAppCache, key, &valueCached)
	if err == nil {
		t.Errorf("CacheGet() key should be expired, error = %v", err)
		return
	}

	if valueCached == value {
		t.Errorf("CacheGet() valueCached = %v, want %v", valueCached, value)
		return
	}
}
