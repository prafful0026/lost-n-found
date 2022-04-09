package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDB() *mongo.Client{

	client,err:=mongo.NewClient(options.Client().ApplyURI(getEnvVar(mongoUriKey)))

	if err!=nil{
		log.Fatal(err)
	}

    ctx,_:=context.WithTimeout(context.Background(),connectionTime*time.Second)

	err = client.Connect(ctx)

	if err!=nil{
		log.Fatal(err)
	}

    err = client.Ping(ctx, nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB")

    return client
}

func DisconnectDB(client *mongo.Client){
    ctx,_:=context.WithTimeout(context.Background(),connectionTime*time.Second)
    client.Disconnect(ctx)
}

func GetCollections(client *mongo.Client, collectionName string) *mongo.Collection{
	collections:=client.Database("lost-n-found").Collection(collectionName)
	return collections
}