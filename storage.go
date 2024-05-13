package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount (*Account) error
	DeleteAccount (id uint) error
	UpdateAccount (*Account) error
	GetAccountById (id uint) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error){
	connStr := "user=postgres dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error{
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error{
	query := `create table if not exists accounts (
			id  SERIAL 	PRIMARY KEY,
			first_name	VARCHAR (50),
			last_name 	VARCHAR	(50),
			balance		DECIMAL	
		)
	`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}















