
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicecatalog/servicecatalogiface"
)

type ServiceCatalogActivities struct {
	client servicecatalogiface.ServiceCatalogAPI
}

func NewServiceCatalogActivities(client servicecatalogiface.ServiceCatalogAPI) *ServiceCatalogActivities {
    return &ServiceCatalogActivities{client: client}
}

func (a *ServiceCatalogActivities) AcceptPortfolioShare(input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
    return a.client.AcceptPortfolioShare(input)
}

func (a *ServiceCatalogActivities) AssociateBudgetWithResource(input *servicecatalog.AssociateBudgetWithResourceInput) (*servicecatalog.AssociateBudgetWithResourceOutput, error) {
    return a.client.AssociateBudgetWithResource(input)
}

func (a *ServiceCatalogActivities) AssociatePrincipalWithPortfolio(input *servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error) {
    return a.client.AssociatePrincipalWithPortfolio(input)
}

func (a *ServiceCatalogActivities) AssociateProductWithPortfolio(input *servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error) {
    return a.client.AssociateProductWithPortfolio(input)
}

func (a *ServiceCatalogActivities) AssociateServiceActionWithProvisioningArtifact(input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error) {
    return a.client.AssociateServiceActionWithProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) AssociateTagOptionWithResource(input *servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error) {
    return a.client.AssociateTagOptionWithResource(input)
}

func (a *ServiceCatalogActivities) BatchAssociateServiceActionWithProvisioningArtifact(input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
    return a.client.BatchAssociateServiceActionWithProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) BatchDisassociateServiceActionFromProvisioningArtifact(input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
    return a.client.BatchDisassociateServiceActionFromProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) CopyProduct(input *servicecatalog.CopyProductInput) (*servicecatalog.CopyProductOutput, error) {
    return a.client.CopyProduct(input)
}

func (a *ServiceCatalogActivities) CreateConstraint(input *servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error) {
    return a.client.CreateConstraint(input)
}

func (a *ServiceCatalogActivities) CreatePortfolio(input *servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error) {
    return a.client.CreatePortfolio(input)
}

func (a *ServiceCatalogActivities) CreatePortfolioShare(input *servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error) {
    return a.client.CreatePortfolioShare(input)
}

func (a *ServiceCatalogActivities) CreateProduct(input *servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error) {
    return a.client.CreateProduct(input)
}

func (a *ServiceCatalogActivities) CreateProvisionedProductPlan(input *servicecatalog.CreateProvisionedProductPlanInput) (*servicecatalog.CreateProvisionedProductPlanOutput, error) {
    return a.client.CreateProvisionedProductPlan(input)
}

func (a *ServiceCatalogActivities) CreateProvisioningArtifact(input *servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error) {
    return a.client.CreateProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) CreateServiceAction(input *servicecatalog.CreateServiceActionInput) (*servicecatalog.CreateServiceActionOutput, error) {
    return a.client.CreateServiceAction(input)
}

func (a *ServiceCatalogActivities) CreateTagOption(input *servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error) {
    return a.client.CreateTagOption(input)
}

func (a *ServiceCatalogActivities) DeleteConstraint(input *servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error) {
    return a.client.DeleteConstraint(input)
}

func (a *ServiceCatalogActivities) DeletePortfolio(input *servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error) {
    return a.client.DeletePortfolio(input)
}

func (a *ServiceCatalogActivities) DeletePortfolioShare(input *servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error) {
    return a.client.DeletePortfolioShare(input)
}

func (a *ServiceCatalogActivities) DeleteProduct(input *servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error) {
    return a.client.DeleteProduct(input)
}

func (a *ServiceCatalogActivities) DeleteProvisionedProductPlan(input *servicecatalog.DeleteProvisionedProductPlanInput) (*servicecatalog.DeleteProvisionedProductPlanOutput, error) {
    return a.client.DeleteProvisionedProductPlan(input)
}

func (a *ServiceCatalogActivities) DeleteProvisioningArtifact(input *servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error) {
    return a.client.DeleteProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) DeleteServiceAction(input *servicecatalog.DeleteServiceActionInput) (*servicecatalog.DeleteServiceActionOutput, error) {
    return a.client.DeleteServiceAction(input)
}

func (a *ServiceCatalogActivities) DeleteTagOption(input *servicecatalog.DeleteTagOptionInput) (*servicecatalog.DeleteTagOptionOutput, error) {
    return a.client.DeleteTagOption(input)
}

func (a *ServiceCatalogActivities) DescribeConstraint(input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error) {
    return a.client.DescribeConstraint(input)
}

func (a *ServiceCatalogActivities) DescribeCopyProductStatus(input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
    return a.client.DescribeCopyProductStatus(input)
}

func (a *ServiceCatalogActivities) DescribePortfolio(input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error) {
    return a.client.DescribePortfolio(input)
}

func (a *ServiceCatalogActivities) DescribePortfolioShareStatus(input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
    return a.client.DescribePortfolioShareStatus(input)
}

func (a *ServiceCatalogActivities) DescribeProduct(input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error) {
    return a.client.DescribeProduct(input)
}

func (a *ServiceCatalogActivities) DescribeProductAsAdmin(input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error) {
    return a.client.DescribeProductAsAdmin(input)
}

func (a *ServiceCatalogActivities) DescribeProductView(input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error) {
    return a.client.DescribeProductView(input)
}

func (a *ServiceCatalogActivities) DescribeProvisionedProduct(input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error) {
    return a.client.DescribeProvisionedProduct(input)
}

func (a *ServiceCatalogActivities) DescribeProvisionedProductPlan(input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
    return a.client.DescribeProvisionedProductPlan(input)
}

func (a *ServiceCatalogActivities) DescribeProvisioningArtifact(input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
    return a.client.DescribeProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) DescribeProvisioningParameters(input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
    return a.client.DescribeProvisioningParameters(input)
}

func (a *ServiceCatalogActivities) DescribeRecord(input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error) {
    return a.client.DescribeRecord(input)
}

func (a *ServiceCatalogActivities) DescribeServiceAction(input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error) {
    return a.client.DescribeServiceAction(input)
}

func (a *ServiceCatalogActivities) DescribeServiceActionExecutionParameters(input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
    return a.client.DescribeServiceActionExecutionParameters(input)
}

func (a *ServiceCatalogActivities) DescribeTagOption(input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error) {
    return a.client.DescribeTagOption(input)
}

func (a *ServiceCatalogActivities) DisableAWSOrganizationsAccess(input *servicecatalog.DisableAWSOrganizationsAccessInput) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error) {
    return a.client.DisableAWSOrganizationsAccess(input)
}

func (a *ServiceCatalogActivities) DisassociateBudgetFromResource(input *servicecatalog.DisassociateBudgetFromResourceInput) (*servicecatalog.DisassociateBudgetFromResourceOutput, error) {
    return a.client.DisassociateBudgetFromResource(input)
}

func (a *ServiceCatalogActivities) DisassociatePrincipalFromPortfolio(input *servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error) {
    return a.client.DisassociatePrincipalFromPortfolio(input)
}

func (a *ServiceCatalogActivities) DisassociateProductFromPortfolio(input *servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error) {
    return a.client.DisassociateProductFromPortfolio(input)
}

func (a *ServiceCatalogActivities) DisassociateServiceActionFromProvisioningArtifact(input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error) {
    return a.client.DisassociateServiceActionFromProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) DisassociateTagOptionFromResource(input *servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error) {
    return a.client.DisassociateTagOptionFromResource(input)
}

func (a *ServiceCatalogActivities) EnableAWSOrganizationsAccess(input *servicecatalog.EnableAWSOrganizationsAccessInput) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error) {
    return a.client.EnableAWSOrganizationsAccess(input)
}

func (a *ServiceCatalogActivities) ExecuteProvisionedProductPlan(input *servicecatalog.ExecuteProvisionedProductPlanInput) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error) {
    return a.client.ExecuteProvisionedProductPlan(input)
}

func (a *ServiceCatalogActivities) ExecuteProvisionedProductServiceAction(input *servicecatalog.ExecuteProvisionedProductServiceActionInput) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error) {
    return a.client.ExecuteProvisionedProductServiceAction(input)
}

func (a *ServiceCatalogActivities) GetAWSOrganizationsAccessStatus(input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
    return a.client.GetAWSOrganizationsAccessStatus(input)
}

func (a *ServiceCatalogActivities) ListAcceptedPortfolioShares(input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
    return a.client.ListAcceptedPortfolioShares(input)
}

func (a *ServiceCatalogActivities) ListBudgetsForResource(input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error) {
    return a.client.ListBudgetsForResource(input)
}

func (a *ServiceCatalogActivities) ListConstraintsForPortfolio(input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
    return a.client.ListConstraintsForPortfolio(input)
}

func (a *ServiceCatalogActivities) ListLaunchPaths(input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error) {
    return a.client.ListLaunchPaths(input)
}

func (a *ServiceCatalogActivities) ListOrganizationPortfolioAccess(input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
    return a.client.ListOrganizationPortfolioAccess(input)
}

func (a *ServiceCatalogActivities) ListPortfolioAccess(input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error) {
    return a.client.ListPortfolioAccess(input)
}

func (a *ServiceCatalogActivities) ListPortfolios(input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error) {
    return a.client.ListPortfolios(input)
}

func (a *ServiceCatalogActivities) ListPortfoliosForProduct(input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error) {
    return a.client.ListPortfoliosForProduct(input)
}

func (a *ServiceCatalogActivities) ListPrincipalsForPortfolio(input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
    return a.client.ListPrincipalsForPortfolio(input)
}

func (a *ServiceCatalogActivities) ListProvisionedProductPlans(input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
    return a.client.ListProvisionedProductPlans(input)
}

func (a *ServiceCatalogActivities) ListProvisioningArtifacts(input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
    return a.client.ListProvisioningArtifacts(input)
}

func (a *ServiceCatalogActivities) ListProvisioningArtifactsForServiceAction(input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
    return a.client.ListProvisioningArtifactsForServiceAction(input)
}

func (a *ServiceCatalogActivities) ListRecordHistory(input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error) {
    return a.client.ListRecordHistory(input)
}

func (a *ServiceCatalogActivities) ListResourcesForTagOption(input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
    return a.client.ListResourcesForTagOption(input)
}

func (a *ServiceCatalogActivities) ListServiceActions(input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error) {
    return a.client.ListServiceActions(input)
}

func (a *ServiceCatalogActivities) ListServiceActionsForProvisioningArtifact(input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
    return a.client.ListServiceActionsForProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) ListStackInstancesForProvisionedProduct(input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
    return a.client.ListStackInstancesForProvisionedProduct(input)
}

func (a *ServiceCatalogActivities) ListTagOptions(input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error) {
    return a.client.ListTagOptions(input)
}

func (a *ServiceCatalogActivities) ProvisionProduct(input *servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error) {
    return a.client.ProvisionProduct(input)
}

func (a *ServiceCatalogActivities) RejectPortfolioShare(input *servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error) {
    return a.client.RejectPortfolioShare(input)
}

func (a *ServiceCatalogActivities) ScanProvisionedProducts(input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error) {
    return a.client.ScanProvisionedProducts(input)
}

func (a *ServiceCatalogActivities) SearchProducts(input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error) {
    return a.client.SearchProducts(input)
}

func (a *ServiceCatalogActivities) SearchProductsAsAdmin(input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error) {
    return a.client.SearchProductsAsAdmin(input)
}

func (a *ServiceCatalogActivities) SearchProvisionedProducts(input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error) {
    return a.client.SearchProvisionedProducts(input)
}

func (a *ServiceCatalogActivities) TerminateProvisionedProduct(input *servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error) {
    return a.client.TerminateProvisionedProduct(input)
}

func (a *ServiceCatalogActivities) UpdateConstraint(input *servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error) {
    return a.client.UpdateConstraint(input)
}

func (a *ServiceCatalogActivities) UpdatePortfolio(input *servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error) {
    return a.client.UpdatePortfolio(input)
}

func (a *ServiceCatalogActivities) UpdateProduct(input *servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error) {
    return a.client.UpdateProduct(input)
}

func (a *ServiceCatalogActivities) UpdateProvisionedProduct(input *servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error) {
    return a.client.UpdateProvisionedProduct(input)
}

func (a *ServiceCatalogActivities) UpdateProvisionedProductProperties(input *servicecatalog.UpdateProvisionedProductPropertiesInput) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error) {
    return a.client.UpdateProvisionedProductProperties(input)
}

func (a *ServiceCatalogActivities) UpdateProvisioningArtifact(input *servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error) {
    return a.client.UpdateProvisioningArtifact(input)
}

func (a *ServiceCatalogActivities) UpdateServiceAction(input *servicecatalog.UpdateServiceActionInput) (*servicecatalog.UpdateServiceActionOutput, error) {
    return a.client.UpdateServiceAction(input)
}

func (a *ServiceCatalogActivities) UpdateTagOption(input *servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error) {
    return a.client.UpdateTagOption(input)
}
