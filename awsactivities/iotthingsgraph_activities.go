package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph/iotthingsgraphiface"
)

type IoTThingsGraphActivities struct {
	client iotthingsgraphiface.IoTThingsGraphAPI
}

func NewIoTThingsGraphActivities(client iotthingsgraphiface.IoTThingsGraphAPI) *IoTThingsGraphActivities {
    return &IoTThingsGraphActivities{client: client}
}

func (a *IoTThingsGraphActivities) AssociateEntityToThing(input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
    return a.client.AssociateEntityToThing(input)
}

func (a *IoTThingsGraphActivities) CreateFlowTemplate(input *iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error) {
    return a.client.CreateFlowTemplate(input)
}

func (a *IoTThingsGraphActivities) CreateSystemInstance(input *iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error) {
    return a.client.CreateSystemInstance(input)
}

func (a *IoTThingsGraphActivities) CreateSystemTemplate(input *iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error) {
    return a.client.CreateSystemTemplate(input)
}

func (a *IoTThingsGraphActivities) DeleteFlowTemplate(input *iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error) {
    return a.client.DeleteFlowTemplate(input)
}

func (a *IoTThingsGraphActivities) DeleteNamespace(input *iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error) {
    return a.client.DeleteNamespace(input)
}

func (a *IoTThingsGraphActivities) DeleteSystemInstance(input *iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error) {
    return a.client.DeleteSystemInstance(input)
}

func (a *IoTThingsGraphActivities) DeleteSystemTemplate(input *iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error) {
    return a.client.DeleteSystemTemplate(input)
}

func (a *IoTThingsGraphActivities) DeploySystemInstance(input *iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error) {
    return a.client.DeploySystemInstance(input)
}

func (a *IoTThingsGraphActivities) DeprecateFlowTemplate(input *iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error) {
    return a.client.DeprecateFlowTemplate(input)
}

func (a *IoTThingsGraphActivities) DeprecateSystemTemplate(input *iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error) {
    return a.client.DeprecateSystemTemplate(input)
}

func (a *IoTThingsGraphActivities) DescribeNamespace(input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error) {
    return a.client.DescribeNamespace(input)
}

func (a *IoTThingsGraphActivities) DissociateEntityFromThing(input *iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error) {
    return a.client.DissociateEntityFromThing(input)
}

func (a *IoTThingsGraphActivities) GetEntities(input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error) {
    return a.client.GetEntities(input)
}

func (a *IoTThingsGraphActivities) GetFlowTemplate(input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error) {
    return a.client.GetFlowTemplate(input)
}

func (a *IoTThingsGraphActivities) GetFlowTemplateRevisions(input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
    return a.client.GetFlowTemplateRevisions(input)
}

func (a *IoTThingsGraphActivities) GetNamespaceDeletionStatus(input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
    return a.client.GetNamespaceDeletionStatus(input)
}

func (a *IoTThingsGraphActivities) GetSystemInstance(input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error) {
    return a.client.GetSystemInstance(input)
}

func (a *IoTThingsGraphActivities) GetSystemTemplate(input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error) {
    return a.client.GetSystemTemplate(input)
}

func (a *IoTThingsGraphActivities) GetSystemTemplateRevisions(input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
    return a.client.GetSystemTemplateRevisions(input)
}

func (a *IoTThingsGraphActivities) GetUploadStatus(input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error) {
    return a.client.GetUploadStatus(input)
}

func (a *IoTThingsGraphActivities) ListFlowExecutionMessages(input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
    return a.client.ListFlowExecutionMessages(input)
}

func (a *IoTThingsGraphActivities) ListTagsForResource(input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *IoTThingsGraphActivities) SearchEntities(input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error) {
    return a.client.SearchEntities(input)
}

func (a *IoTThingsGraphActivities) SearchFlowExecutions(input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
    return a.client.SearchFlowExecutions(input)
}

func (a *IoTThingsGraphActivities) SearchFlowTemplates(input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
    return a.client.SearchFlowTemplates(input)
}

func (a *IoTThingsGraphActivities) SearchSystemInstances(input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
    return a.client.SearchSystemInstances(input)
}

func (a *IoTThingsGraphActivities) SearchSystemTemplates(input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
    return a.client.SearchSystemTemplates(input)
}

func (a *IoTThingsGraphActivities) SearchThings(input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error) {
    return a.client.SearchThings(input)
}

func (a *IoTThingsGraphActivities) TagResource(input *iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *IoTThingsGraphActivities) UndeploySystemInstance(input *iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error) {
    return a.client.UndeploySystemInstance(input)
}

func (a *IoTThingsGraphActivities) UntagResource(input *iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *IoTThingsGraphActivities) UpdateFlowTemplate(input *iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error) {
    return a.client.UpdateFlowTemplate(input)
}

func (a *IoTThingsGraphActivities) UpdateSystemTemplate(input *iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error) {
    return a.client.UpdateSystemTemplate(input)
}

func (a *IoTThingsGraphActivities) UploadEntityDefinitions(input *iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error) {
    return a.client.UploadEntityDefinitions(input)
}
