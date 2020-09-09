package awsclients

import (
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type LexModelBuildingServiceClient interface {
       CreateBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error)
       CreateBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) *LexmodelbuildingserviceCreateBotVersionResult

       CreateIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error)
       CreateIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) *LexmodelbuildingserviceCreateIntentVersionResult

       CreateSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error)
       CreateSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) *LexmodelbuildingserviceCreateSlotTypeVersionResult

       DeleteBot(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error)
       DeleteBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) *LexmodelbuildingserviceDeleteBotResult

       DeleteBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error)
       DeleteBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) *LexmodelbuildingserviceDeleteBotAliasResult

       DeleteBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error)
       DeleteBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) *LexmodelbuildingserviceDeleteBotChannelAssociationResult

       DeleteBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error)
       DeleteBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) *LexmodelbuildingserviceDeleteBotVersionResult

       DeleteIntent(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error)
       DeleteIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) *LexmodelbuildingserviceDeleteIntentResult

       DeleteIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error)
       DeleteIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) *LexmodelbuildingserviceDeleteIntentVersionResult

       DeleteSlotType(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error)
       DeleteSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) *LexmodelbuildingserviceDeleteSlotTypeResult

       DeleteSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error)
       DeleteSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) *LexmodelbuildingserviceDeleteSlotTypeVersionResult

       DeleteUtterances(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error)
       DeleteUtterancesAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) *LexmodelbuildingserviceDeleteUtterancesResult

       GetBot(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error)
       GetBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) *LexmodelbuildingserviceGetBotResult

       GetBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error)
       GetBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) *LexmodelbuildingserviceGetBotAliasResult

       GetBotAliases(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error)
       GetBotAliasesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) *LexmodelbuildingserviceGetBotAliasesResult

       GetBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error)
       GetBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) *LexmodelbuildingserviceGetBotChannelAssociationResult

       GetBotChannelAssociations(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error)
       GetBotChannelAssociationsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) *LexmodelbuildingserviceGetBotChannelAssociationsResult

       GetBotVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error)
       GetBotVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) *LexmodelbuildingserviceGetBotVersionsResult

       GetBots(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error)
       GetBotsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) *LexmodelbuildingserviceGetBotsResult

       GetBuiltinIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error)
       GetBuiltinIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) *LexmodelbuildingserviceGetBuiltinIntentResult

       GetBuiltinIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error)
       GetBuiltinIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) *LexmodelbuildingserviceGetBuiltinIntentsResult

       GetBuiltinSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error)
       GetBuiltinSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) *LexmodelbuildingserviceGetBuiltinSlotTypesResult

       GetExport(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error)
       GetExportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) *LexmodelbuildingserviceGetExportResult

       GetImport(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error)
       GetImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) *LexmodelbuildingserviceGetImportResult

       GetIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error)
       GetIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) *LexmodelbuildingserviceGetIntentResult

       GetIntentVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error)
       GetIntentVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) *LexmodelbuildingserviceGetIntentVersionsResult

       GetIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error)
       GetIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) *LexmodelbuildingserviceGetIntentsResult

       GetSlotType(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error)
       GetSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) *LexmodelbuildingserviceGetSlotTypeResult

       GetSlotTypeVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error)
       GetSlotTypeVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) *LexmodelbuildingserviceGetSlotTypeVersionsResult

       GetSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error)
       GetSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) *LexmodelbuildingserviceGetSlotTypesResult

       GetUtterancesView(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error)
       GetUtterancesViewAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) *LexmodelbuildingserviceGetUtterancesViewResult

       ListTagsForResource(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) *LexmodelbuildingserviceListTagsForResourceResult

       PutBot(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error)
       PutBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) *LexmodelbuildingservicePutBotResult

       PutBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error)
       PutBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) *LexmodelbuildingservicePutBotAliasResult

       PutIntent(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error)
       PutIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) *LexmodelbuildingservicePutIntentResult

       PutSlotType(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error)
       PutSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) *LexmodelbuildingservicePutSlotTypeResult

       StartImport(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error)
       StartImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) *LexmodelbuildingserviceStartImportResult

       TagResource(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) (*lexmodelbuildingservice.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) *LexmodelbuildingserviceTagResourceResult

       UntagResource(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) (*lexmodelbuildingservice.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) *LexmodelbuildingserviceUntagResourceResult
}

type LexmodelbuildingserviceCreateBotVersionResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceCreateBotVersionResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
    var output lexmodelbuildingservice.CreateBotVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceCreateIntentVersionResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceCreateIntentVersionResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.CreateIntentVersionOutput, error) {
    var output lexmodelbuildingservice.CreateIntentVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceCreateSlotTypeVersionResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceCreateSlotTypeVersionResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error) {
    var output lexmodelbuildingservice.CreateSlotTypeVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteBotResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteBotResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotOutput, error) {
    var output lexmodelbuildingservice.DeleteBotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteBotAliasResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteBotAliasResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotAliasOutput, error) {
    var output lexmodelbuildingservice.DeleteBotAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteBotChannelAssociationResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteBotChannelAssociationResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error) {
    var output lexmodelbuildingservice.DeleteBotChannelAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteBotVersionResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteBotVersionResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteBotVersionOutput, error) {
    var output lexmodelbuildingservice.DeleteBotVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteIntentResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteIntentResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteIntentOutput, error) {
    var output lexmodelbuildingservice.DeleteIntentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteIntentVersionResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteIntentVersionResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error) {
    var output lexmodelbuildingservice.DeleteIntentVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteSlotTypeResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteSlotTypeResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error) {
    var output lexmodelbuildingservice.DeleteSlotTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteSlotTypeVersionResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteSlotTypeVersionResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error) {
    var output lexmodelbuildingservice.DeleteSlotTypeVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceDeleteUtterancesResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceDeleteUtterancesResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.DeleteUtterancesOutput, error) {
    var output lexmodelbuildingservice.DeleteUtterancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBotResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBotResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotOutput, error) {
    var output lexmodelbuildingservice.GetBotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBotAliasResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBotAliasResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
    var output lexmodelbuildingservice.GetBotAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBotAliasesResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBotAliasesResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
    var output lexmodelbuildingservice.GetBotAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBotChannelAssociationResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBotChannelAssociationResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
    var output lexmodelbuildingservice.GetBotChannelAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBotChannelAssociationsResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBotChannelAssociationsResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
    var output lexmodelbuildingservice.GetBotChannelAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBotVersionsResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBotVersionsResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
    var output lexmodelbuildingservice.GetBotVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBotsResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBotsResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBotsOutput, error) {
    var output lexmodelbuildingservice.GetBotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBuiltinIntentResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBuiltinIntentResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
    var output lexmodelbuildingservice.GetBuiltinIntentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBuiltinIntentsResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBuiltinIntentsResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
    var output lexmodelbuildingservice.GetBuiltinIntentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetBuiltinSlotTypesResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetBuiltinSlotTypesResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
    var output lexmodelbuildingservice.GetBuiltinSlotTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetExportResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetExportResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetExportOutput, error) {
    var output lexmodelbuildingservice.GetExportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetImportResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetImportResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetImportOutput, error) {
    var output lexmodelbuildingservice.GetImportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetIntentResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetIntentResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetIntentOutput, error) {
    var output lexmodelbuildingservice.GetIntentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetIntentVersionsResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetIntentVersionsResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
    var output lexmodelbuildingservice.GetIntentVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetIntentsResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetIntentsResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetIntentsOutput, error) {
    var output lexmodelbuildingservice.GetIntentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetSlotTypeResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetSlotTypeResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
    var output lexmodelbuildingservice.GetSlotTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetSlotTypeVersionsResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetSlotTypeVersionsResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
    var output lexmodelbuildingservice.GetSlotTypeVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetSlotTypesResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetSlotTypesResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
    var output lexmodelbuildingservice.GetSlotTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceGetUtterancesViewResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceGetUtterancesViewResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
    var output lexmodelbuildingservice.GetUtterancesViewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceListTagsForResourceResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
    var output lexmodelbuildingservice.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingservicePutBotResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingservicePutBotResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutBotOutput, error) {
    var output lexmodelbuildingservice.PutBotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingservicePutBotAliasResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingservicePutBotAliasResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutBotAliasOutput, error) {
    var output lexmodelbuildingservice.PutBotAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingservicePutIntentResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingservicePutIntentResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutIntentOutput, error) {
    var output lexmodelbuildingservice.PutIntentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingservicePutSlotTypeResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingservicePutSlotTypeResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.PutSlotTypeOutput, error) {
    var output lexmodelbuildingservice.PutSlotTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceStartImportResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceStartImportResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.StartImportOutput, error) {
    var output lexmodelbuildingservice.StartImportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceTagResourceResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceTagResourceResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.TagResourceOutput, error) {
    var output lexmodelbuildingservice.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexmodelbuildingserviceUntagResourceResult struct {
	Result workflow.Future
}

func (r *LexmodelbuildingserviceUntagResourceResult) Get(ctx workflow.Context) (*lexmodelbuildingservice.UntagResourceOutput, error) {
    var output lexmodelbuildingservice.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexModelBuildingServiceStub struct {
    activities awsactivities.LexModelBuildingServiceActivities
}

func NewLexModelBuildingServiceStub() LexModelBuildingServiceClient {
    return &LexModelBuildingServiceStub{}
}

func (a *LexModelBuildingServiceStub) CreateBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
    var output lexmodelbuildingservice.CreateBotVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBotVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) CreateBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateBotVersionInput) *LexmodelbuildingserviceCreateBotVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBotVersion, input)
    return &LexmodelbuildingserviceCreateBotVersionResult{Result: future}
}

func (a *LexModelBuildingServiceStub) CreateIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error) {
    var output lexmodelbuildingservice.CreateIntentVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIntentVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) CreateIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) *LexmodelbuildingserviceCreateIntentVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateIntentVersion, input)
    return &LexmodelbuildingserviceCreateIntentVersionResult{Result: future}
}

func (a *LexModelBuildingServiceStub) CreateSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error) {
    var output lexmodelbuildingservice.CreateSlotTypeVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSlotTypeVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) CreateSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) *LexmodelbuildingserviceCreateSlotTypeVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSlotTypeVersion, input)
    return &LexmodelbuildingserviceCreateSlotTypeVersionResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteBot(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error) {
    var output lexmodelbuildingservice.DeleteBotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBot, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotInput) *LexmodelbuildingserviceDeleteBotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBot, input)
    return &LexmodelbuildingserviceDeleteBotResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error) {
    var output lexmodelbuildingservice.DeleteBotAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBotAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) *LexmodelbuildingserviceDeleteBotAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBotAlias, input)
    return &LexmodelbuildingserviceDeleteBotAliasResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error) {
    var output lexmodelbuildingservice.DeleteBotChannelAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBotChannelAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) *LexmodelbuildingserviceDeleteBotChannelAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBotChannelAssociation, input)
    return &LexmodelbuildingserviceDeleteBotChannelAssociationResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteBotVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error) {
    var output lexmodelbuildingservice.DeleteBotVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBotVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteBotVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) *LexmodelbuildingserviceDeleteBotVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBotVersion, input)
    return &LexmodelbuildingserviceDeleteBotVersionResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteIntent(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error) {
    var output lexmodelbuildingservice.DeleteIntentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIntent, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentInput) *LexmodelbuildingserviceDeleteIntentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteIntent, input)
    return &LexmodelbuildingserviceDeleteIntentResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteIntentVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error) {
    var output lexmodelbuildingservice.DeleteIntentVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIntentVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteIntentVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) *LexmodelbuildingserviceDeleteIntentVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteIntentVersion, input)
    return &LexmodelbuildingserviceDeleteIntentVersionResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteSlotType(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error) {
    var output lexmodelbuildingservice.DeleteSlotTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSlotType, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) *LexmodelbuildingserviceDeleteSlotTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSlotType, input)
    return &LexmodelbuildingserviceDeleteSlotTypeResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteSlotTypeVersion(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error) {
    var output lexmodelbuildingservice.DeleteSlotTypeVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSlotTypeVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteSlotTypeVersionAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) *LexmodelbuildingserviceDeleteSlotTypeVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSlotTypeVersion, input)
    return &LexmodelbuildingserviceDeleteSlotTypeVersionResult{Result: future}
}

func (a *LexModelBuildingServiceStub) DeleteUtterances(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error) {
    var output lexmodelbuildingservice.DeleteUtterancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUtterances, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) DeleteUtterancesAsync(ctx workflow.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) *LexmodelbuildingserviceDeleteUtterancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUtterances, input)
    return &LexmodelbuildingserviceDeleteUtterancesResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBot(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error) {
    var output lexmodelbuildingservice.GetBotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBot, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotInput) *LexmodelbuildingserviceGetBotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBot, input)
    return &LexmodelbuildingserviceGetBotResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
    var output lexmodelbuildingservice.GetBotAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBotAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasInput) *LexmodelbuildingserviceGetBotAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBotAlias, input)
    return &LexmodelbuildingserviceGetBotAliasResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBotAliases(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
    var output lexmodelbuildingservice.GetBotAliasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBotAliases, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBotAliasesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotAliasesInput) *LexmodelbuildingserviceGetBotAliasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBotAliases, input)
    return &LexmodelbuildingserviceGetBotAliasesResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBotChannelAssociation(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
    var output lexmodelbuildingservice.GetBotChannelAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBotChannelAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBotChannelAssociationAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) *LexmodelbuildingserviceGetBotChannelAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBotChannelAssociation, input)
    return &LexmodelbuildingserviceGetBotChannelAssociationResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBotChannelAssociations(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
    var output lexmodelbuildingservice.GetBotChannelAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBotChannelAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBotChannelAssociationsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) *LexmodelbuildingserviceGetBotChannelAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBotChannelAssociations, input)
    return &LexmodelbuildingserviceGetBotChannelAssociationsResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBotVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
    var output lexmodelbuildingservice.GetBotVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBotVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBotVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotVersionsInput) *LexmodelbuildingserviceGetBotVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBotVersions, input)
    return &LexmodelbuildingserviceGetBotVersionsResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBots(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error) {
    var output lexmodelbuildingservice.GetBotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBots, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBotsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBotsInput) *LexmodelbuildingserviceGetBotsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBots, input)
    return &LexmodelbuildingserviceGetBotsResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBuiltinIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
    var output lexmodelbuildingservice.GetBuiltinIntentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBuiltinIntent, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBuiltinIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) *LexmodelbuildingserviceGetBuiltinIntentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBuiltinIntent, input)
    return &LexmodelbuildingserviceGetBuiltinIntentResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBuiltinIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
    var output lexmodelbuildingservice.GetBuiltinIntentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBuiltinIntents, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBuiltinIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) *LexmodelbuildingserviceGetBuiltinIntentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBuiltinIntents, input)
    return &LexmodelbuildingserviceGetBuiltinIntentsResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetBuiltinSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
    var output lexmodelbuildingservice.GetBuiltinSlotTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBuiltinSlotTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetBuiltinSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) *LexmodelbuildingserviceGetBuiltinSlotTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBuiltinSlotTypes, input)
    return &LexmodelbuildingserviceGetBuiltinSlotTypesResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetExport(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error) {
    var output lexmodelbuildingservice.GetExportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetExport, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetExportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetExportInput) *LexmodelbuildingserviceGetExportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetExport, input)
    return &LexmodelbuildingserviceGetExportResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetImport(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error) {
    var output lexmodelbuildingservice.GetImportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImport, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetImportInput) *LexmodelbuildingserviceGetImportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetImport, input)
    return &LexmodelbuildingserviceGetImportResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetIntent(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error) {
    var output lexmodelbuildingservice.GetIntentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIntent, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentInput) *LexmodelbuildingserviceGetIntentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetIntent, input)
    return &LexmodelbuildingserviceGetIntentResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetIntentVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
    var output lexmodelbuildingservice.GetIntentVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIntentVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetIntentVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) *LexmodelbuildingserviceGetIntentVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetIntentVersions, input)
    return &LexmodelbuildingserviceGetIntentVersionsResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetIntents(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error) {
    var output lexmodelbuildingservice.GetIntentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIntents, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetIntentsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetIntentsInput) *LexmodelbuildingserviceGetIntentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetIntents, input)
    return &LexmodelbuildingserviceGetIntentsResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetSlotType(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
    var output lexmodelbuildingservice.GetSlotTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSlotType, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeInput) *LexmodelbuildingserviceGetSlotTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSlotType, input)
    return &LexmodelbuildingserviceGetSlotTypeResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetSlotTypeVersions(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
    var output lexmodelbuildingservice.GetSlotTypeVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSlotTypeVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetSlotTypeVersionsAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) *LexmodelbuildingserviceGetSlotTypeVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSlotTypeVersions, input)
    return &LexmodelbuildingserviceGetSlotTypeVersionsResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetSlotTypes(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
    var output lexmodelbuildingservice.GetSlotTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSlotTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetSlotTypesAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetSlotTypesInput) *LexmodelbuildingserviceGetSlotTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSlotTypes, input)
    return &LexmodelbuildingserviceGetSlotTypesResult{Result: future}
}

func (a *LexModelBuildingServiceStub) GetUtterancesView(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
    var output lexmodelbuildingservice.GetUtterancesViewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUtterancesView, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) GetUtterancesViewAsync(ctx workflow.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) *LexmodelbuildingserviceGetUtterancesViewResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUtterancesView, input)
    return &LexmodelbuildingserviceGetUtterancesViewResult{Result: future}
}

func (a *LexModelBuildingServiceStub) ListTagsForResource(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
    var output lexmodelbuildingservice.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) ListTagsForResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) *LexmodelbuildingserviceListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &LexmodelbuildingserviceListTagsForResourceResult{Result: future}
}

func (a *LexModelBuildingServiceStub) PutBot(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error) {
    var output lexmodelbuildingservice.PutBotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutBot, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) PutBotAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotInput) *LexmodelbuildingservicePutBotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutBot, input)
    return &LexmodelbuildingservicePutBotResult{Result: future}
}

func (a *LexModelBuildingServiceStub) PutBotAlias(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error) {
    var output lexmodelbuildingservice.PutBotAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutBotAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) PutBotAliasAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutBotAliasInput) *LexmodelbuildingservicePutBotAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutBotAlias, input)
    return &LexmodelbuildingservicePutBotAliasResult{Result: future}
}

func (a *LexModelBuildingServiceStub) PutIntent(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error) {
    var output lexmodelbuildingservice.PutIntentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutIntent, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) PutIntentAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutIntentInput) *LexmodelbuildingservicePutIntentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutIntent, input)
    return &LexmodelbuildingservicePutIntentResult{Result: future}
}

func (a *LexModelBuildingServiceStub) PutSlotType(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error) {
    var output lexmodelbuildingservice.PutSlotTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSlotType, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) PutSlotTypeAsync(ctx workflow.Context, input *lexmodelbuildingservice.PutSlotTypeInput) *LexmodelbuildingservicePutSlotTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutSlotType, input)
    return &LexmodelbuildingservicePutSlotTypeResult{Result: future}
}

func (a *LexModelBuildingServiceStub) StartImport(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error) {
    var output lexmodelbuildingservice.StartImportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartImport, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) StartImportAsync(ctx workflow.Context, input *lexmodelbuildingservice.StartImportInput) *LexmodelbuildingserviceStartImportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartImport, input)
    return &LexmodelbuildingserviceStartImportResult{Result: future}
}

func (a *LexModelBuildingServiceStub) TagResource(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) (*lexmodelbuildingservice.TagResourceOutput, error) {
    var output lexmodelbuildingservice.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) TagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.TagResourceInput) *LexmodelbuildingserviceTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &LexmodelbuildingserviceTagResourceResult{Result: future}
}

func (a *LexModelBuildingServiceStub) UntagResource(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) (*lexmodelbuildingservice.UntagResourceOutput, error) {
    var output lexmodelbuildingservice.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LexModelBuildingServiceStub) UntagResourceAsync(ctx workflow.Context, input *lexmodelbuildingservice.UntagResourceInput) *LexmodelbuildingserviceUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &LexmodelbuildingserviceUntagResourceResult{Result: future}
}
