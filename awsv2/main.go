package awsv2

import (
	myconfig "aws_test_iam_role/config"

	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func NewConfig(ctx context.Context, cfg *myconfig.AWSConfig) (aws.Config, error) {
	awsCfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithDefaultRegion(cfg.Region),
	)
	if err != nil {
		return aws.Config{}, err
	}

	if cfg.AccessKeyID != "" && cfg.SecretAccessKey != "" {
		awsCfg.Credentials = aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
			return aws.Credentials{AccessKeyID: cfg.AccessKeyID, SecretAccessKey: cfg.SecretAccessKey}, nil
		})
	}
	if cfg.SdkEndpoint != "" && cfg.SdkEndpoint != "default" {
		awsCfg.BaseEndpoint = aws.String(cfg.SdkEndpoint)
	}

	return awsCfg, nil
}

func QueryS3v2(awscfg *myconfig.AWSConfig) {
	fmt.Println("========= START QueryS3withSetup V2 ===============")
	cfg, err := NewConfig(context.Background(), awscfg)
	if err != nil {
		log.Fatalf("QueryS3withSetupV2 Unable to create session, %v", err)
	}
	svc := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if awscfg.SdkEndpoint != "" && awscfg.SdkEndpoint != "default" {
			o.UsePathStyle = true
		}
	})

	bucket := awscfg.S3BucketTest
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}
	//fmt.Printf("%v %v", input, svc)

	resp, err := svc.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Fatalf("unable to list items in bucket %q, %v", bucket, err)
	}

	fmt.Println("Objects in bucket:")
	for _, item := range resp.Contents {
		fmt.Printf("Name: %s, Last modified: %s, Size: %d\n", *item.Key, item.LastModified, item.Size)
	}

	fmt.Println("========= END QueryS3withSetupV2 ===============")
}
