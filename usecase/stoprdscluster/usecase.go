package stoprdscluster

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/nekochans/aws-cost-saving-lambda/infrastructure"
)

type TargetRdsClusterList struct {
	TargetClusterList []string `json:"targetClusterList"`
}

type UseCase struct {
	RdsClient infrastructure.RdsClient
}

func (
	u *UseCase,
) StopRdsCluster(
	ctx context.Context,
	targetRdsClusterList TargetRdsClusterList,
) error {
	for _, v := range targetRdsClusterList.TargetClusterList {
		input := &rds.StopDBClusterInput{DBClusterIdentifier: aws.String(v)}

		_, err := u.RdsClient.StopDBCluster(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}
