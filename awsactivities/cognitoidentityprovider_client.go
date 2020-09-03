package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"go.temporal.io/sdk/workflow"
)

type CognitoIdentityProviderClient interface {
    AddCustomAttributes(ctx workflow.Context, input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
    AddCustomAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.AddCustomAttributesInput) *CognitoidentityproviderAddCustomAttributesResult

    AdminAddUserToGroup(ctx workflow.Context, input *cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
    AdminAddUserToGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminAddUserToGroupInput) *CognitoidentityproviderAdminAddUserToGroupResult

    AdminConfirmSignUp(ctx workflow.Context, input *cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
    AdminConfirmSignUpAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminConfirmSignUpInput) *CognitoidentityproviderAdminConfirmSignUpResult

    AdminCreateUser(ctx workflow.Context, input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error)
    AdminCreateUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminCreateUserInput) *CognitoidentityproviderAdminCreateUserResult

    AdminDeleteUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
    AdminDeleteUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserInput) *CognitoidentityproviderAdminDeleteUserResult

    AdminDeleteUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
    AdminDeleteUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserAttributesInput) *CognitoidentityproviderAdminDeleteUserAttributesResult

    AdminDisableProviderForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableProviderForUserInput) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error)
    AdminDisableProviderForUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableProviderForUserInput) *CognitoidentityproviderAdminDisableProviderForUserResult

    AdminDisableUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error)
    AdminDisableUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableUserInput) *CognitoidentityproviderAdminDisableUserResult

    AdminEnableUser(ctx workflow.Context, input *cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error)
    AdminEnableUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminEnableUserInput) *CognitoidentityproviderAdminEnableUserResult

    AdminForgetDevice(ctx workflow.Context, input *cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
    AdminForgetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminForgetDeviceInput) *CognitoidentityproviderAdminForgetDeviceResult

    AdminGetDevice(ctx workflow.Context, input *cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
    AdminGetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminGetDeviceInput) *CognitoidentityproviderAdminGetDeviceResult

    AdminGetUser(ctx workflow.Context, input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error)
    AdminGetUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminGetUserInput) *CognitoidentityproviderAdminGetUserResult

    AdminInitiateAuth(ctx workflow.Context, input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
    AdminInitiateAuthAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminInitiateAuthInput) *CognitoidentityproviderAdminInitiateAuthResult

    AdminLinkProviderForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminLinkProviderForUserInput) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error)
    AdminLinkProviderForUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminLinkProviderForUserInput) *CognitoidentityproviderAdminLinkProviderForUserResult

    AdminListDevices(ctx workflow.Context, input *cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error)
    AdminListDevicesAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminListDevicesInput) *CognitoidentityproviderAdminListDevicesResult

    AdminListGroupsForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
    AdminListGroupsForUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminListGroupsForUserInput) *CognitoidentityproviderAdminListGroupsForUserResult

    AdminListUserAuthEvents(ctx workflow.Context, input *cognitoidentityprovider.AdminListUserAuthEventsInput) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error)
    AdminListUserAuthEventsAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminListUserAuthEventsInput) *CognitoidentityproviderAdminListUserAuthEventsResult

    AdminRemoveUserFromGroup(ctx workflow.Context, input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
    AdminRemoveUserFromGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) *CognitoidentityproviderAdminRemoveUserFromGroupResult

    AdminResetUserPassword(ctx workflow.Context, input *cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
    AdminResetUserPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminResetUserPasswordInput) *CognitoidentityproviderAdminResetUserPasswordResult

    AdminRespondToAuthChallenge(ctx workflow.Context, input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
    AdminRespondToAuthChallengeAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) *CognitoidentityproviderAdminRespondToAuthChallengeResult

    AdminSetUserMFAPreference(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error)
    AdminSetUserMFAPreferenceAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) *CognitoidentityproviderAdminSetUserMFAPreferenceResult

    AdminSetUserPassword(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserPasswordInput) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error)
    AdminSetUserPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserPasswordInput) *CognitoidentityproviderAdminSetUserPasswordResult

    AdminSetUserSettings(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
    AdminSetUserSettingsAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserSettingsInput) *CognitoidentityproviderAdminSetUserSettingsResult

    AdminUpdateAuthEventFeedback(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error)
    AdminUpdateAuthEventFeedbackAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) *CognitoidentityproviderAdminUpdateAuthEventFeedbackResult

    AdminUpdateDeviceStatus(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
    AdminUpdateDeviceStatusAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) *CognitoidentityproviderAdminUpdateDeviceStatusResult

    AdminUpdateUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
    AdminUpdateUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateUserAttributesInput) *CognitoidentityproviderAdminUpdateUserAttributesResult

    AdminUserGlobalSignOut(ctx workflow.Context, input *cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
    AdminUserGlobalSignOutAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUserGlobalSignOutInput) *CognitoidentityproviderAdminUserGlobalSignOutResult

    AssociateSoftwareToken(ctx workflow.Context, input *cognitoidentityprovider.AssociateSoftwareTokenInput) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error)
    AssociateSoftwareTokenAsync(ctx workflow.Context, input *cognitoidentityprovider.AssociateSoftwareTokenInput) *CognitoidentityproviderAssociateSoftwareTokenResult

    ChangePassword(ctx workflow.Context, input *cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error)
    ChangePasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.ChangePasswordInput) *CognitoidentityproviderChangePasswordResult

    ConfirmDevice(ctx workflow.Context, input *cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
    ConfirmDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.ConfirmDeviceInput) *CognitoidentityproviderConfirmDeviceResult

    ConfirmForgotPassword(ctx workflow.Context, input *cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
    ConfirmForgotPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.ConfirmForgotPasswordInput) *CognitoidentityproviderConfirmForgotPasswordResult

    ConfirmSignUp(ctx workflow.Context, input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
    ConfirmSignUpAsync(ctx workflow.Context, input *cognitoidentityprovider.ConfirmSignUpInput) *CognitoidentityproviderConfirmSignUpResult

    CreateGroup(ctx workflow.Context, input *cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error)
    CreateGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateGroupInput) *CognitoidentityproviderCreateGroupResult

    CreateIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.CreateIdentityProviderInput) (*cognitoidentityprovider.CreateIdentityProviderOutput, error)
    CreateIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateIdentityProviderInput) *CognitoidentityproviderCreateIdentityProviderResult

    CreateResourceServer(ctx workflow.Context, input *cognitoidentityprovider.CreateResourceServerInput) (*cognitoidentityprovider.CreateResourceServerOutput, error)
    CreateResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateResourceServerInput) *CognitoidentityproviderCreateResourceServerResult

    CreateUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
    CreateUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserImportJobInput) *CognitoidentityproviderCreateUserImportJobResult

    CreateUserPool(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error)
    CreateUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolInput) *CognitoidentityproviderCreateUserPoolResult

    CreateUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
    CreateUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolClientInput) *CognitoidentityproviderCreateUserPoolClientResult

    CreateUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolDomainInput) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error)
    CreateUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolDomainInput) *CognitoidentityproviderCreateUserPoolDomainResult

    DeleteGroup(ctx workflow.Context, input *cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error)
    DeleteGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteGroupInput) *CognitoidentityproviderDeleteGroupResult

    DeleteIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.DeleteIdentityProviderInput) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error)
    DeleteIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteIdentityProviderInput) *CognitoidentityproviderDeleteIdentityProviderResult

    DeleteResourceServer(ctx workflow.Context, input *cognitoidentityprovider.DeleteResourceServerInput) (*cognitoidentityprovider.DeleteResourceServerOutput, error)
    DeleteResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteResourceServerInput) *CognitoidentityproviderDeleteResourceServerResult

    DeleteUser(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error)
    DeleteUserAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserInput) *CognitoidentityproviderDeleteUserResult

    DeleteUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
    DeleteUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserAttributesInput) *CognitoidentityproviderDeleteUserAttributesResult

    DeleteUserPool(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
    DeleteUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolInput) *CognitoidentityproviderDeleteUserPoolResult

    DeleteUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
    DeleteUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolClientInput) *CognitoidentityproviderDeleteUserPoolClientResult

    DeleteUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolDomainInput) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error)
    DeleteUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolDomainInput) *CognitoidentityproviderDeleteUserPoolDomainResult

    DescribeIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
    DescribeIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput) *CognitoidentityproviderDescribeIdentityProviderResult

    DescribeResourceServer(ctx workflow.Context, input *cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
    DescribeResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeResourceServerInput) *CognitoidentityproviderDescribeResourceServerResult

    DescribeRiskConfiguration(ctx workflow.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error)
    DescribeRiskConfigurationAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput) *CognitoidentityproviderDescribeRiskConfigurationResult

    DescribeUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
    DescribeUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserImportJobInput) *CognitoidentityproviderDescribeUserImportJobResult

    DescribeUserPool(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
    DescribeUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolInput) *CognitoidentityproviderDescribeUserPoolResult

    DescribeUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
    DescribeUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput) *CognitoidentityproviderDescribeUserPoolClientResult

    DescribeUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
    DescribeUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput) *CognitoidentityproviderDescribeUserPoolDomainResult

    ForgetDevice(ctx workflow.Context, input *cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error)
    ForgetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.ForgetDeviceInput) *CognitoidentityproviderForgetDeviceResult

    ForgotPassword(ctx workflow.Context, input *cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error)
    ForgotPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.ForgotPasswordInput) *CognitoidentityproviderForgotPasswordResult

    GetCSVHeader(ctx workflow.Context, input *cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
    GetCSVHeaderAsync(ctx workflow.Context, input *cognitoidentityprovider.GetCSVHeaderInput) *CognitoidentityproviderGetCSVHeaderResult

    GetDevice(ctx workflow.Context, input *cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error)
    GetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.GetDeviceInput) *CognitoidentityproviderGetDeviceResult

    GetGroup(ctx workflow.Context, input *cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error)
    GetGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.GetGroupInput) *CognitoidentityproviderGetGroupResult

    GetIdentityProviderByIdentifier(ctx workflow.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
    GetIdentityProviderByIdentifierAsync(ctx workflow.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) *CognitoidentityproviderGetIdentityProviderByIdentifierResult

    GetSigningCertificate(ctx workflow.Context, input *cognitoidentityprovider.GetSigningCertificateInput) (*cognitoidentityprovider.GetSigningCertificateOutput, error)
    GetSigningCertificateAsync(ctx workflow.Context, input *cognitoidentityprovider.GetSigningCertificateInput) *CognitoidentityproviderGetSigningCertificateResult

    GetUICustomization(ctx workflow.Context, input *cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error)
    GetUICustomizationAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUICustomizationInput) *CognitoidentityproviderGetUICustomizationResult

    GetUser(ctx workflow.Context, input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error)
    GetUserAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUserInput) *CognitoidentityproviderGetUserResult

    GetUserAttributeVerificationCode(ctx workflow.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
    GetUserAttributeVerificationCodeAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) *CognitoidentityproviderGetUserAttributeVerificationCodeResult

    GetUserPoolMfaConfig(ctx workflow.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error)
    GetUserPoolMfaConfigAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput) *CognitoidentityproviderGetUserPoolMfaConfigResult

    GlobalSignOut(ctx workflow.Context, input *cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error)
    GlobalSignOutAsync(ctx workflow.Context, input *cognitoidentityprovider.GlobalSignOutInput) *CognitoidentityproviderGlobalSignOutResult

    InitiateAuth(ctx workflow.Context, input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error)
    InitiateAuthAsync(ctx workflow.Context, input *cognitoidentityprovider.InitiateAuthInput) *CognitoidentityproviderInitiateAuthResult

    ListDevices(ctx workflow.Context, input *cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error)
    ListDevicesAsync(ctx workflow.Context, input *cognitoidentityprovider.ListDevicesInput) *CognitoidentityproviderListDevicesResult

    ListGroups(ctx workflow.Context, input *cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error)
    ListGroupsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListGroupsInput) *CognitoidentityproviderListGroupsResult

    ListIdentityProviders(ctx workflow.Context, input *cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
    ListIdentityProvidersAsync(ctx workflow.Context, input *cognitoidentityprovider.ListIdentityProvidersInput) *CognitoidentityproviderListIdentityProvidersResult

    ListResourceServers(ctx workflow.Context, input *cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error)
    ListResourceServersAsync(ctx workflow.Context, input *cognitoidentityprovider.ListResourceServersInput) *CognitoidentityproviderListResourceServersResult

    ListTagsForResource(ctx workflow.Context, input *cognitoidentityprovider.ListTagsForResourceInput) (*cognitoidentityprovider.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *cognitoidentityprovider.ListTagsForResourceInput) *CognitoidentityproviderListTagsForResourceResult

    ListUserImportJobs(ctx workflow.Context, input *cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
    ListUserImportJobsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUserImportJobsInput) *CognitoidentityproviderListUserImportJobsResult

    ListUserPoolClients(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
    ListUserPoolClientsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolClientsInput) *CognitoidentityproviderListUserPoolClientsResult

    ListUserPools(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error)
    ListUserPoolsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolsInput) *CognitoidentityproviderListUserPoolsResult

    ListUsers(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error)
    ListUsersAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInput) *CognitoidentityproviderListUsersResult

    ListUsersInGroup(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
    ListUsersInGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInGroupInput) *CognitoidentityproviderListUsersInGroupResult

    ResendConfirmationCode(ctx workflow.Context, input *cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
    ResendConfirmationCodeAsync(ctx workflow.Context, input *cognitoidentityprovider.ResendConfirmationCodeInput) *CognitoidentityproviderResendConfirmationCodeResult

    RespondToAuthChallenge(ctx workflow.Context, input *cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
    RespondToAuthChallengeAsync(ctx workflow.Context, input *cognitoidentityprovider.RespondToAuthChallengeInput) *CognitoidentityproviderRespondToAuthChallengeResult

    SetRiskConfiguration(ctx workflow.Context, input *cognitoidentityprovider.SetRiskConfigurationInput) (*cognitoidentityprovider.SetRiskConfigurationOutput, error)
    SetRiskConfigurationAsync(ctx workflow.Context, input *cognitoidentityprovider.SetRiskConfigurationInput) *CognitoidentityproviderSetRiskConfigurationResult

    SetUICustomization(ctx workflow.Context, input *cognitoidentityprovider.SetUICustomizationInput) (*cognitoidentityprovider.SetUICustomizationOutput, error)
    SetUICustomizationAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUICustomizationInput) *CognitoidentityproviderSetUICustomizationResult

    SetUserMFAPreference(ctx workflow.Context, input *cognitoidentityprovider.SetUserMFAPreferenceInput) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error)
    SetUserMFAPreferenceAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUserMFAPreferenceInput) *CognitoidentityproviderSetUserMFAPreferenceResult

    SetUserPoolMfaConfig(ctx workflow.Context, input *cognitoidentityprovider.SetUserPoolMfaConfigInput) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error)
    SetUserPoolMfaConfigAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUserPoolMfaConfigInput) *CognitoidentityproviderSetUserPoolMfaConfigResult

    SetUserSettings(ctx workflow.Context, input *cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error)
    SetUserSettingsAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUserSettingsInput) *CognitoidentityproviderSetUserSettingsResult

    SignUp(ctx workflow.Context, input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error)
    SignUpAsync(ctx workflow.Context, input *cognitoidentityprovider.SignUpInput) *CognitoidentityproviderSignUpResult

    StartUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error)
    StartUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.StartUserImportJobInput) *CognitoidentityproviderStartUserImportJobResult

    StopUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error)
    StopUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.StopUserImportJobInput) *CognitoidentityproviderStopUserImportJobResult

    TagResource(ctx workflow.Context, input *cognitoidentityprovider.TagResourceInput) (*cognitoidentityprovider.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *cognitoidentityprovider.TagResourceInput) *CognitoidentityproviderTagResourceResult

    UntagResource(ctx workflow.Context, input *cognitoidentityprovider.UntagResourceInput) (*cognitoidentityprovider.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *cognitoidentityprovider.UntagResourceInput) *CognitoidentityproviderUntagResourceResult

    UpdateAuthEventFeedback(ctx workflow.Context, input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error)
    UpdateAuthEventFeedbackAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) *CognitoidentityproviderUpdateAuthEventFeedbackResult

    UpdateDeviceStatus(ctx workflow.Context, input *cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
    UpdateDeviceStatusAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateDeviceStatusInput) *CognitoidentityproviderUpdateDeviceStatusResult

    UpdateGroup(ctx workflow.Context, input *cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error)
    UpdateGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateGroupInput) *CognitoidentityproviderUpdateGroupResult

    UpdateIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.UpdateIdentityProviderInput) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error)
    UpdateIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateIdentityProviderInput) *CognitoidentityproviderUpdateIdentityProviderResult

    UpdateResourceServer(ctx workflow.Context, input *cognitoidentityprovider.UpdateResourceServerInput) (*cognitoidentityprovider.UpdateResourceServerOutput, error)
    UpdateResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateResourceServerInput) *CognitoidentityproviderUpdateResourceServerResult

    UpdateUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
    UpdateUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserAttributesInput) *CognitoidentityproviderUpdateUserAttributesResult

    UpdateUserPool(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
    UpdateUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolInput) *CognitoidentityproviderUpdateUserPoolResult

    UpdateUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
    UpdateUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolClientInput) *CognitoidentityproviderUpdateUserPoolClientResult

    UpdateUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolDomainInput) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error)
    UpdateUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolDomainInput) *CognitoidentityproviderUpdateUserPoolDomainResult

    VerifySoftwareToken(ctx workflow.Context, input *cognitoidentityprovider.VerifySoftwareTokenInput) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error)
    VerifySoftwareTokenAsync(ctx workflow.Context, input *cognitoidentityprovider.VerifySoftwareTokenInput) *CognitoidentityproviderVerifySoftwareTokenResult

    VerifyUserAttribute(ctx workflow.Context, input *cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
    VerifyUserAttributeAsync(ctx workflow.Context, input *cognitoidentityprovider.VerifyUserAttributeInput) *CognitoidentityproviderVerifyUserAttributeResult
}
type CognitoidentityproviderAddCustomAttributesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAddCustomAttributesResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
    var output cognitoidentityprovider.AddCustomAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminAddUserToGroupResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminAddUserToGroupResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
    var output cognitoidentityprovider.AdminAddUserToGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminConfirmSignUpResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminConfirmSignUpResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error) {
    var output cognitoidentityprovider.AdminConfirmSignUpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminCreateUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminCreateUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
    var output cognitoidentityprovider.AdminCreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminDeleteUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminDeleteUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
    var output cognitoidentityprovider.AdminDeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminDeleteUserAttributesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminDeleteUserAttributesResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error) {
    var output cognitoidentityprovider.AdminDeleteUserAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminDisableProviderForUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminDisableProviderForUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error) {
    var output cognitoidentityprovider.AdminDisableProviderForUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminDisableUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminDisableUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminDisableUserOutput, error) {
    var output cognitoidentityprovider.AdminDisableUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminEnableUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminEnableUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminEnableUserOutput, error) {
    var output cognitoidentityprovider.AdminEnableUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminForgetDeviceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminForgetDeviceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminForgetDeviceOutput, error) {
    var output cognitoidentityprovider.AdminForgetDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminGetDeviceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminGetDeviceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminGetDeviceOutput, error) {
    var output cognitoidentityprovider.AdminGetDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminGetUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminGetUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminGetUserOutput, error) {
    var output cognitoidentityprovider.AdminGetUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminInitiateAuthResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminInitiateAuthResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
    var output cognitoidentityprovider.AdminInitiateAuthOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminLinkProviderForUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminLinkProviderForUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error) {
    var output cognitoidentityprovider.AdminLinkProviderForUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminListDevicesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminListDevicesResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminListDevicesOutput, error) {
    var output cognitoidentityprovider.AdminListDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminListGroupsForUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminListGroupsForUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
    var output cognitoidentityprovider.AdminListGroupsForUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminListUserAuthEventsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminListUserAuthEventsResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error) {
    var output cognitoidentityprovider.AdminListUserAuthEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminRemoveUserFromGroupResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminRemoveUserFromGroupResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error) {
    var output cognitoidentityprovider.AdminRemoveUserFromGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminResetUserPasswordResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminResetUserPasswordResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error) {
    var output cognitoidentityprovider.AdminResetUserPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminRespondToAuthChallengeResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminRespondToAuthChallengeResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {
    var output cognitoidentityprovider.AdminRespondToAuthChallengeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminSetUserMFAPreferenceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminSetUserMFAPreferenceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error) {
    var output cognitoidentityprovider.AdminSetUserMFAPreferenceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminSetUserPasswordResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminSetUserPasswordResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
    var output cognitoidentityprovider.AdminSetUserPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminSetUserSettingsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminSetUserSettingsResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error) {
    var output cognitoidentityprovider.AdminSetUserSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminUpdateAuthEventFeedbackResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminUpdateAuthEventFeedbackResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error) {
    var output cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminUpdateDeviceStatusResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminUpdateDeviceStatusResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error) {
    var output cognitoidentityprovider.AdminUpdateDeviceStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminUpdateUserAttributesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminUpdateUserAttributesResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
    var output cognitoidentityprovider.AdminUpdateUserAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAdminUserGlobalSignOutResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAdminUserGlobalSignOutResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error) {
    var output cognitoidentityprovider.AdminUserGlobalSignOutOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderAssociateSoftwareTokenResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderAssociateSoftwareTokenResult) Get(ctx workflow.Context) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error) {
    var output cognitoidentityprovider.AssociateSoftwareTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderChangePasswordResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderChangePasswordResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ChangePasswordOutput, error) {
    var output cognitoidentityprovider.ChangePasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderConfirmDeviceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderConfirmDeviceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ConfirmDeviceOutput, error) {
    var output cognitoidentityprovider.ConfirmDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderConfirmForgotPasswordResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderConfirmForgotPasswordResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {
    var output cognitoidentityprovider.ConfirmForgotPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderConfirmSignUpResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderConfirmSignUpResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
    var output cognitoidentityprovider.ConfirmSignUpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderCreateGroupResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderCreateGroupResult) Get(ctx workflow.Context) (*cognitoidentityprovider.CreateGroupOutput, error) {
    var output cognitoidentityprovider.CreateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderCreateIdentityProviderResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderCreateIdentityProviderResult) Get(ctx workflow.Context) (*cognitoidentityprovider.CreateIdentityProviderOutput, error) {
    var output cognitoidentityprovider.CreateIdentityProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderCreateResourceServerResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderCreateResourceServerResult) Get(ctx workflow.Context) (*cognitoidentityprovider.CreateResourceServerOutput, error) {
    var output cognitoidentityprovider.CreateResourceServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderCreateUserImportJobResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderCreateUserImportJobResult) Get(ctx workflow.Context) (*cognitoidentityprovider.CreateUserImportJobOutput, error) {
    var output cognitoidentityprovider.CreateUserImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderCreateUserPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderCreateUserPoolResult) Get(ctx workflow.Context) (*cognitoidentityprovider.CreateUserPoolOutput, error) {
    var output cognitoidentityprovider.CreateUserPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderCreateUserPoolClientResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderCreateUserPoolClientResult) Get(ctx workflow.Context) (*cognitoidentityprovider.CreateUserPoolClientOutput, error) {
    var output cognitoidentityprovider.CreateUserPoolClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderCreateUserPoolDomainResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderCreateUserPoolDomainResult) Get(ctx workflow.Context) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.CreateUserPoolDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteGroupResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteGroupResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteGroupOutput, error) {
    var output cognitoidentityprovider.DeleteGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteIdentityProviderResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteIdentityProviderResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error) {
    var output cognitoidentityprovider.DeleteIdentityProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteResourceServerResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteResourceServerResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteResourceServerOutput, error) {
    var output cognitoidentityprovider.DeleteResourceServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteUserOutput, error) {
    var output cognitoidentityprovider.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteUserAttributesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteUserAttributesResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteUserAttributesOutput, error) {
    var output cognitoidentityprovider.DeleteUserAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteUserPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteUserPoolResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteUserPoolOutput, error) {
    var output cognitoidentityprovider.DeleteUserPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteUserPoolClientResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteUserPoolClientResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error) {
    var output cognitoidentityprovider.DeleteUserPoolClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDeleteUserPoolDomainResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDeleteUserPoolDomainResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.DeleteUserPoolDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDescribeIdentityProviderResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDescribeIdentityProviderResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
    var output cognitoidentityprovider.DescribeIdentityProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDescribeResourceServerResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDescribeResourceServerResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
    var output cognitoidentityprovider.DescribeResourceServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDescribeRiskConfigurationResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDescribeRiskConfigurationResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
    var output cognitoidentityprovider.DescribeRiskConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDescribeUserImportJobResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDescribeUserImportJobResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
    var output cognitoidentityprovider.DescribeUserImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDescribeUserPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDescribeUserPoolResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
    var output cognitoidentityprovider.DescribeUserPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDescribeUserPoolClientResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDescribeUserPoolClientResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
    var output cognitoidentityprovider.DescribeUserPoolClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderDescribeUserPoolDomainResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderDescribeUserPoolDomainResult) Get(ctx workflow.Context) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.DescribeUserPoolDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderForgetDeviceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderForgetDeviceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ForgetDeviceOutput, error) {
    var output cognitoidentityprovider.ForgetDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderForgotPasswordResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderForgotPasswordResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
    var output cognitoidentityprovider.ForgotPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetCSVHeaderResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetCSVHeaderResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
    var output cognitoidentityprovider.GetCSVHeaderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetDeviceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetDeviceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetDeviceOutput, error) {
    var output cognitoidentityprovider.GetDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetGroupResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetGroupResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetGroupOutput, error) {
    var output cognitoidentityprovider.GetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetIdentityProviderByIdentifierResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetIdentityProviderByIdentifierResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
    var output cognitoidentityprovider.GetIdentityProviderByIdentifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetSigningCertificateResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetSigningCertificateResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
    var output cognitoidentityprovider.GetSigningCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetUICustomizationResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetUICustomizationResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
    var output cognitoidentityprovider.GetUICustomizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetUserResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetUserResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetUserOutput, error) {
    var output cognitoidentityprovider.GetUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetUserAttributeVerificationCodeResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetUserAttributeVerificationCodeResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
    var output cognitoidentityprovider.GetUserAttributeVerificationCodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGetUserPoolMfaConfigResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGetUserPoolMfaConfigResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
    var output cognitoidentityprovider.GetUserPoolMfaConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderGlobalSignOutResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderGlobalSignOutResult) Get(ctx workflow.Context) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
    var output cognitoidentityprovider.GlobalSignOutOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderInitiateAuthResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderInitiateAuthResult) Get(ctx workflow.Context) (*cognitoidentityprovider.InitiateAuthOutput, error) {
    var output cognitoidentityprovider.InitiateAuthOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListDevicesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListDevicesResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListDevicesOutput, error) {
    var output cognitoidentityprovider.ListDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListGroupsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListGroupsResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListGroupsOutput, error) {
    var output cognitoidentityprovider.ListGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListIdentityProvidersResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListIdentityProvidersResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
    var output cognitoidentityprovider.ListIdentityProvidersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListResourceServersResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListResourceServersResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListResourceServersOutput, error) {
    var output cognitoidentityprovider.ListResourceServersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListTagsForResourceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
    var output cognitoidentityprovider.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListUserImportJobsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListUserImportJobsResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
    var output cognitoidentityprovider.ListUserImportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListUserPoolClientsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListUserPoolClientsResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
    var output cognitoidentityprovider.ListUserPoolClientsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListUserPoolsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListUserPoolsResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
    var output cognitoidentityprovider.ListUserPoolsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListUsersResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListUsersResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListUsersOutput, error) {
    var output cognitoidentityprovider.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderListUsersInGroupResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderListUsersInGroupResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
    var output cognitoidentityprovider.ListUsersInGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderResendConfirmationCodeResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderResendConfirmationCodeResult) Get(ctx workflow.Context) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error) {
    var output cognitoidentityprovider.ResendConfirmationCodeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderRespondToAuthChallengeResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderRespondToAuthChallengeResult) Get(ctx workflow.Context) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
    var output cognitoidentityprovider.RespondToAuthChallengeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderSetRiskConfigurationResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderSetRiskConfigurationResult) Get(ctx workflow.Context) (*cognitoidentityprovider.SetRiskConfigurationOutput, error) {
    var output cognitoidentityprovider.SetRiskConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderSetUICustomizationResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderSetUICustomizationResult) Get(ctx workflow.Context) (*cognitoidentityprovider.SetUICustomizationOutput, error) {
    var output cognitoidentityprovider.SetUICustomizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderSetUserMFAPreferenceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderSetUserMFAPreferenceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error) {
    var output cognitoidentityprovider.SetUserMFAPreferenceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderSetUserPoolMfaConfigResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderSetUserPoolMfaConfigResult) Get(ctx workflow.Context) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error) {
    var output cognitoidentityprovider.SetUserPoolMfaConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderSetUserSettingsResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderSetUserSettingsResult) Get(ctx workflow.Context) (*cognitoidentityprovider.SetUserSettingsOutput, error) {
    var output cognitoidentityprovider.SetUserSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderSignUpResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderSignUpResult) Get(ctx workflow.Context) (*cognitoidentityprovider.SignUpOutput, error) {
    var output cognitoidentityprovider.SignUpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderStartUserImportJobResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderStartUserImportJobResult) Get(ctx workflow.Context) (*cognitoidentityprovider.StartUserImportJobOutput, error) {
    var output cognitoidentityprovider.StartUserImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderStopUserImportJobResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderStopUserImportJobResult) Get(ctx workflow.Context) (*cognitoidentityprovider.StopUserImportJobOutput, error) {
    var output cognitoidentityprovider.StopUserImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderTagResourceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderTagResourceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.TagResourceOutput, error) {
    var output cognitoidentityprovider.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUntagResourceResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUntagResourceResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UntagResourceOutput, error) {
    var output cognitoidentityprovider.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateAuthEventFeedbackResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateAuthEventFeedbackResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error) {
    var output cognitoidentityprovider.UpdateAuthEventFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateDeviceStatusResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateDeviceStatusResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error) {
    var output cognitoidentityprovider.UpdateDeviceStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateGroupResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateGroupResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateGroupOutput, error) {
    var output cognitoidentityprovider.UpdateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateIdentityProviderResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateIdentityProviderResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error) {
    var output cognitoidentityprovider.UpdateIdentityProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateResourceServerResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateResourceServerResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateResourceServerOutput, error) {
    var output cognitoidentityprovider.UpdateResourceServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateUserAttributesResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateUserAttributesResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateUserAttributesOutput, error) {
    var output cognitoidentityprovider.UpdateUserAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateUserPoolResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateUserPoolResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateUserPoolOutput, error) {
    var output cognitoidentityprovider.UpdateUserPoolOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateUserPoolClientResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateUserPoolClientResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error) {
    var output cognitoidentityprovider.UpdateUserPoolClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderUpdateUserPoolDomainResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderUpdateUserPoolDomainResult) Get(ctx workflow.Context) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.UpdateUserPoolDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderVerifySoftwareTokenResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderVerifySoftwareTokenResult) Get(ctx workflow.Context) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error) {
    var output cognitoidentityprovider.VerifySoftwareTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitoidentityproviderVerifyUserAttributeResult struct {
	Result workflow.Future
}

func (r *CognitoidentityproviderVerifyUserAttributeResult) Get(ctx workflow.Context) (*cognitoidentityprovider.VerifyUserAttributeOutput, error) {
    var output cognitoidentityprovider.VerifyUserAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CognitoIdentityProviderStub struct {
    activities CognitoIdentityProviderClient
}

func NewCognitoIdentityProviderStub() CognitoIdentityProviderClient {
    return &CognitoIdentityProviderStub{}
}

func (a *CognitoIdentityProviderStub) AddCustomAttributes(ctx workflow.Context, input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
    var output cognitoidentityprovider.AddCustomAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddCustomAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminAddUserToGroup(ctx workflow.Context, input *cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
    var output cognitoidentityprovider.AdminAddUserToGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminAddUserToGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminConfirmSignUp(ctx workflow.Context, input *cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error) {
    var output cognitoidentityprovider.AdminConfirmSignUpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminConfirmSignUp, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminCreateUser(ctx workflow.Context, input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
    var output cognitoidentityprovider.AdminCreateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminCreateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminDeleteUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
    var output cognitoidentityprovider.AdminDeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminDeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminDeleteUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error) {
    var output cognitoidentityprovider.AdminDeleteUserAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminDeleteUserAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminDisableProviderForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableProviderForUserInput) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error) {
    var output cognitoidentityprovider.AdminDisableProviderForUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminDisableProviderForUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminDisableUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error) {
    var output cognitoidentityprovider.AdminDisableUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminDisableUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminEnableUser(ctx workflow.Context, input *cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error) {
    var output cognitoidentityprovider.AdminEnableUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminEnableUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminForgetDevice(ctx workflow.Context, input *cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error) {
    var output cognitoidentityprovider.AdminForgetDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminForgetDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminGetDevice(ctx workflow.Context, input *cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error) {
    var output cognitoidentityprovider.AdminGetDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminGetDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminGetUser(ctx workflow.Context, input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error) {
    var output cognitoidentityprovider.AdminGetUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminGetUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminInitiateAuth(ctx workflow.Context, input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
    var output cognitoidentityprovider.AdminInitiateAuthOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminInitiateAuth, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminLinkProviderForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminLinkProviderForUserInput) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error) {
    var output cognitoidentityprovider.AdminLinkProviderForUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminLinkProviderForUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminListDevices(ctx workflow.Context, input *cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error) {
    var output cognitoidentityprovider.AdminListDevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminListDevices, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminListGroupsForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
    var output cognitoidentityprovider.AdminListGroupsForUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminListGroupsForUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminListUserAuthEvents(ctx workflow.Context, input *cognitoidentityprovider.AdminListUserAuthEventsInput) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error) {
    var output cognitoidentityprovider.AdminListUserAuthEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminListUserAuthEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminRemoveUserFromGroup(ctx workflow.Context, input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error) {
    var output cognitoidentityprovider.AdminRemoveUserFromGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminRemoveUserFromGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminResetUserPassword(ctx workflow.Context, input *cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error) {
    var output cognitoidentityprovider.AdminResetUserPasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminResetUserPassword, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminRespondToAuthChallenge(ctx workflow.Context, input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {
    var output cognitoidentityprovider.AdminRespondToAuthChallengeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminRespondToAuthChallenge, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminSetUserMFAPreference(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error) {
    var output cognitoidentityprovider.AdminSetUserMFAPreferenceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminSetUserMFAPreference, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminSetUserPassword(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserPasswordInput) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
    var output cognitoidentityprovider.AdminSetUserPasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminSetUserPassword, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminSetUserSettings(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error) {
    var output cognitoidentityprovider.AdminSetUserSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminSetUserSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminUpdateAuthEventFeedback(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error) {
    var output cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminUpdateAuthEventFeedback, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminUpdateDeviceStatus(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error) {
    var output cognitoidentityprovider.AdminUpdateDeviceStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminUpdateDeviceStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminUpdateUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
    var output cognitoidentityprovider.AdminUpdateUserAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminUpdateUserAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AdminUserGlobalSignOut(ctx workflow.Context, input *cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error) {
    var output cognitoidentityprovider.AdminUserGlobalSignOutOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AdminUserGlobalSignOut, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) AssociateSoftwareToken(ctx workflow.Context, input *cognitoidentityprovider.AssociateSoftwareTokenInput) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error) {
    var output cognitoidentityprovider.AssociateSoftwareTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateSoftwareToken, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ChangePassword(ctx workflow.Context, input *cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error) {
    var output cognitoidentityprovider.ChangePasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ChangePassword, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ConfirmDevice(ctx workflow.Context, input *cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error) {
    var output cognitoidentityprovider.ConfirmDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ConfirmForgotPassword(ctx workflow.Context, input *cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {
    var output cognitoidentityprovider.ConfirmForgotPasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmForgotPassword, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ConfirmSignUp(ctx workflow.Context, input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
    var output cognitoidentityprovider.ConfirmSignUpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmSignUp, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) CreateGroup(ctx workflow.Context, input *cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error) {
    var output cognitoidentityprovider.CreateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) CreateIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.CreateIdentityProviderInput) (*cognitoidentityprovider.CreateIdentityProviderOutput, error) {
    var output cognitoidentityprovider.CreateIdentityProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIdentityProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) CreateResourceServer(ctx workflow.Context, input *cognitoidentityprovider.CreateResourceServerInput) (*cognitoidentityprovider.CreateResourceServerOutput, error) {
    var output cognitoidentityprovider.CreateResourceServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResourceServer, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) CreateUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error) {
    var output cognitoidentityprovider.CreateUserImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) CreateUserPool(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error) {
    var output cognitoidentityprovider.CreateUserPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserPool, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) CreateUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error) {
    var output cognitoidentityprovider.CreateUserPoolClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserPoolClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) CreateUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolDomainInput) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.CreateUserPoolDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserPoolDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteGroup(ctx workflow.Context, input *cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error) {
    var output cognitoidentityprovider.DeleteGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.DeleteIdentityProviderInput) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error) {
    var output cognitoidentityprovider.DeleteIdentityProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIdentityProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteResourceServer(ctx workflow.Context, input *cognitoidentityprovider.DeleteResourceServerInput) (*cognitoidentityprovider.DeleteResourceServerOutput, error) {
    var output cognitoidentityprovider.DeleteResourceServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourceServer, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteUser(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error) {
    var output cognitoidentityprovider.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error) {
    var output cognitoidentityprovider.DeleteUserAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteUserPool(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error) {
    var output cognitoidentityprovider.DeleteUserPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserPool, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error) {
    var output cognitoidentityprovider.DeleteUserPoolClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserPoolClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DeleteUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolDomainInput) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.DeleteUserPoolDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserPoolDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DescribeIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
    var output cognitoidentityprovider.DescribeIdentityProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DescribeResourceServer(ctx workflow.Context, input *cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
    var output cognitoidentityprovider.DescribeResourceServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeResourceServer, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DescribeRiskConfiguration(ctx workflow.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
    var output cognitoidentityprovider.DescribeRiskConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRiskConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DescribeUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
    var output cognitoidentityprovider.DescribeUserImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DescribeUserPool(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
    var output cognitoidentityprovider.DescribeUserPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserPool, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DescribeUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
    var output cognitoidentityprovider.DescribeUserPoolClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserPoolClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) DescribeUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.DescribeUserPoolDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserPoolDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ForgetDevice(ctx workflow.Context, input *cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error) {
    var output cognitoidentityprovider.ForgetDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ForgetDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ForgotPassword(ctx workflow.Context, input *cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
    var output cognitoidentityprovider.ForgotPasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ForgotPassword, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetCSVHeader(ctx workflow.Context, input *cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
    var output cognitoidentityprovider.GetCSVHeaderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCSVHeader, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetDevice(ctx workflow.Context, input *cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error) {
    var output cognitoidentityprovider.GetDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetGroup(ctx workflow.Context, input *cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error) {
    var output cognitoidentityprovider.GetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetIdentityProviderByIdentifier(ctx workflow.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
    var output cognitoidentityprovider.GetIdentityProviderByIdentifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityProviderByIdentifier, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetSigningCertificate(ctx workflow.Context, input *cognitoidentityprovider.GetSigningCertificateInput) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
    var output cognitoidentityprovider.GetSigningCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSigningCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetUICustomization(ctx workflow.Context, input *cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
    var output cognitoidentityprovider.GetUICustomizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUICustomization, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetUser(ctx workflow.Context, input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
    var output cognitoidentityprovider.GetUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUser, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetUserAttributeVerificationCode(ctx workflow.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
    var output cognitoidentityprovider.GetUserAttributeVerificationCodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUserAttributeVerificationCode, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GetUserPoolMfaConfig(ctx workflow.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
    var output cognitoidentityprovider.GetUserPoolMfaConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUserPoolMfaConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) GlobalSignOut(ctx workflow.Context, input *cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
    var output cognitoidentityprovider.GlobalSignOutOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GlobalSignOut, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) InitiateAuth(ctx workflow.Context, input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
    var output cognitoidentityprovider.InitiateAuthOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InitiateAuth, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListDevices(ctx workflow.Context, input *cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error) {
    var output cognitoidentityprovider.ListDevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDevices, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListGroups(ctx workflow.Context, input *cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error) {
    var output cognitoidentityprovider.ListGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListIdentityProviders(ctx workflow.Context, input *cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
    var output cognitoidentityprovider.ListIdentityProvidersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIdentityProviders, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListResourceServers(ctx workflow.Context, input *cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error) {
    var output cognitoidentityprovider.ListResourceServersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceServers, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListTagsForResource(ctx workflow.Context, input *cognitoidentityprovider.ListTagsForResourceInput) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
    var output cognitoidentityprovider.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListUserImportJobs(ctx workflow.Context, input *cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
    var output cognitoidentityprovider.ListUserImportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserImportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListUserPoolClients(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
    var output cognitoidentityprovider.ListUserPoolClientsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserPoolClients, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListUserPools(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
    var output cognitoidentityprovider.ListUserPoolsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserPools, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListUsers(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error) {
    var output cognitoidentityprovider.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ListUsersInGroup(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
    var output cognitoidentityprovider.ListUsersInGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsersInGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) ResendConfirmationCode(ctx workflow.Context, input *cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error) {
    var output cognitoidentityprovider.ResendConfirmationCodeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResendConfirmationCode, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) RespondToAuthChallenge(ctx workflow.Context, input *cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
    var output cognitoidentityprovider.RespondToAuthChallengeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RespondToAuthChallenge, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) SetRiskConfiguration(ctx workflow.Context, input *cognitoidentityprovider.SetRiskConfigurationInput) (*cognitoidentityprovider.SetRiskConfigurationOutput, error) {
    var output cognitoidentityprovider.SetRiskConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetRiskConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) SetUICustomization(ctx workflow.Context, input *cognitoidentityprovider.SetUICustomizationInput) (*cognitoidentityprovider.SetUICustomizationOutput, error) {
    var output cognitoidentityprovider.SetUICustomizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetUICustomization, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) SetUserMFAPreference(ctx workflow.Context, input *cognitoidentityprovider.SetUserMFAPreferenceInput) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error) {
    var output cognitoidentityprovider.SetUserMFAPreferenceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetUserMFAPreference, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) SetUserPoolMfaConfig(ctx workflow.Context, input *cognitoidentityprovider.SetUserPoolMfaConfigInput) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error) {
    var output cognitoidentityprovider.SetUserPoolMfaConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetUserPoolMfaConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) SetUserSettings(ctx workflow.Context, input *cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error) {
    var output cognitoidentityprovider.SetUserSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetUserSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) SignUp(ctx workflow.Context, input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error) {
    var output cognitoidentityprovider.SignUpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SignUp, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) StartUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error) {
    var output cognitoidentityprovider.StartUserImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartUserImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) StopUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error) {
    var output cognitoidentityprovider.StopUserImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopUserImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) TagResource(ctx workflow.Context, input *cognitoidentityprovider.TagResourceInput) (*cognitoidentityprovider.TagResourceOutput, error) {
    var output cognitoidentityprovider.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UntagResource(ctx workflow.Context, input *cognitoidentityprovider.UntagResourceInput) (*cognitoidentityprovider.UntagResourceOutput, error) {
    var output cognitoidentityprovider.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateAuthEventFeedback(ctx workflow.Context, input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error) {
    var output cognitoidentityprovider.UpdateAuthEventFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAuthEventFeedback, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateDeviceStatus(ctx workflow.Context, input *cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error) {
    var output cognitoidentityprovider.UpdateDeviceStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeviceStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateGroup(ctx workflow.Context, input *cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error) {
    var output cognitoidentityprovider.UpdateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.UpdateIdentityProviderInput) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error) {
    var output cognitoidentityprovider.UpdateIdentityProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIdentityProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateResourceServer(ctx workflow.Context, input *cognitoidentityprovider.UpdateResourceServerInput) (*cognitoidentityprovider.UpdateResourceServerOutput, error) {
    var output cognitoidentityprovider.UpdateResourceServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateResourceServer, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error) {
    var output cognitoidentityprovider.UpdateUserAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateUserPool(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error) {
    var output cognitoidentityprovider.UpdateUserPoolOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserPool, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error) {
    var output cognitoidentityprovider.UpdateUserPoolClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserPoolClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) UpdateUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolDomainInput) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error) {
    var output cognitoidentityprovider.UpdateUserPoolDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserPoolDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) VerifySoftwareToken(ctx workflow.Context, input *cognitoidentityprovider.VerifySoftwareTokenInput) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error) {
    var output cognitoidentityprovider.VerifySoftwareTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.VerifySoftwareToken, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoIdentityProviderStub) VerifyUserAttribute(ctx workflow.Context, input *cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error) {
    var output cognitoidentityprovider.VerifyUserAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.VerifyUserAttribute, input).Get(ctx, &output)
    return &output, err
}
