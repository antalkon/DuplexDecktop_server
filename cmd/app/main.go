package main

import (
	"fmt"
	"github.com/antalkon/DuplexDecktop_srver/internal/config"
	"github.com/antalkon/DuplexDecktop_srver/internal/handler"
	"github.com/antalkon/DuplexDecktop_srver/internal/router"
	"github.com/gin-gonic/gin"
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
	db := config.MustDatabase()

	address := cfg.HTTPServer.Address

	log := setupLogger(cfg.Env)
	fmt.Printf("Db final data: %s\n", db)
	log.Info("Starting app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// Установка режима работы Gin
	if cfg.Env == envProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// Создание нового роутера Gin и инициализация маршрутов
	h := &handler.Handler{}
	r := router.InitRouters(h)

	// Получаем хост и порт из конфигурации

	// Проверка, что адрес включает порт
	if address == "" {
		//log.Fatal("HTTPServer address is not set in the configuration")
	}

	// Запуск сервера на указанном хосте и порте
	if err := r.Run(address); err != nil {
		//log.Fatalf("Failed to run server: %s", err)
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
