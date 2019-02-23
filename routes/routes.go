package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../db"
	"../models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandleGetAll returns all documents
func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	payments, err := db.GetAllPayments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Unable to get payments: %+v", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	e := json.NewEncoder(w)
	e.Encode(payments)
}

// HandleGetByID returns a single document
func HandleGetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID to Object ID: %+v", err)
	}

	payment, err := db.GetPaymentByID(objectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Unable to get payment: %+v", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	e := json.NewEncoder(w)
	e.Encode(payment)
}

// HandleInsert inserts a document according to the request body
func HandleInsert(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload. If you have specified an ID, remove it and let the system generate one."))
		return
	}

	defer r.Body.Close()

	_, err := db.InsertPayment(payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleUpdateByID updates a document
func HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID to Object ID: %+v", err)
	}

	var payment models.Payment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	defer r.Body.Close()

	err = db.UpdatePaymentByID(objectID, payment)
	if err != nil {
		log.Printf("Error updating payment: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

// HandleDelete deletes a document by an ID
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID to Object ID: %+v", err)
	}

	deleteCount, err := db.DeletePaymentByID(objectID)
	if err != nil {
		log.Printf("Error deleting payment: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if deleteCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ID not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
