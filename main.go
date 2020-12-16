package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Env lookup
func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value

	}
	return fallback
}

//FlagsAll for parsing flags
func FlagsAll(port string) string {
	portNo := port
	var portNoP = flag.String("port", portNo, "HTTP port. HTTPPORT environment variable can also be used.")
	flag.Parse()
	PortHTTP := *portNoP
	return PortHTTP
}

// Lists the items in the specified S3 Bucket
func s3List(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("name")

	bucket := value

	if len(value) <= 0 {

		fmt.Fprintf(w, "Add parameter")
	} else {

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		// Get the list of items
		resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
		if err != nil {
			exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
			fmt.Fprintf(w, "Bad server. Sorry we cannot list this bucket\n")
		}

		for _, item := range resp.Contents {
			fmt.Fprintf(w, "Name:         "+*item.Key+"\n")
			fmt.Fprintf(w, "Storage class:"+*item.StorageClass+"\n")
		}

		fmt.Fprintf(w, "\nFound "+strconv.Itoa(len(resp.Contents))+" items in bucket \n")

	}

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

func main() {
	var port = "8080"
	var portNo = getEnv("HTTPPORT", port)
	portHTTP := FlagsAll(portNo)
	fmt.Println("HTTP PORT:", portHTTP)
	http.HandleFunc("/list", s3List)
	log.Fatal(http.ListenAndServe(":"+portHTTP, nil))
}
