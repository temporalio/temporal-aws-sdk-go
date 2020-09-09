package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface"
)

type ComprehendMedicalActivities struct {
	client comprehendmedicaliface.ComprehendMedicalAPI
}

func NewComprehendMedicalActivities(session *session.Session, config ...*aws.Config) *ComprehendMedicalActivities {
	client := comprehendmedical.New(session, config...)
	return &ComprehendMedicalActivities{client: client}
}

func (a *ComprehendMedicalActivities) DescribeEntitiesDetectionV2Job(input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	return a.client.DescribeEntitiesDetectionV2Job(input)
}

func (a *ComprehendMedicalActivities) DescribeICD10CMInferenceJob(input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	return a.client.DescribeICD10CMInferenceJob(input)
}

func (a *ComprehendMedicalActivities) DescribePHIDetectionJob(input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	return a.client.DescribePHIDetectionJob(input)
}

func (a *ComprehendMedicalActivities) DescribeRxNormInferenceJob(input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	return a.client.DescribeRxNormInferenceJob(input)
}

func (a *ComprehendMedicalActivities) DetectEntities(input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
	return a.client.DetectEntities(input)
}

func (a *ComprehendMedicalActivities) DetectEntitiesV2(input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error) {
	return a.client.DetectEntitiesV2(input)
}

func (a *ComprehendMedicalActivities) DetectPHI(input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error) {
	return a.client.DetectPHI(input)
}

func (a *ComprehendMedicalActivities) InferICD10CM(input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error) {
	return a.client.InferICD10CM(input)
}

func (a *ComprehendMedicalActivities) InferRxNorm(input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error) {
	return a.client.InferRxNorm(input)
}

func (a *ComprehendMedicalActivities) ListEntitiesDetectionV2Jobs(input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	return a.client.ListEntitiesDetectionV2Jobs(input)
}

func (a *ComprehendMedicalActivities) ListICD10CMInferenceJobs(input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	return a.client.ListICD10CMInferenceJobs(input)
}

func (a *ComprehendMedicalActivities) ListPHIDetectionJobs(input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	return a.client.ListPHIDetectionJobs(input)
}

func (a *ComprehendMedicalActivities) ListRxNormInferenceJobs(input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	return a.client.ListRxNormInferenceJobs(input)
}

func (a *ComprehendMedicalActivities) StartEntitiesDetectionV2Job(input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	return a.client.StartEntitiesDetectionV2Job(input)
}

func (a *ComprehendMedicalActivities) StartICD10CMInferenceJob(input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	return a.client.StartICD10CMInferenceJob(input)
}

func (a *ComprehendMedicalActivities) StartPHIDetectionJob(input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	return a.client.StartPHIDetectionJob(input)
}

func (a *ComprehendMedicalActivities) StartRxNormInferenceJob(input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	return a.client.StartRxNormInferenceJob(input)
}

func (a *ComprehendMedicalActivities) StopEntitiesDetectionV2Job(input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	return a.client.StopEntitiesDetectionV2Job(input)
}

func (a *ComprehendMedicalActivities) StopICD10CMInferenceJob(input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	return a.client.StopICD10CMInferenceJob(input)
}

func (a *ComprehendMedicalActivities) StopPHIDetectionJob(input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	return a.client.StopPHIDetectionJob(input)
}

func (a *ComprehendMedicalActivities) StopRxNormInferenceJob(input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	return a.client.StopRxNormInferenceJob(input)
}
