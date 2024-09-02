package database

import (
    "context"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
    MONGODB_URI := os.Getenv("MONGODB_URI")
    clientOptions := options.Client().ApplyURI(MONGODB_URI)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }
    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }

    return client.Database("goland_db"), nil
}
