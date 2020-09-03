package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/iam"
	"go.temporal.io/sdk/workflow"
)

type IAMClient interface {
    AddClientIDToOpenIDConnectProvider(ctx workflow.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)
    AddClientIDToOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) *IamAddClientIDToOpenIDConnectProviderResult

    AddRoleToInstanceProfile(ctx workflow.Context, input *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error)
    AddRoleToInstanceProfileAsync(ctx workflow.Context, input *iam.AddRoleToInstanceProfileInput) *IamAddRoleToInstanceProfileResult

    AddUserToGroup(ctx workflow.Context, input *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error)
    AddUserToGroupAsync(ctx workflow.Context, input *iam.AddUserToGroupInput) *IamAddUserToGroupResult

    AttachGroupPolicy(ctx workflow.Context, input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error)
    AttachGroupPolicyAsync(ctx workflow.Context, input *iam.AttachGroupPolicyInput) *IamAttachGroupPolicyResult

    AttachRolePolicy(ctx workflow.Context, input *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error)
    AttachRolePolicyAsync(ctx workflow.Context, input *iam.AttachRolePolicyInput) *IamAttachRolePolicyResult

    AttachUserPolicy(ctx workflow.Context, input *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error)
    AttachUserPolicyAsync(ctx workflow.Context, input *iam.AttachUserPolicyInput) *IamAttachUserPolicyResult

    ChangePassword(ctx workflow.Context, input *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error)
    ChangePasswordAsync(ctx workflow.Context, input *iam.ChangePasswordInput) *IamChangePasswordResult

    CreateAccessKey(ctx workflow.Context, input *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error)
    CreateAccessKeyAsync(ctx workflow.Context, input *iam.CreateAccessKeyInput) *IamCreateAccessKeyResult

    CreateAccountAlias(ctx workflow.Context, input *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error)
    CreateAccountAliasAsync(ctx workflow.Context, input *iam.CreateAccountAliasInput) *IamCreateAccountAliasResult

    CreateGroup(ctx workflow.Context, input *iam.CreateGroupInput) (*iam.CreateGroupOutput, error)
    CreateGroupAsync(ctx workflow.Context, input *iam.CreateGroupInput) *IamCreateGroupResult

    CreateInstanceProfile(ctx workflow.Context, input *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error)
    CreateInstanceProfileAsync(ctx workflow.Context, input *iam.CreateInstanceProfileInput) *IamCreateInstanceProfileResult

    CreateLoginProfile(ctx workflow.Context, input *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error)
    CreateLoginProfileAsync(ctx workflow.Context, input *iam.CreateLoginProfileInput) *IamCreateLoginProfileResult

    CreateOpenIDConnectProvider(ctx workflow.Context, input *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error)
    CreateOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.CreateOpenIDConnectProviderInput) *IamCreateOpenIDConnectProviderResult

    CreatePolicy(ctx workflow.Context, input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)
    CreatePolicyAsync(ctx workflow.Context, input *iam.CreatePolicyInput) *IamCreatePolicyResult

    CreatePolicyVersion(ctx workflow.Context, input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error)
    CreatePolicyVersionAsync(ctx workflow.Context, input *iam.CreatePolicyVersionInput) *IamCreatePolicyVersionResult

    CreateRole(ctx workflow.Context, input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error)
    CreateRoleAsync(ctx workflow.Context, input *iam.CreateRoleInput) *IamCreateRoleResult

    CreateSAMLProvider(ctx workflow.Context, input *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error)
    CreateSAMLProviderAsync(ctx workflow.Context, input *iam.CreateSAMLProviderInput) *IamCreateSAMLProviderResult

    CreateServiceLinkedRole(ctx workflow.Context, input *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error)
    CreateServiceLinkedRoleAsync(ctx workflow.Context, input *iam.CreateServiceLinkedRoleInput) *IamCreateServiceLinkedRoleResult

    CreateServiceSpecificCredential(ctx workflow.Context, input *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error)
    CreateServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.CreateServiceSpecificCredentialInput) *IamCreateServiceSpecificCredentialResult

    CreateUser(ctx workflow.Context, input *iam.CreateUserInput) (*iam.CreateUserOutput, error)
    CreateUserAsync(ctx workflow.Context, input *iam.CreateUserInput) *IamCreateUserResult

    CreateVirtualMFADevice(ctx workflow.Context, input *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error)
    CreateVirtualMFADeviceAsync(ctx workflow.Context, input *iam.CreateVirtualMFADeviceInput) *IamCreateVirtualMFADeviceResult

    DeactivateMFADevice(ctx workflow.Context, input *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error)
    DeactivateMFADeviceAsync(ctx workflow.Context, input *iam.DeactivateMFADeviceInput) *IamDeactivateMFADeviceResult

    DeleteAccessKey(ctx workflow.Context, input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error)
    DeleteAccessKeyAsync(ctx workflow.Context, input *iam.DeleteAccessKeyInput) *IamDeleteAccessKeyResult

    DeleteAccountAlias(ctx workflow.Context, input *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error)
    DeleteAccountAliasAsync(ctx workflow.Context, input *iam.DeleteAccountAliasInput) *IamDeleteAccountAliasResult

    DeleteAccountPasswordPolicy(ctx workflow.Context, input *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error)
    DeleteAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.DeleteAccountPasswordPolicyInput) *IamDeleteAccountPasswordPolicyResult

    DeleteGroup(ctx workflow.Context, input *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error)
    DeleteGroupAsync(ctx workflow.Context, input *iam.DeleteGroupInput) *IamDeleteGroupResult

    DeleteGroupPolicy(ctx workflow.Context, input *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error)
    DeleteGroupPolicyAsync(ctx workflow.Context, input *iam.DeleteGroupPolicyInput) *IamDeleteGroupPolicyResult

    DeleteInstanceProfile(ctx workflow.Context, input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)
    DeleteInstanceProfileAsync(ctx workflow.Context, input *iam.DeleteInstanceProfileInput) *IamDeleteInstanceProfileResult

    DeleteLoginProfile(ctx workflow.Context, input *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error)
    DeleteLoginProfileAsync(ctx workflow.Context, input *iam.DeleteLoginProfileInput) *IamDeleteLoginProfileResult

    DeleteOpenIDConnectProvider(ctx workflow.Context, input *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error)
    DeleteOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.DeleteOpenIDConnectProviderInput) *IamDeleteOpenIDConnectProviderResult

    DeletePolicy(ctx workflow.Context, input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error)
    DeletePolicyAsync(ctx workflow.Context, input *iam.DeletePolicyInput) *IamDeletePolicyResult

    DeletePolicyVersion(ctx workflow.Context, input *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error)
    DeletePolicyVersionAsync(ctx workflow.Context, input *iam.DeletePolicyVersionInput) *IamDeletePolicyVersionResult

    DeleteRole(ctx workflow.Context, input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)
    DeleteRoleAsync(ctx workflow.Context, input *iam.DeleteRoleInput) *IamDeleteRoleResult

    DeleteRolePermissionsBoundary(ctx workflow.Context, input *iam.DeleteRolePermissionsBoundaryInput) (*iam.DeleteRolePermissionsBoundaryOutput, error)
    DeleteRolePermissionsBoundaryAsync(ctx workflow.Context, input *iam.DeleteRolePermissionsBoundaryInput) *IamDeleteRolePermissionsBoundaryResult

    DeleteRolePolicy(ctx workflow.Context, input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error)
    DeleteRolePolicyAsync(ctx workflow.Context, input *iam.DeleteRolePolicyInput) *IamDeleteRolePolicyResult

    DeleteSAMLProvider(ctx workflow.Context, input *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error)
    DeleteSAMLProviderAsync(ctx workflow.Context, input *iam.DeleteSAMLProviderInput) *IamDeleteSAMLProviderResult

    DeleteSSHPublicKey(ctx workflow.Context, input *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error)
    DeleteSSHPublicKeyAsync(ctx workflow.Context, input *iam.DeleteSSHPublicKeyInput) *IamDeleteSSHPublicKeyResult

    DeleteServerCertificate(ctx workflow.Context, input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)
    DeleteServerCertificateAsync(ctx workflow.Context, input *iam.DeleteServerCertificateInput) *IamDeleteServerCertificateResult

    DeleteServiceLinkedRole(ctx workflow.Context, input *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error)
    DeleteServiceLinkedRoleAsync(ctx workflow.Context, input *iam.DeleteServiceLinkedRoleInput) *IamDeleteServiceLinkedRoleResult

    DeleteServiceSpecificCredential(ctx workflow.Context, input *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error)
    DeleteServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.DeleteServiceSpecificCredentialInput) *IamDeleteServiceSpecificCredentialResult

    DeleteSigningCertificate(ctx workflow.Context, input *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error)
    DeleteSigningCertificateAsync(ctx workflow.Context, input *iam.DeleteSigningCertificateInput) *IamDeleteSigningCertificateResult

    DeleteUser(ctx workflow.Context, input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error)
    DeleteUserAsync(ctx workflow.Context, input *iam.DeleteUserInput) *IamDeleteUserResult

    DeleteUserPermissionsBoundary(ctx workflow.Context, input *iam.DeleteUserPermissionsBoundaryInput) (*iam.DeleteUserPermissionsBoundaryOutput, error)
    DeleteUserPermissionsBoundaryAsync(ctx workflow.Context, input *iam.DeleteUserPermissionsBoundaryInput) *IamDeleteUserPermissionsBoundaryResult

    DeleteUserPolicy(ctx workflow.Context, input *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error)
    DeleteUserPolicyAsync(ctx workflow.Context, input *iam.DeleteUserPolicyInput) *IamDeleteUserPolicyResult

    DeleteVirtualMFADevice(ctx workflow.Context, input *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error)
    DeleteVirtualMFADeviceAsync(ctx workflow.Context, input *iam.DeleteVirtualMFADeviceInput) *IamDeleteVirtualMFADeviceResult

    DetachGroupPolicy(ctx workflow.Context, input *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error)
    DetachGroupPolicyAsync(ctx workflow.Context, input *iam.DetachGroupPolicyInput) *IamDetachGroupPolicyResult

    DetachRolePolicy(ctx workflow.Context, input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error)
    DetachRolePolicyAsync(ctx workflow.Context, input *iam.DetachRolePolicyInput) *IamDetachRolePolicyResult

    DetachUserPolicy(ctx workflow.Context, input *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error)
    DetachUserPolicyAsync(ctx workflow.Context, input *iam.DetachUserPolicyInput) *IamDetachUserPolicyResult

    EnableMFADevice(ctx workflow.Context, input *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error)
    EnableMFADeviceAsync(ctx workflow.Context, input *iam.EnableMFADeviceInput) *IamEnableMFADeviceResult

    GenerateCredentialReport(ctx workflow.Context, input *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error)
    GenerateCredentialReportAsync(ctx workflow.Context, input *iam.GenerateCredentialReportInput) *IamGenerateCredentialReportResult

    GenerateOrganizationsAccessReport(ctx workflow.Context, input *iam.GenerateOrganizationsAccessReportInput) (*iam.GenerateOrganizationsAccessReportOutput, error)
    GenerateOrganizationsAccessReportAsync(ctx workflow.Context, input *iam.GenerateOrganizationsAccessReportInput) *IamGenerateOrganizationsAccessReportResult

    GenerateServiceLastAccessedDetails(ctx workflow.Context, input *iam.GenerateServiceLastAccessedDetailsInput) (*iam.GenerateServiceLastAccessedDetailsOutput, error)
    GenerateServiceLastAccessedDetailsAsync(ctx workflow.Context, input *iam.GenerateServiceLastAccessedDetailsInput) *IamGenerateServiceLastAccessedDetailsResult

    GetAccessKeyLastUsed(ctx workflow.Context, input *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error)
    GetAccessKeyLastUsedAsync(ctx workflow.Context, input *iam.GetAccessKeyLastUsedInput) *IamGetAccessKeyLastUsedResult

    GetAccountAuthorizationDetails(ctx workflow.Context, input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error)
    GetAccountAuthorizationDetailsAsync(ctx workflow.Context, input *iam.GetAccountAuthorizationDetailsInput) *IamGetAccountAuthorizationDetailsResult

    GetAccountPasswordPolicy(ctx workflow.Context, input *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error)
    GetAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.GetAccountPasswordPolicyInput) *IamGetAccountPasswordPolicyResult

    GetAccountSummary(ctx workflow.Context, input *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error)
    GetAccountSummaryAsync(ctx workflow.Context, input *iam.GetAccountSummaryInput) *IamGetAccountSummaryResult

    GetContextKeysForCustomPolicy(ctx workflow.Context, input *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
    GetContextKeysForCustomPolicyAsync(ctx workflow.Context, input *iam.GetContextKeysForCustomPolicyInput) *IamGetContextKeysForCustomPolicyResult

    GetContextKeysForPrincipalPolicy(ctx workflow.Context, input *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
    GetContextKeysForPrincipalPolicyAsync(ctx workflow.Context, input *iam.GetContextKeysForPrincipalPolicyInput) *IamGetContextKeysForPrincipalPolicyResult

    GetCredentialReport(ctx workflow.Context, input *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error)
    GetCredentialReportAsync(ctx workflow.Context, input *iam.GetCredentialReportInput) *IamGetCredentialReportResult

    GetGroup(ctx workflow.Context, input *iam.GetGroupInput) (*iam.GetGroupOutput, error)
    GetGroupAsync(ctx workflow.Context, input *iam.GetGroupInput) *IamGetGroupResult

    GetGroupPolicy(ctx workflow.Context, input *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error)
    GetGroupPolicyAsync(ctx workflow.Context, input *iam.GetGroupPolicyInput) *IamGetGroupPolicyResult

    GetInstanceProfile(ctx workflow.Context, input *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error)
    GetInstanceProfileAsync(ctx workflow.Context, input *iam.GetInstanceProfileInput) *IamGetInstanceProfileResult

    GetLoginProfile(ctx workflow.Context, input *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error)
    GetLoginProfileAsync(ctx workflow.Context, input *iam.GetLoginProfileInput) *IamGetLoginProfileResult

    GetOpenIDConnectProvider(ctx workflow.Context, input *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error)
    GetOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.GetOpenIDConnectProviderInput) *IamGetOpenIDConnectProviderResult

    GetOrganizationsAccessReport(ctx workflow.Context, input *iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error)
    GetOrganizationsAccessReportAsync(ctx workflow.Context, input *iam.GetOrganizationsAccessReportInput) *IamGetOrganizationsAccessReportResult

    GetPolicy(ctx workflow.Context, input *iam.GetPolicyInput) (*iam.GetPolicyOutput, error)
    GetPolicyAsync(ctx workflow.Context, input *iam.GetPolicyInput) *IamGetPolicyResult

    GetPolicyVersion(ctx workflow.Context, input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error)
    GetPolicyVersionAsync(ctx workflow.Context, input *iam.GetPolicyVersionInput) *IamGetPolicyVersionResult

    GetRole(ctx workflow.Context, input *iam.GetRoleInput) (*iam.GetRoleOutput, error)
    GetRoleAsync(ctx workflow.Context, input *iam.GetRoleInput) *IamGetRoleResult

    GetRolePolicy(ctx workflow.Context, input *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error)
    GetRolePolicyAsync(ctx workflow.Context, input *iam.GetRolePolicyInput) *IamGetRolePolicyResult

    GetSAMLProvider(ctx workflow.Context, input *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error)
    GetSAMLProviderAsync(ctx workflow.Context, input *iam.GetSAMLProviderInput) *IamGetSAMLProviderResult

    GetSSHPublicKey(ctx workflow.Context, input *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error)
    GetSSHPublicKeyAsync(ctx workflow.Context, input *iam.GetSSHPublicKeyInput) *IamGetSSHPublicKeyResult

    GetServerCertificate(ctx workflow.Context, input *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error)
    GetServerCertificateAsync(ctx workflow.Context, input *iam.GetServerCertificateInput) *IamGetServerCertificateResult

    GetServiceLastAccessedDetails(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error)
    GetServiceLastAccessedDetailsAsync(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsInput) *IamGetServiceLastAccessedDetailsResult

    GetServiceLastAccessedDetailsWithEntities(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error)
    GetServiceLastAccessedDetailsWithEntitiesAsync(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) *IamGetServiceLastAccessedDetailsWithEntitiesResult

    GetServiceLinkedRoleDeletionStatus(ctx workflow.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
    GetServiceLinkedRoleDeletionStatusAsync(ctx workflow.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) *IamGetServiceLinkedRoleDeletionStatusResult

    GetUser(ctx workflow.Context, input *iam.GetUserInput) (*iam.GetUserOutput, error)
    GetUserAsync(ctx workflow.Context, input *iam.GetUserInput) *IamGetUserResult

    GetUserPolicy(ctx workflow.Context, input *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error)
    GetUserPolicyAsync(ctx workflow.Context, input *iam.GetUserPolicyInput) *IamGetUserPolicyResult

    ListAccessKeys(ctx workflow.Context, input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)
    ListAccessKeysAsync(ctx workflow.Context, input *iam.ListAccessKeysInput) *IamListAccessKeysResult

    ListAccountAliases(ctx workflow.Context, input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error)
    ListAccountAliasesAsync(ctx workflow.Context, input *iam.ListAccountAliasesInput) *IamListAccountAliasesResult

    ListAttachedGroupPolicies(ctx workflow.Context, input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error)
    ListAttachedGroupPoliciesAsync(ctx workflow.Context, input *iam.ListAttachedGroupPoliciesInput) *IamListAttachedGroupPoliciesResult

    ListAttachedRolePolicies(ctx workflow.Context, input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error)
    ListAttachedRolePoliciesAsync(ctx workflow.Context, input *iam.ListAttachedRolePoliciesInput) *IamListAttachedRolePoliciesResult

    ListAttachedUserPolicies(ctx workflow.Context, input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error)
    ListAttachedUserPoliciesAsync(ctx workflow.Context, input *iam.ListAttachedUserPoliciesInput) *IamListAttachedUserPoliciesResult

    ListEntitiesForPolicy(ctx workflow.Context, input *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error)
    ListEntitiesForPolicyAsync(ctx workflow.Context, input *iam.ListEntitiesForPolicyInput) *IamListEntitiesForPolicyResult

    ListGroupPolicies(ctx workflow.Context, input *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error)
    ListGroupPoliciesAsync(ctx workflow.Context, input *iam.ListGroupPoliciesInput) *IamListGroupPoliciesResult

    ListGroups(ctx workflow.Context, input *iam.ListGroupsInput) (*iam.ListGroupsOutput, error)
    ListGroupsAsync(ctx workflow.Context, input *iam.ListGroupsInput) *IamListGroupsResult

    ListGroupsForUser(ctx workflow.Context, input *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error)
    ListGroupsForUserAsync(ctx workflow.Context, input *iam.ListGroupsForUserInput) *IamListGroupsForUserResult

    ListInstanceProfiles(ctx workflow.Context, input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)
    ListInstanceProfilesAsync(ctx workflow.Context, input *iam.ListInstanceProfilesInput) *IamListInstanceProfilesResult

    ListInstanceProfilesForRole(ctx workflow.Context, input *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error)
    ListInstanceProfilesForRoleAsync(ctx workflow.Context, input *iam.ListInstanceProfilesForRoleInput) *IamListInstanceProfilesForRoleResult

    ListMFADevices(ctx workflow.Context, input *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error)
    ListMFADevicesAsync(ctx workflow.Context, input *iam.ListMFADevicesInput) *IamListMFADevicesResult

    ListOpenIDConnectProviders(ctx workflow.Context, input *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error)
    ListOpenIDConnectProvidersAsync(ctx workflow.Context, input *iam.ListOpenIDConnectProvidersInput) *IamListOpenIDConnectProvidersResult

    ListPolicies(ctx workflow.Context, input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error)
    ListPoliciesAsync(ctx workflow.Context, input *iam.ListPoliciesInput) *IamListPoliciesResult

    ListPoliciesGrantingServiceAccess(ctx workflow.Context, input *iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error)
    ListPoliciesGrantingServiceAccessAsync(ctx workflow.Context, input *iam.ListPoliciesGrantingServiceAccessInput) *IamListPoliciesGrantingServiceAccessResult

    ListPolicyVersions(ctx workflow.Context, input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error)
    ListPolicyVersionsAsync(ctx workflow.Context, input *iam.ListPolicyVersionsInput) *IamListPolicyVersionsResult

    ListRolePolicies(ctx workflow.Context, input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error)
    ListRolePoliciesAsync(ctx workflow.Context, input *iam.ListRolePoliciesInput) *IamListRolePoliciesResult

    ListRoleTags(ctx workflow.Context, input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error)
    ListRoleTagsAsync(ctx workflow.Context, input *iam.ListRoleTagsInput) *IamListRoleTagsResult

    ListRoles(ctx workflow.Context, input *iam.ListRolesInput) (*iam.ListRolesOutput, error)
    ListRolesAsync(ctx workflow.Context, input *iam.ListRolesInput) *IamListRolesResult

    ListSAMLProviders(ctx workflow.Context, input *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error)
    ListSAMLProvidersAsync(ctx workflow.Context, input *iam.ListSAMLProvidersInput) *IamListSAMLProvidersResult

    ListSSHPublicKeys(ctx workflow.Context, input *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error)
    ListSSHPublicKeysAsync(ctx workflow.Context, input *iam.ListSSHPublicKeysInput) *IamListSSHPublicKeysResult

    ListServerCertificates(ctx workflow.Context, input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)
    ListServerCertificatesAsync(ctx workflow.Context, input *iam.ListServerCertificatesInput) *IamListServerCertificatesResult

    ListServiceSpecificCredentials(ctx workflow.Context, input *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error)
    ListServiceSpecificCredentialsAsync(ctx workflow.Context, input *iam.ListServiceSpecificCredentialsInput) *IamListServiceSpecificCredentialsResult

    ListSigningCertificates(ctx workflow.Context, input *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error)
    ListSigningCertificatesAsync(ctx workflow.Context, input *iam.ListSigningCertificatesInput) *IamListSigningCertificatesResult

    ListUserPolicies(ctx workflow.Context, input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error)
    ListUserPoliciesAsync(ctx workflow.Context, input *iam.ListUserPoliciesInput) *IamListUserPoliciesResult

    ListUserTags(ctx workflow.Context, input *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error)
    ListUserTagsAsync(ctx workflow.Context, input *iam.ListUserTagsInput) *IamListUserTagsResult

    ListUsers(ctx workflow.Context, input *iam.ListUsersInput) (*iam.ListUsersOutput, error)
    ListUsersAsync(ctx workflow.Context, input *iam.ListUsersInput) *IamListUsersResult

    ListVirtualMFADevices(ctx workflow.Context, input *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error)
    ListVirtualMFADevicesAsync(ctx workflow.Context, input *iam.ListVirtualMFADevicesInput) *IamListVirtualMFADevicesResult

    PutGroupPolicy(ctx workflow.Context, input *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error)
    PutGroupPolicyAsync(ctx workflow.Context, input *iam.PutGroupPolicyInput) *IamPutGroupPolicyResult

    PutRolePermissionsBoundary(ctx workflow.Context, input *iam.PutRolePermissionsBoundaryInput) (*iam.PutRolePermissionsBoundaryOutput, error)
    PutRolePermissionsBoundaryAsync(ctx workflow.Context, input *iam.PutRolePermissionsBoundaryInput) *IamPutRolePermissionsBoundaryResult

    PutRolePolicy(ctx workflow.Context, input *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error)
    PutRolePolicyAsync(ctx workflow.Context, input *iam.PutRolePolicyInput) *IamPutRolePolicyResult

    PutUserPermissionsBoundary(ctx workflow.Context, input *iam.PutUserPermissionsBoundaryInput) (*iam.PutUserPermissionsBoundaryOutput, error)
    PutUserPermissionsBoundaryAsync(ctx workflow.Context, input *iam.PutUserPermissionsBoundaryInput) *IamPutUserPermissionsBoundaryResult

    PutUserPolicy(ctx workflow.Context, input *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error)
    PutUserPolicyAsync(ctx workflow.Context, input *iam.PutUserPolicyInput) *IamPutUserPolicyResult

    RemoveClientIDFromOpenIDConnectProvider(ctx workflow.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)
    RemoveClientIDFromOpenIDConnectProviderAsync(ctx workflow.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) *IamRemoveClientIDFromOpenIDConnectProviderResult

    RemoveRoleFromInstanceProfile(ctx workflow.Context, input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error)
    RemoveRoleFromInstanceProfileAsync(ctx workflow.Context, input *iam.RemoveRoleFromInstanceProfileInput) *IamRemoveRoleFromInstanceProfileResult

    RemoveUserFromGroup(ctx workflow.Context, input *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error)
    RemoveUserFromGroupAsync(ctx workflow.Context, input *iam.RemoveUserFromGroupInput) *IamRemoveUserFromGroupResult

    ResetServiceSpecificCredential(ctx workflow.Context, input *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error)
    ResetServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.ResetServiceSpecificCredentialInput) *IamResetServiceSpecificCredentialResult

    ResyncMFADevice(ctx workflow.Context, input *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error)
    ResyncMFADeviceAsync(ctx workflow.Context, input *iam.ResyncMFADeviceInput) *IamResyncMFADeviceResult

    SetDefaultPolicyVersion(ctx workflow.Context, input *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error)
    SetDefaultPolicyVersionAsync(ctx workflow.Context, input *iam.SetDefaultPolicyVersionInput) *IamSetDefaultPolicyVersionResult

    SetSecurityTokenServicePreferences(ctx workflow.Context, input *iam.SetSecurityTokenServicePreferencesInput) (*iam.SetSecurityTokenServicePreferencesOutput, error)
    SetSecurityTokenServicePreferencesAsync(ctx workflow.Context, input *iam.SetSecurityTokenServicePreferencesInput) *IamSetSecurityTokenServicePreferencesResult

    SimulateCustomPolicy(ctx workflow.Context, input *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error)
    SimulateCustomPolicyAsync(ctx workflow.Context, input *iam.SimulateCustomPolicyInput) *IamSimulateCustomPolicyResult

    SimulatePrincipalPolicy(ctx workflow.Context, input *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error)
    SimulatePrincipalPolicyAsync(ctx workflow.Context, input *iam.SimulatePrincipalPolicyInput) *IamSimulatePrincipalPolicyResult

    TagRole(ctx workflow.Context, input *iam.TagRoleInput) (*iam.TagRoleOutput, error)
    TagRoleAsync(ctx workflow.Context, input *iam.TagRoleInput) *IamTagRoleResult

    TagUser(ctx workflow.Context, input *iam.TagUserInput) (*iam.TagUserOutput, error)
    TagUserAsync(ctx workflow.Context, input *iam.TagUserInput) *IamTagUserResult

    UntagRole(ctx workflow.Context, input *iam.UntagRoleInput) (*iam.UntagRoleOutput, error)
    UntagRoleAsync(ctx workflow.Context, input *iam.UntagRoleInput) *IamUntagRoleResult

    UntagUser(ctx workflow.Context, input *iam.UntagUserInput) (*iam.UntagUserOutput, error)
    UntagUserAsync(ctx workflow.Context, input *iam.UntagUserInput) *IamUntagUserResult

    UpdateAccessKey(ctx workflow.Context, input *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error)
    UpdateAccessKeyAsync(ctx workflow.Context, input *iam.UpdateAccessKeyInput) *IamUpdateAccessKeyResult

    UpdateAccountPasswordPolicy(ctx workflow.Context, input *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error)
    UpdateAccountPasswordPolicyAsync(ctx workflow.Context, input *iam.UpdateAccountPasswordPolicyInput) *IamUpdateAccountPasswordPolicyResult

    UpdateAssumeRolePolicy(ctx workflow.Context, input *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error)
    UpdateAssumeRolePolicyAsync(ctx workflow.Context, input *iam.UpdateAssumeRolePolicyInput) *IamUpdateAssumeRolePolicyResult

    UpdateGroup(ctx workflow.Context, input *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error)
    UpdateGroupAsync(ctx workflow.Context, input *iam.UpdateGroupInput) *IamUpdateGroupResult

    UpdateLoginProfile(ctx workflow.Context, input *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error)
    UpdateLoginProfileAsync(ctx workflow.Context, input *iam.UpdateLoginProfileInput) *IamUpdateLoginProfileResult

    UpdateOpenIDConnectProviderThumbprint(ctx workflow.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)
    UpdateOpenIDConnectProviderThumbprintAsync(ctx workflow.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) *IamUpdateOpenIDConnectProviderThumbprintResult

    UpdateRole(ctx workflow.Context, input *iam.UpdateRoleInput) (*iam.UpdateRoleOutput, error)
    UpdateRoleAsync(ctx workflow.Context, input *iam.UpdateRoleInput) *IamUpdateRoleResult

    UpdateRoleDescription(ctx workflow.Context, input *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error)
    UpdateRoleDescriptionAsync(ctx workflow.Context, input *iam.UpdateRoleDescriptionInput) *IamUpdateRoleDescriptionResult

    UpdateSAMLProvider(ctx workflow.Context, input *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error)
    UpdateSAMLProviderAsync(ctx workflow.Context, input *iam.UpdateSAMLProviderInput) *IamUpdateSAMLProviderResult

    UpdateSSHPublicKey(ctx workflow.Context, input *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error)
    UpdateSSHPublicKeyAsync(ctx workflow.Context, input *iam.UpdateSSHPublicKeyInput) *IamUpdateSSHPublicKeyResult

    UpdateServerCertificate(ctx workflow.Context, input *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error)
    UpdateServerCertificateAsync(ctx workflow.Context, input *iam.UpdateServerCertificateInput) *IamUpdateServerCertificateResult

    UpdateServiceSpecificCredential(ctx workflow.Context, input *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error)
    UpdateServiceSpecificCredentialAsync(ctx workflow.Context, input *iam.UpdateServiceSpecificCredentialInput) *IamUpdateServiceSpecificCredentialResult

    UpdateSigningCertificate(ctx workflow.Context, input *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error)
    UpdateSigningCertificateAsync(ctx workflow.Context, input *iam.UpdateSigningCertificateInput) *IamUpdateSigningCertificateResult

    UpdateUser(ctx workflow.Context, input *iam.UpdateUserInput) (*iam.UpdateUserOutput, error)
    UpdateUserAsync(ctx workflow.Context, input *iam.UpdateUserInput) *IamUpdateUserResult

    UploadSSHPublicKey(ctx workflow.Context, input *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error)
    UploadSSHPublicKeyAsync(ctx workflow.Context, input *iam.UploadSSHPublicKeyInput) *IamUploadSSHPublicKeyResult

    UploadServerCertificate(ctx workflow.Context, input *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error)
    UploadServerCertificateAsync(ctx workflow.Context, input *iam.UploadServerCertificateInput) *IamUploadServerCertificateResult

    UploadSigningCertificate(ctx workflow.Context, input *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error)
    UploadSigningCertificateAsync(ctx workflow.Context, input *iam.UploadSigningCertificateInput) *IamUploadSigningCertificateResult

    WaitUntilInstanceProfileExists(ctx workflow.Context, input *iam.GetInstanceProfileInput) error
    WaitUntilPolicyExists(ctx workflow.Context, input *iam.GetPolicyInput) error
    WaitUntilRoleExists(ctx workflow.Context, input *iam.GetRoleInput) error
    WaitUntilUserExists(ctx workflow.Context, input *iam.GetUserInput) error}
type IamAddClientIDToOpenIDConnectProviderResult struct {
	Result workflow.Future
}

func (r *IamAddClientIDToOpenIDConnectProviderResult) Get(ctx workflow.Context) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
    var output iam.AddClientIDToOpenIDConnectProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamAddRoleToInstanceProfileResult struct {
	Result workflow.Future
}

func (r *IamAddRoleToInstanceProfileResult) Get(ctx workflow.Context) (*iam.AddRoleToInstanceProfileOutput, error) {
    var output iam.AddRoleToInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamAddUserToGroupResult struct {
	Result workflow.Future
}

func (r *IamAddUserToGroupResult) Get(ctx workflow.Context) (*iam.AddUserToGroupOutput, error) {
    var output iam.AddUserToGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamAttachGroupPolicyResult struct {
	Result workflow.Future
}

func (r *IamAttachGroupPolicyResult) Get(ctx workflow.Context) (*iam.AttachGroupPolicyOutput, error) {
    var output iam.AttachGroupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamAttachRolePolicyResult struct {
	Result workflow.Future
}

func (r *IamAttachRolePolicyResult) Get(ctx workflow.Context) (*iam.AttachRolePolicyOutput, error) {
    var output iam.AttachRolePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamAttachUserPolicyResult struct {
	Result workflow.Future
}

func (r *IamAttachUserPolicyResult) Get(ctx workflow.Context) (*iam.AttachUserPolicyOutput, error) {
    var output iam.AttachUserPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamChangePasswordResult struct {
	Result workflow.Future
}

func (r *IamChangePasswordResult) Get(ctx workflow.Context) (*iam.ChangePasswordOutput, error) {
    var output iam.ChangePasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateAccessKeyResult struct {
	Result workflow.Future
}

func (r *IamCreateAccessKeyResult) Get(ctx workflow.Context) (*iam.CreateAccessKeyOutput, error) {
    var output iam.CreateAccessKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateAccountAliasResult struct {
	Result workflow.Future
}

func (r *IamCreateAccountAliasResult) Get(ctx workflow.Context) (*iam.CreateAccountAliasOutput, error) {
    var output iam.CreateAccountAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateGroupResult struct {
	Result workflow.Future
}

func (r *IamCreateGroupResult) Get(ctx workflow.Context) (*iam.CreateGroupOutput, error) {
    var output iam.CreateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateInstanceProfileResult struct {
	Result workflow.Future
}

func (r *IamCreateInstanceProfileResult) Get(ctx workflow.Context) (*iam.CreateInstanceProfileOutput, error) {
    var output iam.CreateInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateLoginProfileResult struct {
	Result workflow.Future
}

func (r *IamCreateLoginProfileResult) Get(ctx workflow.Context) (*iam.CreateLoginProfileOutput, error) {
    var output iam.CreateLoginProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateOpenIDConnectProviderResult struct {
	Result workflow.Future
}

func (r *IamCreateOpenIDConnectProviderResult) Get(ctx workflow.Context) (*iam.CreateOpenIDConnectProviderOutput, error) {
    var output iam.CreateOpenIDConnectProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreatePolicyResult struct {
	Result workflow.Future
}

func (r *IamCreatePolicyResult) Get(ctx workflow.Context) (*iam.CreatePolicyOutput, error) {
    var output iam.CreatePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreatePolicyVersionResult struct {
	Result workflow.Future
}

func (r *IamCreatePolicyVersionResult) Get(ctx workflow.Context) (*iam.CreatePolicyVersionOutput, error) {
    var output iam.CreatePolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateRoleResult struct {
	Result workflow.Future
}

func (r *IamCreateRoleResult) Get(ctx workflow.Context) (*iam.CreateRoleOutput, error) {
    var output iam.CreateRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateSAMLProviderResult struct {
	Result workflow.Future
}

func (r *IamCreateSAMLProviderResult) Get(ctx workflow.Context) (*iam.CreateSAMLProviderOutput, error) {
    var output iam.CreateSAMLProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateServiceLinkedRoleResult struct {
	Result workflow.Future
}

func (r *IamCreateServiceLinkedRoleResult) Get(ctx workflow.Context) (*iam.CreateServiceLinkedRoleOutput, error) {
    var output iam.CreateServiceLinkedRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateServiceSpecificCredentialResult struct {
	Result workflow.Future
}

func (r *IamCreateServiceSpecificCredentialResult) Get(ctx workflow.Context) (*iam.CreateServiceSpecificCredentialOutput, error) {
    var output iam.CreateServiceSpecificCredentialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateUserResult struct {
	Result workflow.Future
}

func (r *IamCreateUserResult) Get(ctx workflow.Context) (*iam.CreateUserOutput, error) {
    var output iam.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamCreateVirtualMFADeviceResult struct {
	Result workflow.Future
}

func (r *IamCreateVirtualMFADeviceResult) Get(ctx workflow.Context) (*iam.CreateVirtualMFADeviceOutput, error) {
    var output iam.CreateVirtualMFADeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeactivateMFADeviceResult struct {
	Result workflow.Future
}

func (r *IamDeactivateMFADeviceResult) Get(ctx workflow.Context) (*iam.DeactivateMFADeviceOutput, error) {
    var output iam.DeactivateMFADeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteAccessKeyResult struct {
	Result workflow.Future
}

func (r *IamDeleteAccessKeyResult) Get(ctx workflow.Context) (*iam.DeleteAccessKeyOutput, error) {
    var output iam.DeleteAccessKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteAccountAliasResult struct {
	Result workflow.Future
}

func (r *IamDeleteAccountAliasResult) Get(ctx workflow.Context) (*iam.DeleteAccountAliasOutput, error) {
    var output iam.DeleteAccountAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteAccountPasswordPolicyResult struct {
	Result workflow.Future
}

func (r *IamDeleteAccountPasswordPolicyResult) Get(ctx workflow.Context) (*iam.DeleteAccountPasswordPolicyOutput, error) {
    var output iam.DeleteAccountPasswordPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteGroupResult struct {
	Result workflow.Future
}

func (r *IamDeleteGroupResult) Get(ctx workflow.Context) (*iam.DeleteGroupOutput, error) {
    var output iam.DeleteGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteGroupPolicyResult struct {
	Result workflow.Future
}

func (r *IamDeleteGroupPolicyResult) Get(ctx workflow.Context) (*iam.DeleteGroupPolicyOutput, error) {
    var output iam.DeleteGroupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteInstanceProfileResult struct {
	Result workflow.Future
}

func (r *IamDeleteInstanceProfileResult) Get(ctx workflow.Context) (*iam.DeleteInstanceProfileOutput, error) {
    var output iam.DeleteInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteLoginProfileResult struct {
	Result workflow.Future
}

func (r *IamDeleteLoginProfileResult) Get(ctx workflow.Context) (*iam.DeleteLoginProfileOutput, error) {
    var output iam.DeleteLoginProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteOpenIDConnectProviderResult struct {
	Result workflow.Future
}

func (r *IamDeleteOpenIDConnectProviderResult) Get(ctx workflow.Context) (*iam.DeleteOpenIDConnectProviderOutput, error) {
    var output iam.DeleteOpenIDConnectProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeletePolicyResult struct {
	Result workflow.Future
}

func (r *IamDeletePolicyResult) Get(ctx workflow.Context) (*iam.DeletePolicyOutput, error) {
    var output iam.DeletePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeletePolicyVersionResult struct {
	Result workflow.Future
}

func (r *IamDeletePolicyVersionResult) Get(ctx workflow.Context) (*iam.DeletePolicyVersionOutput, error) {
    var output iam.DeletePolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteRoleResult struct {
	Result workflow.Future
}

func (r *IamDeleteRoleResult) Get(ctx workflow.Context) (*iam.DeleteRoleOutput, error) {
    var output iam.DeleteRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteRolePermissionsBoundaryResult struct {
	Result workflow.Future
}

func (r *IamDeleteRolePermissionsBoundaryResult) Get(ctx workflow.Context) (*iam.DeleteRolePermissionsBoundaryOutput, error) {
    var output iam.DeleteRolePermissionsBoundaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteRolePolicyResult struct {
	Result workflow.Future
}

func (r *IamDeleteRolePolicyResult) Get(ctx workflow.Context) (*iam.DeleteRolePolicyOutput, error) {
    var output iam.DeleteRolePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteSAMLProviderResult struct {
	Result workflow.Future
}

func (r *IamDeleteSAMLProviderResult) Get(ctx workflow.Context) (*iam.DeleteSAMLProviderOutput, error) {
    var output iam.DeleteSAMLProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteSSHPublicKeyResult struct {
	Result workflow.Future
}

func (r *IamDeleteSSHPublicKeyResult) Get(ctx workflow.Context) (*iam.DeleteSSHPublicKeyOutput, error) {
    var output iam.DeleteSSHPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteServerCertificateResult struct {
	Result workflow.Future
}

func (r *IamDeleteServerCertificateResult) Get(ctx workflow.Context) (*iam.DeleteServerCertificateOutput, error) {
    var output iam.DeleteServerCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteServiceLinkedRoleResult struct {
	Result workflow.Future
}

func (r *IamDeleteServiceLinkedRoleResult) Get(ctx workflow.Context) (*iam.DeleteServiceLinkedRoleOutput, error) {
    var output iam.DeleteServiceLinkedRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteServiceSpecificCredentialResult struct {
	Result workflow.Future
}

func (r *IamDeleteServiceSpecificCredentialResult) Get(ctx workflow.Context) (*iam.DeleteServiceSpecificCredentialOutput, error) {
    var output iam.DeleteServiceSpecificCredentialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteSigningCertificateResult struct {
	Result workflow.Future
}

func (r *IamDeleteSigningCertificateResult) Get(ctx workflow.Context) (*iam.DeleteSigningCertificateOutput, error) {
    var output iam.DeleteSigningCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteUserResult struct {
	Result workflow.Future
}

func (r *IamDeleteUserResult) Get(ctx workflow.Context) (*iam.DeleteUserOutput, error) {
    var output iam.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteUserPermissionsBoundaryResult struct {
	Result workflow.Future
}

func (r *IamDeleteUserPermissionsBoundaryResult) Get(ctx workflow.Context) (*iam.DeleteUserPermissionsBoundaryOutput, error) {
    var output iam.DeleteUserPermissionsBoundaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteUserPolicyResult struct {
	Result workflow.Future
}

func (r *IamDeleteUserPolicyResult) Get(ctx workflow.Context) (*iam.DeleteUserPolicyOutput, error) {
    var output iam.DeleteUserPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDeleteVirtualMFADeviceResult struct {
	Result workflow.Future
}

func (r *IamDeleteVirtualMFADeviceResult) Get(ctx workflow.Context) (*iam.DeleteVirtualMFADeviceOutput, error) {
    var output iam.DeleteVirtualMFADeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDetachGroupPolicyResult struct {
	Result workflow.Future
}

func (r *IamDetachGroupPolicyResult) Get(ctx workflow.Context) (*iam.DetachGroupPolicyOutput, error) {
    var output iam.DetachGroupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDetachRolePolicyResult struct {
	Result workflow.Future
}

func (r *IamDetachRolePolicyResult) Get(ctx workflow.Context) (*iam.DetachRolePolicyOutput, error) {
    var output iam.DetachRolePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamDetachUserPolicyResult struct {
	Result workflow.Future
}

func (r *IamDetachUserPolicyResult) Get(ctx workflow.Context) (*iam.DetachUserPolicyOutput, error) {
    var output iam.DetachUserPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamEnableMFADeviceResult struct {
	Result workflow.Future
}

func (r *IamEnableMFADeviceResult) Get(ctx workflow.Context) (*iam.EnableMFADeviceOutput, error) {
    var output iam.EnableMFADeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGenerateCredentialReportResult struct {
	Result workflow.Future
}

func (r *IamGenerateCredentialReportResult) Get(ctx workflow.Context) (*iam.GenerateCredentialReportOutput, error) {
    var output iam.GenerateCredentialReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGenerateOrganizationsAccessReportResult struct {
	Result workflow.Future
}

func (r *IamGenerateOrganizationsAccessReportResult) Get(ctx workflow.Context) (*iam.GenerateOrganizationsAccessReportOutput, error) {
    var output iam.GenerateOrganizationsAccessReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGenerateServiceLastAccessedDetailsResult struct {
	Result workflow.Future
}

func (r *IamGenerateServiceLastAccessedDetailsResult) Get(ctx workflow.Context) (*iam.GenerateServiceLastAccessedDetailsOutput, error) {
    var output iam.GenerateServiceLastAccessedDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetAccessKeyLastUsedResult struct {
	Result workflow.Future
}

func (r *IamGetAccessKeyLastUsedResult) Get(ctx workflow.Context) (*iam.GetAccessKeyLastUsedOutput, error) {
    var output iam.GetAccessKeyLastUsedOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetAccountAuthorizationDetailsResult struct {
	Result workflow.Future
}

func (r *IamGetAccountAuthorizationDetailsResult) Get(ctx workflow.Context) (*iam.GetAccountAuthorizationDetailsOutput, error) {
    var output iam.GetAccountAuthorizationDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetAccountPasswordPolicyResult struct {
	Result workflow.Future
}

func (r *IamGetAccountPasswordPolicyResult) Get(ctx workflow.Context) (*iam.GetAccountPasswordPolicyOutput, error) {
    var output iam.GetAccountPasswordPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetAccountSummaryResult struct {
	Result workflow.Future
}

func (r *IamGetAccountSummaryResult) Get(ctx workflow.Context) (*iam.GetAccountSummaryOutput, error) {
    var output iam.GetAccountSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetContextKeysForCustomPolicyResult struct {
	Result workflow.Future
}

func (r *IamGetContextKeysForCustomPolicyResult) Get(ctx workflow.Context) (*iam.GetContextKeysForPolicyResponse, error) {
    var output iam.GetContextKeysForPolicyResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetContextKeysForPrincipalPolicyResult struct {
	Result workflow.Future
}

func (r *IamGetContextKeysForPrincipalPolicyResult) Get(ctx workflow.Context) (*iam.GetContextKeysForPolicyResponse, error) {
    var output iam.GetContextKeysForPolicyResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetCredentialReportResult struct {
	Result workflow.Future
}

func (r *IamGetCredentialReportResult) Get(ctx workflow.Context) (*iam.GetCredentialReportOutput, error) {
    var output iam.GetCredentialReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetGroupResult struct {
	Result workflow.Future
}

func (r *IamGetGroupResult) Get(ctx workflow.Context) (*iam.GetGroupOutput, error) {
    var output iam.GetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetGroupPolicyResult struct {
	Result workflow.Future
}

func (r *IamGetGroupPolicyResult) Get(ctx workflow.Context) (*iam.GetGroupPolicyOutput, error) {
    var output iam.GetGroupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetInstanceProfileResult struct {
	Result workflow.Future
}

func (r *IamGetInstanceProfileResult) Get(ctx workflow.Context) (*iam.GetInstanceProfileOutput, error) {
    var output iam.GetInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetLoginProfileResult struct {
	Result workflow.Future
}

func (r *IamGetLoginProfileResult) Get(ctx workflow.Context) (*iam.GetLoginProfileOutput, error) {
    var output iam.GetLoginProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetOpenIDConnectProviderResult struct {
	Result workflow.Future
}

func (r *IamGetOpenIDConnectProviderResult) Get(ctx workflow.Context) (*iam.GetOpenIDConnectProviderOutput, error) {
    var output iam.GetOpenIDConnectProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetOrganizationsAccessReportResult struct {
	Result workflow.Future
}

func (r *IamGetOrganizationsAccessReportResult) Get(ctx workflow.Context) (*iam.GetOrganizationsAccessReportOutput, error) {
    var output iam.GetOrganizationsAccessReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetPolicyResult struct {
	Result workflow.Future
}

func (r *IamGetPolicyResult) Get(ctx workflow.Context) (*iam.GetPolicyOutput, error) {
    var output iam.GetPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetPolicyVersionResult struct {
	Result workflow.Future
}

func (r *IamGetPolicyVersionResult) Get(ctx workflow.Context) (*iam.GetPolicyVersionOutput, error) {
    var output iam.GetPolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetRoleResult struct {
	Result workflow.Future
}

func (r *IamGetRoleResult) Get(ctx workflow.Context) (*iam.GetRoleOutput, error) {
    var output iam.GetRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetRolePolicyResult struct {
	Result workflow.Future
}

func (r *IamGetRolePolicyResult) Get(ctx workflow.Context) (*iam.GetRolePolicyOutput, error) {
    var output iam.GetRolePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetSAMLProviderResult struct {
	Result workflow.Future
}

func (r *IamGetSAMLProviderResult) Get(ctx workflow.Context) (*iam.GetSAMLProviderOutput, error) {
    var output iam.GetSAMLProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetSSHPublicKeyResult struct {
	Result workflow.Future
}

func (r *IamGetSSHPublicKeyResult) Get(ctx workflow.Context) (*iam.GetSSHPublicKeyOutput, error) {
    var output iam.GetSSHPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetServerCertificateResult struct {
	Result workflow.Future
}

func (r *IamGetServerCertificateResult) Get(ctx workflow.Context) (*iam.GetServerCertificateOutput, error) {
    var output iam.GetServerCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetServiceLastAccessedDetailsResult struct {
	Result workflow.Future
}

func (r *IamGetServiceLastAccessedDetailsResult) Get(ctx workflow.Context) (*iam.GetServiceLastAccessedDetailsOutput, error) {
    var output iam.GetServiceLastAccessedDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetServiceLastAccessedDetailsWithEntitiesResult struct {
	Result workflow.Future
}

func (r *IamGetServiceLastAccessedDetailsWithEntitiesResult) Get(ctx workflow.Context) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error) {
    var output iam.GetServiceLastAccessedDetailsWithEntitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetServiceLinkedRoleDeletionStatusResult struct {
	Result workflow.Future
}

func (r *IamGetServiceLinkedRoleDeletionStatusResult) Get(ctx workflow.Context) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
    var output iam.GetServiceLinkedRoleDeletionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetUserResult struct {
	Result workflow.Future
}

func (r *IamGetUserResult) Get(ctx workflow.Context) (*iam.GetUserOutput, error) {
    var output iam.GetUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamGetUserPolicyResult struct {
	Result workflow.Future
}

func (r *IamGetUserPolicyResult) Get(ctx workflow.Context) (*iam.GetUserPolicyOutput, error) {
    var output iam.GetUserPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListAccessKeysResult struct {
	Result workflow.Future
}

func (r *IamListAccessKeysResult) Get(ctx workflow.Context) (*iam.ListAccessKeysOutput, error) {
    var output iam.ListAccessKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListAccountAliasesResult struct {
	Result workflow.Future
}

func (r *IamListAccountAliasesResult) Get(ctx workflow.Context) (*iam.ListAccountAliasesOutput, error) {
    var output iam.ListAccountAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListAttachedGroupPoliciesResult struct {
	Result workflow.Future
}

func (r *IamListAttachedGroupPoliciesResult) Get(ctx workflow.Context) (*iam.ListAttachedGroupPoliciesOutput, error) {
    var output iam.ListAttachedGroupPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListAttachedRolePoliciesResult struct {
	Result workflow.Future
}

func (r *IamListAttachedRolePoliciesResult) Get(ctx workflow.Context) (*iam.ListAttachedRolePoliciesOutput, error) {
    var output iam.ListAttachedRolePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListAttachedUserPoliciesResult struct {
	Result workflow.Future
}

func (r *IamListAttachedUserPoliciesResult) Get(ctx workflow.Context) (*iam.ListAttachedUserPoliciesOutput, error) {
    var output iam.ListAttachedUserPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListEntitiesForPolicyResult struct {
	Result workflow.Future
}

func (r *IamListEntitiesForPolicyResult) Get(ctx workflow.Context) (*iam.ListEntitiesForPolicyOutput, error) {
    var output iam.ListEntitiesForPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListGroupPoliciesResult struct {
	Result workflow.Future
}

func (r *IamListGroupPoliciesResult) Get(ctx workflow.Context) (*iam.ListGroupPoliciesOutput, error) {
    var output iam.ListGroupPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListGroupsResult struct {
	Result workflow.Future
}

func (r *IamListGroupsResult) Get(ctx workflow.Context) (*iam.ListGroupsOutput, error) {
    var output iam.ListGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListGroupsForUserResult struct {
	Result workflow.Future
}

func (r *IamListGroupsForUserResult) Get(ctx workflow.Context) (*iam.ListGroupsForUserOutput, error) {
    var output iam.ListGroupsForUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListInstanceProfilesResult struct {
	Result workflow.Future
}

func (r *IamListInstanceProfilesResult) Get(ctx workflow.Context) (*iam.ListInstanceProfilesOutput, error) {
    var output iam.ListInstanceProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListInstanceProfilesForRoleResult struct {
	Result workflow.Future
}

func (r *IamListInstanceProfilesForRoleResult) Get(ctx workflow.Context) (*iam.ListInstanceProfilesForRoleOutput, error) {
    var output iam.ListInstanceProfilesForRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListMFADevicesResult struct {
	Result workflow.Future
}

func (r *IamListMFADevicesResult) Get(ctx workflow.Context) (*iam.ListMFADevicesOutput, error) {
    var output iam.ListMFADevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListOpenIDConnectProvidersResult struct {
	Result workflow.Future
}

func (r *IamListOpenIDConnectProvidersResult) Get(ctx workflow.Context) (*iam.ListOpenIDConnectProvidersOutput, error) {
    var output iam.ListOpenIDConnectProvidersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListPoliciesResult struct {
	Result workflow.Future
}

func (r *IamListPoliciesResult) Get(ctx workflow.Context) (*iam.ListPoliciesOutput, error) {
    var output iam.ListPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListPoliciesGrantingServiceAccessResult struct {
	Result workflow.Future
}

func (r *IamListPoliciesGrantingServiceAccessResult) Get(ctx workflow.Context) (*iam.ListPoliciesGrantingServiceAccessOutput, error) {
    var output iam.ListPoliciesGrantingServiceAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListPolicyVersionsResult struct {
	Result workflow.Future
}

func (r *IamListPolicyVersionsResult) Get(ctx workflow.Context) (*iam.ListPolicyVersionsOutput, error) {
    var output iam.ListPolicyVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListRolePoliciesResult struct {
	Result workflow.Future
}

func (r *IamListRolePoliciesResult) Get(ctx workflow.Context) (*iam.ListRolePoliciesOutput, error) {
    var output iam.ListRolePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListRoleTagsResult struct {
	Result workflow.Future
}

func (r *IamListRoleTagsResult) Get(ctx workflow.Context) (*iam.ListRoleTagsOutput, error) {
    var output iam.ListRoleTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListRolesResult struct {
	Result workflow.Future
}

func (r *IamListRolesResult) Get(ctx workflow.Context) (*iam.ListRolesOutput, error) {
    var output iam.ListRolesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListSAMLProvidersResult struct {
	Result workflow.Future
}

func (r *IamListSAMLProvidersResult) Get(ctx workflow.Context) (*iam.ListSAMLProvidersOutput, error) {
    var output iam.ListSAMLProvidersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListSSHPublicKeysResult struct {
	Result workflow.Future
}

func (r *IamListSSHPublicKeysResult) Get(ctx workflow.Context) (*iam.ListSSHPublicKeysOutput, error) {
    var output iam.ListSSHPublicKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListServerCertificatesResult struct {
	Result workflow.Future
}

func (r *IamListServerCertificatesResult) Get(ctx workflow.Context) (*iam.ListServerCertificatesOutput, error) {
    var output iam.ListServerCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListServiceSpecificCredentialsResult struct {
	Result workflow.Future
}

func (r *IamListServiceSpecificCredentialsResult) Get(ctx workflow.Context) (*iam.ListServiceSpecificCredentialsOutput, error) {
    var output iam.ListServiceSpecificCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListSigningCertificatesResult struct {
	Result workflow.Future
}

func (r *IamListSigningCertificatesResult) Get(ctx workflow.Context) (*iam.ListSigningCertificatesOutput, error) {
    var output iam.ListSigningCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListUserPoliciesResult struct {
	Result workflow.Future
}

func (r *IamListUserPoliciesResult) Get(ctx workflow.Context) (*iam.ListUserPoliciesOutput, error) {
    var output iam.ListUserPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListUserTagsResult struct {
	Result workflow.Future
}

func (r *IamListUserTagsResult) Get(ctx workflow.Context) (*iam.ListUserTagsOutput, error) {
    var output iam.ListUserTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListUsersResult struct {
	Result workflow.Future
}

func (r *IamListUsersResult) Get(ctx workflow.Context) (*iam.ListUsersOutput, error) {
    var output iam.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamListVirtualMFADevicesResult struct {
	Result workflow.Future
}

func (r *IamListVirtualMFADevicesResult) Get(ctx workflow.Context) (*iam.ListVirtualMFADevicesOutput, error) {
    var output iam.ListVirtualMFADevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamPutGroupPolicyResult struct {
	Result workflow.Future
}

func (r *IamPutGroupPolicyResult) Get(ctx workflow.Context) (*iam.PutGroupPolicyOutput, error) {
    var output iam.PutGroupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamPutRolePermissionsBoundaryResult struct {
	Result workflow.Future
}

func (r *IamPutRolePermissionsBoundaryResult) Get(ctx workflow.Context) (*iam.PutRolePermissionsBoundaryOutput, error) {
    var output iam.PutRolePermissionsBoundaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamPutRolePolicyResult struct {
	Result workflow.Future
}

func (r *IamPutRolePolicyResult) Get(ctx workflow.Context) (*iam.PutRolePolicyOutput, error) {
    var output iam.PutRolePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamPutUserPermissionsBoundaryResult struct {
	Result workflow.Future
}

func (r *IamPutUserPermissionsBoundaryResult) Get(ctx workflow.Context) (*iam.PutUserPermissionsBoundaryOutput, error) {
    var output iam.PutUserPermissionsBoundaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamPutUserPolicyResult struct {
	Result workflow.Future
}

func (r *IamPutUserPolicyResult) Get(ctx workflow.Context) (*iam.PutUserPolicyOutput, error) {
    var output iam.PutUserPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamRemoveClientIDFromOpenIDConnectProviderResult struct {
	Result workflow.Future
}

func (r *IamRemoveClientIDFromOpenIDConnectProviderResult) Get(ctx workflow.Context) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {
    var output iam.RemoveClientIDFromOpenIDConnectProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamRemoveRoleFromInstanceProfileResult struct {
	Result workflow.Future
}

func (r *IamRemoveRoleFromInstanceProfileResult) Get(ctx workflow.Context) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
    var output iam.RemoveRoleFromInstanceProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamRemoveUserFromGroupResult struct {
	Result workflow.Future
}

func (r *IamRemoveUserFromGroupResult) Get(ctx workflow.Context) (*iam.RemoveUserFromGroupOutput, error) {
    var output iam.RemoveUserFromGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamResetServiceSpecificCredentialResult struct {
	Result workflow.Future
}

func (r *IamResetServiceSpecificCredentialResult) Get(ctx workflow.Context) (*iam.ResetServiceSpecificCredentialOutput, error) {
    var output iam.ResetServiceSpecificCredentialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamResyncMFADeviceResult struct {
	Result workflow.Future
}

func (r *IamResyncMFADeviceResult) Get(ctx workflow.Context) (*iam.ResyncMFADeviceOutput, error) {
    var output iam.ResyncMFADeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamSetDefaultPolicyVersionResult struct {
	Result workflow.Future
}

func (r *IamSetDefaultPolicyVersionResult) Get(ctx workflow.Context) (*iam.SetDefaultPolicyVersionOutput, error) {
    var output iam.SetDefaultPolicyVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamSetSecurityTokenServicePreferencesResult struct {
	Result workflow.Future
}

func (r *IamSetSecurityTokenServicePreferencesResult) Get(ctx workflow.Context) (*iam.SetSecurityTokenServicePreferencesOutput, error) {
    var output iam.SetSecurityTokenServicePreferencesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamSimulateCustomPolicyResult struct {
	Result workflow.Future
}

func (r *IamSimulateCustomPolicyResult) Get(ctx workflow.Context) (*iam.SimulatePolicyResponse, error) {
    var output iam.SimulatePolicyResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamSimulatePrincipalPolicyResult struct {
	Result workflow.Future
}

func (r *IamSimulatePrincipalPolicyResult) Get(ctx workflow.Context) (*iam.SimulatePolicyResponse, error) {
    var output iam.SimulatePolicyResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamTagRoleResult struct {
	Result workflow.Future
}

func (r *IamTagRoleResult) Get(ctx workflow.Context) (*iam.TagRoleOutput, error) {
    var output iam.TagRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamTagUserResult struct {
	Result workflow.Future
}

func (r *IamTagUserResult) Get(ctx workflow.Context) (*iam.TagUserOutput, error) {
    var output iam.TagUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUntagRoleResult struct {
	Result workflow.Future
}

func (r *IamUntagRoleResult) Get(ctx workflow.Context) (*iam.UntagRoleOutput, error) {
    var output iam.UntagRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUntagUserResult struct {
	Result workflow.Future
}

func (r *IamUntagUserResult) Get(ctx workflow.Context) (*iam.UntagUserOutput, error) {
    var output iam.UntagUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateAccessKeyResult struct {
	Result workflow.Future
}

func (r *IamUpdateAccessKeyResult) Get(ctx workflow.Context) (*iam.UpdateAccessKeyOutput, error) {
    var output iam.UpdateAccessKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateAccountPasswordPolicyResult struct {
	Result workflow.Future
}

func (r *IamUpdateAccountPasswordPolicyResult) Get(ctx workflow.Context) (*iam.UpdateAccountPasswordPolicyOutput, error) {
    var output iam.UpdateAccountPasswordPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateAssumeRolePolicyResult struct {
	Result workflow.Future
}

func (r *IamUpdateAssumeRolePolicyResult) Get(ctx workflow.Context) (*iam.UpdateAssumeRolePolicyOutput, error) {
    var output iam.UpdateAssumeRolePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateGroupResult struct {
	Result workflow.Future
}

func (r *IamUpdateGroupResult) Get(ctx workflow.Context) (*iam.UpdateGroupOutput, error) {
    var output iam.UpdateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateLoginProfileResult struct {
	Result workflow.Future
}

func (r *IamUpdateLoginProfileResult) Get(ctx workflow.Context) (*iam.UpdateLoginProfileOutput, error) {
    var output iam.UpdateLoginProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateOpenIDConnectProviderThumbprintResult struct {
	Result workflow.Future
}

func (r *IamUpdateOpenIDConnectProviderThumbprintResult) Get(ctx workflow.Context) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {
    var output iam.UpdateOpenIDConnectProviderThumbprintOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateRoleResult struct {
	Result workflow.Future
}

func (r *IamUpdateRoleResult) Get(ctx workflow.Context) (*iam.UpdateRoleOutput, error) {
    var output iam.UpdateRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateRoleDescriptionResult struct {
	Result workflow.Future
}

func (r *IamUpdateRoleDescriptionResult) Get(ctx workflow.Context) (*iam.UpdateRoleDescriptionOutput, error) {
    var output iam.UpdateRoleDescriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateSAMLProviderResult struct {
	Result workflow.Future
}

func (r *IamUpdateSAMLProviderResult) Get(ctx workflow.Context) (*iam.UpdateSAMLProviderOutput, error) {
    var output iam.UpdateSAMLProviderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateSSHPublicKeyResult struct {
	Result workflow.Future
}

func (r *IamUpdateSSHPublicKeyResult) Get(ctx workflow.Context) (*iam.UpdateSSHPublicKeyOutput, error) {
    var output iam.UpdateSSHPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateServerCertificateResult struct {
	Result workflow.Future
}

func (r *IamUpdateServerCertificateResult) Get(ctx workflow.Context) (*iam.UpdateServerCertificateOutput, error) {
    var output iam.UpdateServerCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateServiceSpecificCredentialResult struct {
	Result workflow.Future
}

func (r *IamUpdateServiceSpecificCredentialResult) Get(ctx workflow.Context) (*iam.UpdateServiceSpecificCredentialOutput, error) {
    var output iam.UpdateServiceSpecificCredentialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateSigningCertificateResult struct {
	Result workflow.Future
}

func (r *IamUpdateSigningCertificateResult) Get(ctx workflow.Context) (*iam.UpdateSigningCertificateOutput, error) {
    var output iam.UpdateSigningCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUpdateUserResult struct {
	Result workflow.Future
}

func (r *IamUpdateUserResult) Get(ctx workflow.Context) (*iam.UpdateUserOutput, error) {
    var output iam.UpdateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUploadSSHPublicKeyResult struct {
	Result workflow.Future
}

func (r *IamUploadSSHPublicKeyResult) Get(ctx workflow.Context) (*iam.UploadSSHPublicKeyOutput, error) {
    var output iam.UploadSSHPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUploadServerCertificateResult struct {
	Result workflow.Future
}

func (r *IamUploadServerCertificateResult) Get(ctx workflow.Context) (*iam.UploadServerCertificateOutput, error) {
    var output iam.UploadServerCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IamUploadSigningCertificateResult struct {
	Result workflow.Future
}

func (r *IamUploadSigningCertificateResult) Get(ctx workflow.Context) (*iam.UploadSigningCertificateOutput, error) {
    var output iam.UploadSigningCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type IAMStub struct {
    activities IAMClient
}

func NewIAMStub() IAMClient {
    return &IAMStub{}
}

func (a *IAMStub) AddClientIDToOpenIDConnectProvider(ctx workflow.Context, input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
    var output iam.AddClientIDToOpenIDConnectProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddClientIDToOpenIDConnectProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) AddRoleToInstanceProfile(ctx workflow.Context, input *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error) {
    var output iam.AddRoleToInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddRoleToInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) AddUserToGroup(ctx workflow.Context, input *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error) {
    var output iam.AddUserToGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddUserToGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) AttachGroupPolicy(ctx workflow.Context, input *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error) {
    var output iam.AttachGroupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachGroupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) AttachRolePolicy(ctx workflow.Context, input *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error) {
    var output iam.AttachRolePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachRolePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) AttachUserPolicy(ctx workflow.Context, input *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error) {
    var output iam.AttachUserPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachUserPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ChangePassword(ctx workflow.Context, input *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error) {
    var output iam.ChangePasswordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ChangePassword, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateAccessKey(ctx workflow.Context, input *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
    var output iam.CreateAccessKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAccessKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateAccountAlias(ctx workflow.Context, input *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error) {
    var output iam.CreateAccountAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAccountAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateGroup(ctx workflow.Context, input *iam.CreateGroupInput) (*iam.CreateGroupOutput, error) {
    var output iam.CreateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateInstanceProfile(ctx workflow.Context, input *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error) {
    var output iam.CreateInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateLoginProfile(ctx workflow.Context, input *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error) {
    var output iam.CreateLoginProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLoginProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateOpenIDConnectProvider(ctx workflow.Context, input *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error) {
    var output iam.CreateOpenIDConnectProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOpenIDConnectProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreatePolicy(ctx workflow.Context, input *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
    var output iam.CreatePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreatePolicyVersion(ctx workflow.Context, input *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error) {
    var output iam.CreatePolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateRole(ctx workflow.Context, input *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {
    var output iam.CreateRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateSAMLProvider(ctx workflow.Context, input *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error) {
    var output iam.CreateSAMLProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSAMLProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateServiceLinkedRole(ctx workflow.Context, input *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error) {
    var output iam.CreateServiceLinkedRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateServiceLinkedRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateServiceSpecificCredential(ctx workflow.Context, input *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error) {
    var output iam.CreateServiceSpecificCredentialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateServiceSpecificCredential, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateUser(ctx workflow.Context, input *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
    var output iam.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) CreateVirtualMFADevice(ctx workflow.Context, input *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error) {
    var output iam.CreateVirtualMFADeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVirtualMFADevice, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeactivateMFADevice(ctx workflow.Context, input *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error) {
    var output iam.DeactivateMFADeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeactivateMFADevice, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteAccessKey(ctx workflow.Context, input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
    var output iam.DeleteAccessKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteAccountAlias(ctx workflow.Context, input *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error) {
    var output iam.DeleteAccountAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccountAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteAccountPasswordPolicy(ctx workflow.Context, input *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error) {
    var output iam.DeleteAccountPasswordPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccountPasswordPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteGroup(ctx workflow.Context, input *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error) {
    var output iam.DeleteGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteGroupPolicy(ctx workflow.Context, input *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error) {
    var output iam.DeleteGroupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteInstanceProfile(ctx workflow.Context, input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {
    var output iam.DeleteInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteLoginProfile(ctx workflow.Context, input *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error) {
    var output iam.DeleteLoginProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLoginProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteOpenIDConnectProvider(ctx workflow.Context, input *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error) {
    var output iam.DeleteOpenIDConnectProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOpenIDConnectProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeletePolicy(ctx workflow.Context, input *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
    var output iam.DeletePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeletePolicyVersion(ctx workflow.Context, input *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error) {
    var output iam.DeletePolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteRole(ctx workflow.Context, input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
    var output iam.DeleteRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteRolePermissionsBoundary(ctx workflow.Context, input *iam.DeleteRolePermissionsBoundaryInput) (*iam.DeleteRolePermissionsBoundaryOutput, error) {
    var output iam.DeleteRolePermissionsBoundaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRolePermissionsBoundary, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteRolePolicy(ctx workflow.Context, input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {
    var output iam.DeleteRolePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRolePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteSAMLProvider(ctx workflow.Context, input *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error) {
    var output iam.DeleteSAMLProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSAMLProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteSSHPublicKey(ctx workflow.Context, input *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error) {
    var output iam.DeleteSSHPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSSHPublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteServerCertificate(ctx workflow.Context, input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {
    var output iam.DeleteServerCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteServerCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteServiceLinkedRole(ctx workflow.Context, input *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error) {
    var output iam.DeleteServiceLinkedRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteServiceLinkedRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteServiceSpecificCredential(ctx workflow.Context, input *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error) {
    var output iam.DeleteServiceSpecificCredentialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteServiceSpecificCredential, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteSigningCertificate(ctx workflow.Context, input *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error) {
    var output iam.DeleteSigningCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSigningCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteUser(ctx workflow.Context, input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
    var output iam.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteUserPermissionsBoundary(ctx workflow.Context, input *iam.DeleteUserPermissionsBoundaryInput) (*iam.DeleteUserPermissionsBoundaryOutput, error) {
    var output iam.DeleteUserPermissionsBoundaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserPermissionsBoundary, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteUserPolicy(ctx workflow.Context, input *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
    var output iam.DeleteUserPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DeleteVirtualMFADevice(ctx workflow.Context, input *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error) {
    var output iam.DeleteVirtualMFADeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVirtualMFADevice, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DetachGroupPolicy(ctx workflow.Context, input *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error) {
    var output iam.DetachGroupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachGroupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DetachRolePolicy(ctx workflow.Context, input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
    var output iam.DetachRolePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachRolePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) DetachUserPolicy(ctx workflow.Context, input *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error) {
    var output iam.DetachUserPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachUserPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) EnableMFADevice(ctx workflow.Context, input *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error) {
    var output iam.EnableMFADeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableMFADevice, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GenerateCredentialReport(ctx workflow.Context, input *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error) {
    var output iam.GenerateCredentialReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GenerateCredentialReport, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GenerateOrganizationsAccessReport(ctx workflow.Context, input *iam.GenerateOrganizationsAccessReportInput) (*iam.GenerateOrganizationsAccessReportOutput, error) {
    var output iam.GenerateOrganizationsAccessReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GenerateOrganizationsAccessReport, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GenerateServiceLastAccessedDetails(ctx workflow.Context, input *iam.GenerateServiceLastAccessedDetailsInput) (*iam.GenerateServiceLastAccessedDetailsOutput, error) {
    var output iam.GenerateServiceLastAccessedDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GenerateServiceLastAccessedDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetAccessKeyLastUsed(ctx workflow.Context, input *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error) {
    var output iam.GetAccessKeyLastUsedOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccessKeyLastUsed, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetAccountAuthorizationDetails(ctx workflow.Context, input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error) {
    var output iam.GetAccountAuthorizationDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountAuthorizationDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetAccountPasswordPolicy(ctx workflow.Context, input *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error) {
    var output iam.GetAccountPasswordPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountPasswordPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetAccountSummary(ctx workflow.Context, input *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error) {
    var output iam.GetAccountSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetContextKeysForCustomPolicy(ctx workflow.Context, input *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
    var output iam.GetContextKeysForPolicyResponse
    err := workflow.ExecuteActivity(ctx, a.activities.GetContextKeysForCustomPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetContextKeysForPrincipalPolicy(ctx workflow.Context, input *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {
    var output iam.GetContextKeysForPolicyResponse
    err := workflow.ExecuteActivity(ctx, a.activities.GetContextKeysForPrincipalPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetCredentialReport(ctx workflow.Context, input *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error) {
    var output iam.GetCredentialReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCredentialReport, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetGroup(ctx workflow.Context, input *iam.GetGroupInput) (*iam.GetGroupOutput, error) {
    var output iam.GetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetGroupPolicy(ctx workflow.Context, input *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error) {
    var output iam.GetGroupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetInstanceProfile(ctx workflow.Context, input *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error) {
    var output iam.GetInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetLoginProfile(ctx workflow.Context, input *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error) {
    var output iam.GetLoginProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLoginProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetOpenIDConnectProvider(ctx workflow.Context, input *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error) {
    var output iam.GetOpenIDConnectProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOpenIDConnectProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetOrganizationsAccessReport(ctx workflow.Context, input *iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error) {
    var output iam.GetOrganizationsAccessReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOrganizationsAccessReport, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetPolicy(ctx workflow.Context, input *iam.GetPolicyInput) (*iam.GetPolicyOutput, error) {
    var output iam.GetPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetPolicyVersion(ctx workflow.Context, input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error) {
    var output iam.GetPolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetRole(ctx workflow.Context, input *iam.GetRoleInput) (*iam.GetRoleOutput, error) {
    var output iam.GetRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetRolePolicy(ctx workflow.Context, input *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error) {
    var output iam.GetRolePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRolePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetSAMLProvider(ctx workflow.Context, input *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error) {
    var output iam.GetSAMLProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSAMLProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetSSHPublicKey(ctx workflow.Context, input *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error) {
    var output iam.GetSSHPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSSHPublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetServerCertificate(ctx workflow.Context, input *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error) {
    var output iam.GetServerCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServerCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetServiceLastAccessedDetails(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error) {
    var output iam.GetServiceLastAccessedDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceLastAccessedDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetServiceLastAccessedDetailsWithEntities(ctx workflow.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error) {
    var output iam.GetServiceLastAccessedDetailsWithEntitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceLastAccessedDetailsWithEntities, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetServiceLinkedRoleDeletionStatus(ctx workflow.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
    var output iam.GetServiceLinkedRoleDeletionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceLinkedRoleDeletionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetUser(ctx workflow.Context, input *iam.GetUserInput) (*iam.GetUserOutput, error) {
    var output iam.GetUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) GetUserPolicy(ctx workflow.Context, input *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error) {
    var output iam.GetUserPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUserPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListAccessKeys(ctx workflow.Context, input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
    var output iam.ListAccessKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAccessKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListAccountAliases(ctx workflow.Context, input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error) {
    var output iam.ListAccountAliasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAccountAliases, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListAttachedGroupPolicies(ctx workflow.Context, input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error) {
    var output iam.ListAttachedGroupPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttachedGroupPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListAttachedRolePolicies(ctx workflow.Context, input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
    var output iam.ListAttachedRolePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttachedRolePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListAttachedUserPolicies(ctx workflow.Context, input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {
    var output iam.ListAttachedUserPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAttachedUserPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListEntitiesForPolicy(ctx workflow.Context, input *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error) {
    var output iam.ListEntitiesForPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEntitiesForPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListGroupPolicies(ctx workflow.Context, input *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error) {
    var output iam.ListGroupPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroupPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListGroups(ctx workflow.Context, input *iam.ListGroupsInput) (*iam.ListGroupsOutput, error) {
    var output iam.ListGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListGroupsForUser(ctx workflow.Context, input *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error) {
    var output iam.ListGroupsForUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroupsForUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListInstanceProfiles(ctx workflow.Context, input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {
    var output iam.ListInstanceProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInstanceProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListInstanceProfilesForRole(ctx workflow.Context, input *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error) {
    var output iam.ListInstanceProfilesForRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInstanceProfilesForRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListMFADevices(ctx workflow.Context, input *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error) {
    var output iam.ListMFADevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMFADevices, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListOpenIDConnectProviders(ctx workflow.Context, input *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error) {
    var output iam.ListOpenIDConnectProvidersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOpenIDConnectProviders, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListPolicies(ctx workflow.Context, input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
    var output iam.ListPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListPoliciesGrantingServiceAccess(ctx workflow.Context, input *iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error) {
    var output iam.ListPoliciesGrantingServiceAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPoliciesGrantingServiceAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListPolicyVersions(ctx workflow.Context, input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error) {
    var output iam.ListPolicyVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicyVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListRolePolicies(ctx workflow.Context, input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
    var output iam.ListRolePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRolePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListRoleTags(ctx workflow.Context, input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error) {
    var output iam.ListRoleTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRoleTags, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListRoles(ctx workflow.Context, input *iam.ListRolesInput) (*iam.ListRolesOutput, error) {
    var output iam.ListRolesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRoles, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListSAMLProviders(ctx workflow.Context, input *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error) {
    var output iam.ListSAMLProvidersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSAMLProviders, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListSSHPublicKeys(ctx workflow.Context, input *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error) {
    var output iam.ListSSHPublicKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSSHPublicKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListServerCertificates(ctx workflow.Context, input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
    var output iam.ListServerCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServerCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListServiceSpecificCredentials(ctx workflow.Context, input *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error) {
    var output iam.ListServiceSpecificCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServiceSpecificCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListSigningCertificates(ctx workflow.Context, input *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error) {
    var output iam.ListSigningCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSigningCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListUserPolicies(ctx workflow.Context, input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
    var output iam.ListUserPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListUserTags(ctx workflow.Context, input *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error) {
    var output iam.ListUserTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserTags, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListUsers(ctx workflow.Context, input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
    var output iam.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ListVirtualMFADevices(ctx workflow.Context, input *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error) {
    var output iam.ListVirtualMFADevicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVirtualMFADevices, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) PutGroupPolicy(ctx workflow.Context, input *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error) {
    var output iam.PutGroupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutGroupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) PutRolePermissionsBoundary(ctx workflow.Context, input *iam.PutRolePermissionsBoundaryInput) (*iam.PutRolePermissionsBoundaryOutput, error) {
    var output iam.PutRolePermissionsBoundaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRolePermissionsBoundary, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) PutRolePolicy(ctx workflow.Context, input *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error) {
    var output iam.PutRolePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRolePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) PutUserPermissionsBoundary(ctx workflow.Context, input *iam.PutUserPermissionsBoundaryInput) (*iam.PutUserPermissionsBoundaryOutput, error) {
    var output iam.PutUserPermissionsBoundaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutUserPermissionsBoundary, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) PutUserPolicy(ctx workflow.Context, input *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {
    var output iam.PutUserPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutUserPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) RemoveClientIDFromOpenIDConnectProvider(ctx workflow.Context, input *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {
    var output iam.RemoveClientIDFromOpenIDConnectProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveClientIDFromOpenIDConnectProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) RemoveRoleFromInstanceProfile(ctx workflow.Context, input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
    var output iam.RemoveRoleFromInstanceProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveRoleFromInstanceProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) RemoveUserFromGroup(ctx workflow.Context, input *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error) {
    var output iam.RemoveUserFromGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveUserFromGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ResetServiceSpecificCredential(ctx workflow.Context, input *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error) {
    var output iam.ResetServiceSpecificCredentialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetServiceSpecificCredential, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) ResyncMFADevice(ctx workflow.Context, input *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error) {
    var output iam.ResyncMFADeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResyncMFADevice, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) SetDefaultPolicyVersion(ctx workflow.Context, input *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error) {
    var output iam.SetDefaultPolicyVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetDefaultPolicyVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) SetSecurityTokenServicePreferences(ctx workflow.Context, input *iam.SetSecurityTokenServicePreferencesInput) (*iam.SetSecurityTokenServicePreferencesOutput, error) {
    var output iam.SetSecurityTokenServicePreferencesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetSecurityTokenServicePreferences, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) SimulateCustomPolicy(ctx workflow.Context, input *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error) {
    var output iam.SimulatePolicyResponse
    err := workflow.ExecuteActivity(ctx, a.activities.SimulateCustomPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) SimulatePrincipalPolicy(ctx workflow.Context, input *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error) {
    var output iam.SimulatePolicyResponse
    err := workflow.ExecuteActivity(ctx, a.activities.SimulatePrincipalPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) TagRole(ctx workflow.Context, input *iam.TagRoleInput) (*iam.TagRoleOutput, error) {
    var output iam.TagRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) TagUser(ctx workflow.Context, input *iam.TagUserInput) (*iam.TagUserOutput, error) {
    var output iam.TagUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UntagRole(ctx workflow.Context, input *iam.UntagRoleInput) (*iam.UntagRoleOutput, error) {
    var output iam.UntagRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UntagUser(ctx workflow.Context, input *iam.UntagUserInput) (*iam.UntagUserOutput, error) {
    var output iam.UntagUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateAccessKey(ctx workflow.Context, input *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error) {
    var output iam.UpdateAccessKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccessKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateAccountPasswordPolicy(ctx workflow.Context, input *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error) {
    var output iam.UpdateAccountPasswordPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccountPasswordPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateAssumeRolePolicy(ctx workflow.Context, input *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error) {
    var output iam.UpdateAssumeRolePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAssumeRolePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateGroup(ctx workflow.Context, input *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error) {
    var output iam.UpdateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateLoginProfile(ctx workflow.Context, input *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error) {
    var output iam.UpdateLoginProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLoginProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateOpenIDConnectProviderThumbprint(ctx workflow.Context, input *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {
    var output iam.UpdateOpenIDConnectProviderThumbprintOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateOpenIDConnectProviderThumbprint, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateRole(ctx workflow.Context, input *iam.UpdateRoleInput) (*iam.UpdateRoleOutput, error) {
    var output iam.UpdateRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRole, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateRoleDescription(ctx workflow.Context, input *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error) {
    var output iam.UpdateRoleDescriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRoleDescription, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateSAMLProvider(ctx workflow.Context, input *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error) {
    var output iam.UpdateSAMLProviderOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSAMLProvider, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateSSHPublicKey(ctx workflow.Context, input *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error) {
    var output iam.UpdateSSHPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSSHPublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateServerCertificate(ctx workflow.Context, input *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error) {
    var output iam.UpdateServerCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServerCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateServiceSpecificCredential(ctx workflow.Context, input *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error) {
    var output iam.UpdateServiceSpecificCredentialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServiceSpecificCredential, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateSigningCertificate(ctx workflow.Context, input *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error) {
    var output iam.UpdateSigningCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSigningCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UpdateUser(ctx workflow.Context, input *iam.UpdateUserInput) (*iam.UpdateUserOutput, error) {
    var output iam.UpdateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UploadSSHPublicKey(ctx workflow.Context, input *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error) {
    var output iam.UploadSSHPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadSSHPublicKey, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UploadServerCertificate(ctx workflow.Context, input *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error) {
    var output iam.UploadServerCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadServerCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *IAMStub) UploadSigningCertificate(ctx workflow.Context, input *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error) {
    var output iam.UploadSigningCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadSigningCertificate, input).Get(ctx, &output)
    return &output, err
}


func (a *IAMStub) WaitUntilInstanceProfileExists(ctx workflow.Context, input *iam.GetInstanceProfileInput) error {
    return a.client.WaitUntilInstanceProfileExists(input)
}


func (a *IAMStub) WaitUntilPolicyExists(ctx workflow.Context, input *iam.GetPolicyInput) error {
    return a.client.WaitUntilPolicyExists(input)
}


func (a *IAMStub) WaitUntilRoleExists(ctx workflow.Context, input *iam.GetRoleInput) error {
    return a.client.WaitUntilRoleExists(input)
}


func (a *IAMStub) WaitUntilUserExists(ctx workflow.Context, input *iam.GetUserInput) error {
    return a.client.WaitUntilUserExists(input)
}
