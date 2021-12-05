package infrastructure

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type RdsClient interface {
	StopDBCluster(
		ctx context.Context,
		params *rds.StopDBClusterInput,
		optFns ...func(*rds.Options),
	) (*rds.StopDBClusterOutput, error)
}
