package infrastructure

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionPostgres(dbUrl string) *gorm.DB {
	postgresDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected successfully to Postgres")

	return postgresDB
}

func ConnectMongo(dbUrl string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(dbUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
