package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"github.com/todo_app/db"
	"github.com/todo_app/handlers"
	"github.com/todo_app/services"
)

type Application struct {
	Models services.Models
}

func main() {
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	services.New(mongoClient)

	log.Println("Server running in port", 8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))

}

