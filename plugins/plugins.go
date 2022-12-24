package plugins

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	// cache
	"github.com/go-redis/redis/v8"

	// nats
	"github.com/nats-io/nats.go"

	// in-app-cache
	"github.com/gin-contrib/cache/persistence"
)

func NewRedisCacheConnection(redisUrl string) (*redis.Client, context.Context) {
	// parse redis url to username password and host and port and db
	// redis://username:password@host:port/db
	_, username, password, host, port, db := UrlStringToOptions(redisUrl)
	// convert db to int
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		log.Fatalln("Redis db option is not a number: ", err)
	}

	rContext := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Username: username,
		Password: password,
		DB:       dbInt,
	})
	// ping redis for check connection
	_, err = rdb.Ping(rContext).Result()
	if err != nil {
		log.Fatalln("Redis initial conn error: ", err)
	}
	return rdb, rContext
}

func NewNATSConnection(natsUrl string) *nats.Conn {
	_, username, password, host, port, _ := UrlStringToOptions(natsUrl)

	// create nats connection
	nc, err := nats.Connect(host+":"+port, nats.UserInfo(username, password))
	if err != nil {
		log.Fatalln("NATS initial conn error: ", err)
	}

	return nc
}

func NewInAppCacheStore(defaultTTL time.Duration) *persistence.InMemoryStore {
	return persistence.NewInMemoryStore(defaultTTL)
}

func UrlStringToOptions(url string) (string, string, string, string, string, string) {
	// parse redis url to username password and host and port and db
	// app_name://username:password@host:port/db

	// split all options (app_name, username, password, host, port, db)
	options := strings.Split(url, "://")
	// split protocol and info
	protocol := options[0]
	info := options[1]
	// split info to username, password, host, port, db
	infoOptions := strings.Split(info, "@")
	// split username and password
	usernamePassword := infoOptions[0]
	hostPortDb := infoOptions[1]
	// split username and password
	usernamePasswordOptions := strings.Split(usernamePassword, ":")
	username := usernamePasswordOptions[0]
	password := usernamePasswordOptions[1]
	// split host, port and db
	hostPortDbOptions := strings.Split(hostPortDb, "/")
	hostPort := hostPortDbOptions[0]
	db := hostPortDbOptions[1]
	// split host and port
	hostPortOptions := strings.Split(hostPort, ":")
	host := hostPortOptions[0]
	port := hostPortOptions[1]

	return protocol, username, password, host, port, db
}
