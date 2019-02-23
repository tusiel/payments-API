package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./config"
	"./middleware"
	"./routes"

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

	api.HandleFunc("/payments", routes.HandleGetAll).Methods("GET")
	api.HandleFunc("/payment/{id}", routes.HandleGetByID).Methods("GET")
	api.HandleFunc("/payment", routes.HandleInsert).Methods("POST")
	api.HandleFunc("/payment/{id}", routes.HandleUpdateByID).Methods("PUT")
	api.HandleFunc("/payment/{id}", routes.HandleDelete).Methods("DELETE")

	var handler http.Handler
	handler = router
	handler = handlers.LoggingHandler(os.Stdout, handler)
	handler = middleware.RemoveTrailingSlash(handler)

	srv := &http.Server{
		Handler:      handler,
		Addr:         config.GetString("listenAddress"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
