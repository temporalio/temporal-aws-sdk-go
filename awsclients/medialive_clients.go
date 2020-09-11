package awsclients

import (
	"github.com/aws/aws-sdk-go/service/medialive"
	"go.temporal.io/sdk/workflow"
)

type MediaLiveClient interface {
       BatchUpdateSchedule(ctx workflow.Context, input *medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error)
       BatchUpdateScheduleAsync(ctx workflow.Context, input *medialive.BatchUpdateScheduleInput) *MedialiveBatchUpdateScheduleResult

       CreateChannel(ctx workflow.Context, input *medialive.CreateChannelInput) (*medialive.CreateChannelOutput, error)
       CreateChannelAsync(ctx workflow.Context, input *medialive.CreateChannelInput) *MedialiveCreateChannelResult

       CreateInput(ctx workflow.Context, input *medialive.CreateInputInput) (*medialive.CreateInputOutput, error)
       CreateInputAsync(ctx workflow.Context, input *medialive.CreateInputInput) *MedialiveCreateInputResult

       CreateInputSecurityGroup(ctx workflow.Context, input *medialive.CreateInputSecurityGroupInput) (*medialive.CreateInputSecurityGroupOutput, error)
       CreateInputSecurityGroupAsync(ctx workflow.Context, input *medialive.CreateInputSecurityGroupInput) *MedialiveCreateInputSecurityGroupResult

       CreateMultiplex(ctx workflow.Context, input *medialive.CreateMultiplexInput) (*medialive.CreateMultiplexOutput, error)
       CreateMultiplexAsync(ctx workflow.Context, input *medialive.CreateMultiplexInput) *MedialiveCreateMultiplexResult

       CreateMultiplexProgram(ctx workflow.Context, input *medialive.CreateMultiplexProgramInput) (*medialive.CreateMultiplexProgramOutput, error)
       CreateMultiplexProgramAsync(ctx workflow.Context, input *medialive.CreateMultiplexProgramInput) *MedialiveCreateMultiplexProgramResult

       CreateTags(ctx workflow.Context, input *medialive.CreateTagsInput) (*medialive.CreateTagsOutput, error)
       CreateTagsAsync(ctx workflow.Context, input *medialive.CreateTagsInput) *MedialiveCreateTagsResult

       DeleteChannel(ctx workflow.Context, input *medialive.DeleteChannelInput) (*medialive.DeleteChannelOutput, error)
       DeleteChannelAsync(ctx workflow.Context, input *medialive.DeleteChannelInput) *MedialiveDeleteChannelResult

       DeleteInput(ctx workflow.Context, input *medialive.DeleteInputInput) (*medialive.DeleteInputOutput, error)
       DeleteInputAsync(ctx workflow.Context, input *medialive.DeleteInputInput) *MedialiveDeleteInputResult

       DeleteInputSecurityGroup(ctx workflow.Context, input *medialive.DeleteInputSecurityGroupInput) (*medialive.DeleteInputSecurityGroupOutput, error)
       DeleteInputSecurityGroupAsync(ctx workflow.Context, input *medialive.DeleteInputSecurityGroupInput) *MedialiveDeleteInputSecurityGroupResult

       DeleteMultiplex(ctx workflow.Context, input *medialive.DeleteMultiplexInput) (*medialive.DeleteMultiplexOutput, error)
       DeleteMultiplexAsync(ctx workflow.Context, input *medialive.DeleteMultiplexInput) *MedialiveDeleteMultiplexResult

       DeleteMultiplexProgram(ctx workflow.Context, input *medialive.DeleteMultiplexProgramInput) (*medialive.DeleteMultiplexProgramOutput, error)
       DeleteMultiplexProgramAsync(ctx workflow.Context, input *medialive.DeleteMultiplexProgramInput) *MedialiveDeleteMultiplexProgramResult

       DeleteReservation(ctx workflow.Context, input *medialive.DeleteReservationInput) (*medialive.DeleteReservationOutput, error)
       DeleteReservationAsync(ctx workflow.Context, input *medialive.DeleteReservationInput) *MedialiveDeleteReservationResult

       DeleteSchedule(ctx workflow.Context, input *medialive.DeleteScheduleInput) (*medialive.DeleteScheduleOutput, error)
       DeleteScheduleAsync(ctx workflow.Context, input *medialive.DeleteScheduleInput) *MedialiveDeleteScheduleResult

       DeleteTags(ctx workflow.Context, input *medialive.DeleteTagsInput) (*medialive.DeleteTagsOutput, error)
       DeleteTagsAsync(ctx workflow.Context, input *medialive.DeleteTagsInput) *MedialiveDeleteTagsResult

       DescribeChannel(ctx workflow.Context, input *medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error)
       DescribeChannelAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) *MedialiveDescribeChannelResult

       DescribeInput(ctx workflow.Context, input *medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error)
       DescribeInputAsync(ctx workflow.Context, input *medialive.DescribeInputInput) *MedialiveDescribeInputResult

       DescribeInputDevice(ctx workflow.Context, input *medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error)
       DescribeInputDeviceAsync(ctx workflow.Context, input *medialive.DescribeInputDeviceInput) *MedialiveDescribeInputDeviceResult

       DescribeInputDeviceThumbnail(ctx workflow.Context, input *medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error)
       DescribeInputDeviceThumbnailAsync(ctx workflow.Context, input *medialive.DescribeInputDeviceThumbnailInput) *MedialiveDescribeInputDeviceThumbnailResult

       DescribeInputSecurityGroup(ctx workflow.Context, input *medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error)
       DescribeInputSecurityGroupAsync(ctx workflow.Context, input *medialive.DescribeInputSecurityGroupInput) *MedialiveDescribeInputSecurityGroupResult

       DescribeMultiplex(ctx workflow.Context, input *medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error)
       DescribeMultiplexAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) *MedialiveDescribeMultiplexResult

       DescribeMultiplexProgram(ctx workflow.Context, input *medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error)
       DescribeMultiplexProgramAsync(ctx workflow.Context, input *medialive.DescribeMultiplexProgramInput) *MedialiveDescribeMultiplexProgramResult

       DescribeOffering(ctx workflow.Context, input *medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error)
       DescribeOfferingAsync(ctx workflow.Context, input *medialive.DescribeOfferingInput) *MedialiveDescribeOfferingResult

       DescribeReservation(ctx workflow.Context, input *medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error)
       DescribeReservationAsync(ctx workflow.Context, input *medialive.DescribeReservationInput) *MedialiveDescribeReservationResult

       DescribeSchedule(ctx workflow.Context, input *medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error)
       DescribeScheduleAsync(ctx workflow.Context, input *medialive.DescribeScheduleInput) *MedialiveDescribeScheduleResult

       ListChannels(ctx workflow.Context, input *medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error)
       ListChannelsAsync(ctx workflow.Context, input *medialive.ListChannelsInput) *MedialiveListChannelsResult

       ListInputDevices(ctx workflow.Context, input *medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error)
       ListInputDevicesAsync(ctx workflow.Context, input *medialive.ListInputDevicesInput) *MedialiveListInputDevicesResult

       ListInputSecurityGroups(ctx workflow.Context, input *medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error)
       ListInputSecurityGroupsAsync(ctx workflow.Context, input *medialive.ListInputSecurityGroupsInput) *MedialiveListInputSecurityGroupsResult

       ListInputs(ctx workflow.Context, input *medialive.ListInputsInput) (*medialive.ListInputsOutput, error)
       ListInputsAsync(ctx workflow.Context, input *medialive.ListInputsInput) *MedialiveListInputsResult

       ListMultiplexPrograms(ctx workflow.Context, input *medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error)
       ListMultiplexProgramsAsync(ctx workflow.Context, input *medialive.ListMultiplexProgramsInput) *MedialiveListMultiplexProgramsResult

       ListMultiplexes(ctx workflow.Context, input *medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error)
       ListMultiplexesAsync(ctx workflow.Context, input *medialive.ListMultiplexesInput) *MedialiveListMultiplexesResult

       ListOfferings(ctx workflow.Context, input *medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error)
       ListOfferingsAsync(ctx workflow.Context, input *medialive.ListOfferingsInput) *MedialiveListOfferingsResult

       ListReservations(ctx workflow.Context, input *medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error)
       ListReservationsAsync(ctx workflow.Context, input *medialive.ListReservationsInput) *MedialiveListReservationsResult

       ListTagsForResource(ctx workflow.Context, input *medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *medialive.ListTagsForResourceInput) *MedialiveListTagsForResourceResult

       PurchaseOffering(ctx workflow.Context, input *medialive.PurchaseOfferingInput) (*medialive.PurchaseOfferingOutput, error)
       PurchaseOfferingAsync(ctx workflow.Context, input *medialive.PurchaseOfferingInput) *MedialivePurchaseOfferingResult

       StartChannel(ctx workflow.Context, input *medialive.StartChannelInput) (*medialive.StartChannelOutput, error)
       StartChannelAsync(ctx workflow.Context, input *medialive.StartChannelInput) *MedialiveStartChannelResult

       StartMultiplex(ctx workflow.Context, input *medialive.StartMultiplexInput) (*medialive.StartMultiplexOutput, error)
       StartMultiplexAsync(ctx workflow.Context, input *medialive.StartMultiplexInput) *MedialiveStartMultiplexResult

       StopChannel(ctx workflow.Context, input *medialive.StopChannelInput) (*medialive.StopChannelOutput, error)
       StopChannelAsync(ctx workflow.Context, input *medialive.StopChannelInput) *MedialiveStopChannelResult

       StopMultiplex(ctx workflow.Context, input *medialive.StopMultiplexInput) (*medialive.StopMultiplexOutput, error)
       StopMultiplexAsync(ctx workflow.Context, input *medialive.StopMultiplexInput) *MedialiveStopMultiplexResult

       UpdateChannel(ctx workflow.Context, input *medialive.UpdateChannelInput) (*medialive.UpdateChannelOutput, error)
       UpdateChannelAsync(ctx workflow.Context, input *medialive.UpdateChannelInput) *MedialiveUpdateChannelResult

       UpdateChannelClass(ctx workflow.Context, input *medialive.UpdateChannelClassInput) (*medialive.UpdateChannelClassOutput, error)
       UpdateChannelClassAsync(ctx workflow.Context, input *medialive.UpdateChannelClassInput) *MedialiveUpdateChannelClassResult

       UpdateInput(ctx workflow.Context, input *medialive.UpdateInputInput) (*medialive.UpdateInputOutput, error)
       UpdateInputAsync(ctx workflow.Context, input *medialive.UpdateInputInput) *MedialiveUpdateInputResult

       UpdateInputDevice(ctx workflow.Context, input *medialive.UpdateInputDeviceInput) (*medialive.UpdateInputDeviceOutput, error)
       UpdateInputDeviceAsync(ctx workflow.Context, input *medialive.UpdateInputDeviceInput) *MedialiveUpdateInputDeviceResult

       UpdateInputSecurityGroup(ctx workflow.Context, input *medialive.UpdateInputSecurityGroupInput) (*medialive.UpdateInputSecurityGroupOutput, error)
       UpdateInputSecurityGroupAsync(ctx workflow.Context, input *medialive.UpdateInputSecurityGroupInput) *MedialiveUpdateInputSecurityGroupResult

       UpdateMultiplex(ctx workflow.Context, input *medialive.UpdateMultiplexInput) (*medialive.UpdateMultiplexOutput, error)
       UpdateMultiplexAsync(ctx workflow.Context, input *medialive.UpdateMultiplexInput) *MedialiveUpdateMultiplexResult

       UpdateMultiplexProgram(ctx workflow.Context, input *medialive.UpdateMultiplexProgramInput) (*medialive.UpdateMultiplexProgramOutput, error)
       UpdateMultiplexProgramAsync(ctx workflow.Context, input *medialive.UpdateMultiplexProgramInput) *MedialiveUpdateMultiplexProgramResult

       UpdateReservation(ctx workflow.Context, input *medialive.UpdateReservationInput) (*medialive.UpdateReservationOutput, error)
       UpdateReservationAsync(ctx workflow.Context, input *medialive.UpdateReservationInput) *MedialiveUpdateReservationResult

       WaitUntilChannelCreated(ctx workflow.Context, input *medialive.DescribeChannelInput) error
       WaitUntilChannelDeleted(ctx workflow.Context, input *medialive.DescribeChannelInput) error
       WaitUntilChannelRunning(ctx workflow.Context, input *medialive.DescribeChannelInput) error
       WaitUntilChannelStopped(ctx workflow.Context, input *medialive.DescribeChannelInput) error
       WaitUntilInputAttached(ctx workflow.Context, input *medialive.DescribeInputInput) error
       WaitUntilInputDeleted(ctx workflow.Context, input *medialive.DescribeInputInput) error
       WaitUntilInputDetached(ctx workflow.Context, input *medialive.DescribeInputInput) error
       WaitUntilMultiplexCreated(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error
       WaitUntilMultiplexDeleted(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error
       WaitUntilMultiplexRunning(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error
       WaitUntilMultiplexStopped(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error}

type MediaLiveStub struct {
}

func NewMediaLiveStub() MediaLiveClient {
    return &MediaLiveStub{}
}

type MedialiveBatchUpdateScheduleResult struct {
	Result workflow.Future
}

func (r *MedialiveBatchUpdateScheduleResult) Get(ctx workflow.Context) (*medialive.BatchUpdateScheduleOutput, error) {
    var output medialive.BatchUpdateScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveCreateChannelResult struct {
	Result workflow.Future
}

func (r *MedialiveCreateChannelResult) Get(ctx workflow.Context) (*medialive.CreateChannelOutput, error) {
    var output medialive.CreateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveCreateInputResult struct {
	Result workflow.Future
}

func (r *MedialiveCreateInputResult) Get(ctx workflow.Context) (*medialive.CreateInputOutput, error) {
    var output medialive.CreateInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveCreateInputSecurityGroupResult struct {
	Result workflow.Future
}

func (r *MedialiveCreateInputSecurityGroupResult) Get(ctx workflow.Context) (*medialive.CreateInputSecurityGroupOutput, error) {
    var output medialive.CreateInputSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveCreateMultiplexResult struct {
	Result workflow.Future
}

func (r *MedialiveCreateMultiplexResult) Get(ctx workflow.Context) (*medialive.CreateMultiplexOutput, error) {
    var output medialive.CreateMultiplexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveCreateMultiplexProgramResult struct {
	Result workflow.Future
}

func (r *MedialiveCreateMultiplexProgramResult) Get(ctx workflow.Context) (*medialive.CreateMultiplexProgramOutput, error) {
    var output medialive.CreateMultiplexProgramOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveCreateTagsResult struct {
	Result workflow.Future
}

func (r *MedialiveCreateTagsResult) Get(ctx workflow.Context) (*medialive.CreateTagsOutput, error) {
    var output medialive.CreateTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteChannelResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteChannelResult) Get(ctx workflow.Context) (*medialive.DeleteChannelOutput, error) {
    var output medialive.DeleteChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteInputResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteInputResult) Get(ctx workflow.Context) (*medialive.DeleteInputOutput, error) {
    var output medialive.DeleteInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteInputSecurityGroupResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteInputSecurityGroupResult) Get(ctx workflow.Context) (*medialive.DeleteInputSecurityGroupOutput, error) {
    var output medialive.DeleteInputSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteMultiplexResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteMultiplexResult) Get(ctx workflow.Context) (*medialive.DeleteMultiplexOutput, error) {
    var output medialive.DeleteMultiplexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteMultiplexProgramResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteMultiplexProgramResult) Get(ctx workflow.Context) (*medialive.DeleteMultiplexProgramOutput, error) {
    var output medialive.DeleteMultiplexProgramOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteReservationResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteReservationResult) Get(ctx workflow.Context) (*medialive.DeleteReservationOutput, error) {
    var output medialive.DeleteReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteScheduleResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteScheduleResult) Get(ctx workflow.Context) (*medialive.DeleteScheduleOutput, error) {
    var output medialive.DeleteScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDeleteTagsResult struct {
	Result workflow.Future
}

func (r *MedialiveDeleteTagsResult) Get(ctx workflow.Context) (*medialive.DeleteTagsOutput, error) {
    var output medialive.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeChannelResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeChannelResult) Get(ctx workflow.Context) (*medialive.DescribeChannelOutput, error) {
    var output medialive.DescribeChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeInputResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeInputResult) Get(ctx workflow.Context) (*medialive.DescribeInputOutput, error) {
    var output medialive.DescribeInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeInputDeviceResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeInputDeviceResult) Get(ctx workflow.Context) (*medialive.DescribeInputDeviceOutput, error) {
    var output medialive.DescribeInputDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeInputDeviceThumbnailResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeInputDeviceThumbnailResult) Get(ctx workflow.Context) (*medialive.DescribeInputDeviceThumbnailOutput, error) {
    var output medialive.DescribeInputDeviceThumbnailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeInputSecurityGroupResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeInputSecurityGroupResult) Get(ctx workflow.Context) (*medialive.DescribeInputSecurityGroupOutput, error) {
    var output medialive.DescribeInputSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeMultiplexResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeMultiplexResult) Get(ctx workflow.Context) (*medialive.DescribeMultiplexOutput, error) {
    var output medialive.DescribeMultiplexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeMultiplexProgramResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeMultiplexProgramResult) Get(ctx workflow.Context) (*medialive.DescribeMultiplexProgramOutput, error) {
    var output medialive.DescribeMultiplexProgramOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeOfferingResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeOfferingResult) Get(ctx workflow.Context) (*medialive.DescribeOfferingOutput, error) {
    var output medialive.DescribeOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeReservationResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeReservationResult) Get(ctx workflow.Context) (*medialive.DescribeReservationOutput, error) {
    var output medialive.DescribeReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveDescribeScheduleResult struct {
	Result workflow.Future
}

func (r *MedialiveDescribeScheduleResult) Get(ctx workflow.Context) (*medialive.DescribeScheduleOutput, error) {
    var output medialive.DescribeScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListChannelsResult struct {
	Result workflow.Future
}

func (r *MedialiveListChannelsResult) Get(ctx workflow.Context) (*medialive.ListChannelsOutput, error) {
    var output medialive.ListChannelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListInputDevicesResult struct {
	Result workflow.Future
}

func (r *MedialiveListInputDevicesResult) Get(ctx workflow.Context) (*medialive.ListInputDevicesOutput, error) {
    var output medialive.ListInputDevicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListInputSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *MedialiveListInputSecurityGroupsResult) Get(ctx workflow.Context) (*medialive.ListInputSecurityGroupsOutput, error) {
    var output medialive.ListInputSecurityGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListInputsResult struct {
	Result workflow.Future
}

func (r *MedialiveListInputsResult) Get(ctx workflow.Context) (*medialive.ListInputsOutput, error) {
    var output medialive.ListInputsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListMultiplexProgramsResult struct {
	Result workflow.Future
}

func (r *MedialiveListMultiplexProgramsResult) Get(ctx workflow.Context) (*medialive.ListMultiplexProgramsOutput, error) {
    var output medialive.ListMultiplexProgramsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListMultiplexesResult struct {
	Result workflow.Future
}

func (r *MedialiveListMultiplexesResult) Get(ctx workflow.Context) (*medialive.ListMultiplexesOutput, error) {
    var output medialive.ListMultiplexesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListOfferingsResult struct {
	Result workflow.Future
}

func (r *MedialiveListOfferingsResult) Get(ctx workflow.Context) (*medialive.ListOfferingsOutput, error) {
    var output medialive.ListOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListReservationsResult struct {
	Result workflow.Future
}

func (r *MedialiveListReservationsResult) Get(ctx workflow.Context) (*medialive.ListReservationsOutput, error) {
    var output medialive.ListReservationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *MedialiveListTagsForResourceResult) Get(ctx workflow.Context) (*medialive.ListTagsForResourceOutput, error) {
    var output medialive.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialivePurchaseOfferingResult struct {
	Result workflow.Future
}

func (r *MedialivePurchaseOfferingResult) Get(ctx workflow.Context) (*medialive.PurchaseOfferingOutput, error) {
    var output medialive.PurchaseOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveStartChannelResult struct {
	Result workflow.Future
}

func (r *MedialiveStartChannelResult) Get(ctx workflow.Context) (*medialive.StartChannelOutput, error) {
    var output medialive.StartChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveStartMultiplexResult struct {
	Result workflow.Future
}

func (r *MedialiveStartMultiplexResult) Get(ctx workflow.Context) (*medialive.StartMultiplexOutput, error) {
    var output medialive.StartMultiplexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveStopChannelResult struct {
	Result workflow.Future
}

func (r *MedialiveStopChannelResult) Get(ctx workflow.Context) (*medialive.StopChannelOutput, error) {
    var output medialive.StopChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveStopMultiplexResult struct {
	Result workflow.Future
}

func (r *MedialiveStopMultiplexResult) Get(ctx workflow.Context) (*medialive.StopMultiplexOutput, error) {
    var output medialive.StopMultiplexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateChannelResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateChannelResult) Get(ctx workflow.Context) (*medialive.UpdateChannelOutput, error) {
    var output medialive.UpdateChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateChannelClassResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateChannelClassResult) Get(ctx workflow.Context) (*medialive.UpdateChannelClassOutput, error) {
    var output medialive.UpdateChannelClassOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateInputResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateInputResult) Get(ctx workflow.Context) (*medialive.UpdateInputOutput, error) {
    var output medialive.UpdateInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateInputDeviceResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateInputDeviceResult) Get(ctx workflow.Context) (*medialive.UpdateInputDeviceOutput, error) {
    var output medialive.UpdateInputDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateInputSecurityGroupResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateInputSecurityGroupResult) Get(ctx workflow.Context) (*medialive.UpdateInputSecurityGroupOutput, error) {
    var output medialive.UpdateInputSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateMultiplexResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateMultiplexResult) Get(ctx workflow.Context) (*medialive.UpdateMultiplexOutput, error) {
    var output medialive.UpdateMultiplexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateMultiplexProgramResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateMultiplexProgramResult) Get(ctx workflow.Context) (*medialive.UpdateMultiplexProgramOutput, error) {
    var output medialive.UpdateMultiplexProgramOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type MedialiveUpdateReservationResult struct {
	Result workflow.Future
}

func (r *MedialiveUpdateReservationResult) Get(ctx workflow.Context) (*medialive.UpdateReservationOutput, error) {
    var output medialive.UpdateReservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) BatchUpdateSchedule(ctx workflow.Context, input *medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error) {
    var output medialive.BatchUpdateScheduleOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.BatchUpdateSchedule", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) BatchUpdateScheduleAsync(ctx workflow.Context, input *medialive.BatchUpdateScheduleInput) *MedialiveBatchUpdateScheduleResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.BatchUpdateSchedule", input)
    return &MedialiveBatchUpdateScheduleResult{Result: future}
}

func (a *MediaLiveStub) CreateChannel(ctx workflow.Context, input *medialive.CreateChannelInput) (*medialive.CreateChannelOutput, error) {
    var output medialive.CreateChannelOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.CreateChannel", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) CreateChannelAsync(ctx workflow.Context, input *medialive.CreateChannelInput) *MedialiveCreateChannelResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.CreateChannel", input)
    return &MedialiveCreateChannelResult{Result: future}
}

func (a *MediaLiveStub) CreateInput(ctx workflow.Context, input *medialive.CreateInputInput) (*medialive.CreateInputOutput, error) {
    var output medialive.CreateInputOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.CreateInput", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) CreateInputAsync(ctx workflow.Context, input *medialive.CreateInputInput) *MedialiveCreateInputResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.CreateInput", input)
    return &MedialiveCreateInputResult{Result: future}
}

func (a *MediaLiveStub) CreateInputSecurityGroup(ctx workflow.Context, input *medialive.CreateInputSecurityGroupInput) (*medialive.CreateInputSecurityGroupOutput, error) {
    var output medialive.CreateInputSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.CreateInputSecurityGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) CreateInputSecurityGroupAsync(ctx workflow.Context, input *medialive.CreateInputSecurityGroupInput) *MedialiveCreateInputSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.CreateInputSecurityGroup", input)
    return &MedialiveCreateInputSecurityGroupResult{Result: future}
}

func (a *MediaLiveStub) CreateMultiplex(ctx workflow.Context, input *medialive.CreateMultiplexInput) (*medialive.CreateMultiplexOutput, error) {
    var output medialive.CreateMultiplexOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.CreateMultiplex", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) CreateMultiplexAsync(ctx workflow.Context, input *medialive.CreateMultiplexInput) *MedialiveCreateMultiplexResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.CreateMultiplex", input)
    return &MedialiveCreateMultiplexResult{Result: future}
}

func (a *MediaLiveStub) CreateMultiplexProgram(ctx workflow.Context, input *medialive.CreateMultiplexProgramInput) (*medialive.CreateMultiplexProgramOutput, error) {
    var output medialive.CreateMultiplexProgramOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.CreateMultiplexProgram", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) CreateMultiplexProgramAsync(ctx workflow.Context, input *medialive.CreateMultiplexProgramInput) *MedialiveCreateMultiplexProgramResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.CreateMultiplexProgram", input)
    return &MedialiveCreateMultiplexProgramResult{Result: future}
}

func (a *MediaLiveStub) CreateTags(ctx workflow.Context, input *medialive.CreateTagsInput) (*medialive.CreateTagsOutput, error) {
    var output medialive.CreateTagsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.CreateTags", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) CreateTagsAsync(ctx workflow.Context, input *medialive.CreateTagsInput) *MedialiveCreateTagsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.CreateTags", input)
    return &MedialiveCreateTagsResult{Result: future}
}

func (a *MediaLiveStub) DeleteChannel(ctx workflow.Context, input *medialive.DeleteChannelInput) (*medialive.DeleteChannelOutput, error) {
    var output medialive.DeleteChannelOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteChannel", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteChannelAsync(ctx workflow.Context, input *medialive.DeleteChannelInput) *MedialiveDeleteChannelResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteChannel", input)
    return &MedialiveDeleteChannelResult{Result: future}
}

func (a *MediaLiveStub) DeleteInput(ctx workflow.Context, input *medialive.DeleteInputInput) (*medialive.DeleteInputOutput, error) {
    var output medialive.DeleteInputOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteInput", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteInputAsync(ctx workflow.Context, input *medialive.DeleteInputInput) *MedialiveDeleteInputResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteInput", input)
    return &MedialiveDeleteInputResult{Result: future}
}

func (a *MediaLiveStub) DeleteInputSecurityGroup(ctx workflow.Context, input *medialive.DeleteInputSecurityGroupInput) (*medialive.DeleteInputSecurityGroupOutput, error) {
    var output medialive.DeleteInputSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteInputSecurityGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteInputSecurityGroupAsync(ctx workflow.Context, input *medialive.DeleteInputSecurityGroupInput) *MedialiveDeleteInputSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteInputSecurityGroup", input)
    return &MedialiveDeleteInputSecurityGroupResult{Result: future}
}

func (a *MediaLiveStub) DeleteMultiplex(ctx workflow.Context, input *medialive.DeleteMultiplexInput) (*medialive.DeleteMultiplexOutput, error) {
    var output medialive.DeleteMultiplexOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteMultiplex", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteMultiplexAsync(ctx workflow.Context, input *medialive.DeleteMultiplexInput) *MedialiveDeleteMultiplexResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteMultiplex", input)
    return &MedialiveDeleteMultiplexResult{Result: future}
}

func (a *MediaLiveStub) DeleteMultiplexProgram(ctx workflow.Context, input *medialive.DeleteMultiplexProgramInput) (*medialive.DeleteMultiplexProgramOutput, error) {
    var output medialive.DeleteMultiplexProgramOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteMultiplexProgram", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteMultiplexProgramAsync(ctx workflow.Context, input *medialive.DeleteMultiplexProgramInput) *MedialiveDeleteMultiplexProgramResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteMultiplexProgram", input)
    return &MedialiveDeleteMultiplexProgramResult{Result: future}
}

func (a *MediaLiveStub) DeleteReservation(ctx workflow.Context, input *medialive.DeleteReservationInput) (*medialive.DeleteReservationOutput, error) {
    var output medialive.DeleteReservationOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteReservation", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteReservationAsync(ctx workflow.Context, input *medialive.DeleteReservationInput) *MedialiveDeleteReservationResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteReservation", input)
    return &MedialiveDeleteReservationResult{Result: future}
}

func (a *MediaLiveStub) DeleteSchedule(ctx workflow.Context, input *medialive.DeleteScheduleInput) (*medialive.DeleteScheduleOutput, error) {
    var output medialive.DeleteScheduleOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteSchedule", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteScheduleAsync(ctx workflow.Context, input *medialive.DeleteScheduleInput) *MedialiveDeleteScheduleResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteSchedule", input)
    return &MedialiveDeleteScheduleResult{Result: future}
}

func (a *MediaLiveStub) DeleteTags(ctx workflow.Context, input *medialive.DeleteTagsInput) (*medialive.DeleteTagsOutput, error) {
    var output medialive.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DeleteTags", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DeleteTagsAsync(ctx workflow.Context, input *medialive.DeleteTagsInput) *MedialiveDeleteTagsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DeleteTags", input)
    return &MedialiveDeleteTagsResult{Result: future}
}

func (a *MediaLiveStub) DescribeChannel(ctx workflow.Context, input *medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error) {
    var output medialive.DescribeChannelOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeChannel", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeChannelAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) *MedialiveDescribeChannelResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeChannel", input)
    return &MedialiveDescribeChannelResult{Result: future}
}

func (a *MediaLiveStub) DescribeInput(ctx workflow.Context, input *medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error) {
    var output medialive.DescribeInputOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInput", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeInputAsync(ctx workflow.Context, input *medialive.DescribeInputInput) *MedialiveDescribeInputResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInput", input)
    return &MedialiveDescribeInputResult{Result: future}
}

func (a *MediaLiveStub) DescribeInputDevice(ctx workflow.Context, input *medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error) {
    var output medialive.DescribeInputDeviceOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInputDevice", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeInputDeviceAsync(ctx workflow.Context, input *medialive.DescribeInputDeviceInput) *MedialiveDescribeInputDeviceResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInputDevice", input)
    return &MedialiveDescribeInputDeviceResult{Result: future}
}

func (a *MediaLiveStub) DescribeInputDeviceThumbnail(ctx workflow.Context, input *medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error) {
    var output medialive.DescribeInputDeviceThumbnailOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInputDeviceThumbnail", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeInputDeviceThumbnailAsync(ctx workflow.Context, input *medialive.DescribeInputDeviceThumbnailInput) *MedialiveDescribeInputDeviceThumbnailResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInputDeviceThumbnail", input)
    return &MedialiveDescribeInputDeviceThumbnailResult{Result: future}
}

func (a *MediaLiveStub) DescribeInputSecurityGroup(ctx workflow.Context, input *medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error) {
    var output medialive.DescribeInputSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInputSecurityGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeInputSecurityGroupAsync(ctx workflow.Context, input *medialive.DescribeInputSecurityGroupInput) *MedialiveDescribeInputSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeInputSecurityGroup", input)
    return &MedialiveDescribeInputSecurityGroupResult{Result: future}
}

func (a *MediaLiveStub) DescribeMultiplex(ctx workflow.Context, input *medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error) {
    var output medialive.DescribeMultiplexOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeMultiplex", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeMultiplexAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) *MedialiveDescribeMultiplexResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeMultiplex", input)
    return &MedialiveDescribeMultiplexResult{Result: future}
}

func (a *MediaLiveStub) DescribeMultiplexProgram(ctx workflow.Context, input *medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error) {
    var output medialive.DescribeMultiplexProgramOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeMultiplexProgram", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeMultiplexProgramAsync(ctx workflow.Context, input *medialive.DescribeMultiplexProgramInput) *MedialiveDescribeMultiplexProgramResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeMultiplexProgram", input)
    return &MedialiveDescribeMultiplexProgramResult{Result: future}
}

func (a *MediaLiveStub) DescribeOffering(ctx workflow.Context, input *medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error) {
    var output medialive.DescribeOfferingOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeOffering", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeOfferingAsync(ctx workflow.Context, input *medialive.DescribeOfferingInput) *MedialiveDescribeOfferingResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeOffering", input)
    return &MedialiveDescribeOfferingResult{Result: future}
}

func (a *MediaLiveStub) DescribeReservation(ctx workflow.Context, input *medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error) {
    var output medialive.DescribeReservationOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeReservation", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeReservationAsync(ctx workflow.Context, input *medialive.DescribeReservationInput) *MedialiveDescribeReservationResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeReservation", input)
    return &MedialiveDescribeReservationResult{Result: future}
}

func (a *MediaLiveStub) DescribeSchedule(ctx workflow.Context, input *medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error) {
    var output medialive.DescribeScheduleOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.DescribeSchedule", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) DescribeScheduleAsync(ctx workflow.Context, input *medialive.DescribeScheduleInput) *MedialiveDescribeScheduleResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.DescribeSchedule", input)
    return &MedialiveDescribeScheduleResult{Result: future}
}

func (a *MediaLiveStub) ListChannels(ctx workflow.Context, input *medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error) {
    var output medialive.ListChannelsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListChannels", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListChannelsAsync(ctx workflow.Context, input *medialive.ListChannelsInput) *MedialiveListChannelsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListChannels", input)
    return &MedialiveListChannelsResult{Result: future}
}

func (a *MediaLiveStub) ListInputDevices(ctx workflow.Context, input *medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error) {
    var output medialive.ListInputDevicesOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListInputDevices", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListInputDevicesAsync(ctx workflow.Context, input *medialive.ListInputDevicesInput) *MedialiveListInputDevicesResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListInputDevices", input)
    return &MedialiveListInputDevicesResult{Result: future}
}

func (a *MediaLiveStub) ListInputSecurityGroups(ctx workflow.Context, input *medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error) {
    var output medialive.ListInputSecurityGroupsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListInputSecurityGroups", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListInputSecurityGroupsAsync(ctx workflow.Context, input *medialive.ListInputSecurityGroupsInput) *MedialiveListInputSecurityGroupsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListInputSecurityGroups", input)
    return &MedialiveListInputSecurityGroupsResult{Result: future}
}

func (a *MediaLiveStub) ListInputs(ctx workflow.Context, input *medialive.ListInputsInput) (*medialive.ListInputsOutput, error) {
    var output medialive.ListInputsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListInputs", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListInputsAsync(ctx workflow.Context, input *medialive.ListInputsInput) *MedialiveListInputsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListInputs", input)
    return &MedialiveListInputsResult{Result: future}
}

func (a *MediaLiveStub) ListMultiplexPrograms(ctx workflow.Context, input *medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error) {
    var output medialive.ListMultiplexProgramsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListMultiplexPrograms", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListMultiplexProgramsAsync(ctx workflow.Context, input *medialive.ListMultiplexProgramsInput) *MedialiveListMultiplexProgramsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListMultiplexPrograms", input)
    return &MedialiveListMultiplexProgramsResult{Result: future}
}

func (a *MediaLiveStub) ListMultiplexes(ctx workflow.Context, input *medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error) {
    var output medialive.ListMultiplexesOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListMultiplexes", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListMultiplexesAsync(ctx workflow.Context, input *medialive.ListMultiplexesInput) *MedialiveListMultiplexesResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListMultiplexes", input)
    return &MedialiveListMultiplexesResult{Result: future}
}

func (a *MediaLiveStub) ListOfferings(ctx workflow.Context, input *medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error) {
    var output medialive.ListOfferingsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListOfferings", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListOfferingsAsync(ctx workflow.Context, input *medialive.ListOfferingsInput) *MedialiveListOfferingsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListOfferings", input)
    return &MedialiveListOfferingsResult{Result: future}
}

func (a *MediaLiveStub) ListReservations(ctx workflow.Context, input *medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error) {
    var output medialive.ListReservationsOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListReservations", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListReservationsAsync(ctx workflow.Context, input *medialive.ListReservationsInput) *MedialiveListReservationsResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListReservations", input)
    return &MedialiveListReservationsResult{Result: future}
}

func (a *MediaLiveStub) ListTagsForResource(ctx workflow.Context, input *medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error) {
    var output medialive.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) ListTagsForResourceAsync(ctx workflow.Context, input *medialive.ListTagsForResourceInput) *MedialiveListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.ListTagsForResource", input)
    return &MedialiveListTagsForResourceResult{Result: future}
}

func (a *MediaLiveStub) PurchaseOffering(ctx workflow.Context, input *medialive.PurchaseOfferingInput) (*medialive.PurchaseOfferingOutput, error) {
    var output medialive.PurchaseOfferingOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.PurchaseOffering", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) PurchaseOfferingAsync(ctx workflow.Context, input *medialive.PurchaseOfferingInput) *MedialivePurchaseOfferingResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.PurchaseOffering", input)
    return &MedialivePurchaseOfferingResult{Result: future}
}

func (a *MediaLiveStub) StartChannel(ctx workflow.Context, input *medialive.StartChannelInput) (*medialive.StartChannelOutput, error) {
    var output medialive.StartChannelOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.StartChannel", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) StartChannelAsync(ctx workflow.Context, input *medialive.StartChannelInput) *MedialiveStartChannelResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.StartChannel", input)
    return &MedialiveStartChannelResult{Result: future}
}

func (a *MediaLiveStub) StartMultiplex(ctx workflow.Context, input *medialive.StartMultiplexInput) (*medialive.StartMultiplexOutput, error) {
    var output medialive.StartMultiplexOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.StartMultiplex", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) StartMultiplexAsync(ctx workflow.Context, input *medialive.StartMultiplexInput) *MedialiveStartMultiplexResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.StartMultiplex", input)
    return &MedialiveStartMultiplexResult{Result: future}
}

func (a *MediaLiveStub) StopChannel(ctx workflow.Context, input *medialive.StopChannelInput) (*medialive.StopChannelOutput, error) {
    var output medialive.StopChannelOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.StopChannel", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) StopChannelAsync(ctx workflow.Context, input *medialive.StopChannelInput) *MedialiveStopChannelResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.StopChannel", input)
    return &MedialiveStopChannelResult{Result: future}
}

func (a *MediaLiveStub) StopMultiplex(ctx workflow.Context, input *medialive.StopMultiplexInput) (*medialive.StopMultiplexOutput, error) {
    var output medialive.StopMultiplexOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.StopMultiplex", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) StopMultiplexAsync(ctx workflow.Context, input *medialive.StopMultiplexInput) *MedialiveStopMultiplexResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.StopMultiplex", input)
    return &MedialiveStopMultiplexResult{Result: future}
}

func (a *MediaLiveStub) UpdateChannel(ctx workflow.Context, input *medialive.UpdateChannelInput) (*medialive.UpdateChannelOutput, error) {
    var output medialive.UpdateChannelOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateChannel", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateChannelAsync(ctx workflow.Context, input *medialive.UpdateChannelInput) *MedialiveUpdateChannelResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateChannel", input)
    return &MedialiveUpdateChannelResult{Result: future}
}

func (a *MediaLiveStub) UpdateChannelClass(ctx workflow.Context, input *medialive.UpdateChannelClassInput) (*medialive.UpdateChannelClassOutput, error) {
    var output medialive.UpdateChannelClassOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateChannelClass", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateChannelClassAsync(ctx workflow.Context, input *medialive.UpdateChannelClassInput) *MedialiveUpdateChannelClassResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateChannelClass", input)
    return &MedialiveUpdateChannelClassResult{Result: future}
}

func (a *MediaLiveStub) UpdateInput(ctx workflow.Context, input *medialive.UpdateInputInput) (*medialive.UpdateInputOutput, error) {
    var output medialive.UpdateInputOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateInput", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateInputAsync(ctx workflow.Context, input *medialive.UpdateInputInput) *MedialiveUpdateInputResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateInput", input)
    return &MedialiveUpdateInputResult{Result: future}
}

func (a *MediaLiveStub) UpdateInputDevice(ctx workflow.Context, input *medialive.UpdateInputDeviceInput) (*medialive.UpdateInputDeviceOutput, error) {
    var output medialive.UpdateInputDeviceOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateInputDevice", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateInputDeviceAsync(ctx workflow.Context, input *medialive.UpdateInputDeviceInput) *MedialiveUpdateInputDeviceResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateInputDevice", input)
    return &MedialiveUpdateInputDeviceResult{Result: future}
}

func (a *MediaLiveStub) UpdateInputSecurityGroup(ctx workflow.Context, input *medialive.UpdateInputSecurityGroupInput) (*medialive.UpdateInputSecurityGroupOutput, error) {
    var output medialive.UpdateInputSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateInputSecurityGroup", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateInputSecurityGroupAsync(ctx workflow.Context, input *medialive.UpdateInputSecurityGroupInput) *MedialiveUpdateInputSecurityGroupResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateInputSecurityGroup", input)
    return &MedialiveUpdateInputSecurityGroupResult{Result: future}
}

func (a *MediaLiveStub) UpdateMultiplex(ctx workflow.Context, input *medialive.UpdateMultiplexInput) (*medialive.UpdateMultiplexOutput, error) {
    var output medialive.UpdateMultiplexOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateMultiplex", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateMultiplexAsync(ctx workflow.Context, input *medialive.UpdateMultiplexInput) *MedialiveUpdateMultiplexResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateMultiplex", input)
    return &MedialiveUpdateMultiplexResult{Result: future}
}

func (a *MediaLiveStub) UpdateMultiplexProgram(ctx workflow.Context, input *medialive.UpdateMultiplexProgramInput) (*medialive.UpdateMultiplexProgramOutput, error) {
    var output medialive.UpdateMultiplexProgramOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateMultiplexProgram", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateMultiplexProgramAsync(ctx workflow.Context, input *medialive.UpdateMultiplexProgramInput) *MedialiveUpdateMultiplexProgramResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateMultiplexProgram", input)
    return &MedialiveUpdateMultiplexProgramResult{Result: future}
}

func (a *MediaLiveStub) UpdateReservation(ctx workflow.Context, input *medialive.UpdateReservationInput) (*medialive.UpdateReservationOutput, error) {
    var output medialive.UpdateReservationOutput
    err := workflow.ExecuteActivity(ctx, "MediaLive.UpdateReservation", input).Get(ctx, &output)
    return &output, err
}

func (a *MediaLiveStub) UpdateReservationAsync(ctx workflow.Context, input *medialive.UpdateReservationInput) *MedialiveUpdateReservationResult {
    future := workflow.ExecuteActivity(ctx, "MediaLive.UpdateReservation", input)
    return &MedialiveUpdateReservationResult{Result: future}
}

func (a *MediaLiveStub) WaitUntilChannelCreated(ctx workflow.Context, input *medialive.DescribeChannelInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelCreated", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilChannelCreatedAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelCreated", input)
}


func (a *MediaLiveStub) WaitUntilChannelDeleted(ctx workflow.Context, input *medialive.DescribeChannelInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelDeleted", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilChannelDeletedAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelDeleted", input)
}


func (a *MediaLiveStub) WaitUntilChannelRunning(ctx workflow.Context, input *medialive.DescribeChannelInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelRunning", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilChannelRunningAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelRunning", input)
}


func (a *MediaLiveStub) WaitUntilChannelStopped(ctx workflow.Context, input *medialive.DescribeChannelInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelStopped", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilChannelStoppedAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilChannelStopped", input)
}


func (a *MediaLiveStub) WaitUntilInputAttached(ctx workflow.Context, input *medialive.DescribeInputInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilInputAttached", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilInputAttachedAsync(ctx workflow.Context, input *medialive.DescribeInputInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilInputAttached", input)
}


func (a *MediaLiveStub) WaitUntilInputDeleted(ctx workflow.Context, input *medialive.DescribeInputInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilInputDeleted", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilInputDeletedAsync(ctx workflow.Context, input *medialive.DescribeInputInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilInputDeleted", input)
}


func (a *MediaLiveStub) WaitUntilInputDetached(ctx workflow.Context, input *medialive.DescribeInputInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilInputDetached", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilInputDetachedAsync(ctx workflow.Context, input *medialive.DescribeInputInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilInputDetached", input)
}


func (a *MediaLiveStub) WaitUntilMultiplexCreated(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexCreated", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilMultiplexCreatedAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexCreated", input)
}


func (a *MediaLiveStub) WaitUntilMultiplexDeleted(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexDeleted", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilMultiplexDeletedAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexDeleted", input)
}


func (a *MediaLiveStub) WaitUntilMultiplexRunning(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexRunning", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilMultiplexRunningAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexRunning", input)
}


func (a *MediaLiveStub) WaitUntilMultiplexStopped(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexStopped", input).Get(ctx, nil)
}

func (a *MediaLiveStub) WaitUntilMultiplexStoppedAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, "MediaLive.WaitUntilMultiplexStopped", input)
}

