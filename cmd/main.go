package main

import (
	"flag"
	"fmt"

	"tg_bot_proxy/internal/config"
	"tg_bot_proxy/internal/database"
	"tg_bot_proxy/internal/logger"
	"tg_bot_proxy/internal/repository"
	"tg_bot_proxy/internal/service"

	// "github.com/golang-migrate/migrate/v4/database"
	"go.uber.org/zap"
)


func main() {
	configPath := flag.String("c", "./cmd/go-telegram-bot-example/config.yaml", "path to go-telegram-bot-example config")
	flag.Parse()

	logger, err := logger.GetLogger()

	if err != nil {
		panic(fmt.Sprintf("failed setting up logger: %s", err.Error()))
	}

	cfg := config.Config{}
	err = config.GetConfig(*configPath, cfg)
	if err != nil {
		panic(fmt.Sprint("failed get configuration", zap.String("reason", err.Error())))
	}
	logger.Info("configured", zap.Any("config", cfg))

	db, err := database.NewPostgresDB(cfg.Database)
	// defer db.Close()
	if err != nil {
		panic(fmt.Sprint("failed connect to DB", zap.String("reason", err.Error())))

	}
	logger.Info("success connected to database")
	repo := repository.Init(db)
	svc := service.Init(repo)
	// TODO(): create flow like state machine using map

}
