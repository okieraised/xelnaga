package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

const (
	MetadataCollection = "metadata"
)

type MongoWrapperClient struct {
	mongoClient *mongo.Client
	logger      *zap.Logger
}

type DicomMetadata struct {
	ID                string    `bson:"_id"`
	StudyInstanceUID  string    `bson:"study_instance_uid"`
	SeriesInstanceUID string    `bson:"series_instance_uid"`
	SOPInstanceUID    string    `bson:"sop_instance_uid"`
	CreatedAt         time.Time `bson:"created_at"`
	//Metadata          go2com.MappedTag `bson:"metadata"`
}

// NewMongoWrapperClient returns the pointer to a MongoWrapperClient
func NewMongoWrapperClient(logger *zap.Logger, opts ...*options.ClientOptions) (*MongoWrapperClient, error) {
	client, err := mongo.Connect(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoWrapperClient{mongoClient: client}, nil
}

func (c *MongoWrapperClient) InsertMetadata() {

}

func (c *MongoWrapperClient) DeleteMetadata() {

}

func (c *MongoWrapperClient) DeleteAllMetadata() {

}
