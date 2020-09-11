package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"go.temporal.io/sdk/workflow"
)

type IoTEventsDataClient interface {
       BatchPutMessage(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error)
       BatchPutMessageAsync(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) *IoteventsdataBatchPutMessageResult

       BatchUpdateDetector(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error)
       BatchUpdateDetectorAsync(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) *IoteventsdataBatchUpdateDetectorResult

       DescribeDetector(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error)
       DescribeDetectorAsync(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) *IoteventsdataDescribeDetectorResult

       ListDetectors(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error)
       ListDetectorsAsync(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) *IoteventsdataListDetectorsResult
}

type IoTEventsDataStub struct {
}

func NewIoTEventsDataStub() IoTEventsDataClient {
    return &IoTEventsDataStub{}
}

type IoteventsdataBatchPutMessageResult struct {
	Result workflow.Future
}

func (r *IoteventsdataBatchPutMessageResult) Get(ctx workflow.Context) (*ioteventsdata.BatchPutMessageOutput, error) {
    var output ioteventsdata.BatchPutMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsdataBatchUpdateDetectorResult struct {
	Result workflow.Future
}

func (r *IoteventsdataBatchUpdateDetectorResult) Get(ctx workflow.Context) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
    var output ioteventsdata.BatchUpdateDetectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsdataDescribeDetectorResult struct {
	Result workflow.Future
}

func (r *IoteventsdataDescribeDetectorResult) Get(ctx workflow.Context) (*ioteventsdata.DescribeDetectorOutput, error) {
    var output ioteventsdata.DescribeDetectorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type IoteventsdataListDetectorsResult struct {
	Result workflow.Future
}

func (r *IoteventsdataListDetectorsResult) Get(ctx workflow.Context) (*ioteventsdata.ListDetectorsOutput, error) {
    var output ioteventsdata.ListDetectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsDataStub) BatchPutMessage(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error) {
    var output ioteventsdata.BatchPutMessageOutput
    err := workflow.ExecuteActivity(ctx, "IoTEventsData.BatchPutMessage", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsDataStub) BatchPutMessageAsync(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) *IoteventsdataBatchPutMessageResult {
    future := workflow.ExecuteActivity(ctx, "IoTEventsData.BatchPutMessage", input)
    return &IoteventsdataBatchPutMessageResult{Result: future}
}

func (a *IoTEventsDataStub) BatchUpdateDetector(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
    var output ioteventsdata.BatchUpdateDetectorOutput
    err := workflow.ExecuteActivity(ctx, "IoTEventsData.BatchUpdateDetector", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsDataStub) BatchUpdateDetectorAsync(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) *IoteventsdataBatchUpdateDetectorResult {
    future := workflow.ExecuteActivity(ctx, "IoTEventsData.BatchUpdateDetector", input)
    return &IoteventsdataBatchUpdateDetectorResult{Result: future}
}

func (a *IoTEventsDataStub) DescribeDetector(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error) {
    var output ioteventsdata.DescribeDetectorOutput
    err := workflow.ExecuteActivity(ctx, "IoTEventsData.DescribeDetector", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsDataStub) DescribeDetectorAsync(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) *IoteventsdataDescribeDetectorResult {
    future := workflow.ExecuteActivity(ctx, "IoTEventsData.DescribeDetector", input)
    return &IoteventsdataDescribeDetectorResult{Result: future}
}

func (a *IoTEventsDataStub) ListDetectors(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error) {
    var output ioteventsdata.ListDetectorsOutput
    err := workflow.ExecuteActivity(ctx, "IoTEventsData.ListDetectors", input).Get(ctx, &output)
    return &output, err
}

func (a *IoTEventsDataStub) ListDetectorsAsync(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) *IoteventsdataListDetectorsResult {
    future := workflow.ExecuteActivity(ctx, "IoTEventsData.ListDetectors", input)
    return &IoteventsdataListDetectorsResult{Result: future}
}
