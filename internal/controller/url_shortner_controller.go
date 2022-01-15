package controller

import (
	"dwibedis/url-shortener/internal/repository"
	"dwibedis/url-shortener/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)


func GetRedirectUrlFromShortenedUrl(c *gin.Context) {
	shortenedUrl, status:= c.GetQuery("url")

	if len(shortenedUrl) == 0 || !status {
		log.Println("Empty short URL in request, not status")
		c.IndentedJSON(http.StatusUnprocessableEntity, repository.Url{})
		return
	}

	c.IndentedJSON(http.StatusOK, service.GetRedirectUrlFromShortUrl(shortenedUrl))
}

func GenerateRedirectUrl(c *gin.Context) {
	url, status:= c.GetQuery("url")

	if len(url) == 0 || !status {
		log.Println("Empty URL in request!!")
		c.IndentedJSON(http.StatusUnprocessableEntity, repository.Url{})
		return
	}

	log.Println("Received request for url: " + url)
	// Add the new album to the slice.
	c.IndentedJSON(http.StatusCreated, service.GenerateAndStoreUrl(c, url))
}

func RedisPing(c *gin.Context) {
	log.Println("redis ping controller invoked!!")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "docker.for.mac.localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	c.IndentedJSON(http.StatusAccepted,rdb.Ping(c).Val())
}