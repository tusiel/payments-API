package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan

		appCleanup()
		os.Exit(1)
	}()

	start()
}

func appCleanup() {
	log.Println("Shutting down server...")
}

func start() {
	router := mux.NewRouter()

	/** API Routes */
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/payments", simpleHandler).Methods("GET")
	api.HandleFunc("/payments/{id}", simpleHandler).Methods("GET")
	api.HandleFunc("/payment", simpleHandler).Methods("POST")
	api.HandleFunc("/payment/{id}", simpleHandler).Methods("PUT")
	api.HandleFunc("/payment/{id}", simpleHandler).Methods("DELETE")

	var handler http.Handler
	handler = router
	handler = handlers.LoggingHandler(os.Stdout, handler)

	srv := &http.Server{
		Handler:      handler,
		Addr:         config.GetString("listenAddress"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}
