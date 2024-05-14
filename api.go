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

	http.HandleFunc("GET /accounts", makeHttpHandleFunc(s.handleGetAccounts))
	http.HandleFunc("GET /accounts/{id}", makeHttpHandleFunc(s.handleGetAccountById))
	http.HandleFunc("POST /accounts", makeHttpHandleFunc(s.handleCreateAccount))
	http.HandleFunc("DELETE /accounts/{id}", makeHttpHandleFunc(s.handleDeleteAccount))

	log.Fatal(http.ListenAndServe(s.listenAddr, nil))
}

func (s *ApiServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error{
	query := "select * from accounts"
	rows, err := s.store.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var account Account	
		if err := rows.Scan(&account.Id, &account.FirstName, &account.LastName, &account.Balance); err != nil {
			log.Fatal(err)
			return err
		}
		log.Printf("id: %v, firstName: %s, lastName: %s, balance: %v", account.Id, account.FirstName, account.LastName, account.Balance)
	}
	return nil
}

func (s *ApiServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error{
	id := r.PathValue("id")	
	if len(id) > 0 {
		fmt.Println("id is", id)
	}
	return nil
}


func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	fmt.Println("This is Delete method")	
	return nil
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error{
	decoder := json.NewDecoder(r.Body)
	var body AddAccountRequest
	if err := decoder.Decode(&body); err != nil {
		return err
	}
	
	query := fmt.Sprintf("insert into accounts (first_name, last_name, balance) values ('%v', '%v', %v)", body.FirstName, body.LastName, 0) 
	
	if _, err := s.store.db.Exec(query); err != nil {
		return err 
	}
	return nil

}








