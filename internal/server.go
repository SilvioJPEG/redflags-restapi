package apiserver

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Router *mux.Router
	DB     *sql.DB
}

func Initialize(config *Config) *APIServer {
	DB, err := newDB(config.connectionStringDB)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	server := &APIServer{
		Router: mux.NewRouter(),
		DB:     DB,
	}
	server.configureRouter()
	return server
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Successfully connected to redflags database")
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *APIServer) configureRouter() {
	gameRouter := s.Router.PathPrefix("/game").Subrouter()
	gameRouter.HandleFunc("", s.getGameInfo()).Methods("GET")
	gameRouter.HandleFunc("", s.updateGame()).Methods("PATCH")
	gameRouter.HandleFunc("", s.deleteGameRoom()).Methods("DELETE")

	playerRouter := s.Router.PathPrefix("/player").Subrouter()
	playerRouter.HandleFunc("", s.getPlayerInfo()).Methods("GET")
}
