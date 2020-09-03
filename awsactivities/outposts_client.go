package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/outposts"
	"go.temporal.io/sdk/workflow"
)

type OutpostsClient interface {
    CreateOutpost(ctx workflow.Context, input *outposts.CreateOutpostInput) (*outposts.CreateOutpostOutput, error)
    CreateOutpostAsync(ctx workflow.Context, input *outposts.CreateOutpostInput) *OutpostsCreateOutpostResult

    DeleteOutpost(ctx workflow.Context, input *outposts.DeleteOutpostInput) (*outposts.DeleteOutpostOutput, error)
    DeleteOutpostAsync(ctx workflow.Context, input *outposts.DeleteOutpostInput) *OutpostsDeleteOutpostResult

    DeleteSite(ctx workflow.Context, input *outposts.DeleteSiteInput) (*outposts.DeleteSiteOutput, error)
    DeleteSiteAsync(ctx workflow.Context, input *outposts.DeleteSiteInput) *OutpostsDeleteSiteResult

    GetOutpost(ctx workflow.Context, input *outposts.GetOutpostInput) (*outposts.GetOutpostOutput, error)
    GetOutpostAsync(ctx workflow.Context, input *outposts.GetOutpostInput) *OutpostsGetOutpostResult

    GetOutpostInstanceTypes(ctx workflow.Context, input *outposts.GetOutpostInstanceTypesInput) (*outposts.GetOutpostInstanceTypesOutput, error)
    GetOutpostInstanceTypesAsync(ctx workflow.Context, input *outposts.GetOutpostInstanceTypesInput) *OutpostsGetOutpostInstanceTypesResult

    ListOutposts(ctx workflow.Context, input *outposts.ListOutpostsInput) (*outposts.ListOutpostsOutput, error)
    ListOutpostsAsync(ctx workflow.Context, input *outposts.ListOutpostsInput) *OutpostsListOutpostsResult

    ListSites(ctx workflow.Context, input *outposts.ListSitesInput) (*outposts.ListSitesOutput, error)
    ListSitesAsync(ctx workflow.Context, input *outposts.ListSitesInput) *OutpostsListSitesResult
}
type OutpostsCreateOutpostResult struct {
	Result workflow.Future
}

func (r *OutpostsCreateOutpostResult) Get(ctx workflow.Context) (*outposts.CreateOutpostOutput, error) {
    var output outposts.CreateOutpostOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OutpostsDeleteOutpostResult struct {
	Result workflow.Future
}

func (r *OutpostsDeleteOutpostResult) Get(ctx workflow.Context) (*outposts.DeleteOutpostOutput, error) {
    var output outposts.DeleteOutpostOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OutpostsDeleteSiteResult struct {
	Result workflow.Future
}

func (r *OutpostsDeleteSiteResult) Get(ctx workflow.Context) (*outposts.DeleteSiteOutput, error) {
    var output outposts.DeleteSiteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OutpostsGetOutpostResult struct {
	Result workflow.Future
}

func (r *OutpostsGetOutpostResult) Get(ctx workflow.Context) (*outposts.GetOutpostOutput, error) {
    var output outposts.GetOutpostOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OutpostsGetOutpostInstanceTypesResult struct {
	Result workflow.Future
}

func (r *OutpostsGetOutpostInstanceTypesResult) Get(ctx workflow.Context) (*outposts.GetOutpostInstanceTypesOutput, error) {
    var output outposts.GetOutpostInstanceTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OutpostsListOutpostsResult struct {
	Result workflow.Future
}

func (r *OutpostsListOutpostsResult) Get(ctx workflow.Context) (*outposts.ListOutpostsOutput, error) {
    var output outposts.ListOutpostsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OutpostsListSitesResult struct {
	Result workflow.Future
}

func (r *OutpostsListSitesResult) Get(ctx workflow.Context) (*outposts.ListSitesOutput, error) {
    var output outposts.ListSitesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type OutpostsStub struct {
    activities OutpostsClient
}

func NewOutpostsStub() OutpostsClient {
    return &OutpostsStub{}
}

func (a *OutpostsStub) CreateOutpost(ctx workflow.Context, input *outposts.CreateOutpostInput) (*outposts.CreateOutpostOutput, error) {
    var output outposts.CreateOutpostOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOutpost, input).Get(ctx, &output)
    return &output, err
}

func (a *OutpostsStub) DeleteOutpost(ctx workflow.Context, input *outposts.DeleteOutpostInput) (*outposts.DeleteOutpostOutput, error) {
    var output outposts.DeleteOutpostOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOutpost, input).Get(ctx, &output)
    return &output, err
}

func (a *OutpostsStub) DeleteSite(ctx workflow.Context, input *outposts.DeleteSiteInput) (*outposts.DeleteSiteOutput, error) {
    var output outposts.DeleteSiteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSite, input).Get(ctx, &output)
    return &output, err
}

func (a *OutpostsStub) GetOutpost(ctx workflow.Context, input *outposts.GetOutpostInput) (*outposts.GetOutpostOutput, error) {
    var output outposts.GetOutpostOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOutpost, input).Get(ctx, &output)
    return &output, err
}

func (a *OutpostsStub) GetOutpostInstanceTypes(ctx workflow.Context, input *outposts.GetOutpostInstanceTypesInput) (*outposts.GetOutpostInstanceTypesOutput, error) {
    var output outposts.GetOutpostInstanceTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOutpostInstanceTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *OutpostsStub) ListOutposts(ctx workflow.Context, input *outposts.ListOutpostsInput) (*outposts.ListOutpostsOutput, error) {
    var output outposts.ListOutpostsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOutposts, input).Get(ctx, &output)
    return &output, err
}

func (a *OutpostsStub) ListSites(ctx workflow.Context, input *outposts.ListSitesInput) (*outposts.ListSitesOutput, error) {
    var output outposts.ListSitesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSites, input).Get(ctx, &output)
    return &output, err
}
