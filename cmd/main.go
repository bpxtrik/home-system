package main

import (
	"home-system/internal/api"
	"home-system/internal/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	http.ListenAndServe("0.0.0.0:"+port, mux)
}
