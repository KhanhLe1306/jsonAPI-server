package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ApiServer struct {
	listenAddr string
	store *PostgresStore
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJson(w http.ResponseWriter, status int, v any) error{
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandleFunc(f ApiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewApiServer(listenAddr string, store *PostgresStore) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store: store,
	}
} 

func (s *ApiServer) Run(){
	s.store.Init()
	
	http.HandleFunc("GET /", makeHttpHandleFunc(s.handleGetAccount))
	http.HandleFunc("DELETE /", makeHttpHandleFunc(s.handleDeleteAccount))

	log.Fatal(http.ListenAndServe(s.listenAddr, nil))
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error{
	account := NewAccount("Khanh", "Le")
	WriteJson(w, http.StatusOK, account)
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	fmt.Println("This is Delete method")	
	return nil
}

func (s *ApiServer) hanleCreateAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}








