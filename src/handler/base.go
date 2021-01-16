package handler

import (
	"encoding/json"
	"net/http"
)

// BaseHandler here's where the incoming read request will be received
func (handler *handler) BaseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"alive": true})
	}
}
