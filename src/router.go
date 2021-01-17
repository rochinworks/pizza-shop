package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func httpRouter(handler Handler) chi.Router {

	// chi router is easy to use and lightweight
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// base routes
	r.Get("/", handler.BaseHandler())
	// create a user
	r.Post("/user/create", handler.AddUserHandler())
	// order a pizza
	r.Post("/order/start", handler.PizzaHandler())
	// check the order status
	r.Get("/order/status", handler.StatusHandler())

	r.MethodNotAllowed(methodNotAllowedHandler)
	r.NotFound(notFoundHandler)

	return r
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
}
