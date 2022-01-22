package service

import (
	"crypto/md5"
	"dwibedis/url-shortener/internal/repository"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
)

func GenerateAndStoreUrl(c *gin.Context, url string) repository.Url {
	log.Println("Generating shortened url for url: " + url)
	hash := md5.Sum([]byte(url))
	shortenedUrl := "url-shortner/r/" + hex.EncodeToString(hash[:])
	log.Print("url: " + url +  ", shortened: " + shortenedUrl)
	status := repository.Store(c, repository.UrlDb{
		ID: hex.EncodeToString(hash[:]),
		ParentUrl: url,
	})
	if !status {
		return repository.Url{}
	}
	return repository.Url{
		URL: shortenedUrl,
	}
}

func GetRedirectUrlFromShortUrl(shortUrlId string) repository.Url {
	parentUrl:= repository.Get(shortUrlId)
	return repository.Url{
		URL: parentUrl,
	}
}
