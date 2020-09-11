
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CognitoIdentityProviderActivities struct {
    client cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func NewCognitoIdentityProviderActivities(session *session.Session, config ...*aws.Config) *CognitoIdentityProviderActivities {
    client := cognitoidentityprovider.New(session, config...)
    return &CognitoIdentityProviderActivities{client: client}
}

func (a *CognitoIdentityProviderActivities) AddCustomAttributes(ctx context.Context, input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
    return a.client.AddCustomAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminAddUserToGroup(ctx context.Context, input *cognitoidentityprovider.AdminAddUserToGroupInput) (*cognitoidentityprovider.AdminAddUserToGroupOutput, error) {
    return a.client.AdminAddUserToGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminConfirmSignUp(ctx context.Context, input *cognitoidentityprovider.AdminConfirmSignUpInput) (*cognitoidentityprovider.AdminConfirmSignUpOutput, error) {
    return a.client.AdminConfirmSignUpWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminCreateUser(ctx context.Context, input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
    return a.client.AdminCreateUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDeleteUser(ctx context.Context, input *cognitoidentityprovider.AdminDeleteUserInput) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
    return a.client.AdminDeleteUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDeleteUserAttributes(ctx context.Context, input *cognitoidentityprovider.AdminDeleteUserAttributesInput) (*cognitoidentityprovider.AdminDeleteUserAttributesOutput, error) {
    return a.client.AdminDeleteUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDisableProviderForUser(ctx context.Context, input *cognitoidentityprovider.AdminDisableProviderForUserInput) (*cognitoidentityprovider.AdminDisableProviderForUserOutput, error) {
    return a.client.AdminDisableProviderForUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminDisableUser(ctx context.Context, input *cognitoidentityprovider.AdminDisableUserInput) (*cognitoidentityprovider.AdminDisableUserOutput, error) {
    return a.client.AdminDisableUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminEnableUser(ctx context.Context, input *cognitoidentityprovider.AdminEnableUserInput) (*cognitoidentityprovider.AdminEnableUserOutput, error) {
    return a.client.AdminEnableUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminForgetDevice(ctx context.Context, input *cognitoidentityprovider.AdminForgetDeviceInput) (*cognitoidentityprovider.AdminForgetDeviceOutput, error) {
    return a.client.AdminForgetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminGetDevice(ctx context.Context, input *cognitoidentityprovider.AdminGetDeviceInput) (*cognitoidentityprovider.AdminGetDeviceOutput, error) {
    return a.client.AdminGetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminGetUser(ctx context.Context, input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error) {
    return a.client.AdminGetUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminInitiateAuth(ctx context.Context, input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
    return a.client.AdminInitiateAuthWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminLinkProviderForUser(ctx context.Context, input *cognitoidentityprovider.AdminLinkProviderForUserInput) (*cognitoidentityprovider.AdminLinkProviderForUserOutput, error) {
    return a.client.AdminLinkProviderForUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminListDevices(ctx context.Context, input *cognitoidentityprovider.AdminListDevicesInput) (*cognitoidentityprovider.AdminListDevicesOutput, error) {
    return a.client.AdminListDevicesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminListGroupsForUser(ctx context.Context, input *cognitoidentityprovider.AdminListGroupsForUserInput) (*cognitoidentityprovider.AdminListGroupsForUserOutput, error) {
    return a.client.AdminListGroupsForUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminListUserAuthEvents(ctx context.Context, input *cognitoidentityprovider.AdminListUserAuthEventsInput) (*cognitoidentityprovider.AdminListUserAuthEventsOutput, error) {
    return a.client.AdminListUserAuthEventsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminRemoveUserFromGroup(ctx context.Context, input *cognitoidentityprovider.AdminRemoveUserFromGroupInput) (*cognitoidentityprovider.AdminRemoveUserFromGroupOutput, error) {
    return a.client.AdminRemoveUserFromGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminResetUserPassword(ctx context.Context, input *cognitoidentityprovider.AdminResetUserPasswordInput) (*cognitoidentityprovider.AdminResetUserPasswordOutput, error) {
    return a.client.AdminResetUserPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminRespondToAuthChallenge(ctx context.Context, input *cognitoidentityprovider.AdminRespondToAuthChallengeInput) (*cognitoidentityprovider.AdminRespondToAuthChallengeOutput, error) {
    return a.client.AdminRespondToAuthChallengeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserMFAPreference(ctx context.Context, input *cognitoidentityprovider.AdminSetUserMFAPreferenceInput) (*cognitoidentityprovider.AdminSetUserMFAPreferenceOutput, error) {
    return a.client.AdminSetUserMFAPreferenceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserPassword(ctx context.Context, input *cognitoidentityprovider.AdminSetUserPasswordInput) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
    return a.client.AdminSetUserPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminSetUserSettings(ctx context.Context, input *cognitoidentityprovider.AdminSetUserSettingsInput) (*cognitoidentityprovider.AdminSetUserSettingsOutput, error) {
    return a.client.AdminSetUserSettingsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateAuthEventFeedback(ctx context.Context, input *cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) (*cognitoidentityprovider.AdminUpdateAuthEventFeedbackOutput, error) {
    return a.client.AdminUpdateAuthEventFeedbackWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateDeviceStatus(ctx context.Context, input *cognitoidentityprovider.AdminUpdateDeviceStatusInput) (*cognitoidentityprovider.AdminUpdateDeviceStatusOutput, error) {
    return a.client.AdminUpdateDeviceStatusWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUpdateUserAttributes(ctx context.Context, input *cognitoidentityprovider.AdminUpdateUserAttributesInput) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
    return a.client.AdminUpdateUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AdminUserGlobalSignOut(ctx context.Context, input *cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error) {
    return a.client.AdminUserGlobalSignOutWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) AssociateSoftwareToken(ctx context.Context, input *cognitoidentityprovider.AssociateSoftwareTokenInput) (*cognitoidentityprovider.AssociateSoftwareTokenOutput, error) {
    return a.client.AssociateSoftwareTokenWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ChangePassword(ctx context.Context, input *cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error) {
    return a.client.ChangePasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ConfirmDevice(ctx context.Context, input *cognitoidentityprovider.ConfirmDeviceInput) (*cognitoidentityprovider.ConfirmDeviceOutput, error) {
    return a.client.ConfirmDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ConfirmForgotPassword(ctx context.Context, input *cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {
    return a.client.ConfirmForgotPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ConfirmSignUp(ctx context.Context, input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
    return a.client.ConfirmSignUpWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateGroup(ctx context.Context, input *cognitoidentityprovider.CreateGroupInput) (*cognitoidentityprovider.CreateGroupOutput, error) {
    return a.client.CreateGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateIdentityProvider(ctx context.Context, input *cognitoidentityprovider.CreateIdentityProviderInput) (*cognitoidentityprovider.CreateIdentityProviderOutput, error) {
    return a.client.CreateIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateResourceServer(ctx context.Context, input *cognitoidentityprovider.CreateResourceServerInput) (*cognitoidentityprovider.CreateResourceServerOutput, error) {
    return a.client.CreateResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserImportJob(ctx context.Context, input *cognitoidentityprovider.CreateUserImportJobInput) (*cognitoidentityprovider.CreateUserImportJobOutput, error) {
    return a.client.CreateUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPool(ctx context.Context, input *cognitoidentityprovider.CreateUserPoolInput) (*cognitoidentityprovider.CreateUserPoolOutput, error) {
    return a.client.CreateUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPoolClient(ctx context.Context, input *cognitoidentityprovider.CreateUserPoolClientInput) (*cognitoidentityprovider.CreateUserPoolClientOutput, error) {
    return a.client.CreateUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) CreateUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.CreateUserPoolDomainInput) (*cognitoidentityprovider.CreateUserPoolDomainOutput, error) {
    return a.client.CreateUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteGroup(ctx context.Context, input *cognitoidentityprovider.DeleteGroupInput) (*cognitoidentityprovider.DeleteGroupOutput, error) {
    return a.client.DeleteGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteIdentityProvider(ctx context.Context, input *cognitoidentityprovider.DeleteIdentityProviderInput) (*cognitoidentityprovider.DeleteIdentityProviderOutput, error) {
    return a.client.DeleteIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteResourceServer(ctx context.Context, input *cognitoidentityprovider.DeleteResourceServerInput) (*cognitoidentityprovider.DeleteResourceServerOutput, error) {
    return a.client.DeleteResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUser(ctx context.Context, input *cognitoidentityprovider.DeleteUserInput) (*cognitoidentityprovider.DeleteUserOutput, error) {
    return a.client.DeleteUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserAttributes(ctx context.Context, input *cognitoidentityprovider.DeleteUserAttributesInput) (*cognitoidentityprovider.DeleteUserAttributesOutput, error) {
    return a.client.DeleteUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPool(ctx context.Context, input *cognitoidentityprovider.DeleteUserPoolInput) (*cognitoidentityprovider.DeleteUserPoolOutput, error) {
    return a.client.DeleteUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPoolClient(ctx context.Context, input *cognitoidentityprovider.DeleteUserPoolClientInput) (*cognitoidentityprovider.DeleteUserPoolClientOutput, error) {
    return a.client.DeleteUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DeleteUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.DeleteUserPoolDomainInput) (*cognitoidentityprovider.DeleteUserPoolDomainOutput, error) {
    return a.client.DeleteUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeIdentityProvider(ctx context.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
    return a.client.DescribeIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeResourceServer(ctx context.Context, input *cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
    return a.client.DescribeResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeRiskConfiguration(ctx context.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
    return a.client.DescribeRiskConfigurationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserImportJob(ctx context.Context, input *cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
    return a.client.DescribeUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPool(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
    return a.client.DescribeUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPoolClient(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
    return a.client.DescribeUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) DescribeUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
    return a.client.DescribeUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ForgetDevice(ctx context.Context, input *cognitoidentityprovider.ForgetDeviceInput) (*cognitoidentityprovider.ForgetDeviceOutput, error) {
    return a.client.ForgetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ForgotPassword(ctx context.Context, input *cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
    return a.client.ForgotPasswordWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetCSVHeader(ctx context.Context, input *cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
    return a.client.GetCSVHeaderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetDevice(ctx context.Context, input *cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error) {
    return a.client.GetDeviceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetGroup(ctx context.Context, input *cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error) {
    return a.client.GetGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetIdentityProviderByIdentifier(ctx context.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
    return a.client.GetIdentityProviderByIdentifierWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetSigningCertificate(ctx context.Context, input *cognitoidentityprovider.GetSigningCertificateInput) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
    return a.client.GetSigningCertificateWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUICustomization(ctx context.Context, input *cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
    return a.client.GetUICustomizationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUser(ctx context.Context, input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
    return a.client.GetUserWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUserAttributeVerificationCode(ctx context.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
    return a.client.GetUserAttributeVerificationCodeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GetUserPoolMfaConfig(ctx context.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
    return a.client.GetUserPoolMfaConfigWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) GlobalSignOut(ctx context.Context, input *cognitoidentityprovider.GlobalSignOutInput) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
    return a.client.GlobalSignOutWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) InitiateAuth(ctx context.Context, input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
    return a.client.InitiateAuthWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListDevices(ctx context.Context, input *cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error) {
    return a.client.ListDevicesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListGroups(ctx context.Context, input *cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error) {
    return a.client.ListGroupsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListIdentityProviders(ctx context.Context, input *cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
    return a.client.ListIdentityProvidersWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListResourceServers(ctx context.Context, input *cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error) {
    return a.client.ListResourceServersWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListTagsForResource(ctx context.Context, input *cognitoidentityprovider.ListTagsForResourceInput) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUserImportJobs(ctx context.Context, input *cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
    return a.client.ListUserImportJobsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUserPoolClients(ctx context.Context, input *cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
    return a.client.ListUserPoolClientsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUserPools(ctx context.Context, input *cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
    return a.client.ListUserPoolsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUsers(ctx context.Context, input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error) {
    return a.client.ListUsersWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ListUsersInGroup(ctx context.Context, input *cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
    return a.client.ListUsersInGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) ResendConfirmationCode(ctx context.Context, input *cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error) {
    return a.client.ResendConfirmationCodeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) RespondToAuthChallenge(ctx context.Context, input *cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
    return a.client.RespondToAuthChallengeWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetRiskConfiguration(ctx context.Context, input *cognitoidentityprovider.SetRiskConfigurationInput) (*cognitoidentityprovider.SetRiskConfigurationOutput, error) {
    return a.client.SetRiskConfigurationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUICustomization(ctx context.Context, input *cognitoidentityprovider.SetUICustomizationInput) (*cognitoidentityprovider.SetUICustomizationOutput, error) {
    return a.client.SetUICustomizationWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUserMFAPreference(ctx context.Context, input *cognitoidentityprovider.SetUserMFAPreferenceInput) (*cognitoidentityprovider.SetUserMFAPreferenceOutput, error) {
    return a.client.SetUserMFAPreferenceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUserPoolMfaConfig(ctx context.Context, input *cognitoidentityprovider.SetUserPoolMfaConfigInput) (*cognitoidentityprovider.SetUserPoolMfaConfigOutput, error) {
    return a.client.SetUserPoolMfaConfigWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SetUserSettings(ctx context.Context, input *cognitoidentityprovider.SetUserSettingsInput) (*cognitoidentityprovider.SetUserSettingsOutput, error) {
    return a.client.SetUserSettingsWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) SignUp(ctx context.Context, input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error) {
    return a.client.SignUpWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) StartUserImportJob(ctx context.Context, input *cognitoidentityprovider.StartUserImportJobInput) (*cognitoidentityprovider.StartUserImportJobOutput, error) {
    return a.client.StartUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) StopUserImportJob(ctx context.Context, input *cognitoidentityprovider.StopUserImportJobInput) (*cognitoidentityprovider.StopUserImportJobOutput, error) {
    return a.client.StopUserImportJobWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) TagResource(ctx context.Context, input *cognitoidentityprovider.TagResourceInput) (*cognitoidentityprovider.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UntagResource(ctx context.Context, input *cognitoidentityprovider.UntagResourceInput) (*cognitoidentityprovider.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateAuthEventFeedback(ctx context.Context, input *cognitoidentityprovider.UpdateAuthEventFeedbackInput) (*cognitoidentityprovider.UpdateAuthEventFeedbackOutput, error) {
    return a.client.UpdateAuthEventFeedbackWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateDeviceStatus(ctx context.Context, input *cognitoidentityprovider.UpdateDeviceStatusInput) (*cognitoidentityprovider.UpdateDeviceStatusOutput, error) {
    return a.client.UpdateDeviceStatusWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateGroup(ctx context.Context, input *cognitoidentityprovider.UpdateGroupInput) (*cognitoidentityprovider.UpdateGroupOutput, error) {
    return a.client.UpdateGroupWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateIdentityProvider(ctx context.Context, input *cognitoidentityprovider.UpdateIdentityProviderInput) (*cognitoidentityprovider.UpdateIdentityProviderOutput, error) {
    return a.client.UpdateIdentityProviderWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateResourceServer(ctx context.Context, input *cognitoidentityprovider.UpdateResourceServerInput) (*cognitoidentityprovider.UpdateResourceServerOutput, error) {
    return a.client.UpdateResourceServerWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserAttributes(ctx context.Context, input *cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error) {
    return a.client.UpdateUserAttributesWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPool(ctx context.Context, input *cognitoidentityprovider.UpdateUserPoolInput) (*cognitoidentityprovider.UpdateUserPoolOutput, error) {
    return a.client.UpdateUserPoolWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPoolClient(ctx context.Context, input *cognitoidentityprovider.UpdateUserPoolClientInput) (*cognitoidentityprovider.UpdateUserPoolClientOutput, error) {
    return a.client.UpdateUserPoolClientWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) UpdateUserPoolDomain(ctx context.Context, input *cognitoidentityprovider.UpdateUserPoolDomainInput) (*cognitoidentityprovider.UpdateUserPoolDomainOutput, error) {
    return a.client.UpdateUserPoolDomainWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) VerifySoftwareToken(ctx context.Context, input *cognitoidentityprovider.VerifySoftwareTokenInput) (*cognitoidentityprovider.VerifySoftwareTokenOutput, error) {
    return a.client.VerifySoftwareTokenWithContext(ctx, input)
}

func (a *CognitoIdentityProviderActivities) VerifyUserAttribute(ctx context.Context, input *cognitoidentityprovider.VerifyUserAttributeInput) (*cognitoidentityprovider.VerifyUserAttributeOutput, error) {
    return a.client.VerifyUserAttributeWithContext(ctx, input)
}
