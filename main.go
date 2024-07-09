package main

import (
	"net/http"

	_ "github.com/canyouhearthemusic/proxy-server/docs"
	"github.com/canyouhearthemusic/proxy-server/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      https://proxy-server-u78h.onrender.com

func main() {
	log.Println("Starting App...")

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.Heartbeat("/health"))

	routes.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
