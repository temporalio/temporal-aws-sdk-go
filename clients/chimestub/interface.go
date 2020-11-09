// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package chimestub

import (
	"github.com/aws/aws-sdk-go/service/chime"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociatePhoneNumberWithUser(ctx workflow.Context, input *chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error)
	AssociatePhoneNumberWithUserAsync(ctx workflow.Context, input *chime.AssociatePhoneNumberWithUserInput) *AssociatePhoneNumberWithUserFuture

	AssociatePhoneNumbersWithVoiceConnector(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error)
	AssociatePhoneNumbersWithVoiceConnectorAsync(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) *AssociatePhoneNumbersWithVoiceConnectorFuture

	AssociatePhoneNumbersWithVoiceConnectorGroup(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error)
	AssociatePhoneNumbersWithVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) *AssociatePhoneNumbersWithVoiceConnectorGroupFuture

	AssociateSigninDelegateGroupsWithAccount(ctx workflow.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error)
	AssociateSigninDelegateGroupsWithAccountAsync(ctx workflow.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) *AssociateSigninDelegateGroupsWithAccountFuture

	BatchCreateAttendee(ctx workflow.Context, input *chime.BatchCreateAttendeeInput) (*chime.BatchCreateAttendeeOutput, error)
	BatchCreateAttendeeAsync(ctx workflow.Context, input *chime.BatchCreateAttendeeInput) *BatchCreateAttendeeFuture

	BatchCreateRoomMembership(ctx workflow.Context, input *chime.BatchCreateRoomMembershipInput) (*chime.BatchCreateRoomMembershipOutput, error)
	BatchCreateRoomMembershipAsync(ctx workflow.Context, input *chime.BatchCreateRoomMembershipInput) *BatchCreateRoomMembershipFuture

	BatchDeletePhoneNumber(ctx workflow.Context, input *chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error)
	BatchDeletePhoneNumberAsync(ctx workflow.Context, input *chime.BatchDeletePhoneNumberInput) *BatchDeletePhoneNumberFuture

	BatchSuspendUser(ctx workflow.Context, input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error)
	BatchSuspendUserAsync(ctx workflow.Context, input *chime.BatchSuspendUserInput) *BatchSuspendUserFuture

	BatchUnsuspendUser(ctx workflow.Context, input *chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error)
	BatchUnsuspendUserAsync(ctx workflow.Context, input *chime.BatchUnsuspendUserInput) *BatchUnsuspendUserFuture

	BatchUpdatePhoneNumber(ctx workflow.Context, input *chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error)
	BatchUpdatePhoneNumberAsync(ctx workflow.Context, input *chime.BatchUpdatePhoneNumberInput) *BatchUpdatePhoneNumberFuture

	BatchUpdateUser(ctx workflow.Context, input *chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error)
	BatchUpdateUserAsync(ctx workflow.Context, input *chime.BatchUpdateUserInput) *BatchUpdateUserFuture

	CreateAccount(ctx workflow.Context, input *chime.CreateAccountInput) (*chime.CreateAccountOutput, error)
	CreateAccountAsync(ctx workflow.Context, input *chime.CreateAccountInput) *CreateAccountFuture

	CreateAttendee(ctx workflow.Context, input *chime.CreateAttendeeInput) (*chime.CreateAttendeeOutput, error)
	CreateAttendeeAsync(ctx workflow.Context, input *chime.CreateAttendeeInput) *CreateAttendeeFuture

	CreateBot(ctx workflow.Context, input *chime.CreateBotInput) (*chime.CreateBotOutput, error)
	CreateBotAsync(ctx workflow.Context, input *chime.CreateBotInput) *CreateBotFuture

	CreateMeeting(ctx workflow.Context, input *chime.CreateMeetingInput) (*chime.CreateMeetingOutput, error)
	CreateMeetingAsync(ctx workflow.Context, input *chime.CreateMeetingInput) *CreateMeetingFuture

	CreateMeetingWithAttendees(ctx workflow.Context, input *chime.CreateMeetingWithAttendeesInput) (*chime.CreateMeetingWithAttendeesOutput, error)
	CreateMeetingWithAttendeesAsync(ctx workflow.Context, input *chime.CreateMeetingWithAttendeesInput) *CreateMeetingWithAttendeesFuture

	CreatePhoneNumberOrder(ctx workflow.Context, input *chime.CreatePhoneNumberOrderInput) (*chime.CreatePhoneNumberOrderOutput, error)
	CreatePhoneNumberOrderAsync(ctx workflow.Context, input *chime.CreatePhoneNumberOrderInput) *CreatePhoneNumberOrderFuture

	CreateProxySession(ctx workflow.Context, input *chime.CreateProxySessionInput) (*chime.CreateProxySessionOutput, error)
	CreateProxySessionAsync(ctx workflow.Context, input *chime.CreateProxySessionInput) *CreateProxySessionFuture

	CreateRoom(ctx workflow.Context, input *chime.CreateRoomInput) (*chime.CreateRoomOutput, error)
	CreateRoomAsync(ctx workflow.Context, input *chime.CreateRoomInput) *CreateRoomFuture

	CreateRoomMembership(ctx workflow.Context, input *chime.CreateRoomMembershipInput) (*chime.CreateRoomMembershipOutput, error)
	CreateRoomMembershipAsync(ctx workflow.Context, input *chime.CreateRoomMembershipInput) *CreateRoomMembershipFuture

	CreateUser(ctx workflow.Context, input *chime.CreateUserInput) (*chime.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *chime.CreateUserInput) *CreateUserFuture

	CreateVoiceConnector(ctx workflow.Context, input *chime.CreateVoiceConnectorInput) (*chime.CreateVoiceConnectorOutput, error)
	CreateVoiceConnectorAsync(ctx workflow.Context, input *chime.CreateVoiceConnectorInput) *CreateVoiceConnectorFuture

	CreateVoiceConnectorGroup(ctx workflow.Context, input *chime.CreateVoiceConnectorGroupInput) (*chime.CreateVoiceConnectorGroupOutput, error)
	CreateVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.CreateVoiceConnectorGroupInput) *CreateVoiceConnectorGroupFuture

	DeleteAccount(ctx workflow.Context, input *chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error)
	DeleteAccountAsync(ctx workflow.Context, input *chime.DeleteAccountInput) *DeleteAccountFuture

	DeleteAttendee(ctx workflow.Context, input *chime.DeleteAttendeeInput) (*chime.DeleteAttendeeOutput, error)
	DeleteAttendeeAsync(ctx workflow.Context, input *chime.DeleteAttendeeInput) *DeleteAttendeeFuture

	DeleteEventsConfiguration(ctx workflow.Context, input *chime.DeleteEventsConfigurationInput) (*chime.DeleteEventsConfigurationOutput, error)
	DeleteEventsConfigurationAsync(ctx workflow.Context, input *chime.DeleteEventsConfigurationInput) *DeleteEventsConfigurationFuture

	DeleteMeeting(ctx workflow.Context, input *chime.DeleteMeetingInput) (*chime.DeleteMeetingOutput, error)
	DeleteMeetingAsync(ctx workflow.Context, input *chime.DeleteMeetingInput) *DeleteMeetingFuture

	DeletePhoneNumber(ctx workflow.Context, input *chime.DeletePhoneNumberInput) (*chime.DeletePhoneNumberOutput, error)
	DeletePhoneNumberAsync(ctx workflow.Context, input *chime.DeletePhoneNumberInput) *DeletePhoneNumberFuture

	DeleteProxySession(ctx workflow.Context, input *chime.DeleteProxySessionInput) (*chime.DeleteProxySessionOutput, error)
	DeleteProxySessionAsync(ctx workflow.Context, input *chime.DeleteProxySessionInput) *DeleteProxySessionFuture

	DeleteRoom(ctx workflow.Context, input *chime.DeleteRoomInput) (*chime.DeleteRoomOutput, error)
	DeleteRoomAsync(ctx workflow.Context, input *chime.DeleteRoomInput) *DeleteRoomFuture

	DeleteRoomMembership(ctx workflow.Context, input *chime.DeleteRoomMembershipInput) (*chime.DeleteRoomMembershipOutput, error)
	DeleteRoomMembershipAsync(ctx workflow.Context, input *chime.DeleteRoomMembershipInput) *DeleteRoomMembershipFuture

	DeleteVoiceConnector(ctx workflow.Context, input *chime.DeleteVoiceConnectorInput) (*chime.DeleteVoiceConnectorOutput, error)
	DeleteVoiceConnectorAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorInput) *DeleteVoiceConnectorFuture

	DeleteVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error)
	DeleteVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) *DeleteVoiceConnectorEmergencyCallingConfigurationFuture

	DeleteVoiceConnectorGroup(ctx workflow.Context, input *chime.DeleteVoiceConnectorGroupInput) (*chime.DeleteVoiceConnectorGroupOutput, error)
	DeleteVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorGroupInput) *DeleteVoiceConnectorGroupFuture

	DeleteVoiceConnectorOrigination(ctx workflow.Context, input *chime.DeleteVoiceConnectorOriginationInput) (*chime.DeleteVoiceConnectorOriginationOutput, error)
	DeleteVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorOriginationInput) *DeleteVoiceConnectorOriginationFuture

	DeleteVoiceConnectorProxy(ctx workflow.Context, input *chime.DeleteVoiceConnectorProxyInput) (*chime.DeleteVoiceConnectorProxyOutput, error)
	DeleteVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorProxyInput) *DeleteVoiceConnectorProxyFuture

	DeleteVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error)
	DeleteVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) *DeleteVoiceConnectorStreamingConfigurationFuture

	DeleteVoiceConnectorTermination(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationInput) (*chime.DeleteVoiceConnectorTerminationOutput, error)
	DeleteVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationInput) *DeleteVoiceConnectorTerminationFuture

	DeleteVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error)
	DeleteVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) *DeleteVoiceConnectorTerminationCredentialsFuture

	DisassociatePhoneNumberFromUser(ctx workflow.Context, input *chime.DisassociatePhoneNumberFromUserInput) (*chime.DisassociatePhoneNumberFromUserOutput, error)
	DisassociatePhoneNumberFromUserAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumberFromUserInput) *DisassociatePhoneNumberFromUserFuture

	DisassociatePhoneNumbersFromVoiceConnector(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error)
	DisassociatePhoneNumbersFromVoiceConnectorAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) *DisassociatePhoneNumbersFromVoiceConnectorFuture

	DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error)
	DisassociatePhoneNumbersFromVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) *DisassociatePhoneNumbersFromVoiceConnectorGroupFuture

	DisassociateSigninDelegateGroupsFromAccount(ctx workflow.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error)
	DisassociateSigninDelegateGroupsFromAccountAsync(ctx workflow.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) *DisassociateSigninDelegateGroupsFromAccountFuture

	GetAccount(ctx workflow.Context, input *chime.GetAccountInput) (*chime.GetAccountOutput, error)
	GetAccountAsync(ctx workflow.Context, input *chime.GetAccountInput) *GetAccountFuture

	GetAccountSettings(ctx workflow.Context, input *chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error)
	GetAccountSettingsAsync(ctx workflow.Context, input *chime.GetAccountSettingsInput) *GetAccountSettingsFuture

	GetAttendee(ctx workflow.Context, input *chime.GetAttendeeInput) (*chime.GetAttendeeOutput, error)
	GetAttendeeAsync(ctx workflow.Context, input *chime.GetAttendeeInput) *GetAttendeeFuture

	GetBot(ctx workflow.Context, input *chime.GetBotInput) (*chime.GetBotOutput, error)
	GetBotAsync(ctx workflow.Context, input *chime.GetBotInput) *GetBotFuture

	GetEventsConfiguration(ctx workflow.Context, input *chime.GetEventsConfigurationInput) (*chime.GetEventsConfigurationOutput, error)
	GetEventsConfigurationAsync(ctx workflow.Context, input *chime.GetEventsConfigurationInput) *GetEventsConfigurationFuture

	GetGlobalSettings(ctx workflow.Context, input *chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error)
	GetGlobalSettingsAsync(ctx workflow.Context, input *chime.GetGlobalSettingsInput) *GetGlobalSettingsFuture

	GetMeeting(ctx workflow.Context, input *chime.GetMeetingInput) (*chime.GetMeetingOutput, error)
	GetMeetingAsync(ctx workflow.Context, input *chime.GetMeetingInput) *GetMeetingFuture

	GetPhoneNumber(ctx workflow.Context, input *chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error)
	GetPhoneNumberAsync(ctx workflow.Context, input *chime.GetPhoneNumberInput) *GetPhoneNumberFuture

	GetPhoneNumberOrder(ctx workflow.Context, input *chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error)
	GetPhoneNumberOrderAsync(ctx workflow.Context, input *chime.GetPhoneNumberOrderInput) *GetPhoneNumberOrderFuture

	GetPhoneNumberSettings(ctx workflow.Context, input *chime.GetPhoneNumberSettingsInput) (*chime.GetPhoneNumberSettingsOutput, error)
	GetPhoneNumberSettingsAsync(ctx workflow.Context, input *chime.GetPhoneNumberSettingsInput) *GetPhoneNumberSettingsFuture

	GetProxySession(ctx workflow.Context, input *chime.GetProxySessionInput) (*chime.GetProxySessionOutput, error)
	GetProxySessionAsync(ctx workflow.Context, input *chime.GetProxySessionInput) *GetProxySessionFuture

	GetRetentionSettings(ctx workflow.Context, input *chime.GetRetentionSettingsInput) (*chime.GetRetentionSettingsOutput, error)
	GetRetentionSettingsAsync(ctx workflow.Context, input *chime.GetRetentionSettingsInput) *GetRetentionSettingsFuture

	GetRoom(ctx workflow.Context, input *chime.GetRoomInput) (*chime.GetRoomOutput, error)
	GetRoomAsync(ctx workflow.Context, input *chime.GetRoomInput) *GetRoomFuture

	GetUser(ctx workflow.Context, input *chime.GetUserInput) (*chime.GetUserOutput, error)
	GetUserAsync(ctx workflow.Context, input *chime.GetUserInput) *GetUserFuture

	GetUserSettings(ctx workflow.Context, input *chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error)
	GetUserSettingsAsync(ctx workflow.Context, input *chime.GetUserSettingsInput) *GetUserSettingsFuture

	GetVoiceConnector(ctx workflow.Context, input *chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error)
	GetVoiceConnectorAsync(ctx workflow.Context, input *chime.GetVoiceConnectorInput) *GetVoiceConnectorFuture

	GetVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error)
	GetVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) *GetVoiceConnectorEmergencyCallingConfigurationFuture

	GetVoiceConnectorGroup(ctx workflow.Context, input *chime.GetVoiceConnectorGroupInput) (*chime.GetVoiceConnectorGroupOutput, error)
	GetVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.GetVoiceConnectorGroupInput) *GetVoiceConnectorGroupFuture

	GetVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error)
	GetVoiceConnectorLoggingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) *GetVoiceConnectorLoggingConfigurationFuture

	GetVoiceConnectorOrigination(ctx workflow.Context, input *chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error)
	GetVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorOriginationInput) *GetVoiceConnectorOriginationFuture

	GetVoiceConnectorProxy(ctx workflow.Context, input *chime.GetVoiceConnectorProxyInput) (*chime.GetVoiceConnectorProxyOutput, error)
	GetVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.GetVoiceConnectorProxyInput) *GetVoiceConnectorProxyFuture

	GetVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error)
	GetVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) *GetVoiceConnectorStreamingConfigurationFuture

	GetVoiceConnectorTermination(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error)
	GetVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationInput) *GetVoiceConnectorTerminationFuture

	GetVoiceConnectorTerminationHealth(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error)
	GetVoiceConnectorTerminationHealthAsync(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationHealthInput) *GetVoiceConnectorTerminationHealthFuture

	InviteUsers(ctx workflow.Context, input *chime.InviteUsersInput) (*chime.InviteUsersOutput, error)
	InviteUsersAsync(ctx workflow.Context, input *chime.InviteUsersInput) *InviteUsersFuture

	ListAccounts(ctx workflow.Context, input *chime.ListAccountsInput) (*chime.ListAccountsOutput, error)
	ListAccountsAsync(ctx workflow.Context, input *chime.ListAccountsInput) *ListAccountsFuture

	ListAttendeeTags(ctx workflow.Context, input *chime.ListAttendeeTagsInput) (*chime.ListAttendeeTagsOutput, error)
	ListAttendeeTagsAsync(ctx workflow.Context, input *chime.ListAttendeeTagsInput) *ListAttendeeTagsFuture

	ListAttendees(ctx workflow.Context, input *chime.ListAttendeesInput) (*chime.ListAttendeesOutput, error)
	ListAttendeesAsync(ctx workflow.Context, input *chime.ListAttendeesInput) *ListAttendeesFuture

	ListBots(ctx workflow.Context, input *chime.ListBotsInput) (*chime.ListBotsOutput, error)
	ListBotsAsync(ctx workflow.Context, input *chime.ListBotsInput) *ListBotsFuture

	ListMeetingTags(ctx workflow.Context, input *chime.ListMeetingTagsInput) (*chime.ListMeetingTagsOutput, error)
	ListMeetingTagsAsync(ctx workflow.Context, input *chime.ListMeetingTagsInput) *ListMeetingTagsFuture

	ListMeetings(ctx workflow.Context, input *chime.ListMeetingsInput) (*chime.ListMeetingsOutput, error)
	ListMeetingsAsync(ctx workflow.Context, input *chime.ListMeetingsInput) *ListMeetingsFuture

	ListPhoneNumberOrders(ctx workflow.Context, input *chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error)
	ListPhoneNumberOrdersAsync(ctx workflow.Context, input *chime.ListPhoneNumberOrdersInput) *ListPhoneNumberOrdersFuture

	ListPhoneNumbers(ctx workflow.Context, input *chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error)
	ListPhoneNumbersAsync(ctx workflow.Context, input *chime.ListPhoneNumbersInput) *ListPhoneNumbersFuture

	ListProxySessions(ctx workflow.Context, input *chime.ListProxySessionsInput) (*chime.ListProxySessionsOutput, error)
	ListProxySessionsAsync(ctx workflow.Context, input *chime.ListProxySessionsInput) *ListProxySessionsFuture

	ListRoomMemberships(ctx workflow.Context, input *chime.ListRoomMembershipsInput) (*chime.ListRoomMembershipsOutput, error)
	ListRoomMembershipsAsync(ctx workflow.Context, input *chime.ListRoomMembershipsInput) *ListRoomMembershipsFuture

	ListRooms(ctx workflow.Context, input *chime.ListRoomsInput) (*chime.ListRoomsOutput, error)
	ListRoomsAsync(ctx workflow.Context, input *chime.ListRoomsInput) *ListRoomsFuture

	ListTagsForResource(ctx workflow.Context, input *chime.ListTagsForResourceInput) (*chime.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *chime.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListUsers(ctx workflow.Context, input *chime.ListUsersInput) (*chime.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *chime.ListUsersInput) *ListUsersFuture

	ListVoiceConnectorGroups(ctx workflow.Context, input *chime.ListVoiceConnectorGroupsInput) (*chime.ListVoiceConnectorGroupsOutput, error)
	ListVoiceConnectorGroupsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorGroupsInput) *ListVoiceConnectorGroupsFuture

	ListVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error)
	ListVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) *ListVoiceConnectorTerminationCredentialsFuture

	ListVoiceConnectors(ctx workflow.Context, input *chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error)
	ListVoiceConnectorsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorsInput) *ListVoiceConnectorsFuture

	LogoutUser(ctx workflow.Context, input *chime.LogoutUserInput) (*chime.LogoutUserOutput, error)
	LogoutUserAsync(ctx workflow.Context, input *chime.LogoutUserInput) *LogoutUserFuture

	PutEventsConfiguration(ctx workflow.Context, input *chime.PutEventsConfigurationInput) (*chime.PutEventsConfigurationOutput, error)
	PutEventsConfigurationAsync(ctx workflow.Context, input *chime.PutEventsConfigurationInput) *PutEventsConfigurationFuture

	PutRetentionSettings(ctx workflow.Context, input *chime.PutRetentionSettingsInput) (*chime.PutRetentionSettingsOutput, error)
	PutRetentionSettingsAsync(ctx workflow.Context, input *chime.PutRetentionSettingsInput) *PutRetentionSettingsFuture

	PutVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error)
	PutVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) *PutVoiceConnectorEmergencyCallingConfigurationFuture

	PutVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error)
	PutVoiceConnectorLoggingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) *PutVoiceConnectorLoggingConfigurationFuture

	PutVoiceConnectorOrigination(ctx workflow.Context, input *chime.PutVoiceConnectorOriginationInput) (*chime.PutVoiceConnectorOriginationOutput, error)
	PutVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorOriginationInput) *PutVoiceConnectorOriginationFuture

	PutVoiceConnectorProxy(ctx workflow.Context, input *chime.PutVoiceConnectorProxyInput) (*chime.PutVoiceConnectorProxyOutput, error)
	PutVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.PutVoiceConnectorProxyInput) *PutVoiceConnectorProxyFuture

	PutVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error)
	PutVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) *PutVoiceConnectorStreamingConfigurationFuture

	PutVoiceConnectorTermination(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationInput) (*chime.PutVoiceConnectorTerminationOutput, error)
	PutVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationInput) *PutVoiceConnectorTerminationFuture

	PutVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error)
	PutVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) *PutVoiceConnectorTerminationCredentialsFuture

	RedactConversationMessage(ctx workflow.Context, input *chime.RedactConversationMessageInput) (*chime.RedactConversationMessageOutput, error)
	RedactConversationMessageAsync(ctx workflow.Context, input *chime.RedactConversationMessageInput) *RedactConversationMessageFuture

	RedactRoomMessage(ctx workflow.Context, input *chime.RedactRoomMessageInput) (*chime.RedactRoomMessageOutput, error)
	RedactRoomMessageAsync(ctx workflow.Context, input *chime.RedactRoomMessageInput) *RedactRoomMessageFuture

	RegenerateSecurityToken(ctx workflow.Context, input *chime.RegenerateSecurityTokenInput) (*chime.RegenerateSecurityTokenOutput, error)
	RegenerateSecurityTokenAsync(ctx workflow.Context, input *chime.RegenerateSecurityTokenInput) *RegenerateSecurityTokenFuture

	ResetPersonalPIN(ctx workflow.Context, input *chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error)
	ResetPersonalPINAsync(ctx workflow.Context, input *chime.ResetPersonalPINInput) *ResetPersonalPINFuture

	RestorePhoneNumber(ctx workflow.Context, input *chime.RestorePhoneNumberInput) (*chime.RestorePhoneNumberOutput, error)
	RestorePhoneNumberAsync(ctx workflow.Context, input *chime.RestorePhoneNumberInput) *RestorePhoneNumberFuture

	SearchAvailablePhoneNumbers(ctx workflow.Context, input *chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error)
	SearchAvailablePhoneNumbersAsync(ctx workflow.Context, input *chime.SearchAvailablePhoneNumbersInput) *SearchAvailablePhoneNumbersFuture

	TagAttendee(ctx workflow.Context, input *chime.TagAttendeeInput) (*chime.TagAttendeeOutput, error)
	TagAttendeeAsync(ctx workflow.Context, input *chime.TagAttendeeInput) *TagAttendeeFuture

	TagMeeting(ctx workflow.Context, input *chime.TagMeetingInput) (*chime.TagMeetingOutput, error)
	TagMeetingAsync(ctx workflow.Context, input *chime.TagMeetingInput) *TagMeetingFuture

	TagResource(ctx workflow.Context, input *chime.TagResourceInput) (*chime.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *chime.TagResourceInput) *TagResourceFuture

	UntagAttendee(ctx workflow.Context, input *chime.UntagAttendeeInput) (*chime.UntagAttendeeOutput, error)
	UntagAttendeeAsync(ctx workflow.Context, input *chime.UntagAttendeeInput) *UntagAttendeeFuture

	UntagMeeting(ctx workflow.Context, input *chime.UntagMeetingInput) (*chime.UntagMeetingOutput, error)
	UntagMeetingAsync(ctx workflow.Context, input *chime.UntagMeetingInput) *UntagMeetingFuture

	UntagResource(ctx workflow.Context, input *chime.UntagResourceInput) (*chime.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *chime.UntagResourceInput) *UntagResourceFuture

	UpdateAccount(ctx workflow.Context, input *chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error)
	UpdateAccountAsync(ctx workflow.Context, input *chime.UpdateAccountInput) *UpdateAccountFuture

	UpdateAccountSettings(ctx workflow.Context, input *chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error)
	UpdateAccountSettingsAsync(ctx workflow.Context, input *chime.UpdateAccountSettingsInput) *UpdateAccountSettingsFuture

	UpdateBot(ctx workflow.Context, input *chime.UpdateBotInput) (*chime.UpdateBotOutput, error)
	UpdateBotAsync(ctx workflow.Context, input *chime.UpdateBotInput) *UpdateBotFuture

	UpdateGlobalSettings(ctx workflow.Context, input *chime.UpdateGlobalSettingsInput) (*chime.UpdateGlobalSettingsOutput, error)
	UpdateGlobalSettingsAsync(ctx workflow.Context, input *chime.UpdateGlobalSettingsInput) *UpdateGlobalSettingsFuture

	UpdatePhoneNumber(ctx workflow.Context, input *chime.UpdatePhoneNumberInput) (*chime.UpdatePhoneNumberOutput, error)
	UpdatePhoneNumberAsync(ctx workflow.Context, input *chime.UpdatePhoneNumberInput) *UpdatePhoneNumberFuture

	UpdatePhoneNumberSettings(ctx workflow.Context, input *chime.UpdatePhoneNumberSettingsInput) (*chime.UpdatePhoneNumberSettingsOutput, error)
	UpdatePhoneNumberSettingsAsync(ctx workflow.Context, input *chime.UpdatePhoneNumberSettingsInput) *UpdatePhoneNumberSettingsFuture

	UpdateProxySession(ctx workflow.Context, input *chime.UpdateProxySessionInput) (*chime.UpdateProxySessionOutput, error)
	UpdateProxySessionAsync(ctx workflow.Context, input *chime.UpdateProxySessionInput) *UpdateProxySessionFuture

	UpdateRoom(ctx workflow.Context, input *chime.UpdateRoomInput) (*chime.UpdateRoomOutput, error)
	UpdateRoomAsync(ctx workflow.Context, input *chime.UpdateRoomInput) *UpdateRoomFuture

	UpdateRoomMembership(ctx workflow.Context, input *chime.UpdateRoomMembershipInput) (*chime.UpdateRoomMembershipOutput, error)
	UpdateRoomMembershipAsync(ctx workflow.Context, input *chime.UpdateRoomMembershipInput) *UpdateRoomMembershipFuture

	UpdateUser(ctx workflow.Context, input *chime.UpdateUserInput) (*chime.UpdateUserOutput, error)
	UpdateUserAsync(ctx workflow.Context, input *chime.UpdateUserInput) *UpdateUserFuture

	UpdateUserSettings(ctx workflow.Context, input *chime.UpdateUserSettingsInput) (*chime.UpdateUserSettingsOutput, error)
	UpdateUserSettingsAsync(ctx workflow.Context, input *chime.UpdateUserSettingsInput) *UpdateUserSettingsFuture

	UpdateVoiceConnector(ctx workflow.Context, input *chime.UpdateVoiceConnectorInput) (*chime.UpdateVoiceConnectorOutput, error)
	UpdateVoiceConnectorAsync(ctx workflow.Context, input *chime.UpdateVoiceConnectorInput) *UpdateVoiceConnectorFuture

	UpdateVoiceConnectorGroup(ctx workflow.Context, input *chime.UpdateVoiceConnectorGroupInput) (*chime.UpdateVoiceConnectorGroupOutput, error)
	UpdateVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.UpdateVoiceConnectorGroupInput) *UpdateVoiceConnectorGroupFuture
}

func NewClient() Client {
	return &stub{}
}
