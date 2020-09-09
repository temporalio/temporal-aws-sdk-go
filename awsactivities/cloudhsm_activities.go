
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface"
)

type CloudHSMActivities struct {
	client cloudhsmiface.CloudHSMAPI
}

func NewCloudHSMActivities(session *session.Session, config... *aws.Config) *CloudHSMActivities {
    client := cloudhsm.New(session, config...)
    return &CloudHSMActivities{client: client}
}

func (a *CloudHSMActivities) AddTagsToResource(input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResource(input)
}

func (a *CloudHSMActivities) CreateHapg(input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error) {
    return a.client.CreateHapg(input)
}

func (a *CloudHSMActivities) CreateHsm(input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error) {
    return a.client.CreateHsm(input)
}

func (a *CloudHSMActivities) CreateLunaClient(input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error) {
    return a.client.CreateLunaClient(input)
}

func (a *CloudHSMActivities) DeleteHapg(input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error) {
    return a.client.DeleteHapg(input)
}

func (a *CloudHSMActivities) DeleteHsm(input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error) {
    return a.client.DeleteHsm(input)
}

func (a *CloudHSMActivities) DeleteLunaClient(input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error) {
    return a.client.DeleteLunaClient(input)
}

func (a *CloudHSMActivities) DescribeHapg(input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
    return a.client.DescribeHapg(input)
}

func (a *CloudHSMActivities) DescribeHsm(input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
    return a.client.DescribeHsm(input)
}

func (a *CloudHSMActivities) DescribeLunaClient(input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
    return a.client.DescribeLunaClient(input)
}

func (a *CloudHSMActivities) GetConfig(input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
    return a.client.GetConfig(input)
}

func (a *CloudHSMActivities) ListAvailableZones(input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
    return a.client.ListAvailableZones(input)
}

func (a *CloudHSMActivities) ListHapgs(input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
    return a.client.ListHapgs(input)
}

func (a *CloudHSMActivities) ListHsms(input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
    return a.client.ListHsms(input)
}

func (a *CloudHSMActivities) ListLunaClients(input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
    return a.client.ListLunaClients(input)
}

func (a *CloudHSMActivities) ListTagsForResource(input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CloudHSMActivities) ModifyHapg(input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error) {
    return a.client.ModifyHapg(input)
}

func (a *CloudHSMActivities) ModifyHsm(input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error) {
    return a.client.ModifyHsm(input)
}

func (a *CloudHSMActivities) ModifyLunaClient(input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error) {
    return a.client.ModifyLunaClient(input)
}

func (a *CloudHSMActivities) RemoveTagsFromResource(input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResource(input)
}
