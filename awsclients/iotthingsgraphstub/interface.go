// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iotthingsgraphstub

import (
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	AssociateEntityToThing(ctx workflow.Context, input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error)
	AssociateEntityToThingAsync(ctx workflow.Context, input *iotthingsgraph.AssociateEntityToThingInput) *IoTThingsGraphAssociateEntityToThingFuture

	CreateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error)
	CreateFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.CreateFlowTemplateInput) *IoTThingsGraphCreateFlowTemplateFuture

	CreateSystemInstance(ctx workflow.Context, input *iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error)
	CreateSystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.CreateSystemInstanceInput) *IoTThingsGraphCreateSystemInstanceFuture

	CreateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error)
	CreateSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.CreateSystemTemplateInput) *IoTThingsGraphCreateSystemTemplateFuture

	DeleteFlowTemplate(ctx workflow.Context, input *iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error)
	DeleteFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeleteFlowTemplateInput) *IoTThingsGraphDeleteFlowTemplateFuture

	DeleteNamespace(ctx workflow.Context, input *iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error)
	DeleteNamespaceAsync(ctx workflow.Context, input *iotthingsgraph.DeleteNamespaceInput) *IoTThingsGraphDeleteNamespaceFuture

	DeleteSystemInstance(ctx workflow.Context, input *iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error)
	DeleteSystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.DeleteSystemInstanceInput) *IoTThingsGraphDeleteSystemInstanceFuture

	DeleteSystemTemplate(ctx workflow.Context, input *iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error)
	DeleteSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeleteSystemTemplateInput) *IoTThingsGraphDeleteSystemTemplateFuture

	DeploySystemInstance(ctx workflow.Context, input *iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error)
	DeploySystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.DeploySystemInstanceInput) *IoTThingsGraphDeploySystemInstanceFuture

	DeprecateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error)
	DeprecateFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) *IoTThingsGraphDeprecateFlowTemplateFuture

	DeprecateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error)
	DeprecateSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) *IoTThingsGraphDeprecateSystemTemplateFuture

	DescribeNamespace(ctx workflow.Context, input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error)
	DescribeNamespaceAsync(ctx workflow.Context, input *iotthingsgraph.DescribeNamespaceInput) *IoTThingsGraphDescribeNamespaceFuture

	DissociateEntityFromThing(ctx workflow.Context, input *iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error)
	DissociateEntityFromThingAsync(ctx workflow.Context, input *iotthingsgraph.DissociateEntityFromThingInput) *IoTThingsGraphDissociateEntityFromThingFuture

	GetEntities(ctx workflow.Context, input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error)
	GetEntitiesAsync(ctx workflow.Context, input *iotthingsgraph.GetEntitiesInput) *IoTThingsGraphGetEntitiesFuture

	GetFlowTemplate(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error)
	GetFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateInput) *IoTThingsGraphGetFlowTemplateFuture

	GetFlowTemplateRevisions(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error)
	GetFlowTemplateRevisionsAsync(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) *IoTThingsGraphGetFlowTemplateRevisionsFuture

	GetNamespaceDeletionStatus(ctx workflow.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error)
	GetNamespaceDeletionStatusAsync(ctx workflow.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) *IoTThingsGraphGetNamespaceDeletionStatusFuture

	GetSystemInstance(ctx workflow.Context, input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error)
	GetSystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.GetSystemInstanceInput) *IoTThingsGraphGetSystemInstanceFuture

	GetSystemTemplate(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error)
	GetSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateInput) *IoTThingsGraphGetSystemTemplateFuture

	GetSystemTemplateRevisions(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error)
	GetSystemTemplateRevisionsAsync(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) *IoTThingsGraphGetSystemTemplateRevisionsFuture

	GetUploadStatus(ctx workflow.Context, input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error)
	GetUploadStatusAsync(ctx workflow.Context, input *iotthingsgraph.GetUploadStatusInput) *IoTThingsGraphGetUploadStatusFuture

	ListFlowExecutionMessages(ctx workflow.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error)
	ListFlowExecutionMessagesAsync(ctx workflow.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) *IoTThingsGraphListFlowExecutionMessagesFuture

	ListTagsForResource(ctx workflow.Context, input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *iotthingsgraph.ListTagsForResourceInput) *IoTThingsGraphListTagsForResourceFuture

	SearchEntities(ctx workflow.Context, input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error)
	SearchEntitiesAsync(ctx workflow.Context, input *iotthingsgraph.SearchEntitiesInput) *IoTThingsGraphSearchEntitiesFuture

	SearchFlowExecutions(ctx workflow.Context, input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error)
	SearchFlowExecutionsAsync(ctx workflow.Context, input *iotthingsgraph.SearchFlowExecutionsInput) *IoTThingsGraphSearchFlowExecutionsFuture

	SearchFlowTemplates(ctx workflow.Context, input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error)
	SearchFlowTemplatesAsync(ctx workflow.Context, input *iotthingsgraph.SearchFlowTemplatesInput) *IoTThingsGraphSearchFlowTemplatesFuture

	SearchSystemInstances(ctx workflow.Context, input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error)
	SearchSystemInstancesAsync(ctx workflow.Context, input *iotthingsgraph.SearchSystemInstancesInput) *IoTThingsGraphSearchSystemInstancesFuture

	SearchSystemTemplates(ctx workflow.Context, input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error)
	SearchSystemTemplatesAsync(ctx workflow.Context, input *iotthingsgraph.SearchSystemTemplatesInput) *IoTThingsGraphSearchSystemTemplatesFuture

	SearchThings(ctx workflow.Context, input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error)
	SearchThingsAsync(ctx workflow.Context, input *iotthingsgraph.SearchThingsInput) *IoTThingsGraphSearchThingsFuture

	TagResource(ctx workflow.Context, input *iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *iotthingsgraph.TagResourceInput) *IoTThingsGraphTagResourceFuture

	UndeploySystemInstance(ctx workflow.Context, input *iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error)
	UndeploySystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.UndeploySystemInstanceInput) *IoTThingsGraphUndeploySystemInstanceFuture

	UntagResource(ctx workflow.Context, input *iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *iotthingsgraph.UntagResourceInput) *IoTThingsGraphUntagResourceFuture

	UpdateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error)
	UpdateFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.UpdateFlowTemplateInput) *IoTThingsGraphUpdateFlowTemplateFuture

	UpdateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error)
	UpdateSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.UpdateSystemTemplateInput) *IoTThingsGraphUpdateSystemTemplateFuture

	UploadEntityDefinitions(ctx workflow.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error)
	UploadEntityDefinitionsAsync(ctx workflow.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) *IoTThingsGraphUploadEntityDefinitionsFuture
}

func NewClient() Client {
	return &stub{}
}

