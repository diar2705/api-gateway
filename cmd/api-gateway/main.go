package main

import (
	"os"

	routerApi "github.com/BetterGR/api-gateway/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()

	err := godotenv.Load()
	if err != nil {
		klog.Fatalf("Error loading .env file")
	}

	// Get the port from the environment variable, default to 1234 if not set
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		klog.Fatalf("API_GATEWAY_PORT is not set")
	}

	router := gin.New()

	routerApi.InitiateRoutes(router)

	err = router.Run(":" + port)
	if err != nil {
		klog.Fatalf("Failed to start the server, %v", err)
	}
}
