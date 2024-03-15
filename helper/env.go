package helper

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetEnvironment() {
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
