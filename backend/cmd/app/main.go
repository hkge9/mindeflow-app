package main

import (
	"log"
	"mindeflow-app/backend/internal/config"
	"mindeflow-app/backend/internal/db"
	"mindeflow-app/backend/internal/user/repository"
	"mindeflow-app/backend/internal/user/service"
	"mindeflow-app/backend/internal/user/transport"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	pool, err := db.NewPostgresPool(cfg.DatabaseURL())

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepo := repository.NewPostgresRepository(pool)
	userService := service.New(userRepo)
	userHandler := transport.NewHandler(userService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		transport.RegisterRoutes(r, userHandler)
	})

	addr := ":" + cfg.AppPort
	log.Printf("server started on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
