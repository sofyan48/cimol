package redis

import (
	"fmt"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// Store variabels
var Store redis.Conn

// InitRedis func
// return: redis.Conn
func InitRedis() (redis.Conn, error) {
	rdhost := os.Getenv("REDIS_HOST")
	rdport := os.Getenv("REDIS_PORT")

	connRedis, err := redis.Dial("tcp", fmt.Sprintf(
		"%s:%s", rdhost, rdport))
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to connect to redis from environment: %v", err))
		fmt.Println("Trying Local Connection")
		connRedis, err := redis.Dial("tcp", fmt.Sprintf(
			"localhost:6379"))
		return connRedis, err
	}
	return connRedis, nil
}

// GetConnection function
// return store
func GetConnection() redis.Conn {
	if Store == nil {
		Store, _ = InitRedis()
	}
	return Store
}

// RowsCached params
// @keys: string
// @data: []byte
// @ttl: int
// return []byte, error
func RowsCached(keys string, data []byte, ttl int) ([]byte, error) {
	_, err := redis.String(Store.Do("SET", keys, data))
	if err != nil {
		return nil, err
	}
	redis.String(Store.Do("EXPIRE", keys, ttl))
	return data, nil
}

// GetRowsCached params
// @keys: string
// return string, error
func GetRowsCached(keys string) (string, error) {
	value, err := redis.String(Store.Do("GET", keys))
	if err != nil {
		return "", err
	}
	return value, nil
}

// GetSessionData params
// @keys: string
// @scheme: scheme
// return interface , error
func GetTokenData(context *gin.Context) string {
	token := GetToken(context)
	return token
}

// GetToken params
// @context: *gin.Context
// return string
func GetToken(context *gin.Context) string {
	token := context.Request.Header["Authorization"]
	if len(token) < 1 {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token Not Compatible",
		})
		context.Abort()
	}
	return token[0]
}
