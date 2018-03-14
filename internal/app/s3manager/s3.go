package s3manager

import (
	"io"

	minio "github.com/minio/minio-go"
)

// S3 is a client to interact with S3 storage.
type S3 interface {
	CopyObject(minio.DestinationInfo, minio.SourceInfo) error
	GetObject(string, string, minio.GetObjectOptions) (*minio.Object, error)
	ListBuckets() ([]minio.BucketInfo, error)
	ListObjectsV2(string, string, bool, <-chan struct{}) <-chan minio.ObjectInfo
	MakeBucket(string, string) error
	PutObject(string, string, io.Reader, int64, minio.PutObjectOptions) (int64, error)
	RemoveBucket(string) error
	RemoveObject(string, string) error
}