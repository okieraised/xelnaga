package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"io"
)

type S3StorageCredential struct {
	Endpoint    string
	AccessKeyID string
	SecretKeyID string
}

type s3ClientWrapper struct {
	client *minio.Client
	logger *zap.Logger
}

func NewS3Storage(endpoint string, opts *minio.Options, logger *zap.Logger) (*s3ClientWrapper, error) {
	// Initialize minio client object.
	client, err := minio.New(endpoint, opts)
	if err != nil {
		return nil, err
	}
	return &s3ClientWrapper{
		client: client,
		logger: logger,
	}, nil
}

func (c *s3ClientWrapper) CreateStorage(bucketName string, opts minio.MakeBucketOptions) error {
	err := c.client.MakeBucket(context.Background(), bucketName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (c *s3ClientWrapper) ListBuckets() ([]minio.BucketInfo, error) {
	bucketInfo, err := c.client.ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}
	return bucketInfo, nil
}

func (c *s3ClientWrapper) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error) {
	info, err := c.client.PutObject(context.Background(), bucketName, objectName, reader, objectSize, opts)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (c *s3ClientWrapper) GetObject(bucketName string, objectName string, opts minio.GetObjectOptions) (*minio.Object, error) {
	obj, err := c.client.GetObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
