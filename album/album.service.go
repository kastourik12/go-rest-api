package album

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AlbumService struct {
	AlbumCollection  *mongo.Collection
	Context          context.Context
	ArtistCollection *mongo.Collection
}

func NewAlbumService(albumCollection *mongo.Collection, ctx context.Context) AlbumService {
	return AlbumService{
		AlbumCollection: albumCollection,
		Context:         ctx,
	}
}

func (s *AlbumService) CreateAlbum(album *AlbumDTO) error {
	album.CreateAt = time.Now()
	album.UpdateAt = time.Now()
	_, err := s.AlbumCollection.InsertOne(s.Context, album)
	return err
}
func (s *AlbumService) GetAlbum(id string) (*Album, error) {
	var album Album
	objID, _ := primitive.ObjectIDFromHex(id)
	err := s.AlbumCollection.FindOne(s.Context, bson.M{"_id": objID}).Decode(&album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}
func (s *AlbumService) UpdateAlbum(id string, album *Album) error {
	album.UpdateAt = time.Now()
	_, err := s.AlbumCollection.UpdateOne(s.Context, bson.M{"_id": id}, bson.M{"$set": album})
	if err != nil {
		return err
	}
	return nil
}
func (s *AlbumService) DeleteAlbum(id string) error {
	filter := bson.D{{"_id", id}}
	_, err := s.AlbumCollection.DeleteOne(s.Context, filter)
	if err != nil {
		return err
	}
	return nil
}
func (s *AlbumService) GetAlbums() ([]Album, error) {
	var albums []Album
	cursor, err := s.AlbumCollection.Find(s.Context, bson.D{{}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.Context, &albums)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *AlbumService) GetAlbumsByArtist(id string) ([]Album, error) {
	var albums []Album
	objID, _ := primitive.ObjectIDFromHex(id)
	cursor, err := s.AlbumCollection.Find(s.Context, bson.M{"artist": objID})
	if err != nil {
		return nil, err
	}
	for cursor.Next(s.Context) {
		var album Album
		err := cursor.Decode(&album)
		if err != nil {
			return nil, err
		}

	}
	return albums, err
}
