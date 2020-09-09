
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/aws/aws-sdk-go/service/codeartifact/codeartifactiface"
)

type CodeArtifactActivities struct {
    client codeartifactiface.CodeArtifactAPI
}

func NewCodeArtifactActivities(session *session.Session, config ...*aws.Config) *CodeArtifactActivities {
    client := codeartifact.New(session, config...)
    return &CodeArtifactActivities{client: client}
}

func (a *CodeArtifactActivities) AssociateExternalConnection(input *codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error) {
    return a.client.AssociateExternalConnection(input)
}

func (a *CodeArtifactActivities) CopyPackageVersions(input *codeartifact.CopyPackageVersionsInput) (*codeartifact.CopyPackageVersionsOutput, error) {
    return a.client.CopyPackageVersions(input)
}

func (a *CodeArtifactActivities) CreateDomain(input *codeartifact.CreateDomainInput) (*codeartifact.CreateDomainOutput, error) {
    return a.client.CreateDomain(input)
}

func (a *CodeArtifactActivities) CreateRepository(input *codeartifact.CreateRepositoryInput) (*codeartifact.CreateRepositoryOutput, error) {
    return a.client.CreateRepository(input)
}

func (a *CodeArtifactActivities) DeleteDomain(input *codeartifact.DeleteDomainInput) (*codeartifact.DeleteDomainOutput, error) {
    return a.client.DeleteDomain(input)
}

func (a *CodeArtifactActivities) DeleteDomainPermissionsPolicy(input *codeartifact.DeleteDomainPermissionsPolicyInput) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error) {
    return a.client.DeleteDomainPermissionsPolicy(input)
}

func (a *CodeArtifactActivities) DeletePackageVersions(input *codeartifact.DeletePackageVersionsInput) (*codeartifact.DeletePackageVersionsOutput, error) {
    return a.client.DeletePackageVersions(input)
}

func (a *CodeArtifactActivities) DeleteRepository(input *codeartifact.DeleteRepositoryInput) (*codeartifact.DeleteRepositoryOutput, error) {
    return a.client.DeleteRepository(input)
}

func (a *CodeArtifactActivities) DeleteRepositoryPermissionsPolicy(input *codeartifact.DeleteRepositoryPermissionsPolicyInput) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error) {
    return a.client.DeleteRepositoryPermissionsPolicy(input)
}

func (a *CodeArtifactActivities) DescribeDomain(input *codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error) {
    return a.client.DescribeDomain(input)
}

func (a *CodeArtifactActivities) DescribePackageVersion(input *codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error) {
    return a.client.DescribePackageVersion(input)
}

func (a *CodeArtifactActivities) DescribeRepository(input *codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error) {
    return a.client.DescribeRepository(input)
}

func (a *CodeArtifactActivities) DisassociateExternalConnection(input *codeartifact.DisassociateExternalConnectionInput) (*codeartifact.DisassociateExternalConnectionOutput, error) {
    return a.client.DisassociateExternalConnection(input)
}

func (a *CodeArtifactActivities) DisposePackageVersions(input *codeartifact.DisposePackageVersionsInput) (*codeartifact.DisposePackageVersionsOutput, error) {
    return a.client.DisposePackageVersions(input)
}

func (a *CodeArtifactActivities) GetAuthorizationToken(input *codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error) {
    return a.client.GetAuthorizationToken(input)
}

func (a *CodeArtifactActivities) GetDomainPermissionsPolicy(input *codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error) {
    return a.client.GetDomainPermissionsPolicy(input)
}

func (a *CodeArtifactActivities) GetPackageVersionAsset(input *codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error) {
    return a.client.GetPackageVersionAsset(input)
}

func (a *CodeArtifactActivities) GetPackageVersionReadme(input *codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error) {
    return a.client.GetPackageVersionReadme(input)
}

func (a *CodeArtifactActivities) GetRepositoryEndpoint(input *codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error) {
    return a.client.GetRepositoryEndpoint(input)
}

func (a *CodeArtifactActivities) GetRepositoryPermissionsPolicy(input *codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error) {
    return a.client.GetRepositoryPermissionsPolicy(input)
}

func (a *CodeArtifactActivities) ListDomains(input *codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error) {
    return a.client.ListDomains(input)
}

func (a *CodeArtifactActivities) ListPackageVersionAssets(input *codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error) {
    return a.client.ListPackageVersionAssets(input)
}

func (a *CodeArtifactActivities) ListPackageVersionDependencies(input *codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error) {
    return a.client.ListPackageVersionDependencies(input)
}

func (a *CodeArtifactActivities) ListPackageVersions(input *codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error) {
    return a.client.ListPackageVersions(input)
}

func (a *CodeArtifactActivities) ListPackages(input *codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error) {
    return a.client.ListPackages(input)
}

func (a *CodeArtifactActivities) ListRepositories(input *codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error) {
    return a.client.ListRepositories(input)
}

func (a *CodeArtifactActivities) ListRepositoriesInDomain(input *codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error) {
    return a.client.ListRepositoriesInDomain(input)
}

func (a *CodeArtifactActivities) PutDomainPermissionsPolicy(input *codeartifact.PutDomainPermissionsPolicyInput) (*codeartifact.PutDomainPermissionsPolicyOutput, error) {
    return a.client.PutDomainPermissionsPolicy(input)
}

func (a *CodeArtifactActivities) PutRepositoryPermissionsPolicy(input *codeartifact.PutRepositoryPermissionsPolicyInput) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error) {
    return a.client.PutRepositoryPermissionsPolicy(input)
}

func (a *CodeArtifactActivities) UpdatePackageVersionsStatus(input *codeartifact.UpdatePackageVersionsStatusInput) (*codeartifact.UpdatePackageVersionsStatusOutput, error) {
    return a.client.UpdatePackageVersionsStatus(input)
}

func (a *CodeArtifactActivities) UpdateRepository(input *codeartifact.UpdateRepositoryInput) (*codeartifact.UpdateRepositoryOutput, error) {
    return a.client.UpdateRepository(input)
}
