package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/codestar/codestariface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CodeStarActivities struct {
	client codestariface.CodeStarAPI
}

func NewCodeStarActivities(session *session.Session, config ...*aws.Config) *CodeStarActivities {
	client := codestar.New(session, config...)
	return &CodeStarActivities{client: client}
}

func (a *CodeStarActivities) AssociateTeamMember(ctx context.Context, input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error) {
	return a.client.AssociateTeamMemberWithContext(ctx, input)
}

func (a *CodeStarActivities) CreateProject(ctx context.Context, input *codestar.CreateProjectInput) (*codestar.CreateProjectOutput, error) {
	return a.client.CreateProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) CreateUserProfile(ctx context.Context, input *codestar.CreateUserProfileInput) (*codestar.CreateUserProfileOutput, error) {
	return a.client.CreateUserProfileWithContext(ctx, input)
}

func (a *CodeStarActivities) DeleteProject(ctx context.Context, input *codestar.DeleteProjectInput) (*codestar.DeleteProjectOutput, error) {
	return a.client.DeleteProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) DeleteUserProfile(ctx context.Context, input *codestar.DeleteUserProfileInput) (*codestar.DeleteUserProfileOutput, error) {
	return a.client.DeleteUserProfileWithContext(ctx, input)
}

func (a *CodeStarActivities) DescribeProject(ctx context.Context, input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error) {
	return a.client.DescribeProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) DescribeUserProfile(ctx context.Context, input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error) {
	return a.client.DescribeUserProfileWithContext(ctx, input)
}

func (a *CodeStarActivities) DisassociateTeamMember(ctx context.Context, input *codestar.DisassociateTeamMemberInput) (*codestar.DisassociateTeamMemberOutput, error) {
	return a.client.DisassociateTeamMemberWithContext(ctx, input)
}

func (a *CodeStarActivities) ListProjects(ctx context.Context, input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error) {
	return a.client.ListProjectsWithContext(ctx, input)
}

func (a *CodeStarActivities) ListResources(ctx context.Context, input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error) {
	return a.client.ListResourcesWithContext(ctx, input)
}

func (a *CodeStarActivities) ListTagsForProject(ctx context.Context, input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error) {
	return a.client.ListTagsForProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) ListTeamMembers(ctx context.Context, input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error) {
	return a.client.ListTeamMembersWithContext(ctx, input)
}

func (a *CodeStarActivities) ListUserProfiles(ctx context.Context, input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error) {
	return a.client.ListUserProfilesWithContext(ctx, input)
}

func (a *CodeStarActivities) TagProject(ctx context.Context, input *codestar.TagProjectInput) (*codestar.TagProjectOutput, error) {
	return a.client.TagProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) UntagProject(ctx context.Context, input *codestar.UntagProjectInput) (*codestar.UntagProjectOutput, error) {
	return a.client.UntagProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) UpdateProject(ctx context.Context, input *codestar.UpdateProjectInput) (*codestar.UpdateProjectOutput, error) {
	return a.client.UpdateProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) UpdateTeamMember(ctx context.Context, input *codestar.UpdateTeamMemberInput) (*codestar.UpdateTeamMemberOutput, error) {
	return a.client.UpdateTeamMemberWithContext(ctx, input)
}

func (a *CodeStarActivities) UpdateUserProfile(ctx context.Context, input *codestar.UpdateUserProfileInput) (*codestar.UpdateUserProfileOutput, error) {
	return a.client.UpdateUserProfileWithContext(ctx, input)
}
