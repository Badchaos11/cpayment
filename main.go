package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Badchaos11/cpayment/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	tr := handlers.NewTransactions(l)

	Db, err := sql.Open("postgres", "host=localhost dbname=testdb user=user password=secret sslmode=disable")

	if err != nil {
		log.Fatal("Cannot connect to databaase")
	}
	defer Db.Close()

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/onetrbid/{id}", tr.GetOneById)
	getRouter.HandleFunc("/alltrbid/{id_user}", tr.GetAllById)
	getRouter.HandleFunc("/alltrbem/{email}", tr.GetAllByEmail)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/create", tr.CreateTransaction)

	patchRouter := sm.Methods("PATCH").Subrouter()
	patchRouter.HandleFunc("/reject/{id}", tr.RejectTransactions)
	patchRouter.HandleFunc("/changest/{id}", tr.ChangeTransactionStatus)

	s := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
}
