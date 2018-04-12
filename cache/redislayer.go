// This package allows you to Utilize the Redis Database through its Get and Post methods.
// The Redis service must be running for this to perform as expected, else connection will fail.
package cache

import (
	"github.com/gomodule/redigo/redis"
	"strings"
	"log"
)

func handleError(err error, messages ... string) {
	comboMessage := strings.Join(messages, ", ")
	if err != nil {
		log.Fatal("ERR "+comboMessage, err)
		panic(err)
	}
}

func redisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	handleError(err)
	return c
}

func ClearDB() {
	c := redisConnect()
	defer c.Close()
	_, err := c.Do("FLUSHALL")
	handleError(err)
}

func Post(key, value string) {
	c := redisConnect()
	defer c.Close()
	_, err := c.Do("SET", key, value)
	handleError(err)
}

func Get(key string) (string, error) {
	c := redisConnect()
	defer c.Close()
	result, err := redis.String(c.Do("GET", key))
	return result, err
}

func AddToUnorderedSet(key string, members ... string) {
	c := redisConnect()
	defer c.Close()
	for _, member := range members {
		_, err := c.Do("SADD", key, member)
		if err != nil {
			handleError(err)
			break
		}
	}
}

func GetFromUnorderedSet(key string) []string {
	c := redisConnect()
	defer c.Close()
	result, err := redis.Strings(c.Do("SMEMBERS", key))
	handleError(err, "Failed getting SMembers")
	return result
}
