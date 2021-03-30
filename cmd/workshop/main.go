package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/api/jokes"
	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)

	if err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient, cfg.CustomJoke)
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	log.Printf("starting server %s", path)
	log.Fatal(http.ListenAndServe(path, r))
	log.Print("shutting down")
}
