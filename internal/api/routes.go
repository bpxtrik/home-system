package api

import "net/http"

func RegisterRoutes(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("POST /motion", h.MotionTrigger)

	return mux
}
