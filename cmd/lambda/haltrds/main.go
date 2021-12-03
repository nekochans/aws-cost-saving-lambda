package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type TargetRdsClusterList struct {
	TargetClusterList []string `json:"targetClusterList"`
}

func HandleRequest(ctx context.Context, request events.CloudWatchEvent) error {
	log.Println("🐱")
	log.Println(os.Getenv("TARGET_RDS_CLUSTER_LIST"))
	log.Println("🐱")

	var targetRdsClusterList TargetRdsClusterList
	if err := json.Unmarshal([]byte(os.Getenv("TARGET_RDS_CLUSTER_LIST")), &targetRdsClusterList); err != nil {
	}

	log.Println("🐹")
	log.Println(targetRdsClusterList)
	log.Println("🐹")

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
