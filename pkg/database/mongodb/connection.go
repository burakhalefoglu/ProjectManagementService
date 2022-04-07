package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func GetMongodbClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url = "mongodb://" + os.Getenv("MONGODB_USER") + ":" + os.Getenv("MONGODB_PASS") + "@" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("MongoConnection", "ConnectMongodb", err)
	}
	return client
}
