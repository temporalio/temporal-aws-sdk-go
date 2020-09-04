package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codestar"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CodeStarClient interface {
    AssociateTeamMember(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error)
    AssociateTeamMemberAsync(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) *CodestarAssociateTeamMemberResult

    CreateProject(ctx workflow.Context, input *codestar.CreateProjectInput) (*codestar.CreateProjectOutput, error)
    CreateProjectAsync(ctx workflow.Context, input *codestar.CreateProjectInput) *CodestarCreateProjectResult

    CreateUserProfile(ctx workflow.Context, input *codestar.CreateUserProfileInput) (*codestar.CreateUserProfileOutput, error)
    CreateUserProfileAsync(ctx workflow.Context, input *codestar.CreateUserProfileInput) *CodestarCreateUserProfileResult

    DeleteProject(ctx workflow.Context, input *codestar.DeleteProjectInput) (*codestar.DeleteProjectOutput, error)
    DeleteProjectAsync(ctx workflow.Context, input *codestar.DeleteProjectInput) *CodestarDeleteProjectResult

    DeleteUserProfile(ctx workflow.Context, input *codestar.DeleteUserProfileInput) (*codestar.DeleteUserProfileOutput, error)
    DeleteUserProfileAsync(ctx workflow.Context, input *codestar.DeleteUserProfileInput) *CodestarDeleteUserProfileResult

    DescribeProject(ctx workflow.Context, input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error)
    DescribeProjectAsync(ctx workflow.Context, input *codestar.DescribeProjectInput) *CodestarDescribeProjectResult

    DescribeUserProfile(ctx workflow.Context, input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error)
    DescribeUserProfileAsync(ctx workflow.Context, input *codestar.DescribeUserProfileInput) *CodestarDescribeUserProfileResult

    DisassociateTeamMember(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) (*codestar.DisassociateTeamMemberOutput, error)
    DisassociateTeamMemberAsync(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) *CodestarDisassociateTeamMemberResult

    ListProjects(ctx workflow.Context, input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error)
    ListProjectsAsync(ctx workflow.Context, input *codestar.ListProjectsInput) *CodestarListProjectsResult

    ListResources(ctx workflow.Context, input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error)
    ListResourcesAsync(ctx workflow.Context, input *codestar.ListResourcesInput) *CodestarListResourcesResult

    ListTagsForProject(ctx workflow.Context, input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error)
    ListTagsForProjectAsync(ctx workflow.Context, input *codestar.ListTagsForProjectInput) *CodestarListTagsForProjectResult

    ListTeamMembers(ctx workflow.Context, input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error)
    ListTeamMembersAsync(ctx workflow.Context, input *codestar.ListTeamMembersInput) *CodestarListTeamMembersResult

    ListUserProfiles(ctx workflow.Context, input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error)
    ListUserProfilesAsync(ctx workflow.Context, input *codestar.ListUserProfilesInput) *CodestarListUserProfilesResult

    TagProject(ctx workflow.Context, input *codestar.TagProjectInput) (*codestar.TagProjectOutput, error)
    TagProjectAsync(ctx workflow.Context, input *codestar.TagProjectInput) *CodestarTagProjectResult

    UntagProject(ctx workflow.Context, input *codestar.UntagProjectInput) (*codestar.UntagProjectOutput, error)
    UntagProjectAsync(ctx workflow.Context, input *codestar.UntagProjectInput) *CodestarUntagProjectResult

    UpdateProject(ctx workflow.Context, input *codestar.UpdateProjectInput) (*codestar.UpdateProjectOutput, error)
    UpdateProjectAsync(ctx workflow.Context, input *codestar.UpdateProjectInput) *CodestarUpdateProjectResult

    UpdateTeamMember(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) (*codestar.UpdateTeamMemberOutput, error)
    UpdateTeamMemberAsync(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) *CodestarUpdateTeamMemberResult

    UpdateUserProfile(ctx workflow.Context, input *codestar.UpdateUserProfileInput) (*codestar.UpdateUserProfileOutput, error)
    UpdateUserProfileAsync(ctx workflow.Context, input *codestar.UpdateUserProfileInput) *CodestarUpdateUserProfileResult
}

type CodestarAssociateTeamMemberResult struct {
	Result workflow.Future
}

func (r *CodestarAssociateTeamMemberResult) Get(ctx workflow.Context) (*codestar.AssociateTeamMemberOutput, error) {
    var output codestar.AssociateTeamMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarCreateProjectResult struct {
	Result workflow.Future
}

func (r *CodestarCreateProjectResult) Get(ctx workflow.Context) (*codestar.CreateProjectOutput, error) {
    var output codestar.CreateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarCreateUserProfileResult struct {
	Result workflow.Future
}

func (r *CodestarCreateUserProfileResult) Get(ctx workflow.Context) (*codestar.CreateUserProfileOutput, error) {
    var output codestar.CreateUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarDeleteProjectResult struct {
	Result workflow.Future
}

func (r *CodestarDeleteProjectResult) Get(ctx workflow.Context) (*codestar.DeleteProjectOutput, error) {
    var output codestar.DeleteProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarDeleteUserProfileResult struct {
	Result workflow.Future
}

func (r *CodestarDeleteUserProfileResult) Get(ctx workflow.Context) (*codestar.DeleteUserProfileOutput, error) {
    var output codestar.DeleteUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarDescribeProjectResult struct {
	Result workflow.Future
}

func (r *CodestarDescribeProjectResult) Get(ctx workflow.Context) (*codestar.DescribeProjectOutput, error) {
    var output codestar.DescribeProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarDescribeUserProfileResult struct {
	Result workflow.Future
}

func (r *CodestarDescribeUserProfileResult) Get(ctx workflow.Context) (*codestar.DescribeUserProfileOutput, error) {
    var output codestar.DescribeUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarDisassociateTeamMemberResult struct {
	Result workflow.Future
}

func (r *CodestarDisassociateTeamMemberResult) Get(ctx workflow.Context) (*codestar.DisassociateTeamMemberOutput, error) {
    var output codestar.DisassociateTeamMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarListProjectsResult struct {
	Result workflow.Future
}

func (r *CodestarListProjectsResult) Get(ctx workflow.Context) (*codestar.ListProjectsOutput, error) {
    var output codestar.ListProjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarListResourcesResult struct {
	Result workflow.Future
}

func (r *CodestarListResourcesResult) Get(ctx workflow.Context) (*codestar.ListResourcesOutput, error) {
    var output codestar.ListResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarListTagsForProjectResult struct {
	Result workflow.Future
}

func (r *CodestarListTagsForProjectResult) Get(ctx workflow.Context) (*codestar.ListTagsForProjectOutput, error) {
    var output codestar.ListTagsForProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarListTeamMembersResult struct {
	Result workflow.Future
}

func (r *CodestarListTeamMembersResult) Get(ctx workflow.Context) (*codestar.ListTeamMembersOutput, error) {
    var output codestar.ListTeamMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarListUserProfilesResult struct {
	Result workflow.Future
}

func (r *CodestarListUserProfilesResult) Get(ctx workflow.Context) (*codestar.ListUserProfilesOutput, error) {
    var output codestar.ListUserProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarTagProjectResult struct {
	Result workflow.Future
}

func (r *CodestarTagProjectResult) Get(ctx workflow.Context) (*codestar.TagProjectOutput, error) {
    var output codestar.TagProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarUntagProjectResult struct {
	Result workflow.Future
}

func (r *CodestarUntagProjectResult) Get(ctx workflow.Context) (*codestar.UntagProjectOutput, error) {
    var output codestar.UntagProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarUpdateProjectResult struct {
	Result workflow.Future
}

func (r *CodestarUpdateProjectResult) Get(ctx workflow.Context) (*codestar.UpdateProjectOutput, error) {
    var output codestar.UpdateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarUpdateTeamMemberResult struct {
	Result workflow.Future
}

func (r *CodestarUpdateTeamMemberResult) Get(ctx workflow.Context) (*codestar.UpdateTeamMemberOutput, error) {
    var output codestar.UpdateTeamMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodestarUpdateUserProfileResult struct {
	Result workflow.Future
}

func (r *CodestarUpdateUserProfileResult) Get(ctx workflow.Context) (*codestar.UpdateUserProfileOutput, error) {
    var output codestar.UpdateUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeStarStub struct {
    activities awsactivities.CodeStarActivities
}

func NewCodeStarStub() CodeStarClient {
    return &CodeStarStub{}
}

func (a *CodeStarStub) AssociateTeamMember(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error) {
    var output codestar.AssociateTeamMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateTeamMember, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) AssociateTeamMemberAsync(ctx workflow.Context, input *codestar.AssociateTeamMemberInput) *CodestarAssociateTeamMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateTeamMember, input)
    return &CodestarAssociateTeamMemberResult{Result: future}
}

func (a *CodeStarStub) CreateProject(ctx workflow.Context, input *codestar.CreateProjectInput) (*codestar.CreateProjectOutput, error) {
    var output codestar.CreateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) CreateProjectAsync(ctx workflow.Context, input *codestar.CreateProjectInput) *CodestarCreateProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input)
    return &CodestarCreateProjectResult{Result: future}
}

func (a *CodeStarStub) CreateUserProfile(ctx workflow.Context, input *codestar.CreateUserProfileInput) (*codestar.CreateUserProfileOutput, error) {
    var output codestar.CreateUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) CreateUserProfileAsync(ctx workflow.Context, input *codestar.CreateUserProfileInput) *CodestarCreateUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUserProfile, input)
    return &CodestarCreateUserProfileResult{Result: future}
}

func (a *CodeStarStub) DeleteProject(ctx workflow.Context, input *codestar.DeleteProjectInput) (*codestar.DeleteProjectOutput, error) {
    var output codestar.DeleteProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) DeleteProjectAsync(ctx workflow.Context, input *codestar.DeleteProjectInput) *CodestarDeleteProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input)
    return &CodestarDeleteProjectResult{Result: future}
}

func (a *CodeStarStub) DeleteUserProfile(ctx workflow.Context, input *codestar.DeleteUserProfileInput) (*codestar.DeleteUserProfileOutput, error) {
    var output codestar.DeleteUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) DeleteUserProfileAsync(ctx workflow.Context, input *codestar.DeleteUserProfileInput) *CodestarDeleteUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUserProfile, input)
    return &CodestarDeleteUserProfileResult{Result: future}
}

func (a *CodeStarStub) DescribeProject(ctx workflow.Context, input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error) {
    var output codestar.DescribeProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProject, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) DescribeProjectAsync(ctx workflow.Context, input *codestar.DescribeProjectInput) *CodestarDescribeProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProject, input)
    return &CodestarDescribeProjectResult{Result: future}
}

func (a *CodeStarStub) DescribeUserProfile(ctx workflow.Context, input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error) {
    var output codestar.DescribeUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) DescribeUserProfileAsync(ctx workflow.Context, input *codestar.DescribeUserProfileInput) *CodestarDescribeUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUserProfile, input)
    return &CodestarDescribeUserProfileResult{Result: future}
}

func (a *CodeStarStub) DisassociateTeamMember(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) (*codestar.DisassociateTeamMemberOutput, error) {
    var output codestar.DisassociateTeamMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateTeamMember, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) DisassociateTeamMemberAsync(ctx workflow.Context, input *codestar.DisassociateTeamMemberInput) *CodestarDisassociateTeamMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateTeamMember, input)
    return &CodestarDisassociateTeamMemberResult{Result: future}
}

func (a *CodeStarStub) ListProjects(ctx workflow.Context, input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error) {
    var output codestar.ListProjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) ListProjectsAsync(ctx workflow.Context, input *codestar.ListProjectsInput) *CodestarListProjectsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input)
    return &CodestarListProjectsResult{Result: future}
}

func (a *CodeStarStub) ListResources(ctx workflow.Context, input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error) {
    var output codestar.ListResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResources, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) ListResourcesAsync(ctx workflow.Context, input *codestar.ListResourcesInput) *CodestarListResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResources, input)
    return &CodestarListResourcesResult{Result: future}
}

func (a *CodeStarStub) ListTagsForProject(ctx workflow.Context, input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error) {
    var output codestar.ListTagsForProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForProject, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) ListTagsForProjectAsync(ctx workflow.Context, input *codestar.ListTagsForProjectInput) *CodestarListTagsForProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForProject, input)
    return &CodestarListTagsForProjectResult{Result: future}
}

func (a *CodeStarStub) ListTeamMembers(ctx workflow.Context, input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error) {
    var output codestar.ListTeamMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTeamMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) ListTeamMembersAsync(ctx workflow.Context, input *codestar.ListTeamMembersInput) *CodestarListTeamMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTeamMembers, input)
    return &CodestarListTeamMembersResult{Result: future}
}

func (a *CodeStarStub) ListUserProfiles(ctx workflow.Context, input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error) {
    var output codestar.ListUserProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) ListUserProfilesAsync(ctx workflow.Context, input *codestar.ListUserProfilesInput) *CodestarListUserProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUserProfiles, input)
    return &CodestarListUserProfilesResult{Result: future}
}

func (a *CodeStarStub) TagProject(ctx workflow.Context, input *codestar.TagProjectInput) (*codestar.TagProjectOutput, error) {
    var output codestar.TagProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagProject, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) TagProjectAsync(ctx workflow.Context, input *codestar.TagProjectInput) *CodestarTagProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagProject, input)
    return &CodestarTagProjectResult{Result: future}
}

func (a *CodeStarStub) UntagProject(ctx workflow.Context, input *codestar.UntagProjectInput) (*codestar.UntagProjectOutput, error) {
    var output codestar.UntagProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagProject, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) UntagProjectAsync(ctx workflow.Context, input *codestar.UntagProjectInput) *CodestarUntagProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagProject, input)
    return &CodestarUntagProjectResult{Result: future}
}

func (a *CodeStarStub) UpdateProject(ctx workflow.Context, input *codestar.UpdateProjectInput) (*codestar.UpdateProjectOutput, error) {
    var output codestar.UpdateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) UpdateProjectAsync(ctx workflow.Context, input *codestar.UpdateProjectInput) *CodestarUpdateProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input)
    return &CodestarUpdateProjectResult{Result: future}
}

func (a *CodeStarStub) UpdateTeamMember(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) (*codestar.UpdateTeamMemberOutput, error) {
    var output codestar.UpdateTeamMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTeamMember, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) UpdateTeamMemberAsync(ctx workflow.Context, input *codestar.UpdateTeamMemberInput) *CodestarUpdateTeamMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTeamMember, input)
    return &CodestarUpdateTeamMemberResult{Result: future}
}

func (a *CodeStarStub) UpdateUserProfile(ctx workflow.Context, input *codestar.UpdateUserProfileInput) (*codestar.UpdateUserProfileOutput, error) {
    var output codestar.UpdateUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarStub) UpdateUserProfileAsync(ctx workflow.Context, input *codestar.UpdateUserProfileInput) *CodestarUpdateUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserProfile, input)
    return &CodestarUpdateUserProfileResult{Result: future}
}
