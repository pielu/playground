package sure

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Challenge(bucketName string, numberRecent int) {
	awsCfg := NewCfg(os.Getenv("AWS_REGION"), os.Getenv("AWS_ENDPOINT"))

	s3client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		// Seems like we need to allow non-default path-style to s3 service
		o.UsePathStyle = true
	})
	s3svc := NewS3(s3client)

	if s3bn, err := s3svc.Exists(context.TODO(), bucketName); s3bn == false || err != nil {
		log.Fatal(err)
	}

	s3o, err := s3svc.ListObjects(context.TODO(), bucketName)
	if err != nil {
		log.Fatal(err)
	}
	Printer("ORIGINAL", s3o)

	s3rm := FindX(numberRecent, s3o)
	Printer("TO REMOVE", s3rm)

	if len(s3rm) <= numberRecent {
		log.Printf("Nothing to do! Asked to leave %d most recent folders in the bucket `%s`", numberRecent, bucketName)
		os.Exit(0)
	}

	if err := s3svc.DeleteObjects(context.TODO(), bucketName, s3rm); err != nil {
		log.Fatal(err)
	}

	s3or, err := s3svc.ListObjects(context.TODO(), bucketName)
	if err != nil {
		log.Fatal(err)
	}
	Printer("REMAINED", s3or)
}
