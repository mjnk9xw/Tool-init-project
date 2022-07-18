package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ConnectionURI string       `json:"connection_uri" mapstructure:"connection_uri"`
	Database      string       `json:"database" mapstructure:"database"`
	Timeout       int          `json:"timeout" mapstructure:"timeout"`
	Options       MongoOptions `json:"options" mapstructure:"options"`
}

type MongoOptions struct {
	ReplicaSet string `json:"replica_set" mapstructure:"replica_set"`
	SSL        bool   `json:"ssl" mapstructure:"ssl"`
}

// DB object
type DB struct {
	Client   *mongo.Client
	Context  context.Context
	Database *mongo.Database
}

// Init mongodb client
func Init(config Config) (*DB, error) {
	// error ? mongodb://test:test@mongo:27017/livestream
	// sliceStr := strings.Split(config.ConnectionURI, "/")
	// config.Database = sliceStr[len(sliceStr)-1]
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Set client options
	clientOptions := options.Client().ApplyURI(config.ConnectionURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Connect to database
	database := client.Database(config.Database)

	return &DB{
		Database: database,
		Client:   client,
		Context:  context.Background(),
	}, nil
}

// Disconnect - disconnect to database
func (db *DB) Disconnect() error {
	return db.Client.Disconnect(db.Context)
}

// EnsureIndex will ensure the index model provided is on the given collection.
func (db *DB) EnsureIndex(ctx context.Context, c *mongo.Collection, keys []string, opts *options.IndexOptions) error {
	ks := bson.M{}
	for _, item := range keys {
		ks[item] = 1
	}
	index := mongo.IndexModel{Keys: ks, Options: opts}
	if _, err := c.Indexes().CreateOne(ctx, index); err != nil {
		return err
	}
	return nil
}
