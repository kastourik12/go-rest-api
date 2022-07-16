package artist

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Service struct {
	ArtistCollection *mongo.Collection
	Context          context.Context
}

func NewArtistService(artistCollection *mongo.Collection, ctx context.Context) Service {
	return Service{
		ArtistCollection: artistCollection,
		Context:          ctx,
	}
}
func (s *Service) CreateArtist(artist *DTO) error {
	artist.CreateAt = time.Now()
	artist.UpdateAt = time.Now()
	_, err := s.ArtistCollection.InsertOne(s.Context, artist)
	return err
}

func (s *Service) GetArtist(id string) (*Artist, error) {
	var artist Artist
	objID, _ := primitive.ObjectIDFromHex(id)
	err := s.ArtistCollection.FindOne(s.Context, bson.M{"_id": objID}).Decode(&artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (s *Service) UpdateArtist(id string, artist *Artist) error {
	artist.UpdateAt = time.Now()
	_, err := s.ArtistCollection.UpdateOne(s.Context, bson.M{"_id": id}, bson.M{"$set": artist})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteArtist(id string) error {
	filter := bson.D{{"_id", id}}
	_, err := s.ArtistCollection.DeleteOne(s.Context, filter)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetArtists() ([]Artist, error) {
	var artists []Artist
	cursor, err := s.ArtistCollection.Find(s.Context, bson.D{{}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.Context, &artists)
	cursor.Close(s.Context)
	if err != nil {
		return nil, err
	}
	return artists, nil
}
