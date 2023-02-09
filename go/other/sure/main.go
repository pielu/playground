package sure

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	var bucketName string
	var numberRecent int

	flag.StringVar(&bucketName, "b", "", "Specify S3 bucket name.")
	flag.IntVar(&numberRecent, "n", 0, "Specify number of most recent folders to spare.")

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

	s3rm := FindX(numberRecent, s3o)
	if err := s3svc.DeleteObjects(context.TODO(), bucketName, s3rm); err != nil {
		log.Fatal(err)
	}
}
