package helper

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetEnvironment() {
	if os.Getenv("FLY_APP_NAME") == "" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	GinMode()
}

func GinMode() {
	mode := os.Getenv("GIN_MODE")

	if mode != "debug" && mode != "release" {
		panic("Invalid Gin mode: " + mode)
	}
	if mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func NeonConn() string {
	return os.Getenv("CON")
}
