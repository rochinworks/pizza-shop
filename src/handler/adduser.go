package handler

import (
	"encoding/json"
	"fmt"
	"go-web-api/src/pg"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (handler *handler) AddUserHandler() http.HandlerFunc {
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
		user := pg.User{}
		err = json.Unmarshal(req, &user)
		if err != nil {
			log.Error("could not unmarshal json correctly: ", err)
			// respond with 400s
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}

		// save to db
		userID, err := handler.Repo.Store(r.Context(), user)
		if err != nil {
			log.Error("could not store user: ", err)
			// respond with 500s
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		}

		response := map[string]interface{}{
			"id": userID,
		}

		// respond with id and 200s
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
