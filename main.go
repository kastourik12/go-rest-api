package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"kastouri/web-service-gin/controllers"
	"kastouri/web-service-gin/services"
	"log"
)

var (
	server           *gin.Engine
	artistService    services.ArtistService
	artistController controllers.ArtistController
	ctx              context.Context
	artistCollection *mongo.Collection
	mongoclient      *mongo.Client
	err              error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	artistCollection = mongoclient.Database("userdb").Collection("users")
	artistService = services.NewArtistService(artistCollection, ctx)
	artistController = controllers.New(artistService)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basePath := server.Group("/v1")
	artistController.RegisterRoutes(basePath)

	log.Fatal(server.Run(":9090"))
}
