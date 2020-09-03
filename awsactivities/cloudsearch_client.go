package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"go.temporal.io/sdk/workflow"
)

type CloudSearchClient interface {
    BuildSuggesters(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error)
    BuildSuggestersAsync(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) *CloudsearchBuildSuggestersResult

    CreateDomain(ctx workflow.Context, input *cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error)
    CreateDomainAsync(ctx workflow.Context, input *cloudsearch.CreateDomainInput) *CloudsearchCreateDomainResult

    DefineAnalysisScheme(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error)
    DefineAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) *CloudsearchDefineAnalysisSchemeResult

    DefineExpression(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error)
    DefineExpressionAsync(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) *CloudsearchDefineExpressionResult

    DefineIndexField(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error)
    DefineIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) *CloudsearchDefineIndexFieldResult

    DefineSuggester(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error)
    DefineSuggesterAsync(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) *CloudsearchDefineSuggesterResult

    DeleteAnalysisScheme(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error)
    DeleteAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) *CloudsearchDeleteAnalysisSchemeResult

    DeleteDomain(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error)
    DeleteDomainAsync(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) *CloudsearchDeleteDomainResult

    DeleteExpression(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error)
    DeleteExpressionAsync(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) *CloudsearchDeleteExpressionResult

    DeleteIndexField(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error)
    DeleteIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) *CloudsearchDeleteIndexFieldResult

    DeleteSuggester(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error)
    DeleteSuggesterAsync(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) *CloudsearchDeleteSuggesterResult

    DescribeAnalysisSchemes(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error)
    DescribeAnalysisSchemesAsync(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) *CloudsearchDescribeAnalysisSchemesResult

    DescribeAvailabilityOptions(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)
    DescribeAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) *CloudsearchDescribeAvailabilityOptionsResult

    DescribeDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error)
    DescribeDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) *CloudsearchDescribeDomainEndpointOptionsResult

    DescribeDomains(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error)
    DescribeDomainsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) *CloudsearchDescribeDomainsResult

    DescribeExpressions(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error)
    DescribeExpressionsAsync(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) *CloudsearchDescribeExpressionsResult

    DescribeIndexFields(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error)
    DescribeIndexFieldsAsync(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) *CloudsearchDescribeIndexFieldsResult

    DescribeScalingParameters(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error)
    DescribeScalingParametersAsync(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) *CloudsearchDescribeScalingParametersResult

    DescribeServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)
    DescribeServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) *CloudsearchDescribeServiceAccessPoliciesResult

    DescribeSuggesters(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error)
    DescribeSuggestersAsync(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) *CloudsearchDescribeSuggestersResult

    IndexDocuments(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error)
    IndexDocumentsAsync(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) *CloudsearchIndexDocumentsResult

    ListDomainNames(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error)
    ListDomainNamesAsync(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) *CloudsearchListDomainNamesResult

    UpdateAvailabilityOptions(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)
    UpdateAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) *CloudsearchUpdateAvailabilityOptionsResult

    UpdateDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error)
    UpdateDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) *CloudsearchUpdateDomainEndpointOptionsResult

    UpdateScalingParameters(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error)
    UpdateScalingParametersAsync(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) *CloudsearchUpdateScalingParametersResult

    UpdateServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
    UpdateServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) *CloudsearchUpdateServiceAccessPoliciesResult
}
type CloudsearchBuildSuggestersResult struct {
	Result workflow.Future
}

func (r *CloudsearchBuildSuggestersResult) Get(ctx workflow.Context) (*cloudsearch.BuildSuggestersOutput, error) {
    var output cloudsearch.BuildSuggestersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchCreateDomainResult struct {
	Result workflow.Future
}

func (r *CloudsearchCreateDomainResult) Get(ctx workflow.Context) (*cloudsearch.CreateDomainOutput, error) {
    var output cloudsearch.CreateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDefineAnalysisSchemeResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineAnalysisSchemeResult) Get(ctx workflow.Context) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
    var output cloudsearch.DefineAnalysisSchemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDefineExpressionResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineExpressionResult) Get(ctx workflow.Context) (*cloudsearch.DefineExpressionOutput, error) {
    var output cloudsearch.DefineExpressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDefineIndexFieldResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineIndexFieldResult) Get(ctx workflow.Context) (*cloudsearch.DefineIndexFieldOutput, error) {
    var output cloudsearch.DefineIndexFieldOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDefineSuggesterResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineSuggesterResult) Get(ctx workflow.Context) (*cloudsearch.DefineSuggesterOutput, error) {
    var output cloudsearch.DefineSuggesterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDeleteAnalysisSchemeResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteAnalysisSchemeResult) Get(ctx workflow.Context) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
    var output cloudsearch.DeleteAnalysisSchemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDeleteDomainResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteDomainResult) Get(ctx workflow.Context) (*cloudsearch.DeleteDomainOutput, error) {
    var output cloudsearch.DeleteDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDeleteExpressionResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteExpressionResult) Get(ctx workflow.Context) (*cloudsearch.DeleteExpressionOutput, error) {
    var output cloudsearch.DeleteExpressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDeleteIndexFieldResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteIndexFieldResult) Get(ctx workflow.Context) (*cloudsearch.DeleteIndexFieldOutput, error) {
    var output cloudsearch.DeleteIndexFieldOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDeleteSuggesterResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteSuggesterResult) Get(ctx workflow.Context) (*cloudsearch.DeleteSuggesterOutput, error) {
    var output cloudsearch.DeleteSuggesterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeAnalysisSchemesResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeAnalysisSchemesResult) Get(ctx workflow.Context) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
    var output cloudsearch.DescribeAnalysisSchemesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeAvailabilityOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeAvailabilityOptionsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
    var output cloudsearch.DescribeAvailabilityOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeDomainEndpointOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeDomainEndpointOptionsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
    var output cloudsearch.DescribeDomainEndpointOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeDomainsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeDomainsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeDomainsOutput, error) {
    var output cloudsearch.DescribeDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeExpressionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeExpressionsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeExpressionsOutput, error) {
    var output cloudsearch.DescribeExpressionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeIndexFieldsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeIndexFieldsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeIndexFieldsOutput, error) {
    var output cloudsearch.DescribeIndexFieldsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeScalingParametersResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeScalingParametersResult) Get(ctx workflow.Context) (*cloudsearch.DescribeScalingParametersOutput, error) {
    var output cloudsearch.DescribeScalingParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeServiceAccessPoliciesResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeServiceAccessPoliciesResult) Get(ctx workflow.Context) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
    var output cloudsearch.DescribeServiceAccessPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchDescribeSuggestersResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeSuggestersResult) Get(ctx workflow.Context) (*cloudsearch.DescribeSuggestersOutput, error) {
    var output cloudsearch.DescribeSuggestersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchIndexDocumentsResult struct {
	Result workflow.Future
}

func (r *CloudsearchIndexDocumentsResult) Get(ctx workflow.Context) (*cloudsearch.IndexDocumentsOutput, error) {
    var output cloudsearch.IndexDocumentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchListDomainNamesResult struct {
	Result workflow.Future
}

func (r *CloudsearchListDomainNamesResult) Get(ctx workflow.Context) (*cloudsearch.ListDomainNamesOutput, error) {
    var output cloudsearch.ListDomainNamesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchUpdateAvailabilityOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateAvailabilityOptionsResult) Get(ctx workflow.Context) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
    var output cloudsearch.UpdateAvailabilityOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchUpdateDomainEndpointOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateDomainEndpointOptionsResult) Get(ctx workflow.Context) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
    var output cloudsearch.UpdateDomainEndpointOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchUpdateScalingParametersResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateScalingParametersResult) Get(ctx workflow.Context) (*cloudsearch.UpdateScalingParametersOutput, error) {
    var output cloudsearch.UpdateScalingParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchUpdateServiceAccessPoliciesResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateServiceAccessPoliciesResult) Get(ctx workflow.Context) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
    var output cloudsearch.UpdateServiceAccessPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CloudSearchStub struct {
    activities CloudSearchClient
}

func NewCloudSearchStub() CloudSearchClient {
    return &CloudSearchStub{}
}

func (a *CloudSearchStub) BuildSuggesters(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error) {
    var output cloudsearch.BuildSuggestersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BuildSuggesters, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) CreateDomain(ctx workflow.Context, input *cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error) {
    var output cloudsearch.CreateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineAnalysisScheme(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
    var output cloudsearch.DefineAnalysisSchemeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DefineAnalysisScheme, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineExpression(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error) {
    var output cloudsearch.DefineExpressionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DefineExpression, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineIndexField(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error) {
    var output cloudsearch.DefineIndexFieldOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DefineIndexField, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineSuggester(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error) {
    var output cloudsearch.DefineSuggesterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DefineSuggester, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteAnalysisScheme(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
    var output cloudsearch.DeleteAnalysisSchemeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAnalysisScheme, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteDomain(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error) {
    var output cloudsearch.DeleteDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteExpression(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error) {
    var output cloudsearch.DeleteExpressionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteExpression, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteIndexField(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error) {
    var output cloudsearch.DeleteIndexFieldOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIndexField, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteSuggester(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error) {
    var output cloudsearch.DeleteSuggesterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSuggester, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeAnalysisSchemes(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
    var output cloudsearch.DescribeAnalysisSchemesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAnalysisSchemes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeAvailabilityOptions(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
    var output cloudsearch.DescribeAvailabilityOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAvailabilityOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
    var output cloudsearch.DescribeDomainEndpointOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDomainEndpointOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeDomains(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error) {
    var output cloudsearch.DescribeDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeExpressions(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error) {
    var output cloudsearch.DescribeExpressionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExpressions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeIndexFields(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error) {
    var output cloudsearch.DescribeIndexFieldsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIndexFields, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeScalingParameters(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error) {
    var output cloudsearch.DescribeScalingParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
    var output cloudsearch.DescribeServiceAccessPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServiceAccessPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeSuggesters(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error) {
    var output cloudsearch.DescribeSuggestersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSuggesters, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) IndexDocuments(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error) {
    var output cloudsearch.IndexDocumentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.IndexDocuments, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) ListDomainNames(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error) {
    var output cloudsearch.ListDomainNamesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomainNames, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateAvailabilityOptions(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
    var output cloudsearch.UpdateAvailabilityOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAvailabilityOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
    var output cloudsearch.UpdateDomainEndpointOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainEndpointOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateScalingParameters(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error) {
    var output cloudsearch.UpdateScalingParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateScalingParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
    var output cloudsearch.UpdateServiceAccessPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServiceAccessPolicies, input).Get(ctx, &output)
    return &output, err
}
