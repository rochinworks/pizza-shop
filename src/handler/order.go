package handler

import (
	"encoding/json"
	"fmt"
	"go-web-api/src/pg"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (handler *handler) PizzaHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read request
		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("could not read request body: ", err)
			// respond with 400s
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}
		defer r.Body.Close()

		// parse name
		pizza := pg.Pizza{
			Status: "starting",
		}

		err = json.Unmarshal(req, &pizza)
		if err != nil {
			log.Error("could not unmarshal json: ", err)
			// respond with 400s
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}

		// save to db
		orderID, err := handler.Repo.StoreOrder(r.Context(), pizza)
		if err != nil {
			log.Error("could not store user: ", err)
			// respond with 500s
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}

		// start async pizza processing

		response := map[string]interface{}{
			"orderId": orderID,
			"status":  pizza.Status,
		}

		// respond with id and 200s
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
