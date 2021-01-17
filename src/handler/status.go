package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func (handler *handler) StatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get params from url
		orderID := r.URL.Query().Get("orderId")
		userID := r.URL.Query().Get("userId")

		oID, err := uuid.Parse(orderID)
		if err != nil {
			log.Error("could not verify order id: ", err)
			// respond with 400s
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}

		uID, err := uuid.Parse(userID)
		if err != nil {
			log.Error("could not verify user id: ", err)
			// respond with 400s
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}

		// save to db
		status, err := handler.Repo.GetStatus(r.Context(), uID, oID)
		if err != nil {
			log.Error("could not get order status: ", err)
			// respond with 500s
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}

		response := map[string]interface{}{
			"orderId": orderID,
			"status":  status,
		}

		// respond with id and 200s
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
