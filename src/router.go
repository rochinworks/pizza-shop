package main

import (
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

	// base routes
	r.Get("/", handler.BaseHandler())
	// create a user
	r.Post("/user/create", handler.AddUserHandler())
	// order a pizza
	r.Post("/order/start", handler.PizzaHandler())
	// check the order status
	r.Get("/order/status?userId={userId}&orderId={orderId}", handler.StatusHandler())

	return r
}
