package main

import (
	"context"
	"dwibedis/url-shortener/internal/controller"
	"dwibedis/url-shortener/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := gin.Default()
	log.Println("invoking main!!")
	router.GET("/generate-redirect-url", controller.GenerateRedirectUrl).Use(middleware.Authenticate)
	router.GET("/r/:urlId", controller.GetRedirectUrlFromShortenedUrl).Use(middleware.Authenticate)
	log.Println("invoking redis ping!!")
	router.GET("/redis-ping", controller.RedisPing)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
