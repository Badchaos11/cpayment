package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Badchaos11/cpayment/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "Constanta Payment Gateway ", log.LstdFlags)

	tr := handlers.NewTransactions(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/onetrbid/{id}", tr.GetOneById)
	getRouter.HandleFunc("/alltrbid/{userid}", tr.GetAllById)
	getRouter.HandleFunc("/alltrbem/{email}", tr.GetAllByEmail)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/create", tr.CreateTransaction)

	puthRouter := sm.Methods("PUT").Subrouter()
	puthRouter.HandleFunc("/reject/{id}", tr.RejectTransaction)

	protectedRouter := sm.Methods("PUT").Subrouter()
	protectedRouter.HandleFunc("/changest/{id}", tr.ChangeTransactionStatus)

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
