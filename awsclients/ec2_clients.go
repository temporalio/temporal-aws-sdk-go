package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type EC2Client interface {
    AcceptReservedInstancesExchangeQuote(ctx workflow.Context, input *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error)
    AcceptReservedInstancesExchangeQuoteAsync(ctx workflow.Context, input *ec2.AcceptReservedInstancesExchangeQuoteInput) *Ec2AcceptReservedInstancesExchangeQuoteResult

    AcceptTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.AcceptTransitGatewayPeeringAttachmentInput) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error)
    AcceptTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.AcceptTransitGatewayPeeringAttachmentInput) *Ec2AcceptTransitGatewayPeeringAttachmentResult

    AcceptTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.AcceptTransitGatewayVpcAttachmentInput) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error)
    AcceptTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.AcceptTransitGatewayVpcAttachmentInput) *Ec2AcceptTransitGatewayVpcAttachmentResult

    AcceptVpcEndpointConnections(ctx workflow.Context, input *ec2.AcceptVpcEndpointConnectionsInput) (*ec2.AcceptVpcEndpointConnectionsOutput, error)
    AcceptVpcEndpointConnectionsAsync(ctx workflow.Context, input *ec2.AcceptVpcEndpointConnectionsInput) *Ec2AcceptVpcEndpointConnectionsResult

    AcceptVpcPeeringConnection(ctx workflow.Context, input *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error)
    AcceptVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.AcceptVpcPeeringConnectionInput) *Ec2AcceptVpcPeeringConnectionResult

    AdvertiseByoipCidr(ctx workflow.Context, input *ec2.AdvertiseByoipCidrInput) (*ec2.AdvertiseByoipCidrOutput, error)
    AdvertiseByoipCidrAsync(ctx workflow.Context, input *ec2.AdvertiseByoipCidrInput) *Ec2AdvertiseByoipCidrResult

    AllocateAddress(ctx workflow.Context, input *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error)
    AllocateAddressAsync(ctx workflow.Context, input *ec2.AllocateAddressInput) *Ec2AllocateAddressResult

    AllocateHosts(ctx workflow.Context, input *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error)
    AllocateHostsAsync(ctx workflow.Context, input *ec2.AllocateHostsInput) *Ec2AllocateHostsResult

    ApplySecurityGroupsToClientVpnTargetNetwork(ctx workflow.Context, input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error)
    ApplySecurityGroupsToClientVpnTargetNetworkAsync(ctx workflow.Context, input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) *Ec2ApplySecurityGroupsToClientVpnTargetNetworkResult

    AssignIpv6Addresses(ctx workflow.Context, input *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error)
    AssignIpv6AddressesAsync(ctx workflow.Context, input *ec2.AssignIpv6AddressesInput) *Ec2AssignIpv6AddressesResult

    AssignPrivateIpAddresses(ctx workflow.Context, input *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error)
    AssignPrivateIpAddressesAsync(ctx workflow.Context, input *ec2.AssignPrivateIpAddressesInput) *Ec2AssignPrivateIpAddressesResult

    AssociateAddress(ctx workflow.Context, input *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error)
    AssociateAddressAsync(ctx workflow.Context, input *ec2.AssociateAddressInput) *Ec2AssociateAddressResult

    AssociateClientVpnTargetNetwork(ctx workflow.Context, input *ec2.AssociateClientVpnTargetNetworkInput) (*ec2.AssociateClientVpnTargetNetworkOutput, error)
    AssociateClientVpnTargetNetworkAsync(ctx workflow.Context, input *ec2.AssociateClientVpnTargetNetworkInput) *Ec2AssociateClientVpnTargetNetworkResult

    AssociateDhcpOptions(ctx workflow.Context, input *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error)
    AssociateDhcpOptionsAsync(ctx workflow.Context, input *ec2.AssociateDhcpOptionsInput) *Ec2AssociateDhcpOptionsResult

    AssociateIamInstanceProfile(ctx workflow.Context, input *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error)
    AssociateIamInstanceProfileAsync(ctx workflow.Context, input *ec2.AssociateIamInstanceProfileInput) *Ec2AssociateIamInstanceProfileResult

    AssociateRouteTable(ctx workflow.Context, input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error)
    AssociateRouteTableAsync(ctx workflow.Context, input *ec2.AssociateRouteTableInput) *Ec2AssociateRouteTableResult

    AssociateSubnetCidrBlock(ctx workflow.Context, input *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error)
    AssociateSubnetCidrBlockAsync(ctx workflow.Context, input *ec2.AssociateSubnetCidrBlockInput) *Ec2AssociateSubnetCidrBlockResult

    AssociateTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.AssociateTransitGatewayMulticastDomainInput) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error)
    AssociateTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.AssociateTransitGatewayMulticastDomainInput) *Ec2AssociateTransitGatewayMulticastDomainResult

    AssociateTransitGatewayRouteTable(ctx workflow.Context, input *ec2.AssociateTransitGatewayRouteTableInput) (*ec2.AssociateTransitGatewayRouteTableOutput, error)
    AssociateTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.AssociateTransitGatewayRouteTableInput) *Ec2AssociateTransitGatewayRouteTableResult

    AssociateVpcCidrBlock(ctx workflow.Context, input *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error)
    AssociateVpcCidrBlockAsync(ctx workflow.Context, input *ec2.AssociateVpcCidrBlockInput) *Ec2AssociateVpcCidrBlockResult

    AttachClassicLinkVpc(ctx workflow.Context, input *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error)
    AttachClassicLinkVpcAsync(ctx workflow.Context, input *ec2.AttachClassicLinkVpcInput) *Ec2AttachClassicLinkVpcResult

    AttachInternetGateway(ctx workflow.Context, input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error)
    AttachInternetGatewayAsync(ctx workflow.Context, input *ec2.AttachInternetGatewayInput) *Ec2AttachInternetGatewayResult

    AttachNetworkInterface(ctx workflow.Context, input *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error)
    AttachNetworkInterfaceAsync(ctx workflow.Context, input *ec2.AttachNetworkInterfaceInput) *Ec2AttachNetworkInterfaceResult

    AttachVolume(ctx workflow.Context, input *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error)
    AttachVolumeAsync(ctx workflow.Context, input *ec2.AttachVolumeInput) *Ec2AttachVolumeResult

    AttachVpnGateway(ctx workflow.Context, input *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error)
    AttachVpnGatewayAsync(ctx workflow.Context, input *ec2.AttachVpnGatewayInput) *Ec2AttachVpnGatewayResult

    AuthorizeClientVpnIngress(ctx workflow.Context, input *ec2.AuthorizeClientVpnIngressInput) (*ec2.AuthorizeClientVpnIngressOutput, error)
    AuthorizeClientVpnIngressAsync(ctx workflow.Context, input *ec2.AuthorizeClientVpnIngressInput) *Ec2AuthorizeClientVpnIngressResult

    AuthorizeSecurityGroupEgress(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error)
    AuthorizeSecurityGroupEgressAsync(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupEgressInput) *Ec2AuthorizeSecurityGroupEgressResult

    AuthorizeSecurityGroupIngress(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error)
    AuthorizeSecurityGroupIngressAsync(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupIngressInput) *Ec2AuthorizeSecurityGroupIngressResult

    BundleInstance(ctx workflow.Context, input *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error)
    BundleInstanceAsync(ctx workflow.Context, input *ec2.BundleInstanceInput) *Ec2BundleInstanceResult

    CancelBundleTask(ctx workflow.Context, input *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error)
    CancelBundleTaskAsync(ctx workflow.Context, input *ec2.CancelBundleTaskInput) *Ec2CancelBundleTaskResult

    CancelCapacityReservation(ctx workflow.Context, input *ec2.CancelCapacityReservationInput) (*ec2.CancelCapacityReservationOutput, error)
    CancelCapacityReservationAsync(ctx workflow.Context, input *ec2.CancelCapacityReservationInput) *Ec2CancelCapacityReservationResult

    CancelConversionTask(ctx workflow.Context, input *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error)
    CancelConversionTaskAsync(ctx workflow.Context, input *ec2.CancelConversionTaskInput) *Ec2CancelConversionTaskResult

    CancelExportTask(ctx workflow.Context, input *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error)
    CancelExportTaskAsync(ctx workflow.Context, input *ec2.CancelExportTaskInput) *Ec2CancelExportTaskResult

    CancelImportTask(ctx workflow.Context, input *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error)
    CancelImportTaskAsync(ctx workflow.Context, input *ec2.CancelImportTaskInput) *Ec2CancelImportTaskResult

    CancelReservedInstancesListing(ctx workflow.Context, input *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error)
    CancelReservedInstancesListingAsync(ctx workflow.Context, input *ec2.CancelReservedInstancesListingInput) *Ec2CancelReservedInstancesListingResult

    CancelSpotFleetRequests(ctx workflow.Context, input *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error)
    CancelSpotFleetRequestsAsync(ctx workflow.Context, input *ec2.CancelSpotFleetRequestsInput) *Ec2CancelSpotFleetRequestsResult

    CancelSpotInstanceRequests(ctx workflow.Context, input *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error)
    CancelSpotInstanceRequestsAsync(ctx workflow.Context, input *ec2.CancelSpotInstanceRequestsInput) *Ec2CancelSpotInstanceRequestsResult

    ConfirmProductInstance(ctx workflow.Context, input *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error)
    ConfirmProductInstanceAsync(ctx workflow.Context, input *ec2.ConfirmProductInstanceInput) *Ec2ConfirmProductInstanceResult

    CopyFpgaImage(ctx workflow.Context, input *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error)
    CopyFpgaImageAsync(ctx workflow.Context, input *ec2.CopyFpgaImageInput) *Ec2CopyFpgaImageResult

    CopyImage(ctx workflow.Context, input *ec2.CopyImageInput) (*ec2.CopyImageOutput, error)
    CopyImageAsync(ctx workflow.Context, input *ec2.CopyImageInput) *Ec2CopyImageResult

    CopySnapshot(ctx workflow.Context, input *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error)
    CopySnapshotAsync(ctx workflow.Context, input *ec2.CopySnapshotInput) *Ec2CopySnapshotResult

    CreateCapacityReservation(ctx workflow.Context, input *ec2.CreateCapacityReservationInput) (*ec2.CreateCapacityReservationOutput, error)
    CreateCapacityReservationAsync(ctx workflow.Context, input *ec2.CreateCapacityReservationInput) *Ec2CreateCapacityReservationResult

    CreateCarrierGateway(ctx workflow.Context, input *ec2.CreateCarrierGatewayInput) (*ec2.CreateCarrierGatewayOutput, error)
    CreateCarrierGatewayAsync(ctx workflow.Context, input *ec2.CreateCarrierGatewayInput) *Ec2CreateCarrierGatewayResult

    CreateClientVpnEndpoint(ctx workflow.Context, input *ec2.CreateClientVpnEndpointInput) (*ec2.CreateClientVpnEndpointOutput, error)
    CreateClientVpnEndpointAsync(ctx workflow.Context, input *ec2.CreateClientVpnEndpointInput) *Ec2CreateClientVpnEndpointResult

    CreateClientVpnRoute(ctx workflow.Context, input *ec2.CreateClientVpnRouteInput) (*ec2.CreateClientVpnRouteOutput, error)
    CreateClientVpnRouteAsync(ctx workflow.Context, input *ec2.CreateClientVpnRouteInput) *Ec2CreateClientVpnRouteResult

    CreateCustomerGateway(ctx workflow.Context, input *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error)
    CreateCustomerGatewayAsync(ctx workflow.Context, input *ec2.CreateCustomerGatewayInput) *Ec2CreateCustomerGatewayResult

    CreateDefaultSubnet(ctx workflow.Context, input *ec2.CreateDefaultSubnetInput) (*ec2.CreateDefaultSubnetOutput, error)
    CreateDefaultSubnetAsync(ctx workflow.Context, input *ec2.CreateDefaultSubnetInput) *Ec2CreateDefaultSubnetResult

    CreateDefaultVpc(ctx workflow.Context, input *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error)
    CreateDefaultVpcAsync(ctx workflow.Context, input *ec2.CreateDefaultVpcInput) *Ec2CreateDefaultVpcResult

    CreateDhcpOptions(ctx workflow.Context, input *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error)
    CreateDhcpOptionsAsync(ctx workflow.Context, input *ec2.CreateDhcpOptionsInput) *Ec2CreateDhcpOptionsResult

    CreateEgressOnlyInternetGateway(ctx workflow.Context, input *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error)
    CreateEgressOnlyInternetGatewayAsync(ctx workflow.Context, input *ec2.CreateEgressOnlyInternetGatewayInput) *Ec2CreateEgressOnlyInternetGatewayResult

    CreateFleet(ctx workflow.Context, input *ec2.CreateFleetInput) (*ec2.CreateFleetOutput, error)
    CreateFleetAsync(ctx workflow.Context, input *ec2.CreateFleetInput) *Ec2CreateFleetResult

    CreateFlowLogs(ctx workflow.Context, input *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error)
    CreateFlowLogsAsync(ctx workflow.Context, input *ec2.CreateFlowLogsInput) *Ec2CreateFlowLogsResult

    CreateFpgaImage(ctx workflow.Context, input *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error)
    CreateFpgaImageAsync(ctx workflow.Context, input *ec2.CreateFpgaImageInput) *Ec2CreateFpgaImageResult

    CreateImage(ctx workflow.Context, input *ec2.CreateImageInput) (*ec2.CreateImageOutput, error)
    CreateImageAsync(ctx workflow.Context, input *ec2.CreateImageInput) *Ec2CreateImageResult

    CreateInstanceExportTask(ctx workflow.Context, input *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error)
    CreateInstanceExportTaskAsync(ctx workflow.Context, input *ec2.CreateInstanceExportTaskInput) *Ec2CreateInstanceExportTaskResult

    CreateInternetGateway(ctx workflow.Context, input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error)
    CreateInternetGatewayAsync(ctx workflow.Context, input *ec2.CreateInternetGatewayInput) *Ec2CreateInternetGatewayResult

    CreateKeyPair(ctx workflow.Context, input *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error)
    CreateKeyPairAsync(ctx workflow.Context, input *ec2.CreateKeyPairInput) *Ec2CreateKeyPairResult

    CreateLaunchTemplate(ctx workflow.Context, input *ec2.CreateLaunchTemplateInput) (*ec2.CreateLaunchTemplateOutput, error)
    CreateLaunchTemplateAsync(ctx workflow.Context, input *ec2.CreateLaunchTemplateInput) *Ec2CreateLaunchTemplateResult

    CreateLaunchTemplateVersion(ctx workflow.Context, input *ec2.CreateLaunchTemplateVersionInput) (*ec2.CreateLaunchTemplateVersionOutput, error)
    CreateLaunchTemplateVersionAsync(ctx workflow.Context, input *ec2.CreateLaunchTemplateVersionInput) *Ec2CreateLaunchTemplateVersionResult

    CreateLocalGatewayRoute(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteInput) (*ec2.CreateLocalGatewayRouteOutput, error)
    CreateLocalGatewayRouteAsync(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteInput) *Ec2CreateLocalGatewayRouteResult

    CreateLocalGatewayRouteTableVpcAssociation(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error)
    CreateLocalGatewayRouteTableVpcAssociationAsync(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) *Ec2CreateLocalGatewayRouteTableVpcAssociationResult

    CreateManagedPrefixList(ctx workflow.Context, input *ec2.CreateManagedPrefixListInput) (*ec2.CreateManagedPrefixListOutput, error)
    CreateManagedPrefixListAsync(ctx workflow.Context, input *ec2.CreateManagedPrefixListInput) *Ec2CreateManagedPrefixListResult

    CreateNatGateway(ctx workflow.Context, input *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error)
    CreateNatGatewayAsync(ctx workflow.Context, input *ec2.CreateNatGatewayInput) *Ec2CreateNatGatewayResult

    CreateNetworkAcl(ctx workflow.Context, input *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error)
    CreateNetworkAclAsync(ctx workflow.Context, input *ec2.CreateNetworkAclInput) *Ec2CreateNetworkAclResult

    CreateNetworkAclEntry(ctx workflow.Context, input *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error)
    CreateNetworkAclEntryAsync(ctx workflow.Context, input *ec2.CreateNetworkAclEntryInput) *Ec2CreateNetworkAclEntryResult

    CreateNetworkInterface(ctx workflow.Context, input *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error)
    CreateNetworkInterfaceAsync(ctx workflow.Context, input *ec2.CreateNetworkInterfaceInput) *Ec2CreateNetworkInterfaceResult

    CreateNetworkInterfacePermission(ctx workflow.Context, input *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error)
    CreateNetworkInterfacePermissionAsync(ctx workflow.Context, input *ec2.CreateNetworkInterfacePermissionInput) *Ec2CreateNetworkInterfacePermissionResult

    CreatePlacementGroup(ctx workflow.Context, input *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error)
    CreatePlacementGroupAsync(ctx workflow.Context, input *ec2.CreatePlacementGroupInput) *Ec2CreatePlacementGroupResult

    CreateReservedInstancesListing(ctx workflow.Context, input *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error)
    CreateReservedInstancesListingAsync(ctx workflow.Context, input *ec2.CreateReservedInstancesListingInput) *Ec2CreateReservedInstancesListingResult

    CreateRoute(ctx workflow.Context, input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error)
    CreateRouteAsync(ctx workflow.Context, input *ec2.CreateRouteInput) *Ec2CreateRouteResult

    CreateRouteTable(ctx workflow.Context, input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error)
    CreateRouteTableAsync(ctx workflow.Context, input *ec2.CreateRouteTableInput) *Ec2CreateRouteTableResult

    CreateSecurityGroup(ctx workflow.Context, input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error)
    CreateSecurityGroupAsync(ctx workflow.Context, input *ec2.CreateSecurityGroupInput) *Ec2CreateSecurityGroupResult

    CreateSnapshot(ctx workflow.Context, input *ec2.CreateSnapshotInput) (*ec2.Snapshot, error)
    CreateSnapshotAsync(ctx workflow.Context, input *ec2.CreateSnapshotInput) *Ec2CreateSnapshotResult

    CreateSnapshots(ctx workflow.Context, input *ec2.CreateSnapshotsInput) (*ec2.CreateSnapshotsOutput, error)
    CreateSnapshotsAsync(ctx workflow.Context, input *ec2.CreateSnapshotsInput) *Ec2CreateSnapshotsResult

    CreateSpotDatafeedSubscription(ctx workflow.Context, input *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error)
    CreateSpotDatafeedSubscriptionAsync(ctx workflow.Context, input *ec2.CreateSpotDatafeedSubscriptionInput) *Ec2CreateSpotDatafeedSubscriptionResult

    CreateSubnet(ctx workflow.Context, input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error)
    CreateSubnetAsync(ctx workflow.Context, input *ec2.CreateSubnetInput) *Ec2CreateSubnetResult

    CreateTags(ctx workflow.Context, input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error)
    CreateTagsAsync(ctx workflow.Context, input *ec2.CreateTagsInput) *Ec2CreateTagsResult

    CreateTrafficMirrorFilter(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterInput) (*ec2.CreateTrafficMirrorFilterOutput, error)
    CreateTrafficMirrorFilterAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterInput) *Ec2CreateTrafficMirrorFilterResult

    CreateTrafficMirrorFilterRule(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterRuleInput) (*ec2.CreateTrafficMirrorFilterRuleOutput, error)
    CreateTrafficMirrorFilterRuleAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterRuleInput) *Ec2CreateTrafficMirrorFilterRuleResult

    CreateTrafficMirrorSession(ctx workflow.Context, input *ec2.CreateTrafficMirrorSessionInput) (*ec2.CreateTrafficMirrorSessionOutput, error)
    CreateTrafficMirrorSessionAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorSessionInput) *Ec2CreateTrafficMirrorSessionResult

    CreateTrafficMirrorTarget(ctx workflow.Context, input *ec2.CreateTrafficMirrorTargetInput) (*ec2.CreateTrafficMirrorTargetOutput, error)
    CreateTrafficMirrorTargetAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorTargetInput) *Ec2CreateTrafficMirrorTargetResult

    CreateTransitGateway(ctx workflow.Context, input *ec2.CreateTransitGatewayInput) (*ec2.CreateTransitGatewayOutput, error)
    CreateTransitGatewayAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayInput) *Ec2CreateTransitGatewayResult

    CreateTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.CreateTransitGatewayMulticastDomainInput) (*ec2.CreateTransitGatewayMulticastDomainOutput, error)
    CreateTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayMulticastDomainInput) *Ec2CreateTransitGatewayMulticastDomainResult

    CreateTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.CreateTransitGatewayPeeringAttachmentInput) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error)
    CreateTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayPeeringAttachmentInput) *Ec2CreateTransitGatewayPeeringAttachmentResult

    CreateTransitGatewayPrefixListReference(ctx workflow.Context, input *ec2.CreateTransitGatewayPrefixListReferenceInput) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error)
    CreateTransitGatewayPrefixListReferenceAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayPrefixListReferenceInput) *Ec2CreateTransitGatewayPrefixListReferenceResult

    CreateTransitGatewayRoute(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteInput) (*ec2.CreateTransitGatewayRouteOutput, error)
    CreateTransitGatewayRouteAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteInput) *Ec2CreateTransitGatewayRouteResult

    CreateTransitGatewayRouteTable(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteTableInput) (*ec2.CreateTransitGatewayRouteTableOutput, error)
    CreateTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteTableInput) *Ec2CreateTransitGatewayRouteTableResult

    CreateTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.CreateTransitGatewayVpcAttachmentInput) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error)
    CreateTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayVpcAttachmentInput) *Ec2CreateTransitGatewayVpcAttachmentResult

    CreateVolume(ctx workflow.Context, input *ec2.CreateVolumeInput) (*ec2.Volume, error)
    CreateVolumeAsync(ctx workflow.Context, input *ec2.CreateVolumeInput) *Ec2CreateVolumeResult

    CreateVpc(ctx workflow.Context, input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error)
    CreateVpcAsync(ctx workflow.Context, input *ec2.CreateVpcInput) *Ec2CreateVpcResult

    CreateVpcEndpoint(ctx workflow.Context, input *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error)
    CreateVpcEndpointAsync(ctx workflow.Context, input *ec2.CreateVpcEndpointInput) *Ec2CreateVpcEndpointResult

    CreateVpcEndpointConnectionNotification(ctx workflow.Context, input *ec2.CreateVpcEndpointConnectionNotificationInput) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error)
    CreateVpcEndpointConnectionNotificationAsync(ctx workflow.Context, input *ec2.CreateVpcEndpointConnectionNotificationInput) *Ec2CreateVpcEndpointConnectionNotificationResult

    CreateVpcEndpointServiceConfiguration(ctx workflow.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error)
    CreateVpcEndpointServiceConfigurationAsync(ctx workflow.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput) *Ec2CreateVpcEndpointServiceConfigurationResult

    CreateVpcPeeringConnection(ctx workflow.Context, input *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error)
    CreateVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.CreateVpcPeeringConnectionInput) *Ec2CreateVpcPeeringConnectionResult

    CreateVpnConnection(ctx workflow.Context, input *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error)
    CreateVpnConnectionAsync(ctx workflow.Context, input *ec2.CreateVpnConnectionInput) *Ec2CreateVpnConnectionResult

    CreateVpnConnectionRoute(ctx workflow.Context, input *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error)
    CreateVpnConnectionRouteAsync(ctx workflow.Context, input *ec2.CreateVpnConnectionRouteInput) *Ec2CreateVpnConnectionRouteResult

    CreateVpnGateway(ctx workflow.Context, input *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error)
    CreateVpnGatewayAsync(ctx workflow.Context, input *ec2.CreateVpnGatewayInput) *Ec2CreateVpnGatewayResult

    DeleteCarrierGateway(ctx workflow.Context, input *ec2.DeleteCarrierGatewayInput) (*ec2.DeleteCarrierGatewayOutput, error)
    DeleteCarrierGatewayAsync(ctx workflow.Context, input *ec2.DeleteCarrierGatewayInput) *Ec2DeleteCarrierGatewayResult

    DeleteClientVpnEndpoint(ctx workflow.Context, input *ec2.DeleteClientVpnEndpointInput) (*ec2.DeleteClientVpnEndpointOutput, error)
    DeleteClientVpnEndpointAsync(ctx workflow.Context, input *ec2.DeleteClientVpnEndpointInput) *Ec2DeleteClientVpnEndpointResult

    DeleteClientVpnRoute(ctx workflow.Context, input *ec2.DeleteClientVpnRouteInput) (*ec2.DeleteClientVpnRouteOutput, error)
    DeleteClientVpnRouteAsync(ctx workflow.Context, input *ec2.DeleteClientVpnRouteInput) *Ec2DeleteClientVpnRouteResult

    DeleteCustomerGateway(ctx workflow.Context, input *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error)
    DeleteCustomerGatewayAsync(ctx workflow.Context, input *ec2.DeleteCustomerGatewayInput) *Ec2DeleteCustomerGatewayResult

    DeleteDhcpOptions(ctx workflow.Context, input *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error)
    DeleteDhcpOptionsAsync(ctx workflow.Context, input *ec2.DeleteDhcpOptionsInput) *Ec2DeleteDhcpOptionsResult

    DeleteEgressOnlyInternetGateway(ctx workflow.Context, input *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error)
    DeleteEgressOnlyInternetGatewayAsync(ctx workflow.Context, input *ec2.DeleteEgressOnlyInternetGatewayInput) *Ec2DeleteEgressOnlyInternetGatewayResult

    DeleteFleets(ctx workflow.Context, input *ec2.DeleteFleetsInput) (*ec2.DeleteFleetsOutput, error)
    DeleteFleetsAsync(ctx workflow.Context, input *ec2.DeleteFleetsInput) *Ec2DeleteFleetsResult

    DeleteFlowLogs(ctx workflow.Context, input *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error)
    DeleteFlowLogsAsync(ctx workflow.Context, input *ec2.DeleteFlowLogsInput) *Ec2DeleteFlowLogsResult

    DeleteFpgaImage(ctx workflow.Context, input *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error)
    DeleteFpgaImageAsync(ctx workflow.Context, input *ec2.DeleteFpgaImageInput) *Ec2DeleteFpgaImageResult

    DeleteInternetGateway(ctx workflow.Context, input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error)
    DeleteInternetGatewayAsync(ctx workflow.Context, input *ec2.DeleteInternetGatewayInput) *Ec2DeleteInternetGatewayResult

    DeleteKeyPair(ctx workflow.Context, input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error)
    DeleteKeyPairAsync(ctx workflow.Context, input *ec2.DeleteKeyPairInput) *Ec2DeleteKeyPairResult

    DeleteLaunchTemplate(ctx workflow.Context, input *ec2.DeleteLaunchTemplateInput) (*ec2.DeleteLaunchTemplateOutput, error)
    DeleteLaunchTemplateAsync(ctx workflow.Context, input *ec2.DeleteLaunchTemplateInput) *Ec2DeleteLaunchTemplateResult

    DeleteLaunchTemplateVersions(ctx workflow.Context, input *ec2.DeleteLaunchTemplateVersionsInput) (*ec2.DeleteLaunchTemplateVersionsOutput, error)
    DeleteLaunchTemplateVersionsAsync(ctx workflow.Context, input *ec2.DeleteLaunchTemplateVersionsInput) *Ec2DeleteLaunchTemplateVersionsResult

    DeleteLocalGatewayRoute(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteInput) (*ec2.DeleteLocalGatewayRouteOutput, error)
    DeleteLocalGatewayRouteAsync(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteInput) *Ec2DeleteLocalGatewayRouteResult

    DeleteLocalGatewayRouteTableVpcAssociation(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error)
    DeleteLocalGatewayRouteTableVpcAssociationAsync(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) *Ec2DeleteLocalGatewayRouteTableVpcAssociationResult

    DeleteManagedPrefixList(ctx workflow.Context, input *ec2.DeleteManagedPrefixListInput) (*ec2.DeleteManagedPrefixListOutput, error)
    DeleteManagedPrefixListAsync(ctx workflow.Context, input *ec2.DeleteManagedPrefixListInput) *Ec2DeleteManagedPrefixListResult

    DeleteNatGateway(ctx workflow.Context, input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error)
    DeleteNatGatewayAsync(ctx workflow.Context, input *ec2.DeleteNatGatewayInput) *Ec2DeleteNatGatewayResult

    DeleteNetworkAcl(ctx workflow.Context, input *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error)
    DeleteNetworkAclAsync(ctx workflow.Context, input *ec2.DeleteNetworkAclInput) *Ec2DeleteNetworkAclResult

    DeleteNetworkAclEntry(ctx workflow.Context, input *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error)
    DeleteNetworkAclEntryAsync(ctx workflow.Context, input *ec2.DeleteNetworkAclEntryInput) *Ec2DeleteNetworkAclEntryResult

    DeleteNetworkInterface(ctx workflow.Context, input *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error)
    DeleteNetworkInterfaceAsync(ctx workflow.Context, input *ec2.DeleteNetworkInterfaceInput) *Ec2DeleteNetworkInterfaceResult

    DeleteNetworkInterfacePermission(ctx workflow.Context, input *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error)
    DeleteNetworkInterfacePermissionAsync(ctx workflow.Context, input *ec2.DeleteNetworkInterfacePermissionInput) *Ec2DeleteNetworkInterfacePermissionResult

    DeletePlacementGroup(ctx workflow.Context, input *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error)
    DeletePlacementGroupAsync(ctx workflow.Context, input *ec2.DeletePlacementGroupInput) *Ec2DeletePlacementGroupResult

    DeleteQueuedReservedInstances(ctx workflow.Context, input *ec2.DeleteQueuedReservedInstancesInput) (*ec2.DeleteQueuedReservedInstancesOutput, error)
    DeleteQueuedReservedInstancesAsync(ctx workflow.Context, input *ec2.DeleteQueuedReservedInstancesInput) *Ec2DeleteQueuedReservedInstancesResult

    DeleteRoute(ctx workflow.Context, input *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error)
    DeleteRouteAsync(ctx workflow.Context, input *ec2.DeleteRouteInput) *Ec2DeleteRouteResult

    DeleteRouteTable(ctx workflow.Context, input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error)
    DeleteRouteTableAsync(ctx workflow.Context, input *ec2.DeleteRouteTableInput) *Ec2DeleteRouteTableResult

    DeleteSecurityGroup(ctx workflow.Context, input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error)
    DeleteSecurityGroupAsync(ctx workflow.Context, input *ec2.DeleteSecurityGroupInput) *Ec2DeleteSecurityGroupResult

    DeleteSnapshot(ctx workflow.Context, input *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error)
    DeleteSnapshotAsync(ctx workflow.Context, input *ec2.DeleteSnapshotInput) *Ec2DeleteSnapshotResult

    DeleteSpotDatafeedSubscription(ctx workflow.Context, input *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error)
    DeleteSpotDatafeedSubscriptionAsync(ctx workflow.Context, input *ec2.DeleteSpotDatafeedSubscriptionInput) *Ec2DeleteSpotDatafeedSubscriptionResult

    DeleteSubnet(ctx workflow.Context, input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error)
    DeleteSubnetAsync(ctx workflow.Context, input *ec2.DeleteSubnetInput) *Ec2DeleteSubnetResult

    DeleteTags(ctx workflow.Context, input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error)
    DeleteTagsAsync(ctx workflow.Context, input *ec2.DeleteTagsInput) *Ec2DeleteTagsResult

    DeleteTrafficMirrorFilter(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterInput) (*ec2.DeleteTrafficMirrorFilterOutput, error)
    DeleteTrafficMirrorFilterAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterInput) *Ec2DeleteTrafficMirrorFilterResult

    DeleteTrafficMirrorFilterRule(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterRuleInput) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error)
    DeleteTrafficMirrorFilterRuleAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterRuleInput) *Ec2DeleteTrafficMirrorFilterRuleResult

    DeleteTrafficMirrorSession(ctx workflow.Context, input *ec2.DeleteTrafficMirrorSessionInput) (*ec2.DeleteTrafficMirrorSessionOutput, error)
    DeleteTrafficMirrorSessionAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorSessionInput) *Ec2DeleteTrafficMirrorSessionResult

    DeleteTrafficMirrorTarget(ctx workflow.Context, input *ec2.DeleteTrafficMirrorTargetInput) (*ec2.DeleteTrafficMirrorTargetOutput, error)
    DeleteTrafficMirrorTargetAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorTargetInput) *Ec2DeleteTrafficMirrorTargetResult

    DeleteTransitGateway(ctx workflow.Context, input *ec2.DeleteTransitGatewayInput) (*ec2.DeleteTransitGatewayOutput, error)
    DeleteTransitGatewayAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayInput) *Ec2DeleteTransitGatewayResult

    DeleteTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.DeleteTransitGatewayMulticastDomainInput) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error)
    DeleteTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayMulticastDomainInput) *Ec2DeleteTransitGatewayMulticastDomainResult

    DeleteTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.DeleteTransitGatewayPeeringAttachmentInput) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error)
    DeleteTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayPeeringAttachmentInput) *Ec2DeleteTransitGatewayPeeringAttachmentResult

    DeleteTransitGatewayPrefixListReference(ctx workflow.Context, input *ec2.DeleteTransitGatewayPrefixListReferenceInput) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error)
    DeleteTransitGatewayPrefixListReferenceAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayPrefixListReferenceInput) *Ec2DeleteTransitGatewayPrefixListReferenceResult

    DeleteTransitGatewayRoute(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteInput) (*ec2.DeleteTransitGatewayRouteOutput, error)
    DeleteTransitGatewayRouteAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteInput) *Ec2DeleteTransitGatewayRouteResult

    DeleteTransitGatewayRouteTable(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteTableInput) (*ec2.DeleteTransitGatewayRouteTableOutput, error)
    DeleteTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteTableInput) *Ec2DeleteTransitGatewayRouteTableResult

    DeleteTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.DeleteTransitGatewayVpcAttachmentInput) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error)
    DeleteTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayVpcAttachmentInput) *Ec2DeleteTransitGatewayVpcAttachmentResult

    DeleteVolume(ctx workflow.Context, input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error)
    DeleteVolumeAsync(ctx workflow.Context, input *ec2.DeleteVolumeInput) *Ec2DeleteVolumeResult

    DeleteVpc(ctx workflow.Context, input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error)
    DeleteVpcAsync(ctx workflow.Context, input *ec2.DeleteVpcInput) *Ec2DeleteVpcResult

    DeleteVpcEndpointConnectionNotifications(ctx workflow.Context, input *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error)
    DeleteVpcEndpointConnectionNotificationsAsync(ctx workflow.Context, input *ec2.DeleteVpcEndpointConnectionNotificationsInput) *Ec2DeleteVpcEndpointConnectionNotificationsResult

    DeleteVpcEndpointServiceConfigurations(ctx workflow.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error)
    DeleteVpcEndpointServiceConfigurationsAsync(ctx workflow.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput) *Ec2DeleteVpcEndpointServiceConfigurationsResult

    DeleteVpcEndpoints(ctx workflow.Context, input *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error)
    DeleteVpcEndpointsAsync(ctx workflow.Context, input *ec2.DeleteVpcEndpointsInput) *Ec2DeleteVpcEndpointsResult

    DeleteVpcPeeringConnection(ctx workflow.Context, input *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error)
    DeleteVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.DeleteVpcPeeringConnectionInput) *Ec2DeleteVpcPeeringConnectionResult

    DeleteVpnConnection(ctx workflow.Context, input *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error)
    DeleteVpnConnectionAsync(ctx workflow.Context, input *ec2.DeleteVpnConnectionInput) *Ec2DeleteVpnConnectionResult

    DeleteVpnConnectionRoute(ctx workflow.Context, input *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error)
    DeleteVpnConnectionRouteAsync(ctx workflow.Context, input *ec2.DeleteVpnConnectionRouteInput) *Ec2DeleteVpnConnectionRouteResult

    DeleteVpnGateway(ctx workflow.Context, input *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error)
    DeleteVpnGatewayAsync(ctx workflow.Context, input *ec2.DeleteVpnGatewayInput) *Ec2DeleteVpnGatewayResult

    DeprovisionByoipCidr(ctx workflow.Context, input *ec2.DeprovisionByoipCidrInput) (*ec2.DeprovisionByoipCidrOutput, error)
    DeprovisionByoipCidrAsync(ctx workflow.Context, input *ec2.DeprovisionByoipCidrInput) *Ec2DeprovisionByoipCidrResult

    DeregisterImage(ctx workflow.Context, input *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error)
    DeregisterImageAsync(ctx workflow.Context, input *ec2.DeregisterImageInput) *Ec2DeregisterImageResult

    DeregisterInstanceEventNotificationAttributes(ctx workflow.Context, input *ec2.DeregisterInstanceEventNotificationAttributesInput) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error)
    DeregisterInstanceEventNotificationAttributesAsync(ctx workflow.Context, input *ec2.DeregisterInstanceEventNotificationAttributesInput) *Ec2DeregisterInstanceEventNotificationAttributesResult

    DeregisterTransitGatewayMulticastGroupMembers(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error)
    DeregisterTransitGatewayMulticastGroupMembersAsync(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) *Ec2DeregisterTransitGatewayMulticastGroupMembersResult

    DeregisterTransitGatewayMulticastGroupSources(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error)
    DeregisterTransitGatewayMulticastGroupSourcesAsync(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) *Ec2DeregisterTransitGatewayMulticastGroupSourcesResult

    DescribeAccountAttributes(ctx workflow.Context, input *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error)
    DescribeAccountAttributesAsync(ctx workflow.Context, input *ec2.DescribeAccountAttributesInput) *Ec2DescribeAccountAttributesResult

    DescribeAddresses(ctx workflow.Context, input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error)
    DescribeAddressesAsync(ctx workflow.Context, input *ec2.DescribeAddressesInput) *Ec2DescribeAddressesResult

    DescribeAggregateIdFormat(ctx workflow.Context, input *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error)
    DescribeAggregateIdFormatAsync(ctx workflow.Context, input *ec2.DescribeAggregateIdFormatInput) *Ec2DescribeAggregateIdFormatResult

    DescribeAvailabilityZones(ctx workflow.Context, input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error)
    DescribeAvailabilityZonesAsync(ctx workflow.Context, input *ec2.DescribeAvailabilityZonesInput) *Ec2DescribeAvailabilityZonesResult

    DescribeBundleTasks(ctx workflow.Context, input *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error)
    DescribeBundleTasksAsync(ctx workflow.Context, input *ec2.DescribeBundleTasksInput) *Ec2DescribeBundleTasksResult

    DescribeByoipCidrs(ctx workflow.Context, input *ec2.DescribeByoipCidrsInput) (*ec2.DescribeByoipCidrsOutput, error)
    DescribeByoipCidrsAsync(ctx workflow.Context, input *ec2.DescribeByoipCidrsInput) *Ec2DescribeByoipCidrsResult

    DescribeCapacityReservations(ctx workflow.Context, input *ec2.DescribeCapacityReservationsInput) (*ec2.DescribeCapacityReservationsOutput, error)
    DescribeCapacityReservationsAsync(ctx workflow.Context, input *ec2.DescribeCapacityReservationsInput) *Ec2DescribeCapacityReservationsResult

    DescribeCarrierGateways(ctx workflow.Context, input *ec2.DescribeCarrierGatewaysInput) (*ec2.DescribeCarrierGatewaysOutput, error)
    DescribeCarrierGatewaysAsync(ctx workflow.Context, input *ec2.DescribeCarrierGatewaysInput) *Ec2DescribeCarrierGatewaysResult

    DescribeClassicLinkInstances(ctx workflow.Context, input *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error)
    DescribeClassicLinkInstancesAsync(ctx workflow.Context, input *ec2.DescribeClassicLinkInstancesInput) *Ec2DescribeClassicLinkInstancesResult

    DescribeClientVpnAuthorizationRules(ctx workflow.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error)
    DescribeClientVpnAuthorizationRulesAsync(ctx workflow.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput) *Ec2DescribeClientVpnAuthorizationRulesResult

    DescribeClientVpnConnections(ctx workflow.Context, input *ec2.DescribeClientVpnConnectionsInput) (*ec2.DescribeClientVpnConnectionsOutput, error)
    DescribeClientVpnConnectionsAsync(ctx workflow.Context, input *ec2.DescribeClientVpnConnectionsInput) *Ec2DescribeClientVpnConnectionsResult

    DescribeClientVpnEndpoints(ctx workflow.Context, input *ec2.DescribeClientVpnEndpointsInput) (*ec2.DescribeClientVpnEndpointsOutput, error)
    DescribeClientVpnEndpointsAsync(ctx workflow.Context, input *ec2.DescribeClientVpnEndpointsInput) *Ec2DescribeClientVpnEndpointsResult

    DescribeClientVpnRoutes(ctx workflow.Context, input *ec2.DescribeClientVpnRoutesInput) (*ec2.DescribeClientVpnRoutesOutput, error)
    DescribeClientVpnRoutesAsync(ctx workflow.Context, input *ec2.DescribeClientVpnRoutesInput) *Ec2DescribeClientVpnRoutesResult

    DescribeClientVpnTargetNetworks(ctx workflow.Context, input *ec2.DescribeClientVpnTargetNetworksInput) (*ec2.DescribeClientVpnTargetNetworksOutput, error)
    DescribeClientVpnTargetNetworksAsync(ctx workflow.Context, input *ec2.DescribeClientVpnTargetNetworksInput) *Ec2DescribeClientVpnTargetNetworksResult

    DescribeCoipPools(ctx workflow.Context, input *ec2.DescribeCoipPoolsInput) (*ec2.DescribeCoipPoolsOutput, error)
    DescribeCoipPoolsAsync(ctx workflow.Context, input *ec2.DescribeCoipPoolsInput) *Ec2DescribeCoipPoolsResult

    DescribeConversionTasks(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error)
    DescribeConversionTasksAsync(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) *Ec2DescribeConversionTasksResult

    DescribeCustomerGateways(ctx workflow.Context, input *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error)
    DescribeCustomerGatewaysAsync(ctx workflow.Context, input *ec2.DescribeCustomerGatewaysInput) *Ec2DescribeCustomerGatewaysResult

    DescribeDhcpOptions(ctx workflow.Context, input *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error)
    DescribeDhcpOptionsAsync(ctx workflow.Context, input *ec2.DescribeDhcpOptionsInput) *Ec2DescribeDhcpOptionsResult

    DescribeEgressOnlyInternetGateways(ctx workflow.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
    DescribeEgressOnlyInternetGatewaysAsync(ctx workflow.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput) *Ec2DescribeEgressOnlyInternetGatewaysResult

    DescribeElasticGpus(ctx workflow.Context, input *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error)
    DescribeElasticGpusAsync(ctx workflow.Context, input *ec2.DescribeElasticGpusInput) *Ec2DescribeElasticGpusResult

    DescribeExportImageTasks(ctx workflow.Context, input *ec2.DescribeExportImageTasksInput) (*ec2.DescribeExportImageTasksOutput, error)
    DescribeExportImageTasksAsync(ctx workflow.Context, input *ec2.DescribeExportImageTasksInput) *Ec2DescribeExportImageTasksResult

    DescribeExportTasks(ctx workflow.Context, input *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error)
    DescribeExportTasksAsync(ctx workflow.Context, input *ec2.DescribeExportTasksInput) *Ec2DescribeExportTasksResult

    DescribeFastSnapshotRestores(ctx workflow.Context, input *ec2.DescribeFastSnapshotRestoresInput) (*ec2.DescribeFastSnapshotRestoresOutput, error)
    DescribeFastSnapshotRestoresAsync(ctx workflow.Context, input *ec2.DescribeFastSnapshotRestoresInput) *Ec2DescribeFastSnapshotRestoresResult

    DescribeFleetHistory(ctx workflow.Context, input *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error)
    DescribeFleetHistoryAsync(ctx workflow.Context, input *ec2.DescribeFleetHistoryInput) *Ec2DescribeFleetHistoryResult

    DescribeFleetInstances(ctx workflow.Context, input *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error)
    DescribeFleetInstancesAsync(ctx workflow.Context, input *ec2.DescribeFleetInstancesInput) *Ec2DescribeFleetInstancesResult

    DescribeFleets(ctx workflow.Context, input *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error)
    DescribeFleetsAsync(ctx workflow.Context, input *ec2.DescribeFleetsInput) *Ec2DescribeFleetsResult

    DescribeFlowLogs(ctx workflow.Context, input *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error)
    DescribeFlowLogsAsync(ctx workflow.Context, input *ec2.DescribeFlowLogsInput) *Ec2DescribeFlowLogsResult

    DescribeFpgaImageAttribute(ctx workflow.Context, input *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error)
    DescribeFpgaImageAttributeAsync(ctx workflow.Context, input *ec2.DescribeFpgaImageAttributeInput) *Ec2DescribeFpgaImageAttributeResult

    DescribeFpgaImages(ctx workflow.Context, input *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error)
    DescribeFpgaImagesAsync(ctx workflow.Context, input *ec2.DescribeFpgaImagesInput) *Ec2DescribeFpgaImagesResult

    DescribeHostReservationOfferings(ctx workflow.Context, input *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error)
    DescribeHostReservationOfferingsAsync(ctx workflow.Context, input *ec2.DescribeHostReservationOfferingsInput) *Ec2DescribeHostReservationOfferingsResult

    DescribeHostReservations(ctx workflow.Context, input *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error)
    DescribeHostReservationsAsync(ctx workflow.Context, input *ec2.DescribeHostReservationsInput) *Ec2DescribeHostReservationsResult

    DescribeHosts(ctx workflow.Context, input *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error)
    DescribeHostsAsync(ctx workflow.Context, input *ec2.DescribeHostsInput) *Ec2DescribeHostsResult

    DescribeIamInstanceProfileAssociations(ctx workflow.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error)
    DescribeIamInstanceProfileAssociationsAsync(ctx workflow.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput) *Ec2DescribeIamInstanceProfileAssociationsResult

    DescribeIdFormat(ctx workflow.Context, input *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error)
    DescribeIdFormatAsync(ctx workflow.Context, input *ec2.DescribeIdFormatInput) *Ec2DescribeIdFormatResult

    DescribeIdentityIdFormat(ctx workflow.Context, input *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error)
    DescribeIdentityIdFormatAsync(ctx workflow.Context, input *ec2.DescribeIdentityIdFormatInput) *Ec2DescribeIdentityIdFormatResult

    DescribeImageAttribute(ctx workflow.Context, input *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error)
    DescribeImageAttributeAsync(ctx workflow.Context, input *ec2.DescribeImageAttributeInput) *Ec2DescribeImageAttributeResult

    DescribeImages(ctx workflow.Context, input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error)
    DescribeImagesAsync(ctx workflow.Context, input *ec2.DescribeImagesInput) *Ec2DescribeImagesResult

    DescribeImportImageTasks(ctx workflow.Context, input *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error)
    DescribeImportImageTasksAsync(ctx workflow.Context, input *ec2.DescribeImportImageTasksInput) *Ec2DescribeImportImageTasksResult

    DescribeImportSnapshotTasks(ctx workflow.Context, input *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error)
    DescribeImportSnapshotTasksAsync(ctx workflow.Context, input *ec2.DescribeImportSnapshotTasksInput) *Ec2DescribeImportSnapshotTasksResult

    DescribeInstanceAttribute(ctx workflow.Context, input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error)
    DescribeInstanceAttributeAsync(ctx workflow.Context, input *ec2.DescribeInstanceAttributeInput) *Ec2DescribeInstanceAttributeResult

    DescribeInstanceCreditSpecifications(ctx workflow.Context, input *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error)
    DescribeInstanceCreditSpecificationsAsync(ctx workflow.Context, input *ec2.DescribeInstanceCreditSpecificationsInput) *Ec2DescribeInstanceCreditSpecificationsResult

    DescribeInstanceEventNotificationAttributes(ctx workflow.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error)
    DescribeInstanceEventNotificationAttributesAsync(ctx workflow.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput) *Ec2DescribeInstanceEventNotificationAttributesResult

    DescribeInstanceStatus(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error)
    DescribeInstanceStatusAsync(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) *Ec2DescribeInstanceStatusResult

    DescribeInstanceTypeOfferings(ctx workflow.Context, input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error)
    DescribeInstanceTypeOfferingsAsync(ctx workflow.Context, input *ec2.DescribeInstanceTypeOfferingsInput) *Ec2DescribeInstanceTypeOfferingsResult

    DescribeInstanceTypes(ctx workflow.Context, input *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error)
    DescribeInstanceTypesAsync(ctx workflow.Context, input *ec2.DescribeInstanceTypesInput) *Ec2DescribeInstanceTypesResult

    DescribeInstances(ctx workflow.Context, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)
    DescribeInstancesAsync(ctx workflow.Context, input *ec2.DescribeInstancesInput) *Ec2DescribeInstancesResult

    DescribeInternetGateways(ctx workflow.Context, input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error)
    DescribeInternetGatewaysAsync(ctx workflow.Context, input *ec2.DescribeInternetGatewaysInput) *Ec2DescribeInternetGatewaysResult

    DescribeIpv6Pools(ctx workflow.Context, input *ec2.DescribeIpv6PoolsInput) (*ec2.DescribeIpv6PoolsOutput, error)
    DescribeIpv6PoolsAsync(ctx workflow.Context, input *ec2.DescribeIpv6PoolsInput) *Ec2DescribeIpv6PoolsResult

    DescribeKeyPairs(ctx workflow.Context, input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error)
    DescribeKeyPairsAsync(ctx workflow.Context, input *ec2.DescribeKeyPairsInput) *Ec2DescribeKeyPairsResult

    DescribeLaunchTemplateVersions(ctx workflow.Context, input *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error)
    DescribeLaunchTemplateVersionsAsync(ctx workflow.Context, input *ec2.DescribeLaunchTemplateVersionsInput) *Ec2DescribeLaunchTemplateVersionsResult

    DescribeLaunchTemplates(ctx workflow.Context, input *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error)
    DescribeLaunchTemplatesAsync(ctx workflow.Context, input *ec2.DescribeLaunchTemplatesInput) *Ec2DescribeLaunchTemplatesResult

    DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error)
    DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) *Ec2DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResult

    DescribeLocalGatewayRouteTableVpcAssociations(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error)
    DescribeLocalGatewayRouteTableVpcAssociationsAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) *Ec2DescribeLocalGatewayRouteTableVpcAssociationsResult

    DescribeLocalGatewayRouteTables(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTablesInput) (*ec2.DescribeLocalGatewayRouteTablesOutput, error)
    DescribeLocalGatewayRouteTablesAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTablesInput) *Ec2DescribeLocalGatewayRouteTablesResult

    DescribeLocalGatewayVirtualInterfaceGroups(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error)
    DescribeLocalGatewayVirtualInterfaceGroupsAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) *Ec2DescribeLocalGatewayVirtualInterfaceGroupsResult

    DescribeLocalGatewayVirtualInterfaces(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error)
    DescribeLocalGatewayVirtualInterfacesAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput) *Ec2DescribeLocalGatewayVirtualInterfacesResult

    DescribeLocalGateways(ctx workflow.Context, input *ec2.DescribeLocalGatewaysInput) (*ec2.DescribeLocalGatewaysOutput, error)
    DescribeLocalGatewaysAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewaysInput) *Ec2DescribeLocalGatewaysResult

    DescribeManagedPrefixLists(ctx workflow.Context, input *ec2.DescribeManagedPrefixListsInput) (*ec2.DescribeManagedPrefixListsOutput, error)
    DescribeManagedPrefixListsAsync(ctx workflow.Context, input *ec2.DescribeManagedPrefixListsInput) *Ec2DescribeManagedPrefixListsResult

    DescribeMovingAddresses(ctx workflow.Context, input *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error)
    DescribeMovingAddressesAsync(ctx workflow.Context, input *ec2.DescribeMovingAddressesInput) *Ec2DescribeMovingAddressesResult

    DescribeNatGateways(ctx workflow.Context, input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error)
    DescribeNatGatewaysAsync(ctx workflow.Context, input *ec2.DescribeNatGatewaysInput) *Ec2DescribeNatGatewaysResult

    DescribeNetworkAcls(ctx workflow.Context, input *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error)
    DescribeNetworkAclsAsync(ctx workflow.Context, input *ec2.DescribeNetworkAclsInput) *Ec2DescribeNetworkAclsResult

    DescribeNetworkInterfaceAttribute(ctx workflow.Context, input *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)
    DescribeNetworkInterfaceAttributeAsync(ctx workflow.Context, input *ec2.DescribeNetworkInterfaceAttributeInput) *Ec2DescribeNetworkInterfaceAttributeResult

    DescribeNetworkInterfacePermissions(ctx workflow.Context, input *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error)
    DescribeNetworkInterfacePermissionsAsync(ctx workflow.Context, input *ec2.DescribeNetworkInterfacePermissionsInput) *Ec2DescribeNetworkInterfacePermissionsResult

    DescribeNetworkInterfaces(ctx workflow.Context, input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error)
    DescribeNetworkInterfacesAsync(ctx workflow.Context, input *ec2.DescribeNetworkInterfacesInput) *Ec2DescribeNetworkInterfacesResult

    DescribePlacementGroups(ctx workflow.Context, input *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error)
    DescribePlacementGroupsAsync(ctx workflow.Context, input *ec2.DescribePlacementGroupsInput) *Ec2DescribePlacementGroupsResult

    DescribePrefixLists(ctx workflow.Context, input *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error)
    DescribePrefixListsAsync(ctx workflow.Context, input *ec2.DescribePrefixListsInput) *Ec2DescribePrefixListsResult

    DescribePrincipalIdFormat(ctx workflow.Context, input *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error)
    DescribePrincipalIdFormatAsync(ctx workflow.Context, input *ec2.DescribePrincipalIdFormatInput) *Ec2DescribePrincipalIdFormatResult

    DescribePublicIpv4Pools(ctx workflow.Context, input *ec2.DescribePublicIpv4PoolsInput) (*ec2.DescribePublicIpv4PoolsOutput, error)
    DescribePublicIpv4PoolsAsync(ctx workflow.Context, input *ec2.DescribePublicIpv4PoolsInput) *Ec2DescribePublicIpv4PoolsResult

    DescribeRegions(ctx workflow.Context, input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error)
    DescribeRegionsAsync(ctx workflow.Context, input *ec2.DescribeRegionsInput) *Ec2DescribeRegionsResult

    DescribeReservedInstances(ctx workflow.Context, input *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error)
    DescribeReservedInstancesAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesInput) *Ec2DescribeReservedInstancesResult

    DescribeReservedInstancesListings(ctx workflow.Context, input *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error)
    DescribeReservedInstancesListingsAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesListingsInput) *Ec2DescribeReservedInstancesListingsResult

    DescribeReservedInstancesModifications(ctx workflow.Context, input *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error)
    DescribeReservedInstancesModificationsAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesModificationsInput) *Ec2DescribeReservedInstancesModificationsResult

    DescribeReservedInstancesOfferings(ctx workflow.Context, input *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
    DescribeReservedInstancesOfferingsAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesOfferingsInput) *Ec2DescribeReservedInstancesOfferingsResult

    DescribeRouteTables(ctx workflow.Context, input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error)
    DescribeRouteTablesAsync(ctx workflow.Context, input *ec2.DescribeRouteTablesInput) *Ec2DescribeRouteTablesResult

    DescribeScheduledInstanceAvailability(ctx workflow.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error)
    DescribeScheduledInstanceAvailabilityAsync(ctx workflow.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput) *Ec2DescribeScheduledInstanceAvailabilityResult

    DescribeScheduledInstances(ctx workflow.Context, input *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error)
    DescribeScheduledInstancesAsync(ctx workflow.Context, input *ec2.DescribeScheduledInstancesInput) *Ec2DescribeScheduledInstancesResult

    DescribeSecurityGroupReferences(ctx workflow.Context, input *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error)
    DescribeSecurityGroupReferencesAsync(ctx workflow.Context, input *ec2.DescribeSecurityGroupReferencesInput) *Ec2DescribeSecurityGroupReferencesResult

    DescribeSecurityGroups(ctx workflow.Context, input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error)
    DescribeSecurityGroupsAsync(ctx workflow.Context, input *ec2.DescribeSecurityGroupsInput) *Ec2DescribeSecurityGroupsResult

    DescribeSnapshotAttribute(ctx workflow.Context, input *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error)
    DescribeSnapshotAttributeAsync(ctx workflow.Context, input *ec2.DescribeSnapshotAttributeInput) *Ec2DescribeSnapshotAttributeResult

    DescribeSnapshots(ctx workflow.Context, input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error)
    DescribeSnapshotsAsync(ctx workflow.Context, input *ec2.DescribeSnapshotsInput) *Ec2DescribeSnapshotsResult

    DescribeSpotDatafeedSubscription(ctx workflow.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)
    DescribeSpotDatafeedSubscriptionAsync(ctx workflow.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput) *Ec2DescribeSpotDatafeedSubscriptionResult

    DescribeSpotFleetInstances(ctx workflow.Context, input *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error)
    DescribeSpotFleetInstancesAsync(ctx workflow.Context, input *ec2.DescribeSpotFleetInstancesInput) *Ec2DescribeSpotFleetInstancesResult

    DescribeSpotFleetRequestHistory(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)
    DescribeSpotFleetRequestHistoryAsync(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestHistoryInput) *Ec2DescribeSpotFleetRequestHistoryResult

    DescribeSpotFleetRequests(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error)
    DescribeSpotFleetRequestsAsync(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestsInput) *Ec2DescribeSpotFleetRequestsResult

    DescribeSpotInstanceRequests(ctx workflow.Context, input *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error)
    DescribeSpotInstanceRequestsAsync(ctx workflow.Context, input *ec2.DescribeSpotInstanceRequestsInput) *Ec2DescribeSpotInstanceRequestsResult

    DescribeSpotPriceHistory(ctx workflow.Context, input *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error)
    DescribeSpotPriceHistoryAsync(ctx workflow.Context, input *ec2.DescribeSpotPriceHistoryInput) *Ec2DescribeSpotPriceHistoryResult

    DescribeStaleSecurityGroups(ctx workflow.Context, input *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error)
    DescribeStaleSecurityGroupsAsync(ctx workflow.Context, input *ec2.DescribeStaleSecurityGroupsInput) *Ec2DescribeStaleSecurityGroupsResult

    DescribeSubnets(ctx workflow.Context, input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error)
    DescribeSubnetsAsync(ctx workflow.Context, input *ec2.DescribeSubnetsInput) *Ec2DescribeSubnetsResult

    DescribeTags(ctx workflow.Context, input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error)
    DescribeTagsAsync(ctx workflow.Context, input *ec2.DescribeTagsInput) *Ec2DescribeTagsResult

    DescribeTrafficMirrorFilters(ctx workflow.Context, input *ec2.DescribeTrafficMirrorFiltersInput) (*ec2.DescribeTrafficMirrorFiltersOutput, error)
    DescribeTrafficMirrorFiltersAsync(ctx workflow.Context, input *ec2.DescribeTrafficMirrorFiltersInput) *Ec2DescribeTrafficMirrorFiltersResult

    DescribeTrafficMirrorSessions(ctx workflow.Context, input *ec2.DescribeTrafficMirrorSessionsInput) (*ec2.DescribeTrafficMirrorSessionsOutput, error)
    DescribeTrafficMirrorSessionsAsync(ctx workflow.Context, input *ec2.DescribeTrafficMirrorSessionsInput) *Ec2DescribeTrafficMirrorSessionsResult

    DescribeTrafficMirrorTargets(ctx workflow.Context, input *ec2.DescribeTrafficMirrorTargetsInput) (*ec2.DescribeTrafficMirrorTargetsOutput, error)
    DescribeTrafficMirrorTargetsAsync(ctx workflow.Context, input *ec2.DescribeTrafficMirrorTargetsInput) *Ec2DescribeTrafficMirrorTargetsResult

    DescribeTransitGatewayAttachments(ctx workflow.Context, input *ec2.DescribeTransitGatewayAttachmentsInput) (*ec2.DescribeTransitGatewayAttachmentsOutput, error)
    DescribeTransitGatewayAttachmentsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayAttachmentsInput) *Ec2DescribeTransitGatewayAttachmentsResult

    DescribeTransitGatewayMulticastDomains(ctx workflow.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error)
    DescribeTransitGatewayMulticastDomainsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput) *Ec2DescribeTransitGatewayMulticastDomainsResult

    DescribeTransitGatewayPeeringAttachments(ctx workflow.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error)
    DescribeTransitGatewayPeeringAttachmentsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) *Ec2DescribeTransitGatewayPeeringAttachmentsResult

    DescribeTransitGatewayRouteTables(ctx workflow.Context, input *ec2.DescribeTransitGatewayRouteTablesInput) (*ec2.DescribeTransitGatewayRouteTablesOutput, error)
    DescribeTransitGatewayRouteTablesAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayRouteTablesInput) *Ec2DescribeTransitGatewayRouteTablesResult

    DescribeTransitGatewayVpcAttachments(ctx workflow.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error)
    DescribeTransitGatewayVpcAttachmentsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput) *Ec2DescribeTransitGatewayVpcAttachmentsResult

    DescribeTransitGateways(ctx workflow.Context, input *ec2.DescribeTransitGatewaysInput) (*ec2.DescribeTransitGatewaysOutput, error)
    DescribeTransitGatewaysAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewaysInput) *Ec2DescribeTransitGatewaysResult

    DescribeVolumeAttribute(ctx workflow.Context, input *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error)
    DescribeVolumeAttributeAsync(ctx workflow.Context, input *ec2.DescribeVolumeAttributeInput) *Ec2DescribeVolumeAttributeResult

    DescribeVolumeStatus(ctx workflow.Context, input *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error)
    DescribeVolumeStatusAsync(ctx workflow.Context, input *ec2.DescribeVolumeStatusInput) *Ec2DescribeVolumeStatusResult

    DescribeVolumes(ctx workflow.Context, input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error)
    DescribeVolumesAsync(ctx workflow.Context, input *ec2.DescribeVolumesInput) *Ec2DescribeVolumesResult

    DescribeVolumesModifications(ctx workflow.Context, input *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error)
    DescribeVolumesModificationsAsync(ctx workflow.Context, input *ec2.DescribeVolumesModificationsInput) *Ec2DescribeVolumesModificationsResult

    DescribeVpcAttribute(ctx workflow.Context, input *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error)
    DescribeVpcAttributeAsync(ctx workflow.Context, input *ec2.DescribeVpcAttributeInput) *Ec2DescribeVpcAttributeResult

    DescribeVpcClassicLink(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error)
    DescribeVpcClassicLinkAsync(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkInput) *Ec2DescribeVpcClassicLinkResult

    DescribeVpcClassicLinkDnsSupport(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error)
    DescribeVpcClassicLinkDnsSupportAsync(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput) *Ec2DescribeVpcClassicLinkDnsSupportResult

    DescribeVpcEndpointConnectionNotifications(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error)
    DescribeVpcEndpointConnectionNotificationsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput) *Ec2DescribeVpcEndpointConnectionNotificationsResult

    DescribeVpcEndpointConnections(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error)
    DescribeVpcEndpointConnectionsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionsInput) *Ec2DescribeVpcEndpointConnectionsResult

    DescribeVpcEndpointServiceConfigurations(ctx workflow.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error)
    DescribeVpcEndpointServiceConfigurationsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput) *Ec2DescribeVpcEndpointServiceConfigurationsResult

    DescribeVpcEndpointServicePermissions(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error)
    DescribeVpcEndpointServicePermissionsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput) *Ec2DescribeVpcEndpointServicePermissionsResult

    DescribeVpcEndpointServices(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error)
    DescribeVpcEndpointServicesAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicesInput) *Ec2DescribeVpcEndpointServicesResult

    DescribeVpcEndpoints(ctx workflow.Context, input *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error)
    DescribeVpcEndpointsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointsInput) *Ec2DescribeVpcEndpointsResult

    DescribeVpcPeeringConnections(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
    DescribeVpcPeeringConnectionsAsync(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) *Ec2DescribeVpcPeeringConnectionsResult

    DescribeVpcs(ctx workflow.Context, input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)
    DescribeVpcsAsync(ctx workflow.Context, input *ec2.DescribeVpcsInput) *Ec2DescribeVpcsResult

    DescribeVpnConnections(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error)
    DescribeVpnConnectionsAsync(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) *Ec2DescribeVpnConnectionsResult

    DescribeVpnGateways(ctx workflow.Context, input *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error)
    DescribeVpnGatewaysAsync(ctx workflow.Context, input *ec2.DescribeVpnGatewaysInput) *Ec2DescribeVpnGatewaysResult

    DetachClassicLinkVpc(ctx workflow.Context, input *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error)
    DetachClassicLinkVpcAsync(ctx workflow.Context, input *ec2.DetachClassicLinkVpcInput) *Ec2DetachClassicLinkVpcResult

    DetachInternetGateway(ctx workflow.Context, input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error)
    DetachInternetGatewayAsync(ctx workflow.Context, input *ec2.DetachInternetGatewayInput) *Ec2DetachInternetGatewayResult

    DetachNetworkInterface(ctx workflow.Context, input *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error)
    DetachNetworkInterfaceAsync(ctx workflow.Context, input *ec2.DetachNetworkInterfaceInput) *Ec2DetachNetworkInterfaceResult

    DetachVolume(ctx workflow.Context, input *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error)
    DetachVolumeAsync(ctx workflow.Context, input *ec2.DetachVolumeInput) *Ec2DetachVolumeResult

    DetachVpnGateway(ctx workflow.Context, input *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error)
    DetachVpnGatewayAsync(ctx workflow.Context, input *ec2.DetachVpnGatewayInput) *Ec2DetachVpnGatewayResult

    DisableEbsEncryptionByDefault(ctx workflow.Context, input *ec2.DisableEbsEncryptionByDefaultInput) (*ec2.DisableEbsEncryptionByDefaultOutput, error)
    DisableEbsEncryptionByDefaultAsync(ctx workflow.Context, input *ec2.DisableEbsEncryptionByDefaultInput) *Ec2DisableEbsEncryptionByDefaultResult

    DisableFastSnapshotRestores(ctx workflow.Context, input *ec2.DisableFastSnapshotRestoresInput) (*ec2.DisableFastSnapshotRestoresOutput, error)
    DisableFastSnapshotRestoresAsync(ctx workflow.Context, input *ec2.DisableFastSnapshotRestoresInput) *Ec2DisableFastSnapshotRestoresResult

    DisableTransitGatewayRouteTablePropagation(ctx workflow.Context, input *ec2.DisableTransitGatewayRouteTablePropagationInput) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error)
    DisableTransitGatewayRouteTablePropagationAsync(ctx workflow.Context, input *ec2.DisableTransitGatewayRouteTablePropagationInput) *Ec2DisableTransitGatewayRouteTablePropagationResult

    DisableVgwRoutePropagation(ctx workflow.Context, input *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error)
    DisableVgwRoutePropagationAsync(ctx workflow.Context, input *ec2.DisableVgwRoutePropagationInput) *Ec2DisableVgwRoutePropagationResult

    DisableVpcClassicLink(ctx workflow.Context, input *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error)
    DisableVpcClassicLinkAsync(ctx workflow.Context, input *ec2.DisableVpcClassicLinkInput) *Ec2DisableVpcClassicLinkResult

    DisableVpcClassicLinkDnsSupport(ctx workflow.Context, input *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error)
    DisableVpcClassicLinkDnsSupportAsync(ctx workflow.Context, input *ec2.DisableVpcClassicLinkDnsSupportInput) *Ec2DisableVpcClassicLinkDnsSupportResult

    DisassociateAddress(ctx workflow.Context, input *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error)
    DisassociateAddressAsync(ctx workflow.Context, input *ec2.DisassociateAddressInput) *Ec2DisassociateAddressResult

    DisassociateClientVpnTargetNetwork(ctx workflow.Context, input *ec2.DisassociateClientVpnTargetNetworkInput) (*ec2.DisassociateClientVpnTargetNetworkOutput, error)
    DisassociateClientVpnTargetNetworkAsync(ctx workflow.Context, input *ec2.DisassociateClientVpnTargetNetworkInput) *Ec2DisassociateClientVpnTargetNetworkResult

    DisassociateIamInstanceProfile(ctx workflow.Context, input *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error)
    DisassociateIamInstanceProfileAsync(ctx workflow.Context, input *ec2.DisassociateIamInstanceProfileInput) *Ec2DisassociateIamInstanceProfileResult

    DisassociateRouteTable(ctx workflow.Context, input *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error)
    DisassociateRouteTableAsync(ctx workflow.Context, input *ec2.DisassociateRouteTableInput) *Ec2DisassociateRouteTableResult

    DisassociateSubnetCidrBlock(ctx workflow.Context, input *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error)
    DisassociateSubnetCidrBlockAsync(ctx workflow.Context, input *ec2.DisassociateSubnetCidrBlockInput) *Ec2DisassociateSubnetCidrBlockResult

    DisassociateTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.DisassociateTransitGatewayMulticastDomainInput) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error)
    DisassociateTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.DisassociateTransitGatewayMulticastDomainInput) *Ec2DisassociateTransitGatewayMulticastDomainResult

    DisassociateTransitGatewayRouteTable(ctx workflow.Context, input *ec2.DisassociateTransitGatewayRouteTableInput) (*ec2.DisassociateTransitGatewayRouteTableOutput, error)
    DisassociateTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.DisassociateTransitGatewayRouteTableInput) *Ec2DisassociateTransitGatewayRouteTableResult

    DisassociateVpcCidrBlock(ctx workflow.Context, input *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error)
    DisassociateVpcCidrBlockAsync(ctx workflow.Context, input *ec2.DisassociateVpcCidrBlockInput) *Ec2DisassociateVpcCidrBlockResult

    EnableEbsEncryptionByDefault(ctx workflow.Context, input *ec2.EnableEbsEncryptionByDefaultInput) (*ec2.EnableEbsEncryptionByDefaultOutput, error)
    EnableEbsEncryptionByDefaultAsync(ctx workflow.Context, input *ec2.EnableEbsEncryptionByDefaultInput) *Ec2EnableEbsEncryptionByDefaultResult

    EnableFastSnapshotRestores(ctx workflow.Context, input *ec2.EnableFastSnapshotRestoresInput) (*ec2.EnableFastSnapshotRestoresOutput, error)
    EnableFastSnapshotRestoresAsync(ctx workflow.Context, input *ec2.EnableFastSnapshotRestoresInput) *Ec2EnableFastSnapshotRestoresResult

    EnableTransitGatewayRouteTablePropagation(ctx workflow.Context, input *ec2.EnableTransitGatewayRouteTablePropagationInput) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error)
    EnableTransitGatewayRouteTablePropagationAsync(ctx workflow.Context, input *ec2.EnableTransitGatewayRouteTablePropagationInput) *Ec2EnableTransitGatewayRouteTablePropagationResult

    EnableVgwRoutePropagation(ctx workflow.Context, input *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error)
    EnableVgwRoutePropagationAsync(ctx workflow.Context, input *ec2.EnableVgwRoutePropagationInput) *Ec2EnableVgwRoutePropagationResult

    EnableVolumeIO(ctx workflow.Context, input *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error)
    EnableVolumeIOAsync(ctx workflow.Context, input *ec2.EnableVolumeIOInput) *Ec2EnableVolumeIOResult

    EnableVpcClassicLink(ctx workflow.Context, input *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error)
    EnableVpcClassicLinkAsync(ctx workflow.Context, input *ec2.EnableVpcClassicLinkInput) *Ec2EnableVpcClassicLinkResult

    EnableVpcClassicLinkDnsSupport(ctx workflow.Context, input *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error)
    EnableVpcClassicLinkDnsSupportAsync(ctx workflow.Context, input *ec2.EnableVpcClassicLinkDnsSupportInput) *Ec2EnableVpcClassicLinkDnsSupportResult

    ExportClientVpnClientCertificateRevocationList(ctx workflow.Context, input *ec2.ExportClientVpnClientCertificateRevocationListInput) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error)
    ExportClientVpnClientCertificateRevocationListAsync(ctx workflow.Context, input *ec2.ExportClientVpnClientCertificateRevocationListInput) *Ec2ExportClientVpnClientCertificateRevocationListResult

    ExportClientVpnClientConfiguration(ctx workflow.Context, input *ec2.ExportClientVpnClientConfigurationInput) (*ec2.ExportClientVpnClientConfigurationOutput, error)
    ExportClientVpnClientConfigurationAsync(ctx workflow.Context, input *ec2.ExportClientVpnClientConfigurationInput) *Ec2ExportClientVpnClientConfigurationResult

    ExportImage(ctx workflow.Context, input *ec2.ExportImageInput) (*ec2.ExportImageOutput, error)
    ExportImageAsync(ctx workflow.Context, input *ec2.ExportImageInput) *Ec2ExportImageResult

    ExportTransitGatewayRoutes(ctx workflow.Context, input *ec2.ExportTransitGatewayRoutesInput) (*ec2.ExportTransitGatewayRoutesOutput, error)
    ExportTransitGatewayRoutesAsync(ctx workflow.Context, input *ec2.ExportTransitGatewayRoutesInput) *Ec2ExportTransitGatewayRoutesResult

    GetAssociatedIpv6PoolCidrs(ctx workflow.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error)
    GetAssociatedIpv6PoolCidrsAsync(ctx workflow.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput) *Ec2GetAssociatedIpv6PoolCidrsResult

    GetCapacityReservationUsage(ctx workflow.Context, input *ec2.GetCapacityReservationUsageInput) (*ec2.GetCapacityReservationUsageOutput, error)
    GetCapacityReservationUsageAsync(ctx workflow.Context, input *ec2.GetCapacityReservationUsageInput) *Ec2GetCapacityReservationUsageResult

    GetCoipPoolUsage(ctx workflow.Context, input *ec2.GetCoipPoolUsageInput) (*ec2.GetCoipPoolUsageOutput, error)
    GetCoipPoolUsageAsync(ctx workflow.Context, input *ec2.GetCoipPoolUsageInput) *Ec2GetCoipPoolUsageResult

    GetConsoleOutput(ctx workflow.Context, input *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error)
    GetConsoleOutputAsync(ctx workflow.Context, input *ec2.GetConsoleOutputInput) *Ec2GetConsoleOutputResult

    GetConsoleScreenshot(ctx workflow.Context, input *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error)
    GetConsoleScreenshotAsync(ctx workflow.Context, input *ec2.GetConsoleScreenshotInput) *Ec2GetConsoleScreenshotResult

    GetDefaultCreditSpecification(ctx workflow.Context, input *ec2.GetDefaultCreditSpecificationInput) (*ec2.GetDefaultCreditSpecificationOutput, error)
    GetDefaultCreditSpecificationAsync(ctx workflow.Context, input *ec2.GetDefaultCreditSpecificationInput) *Ec2GetDefaultCreditSpecificationResult

    GetEbsDefaultKmsKeyId(ctx workflow.Context, input *ec2.GetEbsDefaultKmsKeyIdInput) (*ec2.GetEbsDefaultKmsKeyIdOutput, error)
    GetEbsDefaultKmsKeyIdAsync(ctx workflow.Context, input *ec2.GetEbsDefaultKmsKeyIdInput) *Ec2GetEbsDefaultKmsKeyIdResult

    GetEbsEncryptionByDefault(ctx workflow.Context, input *ec2.GetEbsEncryptionByDefaultInput) (*ec2.GetEbsEncryptionByDefaultOutput, error)
    GetEbsEncryptionByDefaultAsync(ctx workflow.Context, input *ec2.GetEbsEncryptionByDefaultInput) *Ec2GetEbsEncryptionByDefaultResult

    GetGroupsForCapacityReservation(ctx workflow.Context, input *ec2.GetGroupsForCapacityReservationInput) (*ec2.GetGroupsForCapacityReservationOutput, error)
    GetGroupsForCapacityReservationAsync(ctx workflow.Context, input *ec2.GetGroupsForCapacityReservationInput) *Ec2GetGroupsForCapacityReservationResult

    GetHostReservationPurchasePreview(ctx workflow.Context, input *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error)
    GetHostReservationPurchasePreviewAsync(ctx workflow.Context, input *ec2.GetHostReservationPurchasePreviewInput) *Ec2GetHostReservationPurchasePreviewResult

    GetLaunchTemplateData(ctx workflow.Context, input *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error)
    GetLaunchTemplateDataAsync(ctx workflow.Context, input *ec2.GetLaunchTemplateDataInput) *Ec2GetLaunchTemplateDataResult

    GetManagedPrefixListAssociations(ctx workflow.Context, input *ec2.GetManagedPrefixListAssociationsInput) (*ec2.GetManagedPrefixListAssociationsOutput, error)
    GetManagedPrefixListAssociationsAsync(ctx workflow.Context, input *ec2.GetManagedPrefixListAssociationsInput) *Ec2GetManagedPrefixListAssociationsResult

    GetManagedPrefixListEntries(ctx workflow.Context, input *ec2.GetManagedPrefixListEntriesInput) (*ec2.GetManagedPrefixListEntriesOutput, error)
    GetManagedPrefixListEntriesAsync(ctx workflow.Context, input *ec2.GetManagedPrefixListEntriesInput) *Ec2GetManagedPrefixListEntriesResult

    GetPasswordData(ctx workflow.Context, input *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error)
    GetPasswordDataAsync(ctx workflow.Context, input *ec2.GetPasswordDataInput) *Ec2GetPasswordDataResult

    GetReservedInstancesExchangeQuote(ctx workflow.Context, input *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error)
    GetReservedInstancesExchangeQuoteAsync(ctx workflow.Context, input *ec2.GetReservedInstancesExchangeQuoteInput) *Ec2GetReservedInstancesExchangeQuoteResult

    GetTransitGatewayAttachmentPropagations(ctx workflow.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error)
    GetTransitGatewayAttachmentPropagationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput) *Ec2GetTransitGatewayAttachmentPropagationsResult

    GetTransitGatewayMulticastDomainAssociations(ctx workflow.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error)
    GetTransitGatewayMulticastDomainAssociationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) *Ec2GetTransitGatewayMulticastDomainAssociationsResult

    GetTransitGatewayPrefixListReferences(ctx workflow.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error)
    GetTransitGatewayPrefixListReferencesAsync(ctx workflow.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput) *Ec2GetTransitGatewayPrefixListReferencesResult

    GetTransitGatewayRouteTableAssociations(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error)
    GetTransitGatewayRouteTableAssociationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput) *Ec2GetTransitGatewayRouteTableAssociationsResult

    GetTransitGatewayRouteTablePropagations(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error)
    GetTransitGatewayRouteTablePropagationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput) *Ec2GetTransitGatewayRouteTablePropagationsResult

    ImportClientVpnClientCertificateRevocationList(ctx workflow.Context, input *ec2.ImportClientVpnClientCertificateRevocationListInput) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error)
    ImportClientVpnClientCertificateRevocationListAsync(ctx workflow.Context, input *ec2.ImportClientVpnClientCertificateRevocationListInput) *Ec2ImportClientVpnClientCertificateRevocationListResult

    ImportImage(ctx workflow.Context, input *ec2.ImportImageInput) (*ec2.ImportImageOutput, error)
    ImportImageAsync(ctx workflow.Context, input *ec2.ImportImageInput) *Ec2ImportImageResult

    ImportInstance(ctx workflow.Context, input *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error)
    ImportInstanceAsync(ctx workflow.Context, input *ec2.ImportInstanceInput) *Ec2ImportInstanceResult

    ImportKeyPair(ctx workflow.Context, input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error)
    ImportKeyPairAsync(ctx workflow.Context, input *ec2.ImportKeyPairInput) *Ec2ImportKeyPairResult

    ImportSnapshot(ctx workflow.Context, input *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error)
    ImportSnapshotAsync(ctx workflow.Context, input *ec2.ImportSnapshotInput) *Ec2ImportSnapshotResult

    ImportVolume(ctx workflow.Context, input *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error)
    ImportVolumeAsync(ctx workflow.Context, input *ec2.ImportVolumeInput) *Ec2ImportVolumeResult

    ModifyAvailabilityZoneGroup(ctx workflow.Context, input *ec2.ModifyAvailabilityZoneGroupInput) (*ec2.ModifyAvailabilityZoneGroupOutput, error)
    ModifyAvailabilityZoneGroupAsync(ctx workflow.Context, input *ec2.ModifyAvailabilityZoneGroupInput) *Ec2ModifyAvailabilityZoneGroupResult

    ModifyCapacityReservation(ctx workflow.Context, input *ec2.ModifyCapacityReservationInput) (*ec2.ModifyCapacityReservationOutput, error)
    ModifyCapacityReservationAsync(ctx workflow.Context, input *ec2.ModifyCapacityReservationInput) *Ec2ModifyCapacityReservationResult

    ModifyClientVpnEndpoint(ctx workflow.Context, input *ec2.ModifyClientVpnEndpointInput) (*ec2.ModifyClientVpnEndpointOutput, error)
    ModifyClientVpnEndpointAsync(ctx workflow.Context, input *ec2.ModifyClientVpnEndpointInput) *Ec2ModifyClientVpnEndpointResult

    ModifyDefaultCreditSpecification(ctx workflow.Context, input *ec2.ModifyDefaultCreditSpecificationInput) (*ec2.ModifyDefaultCreditSpecificationOutput, error)
    ModifyDefaultCreditSpecificationAsync(ctx workflow.Context, input *ec2.ModifyDefaultCreditSpecificationInput) *Ec2ModifyDefaultCreditSpecificationResult

    ModifyEbsDefaultKmsKeyId(ctx workflow.Context, input *ec2.ModifyEbsDefaultKmsKeyIdInput) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error)
    ModifyEbsDefaultKmsKeyIdAsync(ctx workflow.Context, input *ec2.ModifyEbsDefaultKmsKeyIdInput) *Ec2ModifyEbsDefaultKmsKeyIdResult

    ModifyFleet(ctx workflow.Context, input *ec2.ModifyFleetInput) (*ec2.ModifyFleetOutput, error)
    ModifyFleetAsync(ctx workflow.Context, input *ec2.ModifyFleetInput) *Ec2ModifyFleetResult

    ModifyFpgaImageAttribute(ctx workflow.Context, input *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error)
    ModifyFpgaImageAttributeAsync(ctx workflow.Context, input *ec2.ModifyFpgaImageAttributeInput) *Ec2ModifyFpgaImageAttributeResult

    ModifyHosts(ctx workflow.Context, input *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error)
    ModifyHostsAsync(ctx workflow.Context, input *ec2.ModifyHostsInput) *Ec2ModifyHostsResult

    ModifyIdFormat(ctx workflow.Context, input *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error)
    ModifyIdFormatAsync(ctx workflow.Context, input *ec2.ModifyIdFormatInput) *Ec2ModifyIdFormatResult

    ModifyIdentityIdFormat(ctx workflow.Context, input *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error)
    ModifyIdentityIdFormatAsync(ctx workflow.Context, input *ec2.ModifyIdentityIdFormatInput) *Ec2ModifyIdentityIdFormatResult

    ModifyImageAttribute(ctx workflow.Context, input *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error)
    ModifyImageAttributeAsync(ctx workflow.Context, input *ec2.ModifyImageAttributeInput) *Ec2ModifyImageAttributeResult

    ModifyInstanceAttribute(ctx workflow.Context, input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error)
    ModifyInstanceAttributeAsync(ctx workflow.Context, input *ec2.ModifyInstanceAttributeInput) *Ec2ModifyInstanceAttributeResult

    ModifyInstanceCapacityReservationAttributes(ctx workflow.Context, input *ec2.ModifyInstanceCapacityReservationAttributesInput) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error)
    ModifyInstanceCapacityReservationAttributesAsync(ctx workflow.Context, input *ec2.ModifyInstanceCapacityReservationAttributesInput) *Ec2ModifyInstanceCapacityReservationAttributesResult

    ModifyInstanceCreditSpecification(ctx workflow.Context, input *ec2.ModifyInstanceCreditSpecificationInput) (*ec2.ModifyInstanceCreditSpecificationOutput, error)
    ModifyInstanceCreditSpecificationAsync(ctx workflow.Context, input *ec2.ModifyInstanceCreditSpecificationInput) *Ec2ModifyInstanceCreditSpecificationResult

    ModifyInstanceEventStartTime(ctx workflow.Context, input *ec2.ModifyInstanceEventStartTimeInput) (*ec2.ModifyInstanceEventStartTimeOutput, error)
    ModifyInstanceEventStartTimeAsync(ctx workflow.Context, input *ec2.ModifyInstanceEventStartTimeInput) *Ec2ModifyInstanceEventStartTimeResult

    ModifyInstanceMetadataOptions(ctx workflow.Context, input *ec2.ModifyInstanceMetadataOptionsInput) (*ec2.ModifyInstanceMetadataOptionsOutput, error)
    ModifyInstanceMetadataOptionsAsync(ctx workflow.Context, input *ec2.ModifyInstanceMetadataOptionsInput) *Ec2ModifyInstanceMetadataOptionsResult

    ModifyInstancePlacement(ctx workflow.Context, input *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error)
    ModifyInstancePlacementAsync(ctx workflow.Context, input *ec2.ModifyInstancePlacementInput) *Ec2ModifyInstancePlacementResult

    ModifyLaunchTemplate(ctx workflow.Context, input *ec2.ModifyLaunchTemplateInput) (*ec2.ModifyLaunchTemplateOutput, error)
    ModifyLaunchTemplateAsync(ctx workflow.Context, input *ec2.ModifyLaunchTemplateInput) *Ec2ModifyLaunchTemplateResult

    ModifyManagedPrefixList(ctx workflow.Context, input *ec2.ModifyManagedPrefixListInput) (*ec2.ModifyManagedPrefixListOutput, error)
    ModifyManagedPrefixListAsync(ctx workflow.Context, input *ec2.ModifyManagedPrefixListInput) *Ec2ModifyManagedPrefixListResult

    ModifyNetworkInterfaceAttribute(ctx workflow.Context, input *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)
    ModifyNetworkInterfaceAttributeAsync(ctx workflow.Context, input *ec2.ModifyNetworkInterfaceAttributeInput) *Ec2ModifyNetworkInterfaceAttributeResult

    ModifyReservedInstances(ctx workflow.Context, input *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error)
    ModifyReservedInstancesAsync(ctx workflow.Context, input *ec2.ModifyReservedInstancesInput) *Ec2ModifyReservedInstancesResult

    ModifySnapshotAttribute(ctx workflow.Context, input *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error)
    ModifySnapshotAttributeAsync(ctx workflow.Context, input *ec2.ModifySnapshotAttributeInput) *Ec2ModifySnapshotAttributeResult

    ModifySubnetAttribute(ctx workflow.Context, input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error)
    ModifySubnetAttributeAsync(ctx workflow.Context, input *ec2.ModifySubnetAttributeInput) *Ec2ModifySubnetAttributeResult

    ModifyTrafficMirrorFilterNetworkServices(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error)
    ModifyTrafficMirrorFilterNetworkServicesAsync(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) *Ec2ModifyTrafficMirrorFilterNetworkServicesResult

    ModifyTrafficMirrorFilterRule(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterRuleInput) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error)
    ModifyTrafficMirrorFilterRuleAsync(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterRuleInput) *Ec2ModifyTrafficMirrorFilterRuleResult

    ModifyTrafficMirrorSession(ctx workflow.Context, input *ec2.ModifyTrafficMirrorSessionInput) (*ec2.ModifyTrafficMirrorSessionOutput, error)
    ModifyTrafficMirrorSessionAsync(ctx workflow.Context, input *ec2.ModifyTrafficMirrorSessionInput) *Ec2ModifyTrafficMirrorSessionResult

    ModifyTransitGateway(ctx workflow.Context, input *ec2.ModifyTransitGatewayInput) (*ec2.ModifyTransitGatewayOutput, error)
    ModifyTransitGatewayAsync(ctx workflow.Context, input *ec2.ModifyTransitGatewayInput) *Ec2ModifyTransitGatewayResult

    ModifyTransitGatewayPrefixListReference(ctx workflow.Context, input *ec2.ModifyTransitGatewayPrefixListReferenceInput) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error)
    ModifyTransitGatewayPrefixListReferenceAsync(ctx workflow.Context, input *ec2.ModifyTransitGatewayPrefixListReferenceInput) *Ec2ModifyTransitGatewayPrefixListReferenceResult

    ModifyTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.ModifyTransitGatewayVpcAttachmentInput) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error)
    ModifyTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.ModifyTransitGatewayVpcAttachmentInput) *Ec2ModifyTransitGatewayVpcAttachmentResult

    ModifyVolume(ctx workflow.Context, input *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error)
    ModifyVolumeAsync(ctx workflow.Context, input *ec2.ModifyVolumeInput) *Ec2ModifyVolumeResult

    ModifyVolumeAttribute(ctx workflow.Context, input *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error)
    ModifyVolumeAttributeAsync(ctx workflow.Context, input *ec2.ModifyVolumeAttributeInput) *Ec2ModifyVolumeAttributeResult

    ModifyVpcAttribute(ctx workflow.Context, input *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error)
    ModifyVpcAttributeAsync(ctx workflow.Context, input *ec2.ModifyVpcAttributeInput) *Ec2ModifyVpcAttributeResult

    ModifyVpcEndpoint(ctx workflow.Context, input *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error)
    ModifyVpcEndpointAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointInput) *Ec2ModifyVpcEndpointResult

    ModifyVpcEndpointConnectionNotification(ctx workflow.Context, input *ec2.ModifyVpcEndpointConnectionNotificationInput) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error)
    ModifyVpcEndpointConnectionNotificationAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointConnectionNotificationInput) *Ec2ModifyVpcEndpointConnectionNotificationResult

    ModifyVpcEndpointServiceConfiguration(ctx workflow.Context, input *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error)
    ModifyVpcEndpointServiceConfigurationAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointServiceConfigurationInput) *Ec2ModifyVpcEndpointServiceConfigurationResult

    ModifyVpcEndpointServicePermissions(ctx workflow.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error)
    ModifyVpcEndpointServicePermissionsAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput) *Ec2ModifyVpcEndpointServicePermissionsResult

    ModifyVpcPeeringConnectionOptions(ctx workflow.Context, input *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error)
    ModifyVpcPeeringConnectionOptionsAsync(ctx workflow.Context, input *ec2.ModifyVpcPeeringConnectionOptionsInput) *Ec2ModifyVpcPeeringConnectionOptionsResult

    ModifyVpcTenancy(ctx workflow.Context, input *ec2.ModifyVpcTenancyInput) (*ec2.ModifyVpcTenancyOutput, error)
    ModifyVpcTenancyAsync(ctx workflow.Context, input *ec2.ModifyVpcTenancyInput) *Ec2ModifyVpcTenancyResult

    ModifyVpnConnection(ctx workflow.Context, input *ec2.ModifyVpnConnectionInput) (*ec2.ModifyVpnConnectionOutput, error)
    ModifyVpnConnectionAsync(ctx workflow.Context, input *ec2.ModifyVpnConnectionInput) *Ec2ModifyVpnConnectionResult

    ModifyVpnConnectionOptions(ctx workflow.Context, input *ec2.ModifyVpnConnectionOptionsInput) (*ec2.ModifyVpnConnectionOptionsOutput, error)
    ModifyVpnConnectionOptionsAsync(ctx workflow.Context, input *ec2.ModifyVpnConnectionOptionsInput) *Ec2ModifyVpnConnectionOptionsResult

    ModifyVpnTunnelCertificate(ctx workflow.Context, input *ec2.ModifyVpnTunnelCertificateInput) (*ec2.ModifyVpnTunnelCertificateOutput, error)
    ModifyVpnTunnelCertificateAsync(ctx workflow.Context, input *ec2.ModifyVpnTunnelCertificateInput) *Ec2ModifyVpnTunnelCertificateResult

    ModifyVpnTunnelOptions(ctx workflow.Context, input *ec2.ModifyVpnTunnelOptionsInput) (*ec2.ModifyVpnTunnelOptionsOutput, error)
    ModifyVpnTunnelOptionsAsync(ctx workflow.Context, input *ec2.ModifyVpnTunnelOptionsInput) *Ec2ModifyVpnTunnelOptionsResult

    MonitorInstances(ctx workflow.Context, input *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error)
    MonitorInstancesAsync(ctx workflow.Context, input *ec2.MonitorInstancesInput) *Ec2MonitorInstancesResult

    MoveAddressToVpc(ctx workflow.Context, input *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error)
    MoveAddressToVpcAsync(ctx workflow.Context, input *ec2.MoveAddressToVpcInput) *Ec2MoveAddressToVpcResult

    ProvisionByoipCidr(ctx workflow.Context, input *ec2.ProvisionByoipCidrInput) (*ec2.ProvisionByoipCidrOutput, error)
    ProvisionByoipCidrAsync(ctx workflow.Context, input *ec2.ProvisionByoipCidrInput) *Ec2ProvisionByoipCidrResult

    PurchaseHostReservation(ctx workflow.Context, input *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error)
    PurchaseHostReservationAsync(ctx workflow.Context, input *ec2.PurchaseHostReservationInput) *Ec2PurchaseHostReservationResult

    PurchaseReservedInstancesOffering(ctx workflow.Context, input *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error)
    PurchaseReservedInstancesOfferingAsync(ctx workflow.Context, input *ec2.PurchaseReservedInstancesOfferingInput) *Ec2PurchaseReservedInstancesOfferingResult

    PurchaseScheduledInstances(ctx workflow.Context, input *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error)
    PurchaseScheduledInstancesAsync(ctx workflow.Context, input *ec2.PurchaseScheduledInstancesInput) *Ec2PurchaseScheduledInstancesResult

    RebootInstances(ctx workflow.Context, input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error)
    RebootInstancesAsync(ctx workflow.Context, input *ec2.RebootInstancesInput) *Ec2RebootInstancesResult

    RegisterImage(ctx workflow.Context, input *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error)
    RegisterImageAsync(ctx workflow.Context, input *ec2.RegisterImageInput) *Ec2RegisterImageResult

    RegisterInstanceEventNotificationAttributes(ctx workflow.Context, input *ec2.RegisterInstanceEventNotificationAttributesInput) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error)
    RegisterInstanceEventNotificationAttributesAsync(ctx workflow.Context, input *ec2.RegisterInstanceEventNotificationAttributesInput) *Ec2RegisterInstanceEventNotificationAttributesResult

    RegisterTransitGatewayMulticastGroupMembers(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error)
    RegisterTransitGatewayMulticastGroupMembersAsync(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) *Ec2RegisterTransitGatewayMulticastGroupMembersResult

    RegisterTransitGatewayMulticastGroupSources(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error)
    RegisterTransitGatewayMulticastGroupSourcesAsync(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) *Ec2RegisterTransitGatewayMulticastGroupSourcesResult

    RejectTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.RejectTransitGatewayPeeringAttachmentInput) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error)
    RejectTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.RejectTransitGatewayPeeringAttachmentInput) *Ec2RejectTransitGatewayPeeringAttachmentResult

    RejectTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.RejectTransitGatewayVpcAttachmentInput) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error)
    RejectTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.RejectTransitGatewayVpcAttachmentInput) *Ec2RejectTransitGatewayVpcAttachmentResult

    RejectVpcEndpointConnections(ctx workflow.Context, input *ec2.RejectVpcEndpointConnectionsInput) (*ec2.RejectVpcEndpointConnectionsOutput, error)
    RejectVpcEndpointConnectionsAsync(ctx workflow.Context, input *ec2.RejectVpcEndpointConnectionsInput) *Ec2RejectVpcEndpointConnectionsResult

    RejectVpcPeeringConnection(ctx workflow.Context, input *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error)
    RejectVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.RejectVpcPeeringConnectionInput) *Ec2RejectVpcPeeringConnectionResult

    ReleaseAddress(ctx workflow.Context, input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error)
    ReleaseAddressAsync(ctx workflow.Context, input *ec2.ReleaseAddressInput) *Ec2ReleaseAddressResult

    ReleaseHosts(ctx workflow.Context, input *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error)
    ReleaseHostsAsync(ctx workflow.Context, input *ec2.ReleaseHostsInput) *Ec2ReleaseHostsResult

    ReplaceIamInstanceProfileAssociation(ctx workflow.Context, input *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error)
    ReplaceIamInstanceProfileAssociationAsync(ctx workflow.Context, input *ec2.ReplaceIamInstanceProfileAssociationInput) *Ec2ReplaceIamInstanceProfileAssociationResult

    ReplaceNetworkAclAssociation(ctx workflow.Context, input *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error)
    ReplaceNetworkAclAssociationAsync(ctx workflow.Context, input *ec2.ReplaceNetworkAclAssociationInput) *Ec2ReplaceNetworkAclAssociationResult

    ReplaceNetworkAclEntry(ctx workflow.Context, input *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error)
    ReplaceNetworkAclEntryAsync(ctx workflow.Context, input *ec2.ReplaceNetworkAclEntryInput) *Ec2ReplaceNetworkAclEntryResult

    ReplaceRoute(ctx workflow.Context, input *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error)
    ReplaceRouteAsync(ctx workflow.Context, input *ec2.ReplaceRouteInput) *Ec2ReplaceRouteResult

    ReplaceRouteTableAssociation(ctx workflow.Context, input *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error)
    ReplaceRouteTableAssociationAsync(ctx workflow.Context, input *ec2.ReplaceRouteTableAssociationInput) *Ec2ReplaceRouteTableAssociationResult

    ReplaceTransitGatewayRoute(ctx workflow.Context, input *ec2.ReplaceTransitGatewayRouteInput) (*ec2.ReplaceTransitGatewayRouteOutput, error)
    ReplaceTransitGatewayRouteAsync(ctx workflow.Context, input *ec2.ReplaceTransitGatewayRouteInput) *Ec2ReplaceTransitGatewayRouteResult

    ReportInstanceStatus(ctx workflow.Context, input *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error)
    ReportInstanceStatusAsync(ctx workflow.Context, input *ec2.ReportInstanceStatusInput) *Ec2ReportInstanceStatusResult

    RequestSpotFleet(ctx workflow.Context, input *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error)
    RequestSpotFleetAsync(ctx workflow.Context, input *ec2.RequestSpotFleetInput) *Ec2RequestSpotFleetResult

    RequestSpotInstances(ctx workflow.Context, input *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error)
    RequestSpotInstancesAsync(ctx workflow.Context, input *ec2.RequestSpotInstancesInput) *Ec2RequestSpotInstancesResult

    ResetEbsDefaultKmsKeyId(ctx workflow.Context, input *ec2.ResetEbsDefaultKmsKeyIdInput) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error)
    ResetEbsDefaultKmsKeyIdAsync(ctx workflow.Context, input *ec2.ResetEbsDefaultKmsKeyIdInput) *Ec2ResetEbsDefaultKmsKeyIdResult

    ResetFpgaImageAttribute(ctx workflow.Context, input *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error)
    ResetFpgaImageAttributeAsync(ctx workflow.Context, input *ec2.ResetFpgaImageAttributeInput) *Ec2ResetFpgaImageAttributeResult

    ResetImageAttribute(ctx workflow.Context, input *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error)
    ResetImageAttributeAsync(ctx workflow.Context, input *ec2.ResetImageAttributeInput) *Ec2ResetImageAttributeResult

    ResetInstanceAttribute(ctx workflow.Context, input *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error)
    ResetInstanceAttributeAsync(ctx workflow.Context, input *ec2.ResetInstanceAttributeInput) *Ec2ResetInstanceAttributeResult

    ResetNetworkInterfaceAttribute(ctx workflow.Context, input *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error)
    ResetNetworkInterfaceAttributeAsync(ctx workflow.Context, input *ec2.ResetNetworkInterfaceAttributeInput) *Ec2ResetNetworkInterfaceAttributeResult

    ResetSnapshotAttribute(ctx workflow.Context, input *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error)
    ResetSnapshotAttributeAsync(ctx workflow.Context, input *ec2.ResetSnapshotAttributeInput) *Ec2ResetSnapshotAttributeResult

    RestoreAddressToClassic(ctx workflow.Context, input *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error)
    RestoreAddressToClassicAsync(ctx workflow.Context, input *ec2.RestoreAddressToClassicInput) *Ec2RestoreAddressToClassicResult

    RestoreManagedPrefixListVersion(ctx workflow.Context, input *ec2.RestoreManagedPrefixListVersionInput) (*ec2.RestoreManagedPrefixListVersionOutput, error)
    RestoreManagedPrefixListVersionAsync(ctx workflow.Context, input *ec2.RestoreManagedPrefixListVersionInput) *Ec2RestoreManagedPrefixListVersionResult

    RevokeClientVpnIngress(ctx workflow.Context, input *ec2.RevokeClientVpnIngressInput) (*ec2.RevokeClientVpnIngressOutput, error)
    RevokeClientVpnIngressAsync(ctx workflow.Context, input *ec2.RevokeClientVpnIngressInput) *Ec2RevokeClientVpnIngressResult

    RevokeSecurityGroupEgress(ctx workflow.Context, input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error)
    RevokeSecurityGroupEgressAsync(ctx workflow.Context, input *ec2.RevokeSecurityGroupEgressInput) *Ec2RevokeSecurityGroupEgressResult

    RevokeSecurityGroupIngress(ctx workflow.Context, input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error)
    RevokeSecurityGroupIngressAsync(ctx workflow.Context, input *ec2.RevokeSecurityGroupIngressInput) *Ec2RevokeSecurityGroupIngressResult

    RunInstances(ctx workflow.Context, input *ec2.RunInstancesInput) (*ec2.Reservation, error)
    RunInstancesAsync(ctx workflow.Context, input *ec2.RunInstancesInput) *Ec2RunInstancesResult

    RunScheduledInstances(ctx workflow.Context, input *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error)
    RunScheduledInstancesAsync(ctx workflow.Context, input *ec2.RunScheduledInstancesInput) *Ec2RunScheduledInstancesResult

    SearchLocalGatewayRoutes(ctx workflow.Context, input *ec2.SearchLocalGatewayRoutesInput) (*ec2.SearchLocalGatewayRoutesOutput, error)
    SearchLocalGatewayRoutesAsync(ctx workflow.Context, input *ec2.SearchLocalGatewayRoutesInput) *Ec2SearchLocalGatewayRoutesResult

    SearchTransitGatewayMulticastGroups(ctx workflow.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error)
    SearchTransitGatewayMulticastGroupsAsync(ctx workflow.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput) *Ec2SearchTransitGatewayMulticastGroupsResult

    SearchTransitGatewayRoutes(ctx workflow.Context, input *ec2.SearchTransitGatewayRoutesInput) (*ec2.SearchTransitGatewayRoutesOutput, error)
    SearchTransitGatewayRoutesAsync(ctx workflow.Context, input *ec2.SearchTransitGatewayRoutesInput) *Ec2SearchTransitGatewayRoutesResult

    SendDiagnosticInterrupt(ctx workflow.Context, input *ec2.SendDiagnosticInterruptInput) (*ec2.SendDiagnosticInterruptOutput, error)
    SendDiagnosticInterruptAsync(ctx workflow.Context, input *ec2.SendDiagnosticInterruptInput) *Ec2SendDiagnosticInterruptResult

    StartInstances(ctx workflow.Context, input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error)
    StartInstancesAsync(ctx workflow.Context, input *ec2.StartInstancesInput) *Ec2StartInstancesResult

    StartVpcEndpointServicePrivateDnsVerification(ctx workflow.Context, input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error)
    StartVpcEndpointServicePrivateDnsVerificationAsync(ctx workflow.Context, input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) *Ec2StartVpcEndpointServicePrivateDnsVerificationResult

    StopInstances(ctx workflow.Context, input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error)
    StopInstancesAsync(ctx workflow.Context, input *ec2.StopInstancesInput) *Ec2StopInstancesResult

    TerminateClientVpnConnections(ctx workflow.Context, input *ec2.TerminateClientVpnConnectionsInput) (*ec2.TerminateClientVpnConnectionsOutput, error)
    TerminateClientVpnConnectionsAsync(ctx workflow.Context, input *ec2.TerminateClientVpnConnectionsInput) *Ec2TerminateClientVpnConnectionsResult

    TerminateInstances(ctx workflow.Context, input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)
    TerminateInstancesAsync(ctx workflow.Context, input *ec2.TerminateInstancesInput) *Ec2TerminateInstancesResult

    UnassignIpv6Addresses(ctx workflow.Context, input *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error)
    UnassignIpv6AddressesAsync(ctx workflow.Context, input *ec2.UnassignIpv6AddressesInput) *Ec2UnassignIpv6AddressesResult

    UnassignPrivateIpAddresses(ctx workflow.Context, input *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error)
    UnassignPrivateIpAddressesAsync(ctx workflow.Context, input *ec2.UnassignPrivateIpAddressesInput) *Ec2UnassignPrivateIpAddressesResult

    UnmonitorInstances(ctx workflow.Context, input *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error)
    UnmonitorInstancesAsync(ctx workflow.Context, input *ec2.UnmonitorInstancesInput) *Ec2UnmonitorInstancesResult

    UpdateSecurityGroupRuleDescriptionsEgress(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error)
    UpdateSecurityGroupRuleDescriptionsEgressAsync(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) *Ec2UpdateSecurityGroupRuleDescriptionsEgressResult

    UpdateSecurityGroupRuleDescriptionsIngress(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error)
    UpdateSecurityGroupRuleDescriptionsIngressAsync(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) *Ec2UpdateSecurityGroupRuleDescriptionsIngressResult

    WithdrawByoipCidr(ctx workflow.Context, input *ec2.WithdrawByoipCidrInput) (*ec2.WithdrawByoipCidrOutput, error)
    WithdrawByoipCidrAsync(ctx workflow.Context, input *ec2.WithdrawByoipCidrInput) *Ec2WithdrawByoipCidrResult

    WaitUntilBundleTaskComplete(ctx workflow.Context, input *ec2.DescribeBundleTasksInput) error
    WaitUntilConversionTaskCancelled(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) error
    WaitUntilConversionTaskCompleted(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) error
    WaitUntilConversionTaskDeleted(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) error
    WaitUntilCustomerGatewayAvailable(ctx workflow.Context, input *ec2.DescribeCustomerGatewaysInput) error
    WaitUntilExportTaskCancelled(ctx workflow.Context, input *ec2.DescribeExportTasksInput) error
    WaitUntilExportTaskCompleted(ctx workflow.Context, input *ec2.DescribeExportTasksInput) error
    WaitUntilImageAvailable(ctx workflow.Context, input *ec2.DescribeImagesInput) error
    WaitUntilImageExists(ctx workflow.Context, input *ec2.DescribeImagesInput) error
    WaitUntilInstanceExists(ctx workflow.Context, input *ec2.DescribeInstancesInput) error
    WaitUntilInstanceRunning(ctx workflow.Context, input *ec2.DescribeInstancesInput) error
    WaitUntilInstanceStatusOk(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) error
    WaitUntilInstanceStopped(ctx workflow.Context, input *ec2.DescribeInstancesInput) error
    WaitUntilInstanceTerminated(ctx workflow.Context, input *ec2.DescribeInstancesInput) error
    WaitUntilKeyPairExists(ctx workflow.Context, input *ec2.DescribeKeyPairsInput) error
    WaitUntilNatGatewayAvailable(ctx workflow.Context, input *ec2.DescribeNatGatewaysInput) error
    WaitUntilNetworkInterfaceAvailable(ctx workflow.Context, input *ec2.DescribeNetworkInterfacesInput) error
    WaitUntilPasswordDataAvailable(ctx workflow.Context, input *ec2.GetPasswordDataInput) error
    WaitUntilSecurityGroupExists(ctx workflow.Context, input *ec2.DescribeSecurityGroupsInput) error
    WaitUntilSnapshotCompleted(ctx workflow.Context, input *ec2.DescribeSnapshotsInput) error
    WaitUntilSpotInstanceRequestFulfilled(ctx workflow.Context, input *ec2.DescribeSpotInstanceRequestsInput) error
    WaitUntilSubnetAvailable(ctx workflow.Context, input *ec2.DescribeSubnetsInput) error
    WaitUntilSystemStatusOk(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) error
    WaitUntilVolumeAvailable(ctx workflow.Context, input *ec2.DescribeVolumesInput) error
    WaitUntilVolumeDeleted(ctx workflow.Context, input *ec2.DescribeVolumesInput) error
    WaitUntilVolumeInUse(ctx workflow.Context, input *ec2.DescribeVolumesInput) error
    WaitUntilVpcAvailable(ctx workflow.Context, input *ec2.DescribeVpcsInput) error
    WaitUntilVpcExists(ctx workflow.Context, input *ec2.DescribeVpcsInput) error
    WaitUntilVpcPeeringConnectionDeleted(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) error
    WaitUntilVpcPeeringConnectionExists(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) error
    WaitUntilVpnConnectionAvailable(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) error
    WaitUntilVpnConnectionDeleted(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) error}

type Ec2AcceptReservedInstancesExchangeQuoteResult struct {
	Result workflow.Future
}

func (r *Ec2AcceptReservedInstancesExchangeQuoteResult) Get(ctx workflow.Context) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
    var output ec2.AcceptReservedInstancesExchangeQuoteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AcceptTransitGatewayPeeringAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2AcceptTransitGatewayPeeringAttachmentResult) Get(ctx workflow.Context) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.AcceptTransitGatewayPeeringAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AcceptTransitGatewayVpcAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2AcceptTransitGatewayVpcAttachmentResult) Get(ctx workflow.Context) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.AcceptTransitGatewayVpcAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AcceptVpcEndpointConnectionsResult struct {
	Result workflow.Future
}

func (r *Ec2AcceptVpcEndpointConnectionsResult) Get(ctx workflow.Context) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
    var output ec2.AcceptVpcEndpointConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AcceptVpcPeeringConnectionResult struct {
	Result workflow.Future
}

func (r *Ec2AcceptVpcPeeringConnectionResult) Get(ctx workflow.Context) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
    var output ec2.AcceptVpcPeeringConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AdvertiseByoipCidrResult struct {
	Result workflow.Future
}

func (r *Ec2AdvertiseByoipCidrResult) Get(ctx workflow.Context) (*ec2.AdvertiseByoipCidrOutput, error) {
    var output ec2.AdvertiseByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AllocateAddressResult struct {
	Result workflow.Future
}

func (r *Ec2AllocateAddressResult) Get(ctx workflow.Context) (*ec2.AllocateAddressOutput, error) {
    var output ec2.AllocateAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AllocateHostsResult struct {
	Result workflow.Future
}

func (r *Ec2AllocateHostsResult) Get(ctx workflow.Context) (*ec2.AllocateHostsOutput, error) {
    var output ec2.AllocateHostsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ApplySecurityGroupsToClientVpnTargetNetworkResult struct {
	Result workflow.Future
}

func (r *Ec2ApplySecurityGroupsToClientVpnTargetNetworkResult) Get(ctx workflow.Context) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
    var output ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssignIpv6AddressesResult struct {
	Result workflow.Future
}

func (r *Ec2AssignIpv6AddressesResult) Get(ctx workflow.Context) (*ec2.AssignIpv6AddressesOutput, error) {
    var output ec2.AssignIpv6AddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssignPrivateIpAddressesResult struct {
	Result workflow.Future
}

func (r *Ec2AssignPrivateIpAddressesResult) Get(ctx workflow.Context) (*ec2.AssignPrivateIpAddressesOutput, error) {
    var output ec2.AssignPrivateIpAddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateAddressResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateAddressResult) Get(ctx workflow.Context) (*ec2.AssociateAddressOutput, error) {
    var output ec2.AssociateAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateClientVpnTargetNetworkResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateClientVpnTargetNetworkResult) Get(ctx workflow.Context) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
    var output ec2.AssociateClientVpnTargetNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateDhcpOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateDhcpOptionsResult) Get(ctx workflow.Context) (*ec2.AssociateDhcpOptionsOutput, error) {
    var output ec2.AssociateDhcpOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateIamInstanceProfileResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateIamInstanceProfileResult) Get(ctx workflow.Context) (*ec2.AssociateIamInstanceProfileOutput, error) {
    var output ec2.AssociateIamInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateRouteTableResult) Get(ctx workflow.Context) (*ec2.AssociateRouteTableOutput, error) {
    var output ec2.AssociateRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateSubnetCidrBlockResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateSubnetCidrBlockResult) Get(ctx workflow.Context) (*ec2.AssociateSubnetCidrBlockOutput, error) {
    var output ec2.AssociateSubnetCidrBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateTransitGatewayMulticastDomainResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateTransitGatewayMulticastDomainResult) Get(ctx workflow.Context) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
    var output ec2.AssociateTransitGatewayMulticastDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateTransitGatewayRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateTransitGatewayRouteTableResult) Get(ctx workflow.Context) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
    var output ec2.AssociateTransitGatewayRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AssociateVpcCidrBlockResult struct {
	Result workflow.Future
}

func (r *Ec2AssociateVpcCidrBlockResult) Get(ctx workflow.Context) (*ec2.AssociateVpcCidrBlockOutput, error) {
    var output ec2.AssociateVpcCidrBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AttachClassicLinkVpcResult struct {
	Result workflow.Future
}

func (r *Ec2AttachClassicLinkVpcResult) Get(ctx workflow.Context) (*ec2.AttachClassicLinkVpcOutput, error) {
    var output ec2.AttachClassicLinkVpcOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AttachInternetGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2AttachInternetGatewayResult) Get(ctx workflow.Context) (*ec2.AttachInternetGatewayOutput, error) {
    var output ec2.AttachInternetGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AttachNetworkInterfaceResult struct {
	Result workflow.Future
}

func (r *Ec2AttachNetworkInterfaceResult) Get(ctx workflow.Context) (*ec2.AttachNetworkInterfaceOutput, error) {
    var output ec2.AttachNetworkInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AttachVolumeResult struct {
	Result workflow.Future
}

func (r *Ec2AttachVolumeResult) Get(ctx workflow.Context) (*ec2.VolumeAttachment, error) {
    var output ec2.VolumeAttachment
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AttachVpnGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2AttachVpnGatewayResult) Get(ctx workflow.Context) (*ec2.AttachVpnGatewayOutput, error) {
    var output ec2.AttachVpnGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AuthorizeClientVpnIngressResult struct {
	Result workflow.Future
}

func (r *Ec2AuthorizeClientVpnIngressResult) Get(ctx workflow.Context) (*ec2.AuthorizeClientVpnIngressOutput, error) {
    var output ec2.AuthorizeClientVpnIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AuthorizeSecurityGroupEgressResult struct {
	Result workflow.Future
}

func (r *Ec2AuthorizeSecurityGroupEgressResult) Get(ctx workflow.Context) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
    var output ec2.AuthorizeSecurityGroupEgressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2AuthorizeSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *Ec2AuthorizeSecurityGroupIngressResult) Get(ctx workflow.Context) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
    var output ec2.AuthorizeSecurityGroupIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2BundleInstanceResult struct {
	Result workflow.Future
}

func (r *Ec2BundleInstanceResult) Get(ctx workflow.Context) (*ec2.BundleInstanceOutput, error) {
    var output ec2.BundleInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelBundleTaskResult struct {
	Result workflow.Future
}

func (r *Ec2CancelBundleTaskResult) Get(ctx workflow.Context) (*ec2.CancelBundleTaskOutput, error) {
    var output ec2.CancelBundleTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelCapacityReservationResult struct {
	Result workflow.Future
}

func (r *Ec2CancelCapacityReservationResult) Get(ctx workflow.Context) (*ec2.CancelCapacityReservationOutput, error) {
    var output ec2.CancelCapacityReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelConversionTaskResult struct {
	Result workflow.Future
}

func (r *Ec2CancelConversionTaskResult) Get(ctx workflow.Context) (*ec2.CancelConversionTaskOutput, error) {
    var output ec2.CancelConversionTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelExportTaskResult struct {
	Result workflow.Future
}

func (r *Ec2CancelExportTaskResult) Get(ctx workflow.Context) (*ec2.CancelExportTaskOutput, error) {
    var output ec2.CancelExportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelImportTaskResult struct {
	Result workflow.Future
}

func (r *Ec2CancelImportTaskResult) Get(ctx workflow.Context) (*ec2.CancelImportTaskOutput, error) {
    var output ec2.CancelImportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelReservedInstancesListingResult struct {
	Result workflow.Future
}

func (r *Ec2CancelReservedInstancesListingResult) Get(ctx workflow.Context) (*ec2.CancelReservedInstancesListingOutput, error) {
    var output ec2.CancelReservedInstancesListingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelSpotFleetRequestsResult struct {
	Result workflow.Future
}

func (r *Ec2CancelSpotFleetRequestsResult) Get(ctx workflow.Context) (*ec2.CancelSpotFleetRequestsOutput, error) {
    var output ec2.CancelSpotFleetRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CancelSpotInstanceRequestsResult struct {
	Result workflow.Future
}

func (r *Ec2CancelSpotInstanceRequestsResult) Get(ctx workflow.Context) (*ec2.CancelSpotInstanceRequestsOutput, error) {
    var output ec2.CancelSpotInstanceRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ConfirmProductInstanceResult struct {
	Result workflow.Future
}

func (r *Ec2ConfirmProductInstanceResult) Get(ctx workflow.Context) (*ec2.ConfirmProductInstanceOutput, error) {
    var output ec2.ConfirmProductInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CopyFpgaImageResult struct {
	Result workflow.Future
}

func (r *Ec2CopyFpgaImageResult) Get(ctx workflow.Context) (*ec2.CopyFpgaImageOutput, error) {
    var output ec2.CopyFpgaImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CopyImageResult struct {
	Result workflow.Future
}

func (r *Ec2CopyImageResult) Get(ctx workflow.Context) (*ec2.CopyImageOutput, error) {
    var output ec2.CopyImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CopySnapshotResult struct {
	Result workflow.Future
}

func (r *Ec2CopySnapshotResult) Get(ctx workflow.Context) (*ec2.CopySnapshotOutput, error) {
    var output ec2.CopySnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateCapacityReservationResult struct {
	Result workflow.Future
}

func (r *Ec2CreateCapacityReservationResult) Get(ctx workflow.Context) (*ec2.CreateCapacityReservationOutput, error) {
    var output ec2.CreateCapacityReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateCarrierGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2CreateCarrierGatewayResult) Get(ctx workflow.Context) (*ec2.CreateCarrierGatewayOutput, error) {
    var output ec2.CreateCarrierGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateClientVpnEndpointResult struct {
	Result workflow.Future
}

func (r *Ec2CreateClientVpnEndpointResult) Get(ctx workflow.Context) (*ec2.CreateClientVpnEndpointOutput, error) {
    var output ec2.CreateClientVpnEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateClientVpnRouteResult struct {
	Result workflow.Future
}

func (r *Ec2CreateClientVpnRouteResult) Get(ctx workflow.Context) (*ec2.CreateClientVpnRouteOutput, error) {
    var output ec2.CreateClientVpnRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateCustomerGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2CreateCustomerGatewayResult) Get(ctx workflow.Context) (*ec2.CreateCustomerGatewayOutput, error) {
    var output ec2.CreateCustomerGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateDefaultSubnetResult struct {
	Result workflow.Future
}

func (r *Ec2CreateDefaultSubnetResult) Get(ctx workflow.Context) (*ec2.CreateDefaultSubnetOutput, error) {
    var output ec2.CreateDefaultSubnetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateDefaultVpcResult struct {
	Result workflow.Future
}

func (r *Ec2CreateDefaultVpcResult) Get(ctx workflow.Context) (*ec2.CreateDefaultVpcOutput, error) {
    var output ec2.CreateDefaultVpcOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateDhcpOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2CreateDhcpOptionsResult) Get(ctx workflow.Context) (*ec2.CreateDhcpOptionsOutput, error) {
    var output ec2.CreateDhcpOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateEgressOnlyInternetGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2CreateEgressOnlyInternetGatewayResult) Get(ctx workflow.Context) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
    var output ec2.CreateEgressOnlyInternetGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateFleetResult struct {
	Result workflow.Future
}

func (r *Ec2CreateFleetResult) Get(ctx workflow.Context) (*ec2.CreateFleetOutput, error) {
    var output ec2.CreateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateFlowLogsResult struct {
	Result workflow.Future
}

func (r *Ec2CreateFlowLogsResult) Get(ctx workflow.Context) (*ec2.CreateFlowLogsOutput, error) {
    var output ec2.CreateFlowLogsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateFpgaImageResult struct {
	Result workflow.Future
}

func (r *Ec2CreateFpgaImageResult) Get(ctx workflow.Context) (*ec2.CreateFpgaImageOutput, error) {
    var output ec2.CreateFpgaImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateImageResult struct {
	Result workflow.Future
}

func (r *Ec2CreateImageResult) Get(ctx workflow.Context) (*ec2.CreateImageOutput, error) {
    var output ec2.CreateImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateInstanceExportTaskResult struct {
	Result workflow.Future
}

func (r *Ec2CreateInstanceExportTaskResult) Get(ctx workflow.Context) (*ec2.CreateInstanceExportTaskOutput, error) {
    var output ec2.CreateInstanceExportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateInternetGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2CreateInternetGatewayResult) Get(ctx workflow.Context) (*ec2.CreateInternetGatewayOutput, error) {
    var output ec2.CreateInternetGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateKeyPairResult struct {
	Result workflow.Future
}

func (r *Ec2CreateKeyPairResult) Get(ctx workflow.Context) (*ec2.CreateKeyPairOutput, error) {
    var output ec2.CreateKeyPairOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateLaunchTemplateResult struct {
	Result workflow.Future
}

func (r *Ec2CreateLaunchTemplateResult) Get(ctx workflow.Context) (*ec2.CreateLaunchTemplateOutput, error) {
    var output ec2.CreateLaunchTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateLaunchTemplateVersionResult struct {
	Result workflow.Future
}

func (r *Ec2CreateLaunchTemplateVersionResult) Get(ctx workflow.Context) (*ec2.CreateLaunchTemplateVersionOutput, error) {
    var output ec2.CreateLaunchTemplateVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateLocalGatewayRouteResult struct {
	Result workflow.Future
}

func (r *Ec2CreateLocalGatewayRouteResult) Get(ctx workflow.Context) (*ec2.CreateLocalGatewayRouteOutput, error) {
    var output ec2.CreateLocalGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateLocalGatewayRouteTableVpcAssociationResult struct {
	Result workflow.Future
}

func (r *Ec2CreateLocalGatewayRouteTableVpcAssociationResult) Get(ctx workflow.Context) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
    var output ec2.CreateLocalGatewayRouteTableVpcAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateManagedPrefixListResult struct {
	Result workflow.Future
}

func (r *Ec2CreateManagedPrefixListResult) Get(ctx workflow.Context) (*ec2.CreateManagedPrefixListOutput, error) {
    var output ec2.CreateManagedPrefixListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateNatGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2CreateNatGatewayResult) Get(ctx workflow.Context) (*ec2.CreateNatGatewayOutput, error) {
    var output ec2.CreateNatGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateNetworkAclResult struct {
	Result workflow.Future
}

func (r *Ec2CreateNetworkAclResult) Get(ctx workflow.Context) (*ec2.CreateNetworkAclOutput, error) {
    var output ec2.CreateNetworkAclOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateNetworkAclEntryResult struct {
	Result workflow.Future
}

func (r *Ec2CreateNetworkAclEntryResult) Get(ctx workflow.Context) (*ec2.CreateNetworkAclEntryOutput, error) {
    var output ec2.CreateNetworkAclEntryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateNetworkInterfaceResult struct {
	Result workflow.Future
}

func (r *Ec2CreateNetworkInterfaceResult) Get(ctx workflow.Context) (*ec2.CreateNetworkInterfaceOutput, error) {
    var output ec2.CreateNetworkInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateNetworkInterfacePermissionResult struct {
	Result workflow.Future
}

func (r *Ec2CreateNetworkInterfacePermissionResult) Get(ctx workflow.Context) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
    var output ec2.CreateNetworkInterfacePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreatePlacementGroupResult struct {
	Result workflow.Future
}

func (r *Ec2CreatePlacementGroupResult) Get(ctx workflow.Context) (*ec2.CreatePlacementGroupOutput, error) {
    var output ec2.CreatePlacementGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateReservedInstancesListingResult struct {
	Result workflow.Future
}

func (r *Ec2CreateReservedInstancesListingResult) Get(ctx workflow.Context) (*ec2.CreateReservedInstancesListingOutput, error) {
    var output ec2.CreateReservedInstancesListingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateRouteResult struct {
	Result workflow.Future
}

func (r *Ec2CreateRouteResult) Get(ctx workflow.Context) (*ec2.CreateRouteOutput, error) {
    var output ec2.CreateRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2CreateRouteTableResult) Get(ctx workflow.Context) (*ec2.CreateRouteTableOutput, error) {
    var output ec2.CreateRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateSecurityGroupResult struct {
	Result workflow.Future
}

func (r *Ec2CreateSecurityGroupResult) Get(ctx workflow.Context) (*ec2.CreateSecurityGroupOutput, error) {
    var output ec2.CreateSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateSnapshotResult struct {
	Result workflow.Future
}

func (r *Ec2CreateSnapshotResult) Get(ctx workflow.Context) (*ec2.Snapshot, error) {
    var output ec2.Snapshot
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateSnapshotsResult struct {
	Result workflow.Future
}

func (r *Ec2CreateSnapshotsResult) Get(ctx workflow.Context) (*ec2.CreateSnapshotsOutput, error) {
    var output ec2.CreateSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateSpotDatafeedSubscriptionResult struct {
	Result workflow.Future
}

func (r *Ec2CreateSpotDatafeedSubscriptionResult) Get(ctx workflow.Context) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
    var output ec2.CreateSpotDatafeedSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateSubnetResult struct {
	Result workflow.Future
}

func (r *Ec2CreateSubnetResult) Get(ctx workflow.Context) (*ec2.CreateSubnetOutput, error) {
    var output ec2.CreateSubnetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTagsResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTagsResult) Get(ctx workflow.Context) (*ec2.CreateTagsOutput, error) {
    var output ec2.CreateTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTrafficMirrorFilterResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTrafficMirrorFilterResult) Get(ctx workflow.Context) (*ec2.CreateTrafficMirrorFilterOutput, error) {
    var output ec2.CreateTrafficMirrorFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTrafficMirrorFilterRuleResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTrafficMirrorFilterRuleResult) Get(ctx workflow.Context) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
    var output ec2.CreateTrafficMirrorFilterRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTrafficMirrorSessionResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTrafficMirrorSessionResult) Get(ctx workflow.Context) (*ec2.CreateTrafficMirrorSessionOutput, error) {
    var output ec2.CreateTrafficMirrorSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTrafficMirrorTargetResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTrafficMirrorTargetResult) Get(ctx workflow.Context) (*ec2.CreateTrafficMirrorTargetOutput, error) {
    var output ec2.CreateTrafficMirrorTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTransitGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTransitGatewayResult) Get(ctx workflow.Context) (*ec2.CreateTransitGatewayOutput, error) {
    var output ec2.CreateTransitGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTransitGatewayMulticastDomainResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTransitGatewayMulticastDomainResult) Get(ctx workflow.Context) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
    var output ec2.CreateTransitGatewayMulticastDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTransitGatewayPeeringAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTransitGatewayPeeringAttachmentResult) Get(ctx workflow.Context) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.CreateTransitGatewayPeeringAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTransitGatewayPrefixListReferenceResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTransitGatewayPrefixListReferenceResult) Get(ctx workflow.Context) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
    var output ec2.CreateTransitGatewayPrefixListReferenceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTransitGatewayRouteResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTransitGatewayRouteResult) Get(ctx workflow.Context) (*ec2.CreateTransitGatewayRouteOutput, error) {
    var output ec2.CreateTransitGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTransitGatewayRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTransitGatewayRouteTableResult) Get(ctx workflow.Context) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
    var output ec2.CreateTransitGatewayRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateTransitGatewayVpcAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2CreateTransitGatewayVpcAttachmentResult) Get(ctx workflow.Context) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.CreateTransitGatewayVpcAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVolumeResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVolumeResult) Get(ctx workflow.Context) (*ec2.Volume, error) {
    var output ec2.Volume
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpcResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpcResult) Get(ctx workflow.Context) (*ec2.CreateVpcOutput, error) {
    var output ec2.CreateVpcOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpcEndpointResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpcEndpointResult) Get(ctx workflow.Context) (*ec2.CreateVpcEndpointOutput, error) {
    var output ec2.CreateVpcEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpcEndpointConnectionNotificationResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpcEndpointConnectionNotificationResult) Get(ctx workflow.Context) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
    var output ec2.CreateVpcEndpointConnectionNotificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpcEndpointServiceConfigurationResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpcEndpointServiceConfigurationResult) Get(ctx workflow.Context) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
    var output ec2.CreateVpcEndpointServiceConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpcPeeringConnectionResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpcPeeringConnectionResult) Get(ctx workflow.Context) (*ec2.CreateVpcPeeringConnectionOutput, error) {
    var output ec2.CreateVpcPeeringConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpnConnectionResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpnConnectionResult) Get(ctx workflow.Context) (*ec2.CreateVpnConnectionOutput, error) {
    var output ec2.CreateVpnConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpnConnectionRouteResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpnConnectionRouteResult) Get(ctx workflow.Context) (*ec2.CreateVpnConnectionRouteOutput, error) {
    var output ec2.CreateVpnConnectionRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2CreateVpnGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2CreateVpnGatewayResult) Get(ctx workflow.Context) (*ec2.CreateVpnGatewayOutput, error) {
    var output ec2.CreateVpnGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteCarrierGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteCarrierGatewayResult) Get(ctx workflow.Context) (*ec2.DeleteCarrierGatewayOutput, error) {
    var output ec2.DeleteCarrierGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteClientVpnEndpointResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteClientVpnEndpointResult) Get(ctx workflow.Context) (*ec2.DeleteClientVpnEndpointOutput, error) {
    var output ec2.DeleteClientVpnEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteClientVpnRouteResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteClientVpnRouteResult) Get(ctx workflow.Context) (*ec2.DeleteClientVpnRouteOutput, error) {
    var output ec2.DeleteClientVpnRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteCustomerGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteCustomerGatewayResult) Get(ctx workflow.Context) (*ec2.DeleteCustomerGatewayOutput, error) {
    var output ec2.DeleteCustomerGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteDhcpOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteDhcpOptionsResult) Get(ctx workflow.Context) (*ec2.DeleteDhcpOptionsOutput, error) {
    var output ec2.DeleteDhcpOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteEgressOnlyInternetGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteEgressOnlyInternetGatewayResult) Get(ctx workflow.Context) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
    var output ec2.DeleteEgressOnlyInternetGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteFleetsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteFleetsResult) Get(ctx workflow.Context) (*ec2.DeleteFleetsOutput, error) {
    var output ec2.DeleteFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteFlowLogsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteFlowLogsResult) Get(ctx workflow.Context) (*ec2.DeleteFlowLogsOutput, error) {
    var output ec2.DeleteFlowLogsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteFpgaImageResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteFpgaImageResult) Get(ctx workflow.Context) (*ec2.DeleteFpgaImageOutput, error) {
    var output ec2.DeleteFpgaImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteInternetGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteInternetGatewayResult) Get(ctx workflow.Context) (*ec2.DeleteInternetGatewayOutput, error) {
    var output ec2.DeleteInternetGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteKeyPairResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteKeyPairResult) Get(ctx workflow.Context) (*ec2.DeleteKeyPairOutput, error) {
    var output ec2.DeleteKeyPairOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteLaunchTemplateResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteLaunchTemplateResult) Get(ctx workflow.Context) (*ec2.DeleteLaunchTemplateOutput, error) {
    var output ec2.DeleteLaunchTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteLaunchTemplateVersionsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteLaunchTemplateVersionsResult) Get(ctx workflow.Context) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
    var output ec2.DeleteLaunchTemplateVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteLocalGatewayRouteResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteLocalGatewayRouteResult) Get(ctx workflow.Context) (*ec2.DeleteLocalGatewayRouteOutput, error) {
    var output ec2.DeleteLocalGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteLocalGatewayRouteTableVpcAssociationResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteLocalGatewayRouteTableVpcAssociationResult) Get(ctx workflow.Context) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
    var output ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteManagedPrefixListResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteManagedPrefixListResult) Get(ctx workflow.Context) (*ec2.DeleteManagedPrefixListOutput, error) {
    var output ec2.DeleteManagedPrefixListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteNatGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteNatGatewayResult) Get(ctx workflow.Context) (*ec2.DeleteNatGatewayOutput, error) {
    var output ec2.DeleteNatGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteNetworkAclResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteNetworkAclResult) Get(ctx workflow.Context) (*ec2.DeleteNetworkAclOutput, error) {
    var output ec2.DeleteNetworkAclOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteNetworkAclEntryResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteNetworkAclEntryResult) Get(ctx workflow.Context) (*ec2.DeleteNetworkAclEntryOutput, error) {
    var output ec2.DeleteNetworkAclEntryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteNetworkInterfaceResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteNetworkInterfaceResult) Get(ctx workflow.Context) (*ec2.DeleteNetworkInterfaceOutput, error) {
    var output ec2.DeleteNetworkInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteNetworkInterfacePermissionResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteNetworkInterfacePermissionResult) Get(ctx workflow.Context) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
    var output ec2.DeleteNetworkInterfacePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeletePlacementGroupResult struct {
	Result workflow.Future
}

func (r *Ec2DeletePlacementGroupResult) Get(ctx workflow.Context) (*ec2.DeletePlacementGroupOutput, error) {
    var output ec2.DeletePlacementGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteQueuedReservedInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteQueuedReservedInstancesResult) Get(ctx workflow.Context) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
    var output ec2.DeleteQueuedReservedInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteRouteResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteRouteResult) Get(ctx workflow.Context) (*ec2.DeleteRouteOutput, error) {
    var output ec2.DeleteRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteRouteTableResult) Get(ctx workflow.Context) (*ec2.DeleteRouteTableOutput, error) {
    var output ec2.DeleteRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteSecurityGroupResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteSecurityGroupResult) Get(ctx workflow.Context) (*ec2.DeleteSecurityGroupOutput, error) {
    var output ec2.DeleteSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteSnapshotResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteSnapshotResult) Get(ctx workflow.Context) (*ec2.DeleteSnapshotOutput, error) {
    var output ec2.DeleteSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteSpotDatafeedSubscriptionResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteSpotDatafeedSubscriptionResult) Get(ctx workflow.Context) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
    var output ec2.DeleteSpotDatafeedSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteSubnetResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteSubnetResult) Get(ctx workflow.Context) (*ec2.DeleteSubnetOutput, error) {
    var output ec2.DeleteSubnetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTagsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTagsResult) Get(ctx workflow.Context) (*ec2.DeleteTagsOutput, error) {
    var output ec2.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTrafficMirrorFilterResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTrafficMirrorFilterResult) Get(ctx workflow.Context) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
    var output ec2.DeleteTrafficMirrorFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTrafficMirrorFilterRuleResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTrafficMirrorFilterRuleResult) Get(ctx workflow.Context) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
    var output ec2.DeleteTrafficMirrorFilterRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTrafficMirrorSessionResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTrafficMirrorSessionResult) Get(ctx workflow.Context) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
    var output ec2.DeleteTrafficMirrorSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTrafficMirrorTargetResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTrafficMirrorTargetResult) Get(ctx workflow.Context) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
    var output ec2.DeleteTrafficMirrorTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTransitGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTransitGatewayResult) Get(ctx workflow.Context) (*ec2.DeleteTransitGatewayOutput, error) {
    var output ec2.DeleteTransitGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTransitGatewayMulticastDomainResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTransitGatewayMulticastDomainResult) Get(ctx workflow.Context) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
    var output ec2.DeleteTransitGatewayMulticastDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTransitGatewayPeeringAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTransitGatewayPeeringAttachmentResult) Get(ctx workflow.Context) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.DeleteTransitGatewayPeeringAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTransitGatewayPrefixListReferenceResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTransitGatewayPrefixListReferenceResult) Get(ctx workflow.Context) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
    var output ec2.DeleteTransitGatewayPrefixListReferenceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTransitGatewayRouteResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTransitGatewayRouteResult) Get(ctx workflow.Context) (*ec2.DeleteTransitGatewayRouteOutput, error) {
    var output ec2.DeleteTransitGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTransitGatewayRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTransitGatewayRouteTableResult) Get(ctx workflow.Context) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
    var output ec2.DeleteTransitGatewayRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteTransitGatewayVpcAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteTransitGatewayVpcAttachmentResult) Get(ctx workflow.Context) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.DeleteTransitGatewayVpcAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVolumeResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVolumeResult) Get(ctx workflow.Context) (*ec2.DeleteVolumeOutput, error) {
    var output ec2.DeleteVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpcResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpcResult) Get(ctx workflow.Context) (*ec2.DeleteVpcOutput, error) {
    var output ec2.DeleteVpcOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpcEndpointConnectionNotificationsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpcEndpointConnectionNotificationsResult) Get(ctx workflow.Context) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
    var output ec2.DeleteVpcEndpointConnectionNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpcEndpointServiceConfigurationsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpcEndpointServiceConfigurationsResult) Get(ctx workflow.Context) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
    var output ec2.DeleteVpcEndpointServiceConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpcEndpointsResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpcEndpointsResult) Get(ctx workflow.Context) (*ec2.DeleteVpcEndpointsOutput, error) {
    var output ec2.DeleteVpcEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpcPeeringConnectionResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpcPeeringConnectionResult) Get(ctx workflow.Context) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
    var output ec2.DeleteVpcPeeringConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpnConnectionResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpnConnectionResult) Get(ctx workflow.Context) (*ec2.DeleteVpnConnectionOutput, error) {
    var output ec2.DeleteVpnConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpnConnectionRouteResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpnConnectionRouteResult) Get(ctx workflow.Context) (*ec2.DeleteVpnConnectionRouteOutput, error) {
    var output ec2.DeleteVpnConnectionRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeleteVpnGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DeleteVpnGatewayResult) Get(ctx workflow.Context) (*ec2.DeleteVpnGatewayOutput, error) {
    var output ec2.DeleteVpnGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeprovisionByoipCidrResult struct {
	Result workflow.Future
}

func (r *Ec2DeprovisionByoipCidrResult) Get(ctx workflow.Context) (*ec2.DeprovisionByoipCidrOutput, error) {
    var output ec2.DeprovisionByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeregisterImageResult struct {
	Result workflow.Future
}

func (r *Ec2DeregisterImageResult) Get(ctx workflow.Context) (*ec2.DeregisterImageOutput, error) {
    var output ec2.DeregisterImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeregisterInstanceEventNotificationAttributesResult struct {
	Result workflow.Future
}

func (r *Ec2DeregisterInstanceEventNotificationAttributesResult) Get(ctx workflow.Context) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
    var output ec2.DeregisterInstanceEventNotificationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeregisterTransitGatewayMulticastGroupMembersResult struct {
	Result workflow.Future
}

func (r *Ec2DeregisterTransitGatewayMulticastGroupMembersResult) Get(ctx workflow.Context) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
    var output ec2.DeregisterTransitGatewayMulticastGroupMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DeregisterTransitGatewayMulticastGroupSourcesResult struct {
	Result workflow.Future
}

func (r *Ec2DeregisterTransitGatewayMulticastGroupSourcesResult) Get(ctx workflow.Context) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
    var output ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeAccountAttributesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeAccountAttributesResult) Get(ctx workflow.Context) (*ec2.DescribeAccountAttributesOutput, error) {
    var output ec2.DescribeAccountAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeAddressesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeAddressesResult) Get(ctx workflow.Context) (*ec2.DescribeAddressesOutput, error) {
    var output ec2.DescribeAddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeAggregateIdFormatResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeAggregateIdFormatResult) Get(ctx workflow.Context) (*ec2.DescribeAggregateIdFormatOutput, error) {
    var output ec2.DescribeAggregateIdFormatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeAvailabilityZonesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeAvailabilityZonesResult) Get(ctx workflow.Context) (*ec2.DescribeAvailabilityZonesOutput, error) {
    var output ec2.DescribeAvailabilityZonesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeBundleTasksResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeBundleTasksResult) Get(ctx workflow.Context) (*ec2.DescribeBundleTasksOutput, error) {
    var output ec2.DescribeBundleTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeByoipCidrsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeByoipCidrsResult) Get(ctx workflow.Context) (*ec2.DescribeByoipCidrsOutput, error) {
    var output ec2.DescribeByoipCidrsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeCapacityReservationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeCapacityReservationsResult) Get(ctx workflow.Context) (*ec2.DescribeCapacityReservationsOutput, error) {
    var output ec2.DescribeCapacityReservationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeCarrierGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeCarrierGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeCarrierGatewaysOutput, error) {
    var output ec2.DescribeCarrierGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeClassicLinkInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeClassicLinkInstancesResult) Get(ctx workflow.Context) (*ec2.DescribeClassicLinkInstancesOutput, error) {
    var output ec2.DescribeClassicLinkInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeClientVpnAuthorizationRulesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeClientVpnAuthorizationRulesResult) Get(ctx workflow.Context) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
    var output ec2.DescribeClientVpnAuthorizationRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeClientVpnConnectionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeClientVpnConnectionsResult) Get(ctx workflow.Context) (*ec2.DescribeClientVpnConnectionsOutput, error) {
    var output ec2.DescribeClientVpnConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeClientVpnEndpointsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeClientVpnEndpointsResult) Get(ctx workflow.Context) (*ec2.DescribeClientVpnEndpointsOutput, error) {
    var output ec2.DescribeClientVpnEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeClientVpnRoutesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeClientVpnRoutesResult) Get(ctx workflow.Context) (*ec2.DescribeClientVpnRoutesOutput, error) {
    var output ec2.DescribeClientVpnRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeClientVpnTargetNetworksResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeClientVpnTargetNetworksResult) Get(ctx workflow.Context) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
    var output ec2.DescribeClientVpnTargetNetworksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeCoipPoolsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeCoipPoolsResult) Get(ctx workflow.Context) (*ec2.DescribeCoipPoolsOutput, error) {
    var output ec2.DescribeCoipPoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeConversionTasksResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeConversionTasksResult) Get(ctx workflow.Context) (*ec2.DescribeConversionTasksOutput, error) {
    var output ec2.DescribeConversionTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeCustomerGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeCustomerGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeCustomerGatewaysOutput, error) {
    var output ec2.DescribeCustomerGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeDhcpOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeDhcpOptionsResult) Get(ctx workflow.Context) (*ec2.DescribeDhcpOptionsOutput, error) {
    var output ec2.DescribeDhcpOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeEgressOnlyInternetGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeEgressOnlyInternetGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
    var output ec2.DescribeEgressOnlyInternetGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeElasticGpusResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeElasticGpusResult) Get(ctx workflow.Context) (*ec2.DescribeElasticGpusOutput, error) {
    var output ec2.DescribeElasticGpusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeExportImageTasksResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeExportImageTasksResult) Get(ctx workflow.Context) (*ec2.DescribeExportImageTasksOutput, error) {
    var output ec2.DescribeExportImageTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeExportTasksResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeExportTasksResult) Get(ctx workflow.Context) (*ec2.DescribeExportTasksOutput, error) {
    var output ec2.DescribeExportTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeFastSnapshotRestoresResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeFastSnapshotRestoresResult) Get(ctx workflow.Context) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
    var output ec2.DescribeFastSnapshotRestoresOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeFleetHistoryResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeFleetHistoryResult) Get(ctx workflow.Context) (*ec2.DescribeFleetHistoryOutput, error) {
    var output ec2.DescribeFleetHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeFleetInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeFleetInstancesResult) Get(ctx workflow.Context) (*ec2.DescribeFleetInstancesOutput, error) {
    var output ec2.DescribeFleetInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeFleetsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeFleetsResult) Get(ctx workflow.Context) (*ec2.DescribeFleetsOutput, error) {
    var output ec2.DescribeFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeFlowLogsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeFlowLogsResult) Get(ctx workflow.Context) (*ec2.DescribeFlowLogsOutput, error) {
    var output ec2.DescribeFlowLogsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeFpgaImageAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeFpgaImageAttributeResult) Get(ctx workflow.Context) (*ec2.DescribeFpgaImageAttributeOutput, error) {
    var output ec2.DescribeFpgaImageAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeFpgaImagesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeFpgaImagesResult) Get(ctx workflow.Context) (*ec2.DescribeFpgaImagesOutput, error) {
    var output ec2.DescribeFpgaImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeHostReservationOfferingsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeHostReservationOfferingsResult) Get(ctx workflow.Context) (*ec2.DescribeHostReservationOfferingsOutput, error) {
    var output ec2.DescribeHostReservationOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeHostReservationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeHostReservationsResult) Get(ctx workflow.Context) (*ec2.DescribeHostReservationsOutput, error) {
    var output ec2.DescribeHostReservationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeHostsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeHostsResult) Get(ctx workflow.Context) (*ec2.DescribeHostsOutput, error) {
    var output ec2.DescribeHostsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeIamInstanceProfileAssociationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeIamInstanceProfileAssociationsResult) Get(ctx workflow.Context) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
    var output ec2.DescribeIamInstanceProfileAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeIdFormatResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeIdFormatResult) Get(ctx workflow.Context) (*ec2.DescribeIdFormatOutput, error) {
    var output ec2.DescribeIdFormatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeIdentityIdFormatResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeIdentityIdFormatResult) Get(ctx workflow.Context) (*ec2.DescribeIdentityIdFormatOutput, error) {
    var output ec2.DescribeIdentityIdFormatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeImageAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeImageAttributeResult) Get(ctx workflow.Context) (*ec2.DescribeImageAttributeOutput, error) {
    var output ec2.DescribeImageAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeImagesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeImagesResult) Get(ctx workflow.Context) (*ec2.DescribeImagesOutput, error) {
    var output ec2.DescribeImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeImportImageTasksResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeImportImageTasksResult) Get(ctx workflow.Context) (*ec2.DescribeImportImageTasksOutput, error) {
    var output ec2.DescribeImportImageTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeImportSnapshotTasksResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeImportSnapshotTasksResult) Get(ctx workflow.Context) (*ec2.DescribeImportSnapshotTasksOutput, error) {
    var output ec2.DescribeImportSnapshotTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInstanceAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInstanceAttributeResult) Get(ctx workflow.Context) (*ec2.DescribeInstanceAttributeOutput, error) {
    var output ec2.DescribeInstanceAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInstanceCreditSpecificationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInstanceCreditSpecificationsResult) Get(ctx workflow.Context) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
    var output ec2.DescribeInstanceCreditSpecificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInstanceEventNotificationAttributesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInstanceEventNotificationAttributesResult) Get(ctx workflow.Context) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
    var output ec2.DescribeInstanceEventNotificationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInstanceStatusResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInstanceStatusResult) Get(ctx workflow.Context) (*ec2.DescribeInstanceStatusOutput, error) {
    var output ec2.DescribeInstanceStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInstanceTypeOfferingsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInstanceTypeOfferingsResult) Get(ctx workflow.Context) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
    var output ec2.DescribeInstanceTypeOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInstanceTypesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInstanceTypesResult) Get(ctx workflow.Context) (*ec2.DescribeInstanceTypesOutput, error) {
    var output ec2.DescribeInstanceTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInstancesResult) Get(ctx workflow.Context) (*ec2.DescribeInstancesOutput, error) {
    var output ec2.DescribeInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeInternetGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeInternetGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeInternetGatewaysOutput, error) {
    var output ec2.DescribeInternetGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeIpv6PoolsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeIpv6PoolsResult) Get(ctx workflow.Context) (*ec2.DescribeIpv6PoolsOutput, error) {
    var output ec2.DescribeIpv6PoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeKeyPairsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeKeyPairsResult) Get(ctx workflow.Context) (*ec2.DescribeKeyPairsOutput, error) {
    var output ec2.DescribeKeyPairsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLaunchTemplateVersionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLaunchTemplateVersionsResult) Get(ctx workflow.Context) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
    var output ec2.DescribeLaunchTemplateVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLaunchTemplatesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLaunchTemplatesResult) Get(ctx workflow.Context) (*ec2.DescribeLaunchTemplatesOutput, error) {
    var output ec2.DescribeLaunchTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResult) Get(ctx workflow.Context) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
    var output ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLocalGatewayRouteTableVpcAssociationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLocalGatewayRouteTableVpcAssociationsResult) Get(ctx workflow.Context) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
    var output ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLocalGatewayRouteTablesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLocalGatewayRouteTablesResult) Get(ctx workflow.Context) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
    var output ec2.DescribeLocalGatewayRouteTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLocalGatewayVirtualInterfaceGroupsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLocalGatewayVirtualInterfaceGroupsResult) Get(ctx workflow.Context) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
    var output ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLocalGatewayVirtualInterfacesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLocalGatewayVirtualInterfacesResult) Get(ctx workflow.Context) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
    var output ec2.DescribeLocalGatewayVirtualInterfacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeLocalGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeLocalGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeLocalGatewaysOutput, error) {
    var output ec2.DescribeLocalGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeManagedPrefixListsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeManagedPrefixListsResult) Get(ctx workflow.Context) (*ec2.DescribeManagedPrefixListsOutput, error) {
    var output ec2.DescribeManagedPrefixListsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeMovingAddressesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeMovingAddressesResult) Get(ctx workflow.Context) (*ec2.DescribeMovingAddressesOutput, error) {
    var output ec2.DescribeMovingAddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeNatGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeNatGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeNatGatewaysOutput, error) {
    var output ec2.DescribeNatGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeNetworkAclsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeNetworkAclsResult) Get(ctx workflow.Context) (*ec2.DescribeNetworkAclsOutput, error) {
    var output ec2.DescribeNetworkAclsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeNetworkInterfaceAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeNetworkInterfaceAttributeResult) Get(ctx workflow.Context) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
    var output ec2.DescribeNetworkInterfaceAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeNetworkInterfacePermissionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeNetworkInterfacePermissionsResult) Get(ctx workflow.Context) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
    var output ec2.DescribeNetworkInterfacePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeNetworkInterfacesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeNetworkInterfacesResult) Get(ctx workflow.Context) (*ec2.DescribeNetworkInterfacesOutput, error) {
    var output ec2.DescribeNetworkInterfacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribePlacementGroupsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribePlacementGroupsResult) Get(ctx workflow.Context) (*ec2.DescribePlacementGroupsOutput, error) {
    var output ec2.DescribePlacementGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribePrefixListsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribePrefixListsResult) Get(ctx workflow.Context) (*ec2.DescribePrefixListsOutput, error) {
    var output ec2.DescribePrefixListsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribePrincipalIdFormatResult struct {
	Result workflow.Future
}

func (r *Ec2DescribePrincipalIdFormatResult) Get(ctx workflow.Context) (*ec2.DescribePrincipalIdFormatOutput, error) {
    var output ec2.DescribePrincipalIdFormatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribePublicIpv4PoolsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribePublicIpv4PoolsResult) Get(ctx workflow.Context) (*ec2.DescribePublicIpv4PoolsOutput, error) {
    var output ec2.DescribePublicIpv4PoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeRegionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeRegionsResult) Get(ctx workflow.Context) (*ec2.DescribeRegionsOutput, error) {
    var output ec2.DescribeRegionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeReservedInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeReservedInstancesResult) Get(ctx workflow.Context) (*ec2.DescribeReservedInstancesOutput, error) {
    var output ec2.DescribeReservedInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeReservedInstancesListingsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeReservedInstancesListingsResult) Get(ctx workflow.Context) (*ec2.DescribeReservedInstancesListingsOutput, error) {
    var output ec2.DescribeReservedInstancesListingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeReservedInstancesModificationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeReservedInstancesModificationsResult) Get(ctx workflow.Context) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
    var output ec2.DescribeReservedInstancesModificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeReservedInstancesOfferingsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeReservedInstancesOfferingsResult) Get(ctx workflow.Context) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
    var output ec2.DescribeReservedInstancesOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeRouteTablesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeRouteTablesResult) Get(ctx workflow.Context) (*ec2.DescribeRouteTablesOutput, error) {
    var output ec2.DescribeRouteTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeScheduledInstanceAvailabilityResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeScheduledInstanceAvailabilityResult) Get(ctx workflow.Context) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
    var output ec2.DescribeScheduledInstanceAvailabilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeScheduledInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeScheduledInstancesResult) Get(ctx workflow.Context) (*ec2.DescribeScheduledInstancesOutput, error) {
    var output ec2.DescribeScheduledInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSecurityGroupReferencesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSecurityGroupReferencesResult) Get(ctx workflow.Context) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
    var output ec2.DescribeSecurityGroupReferencesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSecurityGroupsResult) Get(ctx workflow.Context) (*ec2.DescribeSecurityGroupsOutput, error) {
    var output ec2.DescribeSecurityGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSnapshotAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSnapshotAttributeResult) Get(ctx workflow.Context) (*ec2.DescribeSnapshotAttributeOutput, error) {
    var output ec2.DescribeSnapshotAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSnapshotsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSnapshotsResult) Get(ctx workflow.Context) (*ec2.DescribeSnapshotsOutput, error) {
    var output ec2.DescribeSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSpotDatafeedSubscriptionResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSpotDatafeedSubscriptionResult) Get(ctx workflow.Context) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
    var output ec2.DescribeSpotDatafeedSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSpotFleetInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSpotFleetInstancesResult) Get(ctx workflow.Context) (*ec2.DescribeSpotFleetInstancesOutput, error) {
    var output ec2.DescribeSpotFleetInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSpotFleetRequestHistoryResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSpotFleetRequestHistoryResult) Get(ctx workflow.Context) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
    var output ec2.DescribeSpotFleetRequestHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSpotFleetRequestsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSpotFleetRequestsResult) Get(ctx workflow.Context) (*ec2.DescribeSpotFleetRequestsOutput, error) {
    var output ec2.DescribeSpotFleetRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSpotInstanceRequestsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSpotInstanceRequestsResult) Get(ctx workflow.Context) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
    var output ec2.DescribeSpotInstanceRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSpotPriceHistoryResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSpotPriceHistoryResult) Get(ctx workflow.Context) (*ec2.DescribeSpotPriceHistoryOutput, error) {
    var output ec2.DescribeSpotPriceHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeStaleSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeStaleSecurityGroupsResult) Get(ctx workflow.Context) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
    var output ec2.DescribeStaleSecurityGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeSubnetsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeSubnetsResult) Get(ctx workflow.Context) (*ec2.DescribeSubnetsOutput, error) {
    var output ec2.DescribeSubnetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTagsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTagsResult) Get(ctx workflow.Context) (*ec2.DescribeTagsOutput, error) {
    var output ec2.DescribeTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTrafficMirrorFiltersResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTrafficMirrorFiltersResult) Get(ctx workflow.Context) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
    var output ec2.DescribeTrafficMirrorFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTrafficMirrorSessionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTrafficMirrorSessionsResult) Get(ctx workflow.Context) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
    var output ec2.DescribeTrafficMirrorSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTrafficMirrorTargetsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTrafficMirrorTargetsResult) Get(ctx workflow.Context) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
    var output ec2.DescribeTrafficMirrorTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTransitGatewayAttachmentsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTransitGatewayAttachmentsResult) Get(ctx workflow.Context) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
    var output ec2.DescribeTransitGatewayAttachmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTransitGatewayMulticastDomainsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTransitGatewayMulticastDomainsResult) Get(ctx workflow.Context) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
    var output ec2.DescribeTransitGatewayMulticastDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTransitGatewayPeeringAttachmentsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTransitGatewayPeeringAttachmentsResult) Get(ctx workflow.Context) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
    var output ec2.DescribeTransitGatewayPeeringAttachmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTransitGatewayRouteTablesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTransitGatewayRouteTablesResult) Get(ctx workflow.Context) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
    var output ec2.DescribeTransitGatewayRouteTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTransitGatewayVpcAttachmentsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTransitGatewayVpcAttachmentsResult) Get(ctx workflow.Context) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
    var output ec2.DescribeTransitGatewayVpcAttachmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeTransitGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeTransitGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeTransitGatewaysOutput, error) {
    var output ec2.DescribeTransitGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVolumeAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVolumeAttributeResult) Get(ctx workflow.Context) (*ec2.DescribeVolumeAttributeOutput, error) {
    var output ec2.DescribeVolumeAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVolumeStatusResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVolumeStatusResult) Get(ctx workflow.Context) (*ec2.DescribeVolumeStatusOutput, error) {
    var output ec2.DescribeVolumeStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVolumesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVolumesResult) Get(ctx workflow.Context) (*ec2.DescribeVolumesOutput, error) {
    var output ec2.DescribeVolumesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVolumesModificationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVolumesModificationsResult) Get(ctx workflow.Context) (*ec2.DescribeVolumesModificationsOutput, error) {
    var output ec2.DescribeVolumesModificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcAttributeResult) Get(ctx workflow.Context) (*ec2.DescribeVpcAttributeOutput, error) {
    var output ec2.DescribeVpcAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcClassicLinkResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcClassicLinkResult) Get(ctx workflow.Context) (*ec2.DescribeVpcClassicLinkOutput, error) {
    var output ec2.DescribeVpcClassicLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcClassicLinkDnsSupportResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcClassicLinkDnsSupportResult) Get(ctx workflow.Context) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
    var output ec2.DescribeVpcClassicLinkDnsSupportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcEndpointConnectionNotificationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcEndpointConnectionNotificationsResult) Get(ctx workflow.Context) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
    var output ec2.DescribeVpcEndpointConnectionNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcEndpointConnectionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcEndpointConnectionsResult) Get(ctx workflow.Context) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
    var output ec2.DescribeVpcEndpointConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcEndpointServiceConfigurationsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcEndpointServiceConfigurationsResult) Get(ctx workflow.Context) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
    var output ec2.DescribeVpcEndpointServiceConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcEndpointServicePermissionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcEndpointServicePermissionsResult) Get(ctx workflow.Context) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
    var output ec2.DescribeVpcEndpointServicePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcEndpointServicesResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcEndpointServicesResult) Get(ctx workflow.Context) (*ec2.DescribeVpcEndpointServicesOutput, error) {
    var output ec2.DescribeVpcEndpointServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcEndpointsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcEndpointsResult) Get(ctx workflow.Context) (*ec2.DescribeVpcEndpointsOutput, error) {
    var output ec2.DescribeVpcEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcPeeringConnectionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcPeeringConnectionsResult) Get(ctx workflow.Context) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
    var output ec2.DescribeVpcPeeringConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpcsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpcsResult) Get(ctx workflow.Context) (*ec2.DescribeVpcsOutput, error) {
    var output ec2.DescribeVpcsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpnConnectionsResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpnConnectionsResult) Get(ctx workflow.Context) (*ec2.DescribeVpnConnectionsOutput, error) {
    var output ec2.DescribeVpnConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DescribeVpnGatewaysResult struct {
	Result workflow.Future
}

func (r *Ec2DescribeVpnGatewaysResult) Get(ctx workflow.Context) (*ec2.DescribeVpnGatewaysOutput, error) {
    var output ec2.DescribeVpnGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DetachClassicLinkVpcResult struct {
	Result workflow.Future
}

func (r *Ec2DetachClassicLinkVpcResult) Get(ctx workflow.Context) (*ec2.DetachClassicLinkVpcOutput, error) {
    var output ec2.DetachClassicLinkVpcOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DetachInternetGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DetachInternetGatewayResult) Get(ctx workflow.Context) (*ec2.DetachInternetGatewayOutput, error) {
    var output ec2.DetachInternetGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DetachNetworkInterfaceResult struct {
	Result workflow.Future
}

func (r *Ec2DetachNetworkInterfaceResult) Get(ctx workflow.Context) (*ec2.DetachNetworkInterfaceOutput, error) {
    var output ec2.DetachNetworkInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DetachVolumeResult struct {
	Result workflow.Future
}

func (r *Ec2DetachVolumeResult) Get(ctx workflow.Context) (*ec2.VolumeAttachment, error) {
    var output ec2.VolumeAttachment
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DetachVpnGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2DetachVpnGatewayResult) Get(ctx workflow.Context) (*ec2.DetachVpnGatewayOutput, error) {
    var output ec2.DetachVpnGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisableEbsEncryptionByDefaultResult struct {
	Result workflow.Future
}

func (r *Ec2DisableEbsEncryptionByDefaultResult) Get(ctx workflow.Context) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
    var output ec2.DisableEbsEncryptionByDefaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisableFastSnapshotRestoresResult struct {
	Result workflow.Future
}

func (r *Ec2DisableFastSnapshotRestoresResult) Get(ctx workflow.Context) (*ec2.DisableFastSnapshotRestoresOutput, error) {
    var output ec2.DisableFastSnapshotRestoresOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisableTransitGatewayRouteTablePropagationResult struct {
	Result workflow.Future
}

func (r *Ec2DisableTransitGatewayRouteTablePropagationResult) Get(ctx workflow.Context) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
    var output ec2.DisableTransitGatewayRouteTablePropagationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisableVgwRoutePropagationResult struct {
	Result workflow.Future
}

func (r *Ec2DisableVgwRoutePropagationResult) Get(ctx workflow.Context) (*ec2.DisableVgwRoutePropagationOutput, error) {
    var output ec2.DisableVgwRoutePropagationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisableVpcClassicLinkResult struct {
	Result workflow.Future
}

func (r *Ec2DisableVpcClassicLinkResult) Get(ctx workflow.Context) (*ec2.DisableVpcClassicLinkOutput, error) {
    var output ec2.DisableVpcClassicLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisableVpcClassicLinkDnsSupportResult struct {
	Result workflow.Future
}

func (r *Ec2DisableVpcClassicLinkDnsSupportResult) Get(ctx workflow.Context) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
    var output ec2.DisableVpcClassicLinkDnsSupportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateAddressResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateAddressResult) Get(ctx workflow.Context) (*ec2.DisassociateAddressOutput, error) {
    var output ec2.DisassociateAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateClientVpnTargetNetworkResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateClientVpnTargetNetworkResult) Get(ctx workflow.Context) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
    var output ec2.DisassociateClientVpnTargetNetworkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateIamInstanceProfileResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateIamInstanceProfileResult) Get(ctx workflow.Context) (*ec2.DisassociateIamInstanceProfileOutput, error) {
    var output ec2.DisassociateIamInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateRouteTableResult) Get(ctx workflow.Context) (*ec2.DisassociateRouteTableOutput, error) {
    var output ec2.DisassociateRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateSubnetCidrBlockResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateSubnetCidrBlockResult) Get(ctx workflow.Context) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
    var output ec2.DisassociateSubnetCidrBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateTransitGatewayMulticastDomainResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateTransitGatewayMulticastDomainResult) Get(ctx workflow.Context) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
    var output ec2.DisassociateTransitGatewayMulticastDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateTransitGatewayRouteTableResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateTransitGatewayRouteTableResult) Get(ctx workflow.Context) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
    var output ec2.DisassociateTransitGatewayRouteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2DisassociateVpcCidrBlockResult struct {
	Result workflow.Future
}

func (r *Ec2DisassociateVpcCidrBlockResult) Get(ctx workflow.Context) (*ec2.DisassociateVpcCidrBlockOutput, error) {
    var output ec2.DisassociateVpcCidrBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2EnableEbsEncryptionByDefaultResult struct {
	Result workflow.Future
}

func (r *Ec2EnableEbsEncryptionByDefaultResult) Get(ctx workflow.Context) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
    var output ec2.EnableEbsEncryptionByDefaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2EnableFastSnapshotRestoresResult struct {
	Result workflow.Future
}

func (r *Ec2EnableFastSnapshotRestoresResult) Get(ctx workflow.Context) (*ec2.EnableFastSnapshotRestoresOutput, error) {
    var output ec2.EnableFastSnapshotRestoresOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2EnableTransitGatewayRouteTablePropagationResult struct {
	Result workflow.Future
}

func (r *Ec2EnableTransitGatewayRouteTablePropagationResult) Get(ctx workflow.Context) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
    var output ec2.EnableTransitGatewayRouteTablePropagationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2EnableVgwRoutePropagationResult struct {
	Result workflow.Future
}

func (r *Ec2EnableVgwRoutePropagationResult) Get(ctx workflow.Context) (*ec2.EnableVgwRoutePropagationOutput, error) {
    var output ec2.EnableVgwRoutePropagationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2EnableVolumeIOResult struct {
	Result workflow.Future
}

func (r *Ec2EnableVolumeIOResult) Get(ctx workflow.Context) (*ec2.EnableVolumeIOOutput, error) {
    var output ec2.EnableVolumeIOOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2EnableVpcClassicLinkResult struct {
	Result workflow.Future
}

func (r *Ec2EnableVpcClassicLinkResult) Get(ctx workflow.Context) (*ec2.EnableVpcClassicLinkOutput, error) {
    var output ec2.EnableVpcClassicLinkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2EnableVpcClassicLinkDnsSupportResult struct {
	Result workflow.Future
}

func (r *Ec2EnableVpcClassicLinkDnsSupportResult) Get(ctx workflow.Context) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
    var output ec2.EnableVpcClassicLinkDnsSupportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ExportClientVpnClientCertificateRevocationListResult struct {
	Result workflow.Future
}

func (r *Ec2ExportClientVpnClientCertificateRevocationListResult) Get(ctx workflow.Context) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
    var output ec2.ExportClientVpnClientCertificateRevocationListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ExportClientVpnClientConfigurationResult struct {
	Result workflow.Future
}

func (r *Ec2ExportClientVpnClientConfigurationResult) Get(ctx workflow.Context) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
    var output ec2.ExportClientVpnClientConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ExportImageResult struct {
	Result workflow.Future
}

func (r *Ec2ExportImageResult) Get(ctx workflow.Context) (*ec2.ExportImageOutput, error) {
    var output ec2.ExportImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ExportTransitGatewayRoutesResult struct {
	Result workflow.Future
}

func (r *Ec2ExportTransitGatewayRoutesResult) Get(ctx workflow.Context) (*ec2.ExportTransitGatewayRoutesOutput, error) {
    var output ec2.ExportTransitGatewayRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetAssociatedIpv6PoolCidrsResult struct {
	Result workflow.Future
}

func (r *Ec2GetAssociatedIpv6PoolCidrsResult) Get(ctx workflow.Context) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
    var output ec2.GetAssociatedIpv6PoolCidrsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetCapacityReservationUsageResult struct {
	Result workflow.Future
}

func (r *Ec2GetCapacityReservationUsageResult) Get(ctx workflow.Context) (*ec2.GetCapacityReservationUsageOutput, error) {
    var output ec2.GetCapacityReservationUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetCoipPoolUsageResult struct {
	Result workflow.Future
}

func (r *Ec2GetCoipPoolUsageResult) Get(ctx workflow.Context) (*ec2.GetCoipPoolUsageOutput, error) {
    var output ec2.GetCoipPoolUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetConsoleOutputResult struct {
	Result workflow.Future
}

func (r *Ec2GetConsoleOutputResult) Get(ctx workflow.Context) (*ec2.GetConsoleOutputOutput, error) {
    var output ec2.GetConsoleOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetConsoleScreenshotResult struct {
	Result workflow.Future
}

func (r *Ec2GetConsoleScreenshotResult) Get(ctx workflow.Context) (*ec2.GetConsoleScreenshotOutput, error) {
    var output ec2.GetConsoleScreenshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetDefaultCreditSpecificationResult struct {
	Result workflow.Future
}

func (r *Ec2GetDefaultCreditSpecificationResult) Get(ctx workflow.Context) (*ec2.GetDefaultCreditSpecificationOutput, error) {
    var output ec2.GetDefaultCreditSpecificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetEbsDefaultKmsKeyIdResult struct {
	Result workflow.Future
}

func (r *Ec2GetEbsDefaultKmsKeyIdResult) Get(ctx workflow.Context) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
    var output ec2.GetEbsDefaultKmsKeyIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetEbsEncryptionByDefaultResult struct {
	Result workflow.Future
}

func (r *Ec2GetEbsEncryptionByDefaultResult) Get(ctx workflow.Context) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
    var output ec2.GetEbsEncryptionByDefaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetGroupsForCapacityReservationResult struct {
	Result workflow.Future
}

func (r *Ec2GetGroupsForCapacityReservationResult) Get(ctx workflow.Context) (*ec2.GetGroupsForCapacityReservationOutput, error) {
    var output ec2.GetGroupsForCapacityReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetHostReservationPurchasePreviewResult struct {
	Result workflow.Future
}

func (r *Ec2GetHostReservationPurchasePreviewResult) Get(ctx workflow.Context) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
    var output ec2.GetHostReservationPurchasePreviewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetLaunchTemplateDataResult struct {
	Result workflow.Future
}

func (r *Ec2GetLaunchTemplateDataResult) Get(ctx workflow.Context) (*ec2.GetLaunchTemplateDataOutput, error) {
    var output ec2.GetLaunchTemplateDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetManagedPrefixListAssociationsResult struct {
	Result workflow.Future
}

func (r *Ec2GetManagedPrefixListAssociationsResult) Get(ctx workflow.Context) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
    var output ec2.GetManagedPrefixListAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetManagedPrefixListEntriesResult struct {
	Result workflow.Future
}

func (r *Ec2GetManagedPrefixListEntriesResult) Get(ctx workflow.Context) (*ec2.GetManagedPrefixListEntriesOutput, error) {
    var output ec2.GetManagedPrefixListEntriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetPasswordDataResult struct {
	Result workflow.Future
}

func (r *Ec2GetPasswordDataResult) Get(ctx workflow.Context) (*ec2.GetPasswordDataOutput, error) {
    var output ec2.GetPasswordDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetReservedInstancesExchangeQuoteResult struct {
	Result workflow.Future
}

func (r *Ec2GetReservedInstancesExchangeQuoteResult) Get(ctx workflow.Context) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
    var output ec2.GetReservedInstancesExchangeQuoteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetTransitGatewayAttachmentPropagationsResult struct {
	Result workflow.Future
}

func (r *Ec2GetTransitGatewayAttachmentPropagationsResult) Get(ctx workflow.Context) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
    var output ec2.GetTransitGatewayAttachmentPropagationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetTransitGatewayMulticastDomainAssociationsResult struct {
	Result workflow.Future
}

func (r *Ec2GetTransitGatewayMulticastDomainAssociationsResult) Get(ctx workflow.Context) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
    var output ec2.GetTransitGatewayMulticastDomainAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetTransitGatewayPrefixListReferencesResult struct {
	Result workflow.Future
}

func (r *Ec2GetTransitGatewayPrefixListReferencesResult) Get(ctx workflow.Context) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
    var output ec2.GetTransitGatewayPrefixListReferencesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetTransitGatewayRouteTableAssociationsResult struct {
	Result workflow.Future
}

func (r *Ec2GetTransitGatewayRouteTableAssociationsResult) Get(ctx workflow.Context) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
    var output ec2.GetTransitGatewayRouteTableAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2GetTransitGatewayRouteTablePropagationsResult struct {
	Result workflow.Future
}

func (r *Ec2GetTransitGatewayRouteTablePropagationsResult) Get(ctx workflow.Context) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
    var output ec2.GetTransitGatewayRouteTablePropagationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ImportClientVpnClientCertificateRevocationListResult struct {
	Result workflow.Future
}

func (r *Ec2ImportClientVpnClientCertificateRevocationListResult) Get(ctx workflow.Context) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
    var output ec2.ImportClientVpnClientCertificateRevocationListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ImportImageResult struct {
	Result workflow.Future
}

func (r *Ec2ImportImageResult) Get(ctx workflow.Context) (*ec2.ImportImageOutput, error) {
    var output ec2.ImportImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ImportInstanceResult struct {
	Result workflow.Future
}

func (r *Ec2ImportInstanceResult) Get(ctx workflow.Context) (*ec2.ImportInstanceOutput, error) {
    var output ec2.ImportInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ImportKeyPairResult struct {
	Result workflow.Future
}

func (r *Ec2ImportKeyPairResult) Get(ctx workflow.Context) (*ec2.ImportKeyPairOutput, error) {
    var output ec2.ImportKeyPairOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ImportSnapshotResult struct {
	Result workflow.Future
}

func (r *Ec2ImportSnapshotResult) Get(ctx workflow.Context) (*ec2.ImportSnapshotOutput, error) {
    var output ec2.ImportSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ImportVolumeResult struct {
	Result workflow.Future
}

func (r *Ec2ImportVolumeResult) Get(ctx workflow.Context) (*ec2.ImportVolumeOutput, error) {
    var output ec2.ImportVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyAvailabilityZoneGroupResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyAvailabilityZoneGroupResult) Get(ctx workflow.Context) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
    var output ec2.ModifyAvailabilityZoneGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyCapacityReservationResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyCapacityReservationResult) Get(ctx workflow.Context) (*ec2.ModifyCapacityReservationOutput, error) {
    var output ec2.ModifyCapacityReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyClientVpnEndpointResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyClientVpnEndpointResult) Get(ctx workflow.Context) (*ec2.ModifyClientVpnEndpointOutput, error) {
    var output ec2.ModifyClientVpnEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyDefaultCreditSpecificationResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyDefaultCreditSpecificationResult) Get(ctx workflow.Context) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
    var output ec2.ModifyDefaultCreditSpecificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyEbsDefaultKmsKeyIdResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyEbsDefaultKmsKeyIdResult) Get(ctx workflow.Context) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
    var output ec2.ModifyEbsDefaultKmsKeyIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyFleetResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyFleetResult) Get(ctx workflow.Context) (*ec2.ModifyFleetOutput, error) {
    var output ec2.ModifyFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyFpgaImageAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyFpgaImageAttributeResult) Get(ctx workflow.Context) (*ec2.ModifyFpgaImageAttributeOutput, error) {
    var output ec2.ModifyFpgaImageAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyHostsResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyHostsResult) Get(ctx workflow.Context) (*ec2.ModifyHostsOutput, error) {
    var output ec2.ModifyHostsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyIdFormatResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyIdFormatResult) Get(ctx workflow.Context) (*ec2.ModifyIdFormatOutput, error) {
    var output ec2.ModifyIdFormatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyIdentityIdFormatResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyIdentityIdFormatResult) Get(ctx workflow.Context) (*ec2.ModifyIdentityIdFormatOutput, error) {
    var output ec2.ModifyIdentityIdFormatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyImageAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyImageAttributeResult) Get(ctx workflow.Context) (*ec2.ModifyImageAttributeOutput, error) {
    var output ec2.ModifyImageAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyInstanceAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyInstanceAttributeResult) Get(ctx workflow.Context) (*ec2.ModifyInstanceAttributeOutput, error) {
    var output ec2.ModifyInstanceAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyInstanceCapacityReservationAttributesResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyInstanceCapacityReservationAttributesResult) Get(ctx workflow.Context) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
    var output ec2.ModifyInstanceCapacityReservationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyInstanceCreditSpecificationResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyInstanceCreditSpecificationResult) Get(ctx workflow.Context) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
    var output ec2.ModifyInstanceCreditSpecificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyInstanceEventStartTimeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyInstanceEventStartTimeResult) Get(ctx workflow.Context) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
    var output ec2.ModifyInstanceEventStartTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyInstanceMetadataOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyInstanceMetadataOptionsResult) Get(ctx workflow.Context) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
    var output ec2.ModifyInstanceMetadataOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyInstancePlacementResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyInstancePlacementResult) Get(ctx workflow.Context) (*ec2.ModifyInstancePlacementOutput, error) {
    var output ec2.ModifyInstancePlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyLaunchTemplateResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyLaunchTemplateResult) Get(ctx workflow.Context) (*ec2.ModifyLaunchTemplateOutput, error) {
    var output ec2.ModifyLaunchTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyManagedPrefixListResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyManagedPrefixListResult) Get(ctx workflow.Context) (*ec2.ModifyManagedPrefixListOutput, error) {
    var output ec2.ModifyManagedPrefixListOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyNetworkInterfaceAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyNetworkInterfaceAttributeResult) Get(ctx workflow.Context) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
    var output ec2.ModifyNetworkInterfaceAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyReservedInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyReservedInstancesResult) Get(ctx workflow.Context) (*ec2.ModifyReservedInstancesOutput, error) {
    var output ec2.ModifyReservedInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifySnapshotAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifySnapshotAttributeResult) Get(ctx workflow.Context) (*ec2.ModifySnapshotAttributeOutput, error) {
    var output ec2.ModifySnapshotAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifySubnetAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifySubnetAttributeResult) Get(ctx workflow.Context) (*ec2.ModifySubnetAttributeOutput, error) {
    var output ec2.ModifySubnetAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyTrafficMirrorFilterNetworkServicesResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyTrafficMirrorFilterNetworkServicesResult) Get(ctx workflow.Context) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
    var output ec2.ModifyTrafficMirrorFilterNetworkServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyTrafficMirrorFilterRuleResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyTrafficMirrorFilterRuleResult) Get(ctx workflow.Context) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
    var output ec2.ModifyTrafficMirrorFilterRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyTrafficMirrorSessionResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyTrafficMirrorSessionResult) Get(ctx workflow.Context) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
    var output ec2.ModifyTrafficMirrorSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyTransitGatewayResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyTransitGatewayResult) Get(ctx workflow.Context) (*ec2.ModifyTransitGatewayOutput, error) {
    var output ec2.ModifyTransitGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyTransitGatewayPrefixListReferenceResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyTransitGatewayPrefixListReferenceResult) Get(ctx workflow.Context) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
    var output ec2.ModifyTransitGatewayPrefixListReferenceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyTransitGatewayVpcAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyTransitGatewayVpcAttachmentResult) Get(ctx workflow.Context) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.ModifyTransitGatewayVpcAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVolumeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVolumeResult) Get(ctx workflow.Context) (*ec2.ModifyVolumeOutput, error) {
    var output ec2.ModifyVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVolumeAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVolumeAttributeResult) Get(ctx workflow.Context) (*ec2.ModifyVolumeAttributeOutput, error) {
    var output ec2.ModifyVolumeAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpcAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpcAttributeResult) Get(ctx workflow.Context) (*ec2.ModifyVpcAttributeOutput, error) {
    var output ec2.ModifyVpcAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpcEndpointResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpcEndpointResult) Get(ctx workflow.Context) (*ec2.ModifyVpcEndpointOutput, error) {
    var output ec2.ModifyVpcEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpcEndpointConnectionNotificationResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpcEndpointConnectionNotificationResult) Get(ctx workflow.Context) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
    var output ec2.ModifyVpcEndpointConnectionNotificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpcEndpointServiceConfigurationResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpcEndpointServiceConfigurationResult) Get(ctx workflow.Context) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
    var output ec2.ModifyVpcEndpointServiceConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpcEndpointServicePermissionsResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpcEndpointServicePermissionsResult) Get(ctx workflow.Context) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
    var output ec2.ModifyVpcEndpointServicePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpcPeeringConnectionOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpcPeeringConnectionOptionsResult) Get(ctx workflow.Context) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
    var output ec2.ModifyVpcPeeringConnectionOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpcTenancyResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpcTenancyResult) Get(ctx workflow.Context) (*ec2.ModifyVpcTenancyOutput, error) {
    var output ec2.ModifyVpcTenancyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpnConnectionResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpnConnectionResult) Get(ctx workflow.Context) (*ec2.ModifyVpnConnectionOutput, error) {
    var output ec2.ModifyVpnConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpnConnectionOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpnConnectionOptionsResult) Get(ctx workflow.Context) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
    var output ec2.ModifyVpnConnectionOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpnTunnelCertificateResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpnTunnelCertificateResult) Get(ctx workflow.Context) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
    var output ec2.ModifyVpnTunnelCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ModifyVpnTunnelOptionsResult struct {
	Result workflow.Future
}

func (r *Ec2ModifyVpnTunnelOptionsResult) Get(ctx workflow.Context) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
    var output ec2.ModifyVpnTunnelOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2MonitorInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2MonitorInstancesResult) Get(ctx workflow.Context) (*ec2.MonitorInstancesOutput, error) {
    var output ec2.MonitorInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2MoveAddressToVpcResult struct {
	Result workflow.Future
}

func (r *Ec2MoveAddressToVpcResult) Get(ctx workflow.Context) (*ec2.MoveAddressToVpcOutput, error) {
    var output ec2.MoveAddressToVpcOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ProvisionByoipCidrResult struct {
	Result workflow.Future
}

func (r *Ec2ProvisionByoipCidrResult) Get(ctx workflow.Context) (*ec2.ProvisionByoipCidrOutput, error) {
    var output ec2.ProvisionByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2PurchaseHostReservationResult struct {
	Result workflow.Future
}

func (r *Ec2PurchaseHostReservationResult) Get(ctx workflow.Context) (*ec2.PurchaseHostReservationOutput, error) {
    var output ec2.PurchaseHostReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2PurchaseReservedInstancesOfferingResult struct {
	Result workflow.Future
}

func (r *Ec2PurchaseReservedInstancesOfferingResult) Get(ctx workflow.Context) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
    var output ec2.PurchaseReservedInstancesOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2PurchaseScheduledInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2PurchaseScheduledInstancesResult) Get(ctx workflow.Context) (*ec2.PurchaseScheduledInstancesOutput, error) {
    var output ec2.PurchaseScheduledInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RebootInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2RebootInstancesResult) Get(ctx workflow.Context) (*ec2.RebootInstancesOutput, error) {
    var output ec2.RebootInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RegisterImageResult struct {
	Result workflow.Future
}

func (r *Ec2RegisterImageResult) Get(ctx workflow.Context) (*ec2.RegisterImageOutput, error) {
    var output ec2.RegisterImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RegisterInstanceEventNotificationAttributesResult struct {
	Result workflow.Future
}

func (r *Ec2RegisterInstanceEventNotificationAttributesResult) Get(ctx workflow.Context) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
    var output ec2.RegisterInstanceEventNotificationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RegisterTransitGatewayMulticastGroupMembersResult struct {
	Result workflow.Future
}

func (r *Ec2RegisterTransitGatewayMulticastGroupMembersResult) Get(ctx workflow.Context) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
    var output ec2.RegisterTransitGatewayMulticastGroupMembersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RegisterTransitGatewayMulticastGroupSourcesResult struct {
	Result workflow.Future
}

func (r *Ec2RegisterTransitGatewayMulticastGroupSourcesResult) Get(ctx workflow.Context) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
    var output ec2.RegisterTransitGatewayMulticastGroupSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RejectTransitGatewayPeeringAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2RejectTransitGatewayPeeringAttachmentResult) Get(ctx workflow.Context) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.RejectTransitGatewayPeeringAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RejectTransitGatewayVpcAttachmentResult struct {
	Result workflow.Future
}

func (r *Ec2RejectTransitGatewayVpcAttachmentResult) Get(ctx workflow.Context) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.RejectTransitGatewayVpcAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RejectVpcEndpointConnectionsResult struct {
	Result workflow.Future
}

func (r *Ec2RejectVpcEndpointConnectionsResult) Get(ctx workflow.Context) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
    var output ec2.RejectVpcEndpointConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RejectVpcPeeringConnectionResult struct {
	Result workflow.Future
}

func (r *Ec2RejectVpcPeeringConnectionResult) Get(ctx workflow.Context) (*ec2.RejectVpcPeeringConnectionOutput, error) {
    var output ec2.RejectVpcPeeringConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReleaseAddressResult struct {
	Result workflow.Future
}

func (r *Ec2ReleaseAddressResult) Get(ctx workflow.Context) (*ec2.ReleaseAddressOutput, error) {
    var output ec2.ReleaseAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReleaseHostsResult struct {
	Result workflow.Future
}

func (r *Ec2ReleaseHostsResult) Get(ctx workflow.Context) (*ec2.ReleaseHostsOutput, error) {
    var output ec2.ReleaseHostsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReplaceIamInstanceProfileAssociationResult struct {
	Result workflow.Future
}

func (r *Ec2ReplaceIamInstanceProfileAssociationResult) Get(ctx workflow.Context) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
    var output ec2.ReplaceIamInstanceProfileAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReplaceNetworkAclAssociationResult struct {
	Result workflow.Future
}

func (r *Ec2ReplaceNetworkAclAssociationResult) Get(ctx workflow.Context) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
    var output ec2.ReplaceNetworkAclAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReplaceNetworkAclEntryResult struct {
	Result workflow.Future
}

func (r *Ec2ReplaceNetworkAclEntryResult) Get(ctx workflow.Context) (*ec2.ReplaceNetworkAclEntryOutput, error) {
    var output ec2.ReplaceNetworkAclEntryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReplaceRouteResult struct {
	Result workflow.Future
}

func (r *Ec2ReplaceRouteResult) Get(ctx workflow.Context) (*ec2.ReplaceRouteOutput, error) {
    var output ec2.ReplaceRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReplaceRouteTableAssociationResult struct {
	Result workflow.Future
}

func (r *Ec2ReplaceRouteTableAssociationResult) Get(ctx workflow.Context) (*ec2.ReplaceRouteTableAssociationOutput, error) {
    var output ec2.ReplaceRouteTableAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReplaceTransitGatewayRouteResult struct {
	Result workflow.Future
}

func (r *Ec2ReplaceTransitGatewayRouteResult) Get(ctx workflow.Context) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
    var output ec2.ReplaceTransitGatewayRouteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ReportInstanceStatusResult struct {
	Result workflow.Future
}

func (r *Ec2ReportInstanceStatusResult) Get(ctx workflow.Context) (*ec2.ReportInstanceStatusOutput, error) {
    var output ec2.ReportInstanceStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RequestSpotFleetResult struct {
	Result workflow.Future
}

func (r *Ec2RequestSpotFleetResult) Get(ctx workflow.Context) (*ec2.RequestSpotFleetOutput, error) {
    var output ec2.RequestSpotFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RequestSpotInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2RequestSpotInstancesResult) Get(ctx workflow.Context) (*ec2.RequestSpotInstancesOutput, error) {
    var output ec2.RequestSpotInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ResetEbsDefaultKmsKeyIdResult struct {
	Result workflow.Future
}

func (r *Ec2ResetEbsDefaultKmsKeyIdResult) Get(ctx workflow.Context) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
    var output ec2.ResetEbsDefaultKmsKeyIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ResetFpgaImageAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ResetFpgaImageAttributeResult) Get(ctx workflow.Context) (*ec2.ResetFpgaImageAttributeOutput, error) {
    var output ec2.ResetFpgaImageAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ResetImageAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ResetImageAttributeResult) Get(ctx workflow.Context) (*ec2.ResetImageAttributeOutput, error) {
    var output ec2.ResetImageAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ResetInstanceAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ResetInstanceAttributeResult) Get(ctx workflow.Context) (*ec2.ResetInstanceAttributeOutput, error) {
    var output ec2.ResetInstanceAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ResetNetworkInterfaceAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ResetNetworkInterfaceAttributeResult) Get(ctx workflow.Context) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
    var output ec2.ResetNetworkInterfaceAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2ResetSnapshotAttributeResult struct {
	Result workflow.Future
}

func (r *Ec2ResetSnapshotAttributeResult) Get(ctx workflow.Context) (*ec2.ResetSnapshotAttributeOutput, error) {
    var output ec2.ResetSnapshotAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RestoreAddressToClassicResult struct {
	Result workflow.Future
}

func (r *Ec2RestoreAddressToClassicResult) Get(ctx workflow.Context) (*ec2.RestoreAddressToClassicOutput, error) {
    var output ec2.RestoreAddressToClassicOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RestoreManagedPrefixListVersionResult struct {
	Result workflow.Future
}

func (r *Ec2RestoreManagedPrefixListVersionResult) Get(ctx workflow.Context) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
    var output ec2.RestoreManagedPrefixListVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RevokeClientVpnIngressResult struct {
	Result workflow.Future
}

func (r *Ec2RevokeClientVpnIngressResult) Get(ctx workflow.Context) (*ec2.RevokeClientVpnIngressOutput, error) {
    var output ec2.RevokeClientVpnIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RevokeSecurityGroupEgressResult struct {
	Result workflow.Future
}

func (r *Ec2RevokeSecurityGroupEgressResult) Get(ctx workflow.Context) (*ec2.RevokeSecurityGroupEgressOutput, error) {
    var output ec2.RevokeSecurityGroupEgressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RevokeSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *Ec2RevokeSecurityGroupIngressResult) Get(ctx workflow.Context) (*ec2.RevokeSecurityGroupIngressOutput, error) {
    var output ec2.RevokeSecurityGroupIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RunInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2RunInstancesResult) Get(ctx workflow.Context) (*ec2.Reservation, error) {
    var output ec2.Reservation
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2RunScheduledInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2RunScheduledInstancesResult) Get(ctx workflow.Context) (*ec2.RunScheduledInstancesOutput, error) {
    var output ec2.RunScheduledInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2SearchLocalGatewayRoutesResult struct {
	Result workflow.Future
}

func (r *Ec2SearchLocalGatewayRoutesResult) Get(ctx workflow.Context) (*ec2.SearchLocalGatewayRoutesOutput, error) {
    var output ec2.SearchLocalGatewayRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2SearchTransitGatewayMulticastGroupsResult struct {
	Result workflow.Future
}

func (r *Ec2SearchTransitGatewayMulticastGroupsResult) Get(ctx workflow.Context) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
    var output ec2.SearchTransitGatewayMulticastGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2SearchTransitGatewayRoutesResult struct {
	Result workflow.Future
}

func (r *Ec2SearchTransitGatewayRoutesResult) Get(ctx workflow.Context) (*ec2.SearchTransitGatewayRoutesOutput, error) {
    var output ec2.SearchTransitGatewayRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2SendDiagnosticInterruptResult struct {
	Result workflow.Future
}

func (r *Ec2SendDiagnosticInterruptResult) Get(ctx workflow.Context) (*ec2.SendDiagnosticInterruptOutput, error) {
    var output ec2.SendDiagnosticInterruptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2StartInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2StartInstancesResult) Get(ctx workflow.Context) (*ec2.StartInstancesOutput, error) {
    var output ec2.StartInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2StartVpcEndpointServicePrivateDnsVerificationResult struct {
	Result workflow.Future
}

func (r *Ec2StartVpcEndpointServicePrivateDnsVerificationResult) Get(ctx workflow.Context) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
    var output ec2.StartVpcEndpointServicePrivateDnsVerificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2StopInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2StopInstancesResult) Get(ctx workflow.Context) (*ec2.StopInstancesOutput, error) {
    var output ec2.StopInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2TerminateClientVpnConnectionsResult struct {
	Result workflow.Future
}

func (r *Ec2TerminateClientVpnConnectionsResult) Get(ctx workflow.Context) (*ec2.TerminateClientVpnConnectionsOutput, error) {
    var output ec2.TerminateClientVpnConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2TerminateInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2TerminateInstancesResult) Get(ctx workflow.Context) (*ec2.TerminateInstancesOutput, error) {
    var output ec2.TerminateInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2UnassignIpv6AddressesResult struct {
	Result workflow.Future
}

func (r *Ec2UnassignIpv6AddressesResult) Get(ctx workflow.Context) (*ec2.UnassignIpv6AddressesOutput, error) {
    var output ec2.UnassignIpv6AddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2UnassignPrivateIpAddressesResult struct {
	Result workflow.Future
}

func (r *Ec2UnassignPrivateIpAddressesResult) Get(ctx workflow.Context) (*ec2.UnassignPrivateIpAddressesOutput, error) {
    var output ec2.UnassignPrivateIpAddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2UnmonitorInstancesResult struct {
	Result workflow.Future
}

func (r *Ec2UnmonitorInstancesResult) Get(ctx workflow.Context) (*ec2.UnmonitorInstancesOutput, error) {
    var output ec2.UnmonitorInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2UpdateSecurityGroupRuleDescriptionsEgressResult struct {
	Result workflow.Future
}

func (r *Ec2UpdateSecurityGroupRuleDescriptionsEgressResult) Get(ctx workflow.Context) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
    var output ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2UpdateSecurityGroupRuleDescriptionsIngressResult struct {
	Result workflow.Future
}

func (r *Ec2UpdateSecurityGroupRuleDescriptionsIngressResult) Get(ctx workflow.Context) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
    var output ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Ec2WithdrawByoipCidrResult struct {
	Result workflow.Future
}

func (r *Ec2WithdrawByoipCidrResult) Get(ctx workflow.Context) (*ec2.WithdrawByoipCidrOutput, error) {
    var output ec2.WithdrawByoipCidrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EC2Stub struct {
    activities awsactivities.EC2Activities
}

func NewEC2Stub() EC2Client {
    return &EC2Stub{}
}

func (a *EC2Stub) AcceptReservedInstancesExchangeQuote(ctx workflow.Context, input *ec2.AcceptReservedInstancesExchangeQuoteInput) (*ec2.AcceptReservedInstancesExchangeQuoteOutput, error) {
    var output ec2.AcceptReservedInstancesExchangeQuoteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptReservedInstancesExchangeQuote, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AcceptReservedInstancesExchangeQuoteAsync(ctx workflow.Context, input *ec2.AcceptReservedInstancesExchangeQuoteInput) *Ec2AcceptReservedInstancesExchangeQuoteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptReservedInstancesExchangeQuote, input)
    return &Ec2AcceptReservedInstancesExchangeQuoteResult{Result: future}
}

func (a *EC2Stub) AcceptTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.AcceptTransitGatewayPeeringAttachmentInput) (*ec2.AcceptTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.AcceptTransitGatewayPeeringAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptTransitGatewayPeeringAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AcceptTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.AcceptTransitGatewayPeeringAttachmentInput) *Ec2AcceptTransitGatewayPeeringAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptTransitGatewayPeeringAttachment, input)
    return &Ec2AcceptTransitGatewayPeeringAttachmentResult{Result: future}
}

func (a *EC2Stub) AcceptTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.AcceptTransitGatewayVpcAttachmentInput) (*ec2.AcceptTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.AcceptTransitGatewayVpcAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptTransitGatewayVpcAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AcceptTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.AcceptTransitGatewayVpcAttachmentInput) *Ec2AcceptTransitGatewayVpcAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptTransitGatewayVpcAttachment, input)
    return &Ec2AcceptTransitGatewayVpcAttachmentResult{Result: future}
}

func (a *EC2Stub) AcceptVpcEndpointConnections(ctx workflow.Context, input *ec2.AcceptVpcEndpointConnectionsInput) (*ec2.AcceptVpcEndpointConnectionsOutput, error) {
    var output ec2.AcceptVpcEndpointConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptVpcEndpointConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AcceptVpcEndpointConnectionsAsync(ctx workflow.Context, input *ec2.AcceptVpcEndpointConnectionsInput) *Ec2AcceptVpcEndpointConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptVpcEndpointConnections, input)
    return &Ec2AcceptVpcEndpointConnectionsResult{Result: future}
}

func (a *EC2Stub) AcceptVpcPeeringConnection(ctx workflow.Context, input *ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error) {
    var output ec2.AcceptVpcPeeringConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptVpcPeeringConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AcceptVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.AcceptVpcPeeringConnectionInput) *Ec2AcceptVpcPeeringConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptVpcPeeringConnection, input)
    return &Ec2AcceptVpcPeeringConnectionResult{Result: future}
}

func (a *EC2Stub) AdvertiseByoipCidr(ctx workflow.Context, input *ec2.AdvertiseByoipCidrInput) (*ec2.AdvertiseByoipCidrOutput, error) {
    var output ec2.AdvertiseByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdvertiseByoipCidr, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AdvertiseByoipCidrAsync(ctx workflow.Context, input *ec2.AdvertiseByoipCidrInput) *Ec2AdvertiseByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AdvertiseByoipCidr, input)
    return &Ec2AdvertiseByoipCidrResult{Result: future}
}

func (a *EC2Stub) AllocateAddress(ctx workflow.Context, input *ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error) {
    var output ec2.AllocateAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AllocateAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AllocateAddressAsync(ctx workflow.Context, input *ec2.AllocateAddressInput) *Ec2AllocateAddressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AllocateAddress, input)
    return &Ec2AllocateAddressResult{Result: future}
}

func (a *EC2Stub) AllocateHosts(ctx workflow.Context, input *ec2.AllocateHostsInput) (*ec2.AllocateHostsOutput, error) {
    var output ec2.AllocateHostsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AllocateHosts, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AllocateHostsAsync(ctx workflow.Context, input *ec2.AllocateHostsInput) *Ec2AllocateHostsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AllocateHosts, input)
    return &Ec2AllocateHostsResult{Result: future}
}

func (a *EC2Stub) ApplySecurityGroupsToClientVpnTargetNetwork(ctx workflow.Context, input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) (*ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
    var output ec2.ApplySecurityGroupsToClientVpnTargetNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ApplySecurityGroupsToClientVpnTargetNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ApplySecurityGroupsToClientVpnTargetNetworkAsync(ctx workflow.Context, input *ec2.ApplySecurityGroupsToClientVpnTargetNetworkInput) *Ec2ApplySecurityGroupsToClientVpnTargetNetworkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ApplySecurityGroupsToClientVpnTargetNetwork, input)
    return &Ec2ApplySecurityGroupsToClientVpnTargetNetworkResult{Result: future}
}

func (a *EC2Stub) AssignIpv6Addresses(ctx workflow.Context, input *ec2.AssignIpv6AddressesInput) (*ec2.AssignIpv6AddressesOutput, error) {
    var output ec2.AssignIpv6AddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssignIpv6Addresses, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssignIpv6AddressesAsync(ctx workflow.Context, input *ec2.AssignIpv6AddressesInput) *Ec2AssignIpv6AddressesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssignIpv6Addresses, input)
    return &Ec2AssignIpv6AddressesResult{Result: future}
}

func (a *EC2Stub) AssignPrivateIpAddresses(ctx workflow.Context, input *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
    var output ec2.AssignPrivateIpAddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssignPrivateIpAddresses, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssignPrivateIpAddressesAsync(ctx workflow.Context, input *ec2.AssignPrivateIpAddressesInput) *Ec2AssignPrivateIpAddressesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssignPrivateIpAddresses, input)
    return &Ec2AssignPrivateIpAddressesResult{Result: future}
}

func (a *EC2Stub) AssociateAddress(ctx workflow.Context, input *ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error) {
    var output ec2.AssociateAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateAddressAsync(ctx workflow.Context, input *ec2.AssociateAddressInput) *Ec2AssociateAddressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateAddress, input)
    return &Ec2AssociateAddressResult{Result: future}
}

func (a *EC2Stub) AssociateClientVpnTargetNetwork(ctx workflow.Context, input *ec2.AssociateClientVpnTargetNetworkInput) (*ec2.AssociateClientVpnTargetNetworkOutput, error) {
    var output ec2.AssociateClientVpnTargetNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateClientVpnTargetNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateClientVpnTargetNetworkAsync(ctx workflow.Context, input *ec2.AssociateClientVpnTargetNetworkInput) *Ec2AssociateClientVpnTargetNetworkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateClientVpnTargetNetwork, input)
    return &Ec2AssociateClientVpnTargetNetworkResult{Result: future}
}

func (a *EC2Stub) AssociateDhcpOptions(ctx workflow.Context, input *ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error) {
    var output ec2.AssociateDhcpOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateDhcpOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateDhcpOptionsAsync(ctx workflow.Context, input *ec2.AssociateDhcpOptionsInput) *Ec2AssociateDhcpOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateDhcpOptions, input)
    return &Ec2AssociateDhcpOptionsResult{Result: future}
}

func (a *EC2Stub) AssociateIamInstanceProfile(ctx workflow.Context, input *ec2.AssociateIamInstanceProfileInput) (*ec2.AssociateIamInstanceProfileOutput, error) {
    var output ec2.AssociateIamInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateIamInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateIamInstanceProfileAsync(ctx workflow.Context, input *ec2.AssociateIamInstanceProfileInput) *Ec2AssociateIamInstanceProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateIamInstanceProfile, input)
    return &Ec2AssociateIamInstanceProfileResult{Result: future}
}

func (a *EC2Stub) AssociateRouteTable(ctx workflow.Context, input *ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error) {
    var output ec2.AssociateRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateRouteTableAsync(ctx workflow.Context, input *ec2.AssociateRouteTableInput) *Ec2AssociateRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateRouteTable, input)
    return &Ec2AssociateRouteTableResult{Result: future}
}

func (a *EC2Stub) AssociateSubnetCidrBlock(ctx workflow.Context, input *ec2.AssociateSubnetCidrBlockInput) (*ec2.AssociateSubnetCidrBlockOutput, error) {
    var output ec2.AssociateSubnetCidrBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateSubnetCidrBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateSubnetCidrBlockAsync(ctx workflow.Context, input *ec2.AssociateSubnetCidrBlockInput) *Ec2AssociateSubnetCidrBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateSubnetCidrBlock, input)
    return &Ec2AssociateSubnetCidrBlockResult{Result: future}
}

func (a *EC2Stub) AssociateTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.AssociateTransitGatewayMulticastDomainInput) (*ec2.AssociateTransitGatewayMulticastDomainOutput, error) {
    var output ec2.AssociateTransitGatewayMulticastDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateTransitGatewayMulticastDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.AssociateTransitGatewayMulticastDomainInput) *Ec2AssociateTransitGatewayMulticastDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateTransitGatewayMulticastDomain, input)
    return &Ec2AssociateTransitGatewayMulticastDomainResult{Result: future}
}

func (a *EC2Stub) AssociateTransitGatewayRouteTable(ctx workflow.Context, input *ec2.AssociateTransitGatewayRouteTableInput) (*ec2.AssociateTransitGatewayRouteTableOutput, error) {
    var output ec2.AssociateTransitGatewayRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateTransitGatewayRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.AssociateTransitGatewayRouteTableInput) *Ec2AssociateTransitGatewayRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateTransitGatewayRouteTable, input)
    return &Ec2AssociateTransitGatewayRouteTableResult{Result: future}
}

func (a *EC2Stub) AssociateVpcCidrBlock(ctx workflow.Context, input *ec2.AssociateVpcCidrBlockInput) (*ec2.AssociateVpcCidrBlockOutput, error) {
    var output ec2.AssociateVpcCidrBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateVpcCidrBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AssociateVpcCidrBlockAsync(ctx workflow.Context, input *ec2.AssociateVpcCidrBlockInput) *Ec2AssociateVpcCidrBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateVpcCidrBlock, input)
    return &Ec2AssociateVpcCidrBlockResult{Result: future}
}

func (a *EC2Stub) AttachClassicLinkVpc(ctx workflow.Context, input *ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error) {
    var output ec2.AttachClassicLinkVpcOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachClassicLinkVpc, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AttachClassicLinkVpcAsync(ctx workflow.Context, input *ec2.AttachClassicLinkVpcInput) *Ec2AttachClassicLinkVpcResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachClassicLinkVpc, input)
    return &Ec2AttachClassicLinkVpcResult{Result: future}
}

func (a *EC2Stub) AttachInternetGateway(ctx workflow.Context, input *ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error) {
    var output ec2.AttachInternetGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachInternetGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AttachInternetGatewayAsync(ctx workflow.Context, input *ec2.AttachInternetGatewayInput) *Ec2AttachInternetGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachInternetGateway, input)
    return &Ec2AttachInternetGatewayResult{Result: future}
}

func (a *EC2Stub) AttachNetworkInterface(ctx workflow.Context, input *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
    var output ec2.AttachNetworkInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachNetworkInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AttachNetworkInterfaceAsync(ctx workflow.Context, input *ec2.AttachNetworkInterfaceInput) *Ec2AttachNetworkInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachNetworkInterface, input)
    return &Ec2AttachNetworkInterfaceResult{Result: future}
}

func (a *EC2Stub) AttachVolume(ctx workflow.Context, input *ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error) {
    var output ec2.VolumeAttachment
    err := workflow.ExecuteActivity(ctx, a.activities.AttachVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AttachVolumeAsync(ctx workflow.Context, input *ec2.AttachVolumeInput) *Ec2AttachVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachVolume, input)
    return &Ec2AttachVolumeResult{Result: future}
}

func (a *EC2Stub) AttachVpnGateway(ctx workflow.Context, input *ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error) {
    var output ec2.AttachVpnGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachVpnGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AttachVpnGatewayAsync(ctx workflow.Context, input *ec2.AttachVpnGatewayInput) *Ec2AttachVpnGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachVpnGateway, input)
    return &Ec2AttachVpnGatewayResult{Result: future}
}

func (a *EC2Stub) AuthorizeClientVpnIngress(ctx workflow.Context, input *ec2.AuthorizeClientVpnIngressInput) (*ec2.AuthorizeClientVpnIngressOutput, error) {
    var output ec2.AuthorizeClientVpnIngressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AuthorizeClientVpnIngress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AuthorizeClientVpnIngressAsync(ctx workflow.Context, input *ec2.AuthorizeClientVpnIngressInput) *Ec2AuthorizeClientVpnIngressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AuthorizeClientVpnIngress, input)
    return &Ec2AuthorizeClientVpnIngressResult{Result: future}
}

func (a *EC2Stub) AuthorizeSecurityGroupEgress(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
    var output ec2.AuthorizeSecurityGroupEgressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AuthorizeSecurityGroupEgress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AuthorizeSecurityGroupEgressAsync(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupEgressInput) *Ec2AuthorizeSecurityGroupEgressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AuthorizeSecurityGroupEgress, input)
    return &Ec2AuthorizeSecurityGroupEgressResult{Result: future}
}

func (a *EC2Stub) AuthorizeSecurityGroupIngress(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
    var output ec2.AuthorizeSecurityGroupIngressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AuthorizeSecurityGroupIngress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) AuthorizeSecurityGroupIngressAsync(ctx workflow.Context, input *ec2.AuthorizeSecurityGroupIngressInput) *Ec2AuthorizeSecurityGroupIngressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AuthorizeSecurityGroupIngress, input)
    return &Ec2AuthorizeSecurityGroupIngressResult{Result: future}
}

func (a *EC2Stub) BundleInstance(ctx workflow.Context, input *ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error) {
    var output ec2.BundleInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BundleInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) BundleInstanceAsync(ctx workflow.Context, input *ec2.BundleInstanceInput) *Ec2BundleInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BundleInstance, input)
    return &Ec2BundleInstanceResult{Result: future}
}

func (a *EC2Stub) CancelBundleTask(ctx workflow.Context, input *ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error) {
    var output ec2.CancelBundleTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelBundleTask, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelBundleTaskAsync(ctx workflow.Context, input *ec2.CancelBundleTaskInput) *Ec2CancelBundleTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelBundleTask, input)
    return &Ec2CancelBundleTaskResult{Result: future}
}

func (a *EC2Stub) CancelCapacityReservation(ctx workflow.Context, input *ec2.CancelCapacityReservationInput) (*ec2.CancelCapacityReservationOutput, error) {
    var output ec2.CancelCapacityReservationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelCapacityReservation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelCapacityReservationAsync(ctx workflow.Context, input *ec2.CancelCapacityReservationInput) *Ec2CancelCapacityReservationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelCapacityReservation, input)
    return &Ec2CancelCapacityReservationResult{Result: future}
}

func (a *EC2Stub) CancelConversionTask(ctx workflow.Context, input *ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error) {
    var output ec2.CancelConversionTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelConversionTask, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelConversionTaskAsync(ctx workflow.Context, input *ec2.CancelConversionTaskInput) *Ec2CancelConversionTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelConversionTask, input)
    return &Ec2CancelConversionTaskResult{Result: future}
}

func (a *EC2Stub) CancelExportTask(ctx workflow.Context, input *ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error) {
    var output ec2.CancelExportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelExportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelExportTaskAsync(ctx workflow.Context, input *ec2.CancelExportTaskInput) *Ec2CancelExportTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelExportTask, input)
    return &Ec2CancelExportTaskResult{Result: future}
}

func (a *EC2Stub) CancelImportTask(ctx workflow.Context, input *ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error) {
    var output ec2.CancelImportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelImportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelImportTaskAsync(ctx workflow.Context, input *ec2.CancelImportTaskInput) *Ec2CancelImportTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelImportTask, input)
    return &Ec2CancelImportTaskResult{Result: future}
}

func (a *EC2Stub) CancelReservedInstancesListing(ctx workflow.Context, input *ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error) {
    var output ec2.CancelReservedInstancesListingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelReservedInstancesListing, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelReservedInstancesListingAsync(ctx workflow.Context, input *ec2.CancelReservedInstancesListingInput) *Ec2CancelReservedInstancesListingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelReservedInstancesListing, input)
    return &Ec2CancelReservedInstancesListingResult{Result: future}
}

func (a *EC2Stub) CancelSpotFleetRequests(ctx workflow.Context, input *ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error) {
    var output ec2.CancelSpotFleetRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelSpotFleetRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelSpotFleetRequestsAsync(ctx workflow.Context, input *ec2.CancelSpotFleetRequestsInput) *Ec2CancelSpotFleetRequestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelSpotFleetRequests, input)
    return &Ec2CancelSpotFleetRequestsResult{Result: future}
}

func (a *EC2Stub) CancelSpotInstanceRequests(ctx workflow.Context, input *ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error) {
    var output ec2.CancelSpotInstanceRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelSpotInstanceRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CancelSpotInstanceRequestsAsync(ctx workflow.Context, input *ec2.CancelSpotInstanceRequestsInput) *Ec2CancelSpotInstanceRequestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelSpotInstanceRequests, input)
    return &Ec2CancelSpotInstanceRequestsResult{Result: future}
}

func (a *EC2Stub) ConfirmProductInstance(ctx workflow.Context, input *ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error) {
    var output ec2.ConfirmProductInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmProductInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ConfirmProductInstanceAsync(ctx workflow.Context, input *ec2.ConfirmProductInstanceInput) *Ec2ConfirmProductInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ConfirmProductInstance, input)
    return &Ec2ConfirmProductInstanceResult{Result: future}
}

func (a *EC2Stub) CopyFpgaImage(ctx workflow.Context, input *ec2.CopyFpgaImageInput) (*ec2.CopyFpgaImageOutput, error) {
    var output ec2.CopyFpgaImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyFpgaImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CopyFpgaImageAsync(ctx workflow.Context, input *ec2.CopyFpgaImageInput) *Ec2CopyFpgaImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyFpgaImage, input)
    return &Ec2CopyFpgaImageResult{Result: future}
}

func (a *EC2Stub) CopyImage(ctx workflow.Context, input *ec2.CopyImageInput) (*ec2.CopyImageOutput, error) {
    var output ec2.CopyImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CopyImageAsync(ctx workflow.Context, input *ec2.CopyImageInput) *Ec2CopyImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyImage, input)
    return &Ec2CopyImageResult{Result: future}
}

func (a *EC2Stub) CopySnapshot(ctx workflow.Context, input *ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error) {
    var output ec2.CopySnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopySnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CopySnapshotAsync(ctx workflow.Context, input *ec2.CopySnapshotInput) *Ec2CopySnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopySnapshot, input)
    return &Ec2CopySnapshotResult{Result: future}
}

func (a *EC2Stub) CreateCapacityReservation(ctx workflow.Context, input *ec2.CreateCapacityReservationInput) (*ec2.CreateCapacityReservationOutput, error) {
    var output ec2.CreateCapacityReservationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCapacityReservation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateCapacityReservationAsync(ctx workflow.Context, input *ec2.CreateCapacityReservationInput) *Ec2CreateCapacityReservationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCapacityReservation, input)
    return &Ec2CreateCapacityReservationResult{Result: future}
}

func (a *EC2Stub) CreateCarrierGateway(ctx workflow.Context, input *ec2.CreateCarrierGatewayInput) (*ec2.CreateCarrierGatewayOutput, error) {
    var output ec2.CreateCarrierGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCarrierGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateCarrierGatewayAsync(ctx workflow.Context, input *ec2.CreateCarrierGatewayInput) *Ec2CreateCarrierGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCarrierGateway, input)
    return &Ec2CreateCarrierGatewayResult{Result: future}
}

func (a *EC2Stub) CreateClientVpnEndpoint(ctx workflow.Context, input *ec2.CreateClientVpnEndpointInput) (*ec2.CreateClientVpnEndpointOutput, error) {
    var output ec2.CreateClientVpnEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateClientVpnEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateClientVpnEndpointAsync(ctx workflow.Context, input *ec2.CreateClientVpnEndpointInput) *Ec2CreateClientVpnEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateClientVpnEndpoint, input)
    return &Ec2CreateClientVpnEndpointResult{Result: future}
}

func (a *EC2Stub) CreateClientVpnRoute(ctx workflow.Context, input *ec2.CreateClientVpnRouteInput) (*ec2.CreateClientVpnRouteOutput, error) {
    var output ec2.CreateClientVpnRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateClientVpnRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateClientVpnRouteAsync(ctx workflow.Context, input *ec2.CreateClientVpnRouteInput) *Ec2CreateClientVpnRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateClientVpnRoute, input)
    return &Ec2CreateClientVpnRouteResult{Result: future}
}

func (a *EC2Stub) CreateCustomerGateway(ctx workflow.Context, input *ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error) {
    var output ec2.CreateCustomerGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCustomerGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateCustomerGatewayAsync(ctx workflow.Context, input *ec2.CreateCustomerGatewayInput) *Ec2CreateCustomerGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCustomerGateway, input)
    return &Ec2CreateCustomerGatewayResult{Result: future}
}

func (a *EC2Stub) CreateDefaultSubnet(ctx workflow.Context, input *ec2.CreateDefaultSubnetInput) (*ec2.CreateDefaultSubnetOutput, error) {
    var output ec2.CreateDefaultSubnetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDefaultSubnet, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateDefaultSubnetAsync(ctx workflow.Context, input *ec2.CreateDefaultSubnetInput) *Ec2CreateDefaultSubnetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDefaultSubnet, input)
    return &Ec2CreateDefaultSubnetResult{Result: future}
}

func (a *EC2Stub) CreateDefaultVpc(ctx workflow.Context, input *ec2.CreateDefaultVpcInput) (*ec2.CreateDefaultVpcOutput, error) {
    var output ec2.CreateDefaultVpcOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDefaultVpc, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateDefaultVpcAsync(ctx workflow.Context, input *ec2.CreateDefaultVpcInput) *Ec2CreateDefaultVpcResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDefaultVpc, input)
    return &Ec2CreateDefaultVpcResult{Result: future}
}

func (a *EC2Stub) CreateDhcpOptions(ctx workflow.Context, input *ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error) {
    var output ec2.CreateDhcpOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDhcpOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateDhcpOptionsAsync(ctx workflow.Context, input *ec2.CreateDhcpOptionsInput) *Ec2CreateDhcpOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDhcpOptions, input)
    return &Ec2CreateDhcpOptionsResult{Result: future}
}

func (a *EC2Stub) CreateEgressOnlyInternetGateway(ctx workflow.Context, input *ec2.CreateEgressOnlyInternetGatewayInput) (*ec2.CreateEgressOnlyInternetGatewayOutput, error) {
    var output ec2.CreateEgressOnlyInternetGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEgressOnlyInternetGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateEgressOnlyInternetGatewayAsync(ctx workflow.Context, input *ec2.CreateEgressOnlyInternetGatewayInput) *Ec2CreateEgressOnlyInternetGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEgressOnlyInternetGateway, input)
    return &Ec2CreateEgressOnlyInternetGatewayResult{Result: future}
}

func (a *EC2Stub) CreateFleet(ctx workflow.Context, input *ec2.CreateFleetInput) (*ec2.CreateFleetOutput, error) {
    var output ec2.CreateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateFleetAsync(ctx workflow.Context, input *ec2.CreateFleetInput) *Ec2CreateFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input)
    return &Ec2CreateFleetResult{Result: future}
}

func (a *EC2Stub) CreateFlowLogs(ctx workflow.Context, input *ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error) {
    var output ec2.CreateFlowLogsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFlowLogs, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateFlowLogsAsync(ctx workflow.Context, input *ec2.CreateFlowLogsInput) *Ec2CreateFlowLogsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFlowLogs, input)
    return &Ec2CreateFlowLogsResult{Result: future}
}

func (a *EC2Stub) CreateFpgaImage(ctx workflow.Context, input *ec2.CreateFpgaImageInput) (*ec2.CreateFpgaImageOutput, error) {
    var output ec2.CreateFpgaImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFpgaImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateFpgaImageAsync(ctx workflow.Context, input *ec2.CreateFpgaImageInput) *Ec2CreateFpgaImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFpgaImage, input)
    return &Ec2CreateFpgaImageResult{Result: future}
}

func (a *EC2Stub) CreateImage(ctx workflow.Context, input *ec2.CreateImageInput) (*ec2.CreateImageOutput, error) {
    var output ec2.CreateImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateImageAsync(ctx workflow.Context, input *ec2.CreateImageInput) *Ec2CreateImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateImage, input)
    return &Ec2CreateImageResult{Result: future}
}

func (a *EC2Stub) CreateInstanceExportTask(ctx workflow.Context, input *ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error) {
    var output ec2.CreateInstanceExportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInstanceExportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateInstanceExportTaskAsync(ctx workflow.Context, input *ec2.CreateInstanceExportTaskInput) *Ec2CreateInstanceExportTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateInstanceExportTask, input)
    return &Ec2CreateInstanceExportTaskResult{Result: future}
}

func (a *EC2Stub) CreateInternetGateway(ctx workflow.Context, input *ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error) {
    var output ec2.CreateInternetGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInternetGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateInternetGatewayAsync(ctx workflow.Context, input *ec2.CreateInternetGatewayInput) *Ec2CreateInternetGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateInternetGateway, input)
    return &Ec2CreateInternetGatewayResult{Result: future}
}

func (a *EC2Stub) CreateKeyPair(ctx workflow.Context, input *ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error) {
    var output ec2.CreateKeyPairOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateKeyPair, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateKeyPairAsync(ctx workflow.Context, input *ec2.CreateKeyPairInput) *Ec2CreateKeyPairResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateKeyPair, input)
    return &Ec2CreateKeyPairResult{Result: future}
}

func (a *EC2Stub) CreateLaunchTemplate(ctx workflow.Context, input *ec2.CreateLaunchTemplateInput) (*ec2.CreateLaunchTemplateOutput, error) {
    var output ec2.CreateLaunchTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLaunchTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateLaunchTemplateAsync(ctx workflow.Context, input *ec2.CreateLaunchTemplateInput) *Ec2CreateLaunchTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLaunchTemplate, input)
    return &Ec2CreateLaunchTemplateResult{Result: future}
}

func (a *EC2Stub) CreateLaunchTemplateVersion(ctx workflow.Context, input *ec2.CreateLaunchTemplateVersionInput) (*ec2.CreateLaunchTemplateVersionOutput, error) {
    var output ec2.CreateLaunchTemplateVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLaunchTemplateVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateLaunchTemplateVersionAsync(ctx workflow.Context, input *ec2.CreateLaunchTemplateVersionInput) *Ec2CreateLaunchTemplateVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLaunchTemplateVersion, input)
    return &Ec2CreateLaunchTemplateVersionResult{Result: future}
}

func (a *EC2Stub) CreateLocalGatewayRoute(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteInput) (*ec2.CreateLocalGatewayRouteOutput, error) {
    var output ec2.CreateLocalGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocalGatewayRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateLocalGatewayRouteAsync(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteInput) *Ec2CreateLocalGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLocalGatewayRoute, input)
    return &Ec2CreateLocalGatewayRouteResult{Result: future}
}

func (a *EC2Stub) CreateLocalGatewayRouteTableVpcAssociation(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) (*ec2.CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
    var output ec2.CreateLocalGatewayRouteTableVpcAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLocalGatewayRouteTableVpcAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateLocalGatewayRouteTableVpcAssociationAsync(ctx workflow.Context, input *ec2.CreateLocalGatewayRouteTableVpcAssociationInput) *Ec2CreateLocalGatewayRouteTableVpcAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLocalGatewayRouteTableVpcAssociation, input)
    return &Ec2CreateLocalGatewayRouteTableVpcAssociationResult{Result: future}
}

func (a *EC2Stub) CreateManagedPrefixList(ctx workflow.Context, input *ec2.CreateManagedPrefixListInput) (*ec2.CreateManagedPrefixListOutput, error) {
    var output ec2.CreateManagedPrefixListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateManagedPrefixList, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateManagedPrefixListAsync(ctx workflow.Context, input *ec2.CreateManagedPrefixListInput) *Ec2CreateManagedPrefixListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateManagedPrefixList, input)
    return &Ec2CreateManagedPrefixListResult{Result: future}
}

func (a *EC2Stub) CreateNatGateway(ctx workflow.Context, input *ec2.CreateNatGatewayInput) (*ec2.CreateNatGatewayOutput, error) {
    var output ec2.CreateNatGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNatGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateNatGatewayAsync(ctx workflow.Context, input *ec2.CreateNatGatewayInput) *Ec2CreateNatGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNatGateway, input)
    return &Ec2CreateNatGatewayResult{Result: future}
}

func (a *EC2Stub) CreateNetworkAcl(ctx workflow.Context, input *ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error) {
    var output ec2.CreateNetworkAclOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkAcl, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateNetworkAclAsync(ctx workflow.Context, input *ec2.CreateNetworkAclInput) *Ec2CreateNetworkAclResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkAcl, input)
    return &Ec2CreateNetworkAclResult{Result: future}
}

func (a *EC2Stub) CreateNetworkAclEntry(ctx workflow.Context, input *ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error) {
    var output ec2.CreateNetworkAclEntryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkAclEntry, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateNetworkAclEntryAsync(ctx workflow.Context, input *ec2.CreateNetworkAclEntryInput) *Ec2CreateNetworkAclEntryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkAclEntry, input)
    return &Ec2CreateNetworkAclEntryResult{Result: future}
}

func (a *EC2Stub) CreateNetworkInterface(ctx workflow.Context, input *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
    var output ec2.CreateNetworkInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateNetworkInterfaceAsync(ctx workflow.Context, input *ec2.CreateNetworkInterfaceInput) *Ec2CreateNetworkInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkInterface, input)
    return &Ec2CreateNetworkInterfaceResult{Result: future}
}

func (a *EC2Stub) CreateNetworkInterfacePermission(ctx workflow.Context, input *ec2.CreateNetworkInterfacePermissionInput) (*ec2.CreateNetworkInterfacePermissionOutput, error) {
    var output ec2.CreateNetworkInterfacePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkInterfacePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateNetworkInterfacePermissionAsync(ctx workflow.Context, input *ec2.CreateNetworkInterfacePermissionInput) *Ec2CreateNetworkInterfacePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkInterfacePermission, input)
    return &Ec2CreateNetworkInterfacePermissionResult{Result: future}
}

func (a *EC2Stub) CreatePlacementGroup(ctx workflow.Context, input *ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error) {
    var output ec2.CreatePlacementGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePlacementGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreatePlacementGroupAsync(ctx workflow.Context, input *ec2.CreatePlacementGroupInput) *Ec2CreatePlacementGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePlacementGroup, input)
    return &Ec2CreatePlacementGroupResult{Result: future}
}

func (a *EC2Stub) CreateReservedInstancesListing(ctx workflow.Context, input *ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error) {
    var output ec2.CreateReservedInstancesListingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReservedInstancesListing, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateReservedInstancesListingAsync(ctx workflow.Context, input *ec2.CreateReservedInstancesListingInput) *Ec2CreateReservedInstancesListingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateReservedInstancesListing, input)
    return &Ec2CreateReservedInstancesListingResult{Result: future}
}

func (a *EC2Stub) CreateRoute(ctx workflow.Context, input *ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error) {
    var output ec2.CreateRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateRouteAsync(ctx workflow.Context, input *ec2.CreateRouteInput) *Ec2CreateRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRoute, input)
    return &Ec2CreateRouteResult{Result: future}
}

func (a *EC2Stub) CreateRouteTable(ctx workflow.Context, input *ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error) {
    var output ec2.CreateRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateRouteTableAsync(ctx workflow.Context, input *ec2.CreateRouteTableInput) *Ec2CreateRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRouteTable, input)
    return &Ec2CreateRouteTableResult{Result: future}
}

func (a *EC2Stub) CreateSecurityGroup(ctx workflow.Context, input *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
    var output ec2.CreateSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSecurityGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateSecurityGroupAsync(ctx workflow.Context, input *ec2.CreateSecurityGroupInput) *Ec2CreateSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSecurityGroup, input)
    return &Ec2CreateSecurityGroupResult{Result: future}
}

func (a *EC2Stub) CreateSnapshot(ctx workflow.Context, input *ec2.CreateSnapshotInput) (*ec2.Snapshot, error) {
    var output ec2.Snapshot
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateSnapshotAsync(ctx workflow.Context, input *ec2.CreateSnapshotInput) *Ec2CreateSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshot, input)
    return &Ec2CreateSnapshotResult{Result: future}
}

func (a *EC2Stub) CreateSnapshots(ctx workflow.Context, input *ec2.CreateSnapshotsInput) (*ec2.CreateSnapshotsOutput, error) {
    var output ec2.CreateSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshots, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateSnapshotsAsync(ctx workflow.Context, input *ec2.CreateSnapshotsInput) *Ec2CreateSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSnapshots, input)
    return &Ec2CreateSnapshotsResult{Result: future}
}

func (a *EC2Stub) CreateSpotDatafeedSubscription(ctx workflow.Context, input *ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error) {
    var output ec2.CreateSpotDatafeedSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSpotDatafeedSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateSpotDatafeedSubscriptionAsync(ctx workflow.Context, input *ec2.CreateSpotDatafeedSubscriptionInput) *Ec2CreateSpotDatafeedSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSpotDatafeedSubscription, input)
    return &Ec2CreateSpotDatafeedSubscriptionResult{Result: future}
}

func (a *EC2Stub) CreateSubnet(ctx workflow.Context, input *ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error) {
    var output ec2.CreateSubnetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSubnet, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateSubnetAsync(ctx workflow.Context, input *ec2.CreateSubnetInput) *Ec2CreateSubnetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSubnet, input)
    return &Ec2CreateSubnetResult{Result: future}
}

func (a *EC2Stub) CreateTags(ctx workflow.Context, input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
    var output ec2.CreateTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTagsAsync(ctx workflow.Context, input *ec2.CreateTagsInput) *Ec2CreateTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input)
    return &Ec2CreateTagsResult{Result: future}
}

func (a *EC2Stub) CreateTrafficMirrorFilter(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterInput) (*ec2.CreateTrafficMirrorFilterOutput, error) {
    var output ec2.CreateTrafficMirrorFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTrafficMirrorFilterAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterInput) *Ec2CreateTrafficMirrorFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorFilter, input)
    return &Ec2CreateTrafficMirrorFilterResult{Result: future}
}

func (a *EC2Stub) CreateTrafficMirrorFilterRule(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterRuleInput) (*ec2.CreateTrafficMirrorFilterRuleOutput, error) {
    var output ec2.CreateTrafficMirrorFilterRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorFilterRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTrafficMirrorFilterRuleAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorFilterRuleInput) *Ec2CreateTrafficMirrorFilterRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorFilterRule, input)
    return &Ec2CreateTrafficMirrorFilterRuleResult{Result: future}
}

func (a *EC2Stub) CreateTrafficMirrorSession(ctx workflow.Context, input *ec2.CreateTrafficMirrorSessionInput) (*ec2.CreateTrafficMirrorSessionOutput, error) {
    var output ec2.CreateTrafficMirrorSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorSession, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTrafficMirrorSessionAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorSessionInput) *Ec2CreateTrafficMirrorSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorSession, input)
    return &Ec2CreateTrafficMirrorSessionResult{Result: future}
}

func (a *EC2Stub) CreateTrafficMirrorTarget(ctx workflow.Context, input *ec2.CreateTrafficMirrorTargetInput) (*ec2.CreateTrafficMirrorTargetOutput, error) {
    var output ec2.CreateTrafficMirrorTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTrafficMirrorTargetAsync(ctx workflow.Context, input *ec2.CreateTrafficMirrorTargetInput) *Ec2CreateTrafficMirrorTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrafficMirrorTarget, input)
    return &Ec2CreateTrafficMirrorTargetResult{Result: future}
}

func (a *EC2Stub) CreateTransitGateway(ctx workflow.Context, input *ec2.CreateTransitGatewayInput) (*ec2.CreateTransitGatewayOutput, error) {
    var output ec2.CreateTransitGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTransitGatewayAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayInput) *Ec2CreateTransitGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGateway, input)
    return &Ec2CreateTransitGatewayResult{Result: future}
}

func (a *EC2Stub) CreateTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.CreateTransitGatewayMulticastDomainInput) (*ec2.CreateTransitGatewayMulticastDomainOutput, error) {
    var output ec2.CreateTransitGatewayMulticastDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayMulticastDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayMulticastDomainInput) *Ec2CreateTransitGatewayMulticastDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayMulticastDomain, input)
    return &Ec2CreateTransitGatewayMulticastDomainResult{Result: future}
}

func (a *EC2Stub) CreateTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.CreateTransitGatewayPeeringAttachmentInput) (*ec2.CreateTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.CreateTransitGatewayPeeringAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayPeeringAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayPeeringAttachmentInput) *Ec2CreateTransitGatewayPeeringAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayPeeringAttachment, input)
    return &Ec2CreateTransitGatewayPeeringAttachmentResult{Result: future}
}

func (a *EC2Stub) CreateTransitGatewayPrefixListReference(ctx workflow.Context, input *ec2.CreateTransitGatewayPrefixListReferenceInput) (*ec2.CreateTransitGatewayPrefixListReferenceOutput, error) {
    var output ec2.CreateTransitGatewayPrefixListReferenceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayPrefixListReference, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTransitGatewayPrefixListReferenceAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayPrefixListReferenceInput) *Ec2CreateTransitGatewayPrefixListReferenceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayPrefixListReference, input)
    return &Ec2CreateTransitGatewayPrefixListReferenceResult{Result: future}
}

func (a *EC2Stub) CreateTransitGatewayRoute(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteInput) (*ec2.CreateTransitGatewayRouteOutput, error) {
    var output ec2.CreateTransitGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTransitGatewayRouteAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteInput) *Ec2CreateTransitGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayRoute, input)
    return &Ec2CreateTransitGatewayRouteResult{Result: future}
}

func (a *EC2Stub) CreateTransitGatewayRouteTable(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteTableInput) (*ec2.CreateTransitGatewayRouteTableOutput, error) {
    var output ec2.CreateTransitGatewayRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayRouteTableInput) *Ec2CreateTransitGatewayRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayRouteTable, input)
    return &Ec2CreateTransitGatewayRouteTableResult{Result: future}
}

func (a *EC2Stub) CreateTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.CreateTransitGatewayVpcAttachmentInput) (*ec2.CreateTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.CreateTransitGatewayVpcAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayVpcAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.CreateTransitGatewayVpcAttachmentInput) *Ec2CreateTransitGatewayVpcAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitGatewayVpcAttachment, input)
    return &Ec2CreateTransitGatewayVpcAttachmentResult{Result: future}
}

func (a *EC2Stub) CreateVolume(ctx workflow.Context, input *ec2.CreateVolumeInput) (*ec2.Volume, error) {
    var output ec2.Volume
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVolumeAsync(ctx workflow.Context, input *ec2.CreateVolumeInput) *Ec2CreateVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVolume, input)
    return &Ec2CreateVolumeResult{Result: future}
}

func (a *EC2Stub) CreateVpc(ctx workflow.Context, input *ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error) {
    var output ec2.CreateVpcOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpc, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpcAsync(ctx workflow.Context, input *ec2.CreateVpcInput) *Ec2CreateVpcResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpc, input)
    return &Ec2CreateVpcResult{Result: future}
}

func (a *EC2Stub) CreateVpcEndpoint(ctx workflow.Context, input *ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error) {
    var output ec2.CreateVpcEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpcEndpointAsync(ctx workflow.Context, input *ec2.CreateVpcEndpointInput) *Ec2CreateVpcEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcEndpoint, input)
    return &Ec2CreateVpcEndpointResult{Result: future}
}

func (a *EC2Stub) CreateVpcEndpointConnectionNotification(ctx workflow.Context, input *ec2.CreateVpcEndpointConnectionNotificationInput) (*ec2.CreateVpcEndpointConnectionNotificationOutput, error) {
    var output ec2.CreateVpcEndpointConnectionNotificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcEndpointConnectionNotification, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpcEndpointConnectionNotificationAsync(ctx workflow.Context, input *ec2.CreateVpcEndpointConnectionNotificationInput) *Ec2CreateVpcEndpointConnectionNotificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcEndpointConnectionNotification, input)
    return &Ec2CreateVpcEndpointConnectionNotificationResult{Result: future}
}

func (a *EC2Stub) CreateVpcEndpointServiceConfiguration(ctx workflow.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error) {
    var output ec2.CreateVpcEndpointServiceConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcEndpointServiceConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpcEndpointServiceConfigurationAsync(ctx workflow.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput) *Ec2CreateVpcEndpointServiceConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcEndpointServiceConfiguration, input)
    return &Ec2CreateVpcEndpointServiceConfigurationResult{Result: future}
}

func (a *EC2Stub) CreateVpcPeeringConnection(ctx workflow.Context, input *ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error) {
    var output ec2.CreateVpcPeeringConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcPeeringConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.CreateVpcPeeringConnectionInput) *Ec2CreateVpcPeeringConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcPeeringConnection, input)
    return &Ec2CreateVpcPeeringConnectionResult{Result: future}
}

func (a *EC2Stub) CreateVpnConnection(ctx workflow.Context, input *ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error) {
    var output ec2.CreateVpnConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpnConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpnConnectionAsync(ctx workflow.Context, input *ec2.CreateVpnConnectionInput) *Ec2CreateVpnConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpnConnection, input)
    return &Ec2CreateVpnConnectionResult{Result: future}
}

func (a *EC2Stub) CreateVpnConnectionRoute(ctx workflow.Context, input *ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error) {
    var output ec2.CreateVpnConnectionRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpnConnectionRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpnConnectionRouteAsync(ctx workflow.Context, input *ec2.CreateVpnConnectionRouteInput) *Ec2CreateVpnConnectionRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpnConnectionRoute, input)
    return &Ec2CreateVpnConnectionRouteResult{Result: future}
}

func (a *EC2Stub) CreateVpnGateway(ctx workflow.Context, input *ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error) {
    var output ec2.CreateVpnGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpnGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) CreateVpnGatewayAsync(ctx workflow.Context, input *ec2.CreateVpnGatewayInput) *Ec2CreateVpnGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpnGateway, input)
    return &Ec2CreateVpnGatewayResult{Result: future}
}

func (a *EC2Stub) DeleteCarrierGateway(ctx workflow.Context, input *ec2.DeleteCarrierGatewayInput) (*ec2.DeleteCarrierGatewayOutput, error) {
    var output ec2.DeleteCarrierGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCarrierGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteCarrierGatewayAsync(ctx workflow.Context, input *ec2.DeleteCarrierGatewayInput) *Ec2DeleteCarrierGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCarrierGateway, input)
    return &Ec2DeleteCarrierGatewayResult{Result: future}
}

func (a *EC2Stub) DeleteClientVpnEndpoint(ctx workflow.Context, input *ec2.DeleteClientVpnEndpointInput) (*ec2.DeleteClientVpnEndpointOutput, error) {
    var output ec2.DeleteClientVpnEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteClientVpnEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteClientVpnEndpointAsync(ctx workflow.Context, input *ec2.DeleteClientVpnEndpointInput) *Ec2DeleteClientVpnEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteClientVpnEndpoint, input)
    return &Ec2DeleteClientVpnEndpointResult{Result: future}
}

func (a *EC2Stub) DeleteClientVpnRoute(ctx workflow.Context, input *ec2.DeleteClientVpnRouteInput) (*ec2.DeleteClientVpnRouteOutput, error) {
    var output ec2.DeleteClientVpnRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteClientVpnRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteClientVpnRouteAsync(ctx workflow.Context, input *ec2.DeleteClientVpnRouteInput) *Ec2DeleteClientVpnRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteClientVpnRoute, input)
    return &Ec2DeleteClientVpnRouteResult{Result: future}
}

func (a *EC2Stub) DeleteCustomerGateway(ctx workflow.Context, input *ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error) {
    var output ec2.DeleteCustomerGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomerGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteCustomerGatewayAsync(ctx workflow.Context, input *ec2.DeleteCustomerGatewayInput) *Ec2DeleteCustomerGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomerGateway, input)
    return &Ec2DeleteCustomerGatewayResult{Result: future}
}

func (a *EC2Stub) DeleteDhcpOptions(ctx workflow.Context, input *ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error) {
    var output ec2.DeleteDhcpOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDhcpOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteDhcpOptionsAsync(ctx workflow.Context, input *ec2.DeleteDhcpOptionsInput) *Ec2DeleteDhcpOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDhcpOptions, input)
    return &Ec2DeleteDhcpOptionsResult{Result: future}
}

func (a *EC2Stub) DeleteEgressOnlyInternetGateway(ctx workflow.Context, input *ec2.DeleteEgressOnlyInternetGatewayInput) (*ec2.DeleteEgressOnlyInternetGatewayOutput, error) {
    var output ec2.DeleteEgressOnlyInternetGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEgressOnlyInternetGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteEgressOnlyInternetGatewayAsync(ctx workflow.Context, input *ec2.DeleteEgressOnlyInternetGatewayInput) *Ec2DeleteEgressOnlyInternetGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEgressOnlyInternetGateway, input)
    return &Ec2DeleteEgressOnlyInternetGatewayResult{Result: future}
}

func (a *EC2Stub) DeleteFleets(ctx workflow.Context, input *ec2.DeleteFleetsInput) (*ec2.DeleteFleetsOutput, error) {
    var output ec2.DeleteFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteFleetsAsync(ctx workflow.Context, input *ec2.DeleteFleetsInput) *Ec2DeleteFleetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFleets, input)
    return &Ec2DeleteFleetsResult{Result: future}
}

func (a *EC2Stub) DeleteFlowLogs(ctx workflow.Context, input *ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error) {
    var output ec2.DeleteFlowLogsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFlowLogs, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteFlowLogsAsync(ctx workflow.Context, input *ec2.DeleteFlowLogsInput) *Ec2DeleteFlowLogsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFlowLogs, input)
    return &Ec2DeleteFlowLogsResult{Result: future}
}

func (a *EC2Stub) DeleteFpgaImage(ctx workflow.Context, input *ec2.DeleteFpgaImageInput) (*ec2.DeleteFpgaImageOutput, error) {
    var output ec2.DeleteFpgaImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFpgaImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteFpgaImageAsync(ctx workflow.Context, input *ec2.DeleteFpgaImageInput) *Ec2DeleteFpgaImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFpgaImage, input)
    return &Ec2DeleteFpgaImageResult{Result: future}
}

func (a *EC2Stub) DeleteInternetGateway(ctx workflow.Context, input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
    var output ec2.DeleteInternetGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInternetGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteInternetGatewayAsync(ctx workflow.Context, input *ec2.DeleteInternetGatewayInput) *Ec2DeleteInternetGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteInternetGateway, input)
    return &Ec2DeleteInternetGatewayResult{Result: future}
}

func (a *EC2Stub) DeleteKeyPair(ctx workflow.Context, input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
    var output ec2.DeleteKeyPairOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteKeyPair, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteKeyPairAsync(ctx workflow.Context, input *ec2.DeleteKeyPairInput) *Ec2DeleteKeyPairResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteKeyPair, input)
    return &Ec2DeleteKeyPairResult{Result: future}
}

func (a *EC2Stub) DeleteLaunchTemplate(ctx workflow.Context, input *ec2.DeleteLaunchTemplateInput) (*ec2.DeleteLaunchTemplateOutput, error) {
    var output ec2.DeleteLaunchTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLaunchTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteLaunchTemplateAsync(ctx workflow.Context, input *ec2.DeleteLaunchTemplateInput) *Ec2DeleteLaunchTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLaunchTemplate, input)
    return &Ec2DeleteLaunchTemplateResult{Result: future}
}

func (a *EC2Stub) DeleteLaunchTemplateVersions(ctx workflow.Context, input *ec2.DeleteLaunchTemplateVersionsInput) (*ec2.DeleteLaunchTemplateVersionsOutput, error) {
    var output ec2.DeleteLaunchTemplateVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLaunchTemplateVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteLaunchTemplateVersionsAsync(ctx workflow.Context, input *ec2.DeleteLaunchTemplateVersionsInput) *Ec2DeleteLaunchTemplateVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLaunchTemplateVersions, input)
    return &Ec2DeleteLaunchTemplateVersionsResult{Result: future}
}

func (a *EC2Stub) DeleteLocalGatewayRoute(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteInput) (*ec2.DeleteLocalGatewayRouteOutput, error) {
    var output ec2.DeleteLocalGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLocalGatewayRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteLocalGatewayRouteAsync(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteInput) *Ec2DeleteLocalGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLocalGatewayRoute, input)
    return &Ec2DeleteLocalGatewayRouteResult{Result: future}
}

func (a *EC2Stub) DeleteLocalGatewayRouteTableVpcAssociation(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) (*ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput, error) {
    var output ec2.DeleteLocalGatewayRouteTableVpcAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLocalGatewayRouteTableVpcAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteLocalGatewayRouteTableVpcAssociationAsync(ctx workflow.Context, input *ec2.DeleteLocalGatewayRouteTableVpcAssociationInput) *Ec2DeleteLocalGatewayRouteTableVpcAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLocalGatewayRouteTableVpcAssociation, input)
    return &Ec2DeleteLocalGatewayRouteTableVpcAssociationResult{Result: future}
}

func (a *EC2Stub) DeleteManagedPrefixList(ctx workflow.Context, input *ec2.DeleteManagedPrefixListInput) (*ec2.DeleteManagedPrefixListOutput, error) {
    var output ec2.DeleteManagedPrefixListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteManagedPrefixList, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteManagedPrefixListAsync(ctx workflow.Context, input *ec2.DeleteManagedPrefixListInput) *Ec2DeleteManagedPrefixListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteManagedPrefixList, input)
    return &Ec2DeleteManagedPrefixListResult{Result: future}
}

func (a *EC2Stub) DeleteNatGateway(ctx workflow.Context, input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
    var output ec2.DeleteNatGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNatGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteNatGatewayAsync(ctx workflow.Context, input *ec2.DeleteNatGatewayInput) *Ec2DeleteNatGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNatGateway, input)
    return &Ec2DeleteNatGatewayResult{Result: future}
}

func (a *EC2Stub) DeleteNetworkAcl(ctx workflow.Context, input *ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error) {
    var output ec2.DeleteNetworkAclOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkAcl, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteNetworkAclAsync(ctx workflow.Context, input *ec2.DeleteNetworkAclInput) *Ec2DeleteNetworkAclResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkAcl, input)
    return &Ec2DeleteNetworkAclResult{Result: future}
}

func (a *EC2Stub) DeleteNetworkAclEntry(ctx workflow.Context, input *ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error) {
    var output ec2.DeleteNetworkAclEntryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkAclEntry, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteNetworkAclEntryAsync(ctx workflow.Context, input *ec2.DeleteNetworkAclEntryInput) *Ec2DeleteNetworkAclEntryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkAclEntry, input)
    return &Ec2DeleteNetworkAclEntryResult{Result: future}
}

func (a *EC2Stub) DeleteNetworkInterface(ctx workflow.Context, input *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
    var output ec2.DeleteNetworkInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteNetworkInterfaceAsync(ctx workflow.Context, input *ec2.DeleteNetworkInterfaceInput) *Ec2DeleteNetworkInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkInterface, input)
    return &Ec2DeleteNetworkInterfaceResult{Result: future}
}

func (a *EC2Stub) DeleteNetworkInterfacePermission(ctx workflow.Context, input *ec2.DeleteNetworkInterfacePermissionInput) (*ec2.DeleteNetworkInterfacePermissionOutput, error) {
    var output ec2.DeleteNetworkInterfacePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkInterfacePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteNetworkInterfacePermissionAsync(ctx workflow.Context, input *ec2.DeleteNetworkInterfacePermissionInput) *Ec2DeleteNetworkInterfacePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkInterfacePermission, input)
    return &Ec2DeleteNetworkInterfacePermissionResult{Result: future}
}

func (a *EC2Stub) DeletePlacementGroup(ctx workflow.Context, input *ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error) {
    var output ec2.DeletePlacementGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePlacementGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeletePlacementGroupAsync(ctx workflow.Context, input *ec2.DeletePlacementGroupInput) *Ec2DeletePlacementGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePlacementGroup, input)
    return &Ec2DeletePlacementGroupResult{Result: future}
}

func (a *EC2Stub) DeleteQueuedReservedInstances(ctx workflow.Context, input *ec2.DeleteQueuedReservedInstancesInput) (*ec2.DeleteQueuedReservedInstancesOutput, error) {
    var output ec2.DeleteQueuedReservedInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteQueuedReservedInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteQueuedReservedInstancesAsync(ctx workflow.Context, input *ec2.DeleteQueuedReservedInstancesInput) *Ec2DeleteQueuedReservedInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteQueuedReservedInstances, input)
    return &Ec2DeleteQueuedReservedInstancesResult{Result: future}
}

func (a *EC2Stub) DeleteRoute(ctx workflow.Context, input *ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error) {
    var output ec2.DeleteRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteRouteAsync(ctx workflow.Context, input *ec2.DeleteRouteInput) *Ec2DeleteRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRoute, input)
    return &Ec2DeleteRouteResult{Result: future}
}

func (a *EC2Stub) DeleteRouteTable(ctx workflow.Context, input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
    var output ec2.DeleteRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteRouteTableAsync(ctx workflow.Context, input *ec2.DeleteRouteTableInput) *Ec2DeleteRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRouteTable, input)
    return &Ec2DeleteRouteTableResult{Result: future}
}

func (a *EC2Stub) DeleteSecurityGroup(ctx workflow.Context, input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
    var output ec2.DeleteSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSecurityGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteSecurityGroupAsync(ctx workflow.Context, input *ec2.DeleteSecurityGroupInput) *Ec2DeleteSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSecurityGroup, input)
    return &Ec2DeleteSecurityGroupResult{Result: future}
}

func (a *EC2Stub) DeleteSnapshot(ctx workflow.Context, input *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
    var output ec2.DeleteSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteSnapshotAsync(ctx workflow.Context, input *ec2.DeleteSnapshotInput) *Ec2DeleteSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSnapshot, input)
    return &Ec2DeleteSnapshotResult{Result: future}
}

func (a *EC2Stub) DeleteSpotDatafeedSubscription(ctx workflow.Context, input *ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error) {
    var output ec2.DeleteSpotDatafeedSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSpotDatafeedSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteSpotDatafeedSubscriptionAsync(ctx workflow.Context, input *ec2.DeleteSpotDatafeedSubscriptionInput) *Ec2DeleteSpotDatafeedSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSpotDatafeedSubscription, input)
    return &Ec2DeleteSpotDatafeedSubscriptionResult{Result: future}
}

func (a *EC2Stub) DeleteSubnet(ctx workflow.Context, input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
    var output ec2.DeleteSubnetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSubnet, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteSubnetAsync(ctx workflow.Context, input *ec2.DeleteSubnetInput) *Ec2DeleteSubnetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSubnet, input)
    return &Ec2DeleteSubnetResult{Result: future}
}

func (a *EC2Stub) DeleteTags(ctx workflow.Context, input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
    var output ec2.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTagsAsync(ctx workflow.Context, input *ec2.DeleteTagsInput) *Ec2DeleteTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input)
    return &Ec2DeleteTagsResult{Result: future}
}

func (a *EC2Stub) DeleteTrafficMirrorFilter(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterInput) (*ec2.DeleteTrafficMirrorFilterOutput, error) {
    var output ec2.DeleteTrafficMirrorFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTrafficMirrorFilterAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterInput) *Ec2DeleteTrafficMirrorFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorFilter, input)
    return &Ec2DeleteTrafficMirrorFilterResult{Result: future}
}

func (a *EC2Stub) DeleteTrafficMirrorFilterRule(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterRuleInput) (*ec2.DeleteTrafficMirrorFilterRuleOutput, error) {
    var output ec2.DeleteTrafficMirrorFilterRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorFilterRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTrafficMirrorFilterRuleAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorFilterRuleInput) *Ec2DeleteTrafficMirrorFilterRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorFilterRule, input)
    return &Ec2DeleteTrafficMirrorFilterRuleResult{Result: future}
}

func (a *EC2Stub) DeleteTrafficMirrorSession(ctx workflow.Context, input *ec2.DeleteTrafficMirrorSessionInput) (*ec2.DeleteTrafficMirrorSessionOutput, error) {
    var output ec2.DeleteTrafficMirrorSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorSession, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTrafficMirrorSessionAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorSessionInput) *Ec2DeleteTrafficMirrorSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorSession, input)
    return &Ec2DeleteTrafficMirrorSessionResult{Result: future}
}

func (a *EC2Stub) DeleteTrafficMirrorTarget(ctx workflow.Context, input *ec2.DeleteTrafficMirrorTargetInput) (*ec2.DeleteTrafficMirrorTargetOutput, error) {
    var output ec2.DeleteTrafficMirrorTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTrafficMirrorTargetAsync(ctx workflow.Context, input *ec2.DeleteTrafficMirrorTargetInput) *Ec2DeleteTrafficMirrorTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrafficMirrorTarget, input)
    return &Ec2DeleteTrafficMirrorTargetResult{Result: future}
}

func (a *EC2Stub) DeleteTransitGateway(ctx workflow.Context, input *ec2.DeleteTransitGatewayInput) (*ec2.DeleteTransitGatewayOutput, error) {
    var output ec2.DeleteTransitGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTransitGatewayAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayInput) *Ec2DeleteTransitGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGateway, input)
    return &Ec2DeleteTransitGatewayResult{Result: future}
}

func (a *EC2Stub) DeleteTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.DeleteTransitGatewayMulticastDomainInput) (*ec2.DeleteTransitGatewayMulticastDomainOutput, error) {
    var output ec2.DeleteTransitGatewayMulticastDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayMulticastDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayMulticastDomainInput) *Ec2DeleteTransitGatewayMulticastDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayMulticastDomain, input)
    return &Ec2DeleteTransitGatewayMulticastDomainResult{Result: future}
}

func (a *EC2Stub) DeleteTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.DeleteTransitGatewayPeeringAttachmentInput) (*ec2.DeleteTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.DeleteTransitGatewayPeeringAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayPeeringAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayPeeringAttachmentInput) *Ec2DeleteTransitGatewayPeeringAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayPeeringAttachment, input)
    return &Ec2DeleteTransitGatewayPeeringAttachmentResult{Result: future}
}

func (a *EC2Stub) DeleteTransitGatewayPrefixListReference(ctx workflow.Context, input *ec2.DeleteTransitGatewayPrefixListReferenceInput) (*ec2.DeleteTransitGatewayPrefixListReferenceOutput, error) {
    var output ec2.DeleteTransitGatewayPrefixListReferenceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayPrefixListReference, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTransitGatewayPrefixListReferenceAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayPrefixListReferenceInput) *Ec2DeleteTransitGatewayPrefixListReferenceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayPrefixListReference, input)
    return &Ec2DeleteTransitGatewayPrefixListReferenceResult{Result: future}
}

func (a *EC2Stub) DeleteTransitGatewayRoute(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteInput) (*ec2.DeleteTransitGatewayRouteOutput, error) {
    var output ec2.DeleteTransitGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTransitGatewayRouteAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteInput) *Ec2DeleteTransitGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayRoute, input)
    return &Ec2DeleteTransitGatewayRouteResult{Result: future}
}

func (a *EC2Stub) DeleteTransitGatewayRouteTable(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteTableInput) (*ec2.DeleteTransitGatewayRouteTableOutput, error) {
    var output ec2.DeleteTransitGatewayRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayRouteTableInput) *Ec2DeleteTransitGatewayRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayRouteTable, input)
    return &Ec2DeleteTransitGatewayRouteTableResult{Result: future}
}

func (a *EC2Stub) DeleteTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.DeleteTransitGatewayVpcAttachmentInput) (*ec2.DeleteTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.DeleteTransitGatewayVpcAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayVpcAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.DeleteTransitGatewayVpcAttachmentInput) *Ec2DeleteTransitGatewayVpcAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTransitGatewayVpcAttachment, input)
    return &Ec2DeleteTransitGatewayVpcAttachmentResult{Result: future}
}

func (a *EC2Stub) DeleteVolume(ctx workflow.Context, input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
    var output ec2.DeleteVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVolumeAsync(ctx workflow.Context, input *ec2.DeleteVolumeInput) *Ec2DeleteVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVolume, input)
    return &Ec2DeleteVolumeResult{Result: future}
}

func (a *EC2Stub) DeleteVpc(ctx workflow.Context, input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
    var output ec2.DeleteVpcOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpc, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpcAsync(ctx workflow.Context, input *ec2.DeleteVpcInput) *Ec2DeleteVpcResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpc, input)
    return &Ec2DeleteVpcResult{Result: future}
}

func (a *EC2Stub) DeleteVpcEndpointConnectionNotifications(ctx workflow.Context, input *ec2.DeleteVpcEndpointConnectionNotificationsInput) (*ec2.DeleteVpcEndpointConnectionNotificationsOutput, error) {
    var output ec2.DeleteVpcEndpointConnectionNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcEndpointConnectionNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpcEndpointConnectionNotificationsAsync(ctx workflow.Context, input *ec2.DeleteVpcEndpointConnectionNotificationsInput) *Ec2DeleteVpcEndpointConnectionNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcEndpointConnectionNotifications, input)
    return &Ec2DeleteVpcEndpointConnectionNotificationsResult{Result: future}
}

func (a *EC2Stub) DeleteVpcEndpointServiceConfigurations(ctx workflow.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error) {
    var output ec2.DeleteVpcEndpointServiceConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcEndpointServiceConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpcEndpointServiceConfigurationsAsync(ctx workflow.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput) *Ec2DeleteVpcEndpointServiceConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcEndpointServiceConfigurations, input)
    return &Ec2DeleteVpcEndpointServiceConfigurationsResult{Result: future}
}

func (a *EC2Stub) DeleteVpcEndpoints(ctx workflow.Context, input *ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error) {
    var output ec2.DeleteVpcEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpcEndpointsAsync(ctx workflow.Context, input *ec2.DeleteVpcEndpointsInput) *Ec2DeleteVpcEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcEndpoints, input)
    return &Ec2DeleteVpcEndpointsResult{Result: future}
}

func (a *EC2Stub) DeleteVpcPeeringConnection(ctx workflow.Context, input *ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error) {
    var output ec2.DeleteVpcPeeringConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcPeeringConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.DeleteVpcPeeringConnectionInput) *Ec2DeleteVpcPeeringConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcPeeringConnection, input)
    return &Ec2DeleteVpcPeeringConnectionResult{Result: future}
}

func (a *EC2Stub) DeleteVpnConnection(ctx workflow.Context, input *ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error) {
    var output ec2.DeleteVpnConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpnConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpnConnectionAsync(ctx workflow.Context, input *ec2.DeleteVpnConnectionInput) *Ec2DeleteVpnConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpnConnection, input)
    return &Ec2DeleteVpnConnectionResult{Result: future}
}

func (a *EC2Stub) DeleteVpnConnectionRoute(ctx workflow.Context, input *ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error) {
    var output ec2.DeleteVpnConnectionRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpnConnectionRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpnConnectionRouteAsync(ctx workflow.Context, input *ec2.DeleteVpnConnectionRouteInput) *Ec2DeleteVpnConnectionRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpnConnectionRoute, input)
    return &Ec2DeleteVpnConnectionRouteResult{Result: future}
}

func (a *EC2Stub) DeleteVpnGateway(ctx workflow.Context, input *ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error) {
    var output ec2.DeleteVpnGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpnGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeleteVpnGatewayAsync(ctx workflow.Context, input *ec2.DeleteVpnGatewayInput) *Ec2DeleteVpnGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpnGateway, input)
    return &Ec2DeleteVpnGatewayResult{Result: future}
}

func (a *EC2Stub) DeprovisionByoipCidr(ctx workflow.Context, input *ec2.DeprovisionByoipCidrInput) (*ec2.DeprovisionByoipCidrOutput, error) {
    var output ec2.DeprovisionByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeprovisionByoipCidr, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeprovisionByoipCidrAsync(ctx workflow.Context, input *ec2.DeprovisionByoipCidrInput) *Ec2DeprovisionByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeprovisionByoipCidr, input)
    return &Ec2DeprovisionByoipCidrResult{Result: future}
}

func (a *EC2Stub) DeregisterImage(ctx workflow.Context, input *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
    var output ec2.DeregisterImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeregisterImageAsync(ctx workflow.Context, input *ec2.DeregisterImageInput) *Ec2DeregisterImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterImage, input)
    return &Ec2DeregisterImageResult{Result: future}
}

func (a *EC2Stub) DeregisterInstanceEventNotificationAttributes(ctx workflow.Context, input *ec2.DeregisterInstanceEventNotificationAttributesInput) (*ec2.DeregisterInstanceEventNotificationAttributesOutput, error) {
    var output ec2.DeregisterInstanceEventNotificationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterInstanceEventNotificationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeregisterInstanceEventNotificationAttributesAsync(ctx workflow.Context, input *ec2.DeregisterInstanceEventNotificationAttributesInput) *Ec2DeregisterInstanceEventNotificationAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterInstanceEventNotificationAttributes, input)
    return &Ec2DeregisterInstanceEventNotificationAttributesResult{Result: future}
}

func (a *EC2Stub) DeregisterTransitGatewayMulticastGroupMembers(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) (*ec2.DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
    var output ec2.DeregisterTransitGatewayMulticastGroupMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterTransitGatewayMulticastGroupMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeregisterTransitGatewayMulticastGroupMembersAsync(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupMembersInput) *Ec2DeregisterTransitGatewayMulticastGroupMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterTransitGatewayMulticastGroupMembers, input)
    return &Ec2DeregisterTransitGatewayMulticastGroupMembersResult{Result: future}
}

func (a *EC2Stub) DeregisterTransitGatewayMulticastGroupSources(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) (*ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput, error) {
    var output ec2.DeregisterTransitGatewayMulticastGroupSourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterTransitGatewayMulticastGroupSources, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DeregisterTransitGatewayMulticastGroupSourcesAsync(ctx workflow.Context, input *ec2.DeregisterTransitGatewayMulticastGroupSourcesInput) *Ec2DeregisterTransitGatewayMulticastGroupSourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterTransitGatewayMulticastGroupSources, input)
    return &Ec2DeregisterTransitGatewayMulticastGroupSourcesResult{Result: future}
}

func (a *EC2Stub) DescribeAccountAttributes(ctx workflow.Context, input *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
    var output ec2.DescribeAccountAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeAccountAttributesAsync(ctx workflow.Context, input *ec2.DescribeAccountAttributesInput) *Ec2DescribeAccountAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input)
    return &Ec2DescribeAccountAttributesResult{Result: future}
}

func (a *EC2Stub) DescribeAddresses(ctx workflow.Context, input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
    var output ec2.DescribeAddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAddresses, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeAddressesAsync(ctx workflow.Context, input *ec2.DescribeAddressesInput) *Ec2DescribeAddressesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAddresses, input)
    return &Ec2DescribeAddressesResult{Result: future}
}

func (a *EC2Stub) DescribeAggregateIdFormat(ctx workflow.Context, input *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error) {
    var output ec2.DescribeAggregateIdFormatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAggregateIdFormat, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeAggregateIdFormatAsync(ctx workflow.Context, input *ec2.DescribeAggregateIdFormatInput) *Ec2DescribeAggregateIdFormatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAggregateIdFormat, input)
    return &Ec2DescribeAggregateIdFormatResult{Result: future}
}

func (a *EC2Stub) DescribeAvailabilityZones(ctx workflow.Context, input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
    var output ec2.DescribeAvailabilityZonesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAvailabilityZones, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeAvailabilityZonesAsync(ctx workflow.Context, input *ec2.DescribeAvailabilityZonesInput) *Ec2DescribeAvailabilityZonesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAvailabilityZones, input)
    return &Ec2DescribeAvailabilityZonesResult{Result: future}
}

func (a *EC2Stub) DescribeBundleTasks(ctx workflow.Context, input *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
    var output ec2.DescribeBundleTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBundleTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeBundleTasksAsync(ctx workflow.Context, input *ec2.DescribeBundleTasksInput) *Ec2DescribeBundleTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBundleTasks, input)
    return &Ec2DescribeBundleTasksResult{Result: future}
}

func (a *EC2Stub) DescribeByoipCidrs(ctx workflow.Context, input *ec2.DescribeByoipCidrsInput) (*ec2.DescribeByoipCidrsOutput, error) {
    var output ec2.DescribeByoipCidrsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeByoipCidrs, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeByoipCidrsAsync(ctx workflow.Context, input *ec2.DescribeByoipCidrsInput) *Ec2DescribeByoipCidrsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeByoipCidrs, input)
    return &Ec2DescribeByoipCidrsResult{Result: future}
}

func (a *EC2Stub) DescribeCapacityReservations(ctx workflow.Context, input *ec2.DescribeCapacityReservationsInput) (*ec2.DescribeCapacityReservationsOutput, error) {
    var output ec2.DescribeCapacityReservationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCapacityReservations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeCapacityReservationsAsync(ctx workflow.Context, input *ec2.DescribeCapacityReservationsInput) *Ec2DescribeCapacityReservationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCapacityReservations, input)
    return &Ec2DescribeCapacityReservationsResult{Result: future}
}

func (a *EC2Stub) DescribeCarrierGateways(ctx workflow.Context, input *ec2.DescribeCarrierGatewaysInput) (*ec2.DescribeCarrierGatewaysOutput, error) {
    var output ec2.DescribeCarrierGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCarrierGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeCarrierGatewaysAsync(ctx workflow.Context, input *ec2.DescribeCarrierGatewaysInput) *Ec2DescribeCarrierGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCarrierGateways, input)
    return &Ec2DescribeCarrierGatewaysResult{Result: future}
}

func (a *EC2Stub) DescribeClassicLinkInstances(ctx workflow.Context, input *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
    var output ec2.DescribeClassicLinkInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClassicLinkInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeClassicLinkInstancesAsync(ctx workflow.Context, input *ec2.DescribeClassicLinkInstancesInput) *Ec2DescribeClassicLinkInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClassicLinkInstances, input)
    return &Ec2DescribeClassicLinkInstancesResult{Result: future}
}

func (a *EC2Stub) DescribeClientVpnAuthorizationRules(ctx workflow.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
    var output ec2.DescribeClientVpnAuthorizationRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnAuthorizationRules, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeClientVpnAuthorizationRulesAsync(ctx workflow.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput) *Ec2DescribeClientVpnAuthorizationRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnAuthorizationRules, input)
    return &Ec2DescribeClientVpnAuthorizationRulesResult{Result: future}
}

func (a *EC2Stub) DescribeClientVpnConnections(ctx workflow.Context, input *ec2.DescribeClientVpnConnectionsInput) (*ec2.DescribeClientVpnConnectionsOutput, error) {
    var output ec2.DescribeClientVpnConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeClientVpnConnectionsAsync(ctx workflow.Context, input *ec2.DescribeClientVpnConnectionsInput) *Ec2DescribeClientVpnConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnConnections, input)
    return &Ec2DescribeClientVpnConnectionsResult{Result: future}
}

func (a *EC2Stub) DescribeClientVpnEndpoints(ctx workflow.Context, input *ec2.DescribeClientVpnEndpointsInput) (*ec2.DescribeClientVpnEndpointsOutput, error) {
    var output ec2.DescribeClientVpnEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeClientVpnEndpointsAsync(ctx workflow.Context, input *ec2.DescribeClientVpnEndpointsInput) *Ec2DescribeClientVpnEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnEndpoints, input)
    return &Ec2DescribeClientVpnEndpointsResult{Result: future}
}

func (a *EC2Stub) DescribeClientVpnRoutes(ctx workflow.Context, input *ec2.DescribeClientVpnRoutesInput) (*ec2.DescribeClientVpnRoutesOutput, error) {
    var output ec2.DescribeClientVpnRoutesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnRoutes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeClientVpnRoutesAsync(ctx workflow.Context, input *ec2.DescribeClientVpnRoutesInput) *Ec2DescribeClientVpnRoutesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnRoutes, input)
    return &Ec2DescribeClientVpnRoutesResult{Result: future}
}

func (a *EC2Stub) DescribeClientVpnTargetNetworks(ctx workflow.Context, input *ec2.DescribeClientVpnTargetNetworksInput) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
    var output ec2.DescribeClientVpnTargetNetworksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnTargetNetworks, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeClientVpnTargetNetworksAsync(ctx workflow.Context, input *ec2.DescribeClientVpnTargetNetworksInput) *Ec2DescribeClientVpnTargetNetworksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClientVpnTargetNetworks, input)
    return &Ec2DescribeClientVpnTargetNetworksResult{Result: future}
}

func (a *EC2Stub) DescribeCoipPools(ctx workflow.Context, input *ec2.DescribeCoipPoolsInput) (*ec2.DescribeCoipPoolsOutput, error) {
    var output ec2.DescribeCoipPoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCoipPools, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeCoipPoolsAsync(ctx workflow.Context, input *ec2.DescribeCoipPoolsInput) *Ec2DescribeCoipPoolsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCoipPools, input)
    return &Ec2DescribeCoipPoolsResult{Result: future}
}

func (a *EC2Stub) DescribeConversionTasks(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
    var output ec2.DescribeConversionTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConversionTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeConversionTasksAsync(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) *Ec2DescribeConversionTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConversionTasks, input)
    return &Ec2DescribeConversionTasksResult{Result: future}
}

func (a *EC2Stub) DescribeCustomerGateways(ctx workflow.Context, input *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
    var output ec2.DescribeCustomerGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCustomerGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeCustomerGatewaysAsync(ctx workflow.Context, input *ec2.DescribeCustomerGatewaysInput) *Ec2DescribeCustomerGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCustomerGateways, input)
    return &Ec2DescribeCustomerGatewaysResult{Result: future}
}

func (a *EC2Stub) DescribeDhcpOptions(ctx workflow.Context, input *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
    var output ec2.DescribeDhcpOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDhcpOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeDhcpOptionsAsync(ctx workflow.Context, input *ec2.DescribeDhcpOptionsInput) *Ec2DescribeDhcpOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDhcpOptions, input)
    return &Ec2DescribeDhcpOptionsResult{Result: future}
}

func (a *EC2Stub) DescribeEgressOnlyInternetGateways(ctx workflow.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
    var output ec2.DescribeEgressOnlyInternetGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEgressOnlyInternetGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeEgressOnlyInternetGatewaysAsync(ctx workflow.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput) *Ec2DescribeEgressOnlyInternetGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEgressOnlyInternetGateways, input)
    return &Ec2DescribeEgressOnlyInternetGatewaysResult{Result: future}
}

func (a *EC2Stub) DescribeElasticGpus(ctx workflow.Context, input *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
    var output ec2.DescribeElasticGpusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticGpus, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeElasticGpusAsync(ctx workflow.Context, input *ec2.DescribeElasticGpusInput) *Ec2DescribeElasticGpusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticGpus, input)
    return &Ec2DescribeElasticGpusResult{Result: future}
}

func (a *EC2Stub) DescribeExportImageTasks(ctx workflow.Context, input *ec2.DescribeExportImageTasksInput) (*ec2.DescribeExportImageTasksOutput, error) {
    var output ec2.DescribeExportImageTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExportImageTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeExportImageTasksAsync(ctx workflow.Context, input *ec2.DescribeExportImageTasksInput) *Ec2DescribeExportImageTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeExportImageTasks, input)
    return &Ec2DescribeExportImageTasksResult{Result: future}
}

func (a *EC2Stub) DescribeExportTasks(ctx workflow.Context, input *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
    var output ec2.DescribeExportTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExportTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeExportTasksAsync(ctx workflow.Context, input *ec2.DescribeExportTasksInput) *Ec2DescribeExportTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeExportTasks, input)
    return &Ec2DescribeExportTasksResult{Result: future}
}

func (a *EC2Stub) DescribeFastSnapshotRestores(ctx workflow.Context, input *ec2.DescribeFastSnapshotRestoresInput) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
    var output ec2.DescribeFastSnapshotRestoresOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFastSnapshotRestores, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeFastSnapshotRestoresAsync(ctx workflow.Context, input *ec2.DescribeFastSnapshotRestoresInput) *Ec2DescribeFastSnapshotRestoresResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFastSnapshotRestores, input)
    return &Ec2DescribeFastSnapshotRestoresResult{Result: future}
}

func (a *EC2Stub) DescribeFleetHistory(ctx workflow.Context, input *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error) {
    var output ec2.DescribeFleetHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeFleetHistoryAsync(ctx workflow.Context, input *ec2.DescribeFleetHistoryInput) *Ec2DescribeFleetHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetHistory, input)
    return &Ec2DescribeFleetHistoryResult{Result: future}
}

func (a *EC2Stub) DescribeFleetInstances(ctx workflow.Context, input *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error) {
    var output ec2.DescribeFleetInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeFleetInstancesAsync(ctx workflow.Context, input *ec2.DescribeFleetInstancesInput) *Ec2DescribeFleetInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetInstances, input)
    return &Ec2DescribeFleetInstancesResult{Result: future}
}

func (a *EC2Stub) DescribeFleets(ctx workflow.Context, input *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error) {
    var output ec2.DescribeFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeFleetsAsync(ctx workflow.Context, input *ec2.DescribeFleetsInput) *Ec2DescribeFleetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleets, input)
    return &Ec2DescribeFleetsResult{Result: future}
}

func (a *EC2Stub) DescribeFlowLogs(ctx workflow.Context, input *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
    var output ec2.DescribeFlowLogsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFlowLogs, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeFlowLogsAsync(ctx workflow.Context, input *ec2.DescribeFlowLogsInput) *Ec2DescribeFlowLogsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFlowLogs, input)
    return &Ec2DescribeFlowLogsResult{Result: future}
}

func (a *EC2Stub) DescribeFpgaImageAttribute(ctx workflow.Context, input *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
    var output ec2.DescribeFpgaImageAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFpgaImageAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeFpgaImageAttributeAsync(ctx workflow.Context, input *ec2.DescribeFpgaImageAttributeInput) *Ec2DescribeFpgaImageAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFpgaImageAttribute, input)
    return &Ec2DescribeFpgaImageAttributeResult{Result: future}
}

func (a *EC2Stub) DescribeFpgaImages(ctx workflow.Context, input *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
    var output ec2.DescribeFpgaImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFpgaImages, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeFpgaImagesAsync(ctx workflow.Context, input *ec2.DescribeFpgaImagesInput) *Ec2DescribeFpgaImagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFpgaImages, input)
    return &Ec2DescribeFpgaImagesResult{Result: future}
}

func (a *EC2Stub) DescribeHostReservationOfferings(ctx workflow.Context, input *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
    var output ec2.DescribeHostReservationOfferingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHostReservationOfferings, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeHostReservationOfferingsAsync(ctx workflow.Context, input *ec2.DescribeHostReservationOfferingsInput) *Ec2DescribeHostReservationOfferingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHostReservationOfferings, input)
    return &Ec2DescribeHostReservationOfferingsResult{Result: future}
}

func (a *EC2Stub) DescribeHostReservations(ctx workflow.Context, input *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
    var output ec2.DescribeHostReservationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHostReservations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeHostReservationsAsync(ctx workflow.Context, input *ec2.DescribeHostReservationsInput) *Ec2DescribeHostReservationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHostReservations, input)
    return &Ec2DescribeHostReservationsResult{Result: future}
}

func (a *EC2Stub) DescribeHosts(ctx workflow.Context, input *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
    var output ec2.DescribeHostsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHosts, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeHostsAsync(ctx workflow.Context, input *ec2.DescribeHostsInput) *Ec2DescribeHostsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHosts, input)
    return &Ec2DescribeHostsResult{Result: future}
}

func (a *EC2Stub) DescribeIamInstanceProfileAssociations(ctx workflow.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
    var output ec2.DescribeIamInstanceProfileAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIamInstanceProfileAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeIamInstanceProfileAssociationsAsync(ctx workflow.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput) *Ec2DescribeIamInstanceProfileAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeIamInstanceProfileAssociations, input)
    return &Ec2DescribeIamInstanceProfileAssociationsResult{Result: future}
}

func (a *EC2Stub) DescribeIdFormat(ctx workflow.Context, input *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
    var output ec2.DescribeIdFormatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdFormat, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeIdFormatAsync(ctx workflow.Context, input *ec2.DescribeIdFormatInput) *Ec2DescribeIdFormatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeIdFormat, input)
    return &Ec2DescribeIdFormatResult{Result: future}
}

func (a *EC2Stub) DescribeIdentityIdFormat(ctx workflow.Context, input *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
    var output ec2.DescribeIdentityIdFormatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityIdFormat, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeIdentityIdFormatAsync(ctx workflow.Context, input *ec2.DescribeIdentityIdFormatInput) *Ec2DescribeIdentityIdFormatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityIdFormat, input)
    return &Ec2DescribeIdentityIdFormatResult{Result: future}
}

func (a *EC2Stub) DescribeImageAttribute(ctx workflow.Context, input *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
    var output ec2.DescribeImageAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImageAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeImageAttributeAsync(ctx workflow.Context, input *ec2.DescribeImageAttributeInput) *Ec2DescribeImageAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeImageAttribute, input)
    return &Ec2DescribeImageAttributeResult{Result: future}
}

func (a *EC2Stub) DescribeImages(ctx workflow.Context, input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
    var output ec2.DescribeImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImages, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeImagesAsync(ctx workflow.Context, input *ec2.DescribeImagesInput) *Ec2DescribeImagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeImages, input)
    return &Ec2DescribeImagesResult{Result: future}
}

func (a *EC2Stub) DescribeImportImageTasks(ctx workflow.Context, input *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
    var output ec2.DescribeImportImageTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImportImageTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeImportImageTasksAsync(ctx workflow.Context, input *ec2.DescribeImportImageTasksInput) *Ec2DescribeImportImageTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeImportImageTasks, input)
    return &Ec2DescribeImportImageTasksResult{Result: future}
}

func (a *EC2Stub) DescribeImportSnapshotTasks(ctx workflow.Context, input *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
    var output ec2.DescribeImportSnapshotTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImportSnapshotTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeImportSnapshotTasksAsync(ctx workflow.Context, input *ec2.DescribeImportSnapshotTasksInput) *Ec2DescribeImportSnapshotTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeImportSnapshotTasks, input)
    return &Ec2DescribeImportSnapshotTasksResult{Result: future}
}

func (a *EC2Stub) DescribeInstanceAttribute(ctx workflow.Context, input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
    var output ec2.DescribeInstanceAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInstanceAttributeAsync(ctx workflow.Context, input *ec2.DescribeInstanceAttributeInput) *Ec2DescribeInstanceAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceAttribute, input)
    return &Ec2DescribeInstanceAttributeResult{Result: future}
}

func (a *EC2Stub) DescribeInstanceCreditSpecifications(ctx workflow.Context, input *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
    var output ec2.DescribeInstanceCreditSpecificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceCreditSpecifications, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInstanceCreditSpecificationsAsync(ctx workflow.Context, input *ec2.DescribeInstanceCreditSpecificationsInput) *Ec2DescribeInstanceCreditSpecificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceCreditSpecifications, input)
    return &Ec2DescribeInstanceCreditSpecificationsResult{Result: future}
}

func (a *EC2Stub) DescribeInstanceEventNotificationAttributes(ctx workflow.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
    var output ec2.DescribeInstanceEventNotificationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceEventNotificationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInstanceEventNotificationAttributesAsync(ctx workflow.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput) *Ec2DescribeInstanceEventNotificationAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceEventNotificationAttributes, input)
    return &Ec2DescribeInstanceEventNotificationAttributesResult{Result: future}
}

func (a *EC2Stub) DescribeInstanceStatus(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
    var output ec2.DescribeInstanceStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInstanceStatusAsync(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) *Ec2DescribeInstanceStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceStatus, input)
    return &Ec2DescribeInstanceStatusResult{Result: future}
}

func (a *EC2Stub) DescribeInstanceTypeOfferings(ctx workflow.Context, input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
    var output ec2.DescribeInstanceTypeOfferingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceTypeOfferings, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInstanceTypeOfferingsAsync(ctx workflow.Context, input *ec2.DescribeInstanceTypeOfferingsInput) *Ec2DescribeInstanceTypeOfferingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceTypeOfferings, input)
    return &Ec2DescribeInstanceTypeOfferingsResult{Result: future}
}

func (a *EC2Stub) DescribeInstanceTypes(ctx workflow.Context, input *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
    var output ec2.DescribeInstanceTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInstanceTypesAsync(ctx workflow.Context, input *ec2.DescribeInstanceTypesInput) *Ec2DescribeInstanceTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceTypes, input)
    return &Ec2DescribeInstanceTypesResult{Result: future}
}

func (a *EC2Stub) DescribeInstances(ctx workflow.Context, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
    var output ec2.DescribeInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInstancesAsync(ctx workflow.Context, input *ec2.DescribeInstancesInput) *Ec2DescribeInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstances, input)
    return &Ec2DescribeInstancesResult{Result: future}
}

func (a *EC2Stub) DescribeInternetGateways(ctx workflow.Context, input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
    var output ec2.DescribeInternetGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInternetGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeInternetGatewaysAsync(ctx workflow.Context, input *ec2.DescribeInternetGatewaysInput) *Ec2DescribeInternetGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInternetGateways, input)
    return &Ec2DescribeInternetGatewaysResult{Result: future}
}

func (a *EC2Stub) DescribeIpv6Pools(ctx workflow.Context, input *ec2.DescribeIpv6PoolsInput) (*ec2.DescribeIpv6PoolsOutput, error) {
    var output ec2.DescribeIpv6PoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIpv6Pools, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeIpv6PoolsAsync(ctx workflow.Context, input *ec2.DescribeIpv6PoolsInput) *Ec2DescribeIpv6PoolsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeIpv6Pools, input)
    return &Ec2DescribeIpv6PoolsResult{Result: future}
}

func (a *EC2Stub) DescribeKeyPairs(ctx workflow.Context, input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
    var output ec2.DescribeKeyPairsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeKeyPairs, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeKeyPairsAsync(ctx workflow.Context, input *ec2.DescribeKeyPairsInput) *Ec2DescribeKeyPairsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeKeyPairs, input)
    return &Ec2DescribeKeyPairsResult{Result: future}
}

func (a *EC2Stub) DescribeLaunchTemplateVersions(ctx workflow.Context, input *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
    var output ec2.DescribeLaunchTemplateVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLaunchTemplateVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLaunchTemplateVersionsAsync(ctx workflow.Context, input *ec2.DescribeLaunchTemplateVersionsInput) *Ec2DescribeLaunchTemplateVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLaunchTemplateVersions, input)
    return &Ec2DescribeLaunchTemplateVersionsResult{Result: future}
}

func (a *EC2Stub) DescribeLaunchTemplates(ctx workflow.Context, input *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error) {
    var output ec2.DescribeLaunchTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLaunchTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLaunchTemplatesAsync(ctx workflow.Context, input *ec2.DescribeLaunchTemplatesInput) *Ec2DescribeLaunchTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLaunchTemplates, input)
    return &Ec2DescribeLaunchTemplatesResult{Result: future}
}

func (a *EC2Stub) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
    var output ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) *Ec2DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations, input)
    return &Ec2DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsResult{Result: future}
}

func (a *EC2Stub) DescribeLocalGatewayRouteTableVpcAssociations(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
    var output ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayRouteTableVpcAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLocalGatewayRouteTableVpcAssociationsAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) *Ec2DescribeLocalGatewayRouteTableVpcAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayRouteTableVpcAssociations, input)
    return &Ec2DescribeLocalGatewayRouteTableVpcAssociationsResult{Result: future}
}

func (a *EC2Stub) DescribeLocalGatewayRouteTables(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTablesInput) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
    var output ec2.DescribeLocalGatewayRouteTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayRouteTables, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLocalGatewayRouteTablesAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayRouteTablesInput) *Ec2DescribeLocalGatewayRouteTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayRouteTables, input)
    return &Ec2DescribeLocalGatewayRouteTablesResult{Result: future}
}

func (a *EC2Stub) DescribeLocalGatewayVirtualInterfaceGroups(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
    var output ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayVirtualInterfaceGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLocalGatewayVirtualInterfaceGroupsAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) *Ec2DescribeLocalGatewayVirtualInterfaceGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayVirtualInterfaceGroups, input)
    return &Ec2DescribeLocalGatewayVirtualInterfaceGroupsResult{Result: future}
}

func (a *EC2Stub) DescribeLocalGatewayVirtualInterfaces(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
    var output ec2.DescribeLocalGatewayVirtualInterfacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayVirtualInterfaces, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLocalGatewayVirtualInterfacesAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput) *Ec2DescribeLocalGatewayVirtualInterfacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGatewayVirtualInterfaces, input)
    return &Ec2DescribeLocalGatewayVirtualInterfacesResult{Result: future}
}

func (a *EC2Stub) DescribeLocalGateways(ctx workflow.Context, input *ec2.DescribeLocalGatewaysInput) (*ec2.DescribeLocalGatewaysOutput, error) {
    var output ec2.DescribeLocalGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeLocalGatewaysAsync(ctx workflow.Context, input *ec2.DescribeLocalGatewaysInput) *Ec2DescribeLocalGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLocalGateways, input)
    return &Ec2DescribeLocalGatewaysResult{Result: future}
}

func (a *EC2Stub) DescribeManagedPrefixLists(ctx workflow.Context, input *ec2.DescribeManagedPrefixListsInput) (*ec2.DescribeManagedPrefixListsOutput, error) {
    var output ec2.DescribeManagedPrefixListsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeManagedPrefixLists, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeManagedPrefixListsAsync(ctx workflow.Context, input *ec2.DescribeManagedPrefixListsInput) *Ec2DescribeManagedPrefixListsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeManagedPrefixLists, input)
    return &Ec2DescribeManagedPrefixListsResult{Result: future}
}

func (a *EC2Stub) DescribeMovingAddresses(ctx workflow.Context, input *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
    var output ec2.DescribeMovingAddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMovingAddresses, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeMovingAddressesAsync(ctx workflow.Context, input *ec2.DescribeMovingAddressesInput) *Ec2DescribeMovingAddressesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMovingAddresses, input)
    return &Ec2DescribeMovingAddressesResult{Result: future}
}

func (a *EC2Stub) DescribeNatGateways(ctx workflow.Context, input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
    var output ec2.DescribeNatGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNatGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeNatGatewaysAsync(ctx workflow.Context, input *ec2.DescribeNatGatewaysInput) *Ec2DescribeNatGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNatGateways, input)
    return &Ec2DescribeNatGatewaysResult{Result: future}
}

func (a *EC2Stub) DescribeNetworkAcls(ctx workflow.Context, input *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
    var output ec2.DescribeNetworkAclsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkAcls, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeNetworkAclsAsync(ctx workflow.Context, input *ec2.DescribeNetworkAclsInput) *Ec2DescribeNetworkAclsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkAcls, input)
    return &Ec2DescribeNetworkAclsResult{Result: future}
}

func (a *EC2Stub) DescribeNetworkInterfaceAttribute(ctx workflow.Context, input *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
    var output ec2.DescribeNetworkInterfaceAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkInterfaceAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeNetworkInterfaceAttributeAsync(ctx workflow.Context, input *ec2.DescribeNetworkInterfaceAttributeInput) *Ec2DescribeNetworkInterfaceAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkInterfaceAttribute, input)
    return &Ec2DescribeNetworkInterfaceAttributeResult{Result: future}
}

func (a *EC2Stub) DescribeNetworkInterfacePermissions(ctx workflow.Context, input *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
    var output ec2.DescribeNetworkInterfacePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkInterfacePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeNetworkInterfacePermissionsAsync(ctx workflow.Context, input *ec2.DescribeNetworkInterfacePermissionsInput) *Ec2DescribeNetworkInterfacePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkInterfacePermissions, input)
    return &Ec2DescribeNetworkInterfacePermissionsResult{Result: future}
}

func (a *EC2Stub) DescribeNetworkInterfaces(ctx workflow.Context, input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
    var output ec2.DescribeNetworkInterfacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkInterfaces, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeNetworkInterfacesAsync(ctx workflow.Context, input *ec2.DescribeNetworkInterfacesInput) *Ec2DescribeNetworkInterfacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNetworkInterfaces, input)
    return &Ec2DescribeNetworkInterfacesResult{Result: future}
}

func (a *EC2Stub) DescribePlacementGroups(ctx workflow.Context, input *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
    var output ec2.DescribePlacementGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePlacementGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribePlacementGroupsAsync(ctx workflow.Context, input *ec2.DescribePlacementGroupsInput) *Ec2DescribePlacementGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePlacementGroups, input)
    return &Ec2DescribePlacementGroupsResult{Result: future}
}

func (a *EC2Stub) DescribePrefixLists(ctx workflow.Context, input *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
    var output ec2.DescribePrefixListsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePrefixLists, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribePrefixListsAsync(ctx workflow.Context, input *ec2.DescribePrefixListsInput) *Ec2DescribePrefixListsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePrefixLists, input)
    return &Ec2DescribePrefixListsResult{Result: future}
}

func (a *EC2Stub) DescribePrincipalIdFormat(ctx workflow.Context, input *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error) {
    var output ec2.DescribePrincipalIdFormatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePrincipalIdFormat, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribePrincipalIdFormatAsync(ctx workflow.Context, input *ec2.DescribePrincipalIdFormatInput) *Ec2DescribePrincipalIdFormatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePrincipalIdFormat, input)
    return &Ec2DescribePrincipalIdFormatResult{Result: future}
}

func (a *EC2Stub) DescribePublicIpv4Pools(ctx workflow.Context, input *ec2.DescribePublicIpv4PoolsInput) (*ec2.DescribePublicIpv4PoolsOutput, error) {
    var output ec2.DescribePublicIpv4PoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePublicIpv4Pools, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribePublicIpv4PoolsAsync(ctx workflow.Context, input *ec2.DescribePublicIpv4PoolsInput) *Ec2DescribePublicIpv4PoolsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePublicIpv4Pools, input)
    return &Ec2DescribePublicIpv4PoolsResult{Result: future}
}

func (a *EC2Stub) DescribeRegions(ctx workflow.Context, input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
    var output ec2.DescribeRegionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRegions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeRegionsAsync(ctx workflow.Context, input *ec2.DescribeRegionsInput) *Ec2DescribeRegionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRegions, input)
    return &Ec2DescribeRegionsResult{Result: future}
}

func (a *EC2Stub) DescribeReservedInstances(ctx workflow.Context, input *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
    var output ec2.DescribeReservedInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeReservedInstancesAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesInput) *Ec2DescribeReservedInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstances, input)
    return &Ec2DescribeReservedInstancesResult{Result: future}
}

func (a *EC2Stub) DescribeReservedInstancesListings(ctx workflow.Context, input *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
    var output ec2.DescribeReservedInstancesListingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstancesListings, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeReservedInstancesListingsAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesListingsInput) *Ec2DescribeReservedInstancesListingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstancesListings, input)
    return &Ec2DescribeReservedInstancesListingsResult{Result: future}
}

func (a *EC2Stub) DescribeReservedInstancesModifications(ctx workflow.Context, input *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
    var output ec2.DescribeReservedInstancesModificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstancesModifications, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeReservedInstancesModificationsAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesModificationsInput) *Ec2DescribeReservedInstancesModificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstancesModifications, input)
    return &Ec2DescribeReservedInstancesModificationsResult{Result: future}
}

func (a *EC2Stub) DescribeReservedInstancesOfferings(ctx workflow.Context, input *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
    var output ec2.DescribeReservedInstancesOfferingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstancesOfferings, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeReservedInstancesOfferingsAsync(ctx workflow.Context, input *ec2.DescribeReservedInstancesOfferingsInput) *Ec2DescribeReservedInstancesOfferingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedInstancesOfferings, input)
    return &Ec2DescribeReservedInstancesOfferingsResult{Result: future}
}

func (a *EC2Stub) DescribeRouteTables(ctx workflow.Context, input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
    var output ec2.DescribeRouteTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRouteTables, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeRouteTablesAsync(ctx workflow.Context, input *ec2.DescribeRouteTablesInput) *Ec2DescribeRouteTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRouteTables, input)
    return &Ec2DescribeRouteTablesResult{Result: future}
}

func (a *EC2Stub) DescribeScheduledInstanceAvailability(ctx workflow.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
    var output ec2.DescribeScheduledInstanceAvailabilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledInstanceAvailability, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeScheduledInstanceAvailabilityAsync(ctx workflow.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput) *Ec2DescribeScheduledInstanceAvailabilityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledInstanceAvailability, input)
    return &Ec2DescribeScheduledInstanceAvailabilityResult{Result: future}
}

func (a *EC2Stub) DescribeScheduledInstances(ctx workflow.Context, input *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
    var output ec2.DescribeScheduledInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeScheduledInstancesAsync(ctx workflow.Context, input *ec2.DescribeScheduledInstancesInput) *Ec2DescribeScheduledInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScheduledInstances, input)
    return &Ec2DescribeScheduledInstancesResult{Result: future}
}

func (a *EC2Stub) DescribeSecurityGroupReferences(ctx workflow.Context, input *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
    var output ec2.DescribeSecurityGroupReferencesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityGroupReferences, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSecurityGroupReferencesAsync(ctx workflow.Context, input *ec2.DescribeSecurityGroupReferencesInput) *Ec2DescribeSecurityGroupReferencesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityGroupReferences, input)
    return &Ec2DescribeSecurityGroupReferencesResult{Result: future}
}

func (a *EC2Stub) DescribeSecurityGroups(ctx workflow.Context, input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
    var output ec2.DescribeSecurityGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSecurityGroupsAsync(ctx workflow.Context, input *ec2.DescribeSecurityGroupsInput) *Ec2DescribeSecurityGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSecurityGroups, input)
    return &Ec2DescribeSecurityGroupsResult{Result: future}
}

func (a *EC2Stub) DescribeSnapshotAttribute(ctx workflow.Context, input *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
    var output ec2.DescribeSnapshotAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSnapshotAttributeAsync(ctx workflow.Context, input *ec2.DescribeSnapshotAttributeInput) *Ec2DescribeSnapshotAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshotAttribute, input)
    return &Ec2DescribeSnapshotAttributeResult{Result: future}
}

func (a *EC2Stub) DescribeSnapshots(ctx workflow.Context, input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
    var output ec2.DescribeSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshots, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSnapshotsAsync(ctx workflow.Context, input *ec2.DescribeSnapshotsInput) *Ec2DescribeSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSnapshots, input)
    return &Ec2DescribeSnapshotsResult{Result: future}
}

func (a *EC2Stub) DescribeSpotDatafeedSubscription(ctx workflow.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
    var output ec2.DescribeSpotDatafeedSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotDatafeedSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSpotDatafeedSubscriptionAsync(ctx workflow.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput) *Ec2DescribeSpotDatafeedSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotDatafeedSubscription, input)
    return &Ec2DescribeSpotDatafeedSubscriptionResult{Result: future}
}

func (a *EC2Stub) DescribeSpotFleetInstances(ctx workflow.Context, input *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
    var output ec2.DescribeSpotFleetInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotFleetInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSpotFleetInstancesAsync(ctx workflow.Context, input *ec2.DescribeSpotFleetInstancesInput) *Ec2DescribeSpotFleetInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotFleetInstances, input)
    return &Ec2DescribeSpotFleetInstancesResult{Result: future}
}

func (a *EC2Stub) DescribeSpotFleetRequestHistory(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
    var output ec2.DescribeSpotFleetRequestHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotFleetRequestHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSpotFleetRequestHistoryAsync(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestHistoryInput) *Ec2DescribeSpotFleetRequestHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotFleetRequestHistory, input)
    return &Ec2DescribeSpotFleetRequestHistoryResult{Result: future}
}

func (a *EC2Stub) DescribeSpotFleetRequests(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
    var output ec2.DescribeSpotFleetRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotFleetRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSpotFleetRequestsAsync(ctx workflow.Context, input *ec2.DescribeSpotFleetRequestsInput) *Ec2DescribeSpotFleetRequestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotFleetRequests, input)
    return &Ec2DescribeSpotFleetRequestsResult{Result: future}
}

func (a *EC2Stub) DescribeSpotInstanceRequests(ctx workflow.Context, input *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
    var output ec2.DescribeSpotInstanceRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotInstanceRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSpotInstanceRequestsAsync(ctx workflow.Context, input *ec2.DescribeSpotInstanceRequestsInput) *Ec2DescribeSpotInstanceRequestsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotInstanceRequests, input)
    return &Ec2DescribeSpotInstanceRequestsResult{Result: future}
}

func (a *EC2Stub) DescribeSpotPriceHistory(ctx workflow.Context, input *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
    var output ec2.DescribeSpotPriceHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotPriceHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSpotPriceHistoryAsync(ctx workflow.Context, input *ec2.DescribeSpotPriceHistoryInput) *Ec2DescribeSpotPriceHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSpotPriceHistory, input)
    return &Ec2DescribeSpotPriceHistoryResult{Result: future}
}

func (a *EC2Stub) DescribeStaleSecurityGroups(ctx workflow.Context, input *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
    var output ec2.DescribeStaleSecurityGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStaleSecurityGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeStaleSecurityGroupsAsync(ctx workflow.Context, input *ec2.DescribeStaleSecurityGroupsInput) *Ec2DescribeStaleSecurityGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStaleSecurityGroups, input)
    return &Ec2DescribeStaleSecurityGroupsResult{Result: future}
}

func (a *EC2Stub) DescribeSubnets(ctx workflow.Context, input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
    var output ec2.DescribeSubnetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSubnets, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeSubnetsAsync(ctx workflow.Context, input *ec2.DescribeSubnetsInput) *Ec2DescribeSubnetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSubnets, input)
    return &Ec2DescribeSubnetsResult{Result: future}
}

func (a *EC2Stub) DescribeTags(ctx workflow.Context, input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
    var output ec2.DescribeTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTagsAsync(ctx workflow.Context, input *ec2.DescribeTagsInput) *Ec2DescribeTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input)
    return &Ec2DescribeTagsResult{Result: future}
}

func (a *EC2Stub) DescribeTrafficMirrorFilters(ctx workflow.Context, input *ec2.DescribeTrafficMirrorFiltersInput) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
    var output ec2.DescribeTrafficMirrorFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrafficMirrorFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTrafficMirrorFiltersAsync(ctx workflow.Context, input *ec2.DescribeTrafficMirrorFiltersInput) *Ec2DescribeTrafficMirrorFiltersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrafficMirrorFilters, input)
    return &Ec2DescribeTrafficMirrorFiltersResult{Result: future}
}

func (a *EC2Stub) DescribeTrafficMirrorSessions(ctx workflow.Context, input *ec2.DescribeTrafficMirrorSessionsInput) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
    var output ec2.DescribeTrafficMirrorSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrafficMirrorSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTrafficMirrorSessionsAsync(ctx workflow.Context, input *ec2.DescribeTrafficMirrorSessionsInput) *Ec2DescribeTrafficMirrorSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrafficMirrorSessions, input)
    return &Ec2DescribeTrafficMirrorSessionsResult{Result: future}
}

func (a *EC2Stub) DescribeTrafficMirrorTargets(ctx workflow.Context, input *ec2.DescribeTrafficMirrorTargetsInput) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
    var output ec2.DescribeTrafficMirrorTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrafficMirrorTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTrafficMirrorTargetsAsync(ctx workflow.Context, input *ec2.DescribeTrafficMirrorTargetsInput) *Ec2DescribeTrafficMirrorTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrafficMirrorTargets, input)
    return &Ec2DescribeTrafficMirrorTargetsResult{Result: future}
}

func (a *EC2Stub) DescribeTransitGatewayAttachments(ctx workflow.Context, input *ec2.DescribeTransitGatewayAttachmentsInput) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
    var output ec2.DescribeTransitGatewayAttachmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayAttachments, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTransitGatewayAttachmentsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayAttachmentsInput) *Ec2DescribeTransitGatewayAttachmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayAttachments, input)
    return &Ec2DescribeTransitGatewayAttachmentsResult{Result: future}
}

func (a *EC2Stub) DescribeTransitGatewayMulticastDomains(ctx workflow.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
    var output ec2.DescribeTransitGatewayMulticastDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayMulticastDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTransitGatewayMulticastDomainsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput) *Ec2DescribeTransitGatewayMulticastDomainsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayMulticastDomains, input)
    return &Ec2DescribeTransitGatewayMulticastDomainsResult{Result: future}
}

func (a *EC2Stub) DescribeTransitGatewayPeeringAttachments(ctx workflow.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
    var output ec2.DescribeTransitGatewayPeeringAttachmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayPeeringAttachments, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTransitGatewayPeeringAttachmentsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) *Ec2DescribeTransitGatewayPeeringAttachmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayPeeringAttachments, input)
    return &Ec2DescribeTransitGatewayPeeringAttachmentsResult{Result: future}
}

func (a *EC2Stub) DescribeTransitGatewayRouteTables(ctx workflow.Context, input *ec2.DescribeTransitGatewayRouteTablesInput) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
    var output ec2.DescribeTransitGatewayRouteTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayRouteTables, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTransitGatewayRouteTablesAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayRouteTablesInput) *Ec2DescribeTransitGatewayRouteTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayRouteTables, input)
    return &Ec2DescribeTransitGatewayRouteTablesResult{Result: future}
}

func (a *EC2Stub) DescribeTransitGatewayVpcAttachments(ctx workflow.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
    var output ec2.DescribeTransitGatewayVpcAttachmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayVpcAttachments, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTransitGatewayVpcAttachmentsAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput) *Ec2DescribeTransitGatewayVpcAttachmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGatewayVpcAttachments, input)
    return &Ec2DescribeTransitGatewayVpcAttachmentsResult{Result: future}
}

func (a *EC2Stub) DescribeTransitGateways(ctx workflow.Context, input *ec2.DescribeTransitGatewaysInput) (*ec2.DescribeTransitGatewaysOutput, error) {
    var output ec2.DescribeTransitGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeTransitGatewaysAsync(ctx workflow.Context, input *ec2.DescribeTransitGatewaysInput) *Ec2DescribeTransitGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTransitGateways, input)
    return &Ec2DescribeTransitGatewaysResult{Result: future}
}

func (a *EC2Stub) DescribeVolumeAttribute(ctx workflow.Context, input *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
    var output ec2.DescribeVolumeAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumeAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVolumeAttributeAsync(ctx workflow.Context, input *ec2.DescribeVolumeAttributeInput) *Ec2DescribeVolumeAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumeAttribute, input)
    return &Ec2DescribeVolumeAttributeResult{Result: future}
}

func (a *EC2Stub) DescribeVolumeStatus(ctx workflow.Context, input *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
    var output ec2.DescribeVolumeStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumeStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVolumeStatusAsync(ctx workflow.Context, input *ec2.DescribeVolumeStatusInput) *Ec2DescribeVolumeStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumeStatus, input)
    return &Ec2DescribeVolumeStatusResult{Result: future}
}

func (a *EC2Stub) DescribeVolumes(ctx workflow.Context, input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
    var output ec2.DescribeVolumesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVolumesAsync(ctx workflow.Context, input *ec2.DescribeVolumesInput) *Ec2DescribeVolumesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumes, input)
    return &Ec2DescribeVolumesResult{Result: future}
}

func (a *EC2Stub) DescribeVolumesModifications(ctx workflow.Context, input *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
    var output ec2.DescribeVolumesModificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumesModifications, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVolumesModificationsAsync(ctx workflow.Context, input *ec2.DescribeVolumesModificationsInput) *Ec2DescribeVolumesModificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumesModifications, input)
    return &Ec2DescribeVolumesModificationsResult{Result: future}
}

func (a *EC2Stub) DescribeVpcAttribute(ctx workflow.Context, input *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
    var output ec2.DescribeVpcAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcAttributeAsync(ctx workflow.Context, input *ec2.DescribeVpcAttributeInput) *Ec2DescribeVpcAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcAttribute, input)
    return &Ec2DescribeVpcAttributeResult{Result: future}
}

func (a *EC2Stub) DescribeVpcClassicLink(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
    var output ec2.DescribeVpcClassicLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcClassicLink, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcClassicLinkAsync(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkInput) *Ec2DescribeVpcClassicLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcClassicLink, input)
    return &Ec2DescribeVpcClassicLinkResult{Result: future}
}

func (a *EC2Stub) DescribeVpcClassicLinkDnsSupport(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
    var output ec2.DescribeVpcClassicLinkDnsSupportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcClassicLinkDnsSupport, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcClassicLinkDnsSupportAsync(ctx workflow.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput) *Ec2DescribeVpcClassicLinkDnsSupportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcClassicLinkDnsSupport, input)
    return &Ec2DescribeVpcClassicLinkDnsSupportResult{Result: future}
}

func (a *EC2Stub) DescribeVpcEndpointConnectionNotifications(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
    var output ec2.DescribeVpcEndpointConnectionNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointConnectionNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcEndpointConnectionNotificationsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput) *Ec2DescribeVpcEndpointConnectionNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointConnectionNotifications, input)
    return &Ec2DescribeVpcEndpointConnectionNotificationsResult{Result: future}
}

func (a *EC2Stub) DescribeVpcEndpointConnections(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
    var output ec2.DescribeVpcEndpointConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcEndpointConnectionsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointConnectionsInput) *Ec2DescribeVpcEndpointConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointConnections, input)
    return &Ec2DescribeVpcEndpointConnectionsResult{Result: future}
}

func (a *EC2Stub) DescribeVpcEndpointServiceConfigurations(ctx workflow.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
    var output ec2.DescribeVpcEndpointServiceConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointServiceConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcEndpointServiceConfigurationsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput) *Ec2DescribeVpcEndpointServiceConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointServiceConfigurations, input)
    return &Ec2DescribeVpcEndpointServiceConfigurationsResult{Result: future}
}

func (a *EC2Stub) DescribeVpcEndpointServicePermissions(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
    var output ec2.DescribeVpcEndpointServicePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointServicePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcEndpointServicePermissionsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput) *Ec2DescribeVpcEndpointServicePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointServicePermissions, input)
    return &Ec2DescribeVpcEndpointServicePermissionsResult{Result: future}
}

func (a *EC2Stub) DescribeVpcEndpointServices(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
    var output ec2.DescribeVpcEndpointServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointServices, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcEndpointServicesAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointServicesInput) *Ec2DescribeVpcEndpointServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpointServices, input)
    return &Ec2DescribeVpcEndpointServicesResult{Result: future}
}

func (a *EC2Stub) DescribeVpcEndpoints(ctx workflow.Context, input *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
    var output ec2.DescribeVpcEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcEndpointsAsync(ctx workflow.Context, input *ec2.DescribeVpcEndpointsInput) *Ec2DescribeVpcEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcEndpoints, input)
    return &Ec2DescribeVpcEndpointsResult{Result: future}
}

func (a *EC2Stub) DescribeVpcPeeringConnections(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
    var output ec2.DescribeVpcPeeringConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcPeeringConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcPeeringConnectionsAsync(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) *Ec2DescribeVpcPeeringConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcPeeringConnections, input)
    return &Ec2DescribeVpcPeeringConnectionsResult{Result: future}
}

func (a *EC2Stub) DescribeVpcs(ctx workflow.Context, input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
    var output ec2.DescribeVpcsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcs, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpcsAsync(ctx workflow.Context, input *ec2.DescribeVpcsInput) *Ec2DescribeVpcsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcs, input)
    return &Ec2DescribeVpcsResult{Result: future}
}

func (a *EC2Stub) DescribeVpnConnections(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
    var output ec2.DescribeVpnConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpnConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpnConnectionsAsync(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) *Ec2DescribeVpnConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpnConnections, input)
    return &Ec2DescribeVpnConnectionsResult{Result: future}
}

func (a *EC2Stub) DescribeVpnGateways(ctx workflow.Context, input *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
    var output ec2.DescribeVpnGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpnGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DescribeVpnGatewaysAsync(ctx workflow.Context, input *ec2.DescribeVpnGatewaysInput) *Ec2DescribeVpnGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpnGateways, input)
    return &Ec2DescribeVpnGatewaysResult{Result: future}
}

func (a *EC2Stub) DetachClassicLinkVpc(ctx workflow.Context, input *ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error) {
    var output ec2.DetachClassicLinkVpcOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachClassicLinkVpc, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DetachClassicLinkVpcAsync(ctx workflow.Context, input *ec2.DetachClassicLinkVpcInput) *Ec2DetachClassicLinkVpcResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachClassicLinkVpc, input)
    return &Ec2DetachClassicLinkVpcResult{Result: future}
}

func (a *EC2Stub) DetachInternetGateway(ctx workflow.Context, input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
    var output ec2.DetachInternetGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachInternetGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DetachInternetGatewayAsync(ctx workflow.Context, input *ec2.DetachInternetGatewayInput) *Ec2DetachInternetGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachInternetGateway, input)
    return &Ec2DetachInternetGatewayResult{Result: future}
}

func (a *EC2Stub) DetachNetworkInterface(ctx workflow.Context, input *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
    var output ec2.DetachNetworkInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachNetworkInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DetachNetworkInterfaceAsync(ctx workflow.Context, input *ec2.DetachNetworkInterfaceInput) *Ec2DetachNetworkInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachNetworkInterface, input)
    return &Ec2DetachNetworkInterfaceResult{Result: future}
}

func (a *EC2Stub) DetachVolume(ctx workflow.Context, input *ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error) {
    var output ec2.VolumeAttachment
    err := workflow.ExecuteActivity(ctx, a.activities.DetachVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DetachVolumeAsync(ctx workflow.Context, input *ec2.DetachVolumeInput) *Ec2DetachVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachVolume, input)
    return &Ec2DetachVolumeResult{Result: future}
}

func (a *EC2Stub) DetachVpnGateway(ctx workflow.Context, input *ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error) {
    var output ec2.DetachVpnGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachVpnGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DetachVpnGatewayAsync(ctx workflow.Context, input *ec2.DetachVpnGatewayInput) *Ec2DetachVpnGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachVpnGateway, input)
    return &Ec2DetachVpnGatewayResult{Result: future}
}

func (a *EC2Stub) DisableEbsEncryptionByDefault(ctx workflow.Context, input *ec2.DisableEbsEncryptionByDefaultInput) (*ec2.DisableEbsEncryptionByDefaultOutput, error) {
    var output ec2.DisableEbsEncryptionByDefaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableEbsEncryptionByDefault, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisableEbsEncryptionByDefaultAsync(ctx workflow.Context, input *ec2.DisableEbsEncryptionByDefaultInput) *Ec2DisableEbsEncryptionByDefaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableEbsEncryptionByDefault, input)
    return &Ec2DisableEbsEncryptionByDefaultResult{Result: future}
}

func (a *EC2Stub) DisableFastSnapshotRestores(ctx workflow.Context, input *ec2.DisableFastSnapshotRestoresInput) (*ec2.DisableFastSnapshotRestoresOutput, error) {
    var output ec2.DisableFastSnapshotRestoresOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableFastSnapshotRestores, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisableFastSnapshotRestoresAsync(ctx workflow.Context, input *ec2.DisableFastSnapshotRestoresInput) *Ec2DisableFastSnapshotRestoresResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableFastSnapshotRestores, input)
    return &Ec2DisableFastSnapshotRestoresResult{Result: future}
}

func (a *EC2Stub) DisableTransitGatewayRouteTablePropagation(ctx workflow.Context, input *ec2.DisableTransitGatewayRouteTablePropagationInput) (*ec2.DisableTransitGatewayRouteTablePropagationOutput, error) {
    var output ec2.DisableTransitGatewayRouteTablePropagationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableTransitGatewayRouteTablePropagation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisableTransitGatewayRouteTablePropagationAsync(ctx workflow.Context, input *ec2.DisableTransitGatewayRouteTablePropagationInput) *Ec2DisableTransitGatewayRouteTablePropagationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableTransitGatewayRouteTablePropagation, input)
    return &Ec2DisableTransitGatewayRouteTablePropagationResult{Result: future}
}

func (a *EC2Stub) DisableVgwRoutePropagation(ctx workflow.Context, input *ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error) {
    var output ec2.DisableVgwRoutePropagationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableVgwRoutePropagation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisableVgwRoutePropagationAsync(ctx workflow.Context, input *ec2.DisableVgwRoutePropagationInput) *Ec2DisableVgwRoutePropagationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableVgwRoutePropagation, input)
    return &Ec2DisableVgwRoutePropagationResult{Result: future}
}

func (a *EC2Stub) DisableVpcClassicLink(ctx workflow.Context, input *ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error) {
    var output ec2.DisableVpcClassicLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableVpcClassicLink, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisableVpcClassicLinkAsync(ctx workflow.Context, input *ec2.DisableVpcClassicLinkInput) *Ec2DisableVpcClassicLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableVpcClassicLink, input)
    return &Ec2DisableVpcClassicLinkResult{Result: future}
}

func (a *EC2Stub) DisableVpcClassicLinkDnsSupport(ctx workflow.Context, input *ec2.DisableVpcClassicLinkDnsSupportInput) (*ec2.DisableVpcClassicLinkDnsSupportOutput, error) {
    var output ec2.DisableVpcClassicLinkDnsSupportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableVpcClassicLinkDnsSupport, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisableVpcClassicLinkDnsSupportAsync(ctx workflow.Context, input *ec2.DisableVpcClassicLinkDnsSupportInput) *Ec2DisableVpcClassicLinkDnsSupportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableVpcClassicLinkDnsSupport, input)
    return &Ec2DisableVpcClassicLinkDnsSupportResult{Result: future}
}

func (a *EC2Stub) DisassociateAddress(ctx workflow.Context, input *ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error) {
    var output ec2.DisassociateAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateAddressAsync(ctx workflow.Context, input *ec2.DisassociateAddressInput) *Ec2DisassociateAddressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateAddress, input)
    return &Ec2DisassociateAddressResult{Result: future}
}

func (a *EC2Stub) DisassociateClientVpnTargetNetwork(ctx workflow.Context, input *ec2.DisassociateClientVpnTargetNetworkInput) (*ec2.DisassociateClientVpnTargetNetworkOutput, error) {
    var output ec2.DisassociateClientVpnTargetNetworkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateClientVpnTargetNetwork, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateClientVpnTargetNetworkAsync(ctx workflow.Context, input *ec2.DisassociateClientVpnTargetNetworkInput) *Ec2DisassociateClientVpnTargetNetworkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateClientVpnTargetNetwork, input)
    return &Ec2DisassociateClientVpnTargetNetworkResult{Result: future}
}

func (a *EC2Stub) DisassociateIamInstanceProfile(ctx workflow.Context, input *ec2.DisassociateIamInstanceProfileInput) (*ec2.DisassociateIamInstanceProfileOutput, error) {
    var output ec2.DisassociateIamInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateIamInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateIamInstanceProfileAsync(ctx workflow.Context, input *ec2.DisassociateIamInstanceProfileInput) *Ec2DisassociateIamInstanceProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateIamInstanceProfile, input)
    return &Ec2DisassociateIamInstanceProfileResult{Result: future}
}

func (a *EC2Stub) DisassociateRouteTable(ctx workflow.Context, input *ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error) {
    var output ec2.DisassociateRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateRouteTableAsync(ctx workflow.Context, input *ec2.DisassociateRouteTableInput) *Ec2DisassociateRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateRouteTable, input)
    return &Ec2DisassociateRouteTableResult{Result: future}
}

func (a *EC2Stub) DisassociateSubnetCidrBlock(ctx workflow.Context, input *ec2.DisassociateSubnetCidrBlockInput) (*ec2.DisassociateSubnetCidrBlockOutput, error) {
    var output ec2.DisassociateSubnetCidrBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateSubnetCidrBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateSubnetCidrBlockAsync(ctx workflow.Context, input *ec2.DisassociateSubnetCidrBlockInput) *Ec2DisassociateSubnetCidrBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateSubnetCidrBlock, input)
    return &Ec2DisassociateSubnetCidrBlockResult{Result: future}
}

func (a *EC2Stub) DisassociateTransitGatewayMulticastDomain(ctx workflow.Context, input *ec2.DisassociateTransitGatewayMulticastDomainInput) (*ec2.DisassociateTransitGatewayMulticastDomainOutput, error) {
    var output ec2.DisassociateTransitGatewayMulticastDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateTransitGatewayMulticastDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateTransitGatewayMulticastDomainAsync(ctx workflow.Context, input *ec2.DisassociateTransitGatewayMulticastDomainInput) *Ec2DisassociateTransitGatewayMulticastDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateTransitGatewayMulticastDomain, input)
    return &Ec2DisassociateTransitGatewayMulticastDomainResult{Result: future}
}

func (a *EC2Stub) DisassociateTransitGatewayRouteTable(ctx workflow.Context, input *ec2.DisassociateTransitGatewayRouteTableInput) (*ec2.DisassociateTransitGatewayRouteTableOutput, error) {
    var output ec2.DisassociateTransitGatewayRouteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateTransitGatewayRouteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateTransitGatewayRouteTableAsync(ctx workflow.Context, input *ec2.DisassociateTransitGatewayRouteTableInput) *Ec2DisassociateTransitGatewayRouteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateTransitGatewayRouteTable, input)
    return &Ec2DisassociateTransitGatewayRouteTableResult{Result: future}
}

func (a *EC2Stub) DisassociateVpcCidrBlock(ctx workflow.Context, input *ec2.DisassociateVpcCidrBlockInput) (*ec2.DisassociateVpcCidrBlockOutput, error) {
    var output ec2.DisassociateVpcCidrBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateVpcCidrBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) DisassociateVpcCidrBlockAsync(ctx workflow.Context, input *ec2.DisassociateVpcCidrBlockInput) *Ec2DisassociateVpcCidrBlockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateVpcCidrBlock, input)
    return &Ec2DisassociateVpcCidrBlockResult{Result: future}
}

func (a *EC2Stub) EnableEbsEncryptionByDefault(ctx workflow.Context, input *ec2.EnableEbsEncryptionByDefaultInput) (*ec2.EnableEbsEncryptionByDefaultOutput, error) {
    var output ec2.EnableEbsEncryptionByDefaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableEbsEncryptionByDefault, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) EnableEbsEncryptionByDefaultAsync(ctx workflow.Context, input *ec2.EnableEbsEncryptionByDefaultInput) *Ec2EnableEbsEncryptionByDefaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableEbsEncryptionByDefault, input)
    return &Ec2EnableEbsEncryptionByDefaultResult{Result: future}
}

func (a *EC2Stub) EnableFastSnapshotRestores(ctx workflow.Context, input *ec2.EnableFastSnapshotRestoresInput) (*ec2.EnableFastSnapshotRestoresOutput, error) {
    var output ec2.EnableFastSnapshotRestoresOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableFastSnapshotRestores, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) EnableFastSnapshotRestoresAsync(ctx workflow.Context, input *ec2.EnableFastSnapshotRestoresInput) *Ec2EnableFastSnapshotRestoresResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableFastSnapshotRestores, input)
    return &Ec2EnableFastSnapshotRestoresResult{Result: future}
}

func (a *EC2Stub) EnableTransitGatewayRouteTablePropagation(ctx workflow.Context, input *ec2.EnableTransitGatewayRouteTablePropagationInput) (*ec2.EnableTransitGatewayRouteTablePropagationOutput, error) {
    var output ec2.EnableTransitGatewayRouteTablePropagationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableTransitGatewayRouteTablePropagation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) EnableTransitGatewayRouteTablePropagationAsync(ctx workflow.Context, input *ec2.EnableTransitGatewayRouteTablePropagationInput) *Ec2EnableTransitGatewayRouteTablePropagationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableTransitGatewayRouteTablePropagation, input)
    return &Ec2EnableTransitGatewayRouteTablePropagationResult{Result: future}
}

func (a *EC2Stub) EnableVgwRoutePropagation(ctx workflow.Context, input *ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error) {
    var output ec2.EnableVgwRoutePropagationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableVgwRoutePropagation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) EnableVgwRoutePropagationAsync(ctx workflow.Context, input *ec2.EnableVgwRoutePropagationInput) *Ec2EnableVgwRoutePropagationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableVgwRoutePropagation, input)
    return &Ec2EnableVgwRoutePropagationResult{Result: future}
}

func (a *EC2Stub) EnableVolumeIO(ctx workflow.Context, input *ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error) {
    var output ec2.EnableVolumeIOOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableVolumeIO, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) EnableVolumeIOAsync(ctx workflow.Context, input *ec2.EnableVolumeIOInput) *Ec2EnableVolumeIOResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableVolumeIO, input)
    return &Ec2EnableVolumeIOResult{Result: future}
}

func (a *EC2Stub) EnableVpcClassicLink(ctx workflow.Context, input *ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error) {
    var output ec2.EnableVpcClassicLinkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableVpcClassicLink, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) EnableVpcClassicLinkAsync(ctx workflow.Context, input *ec2.EnableVpcClassicLinkInput) *Ec2EnableVpcClassicLinkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableVpcClassicLink, input)
    return &Ec2EnableVpcClassicLinkResult{Result: future}
}

func (a *EC2Stub) EnableVpcClassicLinkDnsSupport(ctx workflow.Context, input *ec2.EnableVpcClassicLinkDnsSupportInput) (*ec2.EnableVpcClassicLinkDnsSupportOutput, error) {
    var output ec2.EnableVpcClassicLinkDnsSupportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableVpcClassicLinkDnsSupport, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) EnableVpcClassicLinkDnsSupportAsync(ctx workflow.Context, input *ec2.EnableVpcClassicLinkDnsSupportInput) *Ec2EnableVpcClassicLinkDnsSupportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableVpcClassicLinkDnsSupport, input)
    return &Ec2EnableVpcClassicLinkDnsSupportResult{Result: future}
}

func (a *EC2Stub) ExportClientVpnClientCertificateRevocationList(ctx workflow.Context, input *ec2.ExportClientVpnClientCertificateRevocationListInput) (*ec2.ExportClientVpnClientCertificateRevocationListOutput, error) {
    var output ec2.ExportClientVpnClientCertificateRevocationListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportClientVpnClientCertificateRevocationList, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ExportClientVpnClientCertificateRevocationListAsync(ctx workflow.Context, input *ec2.ExportClientVpnClientCertificateRevocationListInput) *Ec2ExportClientVpnClientCertificateRevocationListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExportClientVpnClientCertificateRevocationList, input)
    return &Ec2ExportClientVpnClientCertificateRevocationListResult{Result: future}
}

func (a *EC2Stub) ExportClientVpnClientConfiguration(ctx workflow.Context, input *ec2.ExportClientVpnClientConfigurationInput) (*ec2.ExportClientVpnClientConfigurationOutput, error) {
    var output ec2.ExportClientVpnClientConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportClientVpnClientConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ExportClientVpnClientConfigurationAsync(ctx workflow.Context, input *ec2.ExportClientVpnClientConfigurationInput) *Ec2ExportClientVpnClientConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExportClientVpnClientConfiguration, input)
    return &Ec2ExportClientVpnClientConfigurationResult{Result: future}
}

func (a *EC2Stub) ExportImage(ctx workflow.Context, input *ec2.ExportImageInput) (*ec2.ExportImageOutput, error) {
    var output ec2.ExportImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ExportImageAsync(ctx workflow.Context, input *ec2.ExportImageInput) *Ec2ExportImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExportImage, input)
    return &Ec2ExportImageResult{Result: future}
}

func (a *EC2Stub) ExportTransitGatewayRoutes(ctx workflow.Context, input *ec2.ExportTransitGatewayRoutesInput) (*ec2.ExportTransitGatewayRoutesOutput, error) {
    var output ec2.ExportTransitGatewayRoutesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportTransitGatewayRoutes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ExportTransitGatewayRoutesAsync(ctx workflow.Context, input *ec2.ExportTransitGatewayRoutesInput) *Ec2ExportTransitGatewayRoutesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExportTransitGatewayRoutes, input)
    return &Ec2ExportTransitGatewayRoutesResult{Result: future}
}

func (a *EC2Stub) GetAssociatedIpv6PoolCidrs(ctx workflow.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
    var output ec2.GetAssociatedIpv6PoolCidrsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAssociatedIpv6PoolCidrs, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetAssociatedIpv6PoolCidrsAsync(ctx workflow.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput) *Ec2GetAssociatedIpv6PoolCidrsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAssociatedIpv6PoolCidrs, input)
    return &Ec2GetAssociatedIpv6PoolCidrsResult{Result: future}
}

func (a *EC2Stub) GetCapacityReservationUsage(ctx workflow.Context, input *ec2.GetCapacityReservationUsageInput) (*ec2.GetCapacityReservationUsageOutput, error) {
    var output ec2.GetCapacityReservationUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCapacityReservationUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetCapacityReservationUsageAsync(ctx workflow.Context, input *ec2.GetCapacityReservationUsageInput) *Ec2GetCapacityReservationUsageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCapacityReservationUsage, input)
    return &Ec2GetCapacityReservationUsageResult{Result: future}
}

func (a *EC2Stub) GetCoipPoolUsage(ctx workflow.Context, input *ec2.GetCoipPoolUsageInput) (*ec2.GetCoipPoolUsageOutput, error) {
    var output ec2.GetCoipPoolUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCoipPoolUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetCoipPoolUsageAsync(ctx workflow.Context, input *ec2.GetCoipPoolUsageInput) *Ec2GetCoipPoolUsageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCoipPoolUsage, input)
    return &Ec2GetCoipPoolUsageResult{Result: future}
}

func (a *EC2Stub) GetConsoleOutput(ctx workflow.Context, input *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
    var output ec2.GetConsoleOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConsoleOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetConsoleOutputAsync(ctx workflow.Context, input *ec2.GetConsoleOutputInput) *Ec2GetConsoleOutputResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConsoleOutput, input)
    return &Ec2GetConsoleOutputResult{Result: future}
}

func (a *EC2Stub) GetConsoleScreenshot(ctx workflow.Context, input *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
    var output ec2.GetConsoleScreenshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConsoleScreenshot, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetConsoleScreenshotAsync(ctx workflow.Context, input *ec2.GetConsoleScreenshotInput) *Ec2GetConsoleScreenshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConsoleScreenshot, input)
    return &Ec2GetConsoleScreenshotResult{Result: future}
}

func (a *EC2Stub) GetDefaultCreditSpecification(ctx workflow.Context, input *ec2.GetDefaultCreditSpecificationInput) (*ec2.GetDefaultCreditSpecificationOutput, error) {
    var output ec2.GetDefaultCreditSpecificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDefaultCreditSpecification, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetDefaultCreditSpecificationAsync(ctx workflow.Context, input *ec2.GetDefaultCreditSpecificationInput) *Ec2GetDefaultCreditSpecificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDefaultCreditSpecification, input)
    return &Ec2GetDefaultCreditSpecificationResult{Result: future}
}

func (a *EC2Stub) GetEbsDefaultKmsKeyId(ctx workflow.Context, input *ec2.GetEbsDefaultKmsKeyIdInput) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
    var output ec2.GetEbsDefaultKmsKeyIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEbsDefaultKmsKeyId, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetEbsDefaultKmsKeyIdAsync(ctx workflow.Context, input *ec2.GetEbsDefaultKmsKeyIdInput) *Ec2GetEbsDefaultKmsKeyIdResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEbsDefaultKmsKeyId, input)
    return &Ec2GetEbsDefaultKmsKeyIdResult{Result: future}
}

func (a *EC2Stub) GetEbsEncryptionByDefault(ctx workflow.Context, input *ec2.GetEbsEncryptionByDefaultInput) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
    var output ec2.GetEbsEncryptionByDefaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEbsEncryptionByDefault, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetEbsEncryptionByDefaultAsync(ctx workflow.Context, input *ec2.GetEbsEncryptionByDefaultInput) *Ec2GetEbsEncryptionByDefaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEbsEncryptionByDefault, input)
    return &Ec2GetEbsEncryptionByDefaultResult{Result: future}
}

func (a *EC2Stub) GetGroupsForCapacityReservation(ctx workflow.Context, input *ec2.GetGroupsForCapacityReservationInput) (*ec2.GetGroupsForCapacityReservationOutput, error) {
    var output ec2.GetGroupsForCapacityReservationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroupsForCapacityReservation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetGroupsForCapacityReservationAsync(ctx workflow.Context, input *ec2.GetGroupsForCapacityReservationInput) *Ec2GetGroupsForCapacityReservationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGroupsForCapacityReservation, input)
    return &Ec2GetGroupsForCapacityReservationResult{Result: future}
}

func (a *EC2Stub) GetHostReservationPurchasePreview(ctx workflow.Context, input *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
    var output ec2.GetHostReservationPurchasePreviewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHostReservationPurchasePreview, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetHostReservationPurchasePreviewAsync(ctx workflow.Context, input *ec2.GetHostReservationPurchasePreviewInput) *Ec2GetHostReservationPurchasePreviewResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHostReservationPurchasePreview, input)
    return &Ec2GetHostReservationPurchasePreviewResult{Result: future}
}

func (a *EC2Stub) GetLaunchTemplateData(ctx workflow.Context, input *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error) {
    var output ec2.GetLaunchTemplateDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLaunchTemplateData, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetLaunchTemplateDataAsync(ctx workflow.Context, input *ec2.GetLaunchTemplateDataInput) *Ec2GetLaunchTemplateDataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLaunchTemplateData, input)
    return &Ec2GetLaunchTemplateDataResult{Result: future}
}

func (a *EC2Stub) GetManagedPrefixListAssociations(ctx workflow.Context, input *ec2.GetManagedPrefixListAssociationsInput) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
    var output ec2.GetManagedPrefixListAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetManagedPrefixListAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetManagedPrefixListAssociationsAsync(ctx workflow.Context, input *ec2.GetManagedPrefixListAssociationsInput) *Ec2GetManagedPrefixListAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetManagedPrefixListAssociations, input)
    return &Ec2GetManagedPrefixListAssociationsResult{Result: future}
}

func (a *EC2Stub) GetManagedPrefixListEntries(ctx workflow.Context, input *ec2.GetManagedPrefixListEntriesInput) (*ec2.GetManagedPrefixListEntriesOutput, error) {
    var output ec2.GetManagedPrefixListEntriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetManagedPrefixListEntries, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetManagedPrefixListEntriesAsync(ctx workflow.Context, input *ec2.GetManagedPrefixListEntriesInput) *Ec2GetManagedPrefixListEntriesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetManagedPrefixListEntries, input)
    return &Ec2GetManagedPrefixListEntriesResult{Result: future}
}

func (a *EC2Stub) GetPasswordData(ctx workflow.Context, input *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
    var output ec2.GetPasswordDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPasswordData, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetPasswordDataAsync(ctx workflow.Context, input *ec2.GetPasswordDataInput) *Ec2GetPasswordDataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPasswordData, input)
    return &Ec2GetPasswordDataResult{Result: future}
}

func (a *EC2Stub) GetReservedInstancesExchangeQuote(ctx workflow.Context, input *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
    var output ec2.GetReservedInstancesExchangeQuoteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetReservedInstancesExchangeQuote, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetReservedInstancesExchangeQuoteAsync(ctx workflow.Context, input *ec2.GetReservedInstancesExchangeQuoteInput) *Ec2GetReservedInstancesExchangeQuoteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetReservedInstancesExchangeQuote, input)
    return &Ec2GetReservedInstancesExchangeQuoteResult{Result: future}
}

func (a *EC2Stub) GetTransitGatewayAttachmentPropagations(ctx workflow.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
    var output ec2.GetTransitGatewayAttachmentPropagationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayAttachmentPropagations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetTransitGatewayAttachmentPropagationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput) *Ec2GetTransitGatewayAttachmentPropagationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayAttachmentPropagations, input)
    return &Ec2GetTransitGatewayAttachmentPropagationsResult{Result: future}
}

func (a *EC2Stub) GetTransitGatewayMulticastDomainAssociations(ctx workflow.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
    var output ec2.GetTransitGatewayMulticastDomainAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayMulticastDomainAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetTransitGatewayMulticastDomainAssociationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) *Ec2GetTransitGatewayMulticastDomainAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayMulticastDomainAssociations, input)
    return &Ec2GetTransitGatewayMulticastDomainAssociationsResult{Result: future}
}

func (a *EC2Stub) GetTransitGatewayPrefixListReferences(ctx workflow.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
    var output ec2.GetTransitGatewayPrefixListReferencesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayPrefixListReferences, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetTransitGatewayPrefixListReferencesAsync(ctx workflow.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput) *Ec2GetTransitGatewayPrefixListReferencesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayPrefixListReferences, input)
    return &Ec2GetTransitGatewayPrefixListReferencesResult{Result: future}
}

func (a *EC2Stub) GetTransitGatewayRouteTableAssociations(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
    var output ec2.GetTransitGatewayRouteTableAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayRouteTableAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetTransitGatewayRouteTableAssociationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput) *Ec2GetTransitGatewayRouteTableAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayRouteTableAssociations, input)
    return &Ec2GetTransitGatewayRouteTableAssociationsResult{Result: future}
}

func (a *EC2Stub) GetTransitGatewayRouteTablePropagations(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
    var output ec2.GetTransitGatewayRouteTablePropagationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayRouteTablePropagations, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) GetTransitGatewayRouteTablePropagationsAsync(ctx workflow.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput) *Ec2GetTransitGatewayRouteTablePropagationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTransitGatewayRouteTablePropagations, input)
    return &Ec2GetTransitGatewayRouteTablePropagationsResult{Result: future}
}

func (a *EC2Stub) ImportClientVpnClientCertificateRevocationList(ctx workflow.Context, input *ec2.ImportClientVpnClientCertificateRevocationListInput) (*ec2.ImportClientVpnClientCertificateRevocationListOutput, error) {
    var output ec2.ImportClientVpnClientCertificateRevocationListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportClientVpnClientCertificateRevocationList, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ImportClientVpnClientCertificateRevocationListAsync(ctx workflow.Context, input *ec2.ImportClientVpnClientCertificateRevocationListInput) *Ec2ImportClientVpnClientCertificateRevocationListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportClientVpnClientCertificateRevocationList, input)
    return &Ec2ImportClientVpnClientCertificateRevocationListResult{Result: future}
}

func (a *EC2Stub) ImportImage(ctx workflow.Context, input *ec2.ImportImageInput) (*ec2.ImportImageOutput, error) {
    var output ec2.ImportImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ImportImageAsync(ctx workflow.Context, input *ec2.ImportImageInput) *Ec2ImportImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportImage, input)
    return &Ec2ImportImageResult{Result: future}
}

func (a *EC2Stub) ImportInstance(ctx workflow.Context, input *ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error) {
    var output ec2.ImportInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ImportInstanceAsync(ctx workflow.Context, input *ec2.ImportInstanceInput) *Ec2ImportInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportInstance, input)
    return &Ec2ImportInstanceResult{Result: future}
}

func (a *EC2Stub) ImportKeyPair(ctx workflow.Context, input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
    var output ec2.ImportKeyPairOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportKeyPair, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ImportKeyPairAsync(ctx workflow.Context, input *ec2.ImportKeyPairInput) *Ec2ImportKeyPairResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportKeyPair, input)
    return &Ec2ImportKeyPairResult{Result: future}
}

func (a *EC2Stub) ImportSnapshot(ctx workflow.Context, input *ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error) {
    var output ec2.ImportSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ImportSnapshotAsync(ctx workflow.Context, input *ec2.ImportSnapshotInput) *Ec2ImportSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportSnapshot, input)
    return &Ec2ImportSnapshotResult{Result: future}
}

func (a *EC2Stub) ImportVolume(ctx workflow.Context, input *ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error) {
    var output ec2.ImportVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ImportVolumeAsync(ctx workflow.Context, input *ec2.ImportVolumeInput) *Ec2ImportVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportVolume, input)
    return &Ec2ImportVolumeResult{Result: future}
}

func (a *EC2Stub) ModifyAvailabilityZoneGroup(ctx workflow.Context, input *ec2.ModifyAvailabilityZoneGroupInput) (*ec2.ModifyAvailabilityZoneGroupOutput, error) {
    var output ec2.ModifyAvailabilityZoneGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyAvailabilityZoneGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyAvailabilityZoneGroupAsync(ctx workflow.Context, input *ec2.ModifyAvailabilityZoneGroupInput) *Ec2ModifyAvailabilityZoneGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyAvailabilityZoneGroup, input)
    return &Ec2ModifyAvailabilityZoneGroupResult{Result: future}
}

func (a *EC2Stub) ModifyCapacityReservation(ctx workflow.Context, input *ec2.ModifyCapacityReservationInput) (*ec2.ModifyCapacityReservationOutput, error) {
    var output ec2.ModifyCapacityReservationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyCapacityReservation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyCapacityReservationAsync(ctx workflow.Context, input *ec2.ModifyCapacityReservationInput) *Ec2ModifyCapacityReservationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyCapacityReservation, input)
    return &Ec2ModifyCapacityReservationResult{Result: future}
}

func (a *EC2Stub) ModifyClientVpnEndpoint(ctx workflow.Context, input *ec2.ModifyClientVpnEndpointInput) (*ec2.ModifyClientVpnEndpointOutput, error) {
    var output ec2.ModifyClientVpnEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyClientVpnEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyClientVpnEndpointAsync(ctx workflow.Context, input *ec2.ModifyClientVpnEndpointInput) *Ec2ModifyClientVpnEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyClientVpnEndpoint, input)
    return &Ec2ModifyClientVpnEndpointResult{Result: future}
}

func (a *EC2Stub) ModifyDefaultCreditSpecification(ctx workflow.Context, input *ec2.ModifyDefaultCreditSpecificationInput) (*ec2.ModifyDefaultCreditSpecificationOutput, error) {
    var output ec2.ModifyDefaultCreditSpecificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDefaultCreditSpecification, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyDefaultCreditSpecificationAsync(ctx workflow.Context, input *ec2.ModifyDefaultCreditSpecificationInput) *Ec2ModifyDefaultCreditSpecificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyDefaultCreditSpecification, input)
    return &Ec2ModifyDefaultCreditSpecificationResult{Result: future}
}

func (a *EC2Stub) ModifyEbsDefaultKmsKeyId(ctx workflow.Context, input *ec2.ModifyEbsDefaultKmsKeyIdInput) (*ec2.ModifyEbsDefaultKmsKeyIdOutput, error) {
    var output ec2.ModifyEbsDefaultKmsKeyIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyEbsDefaultKmsKeyId, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyEbsDefaultKmsKeyIdAsync(ctx workflow.Context, input *ec2.ModifyEbsDefaultKmsKeyIdInput) *Ec2ModifyEbsDefaultKmsKeyIdResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyEbsDefaultKmsKeyId, input)
    return &Ec2ModifyEbsDefaultKmsKeyIdResult{Result: future}
}

func (a *EC2Stub) ModifyFleet(ctx workflow.Context, input *ec2.ModifyFleetInput) (*ec2.ModifyFleetOutput, error) {
    var output ec2.ModifyFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyFleetAsync(ctx workflow.Context, input *ec2.ModifyFleetInput) *Ec2ModifyFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyFleet, input)
    return &Ec2ModifyFleetResult{Result: future}
}

func (a *EC2Stub) ModifyFpgaImageAttribute(ctx workflow.Context, input *ec2.ModifyFpgaImageAttributeInput) (*ec2.ModifyFpgaImageAttributeOutput, error) {
    var output ec2.ModifyFpgaImageAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyFpgaImageAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyFpgaImageAttributeAsync(ctx workflow.Context, input *ec2.ModifyFpgaImageAttributeInput) *Ec2ModifyFpgaImageAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyFpgaImageAttribute, input)
    return &Ec2ModifyFpgaImageAttributeResult{Result: future}
}

func (a *EC2Stub) ModifyHosts(ctx workflow.Context, input *ec2.ModifyHostsInput) (*ec2.ModifyHostsOutput, error) {
    var output ec2.ModifyHostsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyHosts, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyHostsAsync(ctx workflow.Context, input *ec2.ModifyHostsInput) *Ec2ModifyHostsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyHosts, input)
    return &Ec2ModifyHostsResult{Result: future}
}

func (a *EC2Stub) ModifyIdFormat(ctx workflow.Context, input *ec2.ModifyIdFormatInput) (*ec2.ModifyIdFormatOutput, error) {
    var output ec2.ModifyIdFormatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyIdFormat, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyIdFormatAsync(ctx workflow.Context, input *ec2.ModifyIdFormatInput) *Ec2ModifyIdFormatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyIdFormat, input)
    return &Ec2ModifyIdFormatResult{Result: future}
}

func (a *EC2Stub) ModifyIdentityIdFormat(ctx workflow.Context, input *ec2.ModifyIdentityIdFormatInput) (*ec2.ModifyIdentityIdFormatOutput, error) {
    var output ec2.ModifyIdentityIdFormatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyIdentityIdFormat, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyIdentityIdFormatAsync(ctx workflow.Context, input *ec2.ModifyIdentityIdFormatInput) *Ec2ModifyIdentityIdFormatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyIdentityIdFormat, input)
    return &Ec2ModifyIdentityIdFormatResult{Result: future}
}

func (a *EC2Stub) ModifyImageAttribute(ctx workflow.Context, input *ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error) {
    var output ec2.ModifyImageAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyImageAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyImageAttributeAsync(ctx workflow.Context, input *ec2.ModifyImageAttributeInput) *Ec2ModifyImageAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyImageAttribute, input)
    return &Ec2ModifyImageAttributeResult{Result: future}
}

func (a *EC2Stub) ModifyInstanceAttribute(ctx workflow.Context, input *ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error) {
    var output ec2.ModifyInstanceAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyInstanceAttributeAsync(ctx workflow.Context, input *ec2.ModifyInstanceAttributeInput) *Ec2ModifyInstanceAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceAttribute, input)
    return &Ec2ModifyInstanceAttributeResult{Result: future}
}

func (a *EC2Stub) ModifyInstanceCapacityReservationAttributes(ctx workflow.Context, input *ec2.ModifyInstanceCapacityReservationAttributesInput) (*ec2.ModifyInstanceCapacityReservationAttributesOutput, error) {
    var output ec2.ModifyInstanceCapacityReservationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceCapacityReservationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyInstanceCapacityReservationAttributesAsync(ctx workflow.Context, input *ec2.ModifyInstanceCapacityReservationAttributesInput) *Ec2ModifyInstanceCapacityReservationAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceCapacityReservationAttributes, input)
    return &Ec2ModifyInstanceCapacityReservationAttributesResult{Result: future}
}

func (a *EC2Stub) ModifyInstanceCreditSpecification(ctx workflow.Context, input *ec2.ModifyInstanceCreditSpecificationInput) (*ec2.ModifyInstanceCreditSpecificationOutput, error) {
    var output ec2.ModifyInstanceCreditSpecificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceCreditSpecification, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyInstanceCreditSpecificationAsync(ctx workflow.Context, input *ec2.ModifyInstanceCreditSpecificationInput) *Ec2ModifyInstanceCreditSpecificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceCreditSpecification, input)
    return &Ec2ModifyInstanceCreditSpecificationResult{Result: future}
}

func (a *EC2Stub) ModifyInstanceEventStartTime(ctx workflow.Context, input *ec2.ModifyInstanceEventStartTimeInput) (*ec2.ModifyInstanceEventStartTimeOutput, error) {
    var output ec2.ModifyInstanceEventStartTimeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceEventStartTime, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyInstanceEventStartTimeAsync(ctx workflow.Context, input *ec2.ModifyInstanceEventStartTimeInput) *Ec2ModifyInstanceEventStartTimeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceEventStartTime, input)
    return &Ec2ModifyInstanceEventStartTimeResult{Result: future}
}

func (a *EC2Stub) ModifyInstanceMetadataOptions(ctx workflow.Context, input *ec2.ModifyInstanceMetadataOptionsInput) (*ec2.ModifyInstanceMetadataOptionsOutput, error) {
    var output ec2.ModifyInstanceMetadataOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceMetadataOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyInstanceMetadataOptionsAsync(ctx workflow.Context, input *ec2.ModifyInstanceMetadataOptionsInput) *Ec2ModifyInstanceMetadataOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyInstanceMetadataOptions, input)
    return &Ec2ModifyInstanceMetadataOptionsResult{Result: future}
}

func (a *EC2Stub) ModifyInstancePlacement(ctx workflow.Context, input *ec2.ModifyInstancePlacementInput) (*ec2.ModifyInstancePlacementOutput, error) {
    var output ec2.ModifyInstancePlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyInstancePlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyInstancePlacementAsync(ctx workflow.Context, input *ec2.ModifyInstancePlacementInput) *Ec2ModifyInstancePlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyInstancePlacement, input)
    return &Ec2ModifyInstancePlacementResult{Result: future}
}

func (a *EC2Stub) ModifyLaunchTemplate(ctx workflow.Context, input *ec2.ModifyLaunchTemplateInput) (*ec2.ModifyLaunchTemplateOutput, error) {
    var output ec2.ModifyLaunchTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyLaunchTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyLaunchTemplateAsync(ctx workflow.Context, input *ec2.ModifyLaunchTemplateInput) *Ec2ModifyLaunchTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyLaunchTemplate, input)
    return &Ec2ModifyLaunchTemplateResult{Result: future}
}

func (a *EC2Stub) ModifyManagedPrefixList(ctx workflow.Context, input *ec2.ModifyManagedPrefixListInput) (*ec2.ModifyManagedPrefixListOutput, error) {
    var output ec2.ModifyManagedPrefixListOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyManagedPrefixList, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyManagedPrefixListAsync(ctx workflow.Context, input *ec2.ModifyManagedPrefixListInput) *Ec2ModifyManagedPrefixListResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyManagedPrefixList, input)
    return &Ec2ModifyManagedPrefixListResult{Result: future}
}

func (a *EC2Stub) ModifyNetworkInterfaceAttribute(ctx workflow.Context, input *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
    var output ec2.ModifyNetworkInterfaceAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyNetworkInterfaceAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyNetworkInterfaceAttributeAsync(ctx workflow.Context, input *ec2.ModifyNetworkInterfaceAttributeInput) *Ec2ModifyNetworkInterfaceAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyNetworkInterfaceAttribute, input)
    return &Ec2ModifyNetworkInterfaceAttributeResult{Result: future}
}

func (a *EC2Stub) ModifyReservedInstances(ctx workflow.Context, input *ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error) {
    var output ec2.ModifyReservedInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyReservedInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyReservedInstancesAsync(ctx workflow.Context, input *ec2.ModifyReservedInstancesInput) *Ec2ModifyReservedInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyReservedInstances, input)
    return &Ec2ModifyReservedInstancesResult{Result: future}
}

func (a *EC2Stub) ModifySnapshotAttribute(ctx workflow.Context, input *ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error) {
    var output ec2.ModifySnapshotAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifySnapshotAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifySnapshotAttributeAsync(ctx workflow.Context, input *ec2.ModifySnapshotAttributeInput) *Ec2ModifySnapshotAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifySnapshotAttribute, input)
    return &Ec2ModifySnapshotAttributeResult{Result: future}
}

func (a *EC2Stub) ModifySubnetAttribute(ctx workflow.Context, input *ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error) {
    var output ec2.ModifySubnetAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifySubnetAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifySubnetAttributeAsync(ctx workflow.Context, input *ec2.ModifySubnetAttributeInput) *Ec2ModifySubnetAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifySubnetAttribute, input)
    return &Ec2ModifySubnetAttributeResult{Result: future}
}

func (a *EC2Stub) ModifyTrafficMirrorFilterNetworkServices(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) (*ec2.ModifyTrafficMirrorFilterNetworkServicesOutput, error) {
    var output ec2.ModifyTrafficMirrorFilterNetworkServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyTrafficMirrorFilterNetworkServices, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyTrafficMirrorFilterNetworkServicesAsync(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterNetworkServicesInput) *Ec2ModifyTrafficMirrorFilterNetworkServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyTrafficMirrorFilterNetworkServices, input)
    return &Ec2ModifyTrafficMirrorFilterNetworkServicesResult{Result: future}
}

func (a *EC2Stub) ModifyTrafficMirrorFilterRule(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterRuleInput) (*ec2.ModifyTrafficMirrorFilterRuleOutput, error) {
    var output ec2.ModifyTrafficMirrorFilterRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyTrafficMirrorFilterRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyTrafficMirrorFilterRuleAsync(ctx workflow.Context, input *ec2.ModifyTrafficMirrorFilterRuleInput) *Ec2ModifyTrafficMirrorFilterRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyTrafficMirrorFilterRule, input)
    return &Ec2ModifyTrafficMirrorFilterRuleResult{Result: future}
}

func (a *EC2Stub) ModifyTrafficMirrorSession(ctx workflow.Context, input *ec2.ModifyTrafficMirrorSessionInput) (*ec2.ModifyTrafficMirrorSessionOutput, error) {
    var output ec2.ModifyTrafficMirrorSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyTrafficMirrorSession, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyTrafficMirrorSessionAsync(ctx workflow.Context, input *ec2.ModifyTrafficMirrorSessionInput) *Ec2ModifyTrafficMirrorSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyTrafficMirrorSession, input)
    return &Ec2ModifyTrafficMirrorSessionResult{Result: future}
}

func (a *EC2Stub) ModifyTransitGateway(ctx workflow.Context, input *ec2.ModifyTransitGatewayInput) (*ec2.ModifyTransitGatewayOutput, error) {
    var output ec2.ModifyTransitGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyTransitGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyTransitGatewayAsync(ctx workflow.Context, input *ec2.ModifyTransitGatewayInput) *Ec2ModifyTransitGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyTransitGateway, input)
    return &Ec2ModifyTransitGatewayResult{Result: future}
}

func (a *EC2Stub) ModifyTransitGatewayPrefixListReference(ctx workflow.Context, input *ec2.ModifyTransitGatewayPrefixListReferenceInput) (*ec2.ModifyTransitGatewayPrefixListReferenceOutput, error) {
    var output ec2.ModifyTransitGatewayPrefixListReferenceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyTransitGatewayPrefixListReference, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyTransitGatewayPrefixListReferenceAsync(ctx workflow.Context, input *ec2.ModifyTransitGatewayPrefixListReferenceInput) *Ec2ModifyTransitGatewayPrefixListReferenceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyTransitGatewayPrefixListReference, input)
    return &Ec2ModifyTransitGatewayPrefixListReferenceResult{Result: future}
}

func (a *EC2Stub) ModifyTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.ModifyTransitGatewayVpcAttachmentInput) (*ec2.ModifyTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.ModifyTransitGatewayVpcAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyTransitGatewayVpcAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.ModifyTransitGatewayVpcAttachmentInput) *Ec2ModifyTransitGatewayVpcAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyTransitGatewayVpcAttachment, input)
    return &Ec2ModifyTransitGatewayVpcAttachmentResult{Result: future}
}

func (a *EC2Stub) ModifyVolume(ctx workflow.Context, input *ec2.ModifyVolumeInput) (*ec2.ModifyVolumeOutput, error) {
    var output ec2.ModifyVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVolumeAsync(ctx workflow.Context, input *ec2.ModifyVolumeInput) *Ec2ModifyVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVolume, input)
    return &Ec2ModifyVolumeResult{Result: future}
}

func (a *EC2Stub) ModifyVolumeAttribute(ctx workflow.Context, input *ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error) {
    var output ec2.ModifyVolumeAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVolumeAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVolumeAttributeAsync(ctx workflow.Context, input *ec2.ModifyVolumeAttributeInput) *Ec2ModifyVolumeAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVolumeAttribute, input)
    return &Ec2ModifyVolumeAttributeResult{Result: future}
}

func (a *EC2Stub) ModifyVpcAttribute(ctx workflow.Context, input *ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error) {
    var output ec2.ModifyVpcAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpcAttributeAsync(ctx workflow.Context, input *ec2.ModifyVpcAttributeInput) *Ec2ModifyVpcAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcAttribute, input)
    return &Ec2ModifyVpcAttributeResult{Result: future}
}

func (a *EC2Stub) ModifyVpcEndpoint(ctx workflow.Context, input *ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error) {
    var output ec2.ModifyVpcEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpcEndpointAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointInput) *Ec2ModifyVpcEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpoint, input)
    return &Ec2ModifyVpcEndpointResult{Result: future}
}

func (a *EC2Stub) ModifyVpcEndpointConnectionNotification(ctx workflow.Context, input *ec2.ModifyVpcEndpointConnectionNotificationInput) (*ec2.ModifyVpcEndpointConnectionNotificationOutput, error) {
    var output ec2.ModifyVpcEndpointConnectionNotificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpointConnectionNotification, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpcEndpointConnectionNotificationAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointConnectionNotificationInput) *Ec2ModifyVpcEndpointConnectionNotificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpointConnectionNotification, input)
    return &Ec2ModifyVpcEndpointConnectionNotificationResult{Result: future}
}

func (a *EC2Stub) ModifyVpcEndpointServiceConfiguration(ctx workflow.Context, input *ec2.ModifyVpcEndpointServiceConfigurationInput) (*ec2.ModifyVpcEndpointServiceConfigurationOutput, error) {
    var output ec2.ModifyVpcEndpointServiceConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpointServiceConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpcEndpointServiceConfigurationAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointServiceConfigurationInput) *Ec2ModifyVpcEndpointServiceConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpointServiceConfiguration, input)
    return &Ec2ModifyVpcEndpointServiceConfigurationResult{Result: future}
}

func (a *EC2Stub) ModifyVpcEndpointServicePermissions(ctx workflow.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error) {
    var output ec2.ModifyVpcEndpointServicePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpointServicePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpcEndpointServicePermissionsAsync(ctx workflow.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput) *Ec2ModifyVpcEndpointServicePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcEndpointServicePermissions, input)
    return &Ec2ModifyVpcEndpointServicePermissionsResult{Result: future}
}

func (a *EC2Stub) ModifyVpcPeeringConnectionOptions(ctx workflow.Context, input *ec2.ModifyVpcPeeringConnectionOptionsInput) (*ec2.ModifyVpcPeeringConnectionOptionsOutput, error) {
    var output ec2.ModifyVpcPeeringConnectionOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcPeeringConnectionOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpcPeeringConnectionOptionsAsync(ctx workflow.Context, input *ec2.ModifyVpcPeeringConnectionOptionsInput) *Ec2ModifyVpcPeeringConnectionOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcPeeringConnectionOptions, input)
    return &Ec2ModifyVpcPeeringConnectionOptionsResult{Result: future}
}

func (a *EC2Stub) ModifyVpcTenancy(ctx workflow.Context, input *ec2.ModifyVpcTenancyInput) (*ec2.ModifyVpcTenancyOutput, error) {
    var output ec2.ModifyVpcTenancyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcTenancy, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpcTenancyAsync(ctx workflow.Context, input *ec2.ModifyVpcTenancyInput) *Ec2ModifyVpcTenancyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpcTenancy, input)
    return &Ec2ModifyVpcTenancyResult{Result: future}
}

func (a *EC2Stub) ModifyVpnConnection(ctx workflow.Context, input *ec2.ModifyVpnConnectionInput) (*ec2.ModifyVpnConnectionOutput, error) {
    var output ec2.ModifyVpnConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpnConnectionAsync(ctx workflow.Context, input *ec2.ModifyVpnConnectionInput) *Ec2ModifyVpnConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnConnection, input)
    return &Ec2ModifyVpnConnectionResult{Result: future}
}

func (a *EC2Stub) ModifyVpnConnectionOptions(ctx workflow.Context, input *ec2.ModifyVpnConnectionOptionsInput) (*ec2.ModifyVpnConnectionOptionsOutput, error) {
    var output ec2.ModifyVpnConnectionOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnConnectionOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpnConnectionOptionsAsync(ctx workflow.Context, input *ec2.ModifyVpnConnectionOptionsInput) *Ec2ModifyVpnConnectionOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnConnectionOptions, input)
    return &Ec2ModifyVpnConnectionOptionsResult{Result: future}
}

func (a *EC2Stub) ModifyVpnTunnelCertificate(ctx workflow.Context, input *ec2.ModifyVpnTunnelCertificateInput) (*ec2.ModifyVpnTunnelCertificateOutput, error) {
    var output ec2.ModifyVpnTunnelCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnTunnelCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpnTunnelCertificateAsync(ctx workflow.Context, input *ec2.ModifyVpnTunnelCertificateInput) *Ec2ModifyVpnTunnelCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnTunnelCertificate, input)
    return &Ec2ModifyVpnTunnelCertificateResult{Result: future}
}

func (a *EC2Stub) ModifyVpnTunnelOptions(ctx workflow.Context, input *ec2.ModifyVpnTunnelOptionsInput) (*ec2.ModifyVpnTunnelOptionsOutput, error) {
    var output ec2.ModifyVpnTunnelOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnTunnelOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ModifyVpnTunnelOptionsAsync(ctx workflow.Context, input *ec2.ModifyVpnTunnelOptionsInput) *Ec2ModifyVpnTunnelOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyVpnTunnelOptions, input)
    return &Ec2ModifyVpnTunnelOptionsResult{Result: future}
}

func (a *EC2Stub) MonitorInstances(ctx workflow.Context, input *ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error) {
    var output ec2.MonitorInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.MonitorInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) MonitorInstancesAsync(ctx workflow.Context, input *ec2.MonitorInstancesInput) *Ec2MonitorInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.MonitorInstances, input)
    return &Ec2MonitorInstancesResult{Result: future}
}

func (a *EC2Stub) MoveAddressToVpc(ctx workflow.Context, input *ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error) {
    var output ec2.MoveAddressToVpcOutput
    err := workflow.ExecuteActivity(ctx, a.activities.MoveAddressToVpc, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) MoveAddressToVpcAsync(ctx workflow.Context, input *ec2.MoveAddressToVpcInput) *Ec2MoveAddressToVpcResult {
    future := workflow.ExecuteActivity(ctx, a.activities.MoveAddressToVpc, input)
    return &Ec2MoveAddressToVpcResult{Result: future}
}

func (a *EC2Stub) ProvisionByoipCidr(ctx workflow.Context, input *ec2.ProvisionByoipCidrInput) (*ec2.ProvisionByoipCidrOutput, error) {
    var output ec2.ProvisionByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ProvisionByoipCidr, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ProvisionByoipCidrAsync(ctx workflow.Context, input *ec2.ProvisionByoipCidrInput) *Ec2ProvisionByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ProvisionByoipCidr, input)
    return &Ec2ProvisionByoipCidrResult{Result: future}
}

func (a *EC2Stub) PurchaseHostReservation(ctx workflow.Context, input *ec2.PurchaseHostReservationInput) (*ec2.PurchaseHostReservationOutput, error) {
    var output ec2.PurchaseHostReservationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurchaseHostReservation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) PurchaseHostReservationAsync(ctx workflow.Context, input *ec2.PurchaseHostReservationInput) *Ec2PurchaseHostReservationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PurchaseHostReservation, input)
    return &Ec2PurchaseHostReservationResult{Result: future}
}

func (a *EC2Stub) PurchaseReservedInstancesOffering(ctx workflow.Context, input *ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error) {
    var output ec2.PurchaseReservedInstancesOfferingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurchaseReservedInstancesOffering, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) PurchaseReservedInstancesOfferingAsync(ctx workflow.Context, input *ec2.PurchaseReservedInstancesOfferingInput) *Ec2PurchaseReservedInstancesOfferingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PurchaseReservedInstancesOffering, input)
    return &Ec2PurchaseReservedInstancesOfferingResult{Result: future}
}

func (a *EC2Stub) PurchaseScheduledInstances(ctx workflow.Context, input *ec2.PurchaseScheduledInstancesInput) (*ec2.PurchaseScheduledInstancesOutput, error) {
    var output ec2.PurchaseScheduledInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurchaseScheduledInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) PurchaseScheduledInstancesAsync(ctx workflow.Context, input *ec2.PurchaseScheduledInstancesInput) *Ec2PurchaseScheduledInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PurchaseScheduledInstances, input)
    return &Ec2PurchaseScheduledInstancesResult{Result: future}
}

func (a *EC2Stub) RebootInstances(ctx workflow.Context, input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
    var output ec2.RebootInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RebootInstancesAsync(ctx workflow.Context, input *ec2.RebootInstancesInput) *Ec2RebootInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebootInstances, input)
    return &Ec2RebootInstancesResult{Result: future}
}

func (a *EC2Stub) RegisterImage(ctx workflow.Context, input *ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error) {
    var output ec2.RegisterImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterImage, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RegisterImageAsync(ctx workflow.Context, input *ec2.RegisterImageInput) *Ec2RegisterImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterImage, input)
    return &Ec2RegisterImageResult{Result: future}
}

func (a *EC2Stub) RegisterInstanceEventNotificationAttributes(ctx workflow.Context, input *ec2.RegisterInstanceEventNotificationAttributesInput) (*ec2.RegisterInstanceEventNotificationAttributesOutput, error) {
    var output ec2.RegisterInstanceEventNotificationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterInstanceEventNotificationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RegisterInstanceEventNotificationAttributesAsync(ctx workflow.Context, input *ec2.RegisterInstanceEventNotificationAttributesInput) *Ec2RegisterInstanceEventNotificationAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterInstanceEventNotificationAttributes, input)
    return &Ec2RegisterInstanceEventNotificationAttributesResult{Result: future}
}

func (a *EC2Stub) RegisterTransitGatewayMulticastGroupMembers(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) (*ec2.RegisterTransitGatewayMulticastGroupMembersOutput, error) {
    var output ec2.RegisterTransitGatewayMulticastGroupMembersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterTransitGatewayMulticastGroupMembers, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RegisterTransitGatewayMulticastGroupMembersAsync(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupMembersInput) *Ec2RegisterTransitGatewayMulticastGroupMembersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterTransitGatewayMulticastGroupMembers, input)
    return &Ec2RegisterTransitGatewayMulticastGroupMembersResult{Result: future}
}

func (a *EC2Stub) RegisterTransitGatewayMulticastGroupSources(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) (*ec2.RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
    var output ec2.RegisterTransitGatewayMulticastGroupSourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterTransitGatewayMulticastGroupSources, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RegisterTransitGatewayMulticastGroupSourcesAsync(ctx workflow.Context, input *ec2.RegisterTransitGatewayMulticastGroupSourcesInput) *Ec2RegisterTransitGatewayMulticastGroupSourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterTransitGatewayMulticastGroupSources, input)
    return &Ec2RegisterTransitGatewayMulticastGroupSourcesResult{Result: future}
}

func (a *EC2Stub) RejectTransitGatewayPeeringAttachment(ctx workflow.Context, input *ec2.RejectTransitGatewayPeeringAttachmentInput) (*ec2.RejectTransitGatewayPeeringAttachmentOutput, error) {
    var output ec2.RejectTransitGatewayPeeringAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectTransitGatewayPeeringAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RejectTransitGatewayPeeringAttachmentAsync(ctx workflow.Context, input *ec2.RejectTransitGatewayPeeringAttachmentInput) *Ec2RejectTransitGatewayPeeringAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectTransitGatewayPeeringAttachment, input)
    return &Ec2RejectTransitGatewayPeeringAttachmentResult{Result: future}
}

func (a *EC2Stub) RejectTransitGatewayVpcAttachment(ctx workflow.Context, input *ec2.RejectTransitGatewayVpcAttachmentInput) (*ec2.RejectTransitGatewayVpcAttachmentOutput, error) {
    var output ec2.RejectTransitGatewayVpcAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectTransitGatewayVpcAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RejectTransitGatewayVpcAttachmentAsync(ctx workflow.Context, input *ec2.RejectTransitGatewayVpcAttachmentInput) *Ec2RejectTransitGatewayVpcAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectTransitGatewayVpcAttachment, input)
    return &Ec2RejectTransitGatewayVpcAttachmentResult{Result: future}
}

func (a *EC2Stub) RejectVpcEndpointConnections(ctx workflow.Context, input *ec2.RejectVpcEndpointConnectionsInput) (*ec2.RejectVpcEndpointConnectionsOutput, error) {
    var output ec2.RejectVpcEndpointConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectVpcEndpointConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RejectVpcEndpointConnectionsAsync(ctx workflow.Context, input *ec2.RejectVpcEndpointConnectionsInput) *Ec2RejectVpcEndpointConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectVpcEndpointConnections, input)
    return &Ec2RejectVpcEndpointConnectionsResult{Result: future}
}

func (a *EC2Stub) RejectVpcPeeringConnection(ctx workflow.Context, input *ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error) {
    var output ec2.RejectVpcPeeringConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectVpcPeeringConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RejectVpcPeeringConnectionAsync(ctx workflow.Context, input *ec2.RejectVpcPeeringConnectionInput) *Ec2RejectVpcPeeringConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectVpcPeeringConnection, input)
    return &Ec2RejectVpcPeeringConnectionResult{Result: future}
}

func (a *EC2Stub) ReleaseAddress(ctx workflow.Context, input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
    var output ec2.ReleaseAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReleaseAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReleaseAddressAsync(ctx workflow.Context, input *ec2.ReleaseAddressInput) *Ec2ReleaseAddressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReleaseAddress, input)
    return &Ec2ReleaseAddressResult{Result: future}
}

func (a *EC2Stub) ReleaseHosts(ctx workflow.Context, input *ec2.ReleaseHostsInput) (*ec2.ReleaseHostsOutput, error) {
    var output ec2.ReleaseHostsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReleaseHosts, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReleaseHostsAsync(ctx workflow.Context, input *ec2.ReleaseHostsInput) *Ec2ReleaseHostsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReleaseHosts, input)
    return &Ec2ReleaseHostsResult{Result: future}
}

func (a *EC2Stub) ReplaceIamInstanceProfileAssociation(ctx workflow.Context, input *ec2.ReplaceIamInstanceProfileAssociationInput) (*ec2.ReplaceIamInstanceProfileAssociationOutput, error) {
    var output ec2.ReplaceIamInstanceProfileAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReplaceIamInstanceProfileAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReplaceIamInstanceProfileAssociationAsync(ctx workflow.Context, input *ec2.ReplaceIamInstanceProfileAssociationInput) *Ec2ReplaceIamInstanceProfileAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReplaceIamInstanceProfileAssociation, input)
    return &Ec2ReplaceIamInstanceProfileAssociationResult{Result: future}
}

func (a *EC2Stub) ReplaceNetworkAclAssociation(ctx workflow.Context, input *ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error) {
    var output ec2.ReplaceNetworkAclAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReplaceNetworkAclAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReplaceNetworkAclAssociationAsync(ctx workflow.Context, input *ec2.ReplaceNetworkAclAssociationInput) *Ec2ReplaceNetworkAclAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReplaceNetworkAclAssociation, input)
    return &Ec2ReplaceNetworkAclAssociationResult{Result: future}
}

func (a *EC2Stub) ReplaceNetworkAclEntry(ctx workflow.Context, input *ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error) {
    var output ec2.ReplaceNetworkAclEntryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReplaceNetworkAclEntry, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReplaceNetworkAclEntryAsync(ctx workflow.Context, input *ec2.ReplaceNetworkAclEntryInput) *Ec2ReplaceNetworkAclEntryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReplaceNetworkAclEntry, input)
    return &Ec2ReplaceNetworkAclEntryResult{Result: future}
}

func (a *EC2Stub) ReplaceRoute(ctx workflow.Context, input *ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error) {
    var output ec2.ReplaceRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReplaceRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReplaceRouteAsync(ctx workflow.Context, input *ec2.ReplaceRouteInput) *Ec2ReplaceRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReplaceRoute, input)
    return &Ec2ReplaceRouteResult{Result: future}
}

func (a *EC2Stub) ReplaceRouteTableAssociation(ctx workflow.Context, input *ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error) {
    var output ec2.ReplaceRouteTableAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReplaceRouteTableAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReplaceRouteTableAssociationAsync(ctx workflow.Context, input *ec2.ReplaceRouteTableAssociationInput) *Ec2ReplaceRouteTableAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReplaceRouteTableAssociation, input)
    return &Ec2ReplaceRouteTableAssociationResult{Result: future}
}

func (a *EC2Stub) ReplaceTransitGatewayRoute(ctx workflow.Context, input *ec2.ReplaceTransitGatewayRouteInput) (*ec2.ReplaceTransitGatewayRouteOutput, error) {
    var output ec2.ReplaceTransitGatewayRouteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReplaceTransitGatewayRoute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReplaceTransitGatewayRouteAsync(ctx workflow.Context, input *ec2.ReplaceTransitGatewayRouteInput) *Ec2ReplaceTransitGatewayRouteResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReplaceTransitGatewayRoute, input)
    return &Ec2ReplaceTransitGatewayRouteResult{Result: future}
}

func (a *EC2Stub) ReportInstanceStatus(ctx workflow.Context, input *ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error) {
    var output ec2.ReportInstanceStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReportInstanceStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ReportInstanceStatusAsync(ctx workflow.Context, input *ec2.ReportInstanceStatusInput) *Ec2ReportInstanceStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReportInstanceStatus, input)
    return &Ec2ReportInstanceStatusResult{Result: future}
}

func (a *EC2Stub) RequestSpotFleet(ctx workflow.Context, input *ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error) {
    var output ec2.RequestSpotFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RequestSpotFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RequestSpotFleetAsync(ctx workflow.Context, input *ec2.RequestSpotFleetInput) *Ec2RequestSpotFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RequestSpotFleet, input)
    return &Ec2RequestSpotFleetResult{Result: future}
}

func (a *EC2Stub) RequestSpotInstances(ctx workflow.Context, input *ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error) {
    var output ec2.RequestSpotInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RequestSpotInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RequestSpotInstancesAsync(ctx workflow.Context, input *ec2.RequestSpotInstancesInput) *Ec2RequestSpotInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RequestSpotInstances, input)
    return &Ec2RequestSpotInstancesResult{Result: future}
}

func (a *EC2Stub) ResetEbsDefaultKmsKeyId(ctx workflow.Context, input *ec2.ResetEbsDefaultKmsKeyIdInput) (*ec2.ResetEbsDefaultKmsKeyIdOutput, error) {
    var output ec2.ResetEbsDefaultKmsKeyIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetEbsDefaultKmsKeyId, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ResetEbsDefaultKmsKeyIdAsync(ctx workflow.Context, input *ec2.ResetEbsDefaultKmsKeyIdInput) *Ec2ResetEbsDefaultKmsKeyIdResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetEbsDefaultKmsKeyId, input)
    return &Ec2ResetEbsDefaultKmsKeyIdResult{Result: future}
}

func (a *EC2Stub) ResetFpgaImageAttribute(ctx workflow.Context, input *ec2.ResetFpgaImageAttributeInput) (*ec2.ResetFpgaImageAttributeOutput, error) {
    var output ec2.ResetFpgaImageAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetFpgaImageAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ResetFpgaImageAttributeAsync(ctx workflow.Context, input *ec2.ResetFpgaImageAttributeInput) *Ec2ResetFpgaImageAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetFpgaImageAttribute, input)
    return &Ec2ResetFpgaImageAttributeResult{Result: future}
}

func (a *EC2Stub) ResetImageAttribute(ctx workflow.Context, input *ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error) {
    var output ec2.ResetImageAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetImageAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ResetImageAttributeAsync(ctx workflow.Context, input *ec2.ResetImageAttributeInput) *Ec2ResetImageAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetImageAttribute, input)
    return &Ec2ResetImageAttributeResult{Result: future}
}

func (a *EC2Stub) ResetInstanceAttribute(ctx workflow.Context, input *ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error) {
    var output ec2.ResetInstanceAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetInstanceAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ResetInstanceAttributeAsync(ctx workflow.Context, input *ec2.ResetInstanceAttributeInput) *Ec2ResetInstanceAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetInstanceAttribute, input)
    return &Ec2ResetInstanceAttributeResult{Result: future}
}

func (a *EC2Stub) ResetNetworkInterfaceAttribute(ctx workflow.Context, input *ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error) {
    var output ec2.ResetNetworkInterfaceAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetNetworkInterfaceAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ResetNetworkInterfaceAttributeAsync(ctx workflow.Context, input *ec2.ResetNetworkInterfaceAttributeInput) *Ec2ResetNetworkInterfaceAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetNetworkInterfaceAttribute, input)
    return &Ec2ResetNetworkInterfaceAttributeResult{Result: future}
}

func (a *EC2Stub) ResetSnapshotAttribute(ctx workflow.Context, input *ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error) {
    var output ec2.ResetSnapshotAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetSnapshotAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) ResetSnapshotAttributeAsync(ctx workflow.Context, input *ec2.ResetSnapshotAttributeInput) *Ec2ResetSnapshotAttributeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetSnapshotAttribute, input)
    return &Ec2ResetSnapshotAttributeResult{Result: future}
}

func (a *EC2Stub) RestoreAddressToClassic(ctx workflow.Context, input *ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error) {
    var output ec2.RestoreAddressToClassicOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreAddressToClassic, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RestoreAddressToClassicAsync(ctx workflow.Context, input *ec2.RestoreAddressToClassicInput) *Ec2RestoreAddressToClassicResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreAddressToClassic, input)
    return &Ec2RestoreAddressToClassicResult{Result: future}
}

func (a *EC2Stub) RestoreManagedPrefixListVersion(ctx workflow.Context, input *ec2.RestoreManagedPrefixListVersionInput) (*ec2.RestoreManagedPrefixListVersionOutput, error) {
    var output ec2.RestoreManagedPrefixListVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreManagedPrefixListVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RestoreManagedPrefixListVersionAsync(ctx workflow.Context, input *ec2.RestoreManagedPrefixListVersionInput) *Ec2RestoreManagedPrefixListVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreManagedPrefixListVersion, input)
    return &Ec2RestoreManagedPrefixListVersionResult{Result: future}
}

func (a *EC2Stub) RevokeClientVpnIngress(ctx workflow.Context, input *ec2.RevokeClientVpnIngressInput) (*ec2.RevokeClientVpnIngressOutput, error) {
    var output ec2.RevokeClientVpnIngressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeClientVpnIngress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RevokeClientVpnIngressAsync(ctx workflow.Context, input *ec2.RevokeClientVpnIngressInput) *Ec2RevokeClientVpnIngressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RevokeClientVpnIngress, input)
    return &Ec2RevokeClientVpnIngressResult{Result: future}
}

func (a *EC2Stub) RevokeSecurityGroupEgress(ctx workflow.Context, input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
    var output ec2.RevokeSecurityGroupEgressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeSecurityGroupEgress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RevokeSecurityGroupEgressAsync(ctx workflow.Context, input *ec2.RevokeSecurityGroupEgressInput) *Ec2RevokeSecurityGroupEgressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RevokeSecurityGroupEgress, input)
    return &Ec2RevokeSecurityGroupEgressResult{Result: future}
}

func (a *EC2Stub) RevokeSecurityGroupIngress(ctx workflow.Context, input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
    var output ec2.RevokeSecurityGroupIngressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeSecurityGroupIngress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RevokeSecurityGroupIngressAsync(ctx workflow.Context, input *ec2.RevokeSecurityGroupIngressInput) *Ec2RevokeSecurityGroupIngressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RevokeSecurityGroupIngress, input)
    return &Ec2RevokeSecurityGroupIngressResult{Result: future}
}

func (a *EC2Stub) RunInstances(ctx workflow.Context, input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
    var output ec2.Reservation
    err := workflow.ExecuteActivity(ctx, a.activities.RunInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RunInstancesAsync(ctx workflow.Context, input *ec2.RunInstancesInput) *Ec2RunInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RunInstances, input)
    return &Ec2RunInstancesResult{Result: future}
}

func (a *EC2Stub) RunScheduledInstances(ctx workflow.Context, input *ec2.RunScheduledInstancesInput) (*ec2.RunScheduledInstancesOutput, error) {
    var output ec2.RunScheduledInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RunScheduledInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) RunScheduledInstancesAsync(ctx workflow.Context, input *ec2.RunScheduledInstancesInput) *Ec2RunScheduledInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RunScheduledInstances, input)
    return &Ec2RunScheduledInstancesResult{Result: future}
}

func (a *EC2Stub) SearchLocalGatewayRoutes(ctx workflow.Context, input *ec2.SearchLocalGatewayRoutesInput) (*ec2.SearchLocalGatewayRoutesOutput, error) {
    var output ec2.SearchLocalGatewayRoutesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchLocalGatewayRoutes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) SearchLocalGatewayRoutesAsync(ctx workflow.Context, input *ec2.SearchLocalGatewayRoutesInput) *Ec2SearchLocalGatewayRoutesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchLocalGatewayRoutes, input)
    return &Ec2SearchLocalGatewayRoutesResult{Result: future}
}

func (a *EC2Stub) SearchTransitGatewayMulticastGroups(ctx workflow.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
    var output ec2.SearchTransitGatewayMulticastGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchTransitGatewayMulticastGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) SearchTransitGatewayMulticastGroupsAsync(ctx workflow.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput) *Ec2SearchTransitGatewayMulticastGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchTransitGatewayMulticastGroups, input)
    return &Ec2SearchTransitGatewayMulticastGroupsResult{Result: future}
}

func (a *EC2Stub) SearchTransitGatewayRoutes(ctx workflow.Context, input *ec2.SearchTransitGatewayRoutesInput) (*ec2.SearchTransitGatewayRoutesOutput, error) {
    var output ec2.SearchTransitGatewayRoutesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchTransitGatewayRoutes, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) SearchTransitGatewayRoutesAsync(ctx workflow.Context, input *ec2.SearchTransitGatewayRoutesInput) *Ec2SearchTransitGatewayRoutesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchTransitGatewayRoutes, input)
    return &Ec2SearchTransitGatewayRoutesResult{Result: future}
}

func (a *EC2Stub) SendDiagnosticInterrupt(ctx workflow.Context, input *ec2.SendDiagnosticInterruptInput) (*ec2.SendDiagnosticInterruptOutput, error) {
    var output ec2.SendDiagnosticInterruptOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendDiagnosticInterrupt, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) SendDiagnosticInterruptAsync(ctx workflow.Context, input *ec2.SendDiagnosticInterruptInput) *Ec2SendDiagnosticInterruptResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendDiagnosticInterrupt, input)
    return &Ec2SendDiagnosticInterruptResult{Result: future}
}

func (a *EC2Stub) StartInstances(ctx workflow.Context, input *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
    var output ec2.StartInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) StartInstancesAsync(ctx workflow.Context, input *ec2.StartInstancesInput) *Ec2StartInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartInstances, input)
    return &Ec2StartInstancesResult{Result: future}
}

func (a *EC2Stub) StartVpcEndpointServicePrivateDnsVerification(ctx workflow.Context, input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) (*ec2.StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
    var output ec2.StartVpcEndpointServicePrivateDnsVerificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartVpcEndpointServicePrivateDnsVerification, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) StartVpcEndpointServicePrivateDnsVerificationAsync(ctx workflow.Context, input *ec2.StartVpcEndpointServicePrivateDnsVerificationInput) *Ec2StartVpcEndpointServicePrivateDnsVerificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartVpcEndpointServicePrivateDnsVerification, input)
    return &Ec2StartVpcEndpointServicePrivateDnsVerificationResult{Result: future}
}

func (a *EC2Stub) StopInstances(ctx workflow.Context, input *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
    var output ec2.StopInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) StopInstancesAsync(ctx workflow.Context, input *ec2.StopInstancesInput) *Ec2StopInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopInstances, input)
    return &Ec2StopInstancesResult{Result: future}
}

func (a *EC2Stub) TerminateClientVpnConnections(ctx workflow.Context, input *ec2.TerminateClientVpnConnectionsInput) (*ec2.TerminateClientVpnConnectionsOutput, error) {
    var output ec2.TerminateClientVpnConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateClientVpnConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) TerminateClientVpnConnectionsAsync(ctx workflow.Context, input *ec2.TerminateClientVpnConnectionsInput) *Ec2TerminateClientVpnConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateClientVpnConnections, input)
    return &Ec2TerminateClientVpnConnectionsResult{Result: future}
}

func (a *EC2Stub) TerminateInstances(ctx workflow.Context, input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
    var output ec2.TerminateInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) TerminateInstancesAsync(ctx workflow.Context, input *ec2.TerminateInstancesInput) *Ec2TerminateInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateInstances, input)
    return &Ec2TerminateInstancesResult{Result: future}
}

func (a *EC2Stub) UnassignIpv6Addresses(ctx workflow.Context, input *ec2.UnassignIpv6AddressesInput) (*ec2.UnassignIpv6AddressesOutput, error) {
    var output ec2.UnassignIpv6AddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnassignIpv6Addresses, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) UnassignIpv6AddressesAsync(ctx workflow.Context, input *ec2.UnassignIpv6AddressesInput) *Ec2UnassignIpv6AddressesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnassignIpv6Addresses, input)
    return &Ec2UnassignIpv6AddressesResult{Result: future}
}

func (a *EC2Stub) UnassignPrivateIpAddresses(ctx workflow.Context, input *ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error) {
    var output ec2.UnassignPrivateIpAddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnassignPrivateIpAddresses, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) UnassignPrivateIpAddressesAsync(ctx workflow.Context, input *ec2.UnassignPrivateIpAddressesInput) *Ec2UnassignPrivateIpAddressesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnassignPrivateIpAddresses, input)
    return &Ec2UnassignPrivateIpAddressesResult{Result: future}
}

func (a *EC2Stub) UnmonitorInstances(ctx workflow.Context, input *ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error) {
    var output ec2.UnmonitorInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnmonitorInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) UnmonitorInstancesAsync(ctx workflow.Context, input *ec2.UnmonitorInstancesInput) *Ec2UnmonitorInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnmonitorInstances, input)
    return &Ec2UnmonitorInstancesResult{Result: future}
}

func (a *EC2Stub) UpdateSecurityGroupRuleDescriptionsEgress(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput, error) {
    var output ec2.UpdateSecurityGroupRuleDescriptionsEgressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSecurityGroupRuleDescriptionsEgress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) UpdateSecurityGroupRuleDescriptionsEgressAsync(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsEgressInput) *Ec2UpdateSecurityGroupRuleDescriptionsEgressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSecurityGroupRuleDescriptionsEgress, input)
    return &Ec2UpdateSecurityGroupRuleDescriptionsEgressResult{Result: future}
}

func (a *EC2Stub) UpdateSecurityGroupRuleDescriptionsIngress(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) (*ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput, error) {
    var output ec2.UpdateSecurityGroupRuleDescriptionsIngressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSecurityGroupRuleDescriptionsIngress, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) UpdateSecurityGroupRuleDescriptionsIngressAsync(ctx workflow.Context, input *ec2.UpdateSecurityGroupRuleDescriptionsIngressInput) *Ec2UpdateSecurityGroupRuleDescriptionsIngressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSecurityGroupRuleDescriptionsIngress, input)
    return &Ec2UpdateSecurityGroupRuleDescriptionsIngressResult{Result: future}
}

func (a *EC2Stub) WithdrawByoipCidr(ctx workflow.Context, input *ec2.WithdrawByoipCidrInput) (*ec2.WithdrawByoipCidrOutput, error) {
    var output ec2.WithdrawByoipCidrOutput
    err := workflow.ExecuteActivity(ctx, a.activities.WithdrawByoipCidr, input).Get(ctx, &output)
    return &output, err
}

func (a *EC2Stub) WithdrawByoipCidrAsync(ctx workflow.Context, input *ec2.WithdrawByoipCidrInput) *Ec2WithdrawByoipCidrResult {
    future := workflow.ExecuteActivity(ctx, a.activities.WithdrawByoipCidr, input)
    return &Ec2WithdrawByoipCidrResult{Result: future}
}

func (a *EC2Stub) WaitUntilBundleTaskComplete(ctx workflow.Context, input *ec2.DescribeBundleTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBundleTaskComplete, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilBundleTaskCompleteAsync(ctx workflow.Context, input *ec2.DescribeBundleTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBundleTaskComplete, input)
}


func (a *EC2Stub) WaitUntilConversionTaskCancelled(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilConversionTaskCancelled, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilConversionTaskCancelledAsync(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilConversionTaskCancelled, input)
}


func (a *EC2Stub) WaitUntilConversionTaskCompleted(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilConversionTaskCompleted, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilConversionTaskCompletedAsync(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilConversionTaskCompleted, input)
}


func (a *EC2Stub) WaitUntilConversionTaskDeleted(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilConversionTaskDeleted, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilConversionTaskDeletedAsync(ctx workflow.Context, input *ec2.DescribeConversionTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilConversionTaskDeleted, input)
}


func (a *EC2Stub) WaitUntilCustomerGatewayAvailable(ctx workflow.Context, input *ec2.DescribeCustomerGatewaysInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilCustomerGatewayAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilCustomerGatewayAvailableAsync(ctx workflow.Context, input *ec2.DescribeCustomerGatewaysInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilCustomerGatewayAvailable, input)
}


func (a *EC2Stub) WaitUntilExportTaskCancelled(ctx workflow.Context, input *ec2.DescribeExportTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilExportTaskCancelled, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilExportTaskCancelledAsync(ctx workflow.Context, input *ec2.DescribeExportTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilExportTaskCancelled, input)
}


func (a *EC2Stub) WaitUntilExportTaskCompleted(ctx workflow.Context, input *ec2.DescribeExportTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilExportTaskCompleted, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilExportTaskCompletedAsync(ctx workflow.Context, input *ec2.DescribeExportTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilExportTaskCompleted, input)
}


func (a *EC2Stub) WaitUntilImageAvailable(ctx workflow.Context, input *ec2.DescribeImagesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilImageAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilImageAvailableAsync(ctx workflow.Context, input *ec2.DescribeImagesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilImageAvailable, input)
}


func (a *EC2Stub) WaitUntilImageExists(ctx workflow.Context, input *ec2.DescribeImagesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilImageExists, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilImageExistsAsync(ctx workflow.Context, input *ec2.DescribeImagesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilImageExists, input)
}


func (a *EC2Stub) WaitUntilInstanceExists(ctx workflow.Context, input *ec2.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceExists, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilInstanceExistsAsync(ctx workflow.Context, input *ec2.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceExists, input)
}


func (a *EC2Stub) WaitUntilInstanceRunning(ctx workflow.Context, input *ec2.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceRunning, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilInstanceRunningAsync(ctx workflow.Context, input *ec2.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceRunning, input)
}


func (a *EC2Stub) WaitUntilInstanceStatusOk(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceStatusOk, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilInstanceStatusOkAsync(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceStatusOk, input)
}


func (a *EC2Stub) WaitUntilInstanceStopped(ctx workflow.Context, input *ec2.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceStopped, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilInstanceStoppedAsync(ctx workflow.Context, input *ec2.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceStopped, input)
}


func (a *EC2Stub) WaitUntilInstanceTerminated(ctx workflow.Context, input *ec2.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceTerminated, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilInstanceTerminatedAsync(ctx workflow.Context, input *ec2.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceTerminated, input)
}


func (a *EC2Stub) WaitUntilKeyPairExists(ctx workflow.Context, input *ec2.DescribeKeyPairsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilKeyPairExists, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilKeyPairExistsAsync(ctx workflow.Context, input *ec2.DescribeKeyPairsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilKeyPairExists, input)
}


func (a *EC2Stub) WaitUntilNatGatewayAvailable(ctx workflow.Context, input *ec2.DescribeNatGatewaysInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilNatGatewayAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilNatGatewayAvailableAsync(ctx workflow.Context, input *ec2.DescribeNatGatewaysInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilNatGatewayAvailable, input)
}


func (a *EC2Stub) WaitUntilNetworkInterfaceAvailable(ctx workflow.Context, input *ec2.DescribeNetworkInterfacesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilNetworkInterfaceAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilNetworkInterfaceAvailableAsync(ctx workflow.Context, input *ec2.DescribeNetworkInterfacesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilNetworkInterfaceAvailable, input)
}


func (a *EC2Stub) WaitUntilPasswordDataAvailable(ctx workflow.Context, input *ec2.GetPasswordDataInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilPasswordDataAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilPasswordDataAvailableAsync(ctx workflow.Context, input *ec2.GetPasswordDataInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilPasswordDataAvailable, input)
}


func (a *EC2Stub) WaitUntilSecurityGroupExists(ctx workflow.Context, input *ec2.DescribeSecurityGroupsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSecurityGroupExists, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilSecurityGroupExistsAsync(ctx workflow.Context, input *ec2.DescribeSecurityGroupsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSecurityGroupExists, input)
}


func (a *EC2Stub) WaitUntilSnapshotCompleted(ctx workflow.Context, input *ec2.DescribeSnapshotsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSnapshotCompleted, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilSnapshotCompletedAsync(ctx workflow.Context, input *ec2.DescribeSnapshotsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSnapshotCompleted, input)
}


func (a *EC2Stub) WaitUntilSpotInstanceRequestFulfilled(ctx workflow.Context, input *ec2.DescribeSpotInstanceRequestsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSpotInstanceRequestFulfilled, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilSpotInstanceRequestFulfilledAsync(ctx workflow.Context, input *ec2.DescribeSpotInstanceRequestsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSpotInstanceRequestFulfilled, input)
}


func (a *EC2Stub) WaitUntilSubnetAvailable(ctx workflow.Context, input *ec2.DescribeSubnetsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSubnetAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilSubnetAvailableAsync(ctx workflow.Context, input *ec2.DescribeSubnetsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSubnetAvailable, input)
}


func (a *EC2Stub) WaitUntilSystemStatusOk(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSystemStatusOk, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilSystemStatusOkAsync(ctx workflow.Context, input *ec2.DescribeInstanceStatusInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilSystemStatusOk, input)
}


func (a *EC2Stub) WaitUntilVolumeAvailable(ctx workflow.Context, input *ec2.DescribeVolumesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVolumeAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVolumeAvailableAsync(ctx workflow.Context, input *ec2.DescribeVolumesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVolumeAvailable, input)
}


func (a *EC2Stub) WaitUntilVolumeDeleted(ctx workflow.Context, input *ec2.DescribeVolumesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVolumeDeleted, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVolumeDeletedAsync(ctx workflow.Context, input *ec2.DescribeVolumesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVolumeDeleted, input)
}


func (a *EC2Stub) WaitUntilVolumeInUse(ctx workflow.Context, input *ec2.DescribeVolumesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVolumeInUse, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVolumeInUseAsync(ctx workflow.Context, input *ec2.DescribeVolumesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVolumeInUse, input)
}


func (a *EC2Stub) WaitUntilVpcAvailable(ctx workflow.Context, input *ec2.DescribeVpcsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVpcAvailableAsync(ctx workflow.Context, input *ec2.DescribeVpcsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcAvailable, input)
}


func (a *EC2Stub) WaitUntilVpcExists(ctx workflow.Context, input *ec2.DescribeVpcsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcExists, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVpcExistsAsync(ctx workflow.Context, input *ec2.DescribeVpcsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcExists, input)
}


func (a *EC2Stub) WaitUntilVpcPeeringConnectionDeleted(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcPeeringConnectionDeleted, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVpcPeeringConnectionDeletedAsync(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcPeeringConnectionDeleted, input)
}


func (a *EC2Stub) WaitUntilVpcPeeringConnectionExists(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcPeeringConnectionExists, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVpcPeeringConnectionExistsAsync(ctx workflow.Context, input *ec2.DescribeVpcPeeringConnectionsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpcPeeringConnectionExists, input)
}


func (a *EC2Stub) WaitUntilVpnConnectionAvailable(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpnConnectionAvailable, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVpnConnectionAvailableAsync(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpnConnectionAvailable, input)
}


func (a *EC2Stub) WaitUntilVpnConnectionDeleted(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpnConnectionDeleted, input).Get(ctx, nil)
}

func (a *EC2Stub) WaitUntilVpnConnectionDeletedAsync(ctx workflow.Context, input *ec2.DescribeVpnConnectionsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVpnConnectionDeleted, input)
}

