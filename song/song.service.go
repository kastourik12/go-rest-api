package song

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type SongService struct {
	SongCollection   *mongo.Collection
	ArtistCollection *mongo.Collection
	Context          context.Context
}

func NewSongService(songCollection *mongo.Collection, ctx context.Context) SongService {
	return SongService{
		SongCollection: songCollection,
		Context:        ctx,
	}
}
func (s *SongService) CreateSong(song *SongDTO) error {
	song.CreateAt = time.Now()
	song.UpdateAt = time.Now()
	_, err := s.SongCollection.InsertOne(s.Context, song)
	return err
}
func (s *SongService) GetSong(id string) (*Song, error) {
	var song Song
	objID, _ := primitive.ObjectIDFromHex(id)
	err := s.SongCollection.FindOne(s.Context, bson.M{"_id": objID}).Decode(&song)
	if err != nil {
		return nil, err
	}
	return &song, nil
}
func (s *SongService) UpdateSong(id string, song *Song) error {
	song.UpdateAt = time.Now()
	_, err := s.SongCollection.UpdateOne(s.Context, bson.M{"_id": id}, bson.M{"$set": song})
	if err != nil {
		return err
	}
	return nil
}
func (s *SongService) DeleteSong(id string) error {
	filter := bson.D{{"_id", id}}
	_, err := s.SongCollection.DeleteOne(s.Context, filter)
	if err != nil {
		return err
	}
	return nil
}
func (s *SongService) GetSongs() ([]Song, error) {
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
func (s *SongService) GetSongsByArtist(artistId string) ([]Song, error) {
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

func (s *SongService) GetSongsByAlbum(id string) ([]Song, error) {
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
