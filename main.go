package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const logPath = "./logs/go.log"

var logger *zap.Logger

func setupLog() {
	var err error
	_, err = os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	logger, err = c.Build()
	if err != nil {
		panic(err)
	}
}

func main() {
	setupLog()

	r := gin.Default()

	// Setting GIN to use zap as logger
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/ping", func(c *gin.Context) {
		logger.Info("my own info", zap.Int64("key", rand.Int63()))
		logger.Error("myerrror level", zap.Int64("key", rand.Int63()), zap.String("ran", "my"+strconv.Itoa(rand.Int())))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
