package api

import (
	"context"
	"encoding/json"
	"fmt"
	"home-system/internal"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	DB *pgxpool.Pool
}

var sessions = struct {
	sync.RWMutex
	store map[string]string // token -> username
}{store: make(map[string]string)}

func writeJSON(w *http.ResponseWriter, status int, body any) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(body)
}

func checkCookie(w *http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session_token")
	if err != nil {
		http.Error(*w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessions.RLock()
	username := sessions.store[cookie.Value]
	sessions.RUnlock()

	if username == "" {
		http.Error(*w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (h Handler) HealthCheck(w http.ResponseWriter, req *http.Request) {
	resp := internal.Response{Detail: "OK"}
	writeJSON(&w, http.StatusOK, resp)
}

func (h Handler) MotionTrigger(w http.ResponseWriter, req *http.Request) {
	var mr internal.MotionRequest
	err := json.NewDecoder(req.Body).Decode(&mr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mr.Timestamp.IsZero() {
		resp := internal.Response{Detail: "Invalid Request"}
		writeJSON(&w, http.StatusBadRequest, resp)
		return
	}

	m := internal.Motion{
		ID:        uuid.New(),
		Timestamp: mr.Timestamp,
	}

	stmt := `INSERT INTO motions (id, timestamp) VALUES ($1, $2)`
	_, err = h.DB.Exec(context.Background(), stmt, m.ID, m.Timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := internal.Response{Detail: "Saved"}
	writeJSON(&w, http.StatusCreated, resp)
}

func (h Handler) Login(w http.ResponseWriter, req *http.Request) {
	var lr internal.LoginRequest
	err := json.NewDecoder(req.Body).Decode(&lr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if lr.Username != os.Getenv("USERNAME") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	hashedPassword := os.Getenv("PASSWORD_HASH")

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(lr.Password)); err != nil {
		fmt.Println("Failing here")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := uuid.NewString()
	sessions.Lock()
	sessions.store[token] = lr.Username
	sessions.Unlock()

	isProd := os.Getenv("ENV") == "production"

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(2 * time.Hour),
		HttpOnly: true,
		Secure:   isProd,
		Path:     "/",
		SameSite: func() http.SameSite {
			if isProd {
				return http.SameSiteNoneMode
			}
			return http.SameSiteLaxMode
		}(),
	}

	http.SetCookie(w, &cookie)
	resp := internal.Response{Detail: "Login successful"}
	writeJSON(&w, http.StatusOK, resp)
}

func (h Handler) GetMotion(w http.ResponseWriter, req *http.Request) {
	checkCookie(&w, req)
	stmt := `SELECT * FROM motions m ORDER BY m.timestamp DESC`

	rows, err := h.DB.Query(context.Background(), stmt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var motions []internal.Motion
	for rows.Next() {
		var id uuid.UUID
		var timestamp time.Time

		err := rows.Scan(&id, &timestamp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		motion := internal.Motion{
			ID:        id,
			Timestamp: timestamp,
		}
		motions = append(motions, motion)
	}

	writeJSON(&w, http.StatusOK, motions)

}

func (h Handler) Auth(w http.ResponseWriter, req *http.Request) {
	checkCookie(&w, req)
}
