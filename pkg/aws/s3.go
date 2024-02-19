package aws

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type S3 interface {
	SetConfig() error
	Upload(key string, bucketName string, file io.ReadSeeker) (resp *manager.UploadOutput, err error)
	GetObject(ctx context.Context, key string, bucketName string) (resp *awss3.GetObjectOutput, err error)
}

type s3 struct {
	AwsS3Region  string
	AwsAccessKey string
	AwsSecretKey string
	client       *awss3.Client
}

func NewS3(region, accessKey, secretKey string) S3 {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := awss3.NewFromConfig(cfg)

	return &s3{
		AwsS3Region:  region,
		AwsAccessKey: accessKey,
		AwsSecretKey: secretKey,
		client:       client,
	}
}

func (s *s3) SetConfig() error {
	creds := credentials.NewStaticCredentialsProvider(s.AwsAccessKey, s.AwsSecretKey, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds),
		config.WithRegion(s.AwsS3Region),
	)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	s.client = awss3.NewFromConfig(cfg)
	return nil
}

func (s *s3) GetObject(ctx context.Context, key string, bucketName string) (resp *awss3.GetObjectOutput, err error) {
	downloader := manager.NewDownloader(s.client)
	return downloader.S3.GetObject(ctx, &awss3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

}

func (s *s3) Upload(key string, bucketName string, file io.ReadSeeker) (resp *manager.UploadOutput, err error) {
	uploader := manager.NewUploader(s.client)
	resp, err = uploader.Upload(context.TODO(), &awss3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	return
}
