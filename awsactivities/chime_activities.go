
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/chime"
	"github.com/aws/aws-sdk-go/service/chime/chimeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ChimeActivities struct {
    client chimeiface.ChimeAPI
}

func NewChimeActivities(session *session.Session, config ...*aws.Config) *ChimeActivities {
    client := chime.New(session, config...)
    return &ChimeActivities{client: client}
}

func (a *ChimeActivities) AssociatePhoneNumberWithUser(ctx context.Context, input *chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error) {
    return a.client.AssociatePhoneNumberWithUserWithContext(ctx, input)
}

func (a *ChimeActivities) AssociatePhoneNumbersWithVoiceConnector(ctx context.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error) {
    return a.client.AssociatePhoneNumbersWithVoiceConnectorWithContext(ctx, input)
}

func (a *ChimeActivities) AssociatePhoneNumbersWithVoiceConnectorGroup(ctx context.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) {
    return a.client.AssociatePhoneNumbersWithVoiceConnectorGroupWithContext(ctx, input)
}

func (a *ChimeActivities) AssociateSigninDelegateGroupsWithAccount(ctx context.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error) {
    return a.client.AssociateSigninDelegateGroupsWithAccountWithContext(ctx, input)
}

func (a *ChimeActivities) BatchCreateAttendee(ctx context.Context, input *chime.BatchCreateAttendeeInput) (*chime.BatchCreateAttendeeOutput, error) {
    return a.client.BatchCreateAttendeeWithContext(ctx, input)
}

func (a *ChimeActivities) BatchCreateRoomMembership(ctx context.Context, input *chime.BatchCreateRoomMembershipInput) (*chime.BatchCreateRoomMembershipOutput, error) {
    return a.client.BatchCreateRoomMembershipWithContext(ctx, input)
}

func (a *ChimeActivities) BatchDeletePhoneNumber(ctx context.Context, input *chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error) {
    return a.client.BatchDeletePhoneNumberWithContext(ctx, input)
}

func (a *ChimeActivities) BatchSuspendUser(ctx context.Context, input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error) {
    return a.client.BatchSuspendUserWithContext(ctx, input)
}

func (a *ChimeActivities) BatchUnsuspendUser(ctx context.Context, input *chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error) {
    return a.client.BatchUnsuspendUserWithContext(ctx, input)
}

func (a *ChimeActivities) BatchUpdatePhoneNumber(ctx context.Context, input *chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error) {
    return a.client.BatchUpdatePhoneNumberWithContext(ctx, input)
}

func (a *ChimeActivities) BatchUpdateUser(ctx context.Context, input *chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error) {
    return a.client.BatchUpdateUserWithContext(ctx, input)
}

func (a *ChimeActivities) CreateAccount(ctx context.Context, input *chime.CreateAccountInput) (*chime.CreateAccountOutput, error) {
    return a.client.CreateAccountWithContext(ctx, input)
}

func (a *ChimeActivities) CreateAttendee(ctx context.Context, input *chime.CreateAttendeeInput) (*chime.CreateAttendeeOutput, error) {
    return a.client.CreateAttendeeWithContext(ctx, input)
}

func (a *ChimeActivities) CreateBot(ctx context.Context, input *chime.CreateBotInput) (*chime.CreateBotOutput, error) {
    return a.client.CreateBotWithContext(ctx, input)
}

func (a *ChimeActivities) CreateMeeting(ctx context.Context, input *chime.CreateMeetingInput) (*chime.CreateMeetingOutput, error) {
    return a.client.CreateMeetingWithContext(ctx, input)
}

func (a *ChimeActivities) CreateMeetingWithAttendees(ctx context.Context, input *chime.CreateMeetingWithAttendeesInput) (*chime.CreateMeetingWithAttendeesOutput, error) {
    return a.client.CreateMeetingWithAttendeesWithContext(ctx, input)
}

func (a *ChimeActivities) CreatePhoneNumberOrder(ctx context.Context, input *chime.CreatePhoneNumberOrderInput) (*chime.CreatePhoneNumberOrderOutput, error) {
    return a.client.CreatePhoneNumberOrderWithContext(ctx, input)
}

func (a *ChimeActivities) CreateProxySession(ctx context.Context, input *chime.CreateProxySessionInput) (*chime.CreateProxySessionOutput, error) {
    return a.client.CreateProxySessionWithContext(ctx, input)
}

func (a *ChimeActivities) CreateRoom(ctx context.Context, input *chime.CreateRoomInput) (*chime.CreateRoomOutput, error) {
    return a.client.CreateRoomWithContext(ctx, input)
}

func (a *ChimeActivities) CreateRoomMembership(ctx context.Context, input *chime.CreateRoomMembershipInput) (*chime.CreateRoomMembershipOutput, error) {
    return a.client.CreateRoomMembershipWithContext(ctx, input)
}

func (a *ChimeActivities) CreateUser(ctx context.Context, input *chime.CreateUserInput) (*chime.CreateUserOutput, error) {
    return a.client.CreateUserWithContext(ctx, input)
}

func (a *ChimeActivities) CreateVoiceConnector(ctx context.Context, input *chime.CreateVoiceConnectorInput) (*chime.CreateVoiceConnectorOutput, error) {
    return a.client.CreateVoiceConnectorWithContext(ctx, input)
}

func (a *ChimeActivities) CreateVoiceConnectorGroup(ctx context.Context, input *chime.CreateVoiceConnectorGroupInput) (*chime.CreateVoiceConnectorGroupOutput, error) {
    return a.client.CreateVoiceConnectorGroupWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteAccount(ctx context.Context, input *chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error) {
    return a.client.DeleteAccountWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteAttendee(ctx context.Context, input *chime.DeleteAttendeeInput) (*chime.DeleteAttendeeOutput, error) {
    return a.client.DeleteAttendeeWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteEventsConfiguration(ctx context.Context, input *chime.DeleteEventsConfigurationInput) (*chime.DeleteEventsConfigurationOutput, error) {
    return a.client.DeleteEventsConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteMeeting(ctx context.Context, input *chime.DeleteMeetingInput) (*chime.DeleteMeetingOutput, error) {
    return a.client.DeleteMeetingWithContext(ctx, input)
}

func (a *ChimeActivities) DeletePhoneNumber(ctx context.Context, input *chime.DeletePhoneNumberInput) (*chime.DeletePhoneNumberOutput, error) {
    return a.client.DeletePhoneNumberWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteProxySession(ctx context.Context, input *chime.DeleteProxySessionInput) (*chime.DeleteProxySessionOutput, error) {
    return a.client.DeleteProxySessionWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteRoom(ctx context.Context, input *chime.DeleteRoomInput) (*chime.DeleteRoomOutput, error) {
    return a.client.DeleteRoomWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteRoomMembership(ctx context.Context, input *chime.DeleteRoomMembershipInput) (*chime.DeleteRoomMembershipOutput, error) {
    return a.client.DeleteRoomMembershipWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnector(ctx context.Context, input *chime.DeleteVoiceConnectorInput) (*chime.DeleteVoiceConnectorOutput, error) {
    return a.client.DeleteVoiceConnectorWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    return a.client.DeleteVoiceConnectorEmergencyCallingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnectorGroup(ctx context.Context, input *chime.DeleteVoiceConnectorGroupInput) (*chime.DeleteVoiceConnectorGroupOutput, error) {
    return a.client.DeleteVoiceConnectorGroupWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnectorOrigination(ctx context.Context, input *chime.DeleteVoiceConnectorOriginationInput) (*chime.DeleteVoiceConnectorOriginationOutput, error) {
    return a.client.DeleteVoiceConnectorOriginationWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnectorProxy(ctx context.Context, input *chime.DeleteVoiceConnectorProxyInput) (*chime.DeleteVoiceConnectorProxyOutput, error) {
    return a.client.DeleteVoiceConnectorProxyWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnectorStreamingConfiguration(ctx context.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error) {
    return a.client.DeleteVoiceConnectorStreamingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnectorTermination(ctx context.Context, input *chime.DeleteVoiceConnectorTerminationInput) (*chime.DeleteVoiceConnectorTerminationOutput, error) {
    return a.client.DeleteVoiceConnectorTerminationWithContext(ctx, input)
}

func (a *ChimeActivities) DeleteVoiceConnectorTerminationCredentials(ctx context.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error) {
    return a.client.DeleteVoiceConnectorTerminationCredentialsWithContext(ctx, input)
}

func (a *ChimeActivities) DisassociatePhoneNumberFromUser(ctx context.Context, input *chime.DisassociatePhoneNumberFromUserInput) (*chime.DisassociatePhoneNumberFromUserOutput, error) {
    return a.client.DisassociatePhoneNumberFromUserWithContext(ctx, input)
}

func (a *ChimeActivities) DisassociatePhoneNumbersFromVoiceConnector(ctx context.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error) {
    return a.client.DisassociatePhoneNumbersFromVoiceConnectorWithContext(ctx, input)
}

func (a *ChimeActivities) DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx context.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error) {
    return a.client.DisassociatePhoneNumbersFromVoiceConnectorGroupWithContext(ctx, input)
}

func (a *ChimeActivities) DisassociateSigninDelegateGroupsFromAccount(ctx context.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error) {
    return a.client.DisassociateSigninDelegateGroupsFromAccountWithContext(ctx, input)
}

func (a *ChimeActivities) GetAccount(ctx context.Context, input *chime.GetAccountInput) (*chime.GetAccountOutput, error) {
    return a.client.GetAccountWithContext(ctx, input)
}

func (a *ChimeActivities) GetAccountSettings(ctx context.Context, input *chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error) {
    return a.client.GetAccountSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) GetAttendee(ctx context.Context, input *chime.GetAttendeeInput) (*chime.GetAttendeeOutput, error) {
    return a.client.GetAttendeeWithContext(ctx, input)
}

func (a *ChimeActivities) GetBot(ctx context.Context, input *chime.GetBotInput) (*chime.GetBotOutput, error) {
    return a.client.GetBotWithContext(ctx, input)
}

func (a *ChimeActivities) GetEventsConfiguration(ctx context.Context, input *chime.GetEventsConfigurationInput) (*chime.GetEventsConfigurationOutput, error) {
    return a.client.GetEventsConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) GetGlobalSettings(ctx context.Context, input *chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error) {
    return a.client.GetGlobalSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) GetMeeting(ctx context.Context, input *chime.GetMeetingInput) (*chime.GetMeetingOutput, error) {
    return a.client.GetMeetingWithContext(ctx, input)
}

func (a *ChimeActivities) GetPhoneNumber(ctx context.Context, input *chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error) {
    return a.client.GetPhoneNumberWithContext(ctx, input)
}

func (a *ChimeActivities) GetPhoneNumberOrder(ctx context.Context, input *chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error) {
    return a.client.GetPhoneNumberOrderWithContext(ctx, input)
}

func (a *ChimeActivities) GetPhoneNumberSettings(ctx context.Context, input *chime.GetPhoneNumberSettingsInput) (*chime.GetPhoneNumberSettingsOutput, error) {
    return a.client.GetPhoneNumberSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) GetProxySession(ctx context.Context, input *chime.GetProxySessionInput) (*chime.GetProxySessionOutput, error) {
    return a.client.GetProxySessionWithContext(ctx, input)
}

func (a *ChimeActivities) GetRetentionSettings(ctx context.Context, input *chime.GetRetentionSettingsInput) (*chime.GetRetentionSettingsOutput, error) {
    return a.client.GetRetentionSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) GetRoom(ctx context.Context, input *chime.GetRoomInput) (*chime.GetRoomOutput, error) {
    return a.client.GetRoomWithContext(ctx, input)
}

func (a *ChimeActivities) GetUser(ctx context.Context, input *chime.GetUserInput) (*chime.GetUserOutput, error) {
    return a.client.GetUserWithContext(ctx, input)
}

func (a *ChimeActivities) GetUserSettings(ctx context.Context, input *chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error) {
    return a.client.GetUserSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnector(ctx context.Context, input *chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error) {
    return a.client.GetVoiceConnectorWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    return a.client.GetVoiceConnectorEmergencyCallingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorGroup(ctx context.Context, input *chime.GetVoiceConnectorGroupInput) (*chime.GetVoiceConnectorGroupOutput, error) {
    return a.client.GetVoiceConnectorGroupWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorLoggingConfiguration(ctx context.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error) {
    return a.client.GetVoiceConnectorLoggingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorOrigination(ctx context.Context, input *chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error) {
    return a.client.GetVoiceConnectorOriginationWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorProxy(ctx context.Context, input *chime.GetVoiceConnectorProxyInput) (*chime.GetVoiceConnectorProxyOutput, error) {
    return a.client.GetVoiceConnectorProxyWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorStreamingConfiguration(ctx context.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error) {
    return a.client.GetVoiceConnectorStreamingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorTermination(ctx context.Context, input *chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error) {
    return a.client.GetVoiceConnectorTerminationWithContext(ctx, input)
}

func (a *ChimeActivities) GetVoiceConnectorTerminationHealth(ctx context.Context, input *chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error) {
    return a.client.GetVoiceConnectorTerminationHealthWithContext(ctx, input)
}

func (a *ChimeActivities) InviteUsers(ctx context.Context, input *chime.InviteUsersInput) (*chime.InviteUsersOutput, error) {
    return a.client.InviteUsersWithContext(ctx, input)
}

func (a *ChimeActivities) ListAccounts(ctx context.Context, input *chime.ListAccountsInput) (*chime.ListAccountsOutput, error) {
    return a.client.ListAccountsWithContext(ctx, input)
}

func (a *ChimeActivities) ListAttendeeTags(ctx context.Context, input *chime.ListAttendeeTagsInput) (*chime.ListAttendeeTagsOutput, error) {
    return a.client.ListAttendeeTagsWithContext(ctx, input)
}

func (a *ChimeActivities) ListAttendees(ctx context.Context, input *chime.ListAttendeesInput) (*chime.ListAttendeesOutput, error) {
    return a.client.ListAttendeesWithContext(ctx, input)
}

func (a *ChimeActivities) ListBots(ctx context.Context, input *chime.ListBotsInput) (*chime.ListBotsOutput, error) {
    return a.client.ListBotsWithContext(ctx, input)
}

func (a *ChimeActivities) ListMeetingTags(ctx context.Context, input *chime.ListMeetingTagsInput) (*chime.ListMeetingTagsOutput, error) {
    return a.client.ListMeetingTagsWithContext(ctx, input)
}

func (a *ChimeActivities) ListMeetings(ctx context.Context, input *chime.ListMeetingsInput) (*chime.ListMeetingsOutput, error) {
    return a.client.ListMeetingsWithContext(ctx, input)
}

func (a *ChimeActivities) ListPhoneNumberOrders(ctx context.Context, input *chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error) {
    return a.client.ListPhoneNumberOrdersWithContext(ctx, input)
}

func (a *ChimeActivities) ListPhoneNumbers(ctx context.Context, input *chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error) {
    return a.client.ListPhoneNumbersWithContext(ctx, input)
}

func (a *ChimeActivities) ListProxySessions(ctx context.Context, input *chime.ListProxySessionsInput) (*chime.ListProxySessionsOutput, error) {
    return a.client.ListProxySessionsWithContext(ctx, input)
}

func (a *ChimeActivities) ListRoomMemberships(ctx context.Context, input *chime.ListRoomMembershipsInput) (*chime.ListRoomMembershipsOutput, error) {
    return a.client.ListRoomMembershipsWithContext(ctx, input)
}

func (a *ChimeActivities) ListRooms(ctx context.Context, input *chime.ListRoomsInput) (*chime.ListRoomsOutput, error) {
    return a.client.ListRoomsWithContext(ctx, input)
}

func (a *ChimeActivities) ListTagsForResource(ctx context.Context, input *chime.ListTagsForResourceInput) (*chime.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ChimeActivities) ListUsers(ctx context.Context, input *chime.ListUsersInput) (*chime.ListUsersOutput, error) {
    return a.client.ListUsersWithContext(ctx, input)
}

func (a *ChimeActivities) ListVoiceConnectorGroups(ctx context.Context, input *chime.ListVoiceConnectorGroupsInput) (*chime.ListVoiceConnectorGroupsOutput, error) {
    return a.client.ListVoiceConnectorGroupsWithContext(ctx, input)
}

func (a *ChimeActivities) ListVoiceConnectorTerminationCredentials(ctx context.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error) {
    return a.client.ListVoiceConnectorTerminationCredentialsWithContext(ctx, input)
}

func (a *ChimeActivities) ListVoiceConnectors(ctx context.Context, input *chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error) {
    return a.client.ListVoiceConnectorsWithContext(ctx, input)
}

func (a *ChimeActivities) LogoutUser(ctx context.Context, input *chime.LogoutUserInput) (*chime.LogoutUserOutput, error) {
    return a.client.LogoutUserWithContext(ctx, input)
}

func (a *ChimeActivities) PutEventsConfiguration(ctx context.Context, input *chime.PutEventsConfigurationInput) (*chime.PutEventsConfigurationOutput, error) {
    return a.client.PutEventsConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) PutRetentionSettings(ctx context.Context, input *chime.PutRetentionSettingsInput) (*chime.PutRetentionSettingsOutput, error) {
    return a.client.PutRetentionSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) PutVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    return a.client.PutVoiceConnectorEmergencyCallingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) PutVoiceConnectorLoggingConfiguration(ctx context.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error) {
    return a.client.PutVoiceConnectorLoggingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) PutVoiceConnectorOrigination(ctx context.Context, input *chime.PutVoiceConnectorOriginationInput) (*chime.PutVoiceConnectorOriginationOutput, error) {
    return a.client.PutVoiceConnectorOriginationWithContext(ctx, input)
}

func (a *ChimeActivities) PutVoiceConnectorProxy(ctx context.Context, input *chime.PutVoiceConnectorProxyInput) (*chime.PutVoiceConnectorProxyOutput, error) {
    return a.client.PutVoiceConnectorProxyWithContext(ctx, input)
}

func (a *ChimeActivities) PutVoiceConnectorStreamingConfiguration(ctx context.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error) {
    return a.client.PutVoiceConnectorStreamingConfigurationWithContext(ctx, input)
}

func (a *ChimeActivities) PutVoiceConnectorTermination(ctx context.Context, input *chime.PutVoiceConnectorTerminationInput) (*chime.PutVoiceConnectorTerminationOutput, error) {
    return a.client.PutVoiceConnectorTerminationWithContext(ctx, input)
}

func (a *ChimeActivities) PutVoiceConnectorTerminationCredentials(ctx context.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error) {
    return a.client.PutVoiceConnectorTerminationCredentialsWithContext(ctx, input)
}

func (a *ChimeActivities) RedactConversationMessage(ctx context.Context, input *chime.RedactConversationMessageInput) (*chime.RedactConversationMessageOutput, error) {
    return a.client.RedactConversationMessageWithContext(ctx, input)
}

func (a *ChimeActivities) RedactRoomMessage(ctx context.Context, input *chime.RedactRoomMessageInput) (*chime.RedactRoomMessageOutput, error) {
    return a.client.RedactRoomMessageWithContext(ctx, input)
}

func (a *ChimeActivities) RegenerateSecurityToken(ctx context.Context, input *chime.RegenerateSecurityTokenInput) (*chime.RegenerateSecurityTokenOutput, error) {
    return a.client.RegenerateSecurityTokenWithContext(ctx, input)
}

func (a *ChimeActivities) ResetPersonalPIN(ctx context.Context, input *chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error) {
    return a.client.ResetPersonalPINWithContext(ctx, input)
}

func (a *ChimeActivities) RestorePhoneNumber(ctx context.Context, input *chime.RestorePhoneNumberInput) (*chime.RestorePhoneNumberOutput, error) {
    return a.client.RestorePhoneNumberWithContext(ctx, input)
}

func (a *ChimeActivities) SearchAvailablePhoneNumbers(ctx context.Context, input *chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error) {
    return a.client.SearchAvailablePhoneNumbersWithContext(ctx, input)
}

func (a *ChimeActivities) TagAttendee(ctx context.Context, input *chime.TagAttendeeInput) (*chime.TagAttendeeOutput, error) {
    return a.client.TagAttendeeWithContext(ctx, input)
}

func (a *ChimeActivities) TagMeeting(ctx context.Context, input *chime.TagMeetingInput) (*chime.TagMeetingOutput, error) {
    return a.client.TagMeetingWithContext(ctx, input)
}

func (a *ChimeActivities) TagResource(ctx context.Context, input *chime.TagResourceInput) (*chime.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *ChimeActivities) UntagAttendee(ctx context.Context, input *chime.UntagAttendeeInput) (*chime.UntagAttendeeOutput, error) {
    return a.client.UntagAttendeeWithContext(ctx, input)
}

func (a *ChimeActivities) UntagMeeting(ctx context.Context, input *chime.UntagMeetingInput) (*chime.UntagMeetingOutput, error) {
    return a.client.UntagMeetingWithContext(ctx, input)
}

func (a *ChimeActivities) UntagResource(ctx context.Context, input *chime.UntagResourceInput) (*chime.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateAccount(ctx context.Context, input *chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error) {
    return a.client.UpdateAccountWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateAccountSettings(ctx context.Context, input *chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error) {
    return a.client.UpdateAccountSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateBot(ctx context.Context, input *chime.UpdateBotInput) (*chime.UpdateBotOutput, error) {
    return a.client.UpdateBotWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateGlobalSettings(ctx context.Context, input *chime.UpdateGlobalSettingsInput) (*chime.UpdateGlobalSettingsOutput, error) {
    return a.client.UpdateGlobalSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) UpdatePhoneNumber(ctx context.Context, input *chime.UpdatePhoneNumberInput) (*chime.UpdatePhoneNumberOutput, error) {
    return a.client.UpdatePhoneNumberWithContext(ctx, input)
}

func (a *ChimeActivities) UpdatePhoneNumberSettings(ctx context.Context, input *chime.UpdatePhoneNumberSettingsInput) (*chime.UpdatePhoneNumberSettingsOutput, error) {
    return a.client.UpdatePhoneNumberSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateProxySession(ctx context.Context, input *chime.UpdateProxySessionInput) (*chime.UpdateProxySessionOutput, error) {
    return a.client.UpdateProxySessionWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateRoom(ctx context.Context, input *chime.UpdateRoomInput) (*chime.UpdateRoomOutput, error) {
    return a.client.UpdateRoomWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateRoomMembership(ctx context.Context, input *chime.UpdateRoomMembershipInput) (*chime.UpdateRoomMembershipOutput, error) {
    return a.client.UpdateRoomMembershipWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateUser(ctx context.Context, input *chime.UpdateUserInput) (*chime.UpdateUserOutput, error) {
    return a.client.UpdateUserWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateUserSettings(ctx context.Context, input *chime.UpdateUserSettingsInput) (*chime.UpdateUserSettingsOutput, error) {
    return a.client.UpdateUserSettingsWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateVoiceConnector(ctx context.Context, input *chime.UpdateVoiceConnectorInput) (*chime.UpdateVoiceConnectorOutput, error) {
    return a.client.UpdateVoiceConnectorWithContext(ctx, input)
}

func (a *ChimeActivities) UpdateVoiceConnectorGroup(ctx context.Context, input *chime.UpdateVoiceConnectorGroupInput) (*chime.UpdateVoiceConnectorGroupOutput, error) {
    return a.client.UpdateVoiceConnectorGroupWithContext(ctx, input)
}
