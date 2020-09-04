
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemaker/sagemakeriface"
)

type SageMakerActivities struct {
	client sagemakeriface.SageMakerAPI
}

func NewSageMakerActivities(client sagemakeriface.SageMakerAPI) *SageMakerActivities {
    return &SageMakerActivities{client: client}
}

func (a *SageMakerActivities) AddTags(input *sagemaker.AddTagsInput) (*sagemaker.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *SageMakerActivities) AssociateTrialComponent(input *sagemaker.AssociateTrialComponentInput) (*sagemaker.AssociateTrialComponentOutput, error) {
    return a.client.AssociateTrialComponent(input)
}

func (a *SageMakerActivities) CreateAlgorithm(input *sagemaker.CreateAlgorithmInput) (*sagemaker.CreateAlgorithmOutput, error) {
    return a.client.CreateAlgorithm(input)
}

func (a *SageMakerActivities) CreateApp(input *sagemaker.CreateAppInput) (*sagemaker.CreateAppOutput, error) {
    return a.client.CreateApp(input)
}

func (a *SageMakerActivities) CreateAutoMLJob(input *sagemaker.CreateAutoMLJobInput) (*sagemaker.CreateAutoMLJobOutput, error) {
    return a.client.CreateAutoMLJob(input)
}

func (a *SageMakerActivities) CreateCodeRepository(input *sagemaker.CreateCodeRepositoryInput) (*sagemaker.CreateCodeRepositoryOutput, error) {
    return a.client.CreateCodeRepository(input)
}

func (a *SageMakerActivities) CreateCompilationJob(input *sagemaker.CreateCompilationJobInput) (*sagemaker.CreateCompilationJobOutput, error) {
    return a.client.CreateCompilationJob(input)
}

func (a *SageMakerActivities) CreateDomain(input *sagemaker.CreateDomainInput) (*sagemaker.CreateDomainOutput, error) {
    return a.client.CreateDomain(input)
}

func (a *SageMakerActivities) CreateEndpoint(input *sagemaker.CreateEndpointInput) (*sagemaker.CreateEndpointOutput, error) {
    return a.client.CreateEndpoint(input)
}

func (a *SageMakerActivities) CreateEndpointConfig(input *sagemaker.CreateEndpointConfigInput) (*sagemaker.CreateEndpointConfigOutput, error) {
    return a.client.CreateEndpointConfig(input)
}

func (a *SageMakerActivities) CreateExperiment(input *sagemaker.CreateExperimentInput) (*sagemaker.CreateExperimentOutput, error) {
    return a.client.CreateExperiment(input)
}

func (a *SageMakerActivities) CreateFlowDefinition(input *sagemaker.CreateFlowDefinitionInput) (*sagemaker.CreateFlowDefinitionOutput, error) {
    return a.client.CreateFlowDefinition(input)
}

func (a *SageMakerActivities) CreateHumanTaskUi(input *sagemaker.CreateHumanTaskUiInput) (*sagemaker.CreateHumanTaskUiOutput, error) {
    return a.client.CreateHumanTaskUi(input)
}

func (a *SageMakerActivities) CreateHyperParameterTuningJob(input *sagemaker.CreateHyperParameterTuningJobInput) (*sagemaker.CreateHyperParameterTuningJobOutput, error) {
    return a.client.CreateHyperParameterTuningJob(input)
}

func (a *SageMakerActivities) CreateLabelingJob(input *sagemaker.CreateLabelingJobInput) (*sagemaker.CreateLabelingJobOutput, error) {
    return a.client.CreateLabelingJob(input)
}

func (a *SageMakerActivities) CreateModel(input *sagemaker.CreateModelInput) (*sagemaker.CreateModelOutput, error) {
    return a.client.CreateModel(input)
}

func (a *SageMakerActivities) CreateModelPackage(input *sagemaker.CreateModelPackageInput) (*sagemaker.CreateModelPackageOutput, error) {
    return a.client.CreateModelPackage(input)
}

func (a *SageMakerActivities) CreateMonitoringSchedule(input *sagemaker.CreateMonitoringScheduleInput) (*sagemaker.CreateMonitoringScheduleOutput, error) {
    return a.client.CreateMonitoringSchedule(input)
}

func (a *SageMakerActivities) CreateNotebookInstance(input *sagemaker.CreateNotebookInstanceInput) (*sagemaker.CreateNotebookInstanceOutput, error) {
    return a.client.CreateNotebookInstance(input)
}

func (a *SageMakerActivities) CreateNotebookInstanceLifecycleConfig(input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) (*sagemaker.CreateNotebookInstanceLifecycleConfigOutput, error) {
    return a.client.CreateNotebookInstanceLifecycleConfig(input)
}

func (a *SageMakerActivities) CreatePresignedDomainUrl(input *sagemaker.CreatePresignedDomainUrlInput) (*sagemaker.CreatePresignedDomainUrlOutput, error) {
    return a.client.CreatePresignedDomainUrl(input)
}

func (a *SageMakerActivities) CreatePresignedNotebookInstanceUrl(input *sagemaker.CreatePresignedNotebookInstanceUrlInput) (*sagemaker.CreatePresignedNotebookInstanceUrlOutput, error) {
    return a.client.CreatePresignedNotebookInstanceUrl(input)
}

func (a *SageMakerActivities) CreateProcessingJob(input *sagemaker.CreateProcessingJobInput) (*sagemaker.CreateProcessingJobOutput, error) {
    return a.client.CreateProcessingJob(input)
}

func (a *SageMakerActivities) CreateTrainingJob(input *sagemaker.CreateTrainingJobInput) (*sagemaker.CreateTrainingJobOutput, error) {
    return a.client.CreateTrainingJob(input)
}

func (a *SageMakerActivities) CreateTransformJob(input *sagemaker.CreateTransformJobInput) (*sagemaker.CreateTransformJobOutput, error) {
    return a.client.CreateTransformJob(input)
}

func (a *SageMakerActivities) CreateTrial(input *sagemaker.CreateTrialInput) (*sagemaker.CreateTrialOutput, error) {
    return a.client.CreateTrial(input)
}

func (a *SageMakerActivities) CreateTrialComponent(input *sagemaker.CreateTrialComponentInput) (*sagemaker.CreateTrialComponentOutput, error) {
    return a.client.CreateTrialComponent(input)
}

func (a *SageMakerActivities) CreateUserProfile(input *sagemaker.CreateUserProfileInput) (*sagemaker.CreateUserProfileOutput, error) {
    return a.client.CreateUserProfile(input)
}

func (a *SageMakerActivities) CreateWorkforce(input *sagemaker.CreateWorkforceInput) (*sagemaker.CreateWorkforceOutput, error) {
    return a.client.CreateWorkforce(input)
}

func (a *SageMakerActivities) CreateWorkteam(input *sagemaker.CreateWorkteamInput) (*sagemaker.CreateWorkteamOutput, error) {
    return a.client.CreateWorkteam(input)
}

func (a *SageMakerActivities) DeleteAlgorithm(input *sagemaker.DeleteAlgorithmInput) (*sagemaker.DeleteAlgorithmOutput, error) {
    return a.client.DeleteAlgorithm(input)
}

func (a *SageMakerActivities) DeleteApp(input *sagemaker.DeleteAppInput) (*sagemaker.DeleteAppOutput, error) {
    return a.client.DeleteApp(input)
}

func (a *SageMakerActivities) DeleteCodeRepository(input *sagemaker.DeleteCodeRepositoryInput) (*sagemaker.DeleteCodeRepositoryOutput, error) {
    return a.client.DeleteCodeRepository(input)
}

func (a *SageMakerActivities) DeleteDomain(input *sagemaker.DeleteDomainInput) (*sagemaker.DeleteDomainOutput, error) {
    return a.client.DeleteDomain(input)
}

func (a *SageMakerActivities) DeleteEndpoint(input *sagemaker.DeleteEndpointInput) (*sagemaker.DeleteEndpointOutput, error) {
    return a.client.DeleteEndpoint(input)
}

func (a *SageMakerActivities) DeleteEndpointConfig(input *sagemaker.DeleteEndpointConfigInput) (*sagemaker.DeleteEndpointConfigOutput, error) {
    return a.client.DeleteEndpointConfig(input)
}

func (a *SageMakerActivities) DeleteExperiment(input *sagemaker.DeleteExperimentInput) (*sagemaker.DeleteExperimentOutput, error) {
    return a.client.DeleteExperiment(input)
}

func (a *SageMakerActivities) DeleteFlowDefinition(input *sagemaker.DeleteFlowDefinitionInput) (*sagemaker.DeleteFlowDefinitionOutput, error) {
    return a.client.DeleteFlowDefinition(input)
}

func (a *SageMakerActivities) DeleteHumanTaskUi(input *sagemaker.DeleteHumanTaskUiInput) (*sagemaker.DeleteHumanTaskUiOutput, error) {
    return a.client.DeleteHumanTaskUi(input)
}

func (a *SageMakerActivities) DeleteModel(input *sagemaker.DeleteModelInput) (*sagemaker.DeleteModelOutput, error) {
    return a.client.DeleteModel(input)
}

func (a *SageMakerActivities) DeleteModelPackage(input *sagemaker.DeleteModelPackageInput) (*sagemaker.DeleteModelPackageOutput, error) {
    return a.client.DeleteModelPackage(input)
}

func (a *SageMakerActivities) DeleteMonitoringSchedule(input *sagemaker.DeleteMonitoringScheduleInput) (*sagemaker.DeleteMonitoringScheduleOutput, error) {
    return a.client.DeleteMonitoringSchedule(input)
}

func (a *SageMakerActivities) DeleteNotebookInstance(input *sagemaker.DeleteNotebookInstanceInput) (*sagemaker.DeleteNotebookInstanceOutput, error) {
    return a.client.DeleteNotebookInstance(input)
}

func (a *SageMakerActivities) DeleteNotebookInstanceLifecycleConfig(input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) (*sagemaker.DeleteNotebookInstanceLifecycleConfigOutput, error) {
    return a.client.DeleteNotebookInstanceLifecycleConfig(input)
}

func (a *SageMakerActivities) DeleteTags(input *sagemaker.DeleteTagsInput) (*sagemaker.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}

func (a *SageMakerActivities) DeleteTrial(input *sagemaker.DeleteTrialInput) (*sagemaker.DeleteTrialOutput, error) {
    return a.client.DeleteTrial(input)
}

func (a *SageMakerActivities) DeleteTrialComponent(input *sagemaker.DeleteTrialComponentInput) (*sagemaker.DeleteTrialComponentOutput, error) {
    return a.client.DeleteTrialComponent(input)
}

func (a *SageMakerActivities) DeleteUserProfile(input *sagemaker.DeleteUserProfileInput) (*sagemaker.DeleteUserProfileOutput, error) {
    return a.client.DeleteUserProfile(input)
}

func (a *SageMakerActivities) DeleteWorkforce(input *sagemaker.DeleteWorkforceInput) (*sagemaker.DeleteWorkforceOutput, error) {
    return a.client.DeleteWorkforce(input)
}

func (a *SageMakerActivities) DeleteWorkteam(input *sagemaker.DeleteWorkteamInput) (*sagemaker.DeleteWorkteamOutput, error) {
    return a.client.DeleteWorkteam(input)
}

func (a *SageMakerActivities) DescribeAlgorithm(input *sagemaker.DescribeAlgorithmInput) (*sagemaker.DescribeAlgorithmOutput, error) {
    return a.client.DescribeAlgorithm(input)
}

func (a *SageMakerActivities) DescribeApp(input *sagemaker.DescribeAppInput) (*sagemaker.DescribeAppOutput, error) {
    return a.client.DescribeApp(input)
}

func (a *SageMakerActivities) DescribeAutoMLJob(input *sagemaker.DescribeAutoMLJobInput) (*sagemaker.DescribeAutoMLJobOutput, error) {
    return a.client.DescribeAutoMLJob(input)
}

func (a *SageMakerActivities) DescribeCodeRepository(input *sagemaker.DescribeCodeRepositoryInput) (*sagemaker.DescribeCodeRepositoryOutput, error) {
    return a.client.DescribeCodeRepository(input)
}

func (a *SageMakerActivities) DescribeCompilationJob(input *sagemaker.DescribeCompilationJobInput) (*sagemaker.DescribeCompilationJobOutput, error) {
    return a.client.DescribeCompilationJob(input)
}

func (a *SageMakerActivities) DescribeDomain(input *sagemaker.DescribeDomainInput) (*sagemaker.DescribeDomainOutput, error) {
    return a.client.DescribeDomain(input)
}

func (a *SageMakerActivities) DescribeEndpoint(input *sagemaker.DescribeEndpointInput) (*sagemaker.DescribeEndpointOutput, error) {
    return a.client.DescribeEndpoint(input)
}

func (a *SageMakerActivities) DescribeEndpointConfig(input *sagemaker.DescribeEndpointConfigInput) (*sagemaker.DescribeEndpointConfigOutput, error) {
    return a.client.DescribeEndpointConfig(input)
}

func (a *SageMakerActivities) DescribeExperiment(input *sagemaker.DescribeExperimentInput) (*sagemaker.DescribeExperimentOutput, error) {
    return a.client.DescribeExperiment(input)
}

func (a *SageMakerActivities) DescribeFlowDefinition(input *sagemaker.DescribeFlowDefinitionInput) (*sagemaker.DescribeFlowDefinitionOutput, error) {
    return a.client.DescribeFlowDefinition(input)
}

func (a *SageMakerActivities) DescribeHumanTaskUi(input *sagemaker.DescribeHumanTaskUiInput) (*sagemaker.DescribeHumanTaskUiOutput, error) {
    return a.client.DescribeHumanTaskUi(input)
}

func (a *SageMakerActivities) DescribeHyperParameterTuningJob(input *sagemaker.DescribeHyperParameterTuningJobInput) (*sagemaker.DescribeHyperParameterTuningJobOutput, error) {
    return a.client.DescribeHyperParameterTuningJob(input)
}

func (a *SageMakerActivities) DescribeLabelingJob(input *sagemaker.DescribeLabelingJobInput) (*sagemaker.DescribeLabelingJobOutput, error) {
    return a.client.DescribeLabelingJob(input)
}

func (a *SageMakerActivities) DescribeModel(input *sagemaker.DescribeModelInput) (*sagemaker.DescribeModelOutput, error) {
    return a.client.DescribeModel(input)
}

func (a *SageMakerActivities) DescribeModelPackage(input *sagemaker.DescribeModelPackageInput) (*sagemaker.DescribeModelPackageOutput, error) {
    return a.client.DescribeModelPackage(input)
}

func (a *SageMakerActivities) DescribeMonitoringSchedule(input *sagemaker.DescribeMonitoringScheduleInput) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
    return a.client.DescribeMonitoringSchedule(input)
}

func (a *SageMakerActivities) DescribeNotebookInstance(input *sagemaker.DescribeNotebookInstanceInput) (*sagemaker.DescribeNotebookInstanceOutput, error) {
    return a.client.DescribeNotebookInstance(input)
}

func (a *SageMakerActivities) DescribeNotebookInstanceLifecycleConfig(input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error) {
    return a.client.DescribeNotebookInstanceLifecycleConfig(input)
}

func (a *SageMakerActivities) DescribeProcessingJob(input *sagemaker.DescribeProcessingJobInput) (*sagemaker.DescribeProcessingJobOutput, error) {
    return a.client.DescribeProcessingJob(input)
}

func (a *SageMakerActivities) DescribeSubscribedWorkteam(input *sagemaker.DescribeSubscribedWorkteamInput) (*sagemaker.DescribeSubscribedWorkteamOutput, error) {
    return a.client.DescribeSubscribedWorkteam(input)
}

func (a *SageMakerActivities) DescribeTrainingJob(input *sagemaker.DescribeTrainingJobInput) (*sagemaker.DescribeTrainingJobOutput, error) {
    return a.client.DescribeTrainingJob(input)
}

func (a *SageMakerActivities) DescribeTransformJob(input *sagemaker.DescribeTransformJobInput) (*sagemaker.DescribeTransformJobOutput, error) {
    return a.client.DescribeTransformJob(input)
}

func (a *SageMakerActivities) DescribeTrial(input *sagemaker.DescribeTrialInput) (*sagemaker.DescribeTrialOutput, error) {
    return a.client.DescribeTrial(input)
}

func (a *SageMakerActivities) DescribeTrialComponent(input *sagemaker.DescribeTrialComponentInput) (*sagemaker.DescribeTrialComponentOutput, error) {
    return a.client.DescribeTrialComponent(input)
}

func (a *SageMakerActivities) DescribeUserProfile(input *sagemaker.DescribeUserProfileInput) (*sagemaker.DescribeUserProfileOutput, error) {
    return a.client.DescribeUserProfile(input)
}

func (a *SageMakerActivities) DescribeWorkforce(input *sagemaker.DescribeWorkforceInput) (*sagemaker.DescribeWorkforceOutput, error) {
    return a.client.DescribeWorkforce(input)
}

func (a *SageMakerActivities) DescribeWorkteam(input *sagemaker.DescribeWorkteamInput) (*sagemaker.DescribeWorkteamOutput, error) {
    return a.client.DescribeWorkteam(input)
}

func (a *SageMakerActivities) DisassociateTrialComponent(input *sagemaker.DisassociateTrialComponentInput) (*sagemaker.DisassociateTrialComponentOutput, error) {
    return a.client.DisassociateTrialComponent(input)
}

func (a *SageMakerActivities) GetSearchSuggestions(input *sagemaker.GetSearchSuggestionsInput) (*sagemaker.GetSearchSuggestionsOutput, error) {
    return a.client.GetSearchSuggestions(input)
}

func (a *SageMakerActivities) ListAlgorithms(input *sagemaker.ListAlgorithmsInput) (*sagemaker.ListAlgorithmsOutput, error) {
    return a.client.ListAlgorithms(input)
}

func (a *SageMakerActivities) ListApps(input *sagemaker.ListAppsInput) (*sagemaker.ListAppsOutput, error) {
    return a.client.ListApps(input)
}

func (a *SageMakerActivities) ListAutoMLJobs(input *sagemaker.ListAutoMLJobsInput) (*sagemaker.ListAutoMLJobsOutput, error) {
    return a.client.ListAutoMLJobs(input)
}

func (a *SageMakerActivities) ListCandidatesForAutoMLJob(input *sagemaker.ListCandidatesForAutoMLJobInput) (*sagemaker.ListCandidatesForAutoMLJobOutput, error) {
    return a.client.ListCandidatesForAutoMLJob(input)
}

func (a *SageMakerActivities) ListCodeRepositories(input *sagemaker.ListCodeRepositoriesInput) (*sagemaker.ListCodeRepositoriesOutput, error) {
    return a.client.ListCodeRepositories(input)
}

func (a *SageMakerActivities) ListCompilationJobs(input *sagemaker.ListCompilationJobsInput) (*sagemaker.ListCompilationJobsOutput, error) {
    return a.client.ListCompilationJobs(input)
}

func (a *SageMakerActivities) ListDomains(input *sagemaker.ListDomainsInput) (*sagemaker.ListDomainsOutput, error) {
    return a.client.ListDomains(input)
}

func (a *SageMakerActivities) ListEndpointConfigs(input *sagemaker.ListEndpointConfigsInput) (*sagemaker.ListEndpointConfigsOutput, error) {
    return a.client.ListEndpointConfigs(input)
}

func (a *SageMakerActivities) ListEndpoints(input *sagemaker.ListEndpointsInput) (*sagemaker.ListEndpointsOutput, error) {
    return a.client.ListEndpoints(input)
}

func (a *SageMakerActivities) ListExperiments(input *sagemaker.ListExperimentsInput) (*sagemaker.ListExperimentsOutput, error) {
    return a.client.ListExperiments(input)
}

func (a *SageMakerActivities) ListFlowDefinitions(input *sagemaker.ListFlowDefinitionsInput) (*sagemaker.ListFlowDefinitionsOutput, error) {
    return a.client.ListFlowDefinitions(input)
}

func (a *SageMakerActivities) ListHumanTaskUis(input *sagemaker.ListHumanTaskUisInput) (*sagemaker.ListHumanTaskUisOutput, error) {
    return a.client.ListHumanTaskUis(input)
}

func (a *SageMakerActivities) ListHyperParameterTuningJobs(input *sagemaker.ListHyperParameterTuningJobsInput) (*sagemaker.ListHyperParameterTuningJobsOutput, error) {
    return a.client.ListHyperParameterTuningJobs(input)
}

func (a *SageMakerActivities) ListLabelingJobs(input *sagemaker.ListLabelingJobsInput) (*sagemaker.ListLabelingJobsOutput, error) {
    return a.client.ListLabelingJobs(input)
}

func (a *SageMakerActivities) ListLabelingJobsForWorkteam(input *sagemaker.ListLabelingJobsForWorkteamInput) (*sagemaker.ListLabelingJobsForWorkteamOutput, error) {
    return a.client.ListLabelingJobsForWorkteam(input)
}

func (a *SageMakerActivities) ListModelPackages(input *sagemaker.ListModelPackagesInput) (*sagemaker.ListModelPackagesOutput, error) {
    return a.client.ListModelPackages(input)
}

func (a *SageMakerActivities) ListModels(input *sagemaker.ListModelsInput) (*sagemaker.ListModelsOutput, error) {
    return a.client.ListModels(input)
}

func (a *SageMakerActivities) ListMonitoringExecutions(input *sagemaker.ListMonitoringExecutionsInput) (*sagemaker.ListMonitoringExecutionsOutput, error) {
    return a.client.ListMonitoringExecutions(input)
}

func (a *SageMakerActivities) ListMonitoringSchedules(input *sagemaker.ListMonitoringSchedulesInput) (*sagemaker.ListMonitoringSchedulesOutput, error) {
    return a.client.ListMonitoringSchedules(input)
}

func (a *SageMakerActivities) ListNotebookInstanceLifecycleConfigs(input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error) {
    return a.client.ListNotebookInstanceLifecycleConfigs(input)
}

func (a *SageMakerActivities) ListNotebookInstances(input *sagemaker.ListNotebookInstancesInput) (*sagemaker.ListNotebookInstancesOutput, error) {
    return a.client.ListNotebookInstances(input)
}

func (a *SageMakerActivities) ListProcessingJobs(input *sagemaker.ListProcessingJobsInput) (*sagemaker.ListProcessingJobsOutput, error) {
    return a.client.ListProcessingJobs(input)
}

func (a *SageMakerActivities) ListSubscribedWorkteams(input *sagemaker.ListSubscribedWorkteamsInput) (*sagemaker.ListSubscribedWorkteamsOutput, error) {
    return a.client.ListSubscribedWorkteams(input)
}

func (a *SageMakerActivities) ListTags(input *sagemaker.ListTagsInput) (*sagemaker.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *SageMakerActivities) ListTrainingJobs(input *sagemaker.ListTrainingJobsInput) (*sagemaker.ListTrainingJobsOutput, error) {
    return a.client.ListTrainingJobs(input)
}

func (a *SageMakerActivities) ListTrainingJobsForHyperParameterTuningJob(input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error) {
    return a.client.ListTrainingJobsForHyperParameterTuningJob(input)
}

func (a *SageMakerActivities) ListTransformJobs(input *sagemaker.ListTransformJobsInput) (*sagemaker.ListTransformJobsOutput, error) {
    return a.client.ListTransformJobs(input)
}

func (a *SageMakerActivities) ListTrialComponents(input *sagemaker.ListTrialComponentsInput) (*sagemaker.ListTrialComponentsOutput, error) {
    return a.client.ListTrialComponents(input)
}

func (a *SageMakerActivities) ListTrials(input *sagemaker.ListTrialsInput) (*sagemaker.ListTrialsOutput, error) {
    return a.client.ListTrials(input)
}

func (a *SageMakerActivities) ListUserProfiles(input *sagemaker.ListUserProfilesInput) (*sagemaker.ListUserProfilesOutput, error) {
    return a.client.ListUserProfiles(input)
}

func (a *SageMakerActivities) ListWorkforces(input *sagemaker.ListWorkforcesInput) (*sagemaker.ListWorkforcesOutput, error) {
    return a.client.ListWorkforces(input)
}

func (a *SageMakerActivities) ListWorkteams(input *sagemaker.ListWorkteamsInput) (*sagemaker.ListWorkteamsOutput, error) {
    return a.client.ListWorkteams(input)
}

func (a *SageMakerActivities) RenderUiTemplate(input *sagemaker.RenderUiTemplateInput) (*sagemaker.RenderUiTemplateOutput, error) {
    return a.client.RenderUiTemplate(input)
}

func (a *SageMakerActivities) Search(input *sagemaker.SearchInput) (*sagemaker.SearchOutput, error) {
    return a.client.Search(input)
}

func (a *SageMakerActivities) StartMonitoringSchedule(input *sagemaker.StartMonitoringScheduleInput) (*sagemaker.StartMonitoringScheduleOutput, error) {
    return a.client.StartMonitoringSchedule(input)
}

func (a *SageMakerActivities) StartNotebookInstance(input *sagemaker.StartNotebookInstanceInput) (*sagemaker.StartNotebookInstanceOutput, error) {
    return a.client.StartNotebookInstance(input)
}

func (a *SageMakerActivities) StopAutoMLJob(input *sagemaker.StopAutoMLJobInput) (*sagemaker.StopAutoMLJobOutput, error) {
    return a.client.StopAutoMLJob(input)
}

func (a *SageMakerActivities) StopCompilationJob(input *sagemaker.StopCompilationJobInput) (*sagemaker.StopCompilationJobOutput, error) {
    return a.client.StopCompilationJob(input)
}

func (a *SageMakerActivities) StopHyperParameterTuningJob(input *sagemaker.StopHyperParameterTuningJobInput) (*sagemaker.StopHyperParameterTuningJobOutput, error) {
    return a.client.StopHyperParameterTuningJob(input)
}

func (a *SageMakerActivities) StopLabelingJob(input *sagemaker.StopLabelingJobInput) (*sagemaker.StopLabelingJobOutput, error) {
    return a.client.StopLabelingJob(input)
}

func (a *SageMakerActivities) StopMonitoringSchedule(input *sagemaker.StopMonitoringScheduleInput) (*sagemaker.StopMonitoringScheduleOutput, error) {
    return a.client.StopMonitoringSchedule(input)
}

func (a *SageMakerActivities) StopNotebookInstance(input *sagemaker.StopNotebookInstanceInput) (*sagemaker.StopNotebookInstanceOutput, error) {
    return a.client.StopNotebookInstance(input)
}

func (a *SageMakerActivities) StopProcessingJob(input *sagemaker.StopProcessingJobInput) (*sagemaker.StopProcessingJobOutput, error) {
    return a.client.StopProcessingJob(input)
}

func (a *SageMakerActivities) StopTrainingJob(input *sagemaker.StopTrainingJobInput) (*sagemaker.StopTrainingJobOutput, error) {
    return a.client.StopTrainingJob(input)
}

func (a *SageMakerActivities) StopTransformJob(input *sagemaker.StopTransformJobInput) (*sagemaker.StopTransformJobOutput, error) {
    return a.client.StopTransformJob(input)
}

func (a *SageMakerActivities) UpdateCodeRepository(input *sagemaker.UpdateCodeRepositoryInput) (*sagemaker.UpdateCodeRepositoryOutput, error) {
    return a.client.UpdateCodeRepository(input)
}

func (a *SageMakerActivities) UpdateDomain(input *sagemaker.UpdateDomainInput) (*sagemaker.UpdateDomainOutput, error) {
    return a.client.UpdateDomain(input)
}

func (a *SageMakerActivities) UpdateEndpoint(input *sagemaker.UpdateEndpointInput) (*sagemaker.UpdateEndpointOutput, error) {
    return a.client.UpdateEndpoint(input)
}

func (a *SageMakerActivities) UpdateEndpointWeightsAndCapacities(input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) (*sagemaker.UpdateEndpointWeightsAndCapacitiesOutput, error) {
    return a.client.UpdateEndpointWeightsAndCapacities(input)
}

func (a *SageMakerActivities) UpdateExperiment(input *sagemaker.UpdateExperimentInput) (*sagemaker.UpdateExperimentOutput, error) {
    return a.client.UpdateExperiment(input)
}

func (a *SageMakerActivities) UpdateMonitoringSchedule(input *sagemaker.UpdateMonitoringScheduleInput) (*sagemaker.UpdateMonitoringScheduleOutput, error) {
    return a.client.UpdateMonitoringSchedule(input)
}

func (a *SageMakerActivities) UpdateNotebookInstance(input *sagemaker.UpdateNotebookInstanceInput) (*sagemaker.UpdateNotebookInstanceOutput, error) {
    return a.client.UpdateNotebookInstance(input)
}

func (a *SageMakerActivities) UpdateNotebookInstanceLifecycleConfig(input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) (*sagemaker.UpdateNotebookInstanceLifecycleConfigOutput, error) {
    return a.client.UpdateNotebookInstanceLifecycleConfig(input)
}

func (a *SageMakerActivities) UpdateTrial(input *sagemaker.UpdateTrialInput) (*sagemaker.UpdateTrialOutput, error) {
    return a.client.UpdateTrial(input)
}

func (a *SageMakerActivities) UpdateTrialComponent(input *sagemaker.UpdateTrialComponentInput) (*sagemaker.UpdateTrialComponentOutput, error) {
    return a.client.UpdateTrialComponent(input)
}

func (a *SageMakerActivities) UpdateUserProfile(input *sagemaker.UpdateUserProfileInput) (*sagemaker.UpdateUserProfileOutput, error) {
    return a.client.UpdateUserProfile(input)
}

func (a *SageMakerActivities) UpdateWorkforce(input *sagemaker.UpdateWorkforceInput) (*sagemaker.UpdateWorkforceOutput, error) {
    return a.client.UpdateWorkforce(input)
}

func (a *SageMakerActivities) UpdateWorkteam(input *sagemaker.UpdateWorkteamInput) (*sagemaker.UpdateWorkteamOutput, error) {
    return a.client.UpdateWorkteam(input)
}

func (a *SageMakerActivities) WaitUntilEndpointDeleted(input *sagemaker.DescribeEndpointInput) error {
    return a.client.WaitUntilEndpointDeleted(input)
}

func (a *SageMakerActivities) WaitUntilEndpointInService(input *sagemaker.DescribeEndpointInput) error {
    return a.client.WaitUntilEndpointInService(input)
}

func (a *SageMakerActivities) WaitUntilNotebookInstanceDeleted(input *sagemaker.DescribeNotebookInstanceInput) error {
    return a.client.WaitUntilNotebookInstanceDeleted(input)
}

func (a *SageMakerActivities) WaitUntilNotebookInstanceInService(input *sagemaker.DescribeNotebookInstanceInput) error {
    return a.client.WaitUntilNotebookInstanceInService(input)
}

func (a *SageMakerActivities) WaitUntilNotebookInstanceStopped(input *sagemaker.DescribeNotebookInstanceInput) error {
    return a.client.WaitUntilNotebookInstanceStopped(input)
}

func (a *SageMakerActivities) WaitUntilProcessingJobCompletedOrStopped(input *sagemaker.DescribeProcessingJobInput) error {
    return a.client.WaitUntilProcessingJobCompletedOrStopped(input)
}

func (a *SageMakerActivities) WaitUntilTrainingJobCompletedOrStopped(input *sagemaker.DescribeTrainingJobInput) error {
    return a.client.WaitUntilTrainingJobCompletedOrStopped(input)
}

func (a *SageMakerActivities) WaitUntilTransformJobCompletedOrStopped(input *sagemaker.DescribeTransformJobInput) error {
    return a.client.WaitUntilTransformJobCompletedOrStopped(input)
}
