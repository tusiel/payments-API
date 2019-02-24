package db

import (
	"context"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"../config"
	"../models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var client *mongo.Client

func init() {
	var err error
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetString("database.connectionAddress")))
	if err != nil {
		log.Fatalf("Error connecting to database")
	}
}

// GetAllPayments returns an array of all payments
func GetAllPayments() (payments []models.Payment, err error) {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collection"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var payment models.Payment
		err := cur.Decode(&payment)
		if err != nil {
			log.Printf("Error decoding payment: %+v", err)
		}
		payments = append(payments, payment)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return
}

// GetPaymentByID returns a single payment by it's ID
func GetPaymentByID(id string) (payment models.Payment, err error) {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collection"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&payment)
	if err != nil {
		return
	}

	return

}

// InsertPayment takes a Payment and adds it to the database
func InsertPayment(payment models.Payment) (insertedID interface{}, err error) {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collection"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, payment)
	if err != nil {
		return "", err
	}

	insertedID = res.InsertedID

	return
}

// UpdatePaymentByID takes an ID and a Payment model and updates it
func UpdatePaymentByID(id string, payment models.Payment) (err error) {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collection"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": payment})
	if err != nil {
		return err
	}

	return
}

// DeletePaymentByID takes an ID and deletes the payment
func DeletePaymentByID(id string) (deleteCount int64, err error) {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collection"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return 0, err
	}

	deleteCount = res.DeletedCount

	return
}

func deleteAll() error {
	collection := client.Database(config.GetString("database.name")).Collection(config.GetString("database.collection"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}
