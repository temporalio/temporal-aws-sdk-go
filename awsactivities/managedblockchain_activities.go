package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/managedblockchain/managedblockchainiface"
)

type ManagedBlockchainActivities struct {
	client managedblockchainiface.ManagedBlockchainAPI
}

func NewManagedBlockchainActivities(session *session.Session, config ...*aws.Config) *ManagedBlockchainActivities {
	client := managedblockchain.New(session, config...)
	return &ManagedBlockchainActivities{client: client}
}

func (a *ManagedBlockchainActivities) CreateMember(input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error) {
	return a.client.CreateMember(input)
}

func (a *ManagedBlockchainActivities) CreateNetwork(input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error) {
	return a.client.CreateNetwork(input)
}

func (a *ManagedBlockchainActivities) CreateNode(input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error) {
	return a.client.CreateNode(input)
}

func (a *ManagedBlockchainActivities) CreateProposal(input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error) {
	return a.client.CreateProposal(input)
}

func (a *ManagedBlockchainActivities) DeleteMember(input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error) {
	return a.client.DeleteMember(input)
}

func (a *ManagedBlockchainActivities) DeleteNode(input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error) {
	return a.client.DeleteNode(input)
}

func (a *ManagedBlockchainActivities) GetMember(input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
	return a.client.GetMember(input)
}

func (a *ManagedBlockchainActivities) GetNetwork(input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
	return a.client.GetNetwork(input)
}

func (a *ManagedBlockchainActivities) GetNode(input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
	return a.client.GetNode(input)
}

func (a *ManagedBlockchainActivities) GetProposal(input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
	return a.client.GetProposal(input)
}

func (a *ManagedBlockchainActivities) ListInvitations(input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
	return a.client.ListInvitations(input)
}

func (a *ManagedBlockchainActivities) ListMembers(input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
	return a.client.ListMembers(input)
}

func (a *ManagedBlockchainActivities) ListNetworks(input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
	return a.client.ListNetworks(input)
}

func (a *ManagedBlockchainActivities) ListNodes(input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
	return a.client.ListNodes(input)
}

func (a *ManagedBlockchainActivities) ListProposalVotes(input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
	return a.client.ListProposalVotes(input)
}

func (a *ManagedBlockchainActivities) ListProposals(input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
	return a.client.ListProposals(input)
}

func (a *ManagedBlockchainActivities) RejectInvitation(input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error) {
	return a.client.RejectInvitation(input)
}

func (a *ManagedBlockchainActivities) UpdateMember(input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error) {
	return a.client.UpdateMember(input)
}

func (a *ManagedBlockchainActivities) UpdateNode(input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error) {
	return a.client.UpdateNode(input)
}

func (a *ManagedBlockchainActivities) VoteOnProposal(input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error) {
	return a.client.VoteOnProposal(input)
}
