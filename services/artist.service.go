package services

import (
	"context"
	_ "context"
	"errors"
	_ "errors"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"kastouri/web-service-gin/models"
	_ "kastouri/web-service-gin/models"
)

type ArtistService struct {
	ArtistCollection *mongo.Collection
	Context          context.Context
}

func NewArtistService(artistCollection *mongo.Collection, ctx context.Context) ArtistService {
	return ArtistService{
		ArtistCollection: artistCollection,
		Context:          ctx,
	}
}
func (s *ArtistService) CreateArtist(artist *models.Artist) error {
	_, err := s.ArtistCollection.InsertOne(s.Context, artist)
	if err != nil {
		return err
	}
	return nil
}

func (s *ArtistService) GetArtist(id string) (*models.Artist, error) {
	var artist models.Artist
	err := s.ArtistCollection.FindOne(s.Context, bson.M{"_id": id}).Decode(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (s *ArtistService) UpdateArtist(id string, artist *models.Artist) error {
	_, err := s.ArtistCollection.UpdateOne(s.Context, bson.M{"_id": id}, bson.M{"$set": artist})
	if err != nil {
		return err
	}
	return nil
}

func (s *ArtistService) DeleteArtist(id string) error {
	_, err := s.ArtistCollection.DeleteOne(s.Context, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
func (s *ArtistService) GetArtists() ([]models.Artist, error) {
	var artists []models.Artist
	cursor, err := s.ArtistCollection.Find(s.Context, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(s.Context) {
		var artist models.Artist
		err := cursor.Decode(&artist)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	cursor.Close(s.Context)
	if len(artists) == 0 {
		return nil, errors.New("no artists found")
	}

	return artists, nil
}
