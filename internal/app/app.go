package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/markovk1n/spoty/configs"
	"github.com/markovk1n/spoty/internal/handler"
	"github.com/markovk1n/spoty/internal/repository"
	"github.com/markovk1n/spoty/internal/service"
	"github.com/markovk1n/spoty/pkg/logger"
	"github.com/markovk1n/spoty/pkg/postgres"
	"github.com/markovk1n/spoty/pkg/server"
)

func Run(cfg *configs.AppConfigs) {
	log := logger.NewLogger("debug")
	log.Info("Wait db creating")
	time.Sleep(10 * time.Second)

	db, err := postgres.NewDB(configs.Postgres{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatal("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err = srv.Run("8080", handler.InitRoutes()); err != nil {
			log.Fatal("error occured while running http server: %s", err.Error())
		}
	}()
	log.Warn("Spoty Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("TodoApp Shutting Down")

	if err := srv.ShutDown(context.Background()); err != nil {
		log.Error("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Error("error occured on db connection close:", err.Error())
	}
}
