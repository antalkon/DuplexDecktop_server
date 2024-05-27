package main

import (
	"github.com/antalkon/DuplexDecktop_srver/internal/config"
	"github.com/antalkon/DuplexDecktop_srver/internal/handler"
	"github.com/antalkon/DuplexDecktop_srver/internal/router"
	"github.com/gin-gonic/gin"
	log2 "log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	address := cfg.HTTPServer.Address

	log := setupLogger(cfg.Env)

	log.Info("Starting app", slog.String("env", cfg.Env))

	if cfg.Env == envProd {
		gin.SetMode(gin.ReleaseMode) // Установка режима работы Gin
	}

	h := &handler.Handler{}
	r := router.InitRouters(h)

	if address == "" {
		log2.Fatal("Failed to run server: Adress is empty. \n pls. check config files")
	}

	if err := r.Run(address); err != nil {
		log2.Fatal("Failed to run server: %s", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
