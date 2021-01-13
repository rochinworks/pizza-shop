package handler

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

//    // decode request
//		json.NewDecoder(r.Body).Decode(&request)
//		ioutil.ReadAll(r.Body)
//		defer r.Body.Close()
//
//		// Return the response
//		// set the header
//		w.Header().Set("Content-Type", "application/json")
//
//    // encode response
//		json.NewEncoder(w).Encode(response)
