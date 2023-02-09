package sure

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestSure(t *testing.T) {
	// Make sure it won't cache
	time.Now()

	awsCfg := NewCfg(os.Getenv("AWS_REGION"), os.Getenv("AWS_ENDPOINT"))
	s3BucketName := os.Getenv("S3_BUCKET")
	s3BucketPopulatePath := os.Getenv("S3_BUCKET_POPULATE_PATH")

	s3client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		// Seems like we need to allow non-default path-style to s3 service
		o.UsePathStyle = true
	})
	s3svc := NewS3(s3client)
	if err := s3svc.Create(context.TODO(), s3BucketName); err != nil {
		t.Error(err)
	}

	if s3bn, err := s3svc.Exists(context.TODO(), s3BucketName); s3bn == false || err != nil {
		t.Error(s3bn, err)
	}

	if _, err := s3BucketPopulate(s3svc.uploader, s3BucketPopulatePath, s3BucketName); err != nil {
		t.Error(err)
	}

	s3o, err := s3svc.ListObjects(context.TODO(), s3BucketName)
	if err != nil {
		t.Error(err)
	}

	Printer("ORIGINAL", s3o)

	// List to remove should not contain y, x and _z folders
	s3rm := FindX(3, s3o)
	s3safe := []string{"y", "x", "z"}

	for _, o := range s3rm {
		tl := strings.Split(o.Key, "/")[0]
		if Contains(s3safe, tl) {
			t.Error("Remove list contains wrong top level directories")
		}
	}

	if err := s3svc.DeleteObjects(context.TODO(), s3BucketName, s3rm); err != nil {
		t.Error(err)
	}

	s3or, err := s3svc.ListObjects(context.TODO(), s3BucketName)
	if err != nil {
		t.Error(err)
	}
	Printer("REMAINED", s3or)
}

// Walk the directory and upload dummy files for testing
// Shamelessly from https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/
type fileWalk chan string

func (f fileWalk) Walk(path string, info os.FileInfo, err error) error {
	// Since AWS does not allow changing timestamp slow it down for timestamp variety
	// https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingMetadata.html#object-metadata
	time.Sleep(5 + time.Second)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		f <- path
	}

	return nil
}

func s3BucketPopulate(mu *manager.Uploader, s3pp string, bucketName string) (map[string]string, error) {
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
