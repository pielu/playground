package main

import (
	"flag"
	"log"
	"main/sure"
)

func main() {
	var bucketName string
	var numberRecent int

	flag.StringVar(&bucketName, "bucketName", "", "Specify S3 bucket name.")
	flag.IntVar(&numberRecent, "numberRecent", 0, "Specify number of most recent folders to spare.")
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" || f.Value.String() == "0" {
			log.Fatal(f.Name, " is not set!")
		}
	})

	sure.Challenge(bucketName, numberRecent)
}
