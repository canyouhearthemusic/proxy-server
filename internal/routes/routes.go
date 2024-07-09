package routes

import (
	"github.com/canyouhearthemusic/proxy-server/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Post("/request", handlers.HandleProxyRequest)
}
