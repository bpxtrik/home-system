package main

import (
	"home-system/internal/api"
	"home-system/internal/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	pool, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	h := &api.Handler{DB: pool}

	mux := api.RegisterRoutes(h)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	http.ListenAndServe("0.0.0.0:"+port, c.Handler(mux))
}
