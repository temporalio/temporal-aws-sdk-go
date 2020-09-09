package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type IoT1ClickProjectsClient interface {
       AssociateDeviceWithPlacement(ctx workflow.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error)
       AssociateDeviceWithPlacementAsync(ctx workflow.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput) *Iot1clickprojectsAssociateDeviceWithPlacementResult

       CreatePlacement(ctx workflow.Context, input *iot1clickprojects.CreatePlacementInput) (*iot1clickprojects.CreatePlacementOutput, error)
       CreatePlacementAsync(ctx workflow.Context, input *iot1clickprojects.CreatePlacementInput) *Iot1clickprojectsCreatePlacementResult

       CreateProject(ctx workflow.Context, input *iot1clickprojects.CreateProjectInput) (*iot1clickprojects.CreateProjectOutput, error)
       CreateProjectAsync(ctx workflow.Context, input *iot1clickprojects.CreateProjectInput) *Iot1clickprojectsCreateProjectResult

       DeletePlacement(ctx workflow.Context, input *iot1clickprojects.DeletePlacementInput) (*iot1clickprojects.DeletePlacementOutput, error)
       DeletePlacementAsync(ctx workflow.Context, input *iot1clickprojects.DeletePlacementInput) *Iot1clickprojectsDeletePlacementResult

       DeleteProject(ctx workflow.Context, input *iot1clickprojects.DeleteProjectInput) (*iot1clickprojects.DeleteProjectOutput, error)
       DeleteProjectAsync(ctx workflow.Context, input *iot1clickprojects.DeleteProjectInput) *Iot1clickprojectsDeleteProjectResult

       DescribePlacement(ctx workflow.Context, input *iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error)
       DescribePlacementAsync(ctx workflow.Context, input *iot1clickprojects.DescribePlacementInput) *Iot1clickprojectsDescribePlacementResult

       DescribeProject(ctx workflow.Context, input *iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error)
       DescribeProjectAsync(ctx workflow.Context, input *iot1clickprojects.DescribeProjectInput) *Iot1clickprojectsDescribeProjectResult

       DisassociateDeviceFromPlacement(ctx workflow.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error)
       DisassociateDeviceFromPlacementAsync(ctx workflow.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput) *Iot1clickprojectsDisassociateDeviceFromPlacementResult

       GetDevicesInPlacement(ctx workflow.Context, input *iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error)
       GetDevicesInPlacementAsync(ctx workflow.Context, input *iot1clickprojects.GetDevicesInPlacementInput) *Iot1clickprojectsGetDevicesInPlacementResult

       ListPlacements(ctx workflow.Context, input *iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error)
       ListPlacementsAsync(ctx workflow.Context, input *iot1clickprojects.ListPlacementsInput) *Iot1clickprojectsListPlacementsResult

       ListProjects(ctx workflow.Context, input *iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error)
       ListProjectsAsync(ctx workflow.Context, input *iot1clickprojects.ListProjectsInput) *Iot1clickprojectsListProjectsResult

       ListTagsForResource(ctx workflow.Context, input *iot1clickprojects.ListTagsForResourceInput) (*iot1clickprojects.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *iot1clickprojects.ListTagsForResourceInput) *Iot1clickprojectsListTagsForResourceResult

       TagResource(ctx workflow.Context, input *iot1clickprojects.TagResourceInput) (*iot1clickprojects.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *iot1clickprojects.TagResourceInput) *Iot1clickprojectsTagResourceResult

       UntagResource(ctx workflow.Context, input *iot1clickprojects.UntagResourceInput) (*iot1clickprojects.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *iot1clickprojects.UntagResourceInput) *Iot1clickprojectsUntagResourceResult

       UpdatePlacement(ctx workflow.Context, input *iot1clickprojects.UpdatePlacementInput) (*iot1clickprojects.UpdatePlacementOutput, error)
       UpdatePlacementAsync(ctx workflow.Context, input *iot1clickprojects.UpdatePlacementInput) *Iot1clickprojectsUpdatePlacementResult

       UpdateProject(ctx workflow.Context, input *iot1clickprojects.UpdateProjectInput) (*iot1clickprojects.UpdateProjectOutput, error)
       UpdateProjectAsync(ctx workflow.Context, input *iot1clickprojects.UpdateProjectInput) *Iot1clickprojectsUpdateProjectResult
}

type Iot1clickprojectsAssociateDeviceWithPlacementResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsAssociateDeviceWithPlacementResult) Get(ctx workflow.Context) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
    var output iot1clickprojects.AssociateDeviceWithPlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsCreatePlacementResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsCreatePlacementResult) Get(ctx workflow.Context) (*iot1clickprojects.CreatePlacementOutput, error) {
    var output iot1clickprojects.CreatePlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsCreateProjectResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsCreateProjectResult) Get(ctx workflow.Context) (*iot1clickprojects.CreateProjectOutput, error) {
    var output iot1clickprojects.CreateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsDeletePlacementResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsDeletePlacementResult) Get(ctx workflow.Context) (*iot1clickprojects.DeletePlacementOutput, error) {
    var output iot1clickprojects.DeletePlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsDeleteProjectResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsDeleteProjectResult) Get(ctx workflow.Context) (*iot1clickprojects.DeleteProjectOutput, error) {
    var output iot1clickprojects.DeleteProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsDescribePlacementResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsDescribePlacementResult) Get(ctx workflow.Context) (*iot1clickprojects.DescribePlacementOutput, error) {
    var output iot1clickprojects.DescribePlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsDescribeProjectResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsDescribeProjectResult) Get(ctx workflow.Context) (*iot1clickprojects.DescribeProjectOutput, error) {
    var output iot1clickprojects.DescribeProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsDisassociateDeviceFromPlacementResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsDisassociateDeviceFromPlacementResult) Get(ctx workflow.Context) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
    var output iot1clickprojects.DisassociateDeviceFromPlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsGetDevicesInPlacementResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsGetDevicesInPlacementResult) Get(ctx workflow.Context) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
    var output iot1clickprojects.GetDevicesInPlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsListPlacementsResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsListPlacementsResult) Get(ctx workflow.Context) (*iot1clickprojects.ListPlacementsOutput, error) {
    var output iot1clickprojects.ListPlacementsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsListProjectsResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsListProjectsResult) Get(ctx workflow.Context) (*iot1clickprojects.ListProjectsOutput, error) {
    var output iot1clickprojects.ListProjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsListTagsForResourceResult) Get(ctx workflow.Context) (*iot1clickprojects.ListTagsForResourceOutput, error) {
    var output iot1clickprojects.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsTagResourceResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsTagResourceResult) Get(ctx workflow.Context) (*iot1clickprojects.TagResourceOutput, error) {
    var output iot1clickprojects.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsUntagResourceResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsUntagResourceResult) Get(ctx workflow.Context) (*iot1clickprojects.UntagResourceOutput, error) {
    var output iot1clickprojects.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsUpdatePlacementResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsUpdatePlacementResult) Get(ctx workflow.Context) (*iot1clickprojects.UpdatePlacementOutput, error) {
    var output iot1clickprojects.UpdatePlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Iot1clickprojectsUpdateProjectResult struct {
	Result workflow.Future
}

func (r *Iot1clickprojectsUpdateProjectResult) Get(ctx workflow.Context) (*iot1clickprojects.UpdateProjectOutput, error) {
    var output iot1clickprojects.UpdateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IoT1ClickProjectsStub struct {
    activities awsactivities.IoT1ClickProjectsActivities
}

func NewIoT1ClickProjectsStub() IoT1ClickProjectsClient {
    return &IoT1ClickProjectsStub{}
}

func (a *IoT1ClickProjectsStub) AssociateDeviceWithPlacement(ctx workflow.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
    var output iot1clickprojects.AssociateDeviceWithPlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateDeviceWithPlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) AssociateDeviceWithPlacementAsync(ctx workflow.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput) *Iot1clickprojectsAssociateDeviceWithPlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateDeviceWithPlacement, input)
    return &Iot1clickprojectsAssociateDeviceWithPlacementResult{Result: future}
}

func (a *IoT1ClickProjectsStub) CreatePlacement(ctx workflow.Context, input *iot1clickprojects.CreatePlacementInput) (*iot1clickprojects.CreatePlacementOutput, error) {
    var output iot1clickprojects.CreatePlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) CreatePlacementAsync(ctx workflow.Context, input *iot1clickprojects.CreatePlacementInput) *Iot1clickprojectsCreatePlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePlacement, input)
    return &Iot1clickprojectsCreatePlacementResult{Result: future}
}

func (a *IoT1ClickProjectsStub) CreateProject(ctx workflow.Context, input *iot1clickprojects.CreateProjectInput) (*iot1clickprojects.CreateProjectOutput, error) {
    var output iot1clickprojects.CreateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) CreateProjectAsync(ctx workflow.Context, input *iot1clickprojects.CreateProjectInput) *Iot1clickprojectsCreateProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input)
    return &Iot1clickprojectsCreateProjectResult{Result: future}
}

func (a *IoT1ClickProjectsStub) DeletePlacement(ctx workflow.Context, input *iot1clickprojects.DeletePlacementInput) (*iot1clickprojects.DeletePlacementOutput, error) {
    var output iot1clickprojects.DeletePlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) DeletePlacementAsync(ctx workflow.Context, input *iot1clickprojects.DeletePlacementInput) *Iot1clickprojectsDeletePlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePlacement, input)
    return &Iot1clickprojectsDeletePlacementResult{Result: future}
}

func (a *IoT1ClickProjectsStub) DeleteProject(ctx workflow.Context, input *iot1clickprojects.DeleteProjectInput) (*iot1clickprojects.DeleteProjectOutput, error) {
    var output iot1clickprojects.DeleteProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) DeleteProjectAsync(ctx workflow.Context, input *iot1clickprojects.DeleteProjectInput) *Iot1clickprojectsDeleteProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input)
    return &Iot1clickprojectsDeleteProjectResult{Result: future}
}

func (a *IoT1ClickProjectsStub) DescribePlacement(ctx workflow.Context, input *iot1clickprojects.DescribePlacementInput) (*iot1clickprojects.DescribePlacementOutput, error) {
    var output iot1clickprojects.DescribePlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) DescribePlacementAsync(ctx workflow.Context, input *iot1clickprojects.DescribePlacementInput) *Iot1clickprojectsDescribePlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePlacement, input)
    return &Iot1clickprojectsDescribePlacementResult{Result: future}
}

func (a *IoT1ClickProjectsStub) DescribeProject(ctx workflow.Context, input *iot1clickprojects.DescribeProjectInput) (*iot1clickprojects.DescribeProjectOutput, error) {
    var output iot1clickprojects.DescribeProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProject, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) DescribeProjectAsync(ctx workflow.Context, input *iot1clickprojects.DescribeProjectInput) *Iot1clickprojectsDescribeProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProject, input)
    return &Iot1clickprojectsDescribeProjectResult{Result: future}
}

func (a *IoT1ClickProjectsStub) DisassociateDeviceFromPlacement(ctx workflow.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
    var output iot1clickprojects.DisassociateDeviceFromPlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateDeviceFromPlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) DisassociateDeviceFromPlacementAsync(ctx workflow.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput) *Iot1clickprojectsDisassociateDeviceFromPlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateDeviceFromPlacement, input)
    return &Iot1clickprojectsDisassociateDeviceFromPlacementResult{Result: future}
}

func (a *IoT1ClickProjectsStub) GetDevicesInPlacement(ctx workflow.Context, input *iot1clickprojects.GetDevicesInPlacementInput) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
    var output iot1clickprojects.GetDevicesInPlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevicesInPlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) GetDevicesInPlacementAsync(ctx workflow.Context, input *iot1clickprojects.GetDevicesInPlacementInput) *Iot1clickprojectsGetDevicesInPlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDevicesInPlacement, input)
    return &Iot1clickprojectsGetDevicesInPlacementResult{Result: future}
}

func (a *IoT1ClickProjectsStub) ListPlacements(ctx workflow.Context, input *iot1clickprojects.ListPlacementsInput) (*iot1clickprojects.ListPlacementsOutput, error) {
    var output iot1clickprojects.ListPlacementsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPlacements, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) ListPlacementsAsync(ctx workflow.Context, input *iot1clickprojects.ListPlacementsInput) *Iot1clickprojectsListPlacementsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPlacements, input)
    return &Iot1clickprojectsListPlacementsResult{Result: future}
}

func (a *IoT1ClickProjectsStub) ListProjects(ctx workflow.Context, input *iot1clickprojects.ListProjectsInput) (*iot1clickprojects.ListProjectsOutput, error) {
    var output iot1clickprojects.ListProjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) ListProjectsAsync(ctx workflow.Context, input *iot1clickprojects.ListProjectsInput) *Iot1clickprojectsListProjectsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input)
    return &Iot1clickprojectsListProjectsResult{Result: future}
}

func (a *IoT1ClickProjectsStub) ListTagsForResource(ctx workflow.Context, input *iot1clickprojects.ListTagsForResourceInput) (*iot1clickprojects.ListTagsForResourceOutput, error) {
    var output iot1clickprojects.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) ListTagsForResourceAsync(ctx workflow.Context, input *iot1clickprojects.ListTagsForResourceInput) *Iot1clickprojectsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Iot1clickprojectsListTagsForResourceResult{Result: future}
}

func (a *IoT1ClickProjectsStub) TagResource(ctx workflow.Context, input *iot1clickprojects.TagResourceInput) (*iot1clickprojects.TagResourceOutput, error) {
    var output iot1clickprojects.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) TagResourceAsync(ctx workflow.Context, input *iot1clickprojects.TagResourceInput) *Iot1clickprojectsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Iot1clickprojectsTagResourceResult{Result: future}
}

func (a *IoT1ClickProjectsStub) UntagResource(ctx workflow.Context, input *iot1clickprojects.UntagResourceInput) (*iot1clickprojects.UntagResourceOutput, error) {
    var output iot1clickprojects.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) UntagResourceAsync(ctx workflow.Context, input *iot1clickprojects.UntagResourceInput) *Iot1clickprojectsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Iot1clickprojectsUntagResourceResult{Result: future}
}

func (a *IoT1ClickProjectsStub) UpdatePlacement(ctx workflow.Context, input *iot1clickprojects.UpdatePlacementInput) (*iot1clickprojects.UpdatePlacementOutput, error) {
    var output iot1clickprojects.UpdatePlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) UpdatePlacementAsync(ctx workflow.Context, input *iot1clickprojects.UpdatePlacementInput) *Iot1clickprojectsUpdatePlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePlacement, input)
    return &Iot1clickprojectsUpdatePlacementResult{Result: future}
}

func (a *IoT1ClickProjectsStub) UpdateProject(ctx workflow.Context, input *iot1clickprojects.UpdateProjectInput) (*iot1clickprojects.UpdateProjectOutput, error) {
    var output iot1clickprojects.UpdateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *IoT1ClickProjectsStub) UpdateProjectAsync(ctx workflow.Context, input *iot1clickprojects.UpdateProjectInput) *Iot1clickprojectsUpdateProjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input)
    return &Iot1clickprojectsUpdateProjectResult{Result: future}
}
