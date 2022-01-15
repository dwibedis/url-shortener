package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func Store(c *gin.Context, db UrlDb) bool {
	db.addedOn = time.Now()
	clusters := []string{"url_shortner_redis:7000", "url_shortner_redis:7001",
		"url_shortner_redis:7002", "url_shortner_redis:7003", "url_shortner_redis:7004"}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              clusters,
	})

	err := rdb.Set(c, db.ID, db.ParentUrl, -1).Err()
	log.Println("id: " + db.ID + ", parentUrl:" + db.ParentUrl)
	if err != nil {
		log.Println("error occurred while saving in redis: ", err.Error())
		return false
	}
	return true
}

func Get(id string) string {
	clusters := []string{"url_shortner_redis:7000", "url_shortner_redis:7001",
		"url_shortner_redis:7002", "url_shortner_redis:7003", "url_shortner_redis:7004"}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              clusters,
	})
	val, err := rdb.Get(context.Background(), id).Result()
	if err != nil {
		return ""
	}
	return val
}