package main

import (
	"context"
	"log"
	"logger/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (rpc *RPCServer) NewLog(payload RPCPayload, rsp *string) error {
	log.Printf("client: %+v", client)
	col := client.Database("logs").Collection("logs")
	_, err := col.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*rsp = "Processed payload via rpc: " + payload.Name
	return nil
}
