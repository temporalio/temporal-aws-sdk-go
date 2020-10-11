// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicecatalog/servicecatalogiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ServiceCatalogActivities struct {
	client servicecatalogiface.ServiceCatalogAPI

	sessionFactory SessionFactory
}

func NewServiceCatalogActivities(sess *session.Session, config ...*aws.Config) *ServiceCatalogActivities {
	client := servicecatalog.New(sess, config...)
	return &ServiceCatalogActivities{client: client}
}

func NewServiceCatalogActivitiesWithSessionFactory(sessionFactory SessionFactory) *ServiceCatalogActivities {
	return &ServiceCatalogActivities{sessionFactory: sessionFactory}
}

func (a *ServiceCatalogActivities) getClient(ctx context.Context) (servicecatalogiface.ServiceCatalogAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return servicecatalog.New(sess), nil
}

func (a *ServiceCatalogActivities) AcceptPortfolioShare(ctx context.Context, input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AcceptPortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateBudgetWithResource(ctx context.Context, input *servicecatalog.AssociateBudgetWithResourceInput) (*servicecatalog.AssociateBudgetWithResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateBudgetWithResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociatePrincipalWithPortfolio(ctx context.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociatePrincipalWithPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateProductWithPortfolio(ctx context.Context, input *servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateProductWithPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateServiceActionWithProvisioningArtifact(ctx context.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateServiceActionWithProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) AssociateTagOptionWithResource(ctx context.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateTagOptionWithResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) BatchAssociateServiceActionWithProvisioningArtifact(ctx context.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchAssociateServiceActionWithProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) BatchDisassociateServiceActionFromProvisioningArtifact(ctx context.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDisassociateServiceActionFromProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CopyProduct(ctx context.Context, input *servicecatalog.CopyProductInput) (*servicecatalog.CopyProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateConstraint(ctx context.Context, input *servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreatePortfolio(ctx context.Context, input *servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreatePortfolioShare(ctx context.Context, input *servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateProduct(ctx context.Context, input *servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateProvisionedProductPlan(ctx context.Context, input *servicecatalog.CreateProvisionedProductPlanInput) (*servicecatalog.CreateProvisionedProductPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateProvisioningArtifact(ctx context.Context, input *servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateServiceAction(ctx context.Context, input *servicecatalog.CreateServiceActionInput) (*servicecatalog.CreateServiceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) CreateTagOption(ctx context.Context, input *servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteConstraint(ctx context.Context, input *servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeletePortfolio(ctx context.Context, input *servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeletePortfolioShare(ctx context.Context, input *servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteProduct(ctx context.Context, input *servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteProvisionedProductPlan(ctx context.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) (*servicecatalog.DeleteProvisionedProductPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteProvisioningArtifact(ctx context.Context, input *servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteServiceAction(ctx context.Context, input *servicecatalog.DeleteServiceActionInput) (*servicecatalog.DeleteServiceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DeleteTagOption(ctx context.Context, input *servicecatalog.DeleteTagOptionInput) (*servicecatalog.DeleteTagOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeConstraint(ctx context.Context, input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeCopyProductStatus(ctx context.Context, input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCopyProductStatusWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribePortfolio(ctx context.Context, input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribePortfolioShareStatus(ctx context.Context, input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePortfolioShareStatusWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProduct(ctx context.Context, input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProductAsAdmin(ctx context.Context, input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProductAsAdminWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProductView(ctx context.Context, input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProductViewWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisionedProduct(ctx context.Context, input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisionedProductPlan(ctx context.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisioningArtifact(ctx context.Context, input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeProvisioningParameters(ctx context.Context, input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProvisioningParametersWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeRecord(ctx context.Context, input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRecordWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeServiceAction(ctx context.Context, input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeServiceActionExecutionParameters(ctx context.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeServiceActionExecutionParametersWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DescribeTagOption(ctx context.Context, input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisableAWSOrganizationsAccess(ctx context.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableAWSOrganizationsAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateBudgetFromResource(ctx context.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) (*servicecatalog.DisassociateBudgetFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateBudgetFromResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociatePrincipalFromPortfolio(ctx context.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociatePrincipalFromPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateProductFromPortfolio(ctx context.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateProductFromPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateServiceActionFromProvisioningArtifact(ctx context.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateServiceActionFromProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) DisassociateTagOptionFromResource(ctx context.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateTagOptionFromResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) EnableAWSOrganizationsAccess(ctx context.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableAWSOrganizationsAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ExecuteProvisionedProductPlan(ctx context.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExecuteProvisionedProductPlanWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ExecuteProvisionedProductServiceAction(ctx context.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExecuteProvisionedProductServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) GetAWSOrganizationsAccessStatus(ctx context.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAWSOrganizationsAccessStatusWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListAcceptedPortfolioShares(ctx context.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAcceptedPortfolioSharesWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListBudgetsForResource(ctx context.Context, input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBudgetsForResourceWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListConstraintsForPortfolio(ctx context.Context, input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConstraintsForPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListLaunchPaths(ctx context.Context, input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLaunchPathsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListOrganizationPortfolioAccess(ctx context.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListOrganizationPortfolioAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPortfolioAccess(ctx context.Context, input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPortfolioAccessWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPortfolios(ctx context.Context, input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPortfoliosWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPortfoliosForProduct(ctx context.Context, input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPortfoliosForProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListPrincipalsForPortfolio(ctx context.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPrincipalsForPortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListProvisionedProductPlans(ctx context.Context, input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProvisionedProductPlansWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListProvisioningArtifacts(ctx context.Context, input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProvisioningArtifactsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListProvisioningArtifactsForServiceAction(ctx context.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProvisioningArtifactsForServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListRecordHistory(ctx context.Context, input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRecordHistoryWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListResourcesForTagOption(ctx context.Context, input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourcesForTagOptionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListServiceActions(ctx context.Context, input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServiceActionsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListServiceActionsForProvisioningArtifact(ctx context.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServiceActionsForProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListStackInstancesForProvisionedProduct(ctx context.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListStackInstancesForProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ListTagOptions(ctx context.Context, input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagOptionsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ProvisionProduct(ctx context.Context, input *servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ProvisionProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) RejectPortfolioShare(ctx context.Context, input *servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RejectPortfolioShareWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) ScanProvisionedProducts(ctx context.Context, input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ScanProvisionedProductsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) SearchProducts(ctx context.Context, input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchProductsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) SearchProductsAsAdmin(ctx context.Context, input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchProductsAsAdminWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) SearchProvisionedProducts(ctx context.Context, input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchProvisionedProductsWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) TerminateProvisionedProduct(ctx context.Context, input *servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TerminateProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateConstraint(ctx context.Context, input *servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConstraintWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdatePortfolio(ctx context.Context, input *servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePortfolioWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProduct(ctx context.Context, input *servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProvisionedProduct(ctx context.Context, input *servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProvisionedProductWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProvisionedProductProperties(ctx context.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProvisionedProductPropertiesWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateProvisioningArtifact(ctx context.Context, input *servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProvisioningArtifactWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateServiceAction(ctx context.Context, input *servicecatalog.UpdateServiceActionInput) (*servicecatalog.UpdateServiceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateServiceActionWithContext(ctx, input)
}

func (a *ServiceCatalogActivities) UpdateTagOption(ctx context.Context, input *servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateTagOptionWithContext(ctx, input)
}
