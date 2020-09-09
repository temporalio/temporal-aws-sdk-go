package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type IoTDataPlaneClient interface {
	DeleteThingShadow(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error)
	DeleteThingShadowAsync(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) *IotdataplaneDeleteThingShadowResult

	GetThingShadow(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error)
	GetThingShadowAsync(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) *IotdataplaneGetThingShadowResult

	ListNamedShadowsForThing(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error)
	ListNamedShadowsForThingAsync(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) *IotdataplaneListNamedShadowsForThingResult

	Publish(ctx workflow.Context, input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error)
	PublishAsync(ctx workflow.Context, input *iotdataplane.PublishInput) *IotdataplanePublishResult

	UpdateThingShadow(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error)
	UpdateThingShadowAsync(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) *IotdataplaneUpdateThingShadowResult
}

type IotdataplaneDeleteThingShadowResult struct {
	Result workflow.Future
}

func (r *IotdataplaneDeleteThingShadowResult) Get(ctx workflow.Context) (*iotdataplane.DeleteThingShadowOutput, error) {
	var output iotdataplane.DeleteThingShadowOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotdataplaneGetThingShadowResult struct {
	Result workflow.Future
}

func (r *IotdataplaneGetThingShadowResult) Get(ctx workflow.Context) (*iotdataplane.GetThingShadowOutput, error) {
	var output iotdataplane.GetThingShadowOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotdataplaneListNamedShadowsForThingResult struct {
	Result workflow.Future
}

func (r *IotdataplaneListNamedShadowsForThingResult) Get(ctx workflow.Context) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	var output iotdataplane.ListNamedShadowsForThingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotdataplanePublishResult struct {
	Result workflow.Future
}

func (r *IotdataplanePublishResult) Get(ctx workflow.Context) (*iotdataplane.PublishOutput, error) {
	var output iotdataplane.PublishOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotdataplaneUpdateThingShadowResult struct {
	Result workflow.Future
}

func (r *IotdataplaneUpdateThingShadowResult) Get(ctx workflow.Context) (*iotdataplane.UpdateThingShadowOutput, error) {
	var output iotdataplane.UpdateThingShadowOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IoTDataPlaneStub struct {
	activities awsactivities.IoTDataPlaneActivities
}

func NewIoTDataPlaneStub() IoTDataPlaneClient {
	return &IoTDataPlaneStub{}
}

func (a *IoTDataPlaneStub) DeleteThingShadow(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
	var output iotdataplane.DeleteThingShadowOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteThingShadow, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTDataPlaneStub) DeleteThingShadowAsync(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) *IotdataplaneDeleteThingShadowResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteThingShadow, input)
	return &IotdataplaneDeleteThingShadowResult{Result: future}
}

func (a *IoTDataPlaneStub) GetThingShadow(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
	var output iotdataplane.GetThingShadowOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetThingShadow, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTDataPlaneStub) GetThingShadowAsync(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) *IotdataplaneGetThingShadowResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetThingShadow, input)
	return &IotdataplaneGetThingShadowResult{Result: future}
}

func (a *IoTDataPlaneStub) ListNamedShadowsForThing(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	var output iotdataplane.ListNamedShadowsForThingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListNamedShadowsForThing, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTDataPlaneStub) ListNamedShadowsForThingAsync(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) *IotdataplaneListNamedShadowsForThingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListNamedShadowsForThing, input)
	return &IotdataplaneListNamedShadowsForThingResult{Result: future}
}

func (a *IoTDataPlaneStub) Publish(ctx workflow.Context, input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {
	var output iotdataplane.PublishOutput
	err := workflow.ExecuteActivity(ctx, a.activities.Publish, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTDataPlaneStub) PublishAsync(ctx workflow.Context, input *iotdataplane.PublishInput) *IotdataplanePublishResult {
	future := workflow.ExecuteActivity(ctx, a.activities.Publish, input)
	return &IotdataplanePublishResult{Result: future}
}

func (a *IoTDataPlaneStub) UpdateThingShadow(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error) {
	var output iotdataplane.UpdateThingShadowOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateThingShadow, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTDataPlaneStub) UpdateThingShadowAsync(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) *IotdataplaneUpdateThingShadowResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateThingShadow, input)
	return &IotdataplaneUpdateThingShadowResult{Result: future}
}
