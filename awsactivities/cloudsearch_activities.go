package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface"
)

type CloudSearchActivities struct {
	client cloudsearchiface.CloudSearchAPI
}

func NewCloudSearchActivities(session *session.Session, config ...*aws.Config) *CloudSearchActivities {
	client := cloudsearch.New(session, config...)
	return &CloudSearchActivities{client: client}
}

func (a *CloudSearchActivities) BuildSuggesters(input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error) {
	return a.client.BuildSuggesters(input)
}

func (a *CloudSearchActivities) CreateDomain(input *cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error) {
	return a.client.CreateDomain(input)
}

func (a *CloudSearchActivities) DefineAnalysisScheme(input *cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
	return a.client.DefineAnalysisScheme(input)
}

func (a *CloudSearchActivities) DefineExpression(input *cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error) {
	return a.client.DefineExpression(input)
}

func (a *CloudSearchActivities) DefineIndexField(input *cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error) {
	return a.client.DefineIndexField(input)
}

func (a *CloudSearchActivities) DefineSuggester(input *cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error) {
	return a.client.DefineSuggester(input)
}

func (a *CloudSearchActivities) DeleteAnalysisScheme(input *cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
	return a.client.DeleteAnalysisScheme(input)
}

func (a *CloudSearchActivities) DeleteDomain(input *cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error) {
	return a.client.DeleteDomain(input)
}

func (a *CloudSearchActivities) DeleteExpression(input *cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error) {
	return a.client.DeleteExpression(input)
}

func (a *CloudSearchActivities) DeleteIndexField(input *cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error) {
	return a.client.DeleteIndexField(input)
}

func (a *CloudSearchActivities) DeleteSuggester(input *cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error) {
	return a.client.DeleteSuggester(input)
}

func (a *CloudSearchActivities) DescribeAnalysisSchemes(input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
	return a.client.DescribeAnalysisSchemes(input)
}

func (a *CloudSearchActivities) DescribeAvailabilityOptions(input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
	return a.client.DescribeAvailabilityOptions(input)
}

func (a *CloudSearchActivities) DescribeDomainEndpointOptions(input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
	return a.client.DescribeDomainEndpointOptions(input)
}

func (a *CloudSearchActivities) DescribeDomains(input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error) {
	return a.client.DescribeDomains(input)
}

func (a *CloudSearchActivities) DescribeExpressions(input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error) {
	return a.client.DescribeExpressions(input)
}

func (a *CloudSearchActivities) DescribeIndexFields(input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error) {
	return a.client.DescribeIndexFields(input)
}

func (a *CloudSearchActivities) DescribeScalingParameters(input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error) {
	return a.client.DescribeScalingParameters(input)
}

func (a *CloudSearchActivities) DescribeServiceAccessPolicies(input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
	return a.client.DescribeServiceAccessPolicies(input)
}

func (a *CloudSearchActivities) DescribeSuggesters(input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error) {
	return a.client.DescribeSuggesters(input)
}

func (a *CloudSearchActivities) IndexDocuments(input *cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error) {
	return a.client.IndexDocuments(input)
}

func (a *CloudSearchActivities) ListDomainNames(input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error) {
	return a.client.ListDomainNames(input)
}

func (a *CloudSearchActivities) UpdateAvailabilityOptions(input *cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
	return a.client.UpdateAvailabilityOptions(input)
}

func (a *CloudSearchActivities) UpdateDomainEndpointOptions(input *cloudsearch.UpdateDomainEndpointOptionsInput) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
	return a.client.UpdateDomainEndpointOptions(input)
}

func (a *CloudSearchActivities) UpdateScalingParameters(input *cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error) {
	return a.client.UpdateScalingParameters(input)
}

func (a *CloudSearchActivities) UpdateServiceAccessPolicies(input *cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
	return a.client.UpdateServiceAccessPolicies(input)
}
