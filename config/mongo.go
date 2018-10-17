package config

import (
	"fmt"
	"os"
	mgo "gopkg.in/mgo.v2"
)

// GetMongoDb ...Method to set ang return the database for connection
func GetMongoDb() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")

	fmt.Println("Connecting to Mongo...")
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)
	return db, nil
}
