package main

import (
	"os"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"fmt"
	"./routes"
)

func Env_load() {
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Critical("Error loading .env file")
	}
}


func loadConfig() {
	logger, err := log.LoggerFromConfigAsFile("server/config/logger.xml")

	if err != nil {
		panic("fail to load config")
	}

	log.ReplaceLogger(logger)
}

func pingPong(c *gin.Context) {
	c.String(200, "pong")
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return port
}

func main() {
	defer log.Flush()
	loadConfig()

	log.Info("Starting Application.")

	Env_load()

	r := gin.Default()
	r.GET("/", routes.Root)
	r.GET("/ping", pingPong)
	r.Use(static.Serve("/", static.LocalFile("assets", true)))

	r.Run(":" + getPort())
}
