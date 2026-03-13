package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"region-fixer/internal/config"
	"region-fixer/internal/infrastructure/database"
	"region-fixer/internal/router"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config:", err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := database.NewConnection(ctx, cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect db: %w", err)
	}
	defer db.Close()

	r := router.SetupRouter()

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	fmt.Printf("server starting on %s", addr)

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("server forced to shutdown: %v", err)
	} else {
		log.Printf("server exited gracefully")
	}

	fmt.Println("application stopped")
}
