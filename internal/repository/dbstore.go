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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.28.1.4:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.28.1.4:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := rdb.Get(context.Background(), id).Result()
	if err != nil {
		return ""
	}
	return val
}