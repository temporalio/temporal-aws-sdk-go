package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/snowball/snowballiface"
)

type SnowballActivities struct {
	client snowballiface.SnowballAPI
}

func NewSnowballActivities(client snowballiface.SnowballAPI) *SnowballActivities {
    return &SnowballActivities{client: client}
}


func (a *SnowballActivities) CancelCluster(input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error) {
    return a.client.CancelCluster(input)
}



func (a *SnowballActivities) CancelJob(input *snowball.CancelJobInput) (*snowball.CancelJobOutput, error) {
    return a.client.CancelJob(input)
}



func (a *SnowballActivities) CreateAddress(input *snowball.CreateAddressInput) (*snowball.CreateAddressOutput, error) {
    return a.client.CreateAddress(input)
}



func (a *SnowballActivities) CreateCluster(input *snowball.CreateClusterInput) (*snowball.CreateClusterOutput, error) {
    return a.client.CreateCluster(input)
}



func (a *SnowballActivities) CreateJob(input *snowball.CreateJobInput) (*snowball.CreateJobOutput, error) {
    return a.client.CreateJob(input)
}



func (a *SnowballActivities) DescribeAddress(input *snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error) {
    return a.client.DescribeAddress(input)
}



func (a *SnowballActivities) DescribeAddresses(input *snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error) {
    return a.client.DescribeAddresses(input)
}



func (a *SnowballActivities) DescribeCluster(input *snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error) {
    return a.client.DescribeCluster(input)
}



func (a *SnowballActivities) DescribeJob(input *snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error) {
    return a.client.DescribeJob(input)
}



func (a *SnowballActivities) GetJobManifest(input *snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error) {
    return a.client.GetJobManifest(input)
}



func (a *SnowballActivities) GetJobUnlockCode(input *snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error) {
    return a.client.GetJobUnlockCode(input)
}



func (a *SnowballActivities) GetSnowballUsage(input *snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error) {
    return a.client.GetSnowballUsage(input)
}



func (a *SnowballActivities) GetSoftwareUpdates(input *snowball.GetSoftwareUpdatesInput) (*snowball.GetSoftwareUpdatesOutput, error) {
    return a.client.GetSoftwareUpdates(input)
}



func (a *SnowballActivities) ListClusterJobs(input *snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error) {
    return a.client.ListClusterJobs(input)
}



func (a *SnowballActivities) ListClusters(input *snowball.ListClustersInput) (*snowball.ListClustersOutput, error) {
    return a.client.ListClusters(input)
}



func (a *SnowballActivities) ListCompatibleImages(input *snowball.ListCompatibleImagesInput) (*snowball.ListCompatibleImagesOutput, error) {
    return a.client.ListCompatibleImages(input)
}



func (a *SnowballActivities) ListJobs(input *snowball.ListJobsInput) (*snowball.ListJobsOutput, error) {
    return a.client.ListJobs(input)
}



func (a *SnowballActivities) UpdateCluster(input *snowball.UpdateClusterInput) (*snowball.UpdateClusterOutput, error) {
    return a.client.UpdateCluster(input)
}



func (a *SnowballActivities) UpdateJob(input *snowball.UpdateJobInput) (*snowball.UpdateJobOutput, error) {
    return a.client.UpdateJob(input)
}

