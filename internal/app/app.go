package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Habeebamoo/tunnl-backend/internal/configs"
	"github.com/Habeebamoo/tunnl-backend/internal/handlers"
	"github.com/Habeebamoo/tunnl-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type App struct {
	server *http.Server
	router *gin.Engine
	config *configs.Config
}

func New() *App {
	cfg := configs.Load()

	// Gin init
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()


	//services init
	notificationService := services.NewNotificationService()

	
	//handlers init
	notificationHandler := handlers.NewNotificationHandler(notificationService)


	// routes injected here
	RegisterRoutes(router, notificationHandler)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &App{
		server: server,
		router: router,
		config: cfg,
	}
}

func (a *App) Run() error {
	go func() {
		log.Printf("Tunnl running on PORT :%s", a.config.Port)

		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return a.server.Shutdown(ctx)
}