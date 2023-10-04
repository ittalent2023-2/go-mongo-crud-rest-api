package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/parsaakbari1209/go-mongo-crud-rest-api/http"
	"github.com/parsaakbari1209/go-mongo-crud-rest-api/repository"
)

func main() {

	db := os.Getenv("MONGO_DB_HOST")

	dbURL := fmt.Sprintf("mongodb://%s:27017", db)

	// create a database connection
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURL))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// create a repository
	repository := repository.NewRepository(client.Database("users"))

	// create an http server
	server := http.NewServer(repository)

	// create a gin router
	router := gin.Default()
	{
		router.GET("/users/:email", server.GetUser)
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:email", server.UpdateUser)
		router.DELETE("/users/:email", server.DeleteUser)
	}

	// start the router
	router.Run(":9080")
}
