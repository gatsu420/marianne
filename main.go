package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gatsu420/marianne/app/handlers"
	"github.com/gatsu420/marianne/app/repository"
	"github.com/gatsu420/marianne/app/usecases/food"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pgPool, err := pgxpool.New(context.Background(), "postgres://mary:mary@localhost:5432/marydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	pgRepo := repository.NewPGRepo(pgPool)
	foodUsecases := food.NewUsecase(pgRepo)
	handlers := handlers.NewHandler(foodUsecases)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/food", handlers.GetFood)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	intrCh := make(chan os.Signal, 1)
	signal.Notify(intrCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-intrCh
	log.Println("stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
