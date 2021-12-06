# aws-cost-saving-lambda
It provides a function to save cost by periodically stopping unnecessary AWS resources.

## Getting Started

### Setting up AWS credentials

Use [Named profiles](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) for this.

### Setting environment variables

Using [direnv](https://github.com/direnv/direnv) is convenient.

```
export DEPLOY_STAGE=dev
export AWS_REGION=ap-northeast-1
export AWS_PROFILE=nekochans-dev
export SLS_DEBUG=*
export TARGET_RDS_CLUSTER_LIST='{"targetClusterList": ["dev-database-cluster"]}'
export STOP_RDS_CLUSTER_SCHEDULE_RATE='cron(0 3 * * ? *)'
export STOP_RDS_CLUSTER_IS_ENABLED=true
```
