package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CloudSearchDomainClient interface {
    Search(ctx workflow.Context, input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error)
    SearchAsync(ctx workflow.Context, input *cloudsearchdomain.SearchInput) *CloudsearchdomainSearchResult

    Suggest(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error)
    SuggestAsync(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) *CloudsearchdomainSuggestResult

    UploadDocuments(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error)
    UploadDocumentsAsync(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) *CloudsearchdomainUploadDocumentsResult
}
type CloudsearchdomainSearchResult struct {
	Result workflow.Future
}

func (r *CloudsearchdomainSearchResult) Get(ctx workflow.Context) (*cloudsearchdomain.SearchOutput, error) {
    var output cloudsearchdomain.SearchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchdomainSuggestResult struct {
	Result workflow.Future
}

func (r *CloudsearchdomainSuggestResult) Get(ctx workflow.Context) (*cloudsearchdomain.SuggestOutput, error) {
    var output cloudsearchdomain.SuggestOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudsearchdomainUploadDocumentsResult struct {
	Result workflow.Future
}

func (r *CloudsearchdomainUploadDocumentsResult) Get(ctx workflow.Context) (*cloudsearchdomain.UploadDocumentsOutput, error) {
    var output cloudsearchdomain.UploadDocumentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CloudSearchDomainStub struct {
    activities awsactivities.CloudSearchDomainActivities
}

func NewCloudSearchDomainStub() CloudSearchDomainClient {
    return &CloudSearchDomainStub{}
}
func (a *CloudSearchDomainStub) Search(ctx workflow.Context, input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
    var output cloudsearchdomain.SearchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Search, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchDomainStub) SearchAsync(ctx workflow.Context, input *cloudsearchdomain.SearchInput) *CloudsearchdomainSearchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Search, input)
    return &CloudsearchdomainSearchResult{Result: future}
}
func (a *CloudSearchDomainStub) Suggest(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
    var output cloudsearchdomain.SuggestOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Suggest, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchDomainStub) SuggestAsync(ctx workflow.Context, input *cloudsearchdomain.SuggestInput) *CloudsearchdomainSuggestResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Suggest, input)
    return &CloudsearchdomainSuggestResult{Result: future}
}
func (a *CloudSearchDomainStub) UploadDocuments(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
    var output cloudsearchdomain.UploadDocumentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadDocuments, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchDomainStub) UploadDocumentsAsync(ctx workflow.Context, input *cloudsearchdomain.UploadDocumentsInput) *CloudsearchdomainUploadDocumentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UploadDocuments, input)
    return &CloudsearchdomainUploadDocumentsResult{Result: future}
}
