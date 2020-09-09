
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2/cloudhsmv2iface"
)

type CloudHSMV2Activities struct {
	client cloudhsmv2iface.CloudHSMV2API
}

func NewCloudHSMV2Activities(session *session.Session, config... *aws.Config) *CloudHSMV2Activities {
    client := cloudhsmv2.New(session, config...)
    return &CloudHSMV2Activities{client: client}
}

func (a *CloudHSMV2Activities) CopyBackupToRegion(input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
    return a.client.CopyBackupToRegion(input)
}

func (a *CloudHSMV2Activities) CreateCluster(input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error) {
    return a.client.CreateCluster(input)
}

func (a *CloudHSMV2Activities) CreateHsm(input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error) {
    return a.client.CreateHsm(input)
}

func (a *CloudHSMV2Activities) DeleteBackup(input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error) {
    return a.client.DeleteBackup(input)
}

func (a *CloudHSMV2Activities) DeleteCluster(input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error) {
    return a.client.DeleteCluster(input)
}

func (a *CloudHSMV2Activities) DeleteHsm(input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error) {
    return a.client.DeleteHsm(input)
}

func (a *CloudHSMV2Activities) DescribeBackups(input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
    return a.client.DescribeBackups(input)
}

func (a *CloudHSMV2Activities) DescribeClusters(input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
    return a.client.DescribeClusters(input)
}

func (a *CloudHSMV2Activities) InitializeCluster(input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error) {
    return a.client.InitializeCluster(input)
}

func (a *CloudHSMV2Activities) ListTags(input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *CloudHSMV2Activities) RestoreBackup(input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error) {
    return a.client.RestoreBackup(input)
}

func (a *CloudHSMV2Activities) TagResource(input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CloudHSMV2Activities) UntagResource(input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
