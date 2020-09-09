package awsclients

import (
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type AlexaForBusinessClient interface {
	ApproveSkill(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error)
	ApproveSkillAsync(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) *AlexaforbusinessApproveSkillResult

	AssociateContactWithAddressBook(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error)
	AssociateContactWithAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) *AlexaforbusinessAssociateContactWithAddressBookResult

	AssociateDeviceWithNetworkProfile(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error)
	AssociateDeviceWithNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) *AlexaforbusinessAssociateDeviceWithNetworkProfileResult

	AssociateDeviceWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error)
	AssociateDeviceWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) *AlexaforbusinessAssociateDeviceWithRoomResult

	AssociateSkillGroupWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error)
	AssociateSkillGroupWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) *AlexaforbusinessAssociateSkillGroupWithRoomResult

	AssociateSkillWithSkillGroup(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error)
	AssociateSkillWithSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) *AlexaforbusinessAssociateSkillWithSkillGroupResult

	AssociateSkillWithUsers(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) (*alexaforbusiness.AssociateSkillWithUsersOutput, error)
	AssociateSkillWithUsersAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) *AlexaforbusinessAssociateSkillWithUsersResult

	CreateAddressBook(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error)
	CreateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) *AlexaforbusinessCreateAddressBookResult

	CreateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error)
	CreateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) *AlexaforbusinessCreateBusinessReportScheduleResult

	CreateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error)
	CreateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) *AlexaforbusinessCreateConferenceProviderResult

	CreateContact(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error)
	CreateContactAsync(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) *AlexaforbusinessCreateContactResult

	CreateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) (*alexaforbusiness.CreateGatewayGroupOutput, error)
	CreateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) *AlexaforbusinessCreateGatewayGroupResult

	CreateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) (*alexaforbusiness.CreateNetworkProfileOutput, error)
	CreateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) *AlexaforbusinessCreateNetworkProfileResult

	CreateProfile(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error)
	CreateProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) *AlexaforbusinessCreateProfileResult

	CreateRoom(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error)
	CreateRoomAsync(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) *AlexaforbusinessCreateRoomResult

	CreateSkillGroup(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error)
	CreateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) *AlexaforbusinessCreateSkillGroupResult

	CreateUser(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) *AlexaforbusinessCreateUserResult

	DeleteAddressBook(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error)
	DeleteAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) *AlexaforbusinessDeleteAddressBookResult

	DeleteBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error)
	DeleteBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) *AlexaforbusinessDeleteBusinessReportScheduleResult

	DeleteConferenceProvider(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error)
	DeleteConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) *AlexaforbusinessDeleteConferenceProviderResult

	DeleteContact(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error)
	DeleteContactAsync(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) *AlexaforbusinessDeleteContactResult

	DeleteDevice(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error)
	DeleteDeviceAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) *AlexaforbusinessDeleteDeviceResult

	DeleteDeviceUsageData(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error)
	DeleteDeviceUsageDataAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) *AlexaforbusinessDeleteDeviceUsageDataResult

	DeleteGatewayGroup(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) (*alexaforbusiness.DeleteGatewayGroupOutput, error)
	DeleteGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) *AlexaforbusinessDeleteGatewayGroupResult

	DeleteNetworkProfile(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) (*alexaforbusiness.DeleteNetworkProfileOutput, error)
	DeleteNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) *AlexaforbusinessDeleteNetworkProfileResult

	DeleteProfile(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error)
	DeleteProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) *AlexaforbusinessDeleteProfileResult

	DeleteRoom(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error)
	DeleteRoomAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) *AlexaforbusinessDeleteRoomResult

	DeleteRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error)
	DeleteRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) *AlexaforbusinessDeleteRoomSkillParameterResult

	DeleteSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error)
	DeleteSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) *AlexaforbusinessDeleteSkillAuthorizationResult

	DeleteSkillGroup(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error)
	DeleteSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) *AlexaforbusinessDeleteSkillGroupResult

	DeleteUser(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) *AlexaforbusinessDeleteUserResult

	DisassociateContactFromAddressBook(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error)
	DisassociateContactFromAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) *AlexaforbusinessDisassociateContactFromAddressBookResult

	DisassociateDeviceFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error)
	DisassociateDeviceFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) *AlexaforbusinessDisassociateDeviceFromRoomResult

	DisassociateSkillFromSkillGroup(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error)
	DisassociateSkillFromSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) *AlexaforbusinessDisassociateSkillFromSkillGroupResult

	DisassociateSkillFromUsers(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error)
	DisassociateSkillFromUsersAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) *AlexaforbusinessDisassociateSkillFromUsersResult

	DisassociateSkillGroupFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error)
	DisassociateSkillGroupFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) *AlexaforbusinessDisassociateSkillGroupFromRoomResult

	ForgetSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error)
	ForgetSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) *AlexaforbusinessForgetSmartHomeAppliancesResult

	GetAddressBook(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error)
	GetAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) *AlexaforbusinessGetAddressBookResult

	GetConferencePreference(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error)
	GetConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) *AlexaforbusinessGetConferencePreferenceResult

	GetConferenceProvider(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error)
	GetConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) *AlexaforbusinessGetConferenceProviderResult

	GetContact(ctx workflow.Context, input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error)
	GetContactAsync(ctx workflow.Context, input *alexaforbusiness.GetContactInput) *AlexaforbusinessGetContactResult

	GetDevice(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error)
	GetDeviceAsync(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) *AlexaforbusinessGetDeviceResult

	GetGateway(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error)
	GetGatewayAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) *AlexaforbusinessGetGatewayResult

	GetGatewayGroup(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error)
	GetGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) *AlexaforbusinessGetGatewayGroupResult

	GetInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error)
	GetInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) *AlexaforbusinessGetInvitationConfigurationResult

	GetNetworkProfile(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error)
	GetNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) *AlexaforbusinessGetNetworkProfileResult

	GetProfile(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error)
	GetProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) *AlexaforbusinessGetProfileResult

	GetRoom(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error)
	GetRoomAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) *AlexaforbusinessGetRoomResult

	GetRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error)
	GetRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) *AlexaforbusinessGetRoomSkillParameterResult

	GetSkillGroup(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error)
	GetSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) *AlexaforbusinessGetSkillGroupResult

	ListBusinessReportSchedules(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error)
	ListBusinessReportSchedulesAsync(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) *AlexaforbusinessListBusinessReportSchedulesResult

	ListConferenceProviders(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error)
	ListConferenceProvidersAsync(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) *AlexaforbusinessListConferenceProvidersResult

	ListDeviceEvents(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error)
	ListDeviceEventsAsync(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) *AlexaforbusinessListDeviceEventsResult

	ListGatewayGroups(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error)
	ListGatewayGroupsAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) *AlexaforbusinessListGatewayGroupsResult

	ListGateways(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error)
	ListGatewaysAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) *AlexaforbusinessListGatewaysResult

	ListSkills(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error)
	ListSkillsAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) *AlexaforbusinessListSkillsResult

	ListSkillsStoreCategories(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error)
	ListSkillsStoreCategoriesAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) *AlexaforbusinessListSkillsStoreCategoriesResult

	ListSkillsStoreSkillsByCategory(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error)
	ListSkillsStoreSkillsByCategoryAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) *AlexaforbusinessListSkillsStoreSkillsByCategoryResult

	ListSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error)
	ListSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) *AlexaforbusinessListSmartHomeAppliancesResult

	ListTags(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) *AlexaforbusinessListTagsResult

	PutConferencePreference(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error)
	PutConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) *AlexaforbusinessPutConferencePreferenceResult

	PutInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) (*alexaforbusiness.PutInvitationConfigurationOutput, error)
	PutInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) *AlexaforbusinessPutInvitationConfigurationResult

	PutRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error)
	PutRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) *AlexaforbusinessPutRoomSkillParameterResult

	PutSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error)
	PutSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) *AlexaforbusinessPutSkillAuthorizationResult

	RegisterAVSDevice(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error)
	RegisterAVSDeviceAsync(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) *AlexaforbusinessRegisterAVSDeviceResult

	RejectSkill(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error)
	RejectSkillAsync(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) *AlexaforbusinessRejectSkillResult

	ResolveRoom(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error)
	ResolveRoomAsync(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) *AlexaforbusinessResolveRoomResult

	RevokeInvitation(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error)
	RevokeInvitationAsync(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) *AlexaforbusinessRevokeInvitationResult

	SearchAddressBooks(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error)
	SearchAddressBooksAsync(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) *AlexaforbusinessSearchAddressBooksResult

	SearchContacts(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error)
	SearchContactsAsync(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) *AlexaforbusinessSearchContactsResult

	SearchDevices(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error)
	SearchDevicesAsync(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) *AlexaforbusinessSearchDevicesResult

	SearchNetworkProfiles(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error)
	SearchNetworkProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) *AlexaforbusinessSearchNetworkProfilesResult

	SearchProfiles(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error)
	SearchProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) *AlexaforbusinessSearchProfilesResult

	SearchRooms(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error)
	SearchRoomsAsync(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) *AlexaforbusinessSearchRoomsResult

	SearchSkillGroups(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error)
	SearchSkillGroupsAsync(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) *AlexaforbusinessSearchSkillGroupsResult

	SearchUsers(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error)
	SearchUsersAsync(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) *AlexaforbusinessSearchUsersResult

	SendAnnouncement(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) (*alexaforbusiness.SendAnnouncementOutput, error)
	SendAnnouncementAsync(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) *AlexaforbusinessSendAnnouncementResult

	SendInvitation(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error)
	SendInvitationAsync(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) *AlexaforbusinessSendInvitationResult

	StartDeviceSync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error)
	StartDeviceSyncAsync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) *AlexaforbusinessStartDeviceSyncResult

	StartSmartHomeApplianceDiscovery(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error)
	StartSmartHomeApplianceDiscoveryAsync(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) *AlexaforbusinessStartSmartHomeApplianceDiscoveryResult

	TagResource(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) *AlexaforbusinessTagResourceResult

	UntagResource(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) *AlexaforbusinessUntagResourceResult

	UpdateAddressBook(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error)
	UpdateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) *AlexaforbusinessUpdateAddressBookResult

	UpdateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error)
	UpdateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) *AlexaforbusinessUpdateBusinessReportScheduleResult

	UpdateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error)
	UpdateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) *AlexaforbusinessUpdateConferenceProviderResult

	UpdateContact(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error)
	UpdateContactAsync(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) *AlexaforbusinessUpdateContactResult

	UpdateDevice(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error)
	UpdateDeviceAsync(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) *AlexaforbusinessUpdateDeviceResult

	UpdateGateway(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) (*alexaforbusiness.UpdateGatewayOutput, error)
	UpdateGatewayAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) *AlexaforbusinessUpdateGatewayResult

	UpdateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) (*alexaforbusiness.UpdateGatewayGroupOutput, error)
	UpdateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) *AlexaforbusinessUpdateGatewayGroupResult

	UpdateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) (*alexaforbusiness.UpdateNetworkProfileOutput, error)
	UpdateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) *AlexaforbusinessUpdateNetworkProfileResult

	UpdateProfile(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error)
	UpdateProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) *AlexaforbusinessUpdateProfileResult

	UpdateRoom(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error)
	UpdateRoomAsync(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) *AlexaforbusinessUpdateRoomResult

	UpdateSkillGroup(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error)
	UpdateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) *AlexaforbusinessUpdateSkillGroupResult
}

type AlexaforbusinessApproveSkillResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessApproveSkillResult) Get(ctx workflow.Context) (*alexaforbusiness.ApproveSkillOutput, error) {
	var output alexaforbusiness.ApproveSkillOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessAssociateContactWithAddressBookResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessAssociateContactWithAddressBookResult) Get(ctx workflow.Context) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error) {
	var output alexaforbusiness.AssociateContactWithAddressBookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessAssociateDeviceWithNetworkProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessAssociateDeviceWithNetworkProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error) {
	var output alexaforbusiness.AssociateDeviceWithNetworkProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessAssociateDeviceWithRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessAssociateDeviceWithRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error) {
	var output alexaforbusiness.AssociateDeviceWithRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessAssociateSkillGroupWithRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessAssociateSkillGroupWithRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error) {
	var output alexaforbusiness.AssociateSkillGroupWithRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessAssociateSkillWithSkillGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessAssociateSkillWithSkillGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error) {
	var output alexaforbusiness.AssociateSkillWithSkillGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessAssociateSkillWithUsersResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessAssociateSkillWithUsersResult) Get(ctx workflow.Context) (*alexaforbusiness.AssociateSkillWithUsersOutput, error) {
	var output alexaforbusiness.AssociateSkillWithUsersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateAddressBookResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateAddressBookResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateAddressBookOutput, error) {
	var output alexaforbusiness.CreateAddressBookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateBusinessReportScheduleResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateBusinessReportScheduleResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error) {
	var output alexaforbusiness.CreateBusinessReportScheduleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateConferenceProviderResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateConferenceProviderResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateConferenceProviderOutput, error) {
	var output alexaforbusiness.CreateConferenceProviderOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateContactResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateContactResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateContactOutput, error) {
	var output alexaforbusiness.CreateContactOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateGatewayGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateGatewayGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateGatewayGroupOutput, error) {
	var output alexaforbusiness.CreateGatewayGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateNetworkProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateNetworkProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateNetworkProfileOutput, error) {
	var output alexaforbusiness.CreateNetworkProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateProfileOutput, error) {
	var output alexaforbusiness.CreateProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateRoomOutput, error) {
	var output alexaforbusiness.CreateRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateSkillGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateSkillGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateSkillGroupOutput, error) {
	var output alexaforbusiness.CreateSkillGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessCreateUserResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessCreateUserResult) Get(ctx workflow.Context) (*alexaforbusiness.CreateUserOutput, error) {
	var output alexaforbusiness.CreateUserOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteAddressBookResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteAddressBookResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteAddressBookOutput, error) {
	var output alexaforbusiness.DeleteAddressBookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteBusinessReportScheduleResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteBusinessReportScheduleResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error) {
	var output alexaforbusiness.DeleteBusinessReportScheduleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteConferenceProviderResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteConferenceProviderResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteConferenceProviderOutput, error) {
	var output alexaforbusiness.DeleteConferenceProviderOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteContactResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteContactResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteContactOutput, error) {
	var output alexaforbusiness.DeleteContactOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteDeviceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteDeviceResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteDeviceOutput, error) {
	var output alexaforbusiness.DeleteDeviceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteDeviceUsageDataResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteDeviceUsageDataResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error) {
	var output alexaforbusiness.DeleteDeviceUsageDataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteGatewayGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteGatewayGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteGatewayGroupOutput, error) {
	var output alexaforbusiness.DeleteGatewayGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteNetworkProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteNetworkProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteNetworkProfileOutput, error) {
	var output alexaforbusiness.DeleteNetworkProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteProfileOutput, error) {
	var output alexaforbusiness.DeleteProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteRoomOutput, error) {
	var output alexaforbusiness.DeleteRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteRoomSkillParameterResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteRoomSkillParameterResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error) {
	var output alexaforbusiness.DeleteRoomSkillParameterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteSkillAuthorizationResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteSkillAuthorizationResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error) {
	var output alexaforbusiness.DeleteSkillAuthorizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteSkillGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteSkillGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteSkillGroupOutput, error) {
	var output alexaforbusiness.DeleteSkillGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDeleteUserResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDeleteUserResult) Get(ctx workflow.Context) (*alexaforbusiness.DeleteUserOutput, error) {
	var output alexaforbusiness.DeleteUserOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDisassociateContactFromAddressBookResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDisassociateContactFromAddressBookResult) Get(ctx workflow.Context) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error) {
	var output alexaforbusiness.DisassociateContactFromAddressBookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDisassociateDeviceFromRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDisassociateDeviceFromRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error) {
	var output alexaforbusiness.DisassociateDeviceFromRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDisassociateSkillFromSkillGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDisassociateSkillFromSkillGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error) {
	var output alexaforbusiness.DisassociateSkillFromSkillGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDisassociateSkillFromUsersResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDisassociateSkillFromUsersResult) Get(ctx workflow.Context) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error) {
	var output alexaforbusiness.DisassociateSkillFromUsersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessDisassociateSkillGroupFromRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessDisassociateSkillGroupFromRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error) {
	var output alexaforbusiness.DisassociateSkillGroupFromRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessForgetSmartHomeAppliancesResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessForgetSmartHomeAppliancesResult) Get(ctx workflow.Context) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error) {
	var output alexaforbusiness.ForgetSmartHomeAppliancesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetAddressBookResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetAddressBookResult) Get(ctx workflow.Context) (*alexaforbusiness.GetAddressBookOutput, error) {
	var output alexaforbusiness.GetAddressBookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetConferencePreferenceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetConferencePreferenceResult) Get(ctx workflow.Context) (*alexaforbusiness.GetConferencePreferenceOutput, error) {
	var output alexaforbusiness.GetConferencePreferenceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetConferenceProviderResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetConferenceProviderResult) Get(ctx workflow.Context) (*alexaforbusiness.GetConferenceProviderOutput, error) {
	var output alexaforbusiness.GetConferenceProviderOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetContactResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetContactResult) Get(ctx workflow.Context) (*alexaforbusiness.GetContactOutput, error) {
	var output alexaforbusiness.GetContactOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetDeviceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetDeviceResult) Get(ctx workflow.Context) (*alexaforbusiness.GetDeviceOutput, error) {
	var output alexaforbusiness.GetDeviceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetGatewayResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetGatewayResult) Get(ctx workflow.Context) (*alexaforbusiness.GetGatewayOutput, error) {
	var output alexaforbusiness.GetGatewayOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetGatewayGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetGatewayGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.GetGatewayGroupOutput, error) {
	var output alexaforbusiness.GetGatewayGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetInvitationConfigurationResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetInvitationConfigurationResult) Get(ctx workflow.Context) (*alexaforbusiness.GetInvitationConfigurationOutput, error) {
	var output alexaforbusiness.GetInvitationConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetNetworkProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetNetworkProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.GetNetworkProfileOutput, error) {
	var output alexaforbusiness.GetNetworkProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.GetProfileOutput, error) {
	var output alexaforbusiness.GetProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.GetRoomOutput, error) {
	var output alexaforbusiness.GetRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetRoomSkillParameterResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetRoomSkillParameterResult) Get(ctx workflow.Context) (*alexaforbusiness.GetRoomSkillParameterOutput, error) {
	var output alexaforbusiness.GetRoomSkillParameterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessGetSkillGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessGetSkillGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.GetSkillGroupOutput, error) {
	var output alexaforbusiness.GetSkillGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListBusinessReportSchedulesResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListBusinessReportSchedulesResult) Get(ctx workflow.Context) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error) {
	var output alexaforbusiness.ListBusinessReportSchedulesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListConferenceProvidersResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListConferenceProvidersResult) Get(ctx workflow.Context) (*alexaforbusiness.ListConferenceProvidersOutput, error) {
	var output alexaforbusiness.ListConferenceProvidersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListDeviceEventsResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListDeviceEventsResult) Get(ctx workflow.Context) (*alexaforbusiness.ListDeviceEventsOutput, error) {
	var output alexaforbusiness.ListDeviceEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListGatewayGroupsResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListGatewayGroupsResult) Get(ctx workflow.Context) (*alexaforbusiness.ListGatewayGroupsOutput, error) {
	var output alexaforbusiness.ListGatewayGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListGatewaysResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListGatewaysResult) Get(ctx workflow.Context) (*alexaforbusiness.ListGatewaysOutput, error) {
	var output alexaforbusiness.ListGatewaysOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListSkillsResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListSkillsResult) Get(ctx workflow.Context) (*alexaforbusiness.ListSkillsOutput, error) {
	var output alexaforbusiness.ListSkillsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListSkillsStoreCategoriesResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListSkillsStoreCategoriesResult) Get(ctx workflow.Context) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error) {
	var output alexaforbusiness.ListSkillsStoreCategoriesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListSkillsStoreSkillsByCategoryResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListSkillsStoreSkillsByCategoryResult) Get(ctx workflow.Context) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error) {
	var output alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListSmartHomeAppliancesResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListSmartHomeAppliancesResult) Get(ctx workflow.Context) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error) {
	var output alexaforbusiness.ListSmartHomeAppliancesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessListTagsResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessListTagsResult) Get(ctx workflow.Context) (*alexaforbusiness.ListTagsOutput, error) {
	var output alexaforbusiness.ListTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessPutConferencePreferenceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessPutConferencePreferenceResult) Get(ctx workflow.Context) (*alexaforbusiness.PutConferencePreferenceOutput, error) {
	var output alexaforbusiness.PutConferencePreferenceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessPutInvitationConfigurationResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessPutInvitationConfigurationResult) Get(ctx workflow.Context) (*alexaforbusiness.PutInvitationConfigurationOutput, error) {
	var output alexaforbusiness.PutInvitationConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessPutRoomSkillParameterResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessPutRoomSkillParameterResult) Get(ctx workflow.Context) (*alexaforbusiness.PutRoomSkillParameterOutput, error) {
	var output alexaforbusiness.PutRoomSkillParameterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessPutSkillAuthorizationResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessPutSkillAuthorizationResult) Get(ctx workflow.Context) (*alexaforbusiness.PutSkillAuthorizationOutput, error) {
	var output alexaforbusiness.PutSkillAuthorizationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessRegisterAVSDeviceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessRegisterAVSDeviceResult) Get(ctx workflow.Context) (*alexaforbusiness.RegisterAVSDeviceOutput, error) {
	var output alexaforbusiness.RegisterAVSDeviceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessRejectSkillResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessRejectSkillResult) Get(ctx workflow.Context) (*alexaforbusiness.RejectSkillOutput, error) {
	var output alexaforbusiness.RejectSkillOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessResolveRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessResolveRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.ResolveRoomOutput, error) {
	var output alexaforbusiness.ResolveRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessRevokeInvitationResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessRevokeInvitationResult) Get(ctx workflow.Context) (*alexaforbusiness.RevokeInvitationOutput, error) {
	var output alexaforbusiness.RevokeInvitationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchAddressBooksResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchAddressBooksResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchAddressBooksOutput, error) {
	var output alexaforbusiness.SearchAddressBooksOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchContactsResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchContactsResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchContactsOutput, error) {
	var output alexaforbusiness.SearchContactsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchDevicesResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchDevicesResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchDevicesOutput, error) {
	var output alexaforbusiness.SearchDevicesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchNetworkProfilesResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchNetworkProfilesResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchNetworkProfilesOutput, error) {
	var output alexaforbusiness.SearchNetworkProfilesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchProfilesResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchProfilesResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchProfilesOutput, error) {
	var output alexaforbusiness.SearchProfilesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchRoomsResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchRoomsResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchRoomsOutput, error) {
	var output alexaforbusiness.SearchRoomsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchSkillGroupsResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchSkillGroupsResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchSkillGroupsOutput, error) {
	var output alexaforbusiness.SearchSkillGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSearchUsersResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSearchUsersResult) Get(ctx workflow.Context) (*alexaforbusiness.SearchUsersOutput, error) {
	var output alexaforbusiness.SearchUsersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSendAnnouncementResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSendAnnouncementResult) Get(ctx workflow.Context) (*alexaforbusiness.SendAnnouncementOutput, error) {
	var output alexaforbusiness.SendAnnouncementOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessSendInvitationResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessSendInvitationResult) Get(ctx workflow.Context) (*alexaforbusiness.SendInvitationOutput, error) {
	var output alexaforbusiness.SendInvitationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessStartDeviceSyncResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessStartDeviceSyncResult) Get(ctx workflow.Context) (*alexaforbusiness.StartDeviceSyncOutput, error) {
	var output alexaforbusiness.StartDeviceSyncOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessStartSmartHomeApplianceDiscoveryResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessStartSmartHomeApplianceDiscoveryResult) Get(ctx workflow.Context) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error) {
	var output alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessTagResourceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessTagResourceResult) Get(ctx workflow.Context) (*alexaforbusiness.TagResourceOutput, error) {
	var output alexaforbusiness.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUntagResourceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUntagResourceResult) Get(ctx workflow.Context) (*alexaforbusiness.UntagResourceOutput, error) {
	var output alexaforbusiness.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateAddressBookResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateAddressBookResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateAddressBookOutput, error) {
	var output alexaforbusiness.UpdateAddressBookOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateBusinessReportScheduleResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateBusinessReportScheduleResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error) {
	var output alexaforbusiness.UpdateBusinessReportScheduleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateConferenceProviderResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateConferenceProviderResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateConferenceProviderOutput, error) {
	var output alexaforbusiness.UpdateConferenceProviderOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateContactResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateContactResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateContactOutput, error) {
	var output alexaforbusiness.UpdateContactOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateDeviceResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateDeviceResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateDeviceOutput, error) {
	var output alexaforbusiness.UpdateDeviceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateGatewayResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateGatewayResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateGatewayOutput, error) {
	var output alexaforbusiness.UpdateGatewayOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateGatewayGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateGatewayGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateGatewayGroupOutput, error) {
	var output alexaforbusiness.UpdateGatewayGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateNetworkProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateNetworkProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateNetworkProfileOutput, error) {
	var output alexaforbusiness.UpdateNetworkProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateProfileResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateProfileResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateProfileOutput, error) {
	var output alexaforbusiness.UpdateProfileOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateRoomResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateRoomResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateRoomOutput, error) {
	var output alexaforbusiness.UpdateRoomOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaforbusinessUpdateSkillGroupResult struct {
	Result workflow.Future
}

func (r *AlexaforbusinessUpdateSkillGroupResult) Get(ctx workflow.Context) (*alexaforbusiness.UpdateSkillGroupOutput, error) {
	var output alexaforbusiness.UpdateSkillGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type AlexaForBusinessStub struct {
	activities awsactivities.AlexaForBusinessActivities
}

func NewAlexaForBusinessStub() AlexaForBusinessClient {
	return &AlexaForBusinessStub{}
}

func (a *AlexaForBusinessStub) ApproveSkill(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error) {
	var output alexaforbusiness.ApproveSkillOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ApproveSkill, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ApproveSkillAsync(ctx workflow.Context, input *alexaforbusiness.ApproveSkillInput) *AlexaforbusinessApproveSkillResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ApproveSkill, input)
	return &AlexaforbusinessApproveSkillResult{Result: future}
}

func (a *AlexaForBusinessStub) AssociateContactWithAddressBook(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error) {
	var output alexaforbusiness.AssociateContactWithAddressBookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateContactWithAddressBook, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) AssociateContactWithAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) *AlexaforbusinessAssociateContactWithAddressBookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateContactWithAddressBook, input)
	return &AlexaforbusinessAssociateContactWithAddressBookResult{Result: future}
}

func (a *AlexaForBusinessStub) AssociateDeviceWithNetworkProfile(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error) {
	var output alexaforbusiness.AssociateDeviceWithNetworkProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateDeviceWithNetworkProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) AssociateDeviceWithNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) *AlexaforbusinessAssociateDeviceWithNetworkProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateDeviceWithNetworkProfile, input)
	return &AlexaforbusinessAssociateDeviceWithNetworkProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) AssociateDeviceWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error) {
	var output alexaforbusiness.AssociateDeviceWithRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateDeviceWithRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) AssociateDeviceWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) *AlexaforbusinessAssociateDeviceWithRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateDeviceWithRoom, input)
	return &AlexaforbusinessAssociateDeviceWithRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) AssociateSkillGroupWithRoom(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error) {
	var output alexaforbusiness.AssociateSkillGroupWithRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateSkillGroupWithRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) AssociateSkillGroupWithRoomAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) *AlexaforbusinessAssociateSkillGroupWithRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateSkillGroupWithRoom, input)
	return &AlexaforbusinessAssociateSkillGroupWithRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) AssociateSkillWithSkillGroup(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error) {
	var output alexaforbusiness.AssociateSkillWithSkillGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateSkillWithSkillGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) AssociateSkillWithSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) *AlexaforbusinessAssociateSkillWithSkillGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateSkillWithSkillGroup, input)
	return &AlexaforbusinessAssociateSkillWithSkillGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) AssociateSkillWithUsers(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) (*alexaforbusiness.AssociateSkillWithUsersOutput, error) {
	var output alexaforbusiness.AssociateSkillWithUsersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateSkillWithUsers, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) AssociateSkillWithUsersAsync(ctx workflow.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) *AlexaforbusinessAssociateSkillWithUsersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateSkillWithUsers, input)
	return &AlexaforbusinessAssociateSkillWithUsersResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateAddressBook(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error) {
	var output alexaforbusiness.CreateAddressBookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateAddressBook, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.CreateAddressBookInput) *AlexaforbusinessCreateAddressBookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateAddressBook, input)
	return &AlexaforbusinessCreateAddressBookResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error) {
	var output alexaforbusiness.CreateBusinessReportScheduleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateBusinessReportSchedule, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) *AlexaforbusinessCreateBusinessReportScheduleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateBusinessReportSchedule, input)
	return &AlexaforbusinessCreateBusinessReportScheduleResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error) {
	var output alexaforbusiness.CreateConferenceProviderOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateConferenceProvider, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.CreateConferenceProviderInput) *AlexaforbusinessCreateConferenceProviderResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateConferenceProvider, input)
	return &AlexaforbusinessCreateConferenceProviderResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateContact(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error) {
	var output alexaforbusiness.CreateContactOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateContact, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateContactAsync(ctx workflow.Context, input *alexaforbusiness.CreateContactInput) *AlexaforbusinessCreateContactResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateContact, input)
	return &AlexaforbusinessCreateContactResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) (*alexaforbusiness.CreateGatewayGroupOutput, error) {
	var output alexaforbusiness.CreateGatewayGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateGatewayGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateGatewayGroupInput) *AlexaforbusinessCreateGatewayGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateGatewayGroup, input)
	return &AlexaforbusinessCreateGatewayGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) (*alexaforbusiness.CreateNetworkProfileOutput, error) {
	var output alexaforbusiness.CreateNetworkProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateNetworkProfileInput) *AlexaforbusinessCreateNetworkProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateNetworkProfile, input)
	return &AlexaforbusinessCreateNetworkProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateProfile(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error) {
	var output alexaforbusiness.CreateProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateProfileAsync(ctx workflow.Context, input *alexaforbusiness.CreateProfileInput) *AlexaforbusinessCreateProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateProfile, input)
	return &AlexaforbusinessCreateProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateRoom(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error) {
	var output alexaforbusiness.CreateRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateRoomAsync(ctx workflow.Context, input *alexaforbusiness.CreateRoomInput) *AlexaforbusinessCreateRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRoom, input)
	return &AlexaforbusinessCreateRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateSkillGroup(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error) {
	var output alexaforbusiness.CreateSkillGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateSkillGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.CreateSkillGroupInput) *AlexaforbusinessCreateSkillGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateSkillGroup, input)
	return &AlexaforbusinessCreateSkillGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) CreateUser(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error) {
	var output alexaforbusiness.CreateUserOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) CreateUserAsync(ctx workflow.Context, input *alexaforbusiness.CreateUserInput) *AlexaforbusinessCreateUserResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input)
	return &AlexaforbusinessCreateUserResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteAddressBook(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error) {
	var output alexaforbusiness.DeleteAddressBookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAddressBook, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DeleteAddressBookInput) *AlexaforbusinessDeleteAddressBookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAddressBook, input)
	return &AlexaforbusinessDeleteAddressBookResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error) {
	var output alexaforbusiness.DeleteBusinessReportScheduleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBusinessReportSchedule, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) *AlexaforbusinessDeleteBusinessReportScheduleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBusinessReportSchedule, input)
	return &AlexaforbusinessDeleteBusinessReportScheduleResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteConferenceProvider(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error) {
	var output alexaforbusiness.DeleteConferenceProviderOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteConferenceProvider, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.DeleteConferenceProviderInput) *AlexaforbusinessDeleteConferenceProviderResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteConferenceProvider, input)
	return &AlexaforbusinessDeleteConferenceProviderResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteContact(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error) {
	var output alexaforbusiness.DeleteContactOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteContact, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteContactAsync(ctx workflow.Context, input *alexaforbusiness.DeleteContactInput) *AlexaforbusinessDeleteContactResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteContact, input)
	return &AlexaforbusinessDeleteContactResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteDevice(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error) {
	var output alexaforbusiness.DeleteDeviceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDevice, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteDeviceAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceInput) *AlexaforbusinessDeleteDeviceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDevice, input)
	return &AlexaforbusinessDeleteDeviceResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteDeviceUsageData(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error) {
	var output alexaforbusiness.DeleteDeviceUsageDataOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeviceUsageData, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteDeviceUsageDataAsync(ctx workflow.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) *AlexaforbusinessDeleteDeviceUsageDataResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDeviceUsageData, input)
	return &AlexaforbusinessDeleteDeviceUsageDataResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteGatewayGroup(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) (*alexaforbusiness.DeleteGatewayGroupOutput, error) {
	var output alexaforbusiness.DeleteGatewayGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteGatewayGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteGatewayGroupInput) *AlexaforbusinessDeleteGatewayGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteGatewayGroup, input)
	return &AlexaforbusinessDeleteGatewayGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteNetworkProfile(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) (*alexaforbusiness.DeleteNetworkProfileOutput, error) {
	var output alexaforbusiness.DeleteNetworkProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteNetworkProfileInput) *AlexaforbusinessDeleteNetworkProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteNetworkProfile, input)
	return &AlexaforbusinessDeleteNetworkProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteProfile(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error) {
	var output alexaforbusiness.DeleteProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteProfileAsync(ctx workflow.Context, input *alexaforbusiness.DeleteProfileInput) *AlexaforbusinessDeleteProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteProfile, input)
	return &AlexaforbusinessDeleteProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteRoom(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error) {
	var output alexaforbusiness.DeleteRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteRoomAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomInput) *AlexaforbusinessDeleteRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRoom, input)
	return &AlexaforbusinessDeleteRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error) {
	var output alexaforbusiness.DeleteRoomSkillParameterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRoomSkillParameter, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) *AlexaforbusinessDeleteRoomSkillParameterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRoomSkillParameter, input)
	return &AlexaforbusinessDeleteRoomSkillParameterResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error) {
	var output alexaforbusiness.DeleteSkillAuthorizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSkillAuthorization, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) *AlexaforbusinessDeleteSkillAuthorizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSkillAuthorization, input)
	return &AlexaforbusinessDeleteSkillAuthorizationResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteSkillGroup(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error) {
	var output alexaforbusiness.DeleteSkillGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSkillGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DeleteSkillGroupInput) *AlexaforbusinessDeleteSkillGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSkillGroup, input)
	return &AlexaforbusinessDeleteSkillGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) DeleteUser(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error) {
	var output alexaforbusiness.DeleteUserOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DeleteUserAsync(ctx workflow.Context, input *alexaforbusiness.DeleteUserInput) *AlexaforbusinessDeleteUserResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input)
	return &AlexaforbusinessDeleteUserResult{Result: future}
}

func (a *AlexaForBusinessStub) DisassociateContactFromAddressBook(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error) {
	var output alexaforbusiness.DisassociateContactFromAddressBookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateContactFromAddressBook, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DisassociateContactFromAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) *AlexaforbusinessDisassociateContactFromAddressBookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateContactFromAddressBook, input)
	return &AlexaforbusinessDisassociateContactFromAddressBookResult{Result: future}
}

func (a *AlexaForBusinessStub) DisassociateDeviceFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error) {
	var output alexaforbusiness.DisassociateDeviceFromRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateDeviceFromRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DisassociateDeviceFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) *AlexaforbusinessDisassociateDeviceFromRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateDeviceFromRoom, input)
	return &AlexaforbusinessDisassociateDeviceFromRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) DisassociateSkillFromSkillGroup(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error) {
	var output alexaforbusiness.DisassociateSkillFromSkillGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateSkillFromSkillGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DisassociateSkillFromSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) *AlexaforbusinessDisassociateSkillFromSkillGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateSkillFromSkillGroup, input)
	return &AlexaforbusinessDisassociateSkillFromSkillGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) DisassociateSkillFromUsers(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error) {
	var output alexaforbusiness.DisassociateSkillFromUsersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateSkillFromUsers, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DisassociateSkillFromUsersAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) *AlexaforbusinessDisassociateSkillFromUsersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateSkillFromUsers, input)
	return &AlexaforbusinessDisassociateSkillFromUsersResult{Result: future}
}

func (a *AlexaForBusinessStub) DisassociateSkillGroupFromRoom(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error) {
	var output alexaforbusiness.DisassociateSkillGroupFromRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateSkillGroupFromRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) DisassociateSkillGroupFromRoomAsync(ctx workflow.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) *AlexaforbusinessDisassociateSkillGroupFromRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateSkillGroupFromRoom, input)
	return &AlexaforbusinessDisassociateSkillGroupFromRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) ForgetSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error) {
	var output alexaforbusiness.ForgetSmartHomeAppliancesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ForgetSmartHomeAppliances, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ForgetSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) *AlexaforbusinessForgetSmartHomeAppliancesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ForgetSmartHomeAppliances, input)
	return &AlexaforbusinessForgetSmartHomeAppliancesResult{Result: future}
}

func (a *AlexaForBusinessStub) GetAddressBook(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error) {
	var output alexaforbusiness.GetAddressBookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAddressBook, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.GetAddressBookInput) *AlexaforbusinessGetAddressBookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAddressBook, input)
	return &AlexaforbusinessGetAddressBookResult{Result: future}
}

func (a *AlexaForBusinessStub) GetConferencePreference(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error) {
	var output alexaforbusiness.GetConferencePreferenceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetConferencePreference, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.GetConferencePreferenceInput) *AlexaforbusinessGetConferencePreferenceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetConferencePreference, input)
	return &AlexaforbusinessGetConferencePreferenceResult{Result: future}
}

func (a *AlexaForBusinessStub) GetConferenceProvider(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error) {
	var output alexaforbusiness.GetConferenceProviderOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetConferenceProvider, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.GetConferenceProviderInput) *AlexaforbusinessGetConferenceProviderResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetConferenceProvider, input)
	return &AlexaforbusinessGetConferenceProviderResult{Result: future}
}

func (a *AlexaForBusinessStub) GetContact(ctx workflow.Context, input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error) {
	var output alexaforbusiness.GetContactOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetContact, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetContactAsync(ctx workflow.Context, input *alexaforbusiness.GetContactInput) *AlexaforbusinessGetContactResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetContact, input)
	return &AlexaforbusinessGetContactResult{Result: future}
}

func (a *AlexaForBusinessStub) GetDevice(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error) {
	var output alexaforbusiness.GetDeviceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDevice, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetDeviceAsync(ctx workflow.Context, input *alexaforbusiness.GetDeviceInput) *AlexaforbusinessGetDeviceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDevice, input)
	return &AlexaforbusinessGetDeviceResult{Result: future}
}

func (a *AlexaForBusinessStub) GetGateway(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error) {
	var output alexaforbusiness.GetGatewayOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetGateway, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetGatewayAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayInput) *AlexaforbusinessGetGatewayResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetGateway, input)
	return &AlexaforbusinessGetGatewayResult{Result: future}
}

func (a *AlexaForBusinessStub) GetGatewayGroup(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error) {
	var output alexaforbusiness.GetGatewayGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetGatewayGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetGatewayGroupInput) *AlexaforbusinessGetGatewayGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetGatewayGroup, input)
	return &AlexaforbusinessGetGatewayGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) GetInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error) {
	var output alexaforbusiness.GetInvitationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInvitationConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.GetInvitationConfigurationInput) *AlexaforbusinessGetInvitationConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInvitationConfiguration, input)
	return &AlexaforbusinessGetInvitationConfigurationResult{Result: future}
}

func (a *AlexaForBusinessStub) GetNetworkProfile(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error) {
	var output alexaforbusiness.GetNetworkProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetNetworkProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetNetworkProfileInput) *AlexaforbusinessGetNetworkProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetNetworkProfile, input)
	return &AlexaforbusinessGetNetworkProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) GetProfile(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error) {
	var output alexaforbusiness.GetProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetProfileAsync(ctx workflow.Context, input *alexaforbusiness.GetProfileInput) *AlexaforbusinessGetProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetProfile, input)
	return &AlexaforbusinessGetProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) GetRoom(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error) {
	var output alexaforbusiness.GetRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetRoomAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomInput) *AlexaforbusinessGetRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRoom, input)
	return &AlexaforbusinessGetRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) GetRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error) {
	var output alexaforbusiness.GetRoomSkillParameterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRoomSkillParameter, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.GetRoomSkillParameterInput) *AlexaforbusinessGetRoomSkillParameterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRoomSkillParameter, input)
	return &AlexaforbusinessGetRoomSkillParameterResult{Result: future}
}

func (a *AlexaForBusinessStub) GetSkillGroup(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error) {
	var output alexaforbusiness.GetSkillGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSkillGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) GetSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.GetSkillGroupInput) *AlexaforbusinessGetSkillGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSkillGroup, input)
	return &AlexaforbusinessGetSkillGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) ListBusinessReportSchedules(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error) {
	var output alexaforbusiness.ListBusinessReportSchedulesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBusinessReportSchedules, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListBusinessReportSchedulesAsync(ctx workflow.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) *AlexaforbusinessListBusinessReportSchedulesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBusinessReportSchedules, input)
	return &AlexaforbusinessListBusinessReportSchedulesResult{Result: future}
}

func (a *AlexaForBusinessStub) ListConferenceProviders(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error) {
	var output alexaforbusiness.ListConferenceProvidersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListConferenceProviders, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListConferenceProvidersAsync(ctx workflow.Context, input *alexaforbusiness.ListConferenceProvidersInput) *AlexaforbusinessListConferenceProvidersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListConferenceProviders, input)
	return &AlexaforbusinessListConferenceProvidersResult{Result: future}
}

func (a *AlexaForBusinessStub) ListDeviceEvents(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error) {
	var output alexaforbusiness.ListDeviceEventsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListDeviceEvents, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListDeviceEventsAsync(ctx workflow.Context, input *alexaforbusiness.ListDeviceEventsInput) *AlexaforbusinessListDeviceEventsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListDeviceEvents, input)
	return &AlexaforbusinessListDeviceEventsResult{Result: future}
}

func (a *AlexaForBusinessStub) ListGatewayGroups(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error) {
	var output alexaforbusiness.ListGatewayGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListGatewayGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListGatewayGroupsAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewayGroupsInput) *AlexaforbusinessListGatewayGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListGatewayGroups, input)
	return &AlexaforbusinessListGatewayGroupsResult{Result: future}
}

func (a *AlexaForBusinessStub) ListGateways(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error) {
	var output alexaforbusiness.ListGatewaysOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListGateways, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListGatewaysAsync(ctx workflow.Context, input *alexaforbusiness.ListGatewaysInput) *AlexaforbusinessListGatewaysResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListGateways, input)
	return &AlexaforbusinessListGatewaysResult{Result: future}
}

func (a *AlexaForBusinessStub) ListSkills(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error) {
	var output alexaforbusiness.ListSkillsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSkills, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListSkillsAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsInput) *AlexaforbusinessListSkillsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSkills, input)
	return &AlexaforbusinessListSkillsResult{Result: future}
}

func (a *AlexaForBusinessStub) ListSkillsStoreCategories(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error) {
	var output alexaforbusiness.ListSkillsStoreCategoriesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSkillsStoreCategories, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListSkillsStoreCategoriesAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) *AlexaforbusinessListSkillsStoreCategoriesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSkillsStoreCategories, input)
	return &AlexaforbusinessListSkillsStoreCategoriesResult{Result: future}
}

func (a *AlexaForBusinessStub) ListSkillsStoreSkillsByCategory(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error) {
	var output alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSkillsStoreSkillsByCategory, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListSkillsStoreSkillsByCategoryAsync(ctx workflow.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) *AlexaforbusinessListSkillsStoreSkillsByCategoryResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSkillsStoreSkillsByCategory, input)
	return &AlexaforbusinessListSkillsStoreSkillsByCategoryResult{Result: future}
}

func (a *AlexaForBusinessStub) ListSmartHomeAppliances(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error) {
	var output alexaforbusiness.ListSmartHomeAppliancesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListSmartHomeAppliances, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListSmartHomeAppliancesAsync(ctx workflow.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) *AlexaforbusinessListSmartHomeAppliancesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListSmartHomeAppliances, input)
	return &AlexaforbusinessListSmartHomeAppliancesResult{Result: future}
}

func (a *AlexaForBusinessStub) ListTags(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error) {
	var output alexaforbusiness.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ListTagsAsync(ctx workflow.Context, input *alexaforbusiness.ListTagsInput) *AlexaforbusinessListTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
	return &AlexaforbusinessListTagsResult{Result: future}
}

func (a *AlexaForBusinessStub) PutConferencePreference(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error) {
	var output alexaforbusiness.PutConferencePreferenceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutConferencePreference, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) PutConferencePreferenceAsync(ctx workflow.Context, input *alexaforbusiness.PutConferencePreferenceInput) *AlexaforbusinessPutConferencePreferenceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutConferencePreference, input)
	return &AlexaforbusinessPutConferencePreferenceResult{Result: future}
}

func (a *AlexaForBusinessStub) PutInvitationConfiguration(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) (*alexaforbusiness.PutInvitationConfigurationOutput, error) {
	var output alexaforbusiness.PutInvitationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutInvitationConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) PutInvitationConfigurationAsync(ctx workflow.Context, input *alexaforbusiness.PutInvitationConfigurationInput) *AlexaforbusinessPutInvitationConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutInvitationConfiguration, input)
	return &AlexaforbusinessPutInvitationConfigurationResult{Result: future}
}

func (a *AlexaForBusinessStub) PutRoomSkillParameter(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error) {
	var output alexaforbusiness.PutRoomSkillParameterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutRoomSkillParameter, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) PutRoomSkillParameterAsync(ctx workflow.Context, input *alexaforbusiness.PutRoomSkillParameterInput) *AlexaforbusinessPutRoomSkillParameterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutRoomSkillParameter, input)
	return &AlexaforbusinessPutRoomSkillParameterResult{Result: future}
}

func (a *AlexaForBusinessStub) PutSkillAuthorization(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error) {
	var output alexaforbusiness.PutSkillAuthorizationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutSkillAuthorization, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) PutSkillAuthorizationAsync(ctx workflow.Context, input *alexaforbusiness.PutSkillAuthorizationInput) *AlexaforbusinessPutSkillAuthorizationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutSkillAuthorization, input)
	return &AlexaforbusinessPutSkillAuthorizationResult{Result: future}
}

func (a *AlexaForBusinessStub) RegisterAVSDevice(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error) {
	var output alexaforbusiness.RegisterAVSDeviceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RegisterAVSDevice, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) RegisterAVSDeviceAsync(ctx workflow.Context, input *alexaforbusiness.RegisterAVSDeviceInput) *AlexaforbusinessRegisterAVSDeviceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RegisterAVSDevice, input)
	return &AlexaforbusinessRegisterAVSDeviceResult{Result: future}
}

func (a *AlexaForBusinessStub) RejectSkill(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error) {
	var output alexaforbusiness.RejectSkillOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RejectSkill, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) RejectSkillAsync(ctx workflow.Context, input *alexaforbusiness.RejectSkillInput) *AlexaforbusinessRejectSkillResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RejectSkill, input)
	return &AlexaforbusinessRejectSkillResult{Result: future}
}

func (a *AlexaForBusinessStub) ResolveRoom(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error) {
	var output alexaforbusiness.ResolveRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ResolveRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) ResolveRoomAsync(ctx workflow.Context, input *alexaforbusiness.ResolveRoomInput) *AlexaforbusinessResolveRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ResolveRoom, input)
	return &AlexaforbusinessResolveRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) RevokeInvitation(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error) {
	var output alexaforbusiness.RevokeInvitationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RevokeInvitation, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) RevokeInvitationAsync(ctx workflow.Context, input *alexaforbusiness.RevokeInvitationInput) *AlexaforbusinessRevokeInvitationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RevokeInvitation, input)
	return &AlexaforbusinessRevokeInvitationResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchAddressBooks(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error) {
	var output alexaforbusiness.SearchAddressBooksOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchAddressBooks, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchAddressBooksAsync(ctx workflow.Context, input *alexaforbusiness.SearchAddressBooksInput) *AlexaforbusinessSearchAddressBooksResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchAddressBooks, input)
	return &AlexaforbusinessSearchAddressBooksResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchContacts(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error) {
	var output alexaforbusiness.SearchContactsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchContacts, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchContactsAsync(ctx workflow.Context, input *alexaforbusiness.SearchContactsInput) *AlexaforbusinessSearchContactsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchContacts, input)
	return &AlexaforbusinessSearchContactsResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchDevices(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error) {
	var output alexaforbusiness.SearchDevicesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchDevices, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchDevicesAsync(ctx workflow.Context, input *alexaforbusiness.SearchDevicesInput) *AlexaforbusinessSearchDevicesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchDevices, input)
	return &AlexaforbusinessSearchDevicesResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchNetworkProfiles(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error) {
	var output alexaforbusiness.SearchNetworkProfilesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchNetworkProfiles, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchNetworkProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchNetworkProfilesInput) *AlexaforbusinessSearchNetworkProfilesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchNetworkProfiles, input)
	return &AlexaforbusinessSearchNetworkProfilesResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchProfiles(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error) {
	var output alexaforbusiness.SearchProfilesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchProfiles, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchProfilesAsync(ctx workflow.Context, input *alexaforbusiness.SearchProfilesInput) *AlexaforbusinessSearchProfilesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchProfiles, input)
	return &AlexaforbusinessSearchProfilesResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchRooms(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error) {
	var output alexaforbusiness.SearchRoomsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchRooms, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchRoomsAsync(ctx workflow.Context, input *alexaforbusiness.SearchRoomsInput) *AlexaforbusinessSearchRoomsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchRooms, input)
	return &AlexaforbusinessSearchRoomsResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchSkillGroups(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error) {
	var output alexaforbusiness.SearchSkillGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchSkillGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchSkillGroupsAsync(ctx workflow.Context, input *alexaforbusiness.SearchSkillGroupsInput) *AlexaforbusinessSearchSkillGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchSkillGroups, input)
	return &AlexaforbusinessSearchSkillGroupsResult{Result: future}
}

func (a *AlexaForBusinessStub) SearchUsers(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error) {
	var output alexaforbusiness.SearchUsersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchUsers, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SearchUsersAsync(ctx workflow.Context, input *alexaforbusiness.SearchUsersInput) *AlexaforbusinessSearchUsersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchUsers, input)
	return &AlexaforbusinessSearchUsersResult{Result: future}
}

func (a *AlexaForBusinessStub) SendAnnouncement(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) (*alexaforbusiness.SendAnnouncementOutput, error) {
	var output alexaforbusiness.SendAnnouncementOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SendAnnouncement, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SendAnnouncementAsync(ctx workflow.Context, input *alexaforbusiness.SendAnnouncementInput) *AlexaforbusinessSendAnnouncementResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SendAnnouncement, input)
	return &AlexaforbusinessSendAnnouncementResult{Result: future}
}

func (a *AlexaForBusinessStub) SendInvitation(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error) {
	var output alexaforbusiness.SendInvitationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SendInvitation, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) SendInvitationAsync(ctx workflow.Context, input *alexaforbusiness.SendInvitationInput) *AlexaforbusinessSendInvitationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SendInvitation, input)
	return &AlexaforbusinessSendInvitationResult{Result: future}
}

func (a *AlexaForBusinessStub) StartDeviceSync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error) {
	var output alexaforbusiness.StartDeviceSyncOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartDeviceSync, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) StartDeviceSyncAsync(ctx workflow.Context, input *alexaforbusiness.StartDeviceSyncInput) *AlexaforbusinessStartDeviceSyncResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartDeviceSync, input)
	return &AlexaforbusinessStartDeviceSyncResult{Result: future}
}

func (a *AlexaForBusinessStub) StartSmartHomeApplianceDiscovery(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error) {
	var output alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartSmartHomeApplianceDiscovery, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) StartSmartHomeApplianceDiscoveryAsync(ctx workflow.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) *AlexaforbusinessStartSmartHomeApplianceDiscoveryResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartSmartHomeApplianceDiscovery, input)
	return &AlexaforbusinessStartSmartHomeApplianceDiscoveryResult{Result: future}
}

func (a *AlexaForBusinessStub) TagResource(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error) {
	var output alexaforbusiness.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) TagResourceAsync(ctx workflow.Context, input *alexaforbusiness.TagResourceInput) *AlexaforbusinessTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &AlexaforbusinessTagResourceResult{Result: future}
}

func (a *AlexaForBusinessStub) UntagResource(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error) {
	var output alexaforbusiness.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UntagResourceAsync(ctx workflow.Context, input *alexaforbusiness.UntagResourceInput) *AlexaforbusinessUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &AlexaforbusinessUntagResourceResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateAddressBook(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error) {
	var output alexaforbusiness.UpdateAddressBookOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAddressBook, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateAddressBookAsync(ctx workflow.Context, input *alexaforbusiness.UpdateAddressBookInput) *AlexaforbusinessUpdateAddressBookResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAddressBook, input)
	return &AlexaforbusinessUpdateAddressBookResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateBusinessReportSchedule(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error) {
	var output alexaforbusiness.UpdateBusinessReportScheduleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateBusinessReportSchedule, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateBusinessReportScheduleAsync(ctx workflow.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) *AlexaforbusinessUpdateBusinessReportScheduleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateBusinessReportSchedule, input)
	return &AlexaforbusinessUpdateBusinessReportScheduleResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateConferenceProvider(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error) {
	var output alexaforbusiness.UpdateConferenceProviderOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateConferenceProvider, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateConferenceProviderAsync(ctx workflow.Context, input *alexaforbusiness.UpdateConferenceProviderInput) *AlexaforbusinessUpdateConferenceProviderResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateConferenceProvider, input)
	return &AlexaforbusinessUpdateConferenceProviderResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateContact(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error) {
	var output alexaforbusiness.UpdateContactOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateContact, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateContactAsync(ctx workflow.Context, input *alexaforbusiness.UpdateContactInput) *AlexaforbusinessUpdateContactResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateContact, input)
	return &AlexaforbusinessUpdateContactResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateDevice(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error) {
	var output alexaforbusiness.UpdateDeviceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDevice, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateDeviceAsync(ctx workflow.Context, input *alexaforbusiness.UpdateDeviceInput) *AlexaforbusinessUpdateDeviceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDevice, input)
	return &AlexaforbusinessUpdateDeviceResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateGateway(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) (*alexaforbusiness.UpdateGatewayOutput, error) {
	var output alexaforbusiness.UpdateGatewayOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateGateway, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateGatewayAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayInput) *AlexaforbusinessUpdateGatewayResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateGateway, input)
	return &AlexaforbusinessUpdateGatewayResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateGatewayGroup(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) (*alexaforbusiness.UpdateGatewayGroupOutput, error) {
	var output alexaforbusiness.UpdateGatewayGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateGatewayGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateGatewayGroupInput) *AlexaforbusinessUpdateGatewayGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayGroup, input)
	return &AlexaforbusinessUpdateGatewayGroupResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateNetworkProfile(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) (*alexaforbusiness.UpdateNetworkProfileOutput, error) {
	var output alexaforbusiness.UpdateNetworkProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateNetworkProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateNetworkProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateNetworkProfileInput) *AlexaforbusinessUpdateNetworkProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateNetworkProfile, input)
	return &AlexaforbusinessUpdateNetworkProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateProfile(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error) {
	var output alexaforbusiness.UpdateProfileOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateProfile, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateProfileAsync(ctx workflow.Context, input *alexaforbusiness.UpdateProfileInput) *AlexaforbusinessUpdateProfileResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateProfile, input)
	return &AlexaforbusinessUpdateProfileResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateRoom(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error) {
	var output alexaforbusiness.UpdateRoomOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRoom, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateRoomAsync(ctx workflow.Context, input *alexaforbusiness.UpdateRoomInput) *AlexaforbusinessUpdateRoomResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRoom, input)
	return &AlexaforbusinessUpdateRoomResult{Result: future}
}

func (a *AlexaForBusinessStub) UpdateSkillGroup(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error) {
	var output alexaforbusiness.UpdateSkillGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateSkillGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *AlexaForBusinessStub) UpdateSkillGroupAsync(ctx workflow.Context, input *alexaforbusiness.UpdateSkillGroupInput) *AlexaforbusinessUpdateSkillGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateSkillGroup, input)
	return &AlexaforbusinessUpdateSkillGroupResult{Result: future}
}
