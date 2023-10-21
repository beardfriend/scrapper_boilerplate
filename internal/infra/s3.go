package infra

import (
	cfg "boilerplate/config"
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	client *awss3.Client
}

var onceS3 sync.Once
var s3 *S3

func S3Instance() *S3 {
	if s3 != nil {
		return s3
	}
	onceS3.Do(func() {
		awscfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		client := awss3.NewFromConfig(awscfg)
		s3.client = client

		creds := credentials.NewStaticCredentialsProvider(cfg.S3AccessKey, cfg.S3SecretKey, "")
		s3NewConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds),
			config.WithRegion(cfg.S3Region),
		)

		if err != nil {
			panic(err)
		}

		s3.client = awss3.NewFromConfig(s3NewConfig)
	})

	return s3
}
