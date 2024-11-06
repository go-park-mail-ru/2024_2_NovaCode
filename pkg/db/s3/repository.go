package s3

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type S3Repo interface {
	Put(ctx context.Context, upload Upload) (*minio.UploadInfo, error)
	Get(ctx context.Context, bucket string, filename string) (*minio.Object, error)
	Remove(ctx context.Context, bucket string, filename string) error
}
