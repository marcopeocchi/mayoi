package main

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcopeocchi/mayoi/internal"
	"github.com/marcopeocchi/mayoi/internal/management"
	"github.com/marcopeocchi/mayoi/internal/middleware"
	"github.com/marcopeocchi/mayoi/internal/registry"
	"github.com/marcopeocchi/mayoi/pkg/config"
	_ "modernc.org/sqlite"
)

var (
	configPath   string
	databasePath string
	adderess     string
	port         int

	//go:embed ui/dist/index.html
	//go:embed ui/dist/assets/*
	ui embed.FS
)

func main() {
	flag.StringVar(&configPath, "c", "./config.yml", "config path")
	flag.StringVar(&databasePath, "d", "mayoi.db", "database path")
	flag.StringVar(&adderess, "bind", "0.0.0.0", "bind to address")
	flag.IntVar(&port, "p", 6969, "port to listen at")
	flag.Parse()

	if err := config.Instance().Load(configPath); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	config.Instance().Database = databasePath
	config.Instance().Address = adderess
	config.Instance().Port = port

	db, err := sql.Open("sqlite", databasePath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}

	run(db)
}

func run(db *sql.DB) {
	reg := registry.New()
	mux := http.NewServeMux()

	management.Module(mux, reg)

	uifs, err := fs.Sub(ui, "ui/dist")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(3)
	}

	mux.Handle("/", http.FileServer(http.FS(uifs)))

	for _, url := range config.Instance().Indexers {
		indexer, err := internal.IndexerFactory(url, db, reg, mux)
		if err != nil {
			slog.Warn(
				"Skipping indexer",
				slog.String("url", url),
				slog.String("err", err.Error()),
			)
			continue
		}

		go indexer.AutoIndex(context.Background(), time.Minute*5)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", adderess, port),
		Handler: middleware.CORS(mux),
	}

	go gracefulShutdown(server, db)

	server.ListenAndServe()
}

func gracefulShutdown(s *http.Server, db *sql.DB) {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()
		slog.Info("shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)

		defer func() {
			db.Close()
			stop()
			cancel()
			fmt.Println("shutdown completed")
		}()

		s.Shutdown(ctxTimeout)
	}()
}
