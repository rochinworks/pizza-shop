package main

import (
	"net/http"

	_ "github.com/golang/mock/mockgen/model"
)

//Handler interface makes mocking out the http layer in tests easier
//go:generate mockgen -source=handlerinterface.go -destination=mocks/handlerinterface.go -package=mocks
type Handler interface {
	BaseHandler() http.HandlerFunc
}
