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
	// Создание логера
	tr := handlers.NewTransactions(l)
	// Создание роутера
	sm := mux.NewRouter()
	// Регистрация отдельных роутеров на различные виды HTTP методов
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/onetransaction/{id}", tr.GetOneById)
	getRouter.HandleFunc("/allbyuserid/{userid}", tr.GetAllById)
	getRouter.HandleFunc("/allbyemail/{email}", tr.GetAllByEmail)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/create", tr.CreateTransaction)
	postRouter.HandleFunc("/createfail", tr.CreateTransactionFail)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/reject", tr.RejectTransaction)

	protectedRouter := sm.Methods("PUT").Subrouter()
	protectedRouter.HandleFunc("/changestatus", tr.ChangeTransactionStatus)
	protectedRouter.Use(tr.MiddlewareAuth)
	// Настрйока сервера
	s := &http.Server{
		Addr:         ":9090",           // Порт сервера
		Handler:      sm,                // Хэндлеры
		ErrorLog:     l,                 // Логи
		ReadTimeout:  5 * time.Second,   // Таймаут запроса клиента
		WriteTimeout: 10 * time.Second,  // Таймаут ответа клиенту
		IdleTimeout:  120 * time.Second, // Таймаут соединения в простое
	}
	// Запуск сервера
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	// Отключение
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Println("Something went wrong with shutdown:", err)
	}
	s.Shutdown(ctx)
}
