package conn

import (
	"context"
	"quiz_app/infra/config"
	log "quiz_app/infra/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func ConnectDb() {
	conf := config.Db()

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db = client

	log.Info("Connected to MongoDB successful...")
}

func Db() *mongo.Client {
	return db
}
