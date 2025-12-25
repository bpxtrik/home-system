package api

import (
	"context"
	"encoding/json"
	"home-system/internal"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	DB *pgxpool.Pool
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func (h Handler) HealthCheck(w http.ResponseWriter, req *http.Request) {
	resp := internal.Response{Detail: "OK"}
	writeJSON(w, http.StatusOK, resp)
}

func (h Handler) MotionTrigger(w http.ResponseWriter, req *http.Request) {
	var mr internal.MotionRequest
	err := json.NewDecoder(req.Body).Decode(&mr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if mr.AccessKey != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if mr.Timestamp.IsZero() {
		resp := internal.Response{Detail: "Invalid Request"}
		writeJSON(w, http.StatusBadRequest, resp)
		return 
	}

	m := internal.Motion{
		ID: uuid.New(),
		Timestamp: mr.Timestamp,
	}

	stmt := `INSERT INTO motions (id, timestamp)
			VALUES ($1, $2) `

	_, err = h.DB.Exec(context.Background(), stmt, m.ID, m.Timestamp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := internal.Response{Detail: "Saved"}
	writeJSON(w, http.StatusCreated, resp)
}
