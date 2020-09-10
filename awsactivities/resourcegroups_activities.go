package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroups/resourcegroupsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ResourceGroupsActivities struct {
	client resourcegroupsiface.ResourceGroupsAPI
}

func NewResourceGroupsActivities(session *session.Session, config ...*aws.Config) *ResourceGroupsActivities {
	client := resourcegroups.New(session, config...)
	return &ResourceGroupsActivities{client: client}
}

func (a *ResourceGroupsActivities) CreateGroup(ctx context.Context, input *resourcegroups.CreateGroupInput) (*resourcegroups.CreateGroupOutput, error) {
	return a.client.CreateGroupWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) DeleteGroup(ctx context.Context, input *resourcegroups.DeleteGroupInput) (*resourcegroups.DeleteGroupOutput, error) {
	return a.client.DeleteGroupWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) GetGroup(ctx context.Context, input *resourcegroups.GetGroupInput) (*resourcegroups.GetGroupOutput, error) {
	return a.client.GetGroupWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) GetGroupConfiguration(ctx context.Context, input *resourcegroups.GetGroupConfigurationInput) (*resourcegroups.GetGroupConfigurationOutput, error) {
	return a.client.GetGroupConfigurationWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) GetGroupQuery(ctx context.Context, input *resourcegroups.GetGroupQueryInput) (*resourcegroups.GetGroupQueryOutput, error) {
	return a.client.GetGroupQueryWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) GetTags(ctx context.Context, input *resourcegroups.GetTagsInput) (*resourcegroups.GetTagsOutput, error) {
	return a.client.GetTagsWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) GroupResources(ctx context.Context, input *resourcegroups.GroupResourcesInput) (*resourcegroups.GroupResourcesOutput, error) {
	return a.client.GroupResourcesWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) ListGroupResources(ctx context.Context, input *resourcegroups.ListGroupResourcesInput) (*resourcegroups.ListGroupResourcesOutput, error) {
	return a.client.ListGroupResourcesWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) ListGroups(ctx context.Context, input *resourcegroups.ListGroupsInput) (*resourcegroups.ListGroupsOutput, error) {
	return a.client.ListGroupsWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) SearchResources(ctx context.Context, input *resourcegroups.SearchResourcesInput) (*resourcegroups.SearchResourcesOutput, error) {
	return a.client.SearchResourcesWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) Tag(ctx context.Context, input *resourcegroups.TagInput) (*resourcegroups.TagOutput, error) {
	return a.client.TagWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) UngroupResources(ctx context.Context, input *resourcegroups.UngroupResourcesInput) (*resourcegroups.UngroupResourcesOutput, error) {
	return a.client.UngroupResourcesWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) Untag(ctx context.Context, input *resourcegroups.UntagInput) (*resourcegroups.UntagOutput, error) {
	return a.client.UntagWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) UpdateGroup(ctx context.Context, input *resourcegroups.UpdateGroupInput) (*resourcegroups.UpdateGroupOutput, error) {
	return a.client.UpdateGroupWithContext(ctx, input)
}

func (a *ResourceGroupsActivities) UpdateGroupQuery(ctx context.Context, input *resourcegroups.UpdateGroupQueryInput) (*resourcegroups.UpdateGroupQueryOutput, error) {
	return a.client.UpdateGroupQueryWithContext(ctx, input)
}
