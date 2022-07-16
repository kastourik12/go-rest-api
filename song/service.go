package song

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Service struct {
	SongCollection   *mongo.Collection
	ArtistCollection *mongo.Collection
	Context          context.Context
}

func NewSongService(songCollection *mongo.Collection, ctx context.Context) Service {
	return Service{
		SongCollection: songCollection,
		Context:        ctx,
	}
}
func (s *Service) CreateSong(song *SongDTO) error {

	createdSong, err := NewSong(*song)
	if err != nil {
		return err
	}
	_, err2 := s.SongCollection.InsertOne(s.Context, createdSong)
	return err2
}
func (s *Service) GetSong(id string) (*Song, error) {
	var song Song
	objID, _ := primitive.ObjectIDFromHex(id)
	err := s.SongCollection.FindOne(s.Context, bson.M{"_id": objID}).Decode(&song)
	if err != nil {
		return nil, err
	}
	return &song, nil
}
func (s *Service) UpdateSong(id string, song *Song) error {
	song.UpdateAt = time.Now()
	_, err := s.SongCollection.UpdateOne(s.Context, bson.M{"_id": id}, bson.M{"$set": song})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteSong(id string) error {
	filter := bson.D{{"_id", id}}
	_, err := s.SongCollection.DeleteOne(s.Context, filter)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetSongs() ([]Song, error) {
	var songs []Song
	cursor, err := s.SongCollection.Find(s.Context, bson.D{{}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.Context, &songs)
	cursor.Close(s.Context)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
func (s *Service) GetSongsByArtist(artistId string) ([]Song, error) {
	var songs []Song
	objID, _ := primitive.ObjectIDFromHex(artistId)
	cursor, err := s.SongCollection.Find(s.Context, bson.M{"artist": objID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.Context, &songs)
	cursor.Close(s.Context)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s *Service) GetSongsByAlbum(id string) ([]Song, error) {
	var songs []Song
	objID, _ := primitive.ObjectIDFromHex(id)
	cursor, err := s.SongCollection.Find(s.Context, bson.M{"album": objID})
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.Context, &songs)
	cursor.Close(s.Context)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
