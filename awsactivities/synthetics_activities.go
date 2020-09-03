package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/synthetics"
	"github.com/aws/aws-sdk-go/service/synthetics/syntheticsiface"
)

type SyntheticsActivities struct {
	client syntheticsiface.SyntheticsAPI
}

func NewSyntheticsActivities(client syntheticsiface.SyntheticsAPI) *SyntheticsActivities {
    return &SyntheticsActivities{client: client}
}

func (a *SyntheticsActivities) CreateCanary(input *synthetics.CreateCanaryInput) (*synthetics.CreateCanaryOutput, error) {
    return a.client.CreateCanary(input)
}

func (a *SyntheticsActivities) DeleteCanary(input *synthetics.DeleteCanaryInput) (*synthetics.DeleteCanaryOutput, error) {
    return a.client.DeleteCanary(input)
}

func (a *SyntheticsActivities) DescribeCanaries(input *synthetics.DescribeCanariesInput) (*synthetics.DescribeCanariesOutput, error) {
    return a.client.DescribeCanaries(input)
}

func (a *SyntheticsActivities) DescribeCanariesLastRun(input *synthetics.DescribeCanariesLastRunInput) (*synthetics.DescribeCanariesLastRunOutput, error) {
    return a.client.DescribeCanariesLastRun(input)
}

func (a *SyntheticsActivities) DescribeRuntimeVersions(input *synthetics.DescribeRuntimeVersionsInput) (*synthetics.DescribeRuntimeVersionsOutput, error) {
    return a.client.DescribeRuntimeVersions(input)
}

func (a *SyntheticsActivities) GetCanary(input *synthetics.GetCanaryInput) (*synthetics.GetCanaryOutput, error) {
    return a.client.GetCanary(input)
}

func (a *SyntheticsActivities) GetCanaryRuns(input *synthetics.GetCanaryRunsInput) (*synthetics.GetCanaryRunsOutput, error) {
    return a.client.GetCanaryRuns(input)
}

func (a *SyntheticsActivities) ListTagsForResource(input *synthetics.ListTagsForResourceInput) (*synthetics.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SyntheticsActivities) StartCanary(input *synthetics.StartCanaryInput) (*synthetics.StartCanaryOutput, error) {
    return a.client.StartCanary(input)
}

func (a *SyntheticsActivities) StopCanary(input *synthetics.StopCanaryInput) (*synthetics.StopCanaryOutput, error) {
    return a.client.StopCanary(input)
}

func (a *SyntheticsActivities) TagResource(input *synthetics.TagResourceInput) (*synthetics.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SyntheticsActivities) UntagResource(input *synthetics.UntagResourceInput) (*synthetics.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *SyntheticsActivities) UpdateCanary(input *synthetics.UpdateCanaryInput) (*synthetics.UpdateCanaryOutput, error) {
    return a.client.UpdateCanary(input)
}
