package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.CloudWatchEvent) error {
	log.Println("🐱")
	log.Println(os.Getenv("TARGET_CLUSTER_LIST"))
	log.Println("🐱")

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
