
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type EC2Activities struct {
	client ec2iface.EC2API
}

func NewEC2Activities(session *session.Session, config... *aws.Config) *EC2Activities {
    client := ec2.New(session, config...)
    return &EC2Activities{client: client}
}

func (a *EC2Activities) AcceptReservedInstancesExchangeQuote(input *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
    return a.client.AcceptReservedInstancesExchangeQuote(input)
}

func (a *EC2Activities) AcceptTransitGatewayPeeringAttachment(input *ec2.AcceptTransitGatewayPeeringAttachmentInput) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
    return a.client.AcceptTransitGatewayPeeringAttachment(input)
}

func (a *EC2Activities) AcceptTransitGatewayVpcAttachment(input *ec2.AcceptTransitGatewayVpcAttachmentInput) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
    return a.client.AcceptTransitGatewayVpcAttachment(input)
}

func (a *EC2Activities) AcceptVpcEndpointConnections(input *ec2.AcceptVpcEndpointConnectionsInput) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
    return a.client.AcceptVpcEndpointConnections(input)
}

func (a *EC2Activities) AcceptVpcPeeringConnection(input *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
    return a.client.AcceptVpcPeeringConnection(input)
}

func (a *EC2Activities) AdvertiseByoipCidr(input *ec2.AdvertiseByoipCidrInput) (*ec2.AdvertiseByoipCidrOutput, error) {
    return a.client.AdvertiseByoipCidr(input)
}

func (a *EC2Activities) AllocateAddress(input *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
    return a.client.AllocateAddress(input)
}

func (a *EC2Activities) AllocateHosts(input *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
    return a.client.AllocateHosts(input)
}

func (a *EC2Activities) ApplySecurityGroupsToClientVpnTargetNetwork(input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
    return a.client.ApplySecurityGroupsToClientVpnTargetNetwork(input)
}

func (a *EC2Activities) AssignIpv6Addresses(input *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error) {
    return a.client.AssignIpv6Addresses(input)
}

func (a *EC2Activities) AssignPrivateIpAddresses(input *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
    return a.client.AssignPrivateIpAddresses(input)
}

func (a *EC2Activities) AssociateAddress(input *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {
    return a.client.AssociateAddress(input)
}

func (a *EC2Activities) AssociateClientVpnTargetNetwork(input *ec2.AssociateClientVpnTargetNetworkInput) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
    return a.client.AssociateClientVpnTargetNetwork(input)
}

func (a *EC2Activities) AssociateDhcpOptions(input *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
    return a.client.AssociateDhcpOptions(input)
}

func (a *EC2Activities) AssociateIamInstanceProfile(input *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error) {
    return a.client.AssociateIamInstanceProfile(input)
}

func (a *EC2Activities) AssociateRouteTable(input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
    return a.client.AssociateRouteTable(input)
}

func (a *EC2Activities) AssociateSubnetCidrBlock(input *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error) {
    return a.client.AssociateSubnetCidrBlock(input)
}

func (a *EC2Activities) AssociateTransitGatewayMulticastDomain(input *ec2.AssociateTransitGatewayMulticastDomainInput) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
    return a.client.AssociateTransitGatewayMulticastDomain(input)
}

func (a *EC2Activities) AssociateTransitGatewayRouteTable(input *ec2.AssociateTransitGatewayRouteTableInput) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
    return a.client.AssociateTransitGatewayRouteTable(input)
}

func (a *EC2Activities) AssociateVpcCidrBlock(input *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
    return a.client.AssociateVpcCidrBlock(input)
}

func (a *EC2Activities) AttachClassicLinkVpc(input *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
    return a.client.AttachClassicLinkVpc(input)
}

func (a *EC2Activities) AttachInternetGateway(input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
    return a.client.AttachInternetGateway(input)
}

func (a *EC2Activities) AttachNetworkInterface(input *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
    return a.client.AttachNetworkInterface(input)
}

func (a *EC2Activities) AttachVolume(input *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
    return a.client.AttachVolume(input)
}

func (a *EC2Activities) AttachVpnGateway(input *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
    return a.client.AttachVpnGateway(input)
}

func (a *EC2Activities) AuthorizeClientVpnIngress(input *ec2.AuthorizeClientVpnIngressInput) (*ec2.AuthorizeClientVpnIngressOutput, error) {
    return a.client.AuthorizeClientVpnIngress(input)
}

func (a *EC2Activities) AuthorizeSecurityGroupEgress(input *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
    return a.client.AuthorizeSecurityGroupEgress(input)
}

func (a *EC2Activities) AuthorizeSecurityGroupIngress(input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
    return a.client.AuthorizeSecurityGroupIngress(input)
}

func (a *EC2Activities) BundleInstance(input *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
    return a.client.BundleInstance(input)
}

func (a *EC2Activities) CancelBundleTask(input *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
    return a.client.CancelBundleTask(input)
}

func (a *EC2Activities) CancelCapacityReservation(input *ec2.CancelCapacityReservationInput) (*ec2.CancelCapacityReservationOutput, error) {
    return a.client.CancelCapacityReservation(input)
}

func (a *EC2Activities) CancelConversionTask(input *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
    return a.client.CancelConversionTask(input)
}

func (a *EC2Activities) CancelExportTask(input *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
    return a.client.CancelExportTask(input)
}

func (a *EC2Activities) CancelImportTask(input *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
    return a.client.CancelImportTask(input)
}

func (a *EC2Activities) CancelReservedInstancesListing(input *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
    return a.client.CancelReservedInstancesListing(input)
}

func (a *EC2Activities) CancelSpotFleetRequests(input *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
    return a.client.CancelSpotFleetRequests(input)
}

func (a *EC2Activities) CancelSpotInstanceRequests(input *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
    return a.client.CancelSpotInstanceRequests(input)
}

func (a *EC2Activities) ConfirmProductInstance(input *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
    return a.client.ConfirmProductInstance(input)
}

func (a *EC2Activities) CopyFpgaImage(input *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error) {
    return a.client.CopyFpgaImage(input)
}

func (a *EC2Activities) CopyImage(input *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
    return a.client.CopyImage(input)
}

func (a *EC2Activities) CopySnapshot(input *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
    return a.client.CopySnapshot(input)
}

func (a *EC2Activities) CreateCapacityReservation(input *ec2.CreateCapacityReservationInput) (*ec2.CreateCapacityReservationOutput, error) {
    return a.client.CreateCapacityReservation(input)
}

func (a *EC2Activities) CreateCarrierGateway(input *ec2.CreateCarrierGatewayInput) (*ec2.CreateCarrierGatewayOutput, error) {
    return a.client.CreateCarrierGateway(input)
}

func (a *EC2Activities) CreateClientVpnEndpoint(input *ec2.CreateClientVpnEndpointInput) (*ec2.CreateClientVpnEndpointOutput, error) {
    return a.client.CreateClientVpnEndpoint(input)
}

func (a *EC2Activities) CreateClientVpnRoute(input *ec2.CreateClientVpnRouteInput) (*ec2.CreateClientVpnRouteOutput, error) {
    return a.client.CreateClientVpnRoute(input)
}

func (a *EC2Activities) CreateCustomerGateway(input *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
    return a.client.CreateCustomerGateway(input)
}

func (a *EC2Activities) CreateDefaultSubnet(input *ec2.CreateDefaultSubnetInput) (*ec2.CreateDefaultSubnetOutput, error) {
    return a.client.CreateDefaultSubnet(input)
}

func (a *EC2Activities) CreateDefaultVpc(input *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error) {
    return a.client.CreateDefaultVpc(input)
}

func (a *EC2Activities) CreateDhcpOptions(input *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
    return a.client.CreateDhcpOptions(input)
}

func (a *EC2Activities) CreateEgressOnlyInternetGateway(input *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
    return a.client.CreateEgressOnlyInternetGateway(input)
}

func (a *EC2Activities) CreateFleet(input *ec2.CreateFleetInput) (*ec2.CreateFleetOutput, error) {
    return a.client.CreateFleet(input)
}

func (a *EC2Activities) CreateFlowLogs(input *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
    return a.client.CreateFlowLogs(input)
}

func (a *EC2Activities) CreateFpgaImage(input *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error) {
    return a.client.CreateFpgaImage(input)
}

func (a *EC2Activities) CreateImage(input *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
    return a.client.CreateImage(input)
}

func (a *EC2Activities) CreateInstanceExportTask(input *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
    return a.client.CreateInstanceExportTask(input)
}

func (a *EC2Activities) CreateInternetGateway(input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
    return a.client.CreateInternetGateway(input)
}

func (a *EC2Activities) CreateKeyPair(input *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
    return a.client.CreateKeyPair(input)
}

func (a *EC2Activities) CreateLaunchTemplate(input *ec2.CreateLaunchTemplateInput) (*ec2.CreateLaunchTemplateOutput, error) {
    return a.client.CreateLaunchTemplate(input)
}

func (a *EC2Activities) CreateLaunchTemplateVersion(input *ec2.CreateLaunchTemplateVersionInput) (*ec2.CreateLaunchTemplateVersionOutput, error) {
    return a.client.CreateLaunchTemplateVersion(input)
}

func (a *EC2Activities) CreateLocalGatewayRoute(input *ec2.CreateLocalGatewayRouteInput) (*ec2.CreateLocalGatewayRouteOutput, error) {
    return a.client.CreateLocalGatewayRoute(input)
}

func (a *EC2Activities) CreateLocalGatewayRouteTableVpcAssociation(input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
    return a.client.CreateLocalGatewayRouteTableVpcAssociation(input)
}

func (a *EC2Activities) CreateManagedPrefixList(input *ec2.CreateManagedPrefixListInput) (*ec2.CreateManagedPrefixListOutput, error) {
    return a.client.CreateManagedPrefixList(input)
}

func (a *EC2Activities) CreateNatGateway(input *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
    return a.client.CreateNatGateway(input)
}

func (a *EC2Activities) CreateNetworkAcl(input *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
    return a.client.CreateNetworkAcl(input)
}

func (a *EC2Activities) CreateNetworkAclEntry(input *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
    return a.client.CreateNetworkAclEntry(input)
}

func (a *EC2Activities) CreateNetworkInterface(input *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
    return a.client.CreateNetworkInterface(input)
}

func (a *EC2Activities) CreateNetworkInterfacePermission(input *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
    return a.client.CreateNetworkInterfacePermission(input)
}

func (a *EC2Activities) CreatePlacementGroup(input *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
    return a.client.CreatePlacementGroup(input)
}

func (a *EC2Activities) CreateReservedInstancesListing(input *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
    return a.client.CreateReservedInstancesListing(input)
}

func (a *EC2Activities) CreateRoute(input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
    return a.client.CreateRoute(input)
}

func (a *EC2Activities) CreateRouteTable(input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
    return a.client.CreateRouteTable(input)
}

func (a *EC2Activities) CreateSecurityGroup(input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
    return a.client.CreateSecurityGroup(input)
}

func (a *EC2Activities) CreateSnapshot(input *ec2.CreateSnapshotInput) (*ec2.Snapshot, error) {
    return a.client.CreateSnapshot(input)
}

func (a *EC2Activities) CreateSnapshots(input *ec2.CreateSnapshotsInput) (*ec2.CreateSnapshotsOutput, error) {
    return a.client.CreateSnapshots(input)
}

func (a *EC2Activities) CreateSpotDatafeedSubscription(input *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
    return a.client.CreateSpotDatafeedSubscription(input)
}

func (a *EC2Activities) CreateSubnet(input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
    return a.client.CreateSubnet(input)
}

func (a *EC2Activities) CreateTags(input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
    return a.client.CreateTags(input)
}

func (a *EC2Activities) CreateTrafficMirrorFilter(input *ec2.CreateTrafficMirrorFilterInput) (*ec2.CreateTrafficMirrorFilterOutput, error) {
    return a.client.CreateTrafficMirrorFilter(input)
}

func (a *EC2Activities) CreateTrafficMirrorFilterRule(input *ec2.CreateTrafficMirrorFilterRuleInput) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
    return a.client.CreateTrafficMirrorFilterRule(input)
}

func (a *EC2Activities) CreateTrafficMirrorSession(input *ec2.CreateTrafficMirrorSessionInput) (*ec2.CreateTrafficMirrorSessionOutput, error) {
    return a.client.CreateTrafficMirrorSession(input)
}

func (a *EC2Activities) CreateTrafficMirrorTarget(input *ec2.CreateTrafficMirrorTargetInput) (*ec2.CreateTrafficMirrorTargetOutput, error) {
    return a.client.CreateTrafficMirrorTarget(input)
}

func (a *EC2Activities) CreateTransitGateway(input *ec2.CreateTransitGatewayInput) (*ec2.CreateTransitGatewayOutput, error) {
    return a.client.CreateTransitGateway(input)
}

func (a *EC2Activities) CreateTransitGatewayMulticastDomain(input *ec2.CreateTransitGatewayMulticastDomainInput) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
    return a.client.CreateTransitGatewayMulticastDomain(input)
}

func (a *EC2Activities) CreateTransitGatewayPeeringAttachment(input *ec2.CreateTransitGatewayPeeringAttachmentInput) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
    return a.client.CreateTransitGatewayPeeringAttachment(input)
}

func (a *EC2Activities) CreateTransitGatewayPrefixListReference(input *ec2.CreateTransitGatewayPrefixListReferenceInput) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
    return a.client.CreateTransitGatewayPrefixListReference(input)
}

func (a *EC2Activities) CreateTransitGatewayRoute(input *ec2.CreateTransitGatewayRouteInput) (*ec2.CreateTransitGatewayRouteOutput, error) {
    return a.client.CreateTransitGatewayRoute(input)
}

func (a *EC2Activities) CreateTransitGatewayRouteTable(input *ec2.CreateTransitGatewayRouteTableInput) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
    return a.client.CreateTransitGatewayRouteTable(input)
}

func (a *EC2Activities) CreateTransitGatewayVpcAttachment(input *ec2.CreateTransitGatewayVpcAttachmentInput) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
    return a.client.CreateTransitGatewayVpcAttachment(input)
}

func (a *EC2Activities) CreateVolume(input *ec2.CreateVolumeInput) (*ec2.Volume, error) {
    return a.client.CreateVolume(input)
}

func (a *EC2Activities) CreateVpc(input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
    return a.client.CreateVpc(input)
}

func (a *EC2Activities) CreateVpcEndpoint(input *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
    return a.client.CreateVpcEndpoint(input)
}

func (a *EC2Activities) CreateVpcEndpointConnectionNotification(input *ec2.CreateVpcEndpointConnectionNotificationInput) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
    return a.client.CreateVpcEndpointConnectionNotification(input)
}

func (a *EC2Activities) CreateVpcEndpointServiceConfiguration(input *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
    return a.client.CreateVpcEndpointServiceConfiguration(input)
}

func (a *EC2Activities) CreateVpcPeeringConnection(input *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
    return a.client.CreateVpcPeeringConnection(input)
}

func (a *EC2Activities) CreateVpnConnection(input *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
    return a.client.CreateVpnConnection(input)
}

func (a *EC2Activities) CreateVpnConnectionRoute(input *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
    return a.client.CreateVpnConnectionRoute(input)
}

func (a *EC2Activities) CreateVpnGateway(input *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
    return a.client.CreateVpnGateway(input)
}

func (a *EC2Activities) DeleteCarrierGateway(input *ec2.DeleteCarrierGatewayInput) (*ec2.DeleteCarrierGatewayOutput, error) {
    return a.client.DeleteCarrierGateway(input)
}

func (a *EC2Activities) DeleteClientVpnEndpoint(input *ec2.DeleteClientVpnEndpointInput) (*ec2.DeleteClientVpnEndpointOutput, error) {
    return a.client.DeleteClientVpnEndpoint(input)
}

func (a *EC2Activities) DeleteClientVpnRoute(input *ec2.DeleteClientVpnRouteInput) (*ec2.DeleteClientVpnRouteOutput, error) {
    return a.client.DeleteClientVpnRoute(input)
}

func (a *EC2Activities) DeleteCustomerGateway(input *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
    return a.client.DeleteCustomerGateway(input)
}

func (a *EC2Activities) DeleteDhcpOptions(input *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
    return a.client.DeleteDhcpOptions(input)
}

func (a *EC2Activities) DeleteEgressOnlyInternetGateway(input *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
    return a.client.DeleteEgressOnlyInternetGateway(input)
}

func (a *EC2Activities) DeleteFleets(input *ec2.DeleteFleetsInput) (*ec2.DeleteFleetsOutput, error) {
    return a.client.DeleteFleets(input)
}

func (a *EC2Activities) DeleteFlowLogs(input *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
    return a.client.DeleteFlowLogs(input)
}

func (a *EC2Activities) DeleteFpgaImage(input *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error) {
    return a.client.DeleteFpgaImage(input)
}

func (a *EC2Activities) DeleteInternetGateway(input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
    return a.client.DeleteInternetGateway(input)
}

func (a *EC2Activities) DeleteKeyPair(input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
    return a.client.DeleteKeyPair(input)
}

func (a *EC2Activities) DeleteLaunchTemplate(input *ec2.DeleteLaunchTemplateInput) (*ec2.DeleteLaunchTemplateOutput, error) {
    return a.client.DeleteLaunchTemplate(input)
}

func (a *EC2Activities) DeleteLaunchTemplateVersions(input *ec2.DeleteLaunchTemplateVersionsInput) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
    return a.client.DeleteLaunchTemplateVersions(input)
}

func (a *EC2Activities) DeleteLocalGatewayRoute(input *ec2.DeleteLocalGatewayRouteInput) (*ec2.DeleteLocalGatewayRouteOutput, error) {
    return a.client.DeleteLocalGatewayRoute(input)
}

func (a *EC2Activities) DeleteLocalGatewayRouteTableVpcAssociation(input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
    return a.client.DeleteLocalGatewayRouteTableVpcAssociation(input)
}

func (a *EC2Activities) DeleteManagedPrefixList(input *ec2.DeleteManagedPrefixListInput) (*ec2.DeleteManagedPrefixListOutput, error) {
    return a.client.DeleteManagedPrefixList(input)
}

func (a *EC2Activities) DeleteNatGateway(input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
    return a.client.DeleteNatGateway(input)
}

func (a *EC2Activities) DeleteNetworkAcl(input *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
    return a.client.DeleteNetworkAcl(input)
}

func (a *EC2Activities) DeleteNetworkAclEntry(input *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
    return a.client.DeleteNetworkAclEntry(input)
}

func (a *EC2Activities) DeleteNetworkInterface(input *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
    return a.client.DeleteNetworkInterface(input)
}

func (a *EC2Activities) DeleteNetworkInterfacePermission(input *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
    return a.client.DeleteNetworkInterfacePermission(input)
}

func (a *EC2Activities) DeletePlacementGroup(input *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
    return a.client.DeletePlacementGroup(input)
}

func (a *EC2Activities) DeleteQueuedReservedInstances(input *ec2.DeleteQueuedReservedInstancesInput) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
    return a.client.DeleteQueuedReservedInstances(input)
}

func (a *EC2Activities) DeleteRoute(input *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
    return a.client.DeleteRoute(input)
}

func (a *EC2Activities) DeleteRouteTable(input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
    return a.client.DeleteRouteTable(input)
}

func (a *EC2Activities) DeleteSecurityGroup(input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
    return a.client.DeleteSecurityGroup(input)
}

func (a *EC2Activities) DeleteSnapshot(input *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
    return a.client.DeleteSnapshot(input)
}

func (a *EC2Activities) DeleteSpotDatafeedSubscription(input *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
    return a.client.DeleteSpotDatafeedSubscription(input)
}

func (a *EC2Activities) DeleteSubnet(input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
    return a.client.DeleteSubnet(input)
}

func (a *EC2Activities) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}

func (a *EC2Activities) DeleteTrafficMirrorFilter(input *ec2.DeleteTrafficMirrorFilterInput) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
    return a.client.DeleteTrafficMirrorFilter(input)
}

func (a *EC2Activities) DeleteTrafficMirrorFilterRule(input *ec2.DeleteTrafficMirrorFilterRuleInput) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
    return a.client.DeleteTrafficMirrorFilterRule(input)
}

func (a *EC2Activities) DeleteTrafficMirrorSession(input *ec2.DeleteTrafficMirrorSessionInput) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
    return a.client.DeleteTrafficMirrorSession(input)
}

func (a *EC2Activities) DeleteTrafficMirrorTarget(input *ec2.DeleteTrafficMirrorTargetInput) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
    return a.client.DeleteTrafficMirrorTarget(input)
}

func (a *EC2Activities) DeleteTransitGateway(input *ec2.DeleteTransitGatewayInput) (*ec2.DeleteTransitGatewayOutput, error) {
    return a.client.DeleteTransitGateway(input)
}

func (a *EC2Activities) DeleteTransitGatewayMulticastDomain(input *ec2.DeleteTransitGatewayMulticastDomainInput) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
    return a.client.DeleteTransitGatewayMulticastDomain(input)
}

func (a *EC2Activities) DeleteTransitGatewayPeeringAttachment(input *ec2.DeleteTransitGatewayPeeringAttachmentInput) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
    return a.client.DeleteTransitGatewayPeeringAttachment(input)
}

func (a *EC2Activities) DeleteTransitGatewayPrefixListReference(input *ec2.DeleteTransitGatewayPrefixListReferenceInput) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
    return a.client.DeleteTransitGatewayPrefixListReference(input)
}

func (a *EC2Activities) DeleteTransitGatewayRoute(input *ec2.DeleteTransitGatewayRouteInput) (*ec2.DeleteTransitGatewayRouteOutput, error) {
    return a.client.DeleteTransitGatewayRoute(input)
}

func (a *EC2Activities) DeleteTransitGatewayRouteTable(input *ec2.DeleteTransitGatewayRouteTableInput) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
    return a.client.DeleteTransitGatewayRouteTable(input)
}

func (a *EC2Activities) DeleteTransitGatewayVpcAttachment(input *ec2.DeleteTransitGatewayVpcAttachmentInput) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
    return a.client.DeleteTransitGatewayVpcAttachment(input)
}

func (a *EC2Activities) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
    return a.client.DeleteVolume(input)
}

func (a *EC2Activities) DeleteVpc(input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
    return a.client.DeleteVpc(input)
}

func (a *EC2Activities) DeleteVpcEndpointConnectionNotifications(input *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
    return a.client.DeleteVpcEndpointConnectionNotifications(input)
}

func (a *EC2Activities) DeleteVpcEndpointServiceConfigurations(input *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
    return a.client.DeleteVpcEndpointServiceConfigurations(input)
}

func (a *EC2Activities) DeleteVpcEndpoints(input *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
    return a.client.DeleteVpcEndpoints(input)
}

func (a *EC2Activities) DeleteVpcPeeringConnection(input *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
    return a.client.DeleteVpcPeeringConnection(input)
}

func (a *EC2Activities) DeleteVpnConnection(input *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
    return a.client.DeleteVpnConnection(input)
}

func (a *EC2Activities) DeleteVpnConnectionRoute(input *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
    return a.client.DeleteVpnConnectionRoute(input)
}

func (a *EC2Activities) DeleteVpnGateway(input *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
    return a.client.DeleteVpnGateway(input)
}

func (a *EC2Activities) DeprovisionByoipCidr(input *ec2.DeprovisionByoipCidrInput) (*ec2.DeprovisionByoipCidrOutput, error) {
    return a.client.DeprovisionByoipCidr(input)
}

func (a *EC2Activities) DeregisterImage(input *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
    return a.client.DeregisterImage(input)
}

func (a *EC2Activities) DeregisterInstanceEventNotificationAttributes(input *ec2.DeregisterInstanceEventNotificationAttributesInput) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
    return a.client.DeregisterInstanceEventNotificationAttributes(input)
}

func (a *EC2Activities) DeregisterTransitGatewayMulticastGroupMembers(input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
    return a.client.DeregisterTransitGatewayMulticastGroupMembers(input)
}

func (a *EC2Activities) DeregisterTransitGatewayMulticastGroupSources(input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
    return a.client.DeregisterTransitGatewayMulticastGroupSources(input)
}

func (a *EC2Activities) DescribeAccountAttributes(input *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributes(input)
}

func (a *EC2Activities) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
    return a.client.DescribeAddresses(input)
}

func (a *EC2Activities) DescribeAggregateIdFormat(input *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error) {
    return a.client.DescribeAggregateIdFormat(input)
}

func (a *EC2Activities) DescribeAvailabilityZones(input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
    return a.client.DescribeAvailabilityZones(input)
}

func (a *EC2Activities) DescribeBundleTasks(input *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
    return a.client.DescribeBundleTasks(input)
}

func (a *EC2Activities) DescribeByoipCidrs(input *ec2.DescribeByoipCidrsInput) (*ec2.DescribeByoipCidrsOutput, error) {
    return a.client.DescribeByoipCidrs(input)
}

func (a *EC2Activities) DescribeCapacityReservations(input *ec2.DescribeCapacityReservationsInput) (*ec2.DescribeCapacityReservationsOutput, error) {
    return a.client.DescribeCapacityReservations(input)
}

func (a *EC2Activities) DescribeCarrierGateways(input *ec2.DescribeCarrierGatewaysInput) (*ec2.DescribeCarrierGatewaysOutput, error) {
    return a.client.DescribeCarrierGateways(input)
}

func (a *EC2Activities) DescribeClassicLinkInstances(input *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
    return a.client.DescribeClassicLinkInstances(input)
}

func (a *EC2Activities) DescribeClientVpnAuthorizationRules(input *ec2.DescribeClientVpnAuthorizationRulesInput) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
    return a.client.DescribeClientVpnAuthorizationRules(input)
}

func (a *EC2Activities) DescribeClientVpnConnections(input *ec2.DescribeClientVpnConnectionsInput) (*ec2.DescribeClientVpnConnectionsOutput, error) {
    return a.client.DescribeClientVpnConnections(input)
}

func (a *EC2Activities) DescribeClientVpnEndpoints(input *ec2.DescribeClientVpnEndpointsInput) (*ec2.DescribeClientVpnEndpointsOutput, error) {
    return a.client.DescribeClientVpnEndpoints(input)
}

func (a *EC2Activities) DescribeClientVpnRoutes(input *ec2.DescribeClientVpnRoutesInput) (*ec2.DescribeClientVpnRoutesOutput, error) {
    return a.client.DescribeClientVpnRoutes(input)
}

func (a *EC2Activities) DescribeClientVpnTargetNetworks(input *ec2.DescribeClientVpnTargetNetworksInput) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
    return a.client.DescribeClientVpnTargetNetworks(input)
}

func (a *EC2Activities) DescribeCoipPools(input *ec2.DescribeCoipPoolsInput) (*ec2.DescribeCoipPoolsOutput, error) {
    return a.client.DescribeCoipPools(input)
}

func (a *EC2Activities) DescribeConversionTasks(input *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
    return a.client.DescribeConversionTasks(input)
}

func (a *EC2Activities) DescribeCustomerGateways(input *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
    return a.client.DescribeCustomerGateways(input)
}

func (a *EC2Activities) DescribeDhcpOptions(input *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
    return a.client.DescribeDhcpOptions(input)
}

func (a *EC2Activities) DescribeEgressOnlyInternetGateways(input *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
    return a.client.DescribeEgressOnlyInternetGateways(input)
}

func (a *EC2Activities) DescribeElasticGpus(input *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
    return a.client.DescribeElasticGpus(input)
}

func (a *EC2Activities) DescribeExportImageTasks(input *ec2.DescribeExportImageTasksInput) (*ec2.DescribeExportImageTasksOutput, error) {
    return a.client.DescribeExportImageTasks(input)
}

func (a *EC2Activities) DescribeExportTasks(input *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
    return a.client.DescribeExportTasks(input)
}

func (a *EC2Activities) DescribeFastSnapshotRestores(input *ec2.DescribeFastSnapshotRestoresInput) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
    return a.client.DescribeFastSnapshotRestores(input)
}

func (a *EC2Activities) DescribeFleetHistory(input *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error) {
    return a.client.DescribeFleetHistory(input)
}

func (a *EC2Activities) DescribeFleetInstances(input *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error) {
    return a.client.DescribeFleetInstances(input)
}

func (a *EC2Activities) DescribeFleets(input *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error) {
    return a.client.DescribeFleets(input)
}

func (a *EC2Activities) DescribeFlowLogs(input *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
    return a.client.DescribeFlowLogs(input)
}

func (a *EC2Activities) DescribeFpgaImageAttribute(input *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
    return a.client.DescribeFpgaImageAttribute(input)
}

func (a *EC2Activities) DescribeFpgaImages(input *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
    return a.client.DescribeFpgaImages(input)
}

func (a *EC2Activities) DescribeHostReservationOfferings(input *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
    return a.client.DescribeHostReservationOfferings(input)
}

func (a *EC2Activities) DescribeHostReservations(input *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
    return a.client.DescribeHostReservations(input)
}

func (a *EC2Activities) DescribeHosts(input *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
    return a.client.DescribeHosts(input)
}

func (a *EC2Activities) DescribeIamInstanceProfileAssociations(input *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
    return a.client.DescribeIamInstanceProfileAssociations(input)
}

func (a *EC2Activities) DescribeIdFormat(input *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
    return a.client.DescribeIdFormat(input)
}

func (a *EC2Activities) DescribeIdentityIdFormat(input *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
    return a.client.DescribeIdentityIdFormat(input)
}

func (a *EC2Activities) DescribeImageAttribute(input *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
    return a.client.DescribeImageAttribute(input)
}

func (a *EC2Activities) DescribeImages(input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
    return a.client.DescribeImages(input)
}

func (a *EC2Activities) DescribeImportImageTasks(input *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
    return a.client.DescribeImportImageTasks(input)
}

func (a *EC2Activities) DescribeImportSnapshotTasks(input *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
    return a.client.DescribeImportSnapshotTasks(input)
}

func (a *EC2Activities) DescribeInstanceAttribute(input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
    return a.client.DescribeInstanceAttribute(input)
}

func (a *EC2Activities) DescribeInstanceCreditSpecifications(input *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
    return a.client.DescribeInstanceCreditSpecifications(input)
}

func (a *EC2Activities) DescribeInstanceEventNotificationAttributes(input *ec2.DescribeInstanceEventNotificationAttributesInput) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
    return a.client.DescribeInstanceEventNotificationAttributes(input)
}

func (a *EC2Activities) DescribeInstanceStatus(input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
    return a.client.DescribeInstanceStatus(input)
}

func (a *EC2Activities) DescribeInstanceTypeOfferings(input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
    return a.client.DescribeInstanceTypeOfferings(input)
}

func (a *EC2Activities) DescribeInstanceTypes(input *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
    return a.client.DescribeInstanceTypes(input)
}

func (a *EC2Activities) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
    return a.client.DescribeInstances(input)
}

func (a *EC2Activities) DescribeInternetGateways(input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
    return a.client.DescribeInternetGateways(input)
}

func (a *EC2Activities) DescribeIpv6Pools(input *ec2.DescribeIpv6PoolsInput) (*ec2.DescribeIpv6PoolsOutput, error) {
    return a.client.DescribeIpv6Pools(input)
}

func (a *EC2Activities) DescribeKeyPairs(input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
    return a.client.DescribeKeyPairs(input)
}

func (a *EC2Activities) DescribeLaunchTemplateVersions(input *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
    return a.client.DescribeLaunchTemplateVersions(input)
}

func (a *EC2Activities) DescribeLaunchTemplates(input *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error) {
    return a.client.DescribeLaunchTemplates(input)
}

func (a *EC2Activities) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
    return a.client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input)
}

func (a *EC2Activities) DescribeLocalGatewayRouteTableVpcAssociations(input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
    return a.client.DescribeLocalGatewayRouteTableVpcAssociations(input)
}

func (a *EC2Activities) DescribeLocalGatewayRouteTables(input *ec2.DescribeLocalGatewayRouteTablesInput) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
    return a.client.DescribeLocalGatewayRouteTables(input)
}

func (a *EC2Activities) DescribeLocalGatewayVirtualInterfaceGroups(input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
    return a.client.DescribeLocalGatewayVirtualInterfaceGroups(input)
}

func (a *EC2Activities) DescribeLocalGatewayVirtualInterfaces(input *ec2.DescribeLocalGatewayVirtualInterfacesInput) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
    return a.client.DescribeLocalGatewayVirtualInterfaces(input)
}

func (a *EC2Activities) DescribeLocalGateways(input *ec2.DescribeLocalGatewaysInput) (*ec2.DescribeLocalGatewaysOutput, error) {
    return a.client.DescribeLocalGateways(input)
}

func (a *EC2Activities) DescribeManagedPrefixLists(input *ec2.DescribeManagedPrefixListsInput) (*ec2.DescribeManagedPrefixListsOutput, error) {
    return a.client.DescribeManagedPrefixLists(input)
}

func (a *EC2Activities) DescribeMovingAddresses(input *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
    return a.client.DescribeMovingAddresses(input)
}

func (a *EC2Activities) DescribeNatGateways(input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
    return a.client.DescribeNatGateways(input)
}

func (a *EC2Activities) DescribeNetworkAcls(input *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
    return a.client.DescribeNetworkAcls(input)
}

func (a *EC2Activities) DescribeNetworkInterfaceAttribute(input *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
    return a.client.DescribeNetworkInterfaceAttribute(input)
}

func (a *EC2Activities) DescribeNetworkInterfacePermissions(input *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
    return a.client.DescribeNetworkInterfacePermissions(input)
}

func (a *EC2Activities) DescribeNetworkInterfaces(input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
    return a.client.DescribeNetworkInterfaces(input)
}

func (a *EC2Activities) DescribePlacementGroups(input *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
    return a.client.DescribePlacementGroups(input)
}

func (a *EC2Activities) DescribePrefixLists(input *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
    return a.client.DescribePrefixLists(input)
}

func (a *EC2Activities) DescribePrincipalIdFormat(input *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error) {
    return a.client.DescribePrincipalIdFormat(input)
}

func (a *EC2Activities) DescribePublicIpv4Pools(input *ec2.DescribePublicIpv4PoolsInput) (*ec2.DescribePublicIpv4PoolsOutput, error) {
    return a.client.DescribePublicIpv4Pools(input)
}

func (a *EC2Activities) DescribeRegions(input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
    return a.client.DescribeRegions(input)
}

func (a *EC2Activities) DescribeReservedInstances(input *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
    return a.client.DescribeReservedInstances(input)
}

func (a *EC2Activities) DescribeReservedInstancesListings(input *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
    return a.client.DescribeReservedInstancesListings(input)
}

func (a *EC2Activities) DescribeReservedInstancesModifications(input *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
    return a.client.DescribeReservedInstancesModifications(input)
}

func (a *EC2Activities) DescribeReservedInstancesOfferings(input *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
    return a.client.DescribeReservedInstancesOfferings(input)
}

func (a *EC2Activities) DescribeRouteTables(input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
    return a.client.DescribeRouteTables(input)
}

func (a *EC2Activities) DescribeScheduledInstanceAvailability(input *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
    return a.client.DescribeScheduledInstanceAvailability(input)
}

func (a *EC2Activities) DescribeScheduledInstances(input *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
    return a.client.DescribeScheduledInstances(input)
}

func (a *EC2Activities) DescribeSecurityGroupReferences(input *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
    return a.client.DescribeSecurityGroupReferences(input)
}

func (a *EC2Activities) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
    return a.client.DescribeSecurityGroups(input)
}

func (a *EC2Activities) DescribeSnapshotAttribute(input *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
    return a.client.DescribeSnapshotAttribute(input)
}

func (a *EC2Activities) DescribeSnapshots(input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
    return a.client.DescribeSnapshots(input)
}

func (a *EC2Activities) DescribeSpotDatafeedSubscription(input *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
    return a.client.DescribeSpotDatafeedSubscription(input)
}

func (a *EC2Activities) DescribeSpotFleetInstances(input *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
    return a.client.DescribeSpotFleetInstances(input)
}

func (a *EC2Activities) DescribeSpotFleetRequestHistory(input *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
    return a.client.DescribeSpotFleetRequestHistory(input)
}

func (a *EC2Activities) DescribeSpotFleetRequests(input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
    return a.client.DescribeSpotFleetRequests(input)
}

func (a *EC2Activities) DescribeSpotInstanceRequests(input *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
    return a.client.DescribeSpotInstanceRequests(input)
}

func (a *EC2Activities) DescribeSpotPriceHistory(input *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
    return a.client.DescribeSpotPriceHistory(input)
}

func (a *EC2Activities) DescribeStaleSecurityGroups(input *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
    return a.client.DescribeStaleSecurityGroups(input)
}

func (a *EC2Activities) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
    return a.client.DescribeSubnets(input)
}

func (a *EC2Activities) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}

func (a *EC2Activities) DescribeTrafficMirrorFilters(input *ec2.DescribeTrafficMirrorFiltersInput) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
    return a.client.DescribeTrafficMirrorFilters(input)
}

func (a *EC2Activities) DescribeTrafficMirrorSessions(input *ec2.DescribeTrafficMirrorSessionsInput) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
    return a.client.DescribeTrafficMirrorSessions(input)
}

func (a *EC2Activities) DescribeTrafficMirrorTargets(input *ec2.DescribeTrafficMirrorTargetsInput) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
    return a.client.DescribeTrafficMirrorTargets(input)
}

func (a *EC2Activities) DescribeTransitGatewayAttachments(input *ec2.DescribeTransitGatewayAttachmentsInput) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
    return a.client.DescribeTransitGatewayAttachments(input)
}

func (a *EC2Activities) DescribeTransitGatewayMulticastDomains(input *ec2.DescribeTransitGatewayMulticastDomainsInput) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
    return a.client.DescribeTransitGatewayMulticastDomains(input)
}

func (a *EC2Activities) DescribeTransitGatewayPeeringAttachments(input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
    return a.client.DescribeTransitGatewayPeeringAttachments(input)
}

func (a *EC2Activities) DescribeTransitGatewayRouteTables(input *ec2.DescribeTransitGatewayRouteTablesInput) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
    return a.client.DescribeTransitGatewayRouteTables(input)
}

func (a *EC2Activities) DescribeTransitGatewayVpcAttachments(input *ec2.DescribeTransitGatewayVpcAttachmentsInput) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
    return a.client.DescribeTransitGatewayVpcAttachments(input)
}

func (a *EC2Activities) DescribeTransitGateways(input *ec2.DescribeTransitGatewaysInput) (*ec2.DescribeTransitGatewaysOutput, error) {
    return a.client.DescribeTransitGateways(input)
}

func (a *EC2Activities) DescribeVolumeAttribute(input *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
    return a.client.DescribeVolumeAttribute(input)
}

func (a *EC2Activities) DescribeVolumeStatus(input *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
    return a.client.DescribeVolumeStatus(input)
}

func (a *EC2Activities) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
    return a.client.DescribeVolumes(input)
}

func (a *EC2Activities) DescribeVolumesModifications(input *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
    return a.client.DescribeVolumesModifications(input)
}

func (a *EC2Activities) DescribeVpcAttribute(input *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
    return a.client.DescribeVpcAttribute(input)
}

func (a *EC2Activities) DescribeVpcClassicLink(input *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
    return a.client.DescribeVpcClassicLink(input)
}

func (a *EC2Activities) DescribeVpcClassicLinkDnsSupport(input *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
    return a.client.DescribeVpcClassicLinkDnsSupport(input)
}

func (a *EC2Activities) DescribeVpcEndpointConnectionNotifications(input *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
    return a.client.DescribeVpcEndpointConnectionNotifications(input)
}

func (a *EC2Activities) DescribeVpcEndpointConnections(input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
    return a.client.DescribeVpcEndpointConnections(input)
}

func (a *EC2Activities) DescribeVpcEndpointServiceConfigurations(input *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
    return a.client.DescribeVpcEndpointServiceConfigurations(input)
}

func (a *EC2Activities) DescribeVpcEndpointServicePermissions(input *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
    return a.client.DescribeVpcEndpointServicePermissions(input)
}

func (a *EC2Activities) DescribeVpcEndpointServices(input *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
    return a.client.DescribeVpcEndpointServices(input)
}

func (a *EC2Activities) DescribeVpcEndpoints(input *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
    return a.client.DescribeVpcEndpoints(input)
}

func (a *EC2Activities) DescribeVpcPeeringConnections(input *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
    return a.client.DescribeVpcPeeringConnections(input)
}

func (a *EC2Activities) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
    return a.client.DescribeVpcs(input)
}

func (a *EC2Activities) DescribeVpnConnections(input *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
    return a.client.DescribeVpnConnections(input)
}

func (a *EC2Activities) DescribeVpnGateways(input *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
    return a.client.DescribeVpnGateways(input)
}

func (a *EC2Activities) DetachClassicLinkVpc(input *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
    return a.client.DetachClassicLinkVpc(input)
}

func (a *EC2Activities) DetachInternetGateway(input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
    return a.client.DetachInternetGateway(input)
}

func (a *EC2Activities) DetachNetworkInterface(input *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
    return a.client.DetachNetworkInterface(input)
}

func (a *EC2Activities) DetachVolume(input *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
    return a.client.DetachVolume(input)
}

func (a *EC2Activities) DetachVpnGateway(input *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
    return a.client.DetachVpnGateway(input)
}

func (a *EC2Activities) DisableEbsEncryptionByDefault(input *ec2.DisableEbsEncryptionByDefaultInput) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
    return a.client.DisableEbsEncryptionByDefault(input)
}

func (a *EC2Activities) DisableFastSnapshotRestores(input *ec2.DisableFastSnapshotRestoresInput) (*ec2.DisableFastSnapshotRestoresOutput, error) {
    return a.client.DisableFastSnapshotRestores(input)
}

func (a *EC2Activities) DisableTransitGatewayRouteTablePropagation(input *ec2.DisableTransitGatewayRouteTablePropagationInput) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
    return a.client.DisableTransitGatewayRouteTablePropagation(input)
}

func (a *EC2Activities) DisableVgwRoutePropagation(input *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
    return a.client.DisableVgwRoutePropagation(input)
}

func (a *EC2Activities) DisableVpcClassicLink(input *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
    return a.client.DisableVpcClassicLink(input)
}

func (a *EC2Activities) DisableVpcClassicLinkDnsSupport(input *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
    return a.client.DisableVpcClassicLinkDnsSupport(input)
}

func (a *EC2Activities) DisassociateAddress(input *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {
    return a.client.DisassociateAddress(input)
}

func (a *EC2Activities) DisassociateClientVpnTargetNetwork(input *ec2.DisassociateClientVpnTargetNetworkInput) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
    return a.client.DisassociateClientVpnTargetNetwork(input)
}

func (a *EC2Activities) DisassociateIamInstanceProfile(input *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error) {
    return a.client.DisassociateIamInstanceProfile(input)
}

func (a *EC2Activities) DisassociateRouteTable(input *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
    return a.client.DisassociateRouteTable(input)
}

func (a *EC2Activities) DisassociateSubnetCidrBlock(input *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
    return a.client.DisassociateSubnetCidrBlock(input)
}

func (a *EC2Activities) DisassociateTransitGatewayMulticastDomain(input *ec2.DisassociateTransitGatewayMulticastDomainInput) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
    return a.client.DisassociateTransitGatewayMulticastDomain(input)
}

func (a *EC2Activities) DisassociateTransitGatewayRouteTable(input *ec2.DisassociateTransitGatewayRouteTableInput) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
    return a.client.DisassociateTransitGatewayRouteTable(input)
}

func (a *EC2Activities) DisassociateVpcCidrBlock(input *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error) {
    return a.client.DisassociateVpcCidrBlock(input)
}

func (a *EC2Activities) EnableEbsEncryptionByDefault(input *ec2.EnableEbsEncryptionByDefaultInput) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
    return a.client.EnableEbsEncryptionByDefault(input)
}

func (a *EC2Activities) EnableFastSnapshotRestores(input *ec2.EnableFastSnapshotRestoresInput) (*ec2.EnableFastSnapshotRestoresOutput, error) {
    return a.client.EnableFastSnapshotRestores(input)
}

func (a *EC2Activities) EnableTransitGatewayRouteTablePropagation(input *ec2.EnableTransitGatewayRouteTablePropagationInput) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
    return a.client.EnableTransitGatewayRouteTablePropagation(input)
}

func (a *EC2Activities) EnableVgwRoutePropagation(input *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
    return a.client.EnableVgwRoutePropagation(input)
}

func (a *EC2Activities) EnableVolumeIO(input *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
    return a.client.EnableVolumeIO(input)
}

func (a *EC2Activities) EnableVpcClassicLink(input *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
    return a.client.EnableVpcClassicLink(input)
}

func (a *EC2Activities) EnableVpcClassicLinkDnsSupport(input *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
    return a.client.EnableVpcClassicLinkDnsSupport(input)
}

func (a *EC2Activities) ExportClientVpnClientCertificateRevocationList(input *ec2.ExportClientVpnClientCertificateRevocationListInput) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
    return a.client.ExportClientVpnClientCertificateRevocationList(input)
}

func (a *EC2Activities) ExportClientVpnClientConfiguration(input *ec2.ExportClientVpnClientConfigurationInput) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
    return a.client.ExportClientVpnClientConfiguration(input)
}

func (a *EC2Activities) ExportImage(input *ec2.ExportImageInput) (*ec2.ExportImageOutput, error) {
    return a.client.ExportImage(input)
}

func (a *EC2Activities) ExportTransitGatewayRoutes(input *ec2.ExportTransitGatewayRoutesInput) (*ec2.ExportTransitGatewayRoutesOutput, error) {
    return a.client.ExportTransitGatewayRoutes(input)
}

func (a *EC2Activities) GetAssociatedIpv6PoolCidrs(input *ec2.GetAssociatedIpv6PoolCidrsInput) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
    return a.client.GetAssociatedIpv6PoolCidrs(input)
}

func (a *EC2Activities) GetCapacityReservationUsage(input *ec2.GetCapacityReservationUsageInput) (*ec2.GetCapacityReservationUsageOutput, error) {
    return a.client.GetCapacityReservationUsage(input)
}

func (a *EC2Activities) GetCoipPoolUsage(input *ec2.GetCoipPoolUsageInput) (*ec2.GetCoipPoolUsageOutput, error) {
    return a.client.GetCoipPoolUsage(input)
}

func (a *EC2Activities) GetConsoleOutput(input *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
    return a.client.GetConsoleOutput(input)
}

func (a *EC2Activities) GetConsoleScreenshot(input *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
    return a.client.GetConsoleScreenshot(input)
}

func (a *EC2Activities) GetDefaultCreditSpecification(input *ec2.GetDefaultCreditSpecificationInput) (*ec2.GetDefaultCreditSpecificationOutput, error) {
    return a.client.GetDefaultCreditSpecification(input)
}

func (a *EC2Activities) GetEbsDefaultKmsKeyId(input *ec2.GetEbsDefaultKmsKeyIdInput) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
    return a.client.GetEbsDefaultKmsKeyId(input)
}

func (a *EC2Activities) GetEbsEncryptionByDefault(input *ec2.GetEbsEncryptionByDefaultInput) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
    return a.client.GetEbsEncryptionByDefault(input)
}

func (a *EC2Activities) GetGroupsForCapacityReservation(input *ec2.GetGroupsForCapacityReservationInput) (*ec2.GetGroupsForCapacityReservationOutput, error) {
    return a.client.GetGroupsForCapacityReservation(input)
}

func (a *EC2Activities) GetHostReservationPurchasePreview(input *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
    return a.client.GetHostReservationPurchasePreview(input)
}

func (a *EC2Activities) GetLaunchTemplateData(input *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error) {
    return a.client.GetLaunchTemplateData(input)
}

func (a *EC2Activities) GetManagedPrefixListAssociations(input *ec2.GetManagedPrefixListAssociationsInput) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
    return a.client.GetManagedPrefixListAssociations(input)
}

func (a *EC2Activities) GetManagedPrefixListEntries(input *ec2.GetManagedPrefixListEntriesInput) (*ec2.GetManagedPrefixListEntriesOutput, error) {
    return a.client.GetManagedPrefixListEntries(input)
}

func (a *EC2Activities) GetPasswordData(input *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
    return a.client.GetPasswordData(input)
}

func (a *EC2Activities) GetReservedInstancesExchangeQuote(input *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
    return a.client.GetReservedInstancesExchangeQuote(input)
}

func (a *EC2Activities) GetTransitGatewayAttachmentPropagations(input *ec2.GetTransitGatewayAttachmentPropagationsInput) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
    return a.client.GetTransitGatewayAttachmentPropagations(input)
}

func (a *EC2Activities) GetTransitGatewayMulticastDomainAssociations(input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
    return a.client.GetTransitGatewayMulticastDomainAssociations(input)
}

func (a *EC2Activities) GetTransitGatewayPrefixListReferences(input *ec2.GetTransitGatewayPrefixListReferencesInput) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
    return a.client.GetTransitGatewayPrefixListReferences(input)
}

func (a *EC2Activities) GetTransitGatewayRouteTableAssociations(input *ec2.GetTransitGatewayRouteTableAssociationsInput) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
    return a.client.GetTransitGatewayRouteTableAssociations(input)
}

func (a *EC2Activities) GetTransitGatewayRouteTablePropagations(input *ec2.GetTransitGatewayRouteTablePropagationsInput) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
    return a.client.GetTransitGatewayRouteTablePropagations(input)
}

func (a *EC2Activities) ImportClientVpnClientCertificateRevocationList(input *ec2.ImportClientVpnClientCertificateRevocationListInput) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
    return a.client.ImportClientVpnClientCertificateRevocationList(input)
}

func (a *EC2Activities) ImportImage(input *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
    return a.client.ImportImage(input)
}

func (a *EC2Activities) ImportInstance(input *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
    return a.client.ImportInstance(input)
}

func (a *EC2Activities) ImportKeyPair(input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
    return a.client.ImportKeyPair(input)
}

func (a *EC2Activities) ImportSnapshot(input *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
    return a.client.ImportSnapshot(input)
}

func (a *EC2Activities) ImportVolume(input *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
    return a.client.ImportVolume(input)
}

func (a *EC2Activities) ModifyAvailabilityZoneGroup(input *ec2.ModifyAvailabilityZoneGroupInput) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
    return a.client.ModifyAvailabilityZoneGroup(input)
}

func (a *EC2Activities) ModifyCapacityReservation(input *ec2.ModifyCapacityReservationInput) (*ec2.ModifyCapacityReservationOutput, error) {
    return a.client.ModifyCapacityReservation(input)
}

func (a *EC2Activities) ModifyClientVpnEndpoint(input *ec2.ModifyClientVpnEndpointInput) (*ec2.ModifyClientVpnEndpointOutput, error) {
    return a.client.ModifyClientVpnEndpoint(input)
}

func (a *EC2Activities) ModifyDefaultCreditSpecification(input *ec2.ModifyDefaultCreditSpecificationInput) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
    return a.client.ModifyDefaultCreditSpecification(input)
}

func (a *EC2Activities) ModifyEbsDefaultKmsKeyId(input *ec2.ModifyEbsDefaultKmsKeyIdInput) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
    return a.client.ModifyEbsDefaultKmsKeyId(input)
}

func (a *EC2Activities) ModifyFleet(input *ec2.ModifyFleetInput) (*ec2.ModifyFleetOutput, error) {
    return a.client.ModifyFleet(input)
}

func (a *EC2Activities) ModifyFpgaImageAttribute(input *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error) {
    return a.client.ModifyFpgaImageAttribute(input)
}

func (a *EC2Activities) ModifyHosts(input *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
    return a.client.ModifyHosts(input)
}

func (a *EC2Activities) ModifyIdFormat(input *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
    return a.client.ModifyIdFormat(input)
}

func (a *EC2Activities) ModifyIdentityIdFormat(input *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {
    return a.client.ModifyIdentityIdFormat(input)
}

func (a *EC2Activities) ModifyImageAttribute(input *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
    return a.client.ModifyImageAttribute(input)
}

func (a *EC2Activities) ModifyInstanceAttribute(input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
    return a.client.ModifyInstanceAttribute(input)
}

func (a *EC2Activities) ModifyInstanceCapacityReservationAttributes(input *ec2.ModifyInstanceCapacityReservationAttributesInput) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
    return a.client.ModifyInstanceCapacityReservationAttributes(input)
}

func (a *EC2Activities) ModifyInstanceCreditSpecification(input *ec2.ModifyInstanceCreditSpecificationInput) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
    return a.client.ModifyInstanceCreditSpecification(input)
}

func (a *EC2Activities) ModifyInstanceEventStartTime(input *ec2.ModifyInstanceEventStartTimeInput) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
    return a.client.ModifyInstanceEventStartTime(input)
}

func (a *EC2Activities) ModifyInstanceMetadataOptions(input *ec2.ModifyInstanceMetadataOptionsInput) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
    return a.client.ModifyInstanceMetadataOptions(input)
}

func (a *EC2Activities) ModifyInstancePlacement(input *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
    return a.client.ModifyInstancePlacement(input)
}

func (a *EC2Activities) ModifyLaunchTemplate(input *ec2.ModifyLaunchTemplateInput) (*ec2.ModifyLaunchTemplateOutput, error) {
    return a.client.ModifyLaunchTemplate(input)
}

func (a *EC2Activities) ModifyManagedPrefixList(input *ec2.ModifyManagedPrefixListInput) (*ec2.ModifyManagedPrefixListOutput, error) {
    return a.client.ModifyManagedPrefixList(input)
}

func (a *EC2Activities) ModifyNetworkInterfaceAttribute(input *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
    return a.client.ModifyNetworkInterfaceAttribute(input)
}

func (a *EC2Activities) ModifyReservedInstances(input *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
    return a.client.ModifyReservedInstances(input)
}

func (a *EC2Activities) ModifySnapshotAttribute(input *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
    return a.client.ModifySnapshotAttribute(input)
}

func (a *EC2Activities) ModifySubnetAttribute(input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
    return a.client.ModifySubnetAttribute(input)
}

func (a *EC2Activities) ModifyTrafficMirrorFilterNetworkServices(input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
    return a.client.ModifyTrafficMirrorFilterNetworkServices(input)
}

func (a *EC2Activities) ModifyTrafficMirrorFilterRule(input *ec2.ModifyTrafficMirrorFilterRuleInput) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
    return a.client.ModifyTrafficMirrorFilterRule(input)
}

func (a *EC2Activities) ModifyTrafficMirrorSession(input *ec2.ModifyTrafficMirrorSessionInput) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
    return a.client.ModifyTrafficMirrorSession(input)
}

func (a *EC2Activities) ModifyTransitGateway(input *ec2.ModifyTransitGatewayInput) (*ec2.ModifyTransitGatewayOutput, error) {
    return a.client.ModifyTransitGateway(input)
}

func (a *EC2Activities) ModifyTransitGatewayPrefixListReference(input *ec2.ModifyTransitGatewayPrefixListReferenceInput) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
    return a.client.ModifyTransitGatewayPrefixListReference(input)
}

func (a *EC2Activities) ModifyTransitGatewayVpcAttachment(input *ec2.ModifyTransitGatewayVpcAttachmentInput) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
    return a.client.ModifyTransitGatewayVpcAttachment(input)
}

func (a *EC2Activities) ModifyVolume(input *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
    return a.client.ModifyVolume(input)
}

func (a *EC2Activities) ModifyVolumeAttribute(input *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
    return a.client.ModifyVolumeAttribute(input)
}

func (a *EC2Activities) ModifyVpcAttribute(input *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
    return a.client.ModifyVpcAttribute(input)
}

func (a *EC2Activities) ModifyVpcEndpoint(input *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
    return a.client.ModifyVpcEndpoint(input)
}

func (a *EC2Activities) ModifyVpcEndpointConnectionNotification(input *ec2.ModifyVpcEndpointConnectionNotificationInput) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
    return a.client.ModifyVpcEndpointConnectionNotification(input)
}

func (a *EC2Activities) ModifyVpcEndpointServiceConfiguration(input *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
    return a.client.ModifyVpcEndpointServiceConfiguration(input)
}

func (a *EC2Activities) ModifyVpcEndpointServicePermissions(input *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
    return a.client.ModifyVpcEndpointServicePermissions(input)
}

func (a *EC2Activities) ModifyVpcPeeringConnectionOptions(input *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
    return a.client.ModifyVpcPeeringConnectionOptions(input)
}

func (a *EC2Activities) ModifyVpcTenancy(input *ec2.ModifyVpcTenancyInput) (*ec2.ModifyVpcTenancyOutput, error) {
    return a.client.ModifyVpcTenancy(input)
}

func (a *EC2Activities) ModifyVpnConnection(input *ec2.ModifyVpnConnectionInput) (*ec2.ModifyVpnConnectionOutput, error) {
    return a.client.ModifyVpnConnection(input)
}

func (a *EC2Activities) ModifyVpnConnectionOptions(input *ec2.ModifyVpnConnectionOptionsInput) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
    return a.client.ModifyVpnConnectionOptions(input)
}

func (a *EC2Activities) ModifyVpnTunnelCertificate(input *ec2.ModifyVpnTunnelCertificateInput) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
    return a.client.ModifyVpnTunnelCertificate(input)
}

func (a *EC2Activities) ModifyVpnTunnelOptions(input *ec2.ModifyVpnTunnelOptionsInput) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
    return a.client.ModifyVpnTunnelOptions(input)
}

func (a *EC2Activities) MonitorInstances(input *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
    return a.client.MonitorInstances(input)
}

func (a *EC2Activities) MoveAddressToVpc(input *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
    return a.client.MoveAddressToVpc(input)
}

func (a *EC2Activities) ProvisionByoipCidr(input *ec2.ProvisionByoipCidrInput) (*ec2.ProvisionByoipCidrOutput, error) {
    return a.client.ProvisionByoipCidr(input)
}

func (a *EC2Activities) PurchaseHostReservation(input *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {
    return a.client.PurchaseHostReservation(input)
}

func (a *EC2Activities) PurchaseReservedInstancesOffering(input *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
    return a.client.PurchaseReservedInstancesOffering(input)
}

func (a *EC2Activities) PurchaseScheduledInstances(input *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
    return a.client.PurchaseScheduledInstances(input)
}

func (a *EC2Activities) RebootInstances(input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
    return a.client.RebootInstances(input)
}

func (a *EC2Activities) RegisterImage(input *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
    return a.client.RegisterImage(input)
}

func (a *EC2Activities) RegisterInstanceEventNotificationAttributes(input *ec2.RegisterInstanceEventNotificationAttributesInput) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
    return a.client.RegisterInstanceEventNotificationAttributes(input)
}

func (a *EC2Activities) RegisterTransitGatewayMulticastGroupMembers(input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
    return a.client.RegisterTransitGatewayMulticastGroupMembers(input)
}

func (a *EC2Activities) RegisterTransitGatewayMulticastGroupSources(input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
    return a.client.RegisterTransitGatewayMulticastGroupSources(input)
}

func (a *EC2Activities) RejectTransitGatewayPeeringAttachment(input *ec2.RejectTransitGatewayPeeringAttachmentInput) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
    return a.client.RejectTransitGatewayPeeringAttachment(input)
}

func (a *EC2Activities) RejectTransitGatewayVpcAttachment(input *ec2.RejectTransitGatewayVpcAttachmentInput) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
    return a.client.RejectTransitGatewayVpcAttachment(input)
}

func (a *EC2Activities) RejectVpcEndpointConnections(input *ec2.RejectVpcEndpointConnectionsInput) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
    return a.client.RejectVpcEndpointConnections(input)
}

func (a *EC2Activities) RejectVpcPeeringConnection(input *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
    return a.client.RejectVpcPeeringConnection(input)
}

func (a *EC2Activities) ReleaseAddress(input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
    return a.client.ReleaseAddress(input)
}

func (a *EC2Activities) ReleaseHosts(input *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
    return a.client.ReleaseHosts(input)
}

func (a *EC2Activities) ReplaceIamInstanceProfileAssociation(input *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
    return a.client.ReplaceIamInstanceProfileAssociation(input)
}

func (a *EC2Activities) ReplaceNetworkAclAssociation(input *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
    return a.client.ReplaceNetworkAclAssociation(input)
}

func (a *EC2Activities) ReplaceNetworkAclEntry(input *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
    return a.client.ReplaceNetworkAclEntry(input)
}

func (a *EC2Activities) ReplaceRoute(input *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
    return a.client.ReplaceRoute(input)
}

func (a *EC2Activities) ReplaceRouteTableAssociation(input *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
    return a.client.ReplaceRouteTableAssociation(input)
}

func (a *EC2Activities) ReplaceTransitGatewayRoute(input *ec2.ReplaceTransitGatewayRouteInput) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
    return a.client.ReplaceTransitGatewayRoute(input)
}

func (a *EC2Activities) ReportInstanceStatus(input *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
    return a.client.ReportInstanceStatus(input)
}

func (a *EC2Activities) RequestSpotFleet(input *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
    return a.client.RequestSpotFleet(input)
}

func (a *EC2Activities) RequestSpotInstances(input *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
    return a.client.RequestSpotInstances(input)
}

func (a *EC2Activities) ResetEbsDefaultKmsKeyId(input *ec2.ResetEbsDefaultKmsKeyIdInput) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
    return a.client.ResetEbsDefaultKmsKeyId(input)
}

func (a *EC2Activities) ResetFpgaImageAttribute(input *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error) {
    return a.client.ResetFpgaImageAttribute(input)
}

func (a *EC2Activities) ResetImageAttribute(input *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
    return a.client.ResetImageAttribute(input)
}

func (a *EC2Activities) ResetInstanceAttribute(input *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
    return a.client.ResetInstanceAttribute(input)
}

func (a *EC2Activities) ResetNetworkInterfaceAttribute(input *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
    return a.client.ResetNetworkInterfaceAttribute(input)
}

func (a *EC2Activities) ResetSnapshotAttribute(input *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
    return a.client.ResetSnapshotAttribute(input)
}

func (a *EC2Activities) RestoreAddressToClassic(input *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
    return a.client.RestoreAddressToClassic(input)
}

func (a *EC2Activities) RestoreManagedPrefixListVersion(input *ec2.RestoreManagedPrefixListVersionInput) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
    return a.client.RestoreManagedPrefixListVersion(input)
}

func (a *EC2Activities) RevokeClientVpnIngress(input *ec2.RevokeClientVpnIngressInput) (*ec2.RevokeClientVpnIngressOutput, error) {
    return a.client.RevokeClientVpnIngress(input)
}

func (a *EC2Activities) RevokeSecurityGroupEgress(input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
    return a.client.RevokeSecurityGroupEgress(input)
}

func (a *EC2Activities) RevokeSecurityGroupIngress(input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
    return a.client.RevokeSecurityGroupIngress(input)
}

func (a *EC2Activities) RunInstances(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
    return a.client.RunInstances(input)
}

func (a *EC2Activities) RunScheduledInstances(input *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
    return a.client.RunScheduledInstances(input)
}

func (a *EC2Activities) SearchLocalGatewayRoutes(input *ec2.SearchLocalGatewayRoutesInput) (*ec2.SearchLocalGatewayRoutesOutput, error) {
    return a.client.SearchLocalGatewayRoutes(input)
}

func (a *EC2Activities) SearchTransitGatewayMulticastGroups(input *ec2.SearchTransitGatewayMulticastGroupsInput) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
    return a.client.SearchTransitGatewayMulticastGroups(input)
}

func (a *EC2Activities) SearchTransitGatewayRoutes(input *ec2.SearchTransitGatewayRoutesInput) (*ec2.SearchTransitGatewayRoutesOutput, error) {
    return a.client.SearchTransitGatewayRoutes(input)
}

func (a *EC2Activities) SendDiagnosticInterrupt(input *ec2.SendDiagnosticInterruptInput) (*ec2.SendDiagnosticInterruptOutput, error) {
    return a.client.SendDiagnosticInterrupt(input)
}

func (a *EC2Activities) StartInstances(input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
    return a.client.StartInstances(input)
}

func (a *EC2Activities) StartVpcEndpointServicePrivateDnsVerification(input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
    return a.client.StartVpcEndpointServicePrivateDnsVerification(input)
}

func (a *EC2Activities) StopInstances(input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
    return a.client.StopInstances(input)
}

func (a *EC2Activities) TerminateClientVpnConnections(input *ec2.TerminateClientVpnConnectionsInput) (*ec2.TerminateClientVpnConnectionsOutput, error) {
    return a.client.TerminateClientVpnConnections(input)
}

func (a *EC2Activities) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
    return a.client.TerminateInstances(input)
}

func (a *EC2Activities) UnassignIpv6Addresses(input *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error) {
    return a.client.UnassignIpv6Addresses(input)
}

func (a *EC2Activities) UnassignPrivateIpAddresses(input *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {
    return a.client.UnassignPrivateIpAddresses(input)
}

func (a *EC2Activities) UnmonitorInstances(input *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
    return a.client.UnmonitorInstances(input)
}

func (a *EC2Activities) UpdateSecurityGroupRuleDescriptionsEgress(input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
    return a.client.UpdateSecurityGroupRuleDescriptionsEgress(input)
}

func (a *EC2Activities) UpdateSecurityGroupRuleDescriptionsIngress(input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
    return a.client.UpdateSecurityGroupRuleDescriptionsIngress(input)
}

func (a *EC2Activities) WithdrawByoipCidr(input *ec2.WithdrawByoipCidrInput) (*ec2.WithdrawByoipCidrOutput, error) {
    return a.client.WithdrawByoipCidr(input)
}

func (a *EC2Activities) WaitUntilBundleTaskComplete(input *ec2.DescribeBundleTasksInput) error {
    return a.client.WaitUntilBundleTaskComplete(input)
}

func (a *EC2Activities) WaitUntilConversionTaskCancelled(input *ec2.DescribeConversionTasksInput) error {
    return a.client.WaitUntilConversionTaskCancelled(input)
}

func (a *EC2Activities) WaitUntilConversionTaskCompleted(input *ec2.DescribeConversionTasksInput) error {
    return a.client.WaitUntilConversionTaskCompleted(input)
}

func (a *EC2Activities) WaitUntilConversionTaskDeleted(input *ec2.DescribeConversionTasksInput) error {
    return a.client.WaitUntilConversionTaskDeleted(input)
}

func (a *EC2Activities) WaitUntilCustomerGatewayAvailable(input *ec2.DescribeCustomerGatewaysInput) error {
    return a.client.WaitUntilCustomerGatewayAvailable(input)
}

func (a *EC2Activities) WaitUntilExportTaskCancelled(input *ec2.DescribeExportTasksInput) error {
    return a.client.WaitUntilExportTaskCancelled(input)
}

func (a *EC2Activities) WaitUntilExportTaskCompleted(input *ec2.DescribeExportTasksInput) error {
    return a.client.WaitUntilExportTaskCompleted(input)
}

func (a *EC2Activities) WaitUntilImageAvailable(input *ec2.DescribeImagesInput) error {
    return a.client.WaitUntilImageAvailable(input)
}

func (a *EC2Activities) WaitUntilImageExists(input *ec2.DescribeImagesInput) error {
    return a.client.WaitUntilImageExists(input)
}

func (a *EC2Activities) WaitUntilInstanceExists(input *ec2.DescribeInstancesInput) error {
    return a.client.WaitUntilInstanceExists(input)
}

func (a *EC2Activities) WaitUntilInstanceRunning(input *ec2.DescribeInstancesInput) error {
    return a.client.WaitUntilInstanceRunning(input)
}

func (a *EC2Activities) WaitUntilInstanceStatusOk(input *ec2.DescribeInstanceStatusInput) error {
    return a.client.WaitUntilInstanceStatusOk(input)
}

func (a *EC2Activities) WaitUntilInstanceStopped(input *ec2.DescribeInstancesInput) error {
    return a.client.WaitUntilInstanceStopped(input)
}

func (a *EC2Activities) WaitUntilInstanceTerminated(input *ec2.DescribeInstancesInput) error {
    return a.client.WaitUntilInstanceTerminated(input)
}

func (a *EC2Activities) WaitUntilKeyPairExists(input *ec2.DescribeKeyPairsInput) error {
    return a.client.WaitUntilKeyPairExists(input)
}

func (a *EC2Activities) WaitUntilNatGatewayAvailable(input *ec2.DescribeNatGatewaysInput) error {
    return a.client.WaitUntilNatGatewayAvailable(input)
}

func (a *EC2Activities) WaitUntilNetworkInterfaceAvailable(input *ec2.DescribeNetworkInterfacesInput) error {
    return a.client.WaitUntilNetworkInterfaceAvailable(input)
}

func (a *EC2Activities) WaitUntilPasswordDataAvailable(input *ec2.GetPasswordDataInput) error {
    return a.client.WaitUntilPasswordDataAvailable(input)
}

func (a *EC2Activities) WaitUntilSecurityGroupExists(input *ec2.DescribeSecurityGroupsInput) error {
    return a.client.WaitUntilSecurityGroupExists(input)
}

func (a *EC2Activities) WaitUntilSnapshotCompleted(input *ec2.DescribeSnapshotsInput) error {
    return a.client.WaitUntilSnapshotCompleted(input)
}

func (a *EC2Activities) WaitUntilSpotInstanceRequestFulfilled(input *ec2.DescribeSpotInstanceRequestsInput) error {
    return a.client.WaitUntilSpotInstanceRequestFulfilled(input)
}

func (a *EC2Activities) WaitUntilSubnetAvailable(input *ec2.DescribeSubnetsInput) error {
    return a.client.WaitUntilSubnetAvailable(input)
}

func (a *EC2Activities) WaitUntilSystemStatusOk(input *ec2.DescribeInstanceStatusInput) error {
    return a.client.WaitUntilSystemStatusOk(input)
}

func (a *EC2Activities) WaitUntilVolumeAvailable(input *ec2.DescribeVolumesInput) error {
    return a.client.WaitUntilVolumeAvailable(input)
}

func (a *EC2Activities) WaitUntilVolumeDeleted(input *ec2.DescribeVolumesInput) error {
    return a.client.WaitUntilVolumeDeleted(input)
}

func (a *EC2Activities) WaitUntilVolumeInUse(input *ec2.DescribeVolumesInput) error {
    return a.client.WaitUntilVolumeInUse(input)
}

func (a *EC2Activities) WaitUntilVpcAvailable(input *ec2.DescribeVpcsInput) error {
    return a.client.WaitUntilVpcAvailable(input)
}

func (a *EC2Activities) WaitUntilVpcExists(input *ec2.DescribeVpcsInput) error {
    return a.client.WaitUntilVpcExists(input)
}

func (a *EC2Activities) WaitUntilVpcPeeringConnectionDeleted(input *ec2.DescribeVpcPeeringConnectionsInput) error {
    return a.client.WaitUntilVpcPeeringConnectionDeleted(input)
}

func (a *EC2Activities) WaitUntilVpcPeeringConnectionExists(input *ec2.DescribeVpcPeeringConnectionsInput) error {
    return a.client.WaitUntilVpcPeeringConnectionExists(input)
}

func (a *EC2Activities) WaitUntilVpnConnectionAvailable(input *ec2.DescribeVpnConnectionsInput) error {
    return a.client.WaitUntilVpnConnectionAvailable(input)
}

func (a *EC2Activities) WaitUntilVpnConnectionDeleted(input *ec2.DescribeVpnConnectionsInput) error {
    return a.client.WaitUntilVpnConnectionDeleted(input)
}
