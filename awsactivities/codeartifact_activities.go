package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/aws/aws-sdk-go/service/codeartifact/codeartifactiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CodeArtifactActivities struct {
	client codeartifactiface.CodeArtifactAPI
}

func NewCodeArtifactActivities(session *session.Session, config ...*aws.Config) *CodeArtifactActivities {
	client := codeartifact.New(session, config...)
	return &CodeArtifactActivities{client: client}
}

func (a *CodeArtifactActivities) AssociateExternalConnection(ctx context.Context, input *codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error) {
	return a.client.AssociateExternalConnectionWithContext(ctx, input)
}

func (a *CodeArtifactActivities) CopyPackageVersions(ctx context.Context, input *codeartifact.CopyPackageVersionsInput) (*codeartifact.CopyPackageVersionsOutput, error) {
	return a.client.CopyPackageVersionsWithContext(ctx, input)
}

func (a *CodeArtifactActivities) CreateDomain(ctx context.Context, input *codeartifact.CreateDomainInput) (*codeartifact.CreateDomainOutput, error) {
	return a.client.CreateDomainWithContext(ctx, input)
}

func (a *CodeArtifactActivities) CreateRepository(ctx context.Context, input *codeartifact.CreateRepositoryInput) (*codeartifact.CreateRepositoryOutput, error) {
	return a.client.CreateRepositoryWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DeleteDomain(ctx context.Context, input *codeartifact.DeleteDomainInput) (*codeartifact.DeleteDomainOutput, error) {
	return a.client.DeleteDomainWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DeleteDomainPermissionsPolicy(ctx context.Context, input *codeartifact.DeleteDomainPermissionsPolicyInput) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error) {
	return a.client.DeleteDomainPermissionsPolicyWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DeletePackageVersions(ctx context.Context, input *codeartifact.DeletePackageVersionsInput) (*codeartifact.DeletePackageVersionsOutput, error) {
	return a.client.DeletePackageVersionsWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DeleteRepository(ctx context.Context, input *codeartifact.DeleteRepositoryInput) (*codeartifact.DeleteRepositoryOutput, error) {
	return a.client.DeleteRepositoryWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DeleteRepositoryPermissionsPolicy(ctx context.Context, input *codeartifact.DeleteRepositoryPermissionsPolicyInput) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error) {
	return a.client.DeleteRepositoryPermissionsPolicyWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DescribeDomain(ctx context.Context, input *codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error) {
	return a.client.DescribeDomainWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DescribePackageVersion(ctx context.Context, input *codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error) {
	return a.client.DescribePackageVersionWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DescribeRepository(ctx context.Context, input *codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error) {
	return a.client.DescribeRepositoryWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DisassociateExternalConnection(ctx context.Context, input *codeartifact.DisassociateExternalConnectionInput) (*codeartifact.DisassociateExternalConnectionOutput, error) {
	return a.client.DisassociateExternalConnectionWithContext(ctx, input)
}

func (a *CodeArtifactActivities) DisposePackageVersions(ctx context.Context, input *codeartifact.DisposePackageVersionsInput) (*codeartifact.DisposePackageVersionsOutput, error) {
	return a.client.DisposePackageVersionsWithContext(ctx, input)
}

func (a *CodeArtifactActivities) GetAuthorizationToken(ctx context.Context, input *codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error) {
	return a.client.GetAuthorizationTokenWithContext(ctx, input)
}

func (a *CodeArtifactActivities) GetDomainPermissionsPolicy(ctx context.Context, input *codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error) {
	return a.client.GetDomainPermissionsPolicyWithContext(ctx, input)
}

func (a *CodeArtifactActivities) GetPackageVersionAsset(ctx context.Context, input *codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error) {
	return a.client.GetPackageVersionAssetWithContext(ctx, input)
}

func (a *CodeArtifactActivities) GetPackageVersionReadme(ctx context.Context, input *codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error) {
	return a.client.GetPackageVersionReadmeWithContext(ctx, input)
}

func (a *CodeArtifactActivities) GetRepositoryEndpoint(ctx context.Context, input *codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error) {
	return a.client.GetRepositoryEndpointWithContext(ctx, input)
}

func (a *CodeArtifactActivities) GetRepositoryPermissionsPolicy(ctx context.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error) {
	return a.client.GetRepositoryPermissionsPolicyWithContext(ctx, input)
}

func (a *CodeArtifactActivities) ListDomains(ctx context.Context, input *codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error) {
	return a.client.ListDomainsWithContext(ctx, input)
}

func (a *CodeArtifactActivities) ListPackageVersionAssets(ctx context.Context, input *codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error) {
	return a.client.ListPackageVersionAssetsWithContext(ctx, input)
}

func (a *CodeArtifactActivities) ListPackageVersionDependencies(ctx context.Context, input *codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error) {
	return a.client.ListPackageVersionDependenciesWithContext(ctx, input)
}

func (a *CodeArtifactActivities) ListPackageVersions(ctx context.Context, input *codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error) {
	return a.client.ListPackageVersionsWithContext(ctx, input)
}

func (a *CodeArtifactActivities) ListPackages(ctx context.Context, input *codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error) {
	return a.client.ListPackagesWithContext(ctx, input)
}

func (a *CodeArtifactActivities) ListRepositories(ctx context.Context, input *codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error) {
	return a.client.ListRepositoriesWithContext(ctx, input)
}

func (a *CodeArtifactActivities) ListRepositoriesInDomain(ctx context.Context, input *codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error) {
	return a.client.ListRepositoriesInDomainWithContext(ctx, input)
}

func (a *CodeArtifactActivities) PutDomainPermissionsPolicy(ctx context.Context, input *codeartifact.PutDomainPermissionsPolicyInput) (*codeartifact.PutDomainPermissionsPolicyOutput, error) {
	return a.client.PutDomainPermissionsPolicyWithContext(ctx, input)
}

func (a *CodeArtifactActivities) PutRepositoryPermissionsPolicy(ctx context.Context, input *codeartifact.PutRepositoryPermissionsPolicyInput) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error) {
	return a.client.PutRepositoryPermissionsPolicyWithContext(ctx, input)
}

func (a *CodeArtifactActivities) UpdatePackageVersionsStatus(ctx context.Context, input *codeartifact.UpdatePackageVersionsStatusInput) (*codeartifact.UpdatePackageVersionsStatusOutput, error) {
	return a.client.UpdatePackageVersionsStatusWithContext(ctx, input)
}

func (a *CodeArtifactActivities) UpdateRepository(ctx context.Context, input *codeartifact.UpdateRepositoryInput) (*codeartifact.UpdateRepositoryOutput, error) {
	return a.client.UpdateRepositoryWithContext(ctx, input)
}
