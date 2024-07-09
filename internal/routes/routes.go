package routes

import (
	"github.com/canyouhearthemusic/proxy-server/internal/handlers"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterRoutes(r *chi.Mux) {
	r.Post("/request", handlers.HandleProxyRequest)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("https://proxy-server-u78h.onrender.com/swagger/doc.json"),
	))
}
