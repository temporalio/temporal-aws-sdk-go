package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type EC2Activities struct {
	client ec2iface.EC2API
}

func NewEC2Activities(session *session.Session, config ...*aws.Config) *EC2Activities {
	client := ec2.New(session, config...)
	return &EC2Activities{client: client}
}

func (a *EC2Activities) AcceptReservedInstancesExchangeQuote(ctx context.Context, input *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
	return a.client.AcceptReservedInstancesExchangeQuoteWithContext(ctx, input)
}

func (a *EC2Activities) AcceptTransitGatewayPeeringAttachment(ctx context.Context, input *ec2.AcceptTransitGatewayPeeringAttachmentInput) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
	return a.client.AcceptTransitGatewayPeeringAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) AcceptTransitGatewayVpcAttachment(ctx context.Context, input *ec2.AcceptTransitGatewayVpcAttachmentInput) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
	return a.client.AcceptTransitGatewayVpcAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) AcceptVpcEndpointConnections(ctx context.Context, input *ec2.AcceptVpcEndpointConnectionsInput) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
	return a.client.AcceptVpcEndpointConnectionsWithContext(ctx, input)
}

func (a *EC2Activities) AcceptVpcPeeringConnection(ctx context.Context, input *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
	return a.client.AcceptVpcPeeringConnectionWithContext(ctx, input)
}

func (a *EC2Activities) AdvertiseByoipCidr(ctx context.Context, input *ec2.AdvertiseByoipCidrInput) (*ec2.AdvertiseByoipCidrOutput, error) {
	return a.client.AdvertiseByoipCidrWithContext(ctx, input)
}

func (a *EC2Activities) AllocateAddress(ctx context.Context, input *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
	return a.client.AllocateAddressWithContext(ctx, input)
}

func (a *EC2Activities) AllocateHosts(ctx context.Context, input *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.AllocateHostsWithContext(ctx, input)
}

func (a *EC2Activities) ApplySecurityGroupsToClientVpnTargetNetwork(ctx context.Context, input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	return a.client.ApplySecurityGroupsToClientVpnTargetNetworkWithContext(ctx, input)
}

func (a *EC2Activities) AssignIpv6Addresses(ctx context.Context, input *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error) {
	return a.client.AssignIpv6AddressesWithContext(ctx, input)
}

func (a *EC2Activities) AssignPrivateIpAddresses(ctx context.Context, input *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
	return a.client.AssignPrivateIpAddressesWithContext(ctx, input)
}

func (a *EC2Activities) AssociateAddress(ctx context.Context, input *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {
	return a.client.AssociateAddressWithContext(ctx, input)
}

func (a *EC2Activities) AssociateClientVpnTargetNetwork(ctx context.Context, input *ec2.AssociateClientVpnTargetNetworkInput) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.AssociateClientVpnTargetNetworkWithContext(ctx, input)
}

func (a *EC2Activities) AssociateDhcpOptions(ctx context.Context, input *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
	return a.client.AssociateDhcpOptionsWithContext(ctx, input)
}

func (a *EC2Activities) AssociateIamInstanceProfile(ctx context.Context, input *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error) {
	return a.client.AssociateIamInstanceProfileWithContext(ctx, input)
}

func (a *EC2Activities) AssociateRouteTable(ctx context.Context, input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
	return a.client.AssociateRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) AssociateSubnetCidrBlock(ctx context.Context, input *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error) {
	return a.client.AssociateSubnetCidrBlockWithContext(ctx, input)
}

func (a *EC2Activities) AssociateTransitGatewayMulticastDomain(ctx context.Context, input *ec2.AssociateTransitGatewayMulticastDomainInput) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
	return a.client.AssociateTransitGatewayMulticastDomainWithContext(ctx, input)
}

func (a *EC2Activities) AssociateTransitGatewayRouteTable(ctx context.Context, input *ec2.AssociateTransitGatewayRouteTableInput) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
	return a.client.AssociateTransitGatewayRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) AssociateVpcCidrBlock(ctx context.Context, input *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
	return a.client.AssociateVpcCidrBlockWithContext(ctx, input)
}

func (a *EC2Activities) AttachClassicLinkVpc(ctx context.Context, input *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
	return a.client.AttachClassicLinkVpcWithContext(ctx, input)
}

func (a *EC2Activities) AttachInternetGateway(ctx context.Context, input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
	return a.client.AttachInternetGatewayWithContext(ctx, input)
}

func (a *EC2Activities) AttachNetworkInterface(ctx context.Context, input *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
	return a.client.AttachNetworkInterfaceWithContext(ctx, input)
}

func (a *EC2Activities) AttachVolume(ctx context.Context, input *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
	return a.client.AttachVolumeWithContext(ctx, input)
}

func (a *EC2Activities) AttachVpnGateway(ctx context.Context, input *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
	return a.client.AttachVpnGatewayWithContext(ctx, input)
}

func (a *EC2Activities) AuthorizeClientVpnIngress(ctx context.Context, input *ec2.AuthorizeClientVpnIngressInput) (*ec2.AuthorizeClientVpnIngressOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.AuthorizeClientVpnIngressWithContext(ctx, input)
}

func (a *EC2Activities) AuthorizeSecurityGroupEgress(ctx context.Context, input *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	return a.client.AuthorizeSecurityGroupEgressWithContext(ctx, input)
}

func (a *EC2Activities) AuthorizeSecurityGroupIngress(ctx context.Context, input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	return a.client.AuthorizeSecurityGroupIngressWithContext(ctx, input)
}

func (a *EC2Activities) BundleInstance(ctx context.Context, input *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
	return a.client.BundleInstanceWithContext(ctx, input)
}

func (a *EC2Activities) CancelBundleTask(ctx context.Context, input *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
	return a.client.CancelBundleTaskWithContext(ctx, input)
}

func (a *EC2Activities) CancelCapacityReservation(ctx context.Context, input *ec2.CancelCapacityReservationInput) (*ec2.CancelCapacityReservationOutput, error) {
	return a.client.CancelCapacityReservationWithContext(ctx, input)
}

func (a *EC2Activities) CancelConversionTask(ctx context.Context, input *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
	return a.client.CancelConversionTaskWithContext(ctx, input)
}

func (a *EC2Activities) CancelExportTask(ctx context.Context, input *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
	return a.client.CancelExportTaskWithContext(ctx, input)
}

func (a *EC2Activities) CancelImportTask(ctx context.Context, input *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
	return a.client.CancelImportTaskWithContext(ctx, input)
}

func (a *EC2Activities) CancelReservedInstancesListing(ctx context.Context, input *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
	return a.client.CancelReservedInstancesListingWithContext(ctx, input)
}

func (a *EC2Activities) CancelSpotFleetRequests(ctx context.Context, input *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
	return a.client.CancelSpotFleetRequestsWithContext(ctx, input)
}

func (a *EC2Activities) CancelSpotInstanceRequests(ctx context.Context, input *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
	return a.client.CancelSpotInstanceRequestsWithContext(ctx, input)
}

func (a *EC2Activities) ConfirmProductInstance(ctx context.Context, input *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
	return a.client.ConfirmProductInstanceWithContext(ctx, input)
}

func (a *EC2Activities) CopyFpgaImage(ctx context.Context, input *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CopyFpgaImageWithContext(ctx, input)
}

func (a *EC2Activities) CopyImage(ctx context.Context, input *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CopyImageWithContext(ctx, input)
}

func (a *EC2Activities) CopySnapshot(ctx context.Context, input *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
	return a.client.CopySnapshotWithContext(ctx, input)
}

func (a *EC2Activities) CreateCapacityReservation(ctx context.Context, input *ec2.CreateCapacityReservationInput) (*ec2.CreateCapacityReservationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateCapacityReservationWithContext(ctx, input)
}

func (a *EC2Activities) CreateCarrierGateway(ctx context.Context, input *ec2.CreateCarrierGatewayInput) (*ec2.CreateCarrierGatewayOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateCarrierGatewayWithContext(ctx, input)
}

func (a *EC2Activities) CreateClientVpnEndpoint(ctx context.Context, input *ec2.CreateClientVpnEndpointInput) (*ec2.CreateClientVpnEndpointOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateClientVpnEndpointWithContext(ctx, input)
}

func (a *EC2Activities) CreateClientVpnRoute(ctx context.Context, input *ec2.CreateClientVpnRouteInput) (*ec2.CreateClientVpnRouteOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateClientVpnRouteWithContext(ctx, input)
}

func (a *EC2Activities) CreateCustomerGateway(ctx context.Context, input *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
	return a.client.CreateCustomerGatewayWithContext(ctx, input)
}

func (a *EC2Activities) CreateDefaultSubnet(ctx context.Context, input *ec2.CreateDefaultSubnetInput) (*ec2.CreateDefaultSubnetOutput, error) {
	return a.client.CreateDefaultSubnetWithContext(ctx, input)
}

func (a *EC2Activities) CreateDefaultVpc(ctx context.Context, input *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error) {
	return a.client.CreateDefaultVpcWithContext(ctx, input)
}

func (a *EC2Activities) CreateDhcpOptions(ctx context.Context, input *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
	return a.client.CreateDhcpOptionsWithContext(ctx, input)
}

func (a *EC2Activities) CreateEgressOnlyInternetGateway(ctx context.Context, input *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateEgressOnlyInternetGatewayWithContext(ctx, input)
}

func (a *EC2Activities) CreateFleet(ctx context.Context, input *ec2.CreateFleetInput) (*ec2.CreateFleetOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateFleetWithContext(ctx, input)
}

func (a *EC2Activities) CreateFlowLogs(ctx context.Context, input *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateFlowLogsWithContext(ctx, input)
}

func (a *EC2Activities) CreateFpgaImage(ctx context.Context, input *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateFpgaImageWithContext(ctx, input)
}

func (a *EC2Activities) CreateImage(ctx context.Context, input *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
	return a.client.CreateImageWithContext(ctx, input)
}

func (a *EC2Activities) CreateInstanceExportTask(ctx context.Context, input *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
	return a.client.CreateInstanceExportTaskWithContext(ctx, input)
}

func (a *EC2Activities) CreateInternetGateway(ctx context.Context, input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
	return a.client.CreateInternetGatewayWithContext(ctx, input)
}

func (a *EC2Activities) CreateKeyPair(ctx context.Context, input *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
	return a.client.CreateKeyPairWithContext(ctx, input)
}

func (a *EC2Activities) CreateLaunchTemplate(ctx context.Context, input *ec2.CreateLaunchTemplateInput) (*ec2.CreateLaunchTemplateOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateLaunchTemplateWithContext(ctx, input)
}

func (a *EC2Activities) CreateLaunchTemplateVersion(ctx context.Context, input *ec2.CreateLaunchTemplateVersionInput) (*ec2.CreateLaunchTemplateVersionOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateLaunchTemplateVersionWithContext(ctx, input)
}

func (a *EC2Activities) CreateLocalGatewayRoute(ctx context.Context, input *ec2.CreateLocalGatewayRouteInput) (*ec2.CreateLocalGatewayRouteOutput, error) {
	return a.client.CreateLocalGatewayRouteWithContext(ctx, input)
}

func (a *EC2Activities) CreateLocalGatewayRouteTableVpcAssociation(ctx context.Context, input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	return a.client.CreateLocalGatewayRouteTableVpcAssociationWithContext(ctx, input)
}

func (a *EC2Activities) CreateManagedPrefixList(ctx context.Context, input *ec2.CreateManagedPrefixListInput) (*ec2.CreateManagedPrefixListOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateManagedPrefixListWithContext(ctx, input)
}

func (a *EC2Activities) CreateNatGateway(ctx context.Context, input *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateNatGatewayWithContext(ctx, input)
}

func (a *EC2Activities) CreateNetworkAcl(ctx context.Context, input *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
	return a.client.CreateNetworkAclWithContext(ctx, input)
}

func (a *EC2Activities) CreateNetworkAclEntry(ctx context.Context, input *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
	return a.client.CreateNetworkAclEntryWithContext(ctx, input)
}

func (a *EC2Activities) CreateNetworkInterface(ctx context.Context, input *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
	return a.client.CreateNetworkInterfaceWithContext(ctx, input)
}

func (a *EC2Activities) CreateNetworkInterfacePermission(ctx context.Context, input *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
	return a.client.CreateNetworkInterfacePermissionWithContext(ctx, input)
}

func (a *EC2Activities) CreatePlacementGroup(ctx context.Context, input *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
	return a.client.CreatePlacementGroupWithContext(ctx, input)
}

func (a *EC2Activities) CreateReservedInstancesListing(ctx context.Context, input *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateReservedInstancesListingWithContext(ctx, input)
}

func (a *EC2Activities) CreateRoute(ctx context.Context, input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
	return a.client.CreateRouteWithContext(ctx, input)
}

func (a *EC2Activities) CreateRouteTable(ctx context.Context, input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
	return a.client.CreateRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) CreateSecurityGroup(ctx context.Context, input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	return a.client.CreateSecurityGroupWithContext(ctx, input)
}

func (a *EC2Activities) CreateSnapshot(ctx context.Context, input *ec2.CreateSnapshotInput) (*ec2.Snapshot, error) {
	return a.client.CreateSnapshotWithContext(ctx, input)
}

func (a *EC2Activities) CreateSnapshots(ctx context.Context, input *ec2.CreateSnapshotsInput) (*ec2.CreateSnapshotsOutput, error) {
	return a.client.CreateSnapshotsWithContext(ctx, input)
}

func (a *EC2Activities) CreateSpotDatafeedSubscription(ctx context.Context, input *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
	return a.client.CreateSpotDatafeedSubscriptionWithContext(ctx, input)
}

func (a *EC2Activities) CreateSubnet(ctx context.Context, input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
	return a.client.CreateSubnetWithContext(ctx, input)
}

func (a *EC2Activities) CreateTags(ctx context.Context, input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	return a.client.CreateTagsWithContext(ctx, input)
}

func (a *EC2Activities) CreateTrafficMirrorFilter(ctx context.Context, input *ec2.CreateTrafficMirrorFilterInput) (*ec2.CreateTrafficMirrorFilterOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateTrafficMirrorFilterWithContext(ctx, input)
}

func (a *EC2Activities) CreateTrafficMirrorFilterRule(ctx context.Context, input *ec2.CreateTrafficMirrorFilterRuleInput) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateTrafficMirrorFilterRuleWithContext(ctx, input)
}

func (a *EC2Activities) CreateTrafficMirrorSession(ctx context.Context, input *ec2.CreateTrafficMirrorSessionInput) (*ec2.CreateTrafficMirrorSessionOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateTrafficMirrorSessionWithContext(ctx, input)
}

func (a *EC2Activities) CreateTrafficMirrorTarget(ctx context.Context, input *ec2.CreateTrafficMirrorTargetInput) (*ec2.CreateTrafficMirrorTargetOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateTrafficMirrorTargetWithContext(ctx, input)
}

func (a *EC2Activities) CreateTransitGateway(ctx context.Context, input *ec2.CreateTransitGatewayInput) (*ec2.CreateTransitGatewayOutput, error) {
	return a.client.CreateTransitGatewayWithContext(ctx, input)
}

func (a *EC2Activities) CreateTransitGatewayMulticastDomain(ctx context.Context, input *ec2.CreateTransitGatewayMulticastDomainInput) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
	return a.client.CreateTransitGatewayMulticastDomainWithContext(ctx, input)
}

func (a *EC2Activities) CreateTransitGatewayPeeringAttachment(ctx context.Context, input *ec2.CreateTransitGatewayPeeringAttachmentInput) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
	return a.client.CreateTransitGatewayPeeringAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) CreateTransitGatewayPrefixListReference(ctx context.Context, input *ec2.CreateTransitGatewayPrefixListReferenceInput) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
	return a.client.CreateTransitGatewayPrefixListReferenceWithContext(ctx, input)
}

func (a *EC2Activities) CreateTransitGatewayRoute(ctx context.Context, input *ec2.CreateTransitGatewayRouteInput) (*ec2.CreateTransitGatewayRouteOutput, error) {
	return a.client.CreateTransitGatewayRouteWithContext(ctx, input)
}

func (a *EC2Activities) CreateTransitGatewayRouteTable(ctx context.Context, input *ec2.CreateTransitGatewayRouteTableInput) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
	return a.client.CreateTransitGatewayRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) CreateTransitGatewayVpcAttachment(ctx context.Context, input *ec2.CreateTransitGatewayVpcAttachmentInput) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
	return a.client.CreateTransitGatewayVpcAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) CreateVolume(ctx context.Context, input *ec2.CreateVolumeInput) (*ec2.Volume, error) {
	return a.client.CreateVolumeWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpc(ctx context.Context, input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
	return a.client.CreateVpcWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpcEndpoint(ctx context.Context, input *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateVpcEndpointWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpcEndpointConnectionNotification(ctx context.Context, input *ec2.CreateVpcEndpointConnectionNotificationInput) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateVpcEndpointConnectionNotificationWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpcEndpointServiceConfiguration(ctx context.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateVpcEndpointServiceConfigurationWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpcPeeringConnection(ctx context.Context, input *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
	return a.client.CreateVpcPeeringConnectionWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpnConnection(ctx context.Context, input *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
	return a.client.CreateVpnConnectionWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpnConnectionRoute(ctx context.Context, input *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
	return a.client.CreateVpnConnectionRouteWithContext(ctx, input)
}

func (a *EC2Activities) CreateVpnGateway(ctx context.Context, input *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
	return a.client.CreateVpnGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeleteCarrierGateway(ctx context.Context, input *ec2.DeleteCarrierGatewayInput) (*ec2.DeleteCarrierGatewayOutput, error) {
	return a.client.DeleteCarrierGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeleteClientVpnEndpoint(ctx context.Context, input *ec2.DeleteClientVpnEndpointInput) (*ec2.DeleteClientVpnEndpointOutput, error) {
	return a.client.DeleteClientVpnEndpointWithContext(ctx, input)
}

func (a *EC2Activities) DeleteClientVpnRoute(ctx context.Context, input *ec2.DeleteClientVpnRouteInput) (*ec2.DeleteClientVpnRouteOutput, error) {
	return a.client.DeleteClientVpnRouteWithContext(ctx, input)
}

func (a *EC2Activities) DeleteCustomerGateway(ctx context.Context, input *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
	return a.client.DeleteCustomerGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeleteDhcpOptions(ctx context.Context, input *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
	return a.client.DeleteDhcpOptionsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteEgressOnlyInternetGateway(ctx context.Context, input *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
	return a.client.DeleteEgressOnlyInternetGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeleteFleets(ctx context.Context, input *ec2.DeleteFleetsInput) (*ec2.DeleteFleetsOutput, error) {
	return a.client.DeleteFleetsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteFlowLogs(ctx context.Context, input *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
	return a.client.DeleteFlowLogsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteFpgaImage(ctx context.Context, input *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error) {
	return a.client.DeleteFpgaImageWithContext(ctx, input)
}

func (a *EC2Activities) DeleteInternetGateway(ctx context.Context, input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	return a.client.DeleteInternetGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeleteKeyPair(ctx context.Context, input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	return a.client.DeleteKeyPairWithContext(ctx, input)
}

func (a *EC2Activities) DeleteLaunchTemplate(ctx context.Context, input *ec2.DeleteLaunchTemplateInput) (*ec2.DeleteLaunchTemplateOutput, error) {
	return a.client.DeleteLaunchTemplateWithContext(ctx, input)
}

func (a *EC2Activities) DeleteLaunchTemplateVersions(ctx context.Context, input *ec2.DeleteLaunchTemplateVersionsInput) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
	return a.client.DeleteLaunchTemplateVersionsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteLocalGatewayRoute(ctx context.Context, input *ec2.DeleteLocalGatewayRouteInput) (*ec2.DeleteLocalGatewayRouteOutput, error) {
	return a.client.DeleteLocalGatewayRouteWithContext(ctx, input)
}

func (a *EC2Activities) DeleteLocalGatewayRouteTableVpcAssociation(ctx context.Context, input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
	return a.client.DeleteLocalGatewayRouteTableVpcAssociationWithContext(ctx, input)
}

func (a *EC2Activities) DeleteManagedPrefixList(ctx context.Context, input *ec2.DeleteManagedPrefixListInput) (*ec2.DeleteManagedPrefixListOutput, error) {
	return a.client.DeleteManagedPrefixListWithContext(ctx, input)
}

func (a *EC2Activities) DeleteNatGateway(ctx context.Context, input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	return a.client.DeleteNatGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeleteNetworkAcl(ctx context.Context, input *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
	return a.client.DeleteNetworkAclWithContext(ctx, input)
}

func (a *EC2Activities) DeleteNetworkAclEntry(ctx context.Context, input *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
	return a.client.DeleteNetworkAclEntryWithContext(ctx, input)
}

func (a *EC2Activities) DeleteNetworkInterface(ctx context.Context, input *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
	return a.client.DeleteNetworkInterfaceWithContext(ctx, input)
}

func (a *EC2Activities) DeleteNetworkInterfacePermission(ctx context.Context, input *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
	return a.client.DeleteNetworkInterfacePermissionWithContext(ctx, input)
}

func (a *EC2Activities) DeletePlacementGroup(ctx context.Context, input *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
	return a.client.DeletePlacementGroupWithContext(ctx, input)
}

func (a *EC2Activities) DeleteQueuedReservedInstances(ctx context.Context, input *ec2.DeleteQueuedReservedInstancesInput) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
	return a.client.DeleteQueuedReservedInstancesWithContext(ctx, input)
}

func (a *EC2Activities) DeleteRoute(ctx context.Context, input *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
	return a.client.DeleteRouteWithContext(ctx, input)
}

func (a *EC2Activities) DeleteRouteTable(ctx context.Context, input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	return a.client.DeleteRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) DeleteSecurityGroup(ctx context.Context, input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	return a.client.DeleteSecurityGroupWithContext(ctx, input)
}

func (a *EC2Activities) DeleteSnapshot(ctx context.Context, input *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	return a.client.DeleteSnapshotWithContext(ctx, input)
}

func (a *EC2Activities) DeleteSpotDatafeedSubscription(ctx context.Context, input *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
	return a.client.DeleteSpotDatafeedSubscriptionWithContext(ctx, input)
}

func (a *EC2Activities) DeleteSubnet(ctx context.Context, input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	return a.client.DeleteSubnetWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTags(ctx context.Context, input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTrafficMirrorFilter(ctx context.Context, input *ec2.DeleteTrafficMirrorFilterInput) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
	return a.client.DeleteTrafficMirrorFilterWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTrafficMirrorFilterRule(ctx context.Context, input *ec2.DeleteTrafficMirrorFilterRuleInput) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
	return a.client.DeleteTrafficMirrorFilterRuleWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTrafficMirrorSession(ctx context.Context, input *ec2.DeleteTrafficMirrorSessionInput) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
	return a.client.DeleteTrafficMirrorSessionWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTrafficMirrorTarget(ctx context.Context, input *ec2.DeleteTrafficMirrorTargetInput) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
	return a.client.DeleteTrafficMirrorTargetWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTransitGateway(ctx context.Context, input *ec2.DeleteTransitGatewayInput) (*ec2.DeleteTransitGatewayOutput, error) {
	return a.client.DeleteTransitGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTransitGatewayMulticastDomain(ctx context.Context, input *ec2.DeleteTransitGatewayMulticastDomainInput) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
	return a.client.DeleteTransitGatewayMulticastDomainWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTransitGatewayPeeringAttachment(ctx context.Context, input *ec2.DeleteTransitGatewayPeeringAttachmentInput) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
	return a.client.DeleteTransitGatewayPeeringAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTransitGatewayPrefixListReference(ctx context.Context, input *ec2.DeleteTransitGatewayPrefixListReferenceInput) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
	return a.client.DeleteTransitGatewayPrefixListReferenceWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTransitGatewayRoute(ctx context.Context, input *ec2.DeleteTransitGatewayRouteInput) (*ec2.DeleteTransitGatewayRouteOutput, error) {
	return a.client.DeleteTransitGatewayRouteWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTransitGatewayRouteTable(ctx context.Context, input *ec2.DeleteTransitGatewayRouteTableInput) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
	return a.client.DeleteTransitGatewayRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) DeleteTransitGatewayVpcAttachment(ctx context.Context, input *ec2.DeleteTransitGatewayVpcAttachmentInput) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
	return a.client.DeleteTransitGatewayVpcAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVolume(ctx context.Context, input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	return a.client.DeleteVolumeWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpc(ctx context.Context, input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	return a.client.DeleteVpcWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpcEndpointConnectionNotifications(ctx context.Context, input *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
	return a.client.DeleteVpcEndpointConnectionNotificationsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpcEndpointServiceConfigurations(ctx context.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
	return a.client.DeleteVpcEndpointServiceConfigurationsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpcEndpoints(ctx context.Context, input *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
	return a.client.DeleteVpcEndpointsWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpcPeeringConnection(ctx context.Context, input *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
	return a.client.DeleteVpcPeeringConnectionWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpnConnection(ctx context.Context, input *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
	return a.client.DeleteVpnConnectionWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpnConnectionRoute(ctx context.Context, input *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
	return a.client.DeleteVpnConnectionRouteWithContext(ctx, input)
}

func (a *EC2Activities) DeleteVpnGateway(ctx context.Context, input *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
	return a.client.DeleteVpnGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DeprovisionByoipCidr(ctx context.Context, input *ec2.DeprovisionByoipCidrInput) (*ec2.DeprovisionByoipCidrOutput, error) {
	return a.client.DeprovisionByoipCidrWithContext(ctx, input)
}

func (a *EC2Activities) DeregisterImage(ctx context.Context, input *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
	return a.client.DeregisterImageWithContext(ctx, input)
}

func (a *EC2Activities) DeregisterInstanceEventNotificationAttributes(ctx context.Context, input *ec2.DeregisterInstanceEventNotificationAttributesInput) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
	return a.client.DeregisterInstanceEventNotificationAttributesWithContext(ctx, input)
}

func (a *EC2Activities) DeregisterTransitGatewayMulticastGroupMembers(ctx context.Context, input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
	return a.client.DeregisterTransitGatewayMulticastGroupMembersWithContext(ctx, input)
}

func (a *EC2Activities) DeregisterTransitGatewayMulticastGroupSources(ctx context.Context, input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
	return a.client.DeregisterTransitGatewayMulticastGroupSourcesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeAccountAttributes(ctx context.Context, input *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
	return a.client.DescribeAccountAttributesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeAddresses(ctx context.Context, input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	return a.client.DescribeAddressesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeAggregateIdFormat(ctx context.Context, input *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error) {
	return a.client.DescribeAggregateIdFormatWithContext(ctx, input)
}

func (a *EC2Activities) DescribeAvailabilityZones(ctx context.Context, input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	return a.client.DescribeAvailabilityZonesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeBundleTasks(ctx context.Context, input *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
	return a.client.DescribeBundleTasksWithContext(ctx, input)
}

func (a *EC2Activities) DescribeByoipCidrs(ctx context.Context, input *ec2.DescribeByoipCidrsInput) (*ec2.DescribeByoipCidrsOutput, error) {
	return a.client.DescribeByoipCidrsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeCapacityReservations(ctx context.Context, input *ec2.DescribeCapacityReservationsInput) (*ec2.DescribeCapacityReservationsOutput, error) {
	return a.client.DescribeCapacityReservationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeCarrierGateways(ctx context.Context, input *ec2.DescribeCarrierGatewaysInput) (*ec2.DescribeCarrierGatewaysOutput, error) {
	return a.client.DescribeCarrierGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DescribeClassicLinkInstances(ctx context.Context, input *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	return a.client.DescribeClassicLinkInstancesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeClientVpnAuthorizationRules(ctx context.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	return a.client.DescribeClientVpnAuthorizationRulesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeClientVpnConnections(ctx context.Context, input *ec2.DescribeClientVpnConnectionsInput) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	return a.client.DescribeClientVpnConnectionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeClientVpnEndpoints(ctx context.Context, input *ec2.DescribeClientVpnEndpointsInput) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	return a.client.DescribeClientVpnEndpointsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeClientVpnRoutes(ctx context.Context, input *ec2.DescribeClientVpnRoutesInput) (*ec2.DescribeClientVpnRoutesOutput, error) {
	return a.client.DescribeClientVpnRoutesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeClientVpnTargetNetworks(ctx context.Context, input *ec2.DescribeClientVpnTargetNetworksInput) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	return a.client.DescribeClientVpnTargetNetworksWithContext(ctx, input)
}

func (a *EC2Activities) DescribeCoipPools(ctx context.Context, input *ec2.DescribeCoipPoolsInput) (*ec2.DescribeCoipPoolsOutput, error) {
	return a.client.DescribeCoipPoolsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeConversionTasks(ctx context.Context, input *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
	return a.client.DescribeConversionTasksWithContext(ctx, input)
}

func (a *EC2Activities) DescribeCustomerGateways(ctx context.Context, input *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
	return a.client.DescribeCustomerGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DescribeDhcpOptions(ctx context.Context, input *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
	return a.client.DescribeDhcpOptionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeEgressOnlyInternetGateways(ctx context.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	return a.client.DescribeEgressOnlyInternetGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DescribeElasticGpus(ctx context.Context, input *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
	return a.client.DescribeElasticGpusWithContext(ctx, input)
}

func (a *EC2Activities) DescribeExportImageTasks(ctx context.Context, input *ec2.DescribeExportImageTasksInput) (*ec2.DescribeExportImageTasksOutput, error) {
	return a.client.DescribeExportImageTasksWithContext(ctx, input)
}

func (a *EC2Activities) DescribeExportTasks(ctx context.Context, input *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
	return a.client.DescribeExportTasksWithContext(ctx, input)
}

func (a *EC2Activities) DescribeFastSnapshotRestores(ctx context.Context, input *ec2.DescribeFastSnapshotRestoresInput) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	return a.client.DescribeFastSnapshotRestoresWithContext(ctx, input)
}

func (a *EC2Activities) DescribeFleetHistory(ctx context.Context, input *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error) {
	return a.client.DescribeFleetHistoryWithContext(ctx, input)
}

func (a *EC2Activities) DescribeFleetInstances(ctx context.Context, input *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error) {
	return a.client.DescribeFleetInstancesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeFleets(ctx context.Context, input *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error) {
	return a.client.DescribeFleetsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeFlowLogs(ctx context.Context, input *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
	return a.client.DescribeFlowLogsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeFpgaImageAttribute(ctx context.Context, input *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	return a.client.DescribeFpgaImageAttributeWithContext(ctx, input)
}

func (a *EC2Activities) DescribeFpgaImages(ctx context.Context, input *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
	return a.client.DescribeFpgaImagesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeHostReservationOfferings(ctx context.Context, input *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	return a.client.DescribeHostReservationOfferingsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeHostReservations(ctx context.Context, input *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
	return a.client.DescribeHostReservationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeHosts(ctx context.Context, input *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
	return a.client.DescribeHostsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeIamInstanceProfileAssociations(ctx context.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	return a.client.DescribeIamInstanceProfileAssociationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeIdFormat(ctx context.Context, input *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
	return a.client.DescribeIdFormatWithContext(ctx, input)
}

func (a *EC2Activities) DescribeIdentityIdFormat(ctx context.Context, input *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
	return a.client.DescribeIdentityIdFormatWithContext(ctx, input)
}

func (a *EC2Activities) DescribeImageAttribute(ctx context.Context, input *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
	return a.client.DescribeImageAttributeWithContext(ctx, input)
}

func (a *EC2Activities) DescribeImages(ctx context.Context, input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	return a.client.DescribeImagesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeImportImageTasks(ctx context.Context, input *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
	return a.client.DescribeImportImageTasksWithContext(ctx, input)
}

func (a *EC2Activities) DescribeImportSnapshotTasks(ctx context.Context, input *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	return a.client.DescribeImportSnapshotTasksWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInstanceAttribute(ctx context.Context, input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
	return a.client.DescribeInstanceAttributeWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInstanceCreditSpecifications(ctx context.Context, input *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	return a.client.DescribeInstanceCreditSpecificationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInstanceEventNotificationAttributes(ctx context.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	return a.client.DescribeInstanceEventNotificationAttributesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInstanceStatus(ctx context.Context, input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	return a.client.DescribeInstanceStatusWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInstanceTypeOfferings(ctx context.Context, input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	return a.client.DescribeInstanceTypeOfferingsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInstanceTypes(ctx context.Context, input *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
	return a.client.DescribeInstanceTypesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInstances(ctx context.Context, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return a.client.DescribeInstancesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeInternetGateways(ctx context.Context, input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	return a.client.DescribeInternetGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DescribeIpv6Pools(ctx context.Context, input *ec2.DescribeIpv6PoolsInput) (*ec2.DescribeIpv6PoolsOutput, error) {
	return a.client.DescribeIpv6PoolsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeKeyPairs(ctx context.Context, input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	return a.client.DescribeKeyPairsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLaunchTemplateVersions(ctx context.Context, input *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	return a.client.DescribeLaunchTemplateVersionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLaunchTemplates(ctx context.Context, input *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error) {
	return a.client.DescribeLaunchTemplatesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	return a.client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLocalGatewayRouteTableVpcAssociations(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	return a.client.DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLocalGatewayRouteTables(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTablesInput) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	return a.client.DescribeLocalGatewayRouteTablesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLocalGatewayVirtualInterfaceGroups(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	return a.client.DescribeLocalGatewayVirtualInterfaceGroupsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLocalGatewayVirtualInterfaces(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	return a.client.DescribeLocalGatewayVirtualInterfacesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeLocalGateways(ctx context.Context, input *ec2.DescribeLocalGatewaysInput) (*ec2.DescribeLocalGatewaysOutput, error) {
	return a.client.DescribeLocalGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DescribeManagedPrefixLists(ctx context.Context, input *ec2.DescribeManagedPrefixListsInput) (*ec2.DescribeManagedPrefixListsOutput, error) {
	return a.client.DescribeManagedPrefixListsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeMovingAddresses(ctx context.Context, input *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
	return a.client.DescribeMovingAddressesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeNatGateways(ctx context.Context, input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	return a.client.DescribeNatGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DescribeNetworkAcls(ctx context.Context, input *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
	return a.client.DescribeNetworkAclsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeNetworkInterfaceAttribute(ctx context.Context, input *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	return a.client.DescribeNetworkInterfaceAttributeWithContext(ctx, input)
}

func (a *EC2Activities) DescribeNetworkInterfacePermissions(ctx context.Context, input *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	return a.client.DescribeNetworkInterfacePermissionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeNetworkInterfaces(ctx context.Context, input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	return a.client.DescribeNetworkInterfacesWithContext(ctx, input)
}

func (a *EC2Activities) DescribePlacementGroups(ctx context.Context, input *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
	return a.client.DescribePlacementGroupsWithContext(ctx, input)
}

func (a *EC2Activities) DescribePrefixLists(ctx context.Context, input *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
	return a.client.DescribePrefixListsWithContext(ctx, input)
}

func (a *EC2Activities) DescribePrincipalIdFormat(ctx context.Context, input *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error) {
	return a.client.DescribePrincipalIdFormatWithContext(ctx, input)
}

func (a *EC2Activities) DescribePublicIpv4Pools(ctx context.Context, input *ec2.DescribePublicIpv4PoolsInput) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	return a.client.DescribePublicIpv4PoolsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeRegions(ctx context.Context, input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	return a.client.DescribeRegionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeReservedInstances(ctx context.Context, input *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
	return a.client.DescribeReservedInstancesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeReservedInstancesListings(ctx context.Context, input *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	return a.client.DescribeReservedInstancesListingsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeReservedInstancesModifications(ctx context.Context, input *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	return a.client.DescribeReservedInstancesModificationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeReservedInstancesOfferings(ctx context.Context, input *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	return a.client.DescribeReservedInstancesOfferingsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeRouteTables(ctx context.Context, input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	return a.client.DescribeRouteTablesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeScheduledInstanceAvailability(ctx context.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	return a.client.DescribeScheduledInstanceAvailabilityWithContext(ctx, input)
}

func (a *EC2Activities) DescribeScheduledInstances(ctx context.Context, input *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
	return a.client.DescribeScheduledInstancesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSecurityGroupReferences(ctx context.Context, input *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	return a.client.DescribeSecurityGroupReferencesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSecurityGroups(ctx context.Context, input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	return a.client.DescribeSecurityGroupsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSnapshotAttribute(ctx context.Context, input *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
	return a.client.DescribeSnapshotAttributeWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSnapshots(ctx context.Context, input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	return a.client.DescribeSnapshotsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSpotDatafeedSubscription(ctx context.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	return a.client.DescribeSpotDatafeedSubscriptionWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSpotFleetInstances(ctx context.Context, input *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	return a.client.DescribeSpotFleetInstancesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSpotFleetRequestHistory(ctx context.Context, input *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	return a.client.DescribeSpotFleetRequestHistoryWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSpotFleetRequests(ctx context.Context, input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	return a.client.DescribeSpotFleetRequestsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSpotInstanceRequests(ctx context.Context, input *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	return a.client.DescribeSpotInstanceRequestsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSpotPriceHistory(ctx context.Context, input *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	return a.client.DescribeSpotPriceHistoryWithContext(ctx, input)
}

func (a *EC2Activities) DescribeStaleSecurityGroups(ctx context.Context, input *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	return a.client.DescribeStaleSecurityGroupsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeSubnets(ctx context.Context, input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	return a.client.DescribeSubnetsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTags(ctx context.Context, input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTrafficMirrorFilters(ctx context.Context, input *ec2.DescribeTrafficMirrorFiltersInput) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	return a.client.DescribeTrafficMirrorFiltersWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTrafficMirrorSessions(ctx context.Context, input *ec2.DescribeTrafficMirrorSessionsInput) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	return a.client.DescribeTrafficMirrorSessionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTrafficMirrorTargets(ctx context.Context, input *ec2.DescribeTrafficMirrorTargetsInput) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	return a.client.DescribeTrafficMirrorTargetsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTransitGatewayAttachments(ctx context.Context, input *ec2.DescribeTransitGatewayAttachmentsInput) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	return a.client.DescribeTransitGatewayAttachmentsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTransitGatewayMulticastDomains(ctx context.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	return a.client.DescribeTransitGatewayMulticastDomainsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTransitGatewayPeeringAttachments(ctx context.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	return a.client.DescribeTransitGatewayPeeringAttachmentsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTransitGatewayRouteTables(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTablesInput) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	return a.client.DescribeTransitGatewayRouteTablesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTransitGatewayVpcAttachments(ctx context.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	return a.client.DescribeTransitGatewayVpcAttachmentsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeTransitGateways(ctx context.Context, input *ec2.DescribeTransitGatewaysInput) (*ec2.DescribeTransitGatewaysOutput, error) {
	return a.client.DescribeTransitGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVolumeAttribute(ctx context.Context, input *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
	return a.client.DescribeVolumeAttributeWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVolumeStatus(ctx context.Context, input *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
	return a.client.DescribeVolumeStatusWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVolumes(ctx context.Context, input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	return a.client.DescribeVolumesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVolumesModifications(ctx context.Context, input *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
	return a.client.DescribeVolumesModificationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcAttribute(ctx context.Context, input *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
	return a.client.DescribeVpcAttributeWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcClassicLink(ctx context.Context, input *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
	return a.client.DescribeVpcClassicLinkWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcClassicLinkDnsSupport(ctx context.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	return a.client.DescribeVpcClassicLinkDnsSupportWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcEndpointConnectionNotifications(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	return a.client.DescribeVpcEndpointConnectionNotificationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcEndpointConnections(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	return a.client.DescribeVpcEndpointConnectionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcEndpointServiceConfigurations(ctx context.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	return a.client.DescribeVpcEndpointServiceConfigurationsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcEndpointServicePermissions(ctx context.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	return a.client.DescribeVpcEndpointServicePermissionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcEndpointServices(ctx context.Context, input *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	return a.client.DescribeVpcEndpointServicesWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcEndpoints(ctx context.Context, input *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	return a.client.DescribeVpcEndpointsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcPeeringConnections(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	return a.client.DescribeVpcPeeringConnectionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpcs(ctx context.Context, input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	return a.client.DescribeVpcsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpnConnections(ctx context.Context, input *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
	return a.client.DescribeVpnConnectionsWithContext(ctx, input)
}

func (a *EC2Activities) DescribeVpnGateways(ctx context.Context, input *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
	return a.client.DescribeVpnGatewaysWithContext(ctx, input)
}

func (a *EC2Activities) DetachClassicLinkVpc(ctx context.Context, input *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
	return a.client.DetachClassicLinkVpcWithContext(ctx, input)
}

func (a *EC2Activities) DetachInternetGateway(ctx context.Context, input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	return a.client.DetachInternetGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DetachNetworkInterface(ctx context.Context, input *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
	return a.client.DetachNetworkInterfaceWithContext(ctx, input)
}

func (a *EC2Activities) DetachVolume(ctx context.Context, input *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
	return a.client.DetachVolumeWithContext(ctx, input)
}

func (a *EC2Activities) DetachVpnGateway(ctx context.Context, input *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
	return a.client.DetachVpnGatewayWithContext(ctx, input)
}

func (a *EC2Activities) DisableEbsEncryptionByDefault(ctx context.Context, input *ec2.DisableEbsEncryptionByDefaultInput) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
	return a.client.DisableEbsEncryptionByDefaultWithContext(ctx, input)
}

func (a *EC2Activities) DisableFastSnapshotRestores(ctx context.Context, input *ec2.DisableFastSnapshotRestoresInput) (*ec2.DisableFastSnapshotRestoresOutput, error) {
	return a.client.DisableFastSnapshotRestoresWithContext(ctx, input)
}

func (a *EC2Activities) DisableTransitGatewayRouteTablePropagation(ctx context.Context, input *ec2.DisableTransitGatewayRouteTablePropagationInput) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
	return a.client.DisableTransitGatewayRouteTablePropagationWithContext(ctx, input)
}

func (a *EC2Activities) DisableVgwRoutePropagation(ctx context.Context, input *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
	return a.client.DisableVgwRoutePropagationWithContext(ctx, input)
}

func (a *EC2Activities) DisableVpcClassicLink(ctx context.Context, input *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
	return a.client.DisableVpcClassicLinkWithContext(ctx, input)
}

func (a *EC2Activities) DisableVpcClassicLinkDnsSupport(ctx context.Context, input *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
	return a.client.DisableVpcClassicLinkDnsSupportWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateAddress(ctx context.Context, input *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {
	return a.client.DisassociateAddressWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateClientVpnTargetNetwork(ctx context.Context, input *ec2.DisassociateClientVpnTargetNetworkInput) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
	return a.client.DisassociateClientVpnTargetNetworkWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateIamInstanceProfile(ctx context.Context, input *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error) {
	return a.client.DisassociateIamInstanceProfileWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateRouteTable(ctx context.Context, input *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
	return a.client.DisassociateRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateSubnetCidrBlock(ctx context.Context, input *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
	return a.client.DisassociateSubnetCidrBlockWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateTransitGatewayMulticastDomain(ctx context.Context, input *ec2.DisassociateTransitGatewayMulticastDomainInput) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
	return a.client.DisassociateTransitGatewayMulticastDomainWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateTransitGatewayRouteTable(ctx context.Context, input *ec2.DisassociateTransitGatewayRouteTableInput) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
	return a.client.DisassociateTransitGatewayRouteTableWithContext(ctx, input)
}

func (a *EC2Activities) DisassociateVpcCidrBlock(ctx context.Context, input *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error) {
	return a.client.DisassociateVpcCidrBlockWithContext(ctx, input)
}

func (a *EC2Activities) EnableEbsEncryptionByDefault(ctx context.Context, input *ec2.EnableEbsEncryptionByDefaultInput) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
	return a.client.EnableEbsEncryptionByDefaultWithContext(ctx, input)
}

func (a *EC2Activities) EnableFastSnapshotRestores(ctx context.Context, input *ec2.EnableFastSnapshotRestoresInput) (*ec2.EnableFastSnapshotRestoresOutput, error) {
	return a.client.EnableFastSnapshotRestoresWithContext(ctx, input)
}

func (a *EC2Activities) EnableTransitGatewayRouteTablePropagation(ctx context.Context, input *ec2.EnableTransitGatewayRouteTablePropagationInput) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
	return a.client.EnableTransitGatewayRouteTablePropagationWithContext(ctx, input)
}

func (a *EC2Activities) EnableVgwRoutePropagation(ctx context.Context, input *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
	return a.client.EnableVgwRoutePropagationWithContext(ctx, input)
}

func (a *EC2Activities) EnableVolumeIO(ctx context.Context, input *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
	return a.client.EnableVolumeIOWithContext(ctx, input)
}

func (a *EC2Activities) EnableVpcClassicLink(ctx context.Context, input *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
	return a.client.EnableVpcClassicLinkWithContext(ctx, input)
}

func (a *EC2Activities) EnableVpcClassicLinkDnsSupport(ctx context.Context, input *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
	return a.client.EnableVpcClassicLinkDnsSupportWithContext(ctx, input)
}

func (a *EC2Activities) ExportClientVpnClientCertificateRevocationList(ctx context.Context, input *ec2.ExportClientVpnClientCertificateRevocationListInput) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
	return a.client.ExportClientVpnClientCertificateRevocationListWithContext(ctx, input)
}

func (a *EC2Activities) ExportClientVpnClientConfiguration(ctx context.Context, input *ec2.ExportClientVpnClientConfigurationInput) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
	return a.client.ExportClientVpnClientConfigurationWithContext(ctx, input)
}

func (a *EC2Activities) ExportImage(ctx context.Context, input *ec2.ExportImageInput) (*ec2.ExportImageOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.ExportImageWithContext(ctx, input)
}

func (a *EC2Activities) ExportTransitGatewayRoutes(ctx context.Context, input *ec2.ExportTransitGatewayRoutesInput) (*ec2.ExportTransitGatewayRoutesOutput, error) {
	return a.client.ExportTransitGatewayRoutesWithContext(ctx, input)
}

func (a *EC2Activities) GetAssociatedIpv6PoolCidrs(ctx context.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	return a.client.GetAssociatedIpv6PoolCidrsWithContext(ctx, input)
}

func (a *EC2Activities) GetCapacityReservationUsage(ctx context.Context, input *ec2.GetCapacityReservationUsageInput) (*ec2.GetCapacityReservationUsageOutput, error) {
	return a.client.GetCapacityReservationUsageWithContext(ctx, input)
}

func (a *EC2Activities) GetCoipPoolUsage(ctx context.Context, input *ec2.GetCoipPoolUsageInput) (*ec2.GetCoipPoolUsageOutput, error) {
	return a.client.GetCoipPoolUsageWithContext(ctx, input)
}

func (a *EC2Activities) GetConsoleOutput(ctx context.Context, input *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
	return a.client.GetConsoleOutputWithContext(ctx, input)
}

func (a *EC2Activities) GetConsoleScreenshot(ctx context.Context, input *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
	return a.client.GetConsoleScreenshotWithContext(ctx, input)
}

func (a *EC2Activities) GetDefaultCreditSpecification(ctx context.Context, input *ec2.GetDefaultCreditSpecificationInput) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	return a.client.GetDefaultCreditSpecificationWithContext(ctx, input)
}

func (a *EC2Activities) GetEbsDefaultKmsKeyId(ctx context.Context, input *ec2.GetEbsDefaultKmsKeyIdInput) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	return a.client.GetEbsDefaultKmsKeyIdWithContext(ctx, input)
}

func (a *EC2Activities) GetEbsEncryptionByDefault(ctx context.Context, input *ec2.GetEbsEncryptionByDefaultInput) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	return a.client.GetEbsEncryptionByDefaultWithContext(ctx, input)
}

func (a *EC2Activities) GetGroupsForCapacityReservation(ctx context.Context, input *ec2.GetGroupsForCapacityReservationInput) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	return a.client.GetGroupsForCapacityReservationWithContext(ctx, input)
}

func (a *EC2Activities) GetHostReservationPurchasePreview(ctx context.Context, input *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	return a.client.GetHostReservationPurchasePreviewWithContext(ctx, input)
}

func (a *EC2Activities) GetLaunchTemplateData(ctx context.Context, input *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error) {
	return a.client.GetLaunchTemplateDataWithContext(ctx, input)
}

func (a *EC2Activities) GetManagedPrefixListAssociations(ctx context.Context, input *ec2.GetManagedPrefixListAssociationsInput) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	return a.client.GetManagedPrefixListAssociationsWithContext(ctx, input)
}

func (a *EC2Activities) GetManagedPrefixListEntries(ctx context.Context, input *ec2.GetManagedPrefixListEntriesInput) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	return a.client.GetManagedPrefixListEntriesWithContext(ctx, input)
}

func (a *EC2Activities) GetPasswordData(ctx context.Context, input *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
	return a.client.GetPasswordDataWithContext(ctx, input)
}

func (a *EC2Activities) GetReservedInstancesExchangeQuote(ctx context.Context, input *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	return a.client.GetReservedInstancesExchangeQuoteWithContext(ctx, input)
}

func (a *EC2Activities) GetTransitGatewayAttachmentPropagations(ctx context.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	return a.client.GetTransitGatewayAttachmentPropagationsWithContext(ctx, input)
}

func (a *EC2Activities) GetTransitGatewayMulticastDomainAssociations(ctx context.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	return a.client.GetTransitGatewayMulticastDomainAssociationsWithContext(ctx, input)
}

func (a *EC2Activities) GetTransitGatewayPrefixListReferences(ctx context.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	return a.client.GetTransitGatewayPrefixListReferencesWithContext(ctx, input)
}

func (a *EC2Activities) GetTransitGatewayRouteTableAssociations(ctx context.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	return a.client.GetTransitGatewayRouteTableAssociationsWithContext(ctx, input)
}

func (a *EC2Activities) GetTransitGatewayRouteTablePropagations(ctx context.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	return a.client.GetTransitGatewayRouteTablePropagationsWithContext(ctx, input)
}

func (a *EC2Activities) ImportClientVpnClientCertificateRevocationList(ctx context.Context, input *ec2.ImportClientVpnClientCertificateRevocationListInput) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
	return a.client.ImportClientVpnClientCertificateRevocationListWithContext(ctx, input)
}

func (a *EC2Activities) ImportImage(ctx context.Context, input *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.ImportImageWithContext(ctx, input)
}

func (a *EC2Activities) ImportInstance(ctx context.Context, input *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
	return a.client.ImportInstanceWithContext(ctx, input)
}

func (a *EC2Activities) ImportKeyPair(ctx context.Context, input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	return a.client.ImportKeyPairWithContext(ctx, input)
}

func (a *EC2Activities) ImportSnapshot(ctx context.Context, input *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.ImportSnapshotWithContext(ctx, input)
}

func (a *EC2Activities) ImportVolume(ctx context.Context, input *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
	return a.client.ImportVolumeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyAvailabilityZoneGroup(ctx context.Context, input *ec2.ModifyAvailabilityZoneGroupInput) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
	return a.client.ModifyAvailabilityZoneGroupWithContext(ctx, input)
}

func (a *EC2Activities) ModifyCapacityReservation(ctx context.Context, input *ec2.ModifyCapacityReservationInput) (*ec2.ModifyCapacityReservationOutput, error) {
	return a.client.ModifyCapacityReservationWithContext(ctx, input)
}

func (a *EC2Activities) ModifyClientVpnEndpoint(ctx context.Context, input *ec2.ModifyClientVpnEndpointInput) (*ec2.ModifyClientVpnEndpointOutput, error) {
	return a.client.ModifyClientVpnEndpointWithContext(ctx, input)
}

func (a *EC2Activities) ModifyDefaultCreditSpecification(ctx context.Context, input *ec2.ModifyDefaultCreditSpecificationInput) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
	return a.client.ModifyDefaultCreditSpecificationWithContext(ctx, input)
}

func (a *EC2Activities) ModifyEbsDefaultKmsKeyId(ctx context.Context, input *ec2.ModifyEbsDefaultKmsKeyIdInput) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
	return a.client.ModifyEbsDefaultKmsKeyIdWithContext(ctx, input)
}

func (a *EC2Activities) ModifyFleet(ctx context.Context, input *ec2.ModifyFleetInput) (*ec2.ModifyFleetOutput, error) {
	return a.client.ModifyFleetWithContext(ctx, input)
}

func (a *EC2Activities) ModifyFpgaImageAttribute(ctx context.Context, input *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error) {
	return a.client.ModifyFpgaImageAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyHosts(ctx context.Context, input *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
	return a.client.ModifyHostsWithContext(ctx, input)
}

func (a *EC2Activities) ModifyIdFormat(ctx context.Context, input *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
	return a.client.ModifyIdFormatWithContext(ctx, input)
}

func (a *EC2Activities) ModifyIdentityIdFormat(ctx context.Context, input *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {
	return a.client.ModifyIdentityIdFormatWithContext(ctx, input)
}

func (a *EC2Activities) ModifyImageAttribute(ctx context.Context, input *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
	return a.client.ModifyImageAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyInstanceAttribute(ctx context.Context, input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
	return a.client.ModifyInstanceAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyInstanceCapacityReservationAttributes(ctx context.Context, input *ec2.ModifyInstanceCapacityReservationAttributesInput) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
	return a.client.ModifyInstanceCapacityReservationAttributesWithContext(ctx, input)
}

func (a *EC2Activities) ModifyInstanceCreditSpecification(ctx context.Context, input *ec2.ModifyInstanceCreditSpecificationInput) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.ModifyInstanceCreditSpecificationWithContext(ctx, input)
}

func (a *EC2Activities) ModifyInstanceEventStartTime(ctx context.Context, input *ec2.ModifyInstanceEventStartTimeInput) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
	return a.client.ModifyInstanceEventStartTimeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyInstanceMetadataOptions(ctx context.Context, input *ec2.ModifyInstanceMetadataOptionsInput) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
	return a.client.ModifyInstanceMetadataOptionsWithContext(ctx, input)
}

func (a *EC2Activities) ModifyInstancePlacement(ctx context.Context, input *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
	return a.client.ModifyInstancePlacementWithContext(ctx, input)
}

func (a *EC2Activities) ModifyLaunchTemplate(ctx context.Context, input *ec2.ModifyLaunchTemplateInput) (*ec2.ModifyLaunchTemplateOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.ModifyLaunchTemplateWithContext(ctx, input)
}

func (a *EC2Activities) ModifyManagedPrefixList(ctx context.Context, input *ec2.ModifyManagedPrefixListInput) (*ec2.ModifyManagedPrefixListOutput, error) {
	return a.client.ModifyManagedPrefixListWithContext(ctx, input)
}

func (a *EC2Activities) ModifyNetworkInterfaceAttribute(ctx context.Context, input *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	return a.client.ModifyNetworkInterfaceAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyReservedInstances(ctx context.Context, input *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.ModifyReservedInstancesWithContext(ctx, input)
}

func (a *EC2Activities) ModifySnapshotAttribute(ctx context.Context, input *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
	return a.client.ModifySnapshotAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifySubnetAttribute(ctx context.Context, input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
	return a.client.ModifySubnetAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyTrafficMirrorFilterNetworkServices(ctx context.Context, input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
	return a.client.ModifyTrafficMirrorFilterNetworkServicesWithContext(ctx, input)
}

func (a *EC2Activities) ModifyTrafficMirrorFilterRule(ctx context.Context, input *ec2.ModifyTrafficMirrorFilterRuleInput) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
	return a.client.ModifyTrafficMirrorFilterRuleWithContext(ctx, input)
}

func (a *EC2Activities) ModifyTrafficMirrorSession(ctx context.Context, input *ec2.ModifyTrafficMirrorSessionInput) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
	return a.client.ModifyTrafficMirrorSessionWithContext(ctx, input)
}

func (a *EC2Activities) ModifyTransitGateway(ctx context.Context, input *ec2.ModifyTransitGatewayInput) (*ec2.ModifyTransitGatewayOutput, error) {
	return a.client.ModifyTransitGatewayWithContext(ctx, input)
}

func (a *EC2Activities) ModifyTransitGatewayPrefixListReference(ctx context.Context, input *ec2.ModifyTransitGatewayPrefixListReferenceInput) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
	return a.client.ModifyTransitGatewayPrefixListReferenceWithContext(ctx, input)
}

func (a *EC2Activities) ModifyTransitGatewayVpcAttachment(ctx context.Context, input *ec2.ModifyTransitGatewayVpcAttachmentInput) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
	return a.client.ModifyTransitGatewayVpcAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVolume(ctx context.Context, input *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
	return a.client.ModifyVolumeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVolumeAttribute(ctx context.Context, input *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
	return a.client.ModifyVolumeAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpcAttribute(ctx context.Context, input *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
	return a.client.ModifyVpcAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpcEndpoint(ctx context.Context, input *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
	return a.client.ModifyVpcEndpointWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpcEndpointConnectionNotification(ctx context.Context, input *ec2.ModifyVpcEndpointConnectionNotificationInput) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
	return a.client.ModifyVpcEndpointConnectionNotificationWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpcEndpointServiceConfiguration(ctx context.Context, input *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
	return a.client.ModifyVpcEndpointServiceConfigurationWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpcEndpointServicePermissions(ctx context.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
	return a.client.ModifyVpcEndpointServicePermissionsWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpcPeeringConnectionOptions(ctx context.Context, input *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
	return a.client.ModifyVpcPeeringConnectionOptionsWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpcTenancy(ctx context.Context, input *ec2.ModifyVpcTenancyInput) (*ec2.ModifyVpcTenancyOutput, error) {
	return a.client.ModifyVpcTenancyWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpnConnection(ctx context.Context, input *ec2.ModifyVpnConnectionInput) (*ec2.ModifyVpnConnectionOutput, error) {
	return a.client.ModifyVpnConnectionWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpnConnectionOptions(ctx context.Context, input *ec2.ModifyVpnConnectionOptionsInput) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
	return a.client.ModifyVpnConnectionOptionsWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpnTunnelCertificate(ctx context.Context, input *ec2.ModifyVpnTunnelCertificateInput) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
	return a.client.ModifyVpnTunnelCertificateWithContext(ctx, input)
}

func (a *EC2Activities) ModifyVpnTunnelOptions(ctx context.Context, input *ec2.ModifyVpnTunnelOptionsInput) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
	return a.client.ModifyVpnTunnelOptionsWithContext(ctx, input)
}

func (a *EC2Activities) MonitorInstances(ctx context.Context, input *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
	return a.client.MonitorInstancesWithContext(ctx, input)
}

func (a *EC2Activities) MoveAddressToVpc(ctx context.Context, input *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
	return a.client.MoveAddressToVpcWithContext(ctx, input)
}

func (a *EC2Activities) ProvisionByoipCidr(ctx context.Context, input *ec2.ProvisionByoipCidrInput) (*ec2.ProvisionByoipCidrOutput, error) {
	return a.client.ProvisionByoipCidrWithContext(ctx, input)
}

func (a *EC2Activities) PurchaseHostReservation(ctx context.Context, input *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.PurchaseHostReservationWithContext(ctx, input)
}

func (a *EC2Activities) PurchaseReservedInstancesOffering(ctx context.Context, input *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
	return a.client.PurchaseReservedInstancesOfferingWithContext(ctx, input)
}

func (a *EC2Activities) PurchaseScheduledInstances(ctx context.Context, input *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.PurchaseScheduledInstancesWithContext(ctx, input)
}

func (a *EC2Activities) RebootInstances(ctx context.Context, input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	return a.client.RebootInstancesWithContext(ctx, input)
}

func (a *EC2Activities) RegisterImage(ctx context.Context, input *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
	return a.client.RegisterImageWithContext(ctx, input)
}

func (a *EC2Activities) RegisterInstanceEventNotificationAttributes(ctx context.Context, input *ec2.RegisterInstanceEventNotificationAttributesInput) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
	return a.client.RegisterInstanceEventNotificationAttributesWithContext(ctx, input)
}

func (a *EC2Activities) RegisterTransitGatewayMulticastGroupMembers(ctx context.Context, input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
	return a.client.RegisterTransitGatewayMulticastGroupMembersWithContext(ctx, input)
}

func (a *EC2Activities) RegisterTransitGatewayMulticastGroupSources(ctx context.Context, input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
	return a.client.RegisterTransitGatewayMulticastGroupSourcesWithContext(ctx, input)
}

func (a *EC2Activities) RejectTransitGatewayPeeringAttachment(ctx context.Context, input *ec2.RejectTransitGatewayPeeringAttachmentInput) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
	return a.client.RejectTransitGatewayPeeringAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) RejectTransitGatewayVpcAttachment(ctx context.Context, input *ec2.RejectTransitGatewayVpcAttachmentInput) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
	return a.client.RejectTransitGatewayVpcAttachmentWithContext(ctx, input)
}

func (a *EC2Activities) RejectVpcEndpointConnections(ctx context.Context, input *ec2.RejectVpcEndpointConnectionsInput) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
	return a.client.RejectVpcEndpointConnectionsWithContext(ctx, input)
}

func (a *EC2Activities) RejectVpcPeeringConnection(ctx context.Context, input *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
	return a.client.RejectVpcPeeringConnectionWithContext(ctx, input)
}

func (a *EC2Activities) ReleaseAddress(ctx context.Context, input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	return a.client.ReleaseAddressWithContext(ctx, input)
}

func (a *EC2Activities) ReleaseHosts(ctx context.Context, input *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
	return a.client.ReleaseHostsWithContext(ctx, input)
}

func (a *EC2Activities) ReplaceIamInstanceProfileAssociation(ctx context.Context, input *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
	return a.client.ReplaceIamInstanceProfileAssociationWithContext(ctx, input)
}

func (a *EC2Activities) ReplaceNetworkAclAssociation(ctx context.Context, input *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
	return a.client.ReplaceNetworkAclAssociationWithContext(ctx, input)
}

func (a *EC2Activities) ReplaceNetworkAclEntry(ctx context.Context, input *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
	return a.client.ReplaceNetworkAclEntryWithContext(ctx, input)
}

func (a *EC2Activities) ReplaceRoute(ctx context.Context, input *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
	return a.client.ReplaceRouteWithContext(ctx, input)
}

func (a *EC2Activities) ReplaceRouteTableAssociation(ctx context.Context, input *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
	return a.client.ReplaceRouteTableAssociationWithContext(ctx, input)
}

func (a *EC2Activities) ReplaceTransitGatewayRoute(ctx context.Context, input *ec2.ReplaceTransitGatewayRouteInput) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
	return a.client.ReplaceTransitGatewayRouteWithContext(ctx, input)
}

func (a *EC2Activities) ReportInstanceStatus(ctx context.Context, input *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
	return a.client.ReportInstanceStatusWithContext(ctx, input)
}

func (a *EC2Activities) RequestSpotFleet(ctx context.Context, input *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
	return a.client.RequestSpotFleetWithContext(ctx, input)
}

func (a *EC2Activities) RequestSpotInstances(ctx context.Context, input *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.RequestSpotInstancesWithContext(ctx, input)
}

func (a *EC2Activities) ResetEbsDefaultKmsKeyId(ctx context.Context, input *ec2.ResetEbsDefaultKmsKeyIdInput) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
	return a.client.ResetEbsDefaultKmsKeyIdWithContext(ctx, input)
}

func (a *EC2Activities) ResetFpgaImageAttribute(ctx context.Context, input *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error) {
	return a.client.ResetFpgaImageAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ResetImageAttribute(ctx context.Context, input *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
	return a.client.ResetImageAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ResetInstanceAttribute(ctx context.Context, input *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
	return a.client.ResetInstanceAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ResetNetworkInterfaceAttribute(ctx context.Context, input *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
	return a.client.ResetNetworkInterfaceAttributeWithContext(ctx, input)
}

func (a *EC2Activities) ResetSnapshotAttribute(ctx context.Context, input *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
	return a.client.ResetSnapshotAttributeWithContext(ctx, input)
}

func (a *EC2Activities) RestoreAddressToClassic(ctx context.Context, input *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
	return a.client.RestoreAddressToClassicWithContext(ctx, input)
}

func (a *EC2Activities) RestoreManagedPrefixListVersion(ctx context.Context, input *ec2.RestoreManagedPrefixListVersionInput) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
	return a.client.RestoreManagedPrefixListVersionWithContext(ctx, input)
}

func (a *EC2Activities) RevokeClientVpnIngress(ctx context.Context, input *ec2.RevokeClientVpnIngressInput) (*ec2.RevokeClientVpnIngressOutput, error) {
	return a.client.RevokeClientVpnIngressWithContext(ctx, input)
}

func (a *EC2Activities) RevokeSecurityGroupEgress(ctx context.Context, input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	return a.client.RevokeSecurityGroupEgressWithContext(ctx, input)
}

func (a *EC2Activities) RevokeSecurityGroupIngress(ctx context.Context, input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	return a.client.RevokeSecurityGroupIngressWithContext(ctx, input)
}

func (a *EC2Activities) RunInstances(ctx context.Context, input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.RunInstancesWithContext(ctx, input)
}

func (a *EC2Activities) RunScheduledInstances(ctx context.Context, input *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.RunScheduledInstancesWithContext(ctx, input)
}

func (a *EC2Activities) SearchLocalGatewayRoutes(ctx context.Context, input *ec2.SearchLocalGatewayRoutesInput) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	return a.client.SearchLocalGatewayRoutesWithContext(ctx, input)
}

func (a *EC2Activities) SearchTransitGatewayMulticastGroups(ctx context.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	return a.client.SearchTransitGatewayMulticastGroupsWithContext(ctx, input)
}

func (a *EC2Activities) SearchTransitGatewayRoutes(ctx context.Context, input *ec2.SearchTransitGatewayRoutesInput) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	return a.client.SearchTransitGatewayRoutesWithContext(ctx, input)
}

func (a *EC2Activities) SendDiagnosticInterrupt(ctx context.Context, input *ec2.SendDiagnosticInterruptInput) (*ec2.SendDiagnosticInterruptOutput, error) {
	return a.client.SendDiagnosticInterruptWithContext(ctx, input)
}

func (a *EC2Activities) StartInstances(ctx context.Context, input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	return a.client.StartInstancesWithContext(ctx, input)
}

func (a *EC2Activities) StartVpcEndpointServicePrivateDnsVerification(ctx context.Context, input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
	return a.client.StartVpcEndpointServicePrivateDnsVerificationWithContext(ctx, input)
}

func (a *EC2Activities) StopInstances(ctx context.Context, input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	return a.client.StopInstancesWithContext(ctx, input)
}

func (a *EC2Activities) TerminateClientVpnConnections(ctx context.Context, input *ec2.TerminateClientVpnConnectionsInput) (*ec2.TerminateClientVpnConnectionsOutput, error) {
	return a.client.TerminateClientVpnConnectionsWithContext(ctx, input)
}

func (a *EC2Activities) TerminateInstances(ctx context.Context, input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	return a.client.TerminateInstancesWithContext(ctx, input)
}

func (a *EC2Activities) UnassignIpv6Addresses(ctx context.Context, input *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error) {
	return a.client.UnassignIpv6AddressesWithContext(ctx, input)
}

func (a *EC2Activities) UnassignPrivateIpAddresses(ctx context.Context, input *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	return a.client.UnassignPrivateIpAddressesWithContext(ctx, input)
}

func (a *EC2Activities) UnmonitorInstances(ctx context.Context, input *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
	return a.client.UnmonitorInstancesWithContext(ctx, input)
}

func (a *EC2Activities) UpdateSecurityGroupRuleDescriptionsEgress(ctx context.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
	return a.client.UpdateSecurityGroupRuleDescriptionsEgressWithContext(ctx, input)
}

func (a *EC2Activities) UpdateSecurityGroupRuleDescriptionsIngress(ctx context.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
	return a.client.UpdateSecurityGroupRuleDescriptionsIngressWithContext(ctx, input)
}

func (a *EC2Activities) WithdrawByoipCidr(ctx context.Context, input *ec2.WithdrawByoipCidrInput) (*ec2.WithdrawByoipCidrOutput, error) {
	return a.client.WithdrawByoipCidrWithContext(ctx, input)
}

func (a *EC2Activities) WaitUntilBundleTaskComplete(ctx context.Context, input *ec2.DescribeBundleTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilBundleTaskCompleteWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilConversionTaskCancelled(ctx context.Context, input *ec2.DescribeConversionTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilConversionTaskCancelledWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilConversionTaskCompleted(ctx context.Context, input *ec2.DescribeConversionTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilConversionTaskCompletedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilConversionTaskDeleted(ctx context.Context, input *ec2.DescribeConversionTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilConversionTaskDeletedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilCustomerGatewayAvailable(ctx context.Context, input *ec2.DescribeCustomerGatewaysInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilCustomerGatewayAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilExportTaskCancelled(ctx context.Context, input *ec2.DescribeExportTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilExportTaskCancelledWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilExportTaskCompleted(ctx context.Context, input *ec2.DescribeExportTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilExportTaskCompletedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilImageAvailable(ctx context.Context, input *ec2.DescribeImagesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilImageAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilImageExists(ctx context.Context, input *ec2.DescribeImagesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilImageExistsWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilInstanceExists(ctx context.Context, input *ec2.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceExistsWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilInstanceRunning(ctx context.Context, input *ec2.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceRunningWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilInstanceStatusOk(ctx context.Context, input *ec2.DescribeInstanceStatusInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceStatusOkWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilInstanceStopped(ctx context.Context, input *ec2.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceStoppedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilInstanceTerminated(ctx context.Context, input *ec2.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceTerminatedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilKeyPairExists(ctx context.Context, input *ec2.DescribeKeyPairsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilKeyPairExistsWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilNatGatewayAvailable(ctx context.Context, input *ec2.DescribeNatGatewaysInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilNatGatewayAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilNetworkInterfaceAvailable(ctx context.Context, input *ec2.DescribeNetworkInterfacesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilNetworkInterfaceAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilPasswordDataAvailable(ctx context.Context, input *ec2.GetPasswordDataInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilPasswordDataAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilSecurityGroupExists(ctx context.Context, input *ec2.DescribeSecurityGroupsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilSecurityGroupExistsWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilSnapshotCompleted(ctx context.Context, input *ec2.DescribeSnapshotsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilSnapshotCompletedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilSpotInstanceRequestFulfilled(ctx context.Context, input *ec2.DescribeSpotInstanceRequestsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilSpotInstanceRequestFulfilledWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilSubnetAvailable(ctx context.Context, input *ec2.DescribeSubnetsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilSubnetAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilSystemStatusOk(ctx context.Context, input *ec2.DescribeInstanceStatusInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilSystemStatusOkWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVolumeAvailable(ctx context.Context, input *ec2.DescribeVolumesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVolumeAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVolumeDeleted(ctx context.Context, input *ec2.DescribeVolumesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVolumeDeletedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVolumeInUse(ctx context.Context, input *ec2.DescribeVolumesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVolumeInUseWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVpcAvailable(ctx context.Context, input *ec2.DescribeVpcsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVpcAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVpcExists(ctx context.Context, input *ec2.DescribeVpcsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVpcExistsWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVpcPeeringConnectionDeleted(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVpcPeeringConnectionDeletedWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVpcPeeringConnectionExists(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVpcPeeringConnectionExistsWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVpnConnectionAvailable(ctx context.Context, input *ec2.DescribeVpnConnectionsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVpnConnectionAvailableWithContext(ctx, input, options...)
	})
}

func (a *EC2Activities) WaitUntilVpnConnectionDeleted(ctx context.Context, input *ec2.DescribeVpnConnectionsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilVpnConnectionDeletedWithContext(ctx, input, options...)
	})
}
