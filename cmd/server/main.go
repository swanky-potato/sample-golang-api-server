package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swanky-potato/sample-goalng-api-server/pkg/http/handlers"
	"github.com/swanky-potato/sample-goalng-api-server/pkg/http/logger"
)

func init() {
	json, err := strconv.ParseBool(os.Getenv("LOG_JSON"))
	if err != nil {
		json = false
	}
	logger.SetLogger(os.Stdout, json, "DEBUG")
}

func main() {
	r := gin.New()
	r.Use(logger.Log())
	r.Use(gin.Recovery())

	log.Debug("Loading API endpoints")

	r.GET("/ping", handlers.Ping())
	r.GET("/env", handlers.Enviroment())

	r.Run() // listen and serve on 0.0.0.0:8080
}
