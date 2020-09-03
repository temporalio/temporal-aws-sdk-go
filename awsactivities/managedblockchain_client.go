package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"go.temporal.io/sdk/workflow"
)

type ManagedBlockchainClient interface {
    CreateMember(ctx workflow.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error)
    CreateMemberAsync(ctx workflow.Context, input *managedblockchain.CreateMemberInput) *ManagedblockchainCreateMemberResult

    CreateNetwork(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error)
    CreateNetworkAsync(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) *ManagedblockchainCreateNetworkResult

    CreateNode(ctx workflow.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error)
    CreateNodeAsync(ctx workflow.Context, input *managedblockchain.CreateNodeInput) *ManagedblockchainCreateNodeResult

    CreateProposal(ctx workflow.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error)
    CreateProposalAsync(ctx workflow.Context, input *managedblockchain.CreateProposalInput) *ManagedblockchainCreateProposalResult

    DeleteMember(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error)
    DeleteMemberAsync(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) *ManagedblockchainDeleteMemberResult

    DeleteNode(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error)
    DeleteNodeAsync(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) *ManagedblockchainDeleteNodeResult

    GetMember(ctx workflow.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error)
    GetMemberAsync(ctx workflow.Context, input *managedblockchain.GetMemberInput) *ManagedblockchainGetMemberResult

    GetNetwork(ctx workflow.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error)
    GetNetworkAsync(ctx workflow.Context, input *managedblockchain.GetNetworkInput) *ManagedblockchainGetNetworkResult

    GetNode(ctx workflow.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error)
    GetNodeAsync(ctx workflow.Context, input *managedblockchain.GetNodeInput) *ManagedblockchainGetNodeResult

    GetProposal(ctx workflow.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error)
    GetProposalAsync(ctx workflow.Context, input *managedblockchain.GetProposalInput) *ManagedblockchainGetProposalResult

    ListInvitations(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error)
    ListInvitationsAsync(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) *ManagedblockchainListInvitationsResult

    ListMembers(ctx workflow.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error)
    ListMembersAsync(ctx workflow.Context, input *managedblockchain.ListMembersInput) *ManagedblockchainListMembersResult

    ListNetworks(ctx workflow.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error)
    ListNetworksAsync(ctx workflow.Context, input *managedblockchain.ListNetworksInput) *ManagedblockchainListNetworksResult

    ListNodes(ctx workflow.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error)
    ListNodesAsync(ctx workflow.Context, input *managedblockchain.ListNodesInput) *ManagedblockchainListNodesResult

    ListProposalVotes(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error)
    ListProposalVotesAsync(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) *ManagedblockchainListProposalVotesResult

    ListProposals(ctx workflow.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error)
    ListProposalsAsync(ctx workflow.Context, input *managedblockchain.ListProposalsInput) *ManagedblockchainListProposalsResult

    RejectInvitation(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error)
    RejectInvitationAsync(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) *ManagedblockchainRejectInvitationResult

    UpdateMember(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error)
    UpdateMemberAsync(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) *ManagedblockchainUpdateMemberResult

    UpdateNode(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error)
    UpdateNodeAsync(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) *ManagedblockchainUpdateNodeResult

    VoteOnProposal(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error)
    VoteOnProposalAsync(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) *ManagedblockchainVoteOnProposalResult
}
type ManagedblockchainCreateMemberResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainCreateMemberResult) Get(ctx workflow.Context) (*managedblockchain.CreateMemberOutput, error) {
    var output managedblockchain.CreateMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainCreateNetworkResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainCreateNetworkResult) Get(ctx workflow.Context) (*managedblockchain.CreateNetworkOutput, error) {
    var output managedblockchain.CreateNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainCreateNodeResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainCreateNodeResult) Get(ctx workflow.Context) (*managedblockchain.CreateNodeOutput, error) {
    var output managedblockchain.CreateNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainCreateProposalResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainCreateProposalResult) Get(ctx workflow.Context) (*managedblockchain.CreateProposalOutput, error) {
    var output managedblockchain.CreateProposalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainDeleteMemberResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainDeleteMemberResult) Get(ctx workflow.Context) (*managedblockchain.DeleteMemberOutput, error) {
    var output managedblockchain.DeleteMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainDeleteNodeResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainDeleteNodeResult) Get(ctx workflow.Context) (*managedblockchain.DeleteNodeOutput, error) {
    var output managedblockchain.DeleteNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainGetMemberResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainGetMemberResult) Get(ctx workflow.Context) (*managedblockchain.GetMemberOutput, error) {
    var output managedblockchain.GetMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainGetNetworkResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainGetNetworkResult) Get(ctx workflow.Context) (*managedblockchain.GetNetworkOutput, error) {
    var output managedblockchain.GetNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainGetNodeResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainGetNodeResult) Get(ctx workflow.Context) (*managedblockchain.GetNodeOutput, error) {
    var output managedblockchain.GetNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainGetProposalResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainGetProposalResult) Get(ctx workflow.Context) (*managedblockchain.GetProposalOutput, error) {
    var output managedblockchain.GetProposalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainListInvitationsResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainListInvitationsResult) Get(ctx workflow.Context) (*managedblockchain.ListInvitationsOutput, error) {
    var output managedblockchain.ListInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainListMembersResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainListMembersResult) Get(ctx workflow.Context) (*managedblockchain.ListMembersOutput, error) {
    var output managedblockchain.ListMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainListNetworksResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainListNetworksResult) Get(ctx workflow.Context) (*managedblockchain.ListNetworksOutput, error) {
    var output managedblockchain.ListNetworksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainListNodesResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainListNodesResult) Get(ctx workflow.Context) (*managedblockchain.ListNodesOutput, error) {
    var output managedblockchain.ListNodesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainListProposalVotesResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainListProposalVotesResult) Get(ctx workflow.Context) (*managedblockchain.ListProposalVotesOutput, error) {
    var output managedblockchain.ListProposalVotesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainListProposalsResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainListProposalsResult) Get(ctx workflow.Context) (*managedblockchain.ListProposalsOutput, error) {
    var output managedblockchain.ListProposalsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainRejectInvitationResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainRejectInvitationResult) Get(ctx workflow.Context) (*managedblockchain.RejectInvitationOutput, error) {
    var output managedblockchain.RejectInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainUpdateMemberResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainUpdateMemberResult) Get(ctx workflow.Context) (*managedblockchain.UpdateMemberOutput, error) {
    var output managedblockchain.UpdateMemberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainUpdateNodeResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainUpdateNodeResult) Get(ctx workflow.Context) (*managedblockchain.UpdateNodeOutput, error) {
    var output managedblockchain.UpdateNodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ManagedblockchainVoteOnProposalResult struct {
	Result workflow.Future
}

func (r *ManagedblockchainVoteOnProposalResult) Get(ctx workflow.Context) (*managedblockchain.VoteOnProposalOutput, error) {
    var output managedblockchain.VoteOnProposalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ManagedBlockchainStub struct {
    activities ManagedBlockchainClient
}

func NewManagedBlockchainStub() ManagedBlockchainClient {
    return &ManagedBlockchainStub{}
}

func (a *ManagedBlockchainStub) CreateMember(ctx workflow.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error) {
    var output managedblockchain.CreateMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMember, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) CreateNetwork(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error) {
    var output managedblockchain.CreateNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) CreateNode(ctx workflow.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error) {
    var output managedblockchain.CreateNodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNode, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) CreateProposal(ctx workflow.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error) {
    var output managedblockchain.CreateProposalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProposal, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) DeleteMember(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error) {
    var output managedblockchain.DeleteMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMember, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) DeleteNode(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error) {
    var output managedblockchain.DeleteNodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNode, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) GetMember(ctx workflow.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
    var output managedblockchain.GetMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMember, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) GetNetwork(ctx workflow.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
    var output managedblockchain.GetNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) GetNode(ctx workflow.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
    var output managedblockchain.GetNodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNode, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) GetProposal(ctx workflow.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
    var output managedblockchain.GetProposalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProposal, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) ListInvitations(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
    var output managedblockchain.ListInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) ListMembers(ctx workflow.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
    var output managedblockchain.ListMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) ListNetworks(ctx workflow.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
    var output managedblockchain.ListNetworksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNetworks, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) ListNodes(ctx workflow.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
    var output managedblockchain.ListNodesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNodes, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) ListProposalVotes(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
    var output managedblockchain.ListProposalVotesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProposalVotes, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) ListProposals(ctx workflow.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
    var output managedblockchain.ListProposalsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProposals, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) RejectInvitation(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error) {
    var output managedblockchain.RejectInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) UpdateMember(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error) {
    var output managedblockchain.UpdateMemberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMember, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) UpdateNode(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error) {
    var output managedblockchain.UpdateNodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNode, input).Get(ctx, &output)
    return &output, err
}

func (a *ManagedBlockchainStub) VoteOnProposal(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error) {
    var output managedblockchain.VoteOnProposalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.VoteOnProposal, input).Get(ctx, &output)
    return &output, err
}
