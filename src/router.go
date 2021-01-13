package main

import (
	"go-web-api/src/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func httpRouter() chi.Router {

	handler := handler.NewHandler()

	// chi router is easy to use and lightweight
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// base routes
	r.Get("/", handler.BaseHandler())

	return r
}
