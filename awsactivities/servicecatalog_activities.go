package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicecatalog/servicecatalogiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ServiceCatalogActivities struct {
	client servicecatalogiface.ServiceCatalogAPI
}

func NewServiceCatalogActivities(session *session.Session, config ...*aws.Config) *ServiceCatalogActivities {
	client := servicecatalog.New(session, config...)
	return &ServiceCatalogActivities{client: client}
}

func (a *ServiceCatalogActivities) AcceptPortfolioShare(ctx context.Context, input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
	return a.client.AcceptPortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateBudgetWithResource(ctx context.Context, input *servicecatalog.AssociateBudgetWithResourceInput) (*servicecatalog.AssociateBudgetWithResourceOutput, error) {
	return a.client.AssociateBudgetWithResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociatePrincipalWithPortfolio(ctx context.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error) {
	return a.client.AssociatePrincipalWithPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateProductWithPortfolio(ctx context.Context, input *servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error) {
	return a.client.AssociateProductWithPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateServiceActionWithProvisioningArtifact(ctx context.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error) {
	return a.client.AssociateServiceActionWithProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateTagOptionWithResource(ctx context.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error) {
	return a.client.AssociateTagOptionWithResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) BatchAssociateServiceActionWithProvisioningArtifact(ctx context.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
	return a.client.BatchAssociateServiceActionWithProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) BatchDisassociateServiceActionFromProvisioningArtifact(ctx context.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
	return a.client.BatchDisassociateServiceActionFromProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CopyProduct(ctx context.Context, input *servicecatalog.CopyProductInput) (*servicecatalog.CopyProductOutput, error) {
	return a.client.CopyProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateConstraint(ctx context.Context, input *servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error) {
	return a.client.CreateConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreatePortfolio(ctx context.Context, input *servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error) {
	return a.client.CreatePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreatePortfolioShare(ctx context.Context, input *servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error) {
	return a.client.CreatePortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateProduct(ctx context.Context, input *servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error) {
	return a.client.CreateProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateProvisionedProductPlan(ctx context.Context, input *servicecatalog.CreateProvisionedProductPlanInput) (*servicecatalog.CreateProvisionedProductPlanOutput, error) {
	return a.client.CreateProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateProvisioningArtifact(ctx context.Context, input *servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error) {
	return a.client.CreateProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateServiceAction(ctx context.Context, input *servicecatalog.CreateServiceActionInput) (*servicecatalog.CreateServiceActionOutput, error) {
	return a.client.CreateServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateTagOption(ctx context.Context, input *servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error) {
	return a.client.CreateTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteConstraint(ctx context.Context, input *servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error) {
	return a.client.DeleteConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeletePortfolio(ctx context.Context, input *servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error) {
	return a.client.DeletePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeletePortfolioShare(ctx context.Context, input *servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error) {
	return a.client.DeletePortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteProduct(ctx context.Context, input *servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error) {
	return a.client.DeleteProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteProvisionedProductPlan(ctx context.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) (*servicecatalog.DeleteProvisionedProductPlanOutput, error) {
	return a.client.DeleteProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteProvisioningArtifact(ctx context.Context, input *servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error) {
	return a.client.DeleteProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteServiceAction(ctx context.Context, input *servicecatalog.DeleteServiceActionInput) (*servicecatalog.DeleteServiceActionOutput, error) {
	return a.client.DeleteServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteTagOption(ctx context.Context, input *servicecatalog.DeleteTagOptionInput) (*servicecatalog.DeleteTagOptionOutput, error) {
	return a.client.DeleteTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeConstraint(ctx context.Context, input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error) {
	return a.client.DescribeConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeCopyProductStatus(ctx context.Context, input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
	return a.client.DescribeCopyProductStatusWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribePortfolio(ctx context.Context, input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error) {
	return a.client.DescribePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribePortfolioShareStatus(ctx context.Context, input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
	return a.client.DescribePortfolioShareStatusWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProduct(ctx context.Context, input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error) {
	return a.client.DescribeProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProductAsAdmin(ctx context.Context, input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error) {
	return a.client.DescribeProductAsAdminWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProductView(ctx context.Context, input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error) {
	return a.client.DescribeProductViewWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisionedProduct(ctx context.Context, input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error) {
	return a.client.DescribeProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisionedProductPlan(ctx context.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
	return a.client.DescribeProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisioningArtifact(ctx context.Context, input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
	return a.client.DescribeProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisioningParameters(ctx context.Context, input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
	return a.client.DescribeProvisioningParametersWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeRecord(ctx context.Context, input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error) {
	return a.client.DescribeRecordWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeServiceAction(ctx context.Context, input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error) {
	return a.client.DescribeServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeServiceActionExecutionParameters(ctx context.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
	return a.client.DescribeServiceActionExecutionParametersWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeTagOption(ctx context.Context, input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error) {
	return a.client.DescribeTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisableAWSOrganizationsAccess(ctx context.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error) {
	return a.client.DisableAWSOrganizationsAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateBudgetFromResource(ctx context.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) (*servicecatalog.DisassociateBudgetFromResourceOutput, error) {
	return a.client.DisassociateBudgetFromResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociatePrincipalFromPortfolio(ctx context.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error) {
	return a.client.DisassociatePrincipalFromPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateProductFromPortfolio(ctx context.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error) {
	return a.client.DisassociateProductFromPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateServiceActionFromProvisioningArtifact(ctx context.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error) {
	return a.client.DisassociateServiceActionFromProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateTagOptionFromResource(ctx context.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error) {
	return a.client.DisassociateTagOptionFromResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) EnableAWSOrganizationsAccess(ctx context.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error) {
	return a.client.EnableAWSOrganizationsAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ExecuteProvisionedProductPlan(ctx context.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error) {
	return a.client.ExecuteProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ExecuteProvisionedProductServiceAction(ctx context.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error) {
	return a.client.ExecuteProvisionedProductServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) GetAWSOrganizationsAccessStatus(ctx context.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
	return a.client.GetAWSOrganizationsAccessStatusWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListAcceptedPortfolioShares(ctx context.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
	return a.client.ListAcceptedPortfolioSharesWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListBudgetsForResource(ctx context.Context, input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error) {
	return a.client.ListBudgetsForResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListConstraintsForPortfolio(ctx context.Context, input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
	return a.client.ListConstraintsForPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListLaunchPaths(ctx context.Context, input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error) {
	return a.client.ListLaunchPathsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListOrganizationPortfolioAccess(ctx context.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
	return a.client.ListOrganizationPortfolioAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPortfolioAccess(ctx context.Context, input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error) {
	return a.client.ListPortfolioAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPortfolios(ctx context.Context, input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error) {
	return a.client.ListPortfoliosWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPortfoliosForProduct(ctx context.Context, input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error) {
	return a.client.ListPortfoliosForProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPrincipalsForPortfolio(ctx context.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
	return a.client.ListPrincipalsForPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListProvisionedProductPlans(ctx context.Context, input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
	return a.client.ListProvisionedProductPlansWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListProvisioningArtifacts(ctx context.Context, input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
	return a.client.ListProvisioningArtifactsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListProvisioningArtifactsForServiceAction(ctx context.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
	return a.client.ListProvisioningArtifactsForServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListRecordHistory(ctx context.Context, input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error) {
	return a.client.ListRecordHistoryWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListResourcesForTagOption(ctx context.Context, input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
	return a.client.ListResourcesForTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListServiceActions(ctx context.Context, input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error) {
	return a.client.ListServiceActionsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListServiceActionsForProvisioningArtifact(ctx context.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
	return a.client.ListServiceActionsForProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListStackInstancesForProvisionedProduct(ctx context.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
	return a.client.ListStackInstancesForProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListTagOptions(ctx context.Context, input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error) {
	return a.client.ListTagOptionsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ProvisionProduct(ctx context.Context, input *servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error) {
	return a.client.ProvisionProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) RejectPortfolioShare(ctx context.Context, input *servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error) {
	return a.client.RejectPortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ScanProvisionedProducts(ctx context.Context, input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error) {
	return a.client.ScanProvisionedProductsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) SearchProducts(ctx context.Context, input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error) {
	return a.client.SearchProductsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) SearchProductsAsAdmin(ctx context.Context, input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error) {
	return a.client.SearchProductsAsAdminWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) SearchProvisionedProducts(ctx context.Context, input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error) {
	return a.client.SearchProvisionedProductsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) TerminateProvisionedProduct(ctx context.Context, input *servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error) {
	return a.client.TerminateProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateConstraint(ctx context.Context, input *servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error) {
	return a.client.UpdateConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdatePortfolio(ctx context.Context, input *servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error) {
	return a.client.UpdatePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProduct(ctx context.Context, input *servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error) {
	return a.client.UpdateProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProvisionedProduct(ctx context.Context, input *servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error) {
	return a.client.UpdateProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProvisionedProductProperties(ctx context.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error) {
	return a.client.UpdateProvisionedProductPropertiesWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProvisioningArtifact(ctx context.Context, input *servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error) {
	return a.client.UpdateProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateServiceAction(ctx context.Context, input *servicecatalog.UpdateServiceActionInput) (*servicecatalog.UpdateServiceActionOutput, error) {
	return a.client.UpdateServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateTagOption(ctx context.Context, input *servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error) {
	return a.client.UpdateTagOptionWithContext(ctx, input)
}
