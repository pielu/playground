package sure

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestSure(t *testing.T) {
	awsCfg := NewAwsCfg()
	s3BucketName := os.Getenv("S3_BUCKET")
	s3PopulatePath := os.Getenv("S3_POPULATE_PATH")

	s3client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		// Seems like we need to allow non-default path-style to s3 service
		o.UsePathStyle = true
	})
	s3svc := NewS3(s3client)
	if err := s3svc.Create(context.TODO(), s3BucketName); err != nil {
		t.Error(err)
	}

	s3bn, _ := s3svc.List(context.TODO())
	if *s3bn.Buckets[0].Name != s3BucketName {
		t.Errorf("bucket name mismatch! have: %s want: %s", *s3bn.Buckets[0].Name, s3BucketName)
	}

	if _, err := s3Populate(s3svc.uploader, s3PopulatePath, s3BucketName); err != nil {
		t.Error(err)
	}

}

// Shamelessly from https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/
type fileWalk chan string

func (f fileWalk) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		f <- path
	}
	return nil
}

func s3Populate(mu *manager.Uploader, s3pp string, bucketName string) (map[string]string, error) {
	walker := make(fileWalk)
	upload := make(map[string]string, 0)
	go func() error {
		// Gather the files to upload by walking the path recursively
		if err := filepath.Walk(s3pp, walker.Walk); err != nil {
			return err
		}
		close(walker)
		return nil
	}()

	// For each file found walking, upload it to Amazon S3
	for path := range walker {
		rel, err := filepath.Rel(s3pp, path)
		if err != nil {
			return nil, err
		}
		file, err := os.Open(path)
		if err != nil {
			return nil, err
			continue
		}
		defer file.Close()
		result, err := mu.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: &bucketName,
			Key:    aws.String(rel),
			Body:   file,
		})
		if err != nil {
			return nil, err
		}
		upload[path] = result.Location
	}

	return upload, nil
}
