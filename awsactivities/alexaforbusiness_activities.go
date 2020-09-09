
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness/alexaforbusinessiface"
)

type AlexaForBusinessActivities struct {
	client alexaforbusinessiface.AlexaForBusinessAPI
}

func NewAlexaForBusinessActivities(session *session.Session, config... *aws.Config) *AlexaForBusinessActivities {
    client := alexaforbusiness.New(session, config...)
    return &AlexaForBusinessActivities{client: client}
}

func (a *AlexaForBusinessActivities) ApproveSkill(input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error) {
    return a.client.ApproveSkill(input)
}

func (a *AlexaForBusinessActivities) AssociateContactWithAddressBook(input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error) {
    return a.client.AssociateContactWithAddressBook(input)
}

func (a *AlexaForBusinessActivities) AssociateDeviceWithNetworkProfile(input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error) {
    return a.client.AssociateDeviceWithNetworkProfile(input)
}

func (a *AlexaForBusinessActivities) AssociateDeviceWithRoom(input *alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error) {
    return a.client.AssociateDeviceWithRoom(input)
}

func (a *AlexaForBusinessActivities) AssociateSkillGroupWithRoom(input *alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error) {
    return a.client.AssociateSkillGroupWithRoom(input)
}

func (a *AlexaForBusinessActivities) AssociateSkillWithSkillGroup(input *alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error) {
    return a.client.AssociateSkillWithSkillGroup(input)
}

func (a *AlexaForBusinessActivities) AssociateSkillWithUsers(input *alexaforbusiness.AssociateSkillWithUsersInput) (*alexaforbusiness.AssociateSkillWithUsersOutput, error) {
    return a.client.AssociateSkillWithUsers(input)
}

func (a *AlexaForBusinessActivities) CreateAddressBook(input *alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error) {
    return a.client.CreateAddressBook(input)
}

func (a *AlexaForBusinessActivities) CreateBusinessReportSchedule(input *alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error) {
    return a.client.CreateBusinessReportSchedule(input)
}

func (a *AlexaForBusinessActivities) CreateConferenceProvider(input *alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error) {
    return a.client.CreateConferenceProvider(input)
}

func (a *AlexaForBusinessActivities) CreateContact(input *alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error) {
    return a.client.CreateContact(input)
}

func (a *AlexaForBusinessActivities) CreateGatewayGroup(input *alexaforbusiness.CreateGatewayGroupInput) (*alexaforbusiness.CreateGatewayGroupOutput, error) {
    return a.client.CreateGatewayGroup(input)
}

func (a *AlexaForBusinessActivities) CreateNetworkProfile(input *alexaforbusiness.CreateNetworkProfileInput) (*alexaforbusiness.CreateNetworkProfileOutput, error) {
    return a.client.CreateNetworkProfile(input)
}

func (a *AlexaForBusinessActivities) CreateProfile(input *alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error) {
    return a.client.CreateProfile(input)
}

func (a *AlexaForBusinessActivities) CreateRoom(input *alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error) {
    return a.client.CreateRoom(input)
}

func (a *AlexaForBusinessActivities) CreateSkillGroup(input *alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error) {
    return a.client.CreateSkillGroup(input)
}

func (a *AlexaForBusinessActivities) CreateUser(input *alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *AlexaForBusinessActivities) DeleteAddressBook(input *alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error) {
    return a.client.DeleteAddressBook(input)
}

func (a *AlexaForBusinessActivities) DeleteBusinessReportSchedule(input *alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error) {
    return a.client.DeleteBusinessReportSchedule(input)
}

func (a *AlexaForBusinessActivities) DeleteConferenceProvider(input *alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error) {
    return a.client.DeleteConferenceProvider(input)
}

func (a *AlexaForBusinessActivities) DeleteContact(input *alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error) {
    return a.client.DeleteContact(input)
}

func (a *AlexaForBusinessActivities) DeleteDevice(input *alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error) {
    return a.client.DeleteDevice(input)
}

func (a *AlexaForBusinessActivities) DeleteDeviceUsageData(input *alexaforbusiness.DeleteDeviceUsageDataInput) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error) {
    return a.client.DeleteDeviceUsageData(input)
}

func (a *AlexaForBusinessActivities) DeleteGatewayGroup(input *alexaforbusiness.DeleteGatewayGroupInput) (*alexaforbusiness.DeleteGatewayGroupOutput, error) {
    return a.client.DeleteGatewayGroup(input)
}

func (a *AlexaForBusinessActivities) DeleteNetworkProfile(input *alexaforbusiness.DeleteNetworkProfileInput) (*alexaforbusiness.DeleteNetworkProfileOutput, error) {
    return a.client.DeleteNetworkProfile(input)
}

func (a *AlexaForBusinessActivities) DeleteProfile(input *alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error) {
    return a.client.DeleteProfile(input)
}

func (a *AlexaForBusinessActivities) DeleteRoom(input *alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error) {
    return a.client.DeleteRoom(input)
}

func (a *AlexaForBusinessActivities) DeleteRoomSkillParameter(input *alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error) {
    return a.client.DeleteRoomSkillParameter(input)
}

func (a *AlexaForBusinessActivities) DeleteSkillAuthorization(input *alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error) {
    return a.client.DeleteSkillAuthorization(input)
}

func (a *AlexaForBusinessActivities) DeleteSkillGroup(input *alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error) {
    return a.client.DeleteSkillGroup(input)
}

func (a *AlexaForBusinessActivities) DeleteUser(input *alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *AlexaForBusinessActivities) DisassociateContactFromAddressBook(input *alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error) {
    return a.client.DisassociateContactFromAddressBook(input)
}

func (a *AlexaForBusinessActivities) DisassociateDeviceFromRoom(input *alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error) {
    return a.client.DisassociateDeviceFromRoom(input)
}

func (a *AlexaForBusinessActivities) DisassociateSkillFromSkillGroup(input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error) {
    return a.client.DisassociateSkillFromSkillGroup(input)
}

func (a *AlexaForBusinessActivities) DisassociateSkillFromUsers(input *alexaforbusiness.DisassociateSkillFromUsersInput) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error) {
    return a.client.DisassociateSkillFromUsers(input)
}

func (a *AlexaForBusinessActivities) DisassociateSkillGroupFromRoom(input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error) {
    return a.client.DisassociateSkillGroupFromRoom(input)
}

func (a *AlexaForBusinessActivities) ForgetSmartHomeAppliances(input *alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error) {
    return a.client.ForgetSmartHomeAppliances(input)
}

func (a *AlexaForBusinessActivities) GetAddressBook(input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error) {
    return a.client.GetAddressBook(input)
}

func (a *AlexaForBusinessActivities) GetConferencePreference(input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error) {
    return a.client.GetConferencePreference(input)
}

func (a *AlexaForBusinessActivities) GetConferenceProvider(input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error) {
    return a.client.GetConferenceProvider(input)
}

func (a *AlexaForBusinessActivities) GetContact(input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error) {
    return a.client.GetContact(input)
}

func (a *AlexaForBusinessActivities) GetDevice(input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error) {
    return a.client.GetDevice(input)
}

func (a *AlexaForBusinessActivities) GetGateway(input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error) {
    return a.client.GetGateway(input)
}

func (a *AlexaForBusinessActivities) GetGatewayGroup(input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error) {
    return a.client.GetGatewayGroup(input)
}

func (a *AlexaForBusinessActivities) GetInvitationConfiguration(input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error) {
    return a.client.GetInvitationConfiguration(input)
}

func (a *AlexaForBusinessActivities) GetNetworkProfile(input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error) {
    return a.client.GetNetworkProfile(input)
}

func (a *AlexaForBusinessActivities) GetProfile(input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error) {
    return a.client.GetProfile(input)
}

func (a *AlexaForBusinessActivities) GetRoom(input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error) {
    return a.client.GetRoom(input)
}

func (a *AlexaForBusinessActivities) GetRoomSkillParameter(input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error) {
    return a.client.GetRoomSkillParameter(input)
}

func (a *AlexaForBusinessActivities) GetSkillGroup(input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error) {
    return a.client.GetSkillGroup(input)
}

func (a *AlexaForBusinessActivities) ListBusinessReportSchedules(input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error) {
    return a.client.ListBusinessReportSchedules(input)
}

func (a *AlexaForBusinessActivities) ListConferenceProviders(input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error) {
    return a.client.ListConferenceProviders(input)
}

func (a *AlexaForBusinessActivities) ListDeviceEvents(input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error) {
    return a.client.ListDeviceEvents(input)
}

func (a *AlexaForBusinessActivities) ListGatewayGroups(input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error) {
    return a.client.ListGatewayGroups(input)
}

func (a *AlexaForBusinessActivities) ListGateways(input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error) {
    return a.client.ListGateways(input)
}

func (a *AlexaForBusinessActivities) ListSkills(input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error) {
    return a.client.ListSkills(input)
}

func (a *AlexaForBusinessActivities) ListSkillsStoreCategories(input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error) {
    return a.client.ListSkillsStoreCategories(input)
}

func (a *AlexaForBusinessActivities) ListSkillsStoreSkillsByCategory(input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error) {
    return a.client.ListSkillsStoreSkillsByCategory(input)
}

func (a *AlexaForBusinessActivities) ListSmartHomeAppliances(input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error) {
    return a.client.ListSmartHomeAppliances(input)
}

func (a *AlexaForBusinessActivities) ListTags(input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *AlexaForBusinessActivities) PutConferencePreference(input *alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error) {
    return a.client.PutConferencePreference(input)
}

func (a *AlexaForBusinessActivities) PutInvitationConfiguration(input *alexaforbusiness.PutInvitationConfigurationInput) (*alexaforbusiness.PutInvitationConfigurationOutput, error) {
    return a.client.PutInvitationConfiguration(input)
}

func (a *AlexaForBusinessActivities) PutRoomSkillParameter(input *alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error) {
    return a.client.PutRoomSkillParameter(input)
}

func (a *AlexaForBusinessActivities) PutSkillAuthorization(input *alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error) {
    return a.client.PutSkillAuthorization(input)
}

func (a *AlexaForBusinessActivities) RegisterAVSDevice(input *alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error) {
    return a.client.RegisterAVSDevice(input)
}

func (a *AlexaForBusinessActivities) RejectSkill(input *alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error) {
    return a.client.RejectSkill(input)
}

func (a *AlexaForBusinessActivities) ResolveRoom(input *alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error) {
    return a.client.ResolveRoom(input)
}

func (a *AlexaForBusinessActivities) RevokeInvitation(input *alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error) {
    return a.client.RevokeInvitation(input)
}

func (a *AlexaForBusinessActivities) SearchAddressBooks(input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error) {
    return a.client.SearchAddressBooks(input)
}

func (a *AlexaForBusinessActivities) SearchContacts(input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error) {
    return a.client.SearchContacts(input)
}

func (a *AlexaForBusinessActivities) SearchDevices(input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error) {
    return a.client.SearchDevices(input)
}

func (a *AlexaForBusinessActivities) SearchNetworkProfiles(input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error) {
    return a.client.SearchNetworkProfiles(input)
}

func (a *AlexaForBusinessActivities) SearchProfiles(input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error) {
    return a.client.SearchProfiles(input)
}

func (a *AlexaForBusinessActivities) SearchRooms(input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error) {
    return a.client.SearchRooms(input)
}

func (a *AlexaForBusinessActivities) SearchSkillGroups(input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error) {
    return a.client.SearchSkillGroups(input)
}

func (a *AlexaForBusinessActivities) SearchUsers(input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error) {
    return a.client.SearchUsers(input)
}

func (a *AlexaForBusinessActivities) SendAnnouncement(input *alexaforbusiness.SendAnnouncementInput) (*alexaforbusiness.SendAnnouncementOutput, error) {
    return a.client.SendAnnouncement(input)
}

func (a *AlexaForBusinessActivities) SendInvitation(input *alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error) {
    return a.client.SendInvitation(input)
}

func (a *AlexaForBusinessActivities) StartDeviceSync(input *alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error) {
    return a.client.StartDeviceSync(input)
}

func (a *AlexaForBusinessActivities) StartSmartHomeApplianceDiscovery(input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error) {
    return a.client.StartSmartHomeApplianceDiscovery(input)
}

func (a *AlexaForBusinessActivities) TagResource(input *alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *AlexaForBusinessActivities) UntagResource(input *alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *AlexaForBusinessActivities) UpdateAddressBook(input *alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error) {
    return a.client.UpdateAddressBook(input)
}

func (a *AlexaForBusinessActivities) UpdateBusinessReportSchedule(input *alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error) {
    return a.client.UpdateBusinessReportSchedule(input)
}

func (a *AlexaForBusinessActivities) UpdateConferenceProvider(input *alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error) {
    return a.client.UpdateConferenceProvider(input)
}

func (a *AlexaForBusinessActivities) UpdateContact(input *alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error) {
    return a.client.UpdateContact(input)
}

func (a *AlexaForBusinessActivities) UpdateDevice(input *alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error) {
    return a.client.UpdateDevice(input)
}

func (a *AlexaForBusinessActivities) UpdateGateway(input *alexaforbusiness.UpdateGatewayInput) (*alexaforbusiness.UpdateGatewayOutput, error) {
    return a.client.UpdateGateway(input)
}

func (a *AlexaForBusinessActivities) UpdateGatewayGroup(input *alexaforbusiness.UpdateGatewayGroupInput) (*alexaforbusiness.UpdateGatewayGroupOutput, error) {
    return a.client.UpdateGatewayGroup(input)
}

func (a *AlexaForBusinessActivities) UpdateNetworkProfile(input *alexaforbusiness.UpdateNetworkProfileInput) (*alexaforbusiness.UpdateNetworkProfileOutput, error) {
    return a.client.UpdateNetworkProfile(input)
}

func (a *AlexaForBusinessActivities) UpdateProfile(input *alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error) {
    return a.client.UpdateProfile(input)
}

func (a *AlexaForBusinessActivities) UpdateRoom(input *alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error) {
    return a.client.UpdateRoom(input)
}

func (a *AlexaForBusinessActivities) UpdateSkillGroup(input *alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error) {
    return a.client.UpdateSkillGroup(input)
}
