package uploader

import (
	"github.com/minio/minio-go"
	"io"
	"net/url"
	"os"
	"time"
)

// Uploader struct
type Uploader struct {
	*minio.Client
	*Config
}

// Config struct
type Config struct {
	BucketName      string
	AccessKey       string
	AccessKeySecret string
}

// InitUploader init uploader struct
func InitUploader() *Uploader {
	// Use a secure connection.
	ssl := true
	config := fetchConfig()

	// Initialize minio client object.
	minioClient, err := minio.New("s3.amazonaws.com", config.AccessKey, config.AccessKeySecret, ssl)
	if err != nil {
		panic("failed to initialize uploader")
	}
	return &Uploader{minioClient, config}
}

// UploadImage upload image file
func (u *Uploader) UploadImage(file io.Reader, path string) (int64, error) {
	n, err := u.PutObject(u.Config.BucketName, path, file, "application/octet-stream")
	return n, err
}

// GetImageURL GetImageURL
func (u *Uploader) GetImageURL(filepath string) (*url.URL, error) {
	url, err := u.PresignedGetObject(u.Config.BucketName, filepath, time.Second*60*60, make(url.Values))
	if err != nil {
		return nil, err
	}
	return url, nil
}

// DeleteImage delete image from bucket
func (u *Uploader) DeleteImage(filepath string) error {
	return nil
}

func fetchConfig() *Config {
	return &Config{
		BucketName:      os.Getenv("BUCKET_NAME"),
		AccessKey:       os.Getenv("AWS_ACCESS_KEY"),
		AccessKeySecret: os.Getenv("AWS_ACCESS_KEY_SECRET"),
	}
}
