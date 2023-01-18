package service

import (
	"SampleGoGin/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
)

func CacheOneUserDecorator(
	h gin.HandlerFunc,
	porm string,
	readKeyPattern string,
	value interface{},
) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyId := c.Param(porm)
		redisKey := fmt.Sprintf(readKeyPattern, keyId)
		conn := database.RedisDefaultPool.Get()
		defer conn.Close()
		data, err := redis.Bytes(conn.Do("GET", redisKey))
		if err != nil {
			h(c)
			dbResult, exists := c.Get("dbResult")
			if !exists {
				dbResult = value
			}
			redisData, _ := ffjson.Marshal(dbResult)
			conn.Do("SETEX", redisKey, 30, redisData)
			c.JSON(http.StatusOK, gin.H{
				"message": "From DB",
				"data":    dbResult,
			})
			return
		}
		var retVal interface{}
		ffjson.Unmarshal(data, &retVal)
		c.JSON(http.StatusOK, gin.H{
			"message": "From Redis",
			"data":    retVal,
		})
	}
}

func CacheUserAllDecorator(
	h gin.HandlerFunc,
	redisKey string,
	value interface{},
) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := database.RedisDefaultPool.Get()
		defer conn.Close()

		data, err := redis.Bytes(conn.Do("GET", redisKey))
		if err != nil {
			h(c)
			dbUserAll, exists := c.Get("dbUserAll")
			if !exists {
				dbUserAll = value
			}
			redisData, _ := ffjson.Marshal(dbUserAll)
			conn.Do("SETEX", redisKey, 30, redisData)
			c.JSON(http.StatusOK, gin.H{
				"message": "From DB",
				"data":    dbUserAll,
			})
			return
		}
		var retVal interface{}
		ffjson.Unmarshal(data, &retVal)
		c.JSON(http.StatusOK, gin.H{
			"message": "From Redis",
			"data":    retVal,
		})
	}
}
