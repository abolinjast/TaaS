package main

import (
	"fmt"
	"log"

	"github.com/abolinjast/taas/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("something went wrong starting the API: %v", err)
	}
	router := gin.Default()
	socket := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)
	if err := router.Run(socket); err != nil {
		log.Printf("Started the router on port %s", cfg.APIPort)
	}
}
