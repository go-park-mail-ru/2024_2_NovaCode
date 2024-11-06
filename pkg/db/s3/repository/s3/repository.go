package s3

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/s3"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type S3Repo struct {
	client *minio.Client
	logger logger.Logger
}

func NewS3Repository(client *minio.Client, logger logger.Logger) *S3Repo {
	return &S3Repo{client, logger}
}

func (repo *S3Repo) Put(ctx context.Context, upload s3.Upload) (*minio.UploadInfo, error) {
	options := minio.PutObjectOptions{
		ContentType:  upload.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	uploadInfo, err := repo.client.PutObject(ctx, upload.Bucket, repo.generateFilename(upload.Filename), upload.File, upload.Size, options)
	if err != nil {
		return nil, fmt.Errorf("failed to put object: %v", err)
	}

	return &uploadInfo, nil
}

func (repo *S3Repo) Get(ctx context.Context, bucket string, filename string) (*minio.Object, error) {
	object, err := repo.client.GetObject(ctx, bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %v", err)
	}

	return object, nil
}

func (repo *S3Repo) Remove(ctx context.Context, bucket string, filename string) error {
	if err := repo.client.RemoveObject(ctx, bucket, filename, minio.RemoveObjectOptions{}); err != nil {
		return fmt.Errorf("faild to remove object: %v", err)
	}

	return nil
}

func (repo *S3Repo) generateFilename(filename string) string {
	uid := uuid.New().String()
	return fmt.Sprintf("%s-%s", uid, filename)
}
