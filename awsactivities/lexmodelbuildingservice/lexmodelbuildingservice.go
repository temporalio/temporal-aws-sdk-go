// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package lexmodelbuildingservice

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/lexmodelbuildingserviceiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client lexmodelbuildingserviceiface.LexModelBuildingServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := lexmodelbuildingservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (lexmodelbuildingserviceiface.LexModelBuildingServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return lexmodelbuildingservice.New(sess), nil
}

func (a *Activities) CreateBotVersion(ctx context.Context, input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBotVersionWithContext(ctx, input)
}

func (a *Activities) CreateIntentVersion(ctx context.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateIntentVersionWithContext(ctx, input)
}

func (a *Activities) CreateSlotTypeVersion(ctx context.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSlotTypeVersionWithContext(ctx, input)
}

func (a *Activities) DeleteBot(ctx context.Context, input *lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBotWithContext(ctx, input)
}

func (a *Activities) DeleteBotAlias(ctx context.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBotAliasWithContext(ctx, input)
}

func (a *Activities) DeleteBotChannelAssociation(ctx context.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBotChannelAssociationWithContext(ctx, input)
}

func (a *Activities) DeleteBotVersion(ctx context.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBotVersionWithContext(ctx, input)
}

func (a *Activities) DeleteIntent(ctx context.Context, input *lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIntentWithContext(ctx, input)
}

func (a *Activities) DeleteIntentVersion(ctx context.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIntentVersionWithContext(ctx, input)
}

func (a *Activities) DeleteSlotType(ctx context.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSlotTypeWithContext(ctx, input)
}

func (a *Activities) DeleteSlotTypeVersion(ctx context.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSlotTypeVersionWithContext(ctx, input)
}

func (a *Activities) DeleteUtterances(ctx context.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUtterancesWithContext(ctx, input)
}

func (a *Activities) GetBot(ctx context.Context, input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBotWithContext(ctx, input)
}

func (a *Activities) GetBotAlias(ctx context.Context, input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBotAliasWithContext(ctx, input)
}

func (a *Activities) GetBotAliases(ctx context.Context, input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBotAliasesWithContext(ctx, input)
}

func (a *Activities) GetBotChannelAssociation(ctx context.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBotChannelAssociationWithContext(ctx, input)
}

func (a *Activities) GetBotChannelAssociations(ctx context.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBotChannelAssociationsWithContext(ctx, input)
}

func (a *Activities) GetBotVersions(ctx context.Context, input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBotVersionsWithContext(ctx, input)
}

func (a *Activities) GetBots(ctx context.Context, input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBotsWithContext(ctx, input)
}

func (a *Activities) GetBuiltinIntent(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBuiltinIntentWithContext(ctx, input)
}

func (a *Activities) GetBuiltinIntents(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBuiltinIntentsWithContext(ctx, input)
}

func (a *Activities) GetBuiltinSlotTypes(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBuiltinSlotTypesWithContext(ctx, input)
}

func (a *Activities) GetExport(ctx context.Context, input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetExportWithContext(ctx, input)
}

func (a *Activities) GetImport(ctx context.Context, input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetImportWithContext(ctx, input)
}

func (a *Activities) GetIntent(ctx context.Context, input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntentWithContext(ctx, input)
}

func (a *Activities) GetIntentVersions(ctx context.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntentVersionsWithContext(ctx, input)
}

func (a *Activities) GetIntents(ctx context.Context, input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntentsWithContext(ctx, input)
}

func (a *Activities) GetSlotType(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSlotTypeWithContext(ctx, input)
}

func (a *Activities) GetSlotTypeVersions(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSlotTypeVersionsWithContext(ctx, input)
}

func (a *Activities) GetSlotTypes(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSlotTypesWithContext(ctx, input)
}

func (a *Activities) GetUtterancesView(ctx context.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUtterancesViewWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) PutBot(ctx context.Context, input *lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBotWithContext(ctx, input)
}

func (a *Activities) PutBotAlias(ctx context.Context, input *lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBotAliasWithContext(ctx, input)
}

func (a *Activities) PutIntent(ctx context.Context, input *lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutIntentWithContext(ctx, input)
}

func (a *Activities) PutSlotType(ctx context.Context, input *lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutSlotTypeWithContext(ctx, input)
}

func (a *Activities) StartImport(ctx context.Context, input *lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartImportWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *lexmodelbuildingservice.TagResourceInput) (*lexmodelbuildingservice.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *lexmodelbuildingservice.UntagResourceInput) (*lexmodelbuildingservice.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}
