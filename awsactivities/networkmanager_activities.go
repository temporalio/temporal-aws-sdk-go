package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/aws/aws-sdk-go/service/networkmanager/networkmanageriface"
)

type NetworkManagerActivities struct {
	client networkmanageriface.NetworkManagerAPI
}

func NewNetworkManagerActivities(client networkmanageriface.NetworkManagerAPI) *NetworkManagerActivities {
    return &NetworkManagerActivities{client: client}
}


func (a *NetworkManagerActivities) AssociateCustomerGateway(input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error) {
    return a.client.AssociateCustomerGateway(input)
}



func (a *NetworkManagerActivities) AssociateLink(input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error) {
    return a.client.AssociateLink(input)
}



func (a *NetworkManagerActivities) CreateDevice(input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error) {
    return a.client.CreateDevice(input)
}



func (a *NetworkManagerActivities) CreateGlobalNetwork(input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error) {
    return a.client.CreateGlobalNetwork(input)
}



func (a *NetworkManagerActivities) CreateLink(input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error) {
    return a.client.CreateLink(input)
}



func (a *NetworkManagerActivities) CreateSite(input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error) {
    return a.client.CreateSite(input)
}



func (a *NetworkManagerActivities) DeleteDevice(input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error) {
    return a.client.DeleteDevice(input)
}



func (a *NetworkManagerActivities) DeleteGlobalNetwork(input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error) {
    return a.client.DeleteGlobalNetwork(input)
}



func (a *NetworkManagerActivities) DeleteLink(input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error) {
    return a.client.DeleteLink(input)
}



func (a *NetworkManagerActivities) DeleteSite(input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error) {
    return a.client.DeleteSite(input)
}



func (a *NetworkManagerActivities) DeregisterTransitGateway(input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error) {
    return a.client.DeregisterTransitGateway(input)
}



func (a *NetworkManagerActivities) DescribeGlobalNetworks(input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error) {
    return a.client.DescribeGlobalNetworks(input)
}



func (a *NetworkManagerActivities) DisassociateCustomerGateway(input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
    return a.client.DisassociateCustomerGateway(input)
}



func (a *NetworkManagerActivities) DisassociateLink(input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error) {
    return a.client.DisassociateLink(input)
}



func (a *NetworkManagerActivities) GetCustomerGatewayAssociations(input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
    return a.client.GetCustomerGatewayAssociations(input)
}



func (a *NetworkManagerActivities) GetDevices(input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error) {
    return a.client.GetDevices(input)
}



func (a *NetworkManagerActivities) GetLinkAssociations(input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error) {
    return a.client.GetLinkAssociations(input)
}



func (a *NetworkManagerActivities) GetLinks(input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error) {
    return a.client.GetLinks(input)
}



func (a *NetworkManagerActivities) GetSites(input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error) {
    return a.client.GetSites(input)
}



func (a *NetworkManagerActivities) GetTransitGatewayRegistrations(input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
    return a.client.GetTransitGatewayRegistrations(input)
}



func (a *NetworkManagerActivities) ListTagsForResource(input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *NetworkManagerActivities) RegisterTransitGateway(input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error) {
    return a.client.RegisterTransitGateway(input)
}



func (a *NetworkManagerActivities) TagResource(input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *NetworkManagerActivities) UntagResource(input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *NetworkManagerActivities) UpdateDevice(input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error) {
    return a.client.UpdateDevice(input)
}



func (a *NetworkManagerActivities) UpdateGlobalNetwork(input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error) {
    return a.client.UpdateGlobalNetwork(input)
}



func (a *NetworkManagerActivities) UpdateLink(input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error) {
    return a.client.UpdateLink(input)
}



func (a *NetworkManagerActivities) UpdateSite(input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error) {
    return a.client.UpdateSite(input)
}

