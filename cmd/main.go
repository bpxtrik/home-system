package main

import (
	"home-system/internal/api"
	"home-system/internal/db"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)


func main() {
	_ = godotenv.Load()

	pool, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	
	h := &api.Handler{DB: pool}

	mux := api.RegisterRoutes(h)
	http.ListenAndServe(":8090", mux)
}
