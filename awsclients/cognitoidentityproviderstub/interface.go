// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cognitoidentityproviderstub

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AddCustomAttributes(ctx workflow.Context, input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error)
	AddCustomAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.AddCustomAttributesInput) *CognitoIdentityProviderAddCustomAttributesFuture

	AdminAddUserToGroup(ctx workflow.Context, input *cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error)
	AdminAddUserToGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminAddUserToGroupInput) *CognitoIdentityProviderAdminAddUserToGroupFuture

	AdminConfirmSignUp(ctx workflow.Context, input *cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error)
	AdminConfirmSignUpAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminConfirmSignUpInput) *CognitoIdentityProviderAdminConfirmSignUpFuture

	AdminCreateUser(ctx workflow.Context, input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error)
	AdminCreateUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminCreateUserInput) *CognitoIdentityProviderAdminCreateUserFuture

	AdminDeleteUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error)
	AdminDeleteUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserInput) *CognitoIdentityProviderAdminDeleteUserFuture

	AdminDeleteUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error)
	AdminDeleteUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDeleteUserAttributesInput) *CognitoIdentityProviderAdminDeleteUserAttributesFuture

	AdminDisableProviderForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableProviderForUserInput) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error)
	AdminDisableProviderForUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableProviderForUserInput) *CognitoIdentityProviderAdminDisableProviderForUserFuture

	AdminDisableUser(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error)
	AdminDisableUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminDisableUserInput) *CognitoIdentityProviderAdminDisableUserFuture

	AdminEnableUser(ctx workflow.Context, input *cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error)
	AdminEnableUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminEnableUserInput) *CognitoIdentityProviderAdminEnableUserFuture

	AdminForgetDevice(ctx workflow.Context, input *cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error)
	AdminForgetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminForgetDeviceInput) *CognitoIdentityProviderAdminForgetDeviceFuture

	AdminGetDevice(ctx workflow.Context, input *cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error)
	AdminGetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminGetDeviceInput) *CognitoIdentityProviderAdminGetDeviceFuture

	AdminGetUser(ctx workflow.Context, input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error)
	AdminGetUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminGetUserInput) *CognitoIdentityProviderAdminGetUserFuture

	AdminInitiateAuth(ctx workflow.Context, input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
	AdminInitiateAuthAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminInitiateAuthInput) *CognitoIdentityProviderAdminInitiateAuthFuture

	AdminLinkProviderForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminLinkProviderForUserInput) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error)
	AdminLinkProviderForUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminLinkProviderForUserInput) *CognitoIdentityProviderAdminLinkProviderForUserFuture

	AdminListDevices(ctx workflow.Context, input *cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error)
	AdminListDevicesAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminListDevicesInput) *CognitoIdentityProviderAdminListDevicesFuture

	AdminListGroupsForUser(ctx workflow.Context, input *cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error)
	AdminListGroupsForUserAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminListGroupsForUserInput) *CognitoIdentityProviderAdminListGroupsForUserFuture

	AdminListUserAuthEvents(ctx workflow.Context, input *cognitoidentityprovider.AdminListUserAuthEventsInput) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error)
	AdminListUserAuthEventsAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminListUserAuthEventsInput) *CognitoIdentityProviderAdminListUserAuthEventsFuture

	AdminRemoveUserFromGroup(ctx workflow.Context, input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error)
	AdminRemoveUserFromGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) *CognitoIdentityProviderAdminRemoveUserFromGroupFuture

	AdminResetUserPassword(ctx workflow.Context, input *cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error)
	AdminResetUserPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminResetUserPasswordInput) *CognitoIdentityProviderAdminResetUserPasswordFuture

	AdminRespondToAuthChallenge(ctx workflow.Context, input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error)
	AdminRespondToAuthChallengeAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) *CognitoIdentityProviderAdminRespondToAuthChallengeFuture

	AdminSetUserMFAPreference(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error)
	AdminSetUserMFAPreferenceAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) *CognitoIdentityProviderAdminSetUserMFAPreferenceFuture

	AdminSetUserPassword(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserPasswordInput) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error)
	AdminSetUserPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserPasswordInput) *CognitoIdentityProviderAdminSetUserPasswordFuture

	AdminSetUserSettings(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error)
	AdminSetUserSettingsAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminSetUserSettingsInput) *CognitoIdentityProviderAdminSetUserSettingsFuture

	AdminUpdateAuthEventFeedback(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error)
	AdminUpdateAuthEventFeedbackAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) *CognitoIdentityProviderAdminUpdateAuthEventFeedbackFuture

	AdminUpdateDeviceStatus(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error)
	AdminUpdateDeviceStatusAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) *CognitoIdentityProviderAdminUpdateDeviceStatusFuture

	AdminUpdateUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error)
	AdminUpdateUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUpdateUserAttributesInput) *CognitoIdentityProviderAdminUpdateUserAttributesFuture

	AdminUserGlobalSignOut(ctx workflow.Context, input *cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error)
	AdminUserGlobalSignOutAsync(ctx workflow.Context, input *cognitoidentityprovider.AdminUserGlobalSignOutInput) *CognitoIdentityProviderAdminUserGlobalSignOutFuture

	AssociateSoftwareToken(ctx workflow.Context, input *cognitoidentityprovider.AssociateSoftwareTokenInput) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error)
	AssociateSoftwareTokenAsync(ctx workflow.Context, input *cognitoidentityprovider.AssociateSoftwareTokenInput) *CognitoIdentityProviderAssociateSoftwareTokenFuture

	ChangePassword(ctx workflow.Context, input *cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error)
	ChangePasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.ChangePasswordInput) *CognitoIdentityProviderChangePasswordFuture

	ConfirmDevice(ctx workflow.Context, input *cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error)
	ConfirmDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.ConfirmDeviceInput) *CognitoIdentityProviderConfirmDeviceFuture

	ConfirmForgotPassword(ctx workflow.Context, input *cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ConfirmForgotPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.ConfirmForgotPasswordInput) *CognitoIdentityProviderConfirmForgotPasswordFuture

	ConfirmSignUp(ctx workflow.Context, input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	ConfirmSignUpAsync(ctx workflow.Context, input *cognitoidentityprovider.ConfirmSignUpInput) *CognitoIdentityProviderConfirmSignUpFuture

	CreateGroup(ctx workflow.Context, input *cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error)
	CreateGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateGroupInput) *CognitoIdentityProviderCreateGroupFuture

	CreateIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.CreateIdentityProviderInput) (*cognitoidentityprovider.CreateIdentityProviderOutput, error)
	CreateIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateIdentityProviderInput) *CognitoIdentityProviderCreateIdentityProviderFuture

	CreateResourceServer(ctx workflow.Context, input *cognitoidentityprovider.CreateResourceServerInput) (*cognitoidentityprovider.CreateResourceServerOutput, error)
	CreateResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateResourceServerInput) *CognitoIdentityProviderCreateResourceServerFuture

	CreateUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error)
	CreateUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserImportJobInput) *CognitoIdentityProviderCreateUserImportJobFuture

	CreateUserPool(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error)
	CreateUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolInput) *CognitoIdentityProviderCreateUserPoolFuture

	CreateUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	CreateUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolClientInput) *CognitoIdentityProviderCreateUserPoolClientFuture

	CreateUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolDomainInput) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error)
	CreateUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.CreateUserPoolDomainInput) *CognitoIdentityProviderCreateUserPoolDomainFuture

	DeleteGroup(ctx workflow.Context, input *cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error)
	DeleteGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteGroupInput) *CognitoIdentityProviderDeleteGroupFuture

	DeleteIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.DeleteIdentityProviderInput) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error)
	DeleteIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteIdentityProviderInput) *CognitoIdentityProviderDeleteIdentityProviderFuture

	DeleteResourceServer(ctx workflow.Context, input *cognitoidentityprovider.DeleteResourceServerInput) (*cognitoidentityprovider.DeleteResourceServerOutput, error)
	DeleteResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteResourceServerInput) *CognitoIdentityProviderDeleteResourceServerFuture

	DeleteUser(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserInput) *CognitoIdentityProviderDeleteUserFuture

	DeleteUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error)
	DeleteUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserAttributesInput) *CognitoIdentityProviderDeleteUserAttributesFuture

	DeleteUserPool(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error)
	DeleteUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolInput) *CognitoIdentityProviderDeleteUserPoolFuture

	DeleteUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error)
	DeleteUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolClientInput) *CognitoIdentityProviderDeleteUserPoolClientFuture

	DeleteUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolDomainInput) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error)
	DeleteUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.DeleteUserPoolDomainInput) *CognitoIdentityProviderDeleteUserPoolDomainFuture

	DescribeIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	DescribeIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput) *CognitoIdentityProviderDescribeIdentityProviderFuture

	DescribeResourceServer(ctx workflow.Context, input *cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error)
	DescribeResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeResourceServerInput) *CognitoIdentityProviderDescribeResourceServerFuture

	DescribeRiskConfiguration(ctx workflow.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error)
	DescribeRiskConfigurationAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput) *CognitoIdentityProviderDescribeRiskConfigurationFuture

	DescribeUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error)
	DescribeUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserImportJobInput) *CognitoIdentityProviderDescribeUserImportJobFuture

	DescribeUserPool(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	DescribeUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolInput) *CognitoIdentityProviderDescribeUserPoolFuture

	DescribeUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error)
	DescribeUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput) *CognitoIdentityProviderDescribeUserPoolClientFuture

	DescribeUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error)
	DescribeUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput) *CognitoIdentityProviderDescribeUserPoolDomainFuture

	ForgetDevice(ctx workflow.Context, input *cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error)
	ForgetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.ForgetDeviceInput) *CognitoIdentityProviderForgetDeviceFuture

	ForgotPassword(ctx workflow.Context, input *cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ForgotPasswordAsync(ctx workflow.Context, input *cognitoidentityprovider.ForgotPasswordInput) *CognitoIdentityProviderForgotPasswordFuture

	GetCSVHeader(ctx workflow.Context, input *cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error)
	GetCSVHeaderAsync(ctx workflow.Context, input *cognitoidentityprovider.GetCSVHeaderInput) *CognitoIdentityProviderGetCSVHeaderFuture

	GetDevice(ctx workflow.Context, input *cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error)
	GetDeviceAsync(ctx workflow.Context, input *cognitoidentityprovider.GetDeviceInput) *CognitoIdentityProviderGetDeviceFuture

	GetGroup(ctx workflow.Context, input *cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error)
	GetGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.GetGroupInput) *CognitoIdentityProviderGetGroupFuture

	GetIdentityProviderByIdentifier(ctx workflow.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error)
	GetIdentityProviderByIdentifierAsync(ctx workflow.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) *CognitoIdentityProviderGetIdentityProviderByIdentifierFuture

	GetSigningCertificate(ctx workflow.Context, input *cognitoidentityprovider.GetSigningCertificateInput) (*cognitoidentityprovider.GetSigningCertificateOutput, error)
	GetSigningCertificateAsync(ctx workflow.Context, input *cognitoidentityprovider.GetSigningCertificateInput) *CognitoIdentityProviderGetSigningCertificateFuture

	GetUICustomization(ctx workflow.Context, input *cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error)
	GetUICustomizationAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUICustomizationInput) *CognitoIdentityProviderGetUICustomizationFuture

	GetUser(ctx workflow.Context, input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error)
	GetUserAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUserInput) *CognitoIdentityProviderGetUserFuture

	GetUserAttributeVerificationCode(ctx workflow.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error)
	GetUserAttributeVerificationCodeAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) *CognitoIdentityProviderGetUserAttributeVerificationCodeFuture

	GetUserPoolMfaConfig(ctx workflow.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error)
	GetUserPoolMfaConfigAsync(ctx workflow.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput) *CognitoIdentityProviderGetUserPoolMfaConfigFuture

	GlobalSignOut(ctx workflow.Context, input *cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error)
	GlobalSignOutAsync(ctx workflow.Context, input *cognitoidentityprovider.GlobalSignOutInput) *CognitoIdentityProviderGlobalSignOutFuture

	InitiateAuth(ctx workflow.Context, input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error)
	InitiateAuthAsync(ctx workflow.Context, input *cognitoidentityprovider.InitiateAuthInput) *CognitoIdentityProviderInitiateAuthFuture

	ListDevices(ctx workflow.Context, input *cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error)
	ListDevicesAsync(ctx workflow.Context, input *cognitoidentityprovider.ListDevicesInput) *CognitoIdentityProviderListDevicesFuture

	ListGroups(ctx workflow.Context, input *cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListGroupsInput) *CognitoIdentityProviderListGroupsFuture

	ListIdentityProviders(ctx workflow.Context, input *cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	ListIdentityProvidersAsync(ctx workflow.Context, input *cognitoidentityprovider.ListIdentityProvidersInput) *CognitoIdentityProviderListIdentityProvidersFuture

	ListResourceServers(ctx workflow.Context, input *cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error)
	ListResourceServersAsync(ctx workflow.Context, input *cognitoidentityprovider.ListResourceServersInput) *CognitoIdentityProviderListResourceServersFuture

	ListTagsForResource(ctx workflow.Context, input *cognitoidentityprovider.ListTagsForResourceInput) (*cognitoidentityprovider.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cognitoidentityprovider.ListTagsForResourceInput) *CognitoIdentityProviderListTagsForResourceFuture

	ListUserImportJobs(ctx workflow.Context, input *cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error)
	ListUserImportJobsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUserImportJobsInput) *CognitoIdentityProviderListUserImportJobsFuture

	ListUserPoolClients(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error)
	ListUserPoolClientsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolClientsInput) *CognitoIdentityProviderListUserPoolClientsFuture

	ListUserPools(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error)
	ListUserPoolsAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUserPoolsInput) *CognitoIdentityProviderListUserPoolsFuture

	ListUsers(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInput) *CognitoIdentityProviderListUsersFuture

	ListUsersInGroup(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error)
	ListUsersInGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.ListUsersInGroupInput) *CognitoIdentityProviderListUsersInGroupFuture

	ResendConfirmationCode(ctx workflow.Context, input *cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	ResendConfirmationCodeAsync(ctx workflow.Context, input *cognitoidentityprovider.ResendConfirmationCodeInput) *CognitoIdentityProviderResendConfirmationCodeFuture

	RespondToAuthChallenge(ctx workflow.Context, input *cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
	RespondToAuthChallengeAsync(ctx workflow.Context, input *cognitoidentityprovider.RespondToAuthChallengeInput) *CognitoIdentityProviderRespondToAuthChallengeFuture

	SetRiskConfiguration(ctx workflow.Context, input *cognitoidentityprovider.SetRiskConfigurationInput) (*cognitoidentityprovider.SetRiskConfigurationOutput, error)
	SetRiskConfigurationAsync(ctx workflow.Context, input *cognitoidentityprovider.SetRiskConfigurationInput) *CognitoIdentityProviderSetRiskConfigurationFuture

	SetUICustomization(ctx workflow.Context, input *cognitoidentityprovider.SetUICustomizationInput) (*cognitoidentityprovider.SetUICustomizationOutput, error)
	SetUICustomizationAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUICustomizationInput) *CognitoIdentityProviderSetUICustomizationFuture

	SetUserMFAPreference(ctx workflow.Context, input *cognitoidentityprovider.SetUserMFAPreferenceInput) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error)
	SetUserMFAPreferenceAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUserMFAPreferenceInput) *CognitoIdentityProviderSetUserMFAPreferenceFuture

	SetUserPoolMfaConfig(ctx workflow.Context, input *cognitoidentityprovider.SetUserPoolMfaConfigInput) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error)
	SetUserPoolMfaConfigAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUserPoolMfaConfigInput) *CognitoIdentityProviderSetUserPoolMfaConfigFuture

	SetUserSettings(ctx workflow.Context, input *cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error)
	SetUserSettingsAsync(ctx workflow.Context, input *cognitoidentityprovider.SetUserSettingsInput) *CognitoIdentityProviderSetUserSettingsFuture

	SignUp(ctx workflow.Context, input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error)
	SignUpAsync(ctx workflow.Context, input *cognitoidentityprovider.SignUpInput) *CognitoIdentityProviderSignUpFuture

	StartUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error)
	StartUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.StartUserImportJobInput) *CognitoIdentityProviderStartUserImportJobFuture

	StopUserImportJob(ctx workflow.Context, input *cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error)
	StopUserImportJobAsync(ctx workflow.Context, input *cognitoidentityprovider.StopUserImportJobInput) *CognitoIdentityProviderStopUserImportJobFuture

	TagResource(ctx workflow.Context, input *cognitoidentityprovider.TagResourceInput) (*cognitoidentityprovider.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cognitoidentityprovider.TagResourceInput) *CognitoIdentityProviderTagResourceFuture

	UntagResource(ctx workflow.Context, input *cognitoidentityprovider.UntagResourceInput) (*cognitoidentityprovider.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cognitoidentityprovider.UntagResourceInput) *CognitoIdentityProviderUntagResourceFuture

	UpdateAuthEventFeedback(ctx workflow.Context, input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error)
	UpdateAuthEventFeedbackAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) *CognitoIdentityProviderUpdateAuthEventFeedbackFuture

	UpdateDeviceStatus(ctx workflow.Context, input *cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error)
	UpdateDeviceStatusAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateDeviceStatusInput) *CognitoIdentityProviderUpdateDeviceStatusFuture

	UpdateGroup(ctx workflow.Context, input *cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error)
	UpdateGroupAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateGroupInput) *CognitoIdentityProviderUpdateGroupFuture

	UpdateIdentityProvider(ctx workflow.Context, input *cognitoidentityprovider.UpdateIdentityProviderInput) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error)
	UpdateIdentityProviderAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateIdentityProviderInput) *CognitoIdentityProviderUpdateIdentityProviderFuture

	UpdateResourceServer(ctx workflow.Context, input *cognitoidentityprovider.UpdateResourceServerInput) (*cognitoidentityprovider.UpdateResourceServerOutput, error)
	UpdateResourceServerAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateResourceServerInput) *CognitoIdentityProviderUpdateResourceServerFuture

	UpdateUserAttributes(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error)
	UpdateUserAttributesAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserAttributesInput) *CognitoIdentityProviderUpdateUserAttributesFuture

	UpdateUserPool(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error)
	UpdateUserPoolAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolInput) *CognitoIdentityProviderUpdateUserPoolFuture

	UpdateUserPoolClient(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error)
	UpdateUserPoolClientAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolClientInput) *CognitoIdentityProviderUpdateUserPoolClientFuture

	UpdateUserPoolDomain(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolDomainInput) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error)
	UpdateUserPoolDomainAsync(ctx workflow.Context, input *cognitoidentityprovider.UpdateUserPoolDomainInput) *CognitoIdentityProviderUpdateUserPoolDomainFuture

	VerifySoftwareToken(ctx workflow.Context, input *cognitoidentityprovider.VerifySoftwareTokenInput) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error)
	VerifySoftwareTokenAsync(ctx workflow.Context, input *cognitoidentityprovider.VerifySoftwareTokenInput) *CognitoIdentityProviderVerifySoftwareTokenFuture

	VerifyUserAttribute(ctx workflow.Context, input *cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error)
	VerifyUserAttributeAsync(ctx workflow.Context, input *cognitoidentityprovider.VerifyUserAttributeInput) *CognitoIdentityProviderVerifyUserAttributeFuture
}

func NewClient() Client {
	return &stub{}
}
