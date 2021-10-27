package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	config *Config
	db     *mongo.Client
	ctx    context.Context
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	clientOptions := options.Client().ApplyURI(s.config.DatabaseURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	s.db = client
	s.ctx = ctx
	return nil
}

func (s *Store) Close() {
	s.db.Disconnect(s.ctx)
}

func (s *Store) InsertOne() error {
	_, err := s.db.Database("csv-restaurant").Collection("test-go").InsertOne(s.ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})
	if err != nil {
		return err
	}
	fmt.Println("InsertOne")
	return nil
}
