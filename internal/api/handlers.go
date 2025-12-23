package api

import (
	"context"
	"encoding/json"
	"home-system/internal"
	"net/http"

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
	var m internal.Motion
	err := json.NewDecoder(req.Body).Decode(&m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	m.ID = uuid.New()

	stmt := `INSERT INTO motions (id, timestamp)
			VALUES ($1, $2) `

	_, err = h.DB.Exec(context.Background(), stmt, m.ID, m.Timestamp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := internal.Response{Detail: "Saved"}
	writeJSON(w, http.StatusCreated, resp)
}
