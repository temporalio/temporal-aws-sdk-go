package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/chime"
	"go.temporal.io/sdk/workflow"
)

type ChimeClient interface {
    AssociatePhoneNumberWithUser(ctx workflow.Context, input *chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error)
    AssociatePhoneNumberWithUserAsync(ctx workflow.Context, input *chime.AssociatePhoneNumberWithUserInput) *ChimeAssociatePhoneNumberWithUserResult

    AssociatePhoneNumbersWithVoiceConnector(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error)
    AssociatePhoneNumbersWithVoiceConnectorAsync(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) *ChimeAssociatePhoneNumbersWithVoiceConnectorResult

    AssociatePhoneNumbersWithVoiceConnectorGroup(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error)
    AssociatePhoneNumbersWithVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) *ChimeAssociatePhoneNumbersWithVoiceConnectorGroupResult

    AssociateSigninDelegateGroupsWithAccount(ctx workflow.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error)
    AssociateSigninDelegateGroupsWithAccountAsync(ctx workflow.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) *ChimeAssociateSigninDelegateGroupsWithAccountResult

    BatchCreateAttendee(ctx workflow.Context, input *chime.BatchCreateAttendeeInput) (*chime.BatchCreateAttendeeOutput, error)
    BatchCreateAttendeeAsync(ctx workflow.Context, input *chime.BatchCreateAttendeeInput) *ChimeBatchCreateAttendeeResult

    BatchCreateRoomMembership(ctx workflow.Context, input *chime.BatchCreateRoomMembershipInput) (*chime.BatchCreateRoomMembershipOutput, error)
    BatchCreateRoomMembershipAsync(ctx workflow.Context, input *chime.BatchCreateRoomMembershipInput) *ChimeBatchCreateRoomMembershipResult

    BatchDeletePhoneNumber(ctx workflow.Context, input *chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error)
    BatchDeletePhoneNumberAsync(ctx workflow.Context, input *chime.BatchDeletePhoneNumberInput) *ChimeBatchDeletePhoneNumberResult

    BatchSuspendUser(ctx workflow.Context, input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error)
    BatchSuspendUserAsync(ctx workflow.Context, input *chime.BatchSuspendUserInput) *ChimeBatchSuspendUserResult

    BatchUnsuspendUser(ctx workflow.Context, input *chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error)
    BatchUnsuspendUserAsync(ctx workflow.Context, input *chime.BatchUnsuspendUserInput) *ChimeBatchUnsuspendUserResult

    BatchUpdatePhoneNumber(ctx workflow.Context, input *chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error)
    BatchUpdatePhoneNumberAsync(ctx workflow.Context, input *chime.BatchUpdatePhoneNumberInput) *ChimeBatchUpdatePhoneNumberResult

    BatchUpdateUser(ctx workflow.Context, input *chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error)
    BatchUpdateUserAsync(ctx workflow.Context, input *chime.BatchUpdateUserInput) *ChimeBatchUpdateUserResult

    CreateAccount(ctx workflow.Context, input *chime.CreateAccountInput) (*chime.CreateAccountOutput, error)
    CreateAccountAsync(ctx workflow.Context, input *chime.CreateAccountInput) *ChimeCreateAccountResult

    CreateAttendee(ctx workflow.Context, input *chime.CreateAttendeeInput) (*chime.CreateAttendeeOutput, error)
    CreateAttendeeAsync(ctx workflow.Context, input *chime.CreateAttendeeInput) *ChimeCreateAttendeeResult

    CreateBot(ctx workflow.Context, input *chime.CreateBotInput) (*chime.CreateBotOutput, error)
    CreateBotAsync(ctx workflow.Context, input *chime.CreateBotInput) *ChimeCreateBotResult

    CreateMeeting(ctx workflow.Context, input *chime.CreateMeetingInput) (*chime.CreateMeetingOutput, error)
    CreateMeetingAsync(ctx workflow.Context, input *chime.CreateMeetingInput) *ChimeCreateMeetingResult

    CreateMeetingWithAttendees(ctx workflow.Context, input *chime.CreateMeetingWithAttendeesInput) (*chime.CreateMeetingWithAttendeesOutput, error)
    CreateMeetingWithAttendeesAsync(ctx workflow.Context, input *chime.CreateMeetingWithAttendeesInput) *ChimeCreateMeetingWithAttendeesResult

    CreatePhoneNumberOrder(ctx workflow.Context, input *chime.CreatePhoneNumberOrderInput) (*chime.CreatePhoneNumberOrderOutput, error)
    CreatePhoneNumberOrderAsync(ctx workflow.Context, input *chime.CreatePhoneNumberOrderInput) *ChimeCreatePhoneNumberOrderResult

    CreateProxySession(ctx workflow.Context, input *chime.CreateProxySessionInput) (*chime.CreateProxySessionOutput, error)
    CreateProxySessionAsync(ctx workflow.Context, input *chime.CreateProxySessionInput) *ChimeCreateProxySessionResult

    CreateRoom(ctx workflow.Context, input *chime.CreateRoomInput) (*chime.CreateRoomOutput, error)
    CreateRoomAsync(ctx workflow.Context, input *chime.CreateRoomInput) *ChimeCreateRoomResult

    CreateRoomMembership(ctx workflow.Context, input *chime.CreateRoomMembershipInput) (*chime.CreateRoomMembershipOutput, error)
    CreateRoomMembershipAsync(ctx workflow.Context, input *chime.CreateRoomMembershipInput) *ChimeCreateRoomMembershipResult

    CreateUser(ctx workflow.Context, input *chime.CreateUserInput) (*chime.CreateUserOutput, error)
    CreateUserAsync(ctx workflow.Context, input *chime.CreateUserInput) *ChimeCreateUserResult

    CreateVoiceConnector(ctx workflow.Context, input *chime.CreateVoiceConnectorInput) (*chime.CreateVoiceConnectorOutput, error)
    CreateVoiceConnectorAsync(ctx workflow.Context, input *chime.CreateVoiceConnectorInput) *ChimeCreateVoiceConnectorResult

    CreateVoiceConnectorGroup(ctx workflow.Context, input *chime.CreateVoiceConnectorGroupInput) (*chime.CreateVoiceConnectorGroupOutput, error)
    CreateVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.CreateVoiceConnectorGroupInput) *ChimeCreateVoiceConnectorGroupResult

    DeleteAccount(ctx workflow.Context, input *chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error)
    DeleteAccountAsync(ctx workflow.Context, input *chime.DeleteAccountInput) *ChimeDeleteAccountResult

    DeleteAttendee(ctx workflow.Context, input *chime.DeleteAttendeeInput) (*chime.DeleteAttendeeOutput, error)
    DeleteAttendeeAsync(ctx workflow.Context, input *chime.DeleteAttendeeInput) *ChimeDeleteAttendeeResult

    DeleteEventsConfiguration(ctx workflow.Context, input *chime.DeleteEventsConfigurationInput) (*chime.DeleteEventsConfigurationOutput, error)
    DeleteEventsConfigurationAsync(ctx workflow.Context, input *chime.DeleteEventsConfigurationInput) *ChimeDeleteEventsConfigurationResult

    DeleteMeeting(ctx workflow.Context, input *chime.DeleteMeetingInput) (*chime.DeleteMeetingOutput, error)
    DeleteMeetingAsync(ctx workflow.Context, input *chime.DeleteMeetingInput) *ChimeDeleteMeetingResult

    DeletePhoneNumber(ctx workflow.Context, input *chime.DeletePhoneNumberInput) (*chime.DeletePhoneNumberOutput, error)
    DeletePhoneNumberAsync(ctx workflow.Context, input *chime.DeletePhoneNumberInput) *ChimeDeletePhoneNumberResult

    DeleteProxySession(ctx workflow.Context, input *chime.DeleteProxySessionInput) (*chime.DeleteProxySessionOutput, error)
    DeleteProxySessionAsync(ctx workflow.Context, input *chime.DeleteProxySessionInput) *ChimeDeleteProxySessionResult

    DeleteRoom(ctx workflow.Context, input *chime.DeleteRoomInput) (*chime.DeleteRoomOutput, error)
    DeleteRoomAsync(ctx workflow.Context, input *chime.DeleteRoomInput) *ChimeDeleteRoomResult

    DeleteRoomMembership(ctx workflow.Context, input *chime.DeleteRoomMembershipInput) (*chime.DeleteRoomMembershipOutput, error)
    DeleteRoomMembershipAsync(ctx workflow.Context, input *chime.DeleteRoomMembershipInput) *ChimeDeleteRoomMembershipResult

    DeleteVoiceConnector(ctx workflow.Context, input *chime.DeleteVoiceConnectorInput) (*chime.DeleteVoiceConnectorOutput, error)
    DeleteVoiceConnectorAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorInput) *ChimeDeleteVoiceConnectorResult

    DeleteVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error)
    DeleteVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) *ChimeDeleteVoiceConnectorEmergencyCallingConfigurationResult

    DeleteVoiceConnectorGroup(ctx workflow.Context, input *chime.DeleteVoiceConnectorGroupInput) (*chime.DeleteVoiceConnectorGroupOutput, error)
    DeleteVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorGroupInput) *ChimeDeleteVoiceConnectorGroupResult

    DeleteVoiceConnectorOrigination(ctx workflow.Context, input *chime.DeleteVoiceConnectorOriginationInput) (*chime.DeleteVoiceConnectorOriginationOutput, error)
    DeleteVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorOriginationInput) *ChimeDeleteVoiceConnectorOriginationResult

    DeleteVoiceConnectorProxy(ctx workflow.Context, input *chime.DeleteVoiceConnectorProxyInput) (*chime.DeleteVoiceConnectorProxyOutput, error)
    DeleteVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorProxyInput) *ChimeDeleteVoiceConnectorProxyResult

    DeleteVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error)
    DeleteVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) *ChimeDeleteVoiceConnectorStreamingConfigurationResult

    DeleteVoiceConnectorTermination(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationInput) (*chime.DeleteVoiceConnectorTerminationOutput, error)
    DeleteVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationInput) *ChimeDeleteVoiceConnectorTerminationResult

    DeleteVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error)
    DeleteVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) *ChimeDeleteVoiceConnectorTerminationCredentialsResult

    DisassociatePhoneNumberFromUser(ctx workflow.Context, input *chime.DisassociatePhoneNumberFromUserInput) (*chime.DisassociatePhoneNumberFromUserOutput, error)
    DisassociatePhoneNumberFromUserAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumberFromUserInput) *ChimeDisassociatePhoneNumberFromUserResult

    DisassociatePhoneNumbersFromVoiceConnector(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error)
    DisassociatePhoneNumbersFromVoiceConnectorAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) *ChimeDisassociatePhoneNumbersFromVoiceConnectorResult

    DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error)
    DisassociatePhoneNumbersFromVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) *ChimeDisassociatePhoneNumbersFromVoiceConnectorGroupResult

    DisassociateSigninDelegateGroupsFromAccount(ctx workflow.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error)
    DisassociateSigninDelegateGroupsFromAccountAsync(ctx workflow.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) *ChimeDisassociateSigninDelegateGroupsFromAccountResult

    GetAccount(ctx workflow.Context, input *chime.GetAccountInput) (*chime.GetAccountOutput, error)
    GetAccountAsync(ctx workflow.Context, input *chime.GetAccountInput) *ChimeGetAccountResult

    GetAccountSettings(ctx workflow.Context, input *chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error)
    GetAccountSettingsAsync(ctx workflow.Context, input *chime.GetAccountSettingsInput) *ChimeGetAccountSettingsResult

    GetAttendee(ctx workflow.Context, input *chime.GetAttendeeInput) (*chime.GetAttendeeOutput, error)
    GetAttendeeAsync(ctx workflow.Context, input *chime.GetAttendeeInput) *ChimeGetAttendeeResult

    GetBot(ctx workflow.Context, input *chime.GetBotInput) (*chime.GetBotOutput, error)
    GetBotAsync(ctx workflow.Context, input *chime.GetBotInput) *ChimeGetBotResult

    GetEventsConfiguration(ctx workflow.Context, input *chime.GetEventsConfigurationInput) (*chime.GetEventsConfigurationOutput, error)
    GetEventsConfigurationAsync(ctx workflow.Context, input *chime.GetEventsConfigurationInput) *ChimeGetEventsConfigurationResult

    GetGlobalSettings(ctx workflow.Context, input *chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error)
    GetGlobalSettingsAsync(ctx workflow.Context, input *chime.GetGlobalSettingsInput) *ChimeGetGlobalSettingsResult

    GetMeeting(ctx workflow.Context, input *chime.GetMeetingInput) (*chime.GetMeetingOutput, error)
    GetMeetingAsync(ctx workflow.Context, input *chime.GetMeetingInput) *ChimeGetMeetingResult

    GetPhoneNumber(ctx workflow.Context, input *chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error)
    GetPhoneNumberAsync(ctx workflow.Context, input *chime.GetPhoneNumberInput) *ChimeGetPhoneNumberResult

    GetPhoneNumberOrder(ctx workflow.Context, input *chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error)
    GetPhoneNumberOrderAsync(ctx workflow.Context, input *chime.GetPhoneNumberOrderInput) *ChimeGetPhoneNumberOrderResult

    GetPhoneNumberSettings(ctx workflow.Context, input *chime.GetPhoneNumberSettingsInput) (*chime.GetPhoneNumberSettingsOutput, error)
    GetPhoneNumberSettingsAsync(ctx workflow.Context, input *chime.GetPhoneNumberSettingsInput) *ChimeGetPhoneNumberSettingsResult

    GetProxySession(ctx workflow.Context, input *chime.GetProxySessionInput) (*chime.GetProxySessionOutput, error)
    GetProxySessionAsync(ctx workflow.Context, input *chime.GetProxySessionInput) *ChimeGetProxySessionResult

    GetRetentionSettings(ctx workflow.Context, input *chime.GetRetentionSettingsInput) (*chime.GetRetentionSettingsOutput, error)
    GetRetentionSettingsAsync(ctx workflow.Context, input *chime.GetRetentionSettingsInput) *ChimeGetRetentionSettingsResult

    GetRoom(ctx workflow.Context, input *chime.GetRoomInput) (*chime.GetRoomOutput, error)
    GetRoomAsync(ctx workflow.Context, input *chime.GetRoomInput) *ChimeGetRoomResult

    GetUser(ctx workflow.Context, input *chime.GetUserInput) (*chime.GetUserOutput, error)
    GetUserAsync(ctx workflow.Context, input *chime.GetUserInput) *ChimeGetUserResult

    GetUserSettings(ctx workflow.Context, input *chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error)
    GetUserSettingsAsync(ctx workflow.Context, input *chime.GetUserSettingsInput) *ChimeGetUserSettingsResult

    GetVoiceConnector(ctx workflow.Context, input *chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error)
    GetVoiceConnectorAsync(ctx workflow.Context, input *chime.GetVoiceConnectorInput) *ChimeGetVoiceConnectorResult

    GetVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error)
    GetVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) *ChimeGetVoiceConnectorEmergencyCallingConfigurationResult

    GetVoiceConnectorGroup(ctx workflow.Context, input *chime.GetVoiceConnectorGroupInput) (*chime.GetVoiceConnectorGroupOutput, error)
    GetVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.GetVoiceConnectorGroupInput) *ChimeGetVoiceConnectorGroupResult

    GetVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error)
    GetVoiceConnectorLoggingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) *ChimeGetVoiceConnectorLoggingConfigurationResult

    GetVoiceConnectorOrigination(ctx workflow.Context, input *chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error)
    GetVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorOriginationInput) *ChimeGetVoiceConnectorOriginationResult

    GetVoiceConnectorProxy(ctx workflow.Context, input *chime.GetVoiceConnectorProxyInput) (*chime.GetVoiceConnectorProxyOutput, error)
    GetVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.GetVoiceConnectorProxyInput) *ChimeGetVoiceConnectorProxyResult

    GetVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error)
    GetVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) *ChimeGetVoiceConnectorStreamingConfigurationResult

    GetVoiceConnectorTermination(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error)
    GetVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationInput) *ChimeGetVoiceConnectorTerminationResult

    GetVoiceConnectorTerminationHealth(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error)
    GetVoiceConnectorTerminationHealthAsync(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationHealthInput) *ChimeGetVoiceConnectorTerminationHealthResult

    InviteUsers(ctx workflow.Context, input *chime.InviteUsersInput) (*chime.InviteUsersOutput, error)
    InviteUsersAsync(ctx workflow.Context, input *chime.InviteUsersInput) *ChimeInviteUsersResult

    ListAccounts(ctx workflow.Context, input *chime.ListAccountsInput) (*chime.ListAccountsOutput, error)
    ListAccountsAsync(ctx workflow.Context, input *chime.ListAccountsInput) *ChimeListAccountsResult

    ListAttendeeTags(ctx workflow.Context, input *chime.ListAttendeeTagsInput) (*chime.ListAttendeeTagsOutput, error)
    ListAttendeeTagsAsync(ctx workflow.Context, input *chime.ListAttendeeTagsInput) *ChimeListAttendeeTagsResult

    ListAttendees(ctx workflow.Context, input *chime.ListAttendeesInput) (*chime.ListAttendeesOutput, error)
    ListAttendeesAsync(ctx workflow.Context, input *chime.ListAttendeesInput) *ChimeListAttendeesResult

    ListBots(ctx workflow.Context, input *chime.ListBotsInput) (*chime.ListBotsOutput, error)
    ListBotsAsync(ctx workflow.Context, input *chime.ListBotsInput) *ChimeListBotsResult

    ListMeetingTags(ctx workflow.Context, input *chime.ListMeetingTagsInput) (*chime.ListMeetingTagsOutput, error)
    ListMeetingTagsAsync(ctx workflow.Context, input *chime.ListMeetingTagsInput) *ChimeListMeetingTagsResult

    ListMeetings(ctx workflow.Context, input *chime.ListMeetingsInput) (*chime.ListMeetingsOutput, error)
    ListMeetingsAsync(ctx workflow.Context, input *chime.ListMeetingsInput) *ChimeListMeetingsResult

    ListPhoneNumberOrders(ctx workflow.Context, input *chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error)
    ListPhoneNumberOrdersAsync(ctx workflow.Context, input *chime.ListPhoneNumberOrdersInput) *ChimeListPhoneNumberOrdersResult

    ListPhoneNumbers(ctx workflow.Context, input *chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error)
    ListPhoneNumbersAsync(ctx workflow.Context, input *chime.ListPhoneNumbersInput) *ChimeListPhoneNumbersResult

    ListProxySessions(ctx workflow.Context, input *chime.ListProxySessionsInput) (*chime.ListProxySessionsOutput, error)
    ListProxySessionsAsync(ctx workflow.Context, input *chime.ListProxySessionsInput) *ChimeListProxySessionsResult

    ListRoomMemberships(ctx workflow.Context, input *chime.ListRoomMembershipsInput) (*chime.ListRoomMembershipsOutput, error)
    ListRoomMembershipsAsync(ctx workflow.Context, input *chime.ListRoomMembershipsInput) *ChimeListRoomMembershipsResult

    ListRooms(ctx workflow.Context, input *chime.ListRoomsInput) (*chime.ListRoomsOutput, error)
    ListRoomsAsync(ctx workflow.Context, input *chime.ListRoomsInput) *ChimeListRoomsResult

    ListTagsForResource(ctx workflow.Context, input *chime.ListTagsForResourceInput) (*chime.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *chime.ListTagsForResourceInput) *ChimeListTagsForResourceResult

    ListUsers(ctx workflow.Context, input *chime.ListUsersInput) (*chime.ListUsersOutput, error)
    ListUsersAsync(ctx workflow.Context, input *chime.ListUsersInput) *ChimeListUsersResult

    ListVoiceConnectorGroups(ctx workflow.Context, input *chime.ListVoiceConnectorGroupsInput) (*chime.ListVoiceConnectorGroupsOutput, error)
    ListVoiceConnectorGroupsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorGroupsInput) *ChimeListVoiceConnectorGroupsResult

    ListVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error)
    ListVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) *ChimeListVoiceConnectorTerminationCredentialsResult

    ListVoiceConnectors(ctx workflow.Context, input *chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error)
    ListVoiceConnectorsAsync(ctx workflow.Context, input *chime.ListVoiceConnectorsInput) *ChimeListVoiceConnectorsResult

    LogoutUser(ctx workflow.Context, input *chime.LogoutUserInput) (*chime.LogoutUserOutput, error)
    LogoutUserAsync(ctx workflow.Context, input *chime.LogoutUserInput) *ChimeLogoutUserResult

    PutEventsConfiguration(ctx workflow.Context, input *chime.PutEventsConfigurationInput) (*chime.PutEventsConfigurationOutput, error)
    PutEventsConfigurationAsync(ctx workflow.Context, input *chime.PutEventsConfigurationInput) *ChimePutEventsConfigurationResult

    PutRetentionSettings(ctx workflow.Context, input *chime.PutRetentionSettingsInput) (*chime.PutRetentionSettingsOutput, error)
    PutRetentionSettingsAsync(ctx workflow.Context, input *chime.PutRetentionSettingsInput) *ChimePutRetentionSettingsResult

    PutVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error)
    PutVoiceConnectorEmergencyCallingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) *ChimePutVoiceConnectorEmergencyCallingConfigurationResult

    PutVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error)
    PutVoiceConnectorLoggingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) *ChimePutVoiceConnectorLoggingConfigurationResult

    PutVoiceConnectorOrigination(ctx workflow.Context, input *chime.PutVoiceConnectorOriginationInput) (*chime.PutVoiceConnectorOriginationOutput, error)
    PutVoiceConnectorOriginationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorOriginationInput) *ChimePutVoiceConnectorOriginationResult

    PutVoiceConnectorProxy(ctx workflow.Context, input *chime.PutVoiceConnectorProxyInput) (*chime.PutVoiceConnectorProxyOutput, error)
    PutVoiceConnectorProxyAsync(ctx workflow.Context, input *chime.PutVoiceConnectorProxyInput) *ChimePutVoiceConnectorProxyResult

    PutVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error)
    PutVoiceConnectorStreamingConfigurationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) *ChimePutVoiceConnectorStreamingConfigurationResult

    PutVoiceConnectorTermination(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationInput) (*chime.PutVoiceConnectorTerminationOutput, error)
    PutVoiceConnectorTerminationAsync(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationInput) *ChimePutVoiceConnectorTerminationResult

    PutVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error)
    PutVoiceConnectorTerminationCredentialsAsync(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) *ChimePutVoiceConnectorTerminationCredentialsResult

    RedactConversationMessage(ctx workflow.Context, input *chime.RedactConversationMessageInput) (*chime.RedactConversationMessageOutput, error)
    RedactConversationMessageAsync(ctx workflow.Context, input *chime.RedactConversationMessageInput) *ChimeRedactConversationMessageResult

    RedactRoomMessage(ctx workflow.Context, input *chime.RedactRoomMessageInput) (*chime.RedactRoomMessageOutput, error)
    RedactRoomMessageAsync(ctx workflow.Context, input *chime.RedactRoomMessageInput) *ChimeRedactRoomMessageResult

    RegenerateSecurityToken(ctx workflow.Context, input *chime.RegenerateSecurityTokenInput) (*chime.RegenerateSecurityTokenOutput, error)
    RegenerateSecurityTokenAsync(ctx workflow.Context, input *chime.RegenerateSecurityTokenInput) *ChimeRegenerateSecurityTokenResult

    ResetPersonalPIN(ctx workflow.Context, input *chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error)
    ResetPersonalPINAsync(ctx workflow.Context, input *chime.ResetPersonalPINInput) *ChimeResetPersonalPINResult

    RestorePhoneNumber(ctx workflow.Context, input *chime.RestorePhoneNumberInput) (*chime.RestorePhoneNumberOutput, error)
    RestorePhoneNumberAsync(ctx workflow.Context, input *chime.RestorePhoneNumberInput) *ChimeRestorePhoneNumberResult

    SearchAvailablePhoneNumbers(ctx workflow.Context, input *chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error)
    SearchAvailablePhoneNumbersAsync(ctx workflow.Context, input *chime.SearchAvailablePhoneNumbersInput) *ChimeSearchAvailablePhoneNumbersResult

    TagAttendee(ctx workflow.Context, input *chime.TagAttendeeInput) (*chime.TagAttendeeOutput, error)
    TagAttendeeAsync(ctx workflow.Context, input *chime.TagAttendeeInput) *ChimeTagAttendeeResult

    TagMeeting(ctx workflow.Context, input *chime.TagMeetingInput) (*chime.TagMeetingOutput, error)
    TagMeetingAsync(ctx workflow.Context, input *chime.TagMeetingInput) *ChimeTagMeetingResult

    TagResource(ctx workflow.Context, input *chime.TagResourceInput) (*chime.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *chime.TagResourceInput) *ChimeTagResourceResult

    UntagAttendee(ctx workflow.Context, input *chime.UntagAttendeeInput) (*chime.UntagAttendeeOutput, error)
    UntagAttendeeAsync(ctx workflow.Context, input *chime.UntagAttendeeInput) *ChimeUntagAttendeeResult

    UntagMeeting(ctx workflow.Context, input *chime.UntagMeetingInput) (*chime.UntagMeetingOutput, error)
    UntagMeetingAsync(ctx workflow.Context, input *chime.UntagMeetingInput) *ChimeUntagMeetingResult

    UntagResource(ctx workflow.Context, input *chime.UntagResourceInput) (*chime.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *chime.UntagResourceInput) *ChimeUntagResourceResult

    UpdateAccount(ctx workflow.Context, input *chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error)
    UpdateAccountAsync(ctx workflow.Context, input *chime.UpdateAccountInput) *ChimeUpdateAccountResult

    UpdateAccountSettings(ctx workflow.Context, input *chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error)
    UpdateAccountSettingsAsync(ctx workflow.Context, input *chime.UpdateAccountSettingsInput) *ChimeUpdateAccountSettingsResult

    UpdateBot(ctx workflow.Context, input *chime.UpdateBotInput) (*chime.UpdateBotOutput, error)
    UpdateBotAsync(ctx workflow.Context, input *chime.UpdateBotInput) *ChimeUpdateBotResult

    UpdateGlobalSettings(ctx workflow.Context, input *chime.UpdateGlobalSettingsInput) (*chime.UpdateGlobalSettingsOutput, error)
    UpdateGlobalSettingsAsync(ctx workflow.Context, input *chime.UpdateGlobalSettingsInput) *ChimeUpdateGlobalSettingsResult

    UpdatePhoneNumber(ctx workflow.Context, input *chime.UpdatePhoneNumberInput) (*chime.UpdatePhoneNumberOutput, error)
    UpdatePhoneNumberAsync(ctx workflow.Context, input *chime.UpdatePhoneNumberInput) *ChimeUpdatePhoneNumberResult

    UpdatePhoneNumberSettings(ctx workflow.Context, input *chime.UpdatePhoneNumberSettingsInput) (*chime.UpdatePhoneNumberSettingsOutput, error)
    UpdatePhoneNumberSettingsAsync(ctx workflow.Context, input *chime.UpdatePhoneNumberSettingsInput) *ChimeUpdatePhoneNumberSettingsResult

    UpdateProxySession(ctx workflow.Context, input *chime.UpdateProxySessionInput) (*chime.UpdateProxySessionOutput, error)
    UpdateProxySessionAsync(ctx workflow.Context, input *chime.UpdateProxySessionInput) *ChimeUpdateProxySessionResult

    UpdateRoom(ctx workflow.Context, input *chime.UpdateRoomInput) (*chime.UpdateRoomOutput, error)
    UpdateRoomAsync(ctx workflow.Context, input *chime.UpdateRoomInput) *ChimeUpdateRoomResult

    UpdateRoomMembership(ctx workflow.Context, input *chime.UpdateRoomMembershipInput) (*chime.UpdateRoomMembershipOutput, error)
    UpdateRoomMembershipAsync(ctx workflow.Context, input *chime.UpdateRoomMembershipInput) *ChimeUpdateRoomMembershipResult

    UpdateUser(ctx workflow.Context, input *chime.UpdateUserInput) (*chime.UpdateUserOutput, error)
    UpdateUserAsync(ctx workflow.Context, input *chime.UpdateUserInput) *ChimeUpdateUserResult

    UpdateUserSettings(ctx workflow.Context, input *chime.UpdateUserSettingsInput) (*chime.UpdateUserSettingsOutput, error)
    UpdateUserSettingsAsync(ctx workflow.Context, input *chime.UpdateUserSettingsInput) *ChimeUpdateUserSettingsResult

    UpdateVoiceConnector(ctx workflow.Context, input *chime.UpdateVoiceConnectorInput) (*chime.UpdateVoiceConnectorOutput, error)
    UpdateVoiceConnectorAsync(ctx workflow.Context, input *chime.UpdateVoiceConnectorInput) *ChimeUpdateVoiceConnectorResult

    UpdateVoiceConnectorGroup(ctx workflow.Context, input *chime.UpdateVoiceConnectorGroupInput) (*chime.UpdateVoiceConnectorGroupOutput, error)
    UpdateVoiceConnectorGroupAsync(ctx workflow.Context, input *chime.UpdateVoiceConnectorGroupInput) *ChimeUpdateVoiceConnectorGroupResult
}
type ChimeAssociatePhoneNumberWithUserResult struct {
	Result workflow.Future
}

func (r *ChimeAssociatePhoneNumberWithUserResult) Get(ctx workflow.Context) (*chime.AssociatePhoneNumberWithUserOutput, error) {
    var output chime.AssociatePhoneNumberWithUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeAssociatePhoneNumbersWithVoiceConnectorResult struct {
	Result workflow.Future
}

func (r *ChimeAssociatePhoneNumbersWithVoiceConnectorResult) Get(ctx workflow.Context) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error) {
    var output chime.AssociatePhoneNumbersWithVoiceConnectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeAssociatePhoneNumbersWithVoiceConnectorGroupResult struct {
	Result workflow.Future
}

func (r *ChimeAssociatePhoneNumbersWithVoiceConnectorGroupResult) Get(ctx workflow.Context) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) {
    var output chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeAssociateSigninDelegateGroupsWithAccountResult struct {
	Result workflow.Future
}

func (r *ChimeAssociateSigninDelegateGroupsWithAccountResult) Get(ctx workflow.Context) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error) {
    var output chime.AssociateSigninDelegateGroupsWithAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeBatchCreateAttendeeResult struct {
	Result workflow.Future
}

func (r *ChimeBatchCreateAttendeeResult) Get(ctx workflow.Context) (*chime.BatchCreateAttendeeOutput, error) {
    var output chime.BatchCreateAttendeeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeBatchCreateRoomMembershipResult struct {
	Result workflow.Future
}

func (r *ChimeBatchCreateRoomMembershipResult) Get(ctx workflow.Context) (*chime.BatchCreateRoomMembershipOutput, error) {
    var output chime.BatchCreateRoomMembershipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeBatchDeletePhoneNumberResult struct {
	Result workflow.Future
}

func (r *ChimeBatchDeletePhoneNumberResult) Get(ctx workflow.Context) (*chime.BatchDeletePhoneNumberOutput, error) {
    var output chime.BatchDeletePhoneNumberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeBatchSuspendUserResult struct {
	Result workflow.Future
}

func (r *ChimeBatchSuspendUserResult) Get(ctx workflow.Context) (*chime.BatchSuspendUserOutput, error) {
    var output chime.BatchSuspendUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeBatchUnsuspendUserResult struct {
	Result workflow.Future
}

func (r *ChimeBatchUnsuspendUserResult) Get(ctx workflow.Context) (*chime.BatchUnsuspendUserOutput, error) {
    var output chime.BatchUnsuspendUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeBatchUpdatePhoneNumberResult struct {
	Result workflow.Future
}

func (r *ChimeBatchUpdatePhoneNumberResult) Get(ctx workflow.Context) (*chime.BatchUpdatePhoneNumberOutput, error) {
    var output chime.BatchUpdatePhoneNumberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeBatchUpdateUserResult struct {
	Result workflow.Future
}

func (r *ChimeBatchUpdateUserResult) Get(ctx workflow.Context) (*chime.BatchUpdateUserOutput, error) {
    var output chime.BatchUpdateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateAccountResult struct {
	Result workflow.Future
}

func (r *ChimeCreateAccountResult) Get(ctx workflow.Context) (*chime.CreateAccountOutput, error) {
    var output chime.CreateAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateAttendeeResult struct {
	Result workflow.Future
}

func (r *ChimeCreateAttendeeResult) Get(ctx workflow.Context) (*chime.CreateAttendeeOutput, error) {
    var output chime.CreateAttendeeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateBotResult struct {
	Result workflow.Future
}

func (r *ChimeCreateBotResult) Get(ctx workflow.Context) (*chime.CreateBotOutput, error) {
    var output chime.CreateBotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateMeetingResult struct {
	Result workflow.Future
}

func (r *ChimeCreateMeetingResult) Get(ctx workflow.Context) (*chime.CreateMeetingOutput, error) {
    var output chime.CreateMeetingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateMeetingWithAttendeesResult struct {
	Result workflow.Future
}

func (r *ChimeCreateMeetingWithAttendeesResult) Get(ctx workflow.Context) (*chime.CreateMeetingWithAttendeesOutput, error) {
    var output chime.CreateMeetingWithAttendeesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreatePhoneNumberOrderResult struct {
	Result workflow.Future
}

func (r *ChimeCreatePhoneNumberOrderResult) Get(ctx workflow.Context) (*chime.CreatePhoneNumberOrderOutput, error) {
    var output chime.CreatePhoneNumberOrderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateProxySessionResult struct {
	Result workflow.Future
}

func (r *ChimeCreateProxySessionResult) Get(ctx workflow.Context) (*chime.CreateProxySessionOutput, error) {
    var output chime.CreateProxySessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateRoomResult struct {
	Result workflow.Future
}

func (r *ChimeCreateRoomResult) Get(ctx workflow.Context) (*chime.CreateRoomOutput, error) {
    var output chime.CreateRoomOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateRoomMembershipResult struct {
	Result workflow.Future
}

func (r *ChimeCreateRoomMembershipResult) Get(ctx workflow.Context) (*chime.CreateRoomMembershipOutput, error) {
    var output chime.CreateRoomMembershipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateUserResult struct {
	Result workflow.Future
}

func (r *ChimeCreateUserResult) Get(ctx workflow.Context) (*chime.CreateUserOutput, error) {
    var output chime.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateVoiceConnectorResult struct {
	Result workflow.Future
}

func (r *ChimeCreateVoiceConnectorResult) Get(ctx workflow.Context) (*chime.CreateVoiceConnectorOutput, error) {
    var output chime.CreateVoiceConnectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeCreateVoiceConnectorGroupResult struct {
	Result workflow.Future
}

func (r *ChimeCreateVoiceConnectorGroupResult) Get(ctx workflow.Context) (*chime.CreateVoiceConnectorGroupOutput, error) {
    var output chime.CreateVoiceConnectorGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteAccountResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteAccountResult) Get(ctx workflow.Context) (*chime.DeleteAccountOutput, error) {
    var output chime.DeleteAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteAttendeeResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteAttendeeResult) Get(ctx workflow.Context) (*chime.DeleteAttendeeOutput, error) {
    var output chime.DeleteAttendeeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteEventsConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteEventsConfigurationResult) Get(ctx workflow.Context) (*chime.DeleteEventsConfigurationOutput, error) {
    var output chime.DeleteEventsConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteMeetingResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteMeetingResult) Get(ctx workflow.Context) (*chime.DeleteMeetingOutput, error) {
    var output chime.DeleteMeetingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeletePhoneNumberResult struct {
	Result workflow.Future
}

func (r *ChimeDeletePhoneNumberResult) Get(ctx workflow.Context) (*chime.DeletePhoneNumberOutput, error) {
    var output chime.DeletePhoneNumberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteProxySessionResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteProxySessionResult) Get(ctx workflow.Context) (*chime.DeleteProxySessionOutput, error) {
    var output chime.DeleteProxySessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteRoomResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteRoomResult) Get(ctx workflow.Context) (*chime.DeleteRoomOutput, error) {
    var output chime.DeleteRoomOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteRoomMembershipResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteRoomMembershipResult) Get(ctx workflow.Context) (*chime.DeleteRoomMembershipOutput, error) {
    var output chime.DeleteRoomMembershipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorOutput, error) {
    var output chime.DeleteVoiceConnectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorEmergencyCallingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorEmergencyCallingConfigurationResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    var output chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorGroupResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorGroupResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorGroupOutput, error) {
    var output chime.DeleteVoiceConnectorGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorOriginationResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorOriginationResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorOriginationOutput, error) {
    var output chime.DeleteVoiceConnectorOriginationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorProxyResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorProxyResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorProxyOutput, error) {
    var output chime.DeleteVoiceConnectorProxyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorStreamingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorStreamingConfigurationResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error) {
    var output chime.DeleteVoiceConnectorStreamingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorTerminationResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorTerminationResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorTerminationOutput, error) {
    var output chime.DeleteVoiceConnectorTerminationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDeleteVoiceConnectorTerminationCredentialsResult struct {
	Result workflow.Future
}

func (r *ChimeDeleteVoiceConnectorTerminationCredentialsResult) Get(ctx workflow.Context) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error) {
    var output chime.DeleteVoiceConnectorTerminationCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDisassociatePhoneNumberFromUserResult struct {
	Result workflow.Future
}

func (r *ChimeDisassociatePhoneNumberFromUserResult) Get(ctx workflow.Context) (*chime.DisassociatePhoneNumberFromUserOutput, error) {
    var output chime.DisassociatePhoneNumberFromUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDisassociatePhoneNumbersFromVoiceConnectorResult struct {
	Result workflow.Future
}

func (r *ChimeDisassociatePhoneNumbersFromVoiceConnectorResult) Get(ctx workflow.Context) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error) {
    var output chime.DisassociatePhoneNumbersFromVoiceConnectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDisassociatePhoneNumbersFromVoiceConnectorGroupResult struct {
	Result workflow.Future
}

func (r *ChimeDisassociatePhoneNumbersFromVoiceConnectorGroupResult) Get(ctx workflow.Context) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error) {
    var output chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeDisassociateSigninDelegateGroupsFromAccountResult struct {
	Result workflow.Future
}

func (r *ChimeDisassociateSigninDelegateGroupsFromAccountResult) Get(ctx workflow.Context) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error) {
    var output chime.DisassociateSigninDelegateGroupsFromAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetAccountResult struct {
	Result workflow.Future
}

func (r *ChimeGetAccountResult) Get(ctx workflow.Context) (*chime.GetAccountOutput, error) {
    var output chime.GetAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetAccountSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeGetAccountSettingsResult) Get(ctx workflow.Context) (*chime.GetAccountSettingsOutput, error) {
    var output chime.GetAccountSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetAttendeeResult struct {
	Result workflow.Future
}

func (r *ChimeGetAttendeeResult) Get(ctx workflow.Context) (*chime.GetAttendeeOutput, error) {
    var output chime.GetAttendeeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetBotResult struct {
	Result workflow.Future
}

func (r *ChimeGetBotResult) Get(ctx workflow.Context) (*chime.GetBotOutput, error) {
    var output chime.GetBotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetEventsConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimeGetEventsConfigurationResult) Get(ctx workflow.Context) (*chime.GetEventsConfigurationOutput, error) {
    var output chime.GetEventsConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetGlobalSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeGetGlobalSettingsResult) Get(ctx workflow.Context) (*chime.GetGlobalSettingsOutput, error) {
    var output chime.GetGlobalSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetMeetingResult struct {
	Result workflow.Future
}

func (r *ChimeGetMeetingResult) Get(ctx workflow.Context) (*chime.GetMeetingOutput, error) {
    var output chime.GetMeetingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetPhoneNumberResult struct {
	Result workflow.Future
}

func (r *ChimeGetPhoneNumberResult) Get(ctx workflow.Context) (*chime.GetPhoneNumberOutput, error) {
    var output chime.GetPhoneNumberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetPhoneNumberOrderResult struct {
	Result workflow.Future
}

func (r *ChimeGetPhoneNumberOrderResult) Get(ctx workflow.Context) (*chime.GetPhoneNumberOrderOutput, error) {
    var output chime.GetPhoneNumberOrderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetPhoneNumberSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeGetPhoneNumberSettingsResult) Get(ctx workflow.Context) (*chime.GetPhoneNumberSettingsOutput, error) {
    var output chime.GetPhoneNumberSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetProxySessionResult struct {
	Result workflow.Future
}

func (r *ChimeGetProxySessionResult) Get(ctx workflow.Context) (*chime.GetProxySessionOutput, error) {
    var output chime.GetProxySessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetRetentionSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeGetRetentionSettingsResult) Get(ctx workflow.Context) (*chime.GetRetentionSettingsOutput, error) {
    var output chime.GetRetentionSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetRoomResult struct {
	Result workflow.Future
}

func (r *ChimeGetRoomResult) Get(ctx workflow.Context) (*chime.GetRoomOutput, error) {
    var output chime.GetRoomOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetUserResult struct {
	Result workflow.Future
}

func (r *ChimeGetUserResult) Get(ctx workflow.Context) (*chime.GetUserOutput, error) {
    var output chime.GetUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetUserSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeGetUserSettingsResult) Get(ctx workflow.Context) (*chime.GetUserSettingsOutput, error) {
    var output chime.GetUserSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorOutput, error) {
    var output chime.GetVoiceConnectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorEmergencyCallingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorEmergencyCallingConfigurationResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    var output chime.GetVoiceConnectorEmergencyCallingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorGroupResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorGroupResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorGroupOutput, error) {
    var output chime.GetVoiceConnectorGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorLoggingConfigurationResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error) {
    var output chime.GetVoiceConnectorLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorOriginationResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorOriginationResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorOriginationOutput, error) {
    var output chime.GetVoiceConnectorOriginationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorProxyResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorProxyResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorProxyOutput, error) {
    var output chime.GetVoiceConnectorProxyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorStreamingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorStreamingConfigurationResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error) {
    var output chime.GetVoiceConnectorStreamingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorTerminationResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorTerminationResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorTerminationOutput, error) {
    var output chime.GetVoiceConnectorTerminationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeGetVoiceConnectorTerminationHealthResult struct {
	Result workflow.Future
}

func (r *ChimeGetVoiceConnectorTerminationHealthResult) Get(ctx workflow.Context) (*chime.GetVoiceConnectorTerminationHealthOutput, error) {
    var output chime.GetVoiceConnectorTerminationHealthOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeInviteUsersResult struct {
	Result workflow.Future
}

func (r *ChimeInviteUsersResult) Get(ctx workflow.Context) (*chime.InviteUsersOutput, error) {
    var output chime.InviteUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListAccountsResult struct {
	Result workflow.Future
}

func (r *ChimeListAccountsResult) Get(ctx workflow.Context) (*chime.ListAccountsOutput, error) {
    var output chime.ListAccountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListAttendeeTagsResult struct {
	Result workflow.Future
}

func (r *ChimeListAttendeeTagsResult) Get(ctx workflow.Context) (*chime.ListAttendeeTagsOutput, error) {
    var output chime.ListAttendeeTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListAttendeesResult struct {
	Result workflow.Future
}

func (r *ChimeListAttendeesResult) Get(ctx workflow.Context) (*chime.ListAttendeesOutput, error) {
    var output chime.ListAttendeesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListBotsResult struct {
	Result workflow.Future
}

func (r *ChimeListBotsResult) Get(ctx workflow.Context) (*chime.ListBotsOutput, error) {
    var output chime.ListBotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListMeetingTagsResult struct {
	Result workflow.Future
}

func (r *ChimeListMeetingTagsResult) Get(ctx workflow.Context) (*chime.ListMeetingTagsOutput, error) {
    var output chime.ListMeetingTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListMeetingsResult struct {
	Result workflow.Future
}

func (r *ChimeListMeetingsResult) Get(ctx workflow.Context) (*chime.ListMeetingsOutput, error) {
    var output chime.ListMeetingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListPhoneNumberOrdersResult struct {
	Result workflow.Future
}

func (r *ChimeListPhoneNumberOrdersResult) Get(ctx workflow.Context) (*chime.ListPhoneNumberOrdersOutput, error) {
    var output chime.ListPhoneNumberOrdersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListPhoneNumbersResult struct {
	Result workflow.Future
}

func (r *ChimeListPhoneNumbersResult) Get(ctx workflow.Context) (*chime.ListPhoneNumbersOutput, error) {
    var output chime.ListPhoneNumbersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListProxySessionsResult struct {
	Result workflow.Future
}

func (r *ChimeListProxySessionsResult) Get(ctx workflow.Context) (*chime.ListProxySessionsOutput, error) {
    var output chime.ListProxySessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListRoomMembershipsResult struct {
	Result workflow.Future
}

func (r *ChimeListRoomMembershipsResult) Get(ctx workflow.Context) (*chime.ListRoomMembershipsOutput, error) {
    var output chime.ListRoomMembershipsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListRoomsResult struct {
	Result workflow.Future
}

func (r *ChimeListRoomsResult) Get(ctx workflow.Context) (*chime.ListRoomsOutput, error) {
    var output chime.ListRoomsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ChimeListTagsForResourceResult) Get(ctx workflow.Context) (*chime.ListTagsForResourceOutput, error) {
    var output chime.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListUsersResult struct {
	Result workflow.Future
}

func (r *ChimeListUsersResult) Get(ctx workflow.Context) (*chime.ListUsersOutput, error) {
    var output chime.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListVoiceConnectorGroupsResult struct {
	Result workflow.Future
}

func (r *ChimeListVoiceConnectorGroupsResult) Get(ctx workflow.Context) (*chime.ListVoiceConnectorGroupsOutput, error) {
    var output chime.ListVoiceConnectorGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListVoiceConnectorTerminationCredentialsResult struct {
	Result workflow.Future
}

func (r *ChimeListVoiceConnectorTerminationCredentialsResult) Get(ctx workflow.Context) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error) {
    var output chime.ListVoiceConnectorTerminationCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeListVoiceConnectorsResult struct {
	Result workflow.Future
}

func (r *ChimeListVoiceConnectorsResult) Get(ctx workflow.Context) (*chime.ListVoiceConnectorsOutput, error) {
    var output chime.ListVoiceConnectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeLogoutUserResult struct {
	Result workflow.Future
}

func (r *ChimeLogoutUserResult) Get(ctx workflow.Context) (*chime.LogoutUserOutput, error) {
    var output chime.LogoutUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutEventsConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimePutEventsConfigurationResult) Get(ctx workflow.Context) (*chime.PutEventsConfigurationOutput, error) {
    var output chime.PutEventsConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutRetentionSettingsResult struct {
	Result workflow.Future
}

func (r *ChimePutRetentionSettingsResult) Get(ctx workflow.Context) (*chime.PutRetentionSettingsOutput, error) {
    var output chime.PutRetentionSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutVoiceConnectorEmergencyCallingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimePutVoiceConnectorEmergencyCallingConfigurationResult) Get(ctx workflow.Context) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    var output chime.PutVoiceConnectorEmergencyCallingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutVoiceConnectorLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimePutVoiceConnectorLoggingConfigurationResult) Get(ctx workflow.Context) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error) {
    var output chime.PutVoiceConnectorLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutVoiceConnectorOriginationResult struct {
	Result workflow.Future
}

func (r *ChimePutVoiceConnectorOriginationResult) Get(ctx workflow.Context) (*chime.PutVoiceConnectorOriginationOutput, error) {
    var output chime.PutVoiceConnectorOriginationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutVoiceConnectorProxyResult struct {
	Result workflow.Future
}

func (r *ChimePutVoiceConnectorProxyResult) Get(ctx workflow.Context) (*chime.PutVoiceConnectorProxyOutput, error) {
    var output chime.PutVoiceConnectorProxyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutVoiceConnectorStreamingConfigurationResult struct {
	Result workflow.Future
}

func (r *ChimePutVoiceConnectorStreamingConfigurationResult) Get(ctx workflow.Context) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error) {
    var output chime.PutVoiceConnectorStreamingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutVoiceConnectorTerminationResult struct {
	Result workflow.Future
}

func (r *ChimePutVoiceConnectorTerminationResult) Get(ctx workflow.Context) (*chime.PutVoiceConnectorTerminationOutput, error) {
    var output chime.PutVoiceConnectorTerminationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimePutVoiceConnectorTerminationCredentialsResult struct {
	Result workflow.Future
}

func (r *ChimePutVoiceConnectorTerminationCredentialsResult) Get(ctx workflow.Context) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error) {
    var output chime.PutVoiceConnectorTerminationCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeRedactConversationMessageResult struct {
	Result workflow.Future
}

func (r *ChimeRedactConversationMessageResult) Get(ctx workflow.Context) (*chime.RedactConversationMessageOutput, error) {
    var output chime.RedactConversationMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeRedactRoomMessageResult struct {
	Result workflow.Future
}

func (r *ChimeRedactRoomMessageResult) Get(ctx workflow.Context) (*chime.RedactRoomMessageOutput, error) {
    var output chime.RedactRoomMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeRegenerateSecurityTokenResult struct {
	Result workflow.Future
}

func (r *ChimeRegenerateSecurityTokenResult) Get(ctx workflow.Context) (*chime.RegenerateSecurityTokenOutput, error) {
    var output chime.RegenerateSecurityTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeResetPersonalPINResult struct {
	Result workflow.Future
}

func (r *ChimeResetPersonalPINResult) Get(ctx workflow.Context) (*chime.ResetPersonalPINOutput, error) {
    var output chime.ResetPersonalPINOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeRestorePhoneNumberResult struct {
	Result workflow.Future
}

func (r *ChimeRestorePhoneNumberResult) Get(ctx workflow.Context) (*chime.RestorePhoneNumberOutput, error) {
    var output chime.RestorePhoneNumberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeSearchAvailablePhoneNumbersResult struct {
	Result workflow.Future
}

func (r *ChimeSearchAvailablePhoneNumbersResult) Get(ctx workflow.Context) (*chime.SearchAvailablePhoneNumbersOutput, error) {
    var output chime.SearchAvailablePhoneNumbersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeTagAttendeeResult struct {
	Result workflow.Future
}

func (r *ChimeTagAttendeeResult) Get(ctx workflow.Context) (*chime.TagAttendeeOutput, error) {
    var output chime.TagAttendeeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeTagMeetingResult struct {
	Result workflow.Future
}

func (r *ChimeTagMeetingResult) Get(ctx workflow.Context) (*chime.TagMeetingOutput, error) {
    var output chime.TagMeetingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeTagResourceResult struct {
	Result workflow.Future
}

func (r *ChimeTagResourceResult) Get(ctx workflow.Context) (*chime.TagResourceOutput, error) {
    var output chime.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUntagAttendeeResult struct {
	Result workflow.Future
}

func (r *ChimeUntagAttendeeResult) Get(ctx workflow.Context) (*chime.UntagAttendeeOutput, error) {
    var output chime.UntagAttendeeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUntagMeetingResult struct {
	Result workflow.Future
}

func (r *ChimeUntagMeetingResult) Get(ctx workflow.Context) (*chime.UntagMeetingOutput, error) {
    var output chime.UntagMeetingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUntagResourceResult struct {
	Result workflow.Future
}

func (r *ChimeUntagResourceResult) Get(ctx workflow.Context) (*chime.UntagResourceOutput, error) {
    var output chime.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateAccountResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateAccountResult) Get(ctx workflow.Context) (*chime.UpdateAccountOutput, error) {
    var output chime.UpdateAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateAccountSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateAccountSettingsResult) Get(ctx workflow.Context) (*chime.UpdateAccountSettingsOutput, error) {
    var output chime.UpdateAccountSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateBotResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateBotResult) Get(ctx workflow.Context) (*chime.UpdateBotOutput, error) {
    var output chime.UpdateBotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateGlobalSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateGlobalSettingsResult) Get(ctx workflow.Context) (*chime.UpdateGlobalSettingsOutput, error) {
    var output chime.UpdateGlobalSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdatePhoneNumberResult struct {
	Result workflow.Future
}

func (r *ChimeUpdatePhoneNumberResult) Get(ctx workflow.Context) (*chime.UpdatePhoneNumberOutput, error) {
    var output chime.UpdatePhoneNumberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdatePhoneNumberSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeUpdatePhoneNumberSettingsResult) Get(ctx workflow.Context) (*chime.UpdatePhoneNumberSettingsOutput, error) {
    var output chime.UpdatePhoneNumberSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateProxySessionResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateProxySessionResult) Get(ctx workflow.Context) (*chime.UpdateProxySessionOutput, error) {
    var output chime.UpdateProxySessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateRoomResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateRoomResult) Get(ctx workflow.Context) (*chime.UpdateRoomOutput, error) {
    var output chime.UpdateRoomOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateRoomMembershipResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateRoomMembershipResult) Get(ctx workflow.Context) (*chime.UpdateRoomMembershipOutput, error) {
    var output chime.UpdateRoomMembershipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateUserResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateUserResult) Get(ctx workflow.Context) (*chime.UpdateUserOutput, error) {
    var output chime.UpdateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateUserSettingsResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateUserSettingsResult) Get(ctx workflow.Context) (*chime.UpdateUserSettingsOutput, error) {
    var output chime.UpdateUserSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateVoiceConnectorResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateVoiceConnectorResult) Get(ctx workflow.Context) (*chime.UpdateVoiceConnectorOutput, error) {
    var output chime.UpdateVoiceConnectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ChimeUpdateVoiceConnectorGroupResult struct {
	Result workflow.Future
}

func (r *ChimeUpdateVoiceConnectorGroupResult) Get(ctx workflow.Context) (*chime.UpdateVoiceConnectorGroupOutput, error) {
    var output chime.UpdateVoiceConnectorGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ChimeStub struct {
    activities ChimeClient
}

func NewChimeStub() ChimeClient {
    return &ChimeStub{}
}

func (a *ChimeStub) AssociatePhoneNumberWithUser(ctx workflow.Context, input *chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error) {
    var output chime.AssociatePhoneNumberWithUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociatePhoneNumberWithUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) AssociatePhoneNumbersWithVoiceConnector(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error) {
    var output chime.AssociatePhoneNumbersWithVoiceConnectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociatePhoneNumbersWithVoiceConnector, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) AssociatePhoneNumbersWithVoiceConnectorGroup(ctx workflow.Context, input *chime.AssociatePhoneNumbersWithVoiceConnectorGroupInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) {
    var output chime.AssociatePhoneNumbersWithVoiceConnectorGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociatePhoneNumbersWithVoiceConnectorGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) AssociateSigninDelegateGroupsWithAccount(ctx workflow.Context, input *chime.AssociateSigninDelegateGroupsWithAccountInput) (*chime.AssociateSigninDelegateGroupsWithAccountOutput, error) {
    var output chime.AssociateSigninDelegateGroupsWithAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateSigninDelegateGroupsWithAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) BatchCreateAttendee(ctx workflow.Context, input *chime.BatchCreateAttendeeInput) (*chime.BatchCreateAttendeeOutput, error) {
    var output chime.BatchCreateAttendeeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchCreateAttendee, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) BatchCreateRoomMembership(ctx workflow.Context, input *chime.BatchCreateRoomMembershipInput) (*chime.BatchCreateRoomMembershipOutput, error) {
    var output chime.BatchCreateRoomMembershipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchCreateRoomMembership, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) BatchDeletePhoneNumber(ctx workflow.Context, input *chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error) {
    var output chime.BatchDeletePhoneNumberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeletePhoneNumber, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) BatchSuspendUser(ctx workflow.Context, input *chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error) {
    var output chime.BatchSuspendUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchSuspendUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) BatchUnsuspendUser(ctx workflow.Context, input *chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error) {
    var output chime.BatchUnsuspendUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchUnsuspendUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) BatchUpdatePhoneNumber(ctx workflow.Context, input *chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error) {
    var output chime.BatchUpdatePhoneNumberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchUpdatePhoneNumber, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) BatchUpdateUser(ctx workflow.Context, input *chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error) {
    var output chime.BatchUpdateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchUpdateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateAccount(ctx workflow.Context, input *chime.CreateAccountInput) (*chime.CreateAccountOutput, error) {
    var output chime.CreateAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateAttendee(ctx workflow.Context, input *chime.CreateAttendeeInput) (*chime.CreateAttendeeOutput, error) {
    var output chime.CreateAttendeeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAttendee, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateBot(ctx workflow.Context, input *chime.CreateBotInput) (*chime.CreateBotOutput, error) {
    var output chime.CreateBotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBot, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateMeeting(ctx workflow.Context, input *chime.CreateMeetingInput) (*chime.CreateMeetingOutput, error) {
    var output chime.CreateMeetingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMeeting, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateMeetingWithAttendees(ctx workflow.Context, input *chime.CreateMeetingWithAttendeesInput) (*chime.CreateMeetingWithAttendeesOutput, error) {
    var output chime.CreateMeetingWithAttendeesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMeetingWithAttendees, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreatePhoneNumberOrder(ctx workflow.Context, input *chime.CreatePhoneNumberOrderInput) (*chime.CreatePhoneNumberOrderOutput, error) {
    var output chime.CreatePhoneNumberOrderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePhoneNumberOrder, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateProxySession(ctx workflow.Context, input *chime.CreateProxySessionInput) (*chime.CreateProxySessionOutput, error) {
    var output chime.CreateProxySessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProxySession, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateRoom(ctx workflow.Context, input *chime.CreateRoomInput) (*chime.CreateRoomOutput, error) {
    var output chime.CreateRoomOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRoom, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateRoomMembership(ctx workflow.Context, input *chime.CreateRoomMembershipInput) (*chime.CreateRoomMembershipOutput, error) {
    var output chime.CreateRoomMembershipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRoomMembership, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateUser(ctx workflow.Context, input *chime.CreateUserInput) (*chime.CreateUserOutput, error) {
    var output chime.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateVoiceConnector(ctx workflow.Context, input *chime.CreateVoiceConnectorInput) (*chime.CreateVoiceConnectorOutput, error) {
    var output chime.CreateVoiceConnectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVoiceConnector, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) CreateVoiceConnectorGroup(ctx workflow.Context, input *chime.CreateVoiceConnectorGroupInput) (*chime.CreateVoiceConnectorGroupOutput, error) {
    var output chime.CreateVoiceConnectorGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVoiceConnectorGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteAccount(ctx workflow.Context, input *chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error) {
    var output chime.DeleteAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteAttendee(ctx workflow.Context, input *chime.DeleteAttendeeInput) (*chime.DeleteAttendeeOutput, error) {
    var output chime.DeleteAttendeeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAttendee, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteEventsConfiguration(ctx workflow.Context, input *chime.DeleteEventsConfigurationInput) (*chime.DeleteEventsConfigurationOutput, error) {
    var output chime.DeleteEventsConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventsConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteMeeting(ctx workflow.Context, input *chime.DeleteMeetingInput) (*chime.DeleteMeetingOutput, error) {
    var output chime.DeleteMeetingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMeeting, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeletePhoneNumber(ctx workflow.Context, input *chime.DeletePhoneNumberInput) (*chime.DeletePhoneNumberOutput, error) {
    var output chime.DeletePhoneNumberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePhoneNumber, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteProxySession(ctx workflow.Context, input *chime.DeleteProxySessionInput) (*chime.DeleteProxySessionOutput, error) {
    var output chime.DeleteProxySessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProxySession, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteRoom(ctx workflow.Context, input *chime.DeleteRoomInput) (*chime.DeleteRoomOutput, error) {
    var output chime.DeleteRoomOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRoom, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteRoomMembership(ctx workflow.Context, input *chime.DeleteRoomMembershipInput) (*chime.DeleteRoomMembershipOutput, error) {
    var output chime.DeleteRoomMembershipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRoomMembership, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnector(ctx workflow.Context, input *chime.DeleteVoiceConnectorInput) (*chime.DeleteVoiceConnectorOutput, error) {
    var output chime.DeleteVoiceConnectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnector, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorEmergencyCallingConfigurationInput) (*chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    var output chime.DeleteVoiceConnectorEmergencyCallingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnectorEmergencyCallingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnectorGroup(ctx workflow.Context, input *chime.DeleteVoiceConnectorGroupInput) (*chime.DeleteVoiceConnectorGroupOutput, error) {
    var output chime.DeleteVoiceConnectorGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnectorGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnectorOrigination(ctx workflow.Context, input *chime.DeleteVoiceConnectorOriginationInput) (*chime.DeleteVoiceConnectorOriginationOutput, error) {
    var output chime.DeleteVoiceConnectorOriginationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnectorOrigination, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnectorProxy(ctx workflow.Context, input *chime.DeleteVoiceConnectorProxyInput) (*chime.DeleteVoiceConnectorProxyOutput, error) {
    var output chime.DeleteVoiceConnectorProxyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnectorProxy, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.DeleteVoiceConnectorStreamingConfigurationInput) (*chime.DeleteVoiceConnectorStreamingConfigurationOutput, error) {
    var output chime.DeleteVoiceConnectorStreamingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnectorStreamingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnectorTermination(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationInput) (*chime.DeleteVoiceConnectorTerminationOutput, error) {
    var output chime.DeleteVoiceConnectorTerminationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnectorTermination, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DeleteVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.DeleteVoiceConnectorTerminationCredentialsInput) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error) {
    var output chime.DeleteVoiceConnectorTerminationCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVoiceConnectorTerminationCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DisassociatePhoneNumberFromUser(ctx workflow.Context, input *chime.DisassociatePhoneNumberFromUserInput) (*chime.DisassociatePhoneNumberFromUserOutput, error) {
    var output chime.DisassociatePhoneNumberFromUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociatePhoneNumberFromUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DisassociatePhoneNumbersFromVoiceConnector(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error) {
    var output chime.DisassociatePhoneNumbersFromVoiceConnectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociatePhoneNumbersFromVoiceConnector, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx workflow.Context, input *chime.DisassociatePhoneNumbersFromVoiceConnectorGroupInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error) {
    var output chime.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociatePhoneNumbersFromVoiceConnectorGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) DisassociateSigninDelegateGroupsFromAccount(ctx workflow.Context, input *chime.DisassociateSigninDelegateGroupsFromAccountInput) (*chime.DisassociateSigninDelegateGroupsFromAccountOutput, error) {
    var output chime.DisassociateSigninDelegateGroupsFromAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateSigninDelegateGroupsFromAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetAccount(ctx workflow.Context, input *chime.GetAccountInput) (*chime.GetAccountOutput, error) {
    var output chime.GetAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetAccountSettings(ctx workflow.Context, input *chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error) {
    var output chime.GetAccountSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetAttendee(ctx workflow.Context, input *chime.GetAttendeeInput) (*chime.GetAttendeeOutput, error) {
    var output chime.GetAttendeeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAttendee, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetBot(ctx workflow.Context, input *chime.GetBotInput) (*chime.GetBotOutput, error) {
    var output chime.GetBotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBot, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetEventsConfiguration(ctx workflow.Context, input *chime.GetEventsConfigurationInput) (*chime.GetEventsConfigurationOutput, error) {
    var output chime.GetEventsConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEventsConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetGlobalSettings(ctx workflow.Context, input *chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error) {
    var output chime.GetGlobalSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGlobalSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetMeeting(ctx workflow.Context, input *chime.GetMeetingInput) (*chime.GetMeetingOutput, error) {
    var output chime.GetMeetingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMeeting, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetPhoneNumber(ctx workflow.Context, input *chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error) {
    var output chime.GetPhoneNumberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPhoneNumber, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetPhoneNumberOrder(ctx workflow.Context, input *chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error) {
    var output chime.GetPhoneNumberOrderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPhoneNumberOrder, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetPhoneNumberSettings(ctx workflow.Context, input *chime.GetPhoneNumberSettingsInput) (*chime.GetPhoneNumberSettingsOutput, error) {
    var output chime.GetPhoneNumberSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPhoneNumberSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetProxySession(ctx workflow.Context, input *chime.GetProxySessionInput) (*chime.GetProxySessionOutput, error) {
    var output chime.GetProxySessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProxySession, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetRetentionSettings(ctx workflow.Context, input *chime.GetRetentionSettingsInput) (*chime.GetRetentionSettingsOutput, error) {
    var output chime.GetRetentionSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRetentionSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetRoom(ctx workflow.Context, input *chime.GetRoomInput) (*chime.GetRoomOutput, error) {
    var output chime.GetRoomOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRoom, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetUser(ctx workflow.Context, input *chime.GetUserInput) (*chime.GetUserOutput, error) {
    var output chime.GetUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetUserSettings(ctx workflow.Context, input *chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error) {
    var output chime.GetUserSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUserSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnector(ctx workflow.Context, input *chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error) {
    var output chime.GetVoiceConnectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnector, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chime.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    var output chime.GetVoiceConnectorEmergencyCallingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorEmergencyCallingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorGroup(ctx workflow.Context, input *chime.GetVoiceConnectorGroupInput) (*chime.GetVoiceConnectorGroupOutput, error) {
    var output chime.GetVoiceConnectorGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorLoggingConfigurationInput) (*chime.GetVoiceConnectorLoggingConfigurationOutput, error) {
    var output chime.GetVoiceConnectorLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorOrigination(ctx workflow.Context, input *chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error) {
    var output chime.GetVoiceConnectorOriginationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorOrigination, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorProxy(ctx workflow.Context, input *chime.GetVoiceConnectorProxyInput) (*chime.GetVoiceConnectorProxyOutput, error) {
    var output chime.GetVoiceConnectorProxyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorProxy, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.GetVoiceConnectorStreamingConfigurationInput) (*chime.GetVoiceConnectorStreamingConfigurationOutput, error) {
    var output chime.GetVoiceConnectorStreamingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorStreamingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorTermination(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error) {
    var output chime.GetVoiceConnectorTerminationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorTermination, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) GetVoiceConnectorTerminationHealth(ctx workflow.Context, input *chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error) {
    var output chime.GetVoiceConnectorTerminationHealthOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVoiceConnectorTerminationHealth, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) InviteUsers(ctx workflow.Context, input *chime.InviteUsersInput) (*chime.InviteUsersOutput, error) {
    var output chime.InviteUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InviteUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListAccounts(ctx workflow.Context, input *chime.ListAccountsInput) (*chime.ListAccountsOutput, error) {
    var output chime.ListAccountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAccounts, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListAttendeeTags(ctx workflow.Context, input *chime.ListAttendeeTagsInput) (*chime.ListAttendeeTagsOutput, error) {
    var output chime.ListAttendeeTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttendeeTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListAttendees(ctx workflow.Context, input *chime.ListAttendeesInput) (*chime.ListAttendeesOutput, error) {
    var output chime.ListAttendeesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttendees, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListBots(ctx workflow.Context, input *chime.ListBotsInput) (*chime.ListBotsOutput, error) {
    var output chime.ListBotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBots, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListMeetingTags(ctx workflow.Context, input *chime.ListMeetingTagsInput) (*chime.ListMeetingTagsOutput, error) {
    var output chime.ListMeetingTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMeetingTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListMeetings(ctx workflow.Context, input *chime.ListMeetingsInput) (*chime.ListMeetingsOutput, error) {
    var output chime.ListMeetingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMeetings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListPhoneNumberOrders(ctx workflow.Context, input *chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error) {
    var output chime.ListPhoneNumberOrdersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPhoneNumberOrders, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListPhoneNumbers(ctx workflow.Context, input *chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error) {
    var output chime.ListPhoneNumbersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPhoneNumbers, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListProxySessions(ctx workflow.Context, input *chime.ListProxySessionsInput) (*chime.ListProxySessionsOutput, error) {
    var output chime.ListProxySessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProxySessions, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListRoomMemberships(ctx workflow.Context, input *chime.ListRoomMembershipsInput) (*chime.ListRoomMembershipsOutput, error) {
    var output chime.ListRoomMembershipsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRoomMemberships, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListRooms(ctx workflow.Context, input *chime.ListRoomsInput) (*chime.ListRoomsOutput, error) {
    var output chime.ListRoomsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRooms, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListTagsForResource(ctx workflow.Context, input *chime.ListTagsForResourceInput) (*chime.ListTagsForResourceOutput, error) {
    var output chime.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListUsers(ctx workflow.Context, input *chime.ListUsersInput) (*chime.ListUsersOutput, error) {
    var output chime.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListVoiceConnectorGroups(ctx workflow.Context, input *chime.ListVoiceConnectorGroupsInput) (*chime.ListVoiceConnectorGroupsOutput, error) {
    var output chime.ListVoiceConnectorGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVoiceConnectorGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error) {
    var output chime.ListVoiceConnectorTerminationCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVoiceConnectorTerminationCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ListVoiceConnectors(ctx workflow.Context, input *chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error) {
    var output chime.ListVoiceConnectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVoiceConnectors, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) LogoutUser(ctx workflow.Context, input *chime.LogoutUserInput) (*chime.LogoutUserOutput, error) {
    var output chime.LogoutUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.LogoutUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutEventsConfiguration(ctx workflow.Context, input *chime.PutEventsConfigurationInput) (*chime.PutEventsConfigurationOutput, error) {
    var output chime.PutEventsConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEventsConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutRetentionSettings(ctx workflow.Context, input *chime.PutRetentionSettingsInput) (*chime.PutRetentionSettingsOutput, error) {
    var output chime.PutRetentionSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRetentionSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutVoiceConnectorEmergencyCallingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorEmergencyCallingConfigurationInput) (*chime.PutVoiceConnectorEmergencyCallingConfigurationOutput, error) {
    var output chime.PutVoiceConnectorEmergencyCallingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutVoiceConnectorEmergencyCallingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutVoiceConnectorLoggingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorLoggingConfigurationInput) (*chime.PutVoiceConnectorLoggingConfigurationOutput, error) {
    var output chime.PutVoiceConnectorLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutVoiceConnectorLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutVoiceConnectorOrigination(ctx workflow.Context, input *chime.PutVoiceConnectorOriginationInput) (*chime.PutVoiceConnectorOriginationOutput, error) {
    var output chime.PutVoiceConnectorOriginationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutVoiceConnectorOrigination, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutVoiceConnectorProxy(ctx workflow.Context, input *chime.PutVoiceConnectorProxyInput) (*chime.PutVoiceConnectorProxyOutput, error) {
    var output chime.PutVoiceConnectorProxyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutVoiceConnectorProxy, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutVoiceConnectorStreamingConfiguration(ctx workflow.Context, input *chime.PutVoiceConnectorStreamingConfigurationInput) (*chime.PutVoiceConnectorStreamingConfigurationOutput, error) {
    var output chime.PutVoiceConnectorStreamingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutVoiceConnectorStreamingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutVoiceConnectorTermination(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationInput) (*chime.PutVoiceConnectorTerminationOutput, error) {
    var output chime.PutVoiceConnectorTerminationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutVoiceConnectorTermination, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) PutVoiceConnectorTerminationCredentials(ctx workflow.Context, input *chime.PutVoiceConnectorTerminationCredentialsInput) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error) {
    var output chime.PutVoiceConnectorTerminationCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutVoiceConnectorTerminationCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) RedactConversationMessage(ctx workflow.Context, input *chime.RedactConversationMessageInput) (*chime.RedactConversationMessageOutput, error) {
    var output chime.RedactConversationMessageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RedactConversationMessage, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) RedactRoomMessage(ctx workflow.Context, input *chime.RedactRoomMessageInput) (*chime.RedactRoomMessageOutput, error) {
    var output chime.RedactRoomMessageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RedactRoomMessage, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) RegenerateSecurityToken(ctx workflow.Context, input *chime.RegenerateSecurityTokenInput) (*chime.RegenerateSecurityTokenOutput, error) {
    var output chime.RegenerateSecurityTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegenerateSecurityToken, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) ResetPersonalPIN(ctx workflow.Context, input *chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error) {
    var output chime.ResetPersonalPINOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetPersonalPIN, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) RestorePhoneNumber(ctx workflow.Context, input *chime.RestorePhoneNumberInput) (*chime.RestorePhoneNumberOutput, error) {
    var output chime.RestorePhoneNumberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestorePhoneNumber, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) SearchAvailablePhoneNumbers(ctx workflow.Context, input *chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error) {
    var output chime.SearchAvailablePhoneNumbersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchAvailablePhoneNumbers, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) TagAttendee(ctx workflow.Context, input *chime.TagAttendeeInput) (*chime.TagAttendeeOutput, error) {
    var output chime.TagAttendeeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagAttendee, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) TagMeeting(ctx workflow.Context, input *chime.TagMeetingInput) (*chime.TagMeetingOutput, error) {
    var output chime.TagMeetingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagMeeting, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) TagResource(ctx workflow.Context, input *chime.TagResourceInput) (*chime.TagResourceOutput, error) {
    var output chime.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UntagAttendee(ctx workflow.Context, input *chime.UntagAttendeeInput) (*chime.UntagAttendeeOutput, error) {
    var output chime.UntagAttendeeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagAttendee, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UntagMeeting(ctx workflow.Context, input *chime.UntagMeetingInput) (*chime.UntagMeetingOutput, error) {
    var output chime.UntagMeetingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagMeeting, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UntagResource(ctx workflow.Context, input *chime.UntagResourceInput) (*chime.UntagResourceOutput, error) {
    var output chime.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateAccount(ctx workflow.Context, input *chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error) {
    var output chime.UpdateAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateAccountSettings(ctx workflow.Context, input *chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error) {
    var output chime.UpdateAccountSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccountSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateBot(ctx workflow.Context, input *chime.UpdateBotInput) (*chime.UpdateBotOutput, error) {
    var output chime.UpdateBotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBot, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateGlobalSettings(ctx workflow.Context, input *chime.UpdateGlobalSettingsInput) (*chime.UpdateGlobalSettingsOutput, error) {
    var output chime.UpdateGlobalSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGlobalSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdatePhoneNumber(ctx workflow.Context, input *chime.UpdatePhoneNumberInput) (*chime.UpdatePhoneNumberOutput, error) {
    var output chime.UpdatePhoneNumberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePhoneNumber, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdatePhoneNumberSettings(ctx workflow.Context, input *chime.UpdatePhoneNumberSettingsInput) (*chime.UpdatePhoneNumberSettingsOutput, error) {
    var output chime.UpdatePhoneNumberSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePhoneNumberSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateProxySession(ctx workflow.Context, input *chime.UpdateProxySessionInput) (*chime.UpdateProxySessionOutput, error) {
    var output chime.UpdateProxySessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProxySession, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateRoom(ctx workflow.Context, input *chime.UpdateRoomInput) (*chime.UpdateRoomOutput, error) {
    var output chime.UpdateRoomOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRoom, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateRoomMembership(ctx workflow.Context, input *chime.UpdateRoomMembershipInput) (*chime.UpdateRoomMembershipOutput, error) {
    var output chime.UpdateRoomMembershipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRoomMembership, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateUser(ctx workflow.Context, input *chime.UpdateUserInput) (*chime.UpdateUserOutput, error) {
    var output chime.UpdateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateUserSettings(ctx workflow.Context, input *chime.UpdateUserSettingsInput) (*chime.UpdateUserSettingsOutput, error) {
    var output chime.UpdateUserSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateVoiceConnector(ctx workflow.Context, input *chime.UpdateVoiceConnectorInput) (*chime.UpdateVoiceConnectorOutput, error) {
    var output chime.UpdateVoiceConnectorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVoiceConnector, input).Get(ctx, &output)
    return &output, err
}

func (a *ChimeStub) UpdateVoiceConnectorGroup(ctx workflow.Context, input *chime.UpdateVoiceConnectorGroupInput) (*chime.UpdateVoiceConnectorGroupOutput, error) {
    var output chime.UpdateVoiceConnectorGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVoiceConnectorGroup, input).Get(ctx, &output)
    return &output, err
}
