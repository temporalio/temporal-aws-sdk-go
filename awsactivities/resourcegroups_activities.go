package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroups/resourcegroupsiface"
)

type ResourceGroupsActivities struct {
	client resourcegroupsiface.ResourceGroupsAPI
}

func NewResourceGroupsActivities(client resourcegroupsiface.ResourceGroupsAPI) *ResourceGroupsActivities {
    return &ResourceGroupsActivities{client: client}
}

func (a *ResourceGroupsActivities) CreateGroup(input *resourcegroups.CreateGroupInput) (*resourcegroups.CreateGroupOutput, error) {
    return a.client.CreateGroup(input)
}

func (a *ResourceGroupsActivities) DeleteGroup(input *resourcegroups.DeleteGroupInput) (*resourcegroups.DeleteGroupOutput, error) {
    return a.client.DeleteGroup(input)
}

func (a *ResourceGroupsActivities) GetGroup(input *resourcegroups.GetGroupInput) (*resourcegroups.GetGroupOutput, error) {
    return a.client.GetGroup(input)
}

func (a *ResourceGroupsActivities) GetGroupConfiguration(input *resourcegroups.GetGroupConfigurationInput) (*resourcegroups.GetGroupConfigurationOutput, error) {
    return a.client.GetGroupConfiguration(input)
}

func (a *ResourceGroupsActivities) GetGroupQuery(input *resourcegroups.GetGroupQueryInput) (*resourcegroups.GetGroupQueryOutput, error) {
    return a.client.GetGroupQuery(input)
}

func (a *ResourceGroupsActivities) GetTags(input *resourcegroups.GetTagsInput) (*resourcegroups.GetTagsOutput, error) {
    return a.client.GetTags(input)
}

func (a *ResourceGroupsActivities) GroupResources(input *resourcegroups.GroupResourcesInput) (*resourcegroups.GroupResourcesOutput, error) {
    return a.client.GroupResources(input)
}

func (a *ResourceGroupsActivities) ListGroupResources(input *resourcegroups.ListGroupResourcesInput) (*resourcegroups.ListGroupResourcesOutput, error) {
    return a.client.ListGroupResources(input)
}

func (a *ResourceGroupsActivities) ListGroups(input *resourcegroups.ListGroupsInput) (*resourcegroups.ListGroupsOutput, error) {
    return a.client.ListGroups(input)
}

func (a *ResourceGroupsActivities) SearchResources(input *resourcegroups.SearchResourcesInput) (*resourcegroups.SearchResourcesOutput, error) {
    return a.client.SearchResources(input)
}

func (a *ResourceGroupsActivities) Tag(input *resourcegroups.TagInput) (*resourcegroups.TagOutput, error) {
    return a.client.Tag(input)
}

func (a *ResourceGroupsActivities) UngroupResources(input *resourcegroups.UngroupResourcesInput) (*resourcegroups.UngroupResourcesOutput, error) {
    return a.client.UngroupResources(input)
}

func (a *ResourceGroupsActivities) Untag(input *resourcegroups.UntagInput) (*resourcegroups.UntagOutput, error) {
    return a.client.Untag(input)
}

func (a *ResourceGroupsActivities) UpdateGroup(input *resourcegroups.UpdateGroupInput) (*resourcegroups.UpdateGroupOutput, error) {
    return a.client.UpdateGroup(input)
}

func (a *ResourceGroupsActivities) UpdateGroupQuery(input *resourcegroups.UpdateGroupQueryInput) (*resourcegroups.UpdateGroupQueryOutput, error) {
    return a.client.UpdateGroupQuery(input)
}
