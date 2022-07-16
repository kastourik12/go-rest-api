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
	"kastouri/web-service-gin/song"
	"log"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client
	err         error

	albumCollection *mongo.Collection
	albumService    album.Service
	albumController album.Controller

	artistCollection *mongo.Collection
	artistService    artist.Service
	artistController artist.Controller

	songCollection *mongo.Collection
	songService    song.Service
	songController song.Controller
)

func init() {
	ctx = context.TODO()

	mongoConnection := options.Client().ApplyURI("mongodb+srv://root:root@cluster0.cd2g5kn.mongodb.net/?retryWrites=true&w=majority")
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
	artistService = artist.NewArtistService(artistCollection, ctx)
	artistController = artist.NewArtistController(artistService)

	albumCollection = mongoclient.Database("test-go").Collection("albums")
	albumService = album.NewAlbumService(albumCollection, ctx)
	albumController = album.NewAlbumController(albumService)

	songCollection = mongoclient.Database("test-go").Collection("songs")
	songService = song.NewSongService(songCollection, ctx)
	songController = song.NewSongController(songService)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basePath := server.Group("/v1")
	artistController.RegisterRoutes(basePath)
	albumController.RegisterRoutes(basePath)
	songController.RegisterRoutes(basePath)
	log.Fatal(server.Run(":9090"))
}
