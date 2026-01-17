package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abolinjast/taas/internal/config"
	"github.com/abolinjast/taas/internal/handler"
	"github.com/abolinjast/taas/internal/service"
	"github.com/abolinjast/taas/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("something went wrong starting the API: %v", err)
	}
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("something went wrong connecting to the db: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("couldn't ping the db: %v", err)
	}
	log.Println("connection to the database was successfully done")

	sessionStore := store.NewPostgresStore(db)
	sessionService := service.NewSessionService(sessionStore)
	sessionHandler := handler.NewSessionHandler(sessionService)

	router := gin.Default()
	apiRouter := router.Group("/api/v1")
	{
		apiRouter.POST("/sessions/start", sessionHandler.Start)
		apiRouter.POST("/sessions/stop", sessionHandler.Stop)
	}
	socket := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)
	if err := router.Run(socket); err != nil {
		log.Printf("Started the router on port %s", cfg.APIPort)
	}
}
