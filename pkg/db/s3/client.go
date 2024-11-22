package s3

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func New(cfg *config.MinioConfig) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.URL, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.User, cfg.Password, ""),
		Secure: cfg.SSLMode,
	})
	if err != nil {
		return nil, err
	}

	if err = ping(minioClient); err != nil {
		return nil, err
	}

	return minioClient, nil
}

func ping(minioClient *minio.Client) error {
	ctx := context.Background()
	_, err := minioClient.ListBuckets(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping minio server: %w", err)
	}

	return nil
}
