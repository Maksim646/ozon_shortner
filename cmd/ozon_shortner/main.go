package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/Maksim646/ozon_shortner/internal/api/server/restapi/handler"
	"github.com/Maksim646/ozon_shortner/internal/config"
	"github.com/Maksim646/ozon_shortner/internal/database/postgresql"
	"github.com/Maksim646/ozon_shortner/internal/model"
	"github.com/Maksim646/ozon_shortner/pkg/logger"
	"github.com/justinas/alice"

	"github.com/Maksim646/ozon_shortner/internal/domain/link_inmemory/repository/inmemory"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	_linkRepo "github.com/Maksim646/ozon_shortner/internal/domain/link_postgresql/repository/postgresql"
	_linkUsecase "github.com/Maksim646/ozon_shortner/internal/domain/link_postgresql/usecase"
)

const (
	httpVersion = "development"

	dbTypePostgres = "postgres"
	dbTypeInMemory = "inmemory"

	shortLinkLength = 10
	allowedChars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	maxRetries      = 5
)

var (
	cfg config.Config
)

func main() {
	envconfig.MustProcess("", &cfg)

	if err := logger.BuildLogger(cfg.LogLevel); err != nil {
		log.Fatal("cannot build logger: ", err)
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	var linkRepo model.ILinkRepository
	var linkUsecase model.ILinkUsecase

	switch cfg.DbType {
	case dbTypePostgres:
		zap.L().Info("Using PostgreSQL database")
		zap.L().Info("PostgresURI: ", zap.String("uri", cfg.PostgresURI))
		zap.L().Info("MigrationsDir: ", zap.String("dir", cfg.MigrationsDir))

		time.Sleep(3 * time.Second)

		migrator := postgresql.NewMigrator(cfg.PostgresURI, cfg.MigrationsDir)
		if err := migrator.Apply(); err != nil {
			log.Fatal("cannot apply migrations: ", err)
		}

		sqlxConn, err := sqlx.Connect("postgres", cfg.PostgresURI)
		if err != nil {
			log.Fatal("cannot connect to postgres db: ", err)
		}

		sqlxConn.SetMaxOpenConns(100)
		sqlxConn.SetMaxIdleConns(100)
		sqlxConn.SetConnMaxLifetime(5 * time.Minute)

		defer sqlxConn.Close()

		sqalxConn, err := sqalx.New(sqlxConn)
		if err != nil {
			log.Fatal("cannot connect to postgres db: ", err)
		}
		defer sqalxConn.Close()

		zap.L().Info("Database manage was process successfully")

		linkRepoImpl := _linkRepo.New(sqalxConn)
		linkUsecase = _linkUsecase.New(linkRepoImpl)
	case dbTypeInMemory:
		zap.L().Info("Using In-Memory database")

		linkRepo = inmemory.NewInMemoryLinkRepository()
		linkUsecase = _linkUsecase.New(linkRepo)
	default:
		log.Fatalf("Invalid DB_TYPE: %s", cfg.DbType)
	}

	appHandler := handler.New(
		linkUsecase,
		rng,
		maxRetries,
		shortLinkLength,
		httpVersion,
	)

	chain := alice.New(appHandler.WsMiddleware).Then(appHandler)
	if chain == nil {
		fmt.Println(chain)
	}
	server := http.Server{
		Handler: chain,
		Addr:    ":" + cfg.Addr,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	zap.L().Info("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	zap.L().Info("Server exiting")
}
