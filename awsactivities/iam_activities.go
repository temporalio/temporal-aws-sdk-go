
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

type IAMActivities struct {
	client iamiface.IAMAPI
}

func NewIAMActivities(session *session.Session, config... *aws.Config) *IAMActivities {
    client := iam.New(session, config...)
    return &IAMActivities{client: client}
}

func (a *IAMActivities) AddClientIDToOpenIDConnectProvider(input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
    return a.client.AddClientIDToOpenIDConnectProvider(input)
}

func (a *IAMActivities) AddRoleToInstanceProfile(input *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error) {
    return a.client.AddRoleToInstanceProfile(input)
}

func (a *IAMActivities) AddUserToGroup(input *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error) {
    return a.client.AddUserToGroup(input)
}

func (a *IAMActivities) AttachGroupPolicy(input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error) {
    return a.client.AttachGroupPolicy(input)
}

func (a *IAMActivities) AttachRolePolicy(input *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error) {
    return a.client.AttachRolePolicy(input)
}

func (a *IAMActivities) AttachUserPolicy(input *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error) {
    return a.client.AttachUserPolicy(input)
}

func (a *IAMActivities) ChangePassword(input *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error) {
    return a.client.ChangePassword(input)
}

func (a *IAMActivities) CreateAccessKey(input *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
    return a.client.CreateAccessKey(input)
}

func (a *IAMActivities) CreateAccountAlias(input *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error) {
    return a.client.CreateAccountAlias(input)
}

func (a *IAMActivities) CreateGroup(input *iam.CreateGroupInput) (*iam.CreateGroupOutput, error) {
    return a.client.CreateGroup(input)
}

func (a *IAMActivities) CreateInstanceProfile(input *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error) {
    return a.client.CreateInstanceProfile(input)
}

func (a *IAMActivities) CreateLoginProfile(input *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error) {
    return a.client.CreateLoginProfile(input)
}

func (a *IAMActivities) CreateOpenIDConnectProvider(input *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error) {
    return a.client.CreateOpenIDConnectProvider(input)
}

func (a *IAMActivities) CreatePolicy(input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
    return a.client.CreatePolicy(input)
}

func (a *IAMActivities) CreatePolicyVersion(input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error) {
    return a.client.CreatePolicyVersion(input)
}

func (a *IAMActivities) CreateRole(input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {
    return a.client.CreateRole(input)
}

func (a *IAMActivities) CreateSAMLProvider(input *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error) {
    return a.client.CreateSAMLProvider(input)
}

func (a *IAMActivities) CreateServiceLinkedRole(input *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error) {
    return a.client.CreateServiceLinkedRole(input)
}

func (a *IAMActivities) CreateServiceSpecificCredential(input *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error) {
    return a.client.CreateServiceSpecificCredential(input)
}

func (a *IAMActivities) CreateUser(input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *IAMActivities) CreateVirtualMFADevice(input *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error) {
    return a.client.CreateVirtualMFADevice(input)
}

func (a *IAMActivities) DeactivateMFADevice(input *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error) {
    return a.client.DeactivateMFADevice(input)
}

func (a *IAMActivities) DeleteAccessKey(input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
    return a.client.DeleteAccessKey(input)
}

func (a *IAMActivities) DeleteAccountAlias(input *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error) {
    return a.client.DeleteAccountAlias(input)
}

func (a *IAMActivities) DeleteAccountPasswordPolicy(input *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error) {
    return a.client.DeleteAccountPasswordPolicy(input)
}

func (a *IAMActivities) DeleteGroup(input *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error) {
    return a.client.DeleteGroup(input)
}

func (a *IAMActivities) DeleteGroupPolicy(input *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error) {
    return a.client.DeleteGroupPolicy(input)
}

func (a *IAMActivities) DeleteInstanceProfile(input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {
    return a.client.DeleteInstanceProfile(input)
}

func (a *IAMActivities) DeleteLoginProfile(input *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error) {
    return a.client.DeleteLoginProfile(input)
}

func (a *IAMActivities) DeleteOpenIDConnectProvider(input *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error) {
    return a.client.DeleteOpenIDConnectProvider(input)
}

func (a *IAMActivities) DeletePolicy(input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
    return a.client.DeletePolicy(input)
}

func (a *IAMActivities) DeletePolicyVersion(input *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error) {
    return a.client.DeletePolicyVersion(input)
}

func (a *IAMActivities) DeleteRole(input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
    return a.client.DeleteRole(input)
}

func (a *IAMActivities) DeleteRolePermissionsBoundary(input *iam.DeleteRolePermissionsBoundaryInput) (*iam.DeleteRolePermissionsBoundaryOutput, error) {
    return a.client.DeleteRolePermissionsBoundary(input)
}

func (a *IAMActivities) DeleteRolePolicy(input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {
    return a.client.DeleteRolePolicy(input)
}

func (a *IAMActivities) DeleteSAMLProvider(input *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error) {
    return a.client.DeleteSAMLProvider(input)
}

func (a *IAMActivities) DeleteSSHPublicKey(input *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error) {
    return a.client.DeleteSSHPublicKey(input)
}

func (a *IAMActivities) DeleteServerCertificate(input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {
    return a.client.DeleteServerCertificate(input)
}

func (a *IAMActivities) DeleteServiceLinkedRole(input *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error) {
    return a.client.DeleteServiceLinkedRole(input)
}

func (a *IAMActivities) DeleteServiceSpecificCredential(input *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error) {
    return a.client.DeleteServiceSpecificCredential(input)
}

func (a *IAMActivities) DeleteSigningCertificate(input *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error) {
    return a.client.DeleteSigningCertificate(input)
}

func (a *IAMActivities) DeleteUser(input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *IAMActivities) DeleteUserPermissionsBoundary(input *iam.DeleteUserPermissionsBoundaryInput) (*iam.DeleteUserPermissionsBoundaryOutput, error) {
    return a.client.DeleteUserPermissionsBoundary(input)
}

func (a *IAMActivities) DeleteUserPolicy(input *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
    return a.client.DeleteUserPolicy(input)
}

func (a *IAMActivities) DeleteVirtualMFADevice(input *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error) {
    return a.client.DeleteVirtualMFADevice(input)
}

func (a *IAMActivities) DetachGroupPolicy(input *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error) {
    return a.client.DetachGroupPolicy(input)
}

func (a *IAMActivities) DetachRolePolicy(input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
    return a.client.DetachRolePolicy(input)
}

func (a *IAMActivities) DetachUserPolicy(input *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error) {
    return a.client.DetachUserPolicy(input)
}

func (a *IAMActivities) EnableMFADevice(input *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error) {
    return a.client.EnableMFADevice(input)
}

func (a *IAMActivities) GenerateCredentialReport(input *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error) {
    return a.client.GenerateCredentialReport(input)
}

func (a *IAMActivities) GenerateOrganizationsAccessReport(input *iam.GenerateOrganizationsAccessReportInput) (*iam.GenerateOrganizationsAccessReportOutput, error) {
    return a.client.GenerateOrganizationsAccessReport(input)
}

func (a *IAMActivities) GenerateServiceLastAccessedDetails(input *iam.GenerateServiceLastAccessedDetailsInput) (*iam.GenerateServiceLastAccessedDetailsOutput, error) {
    return a.client.GenerateServiceLastAccessedDetails(input)
}

func (a *IAMActivities) GetAccessKeyLastUsed(input *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error) {
    return a.client.GetAccessKeyLastUsed(input)
}

func (a *IAMActivities) GetAccountAuthorizationDetails(input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error) {
    return a.client.GetAccountAuthorizationDetails(input)
}

func (a *IAMActivities) GetAccountPasswordPolicy(input *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error) {
    return a.client.GetAccountPasswordPolicy(input)
}

func (a *IAMActivities) GetAccountSummary(input *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error) {
    return a.client.GetAccountSummary(input)
}

func (a *IAMActivities) GetContextKeysForCustomPolicy(input *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
    return a.client.GetContextKeysForCustomPolicy(input)
}

func (a *IAMActivities) GetContextKeysForPrincipalPolicy(input *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
    return a.client.GetContextKeysForPrincipalPolicy(input)
}

func (a *IAMActivities) GetCredentialReport(input *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error) {
    return a.client.GetCredentialReport(input)
}

func (a *IAMActivities) GetGroup(input *iam.GetGroupInput) (*iam.GetGroupOutput, error) {
    return a.client.GetGroup(input)
}

func (a *IAMActivities) GetGroupPolicy(input *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error) {
    return a.client.GetGroupPolicy(input)
}

func (a *IAMActivities) GetInstanceProfile(input *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error) {
    return a.client.GetInstanceProfile(input)
}

func (a *IAMActivities) GetLoginProfile(input *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error) {
    return a.client.GetLoginProfile(input)
}

func (a *IAMActivities) GetOpenIDConnectProvider(input *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error) {
    return a.client.GetOpenIDConnectProvider(input)
}

func (a *IAMActivities) GetOrganizationsAccessReport(input *iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error) {
    return a.client.GetOrganizationsAccessReport(input)
}

func (a *IAMActivities) GetPolicy(input *iam.GetPolicyInput) (*iam.GetPolicyOutput, error) {
    return a.client.GetPolicy(input)
}

func (a *IAMActivities) GetPolicyVersion(input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error) {
    return a.client.GetPolicyVersion(input)
}

func (a *IAMActivities) GetRole(input *iam.GetRoleInput) (*iam.GetRoleOutput, error) {
    return a.client.GetRole(input)
}

func (a *IAMActivities) GetRolePolicy(input *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error) {
    return a.client.GetRolePolicy(input)
}

func (a *IAMActivities) GetSAMLProvider(input *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error) {
    return a.client.GetSAMLProvider(input)
}

func (a *IAMActivities) GetSSHPublicKey(input *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error) {
    return a.client.GetSSHPublicKey(input)
}

func (a *IAMActivities) GetServerCertificate(input *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error) {
    return a.client.GetServerCertificate(input)
}

func (a *IAMActivities) GetServiceLastAccessedDetails(input *iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error) {
    return a.client.GetServiceLastAccessedDetails(input)
}

func (a *IAMActivities) GetServiceLastAccessedDetailsWithEntities(input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error) {
    return a.client.GetServiceLastAccessedDetailsWithEntities(input)
}

func (a *IAMActivities) GetServiceLinkedRoleDeletionStatus(input *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
    return a.client.GetServiceLinkedRoleDeletionStatus(input)
}

func (a *IAMActivities) GetUser(input *iam.GetUserInput) (*iam.GetUserOutput, error) {
    return a.client.GetUser(input)
}

func (a *IAMActivities) GetUserPolicy(input *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error) {
    return a.client.GetUserPolicy(input)
}

func (a *IAMActivities) ListAccessKeys(input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
    return a.client.ListAccessKeys(input)
}

func (a *IAMActivities) ListAccountAliases(input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error) {
    return a.client.ListAccountAliases(input)
}

func (a *IAMActivities) ListAttachedGroupPolicies(input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error) {
    return a.client.ListAttachedGroupPolicies(input)
}

func (a *IAMActivities) ListAttachedRolePolicies(input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
    return a.client.ListAttachedRolePolicies(input)
}

func (a *IAMActivities) ListAttachedUserPolicies(input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {
    return a.client.ListAttachedUserPolicies(input)
}

func (a *IAMActivities) ListEntitiesForPolicy(input *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error) {
    return a.client.ListEntitiesForPolicy(input)
}

func (a *IAMActivities) ListGroupPolicies(input *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error) {
    return a.client.ListGroupPolicies(input)
}

func (a *IAMActivities) ListGroups(input *iam.ListGroupsInput) (*iam.ListGroupsOutput, error) {
    return a.client.ListGroups(input)
}

func (a *IAMActivities) ListGroupsForUser(input *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error) {
    return a.client.ListGroupsForUser(input)
}

func (a *IAMActivities) ListInstanceProfiles(input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {
    return a.client.ListInstanceProfiles(input)
}

func (a *IAMActivities) ListInstanceProfilesForRole(input *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error) {
    return a.client.ListInstanceProfilesForRole(input)
}

func (a *IAMActivities) ListMFADevices(input *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error) {
    return a.client.ListMFADevices(input)
}

func (a *IAMActivities) ListOpenIDConnectProviders(input *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error) {
    return a.client.ListOpenIDConnectProviders(input)
}

func (a *IAMActivities) ListPolicies(input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
    return a.client.ListPolicies(input)
}

func (a *IAMActivities) ListPoliciesGrantingServiceAccess(input *iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error) {
    return a.client.ListPoliciesGrantingServiceAccess(input)
}

func (a *IAMActivities) ListPolicyVersions(input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error) {
    return a.client.ListPolicyVersions(input)
}

func (a *IAMActivities) ListRolePolicies(input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
    return a.client.ListRolePolicies(input)
}

func (a *IAMActivities) ListRoleTags(input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error) {
    return a.client.ListRoleTags(input)
}

func (a *IAMActivities) ListRoles(input *iam.ListRolesInput) (*iam.ListRolesOutput, error) {
    return a.client.ListRoles(input)
}

func (a *IAMActivities) ListSAMLProviders(input *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error) {
    return a.client.ListSAMLProviders(input)
}

func (a *IAMActivities) ListSSHPublicKeys(input *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error) {
    return a.client.ListSSHPublicKeys(input)
}

func (a *IAMActivities) ListServerCertificates(input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
    return a.client.ListServerCertificates(input)
}

func (a *IAMActivities) ListServiceSpecificCredentials(input *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error) {
    return a.client.ListServiceSpecificCredentials(input)
}

func (a *IAMActivities) ListSigningCertificates(input *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error) {
    return a.client.ListSigningCertificates(input)
}

func (a *IAMActivities) ListUserPolicies(input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
    return a.client.ListUserPolicies(input)
}

func (a *IAMActivities) ListUserTags(input *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error) {
    return a.client.ListUserTags(input)
}

func (a *IAMActivities) ListUsers(input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *IAMActivities) ListVirtualMFADevices(input *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error) {
    return a.client.ListVirtualMFADevices(input)
}

func (a *IAMActivities) PutGroupPolicy(input *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error) {
    return a.client.PutGroupPolicy(input)
}

func (a *IAMActivities) PutRolePermissionsBoundary(input *iam.PutRolePermissionsBoundaryInput) (*iam.PutRolePermissionsBoundaryOutput, error) {
    return a.client.PutRolePermissionsBoundary(input)
}

func (a *IAMActivities) PutRolePolicy(input *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error) {
    return a.client.PutRolePolicy(input)
}

func (a *IAMActivities) PutUserPermissionsBoundary(input *iam.PutUserPermissionsBoundaryInput) (*iam.PutUserPermissionsBoundaryOutput, error) {
    return a.client.PutUserPermissionsBoundary(input)
}

func (a *IAMActivities) PutUserPolicy(input *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {
    return a.client.PutUserPolicy(input)
}

func (a *IAMActivities) RemoveClientIDFromOpenIDConnectProvider(input *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {
    return a.client.RemoveClientIDFromOpenIDConnectProvider(input)
}

func (a *IAMActivities) RemoveRoleFromInstanceProfile(input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
    return a.client.RemoveRoleFromInstanceProfile(input)
}

func (a *IAMActivities) RemoveUserFromGroup(input *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error) {
    return a.client.RemoveUserFromGroup(input)
}

func (a *IAMActivities) ResetServiceSpecificCredential(input *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error) {
    return a.client.ResetServiceSpecificCredential(input)
}

func (a *IAMActivities) ResyncMFADevice(input *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error) {
    return a.client.ResyncMFADevice(input)
}

func (a *IAMActivities) SetDefaultPolicyVersion(input *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error) {
    return a.client.SetDefaultPolicyVersion(input)
}

func (a *IAMActivities) SetSecurityTokenServicePreferences(input *iam.SetSecurityTokenServicePreferencesInput) (*iam.SetSecurityTokenServicePreferencesOutput, error) {
    return a.client.SetSecurityTokenServicePreferences(input)
}

func (a *IAMActivities) SimulateCustomPolicy(input *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error) {
    return a.client.SimulateCustomPolicy(input)
}

func (a *IAMActivities) SimulatePrincipalPolicy(input *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error) {
    return a.client.SimulatePrincipalPolicy(input)
}

func (a *IAMActivities) TagRole(input *iam.TagRoleInput) (*iam.TagRoleOutput, error) {
    return a.client.TagRole(input)
}

func (a *IAMActivities) TagUser(input *iam.TagUserInput) (*iam.TagUserOutput, error) {
    return a.client.TagUser(input)
}

func (a *IAMActivities) UntagRole(input *iam.UntagRoleInput) (*iam.UntagRoleOutput, error) {
    return a.client.UntagRole(input)
}

func (a *IAMActivities) UntagUser(input *iam.UntagUserInput) (*iam.UntagUserOutput, error) {
    return a.client.UntagUser(input)
}

func (a *IAMActivities) UpdateAccessKey(input *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error) {
    return a.client.UpdateAccessKey(input)
}

func (a *IAMActivities) UpdateAccountPasswordPolicy(input *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error) {
    return a.client.UpdateAccountPasswordPolicy(input)
}

func (a *IAMActivities) UpdateAssumeRolePolicy(input *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error) {
    return a.client.UpdateAssumeRolePolicy(input)
}

func (a *IAMActivities) UpdateGroup(input *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error) {
    return a.client.UpdateGroup(input)
}

func (a *IAMActivities) UpdateLoginProfile(input *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error) {
    return a.client.UpdateLoginProfile(input)
}

func (a *IAMActivities) UpdateOpenIDConnectProviderThumbprint(input *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {
    return a.client.UpdateOpenIDConnectProviderThumbprint(input)
}

func (a *IAMActivities) UpdateRole(input *iam.UpdateRoleInput) (*iam.UpdateRoleOutput, error) {
    return a.client.UpdateRole(input)
}

func (a *IAMActivities) UpdateRoleDescription(input *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error) {
    return a.client.UpdateRoleDescription(input)
}

func (a *IAMActivities) UpdateSAMLProvider(input *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error) {
    return a.client.UpdateSAMLProvider(input)
}

func (a *IAMActivities) UpdateSSHPublicKey(input *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error) {
    return a.client.UpdateSSHPublicKey(input)
}

func (a *IAMActivities) UpdateServerCertificate(input *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error) {
    return a.client.UpdateServerCertificate(input)
}

func (a *IAMActivities) UpdateServiceSpecificCredential(input *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error) {
    return a.client.UpdateServiceSpecificCredential(input)
}

func (a *IAMActivities) UpdateSigningCertificate(input *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error) {
    return a.client.UpdateSigningCertificate(input)
}

func (a *IAMActivities) UpdateUser(input *iam.UpdateUserInput) (*iam.UpdateUserOutput, error) {
    return a.client.UpdateUser(input)
}

func (a *IAMActivities) UploadSSHPublicKey(input *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error) {
    return a.client.UploadSSHPublicKey(input)
}

func (a *IAMActivities) UploadServerCertificate(input *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error) {
    return a.client.UploadServerCertificate(input)
}

func (a *IAMActivities) UploadSigningCertificate(input *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error) {
    return a.client.UploadSigningCertificate(input)
}

func (a *IAMActivities) WaitUntilInstanceProfileExists(input *iam.GetInstanceProfileInput) error {
    return a.client.WaitUntilInstanceProfileExists(input)
}

func (a *IAMActivities) WaitUntilPolicyExists(input *iam.GetPolicyInput) error {
    return a.client.WaitUntilPolicyExists(input)
}

func (a *IAMActivities) WaitUntilRoleExists(input *iam.GetRoleInput) error {
    return a.client.WaitUntilRoleExists(input)
}

func (a *IAMActivities) WaitUntilUserExists(input *iam.GetUserInput) error {
    return a.client.WaitUntilUserExists(input)
}
