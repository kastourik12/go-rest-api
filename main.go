package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"kastouri/web-service-gin/album"
	"kastouri/web-service-gin/artist"
	"log"
)

var (
	server           *gin.Engine
	albumCollection  *mongo.Collection
	albumService     album.AlbumService
	albumController  album.AlbumController
	artistCollection *mongo.Collection
	artistService    artist.ArtistService
	artistController artist.ArtistController
	ctx              context.Context

	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoConnection := options.Client().ApplyURI("mongodb+srv://test:test@cluster0.bki6ob7.mongodb.net/?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	artistCollection = mongoclient.Database("test-go").Collection("artists")
	albumCollection = mongoclient.Database("test-go").Collection("albums")
	albumService = album.NewAlbumService(albumCollection, ctx)
	albumController = album.NewAlbumController(albumService)
	artistService = artist.NewArtistService(artistCollection, ctx)
	artistController = artist.NewArtistController(artistService)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basePath := server.Group("/v1")
	artistController.RegisterRoutes(basePath)
	albumController.RegisterRoutes(basePath)

	log.Fatal(server.Run(":9090"))
}
