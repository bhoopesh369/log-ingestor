package config

import (
	"context"
	"fmt"

	// "github.com/bhoopesh369/log-injestor/models"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// ConnectDB connects to MongoDB
func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://root:password@log_injestor_mg_db:27017/?authSource=admin")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println(color.RedString("Error connecting to MongoDB"))
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(color.RedString("Failed to ping MongoDB"))
		panic(err)
	}

	db = client.Database("log_db")

	fmt.Println(color.GreenString("Connected to MongoDB"))
}

// GetDB returns the database instance
func GetDB() *mongo.Database {
	return db
}

// Migrations
func MigrateDB() {
	db := GetDB()
	
	logCollection := db.Collection("logs")

	// Create an index on the 'timestamp' field
	_, err := logCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    map[string]interface{}{"timestamp": 1},
			Options: options.Index().SetUnique(false),
		},
	)
	if err != nil {
		panic(err)
	}

	// Create an index on the 'resourceId' field
	_, err = logCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    map[string]interface{}{"resourceId": 1}, // 1 for ascending order, -1 for descending
			Options: options.Index().SetUnique(false),        // Optional: Set additional options for the index
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(color.BlueString("Indexes created successfully"))

}
