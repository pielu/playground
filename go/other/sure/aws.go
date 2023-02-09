package sure

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Define basics; for a more robust abstraction probably should use interface struct
// for the s3 bucket client fitting s3 receiver functions

type BucketObject struct {
	Key        string
	ModifiedAt time.Time
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

func (s S3) Exists(ctx context.Context, bucketName string) (bool, error) {
	_, err := s.client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	exists := true
	if err != nil {
		return false, fmt.Errorf("exists not: %w", err)
	}
	return exists, nil
}

func (s S3) List(ctx context.Context) (*s3.ListBucketsOutput, error) {
	res, err := s.client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}

	return res, nil

}

func (s S3) ListObjects(ctx context.Context, bucketName string) ([]*BucketObject, error) {
	// Result is sorted by key name, not date
	res, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}

	objects := make([]*BucketObject, len(res.Contents))

	for i, object := range res.Contents {
		objects[i] = &BucketObject{
			Key:        *object.Key,
			ModifiedAt: *object.LastModified,
		}
	}

	// Sort by date and make it desc, so latest (recent) is at the beginning
	sort.SliceStable(objects, func(i, j int) bool {
		return objects[j].ModifiedAt.Before(objects[i].ModifiedAt)
	})

	return objects, nil
}

func (s S3) DeleteObjects(ctx context.Context, bucketName string, objects []*BucketObject) error {
	var objectIds []types.ObjectIdentifier

	for _, o := range objects {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(o.Key)})
	}
	_, err := s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return err
}

func NewCfg(region string, endpoint string) aws.Config {
	customResolver := aws.EndpointResolverFunc(func(s, r string) (aws.Endpoint, error) {
		if endpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           endpoint,
				SigningRegion: region,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		fmt.Errorf("aws config: %w", err)
	}

	return awsCfg
}
