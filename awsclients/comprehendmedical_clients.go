package awsclients

import (
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ComprehendMedicalClient interface {
	DescribeEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error)
	DescribeEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) *ComprehendmedicalDescribeEntitiesDetectionV2JobResult

	DescribeICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error)
	DescribeICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) *ComprehendmedicalDescribeICD10CMInferenceJobResult

	DescribePHIDetectionJob(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error)
	DescribePHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) *ComprehendmedicalDescribePHIDetectionJobResult

	DescribeRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error)
	DescribeRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) *ComprehendmedicalDescribeRxNormInferenceJobResult

	DetectEntities(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesAsync(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) *ComprehendmedicalDetectEntitiesResult

	DetectEntitiesV2(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error)
	DetectEntitiesV2Async(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) *ComprehendmedicalDetectEntitiesV2Result

	DetectPHI(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error)
	DetectPHIAsync(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) *ComprehendmedicalDetectPHIResult

	InferICD10CM(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error)
	InferICD10CMAsync(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) *ComprehendmedicalInferICD10CMResult

	InferRxNorm(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error)
	InferRxNormAsync(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) *ComprehendmedicalInferRxNormResult

	ListEntitiesDetectionV2Jobs(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error)
	ListEntitiesDetectionV2JobsAsync(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) *ComprehendmedicalListEntitiesDetectionV2JobsResult

	ListICD10CMInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error)
	ListICD10CMInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) *ComprehendmedicalListICD10CMInferenceJobsResult

	ListPHIDetectionJobs(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error)
	ListPHIDetectionJobsAsync(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) *ComprehendmedicalListPHIDetectionJobsResult

	ListRxNormInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error)
	ListRxNormInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) *ComprehendmedicalListRxNormInferenceJobsResult

	StartEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error)
	StartEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) *ComprehendmedicalStartEntitiesDetectionV2JobResult

	StartICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error)
	StartICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) *ComprehendmedicalStartICD10CMInferenceJobResult

	StartPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error)
	StartPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) *ComprehendmedicalStartPHIDetectionJobResult

	StartRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error)
	StartRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) *ComprehendmedicalStartRxNormInferenceJobResult

	StopEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error)
	StopEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) *ComprehendmedicalStopEntitiesDetectionV2JobResult

	StopICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error)
	StopICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) *ComprehendmedicalStopICD10CMInferenceJobResult

	StopPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error)
	StopPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) *ComprehendmedicalStopPHIDetectionJobResult

	StopRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error)
	StopRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) *ComprehendmedicalStopRxNormInferenceJobResult
}

type ComprehendmedicalDescribeEntitiesDetectionV2JobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalDescribeEntitiesDetectionV2JobResult) Get(ctx workflow.Context) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.DescribeEntitiesDetectionV2JobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalDescribeICD10CMInferenceJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalDescribeICD10CMInferenceJobResult) Get(ctx workflow.Context) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.DescribeICD10CMInferenceJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalDescribePHIDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalDescribePHIDetectionJobResult) Get(ctx workflow.Context) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	var output comprehendmedical.DescribePHIDetectionJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalDescribeRxNormInferenceJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalDescribeRxNormInferenceJobResult) Get(ctx workflow.Context) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	var output comprehendmedical.DescribeRxNormInferenceJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalDetectEntitiesResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalDetectEntitiesResult) Get(ctx workflow.Context) (*comprehendmedical.DetectEntitiesOutput, error) {
	var output comprehendmedical.DetectEntitiesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalDetectEntitiesV2Result struct {
	Result workflow.Future
}

func (r *ComprehendmedicalDetectEntitiesV2Result) Get(ctx workflow.Context) (*comprehendmedical.DetectEntitiesV2Output, error) {
	var output comprehendmedical.DetectEntitiesV2Output
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalDetectPHIResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalDetectPHIResult) Get(ctx workflow.Context) (*comprehendmedical.DetectPHIOutput, error) {
	var output comprehendmedical.DetectPHIOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalInferICD10CMResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalInferICD10CMResult) Get(ctx workflow.Context) (*comprehendmedical.InferICD10CMOutput, error) {
	var output comprehendmedical.InferICD10CMOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalInferRxNormResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalInferRxNormResult) Get(ctx workflow.Context) (*comprehendmedical.InferRxNormOutput, error) {
	var output comprehendmedical.InferRxNormOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalListEntitiesDetectionV2JobsResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalListEntitiesDetectionV2JobsResult) Get(ctx workflow.Context) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	var output comprehendmedical.ListEntitiesDetectionV2JobsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalListICD10CMInferenceJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalListICD10CMInferenceJobsResult) Get(ctx workflow.Context) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	var output comprehendmedical.ListICD10CMInferenceJobsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalListPHIDetectionJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalListPHIDetectionJobsResult) Get(ctx workflow.Context) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	var output comprehendmedical.ListPHIDetectionJobsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalListRxNormInferenceJobsResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalListRxNormInferenceJobsResult) Get(ctx workflow.Context) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	var output comprehendmedical.ListRxNormInferenceJobsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStartEntitiesDetectionV2JobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStartEntitiesDetectionV2JobResult) Get(ctx workflow.Context) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StartEntitiesDetectionV2JobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStartICD10CMInferenceJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStartICD10CMInferenceJobResult) Get(ctx workflow.Context) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StartICD10CMInferenceJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStartPHIDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStartPHIDetectionJobResult) Get(ctx workflow.Context) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	var output comprehendmedical.StartPHIDetectionJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStartRxNormInferenceJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStartRxNormInferenceJobResult) Get(ctx workflow.Context) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StartRxNormInferenceJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStopEntitiesDetectionV2JobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStopEntitiesDetectionV2JobResult) Get(ctx workflow.Context) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StopEntitiesDetectionV2JobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStopICD10CMInferenceJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStopICD10CMInferenceJobResult) Get(ctx workflow.Context) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StopICD10CMInferenceJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStopPHIDetectionJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStopPHIDetectionJobResult) Get(ctx workflow.Context) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	var output comprehendmedical.StopPHIDetectionJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendmedicalStopRxNormInferenceJobResult struct {
	Result workflow.Future
}

func (r *ComprehendmedicalStopRxNormInferenceJobResult) Get(ctx workflow.Context) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StopRxNormInferenceJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStub struct {
	activities awsactivities.ComprehendMedicalActivities
}

func NewComprehendMedicalStub() ComprehendMedicalClient {
	return &ComprehendMedicalStub{}
}

func (a *ComprehendMedicalStub) DescribeEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.DescribeEntitiesDetectionV2JobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEntitiesDetectionV2Job, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribeEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) *ComprehendmedicalDescribeEntitiesDetectionV2JobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEntitiesDetectionV2Job, input)
	return &ComprehendmedicalDescribeEntitiesDetectionV2JobResult{Result: future}
}

func (a *ComprehendMedicalStub) DescribeICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.DescribeICD10CMInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeICD10CMInferenceJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribeICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) *ComprehendmedicalDescribeICD10CMInferenceJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeICD10CMInferenceJob, input)
	return &ComprehendmedicalDescribeICD10CMInferenceJobResult{Result: future}
}

func (a *ComprehendMedicalStub) DescribePHIDetectionJob(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	var output comprehendmedical.DescribePHIDetectionJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribePHIDetectionJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribePHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) *ComprehendmedicalDescribePHIDetectionJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribePHIDetectionJob, input)
	return &ComprehendmedicalDescribePHIDetectionJobResult{Result: future}
}

func (a *ComprehendMedicalStub) DescribeRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	var output comprehendmedical.DescribeRxNormInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeRxNormInferenceJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribeRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) *ComprehendmedicalDescribeRxNormInferenceJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeRxNormInferenceJob, input)
	return &ComprehendmedicalDescribeRxNormInferenceJobResult{Result: future}
}

func (a *ComprehendMedicalStub) DetectEntities(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
	var output comprehendmedical.DetectEntitiesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DetectEntities, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DetectEntitiesAsync(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) *ComprehendmedicalDetectEntitiesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DetectEntities, input)
	return &ComprehendmedicalDetectEntitiesResult{Result: future}
}

func (a *ComprehendMedicalStub) DetectEntitiesV2(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error) {
	var output comprehendmedical.DetectEntitiesV2Output
	err := workflow.ExecuteActivity(ctx, a.activities.DetectEntitiesV2, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DetectEntitiesV2Async(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) *ComprehendmedicalDetectEntitiesV2Result {
	future := workflow.ExecuteActivity(ctx, a.activities.DetectEntitiesV2, input)
	return &ComprehendmedicalDetectEntitiesV2Result{Result: future}
}

func (a *ComprehendMedicalStub) DetectPHI(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error) {
	var output comprehendmedical.DetectPHIOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DetectPHI, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DetectPHIAsync(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) *ComprehendmedicalDetectPHIResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DetectPHI, input)
	return &ComprehendmedicalDetectPHIResult{Result: future}
}

func (a *ComprehendMedicalStub) InferICD10CM(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error) {
	var output comprehendmedical.InferICD10CMOutput
	err := workflow.ExecuteActivity(ctx, a.activities.InferICD10CM, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) InferICD10CMAsync(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) *ComprehendmedicalInferICD10CMResult {
	future := workflow.ExecuteActivity(ctx, a.activities.InferICD10CM, input)
	return &ComprehendmedicalInferICD10CMResult{Result: future}
}

func (a *ComprehendMedicalStub) InferRxNorm(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error) {
	var output comprehendmedical.InferRxNormOutput
	err := workflow.ExecuteActivity(ctx, a.activities.InferRxNorm, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) InferRxNormAsync(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) *ComprehendmedicalInferRxNormResult {
	future := workflow.ExecuteActivity(ctx, a.activities.InferRxNorm, input)
	return &ComprehendmedicalInferRxNormResult{Result: future}
}

func (a *ComprehendMedicalStub) ListEntitiesDetectionV2Jobs(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	var output comprehendmedical.ListEntitiesDetectionV2JobsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListEntitiesDetectionV2Jobs, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListEntitiesDetectionV2JobsAsync(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) *ComprehendmedicalListEntitiesDetectionV2JobsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListEntitiesDetectionV2Jobs, input)
	return &ComprehendmedicalListEntitiesDetectionV2JobsResult{Result: future}
}

func (a *ComprehendMedicalStub) ListICD10CMInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	var output comprehendmedical.ListICD10CMInferenceJobsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListICD10CMInferenceJobs, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListICD10CMInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) *ComprehendmedicalListICD10CMInferenceJobsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListICD10CMInferenceJobs, input)
	return &ComprehendmedicalListICD10CMInferenceJobsResult{Result: future}
}

func (a *ComprehendMedicalStub) ListPHIDetectionJobs(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	var output comprehendmedical.ListPHIDetectionJobsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListPHIDetectionJobs, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListPHIDetectionJobsAsync(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) *ComprehendmedicalListPHIDetectionJobsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListPHIDetectionJobs, input)
	return &ComprehendmedicalListPHIDetectionJobsResult{Result: future}
}

func (a *ComprehendMedicalStub) ListRxNormInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	var output comprehendmedical.ListRxNormInferenceJobsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListRxNormInferenceJobs, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListRxNormInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) *ComprehendmedicalListRxNormInferenceJobsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListRxNormInferenceJobs, input)
	return &ComprehendmedicalListRxNormInferenceJobsResult{Result: future}
}

func (a *ComprehendMedicalStub) StartEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StartEntitiesDetectionV2JobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartEntitiesDetectionV2Job, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) *ComprehendmedicalStartEntitiesDetectionV2JobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartEntitiesDetectionV2Job, input)
	return &ComprehendmedicalStartEntitiesDetectionV2JobResult{Result: future}
}

func (a *ComprehendMedicalStub) StartICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StartICD10CMInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartICD10CMInferenceJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) *ComprehendmedicalStartICD10CMInferenceJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartICD10CMInferenceJob, input)
	return &ComprehendmedicalStartICD10CMInferenceJobResult{Result: future}
}

func (a *ComprehendMedicalStub) StartPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	var output comprehendmedical.StartPHIDetectionJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartPHIDetectionJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) *ComprehendmedicalStartPHIDetectionJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartPHIDetectionJob, input)
	return &ComprehendmedicalStartPHIDetectionJobResult{Result: future}
}

func (a *ComprehendMedicalStub) StartRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StartRxNormInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartRxNormInferenceJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) *ComprehendmedicalStartRxNormInferenceJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartRxNormInferenceJob, input)
	return &ComprehendmedicalStartRxNormInferenceJobResult{Result: future}
}

func (a *ComprehendMedicalStub) StopEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StopEntitiesDetectionV2JobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopEntitiesDetectionV2Job, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) *ComprehendmedicalStopEntitiesDetectionV2JobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopEntitiesDetectionV2Job, input)
	return &ComprehendmedicalStopEntitiesDetectionV2JobResult{Result: future}
}

func (a *ComprehendMedicalStub) StopICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StopICD10CMInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopICD10CMInferenceJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) *ComprehendmedicalStopICD10CMInferenceJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopICD10CMInferenceJob, input)
	return &ComprehendmedicalStopICD10CMInferenceJobResult{Result: future}
}

func (a *ComprehendMedicalStub) StopPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	var output comprehendmedical.StopPHIDetectionJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopPHIDetectionJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) *ComprehendmedicalStopPHIDetectionJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopPHIDetectionJob, input)
	return &ComprehendmedicalStopPHIDetectionJobResult{Result: future}
}

func (a *ComprehendMedicalStub) StopRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StopRxNormInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopRxNormInferenceJob, input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) *ComprehendmedicalStopRxNormInferenceJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopRxNormInferenceJob, input)
	return &ComprehendmedicalStopRxNormInferenceJobResult{Result: future}
}
