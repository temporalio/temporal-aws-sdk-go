
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/lexmodelbuildingserviceiface"
)

type LexModelBuildingServiceActivities struct {
	client lexmodelbuildingserviceiface.LexModelBuildingServiceAPI
}

func NewLexModelBuildingServiceActivities(client lexmodelbuildingserviceiface.LexModelBuildingServiceAPI) *LexModelBuildingServiceActivities {
    return &LexModelBuildingServiceActivities{client: client}
}

func (a *LexModelBuildingServiceActivities) CreateBotVersion(input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
    return a.client.CreateBotVersion(input)
}

func (a *LexModelBuildingServiceActivities) CreateIntentVersion(input *lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error) {
    return a.client.CreateIntentVersion(input)
}

func (a *LexModelBuildingServiceActivities) CreateSlotTypeVersion(input *lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error) {
    return a.client.CreateSlotTypeVersion(input)
}

func (a *LexModelBuildingServiceActivities) DeleteBot(input *lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error) {
    return a.client.DeleteBot(input)
}

func (a *LexModelBuildingServiceActivities) DeleteBotAlias(input *lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error) {
    return a.client.DeleteBotAlias(input)
}

func (a *LexModelBuildingServiceActivities) DeleteBotChannelAssociation(input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error) {
    return a.client.DeleteBotChannelAssociation(input)
}

func (a *LexModelBuildingServiceActivities) DeleteBotVersion(input *lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error) {
    return a.client.DeleteBotVersion(input)
}

func (a *LexModelBuildingServiceActivities) DeleteIntent(input *lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error) {
    return a.client.DeleteIntent(input)
}

func (a *LexModelBuildingServiceActivities) DeleteIntentVersion(input *lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error) {
    return a.client.DeleteIntentVersion(input)
}

func (a *LexModelBuildingServiceActivities) DeleteSlotType(input *lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error) {
    return a.client.DeleteSlotType(input)
}

func (a *LexModelBuildingServiceActivities) DeleteSlotTypeVersion(input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error) {
    return a.client.DeleteSlotTypeVersion(input)
}

func (a *LexModelBuildingServiceActivities) DeleteUtterances(input *lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error) {
    return a.client.DeleteUtterances(input)
}

func (a *LexModelBuildingServiceActivities) GetBot(input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error) {
    return a.client.GetBot(input)
}

func (a *LexModelBuildingServiceActivities) GetBotAlias(input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
    return a.client.GetBotAlias(input)
}

func (a *LexModelBuildingServiceActivities) GetBotAliases(input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
    return a.client.GetBotAliases(input)
}

func (a *LexModelBuildingServiceActivities) GetBotChannelAssociation(input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
    return a.client.GetBotChannelAssociation(input)
}

func (a *LexModelBuildingServiceActivities) GetBotChannelAssociations(input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
    return a.client.GetBotChannelAssociations(input)
}

func (a *LexModelBuildingServiceActivities) GetBotVersions(input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
    return a.client.GetBotVersions(input)
}

func (a *LexModelBuildingServiceActivities) GetBots(input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error) {
    return a.client.GetBots(input)
}

func (a *LexModelBuildingServiceActivities) GetBuiltinIntent(input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
    return a.client.GetBuiltinIntent(input)
}

func (a *LexModelBuildingServiceActivities) GetBuiltinIntents(input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
    return a.client.GetBuiltinIntents(input)
}

func (a *LexModelBuildingServiceActivities) GetBuiltinSlotTypes(input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
    return a.client.GetBuiltinSlotTypes(input)
}

func (a *LexModelBuildingServiceActivities) GetExport(input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error) {
    return a.client.GetExport(input)
}

func (a *LexModelBuildingServiceActivities) GetImport(input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error) {
    return a.client.GetImport(input)
}

func (a *LexModelBuildingServiceActivities) GetIntent(input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error) {
    return a.client.GetIntent(input)
}

func (a *LexModelBuildingServiceActivities) GetIntentVersions(input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
    return a.client.GetIntentVersions(input)
}

func (a *LexModelBuildingServiceActivities) GetIntents(input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error) {
    return a.client.GetIntents(input)
}

func (a *LexModelBuildingServiceActivities) GetSlotType(input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
    return a.client.GetSlotType(input)
}

func (a *LexModelBuildingServiceActivities) GetSlotTypeVersions(input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
    return a.client.GetSlotTypeVersions(input)
}

func (a *LexModelBuildingServiceActivities) GetSlotTypes(input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
    return a.client.GetSlotTypes(input)
}

func (a *LexModelBuildingServiceActivities) GetUtterancesView(input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
    return a.client.GetUtterancesView(input)
}

func (a *LexModelBuildingServiceActivities) ListTagsForResource(input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *LexModelBuildingServiceActivities) PutBot(input *lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error) {
    return a.client.PutBot(input)
}

func (a *LexModelBuildingServiceActivities) PutBotAlias(input *lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error) {
    return a.client.PutBotAlias(input)
}

func (a *LexModelBuildingServiceActivities) PutIntent(input *lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error) {
    return a.client.PutIntent(input)
}

func (a *LexModelBuildingServiceActivities) PutSlotType(input *lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error) {
    return a.client.PutSlotType(input)
}

func (a *LexModelBuildingServiceActivities) StartImport(input *lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error) {
    return a.client.StartImport(input)
}

func (a *LexModelBuildingServiceActivities) TagResource(input *lexmodelbuildingservice.TagResourceInput) (*lexmodelbuildingservice.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *LexModelBuildingServiceActivities) UntagResource(input *lexmodelbuildingservice.UntagResourceInput) (*lexmodelbuildingservice.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
