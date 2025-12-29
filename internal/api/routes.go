package api

import "net/http"

func RegisterRoutes(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("POST /motion", h.MotionTrigger)
	mux.HandleFunc("POST /login", h.Login)
	mux.HandleFunc("GET /motions", h.GetMotion)
	return mux
}
