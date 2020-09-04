
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface"
)

type ResourceGroupsTaggingAPIActivities struct {
	client resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI
}

func NewResourceGroupsTaggingAPIActivities(client resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI) *ResourceGroupsTaggingAPIActivities {
    return &ResourceGroupsTaggingAPIActivities{client: client}
}

func (a *ResourceGroupsTaggingAPIActivities) DescribeReportCreation(input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
    return a.client.DescribeReportCreation(input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetComplianceSummary(input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
    return a.client.GetComplianceSummary(input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetResources(input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
    return a.client.GetResources(input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetTagKeys(input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
    return a.client.GetTagKeys(input)
}

func (a *ResourceGroupsTaggingAPIActivities) GetTagValues(input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
    return a.client.GetTagValues(input)
}

func (a *ResourceGroupsTaggingAPIActivities) StartReportCreation(input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
    return a.client.StartReportCreation(input)
}

func (a *ResourceGroupsTaggingAPIActivities) TagResources(input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
    return a.client.TagResources(input)
}

func (a *ResourceGroupsTaggingAPIActivities) UntagResources(input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
    return a.client.UntagResources(input)
}
