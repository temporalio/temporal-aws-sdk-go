
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/chime"
	"github.com/aws/aws-sdk-go/service/chime/chimeiface"
)

type ChimeActivities struct {
	client chimeiface.ChimeAPI
}

func NewChimeActivities(client chimeiface.ChimeAPI) *ChimeActivities {
    return &ChimeActivities{client: client}
}

func (a *ChimeActivities) AssociatePhoneNumberWithUser(input *chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error) {
    return a.client.AssociatePhoneNumberWithUser(input)
}

func (a *ChimeActivities) AssociatePhoneNumbersWithVoiceConnector(input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error) {
    return a.client.AssociatePhoneNumbersWithVoiceConnector(input)
}

func (a *ChimeActivities) AssociatePhoneNumbersWithVoiceConnectorGroup(input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) {
    return a.client.AssociatePhoneNumbersWithVoiceConnectorGroup(input)
}

func (a *ChimeActivities) AssociateSigninDelegateGroupsWithAccount(input *chime.AssociateSigninDelegateGroupsWithAccountInput) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error) {
    return a.client.AssociateSigninDelegateGroupsWithAccount(input)
}

func (a *ChimeActivities) BatchCreateAttendee(input *chime.BatchCreateAttendeeInput) (*chime.BatchCreateAttendeeOutput, error) {
    return a.client.BatchCreateAttendee(input)
}

func (a *ChimeActivities) BatchCreateRoomMembership(input *chime.BatchCreateRoomMembershipInput) (*chime.BatchCreateRoomMembershipOutput, error) {
    return a.client.BatchCreateRoomMembership(input)
}

func (a *ChimeActivities) BatchDeletePhoneNumber(input *chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error) {
    return a.client.BatchDeletePhoneNumber(input)
}

func (a *ChimeActivities) BatchSuspendUser(input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error) {
    return a.client.BatchSuspendUser(input)
}

func (a *ChimeActivities) BatchUnsuspendUser(input *chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error) {
    return a.client.BatchUnsuspendUser(input)
}

func (a *ChimeActivities) BatchUpdatePhoneNumber(input *chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error) {
    return a.client.BatchUpdatePhoneNumber(input)
}

func (a *ChimeActivities) BatchUpdateUser(input *chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error) {
    return a.client.BatchUpdateUser(input)
}

func (a *ChimeActivities) CreateAccount(input *chime.CreateAccountInput) (*chime.CreateAccountOutput, error) {
    return a.client.CreateAccount(input)
}

func (a *ChimeActivities) CreateAttendee(input *chime.CreateAttendeeInput) (*chime.CreateAttendeeOutput, error) {
    return a.client.CreateAttendee(input)
}

func (a *ChimeActivities) CreateBot(input *chime.CreateBotInput) (*chime.CreateBotOutput, error) {
    return a.client.CreateBot(input)
}

func (a *ChimeActivities) CreateMeeting(input *chime.CreateMeetingInput) (*chime.CreateMeetingOutput, error) {
    return a.client.CreateMeeting(input)
}

func (a *ChimeActivities) CreateMeetingWithAttendees(input *chime.CreateMeetingWithAttendeesInput) (*chime.CreateMeetingWithAttendeesOutput, error) {
    return a.client.CreateMeetingWithAttendees(input)
}

func (a *ChimeActivities) CreatePhoneNumberOrder(input *chime.CreatePhoneNumberOrderInput) (*chime.CreatePhoneNumberOrderOutput, error) {
    return a.client.CreatePhoneNumberOrder(input)
}

func (a *ChimeActivities) CreateProxySession(input *chime.CreateProxySessionInput) (*chime.CreateProxySessionOutput, error) {
    return a.client.CreateProxySession(input)
}

func (a *ChimeActivities) CreateRoom(input *chime.CreateRoomInput) (*chime.CreateRoomOutput, error) {
    return a.client.CreateRoom(input)
}

func (a *ChimeActivities) CreateRoomMembership(input *chime.CreateRoomMembershipInput) (*chime.CreateRoomMembershipOutput, error) {
    return a.client.CreateRoomMembership(input)
}

func (a *ChimeActivities) CreateUser(input *chime.CreateUserInput) (*chime.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *ChimeActivities) CreateVoiceConnector(input *chime.CreateVoiceConnectorInput) (*chime.CreateVoiceConnectorOutput, error) {
    return a.client.CreateVoiceConnector(input)
}

func (a *ChimeActivities) CreateVoiceConnectorGroup(input *chime.CreateVoiceConnectorGroupInput) (*chime.CreateVoiceConnectorGroupOutput, error) {
    return a.client.CreateVoiceConnectorGroup(input)
}

func (a *ChimeActivities) DeleteAccount(input *chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error) {
    return a.client.DeleteAccount(input)
}

func (a *ChimeActivities) DeleteAttendee(input *chime.DeleteAttendeeInput) (*chime.DeleteAttendeeOutput, error) {
    return a.client.DeleteAttendee(input)
}

func (a *ChimeActivities) DeleteEventsConfiguration(input *chime.DeleteEventsConfigurationInput) (*chime.DeleteEventsConfigurationOutput, error) {
    return a.client.DeleteEventsConfiguration(input)
}

func (a *ChimeActivities) DeleteMeeting(input *chime.DeleteMeetingInput) (*chime.DeleteMeetingOutput, error) {
    return a.client.DeleteMeeting(input)
}

func (a *ChimeActivities) DeletePhoneNumber(input *chime.DeletePhoneNumberInput) (*chime.DeletePhoneNumberOutput, error) {
    return a.client.DeletePhoneNumber(input)
}

func (a *ChimeActivities) DeleteProxySession(input *chime.DeleteProxySessionInput) (*chime.DeleteProxySessionOutput, error) {
    return a.client.DeleteProxySession(input)
}

func (a *ChimeActivities) DeleteRoom(input *chime.DeleteRoomInput) (*chime.DeleteRoomOutput, error) {
    return a.client.DeleteRoom(input)
}

func (a *ChimeActivities) DeleteRoomMembership(input *chime.DeleteRoomMembershipInput) (*chime.DeleteRoomMembershipOutput, error) {
    return a.client.DeleteRoomMembership(input)
}

func (a *ChimeActivities) DeleteVoiceConnector(input *chime.DeleteVoiceConnectorInput) (*chime.DeleteVoiceConnectorOutput, error) {
    return a.client.DeleteVoiceConnector(input)
}

func (a *ChimeActivities) DeleteVoiceConnectorEmergencyCallingConfiguration(input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    return a.client.DeleteVoiceConnectorEmergencyCallingConfiguration(input)
}

func (a *ChimeActivities) DeleteVoiceConnectorGroup(input *chime.DeleteVoiceConnectorGroupInput) (*chime.DeleteVoiceConnectorGroupOutput, error) {
    return a.client.DeleteVoiceConnectorGroup(input)
}

func (a *ChimeActivities) DeleteVoiceConnectorOrigination(input *chime.DeleteVoiceConnectorOriginationInput) (*chime.DeleteVoiceConnectorOriginationOutput, error) {
    return a.client.DeleteVoiceConnectorOrigination(input)
}

func (a *ChimeActivities) DeleteVoiceConnectorProxy(input *chime.DeleteVoiceConnectorProxyInput) (*chime.DeleteVoiceConnectorProxyOutput, error) {
    return a.client.DeleteVoiceConnectorProxy(input)
}

func (a *ChimeActivities) DeleteVoiceConnectorStreamingConfiguration(input *chime.DeleteVoiceConnectorStreamingConfigurationInput) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error) {
    return a.client.DeleteVoiceConnectorStreamingConfiguration(input)
}

func (a *ChimeActivities) DeleteVoiceConnectorTermination(input *chime.DeleteVoiceConnectorTerminationInput) (*chime.DeleteVoiceConnectorTerminationOutput, error) {
    return a.client.DeleteVoiceConnectorTermination(input)
}

func (a *ChimeActivities) DeleteVoiceConnectorTerminationCredentials(input *chime.DeleteVoiceConnectorTerminationCredentialsInput) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error) {
    return a.client.DeleteVoiceConnectorTerminationCredentials(input)
}

func (a *ChimeActivities) DisassociatePhoneNumberFromUser(input *chime.DisassociatePhoneNumberFromUserInput) (*chime.DisassociatePhoneNumberFromUserOutput, error) {
    return a.client.DisassociatePhoneNumberFromUser(input)
}

func (a *ChimeActivities) DisassociatePhoneNumbersFromVoiceConnector(input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error) {
    return a.client.DisassociatePhoneNumbersFromVoiceConnector(input)
}

func (a *ChimeActivities) DisassociatePhoneNumbersFromVoiceConnectorGroup(input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error) {
    return a.client.DisassociatePhoneNumbersFromVoiceConnectorGroup(input)
}

func (a *ChimeActivities) DisassociateSigninDelegateGroupsFromAccount(input *chime.DisassociateSigninDelegateGroupsFromAccountInput) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error) {
    return a.client.DisassociateSigninDelegateGroupsFromAccount(input)
}

func (a *ChimeActivities) GetAccount(input *chime.GetAccountInput) (*chime.GetAccountOutput, error) {
    return a.client.GetAccount(input)
}

func (a *ChimeActivities) GetAccountSettings(input *chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error) {
    return a.client.GetAccountSettings(input)
}

func (a *ChimeActivities) GetAttendee(input *chime.GetAttendeeInput) (*chime.GetAttendeeOutput, error) {
    return a.client.GetAttendee(input)
}

func (a *ChimeActivities) GetBot(input *chime.GetBotInput) (*chime.GetBotOutput, error) {
    return a.client.GetBot(input)
}

func (a *ChimeActivities) GetEventsConfiguration(input *chime.GetEventsConfigurationInput) (*chime.GetEventsConfigurationOutput, error) {
    return a.client.GetEventsConfiguration(input)
}

func (a *ChimeActivities) GetGlobalSettings(input *chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error) {
    return a.client.GetGlobalSettings(input)
}

func (a *ChimeActivities) GetMeeting(input *chime.GetMeetingInput) (*chime.GetMeetingOutput, error) {
    return a.client.GetMeeting(input)
}

func (a *ChimeActivities) GetPhoneNumber(input *chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error) {
    return a.client.GetPhoneNumber(input)
}

func (a *ChimeActivities) GetPhoneNumberOrder(input *chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error) {
    return a.client.GetPhoneNumberOrder(input)
}

func (a *ChimeActivities) GetPhoneNumberSettings(input *chime.GetPhoneNumberSettingsInput) (*chime.GetPhoneNumberSettingsOutput, error) {
    return a.client.GetPhoneNumberSettings(input)
}

func (a *ChimeActivities) GetProxySession(input *chime.GetProxySessionInput) (*chime.GetProxySessionOutput, error) {
    return a.client.GetProxySession(input)
}

func (a *ChimeActivities) GetRetentionSettings(input *chime.GetRetentionSettingsInput) (*chime.GetRetentionSettingsOutput, error) {
    return a.client.GetRetentionSettings(input)
}

func (a *ChimeActivities) GetRoom(input *chime.GetRoomInput) (*chime.GetRoomOutput, error) {
    return a.client.GetRoom(input)
}

func (a *ChimeActivities) GetUser(input *chime.GetUserInput) (*chime.GetUserOutput, error) {
    return a.client.GetUser(input)
}

func (a *ChimeActivities) GetUserSettings(input *chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error) {
    return a.client.GetUserSettings(input)
}

func (a *ChimeActivities) GetVoiceConnector(input *chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error) {
    return a.client.GetVoiceConnector(input)
}

func (a *ChimeActivities) GetVoiceConnectorEmergencyCallingConfiguration(input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    return a.client.GetVoiceConnectorEmergencyCallingConfiguration(input)
}

func (a *ChimeActivities) GetVoiceConnectorGroup(input *chime.GetVoiceConnectorGroupInput) (*chime.GetVoiceConnectorGroupOutput, error) {
    return a.client.GetVoiceConnectorGroup(input)
}

func (a *ChimeActivities) GetVoiceConnectorLoggingConfiguration(input *chime.GetVoiceConnectorLoggingConfigurationInput) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error) {
    return a.client.GetVoiceConnectorLoggingConfiguration(input)
}

func (a *ChimeActivities) GetVoiceConnectorOrigination(input *chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error) {
    return a.client.GetVoiceConnectorOrigination(input)
}

func (a *ChimeActivities) GetVoiceConnectorProxy(input *chime.GetVoiceConnectorProxyInput) (*chime.GetVoiceConnectorProxyOutput, error) {
    return a.client.GetVoiceConnectorProxy(input)
}

func (a *ChimeActivities) GetVoiceConnectorStreamingConfiguration(input *chime.GetVoiceConnectorStreamingConfigurationInput) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error) {
    return a.client.GetVoiceConnectorStreamingConfiguration(input)
}

func (a *ChimeActivities) GetVoiceConnectorTermination(input *chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error) {
    return a.client.GetVoiceConnectorTermination(input)
}

func (a *ChimeActivities) GetVoiceConnectorTerminationHealth(input *chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error) {
    return a.client.GetVoiceConnectorTerminationHealth(input)
}

func (a *ChimeActivities) InviteUsers(input *chime.InviteUsersInput) (*chime.InviteUsersOutput, error) {
    return a.client.InviteUsers(input)
}

func (a *ChimeActivities) ListAccounts(input *chime.ListAccountsInput) (*chime.ListAccountsOutput, error) {
    return a.client.ListAccounts(input)
}

func (a *ChimeActivities) ListAttendeeTags(input *chime.ListAttendeeTagsInput) (*chime.ListAttendeeTagsOutput, error) {
    return a.client.ListAttendeeTags(input)
}

func (a *ChimeActivities) ListAttendees(input *chime.ListAttendeesInput) (*chime.ListAttendeesOutput, error) {
    return a.client.ListAttendees(input)
}

func (a *ChimeActivities) ListBots(input *chime.ListBotsInput) (*chime.ListBotsOutput, error) {
    return a.client.ListBots(input)
}

func (a *ChimeActivities) ListMeetingTags(input *chime.ListMeetingTagsInput) (*chime.ListMeetingTagsOutput, error) {
    return a.client.ListMeetingTags(input)
}

func (a *ChimeActivities) ListMeetings(input *chime.ListMeetingsInput) (*chime.ListMeetingsOutput, error) {
    return a.client.ListMeetings(input)
}

func (a *ChimeActivities) ListPhoneNumberOrders(input *chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error) {
    return a.client.ListPhoneNumberOrders(input)
}

func (a *ChimeActivities) ListPhoneNumbers(input *chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error) {
    return a.client.ListPhoneNumbers(input)
}

func (a *ChimeActivities) ListProxySessions(input *chime.ListProxySessionsInput) (*chime.ListProxySessionsOutput, error) {
    return a.client.ListProxySessions(input)
}

func (a *ChimeActivities) ListRoomMemberships(input *chime.ListRoomMembershipsInput) (*chime.ListRoomMembershipsOutput, error) {
    return a.client.ListRoomMemberships(input)
}

func (a *ChimeActivities) ListRooms(input *chime.ListRoomsInput) (*chime.ListRoomsOutput, error) {
    return a.client.ListRooms(input)
}

func (a *ChimeActivities) ListTagsForResource(input *chime.ListTagsForResourceInput) (*chime.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ChimeActivities) ListUsers(input *chime.ListUsersInput) (*chime.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *ChimeActivities) ListVoiceConnectorGroups(input *chime.ListVoiceConnectorGroupsInput) (*chime.ListVoiceConnectorGroupsOutput, error) {
    return a.client.ListVoiceConnectorGroups(input)
}

func (a *ChimeActivities) ListVoiceConnectorTerminationCredentials(input *chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error) {
    return a.client.ListVoiceConnectorTerminationCredentials(input)
}

func (a *ChimeActivities) ListVoiceConnectors(input *chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error) {
    return a.client.ListVoiceConnectors(input)
}

func (a *ChimeActivities) LogoutUser(input *chime.LogoutUserInput) (*chime.LogoutUserOutput, error) {
    return a.client.LogoutUser(input)
}

func (a *ChimeActivities) PutEventsConfiguration(input *chime.PutEventsConfigurationInput) (*chime.PutEventsConfigurationOutput, error) {
    return a.client.PutEventsConfiguration(input)
}

func (a *ChimeActivities) PutRetentionSettings(input *chime.PutRetentionSettingsInput) (*chime.PutRetentionSettingsOutput, error) {
    return a.client.PutRetentionSettings(input)
}

func (a *ChimeActivities) PutVoiceConnectorEmergencyCallingConfiguration(input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    return a.client.PutVoiceConnectorEmergencyCallingConfiguration(input)
}

func (a *ChimeActivities) PutVoiceConnectorLoggingConfiguration(input *chime.PutVoiceConnectorLoggingConfigurationInput) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error) {
    return a.client.PutVoiceConnectorLoggingConfiguration(input)
}

func (a *ChimeActivities) PutVoiceConnectorOrigination(input *chime.PutVoiceConnectorOriginationInput) (*chime.PutVoiceConnectorOriginationOutput, error) {
    return a.client.PutVoiceConnectorOrigination(input)
}

func (a *ChimeActivities) PutVoiceConnectorProxy(input *chime.PutVoiceConnectorProxyInput) (*chime.PutVoiceConnectorProxyOutput, error) {
    return a.client.PutVoiceConnectorProxy(input)
}

func (a *ChimeActivities) PutVoiceConnectorStreamingConfiguration(input *chime.PutVoiceConnectorStreamingConfigurationInput) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error) {
    return a.client.PutVoiceConnectorStreamingConfiguration(input)
}

func (a *ChimeActivities) PutVoiceConnectorTermination(input *chime.PutVoiceConnectorTerminationInput) (*chime.PutVoiceConnectorTerminationOutput, error) {
    return a.client.PutVoiceConnectorTermination(input)
}

func (a *ChimeActivities) PutVoiceConnectorTerminationCredentials(input *chime.PutVoiceConnectorTerminationCredentialsInput) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error) {
    return a.client.PutVoiceConnectorTerminationCredentials(input)
}

func (a *ChimeActivities) RedactConversationMessage(input *chime.RedactConversationMessageInput) (*chime.RedactConversationMessageOutput, error) {
    return a.client.RedactConversationMessage(input)
}

func (a *ChimeActivities) RedactRoomMessage(input *chime.RedactRoomMessageInput) (*chime.RedactRoomMessageOutput, error) {
    return a.client.RedactRoomMessage(input)
}

func (a *ChimeActivities) RegenerateSecurityToken(input *chime.RegenerateSecurityTokenInput) (*chime.RegenerateSecurityTokenOutput, error) {
    return a.client.RegenerateSecurityToken(input)
}

func (a *ChimeActivities) ResetPersonalPIN(input *chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error) {
    return a.client.ResetPersonalPIN(input)
}

func (a *ChimeActivities) RestorePhoneNumber(input *chime.RestorePhoneNumberInput) (*chime.RestorePhoneNumberOutput, error) {
    return a.client.RestorePhoneNumber(input)
}

func (a *ChimeActivities) SearchAvailablePhoneNumbers(input *chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error) {
    return a.client.SearchAvailablePhoneNumbers(input)
}

func (a *ChimeActivities) TagAttendee(input *chime.TagAttendeeInput) (*chime.TagAttendeeOutput, error) {
    return a.client.TagAttendee(input)
}

func (a *ChimeActivities) TagMeeting(input *chime.TagMeetingInput) (*chime.TagMeetingOutput, error) {
    return a.client.TagMeeting(input)
}

func (a *ChimeActivities) TagResource(input *chime.TagResourceInput) (*chime.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ChimeActivities) UntagAttendee(input *chime.UntagAttendeeInput) (*chime.UntagAttendeeOutput, error) {
    return a.client.UntagAttendee(input)
}

func (a *ChimeActivities) UntagMeeting(input *chime.UntagMeetingInput) (*chime.UntagMeetingOutput, error) {
    return a.client.UntagMeeting(input)
}

func (a *ChimeActivities) UntagResource(input *chime.UntagResourceInput) (*chime.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ChimeActivities) UpdateAccount(input *chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error) {
    return a.client.UpdateAccount(input)
}

func (a *ChimeActivities) UpdateAccountSettings(input *chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error) {
    return a.client.UpdateAccountSettings(input)
}

func (a *ChimeActivities) UpdateBot(input *chime.UpdateBotInput) (*chime.UpdateBotOutput, error) {
    return a.client.UpdateBot(input)
}

func (a *ChimeActivities) UpdateGlobalSettings(input *chime.UpdateGlobalSettingsInput) (*chime.UpdateGlobalSettingsOutput, error) {
    return a.client.UpdateGlobalSettings(input)
}

func (a *ChimeActivities) UpdatePhoneNumber(input *chime.UpdatePhoneNumberInput) (*chime.UpdatePhoneNumberOutput, error) {
    return a.client.UpdatePhoneNumber(input)
}

func (a *ChimeActivities) UpdatePhoneNumberSettings(input *chime.UpdatePhoneNumberSettingsInput) (*chime.UpdatePhoneNumberSettingsOutput, error) {
    return a.client.UpdatePhoneNumberSettings(input)
}

func (a *ChimeActivities) UpdateProxySession(input *chime.UpdateProxySessionInput) (*chime.UpdateProxySessionOutput, error) {
    return a.client.UpdateProxySession(input)
}

func (a *ChimeActivities) UpdateRoom(input *chime.UpdateRoomInput) (*chime.UpdateRoomOutput, error) {
    return a.client.UpdateRoom(input)
}

func (a *ChimeActivities) UpdateRoomMembership(input *chime.UpdateRoomMembershipInput) (*chime.UpdateRoomMembershipOutput, error) {
    return a.client.UpdateRoomMembership(input)
}

func (a *ChimeActivities) UpdateUser(input *chime.UpdateUserInput) (*chime.UpdateUserOutput, error) {
    return a.client.UpdateUser(input)
}

func (a *ChimeActivities) UpdateUserSettings(input *chime.UpdateUserSettingsInput) (*chime.UpdateUserSettingsOutput, error) {
    return a.client.UpdateUserSettings(input)
}

func (a *ChimeActivities) UpdateVoiceConnector(input *chime.UpdateVoiceConnectorInput) (*chime.UpdateVoiceConnectorOutput, error) {
    return a.client.UpdateVoiceConnector(input)
}

func (a *ChimeActivities) UpdateVoiceConnectorGroup(input *chime.UpdateVoiceConnectorGroupInput) (*chime.UpdateVoiceConnectorGroupOutput, error) {
    return a.client.UpdateVoiceConnectorGroup(input)
}
