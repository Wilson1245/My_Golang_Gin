package service

import (
	"fmt"
	red "golangAPI/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
)

func CachDecorator(h gin.HandlerFunc, porm string, readKeyPattern string, empty interface{}) gin.HandlerFunc {
	{
		return func(c *gin.Context) {
			keyId := c.Param(porm)
			redisKey := fmt.Sprintf(readKeyPattern, keyId)
			conn := red.RedisDefaultPool.Get()
			defer conn.Close()
			data, err := redis.Bytes(conn.Do("GET", redisKey))
			if err != nil {
				h(c)
				dbResult, exists := c.Get("dbResult")
				if !exists {
					dbResult = empty
				}
				redisData, _ := ffjson.Marshal(dbResult)
				conn.Do("SETEX", redisKey, 30, redisData)
				c.JSON(http.StatusOK, gin.H{
					"Message": "From DB",
					"Data":    dbResult,
				})
				return
			}
			ffjson.Unmarshal(data, &empty)
			c.JSON(http.StatusOK, gin.H{
				"Message": "From Redis",
				"Data":    empty,
			})
		}
	}
}

func CachDecoratorUserAll(h gin.HandlerFunc, readKeyPattern string, empty interface{}) gin.HandlerFunc {
	{
		return func(c *gin.Context) {
			conn := red.RedisDefaultPool.Get()
			defer conn.Close()
			data, err := redis.Bytes(conn.Do("GET", readKeyPattern))
			if err != nil {
				h(c)
				dbUserAll, exists := c.Get("dbUserAll")
				if !exists {
					dbUserAll = empty
				}
				redisData, _ := ffjson.Marshal(dbUserAll)
				conn.Do("SETEX", readKeyPattern, 30, redisData)
				c.JSON(http.StatusOK, gin.H{
					"Message": "From DB",
					"Data":    dbUserAll,
				})
				return
			}
			ffjson.Unmarshal(data, &empty)
			c.JSON(http.StatusOK, gin.H{
				"Message": "From Redis",
				"Data":    empty,
			})
		}
	}
}
