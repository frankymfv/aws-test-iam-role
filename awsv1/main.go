package awsv1

import (
	"aws_test_iam_role/config"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"time"
)

func NewConfigAwsV1(cfg *config.AWSConfig) *aws.Config {
	// nolint: wrapcheck
	awsCfg := aws.NewConfig()
	awsCfg.Region = aws.String(cfg.Region)
	if cfg.AccessKeyID != "" && cfg.SecretAccessKey != "" {
		awsCfg.Credentials = credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, "")
	}
	if cfg.SdkEndpoint != "" && cfg.SdkEndpoint != "default" {
		awsCfg.Endpoint = aws.String(cfg.SdkEndpoint)
	}

	return awsCfg
}

func InitS3Config(cfg *config.AWSConfig) *session.Session {
	awsConf := NewConfigAwsV1(cfg)
	awsConf.S3ForcePathStyle = aws.Bool(true)
	awsConf.WithHTTPClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	})
	sess := session.Must(session.NewSession(awsConf))
	return sess
}

func QueryS3withInitS3Config(awscfg *config.AWSConfig) {
	fmt.Println("========= START InitS3Config ===============")
	sess := InitS3Config(awscfg)
	svc := s3.New(sess)

	// List the objects in the specified bucket
	bucket := awscfg.S3BucketTest
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		log.Fatalf("Unable to list items in bucket %q, %v", bucket, err)
	}

	fmt.Println("func Objects in bucket:")
	for _, item := range resp.Contents {
		fmt.Printf("Name: %s, Last modified: %s, Size: %d\n", *item.Key, item.LastModified, *item.Size)
	}
	fmt.Println("========= END InitS3Config ===============")
}

func QueryS3withSetup(awscfg *config.AWSConfig) {
	fmt.Println("========= START QueryS3withSetup ===============")
	svc := Setup(awscfg)

	// List the objects in the specified bucket
	bucket := awscfg.S3BucketTest
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		log.Fatalf("Unable to list items in bucket %q, %v", bucket, err)
	}

	fmt.Println("func Objects in bucket:")
	for _, item := range resp.Contents {
		fmt.Printf("Name: %s, Last modified: %s, Size: %d\n", *item.Key, item.LastModified, *item.Size)
	}
	fmt.Println("========= END QueryS3withSetup ===============")
}

func Setup(awsCfg *config.AWSConfig) *s3.S3 {
	config := NewConfigAwsV1(awsCfg)
	config.S3ForcePathStyle = aws.Bool(true)
	config.WithHTTPClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
		Timeout: 5 * 60 * time.Second,
	})

	sess := session.Must(session.NewSession(config))
	svc := s3.New(sess)
	return svc
}

func InitSqs(awsCfg *config.AWSConfig) *sqs.SQS {
	// init sqs client
	cfg := NewConfigAwsV1(awsCfg)
	cfg.S3ForcePathStyle = aws.Bool(true)

	sess := session.Must(session.NewSession(cfg))
	s := sqs.New(sess)
	return s
}

func SendMessageToQueue(awsCfg *config.AWSConfig, messageBody string) {
	fmt.Println("========= START SendMessageToQueue ===============")
	svc := InitSqs(awsCfg)
	queueUrl := awsCfg.SqsQueueUrl
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(queueUrl),
	}

	result, err := svc.SendMessage(input)
	if err != nil {
		log.Fatalf("Failed to send message to queue %q, %v", queueUrl, err)
	}

	log.Printf("Message sent to queue %q, message ID: %s", queueUrl, *result.MessageId)
	fmt.Println("========= END SendMessageToQueue ===============")

}
