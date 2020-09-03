package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

type CognitoIdentityProviderActivities struct {
	client cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func NewCognitoIdentityProviderActivities(client cognitoidentityprovideriface.CognitoIdentityProviderAPI) *CognitoIdentityProviderActivities {
    return &CognitoIdentityProviderActivities{client: client}
}

func (a *CognitoIdentityProviderActivities) AddCustomAttributes(input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
    return a.client.AddCustomAttributes(input)
}

func (a *CognitoIdentityProviderActivities) AdminAddUserToGroup(input *cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
    return a.client.AdminAddUserToGroup(input)
}

func (a *CognitoIdentityProviderActivities) AdminConfirmSignUp(input *cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error) {
    return a.client.AdminConfirmSignUp(input)
}

func (a *CognitoIdentityProviderActivities) AdminCreateUser(input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
    return a.client.AdminCreateUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminDeleteUser(input *cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
    return a.client.AdminDeleteUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminDeleteUserAttributes(input *cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error) {
    return a.client.AdminDeleteUserAttributes(input)
}

func (a *CognitoIdentityProviderActivities) AdminDisableProviderForUser(input *cognitoidentityprovider.AdminDisableProviderForUserInput) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error) {
    return a.client.AdminDisableProviderForUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminDisableUser(input *cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error) {
    return a.client.AdminDisableUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminEnableUser(input *cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error) {
    return a.client.AdminEnableUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminForgetDevice(input *cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error) {
    return a.client.AdminForgetDevice(input)
}

func (a *CognitoIdentityProviderActivities) AdminGetDevice(input *cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error) {
    return a.client.AdminGetDevice(input)
}

func (a *CognitoIdentityProviderActivities) AdminGetUser(input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error) {
    return a.client.AdminGetUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminInitiateAuth(input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
    return a.client.AdminInitiateAuth(input)
}

func (a *CognitoIdentityProviderActivities) AdminLinkProviderForUser(input *cognitoidentityprovider.AdminLinkProviderForUserInput) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error) {
    return a.client.AdminLinkProviderForUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminListDevices(input *cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error) {
    return a.client.AdminListDevices(input)
}

func (a *CognitoIdentityProviderActivities) AdminListGroupsForUser(input *cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
    return a.client.AdminListGroupsForUser(input)
}

func (a *CognitoIdentityProviderActivities) AdminListUserAuthEvents(input *cognitoidentityprovider.AdminListUserAuthEventsInput) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error) {
    return a.client.AdminListUserAuthEvents(input)
}

func (a *CognitoIdentityProviderActivities) AdminRemoveUserFromGroup(input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error) {
    return a.client.AdminRemoveUserFromGroup(input)
}

func (a *CognitoIdentityProviderActivities) AdminResetUserPassword(input *cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error) {
    return a.client.AdminResetUserPassword(input)
}

func (a *CognitoIdentityProviderActivities) AdminRespondToAuthChallenge(input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {
    return a.client.AdminRespondToAuthChallenge(input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserMFAPreference(input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error) {
    return a.client.AdminSetUserMFAPreference(input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserPassword(input *cognitoidentityprovider.AdminSetUserPasswordInput) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
    return a.client.AdminSetUserPassword(input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserSettings(input *cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error) {
    return a.client.AdminSetUserSettings(input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateAuthEventFeedback(input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error) {
    return a.client.AdminUpdateAuthEventFeedback(input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateDeviceStatus(input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error) {
    return a.client.AdminUpdateDeviceStatus(input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateUserAttributes(input *cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
    return a.client.AdminUpdateUserAttributes(input)
}

func (a *CognitoIdentityProviderActivities) AdminUserGlobalSignOut(input *cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error) {
    return a.client.AdminUserGlobalSignOut(input)
}

func (a *CognitoIdentityProviderActivities) AssociateSoftwareToken(input *cognitoidentityprovider.AssociateSoftwareTokenInput) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error) {
    return a.client.AssociateSoftwareToken(input)
}

func (a *CognitoIdentityProviderActivities) ChangePassword(input *cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error) {
    return a.client.ChangePassword(input)
}

func (a *CognitoIdentityProviderActivities) ConfirmDevice(input *cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error) {
    return a.client.ConfirmDevice(input)
}

func (a *CognitoIdentityProviderActivities) ConfirmForgotPassword(input *cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {
    return a.client.ConfirmForgotPassword(input)
}

func (a *CognitoIdentityProviderActivities) ConfirmSignUp(input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
    return a.client.ConfirmSignUp(input)
}

func (a *CognitoIdentityProviderActivities) CreateGroup(input *cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error) {
    return a.client.CreateGroup(input)
}

func (a *CognitoIdentityProviderActivities) CreateIdentityProvider(input *cognitoidentityprovider.CreateIdentityProviderInput) (*cognitoidentityprovider.CreateIdentityProviderOutput, error) {
    return a.client.CreateIdentityProvider(input)
}

func (a *CognitoIdentityProviderActivities) CreateResourceServer(input *cognitoidentityprovider.CreateResourceServerInput) (*cognitoidentityprovider.CreateResourceServerOutput, error) {
    return a.client.CreateResourceServer(input)
}

func (a *CognitoIdentityProviderActivities) CreateUserImportJob(input *cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error) {
    return a.client.CreateUserImportJob(input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPool(input *cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error) {
    return a.client.CreateUserPool(input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPoolClient(input *cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error) {
    return a.client.CreateUserPoolClient(input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPoolDomain(input *cognitoidentityprovider.CreateUserPoolDomainInput) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error) {
    return a.client.CreateUserPoolDomain(input)
}

func (a *CognitoIdentityProviderActivities) DeleteGroup(input *cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error) {
    return a.client.DeleteGroup(input)
}

func (a *CognitoIdentityProviderActivities) DeleteIdentityProvider(input *cognitoidentityprovider.DeleteIdentityProviderInput) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error) {
    return a.client.DeleteIdentityProvider(input)
}

func (a *CognitoIdentityProviderActivities) DeleteResourceServer(input *cognitoidentityprovider.DeleteResourceServerInput) (*cognitoidentityprovider.DeleteResourceServerOutput, error) {
    return a.client.DeleteResourceServer(input)
}

func (a *CognitoIdentityProviderActivities) DeleteUser(input *cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserAttributes(input *cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error) {
    return a.client.DeleteUserAttributes(input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPool(input *cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error) {
    return a.client.DeleteUserPool(input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPoolClient(input *cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error) {
    return a.client.DeleteUserPoolClient(input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPoolDomain(input *cognitoidentityprovider.DeleteUserPoolDomainInput) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error) {
    return a.client.DeleteUserPoolDomain(input)
}

func (a *CognitoIdentityProviderActivities) DescribeIdentityProvider(input *cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
    return a.client.DescribeIdentityProvider(input)
}

func (a *CognitoIdentityProviderActivities) DescribeResourceServer(input *cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
    return a.client.DescribeResourceServer(input)
}

func (a *CognitoIdentityProviderActivities) DescribeRiskConfiguration(input *cognitoidentityprovider.DescribeRiskConfigurationInput) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
    return a.client.DescribeRiskConfiguration(input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserImportJob(input *cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
    return a.client.DescribeUserImportJob(input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPool(input *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
    return a.client.DescribeUserPool(input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPoolClient(input *cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
    return a.client.DescribeUserPoolClient(input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPoolDomain(input *cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
    return a.client.DescribeUserPoolDomain(input)
}

func (a *CognitoIdentityProviderActivities) ForgetDevice(input *cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error) {
    return a.client.ForgetDevice(input)
}

func (a *CognitoIdentityProviderActivities) ForgotPassword(input *cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
    return a.client.ForgotPassword(input)
}

func (a *CognitoIdentityProviderActivities) GetCSVHeader(input *cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
    return a.client.GetCSVHeader(input)
}

func (a *CognitoIdentityProviderActivities) GetDevice(input *cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error) {
    return a.client.GetDevice(input)
}

func (a *CognitoIdentityProviderActivities) GetGroup(input *cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error) {
    return a.client.GetGroup(input)
}

func (a *CognitoIdentityProviderActivities) GetIdentityProviderByIdentifier(input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
    return a.client.GetIdentityProviderByIdentifier(input)
}

func (a *CognitoIdentityProviderActivities) GetSigningCertificate(input *cognitoidentityprovider.GetSigningCertificateInput) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
    return a.client.GetSigningCertificate(input)
}

func (a *CognitoIdentityProviderActivities) GetUICustomization(input *cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
    return a.client.GetUICustomization(input)
}

func (a *CognitoIdentityProviderActivities) GetUser(input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
    return a.client.GetUser(input)
}

func (a *CognitoIdentityProviderActivities) GetUserAttributeVerificationCode(input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
    return a.client.GetUserAttributeVerificationCode(input)
}

func (a *CognitoIdentityProviderActivities) GetUserPoolMfaConfig(input *cognitoidentityprovider.GetUserPoolMfaConfigInput) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
    return a.client.GetUserPoolMfaConfig(input)
}

func (a *CognitoIdentityProviderActivities) GlobalSignOut(input *cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
    return a.client.GlobalSignOut(input)
}

func (a *CognitoIdentityProviderActivities) InitiateAuth(input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
    return a.client.InitiateAuth(input)
}

func (a *CognitoIdentityProviderActivities) ListDevices(input *cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error) {
    return a.client.ListDevices(input)
}

func (a *CognitoIdentityProviderActivities) ListGroups(input *cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error) {
    return a.client.ListGroups(input)
}

func (a *CognitoIdentityProviderActivities) ListIdentityProviders(input *cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
    return a.client.ListIdentityProviders(input)
}

func (a *CognitoIdentityProviderActivities) ListResourceServers(input *cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error) {
    return a.client.ListResourceServers(input)
}

func (a *CognitoIdentityProviderActivities) ListTagsForResource(input *cognitoidentityprovider.ListTagsForResourceInput) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CognitoIdentityProviderActivities) ListUserImportJobs(input *cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
    return a.client.ListUserImportJobs(input)
}

func (a *CognitoIdentityProviderActivities) ListUserPoolClients(input *cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
    return a.client.ListUserPoolClients(input)
}

func (a *CognitoIdentityProviderActivities) ListUserPools(input *cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
    return a.client.ListUserPools(input)
}

func (a *CognitoIdentityProviderActivities) ListUsers(input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *CognitoIdentityProviderActivities) ListUsersInGroup(input *cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
    return a.client.ListUsersInGroup(input)
}

func (a *CognitoIdentityProviderActivities) ResendConfirmationCode(input *cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error) {
    return a.client.ResendConfirmationCode(input)
}

func (a *CognitoIdentityProviderActivities) RespondToAuthChallenge(input *cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
    return a.client.RespondToAuthChallenge(input)
}

func (a *CognitoIdentityProviderActivities) SetRiskConfiguration(input *cognitoidentityprovider.SetRiskConfigurationInput) (*cognitoidentityprovider.SetRiskConfigurationOutput, error) {
    return a.client.SetRiskConfiguration(input)
}

func (a *CognitoIdentityProviderActivities) SetUICustomization(input *cognitoidentityprovider.SetUICustomizationInput) (*cognitoidentityprovider.SetUICustomizationOutput, error) {
    return a.client.SetUICustomization(input)
}

func (a *CognitoIdentityProviderActivities) SetUserMFAPreference(input *cognitoidentityprovider.SetUserMFAPreferenceInput) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error) {
    return a.client.SetUserMFAPreference(input)
}

func (a *CognitoIdentityProviderActivities) SetUserPoolMfaConfig(input *cognitoidentityprovider.SetUserPoolMfaConfigInput) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error) {
    return a.client.SetUserPoolMfaConfig(input)
}

func (a *CognitoIdentityProviderActivities) SetUserSettings(input *cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error) {
    return a.client.SetUserSettings(input)
}

func (a *CognitoIdentityProviderActivities) SignUp(input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error) {
    return a.client.SignUp(input)
}

func (a *CognitoIdentityProviderActivities) StartUserImportJob(input *cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error) {
    return a.client.StartUserImportJob(input)
}

func (a *CognitoIdentityProviderActivities) StopUserImportJob(input *cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error) {
    return a.client.StopUserImportJob(input)
}

func (a *CognitoIdentityProviderActivities) TagResource(input *cognitoidentityprovider.TagResourceInput) (*cognitoidentityprovider.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CognitoIdentityProviderActivities) UntagResource(input *cognitoidentityprovider.UntagResourceInput) (*cognitoidentityprovider.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CognitoIdentityProviderActivities) UpdateAuthEventFeedback(input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error) {
    return a.client.UpdateAuthEventFeedback(input)
}

func (a *CognitoIdentityProviderActivities) UpdateDeviceStatus(input *cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error) {
    return a.client.UpdateDeviceStatus(input)
}

func (a *CognitoIdentityProviderActivities) UpdateGroup(input *cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error) {
    return a.client.UpdateGroup(input)
}

func (a *CognitoIdentityProviderActivities) UpdateIdentityProvider(input *cognitoidentityprovider.UpdateIdentityProviderInput) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error) {
    return a.client.UpdateIdentityProvider(input)
}

func (a *CognitoIdentityProviderActivities) UpdateResourceServer(input *cognitoidentityprovider.UpdateResourceServerInput) (*cognitoidentityprovider.UpdateResourceServerOutput, error) {
    return a.client.UpdateResourceServer(input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserAttributes(input *cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error) {
    return a.client.UpdateUserAttributes(input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPool(input *cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error) {
    return a.client.UpdateUserPool(input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPoolClient(input *cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error) {
    return a.client.UpdateUserPoolClient(input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPoolDomain(input *cognitoidentityprovider.UpdateUserPoolDomainInput) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error) {
    return a.client.UpdateUserPoolDomain(input)
}

func (a *CognitoIdentityProviderActivities) VerifySoftwareToken(input *cognitoidentityprovider.VerifySoftwareTokenInput) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error) {
    return a.client.VerifySoftwareToken(input)
}

func (a *CognitoIdentityProviderActivities) VerifyUserAttribute(input *cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error) {
    return a.client.VerifyUserAttribute(input)
}
