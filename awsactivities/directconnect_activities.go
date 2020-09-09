
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directconnect/directconnectiface"
)

type DirectConnectActivities struct {
    client directconnectiface.DirectConnectAPI
}

func NewDirectConnectActivities(session *session.Session, config ...*aws.Config) *DirectConnectActivities {
    client := directconnect.New(session, config...)
    return &DirectConnectActivities{client: client}
}

func (a *DirectConnectActivities) AcceptDirectConnectGatewayAssociationProposal(input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) (*directconnect.AcceptDirectConnectGatewayAssociationProposalOutput, error) {
    return a.client.AcceptDirectConnectGatewayAssociationProposal(input)
}

func (a *DirectConnectActivities) AllocateConnectionOnInterconnect(input *directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error) {
    return a.client.AllocateConnectionOnInterconnect(input)
}

func (a *DirectConnectActivities) AllocateHostedConnection(input *directconnect.AllocateHostedConnectionInput) (*directconnect.Connection, error) {
    return a.client.AllocateHostedConnection(input)
}

func (a *DirectConnectActivities) AllocatePrivateVirtualInterface(input *directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.AllocatePrivateVirtualInterface(input)
}

func (a *DirectConnectActivities) AllocatePublicVirtualInterface(input *directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.AllocatePublicVirtualInterface(input)
}

func (a *DirectConnectActivities) AllocateTransitVirtualInterface(input *directconnect.AllocateTransitVirtualInterfaceInput) (*directconnect.AllocateTransitVirtualInterfaceOutput, error) {
    return a.client.AllocateTransitVirtualInterface(input)
}

func (a *DirectConnectActivities) AssociateConnectionWithLag(input *directconnect.AssociateConnectionWithLagInput) (*directconnect.Connection, error) {
    return a.client.AssociateConnectionWithLag(input)
}

func (a *DirectConnectActivities) AssociateHostedConnection(input *directconnect.AssociateHostedConnectionInput) (*directconnect.Connection, error) {
    return a.client.AssociateHostedConnection(input)
}

func (a *DirectConnectActivities) AssociateVirtualInterface(input *directconnect.AssociateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.AssociateVirtualInterface(input)
}

func (a *DirectConnectActivities) ConfirmConnection(input *directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error) {
    return a.client.ConfirmConnection(input)
}

func (a *DirectConnectActivities) ConfirmPrivateVirtualInterface(input *directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error) {
    return a.client.ConfirmPrivateVirtualInterface(input)
}

func (a *DirectConnectActivities) ConfirmPublicVirtualInterface(input *directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error) {
    return a.client.ConfirmPublicVirtualInterface(input)
}

func (a *DirectConnectActivities) ConfirmTransitVirtualInterface(input *directconnect.ConfirmTransitVirtualInterfaceInput) (*directconnect.ConfirmTransitVirtualInterfaceOutput, error) {
    return a.client.ConfirmTransitVirtualInterface(input)
}

func (a *DirectConnectActivities) CreateBGPPeer(input *directconnect.CreateBGPPeerInput) (*directconnect.CreateBGPPeerOutput, error) {
    return a.client.CreateBGPPeer(input)
}

func (a *DirectConnectActivities) CreateConnection(input *directconnect.CreateConnectionInput) (*directconnect.Connection, error) {
    return a.client.CreateConnection(input)
}

func (a *DirectConnectActivities) CreateDirectConnectGateway(input *directconnect.CreateDirectConnectGatewayInput) (*directconnect.CreateDirectConnectGatewayOutput, error) {
    return a.client.CreateDirectConnectGateway(input)
}

func (a *DirectConnectActivities) CreateDirectConnectGatewayAssociation(input *directconnect.CreateDirectConnectGatewayAssociationInput) (*directconnect.CreateDirectConnectGatewayAssociationOutput, error) {
    return a.client.CreateDirectConnectGatewayAssociation(input)
}

func (a *DirectConnectActivities) CreateDirectConnectGatewayAssociationProposal(input *directconnect.CreateDirectConnectGatewayAssociationProposalInput) (*directconnect.CreateDirectConnectGatewayAssociationProposalOutput, error) {
    return a.client.CreateDirectConnectGatewayAssociationProposal(input)
}

func (a *DirectConnectActivities) CreateInterconnect(input *directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error) {
    return a.client.CreateInterconnect(input)
}

func (a *DirectConnectActivities) CreateLag(input *directconnect.CreateLagInput) (*directconnect.Lag, error) {
    return a.client.CreateLag(input)
}

func (a *DirectConnectActivities) CreatePrivateVirtualInterface(input *directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.CreatePrivateVirtualInterface(input)
}

func (a *DirectConnectActivities) CreatePublicVirtualInterface(input *directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    return a.client.CreatePublicVirtualInterface(input)
}

func (a *DirectConnectActivities) CreateTransitVirtualInterface(input *directconnect.CreateTransitVirtualInterfaceInput) (*directconnect.CreateTransitVirtualInterfaceOutput, error) {
    return a.client.CreateTransitVirtualInterface(input)
}

func (a *DirectConnectActivities) DeleteBGPPeer(input *directconnect.DeleteBGPPeerInput) (*directconnect.DeleteBGPPeerOutput, error) {
    return a.client.DeleteBGPPeer(input)
}

func (a *DirectConnectActivities) DeleteConnection(input *directconnect.DeleteConnectionInput) (*directconnect.Connection, error) {
    return a.client.DeleteConnection(input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGateway(input *directconnect.DeleteDirectConnectGatewayInput) (*directconnect.DeleteDirectConnectGatewayOutput, error) {
    return a.client.DeleteDirectConnectGateway(input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGatewayAssociation(input *directconnect.DeleteDirectConnectGatewayAssociationInput) (*directconnect.DeleteDirectConnectGatewayAssociationOutput, error) {
    return a.client.DeleteDirectConnectGatewayAssociation(input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGatewayAssociationProposal(input *directconnect.DeleteDirectConnectGatewayAssociationProposalInput) (*directconnect.DeleteDirectConnectGatewayAssociationProposalOutput, error) {
    return a.client.DeleteDirectConnectGatewayAssociationProposal(input)
}

func (a *DirectConnectActivities) DeleteInterconnect(input *directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error) {
    return a.client.DeleteInterconnect(input)
}

func (a *DirectConnectActivities) DeleteLag(input *directconnect.DeleteLagInput) (*directconnect.Lag, error) {
    return a.client.DeleteLag(input)
}

func (a *DirectConnectActivities) DeleteVirtualInterface(input *directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error) {
    return a.client.DeleteVirtualInterface(input)
}

func (a *DirectConnectActivities) DescribeConnectionLoa(input *directconnect.DescribeConnectionLoaInput) (*directconnect.DescribeConnectionLoaOutput, error) {
    return a.client.DescribeConnectionLoa(input)
}

func (a *DirectConnectActivities) DescribeConnections(input *directconnect.DescribeConnectionsInput) (*directconnect.Connections, error) {
    return a.client.DescribeConnections(input)
}

func (a *DirectConnectActivities) DescribeConnectionsOnInterconnect(input *directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error) {
    return a.client.DescribeConnectionsOnInterconnect(input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAssociationProposals(input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
    return a.client.DescribeDirectConnectGatewayAssociationProposals(input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAssociations(input *directconnect.DescribeDirectConnectGatewayAssociationsInput) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
    return a.client.DescribeDirectConnectGatewayAssociations(input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAttachments(input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
    return a.client.DescribeDirectConnectGatewayAttachments(input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGateways(input *directconnect.DescribeDirectConnectGatewaysInput) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
    return a.client.DescribeDirectConnectGateways(input)
}

func (a *DirectConnectActivities) DescribeHostedConnections(input *directconnect.DescribeHostedConnectionsInput) (*directconnect.Connections, error) {
    return a.client.DescribeHostedConnections(input)
}

func (a *DirectConnectActivities) DescribeInterconnectLoa(input *directconnect.DescribeInterconnectLoaInput) (*directconnect.DescribeInterconnectLoaOutput, error) {
    return a.client.DescribeInterconnectLoa(input)
}

func (a *DirectConnectActivities) DescribeInterconnects(input *directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error) {
    return a.client.DescribeInterconnects(input)
}

func (a *DirectConnectActivities) DescribeLags(input *directconnect.DescribeLagsInput) (*directconnect.DescribeLagsOutput, error) {
    return a.client.DescribeLags(input)
}

func (a *DirectConnectActivities) DescribeLoa(input *directconnect.DescribeLoaInput) (*directconnect.Loa, error) {
    return a.client.DescribeLoa(input)
}

func (a *DirectConnectActivities) DescribeLocations(input *directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error) {
    return a.client.DescribeLocations(input)
}

func (a *DirectConnectActivities) DescribeTags(input *directconnect.DescribeTagsInput) (*directconnect.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}

func (a *DirectConnectActivities) DescribeVirtualGateways(input *directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error) {
    return a.client.DescribeVirtualGateways(input)
}

func (a *DirectConnectActivities) DescribeVirtualInterfaces(input *directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error) {
    return a.client.DescribeVirtualInterfaces(input)
}

func (a *DirectConnectActivities) DisassociateConnectionFromLag(input *directconnect.DisassociateConnectionFromLagInput) (*directconnect.Connection, error) {
    return a.client.DisassociateConnectionFromLag(input)
}

func (a *DirectConnectActivities) ListVirtualInterfaceTestHistory(input *directconnect.ListVirtualInterfaceTestHistoryInput) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
    return a.client.ListVirtualInterfaceTestHistory(input)
}

func (a *DirectConnectActivities) StartBgpFailoverTest(input *directconnect.StartBgpFailoverTestInput) (*directconnect.StartBgpFailoverTestOutput, error) {
    return a.client.StartBgpFailoverTest(input)
}

func (a *DirectConnectActivities) StopBgpFailoverTest(input *directconnect.StopBgpFailoverTestInput) (*directconnect.StopBgpFailoverTestOutput, error) {
    return a.client.StopBgpFailoverTest(input)
}

func (a *DirectConnectActivities) TagResource(input *directconnect.TagResourceInput) (*directconnect.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *DirectConnectActivities) UntagResource(input *directconnect.UntagResourceInput) (*directconnect.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *DirectConnectActivities) UpdateDirectConnectGatewayAssociation(input *directconnect.UpdateDirectConnectGatewayAssociationInput) (*directconnect.UpdateDirectConnectGatewayAssociationOutput, error) {
    return a.client.UpdateDirectConnectGatewayAssociation(input)
}

func (a *DirectConnectActivities) UpdateLag(input *directconnect.UpdateLagInput) (*directconnect.Lag, error) {
    return a.client.UpdateLag(input)
}

func (a *DirectConnectActivities) UpdateVirtualInterfaceAttributes(input *directconnect.UpdateVirtualInterfaceAttributesInput) (*directconnect.UpdateVirtualInterfaceAttributesOutput, error) {
    return a.client.UpdateVirtualInterfaceAttributes(input)
}
