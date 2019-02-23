package db

import (
	"log"
	"os"
	"testing"

	"../models"
)

func TestMain(m *testing.M) {
	retCode := m.Run()

	err := deleteAll()
	if err != nil {
		log.Printf("Error deleting all records after tests: %+v", err)
	}

	os.Exit(retCode)

}
func TestGetAll(t *testing.T) {
	var err error

	payment1 := models.Payment{ID: "1"}
	payment2 := models.Payment{ID: "2"}
	payment3 := models.Payment{ID: "3"}
	payment4 := models.Payment{ID: "4"}

	_, err = InsertPayment(payment1)
	_, err = InsertPayment(payment2)
	_, err = InsertPayment(payment3)
	_, err = InsertPayment(payment4)
	if err != nil {
		t.Errorf("Insert failed: %+v", err)
	}

	payments, err := GetAllPayments()
	if err != nil {
		t.Errorf("GetAll failed: %+v", err)
	}

	if len(payments) != 4 {
		t.Errorf("Expected 4 payments but got %d", len(payments))
	}
}

func TestGetByID(t *testing.T) {
	var err error

	payment1 := models.Payment{ID: "1"}
	payment2 := models.Payment{ID: "id-to-test", Type: "Payment", OrganisationID: "organisation-to-test"}

	_, err = InsertPayment(payment1)
	_, err = InsertPayment(payment2)
	if err != nil {
		t.Errorf("Insert failed: %+v", err)
	}

	payment, err := GetPaymentByID("id-to-test")
	if err != nil {
		t.Errorf("GetAll failed: %+v", err)
	}

	if payment.ID != "id-to-test" {
		t.Errorf("Expected ID to be %s, but got '%s'", "id-to-test", payment.ID)
	}

	if payment.Type != "Payment" {
		t.Errorf("Expected Type to be %s, but got '%s'", "Payment", payment.Type)
	}

	if payment.OrganisationID != "organisation-to-test" {
		t.Errorf("Expected OrganisationID to be %s, but got '%s'", "organisation-to-test", payment.OrganisationID)
	}
}

func TestInsert(t *testing.T) {
	payment := models.Payment{
		Type:           "Payment",
		OrganisationID: "123456",
	}

	_, err := InsertPayment(payment)
	if err != nil {
		t.Errorf("Insert failed: %+v", err)
	}
}

func TestUpdatePaymentByID(t *testing.T) {
	payment := models.Payment{
		ID:             "id-to-test",
		Type:           "Payment",
		OrganisationID: "shouldnt-change",
	}

	_, err := InsertPayment(payment)
	if err != nil {
		t.Errorf("Insert failed: %+v", err)
	}

	updatedPayment := models.Payment{
		Type: "SomethingElse",
	}

	err = UpdatePaymentByID("id-to-test", updatedPayment)
	if err != nil {
		t.Errorf("Update failed: %+v", err)
	}

	p, err := GetPaymentByID("id-to-test")

	if p.ID != "id-to-test" {
		t.Error("ID updated when it shouldn't have")
	}

	if p.OrganisationID != "shouldnt-change" {
		t.Error("OrganisationID updated when it shouldn't have")
	}

	if p.Type != "SomethingElse" {
		t.Errorf("Type should have updated to %s, but was %s", "SomethingElse", p.Type)
	}

}

func TestDeleteByID(t *testing.T) {
	payment := models.Payment{
		ID:             "id-to-test",
		Type:           "Payment",
		OrganisationID: "org-id",
	}

	insertedID, err := InsertPayment(payment)
	if err != nil {
		t.Errorf("Insert failed: %+v", err)
	}

	deleteCount, err := DeletePaymentByID(insertedID.(string))
	if err != nil {
		t.Errorf("Delete failed: %+v", err)
	}

	if deleteCount != 1 {
		t.Errorf("Expected delete count to be 1, but got %d", deleteCount)
	}

}
