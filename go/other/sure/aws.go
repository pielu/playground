package sure

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketObject struct {
	Key        string
	Size       int64
	ModifiedAt time.Time
}

type BucketClient interface {
	Create(ctx context.Context, bucket string) error
	UploadObject(ctx context.Context, bucket, fileName string, body io.Reader) (string, error)
	DeleteObject(ctx context.Context, bucket, fileName string) error
	ListObjects(ctx context.Context, bucket string) ([]*BucketObject, error)
}

type S3 struct {
	client   *s3.Client
	uploader *manager.Uploader
}

func NewS3(s3client *s3.Client) S3 {
	return S3{
		client:   s3client,
		uploader: manager.NewUploader(s3client),
	}
}

func (s S3) Create(ctx context.Context, bucketName string) error {
	if _, err := s.client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}); err != nil {
		return fmt.Errorf("create: %w", err)
	}

	return nil
}

func (s S3) List(ctx context.Context) (*s3.ListBucketsOutput, error) {
	res, err := s.client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}

	return res, nil

}

func NewAwsCfg() aws.Config {
	awsRegion := os.Getenv("AWS_REGION")
	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	// bucketName := os.Getenv("S3_BUCKET")

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS config: %s", err)
	}

	return awsCfg
}
