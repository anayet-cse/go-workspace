package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":get called}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_type", "applicationi/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message":"put called"}`))
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/anayet/ullah").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8080", r))
}
