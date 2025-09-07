package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"

	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot load .env:", err)
	}

	// init config: cleanenv
	cfg := config.MustLoad()

	// fmt.Println(*cfg)

	// init logger: slog
	log := setupLogger(cfg.Env)

	log.Info("starting url shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init router: chi, "chi render"

	// TODO: run server
}

// go run ./cmd/url-shortener

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		handlerOptions := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler := slog.NewTextHandler(os.Stdout, handlerOptions)
		log = slog.New(handler)
	case envDev:
		handlerOptions := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler := slog.NewJSONHandler(os.Stdout, handlerOptions)
		log = slog.New(handler)
	case envProd:
		handlerOptions := &slog.HandlerOptions{Level: slog.LevelInfo}
		handler := slog.NewJSONHandler(os.Stdout, handlerOptions)
		log = slog.New(handler)
	}

	return log
}
