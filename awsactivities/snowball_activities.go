package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/snowball/snowballiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type SnowballActivities struct {
	client snowballiface.SnowballAPI
}

func NewSnowballActivities(session *session.Session, config ...*aws.Config) *SnowballActivities {
	client := snowball.New(session, config...)
	return &SnowballActivities{client: client}
}

func (a *SnowballActivities) CancelCluster(ctx context.Context, input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error) {
	return a.client.CancelClusterWithContext(ctx, input)
}

func (a *SnowballActivities) CancelJob(ctx context.Context, input *snowball.CancelJobInput) (*snowball.CancelJobOutput, error) {
	return a.client.CancelJobWithContext(ctx, input)
}

func (a *SnowballActivities) CreateAddress(ctx context.Context, input *snowball.CreateAddressInput) (*snowball.CreateAddressOutput, error) {
	return a.client.CreateAddressWithContext(ctx, input)
}

func (a *SnowballActivities) CreateCluster(ctx context.Context, input *snowball.CreateClusterInput) (*snowball.CreateClusterOutput, error) {
	return a.client.CreateClusterWithContext(ctx, input)
}

func (a *SnowballActivities) CreateJob(ctx context.Context, input *snowball.CreateJobInput) (*snowball.CreateJobOutput, error) {
	return a.client.CreateJobWithContext(ctx, input)
}

func (a *SnowballActivities) DescribeAddress(ctx context.Context, input *snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error) {
	return a.client.DescribeAddressWithContext(ctx, input)
}

func (a *SnowballActivities) DescribeAddresses(ctx context.Context, input *snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error) {
	return a.client.DescribeAddressesWithContext(ctx, input)
}

func (a *SnowballActivities) DescribeCluster(ctx context.Context, input *snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error) {
	return a.client.DescribeClusterWithContext(ctx, input)
}

func (a *SnowballActivities) DescribeJob(ctx context.Context, input *snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error) {
	return a.client.DescribeJobWithContext(ctx, input)
}

func (a *SnowballActivities) GetJobManifest(ctx context.Context, input *snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error) {
	return a.client.GetJobManifestWithContext(ctx, input)
}

func (a *SnowballActivities) GetJobUnlockCode(ctx context.Context, input *snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error) {
	return a.client.GetJobUnlockCodeWithContext(ctx, input)
}

func (a *SnowballActivities) GetSnowballUsage(ctx context.Context, input *snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error) {
	return a.client.GetSnowballUsageWithContext(ctx, input)
}

func (a *SnowballActivities) GetSoftwareUpdates(ctx context.Context, input *snowball.GetSoftwareUpdatesInput) (*snowball.GetSoftwareUpdatesOutput, error) {
	return a.client.GetSoftwareUpdatesWithContext(ctx, input)
}

func (a *SnowballActivities) ListClusterJobs(ctx context.Context, input *snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error) {
	return a.client.ListClusterJobsWithContext(ctx, input)
}

func (a *SnowballActivities) ListClusters(ctx context.Context, input *snowball.ListClustersInput) (*snowball.ListClustersOutput, error) {
	return a.client.ListClustersWithContext(ctx, input)
}

func (a *SnowballActivities) ListCompatibleImages(ctx context.Context, input *snowball.ListCompatibleImagesInput) (*snowball.ListCompatibleImagesOutput, error) {
	return a.client.ListCompatibleImagesWithContext(ctx, input)
}

func (a *SnowballActivities) ListJobs(ctx context.Context, input *snowball.ListJobsInput) (*snowball.ListJobsOutput, error) {
	return a.client.ListJobsWithContext(ctx, input)
}

func (a *SnowballActivities) UpdateCluster(ctx context.Context, input *snowball.UpdateClusterInput) (*snowball.UpdateClusterOutput, error) {
	return a.client.UpdateClusterWithContext(ctx, input)
}

func (a *SnowballActivities) UpdateJob(ctx context.Context, input *snowball.UpdateJobInput) (*snowball.UpdateJobOutput, error) {
	return a.client.UpdateJobWithContext(ctx, input)
}
