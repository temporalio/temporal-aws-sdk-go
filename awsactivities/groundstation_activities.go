
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/groundstation"
	"github.com/aws/aws-sdk-go/service/groundstation/groundstationiface"
)

type GroundStationActivities struct {
	client groundstationiface.GroundStationAPI
}

func NewGroundStationActivities(session *session.Session, config... *aws.Config) *GroundStationActivities {
    client := groundstation.New(session, config...)
    return &GroundStationActivities{client: client}
}

func (a *GroundStationActivities) CancelContact(input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error) {
    return a.client.CancelContact(input)
}

func (a *GroundStationActivities) CreateConfig(input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error) {
    return a.client.CreateConfig(input)
}

func (a *GroundStationActivities) CreateDataflowEndpointGroup(input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error) {
    return a.client.CreateDataflowEndpointGroup(input)
}

func (a *GroundStationActivities) CreateMissionProfile(input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error) {
    return a.client.CreateMissionProfile(input)
}

func (a *GroundStationActivities) DeleteConfig(input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error) {
    return a.client.DeleteConfig(input)
}

func (a *GroundStationActivities) DeleteDataflowEndpointGroup(input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error) {
    return a.client.DeleteDataflowEndpointGroup(input)
}

func (a *GroundStationActivities) DeleteMissionProfile(input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error) {
    return a.client.DeleteMissionProfile(input)
}

func (a *GroundStationActivities) DescribeContact(input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error) {
    return a.client.DescribeContact(input)
}

func (a *GroundStationActivities) GetConfig(input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error) {
    return a.client.GetConfig(input)
}

func (a *GroundStationActivities) GetDataflowEndpointGroup(input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error) {
    return a.client.GetDataflowEndpointGroup(input)
}

func (a *GroundStationActivities) GetMinuteUsage(input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error) {
    return a.client.GetMinuteUsage(input)
}

func (a *GroundStationActivities) GetMissionProfile(input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error) {
    return a.client.GetMissionProfile(input)
}

func (a *GroundStationActivities) GetSatellite(input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error) {
    return a.client.GetSatellite(input)
}

func (a *GroundStationActivities) ListConfigs(input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error) {
    return a.client.ListConfigs(input)
}

func (a *GroundStationActivities) ListContacts(input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error) {
    return a.client.ListContacts(input)
}

func (a *GroundStationActivities) ListDataflowEndpointGroups(input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
    return a.client.ListDataflowEndpointGroups(input)
}

func (a *GroundStationActivities) ListGroundStations(input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error) {
    return a.client.ListGroundStations(input)
}

func (a *GroundStationActivities) ListMissionProfiles(input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error) {
    return a.client.ListMissionProfiles(input)
}

func (a *GroundStationActivities) ListSatellites(input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error) {
    return a.client.ListSatellites(input)
}

func (a *GroundStationActivities) ListTagsForResource(input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *GroundStationActivities) ReserveContact(input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error) {
    return a.client.ReserveContact(input)
}

func (a *GroundStationActivities) TagResource(input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *GroundStationActivities) UntagResource(input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *GroundStationActivities) UpdateConfig(input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error) {
    return a.client.UpdateConfig(input)
}

func (a *GroundStationActivities) UpdateMissionProfile(input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error) {
    return a.client.UpdateMissionProfile(input)
}
