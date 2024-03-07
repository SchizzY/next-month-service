package helper

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetEnvironment() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	GinMode()
}

func GinMode() {
	mode := os.Getenv("GIN_MODE")

	fmt.Printf("GIN_MODE=%s\n", mode)

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
