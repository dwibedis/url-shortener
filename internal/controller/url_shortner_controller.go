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
	shortenedUrlId := c.Param("urlId")

	if len(shortenedUrlId) == 0 {
		log.Println("Empty short URL in request, not status")
		c.IndentedJSON(http.StatusUnprocessableEntity, repository.Url{})
		return
	}

	//c.IndentedJSON(http.StatusOK, service.GetRedirectUrlFromShortUrl(shortenedUrlId).URL)
	originalUrl := service.GetRedirectUrlFromShortUrl(shortenedUrlId).URL
	if originalUrl[0:4] != "http" {
		originalUrl = "https://" + originalUrl
	}
	c.Redirect(http.StatusFound, originalUrl)
}

func GenerateRedirectUrl(c *gin.Context) {
	url, status:= c.GetQuery("url")

	if len(url) == 0 || !status {
		log.Println("Empty URL in request!!")
		c.IndentedJSON(http.StatusUnprocessableEntity, repository.Url{})
		return
	}

	log.Println("Received request for url: " + url)
	c.IndentedJSON(http.StatusCreated, service.GenerateAndStoreUrl(c, url))
}

func RedisPing(c *gin.Context) {
	log.Println("redis ping controller invoked!!")
	clusters := []string{"url_shortner_redis:7000", "url_shortner_redis:7001",
		"url_shortner_redis:7002", "url_shortner_redis:7003", "url_shortner_redis:7004"}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              clusters,
	})
	err:= rdb.Ping(c).Err()
	if (err != nil) {
		log.Println("error occurred, err:" , err.Error())
		c.IndentedJSON(http.StatusGone, err)
	}
	c.IndentedJSON(http.StatusAccepted, rdb.Ping(c).Val())
}