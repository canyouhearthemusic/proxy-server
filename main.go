package main

import (
	"net/http"

	"github.com/canyouhearthemusic/proxy-server/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

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
