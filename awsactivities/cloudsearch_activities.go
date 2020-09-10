package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CloudSearchActivities struct {
	client cloudsearchiface.CloudSearchAPI
}

func NewCloudSearchActivities(session *session.Session, config ...*aws.Config) *CloudSearchActivities {
	client := cloudsearch.New(session, config...)
	return &CloudSearchActivities{client: client}
}

func (a *CloudSearchActivities) BuildSuggesters(ctx context.Context, input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error) {
	return a.client.BuildSuggestersWithContext(ctx, input)
}

func (a *CloudSearchActivities) CreateDomain(ctx context.Context, input *cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error) {
	return a.client.CreateDomainWithContext(ctx, input)
}

func (a *CloudSearchActivities) DefineAnalysisScheme(ctx context.Context, input *cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
	return a.client.DefineAnalysisSchemeWithContext(ctx, input)
}

func (a *CloudSearchActivities) DefineExpression(ctx context.Context, input *cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error) {
	return a.client.DefineExpressionWithContext(ctx, input)
}

func (a *CloudSearchActivities) DefineIndexField(ctx context.Context, input *cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error) {
	return a.client.DefineIndexFieldWithContext(ctx, input)
}

func (a *CloudSearchActivities) DefineSuggester(ctx context.Context, input *cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error) {
	return a.client.DefineSuggesterWithContext(ctx, input)
}

func (a *CloudSearchActivities) DeleteAnalysisScheme(ctx context.Context, input *cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
	return a.client.DeleteAnalysisSchemeWithContext(ctx, input)
}

func (a *CloudSearchActivities) DeleteDomain(ctx context.Context, input *cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error) {
	return a.client.DeleteDomainWithContext(ctx, input)
}

func (a *CloudSearchActivities) DeleteExpression(ctx context.Context, input *cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error) {
	return a.client.DeleteExpressionWithContext(ctx, input)
}

func (a *CloudSearchActivities) DeleteIndexField(ctx context.Context, input *cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error) {
	return a.client.DeleteIndexFieldWithContext(ctx, input)
}

func (a *CloudSearchActivities) DeleteSuggester(ctx context.Context, input *cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error) {
	return a.client.DeleteSuggesterWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeAnalysisSchemes(ctx context.Context, input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
	return a.client.DescribeAnalysisSchemesWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeAvailabilityOptions(ctx context.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
	return a.client.DescribeAvailabilityOptionsWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeDomainEndpointOptions(ctx context.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
	return a.client.DescribeDomainEndpointOptionsWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeDomains(ctx context.Context, input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error) {
	return a.client.DescribeDomainsWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeExpressions(ctx context.Context, input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error) {
	return a.client.DescribeExpressionsWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeIndexFields(ctx context.Context, input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error) {
	return a.client.DescribeIndexFieldsWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeScalingParameters(ctx context.Context, input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error) {
	return a.client.DescribeScalingParametersWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeServiceAccessPolicies(ctx context.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
	return a.client.DescribeServiceAccessPoliciesWithContext(ctx, input)
}

func (a *CloudSearchActivities) DescribeSuggesters(ctx context.Context, input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error) {
	return a.client.DescribeSuggestersWithContext(ctx, input)
}

func (a *CloudSearchActivities) IndexDocuments(ctx context.Context, input *cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error) {
	return a.client.IndexDocumentsWithContext(ctx, input)
}

func (a *CloudSearchActivities) ListDomainNames(ctx context.Context, input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error) {
	return a.client.ListDomainNamesWithContext(ctx, input)
}

func (a *CloudSearchActivities) UpdateAvailabilityOptions(ctx context.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
	return a.client.UpdateAvailabilityOptionsWithContext(ctx, input)
}

func (a *CloudSearchActivities) UpdateDomainEndpointOptions(ctx context.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
	return a.client.UpdateDomainEndpointOptionsWithContext(ctx, input)
}

func (a *CloudSearchActivities) UpdateScalingParameters(ctx context.Context, input *cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error) {
	return a.client.UpdateScalingParametersWithContext(ctx, input)
}

func (a *CloudSearchActivities) UpdateServiceAccessPolicies(ctx context.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
	return a.client.UpdateServiceAccessPoliciesWithContext(ctx, input)
}
