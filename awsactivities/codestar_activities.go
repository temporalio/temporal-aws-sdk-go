package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/codestar/codestariface"
)

type CodeStarActivities struct {
	client codestariface.CodeStarAPI
}

func NewCodeStarActivities(session *session.Session, config ...*aws.Config) *CodeStarActivities {
	client := codestar.New(session, config...)
	return &CodeStarActivities{client: client}
}

func (a *CodeStarActivities) AssociateTeamMember(input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error) {
	return a.client.AssociateTeamMember(input)
}

func (a *CodeStarActivities) CreateProject(input *codestar.CreateProjectInput) (*codestar.CreateProjectOutput, error) {
	return a.client.CreateProject(input)
}

func (a *CodeStarActivities) CreateUserProfile(input *codestar.CreateUserProfileInput) (*codestar.CreateUserProfileOutput, error) {
	return a.client.CreateUserProfile(input)
}

func (a *CodeStarActivities) DeleteProject(input *codestar.DeleteProjectInput) (*codestar.DeleteProjectOutput, error) {
	return a.client.DeleteProject(input)
}

func (a *CodeStarActivities) DeleteUserProfile(input *codestar.DeleteUserProfileInput) (*codestar.DeleteUserProfileOutput, error) {
	return a.client.DeleteUserProfile(input)
}

func (a *CodeStarActivities) DescribeProject(input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error) {
	return a.client.DescribeProject(input)
}

func (a *CodeStarActivities) DescribeUserProfile(input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error) {
	return a.client.DescribeUserProfile(input)
}

func (a *CodeStarActivities) DisassociateTeamMember(input *codestar.DisassociateTeamMemberInput) (*codestar.DisassociateTeamMemberOutput, error) {
	return a.client.DisassociateTeamMember(input)
}

func (a *CodeStarActivities) ListProjects(input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error) {
	return a.client.ListProjects(input)
}

func (a *CodeStarActivities) ListResources(input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error) {
	return a.client.ListResources(input)
}

func (a *CodeStarActivities) ListTagsForProject(input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error) {
	return a.client.ListTagsForProject(input)
}

func (a *CodeStarActivities) ListTeamMembers(input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error) {
	return a.client.ListTeamMembers(input)
}

func (a *CodeStarActivities) ListUserProfiles(input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error) {
	return a.client.ListUserProfiles(input)
}

func (a *CodeStarActivities) TagProject(input *codestar.TagProjectInput) (*codestar.TagProjectOutput, error) {
	return a.client.TagProject(input)
}

func (a *CodeStarActivities) UntagProject(input *codestar.UntagProjectInput) (*codestar.UntagProjectOutput, error) {
	return a.client.UntagProject(input)
}

func (a *CodeStarActivities) UpdateProject(input *codestar.UpdateProjectInput) (*codestar.UpdateProjectOutput, error) {
	return a.client.UpdateProject(input)
}

func (a *CodeStarActivities) UpdateTeamMember(input *codestar.UpdateTeamMemberInput) (*codestar.UpdateTeamMemberOutput, error) {
	return a.client.UpdateTeamMember(input)
}

func (a *CodeStarActivities) UpdateUserProfile(input *codestar.UpdateUserProfileInput) (*codestar.UpdateUserProfileOutput, error) {
	return a.client.UpdateUserProfile(input)
}
