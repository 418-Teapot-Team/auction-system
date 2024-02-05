package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	server "auction-system"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	"auction-system/internal/web/handlers/auth"
	"auction-system/internal/web/routes"
	"auction-system/pkg/repository"

	"github.com/BoryslavGlov/logrusx"
)

func main() {
	if err := gotenv.Load(); err != nil {
		log.Fatal(err)
	}

	logx, err := logrusx.New("auction-system")
	if err != nil {
		log.Fatal(err)
	}
	DbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := repository.NewDB(DbUrl)
	if err != nil {
		logx.Fatal("error while trying to create NewDB",
			logrusx.LogField{Key: "context", Value: err},
		)
	}

	app := gin.Default()

	v1 := app.Group("/api/v1")

	authRepo := repository.NewAuthRepository(db)

	authHandler := auth.NewHandler(logx, authRepo)

	routes.AuthRouters(v1, authHandler)

	srv := new(server.Server)

	go func() {
		if err = srv.Run(os.Getenv("PORT"), app); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logx.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logx.Info("App Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
}
