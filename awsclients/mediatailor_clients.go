package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MediaTailorClient interface {
    DeletePlaybackConfiguration(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error)
    DeletePlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) *MediatailorDeletePlaybackConfigurationResult

    GetPlaybackConfiguration(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error)
    GetPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) *MediatailorGetPlaybackConfigurationResult

    ListPlaybackConfigurations(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error)
    ListPlaybackConfigurationsAsync(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) *MediatailorListPlaybackConfigurationsResult

    ListTagsForResource(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) *MediatailorListTagsForResourceResult

    PutPlaybackConfiguration(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error)
    PutPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) *MediatailorPutPlaybackConfigurationResult

    TagResource(ctx workflow.Context, input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *mediatailor.TagResourceInput) *MediatailorTagResourceResult

    UntagResource(ctx workflow.Context, input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *mediatailor.UntagResourceInput) *MediatailorUntagResourceResult
}
type MediatailorDeletePlaybackConfigurationResult struct {
	Result workflow.Future
}

func (r *MediatailorDeletePlaybackConfigurationResult) Get(ctx workflow.Context) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
    var output mediatailor.DeletePlaybackConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediatailorGetPlaybackConfigurationResult struct {
	Result workflow.Future
}

func (r *MediatailorGetPlaybackConfigurationResult) Get(ctx workflow.Context) (*mediatailor.GetPlaybackConfigurationOutput, error) {
    var output mediatailor.GetPlaybackConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediatailorListPlaybackConfigurationsResult struct {
	Result workflow.Future
}

func (r *MediatailorListPlaybackConfigurationsResult) Get(ctx workflow.Context) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
    var output mediatailor.ListPlaybackConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediatailorListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MediatailorListTagsForResourceResult) Get(ctx workflow.Context) (*mediatailor.ListTagsForResourceOutput, error) {
    var output mediatailor.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediatailorPutPlaybackConfigurationResult struct {
	Result workflow.Future
}

func (r *MediatailorPutPlaybackConfigurationResult) Get(ctx workflow.Context) (*mediatailor.PutPlaybackConfigurationOutput, error) {
    var output mediatailor.PutPlaybackConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediatailorTagResourceResult struct {
	Result workflow.Future
}

func (r *MediatailorTagResourceResult) Get(ctx workflow.Context) (*mediatailor.TagResourceOutput, error) {
    var output mediatailor.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediatailorUntagResourceResult struct {
	Result workflow.Future
}

func (r *MediatailorUntagResourceResult) Get(ctx workflow.Context) (*mediatailor.UntagResourceOutput, error) {
    var output mediatailor.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MediaTailorStub struct {
    activities awsactivities.MediaTailorActivities
}

func NewMediaTailorStub() MediaTailorClient {
    return &MediaTailorStub{}
}
func (a *MediaTailorStub) DeletePlaybackConfiguration(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
    var output mediatailor.DeletePlaybackConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePlaybackConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaTailorStub) DeletePlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.DeletePlaybackConfigurationInput) *MediatailorDeletePlaybackConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePlaybackConfiguration, input)
    return &MediatailorDeletePlaybackConfigurationResult{Result: future}
}
func (a *MediaTailorStub) GetPlaybackConfiguration(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
    var output mediatailor.GetPlaybackConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPlaybackConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaTailorStub) GetPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.GetPlaybackConfigurationInput) *MediatailorGetPlaybackConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPlaybackConfiguration, input)
    return &MediatailorGetPlaybackConfigurationResult{Result: future}
}
func (a *MediaTailorStub) ListPlaybackConfigurations(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
    var output mediatailor.ListPlaybackConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPlaybackConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaTailorStub) ListPlaybackConfigurationsAsync(ctx workflow.Context, input *mediatailor.ListPlaybackConfigurationsInput) *MediatailorListPlaybackConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPlaybackConfigurations, input)
    return &MediatailorListPlaybackConfigurationsResult{Result: future}
}
func (a *MediaTailorStub) ListTagsForResource(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
    var output mediatailor.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaTailorStub) ListTagsForResourceAsync(ctx workflow.Context, input *mediatailor.ListTagsForResourceInput) *MediatailorListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &MediatailorListTagsForResourceResult{Result: future}
}
func (a *MediaTailorStub) PutPlaybackConfiguration(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error) {
    var output mediatailor.PutPlaybackConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPlaybackConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaTailorStub) PutPlaybackConfigurationAsync(ctx workflow.Context, input *mediatailor.PutPlaybackConfigurationInput) *MediatailorPutPlaybackConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutPlaybackConfiguration, input)
    return &MediatailorPutPlaybackConfigurationResult{Result: future}
}
func (a *MediaTailorStub) TagResource(ctx workflow.Context, input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error) {
    var output mediatailor.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaTailorStub) TagResourceAsync(ctx workflow.Context, input *mediatailor.TagResourceInput) *MediatailorTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &MediatailorTagResourceResult{Result: future}
}
func (a *MediaTailorStub) UntagResource(ctx workflow.Context, input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error) {
    var output mediatailor.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaTailorStub) UntagResourceAsync(ctx workflow.Context, input *mediatailor.UntagResourceInput) *MediatailorUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &MediatailorUntagResourceResult{Result: future}
}