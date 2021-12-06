package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/nekochans/aws-cost-saving-lambda/usecase/stoprdscluster"
)

var useCase stoprdscluster.UseCase

//nolint:gochecknoinits
func init() {
	region := os.Getenv("AWS_REGION")

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		// TODO We will be replacing this with a logging process of our own design.
		log.Fatalln(err)
	}

	rdsClient := rds.NewFromConfig(cfg)

	useCase = stoprdscluster.UseCase{RdsClient: rdsClient}
}

func HandleRequest(ctx context.Context, request events.CloudWatchEvent) error {
	var targetRdsClusterList stoprdscluster.TargetRdsClusterList
	if err := json.Unmarshal([]byte(os.Getenv("TARGET_RDS_CLUSTER_LIST")), &targetRdsClusterList); err != nil {
		// TODO We will be replacing this with a logging process of our own design.
		log.Fatalln(err)
	}

	if err := useCase.StopRdsCluster(ctx, targetRdsClusterList); err != nil {
		// TODO We will be replacing this with a logging process of our own design.
		log.Fatalln(err)
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
