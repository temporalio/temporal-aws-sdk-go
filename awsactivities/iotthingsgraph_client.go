package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"go.temporal.io/sdk/workflow"
)

type IoTThingsGraphClient interface {
    AssociateEntityToThing(ctx workflow.Context, input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error)
    AssociateEntityToThingAsync(ctx workflow.Context, input *iotthingsgraph.AssociateEntityToThingInput) *IotthingsgraphAssociateEntityToThingResult

    CreateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error)
    CreateFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.CreateFlowTemplateInput) *IotthingsgraphCreateFlowTemplateResult

    CreateSystemInstance(ctx workflow.Context, input *iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error)
    CreateSystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.CreateSystemInstanceInput) *IotthingsgraphCreateSystemInstanceResult

    CreateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error)
    CreateSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.CreateSystemTemplateInput) *IotthingsgraphCreateSystemTemplateResult

    DeleteFlowTemplate(ctx workflow.Context, input *iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error)
    DeleteFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeleteFlowTemplateInput) *IotthingsgraphDeleteFlowTemplateResult

    DeleteNamespace(ctx workflow.Context, input *iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error)
    DeleteNamespaceAsync(ctx workflow.Context, input *iotthingsgraph.DeleteNamespaceInput) *IotthingsgraphDeleteNamespaceResult

    DeleteSystemInstance(ctx workflow.Context, input *iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error)
    DeleteSystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.DeleteSystemInstanceInput) *IotthingsgraphDeleteSystemInstanceResult

    DeleteSystemTemplate(ctx workflow.Context, input *iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error)
    DeleteSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeleteSystemTemplateInput) *IotthingsgraphDeleteSystemTemplateResult

    DeploySystemInstance(ctx workflow.Context, input *iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error)
    DeploySystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.DeploySystemInstanceInput) *IotthingsgraphDeploySystemInstanceResult

    DeprecateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error)
    DeprecateFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) *IotthingsgraphDeprecateFlowTemplateResult

    DeprecateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error)
    DeprecateSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) *IotthingsgraphDeprecateSystemTemplateResult

    DescribeNamespace(ctx workflow.Context, input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error)
    DescribeNamespaceAsync(ctx workflow.Context, input *iotthingsgraph.DescribeNamespaceInput) *IotthingsgraphDescribeNamespaceResult

    DissociateEntityFromThing(ctx workflow.Context, input *iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error)
    DissociateEntityFromThingAsync(ctx workflow.Context, input *iotthingsgraph.DissociateEntityFromThingInput) *IotthingsgraphDissociateEntityFromThingResult

    GetEntities(ctx workflow.Context, input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error)
    GetEntitiesAsync(ctx workflow.Context, input *iotthingsgraph.GetEntitiesInput) *IotthingsgraphGetEntitiesResult

    GetFlowTemplate(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error)
    GetFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateInput) *IotthingsgraphGetFlowTemplateResult

    GetFlowTemplateRevisions(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error)
    GetFlowTemplateRevisionsAsync(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) *IotthingsgraphGetFlowTemplateRevisionsResult

    GetNamespaceDeletionStatus(ctx workflow.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error)
    GetNamespaceDeletionStatusAsync(ctx workflow.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) *IotthingsgraphGetNamespaceDeletionStatusResult

    GetSystemInstance(ctx workflow.Context, input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error)
    GetSystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.GetSystemInstanceInput) *IotthingsgraphGetSystemInstanceResult

    GetSystemTemplate(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error)
    GetSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateInput) *IotthingsgraphGetSystemTemplateResult

    GetSystemTemplateRevisions(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error)
    GetSystemTemplateRevisionsAsync(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) *IotthingsgraphGetSystemTemplateRevisionsResult

    GetUploadStatus(ctx workflow.Context, input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error)
    GetUploadStatusAsync(ctx workflow.Context, input *iotthingsgraph.GetUploadStatusInput) *IotthingsgraphGetUploadStatusResult

    ListFlowExecutionMessages(ctx workflow.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error)
    ListFlowExecutionMessagesAsync(ctx workflow.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) *IotthingsgraphListFlowExecutionMessagesResult

    ListTagsForResource(ctx workflow.Context, input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *iotthingsgraph.ListTagsForResourceInput) *IotthingsgraphListTagsForResourceResult

    SearchEntities(ctx workflow.Context, input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error)
    SearchEntitiesAsync(ctx workflow.Context, input *iotthingsgraph.SearchEntitiesInput) *IotthingsgraphSearchEntitiesResult

    SearchFlowExecutions(ctx workflow.Context, input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error)
    SearchFlowExecutionsAsync(ctx workflow.Context, input *iotthingsgraph.SearchFlowExecutionsInput) *IotthingsgraphSearchFlowExecutionsResult

    SearchFlowTemplates(ctx workflow.Context, input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error)
    SearchFlowTemplatesAsync(ctx workflow.Context, input *iotthingsgraph.SearchFlowTemplatesInput) *IotthingsgraphSearchFlowTemplatesResult

    SearchSystemInstances(ctx workflow.Context, input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error)
    SearchSystemInstancesAsync(ctx workflow.Context, input *iotthingsgraph.SearchSystemInstancesInput) *IotthingsgraphSearchSystemInstancesResult

    SearchSystemTemplates(ctx workflow.Context, input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error)
    SearchSystemTemplatesAsync(ctx workflow.Context, input *iotthingsgraph.SearchSystemTemplatesInput) *IotthingsgraphSearchSystemTemplatesResult

    SearchThings(ctx workflow.Context, input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error)
    SearchThingsAsync(ctx workflow.Context, input *iotthingsgraph.SearchThingsInput) *IotthingsgraphSearchThingsResult

    TagResource(ctx workflow.Context, input *iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *iotthingsgraph.TagResourceInput) *IotthingsgraphTagResourceResult

    UndeploySystemInstance(ctx workflow.Context, input *iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error)
    UndeploySystemInstanceAsync(ctx workflow.Context, input *iotthingsgraph.UndeploySystemInstanceInput) *IotthingsgraphUndeploySystemInstanceResult

    UntagResource(ctx workflow.Context, input *iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *iotthingsgraph.UntagResourceInput) *IotthingsgraphUntagResourceResult

    UpdateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error)
    UpdateFlowTemplateAsync(ctx workflow.Context, input *iotthingsgraph.UpdateFlowTemplateInput) *IotthingsgraphUpdateFlowTemplateResult

    UpdateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error)
    UpdateSystemTemplateAsync(ctx workflow.Context, input *iotthingsgraph.UpdateSystemTemplateInput) *IotthingsgraphUpdateSystemTemplateResult

    UploadEntityDefinitions(ctx workflow.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error)
    UploadEntityDefinitionsAsync(ctx workflow.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) *IotthingsgraphUploadEntityDefinitionsResult
}
type IotthingsgraphAssociateEntityToThingResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphAssociateEntityToThingResult) Get(ctx workflow.Context) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
    var output iotthingsgraph.AssociateEntityToThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphCreateFlowTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphCreateFlowTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.CreateFlowTemplateOutput, error) {
    var output iotthingsgraph.CreateFlowTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphCreateSystemInstanceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphCreateSystemInstanceResult) Get(ctx workflow.Context) (*iotthingsgraph.CreateSystemInstanceOutput, error) {
    var output iotthingsgraph.CreateSystemInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphCreateSystemTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphCreateSystemTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.CreateSystemTemplateOutput, error) {
    var output iotthingsgraph.CreateSystemTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDeleteFlowTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDeleteFlowTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.DeleteFlowTemplateOutput, error) {
    var output iotthingsgraph.DeleteFlowTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDeleteNamespaceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDeleteNamespaceResult) Get(ctx workflow.Context) (*iotthingsgraph.DeleteNamespaceOutput, error) {
    var output iotthingsgraph.DeleteNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDeleteSystemInstanceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDeleteSystemInstanceResult) Get(ctx workflow.Context) (*iotthingsgraph.DeleteSystemInstanceOutput, error) {
    var output iotthingsgraph.DeleteSystemInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDeleteSystemTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDeleteSystemTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.DeleteSystemTemplateOutput, error) {
    var output iotthingsgraph.DeleteSystemTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDeploySystemInstanceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDeploySystemInstanceResult) Get(ctx workflow.Context) (*iotthingsgraph.DeploySystemInstanceOutput, error) {
    var output iotthingsgraph.DeploySystemInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDeprecateFlowTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDeprecateFlowTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.DeprecateFlowTemplateOutput, error) {
    var output iotthingsgraph.DeprecateFlowTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDeprecateSystemTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDeprecateSystemTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.DeprecateSystemTemplateOutput, error) {
    var output iotthingsgraph.DeprecateSystemTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDescribeNamespaceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDescribeNamespaceResult) Get(ctx workflow.Context) (*iotthingsgraph.DescribeNamespaceOutput, error) {
    var output iotthingsgraph.DescribeNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphDissociateEntityFromThingResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphDissociateEntityFromThingResult) Get(ctx workflow.Context) (*iotthingsgraph.DissociateEntityFromThingOutput, error) {
    var output iotthingsgraph.DissociateEntityFromThingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetEntitiesResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetEntitiesResult) Get(ctx workflow.Context) (*iotthingsgraph.GetEntitiesOutput, error) {
    var output iotthingsgraph.GetEntitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetFlowTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetFlowTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.GetFlowTemplateOutput, error) {
    var output iotthingsgraph.GetFlowTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetFlowTemplateRevisionsResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetFlowTemplateRevisionsResult) Get(ctx workflow.Context) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
    var output iotthingsgraph.GetFlowTemplateRevisionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetNamespaceDeletionStatusResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetNamespaceDeletionStatusResult) Get(ctx workflow.Context) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
    var output iotthingsgraph.GetNamespaceDeletionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetSystemInstanceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetSystemInstanceResult) Get(ctx workflow.Context) (*iotthingsgraph.GetSystemInstanceOutput, error) {
    var output iotthingsgraph.GetSystemInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetSystemTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetSystemTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.GetSystemTemplateOutput, error) {
    var output iotthingsgraph.GetSystemTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetSystemTemplateRevisionsResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetSystemTemplateRevisionsResult) Get(ctx workflow.Context) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
    var output iotthingsgraph.GetSystemTemplateRevisionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphGetUploadStatusResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphGetUploadStatusResult) Get(ctx workflow.Context) (*iotthingsgraph.GetUploadStatusOutput, error) {
    var output iotthingsgraph.GetUploadStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphListFlowExecutionMessagesResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphListFlowExecutionMessagesResult) Get(ctx workflow.Context) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
    var output iotthingsgraph.ListFlowExecutionMessagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphListTagsForResourceResult) Get(ctx workflow.Context) (*iotthingsgraph.ListTagsForResourceOutput, error) {
    var output iotthingsgraph.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphSearchEntitiesResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphSearchEntitiesResult) Get(ctx workflow.Context) (*iotthingsgraph.SearchEntitiesOutput, error) {
    var output iotthingsgraph.SearchEntitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphSearchFlowExecutionsResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphSearchFlowExecutionsResult) Get(ctx workflow.Context) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
    var output iotthingsgraph.SearchFlowExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphSearchFlowTemplatesResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphSearchFlowTemplatesResult) Get(ctx workflow.Context) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
    var output iotthingsgraph.SearchFlowTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphSearchSystemInstancesResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphSearchSystemInstancesResult) Get(ctx workflow.Context) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
    var output iotthingsgraph.SearchSystemInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphSearchSystemTemplatesResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphSearchSystemTemplatesResult) Get(ctx workflow.Context) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
    var output iotthingsgraph.SearchSystemTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphSearchThingsResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphSearchThingsResult) Get(ctx workflow.Context) (*iotthingsgraph.SearchThingsOutput, error) {
    var output iotthingsgraph.SearchThingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphTagResourceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphTagResourceResult) Get(ctx workflow.Context) (*iotthingsgraph.TagResourceOutput, error) {
    var output iotthingsgraph.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphUndeploySystemInstanceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphUndeploySystemInstanceResult) Get(ctx workflow.Context) (*iotthingsgraph.UndeploySystemInstanceOutput, error) {
    var output iotthingsgraph.UndeploySystemInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphUntagResourceResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphUntagResourceResult) Get(ctx workflow.Context) (*iotthingsgraph.UntagResourceOutput, error) {
    var output iotthingsgraph.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphUpdateFlowTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphUpdateFlowTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.UpdateFlowTemplateOutput, error) {
    var output iotthingsgraph.UpdateFlowTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphUpdateSystemTemplateResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphUpdateSystemTemplateResult) Get(ctx workflow.Context) (*iotthingsgraph.UpdateSystemTemplateOutput, error) {
    var output iotthingsgraph.UpdateSystemTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotthingsgraphUploadEntityDefinitionsResult struct {
	Result workflow.Future
}

func (r *IotthingsgraphUploadEntityDefinitionsResult) Get(ctx workflow.Context) (*iotthingsgraph.UploadEntityDefinitionsOutput, error) {
    var output iotthingsgraph.UploadEntityDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type IoTThingsGraphStub struct {
    activities IoTThingsGraphClient
}

func NewIoTThingsGraphStub() IoTThingsGraphClient {
    return &IoTThingsGraphStub{}
}

func (a *IoTThingsGraphStub) AssociateEntityToThing(ctx workflow.Context, input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
    var output iotthingsgraph.AssociateEntityToThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateEntityToThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) CreateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error) {
    var output iotthingsgraph.CreateFlowTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFlowTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) CreateSystemInstance(ctx workflow.Context, input *iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error) {
    var output iotthingsgraph.CreateSystemInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSystemInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) CreateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error) {
    var output iotthingsgraph.CreateSystemTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSystemTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DeleteFlowTemplate(ctx workflow.Context, input *iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error) {
    var output iotthingsgraph.DeleteFlowTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFlowTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DeleteNamespace(ctx workflow.Context, input *iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error) {
    var output iotthingsgraph.DeleteNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DeleteSystemInstance(ctx workflow.Context, input *iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error) {
    var output iotthingsgraph.DeleteSystemInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSystemInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DeleteSystemTemplate(ctx workflow.Context, input *iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error) {
    var output iotthingsgraph.DeleteSystemTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSystemTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DeploySystemInstance(ctx workflow.Context, input *iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error) {
    var output iotthingsgraph.DeploySystemInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeploySystemInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DeprecateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error) {
    var output iotthingsgraph.DeprecateFlowTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeprecateFlowTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DeprecateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error) {
    var output iotthingsgraph.DeprecateSystemTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeprecateSystemTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DescribeNamespace(ctx workflow.Context, input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error) {
    var output iotthingsgraph.DescribeNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) DissociateEntityFromThing(ctx workflow.Context, input *iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error) {
    var output iotthingsgraph.DissociateEntityFromThingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DissociateEntityFromThing, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetEntities(ctx workflow.Context, input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error) {
    var output iotthingsgraph.GetEntitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEntities, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetFlowTemplate(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error) {
    var output iotthingsgraph.GetFlowTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFlowTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetFlowTemplateRevisions(ctx workflow.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
    var output iotthingsgraph.GetFlowTemplateRevisionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFlowTemplateRevisions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetNamespaceDeletionStatus(ctx workflow.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
    var output iotthingsgraph.GetNamespaceDeletionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNamespaceDeletionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetSystemInstance(ctx workflow.Context, input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error) {
    var output iotthingsgraph.GetSystemInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSystemInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetSystemTemplate(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error) {
    var output iotthingsgraph.GetSystemTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSystemTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetSystemTemplateRevisions(ctx workflow.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
    var output iotthingsgraph.GetSystemTemplateRevisionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSystemTemplateRevisions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) GetUploadStatus(ctx workflow.Context, input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error) {
    var output iotthingsgraph.GetUploadStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUploadStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) ListFlowExecutionMessages(ctx workflow.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
    var output iotthingsgraph.ListFlowExecutionMessagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFlowExecutionMessages, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) ListTagsForResource(ctx workflow.Context, input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error) {
    var output iotthingsgraph.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) SearchEntities(ctx workflow.Context, input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error) {
    var output iotthingsgraph.SearchEntitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchEntities, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) SearchFlowExecutions(ctx workflow.Context, input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
    var output iotthingsgraph.SearchFlowExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchFlowExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) SearchFlowTemplates(ctx workflow.Context, input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
    var output iotthingsgraph.SearchFlowTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchFlowTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) SearchSystemInstances(ctx workflow.Context, input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
    var output iotthingsgraph.SearchSystemInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchSystemInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) SearchSystemTemplates(ctx workflow.Context, input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
    var output iotthingsgraph.SearchSystemTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchSystemTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) SearchThings(ctx workflow.Context, input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error) {
    var output iotthingsgraph.SearchThingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchThings, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) TagResource(ctx workflow.Context, input *iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error) {
    var output iotthingsgraph.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) UndeploySystemInstance(ctx workflow.Context, input *iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error) {
    var output iotthingsgraph.UndeploySystemInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UndeploySystemInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) UntagResource(ctx workflow.Context, input *iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error) {
    var output iotthingsgraph.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) UpdateFlowTemplate(ctx workflow.Context, input *iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error) {
    var output iotthingsgraph.UpdateFlowTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFlowTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) UpdateSystemTemplate(ctx workflow.Context, input *iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error) {
    var output iotthingsgraph.UpdateSystemTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSystemTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTThingsGraphStub) UploadEntityDefinitions(ctx workflow.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error) {
    var output iotthingsgraph.UploadEntityDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadEntityDefinitions, input).Get(ctx, &output)
    return &output, err
}
