package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"logger/data"
	"net/http"
	"time"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	gRpcPort = "50001"
	mongoUrl = "mongodb://mongo:27017"
)

type Config struct {
	Models data.Models
}

func main() {
	client, err := connectToMongo()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	fmt.Println("starting web server on port ", webPort)
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

//func (app *Config) serve() {
//	server := &http.Server{
//		Addr:    fmt.Sprintf(":%s", webPort),
//		Handler: app.routes(),
//	}
//
//	fmt.Println("starting web server on port ", webPort)
//	err := server.ListenAndServe()
//	if err != nil {
//		log.Panic(err)
//	}
//}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Printf("couldn't connect to mongoDB: %v", err)
		return nil, err
	}
	return client, nil
}
