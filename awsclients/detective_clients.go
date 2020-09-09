package awsclients

import (
	"github.com/aws/aws-sdk-go/service/detective"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DetectiveClient interface {
       AcceptInvitation(ctx workflow.Context, input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error)
       AcceptInvitationAsync(ctx workflow.Context, input *detective.AcceptInvitationInput) *DetectiveAcceptInvitationResult

       CreateGraph(ctx workflow.Context, input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error)
       CreateGraphAsync(ctx workflow.Context, input *detective.CreateGraphInput) *DetectiveCreateGraphResult

       CreateMembers(ctx workflow.Context, input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error)
       CreateMembersAsync(ctx workflow.Context, input *detective.CreateMembersInput) *DetectiveCreateMembersResult

       DeleteGraph(ctx workflow.Context, input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error)
       DeleteGraphAsync(ctx workflow.Context, input *detective.DeleteGraphInput) *DetectiveDeleteGraphResult

       DeleteMembers(ctx workflow.Context, input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error)
       DeleteMembersAsync(ctx workflow.Context, input *detective.DeleteMembersInput) *DetectiveDeleteMembersResult

       DisassociateMembership(ctx workflow.Context, input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error)
       DisassociateMembershipAsync(ctx workflow.Context, input *detective.DisassociateMembershipInput) *DetectiveDisassociateMembershipResult

       GetMembers(ctx workflow.Context, input *detective.GetMembersInput) (*detective.GetMembersOutput, error)
       GetMembersAsync(ctx workflow.Context, input *detective.GetMembersInput) *DetectiveGetMembersResult

       ListGraphs(ctx workflow.Context, input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error)
       ListGraphsAsync(ctx workflow.Context, input *detective.ListGraphsInput) *DetectiveListGraphsResult

       ListInvitations(ctx workflow.Context, input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error)
       ListInvitationsAsync(ctx workflow.Context, input *detective.ListInvitationsInput) *DetectiveListInvitationsResult

       ListMembers(ctx workflow.Context, input *detective.ListMembersInput) (*detective.ListMembersOutput, error)
       ListMembersAsync(ctx workflow.Context, input *detective.ListMembersInput) *DetectiveListMembersResult

       RejectInvitation(ctx workflow.Context, input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error)
       RejectInvitationAsync(ctx workflow.Context, input *detective.RejectInvitationInput) *DetectiveRejectInvitationResult

       StartMonitoringMember(ctx workflow.Context, input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error)
       StartMonitoringMemberAsync(ctx workflow.Context, input *detective.StartMonitoringMemberInput) *DetectiveStartMonitoringMemberResult
}

type DetectiveAcceptInvitationResult struct {
	Result workflow.Future
}

func (r *DetectiveAcceptInvitationResult) Get(ctx workflow.Context) (*detective.AcceptInvitationOutput, error) {
    var output detective.AcceptInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveCreateGraphResult struct {
	Result workflow.Future
}

func (r *DetectiveCreateGraphResult) Get(ctx workflow.Context) (*detective.CreateGraphOutput, error) {
    var output detective.CreateGraphOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveCreateMembersResult struct {
	Result workflow.Future
}

func (r *DetectiveCreateMembersResult) Get(ctx workflow.Context) (*detective.CreateMembersOutput, error) {
    var output detective.CreateMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveDeleteGraphResult struct {
	Result workflow.Future
}

func (r *DetectiveDeleteGraphResult) Get(ctx workflow.Context) (*detective.DeleteGraphOutput, error) {
    var output detective.DeleteGraphOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveDeleteMembersResult struct {
	Result workflow.Future
}

func (r *DetectiveDeleteMembersResult) Get(ctx workflow.Context) (*detective.DeleteMembersOutput, error) {
    var output detective.DeleteMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveDisassociateMembershipResult struct {
	Result workflow.Future
}

func (r *DetectiveDisassociateMembershipResult) Get(ctx workflow.Context) (*detective.DisassociateMembershipOutput, error) {
    var output detective.DisassociateMembershipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveGetMembersResult struct {
	Result workflow.Future
}

func (r *DetectiveGetMembersResult) Get(ctx workflow.Context) (*detective.GetMembersOutput, error) {
    var output detective.GetMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveListGraphsResult struct {
	Result workflow.Future
}

func (r *DetectiveListGraphsResult) Get(ctx workflow.Context) (*detective.ListGraphsOutput, error) {
    var output detective.ListGraphsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveListInvitationsResult struct {
	Result workflow.Future
}

func (r *DetectiveListInvitationsResult) Get(ctx workflow.Context) (*detective.ListInvitationsOutput, error) {
    var output detective.ListInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveListMembersResult struct {
	Result workflow.Future
}

func (r *DetectiveListMembersResult) Get(ctx workflow.Context) (*detective.ListMembersOutput, error) {
    var output detective.ListMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveRejectInvitationResult struct {
	Result workflow.Future
}

func (r *DetectiveRejectInvitationResult) Get(ctx workflow.Context) (*detective.RejectInvitationOutput, error) {
    var output detective.RejectInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveStartMonitoringMemberResult struct {
	Result workflow.Future
}

func (r *DetectiveStartMonitoringMemberResult) Get(ctx workflow.Context) (*detective.StartMonitoringMemberOutput, error) {
    var output detective.StartMonitoringMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DetectiveStub struct {
    activities awsactivities.DetectiveActivities
}

func NewDetectiveStub() DetectiveClient {
    return &DetectiveStub{}
}

func (a *DetectiveStub) AcceptInvitation(ctx workflow.Context, input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error) {
    var output detective.AcceptInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) AcceptInvitationAsync(ctx workflow.Context, input *detective.AcceptInvitationInput) *DetectiveAcceptInvitationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptInvitation, input)
    return &DetectiveAcceptInvitationResult{Result: future}
}

func (a *DetectiveStub) CreateGraph(ctx workflow.Context, input *detective.CreateGraphInput) (*detective.CreateGraphOutput, error) {
    var output detective.CreateGraphOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGraph, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) CreateGraphAsync(ctx workflow.Context, input *detective.CreateGraphInput) *DetectiveCreateGraphResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGraph, input)
    return &DetectiveCreateGraphResult{Result: future}
}

func (a *DetectiveStub) CreateMembers(ctx workflow.Context, input *detective.CreateMembersInput) (*detective.CreateMembersOutput, error) {
    var output detective.CreateMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) CreateMembersAsync(ctx workflow.Context, input *detective.CreateMembersInput) *DetectiveCreateMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMembers, input)
    return &DetectiveCreateMembersResult{Result: future}
}

func (a *DetectiveStub) DeleteGraph(ctx workflow.Context, input *detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error) {
    var output detective.DeleteGraphOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGraph, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) DeleteGraphAsync(ctx workflow.Context, input *detective.DeleteGraphInput) *DetectiveDeleteGraphResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGraph, input)
    return &DetectiveDeleteGraphResult{Result: future}
}

func (a *DetectiveStub) DeleteMembers(ctx workflow.Context, input *detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error) {
    var output detective.DeleteMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) DeleteMembersAsync(ctx workflow.Context, input *detective.DeleteMembersInput) *DetectiveDeleteMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMembers, input)
    return &DetectiveDeleteMembersResult{Result: future}
}

func (a *DetectiveStub) DisassociateMembership(ctx workflow.Context, input *detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error) {
    var output detective.DisassociateMembershipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateMembership, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) DisassociateMembershipAsync(ctx workflow.Context, input *detective.DisassociateMembershipInput) *DetectiveDisassociateMembershipResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateMembership, input)
    return &DetectiveDisassociateMembershipResult{Result: future}
}

func (a *DetectiveStub) GetMembers(ctx workflow.Context, input *detective.GetMembersInput) (*detective.GetMembersOutput, error) {
    var output detective.GetMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) GetMembersAsync(ctx workflow.Context, input *detective.GetMembersInput) *DetectiveGetMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMembers, input)
    return &DetectiveGetMembersResult{Result: future}
}

func (a *DetectiveStub) ListGraphs(ctx workflow.Context, input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error) {
    var output detective.ListGraphsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGraphs, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) ListGraphsAsync(ctx workflow.Context, input *detective.ListGraphsInput) *DetectiveListGraphsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGraphs, input)
    return &DetectiveListGraphsResult{Result: future}
}

func (a *DetectiveStub) ListInvitations(ctx workflow.Context, input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error) {
    var output detective.ListInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) ListInvitationsAsync(ctx workflow.Context, input *detective.ListInvitationsInput) *DetectiveListInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input)
    return &DetectiveListInvitationsResult{Result: future}
}

func (a *DetectiveStub) ListMembers(ctx workflow.Context, input *detective.ListMembersInput) (*detective.ListMembersOutput, error) {
    var output detective.ListMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) ListMembersAsync(ctx workflow.Context, input *detective.ListMembersInput) *DetectiveListMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input)
    return &DetectiveListMembersResult{Result: future}
}

func (a *DetectiveStub) RejectInvitation(ctx workflow.Context, input *detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error) {
    var output detective.RejectInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) RejectInvitationAsync(ctx workflow.Context, input *detective.RejectInvitationInput) *DetectiveRejectInvitationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectInvitation, input)
    return &DetectiveRejectInvitationResult{Result: future}
}

func (a *DetectiveStub) StartMonitoringMember(ctx workflow.Context, input *detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error) {
    var output detective.StartMonitoringMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMonitoringMember, input).Get(ctx, &output)
    return &output, err
}

func (a *DetectiveStub) StartMonitoringMemberAsync(ctx workflow.Context, input *detective.StartMonitoringMemberInput) *DetectiveStartMonitoringMemberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartMonitoringMember, input)
    return &DetectiveStartMonitoringMemberResult{Result: future}
}
