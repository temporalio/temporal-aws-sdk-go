package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"go.temporal.io/sdk/workflow"
)

type SageMakerClient interface {
    AddTags(ctx workflow.Context, input *sagemaker.AddTagsInput) (*sagemaker.AddTagsOutput, error)
    AddTagsAsync(ctx workflow.Context, input *sagemaker.AddTagsInput) *SagemakerAddTagsResult

    AssociateTrialComponent(ctx workflow.Context, input *sagemaker.AssociateTrialComponentInput) (*sagemaker.AssociateTrialComponentOutput, error)
    AssociateTrialComponentAsync(ctx workflow.Context, input *sagemaker.AssociateTrialComponentInput) *SagemakerAssociateTrialComponentResult

    CreateAlgorithm(ctx workflow.Context, input *sagemaker.CreateAlgorithmInput) (*sagemaker.CreateAlgorithmOutput, error)
    CreateAlgorithmAsync(ctx workflow.Context, input *sagemaker.CreateAlgorithmInput) *SagemakerCreateAlgorithmResult

    CreateApp(ctx workflow.Context, input *sagemaker.CreateAppInput) (*sagemaker.CreateAppOutput, error)
    CreateAppAsync(ctx workflow.Context, input *sagemaker.CreateAppInput) *SagemakerCreateAppResult

    CreateAutoMLJob(ctx workflow.Context, input *sagemaker.CreateAutoMLJobInput) (*sagemaker.CreateAutoMLJobOutput, error)
    CreateAutoMLJobAsync(ctx workflow.Context, input *sagemaker.CreateAutoMLJobInput) *SagemakerCreateAutoMLJobResult

    CreateCodeRepository(ctx workflow.Context, input *sagemaker.CreateCodeRepositoryInput) (*sagemaker.CreateCodeRepositoryOutput, error)
    CreateCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.CreateCodeRepositoryInput) *SagemakerCreateCodeRepositoryResult

    CreateCompilationJob(ctx workflow.Context, input *sagemaker.CreateCompilationJobInput) (*sagemaker.CreateCompilationJobOutput, error)
    CreateCompilationJobAsync(ctx workflow.Context, input *sagemaker.CreateCompilationJobInput) *SagemakerCreateCompilationJobResult

    CreateDomain(ctx workflow.Context, input *sagemaker.CreateDomainInput) (*sagemaker.CreateDomainOutput, error)
    CreateDomainAsync(ctx workflow.Context, input *sagemaker.CreateDomainInput) *SagemakerCreateDomainResult

    CreateEndpoint(ctx workflow.Context, input *sagemaker.CreateEndpointInput) (*sagemaker.CreateEndpointOutput, error)
    CreateEndpointAsync(ctx workflow.Context, input *sagemaker.CreateEndpointInput) *SagemakerCreateEndpointResult

    CreateEndpointConfig(ctx workflow.Context, input *sagemaker.CreateEndpointConfigInput) (*sagemaker.CreateEndpointConfigOutput, error)
    CreateEndpointConfigAsync(ctx workflow.Context, input *sagemaker.CreateEndpointConfigInput) *SagemakerCreateEndpointConfigResult

    CreateExperiment(ctx workflow.Context, input *sagemaker.CreateExperimentInput) (*sagemaker.CreateExperimentOutput, error)
    CreateExperimentAsync(ctx workflow.Context, input *sagemaker.CreateExperimentInput) *SagemakerCreateExperimentResult

    CreateFlowDefinition(ctx workflow.Context, input *sagemaker.CreateFlowDefinitionInput) (*sagemaker.CreateFlowDefinitionOutput, error)
    CreateFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.CreateFlowDefinitionInput) *SagemakerCreateFlowDefinitionResult

    CreateHumanTaskUi(ctx workflow.Context, input *sagemaker.CreateHumanTaskUiInput) (*sagemaker.CreateHumanTaskUiOutput, error)
    CreateHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.CreateHumanTaskUiInput) *SagemakerCreateHumanTaskUiResult

    CreateHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.CreateHyperParameterTuningJobInput) (*sagemaker.CreateHyperParameterTuningJobOutput, error)
    CreateHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.CreateHyperParameterTuningJobInput) *SagemakerCreateHyperParameterTuningJobResult

    CreateLabelingJob(ctx workflow.Context, input *sagemaker.CreateLabelingJobInput) (*sagemaker.CreateLabelingJobOutput, error)
    CreateLabelingJobAsync(ctx workflow.Context, input *sagemaker.CreateLabelingJobInput) *SagemakerCreateLabelingJobResult

    CreateModel(ctx workflow.Context, input *sagemaker.CreateModelInput) (*sagemaker.CreateModelOutput, error)
    CreateModelAsync(ctx workflow.Context, input *sagemaker.CreateModelInput) *SagemakerCreateModelResult

    CreateModelPackage(ctx workflow.Context, input *sagemaker.CreateModelPackageInput) (*sagemaker.CreateModelPackageOutput, error)
    CreateModelPackageAsync(ctx workflow.Context, input *sagemaker.CreateModelPackageInput) *SagemakerCreateModelPackageResult

    CreateMonitoringSchedule(ctx workflow.Context, input *sagemaker.CreateMonitoringScheduleInput) (*sagemaker.CreateMonitoringScheduleOutput, error)
    CreateMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.CreateMonitoringScheduleInput) *SagemakerCreateMonitoringScheduleResult

    CreateNotebookInstance(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceInput) (*sagemaker.CreateNotebookInstanceOutput, error)
    CreateNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceInput) *SagemakerCreateNotebookInstanceResult

    CreateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) (*sagemaker.CreateNotebookInstanceLifecycleConfigOutput, error)
    CreateNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) *SagemakerCreateNotebookInstanceLifecycleConfigResult

    CreatePresignedDomainUrl(ctx workflow.Context, input *sagemaker.CreatePresignedDomainUrlInput) (*sagemaker.CreatePresignedDomainUrlOutput, error)
    CreatePresignedDomainUrlAsync(ctx workflow.Context, input *sagemaker.CreatePresignedDomainUrlInput) *SagemakerCreatePresignedDomainUrlResult

    CreatePresignedNotebookInstanceUrl(ctx workflow.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) (*sagemaker.CreatePresignedNotebookInstanceUrlOutput, error)
    CreatePresignedNotebookInstanceUrlAsync(ctx workflow.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) *SagemakerCreatePresignedNotebookInstanceUrlResult

    CreateProcessingJob(ctx workflow.Context, input *sagemaker.CreateProcessingJobInput) (*sagemaker.CreateProcessingJobOutput, error)
    CreateProcessingJobAsync(ctx workflow.Context, input *sagemaker.CreateProcessingJobInput) *SagemakerCreateProcessingJobResult

    CreateTrainingJob(ctx workflow.Context, input *sagemaker.CreateTrainingJobInput) (*sagemaker.CreateTrainingJobOutput, error)
    CreateTrainingJobAsync(ctx workflow.Context, input *sagemaker.CreateTrainingJobInput) *SagemakerCreateTrainingJobResult

    CreateTransformJob(ctx workflow.Context, input *sagemaker.CreateTransformJobInput) (*sagemaker.CreateTransformJobOutput, error)
    CreateTransformJobAsync(ctx workflow.Context, input *sagemaker.CreateTransformJobInput) *SagemakerCreateTransformJobResult

    CreateTrial(ctx workflow.Context, input *sagemaker.CreateTrialInput) (*sagemaker.CreateTrialOutput, error)
    CreateTrialAsync(ctx workflow.Context, input *sagemaker.CreateTrialInput) *SagemakerCreateTrialResult

    CreateTrialComponent(ctx workflow.Context, input *sagemaker.CreateTrialComponentInput) (*sagemaker.CreateTrialComponentOutput, error)
    CreateTrialComponentAsync(ctx workflow.Context, input *sagemaker.CreateTrialComponentInput) *SagemakerCreateTrialComponentResult

    CreateUserProfile(ctx workflow.Context, input *sagemaker.CreateUserProfileInput) (*sagemaker.CreateUserProfileOutput, error)
    CreateUserProfileAsync(ctx workflow.Context, input *sagemaker.CreateUserProfileInput) *SagemakerCreateUserProfileResult

    CreateWorkforce(ctx workflow.Context, input *sagemaker.CreateWorkforceInput) (*sagemaker.CreateWorkforceOutput, error)
    CreateWorkforceAsync(ctx workflow.Context, input *sagemaker.CreateWorkforceInput) *SagemakerCreateWorkforceResult

    CreateWorkteam(ctx workflow.Context, input *sagemaker.CreateWorkteamInput) (*sagemaker.CreateWorkteamOutput, error)
    CreateWorkteamAsync(ctx workflow.Context, input *sagemaker.CreateWorkteamInput) *SagemakerCreateWorkteamResult

    DeleteAlgorithm(ctx workflow.Context, input *sagemaker.DeleteAlgorithmInput) (*sagemaker.DeleteAlgorithmOutput, error)
    DeleteAlgorithmAsync(ctx workflow.Context, input *sagemaker.DeleteAlgorithmInput) *SagemakerDeleteAlgorithmResult

    DeleteApp(ctx workflow.Context, input *sagemaker.DeleteAppInput) (*sagemaker.DeleteAppOutput, error)
    DeleteAppAsync(ctx workflow.Context, input *sagemaker.DeleteAppInput) *SagemakerDeleteAppResult

    DeleteCodeRepository(ctx workflow.Context, input *sagemaker.DeleteCodeRepositoryInput) (*sagemaker.DeleteCodeRepositoryOutput, error)
    DeleteCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.DeleteCodeRepositoryInput) *SagemakerDeleteCodeRepositoryResult

    DeleteDomain(ctx workflow.Context, input *sagemaker.DeleteDomainInput) (*sagemaker.DeleteDomainOutput, error)
    DeleteDomainAsync(ctx workflow.Context, input *sagemaker.DeleteDomainInput) *SagemakerDeleteDomainResult

    DeleteEndpoint(ctx workflow.Context, input *sagemaker.DeleteEndpointInput) (*sagemaker.DeleteEndpointOutput, error)
    DeleteEndpointAsync(ctx workflow.Context, input *sagemaker.DeleteEndpointInput) *SagemakerDeleteEndpointResult

    DeleteEndpointConfig(ctx workflow.Context, input *sagemaker.DeleteEndpointConfigInput) (*sagemaker.DeleteEndpointConfigOutput, error)
    DeleteEndpointConfigAsync(ctx workflow.Context, input *sagemaker.DeleteEndpointConfigInput) *SagemakerDeleteEndpointConfigResult

    DeleteExperiment(ctx workflow.Context, input *sagemaker.DeleteExperimentInput) (*sagemaker.DeleteExperimentOutput, error)
    DeleteExperimentAsync(ctx workflow.Context, input *sagemaker.DeleteExperimentInput) *SagemakerDeleteExperimentResult

    DeleteFlowDefinition(ctx workflow.Context, input *sagemaker.DeleteFlowDefinitionInput) (*sagemaker.DeleteFlowDefinitionOutput, error)
    DeleteFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.DeleteFlowDefinitionInput) *SagemakerDeleteFlowDefinitionResult

    DeleteHumanTaskUi(ctx workflow.Context, input *sagemaker.DeleteHumanTaskUiInput) (*sagemaker.DeleteHumanTaskUiOutput, error)
    DeleteHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.DeleteHumanTaskUiInput) *SagemakerDeleteHumanTaskUiResult

    DeleteModel(ctx workflow.Context, input *sagemaker.DeleteModelInput) (*sagemaker.DeleteModelOutput, error)
    DeleteModelAsync(ctx workflow.Context, input *sagemaker.DeleteModelInput) *SagemakerDeleteModelResult

    DeleteModelPackage(ctx workflow.Context, input *sagemaker.DeleteModelPackageInput) (*sagemaker.DeleteModelPackageOutput, error)
    DeleteModelPackageAsync(ctx workflow.Context, input *sagemaker.DeleteModelPackageInput) *SagemakerDeleteModelPackageResult

    DeleteMonitoringSchedule(ctx workflow.Context, input *sagemaker.DeleteMonitoringScheduleInput) (*sagemaker.DeleteMonitoringScheduleOutput, error)
    DeleteMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.DeleteMonitoringScheduleInput) *SagemakerDeleteMonitoringScheduleResult

    DeleteNotebookInstance(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceInput) (*sagemaker.DeleteNotebookInstanceOutput, error)
    DeleteNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceInput) *SagemakerDeleteNotebookInstanceResult

    DeleteNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) (*sagemaker.DeleteNotebookInstanceLifecycleConfigOutput, error)
    DeleteNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) *SagemakerDeleteNotebookInstanceLifecycleConfigResult

    DeleteTags(ctx workflow.Context, input *sagemaker.DeleteTagsInput) (*sagemaker.DeleteTagsOutput, error)
    DeleteTagsAsync(ctx workflow.Context, input *sagemaker.DeleteTagsInput) *SagemakerDeleteTagsResult

    DeleteTrial(ctx workflow.Context, input *sagemaker.DeleteTrialInput) (*sagemaker.DeleteTrialOutput, error)
    DeleteTrialAsync(ctx workflow.Context, input *sagemaker.DeleteTrialInput) *SagemakerDeleteTrialResult

    DeleteTrialComponent(ctx workflow.Context, input *sagemaker.DeleteTrialComponentInput) (*sagemaker.DeleteTrialComponentOutput, error)
    DeleteTrialComponentAsync(ctx workflow.Context, input *sagemaker.DeleteTrialComponentInput) *SagemakerDeleteTrialComponentResult

    DeleteUserProfile(ctx workflow.Context, input *sagemaker.DeleteUserProfileInput) (*sagemaker.DeleteUserProfileOutput, error)
    DeleteUserProfileAsync(ctx workflow.Context, input *sagemaker.DeleteUserProfileInput) *SagemakerDeleteUserProfileResult

    DeleteWorkforce(ctx workflow.Context, input *sagemaker.DeleteWorkforceInput) (*sagemaker.DeleteWorkforceOutput, error)
    DeleteWorkforceAsync(ctx workflow.Context, input *sagemaker.DeleteWorkforceInput) *SagemakerDeleteWorkforceResult

    DeleteWorkteam(ctx workflow.Context, input *sagemaker.DeleteWorkteamInput) (*sagemaker.DeleteWorkteamOutput, error)
    DeleteWorkteamAsync(ctx workflow.Context, input *sagemaker.DeleteWorkteamInput) *SagemakerDeleteWorkteamResult

    DescribeAlgorithm(ctx workflow.Context, input *sagemaker.DescribeAlgorithmInput) (*sagemaker.DescribeAlgorithmOutput, error)
    DescribeAlgorithmAsync(ctx workflow.Context, input *sagemaker.DescribeAlgorithmInput) *SagemakerDescribeAlgorithmResult

    DescribeApp(ctx workflow.Context, input *sagemaker.DescribeAppInput) (*sagemaker.DescribeAppOutput, error)
    DescribeAppAsync(ctx workflow.Context, input *sagemaker.DescribeAppInput) *SagemakerDescribeAppResult

    DescribeAutoMLJob(ctx workflow.Context, input *sagemaker.DescribeAutoMLJobInput) (*sagemaker.DescribeAutoMLJobOutput, error)
    DescribeAutoMLJobAsync(ctx workflow.Context, input *sagemaker.DescribeAutoMLJobInput) *SagemakerDescribeAutoMLJobResult

    DescribeCodeRepository(ctx workflow.Context, input *sagemaker.DescribeCodeRepositoryInput) (*sagemaker.DescribeCodeRepositoryOutput, error)
    DescribeCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.DescribeCodeRepositoryInput) *SagemakerDescribeCodeRepositoryResult

    DescribeCompilationJob(ctx workflow.Context, input *sagemaker.DescribeCompilationJobInput) (*sagemaker.DescribeCompilationJobOutput, error)
    DescribeCompilationJobAsync(ctx workflow.Context, input *sagemaker.DescribeCompilationJobInput) *SagemakerDescribeCompilationJobResult

    DescribeDomain(ctx workflow.Context, input *sagemaker.DescribeDomainInput) (*sagemaker.DescribeDomainOutput, error)
    DescribeDomainAsync(ctx workflow.Context, input *sagemaker.DescribeDomainInput) *SagemakerDescribeDomainResult

    DescribeEndpoint(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) (*sagemaker.DescribeEndpointOutput, error)
    DescribeEndpointAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) *SagemakerDescribeEndpointResult

    DescribeEndpointConfig(ctx workflow.Context, input *sagemaker.DescribeEndpointConfigInput) (*sagemaker.DescribeEndpointConfigOutput, error)
    DescribeEndpointConfigAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointConfigInput) *SagemakerDescribeEndpointConfigResult

    DescribeExperiment(ctx workflow.Context, input *sagemaker.DescribeExperimentInput) (*sagemaker.DescribeExperimentOutput, error)
    DescribeExperimentAsync(ctx workflow.Context, input *sagemaker.DescribeExperimentInput) *SagemakerDescribeExperimentResult

    DescribeFlowDefinition(ctx workflow.Context, input *sagemaker.DescribeFlowDefinitionInput) (*sagemaker.DescribeFlowDefinitionOutput, error)
    DescribeFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.DescribeFlowDefinitionInput) *SagemakerDescribeFlowDefinitionResult

    DescribeHumanTaskUi(ctx workflow.Context, input *sagemaker.DescribeHumanTaskUiInput) (*sagemaker.DescribeHumanTaskUiOutput, error)
    DescribeHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.DescribeHumanTaskUiInput) *SagemakerDescribeHumanTaskUiResult

    DescribeHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) (*sagemaker.DescribeHyperParameterTuningJobOutput, error)
    DescribeHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) *SagemakerDescribeHyperParameterTuningJobResult

    DescribeLabelingJob(ctx workflow.Context, input *sagemaker.DescribeLabelingJobInput) (*sagemaker.DescribeLabelingJobOutput, error)
    DescribeLabelingJobAsync(ctx workflow.Context, input *sagemaker.DescribeLabelingJobInput) *SagemakerDescribeLabelingJobResult

    DescribeModel(ctx workflow.Context, input *sagemaker.DescribeModelInput) (*sagemaker.DescribeModelOutput, error)
    DescribeModelAsync(ctx workflow.Context, input *sagemaker.DescribeModelInput) *SagemakerDescribeModelResult

    DescribeModelPackage(ctx workflow.Context, input *sagemaker.DescribeModelPackageInput) (*sagemaker.DescribeModelPackageOutput, error)
    DescribeModelPackageAsync(ctx workflow.Context, input *sagemaker.DescribeModelPackageInput) *SagemakerDescribeModelPackageResult

    DescribeMonitoringSchedule(ctx workflow.Context, input *sagemaker.DescribeMonitoringScheduleInput) (*sagemaker.DescribeMonitoringScheduleOutput, error)
    DescribeMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.DescribeMonitoringScheduleInput) *SagemakerDescribeMonitoringScheduleResult

    DescribeNotebookInstance(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) (*sagemaker.DescribeNotebookInstanceOutput, error)
    DescribeNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *SagemakerDescribeNotebookInstanceResult

    DescribeNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error)
    DescribeNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) *SagemakerDescribeNotebookInstanceLifecycleConfigResult

    DescribeProcessingJob(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) (*sagemaker.DescribeProcessingJobOutput, error)
    DescribeProcessingJobAsync(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) *SagemakerDescribeProcessingJobResult

    DescribeSubscribedWorkteam(ctx workflow.Context, input *sagemaker.DescribeSubscribedWorkteamInput) (*sagemaker.DescribeSubscribedWorkteamOutput, error)
    DescribeSubscribedWorkteamAsync(ctx workflow.Context, input *sagemaker.DescribeSubscribedWorkteamInput) *SagemakerDescribeSubscribedWorkteamResult

    DescribeTrainingJob(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) (*sagemaker.DescribeTrainingJobOutput, error)
    DescribeTrainingJobAsync(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) *SagemakerDescribeTrainingJobResult

    DescribeTransformJob(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) (*sagemaker.DescribeTransformJobOutput, error)
    DescribeTransformJobAsync(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) *SagemakerDescribeTransformJobResult

    DescribeTrial(ctx workflow.Context, input *sagemaker.DescribeTrialInput) (*sagemaker.DescribeTrialOutput, error)
    DescribeTrialAsync(ctx workflow.Context, input *sagemaker.DescribeTrialInput) *SagemakerDescribeTrialResult

    DescribeTrialComponent(ctx workflow.Context, input *sagemaker.DescribeTrialComponentInput) (*sagemaker.DescribeTrialComponentOutput, error)
    DescribeTrialComponentAsync(ctx workflow.Context, input *sagemaker.DescribeTrialComponentInput) *SagemakerDescribeTrialComponentResult

    DescribeUserProfile(ctx workflow.Context, input *sagemaker.DescribeUserProfileInput) (*sagemaker.DescribeUserProfileOutput, error)
    DescribeUserProfileAsync(ctx workflow.Context, input *sagemaker.DescribeUserProfileInput) *SagemakerDescribeUserProfileResult

    DescribeWorkforce(ctx workflow.Context, input *sagemaker.DescribeWorkforceInput) (*sagemaker.DescribeWorkforceOutput, error)
    DescribeWorkforceAsync(ctx workflow.Context, input *sagemaker.DescribeWorkforceInput) *SagemakerDescribeWorkforceResult

    DescribeWorkteam(ctx workflow.Context, input *sagemaker.DescribeWorkteamInput) (*sagemaker.DescribeWorkteamOutput, error)
    DescribeWorkteamAsync(ctx workflow.Context, input *sagemaker.DescribeWorkteamInput) *SagemakerDescribeWorkteamResult

    DisassociateTrialComponent(ctx workflow.Context, input *sagemaker.DisassociateTrialComponentInput) (*sagemaker.DisassociateTrialComponentOutput, error)
    DisassociateTrialComponentAsync(ctx workflow.Context, input *sagemaker.DisassociateTrialComponentInput) *SagemakerDisassociateTrialComponentResult

    GetSearchSuggestions(ctx workflow.Context, input *sagemaker.GetSearchSuggestionsInput) (*sagemaker.GetSearchSuggestionsOutput, error)
    GetSearchSuggestionsAsync(ctx workflow.Context, input *sagemaker.GetSearchSuggestionsInput) *SagemakerGetSearchSuggestionsResult

    ListAlgorithms(ctx workflow.Context, input *sagemaker.ListAlgorithmsInput) (*sagemaker.ListAlgorithmsOutput, error)
    ListAlgorithmsAsync(ctx workflow.Context, input *sagemaker.ListAlgorithmsInput) *SagemakerListAlgorithmsResult

    ListApps(ctx workflow.Context, input *sagemaker.ListAppsInput) (*sagemaker.ListAppsOutput, error)
    ListAppsAsync(ctx workflow.Context, input *sagemaker.ListAppsInput) *SagemakerListAppsResult

    ListAutoMLJobs(ctx workflow.Context, input *sagemaker.ListAutoMLJobsInput) (*sagemaker.ListAutoMLJobsOutput, error)
    ListAutoMLJobsAsync(ctx workflow.Context, input *sagemaker.ListAutoMLJobsInput) *SagemakerListAutoMLJobsResult

    ListCandidatesForAutoMLJob(ctx workflow.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) (*sagemaker.ListCandidatesForAutoMLJobOutput, error)
    ListCandidatesForAutoMLJobAsync(ctx workflow.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) *SagemakerListCandidatesForAutoMLJobResult

    ListCodeRepositories(ctx workflow.Context, input *sagemaker.ListCodeRepositoriesInput) (*sagemaker.ListCodeRepositoriesOutput, error)
    ListCodeRepositoriesAsync(ctx workflow.Context, input *sagemaker.ListCodeRepositoriesInput) *SagemakerListCodeRepositoriesResult

    ListCompilationJobs(ctx workflow.Context, input *sagemaker.ListCompilationJobsInput) (*sagemaker.ListCompilationJobsOutput, error)
    ListCompilationJobsAsync(ctx workflow.Context, input *sagemaker.ListCompilationJobsInput) *SagemakerListCompilationJobsResult

    ListDomains(ctx workflow.Context, input *sagemaker.ListDomainsInput) (*sagemaker.ListDomainsOutput, error)
    ListDomainsAsync(ctx workflow.Context, input *sagemaker.ListDomainsInput) *SagemakerListDomainsResult

    ListEndpointConfigs(ctx workflow.Context, input *sagemaker.ListEndpointConfigsInput) (*sagemaker.ListEndpointConfigsOutput, error)
    ListEndpointConfigsAsync(ctx workflow.Context, input *sagemaker.ListEndpointConfigsInput) *SagemakerListEndpointConfigsResult

    ListEndpoints(ctx workflow.Context, input *sagemaker.ListEndpointsInput) (*sagemaker.ListEndpointsOutput, error)
    ListEndpointsAsync(ctx workflow.Context, input *sagemaker.ListEndpointsInput) *SagemakerListEndpointsResult

    ListExperiments(ctx workflow.Context, input *sagemaker.ListExperimentsInput) (*sagemaker.ListExperimentsOutput, error)
    ListExperimentsAsync(ctx workflow.Context, input *sagemaker.ListExperimentsInput) *SagemakerListExperimentsResult

    ListFlowDefinitions(ctx workflow.Context, input *sagemaker.ListFlowDefinitionsInput) (*sagemaker.ListFlowDefinitionsOutput, error)
    ListFlowDefinitionsAsync(ctx workflow.Context, input *sagemaker.ListFlowDefinitionsInput) *SagemakerListFlowDefinitionsResult

    ListHumanTaskUis(ctx workflow.Context, input *sagemaker.ListHumanTaskUisInput) (*sagemaker.ListHumanTaskUisOutput, error)
    ListHumanTaskUisAsync(ctx workflow.Context, input *sagemaker.ListHumanTaskUisInput) *SagemakerListHumanTaskUisResult

    ListHyperParameterTuningJobs(ctx workflow.Context, input *sagemaker.ListHyperParameterTuningJobsInput) (*sagemaker.ListHyperParameterTuningJobsOutput, error)
    ListHyperParameterTuningJobsAsync(ctx workflow.Context, input *sagemaker.ListHyperParameterTuningJobsInput) *SagemakerListHyperParameterTuningJobsResult

    ListLabelingJobs(ctx workflow.Context, input *sagemaker.ListLabelingJobsInput) (*sagemaker.ListLabelingJobsOutput, error)
    ListLabelingJobsAsync(ctx workflow.Context, input *sagemaker.ListLabelingJobsInput) *SagemakerListLabelingJobsResult

    ListLabelingJobsForWorkteam(ctx workflow.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) (*sagemaker.ListLabelingJobsForWorkteamOutput, error)
    ListLabelingJobsForWorkteamAsync(ctx workflow.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) *SagemakerListLabelingJobsForWorkteamResult

    ListModelPackages(ctx workflow.Context, input *sagemaker.ListModelPackagesInput) (*sagemaker.ListModelPackagesOutput, error)
    ListModelPackagesAsync(ctx workflow.Context, input *sagemaker.ListModelPackagesInput) *SagemakerListModelPackagesResult

    ListModels(ctx workflow.Context, input *sagemaker.ListModelsInput) (*sagemaker.ListModelsOutput, error)
    ListModelsAsync(ctx workflow.Context, input *sagemaker.ListModelsInput) *SagemakerListModelsResult

    ListMonitoringExecutions(ctx workflow.Context, input *sagemaker.ListMonitoringExecutionsInput) (*sagemaker.ListMonitoringExecutionsOutput, error)
    ListMonitoringExecutionsAsync(ctx workflow.Context, input *sagemaker.ListMonitoringExecutionsInput) *SagemakerListMonitoringExecutionsResult

    ListMonitoringSchedules(ctx workflow.Context, input *sagemaker.ListMonitoringSchedulesInput) (*sagemaker.ListMonitoringSchedulesOutput, error)
    ListMonitoringSchedulesAsync(ctx workflow.Context, input *sagemaker.ListMonitoringSchedulesInput) *SagemakerListMonitoringSchedulesResult

    ListNotebookInstanceLifecycleConfigs(ctx workflow.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error)
    ListNotebookInstanceLifecycleConfigsAsync(ctx workflow.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) *SagemakerListNotebookInstanceLifecycleConfigsResult

    ListNotebookInstances(ctx workflow.Context, input *sagemaker.ListNotebookInstancesInput) (*sagemaker.ListNotebookInstancesOutput, error)
    ListNotebookInstancesAsync(ctx workflow.Context, input *sagemaker.ListNotebookInstancesInput) *SagemakerListNotebookInstancesResult

    ListProcessingJobs(ctx workflow.Context, input *sagemaker.ListProcessingJobsInput) (*sagemaker.ListProcessingJobsOutput, error)
    ListProcessingJobsAsync(ctx workflow.Context, input *sagemaker.ListProcessingJobsInput) *SagemakerListProcessingJobsResult

    ListSubscribedWorkteams(ctx workflow.Context, input *sagemaker.ListSubscribedWorkteamsInput) (*sagemaker.ListSubscribedWorkteamsOutput, error)
    ListSubscribedWorkteamsAsync(ctx workflow.Context, input *sagemaker.ListSubscribedWorkteamsInput) *SagemakerListSubscribedWorkteamsResult

    ListTags(ctx workflow.Context, input *sagemaker.ListTagsInput) (*sagemaker.ListTagsOutput, error)
    ListTagsAsync(ctx workflow.Context, input *sagemaker.ListTagsInput) *SagemakerListTagsResult

    ListTrainingJobs(ctx workflow.Context, input *sagemaker.ListTrainingJobsInput) (*sagemaker.ListTrainingJobsOutput, error)
    ListTrainingJobsAsync(ctx workflow.Context, input *sagemaker.ListTrainingJobsInput) *SagemakerListTrainingJobsResult

    ListTrainingJobsForHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error)
    ListTrainingJobsForHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) *SagemakerListTrainingJobsForHyperParameterTuningJobResult

    ListTransformJobs(ctx workflow.Context, input *sagemaker.ListTransformJobsInput) (*sagemaker.ListTransformJobsOutput, error)
    ListTransformJobsAsync(ctx workflow.Context, input *sagemaker.ListTransformJobsInput) *SagemakerListTransformJobsResult

    ListTrialComponents(ctx workflow.Context, input *sagemaker.ListTrialComponentsInput) (*sagemaker.ListTrialComponentsOutput, error)
    ListTrialComponentsAsync(ctx workflow.Context, input *sagemaker.ListTrialComponentsInput) *SagemakerListTrialComponentsResult

    ListTrials(ctx workflow.Context, input *sagemaker.ListTrialsInput) (*sagemaker.ListTrialsOutput, error)
    ListTrialsAsync(ctx workflow.Context, input *sagemaker.ListTrialsInput) *SagemakerListTrialsResult

    ListUserProfiles(ctx workflow.Context, input *sagemaker.ListUserProfilesInput) (*sagemaker.ListUserProfilesOutput, error)
    ListUserProfilesAsync(ctx workflow.Context, input *sagemaker.ListUserProfilesInput) *SagemakerListUserProfilesResult

    ListWorkforces(ctx workflow.Context, input *sagemaker.ListWorkforcesInput) (*sagemaker.ListWorkforcesOutput, error)
    ListWorkforcesAsync(ctx workflow.Context, input *sagemaker.ListWorkforcesInput) *SagemakerListWorkforcesResult

    ListWorkteams(ctx workflow.Context, input *sagemaker.ListWorkteamsInput) (*sagemaker.ListWorkteamsOutput, error)
    ListWorkteamsAsync(ctx workflow.Context, input *sagemaker.ListWorkteamsInput) *SagemakerListWorkteamsResult

    RenderUiTemplate(ctx workflow.Context, input *sagemaker.RenderUiTemplateInput) (*sagemaker.RenderUiTemplateOutput, error)
    RenderUiTemplateAsync(ctx workflow.Context, input *sagemaker.RenderUiTemplateInput) *SagemakerRenderUiTemplateResult

    Search(ctx workflow.Context, input *sagemaker.SearchInput) (*sagemaker.SearchOutput, error)
    SearchAsync(ctx workflow.Context, input *sagemaker.SearchInput) *SagemakerSearchResult

    StartMonitoringSchedule(ctx workflow.Context, input *sagemaker.StartMonitoringScheduleInput) (*sagemaker.StartMonitoringScheduleOutput, error)
    StartMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.StartMonitoringScheduleInput) *SagemakerStartMonitoringScheduleResult

    StartNotebookInstance(ctx workflow.Context, input *sagemaker.StartNotebookInstanceInput) (*sagemaker.StartNotebookInstanceOutput, error)
    StartNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.StartNotebookInstanceInput) *SagemakerStartNotebookInstanceResult

    StopAutoMLJob(ctx workflow.Context, input *sagemaker.StopAutoMLJobInput) (*sagemaker.StopAutoMLJobOutput, error)
    StopAutoMLJobAsync(ctx workflow.Context, input *sagemaker.StopAutoMLJobInput) *SagemakerStopAutoMLJobResult

    StopCompilationJob(ctx workflow.Context, input *sagemaker.StopCompilationJobInput) (*sagemaker.StopCompilationJobOutput, error)
    StopCompilationJobAsync(ctx workflow.Context, input *sagemaker.StopCompilationJobInput) *SagemakerStopCompilationJobResult

    StopHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.StopHyperParameterTuningJobInput) (*sagemaker.StopHyperParameterTuningJobOutput, error)
    StopHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.StopHyperParameterTuningJobInput) *SagemakerStopHyperParameterTuningJobResult

    StopLabelingJob(ctx workflow.Context, input *sagemaker.StopLabelingJobInput) (*sagemaker.StopLabelingJobOutput, error)
    StopLabelingJobAsync(ctx workflow.Context, input *sagemaker.StopLabelingJobInput) *SagemakerStopLabelingJobResult

    StopMonitoringSchedule(ctx workflow.Context, input *sagemaker.StopMonitoringScheduleInput) (*sagemaker.StopMonitoringScheduleOutput, error)
    StopMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.StopMonitoringScheduleInput) *SagemakerStopMonitoringScheduleResult

    StopNotebookInstance(ctx workflow.Context, input *sagemaker.StopNotebookInstanceInput) (*sagemaker.StopNotebookInstanceOutput, error)
    StopNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.StopNotebookInstanceInput) *SagemakerStopNotebookInstanceResult

    StopProcessingJob(ctx workflow.Context, input *sagemaker.StopProcessingJobInput) (*sagemaker.StopProcessingJobOutput, error)
    StopProcessingJobAsync(ctx workflow.Context, input *sagemaker.StopProcessingJobInput) *SagemakerStopProcessingJobResult

    StopTrainingJob(ctx workflow.Context, input *sagemaker.StopTrainingJobInput) (*sagemaker.StopTrainingJobOutput, error)
    StopTrainingJobAsync(ctx workflow.Context, input *sagemaker.StopTrainingJobInput) *SagemakerStopTrainingJobResult

    StopTransformJob(ctx workflow.Context, input *sagemaker.StopTransformJobInput) (*sagemaker.StopTransformJobOutput, error)
    StopTransformJobAsync(ctx workflow.Context, input *sagemaker.StopTransformJobInput) *SagemakerStopTransformJobResult

    UpdateCodeRepository(ctx workflow.Context, input *sagemaker.UpdateCodeRepositoryInput) (*sagemaker.UpdateCodeRepositoryOutput, error)
    UpdateCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.UpdateCodeRepositoryInput) *SagemakerUpdateCodeRepositoryResult

    UpdateDomain(ctx workflow.Context, input *sagemaker.UpdateDomainInput) (*sagemaker.UpdateDomainOutput, error)
    UpdateDomainAsync(ctx workflow.Context, input *sagemaker.UpdateDomainInput) *SagemakerUpdateDomainResult

    UpdateEndpoint(ctx workflow.Context, input *sagemaker.UpdateEndpointInput) (*sagemaker.UpdateEndpointOutput, error)
    UpdateEndpointAsync(ctx workflow.Context, input *sagemaker.UpdateEndpointInput) *SagemakerUpdateEndpointResult

    UpdateEndpointWeightsAndCapacities(ctx workflow.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) (*sagemaker.UpdateEndpointWeightsAndCapacitiesOutput, error)
    UpdateEndpointWeightsAndCapacitiesAsync(ctx workflow.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) *SagemakerUpdateEndpointWeightsAndCapacitiesResult

    UpdateExperiment(ctx workflow.Context, input *sagemaker.UpdateExperimentInput) (*sagemaker.UpdateExperimentOutput, error)
    UpdateExperimentAsync(ctx workflow.Context, input *sagemaker.UpdateExperimentInput) *SagemakerUpdateExperimentResult

    UpdateMonitoringSchedule(ctx workflow.Context, input *sagemaker.UpdateMonitoringScheduleInput) (*sagemaker.UpdateMonitoringScheduleOutput, error)
    UpdateMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.UpdateMonitoringScheduleInput) *SagemakerUpdateMonitoringScheduleResult

    UpdateNotebookInstance(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceInput) (*sagemaker.UpdateNotebookInstanceOutput, error)
    UpdateNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceInput) *SagemakerUpdateNotebookInstanceResult

    UpdateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) (*sagemaker.UpdateNotebookInstanceLifecycleConfigOutput, error)
    UpdateNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) *SagemakerUpdateNotebookInstanceLifecycleConfigResult

    UpdateTrial(ctx workflow.Context, input *sagemaker.UpdateTrialInput) (*sagemaker.UpdateTrialOutput, error)
    UpdateTrialAsync(ctx workflow.Context, input *sagemaker.UpdateTrialInput) *SagemakerUpdateTrialResult

    UpdateTrialComponent(ctx workflow.Context, input *sagemaker.UpdateTrialComponentInput) (*sagemaker.UpdateTrialComponentOutput, error)
    UpdateTrialComponentAsync(ctx workflow.Context, input *sagemaker.UpdateTrialComponentInput) *SagemakerUpdateTrialComponentResult

    UpdateUserProfile(ctx workflow.Context, input *sagemaker.UpdateUserProfileInput) (*sagemaker.UpdateUserProfileOutput, error)
    UpdateUserProfileAsync(ctx workflow.Context, input *sagemaker.UpdateUserProfileInput) *SagemakerUpdateUserProfileResult

    UpdateWorkforce(ctx workflow.Context, input *sagemaker.UpdateWorkforceInput) (*sagemaker.UpdateWorkforceOutput, error)
    UpdateWorkforceAsync(ctx workflow.Context, input *sagemaker.UpdateWorkforceInput) *SagemakerUpdateWorkforceResult

    UpdateWorkteam(ctx workflow.Context, input *sagemaker.UpdateWorkteamInput) (*sagemaker.UpdateWorkteamOutput, error)
    UpdateWorkteamAsync(ctx workflow.Context, input *sagemaker.UpdateWorkteamInput) *SagemakerUpdateWorkteamResult

    WaitUntilEndpointDeleted(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error
    WaitUntilEndpointInService(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error
    WaitUntilNotebookInstanceDeleted(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
    WaitUntilNotebookInstanceInService(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
    WaitUntilNotebookInstanceStopped(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
    WaitUntilProcessingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) error
    WaitUntilTrainingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) error
    WaitUntilTransformJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) error}
type SagemakerAddTagsResult struct {
	Result workflow.Future
}

func (r *SagemakerAddTagsResult) Get(ctx workflow.Context) (*sagemaker.AddTagsOutput, error) {
    var output sagemaker.AddTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerAssociateTrialComponentResult struct {
	Result workflow.Future
}

func (r *SagemakerAssociateTrialComponentResult) Get(ctx workflow.Context) (*sagemaker.AssociateTrialComponentOutput, error) {
    var output sagemaker.AssociateTrialComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateAlgorithmResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateAlgorithmResult) Get(ctx workflow.Context) (*sagemaker.CreateAlgorithmOutput, error) {
    var output sagemaker.CreateAlgorithmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateAppResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateAppResult) Get(ctx workflow.Context) (*sagemaker.CreateAppOutput, error) {
    var output sagemaker.CreateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateAutoMLJobResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateAutoMLJobResult) Get(ctx workflow.Context) (*sagemaker.CreateAutoMLJobOutput, error) {
    var output sagemaker.CreateAutoMLJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateCodeRepositoryResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateCodeRepositoryResult) Get(ctx workflow.Context) (*sagemaker.CreateCodeRepositoryOutput, error) {
    var output sagemaker.CreateCodeRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateCompilationJobResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateCompilationJobResult) Get(ctx workflow.Context) (*sagemaker.CreateCompilationJobOutput, error) {
    var output sagemaker.CreateCompilationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateDomainResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateDomainResult) Get(ctx workflow.Context) (*sagemaker.CreateDomainOutput, error) {
    var output sagemaker.CreateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateEndpointResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateEndpointResult) Get(ctx workflow.Context) (*sagemaker.CreateEndpointOutput, error) {
    var output sagemaker.CreateEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateEndpointConfigResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateEndpointConfigResult) Get(ctx workflow.Context) (*sagemaker.CreateEndpointConfigOutput, error) {
    var output sagemaker.CreateEndpointConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateExperimentResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateExperimentResult) Get(ctx workflow.Context) (*sagemaker.CreateExperimentOutput, error) {
    var output sagemaker.CreateExperimentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateFlowDefinitionResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateFlowDefinitionResult) Get(ctx workflow.Context) (*sagemaker.CreateFlowDefinitionOutput, error) {
    var output sagemaker.CreateFlowDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateHumanTaskUiResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateHumanTaskUiResult) Get(ctx workflow.Context) (*sagemaker.CreateHumanTaskUiOutput, error) {
    var output sagemaker.CreateHumanTaskUiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateHyperParameterTuningJobResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateHyperParameterTuningJobResult) Get(ctx workflow.Context) (*sagemaker.CreateHyperParameterTuningJobOutput, error) {
    var output sagemaker.CreateHyperParameterTuningJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateLabelingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateLabelingJobResult) Get(ctx workflow.Context) (*sagemaker.CreateLabelingJobOutput, error) {
    var output sagemaker.CreateLabelingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateModelResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateModelResult) Get(ctx workflow.Context) (*sagemaker.CreateModelOutput, error) {
    var output sagemaker.CreateModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateModelPackageResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateModelPackageResult) Get(ctx workflow.Context) (*sagemaker.CreateModelPackageOutput, error) {
    var output sagemaker.CreateModelPackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateMonitoringScheduleResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateMonitoringScheduleResult) Get(ctx workflow.Context) (*sagemaker.CreateMonitoringScheduleOutput, error) {
    var output sagemaker.CreateMonitoringScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateNotebookInstanceResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateNotebookInstanceResult) Get(ctx workflow.Context) (*sagemaker.CreateNotebookInstanceOutput, error) {
    var output sagemaker.CreateNotebookInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateNotebookInstanceLifecycleConfigResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateNotebookInstanceLifecycleConfigResult) Get(ctx workflow.Context) (*sagemaker.CreateNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.CreateNotebookInstanceLifecycleConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreatePresignedDomainUrlResult struct {
	Result workflow.Future
}

func (r *SagemakerCreatePresignedDomainUrlResult) Get(ctx workflow.Context) (*sagemaker.CreatePresignedDomainUrlOutput, error) {
    var output sagemaker.CreatePresignedDomainUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreatePresignedNotebookInstanceUrlResult struct {
	Result workflow.Future
}

func (r *SagemakerCreatePresignedNotebookInstanceUrlResult) Get(ctx workflow.Context) (*sagemaker.CreatePresignedNotebookInstanceUrlOutput, error) {
    var output sagemaker.CreatePresignedNotebookInstanceUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateProcessingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateProcessingJobResult) Get(ctx workflow.Context) (*sagemaker.CreateProcessingJobOutput, error) {
    var output sagemaker.CreateProcessingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateTrainingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateTrainingJobResult) Get(ctx workflow.Context) (*sagemaker.CreateTrainingJobOutput, error) {
    var output sagemaker.CreateTrainingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateTransformJobResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateTransformJobResult) Get(ctx workflow.Context) (*sagemaker.CreateTransformJobOutput, error) {
    var output sagemaker.CreateTransformJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateTrialResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateTrialResult) Get(ctx workflow.Context) (*sagemaker.CreateTrialOutput, error) {
    var output sagemaker.CreateTrialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateTrialComponentResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateTrialComponentResult) Get(ctx workflow.Context) (*sagemaker.CreateTrialComponentOutput, error) {
    var output sagemaker.CreateTrialComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateUserProfileResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateUserProfileResult) Get(ctx workflow.Context) (*sagemaker.CreateUserProfileOutput, error) {
    var output sagemaker.CreateUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateWorkforceResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateWorkforceResult) Get(ctx workflow.Context) (*sagemaker.CreateWorkforceOutput, error) {
    var output sagemaker.CreateWorkforceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerCreateWorkteamResult struct {
	Result workflow.Future
}

func (r *SagemakerCreateWorkteamResult) Get(ctx workflow.Context) (*sagemaker.CreateWorkteamOutput, error) {
    var output sagemaker.CreateWorkteamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteAlgorithmResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteAlgorithmResult) Get(ctx workflow.Context) (*sagemaker.DeleteAlgorithmOutput, error) {
    var output sagemaker.DeleteAlgorithmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteAppResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteAppResult) Get(ctx workflow.Context) (*sagemaker.DeleteAppOutput, error) {
    var output sagemaker.DeleteAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteCodeRepositoryResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteCodeRepositoryResult) Get(ctx workflow.Context) (*sagemaker.DeleteCodeRepositoryOutput, error) {
    var output sagemaker.DeleteCodeRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteDomainResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteDomainResult) Get(ctx workflow.Context) (*sagemaker.DeleteDomainOutput, error) {
    var output sagemaker.DeleteDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteEndpointResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteEndpointResult) Get(ctx workflow.Context) (*sagemaker.DeleteEndpointOutput, error) {
    var output sagemaker.DeleteEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteEndpointConfigResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteEndpointConfigResult) Get(ctx workflow.Context) (*sagemaker.DeleteEndpointConfigOutput, error) {
    var output sagemaker.DeleteEndpointConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteExperimentResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteExperimentResult) Get(ctx workflow.Context) (*sagemaker.DeleteExperimentOutput, error) {
    var output sagemaker.DeleteExperimentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteFlowDefinitionResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteFlowDefinitionResult) Get(ctx workflow.Context) (*sagemaker.DeleteFlowDefinitionOutput, error) {
    var output sagemaker.DeleteFlowDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteHumanTaskUiResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteHumanTaskUiResult) Get(ctx workflow.Context) (*sagemaker.DeleteHumanTaskUiOutput, error) {
    var output sagemaker.DeleteHumanTaskUiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteModelResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteModelResult) Get(ctx workflow.Context) (*sagemaker.DeleteModelOutput, error) {
    var output sagemaker.DeleteModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteModelPackageResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteModelPackageResult) Get(ctx workflow.Context) (*sagemaker.DeleteModelPackageOutput, error) {
    var output sagemaker.DeleteModelPackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteMonitoringScheduleResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteMonitoringScheduleResult) Get(ctx workflow.Context) (*sagemaker.DeleteMonitoringScheduleOutput, error) {
    var output sagemaker.DeleteMonitoringScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteNotebookInstanceResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteNotebookInstanceResult) Get(ctx workflow.Context) (*sagemaker.DeleteNotebookInstanceOutput, error) {
    var output sagemaker.DeleteNotebookInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteNotebookInstanceLifecycleConfigResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteNotebookInstanceLifecycleConfigResult) Get(ctx workflow.Context) (*sagemaker.DeleteNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.DeleteNotebookInstanceLifecycleConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteTagsResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteTagsResult) Get(ctx workflow.Context) (*sagemaker.DeleteTagsOutput, error) {
    var output sagemaker.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteTrialResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteTrialResult) Get(ctx workflow.Context) (*sagemaker.DeleteTrialOutput, error) {
    var output sagemaker.DeleteTrialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteTrialComponentResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteTrialComponentResult) Get(ctx workflow.Context) (*sagemaker.DeleteTrialComponentOutput, error) {
    var output sagemaker.DeleteTrialComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteUserProfileResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteUserProfileResult) Get(ctx workflow.Context) (*sagemaker.DeleteUserProfileOutput, error) {
    var output sagemaker.DeleteUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteWorkforceResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteWorkforceResult) Get(ctx workflow.Context) (*sagemaker.DeleteWorkforceOutput, error) {
    var output sagemaker.DeleteWorkforceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDeleteWorkteamResult struct {
	Result workflow.Future
}

func (r *SagemakerDeleteWorkteamResult) Get(ctx workflow.Context) (*sagemaker.DeleteWorkteamOutput, error) {
    var output sagemaker.DeleteWorkteamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeAlgorithmResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeAlgorithmResult) Get(ctx workflow.Context) (*sagemaker.DescribeAlgorithmOutput, error) {
    var output sagemaker.DescribeAlgorithmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeAppResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeAppResult) Get(ctx workflow.Context) (*sagemaker.DescribeAppOutput, error) {
    var output sagemaker.DescribeAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeAutoMLJobResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeAutoMLJobResult) Get(ctx workflow.Context) (*sagemaker.DescribeAutoMLJobOutput, error) {
    var output sagemaker.DescribeAutoMLJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeCodeRepositoryResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeCodeRepositoryResult) Get(ctx workflow.Context) (*sagemaker.DescribeCodeRepositoryOutput, error) {
    var output sagemaker.DescribeCodeRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeCompilationJobResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeCompilationJobResult) Get(ctx workflow.Context) (*sagemaker.DescribeCompilationJobOutput, error) {
    var output sagemaker.DescribeCompilationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeDomainResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeDomainResult) Get(ctx workflow.Context) (*sagemaker.DescribeDomainOutput, error) {
    var output sagemaker.DescribeDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeEndpointResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeEndpointResult) Get(ctx workflow.Context) (*sagemaker.DescribeEndpointOutput, error) {
    var output sagemaker.DescribeEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeEndpointConfigResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeEndpointConfigResult) Get(ctx workflow.Context) (*sagemaker.DescribeEndpointConfigOutput, error) {
    var output sagemaker.DescribeEndpointConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeExperimentResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeExperimentResult) Get(ctx workflow.Context) (*sagemaker.DescribeExperimentOutput, error) {
    var output sagemaker.DescribeExperimentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeFlowDefinitionResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeFlowDefinitionResult) Get(ctx workflow.Context) (*sagemaker.DescribeFlowDefinitionOutput, error) {
    var output sagemaker.DescribeFlowDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeHumanTaskUiResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeHumanTaskUiResult) Get(ctx workflow.Context) (*sagemaker.DescribeHumanTaskUiOutput, error) {
    var output sagemaker.DescribeHumanTaskUiOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeHyperParameterTuningJobResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeHyperParameterTuningJobResult) Get(ctx workflow.Context) (*sagemaker.DescribeHyperParameterTuningJobOutput, error) {
    var output sagemaker.DescribeHyperParameterTuningJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeLabelingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeLabelingJobResult) Get(ctx workflow.Context) (*sagemaker.DescribeLabelingJobOutput, error) {
    var output sagemaker.DescribeLabelingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeModelResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeModelResult) Get(ctx workflow.Context) (*sagemaker.DescribeModelOutput, error) {
    var output sagemaker.DescribeModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeModelPackageResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeModelPackageResult) Get(ctx workflow.Context) (*sagemaker.DescribeModelPackageOutput, error) {
    var output sagemaker.DescribeModelPackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeMonitoringScheduleResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeMonitoringScheduleResult) Get(ctx workflow.Context) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
    var output sagemaker.DescribeMonitoringScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeNotebookInstanceResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeNotebookInstanceResult) Get(ctx workflow.Context) (*sagemaker.DescribeNotebookInstanceOutput, error) {
    var output sagemaker.DescribeNotebookInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeNotebookInstanceLifecycleConfigResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeNotebookInstanceLifecycleConfigResult) Get(ctx workflow.Context) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.DescribeNotebookInstanceLifecycleConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeProcessingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeProcessingJobResult) Get(ctx workflow.Context) (*sagemaker.DescribeProcessingJobOutput, error) {
    var output sagemaker.DescribeProcessingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeSubscribedWorkteamResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeSubscribedWorkteamResult) Get(ctx workflow.Context) (*sagemaker.DescribeSubscribedWorkteamOutput, error) {
    var output sagemaker.DescribeSubscribedWorkteamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeTrainingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeTrainingJobResult) Get(ctx workflow.Context) (*sagemaker.DescribeTrainingJobOutput, error) {
    var output sagemaker.DescribeTrainingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeTransformJobResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeTransformJobResult) Get(ctx workflow.Context) (*sagemaker.DescribeTransformJobOutput, error) {
    var output sagemaker.DescribeTransformJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeTrialResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeTrialResult) Get(ctx workflow.Context) (*sagemaker.DescribeTrialOutput, error) {
    var output sagemaker.DescribeTrialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeTrialComponentResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeTrialComponentResult) Get(ctx workflow.Context) (*sagemaker.DescribeTrialComponentOutput, error) {
    var output sagemaker.DescribeTrialComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeUserProfileResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeUserProfileResult) Get(ctx workflow.Context) (*sagemaker.DescribeUserProfileOutput, error) {
    var output sagemaker.DescribeUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeWorkforceResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeWorkforceResult) Get(ctx workflow.Context) (*sagemaker.DescribeWorkforceOutput, error) {
    var output sagemaker.DescribeWorkforceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDescribeWorkteamResult struct {
	Result workflow.Future
}

func (r *SagemakerDescribeWorkteamResult) Get(ctx workflow.Context) (*sagemaker.DescribeWorkteamOutput, error) {
    var output sagemaker.DescribeWorkteamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerDisassociateTrialComponentResult struct {
	Result workflow.Future
}

func (r *SagemakerDisassociateTrialComponentResult) Get(ctx workflow.Context) (*sagemaker.DisassociateTrialComponentOutput, error) {
    var output sagemaker.DisassociateTrialComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerGetSearchSuggestionsResult struct {
	Result workflow.Future
}

func (r *SagemakerGetSearchSuggestionsResult) Get(ctx workflow.Context) (*sagemaker.GetSearchSuggestionsOutput, error) {
    var output sagemaker.GetSearchSuggestionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListAlgorithmsResult struct {
	Result workflow.Future
}

func (r *SagemakerListAlgorithmsResult) Get(ctx workflow.Context) (*sagemaker.ListAlgorithmsOutput, error) {
    var output sagemaker.ListAlgorithmsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListAppsResult struct {
	Result workflow.Future
}

func (r *SagemakerListAppsResult) Get(ctx workflow.Context) (*sagemaker.ListAppsOutput, error) {
    var output sagemaker.ListAppsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListAutoMLJobsResult struct {
	Result workflow.Future
}

func (r *SagemakerListAutoMLJobsResult) Get(ctx workflow.Context) (*sagemaker.ListAutoMLJobsOutput, error) {
    var output sagemaker.ListAutoMLJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListCandidatesForAutoMLJobResult struct {
	Result workflow.Future
}

func (r *SagemakerListCandidatesForAutoMLJobResult) Get(ctx workflow.Context) (*sagemaker.ListCandidatesForAutoMLJobOutput, error) {
    var output sagemaker.ListCandidatesForAutoMLJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListCodeRepositoriesResult struct {
	Result workflow.Future
}

func (r *SagemakerListCodeRepositoriesResult) Get(ctx workflow.Context) (*sagemaker.ListCodeRepositoriesOutput, error) {
    var output sagemaker.ListCodeRepositoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListCompilationJobsResult struct {
	Result workflow.Future
}

func (r *SagemakerListCompilationJobsResult) Get(ctx workflow.Context) (*sagemaker.ListCompilationJobsOutput, error) {
    var output sagemaker.ListCompilationJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListDomainsResult struct {
	Result workflow.Future
}

func (r *SagemakerListDomainsResult) Get(ctx workflow.Context) (*sagemaker.ListDomainsOutput, error) {
    var output sagemaker.ListDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListEndpointConfigsResult struct {
	Result workflow.Future
}

func (r *SagemakerListEndpointConfigsResult) Get(ctx workflow.Context) (*sagemaker.ListEndpointConfigsOutput, error) {
    var output sagemaker.ListEndpointConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListEndpointsResult struct {
	Result workflow.Future
}

func (r *SagemakerListEndpointsResult) Get(ctx workflow.Context) (*sagemaker.ListEndpointsOutput, error) {
    var output sagemaker.ListEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListExperimentsResult struct {
	Result workflow.Future
}

func (r *SagemakerListExperimentsResult) Get(ctx workflow.Context) (*sagemaker.ListExperimentsOutput, error) {
    var output sagemaker.ListExperimentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListFlowDefinitionsResult struct {
	Result workflow.Future
}

func (r *SagemakerListFlowDefinitionsResult) Get(ctx workflow.Context) (*sagemaker.ListFlowDefinitionsOutput, error) {
    var output sagemaker.ListFlowDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListHumanTaskUisResult struct {
	Result workflow.Future
}

func (r *SagemakerListHumanTaskUisResult) Get(ctx workflow.Context) (*sagemaker.ListHumanTaskUisOutput, error) {
    var output sagemaker.ListHumanTaskUisOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListHyperParameterTuningJobsResult struct {
	Result workflow.Future
}

func (r *SagemakerListHyperParameterTuningJobsResult) Get(ctx workflow.Context) (*sagemaker.ListHyperParameterTuningJobsOutput, error) {
    var output sagemaker.ListHyperParameterTuningJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListLabelingJobsResult struct {
	Result workflow.Future
}

func (r *SagemakerListLabelingJobsResult) Get(ctx workflow.Context) (*sagemaker.ListLabelingJobsOutput, error) {
    var output sagemaker.ListLabelingJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListLabelingJobsForWorkteamResult struct {
	Result workflow.Future
}

func (r *SagemakerListLabelingJobsForWorkteamResult) Get(ctx workflow.Context) (*sagemaker.ListLabelingJobsForWorkteamOutput, error) {
    var output sagemaker.ListLabelingJobsForWorkteamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListModelPackagesResult struct {
	Result workflow.Future
}

func (r *SagemakerListModelPackagesResult) Get(ctx workflow.Context) (*sagemaker.ListModelPackagesOutput, error) {
    var output sagemaker.ListModelPackagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListModelsResult struct {
	Result workflow.Future
}

func (r *SagemakerListModelsResult) Get(ctx workflow.Context) (*sagemaker.ListModelsOutput, error) {
    var output sagemaker.ListModelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListMonitoringExecutionsResult struct {
	Result workflow.Future
}

func (r *SagemakerListMonitoringExecutionsResult) Get(ctx workflow.Context) (*sagemaker.ListMonitoringExecutionsOutput, error) {
    var output sagemaker.ListMonitoringExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListMonitoringSchedulesResult struct {
	Result workflow.Future
}

func (r *SagemakerListMonitoringSchedulesResult) Get(ctx workflow.Context) (*sagemaker.ListMonitoringSchedulesOutput, error) {
    var output sagemaker.ListMonitoringSchedulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListNotebookInstanceLifecycleConfigsResult struct {
	Result workflow.Future
}

func (r *SagemakerListNotebookInstanceLifecycleConfigsResult) Get(ctx workflow.Context) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error) {
    var output sagemaker.ListNotebookInstanceLifecycleConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListNotebookInstancesResult struct {
	Result workflow.Future
}

func (r *SagemakerListNotebookInstancesResult) Get(ctx workflow.Context) (*sagemaker.ListNotebookInstancesOutput, error) {
    var output sagemaker.ListNotebookInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListProcessingJobsResult struct {
	Result workflow.Future
}

func (r *SagemakerListProcessingJobsResult) Get(ctx workflow.Context) (*sagemaker.ListProcessingJobsOutput, error) {
    var output sagemaker.ListProcessingJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListSubscribedWorkteamsResult struct {
	Result workflow.Future
}

func (r *SagemakerListSubscribedWorkteamsResult) Get(ctx workflow.Context) (*sagemaker.ListSubscribedWorkteamsOutput, error) {
    var output sagemaker.ListSubscribedWorkteamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListTagsResult struct {
	Result workflow.Future
}

func (r *SagemakerListTagsResult) Get(ctx workflow.Context) (*sagemaker.ListTagsOutput, error) {
    var output sagemaker.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListTrainingJobsResult struct {
	Result workflow.Future
}

func (r *SagemakerListTrainingJobsResult) Get(ctx workflow.Context) (*sagemaker.ListTrainingJobsOutput, error) {
    var output sagemaker.ListTrainingJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListTrainingJobsForHyperParameterTuningJobResult struct {
	Result workflow.Future
}

func (r *SagemakerListTrainingJobsForHyperParameterTuningJobResult) Get(ctx workflow.Context) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error) {
    var output sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListTransformJobsResult struct {
	Result workflow.Future
}

func (r *SagemakerListTransformJobsResult) Get(ctx workflow.Context) (*sagemaker.ListTransformJobsOutput, error) {
    var output sagemaker.ListTransformJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListTrialComponentsResult struct {
	Result workflow.Future
}

func (r *SagemakerListTrialComponentsResult) Get(ctx workflow.Context) (*sagemaker.ListTrialComponentsOutput, error) {
    var output sagemaker.ListTrialComponentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListTrialsResult struct {
	Result workflow.Future
}

func (r *SagemakerListTrialsResult) Get(ctx workflow.Context) (*sagemaker.ListTrialsOutput, error) {
    var output sagemaker.ListTrialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListUserProfilesResult struct {
	Result workflow.Future
}

func (r *SagemakerListUserProfilesResult) Get(ctx workflow.Context) (*sagemaker.ListUserProfilesOutput, error) {
    var output sagemaker.ListUserProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListWorkforcesResult struct {
	Result workflow.Future
}

func (r *SagemakerListWorkforcesResult) Get(ctx workflow.Context) (*sagemaker.ListWorkforcesOutput, error) {
    var output sagemaker.ListWorkforcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerListWorkteamsResult struct {
	Result workflow.Future
}

func (r *SagemakerListWorkteamsResult) Get(ctx workflow.Context) (*sagemaker.ListWorkteamsOutput, error) {
    var output sagemaker.ListWorkteamsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerRenderUiTemplateResult struct {
	Result workflow.Future
}

func (r *SagemakerRenderUiTemplateResult) Get(ctx workflow.Context) (*sagemaker.RenderUiTemplateOutput, error) {
    var output sagemaker.RenderUiTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerSearchResult struct {
	Result workflow.Future
}

func (r *SagemakerSearchResult) Get(ctx workflow.Context) (*sagemaker.SearchOutput, error) {
    var output sagemaker.SearchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStartMonitoringScheduleResult struct {
	Result workflow.Future
}

func (r *SagemakerStartMonitoringScheduleResult) Get(ctx workflow.Context) (*sagemaker.StartMonitoringScheduleOutput, error) {
    var output sagemaker.StartMonitoringScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStartNotebookInstanceResult struct {
	Result workflow.Future
}

func (r *SagemakerStartNotebookInstanceResult) Get(ctx workflow.Context) (*sagemaker.StartNotebookInstanceOutput, error) {
    var output sagemaker.StartNotebookInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopAutoMLJobResult struct {
	Result workflow.Future
}

func (r *SagemakerStopAutoMLJobResult) Get(ctx workflow.Context) (*sagemaker.StopAutoMLJobOutput, error) {
    var output sagemaker.StopAutoMLJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopCompilationJobResult struct {
	Result workflow.Future
}

func (r *SagemakerStopCompilationJobResult) Get(ctx workflow.Context) (*sagemaker.StopCompilationJobOutput, error) {
    var output sagemaker.StopCompilationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopHyperParameterTuningJobResult struct {
	Result workflow.Future
}

func (r *SagemakerStopHyperParameterTuningJobResult) Get(ctx workflow.Context) (*sagemaker.StopHyperParameterTuningJobOutput, error) {
    var output sagemaker.StopHyperParameterTuningJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopLabelingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerStopLabelingJobResult) Get(ctx workflow.Context) (*sagemaker.StopLabelingJobOutput, error) {
    var output sagemaker.StopLabelingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopMonitoringScheduleResult struct {
	Result workflow.Future
}

func (r *SagemakerStopMonitoringScheduleResult) Get(ctx workflow.Context) (*sagemaker.StopMonitoringScheduleOutput, error) {
    var output sagemaker.StopMonitoringScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopNotebookInstanceResult struct {
	Result workflow.Future
}

func (r *SagemakerStopNotebookInstanceResult) Get(ctx workflow.Context) (*sagemaker.StopNotebookInstanceOutput, error) {
    var output sagemaker.StopNotebookInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopProcessingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerStopProcessingJobResult) Get(ctx workflow.Context) (*sagemaker.StopProcessingJobOutput, error) {
    var output sagemaker.StopProcessingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopTrainingJobResult struct {
	Result workflow.Future
}

func (r *SagemakerStopTrainingJobResult) Get(ctx workflow.Context) (*sagemaker.StopTrainingJobOutput, error) {
    var output sagemaker.StopTrainingJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerStopTransformJobResult struct {
	Result workflow.Future
}

func (r *SagemakerStopTransformJobResult) Get(ctx workflow.Context) (*sagemaker.StopTransformJobOutput, error) {
    var output sagemaker.StopTransformJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateCodeRepositoryResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateCodeRepositoryResult) Get(ctx workflow.Context) (*sagemaker.UpdateCodeRepositoryOutput, error) {
    var output sagemaker.UpdateCodeRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateDomainResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateDomainResult) Get(ctx workflow.Context) (*sagemaker.UpdateDomainOutput, error) {
    var output sagemaker.UpdateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateEndpointResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateEndpointResult) Get(ctx workflow.Context) (*sagemaker.UpdateEndpointOutput, error) {
    var output sagemaker.UpdateEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateEndpointWeightsAndCapacitiesResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateEndpointWeightsAndCapacitiesResult) Get(ctx workflow.Context) (*sagemaker.UpdateEndpointWeightsAndCapacitiesOutput, error) {
    var output sagemaker.UpdateEndpointWeightsAndCapacitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateExperimentResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateExperimentResult) Get(ctx workflow.Context) (*sagemaker.UpdateExperimentOutput, error) {
    var output sagemaker.UpdateExperimentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateMonitoringScheduleResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateMonitoringScheduleResult) Get(ctx workflow.Context) (*sagemaker.UpdateMonitoringScheduleOutput, error) {
    var output sagemaker.UpdateMonitoringScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateNotebookInstanceResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateNotebookInstanceResult) Get(ctx workflow.Context) (*sagemaker.UpdateNotebookInstanceOutput, error) {
    var output sagemaker.UpdateNotebookInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateNotebookInstanceLifecycleConfigResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateNotebookInstanceLifecycleConfigResult) Get(ctx workflow.Context) (*sagemaker.UpdateNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.UpdateNotebookInstanceLifecycleConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateTrialResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateTrialResult) Get(ctx workflow.Context) (*sagemaker.UpdateTrialOutput, error) {
    var output sagemaker.UpdateTrialOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateTrialComponentResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateTrialComponentResult) Get(ctx workflow.Context) (*sagemaker.UpdateTrialComponentOutput, error) {
    var output sagemaker.UpdateTrialComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateUserProfileResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateUserProfileResult) Get(ctx workflow.Context) (*sagemaker.UpdateUserProfileOutput, error) {
    var output sagemaker.UpdateUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateWorkforceResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateWorkforceResult) Get(ctx workflow.Context) (*sagemaker.UpdateWorkforceOutput, error) {
    var output sagemaker.UpdateWorkforceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SagemakerUpdateWorkteamResult struct {
	Result workflow.Future
}

func (r *SagemakerUpdateWorkteamResult) Get(ctx workflow.Context) (*sagemaker.UpdateWorkteamOutput, error) {
    var output sagemaker.UpdateWorkteamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SageMakerStub struct {
    activities SageMakerClient
}

func NewSageMakerStub() SageMakerClient {
    return &SageMakerStub{}
}

func (a *SageMakerStub) AddTags(ctx workflow.Context, input *sagemaker.AddTagsInput) (*sagemaker.AddTagsOutput, error) {
    var output sagemaker.AddTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTags, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) AssociateTrialComponent(ctx workflow.Context, input *sagemaker.AssociateTrialComponentInput) (*sagemaker.AssociateTrialComponentOutput, error) {
    var output sagemaker.AssociateTrialComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateTrialComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateAlgorithm(ctx workflow.Context, input *sagemaker.CreateAlgorithmInput) (*sagemaker.CreateAlgorithmOutput, error) {
    var output sagemaker.CreateAlgorithmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAlgorithm, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateApp(ctx workflow.Context, input *sagemaker.CreateAppInput) (*sagemaker.CreateAppOutput, error) {
    var output sagemaker.CreateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateAutoMLJob(ctx workflow.Context, input *sagemaker.CreateAutoMLJobInput) (*sagemaker.CreateAutoMLJobOutput, error) {
    var output sagemaker.CreateAutoMLJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAutoMLJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateCodeRepository(ctx workflow.Context, input *sagemaker.CreateCodeRepositoryInput) (*sagemaker.CreateCodeRepositoryOutput, error) {
    var output sagemaker.CreateCodeRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCodeRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateCompilationJob(ctx workflow.Context, input *sagemaker.CreateCompilationJobInput) (*sagemaker.CreateCompilationJobOutput, error) {
    var output sagemaker.CreateCompilationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCompilationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateDomain(ctx workflow.Context, input *sagemaker.CreateDomainInput) (*sagemaker.CreateDomainOutput, error) {
    var output sagemaker.CreateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateEndpoint(ctx workflow.Context, input *sagemaker.CreateEndpointInput) (*sagemaker.CreateEndpointOutput, error) {
    var output sagemaker.CreateEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateEndpointConfig(ctx workflow.Context, input *sagemaker.CreateEndpointConfigInput) (*sagemaker.CreateEndpointConfigOutput, error) {
    var output sagemaker.CreateEndpointConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEndpointConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateExperiment(ctx workflow.Context, input *sagemaker.CreateExperimentInput) (*sagemaker.CreateExperimentOutput, error) {
    var output sagemaker.CreateExperimentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateExperiment, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateFlowDefinition(ctx workflow.Context, input *sagemaker.CreateFlowDefinitionInput) (*sagemaker.CreateFlowDefinitionOutput, error) {
    var output sagemaker.CreateFlowDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFlowDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateHumanTaskUi(ctx workflow.Context, input *sagemaker.CreateHumanTaskUiInput) (*sagemaker.CreateHumanTaskUiOutput, error) {
    var output sagemaker.CreateHumanTaskUiOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHumanTaskUi, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.CreateHyperParameterTuningJobInput) (*sagemaker.CreateHyperParameterTuningJobOutput, error) {
    var output sagemaker.CreateHyperParameterTuningJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHyperParameterTuningJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateLabelingJob(ctx workflow.Context, input *sagemaker.CreateLabelingJobInput) (*sagemaker.CreateLabelingJobOutput, error) {
    var output sagemaker.CreateLabelingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLabelingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateModel(ctx workflow.Context, input *sagemaker.CreateModelInput) (*sagemaker.CreateModelOutput, error) {
    var output sagemaker.CreateModelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateModel, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateModelPackage(ctx workflow.Context, input *sagemaker.CreateModelPackageInput) (*sagemaker.CreateModelPackageOutput, error) {
    var output sagemaker.CreateModelPackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateModelPackage, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateMonitoringSchedule(ctx workflow.Context, input *sagemaker.CreateMonitoringScheduleInput) (*sagemaker.CreateMonitoringScheduleOutput, error) {
    var output sagemaker.CreateMonitoringScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMonitoringSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateNotebookInstance(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceInput) (*sagemaker.CreateNotebookInstanceOutput, error) {
    var output sagemaker.CreateNotebookInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNotebookInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) (*sagemaker.CreateNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.CreateNotebookInstanceLifecycleConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNotebookInstanceLifecycleConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreatePresignedDomainUrl(ctx workflow.Context, input *sagemaker.CreatePresignedDomainUrlInput) (*sagemaker.CreatePresignedDomainUrlOutput, error) {
    var output sagemaker.CreatePresignedDomainUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePresignedDomainUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreatePresignedNotebookInstanceUrl(ctx workflow.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) (*sagemaker.CreatePresignedNotebookInstanceUrlOutput, error) {
    var output sagemaker.CreatePresignedNotebookInstanceUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePresignedNotebookInstanceUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateProcessingJob(ctx workflow.Context, input *sagemaker.CreateProcessingJobInput) (*sagemaker.CreateProcessingJobOutput, error) {
    var output sagemaker.CreateProcessingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProcessingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateTrainingJob(ctx workflow.Context, input *sagemaker.CreateTrainingJobInput) (*sagemaker.CreateTrainingJobOutput, error) {
    var output sagemaker.CreateTrainingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrainingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateTransformJob(ctx workflow.Context, input *sagemaker.CreateTransformJobInput) (*sagemaker.CreateTransformJobOutput, error) {
    var output sagemaker.CreateTransformJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransformJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateTrial(ctx workflow.Context, input *sagemaker.CreateTrialInput) (*sagemaker.CreateTrialOutput, error) {
    var output sagemaker.CreateTrialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrial, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateTrialComponent(ctx workflow.Context, input *sagemaker.CreateTrialComponentInput) (*sagemaker.CreateTrialComponentOutput, error) {
    var output sagemaker.CreateTrialComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrialComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateUserProfile(ctx workflow.Context, input *sagemaker.CreateUserProfileInput) (*sagemaker.CreateUserProfileOutput, error) {
    var output sagemaker.CreateUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateWorkforce(ctx workflow.Context, input *sagemaker.CreateWorkforceInput) (*sagemaker.CreateWorkforceOutput, error) {
    var output sagemaker.CreateWorkforceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorkforce, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) CreateWorkteam(ctx workflow.Context, input *sagemaker.CreateWorkteamInput) (*sagemaker.CreateWorkteamOutput, error) {
    var output sagemaker.CreateWorkteamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorkteam, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteAlgorithm(ctx workflow.Context, input *sagemaker.DeleteAlgorithmInput) (*sagemaker.DeleteAlgorithmOutput, error) {
    var output sagemaker.DeleteAlgorithmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAlgorithm, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteApp(ctx workflow.Context, input *sagemaker.DeleteAppInput) (*sagemaker.DeleteAppOutput, error) {
    var output sagemaker.DeleteAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteCodeRepository(ctx workflow.Context, input *sagemaker.DeleteCodeRepositoryInput) (*sagemaker.DeleteCodeRepositoryOutput, error) {
    var output sagemaker.DeleteCodeRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCodeRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteDomain(ctx workflow.Context, input *sagemaker.DeleteDomainInput) (*sagemaker.DeleteDomainOutput, error) {
    var output sagemaker.DeleteDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteEndpoint(ctx workflow.Context, input *sagemaker.DeleteEndpointInput) (*sagemaker.DeleteEndpointOutput, error) {
    var output sagemaker.DeleteEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteEndpointConfig(ctx workflow.Context, input *sagemaker.DeleteEndpointConfigInput) (*sagemaker.DeleteEndpointConfigOutput, error) {
    var output sagemaker.DeleteEndpointConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpointConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteExperiment(ctx workflow.Context, input *sagemaker.DeleteExperimentInput) (*sagemaker.DeleteExperimentOutput, error) {
    var output sagemaker.DeleteExperimentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteExperiment, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteFlowDefinition(ctx workflow.Context, input *sagemaker.DeleteFlowDefinitionInput) (*sagemaker.DeleteFlowDefinitionOutput, error) {
    var output sagemaker.DeleteFlowDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFlowDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteHumanTaskUi(ctx workflow.Context, input *sagemaker.DeleteHumanTaskUiInput) (*sagemaker.DeleteHumanTaskUiOutput, error) {
    var output sagemaker.DeleteHumanTaskUiOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHumanTaskUi, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteModel(ctx workflow.Context, input *sagemaker.DeleteModelInput) (*sagemaker.DeleteModelOutput, error) {
    var output sagemaker.DeleteModelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteModel, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteModelPackage(ctx workflow.Context, input *sagemaker.DeleteModelPackageInput) (*sagemaker.DeleteModelPackageOutput, error) {
    var output sagemaker.DeleteModelPackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteModelPackage, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteMonitoringSchedule(ctx workflow.Context, input *sagemaker.DeleteMonitoringScheduleInput) (*sagemaker.DeleteMonitoringScheduleOutput, error) {
    var output sagemaker.DeleteMonitoringScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMonitoringSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteNotebookInstance(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceInput) (*sagemaker.DeleteNotebookInstanceOutput, error) {
    var output sagemaker.DeleteNotebookInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNotebookInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) (*sagemaker.DeleteNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.DeleteNotebookInstanceLifecycleConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNotebookInstanceLifecycleConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteTags(ctx workflow.Context, input *sagemaker.DeleteTagsInput) (*sagemaker.DeleteTagsOutput, error) {
    var output sagemaker.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteTrial(ctx workflow.Context, input *sagemaker.DeleteTrialInput) (*sagemaker.DeleteTrialOutput, error) {
    var output sagemaker.DeleteTrialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrial, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteTrialComponent(ctx workflow.Context, input *sagemaker.DeleteTrialComponentInput) (*sagemaker.DeleteTrialComponentOutput, error) {
    var output sagemaker.DeleteTrialComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrialComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteUserProfile(ctx workflow.Context, input *sagemaker.DeleteUserProfileInput) (*sagemaker.DeleteUserProfileOutput, error) {
    var output sagemaker.DeleteUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteWorkforce(ctx workflow.Context, input *sagemaker.DeleteWorkforceInput) (*sagemaker.DeleteWorkforceOutput, error) {
    var output sagemaker.DeleteWorkforceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkforce, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DeleteWorkteam(ctx workflow.Context, input *sagemaker.DeleteWorkteamInput) (*sagemaker.DeleteWorkteamOutput, error) {
    var output sagemaker.DeleteWorkteamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkteam, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeAlgorithm(ctx workflow.Context, input *sagemaker.DescribeAlgorithmInput) (*sagemaker.DescribeAlgorithmOutput, error) {
    var output sagemaker.DescribeAlgorithmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAlgorithm, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeApp(ctx workflow.Context, input *sagemaker.DescribeAppInput) (*sagemaker.DescribeAppOutput, error) {
    var output sagemaker.DescribeAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeApp, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeAutoMLJob(ctx workflow.Context, input *sagemaker.DescribeAutoMLJobInput) (*sagemaker.DescribeAutoMLJobOutput, error) {
    var output sagemaker.DescribeAutoMLJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAutoMLJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeCodeRepository(ctx workflow.Context, input *sagemaker.DescribeCodeRepositoryInput) (*sagemaker.DescribeCodeRepositoryOutput, error) {
    var output sagemaker.DescribeCodeRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCodeRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeCompilationJob(ctx workflow.Context, input *sagemaker.DescribeCompilationJobInput) (*sagemaker.DescribeCompilationJobOutput, error) {
    var output sagemaker.DescribeCompilationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCompilationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeDomain(ctx workflow.Context, input *sagemaker.DescribeDomainInput) (*sagemaker.DescribeDomainOutput, error) {
    var output sagemaker.DescribeDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeEndpoint(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) (*sagemaker.DescribeEndpointOutput, error) {
    var output sagemaker.DescribeEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeEndpointConfig(ctx workflow.Context, input *sagemaker.DescribeEndpointConfigInput) (*sagemaker.DescribeEndpointConfigOutput, error) {
    var output sagemaker.DescribeEndpointConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpointConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeExperiment(ctx workflow.Context, input *sagemaker.DescribeExperimentInput) (*sagemaker.DescribeExperimentOutput, error) {
    var output sagemaker.DescribeExperimentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExperiment, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeFlowDefinition(ctx workflow.Context, input *sagemaker.DescribeFlowDefinitionInput) (*sagemaker.DescribeFlowDefinitionOutput, error) {
    var output sagemaker.DescribeFlowDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFlowDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeHumanTaskUi(ctx workflow.Context, input *sagemaker.DescribeHumanTaskUiInput) (*sagemaker.DescribeHumanTaskUiOutput, error) {
    var output sagemaker.DescribeHumanTaskUiOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHumanTaskUi, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) (*sagemaker.DescribeHyperParameterTuningJobOutput, error) {
    var output sagemaker.DescribeHyperParameterTuningJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHyperParameterTuningJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeLabelingJob(ctx workflow.Context, input *sagemaker.DescribeLabelingJobInput) (*sagemaker.DescribeLabelingJobOutput, error) {
    var output sagemaker.DescribeLabelingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLabelingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeModel(ctx workflow.Context, input *sagemaker.DescribeModelInput) (*sagemaker.DescribeModelOutput, error) {
    var output sagemaker.DescribeModelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeModel, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeModelPackage(ctx workflow.Context, input *sagemaker.DescribeModelPackageInput) (*sagemaker.DescribeModelPackageOutput, error) {
    var output sagemaker.DescribeModelPackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeModelPackage, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeMonitoringSchedule(ctx workflow.Context, input *sagemaker.DescribeMonitoringScheduleInput) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
    var output sagemaker.DescribeMonitoringScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMonitoringSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeNotebookInstance(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) (*sagemaker.DescribeNotebookInstanceOutput, error) {
    var output sagemaker.DescribeNotebookInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNotebookInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.DescribeNotebookInstanceLifecycleConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNotebookInstanceLifecycleConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeProcessingJob(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) (*sagemaker.DescribeProcessingJobOutput, error) {
    var output sagemaker.DescribeProcessingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProcessingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeSubscribedWorkteam(ctx workflow.Context, input *sagemaker.DescribeSubscribedWorkteamInput) (*sagemaker.DescribeSubscribedWorkteamOutput, error) {
    var output sagemaker.DescribeSubscribedWorkteamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSubscribedWorkteam, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeTrainingJob(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) (*sagemaker.DescribeTrainingJobOutput, error) {
    var output sagemaker.DescribeTrainingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrainingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeTransformJob(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) (*sagemaker.DescribeTransformJobOutput, error) {
    var output sagemaker.DescribeTransformJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTransformJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeTrial(ctx workflow.Context, input *sagemaker.DescribeTrialInput) (*sagemaker.DescribeTrialOutput, error) {
    var output sagemaker.DescribeTrialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrial, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeTrialComponent(ctx workflow.Context, input *sagemaker.DescribeTrialComponentInput) (*sagemaker.DescribeTrialComponentOutput, error) {
    var output sagemaker.DescribeTrialComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrialComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeUserProfile(ctx workflow.Context, input *sagemaker.DescribeUserProfileInput) (*sagemaker.DescribeUserProfileOutput, error) {
    var output sagemaker.DescribeUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeWorkforce(ctx workflow.Context, input *sagemaker.DescribeWorkforceInput) (*sagemaker.DescribeWorkforceOutput, error) {
    var output sagemaker.DescribeWorkforceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkforce, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DescribeWorkteam(ctx workflow.Context, input *sagemaker.DescribeWorkteamInput) (*sagemaker.DescribeWorkteamOutput, error) {
    var output sagemaker.DescribeWorkteamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkteam, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) DisassociateTrialComponent(ctx workflow.Context, input *sagemaker.DisassociateTrialComponentInput) (*sagemaker.DisassociateTrialComponentOutput, error) {
    var output sagemaker.DisassociateTrialComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateTrialComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) GetSearchSuggestions(ctx workflow.Context, input *sagemaker.GetSearchSuggestionsInput) (*sagemaker.GetSearchSuggestionsOutput, error) {
    var output sagemaker.GetSearchSuggestionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSearchSuggestions, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListAlgorithms(ctx workflow.Context, input *sagemaker.ListAlgorithmsInput) (*sagemaker.ListAlgorithmsOutput, error) {
    var output sagemaker.ListAlgorithmsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAlgorithms, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListApps(ctx workflow.Context, input *sagemaker.ListAppsInput) (*sagemaker.ListAppsOutput, error) {
    var output sagemaker.ListAppsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApps, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListAutoMLJobs(ctx workflow.Context, input *sagemaker.ListAutoMLJobsInput) (*sagemaker.ListAutoMLJobsOutput, error) {
    var output sagemaker.ListAutoMLJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAutoMLJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListCandidatesForAutoMLJob(ctx workflow.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) (*sagemaker.ListCandidatesForAutoMLJobOutput, error) {
    var output sagemaker.ListCandidatesForAutoMLJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCandidatesForAutoMLJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListCodeRepositories(ctx workflow.Context, input *sagemaker.ListCodeRepositoriesInput) (*sagemaker.ListCodeRepositoriesOutput, error) {
    var output sagemaker.ListCodeRepositoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCodeRepositories, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListCompilationJobs(ctx workflow.Context, input *sagemaker.ListCompilationJobsInput) (*sagemaker.ListCompilationJobsOutput, error) {
    var output sagemaker.ListCompilationJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCompilationJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListDomains(ctx workflow.Context, input *sagemaker.ListDomainsInput) (*sagemaker.ListDomainsOutput, error) {
    var output sagemaker.ListDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListEndpointConfigs(ctx workflow.Context, input *sagemaker.ListEndpointConfigsInput) (*sagemaker.ListEndpointConfigsOutput, error) {
    var output sagemaker.ListEndpointConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEndpointConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListEndpoints(ctx workflow.Context, input *sagemaker.ListEndpointsInput) (*sagemaker.ListEndpointsOutput, error) {
    var output sagemaker.ListEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListExperiments(ctx workflow.Context, input *sagemaker.ListExperimentsInput) (*sagemaker.ListExperimentsOutput, error) {
    var output sagemaker.ListExperimentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListExperiments, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListFlowDefinitions(ctx workflow.Context, input *sagemaker.ListFlowDefinitionsInput) (*sagemaker.ListFlowDefinitionsOutput, error) {
    var output sagemaker.ListFlowDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFlowDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListHumanTaskUis(ctx workflow.Context, input *sagemaker.ListHumanTaskUisInput) (*sagemaker.ListHumanTaskUisOutput, error) {
    var output sagemaker.ListHumanTaskUisOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHumanTaskUis, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListHyperParameterTuningJobs(ctx workflow.Context, input *sagemaker.ListHyperParameterTuningJobsInput) (*sagemaker.ListHyperParameterTuningJobsOutput, error) {
    var output sagemaker.ListHyperParameterTuningJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHyperParameterTuningJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListLabelingJobs(ctx workflow.Context, input *sagemaker.ListLabelingJobsInput) (*sagemaker.ListLabelingJobsOutput, error) {
    var output sagemaker.ListLabelingJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLabelingJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListLabelingJobsForWorkteam(ctx workflow.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) (*sagemaker.ListLabelingJobsForWorkteamOutput, error) {
    var output sagemaker.ListLabelingJobsForWorkteamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLabelingJobsForWorkteam, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListModelPackages(ctx workflow.Context, input *sagemaker.ListModelPackagesInput) (*sagemaker.ListModelPackagesOutput, error) {
    var output sagemaker.ListModelPackagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListModelPackages, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListModels(ctx workflow.Context, input *sagemaker.ListModelsInput) (*sagemaker.ListModelsOutput, error) {
    var output sagemaker.ListModelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListModels, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListMonitoringExecutions(ctx workflow.Context, input *sagemaker.ListMonitoringExecutionsInput) (*sagemaker.ListMonitoringExecutionsOutput, error) {
    var output sagemaker.ListMonitoringExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMonitoringExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListMonitoringSchedules(ctx workflow.Context, input *sagemaker.ListMonitoringSchedulesInput) (*sagemaker.ListMonitoringSchedulesOutput, error) {
    var output sagemaker.ListMonitoringSchedulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMonitoringSchedules, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListNotebookInstanceLifecycleConfigs(ctx workflow.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error) {
    var output sagemaker.ListNotebookInstanceLifecycleConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNotebookInstanceLifecycleConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListNotebookInstances(ctx workflow.Context, input *sagemaker.ListNotebookInstancesInput) (*sagemaker.ListNotebookInstancesOutput, error) {
    var output sagemaker.ListNotebookInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNotebookInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListProcessingJobs(ctx workflow.Context, input *sagemaker.ListProcessingJobsInput) (*sagemaker.ListProcessingJobsOutput, error) {
    var output sagemaker.ListProcessingJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProcessingJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListSubscribedWorkteams(ctx workflow.Context, input *sagemaker.ListSubscribedWorkteamsInput) (*sagemaker.ListSubscribedWorkteamsOutput, error) {
    var output sagemaker.ListSubscribedWorkteamsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSubscribedWorkteams, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListTags(ctx workflow.Context, input *sagemaker.ListTagsInput) (*sagemaker.ListTagsOutput, error) {
    var output sagemaker.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListTrainingJobs(ctx workflow.Context, input *sagemaker.ListTrainingJobsInput) (*sagemaker.ListTrainingJobsOutput, error) {
    var output sagemaker.ListTrainingJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrainingJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListTrainingJobsForHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error) {
    var output sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrainingJobsForHyperParameterTuningJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListTransformJobs(ctx workflow.Context, input *sagemaker.ListTransformJobsInput) (*sagemaker.ListTransformJobsOutput, error) {
    var output sagemaker.ListTransformJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTransformJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListTrialComponents(ctx workflow.Context, input *sagemaker.ListTrialComponentsInput) (*sagemaker.ListTrialComponentsOutput, error) {
    var output sagemaker.ListTrialComponentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrialComponents, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListTrials(ctx workflow.Context, input *sagemaker.ListTrialsInput) (*sagemaker.ListTrialsOutput, error) {
    var output sagemaker.ListTrialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrials, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListUserProfiles(ctx workflow.Context, input *sagemaker.ListUserProfilesInput) (*sagemaker.ListUserProfilesOutput, error) {
    var output sagemaker.ListUserProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListWorkforces(ctx workflow.Context, input *sagemaker.ListWorkforcesInput) (*sagemaker.ListWorkforcesOutput, error) {
    var output sagemaker.ListWorkforcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorkforces, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) ListWorkteams(ctx workflow.Context, input *sagemaker.ListWorkteamsInput) (*sagemaker.ListWorkteamsOutput, error) {
    var output sagemaker.ListWorkteamsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorkteams, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) RenderUiTemplate(ctx workflow.Context, input *sagemaker.RenderUiTemplateInput) (*sagemaker.RenderUiTemplateOutput, error) {
    var output sagemaker.RenderUiTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RenderUiTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) Search(ctx workflow.Context, input *sagemaker.SearchInput) (*sagemaker.SearchOutput, error) {
    var output sagemaker.SearchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Search, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StartMonitoringSchedule(ctx workflow.Context, input *sagemaker.StartMonitoringScheduleInput) (*sagemaker.StartMonitoringScheduleOutput, error) {
    var output sagemaker.StartMonitoringScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMonitoringSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StartNotebookInstance(ctx workflow.Context, input *sagemaker.StartNotebookInstanceInput) (*sagemaker.StartNotebookInstanceOutput, error) {
    var output sagemaker.StartNotebookInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartNotebookInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopAutoMLJob(ctx workflow.Context, input *sagemaker.StopAutoMLJobInput) (*sagemaker.StopAutoMLJobOutput, error) {
    var output sagemaker.StopAutoMLJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopAutoMLJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopCompilationJob(ctx workflow.Context, input *sagemaker.StopCompilationJobInput) (*sagemaker.StopCompilationJobOutput, error) {
    var output sagemaker.StopCompilationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopCompilationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.StopHyperParameterTuningJobInput) (*sagemaker.StopHyperParameterTuningJobOutput, error) {
    var output sagemaker.StopHyperParameterTuningJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopHyperParameterTuningJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopLabelingJob(ctx workflow.Context, input *sagemaker.StopLabelingJobInput) (*sagemaker.StopLabelingJobOutput, error) {
    var output sagemaker.StopLabelingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopLabelingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopMonitoringSchedule(ctx workflow.Context, input *sagemaker.StopMonitoringScheduleInput) (*sagemaker.StopMonitoringScheduleOutput, error) {
    var output sagemaker.StopMonitoringScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopMonitoringSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopNotebookInstance(ctx workflow.Context, input *sagemaker.StopNotebookInstanceInput) (*sagemaker.StopNotebookInstanceOutput, error) {
    var output sagemaker.StopNotebookInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopNotebookInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopProcessingJob(ctx workflow.Context, input *sagemaker.StopProcessingJobInput) (*sagemaker.StopProcessingJobOutput, error) {
    var output sagemaker.StopProcessingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopProcessingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopTrainingJob(ctx workflow.Context, input *sagemaker.StopTrainingJobInput) (*sagemaker.StopTrainingJobOutput, error) {
    var output sagemaker.StopTrainingJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopTrainingJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) StopTransformJob(ctx workflow.Context, input *sagemaker.StopTransformJobInput) (*sagemaker.StopTransformJobOutput, error) {
    var output sagemaker.StopTransformJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopTransformJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateCodeRepository(ctx workflow.Context, input *sagemaker.UpdateCodeRepositoryInput) (*sagemaker.UpdateCodeRepositoryOutput, error) {
    var output sagemaker.UpdateCodeRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCodeRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateDomain(ctx workflow.Context, input *sagemaker.UpdateDomainInput) (*sagemaker.UpdateDomainOutput, error) {
    var output sagemaker.UpdateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateEndpoint(ctx workflow.Context, input *sagemaker.UpdateEndpointInput) (*sagemaker.UpdateEndpointOutput, error) {
    var output sagemaker.UpdateEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateEndpointWeightsAndCapacities(ctx workflow.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) (*sagemaker.UpdateEndpointWeightsAndCapacitiesOutput, error) {
    var output sagemaker.UpdateEndpointWeightsAndCapacitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEndpointWeightsAndCapacities, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateExperiment(ctx workflow.Context, input *sagemaker.UpdateExperimentInput) (*sagemaker.UpdateExperimentOutput, error) {
    var output sagemaker.UpdateExperimentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateExperiment, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateMonitoringSchedule(ctx workflow.Context, input *sagemaker.UpdateMonitoringScheduleInput) (*sagemaker.UpdateMonitoringScheduleOutput, error) {
    var output sagemaker.UpdateMonitoringScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMonitoringSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateNotebookInstance(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceInput) (*sagemaker.UpdateNotebookInstanceOutput, error) {
    var output sagemaker.UpdateNotebookInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNotebookInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) (*sagemaker.UpdateNotebookInstanceLifecycleConfigOutput, error) {
    var output sagemaker.UpdateNotebookInstanceLifecycleConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNotebookInstanceLifecycleConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateTrial(ctx workflow.Context, input *sagemaker.UpdateTrialInput) (*sagemaker.UpdateTrialOutput, error) {
    var output sagemaker.UpdateTrialOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTrial, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateTrialComponent(ctx workflow.Context, input *sagemaker.UpdateTrialComponentInput) (*sagemaker.UpdateTrialComponentOutput, error) {
    var output sagemaker.UpdateTrialComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTrialComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateUserProfile(ctx workflow.Context, input *sagemaker.UpdateUserProfileInput) (*sagemaker.UpdateUserProfileOutput, error) {
    var output sagemaker.UpdateUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateWorkforce(ctx workflow.Context, input *sagemaker.UpdateWorkforceInput) (*sagemaker.UpdateWorkforceOutput, error) {
    var output sagemaker.UpdateWorkforceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkforce, input).Get(ctx, &output)
    return &output, err
}

func (a *SageMakerStub) UpdateWorkteam(ctx workflow.Context, input *sagemaker.UpdateWorkteamInput) (*sagemaker.UpdateWorkteamOutput, error) {
    var output sagemaker.UpdateWorkteamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkteam, input).Get(ctx, &output)
    return &output, err
}


func (a *SageMakerStub) WaitUntilEndpointDeleted(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error {
    return a.client.WaitUntilEndpointDeleted(input)
}


func (a *SageMakerStub) WaitUntilEndpointInService(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error {
    return a.client.WaitUntilEndpointInService(input)
}


func (a *SageMakerStub) WaitUntilNotebookInstanceDeleted(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error {
    return a.client.WaitUntilNotebookInstanceDeleted(input)
}


func (a *SageMakerStub) WaitUntilNotebookInstanceInService(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error {
    return a.client.WaitUntilNotebookInstanceInService(input)
}


func (a *SageMakerStub) WaitUntilNotebookInstanceStopped(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error {
    return a.client.WaitUntilNotebookInstanceStopped(input)
}


func (a *SageMakerStub) WaitUntilProcessingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) error {
    return a.client.WaitUntilProcessingJobCompletedOrStopped(input)
}


func (a *SageMakerStub) WaitUntilTrainingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) error {
    return a.client.WaitUntilTrainingJobCompletedOrStopped(input)
}


func (a *SageMakerStub) WaitUntilTransformJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) error {
    return a.client.WaitUntilTransformJobCompletedOrStopped(input)
}
