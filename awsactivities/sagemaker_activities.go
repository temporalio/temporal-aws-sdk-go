package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemaker/sagemakeriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type SageMakerActivities struct {
	client sagemakeriface.SageMakerAPI
}

func NewSageMakerActivities(session *session.Session, config ...*aws.Config) *SageMakerActivities {
	client := sagemaker.New(session, config...)
	return &SageMakerActivities{client: client}
}

func (a *SageMakerActivities) AddTags(ctx context.Context, input *sagemaker.AddTagsInput) (*sagemaker.AddTagsOutput, error) {
	return a.client.AddTagsWithContext(ctx, input)
}

func (a *SageMakerActivities) AssociateTrialComponent(ctx context.Context, input *sagemaker.AssociateTrialComponentInput) (*sagemaker.AssociateTrialComponentOutput, error) {
	return a.client.AssociateTrialComponentWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateAlgorithm(ctx context.Context, input *sagemaker.CreateAlgorithmInput) (*sagemaker.CreateAlgorithmOutput, error) {
	return a.client.CreateAlgorithmWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateApp(ctx context.Context, input *sagemaker.CreateAppInput) (*sagemaker.CreateAppOutput, error) {
	return a.client.CreateAppWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateAutoMLJob(ctx context.Context, input *sagemaker.CreateAutoMLJobInput) (*sagemaker.CreateAutoMLJobOutput, error) {
	return a.client.CreateAutoMLJobWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateCodeRepository(ctx context.Context, input *sagemaker.CreateCodeRepositoryInput) (*sagemaker.CreateCodeRepositoryOutput, error) {
	return a.client.CreateCodeRepositoryWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateCompilationJob(ctx context.Context, input *sagemaker.CreateCompilationJobInput) (*sagemaker.CreateCompilationJobOutput, error) {
	return a.client.CreateCompilationJobWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateDomain(ctx context.Context, input *sagemaker.CreateDomainInput) (*sagemaker.CreateDomainOutput, error) {
	return a.client.CreateDomainWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateEndpoint(ctx context.Context, input *sagemaker.CreateEndpointInput) (*sagemaker.CreateEndpointOutput, error) {
	return a.client.CreateEndpointWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateEndpointConfig(ctx context.Context, input *sagemaker.CreateEndpointConfigInput) (*sagemaker.CreateEndpointConfigOutput, error) {
	return a.client.CreateEndpointConfigWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateExperiment(ctx context.Context, input *sagemaker.CreateExperimentInput) (*sagemaker.CreateExperimentOutput, error) {
	return a.client.CreateExperimentWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateFlowDefinition(ctx context.Context, input *sagemaker.CreateFlowDefinitionInput) (*sagemaker.CreateFlowDefinitionOutput, error) {
	return a.client.CreateFlowDefinitionWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateHumanTaskUi(ctx context.Context, input *sagemaker.CreateHumanTaskUiInput) (*sagemaker.CreateHumanTaskUiOutput, error) {
	return a.client.CreateHumanTaskUiWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateHyperParameterTuningJob(ctx context.Context, input *sagemaker.CreateHyperParameterTuningJobInput) (*sagemaker.CreateHyperParameterTuningJobOutput, error) {
	return a.client.CreateHyperParameterTuningJobWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateLabelingJob(ctx context.Context, input *sagemaker.CreateLabelingJobInput) (*sagemaker.CreateLabelingJobOutput, error) {
	return a.client.CreateLabelingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateModel(ctx context.Context, input *sagemaker.CreateModelInput) (*sagemaker.CreateModelOutput, error) {
	return a.client.CreateModelWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateModelPackage(ctx context.Context, input *sagemaker.CreateModelPackageInput) (*sagemaker.CreateModelPackageOutput, error) {
	return a.client.CreateModelPackageWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateMonitoringSchedule(ctx context.Context, input *sagemaker.CreateMonitoringScheduleInput) (*sagemaker.CreateMonitoringScheduleOutput, error) {
	return a.client.CreateMonitoringScheduleWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateNotebookInstance(ctx context.Context, input *sagemaker.CreateNotebookInstanceInput) (*sagemaker.CreateNotebookInstanceOutput, error) {
	return a.client.CreateNotebookInstanceWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateNotebookInstanceLifecycleConfig(ctx context.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) (*sagemaker.CreateNotebookInstanceLifecycleConfigOutput, error) {
	return a.client.CreateNotebookInstanceLifecycleConfigWithContext(ctx, input)
}

func (a *SageMakerActivities) CreatePresignedDomainUrl(ctx context.Context, input *sagemaker.CreatePresignedDomainUrlInput) (*sagemaker.CreatePresignedDomainUrlOutput, error) {
	return a.client.CreatePresignedDomainUrlWithContext(ctx, input)
}

func (a *SageMakerActivities) CreatePresignedNotebookInstanceUrl(ctx context.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) (*sagemaker.CreatePresignedNotebookInstanceUrlOutput, error) {
	return a.client.CreatePresignedNotebookInstanceUrlWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateProcessingJob(ctx context.Context, input *sagemaker.CreateProcessingJobInput) (*sagemaker.CreateProcessingJobOutput, error) {
	return a.client.CreateProcessingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateTrainingJob(ctx context.Context, input *sagemaker.CreateTrainingJobInput) (*sagemaker.CreateTrainingJobOutput, error) {
	return a.client.CreateTrainingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateTransformJob(ctx context.Context, input *sagemaker.CreateTransformJobInput) (*sagemaker.CreateTransformJobOutput, error) {
	return a.client.CreateTransformJobWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateTrial(ctx context.Context, input *sagemaker.CreateTrialInput) (*sagemaker.CreateTrialOutput, error) {
	return a.client.CreateTrialWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateTrialComponent(ctx context.Context, input *sagemaker.CreateTrialComponentInput) (*sagemaker.CreateTrialComponentOutput, error) {
	return a.client.CreateTrialComponentWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateUserProfile(ctx context.Context, input *sagemaker.CreateUserProfileInput) (*sagemaker.CreateUserProfileOutput, error) {
	return a.client.CreateUserProfileWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateWorkforce(ctx context.Context, input *sagemaker.CreateWorkforceInput) (*sagemaker.CreateWorkforceOutput, error) {
	return a.client.CreateWorkforceWithContext(ctx, input)
}

func (a *SageMakerActivities) CreateWorkteam(ctx context.Context, input *sagemaker.CreateWorkteamInput) (*sagemaker.CreateWorkteamOutput, error) {
	return a.client.CreateWorkteamWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteAlgorithm(ctx context.Context, input *sagemaker.DeleteAlgorithmInput) (*sagemaker.DeleteAlgorithmOutput, error) {
	return a.client.DeleteAlgorithmWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteApp(ctx context.Context, input *sagemaker.DeleteAppInput) (*sagemaker.DeleteAppOutput, error) {
	return a.client.DeleteAppWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteCodeRepository(ctx context.Context, input *sagemaker.DeleteCodeRepositoryInput) (*sagemaker.DeleteCodeRepositoryOutput, error) {
	return a.client.DeleteCodeRepositoryWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteDomain(ctx context.Context, input *sagemaker.DeleteDomainInput) (*sagemaker.DeleteDomainOutput, error) {
	return a.client.DeleteDomainWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteEndpoint(ctx context.Context, input *sagemaker.DeleteEndpointInput) (*sagemaker.DeleteEndpointOutput, error) {
	return a.client.DeleteEndpointWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteEndpointConfig(ctx context.Context, input *sagemaker.DeleteEndpointConfigInput) (*sagemaker.DeleteEndpointConfigOutput, error) {
	return a.client.DeleteEndpointConfigWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteExperiment(ctx context.Context, input *sagemaker.DeleteExperimentInput) (*sagemaker.DeleteExperimentOutput, error) {
	return a.client.DeleteExperimentWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteFlowDefinition(ctx context.Context, input *sagemaker.DeleteFlowDefinitionInput) (*sagemaker.DeleteFlowDefinitionOutput, error) {
	return a.client.DeleteFlowDefinitionWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteHumanTaskUi(ctx context.Context, input *sagemaker.DeleteHumanTaskUiInput) (*sagemaker.DeleteHumanTaskUiOutput, error) {
	return a.client.DeleteHumanTaskUiWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteModel(ctx context.Context, input *sagemaker.DeleteModelInput) (*sagemaker.DeleteModelOutput, error) {
	return a.client.DeleteModelWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteModelPackage(ctx context.Context, input *sagemaker.DeleteModelPackageInput) (*sagemaker.DeleteModelPackageOutput, error) {
	return a.client.DeleteModelPackageWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteMonitoringSchedule(ctx context.Context, input *sagemaker.DeleteMonitoringScheduleInput) (*sagemaker.DeleteMonitoringScheduleOutput, error) {
	return a.client.DeleteMonitoringScheduleWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteNotebookInstance(ctx context.Context, input *sagemaker.DeleteNotebookInstanceInput) (*sagemaker.DeleteNotebookInstanceOutput, error) {
	return a.client.DeleteNotebookInstanceWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteNotebookInstanceLifecycleConfig(ctx context.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) (*sagemaker.DeleteNotebookInstanceLifecycleConfigOutput, error) {
	return a.client.DeleteNotebookInstanceLifecycleConfigWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteTags(ctx context.Context, input *sagemaker.DeleteTagsInput) (*sagemaker.DeleteTagsOutput, error) {
	return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteTrial(ctx context.Context, input *sagemaker.DeleteTrialInput) (*sagemaker.DeleteTrialOutput, error) {
	return a.client.DeleteTrialWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteTrialComponent(ctx context.Context, input *sagemaker.DeleteTrialComponentInput) (*sagemaker.DeleteTrialComponentOutput, error) {
	return a.client.DeleteTrialComponentWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteUserProfile(ctx context.Context, input *sagemaker.DeleteUserProfileInput) (*sagemaker.DeleteUserProfileOutput, error) {
	return a.client.DeleteUserProfileWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteWorkforce(ctx context.Context, input *sagemaker.DeleteWorkforceInput) (*sagemaker.DeleteWorkforceOutput, error) {
	return a.client.DeleteWorkforceWithContext(ctx, input)
}

func (a *SageMakerActivities) DeleteWorkteam(ctx context.Context, input *sagemaker.DeleteWorkteamInput) (*sagemaker.DeleteWorkteamOutput, error) {
	return a.client.DeleteWorkteamWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeAlgorithm(ctx context.Context, input *sagemaker.DescribeAlgorithmInput) (*sagemaker.DescribeAlgorithmOutput, error) {
	return a.client.DescribeAlgorithmWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeApp(ctx context.Context, input *sagemaker.DescribeAppInput) (*sagemaker.DescribeAppOutput, error) {
	return a.client.DescribeAppWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeAutoMLJob(ctx context.Context, input *sagemaker.DescribeAutoMLJobInput) (*sagemaker.DescribeAutoMLJobOutput, error) {
	return a.client.DescribeAutoMLJobWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeCodeRepository(ctx context.Context, input *sagemaker.DescribeCodeRepositoryInput) (*sagemaker.DescribeCodeRepositoryOutput, error) {
	return a.client.DescribeCodeRepositoryWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeCompilationJob(ctx context.Context, input *sagemaker.DescribeCompilationJobInput) (*sagemaker.DescribeCompilationJobOutput, error) {
	return a.client.DescribeCompilationJobWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeDomain(ctx context.Context, input *sagemaker.DescribeDomainInput) (*sagemaker.DescribeDomainOutput, error) {
	return a.client.DescribeDomainWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeEndpoint(ctx context.Context, input *sagemaker.DescribeEndpointInput) (*sagemaker.DescribeEndpointOutput, error) {
	return a.client.DescribeEndpointWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeEndpointConfig(ctx context.Context, input *sagemaker.DescribeEndpointConfigInput) (*sagemaker.DescribeEndpointConfigOutput, error) {
	return a.client.DescribeEndpointConfigWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeExperiment(ctx context.Context, input *sagemaker.DescribeExperimentInput) (*sagemaker.DescribeExperimentOutput, error) {
	return a.client.DescribeExperimentWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeFlowDefinition(ctx context.Context, input *sagemaker.DescribeFlowDefinitionInput) (*sagemaker.DescribeFlowDefinitionOutput, error) {
	return a.client.DescribeFlowDefinitionWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeHumanTaskUi(ctx context.Context, input *sagemaker.DescribeHumanTaskUiInput) (*sagemaker.DescribeHumanTaskUiOutput, error) {
	return a.client.DescribeHumanTaskUiWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeHyperParameterTuningJob(ctx context.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) (*sagemaker.DescribeHyperParameterTuningJobOutput, error) {
	return a.client.DescribeHyperParameterTuningJobWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeLabelingJob(ctx context.Context, input *sagemaker.DescribeLabelingJobInput) (*sagemaker.DescribeLabelingJobOutput, error) {
	return a.client.DescribeLabelingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeModel(ctx context.Context, input *sagemaker.DescribeModelInput) (*sagemaker.DescribeModelOutput, error) {
	return a.client.DescribeModelWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeModelPackage(ctx context.Context, input *sagemaker.DescribeModelPackageInput) (*sagemaker.DescribeModelPackageOutput, error) {
	return a.client.DescribeModelPackageWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeMonitoringSchedule(ctx context.Context, input *sagemaker.DescribeMonitoringScheduleInput) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
	return a.client.DescribeMonitoringScheduleWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeNotebookInstance(ctx context.Context, input *sagemaker.DescribeNotebookInstanceInput) (*sagemaker.DescribeNotebookInstanceOutput, error) {
	return a.client.DescribeNotebookInstanceWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeNotebookInstanceLifecycleConfig(ctx context.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error) {
	return a.client.DescribeNotebookInstanceLifecycleConfigWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeProcessingJob(ctx context.Context, input *sagemaker.DescribeProcessingJobInput) (*sagemaker.DescribeProcessingJobOutput, error) {
	return a.client.DescribeProcessingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeSubscribedWorkteam(ctx context.Context, input *sagemaker.DescribeSubscribedWorkteamInput) (*sagemaker.DescribeSubscribedWorkteamOutput, error) {
	return a.client.DescribeSubscribedWorkteamWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeTrainingJob(ctx context.Context, input *sagemaker.DescribeTrainingJobInput) (*sagemaker.DescribeTrainingJobOutput, error) {
	return a.client.DescribeTrainingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeTransformJob(ctx context.Context, input *sagemaker.DescribeTransformJobInput) (*sagemaker.DescribeTransformJobOutput, error) {
	return a.client.DescribeTransformJobWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeTrial(ctx context.Context, input *sagemaker.DescribeTrialInput) (*sagemaker.DescribeTrialOutput, error) {
	return a.client.DescribeTrialWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeTrialComponent(ctx context.Context, input *sagemaker.DescribeTrialComponentInput) (*sagemaker.DescribeTrialComponentOutput, error) {
	return a.client.DescribeTrialComponentWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeUserProfile(ctx context.Context, input *sagemaker.DescribeUserProfileInput) (*sagemaker.DescribeUserProfileOutput, error) {
	return a.client.DescribeUserProfileWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeWorkforce(ctx context.Context, input *sagemaker.DescribeWorkforceInput) (*sagemaker.DescribeWorkforceOutput, error) {
	return a.client.DescribeWorkforceWithContext(ctx, input)
}

func (a *SageMakerActivities) DescribeWorkteam(ctx context.Context, input *sagemaker.DescribeWorkteamInput) (*sagemaker.DescribeWorkteamOutput, error) {
	return a.client.DescribeWorkteamWithContext(ctx, input)
}

func (a *SageMakerActivities) DisassociateTrialComponent(ctx context.Context, input *sagemaker.DisassociateTrialComponentInput) (*sagemaker.DisassociateTrialComponentOutput, error) {
	return a.client.DisassociateTrialComponentWithContext(ctx, input)
}

func (a *SageMakerActivities) GetSearchSuggestions(ctx context.Context, input *sagemaker.GetSearchSuggestionsInput) (*sagemaker.GetSearchSuggestionsOutput, error) {
	return a.client.GetSearchSuggestionsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListAlgorithms(ctx context.Context, input *sagemaker.ListAlgorithmsInput) (*sagemaker.ListAlgorithmsOutput, error) {
	return a.client.ListAlgorithmsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListApps(ctx context.Context, input *sagemaker.ListAppsInput) (*sagemaker.ListAppsOutput, error) {
	return a.client.ListAppsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListAutoMLJobs(ctx context.Context, input *sagemaker.ListAutoMLJobsInput) (*sagemaker.ListAutoMLJobsOutput, error) {
	return a.client.ListAutoMLJobsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListCandidatesForAutoMLJob(ctx context.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) (*sagemaker.ListCandidatesForAutoMLJobOutput, error) {
	return a.client.ListCandidatesForAutoMLJobWithContext(ctx, input)
}

func (a *SageMakerActivities) ListCodeRepositories(ctx context.Context, input *sagemaker.ListCodeRepositoriesInput) (*sagemaker.ListCodeRepositoriesOutput, error) {
	return a.client.ListCodeRepositoriesWithContext(ctx, input)
}

func (a *SageMakerActivities) ListCompilationJobs(ctx context.Context, input *sagemaker.ListCompilationJobsInput) (*sagemaker.ListCompilationJobsOutput, error) {
	return a.client.ListCompilationJobsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListDomains(ctx context.Context, input *sagemaker.ListDomainsInput) (*sagemaker.ListDomainsOutput, error) {
	return a.client.ListDomainsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListEndpointConfigs(ctx context.Context, input *sagemaker.ListEndpointConfigsInput) (*sagemaker.ListEndpointConfigsOutput, error) {
	return a.client.ListEndpointConfigsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListEndpoints(ctx context.Context, input *sagemaker.ListEndpointsInput) (*sagemaker.ListEndpointsOutput, error) {
	return a.client.ListEndpointsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListExperiments(ctx context.Context, input *sagemaker.ListExperimentsInput) (*sagemaker.ListExperimentsOutput, error) {
	return a.client.ListExperimentsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListFlowDefinitions(ctx context.Context, input *sagemaker.ListFlowDefinitionsInput) (*sagemaker.ListFlowDefinitionsOutput, error) {
	return a.client.ListFlowDefinitionsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListHumanTaskUis(ctx context.Context, input *sagemaker.ListHumanTaskUisInput) (*sagemaker.ListHumanTaskUisOutput, error) {
	return a.client.ListHumanTaskUisWithContext(ctx, input)
}

func (a *SageMakerActivities) ListHyperParameterTuningJobs(ctx context.Context, input *sagemaker.ListHyperParameterTuningJobsInput) (*sagemaker.ListHyperParameterTuningJobsOutput, error) {
	return a.client.ListHyperParameterTuningJobsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListLabelingJobs(ctx context.Context, input *sagemaker.ListLabelingJobsInput) (*sagemaker.ListLabelingJobsOutput, error) {
	return a.client.ListLabelingJobsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListLabelingJobsForWorkteam(ctx context.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) (*sagemaker.ListLabelingJobsForWorkteamOutput, error) {
	return a.client.ListLabelingJobsForWorkteamWithContext(ctx, input)
}

func (a *SageMakerActivities) ListModelPackages(ctx context.Context, input *sagemaker.ListModelPackagesInput) (*sagemaker.ListModelPackagesOutput, error) {
	return a.client.ListModelPackagesWithContext(ctx, input)
}

func (a *SageMakerActivities) ListModels(ctx context.Context, input *sagemaker.ListModelsInput) (*sagemaker.ListModelsOutput, error) {
	return a.client.ListModelsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListMonitoringExecutions(ctx context.Context, input *sagemaker.ListMonitoringExecutionsInput) (*sagemaker.ListMonitoringExecutionsOutput, error) {
	return a.client.ListMonitoringExecutionsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListMonitoringSchedules(ctx context.Context, input *sagemaker.ListMonitoringSchedulesInput) (*sagemaker.ListMonitoringSchedulesOutput, error) {
	return a.client.ListMonitoringSchedulesWithContext(ctx, input)
}

func (a *SageMakerActivities) ListNotebookInstanceLifecycleConfigs(ctx context.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error) {
	return a.client.ListNotebookInstanceLifecycleConfigsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListNotebookInstances(ctx context.Context, input *sagemaker.ListNotebookInstancesInput) (*sagemaker.ListNotebookInstancesOutput, error) {
	return a.client.ListNotebookInstancesWithContext(ctx, input)
}

func (a *SageMakerActivities) ListProcessingJobs(ctx context.Context, input *sagemaker.ListProcessingJobsInput) (*sagemaker.ListProcessingJobsOutput, error) {
	return a.client.ListProcessingJobsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListSubscribedWorkteams(ctx context.Context, input *sagemaker.ListSubscribedWorkteamsInput) (*sagemaker.ListSubscribedWorkteamsOutput, error) {
	return a.client.ListSubscribedWorkteamsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListTags(ctx context.Context, input *sagemaker.ListTagsInput) (*sagemaker.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListTrainingJobs(ctx context.Context, input *sagemaker.ListTrainingJobsInput) (*sagemaker.ListTrainingJobsOutput, error) {
	return a.client.ListTrainingJobsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListTrainingJobsForHyperParameterTuningJob(ctx context.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error) {
	return a.client.ListTrainingJobsForHyperParameterTuningJobWithContext(ctx, input)
}

func (a *SageMakerActivities) ListTransformJobs(ctx context.Context, input *sagemaker.ListTransformJobsInput) (*sagemaker.ListTransformJobsOutput, error) {
	return a.client.ListTransformJobsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListTrialComponents(ctx context.Context, input *sagemaker.ListTrialComponentsInput) (*sagemaker.ListTrialComponentsOutput, error) {
	return a.client.ListTrialComponentsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListTrials(ctx context.Context, input *sagemaker.ListTrialsInput) (*sagemaker.ListTrialsOutput, error) {
	return a.client.ListTrialsWithContext(ctx, input)
}

func (a *SageMakerActivities) ListUserProfiles(ctx context.Context, input *sagemaker.ListUserProfilesInput) (*sagemaker.ListUserProfilesOutput, error) {
	return a.client.ListUserProfilesWithContext(ctx, input)
}

func (a *SageMakerActivities) ListWorkforces(ctx context.Context, input *sagemaker.ListWorkforcesInput) (*sagemaker.ListWorkforcesOutput, error) {
	return a.client.ListWorkforcesWithContext(ctx, input)
}

func (a *SageMakerActivities) ListWorkteams(ctx context.Context, input *sagemaker.ListWorkteamsInput) (*sagemaker.ListWorkteamsOutput, error) {
	return a.client.ListWorkteamsWithContext(ctx, input)
}

func (a *SageMakerActivities) RenderUiTemplate(ctx context.Context, input *sagemaker.RenderUiTemplateInput) (*sagemaker.RenderUiTemplateOutput, error) {
	return a.client.RenderUiTemplateWithContext(ctx, input)
}

func (a *SageMakerActivities) Search(ctx context.Context, input *sagemaker.SearchInput) (*sagemaker.SearchOutput, error) {
	return a.client.SearchWithContext(ctx, input)
}

func (a *SageMakerActivities) StartMonitoringSchedule(ctx context.Context, input *sagemaker.StartMonitoringScheduleInput) (*sagemaker.StartMonitoringScheduleOutput, error) {
	return a.client.StartMonitoringScheduleWithContext(ctx, input)
}

func (a *SageMakerActivities) StartNotebookInstance(ctx context.Context, input *sagemaker.StartNotebookInstanceInput) (*sagemaker.StartNotebookInstanceOutput, error) {
	return a.client.StartNotebookInstanceWithContext(ctx, input)
}

func (a *SageMakerActivities) StopAutoMLJob(ctx context.Context, input *sagemaker.StopAutoMLJobInput) (*sagemaker.StopAutoMLJobOutput, error) {
	return a.client.StopAutoMLJobWithContext(ctx, input)
}

func (a *SageMakerActivities) StopCompilationJob(ctx context.Context, input *sagemaker.StopCompilationJobInput) (*sagemaker.StopCompilationJobOutput, error) {
	return a.client.StopCompilationJobWithContext(ctx, input)
}

func (a *SageMakerActivities) StopHyperParameterTuningJob(ctx context.Context, input *sagemaker.StopHyperParameterTuningJobInput) (*sagemaker.StopHyperParameterTuningJobOutput, error) {
	return a.client.StopHyperParameterTuningJobWithContext(ctx, input)
}

func (a *SageMakerActivities) StopLabelingJob(ctx context.Context, input *sagemaker.StopLabelingJobInput) (*sagemaker.StopLabelingJobOutput, error) {
	return a.client.StopLabelingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) StopMonitoringSchedule(ctx context.Context, input *sagemaker.StopMonitoringScheduleInput) (*sagemaker.StopMonitoringScheduleOutput, error) {
	return a.client.StopMonitoringScheduleWithContext(ctx, input)
}

func (a *SageMakerActivities) StopNotebookInstance(ctx context.Context, input *sagemaker.StopNotebookInstanceInput) (*sagemaker.StopNotebookInstanceOutput, error) {
	return a.client.StopNotebookInstanceWithContext(ctx, input)
}

func (a *SageMakerActivities) StopProcessingJob(ctx context.Context, input *sagemaker.StopProcessingJobInput) (*sagemaker.StopProcessingJobOutput, error) {
	return a.client.StopProcessingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) StopTrainingJob(ctx context.Context, input *sagemaker.StopTrainingJobInput) (*sagemaker.StopTrainingJobOutput, error) {
	return a.client.StopTrainingJobWithContext(ctx, input)
}

func (a *SageMakerActivities) StopTransformJob(ctx context.Context, input *sagemaker.StopTransformJobInput) (*sagemaker.StopTransformJobOutput, error) {
	return a.client.StopTransformJobWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateCodeRepository(ctx context.Context, input *sagemaker.UpdateCodeRepositoryInput) (*sagemaker.UpdateCodeRepositoryOutput, error) {
	return a.client.UpdateCodeRepositoryWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateDomain(ctx context.Context, input *sagemaker.UpdateDomainInput) (*sagemaker.UpdateDomainOutput, error) {
	return a.client.UpdateDomainWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateEndpoint(ctx context.Context, input *sagemaker.UpdateEndpointInput) (*sagemaker.UpdateEndpointOutput, error) {
	return a.client.UpdateEndpointWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateEndpointWeightsAndCapacities(ctx context.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) (*sagemaker.UpdateEndpointWeightsAndCapacitiesOutput, error) {
	return a.client.UpdateEndpointWeightsAndCapacitiesWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateExperiment(ctx context.Context, input *sagemaker.UpdateExperimentInput) (*sagemaker.UpdateExperimentOutput, error) {
	return a.client.UpdateExperimentWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateMonitoringSchedule(ctx context.Context, input *sagemaker.UpdateMonitoringScheduleInput) (*sagemaker.UpdateMonitoringScheduleOutput, error) {
	return a.client.UpdateMonitoringScheduleWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateNotebookInstance(ctx context.Context, input *sagemaker.UpdateNotebookInstanceInput) (*sagemaker.UpdateNotebookInstanceOutput, error) {
	return a.client.UpdateNotebookInstanceWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateNotebookInstanceLifecycleConfig(ctx context.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) (*sagemaker.UpdateNotebookInstanceLifecycleConfigOutput, error) {
	return a.client.UpdateNotebookInstanceLifecycleConfigWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateTrial(ctx context.Context, input *sagemaker.UpdateTrialInput) (*sagemaker.UpdateTrialOutput, error) {
	return a.client.UpdateTrialWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateTrialComponent(ctx context.Context, input *sagemaker.UpdateTrialComponentInput) (*sagemaker.UpdateTrialComponentOutput, error) {
	return a.client.UpdateTrialComponentWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateUserProfile(ctx context.Context, input *sagemaker.UpdateUserProfileInput) (*sagemaker.UpdateUserProfileOutput, error) {
	return a.client.UpdateUserProfileWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateWorkforce(ctx context.Context, input *sagemaker.UpdateWorkforceInput) (*sagemaker.UpdateWorkforceOutput, error) {
	return a.client.UpdateWorkforceWithContext(ctx, input)
}

func (a *SageMakerActivities) UpdateWorkteam(ctx context.Context, input *sagemaker.UpdateWorkteamInput) (*sagemaker.UpdateWorkteamOutput, error) {
	return a.client.UpdateWorkteamWithContext(ctx, input)
}

func (a *SageMakerActivities) WaitUntilEndpointDeleted(ctx context.Context, input *sagemaker.DescribeEndpointInput) error {
	return a.client.WaitUntilEndpointDeletedWithContext(ctx, input)

}

func (a *SageMakerActivities) WaitUntilEndpointInService(ctx context.Context, input *sagemaker.DescribeEndpointInput) error {
	return a.client.WaitUntilEndpointInServiceWithContext(ctx, input)

}

func (a *SageMakerActivities) WaitUntilNotebookInstanceDeleted(ctx context.Context, input *sagemaker.DescribeNotebookInstanceInput) error {
	return a.client.WaitUntilNotebookInstanceDeletedWithContext(ctx, input)

}

func (a *SageMakerActivities) WaitUntilNotebookInstanceInService(ctx context.Context, input *sagemaker.DescribeNotebookInstanceInput) error {
	return a.client.WaitUntilNotebookInstanceInServiceWithContext(ctx, input)

}

func (a *SageMakerActivities) WaitUntilNotebookInstanceStopped(ctx context.Context, input *sagemaker.DescribeNotebookInstanceInput) error {
	return a.client.WaitUntilNotebookInstanceStoppedWithContext(ctx, input)

}

func (a *SageMakerActivities) WaitUntilProcessingJobCompletedOrStopped(ctx context.Context, input *sagemaker.DescribeProcessingJobInput) error {
	return a.client.WaitUntilProcessingJobCompletedOrStoppedWithContext(ctx, input)

}

func (a *SageMakerActivities) WaitUntilTrainingJobCompletedOrStopped(ctx context.Context, input *sagemaker.DescribeTrainingJobInput) error {
	return a.client.WaitUntilTrainingJobCompletedOrStoppedWithContext(ctx, input)

}

func (a *SageMakerActivities) WaitUntilTransformJobCompletedOrStopped(ctx context.Context, input *sagemaker.DescribeTransformJobInput) error {
	return a.client.WaitUntilTransformJobCompletedOrStoppedWithContext(ctx, input)

}
