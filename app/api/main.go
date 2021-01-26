package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
	"github.com/orvosi/api/internal/builder"
	"github.com/orvosi/api/internal/config"
	"github.com/orvosi/api/internal/http/middleware"
	"github.com/orvosi/api/internal/http/server"
	"github.com/orvosi/api/internal/tool"
)

const (
	dbDriver = "postgres"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	db, err := builder.BuildSQLDatabase(dbDriver, cfg)
	checkError(err)

	jwtDec := tool.NewIDTokenDecoder(cfg.Google.Audience)
	jwtMidd := middleware.WithJWTDecoder(jwtDec.Decode)

	medRecCreator := builder.BuildMedicalRecordCreator(cfg, db)

	srv := server.NewServer(jwtMidd, medRecCreator)
	runServer(srv, cfg.Port)
	waitForShutdown(srv)
}

func runServer(srv *server.Server, port string) {
	go func() {
		if err := srv.Start(fmt.Sprintf(":%s", port)); err != nil {
			srv.Logger.Info("shutting down the server")
		}
	}()
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
