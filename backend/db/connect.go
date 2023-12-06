package db

import (
	"context"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Dbconnect -> connects mongo
func Dbconnect() (*mongo.Client, error) {
	
	clientOptions := options.Client().ApplyURI("mongodb+srv://bloodbankadmin:bloodbank@cluster0.ggcow1t.mongodb.net/bloodbank")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	color.Green("⛁ Connected to Database")
	return client, nil
}
