package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// AWSConfig hold configuration of aws
type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	SdkEndpoint     string
	S3BucketTest    string
	SqsQueueUrl     string
}

func LoadCfg() *AWSConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get environment variables
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("SECRET_ACCESS_KEY")
	region := os.Getenv("REGION")
	sdkEndpoint := os.Getenv("SDK_ENDPOINT")
	s3BucketTest := os.Getenv("S3_BUCKET_TEST")
	sqsQueueURL := os.Getenv("SQS_QUEUE_URL")
	return &AWSConfig{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Region:          region,
		SdkEndpoint:     sdkEndpoint,
		S3BucketTest:    s3BucketTest,
		SqsQueueUrl:     sqsQueueURL,
	}
}
