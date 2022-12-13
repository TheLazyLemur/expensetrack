package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	Storer     Storer
}

func NewAPIServer(listenAddr string, userStorer Storer) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		Storer:     userStorer,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/users", makeHttpHandleFunc(s.handleUser))
	router.HandleFunc("/users/{id}", makeHttpHandleFunc(s.handleUserById))
	router.HandleFunc("/expenses", makeHttpHandleFunc(s.handleExpense))
	router.HandleFunc("/expenses/{id}", makeHttpHandleFunc(s.handleExpenseById))

	log.Println("Starting API server on", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error
type apiError struct {
	Err string `json:"error"`
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err)
			err := WriteJson(w, http.StatusInternalServerError, apiError{Err: err.Error()})
			if err != nil {
				log.Println(err)
			}
		}
	}
}
