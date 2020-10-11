// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package sagemakerstub

import (
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AddTags(ctx workflow.Context, input *sagemaker.AddTagsInput) (*sagemaker.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *sagemaker.AddTagsInput) *SageMakerAddTagsFuture

	AssociateTrialComponent(ctx workflow.Context, input *sagemaker.AssociateTrialComponentInput) (*sagemaker.AssociateTrialComponentOutput, error)
	AssociateTrialComponentAsync(ctx workflow.Context, input *sagemaker.AssociateTrialComponentInput) *SageMakerAssociateTrialComponentFuture

	CreateAlgorithm(ctx workflow.Context, input *sagemaker.CreateAlgorithmInput) (*sagemaker.CreateAlgorithmOutput, error)
	CreateAlgorithmAsync(ctx workflow.Context, input *sagemaker.CreateAlgorithmInput) *SageMakerCreateAlgorithmFuture

	CreateApp(ctx workflow.Context, input *sagemaker.CreateAppInput) (*sagemaker.CreateAppOutput, error)
	CreateAppAsync(ctx workflow.Context, input *sagemaker.CreateAppInput) *SageMakerCreateAppFuture

	CreateAutoMLJob(ctx workflow.Context, input *sagemaker.CreateAutoMLJobInput) (*sagemaker.CreateAutoMLJobOutput, error)
	CreateAutoMLJobAsync(ctx workflow.Context, input *sagemaker.CreateAutoMLJobInput) *SageMakerCreateAutoMLJobFuture

	CreateCodeRepository(ctx workflow.Context, input *sagemaker.CreateCodeRepositoryInput) (*sagemaker.CreateCodeRepositoryOutput, error)
	CreateCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.CreateCodeRepositoryInput) *SageMakerCreateCodeRepositoryFuture

	CreateCompilationJob(ctx workflow.Context, input *sagemaker.CreateCompilationJobInput) (*sagemaker.CreateCompilationJobOutput, error)
	CreateCompilationJobAsync(ctx workflow.Context, input *sagemaker.CreateCompilationJobInput) *SageMakerCreateCompilationJobFuture

	CreateDomain(ctx workflow.Context, input *sagemaker.CreateDomainInput) (*sagemaker.CreateDomainOutput, error)
	CreateDomainAsync(ctx workflow.Context, input *sagemaker.CreateDomainInput) *SageMakerCreateDomainFuture

	CreateEndpoint(ctx workflow.Context, input *sagemaker.CreateEndpointInput) (*sagemaker.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *sagemaker.CreateEndpointInput) *SageMakerCreateEndpointFuture

	CreateEndpointConfig(ctx workflow.Context, input *sagemaker.CreateEndpointConfigInput) (*sagemaker.CreateEndpointConfigOutput, error)
	CreateEndpointConfigAsync(ctx workflow.Context, input *sagemaker.CreateEndpointConfigInput) *SageMakerCreateEndpointConfigFuture

	CreateExperiment(ctx workflow.Context, input *sagemaker.CreateExperimentInput) (*sagemaker.CreateExperimentOutput, error)
	CreateExperimentAsync(ctx workflow.Context, input *sagemaker.CreateExperimentInput) *SageMakerCreateExperimentFuture

	CreateFlowDefinition(ctx workflow.Context, input *sagemaker.CreateFlowDefinitionInput) (*sagemaker.CreateFlowDefinitionOutput, error)
	CreateFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.CreateFlowDefinitionInput) *SageMakerCreateFlowDefinitionFuture

	CreateHumanTaskUi(ctx workflow.Context, input *sagemaker.CreateHumanTaskUiInput) (*sagemaker.CreateHumanTaskUiOutput, error)
	CreateHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.CreateHumanTaskUiInput) *SageMakerCreateHumanTaskUiFuture

	CreateHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.CreateHyperParameterTuningJobInput) (*sagemaker.CreateHyperParameterTuningJobOutput, error)
	CreateHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.CreateHyperParameterTuningJobInput) *SageMakerCreateHyperParameterTuningJobFuture

	CreateLabelingJob(ctx workflow.Context, input *sagemaker.CreateLabelingJobInput) (*sagemaker.CreateLabelingJobOutput, error)
	CreateLabelingJobAsync(ctx workflow.Context, input *sagemaker.CreateLabelingJobInput) *SageMakerCreateLabelingJobFuture

	CreateModel(ctx workflow.Context, input *sagemaker.CreateModelInput) (*sagemaker.CreateModelOutput, error)
	CreateModelAsync(ctx workflow.Context, input *sagemaker.CreateModelInput) *SageMakerCreateModelFuture

	CreateModelPackage(ctx workflow.Context, input *sagemaker.CreateModelPackageInput) (*sagemaker.CreateModelPackageOutput, error)
	CreateModelPackageAsync(ctx workflow.Context, input *sagemaker.CreateModelPackageInput) *SageMakerCreateModelPackageFuture

	CreateMonitoringSchedule(ctx workflow.Context, input *sagemaker.CreateMonitoringScheduleInput) (*sagemaker.CreateMonitoringScheduleOutput, error)
	CreateMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.CreateMonitoringScheduleInput) *SageMakerCreateMonitoringScheduleFuture

	CreateNotebookInstance(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceInput) (*sagemaker.CreateNotebookInstanceOutput, error)
	CreateNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceInput) *SageMakerCreateNotebookInstanceFuture

	CreateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) (*sagemaker.CreateNotebookInstanceLifecycleConfigOutput, error)
	CreateNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.CreateNotebookInstanceLifecycleConfigInput) *SageMakerCreateNotebookInstanceLifecycleConfigFuture

	CreatePresignedDomainUrl(ctx workflow.Context, input *sagemaker.CreatePresignedDomainUrlInput) (*sagemaker.CreatePresignedDomainUrlOutput, error)
	CreatePresignedDomainUrlAsync(ctx workflow.Context, input *sagemaker.CreatePresignedDomainUrlInput) *SageMakerCreatePresignedDomainUrlFuture

	CreatePresignedNotebookInstanceUrl(ctx workflow.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) (*sagemaker.CreatePresignedNotebookInstanceUrlOutput, error)
	CreatePresignedNotebookInstanceUrlAsync(ctx workflow.Context, input *sagemaker.CreatePresignedNotebookInstanceUrlInput) *SageMakerCreatePresignedNotebookInstanceUrlFuture

	CreateProcessingJob(ctx workflow.Context, input *sagemaker.CreateProcessingJobInput) (*sagemaker.CreateProcessingJobOutput, error)
	CreateProcessingJobAsync(ctx workflow.Context, input *sagemaker.CreateProcessingJobInput) *SageMakerCreateProcessingJobFuture

	CreateTrainingJob(ctx workflow.Context, input *sagemaker.CreateTrainingJobInput) (*sagemaker.CreateTrainingJobOutput, error)
	CreateTrainingJobAsync(ctx workflow.Context, input *sagemaker.CreateTrainingJobInput) *SageMakerCreateTrainingJobFuture

	CreateTransformJob(ctx workflow.Context, input *sagemaker.CreateTransformJobInput) (*sagemaker.CreateTransformJobOutput, error)
	CreateTransformJobAsync(ctx workflow.Context, input *sagemaker.CreateTransformJobInput) *SageMakerCreateTransformJobFuture

	CreateTrial(ctx workflow.Context, input *sagemaker.CreateTrialInput) (*sagemaker.CreateTrialOutput, error)
	CreateTrialAsync(ctx workflow.Context, input *sagemaker.CreateTrialInput) *SageMakerCreateTrialFuture

	CreateTrialComponent(ctx workflow.Context, input *sagemaker.CreateTrialComponentInput) (*sagemaker.CreateTrialComponentOutput, error)
	CreateTrialComponentAsync(ctx workflow.Context, input *sagemaker.CreateTrialComponentInput) *SageMakerCreateTrialComponentFuture

	CreateUserProfile(ctx workflow.Context, input *sagemaker.CreateUserProfileInput) (*sagemaker.CreateUserProfileOutput, error)
	CreateUserProfileAsync(ctx workflow.Context, input *sagemaker.CreateUserProfileInput) *SageMakerCreateUserProfileFuture

	CreateWorkforce(ctx workflow.Context, input *sagemaker.CreateWorkforceInput) (*sagemaker.CreateWorkforceOutput, error)
	CreateWorkforceAsync(ctx workflow.Context, input *sagemaker.CreateWorkforceInput) *SageMakerCreateWorkforceFuture

	CreateWorkteam(ctx workflow.Context, input *sagemaker.CreateWorkteamInput) (*sagemaker.CreateWorkteamOutput, error)
	CreateWorkteamAsync(ctx workflow.Context, input *sagemaker.CreateWorkteamInput) *SageMakerCreateWorkteamFuture

	DeleteAlgorithm(ctx workflow.Context, input *sagemaker.DeleteAlgorithmInput) (*sagemaker.DeleteAlgorithmOutput, error)
	DeleteAlgorithmAsync(ctx workflow.Context, input *sagemaker.DeleteAlgorithmInput) *SageMakerDeleteAlgorithmFuture

	DeleteApp(ctx workflow.Context, input *sagemaker.DeleteAppInput) (*sagemaker.DeleteAppOutput, error)
	DeleteAppAsync(ctx workflow.Context, input *sagemaker.DeleteAppInput) *SageMakerDeleteAppFuture

	DeleteCodeRepository(ctx workflow.Context, input *sagemaker.DeleteCodeRepositoryInput) (*sagemaker.DeleteCodeRepositoryOutput, error)
	DeleteCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.DeleteCodeRepositoryInput) *SageMakerDeleteCodeRepositoryFuture

	DeleteDomain(ctx workflow.Context, input *sagemaker.DeleteDomainInput) (*sagemaker.DeleteDomainOutput, error)
	DeleteDomainAsync(ctx workflow.Context, input *sagemaker.DeleteDomainInput) *SageMakerDeleteDomainFuture

	DeleteEndpoint(ctx workflow.Context, input *sagemaker.DeleteEndpointInput) (*sagemaker.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *sagemaker.DeleteEndpointInput) *SageMakerDeleteEndpointFuture

	DeleteEndpointConfig(ctx workflow.Context, input *sagemaker.DeleteEndpointConfigInput) (*sagemaker.DeleteEndpointConfigOutput, error)
	DeleteEndpointConfigAsync(ctx workflow.Context, input *sagemaker.DeleteEndpointConfigInput) *SageMakerDeleteEndpointConfigFuture

	DeleteExperiment(ctx workflow.Context, input *sagemaker.DeleteExperimentInput) (*sagemaker.DeleteExperimentOutput, error)
	DeleteExperimentAsync(ctx workflow.Context, input *sagemaker.DeleteExperimentInput) *SageMakerDeleteExperimentFuture

	DeleteFlowDefinition(ctx workflow.Context, input *sagemaker.DeleteFlowDefinitionInput) (*sagemaker.DeleteFlowDefinitionOutput, error)
	DeleteFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.DeleteFlowDefinitionInput) *SageMakerDeleteFlowDefinitionFuture

	DeleteHumanTaskUi(ctx workflow.Context, input *sagemaker.DeleteHumanTaskUiInput) (*sagemaker.DeleteHumanTaskUiOutput, error)
	DeleteHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.DeleteHumanTaskUiInput) *SageMakerDeleteHumanTaskUiFuture

	DeleteModel(ctx workflow.Context, input *sagemaker.DeleteModelInput) (*sagemaker.DeleteModelOutput, error)
	DeleteModelAsync(ctx workflow.Context, input *sagemaker.DeleteModelInput) *SageMakerDeleteModelFuture

	DeleteModelPackage(ctx workflow.Context, input *sagemaker.DeleteModelPackageInput) (*sagemaker.DeleteModelPackageOutput, error)
	DeleteModelPackageAsync(ctx workflow.Context, input *sagemaker.DeleteModelPackageInput) *SageMakerDeleteModelPackageFuture

	DeleteMonitoringSchedule(ctx workflow.Context, input *sagemaker.DeleteMonitoringScheduleInput) (*sagemaker.DeleteMonitoringScheduleOutput, error)
	DeleteMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.DeleteMonitoringScheduleInput) *SageMakerDeleteMonitoringScheduleFuture

	DeleteNotebookInstance(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceInput) (*sagemaker.DeleteNotebookInstanceOutput, error)
	DeleteNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceInput) *SageMakerDeleteNotebookInstanceFuture

	DeleteNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) (*sagemaker.DeleteNotebookInstanceLifecycleConfigOutput, error)
	DeleteNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.DeleteNotebookInstanceLifecycleConfigInput) *SageMakerDeleteNotebookInstanceLifecycleConfigFuture

	DeleteTags(ctx workflow.Context, input *sagemaker.DeleteTagsInput) (*sagemaker.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *sagemaker.DeleteTagsInput) *SageMakerDeleteTagsFuture

	DeleteTrial(ctx workflow.Context, input *sagemaker.DeleteTrialInput) (*sagemaker.DeleteTrialOutput, error)
	DeleteTrialAsync(ctx workflow.Context, input *sagemaker.DeleteTrialInput) *SageMakerDeleteTrialFuture

	DeleteTrialComponent(ctx workflow.Context, input *sagemaker.DeleteTrialComponentInput) (*sagemaker.DeleteTrialComponentOutput, error)
	DeleteTrialComponentAsync(ctx workflow.Context, input *sagemaker.DeleteTrialComponentInput) *SageMakerDeleteTrialComponentFuture

	DeleteUserProfile(ctx workflow.Context, input *sagemaker.DeleteUserProfileInput) (*sagemaker.DeleteUserProfileOutput, error)
	DeleteUserProfileAsync(ctx workflow.Context, input *sagemaker.DeleteUserProfileInput) *SageMakerDeleteUserProfileFuture

	DeleteWorkforce(ctx workflow.Context, input *sagemaker.DeleteWorkforceInput) (*sagemaker.DeleteWorkforceOutput, error)
	DeleteWorkforceAsync(ctx workflow.Context, input *sagemaker.DeleteWorkforceInput) *SageMakerDeleteWorkforceFuture

	DeleteWorkteam(ctx workflow.Context, input *sagemaker.DeleteWorkteamInput) (*sagemaker.DeleteWorkteamOutput, error)
	DeleteWorkteamAsync(ctx workflow.Context, input *sagemaker.DeleteWorkteamInput) *SageMakerDeleteWorkteamFuture

	DescribeAlgorithm(ctx workflow.Context, input *sagemaker.DescribeAlgorithmInput) (*sagemaker.DescribeAlgorithmOutput, error)
	DescribeAlgorithmAsync(ctx workflow.Context, input *sagemaker.DescribeAlgorithmInput) *SageMakerDescribeAlgorithmFuture

	DescribeApp(ctx workflow.Context, input *sagemaker.DescribeAppInput) (*sagemaker.DescribeAppOutput, error)
	DescribeAppAsync(ctx workflow.Context, input *sagemaker.DescribeAppInput) *SageMakerDescribeAppFuture

	DescribeAutoMLJob(ctx workflow.Context, input *sagemaker.DescribeAutoMLJobInput) (*sagemaker.DescribeAutoMLJobOutput, error)
	DescribeAutoMLJobAsync(ctx workflow.Context, input *sagemaker.DescribeAutoMLJobInput) *SageMakerDescribeAutoMLJobFuture

	DescribeCodeRepository(ctx workflow.Context, input *sagemaker.DescribeCodeRepositoryInput) (*sagemaker.DescribeCodeRepositoryOutput, error)
	DescribeCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.DescribeCodeRepositoryInput) *SageMakerDescribeCodeRepositoryFuture

	DescribeCompilationJob(ctx workflow.Context, input *sagemaker.DescribeCompilationJobInput) (*sagemaker.DescribeCompilationJobOutput, error)
	DescribeCompilationJobAsync(ctx workflow.Context, input *sagemaker.DescribeCompilationJobInput) *SageMakerDescribeCompilationJobFuture

	DescribeDomain(ctx workflow.Context, input *sagemaker.DescribeDomainInput) (*sagemaker.DescribeDomainOutput, error)
	DescribeDomainAsync(ctx workflow.Context, input *sagemaker.DescribeDomainInput) *SageMakerDescribeDomainFuture

	DescribeEndpoint(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) (*sagemaker.DescribeEndpointOutput, error)
	DescribeEndpointAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) *SageMakerDescribeEndpointFuture

	DescribeEndpointConfig(ctx workflow.Context, input *sagemaker.DescribeEndpointConfigInput) (*sagemaker.DescribeEndpointConfigOutput, error)
	DescribeEndpointConfigAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointConfigInput) *SageMakerDescribeEndpointConfigFuture

	DescribeExperiment(ctx workflow.Context, input *sagemaker.DescribeExperimentInput) (*sagemaker.DescribeExperimentOutput, error)
	DescribeExperimentAsync(ctx workflow.Context, input *sagemaker.DescribeExperimentInput) *SageMakerDescribeExperimentFuture

	DescribeFlowDefinition(ctx workflow.Context, input *sagemaker.DescribeFlowDefinitionInput) (*sagemaker.DescribeFlowDefinitionOutput, error)
	DescribeFlowDefinitionAsync(ctx workflow.Context, input *sagemaker.DescribeFlowDefinitionInput) *SageMakerDescribeFlowDefinitionFuture

	DescribeHumanTaskUi(ctx workflow.Context, input *sagemaker.DescribeHumanTaskUiInput) (*sagemaker.DescribeHumanTaskUiOutput, error)
	DescribeHumanTaskUiAsync(ctx workflow.Context, input *sagemaker.DescribeHumanTaskUiInput) *SageMakerDescribeHumanTaskUiFuture

	DescribeHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) (*sagemaker.DescribeHyperParameterTuningJobOutput, error)
	DescribeHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.DescribeHyperParameterTuningJobInput) *SageMakerDescribeHyperParameterTuningJobFuture

	DescribeLabelingJob(ctx workflow.Context, input *sagemaker.DescribeLabelingJobInput) (*sagemaker.DescribeLabelingJobOutput, error)
	DescribeLabelingJobAsync(ctx workflow.Context, input *sagemaker.DescribeLabelingJobInput) *SageMakerDescribeLabelingJobFuture

	DescribeModel(ctx workflow.Context, input *sagemaker.DescribeModelInput) (*sagemaker.DescribeModelOutput, error)
	DescribeModelAsync(ctx workflow.Context, input *sagemaker.DescribeModelInput) *SageMakerDescribeModelFuture

	DescribeModelPackage(ctx workflow.Context, input *sagemaker.DescribeModelPackageInput) (*sagemaker.DescribeModelPackageOutput, error)
	DescribeModelPackageAsync(ctx workflow.Context, input *sagemaker.DescribeModelPackageInput) *SageMakerDescribeModelPackageFuture

	DescribeMonitoringSchedule(ctx workflow.Context, input *sagemaker.DescribeMonitoringScheduleInput) (*sagemaker.DescribeMonitoringScheduleOutput, error)
	DescribeMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.DescribeMonitoringScheduleInput) *SageMakerDescribeMonitoringScheduleFuture

	DescribeNotebookInstance(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) (*sagemaker.DescribeNotebookInstanceOutput, error)
	DescribeNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *SageMakerDescribeNotebookInstanceFuture

	DescribeNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error)
	DescribeNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) *SageMakerDescribeNotebookInstanceLifecycleConfigFuture

	DescribeProcessingJob(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) (*sagemaker.DescribeProcessingJobOutput, error)
	DescribeProcessingJobAsync(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) *SageMakerDescribeProcessingJobFuture

	DescribeSubscribedWorkteam(ctx workflow.Context, input *sagemaker.DescribeSubscribedWorkteamInput) (*sagemaker.DescribeSubscribedWorkteamOutput, error)
	DescribeSubscribedWorkteamAsync(ctx workflow.Context, input *sagemaker.DescribeSubscribedWorkteamInput) *SageMakerDescribeSubscribedWorkteamFuture

	DescribeTrainingJob(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) (*sagemaker.DescribeTrainingJobOutput, error)
	DescribeTrainingJobAsync(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) *SageMakerDescribeTrainingJobFuture

	DescribeTransformJob(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) (*sagemaker.DescribeTransformJobOutput, error)
	DescribeTransformJobAsync(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) *SageMakerDescribeTransformJobFuture

	DescribeTrial(ctx workflow.Context, input *sagemaker.DescribeTrialInput) (*sagemaker.DescribeTrialOutput, error)
	DescribeTrialAsync(ctx workflow.Context, input *sagemaker.DescribeTrialInput) *SageMakerDescribeTrialFuture

	DescribeTrialComponent(ctx workflow.Context, input *sagemaker.DescribeTrialComponentInput) (*sagemaker.DescribeTrialComponentOutput, error)
	DescribeTrialComponentAsync(ctx workflow.Context, input *sagemaker.DescribeTrialComponentInput) *SageMakerDescribeTrialComponentFuture

	DescribeUserProfile(ctx workflow.Context, input *sagemaker.DescribeUserProfileInput) (*sagemaker.DescribeUserProfileOutput, error)
	DescribeUserProfileAsync(ctx workflow.Context, input *sagemaker.DescribeUserProfileInput) *SageMakerDescribeUserProfileFuture

	DescribeWorkforce(ctx workflow.Context, input *sagemaker.DescribeWorkforceInput) (*sagemaker.DescribeWorkforceOutput, error)
	DescribeWorkforceAsync(ctx workflow.Context, input *sagemaker.DescribeWorkforceInput) *SageMakerDescribeWorkforceFuture

	DescribeWorkteam(ctx workflow.Context, input *sagemaker.DescribeWorkteamInput) (*sagemaker.DescribeWorkteamOutput, error)
	DescribeWorkteamAsync(ctx workflow.Context, input *sagemaker.DescribeWorkteamInput) *SageMakerDescribeWorkteamFuture

	DisassociateTrialComponent(ctx workflow.Context, input *sagemaker.DisassociateTrialComponentInput) (*sagemaker.DisassociateTrialComponentOutput, error)
	DisassociateTrialComponentAsync(ctx workflow.Context, input *sagemaker.DisassociateTrialComponentInput) *SageMakerDisassociateTrialComponentFuture

	GetSearchSuggestions(ctx workflow.Context, input *sagemaker.GetSearchSuggestionsInput) (*sagemaker.GetSearchSuggestionsOutput, error)
	GetSearchSuggestionsAsync(ctx workflow.Context, input *sagemaker.GetSearchSuggestionsInput) *SageMakerGetSearchSuggestionsFuture

	ListAlgorithms(ctx workflow.Context, input *sagemaker.ListAlgorithmsInput) (*sagemaker.ListAlgorithmsOutput, error)
	ListAlgorithmsAsync(ctx workflow.Context, input *sagemaker.ListAlgorithmsInput) *SageMakerListAlgorithmsFuture

	ListApps(ctx workflow.Context, input *sagemaker.ListAppsInput) (*sagemaker.ListAppsOutput, error)
	ListAppsAsync(ctx workflow.Context, input *sagemaker.ListAppsInput) *SageMakerListAppsFuture

	ListAutoMLJobs(ctx workflow.Context, input *sagemaker.ListAutoMLJobsInput) (*sagemaker.ListAutoMLJobsOutput, error)
	ListAutoMLJobsAsync(ctx workflow.Context, input *sagemaker.ListAutoMLJobsInput) *SageMakerListAutoMLJobsFuture

	ListCandidatesForAutoMLJob(ctx workflow.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) (*sagemaker.ListCandidatesForAutoMLJobOutput, error)
	ListCandidatesForAutoMLJobAsync(ctx workflow.Context, input *sagemaker.ListCandidatesForAutoMLJobInput) *SageMakerListCandidatesForAutoMLJobFuture

	ListCodeRepositories(ctx workflow.Context, input *sagemaker.ListCodeRepositoriesInput) (*sagemaker.ListCodeRepositoriesOutput, error)
	ListCodeRepositoriesAsync(ctx workflow.Context, input *sagemaker.ListCodeRepositoriesInput) *SageMakerListCodeRepositoriesFuture

	ListCompilationJobs(ctx workflow.Context, input *sagemaker.ListCompilationJobsInput) (*sagemaker.ListCompilationJobsOutput, error)
	ListCompilationJobsAsync(ctx workflow.Context, input *sagemaker.ListCompilationJobsInput) *SageMakerListCompilationJobsFuture

	ListDomains(ctx workflow.Context, input *sagemaker.ListDomainsInput) (*sagemaker.ListDomainsOutput, error)
	ListDomainsAsync(ctx workflow.Context, input *sagemaker.ListDomainsInput) *SageMakerListDomainsFuture

	ListEndpointConfigs(ctx workflow.Context, input *sagemaker.ListEndpointConfigsInput) (*sagemaker.ListEndpointConfigsOutput, error)
	ListEndpointConfigsAsync(ctx workflow.Context, input *sagemaker.ListEndpointConfigsInput) *SageMakerListEndpointConfigsFuture

	ListEndpoints(ctx workflow.Context, input *sagemaker.ListEndpointsInput) (*sagemaker.ListEndpointsOutput, error)
	ListEndpointsAsync(ctx workflow.Context, input *sagemaker.ListEndpointsInput) *SageMakerListEndpointsFuture

	ListExperiments(ctx workflow.Context, input *sagemaker.ListExperimentsInput) (*sagemaker.ListExperimentsOutput, error)
	ListExperimentsAsync(ctx workflow.Context, input *sagemaker.ListExperimentsInput) *SageMakerListExperimentsFuture

	ListFlowDefinitions(ctx workflow.Context, input *sagemaker.ListFlowDefinitionsInput) (*sagemaker.ListFlowDefinitionsOutput, error)
	ListFlowDefinitionsAsync(ctx workflow.Context, input *sagemaker.ListFlowDefinitionsInput) *SageMakerListFlowDefinitionsFuture

	ListHumanTaskUis(ctx workflow.Context, input *sagemaker.ListHumanTaskUisInput) (*sagemaker.ListHumanTaskUisOutput, error)
	ListHumanTaskUisAsync(ctx workflow.Context, input *sagemaker.ListHumanTaskUisInput) *SageMakerListHumanTaskUisFuture

	ListHyperParameterTuningJobs(ctx workflow.Context, input *sagemaker.ListHyperParameterTuningJobsInput) (*sagemaker.ListHyperParameterTuningJobsOutput, error)
	ListHyperParameterTuningJobsAsync(ctx workflow.Context, input *sagemaker.ListHyperParameterTuningJobsInput) *SageMakerListHyperParameterTuningJobsFuture

	ListLabelingJobs(ctx workflow.Context, input *sagemaker.ListLabelingJobsInput) (*sagemaker.ListLabelingJobsOutput, error)
	ListLabelingJobsAsync(ctx workflow.Context, input *sagemaker.ListLabelingJobsInput) *SageMakerListLabelingJobsFuture

	ListLabelingJobsForWorkteam(ctx workflow.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) (*sagemaker.ListLabelingJobsForWorkteamOutput, error)
	ListLabelingJobsForWorkteamAsync(ctx workflow.Context, input *sagemaker.ListLabelingJobsForWorkteamInput) *SageMakerListLabelingJobsForWorkteamFuture

	ListModelPackages(ctx workflow.Context, input *sagemaker.ListModelPackagesInput) (*sagemaker.ListModelPackagesOutput, error)
	ListModelPackagesAsync(ctx workflow.Context, input *sagemaker.ListModelPackagesInput) *SageMakerListModelPackagesFuture

	ListModels(ctx workflow.Context, input *sagemaker.ListModelsInput) (*sagemaker.ListModelsOutput, error)
	ListModelsAsync(ctx workflow.Context, input *sagemaker.ListModelsInput) *SageMakerListModelsFuture

	ListMonitoringExecutions(ctx workflow.Context, input *sagemaker.ListMonitoringExecutionsInput) (*sagemaker.ListMonitoringExecutionsOutput, error)
	ListMonitoringExecutionsAsync(ctx workflow.Context, input *sagemaker.ListMonitoringExecutionsInput) *SageMakerListMonitoringExecutionsFuture

	ListMonitoringSchedules(ctx workflow.Context, input *sagemaker.ListMonitoringSchedulesInput) (*sagemaker.ListMonitoringSchedulesOutput, error)
	ListMonitoringSchedulesAsync(ctx workflow.Context, input *sagemaker.ListMonitoringSchedulesInput) *SageMakerListMonitoringSchedulesFuture

	ListNotebookInstanceLifecycleConfigs(ctx workflow.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error)
	ListNotebookInstanceLifecycleConfigsAsync(ctx workflow.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) *SageMakerListNotebookInstanceLifecycleConfigsFuture

	ListNotebookInstances(ctx workflow.Context, input *sagemaker.ListNotebookInstancesInput) (*sagemaker.ListNotebookInstancesOutput, error)
	ListNotebookInstancesAsync(ctx workflow.Context, input *sagemaker.ListNotebookInstancesInput) *SageMakerListNotebookInstancesFuture

	ListProcessingJobs(ctx workflow.Context, input *sagemaker.ListProcessingJobsInput) (*sagemaker.ListProcessingJobsOutput, error)
	ListProcessingJobsAsync(ctx workflow.Context, input *sagemaker.ListProcessingJobsInput) *SageMakerListProcessingJobsFuture

	ListSubscribedWorkteams(ctx workflow.Context, input *sagemaker.ListSubscribedWorkteamsInput) (*sagemaker.ListSubscribedWorkteamsOutput, error)
	ListSubscribedWorkteamsAsync(ctx workflow.Context, input *sagemaker.ListSubscribedWorkteamsInput) *SageMakerListSubscribedWorkteamsFuture

	ListTags(ctx workflow.Context, input *sagemaker.ListTagsInput) (*sagemaker.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *sagemaker.ListTagsInput) *SageMakerListTagsFuture

	ListTrainingJobs(ctx workflow.Context, input *sagemaker.ListTrainingJobsInput) (*sagemaker.ListTrainingJobsOutput, error)
	ListTrainingJobsAsync(ctx workflow.Context, input *sagemaker.ListTrainingJobsInput) *SageMakerListTrainingJobsFuture

	ListTrainingJobsForHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error)
	ListTrainingJobsForHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) *SageMakerListTrainingJobsForHyperParameterTuningJobFuture

	ListTransformJobs(ctx workflow.Context, input *sagemaker.ListTransformJobsInput) (*sagemaker.ListTransformJobsOutput, error)
	ListTransformJobsAsync(ctx workflow.Context, input *sagemaker.ListTransformJobsInput) *SageMakerListTransformJobsFuture

	ListTrialComponents(ctx workflow.Context, input *sagemaker.ListTrialComponentsInput) (*sagemaker.ListTrialComponentsOutput, error)
	ListTrialComponentsAsync(ctx workflow.Context, input *sagemaker.ListTrialComponentsInput) *SageMakerListTrialComponentsFuture

	ListTrials(ctx workflow.Context, input *sagemaker.ListTrialsInput) (*sagemaker.ListTrialsOutput, error)
	ListTrialsAsync(ctx workflow.Context, input *sagemaker.ListTrialsInput) *SageMakerListTrialsFuture

	ListUserProfiles(ctx workflow.Context, input *sagemaker.ListUserProfilesInput) (*sagemaker.ListUserProfilesOutput, error)
	ListUserProfilesAsync(ctx workflow.Context, input *sagemaker.ListUserProfilesInput) *SageMakerListUserProfilesFuture

	ListWorkforces(ctx workflow.Context, input *sagemaker.ListWorkforcesInput) (*sagemaker.ListWorkforcesOutput, error)
	ListWorkforcesAsync(ctx workflow.Context, input *sagemaker.ListWorkforcesInput) *SageMakerListWorkforcesFuture

	ListWorkteams(ctx workflow.Context, input *sagemaker.ListWorkteamsInput) (*sagemaker.ListWorkteamsOutput, error)
	ListWorkteamsAsync(ctx workflow.Context, input *sagemaker.ListWorkteamsInput) *SageMakerListWorkteamsFuture

	RenderUiTemplate(ctx workflow.Context, input *sagemaker.RenderUiTemplateInput) (*sagemaker.RenderUiTemplateOutput, error)
	RenderUiTemplateAsync(ctx workflow.Context, input *sagemaker.RenderUiTemplateInput) *SageMakerRenderUiTemplateFuture

	Search(ctx workflow.Context, input *sagemaker.SearchInput) (*sagemaker.SearchOutput, error)
	SearchAsync(ctx workflow.Context, input *sagemaker.SearchInput) *SageMakerSearchFuture

	StartMonitoringSchedule(ctx workflow.Context, input *sagemaker.StartMonitoringScheduleInput) (*sagemaker.StartMonitoringScheduleOutput, error)
	StartMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.StartMonitoringScheduleInput) *SageMakerStartMonitoringScheduleFuture

	StartNotebookInstance(ctx workflow.Context, input *sagemaker.StartNotebookInstanceInput) (*sagemaker.StartNotebookInstanceOutput, error)
	StartNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.StartNotebookInstanceInput) *SageMakerStartNotebookInstanceFuture

	StopAutoMLJob(ctx workflow.Context, input *sagemaker.StopAutoMLJobInput) (*sagemaker.StopAutoMLJobOutput, error)
	StopAutoMLJobAsync(ctx workflow.Context, input *sagemaker.StopAutoMLJobInput) *SageMakerStopAutoMLJobFuture

	StopCompilationJob(ctx workflow.Context, input *sagemaker.StopCompilationJobInput) (*sagemaker.StopCompilationJobOutput, error)
	StopCompilationJobAsync(ctx workflow.Context, input *sagemaker.StopCompilationJobInput) *SageMakerStopCompilationJobFuture

	StopHyperParameterTuningJob(ctx workflow.Context, input *sagemaker.StopHyperParameterTuningJobInput) (*sagemaker.StopHyperParameterTuningJobOutput, error)
	StopHyperParameterTuningJobAsync(ctx workflow.Context, input *sagemaker.StopHyperParameterTuningJobInput) *SageMakerStopHyperParameterTuningJobFuture

	StopLabelingJob(ctx workflow.Context, input *sagemaker.StopLabelingJobInput) (*sagemaker.StopLabelingJobOutput, error)
	StopLabelingJobAsync(ctx workflow.Context, input *sagemaker.StopLabelingJobInput) *SageMakerStopLabelingJobFuture

	StopMonitoringSchedule(ctx workflow.Context, input *sagemaker.StopMonitoringScheduleInput) (*sagemaker.StopMonitoringScheduleOutput, error)
	StopMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.StopMonitoringScheduleInput) *SageMakerStopMonitoringScheduleFuture

	StopNotebookInstance(ctx workflow.Context, input *sagemaker.StopNotebookInstanceInput) (*sagemaker.StopNotebookInstanceOutput, error)
	StopNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.StopNotebookInstanceInput) *SageMakerStopNotebookInstanceFuture

	StopProcessingJob(ctx workflow.Context, input *sagemaker.StopProcessingJobInput) (*sagemaker.StopProcessingJobOutput, error)
	StopProcessingJobAsync(ctx workflow.Context, input *sagemaker.StopProcessingJobInput) *SageMakerStopProcessingJobFuture

	StopTrainingJob(ctx workflow.Context, input *sagemaker.StopTrainingJobInput) (*sagemaker.StopTrainingJobOutput, error)
	StopTrainingJobAsync(ctx workflow.Context, input *sagemaker.StopTrainingJobInput) *SageMakerStopTrainingJobFuture

	StopTransformJob(ctx workflow.Context, input *sagemaker.StopTransformJobInput) (*sagemaker.StopTransformJobOutput, error)
	StopTransformJobAsync(ctx workflow.Context, input *sagemaker.StopTransformJobInput) *SageMakerStopTransformJobFuture

	UpdateCodeRepository(ctx workflow.Context, input *sagemaker.UpdateCodeRepositoryInput) (*sagemaker.UpdateCodeRepositoryOutput, error)
	UpdateCodeRepositoryAsync(ctx workflow.Context, input *sagemaker.UpdateCodeRepositoryInput) *SageMakerUpdateCodeRepositoryFuture

	UpdateDomain(ctx workflow.Context, input *sagemaker.UpdateDomainInput) (*sagemaker.UpdateDomainOutput, error)
	UpdateDomainAsync(ctx workflow.Context, input *sagemaker.UpdateDomainInput) *SageMakerUpdateDomainFuture

	UpdateEndpoint(ctx workflow.Context, input *sagemaker.UpdateEndpointInput) (*sagemaker.UpdateEndpointOutput, error)
	UpdateEndpointAsync(ctx workflow.Context, input *sagemaker.UpdateEndpointInput) *SageMakerUpdateEndpointFuture

	UpdateEndpointWeightsAndCapacities(ctx workflow.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) (*sagemaker.UpdateEndpointWeightsAndCapacitiesOutput, error)
	UpdateEndpointWeightsAndCapacitiesAsync(ctx workflow.Context, input *sagemaker.UpdateEndpointWeightsAndCapacitiesInput) *SageMakerUpdateEndpointWeightsAndCapacitiesFuture

	UpdateExperiment(ctx workflow.Context, input *sagemaker.UpdateExperimentInput) (*sagemaker.UpdateExperimentOutput, error)
	UpdateExperimentAsync(ctx workflow.Context, input *sagemaker.UpdateExperimentInput) *SageMakerUpdateExperimentFuture

	UpdateMonitoringSchedule(ctx workflow.Context, input *sagemaker.UpdateMonitoringScheduleInput) (*sagemaker.UpdateMonitoringScheduleOutput, error)
	UpdateMonitoringScheduleAsync(ctx workflow.Context, input *sagemaker.UpdateMonitoringScheduleInput) *SageMakerUpdateMonitoringScheduleFuture

	UpdateNotebookInstance(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceInput) (*sagemaker.UpdateNotebookInstanceOutput, error)
	UpdateNotebookInstanceAsync(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceInput) *SageMakerUpdateNotebookInstanceFuture

	UpdateNotebookInstanceLifecycleConfig(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) (*sagemaker.UpdateNotebookInstanceLifecycleConfigOutput, error)
	UpdateNotebookInstanceLifecycleConfigAsync(ctx workflow.Context, input *sagemaker.UpdateNotebookInstanceLifecycleConfigInput) *SageMakerUpdateNotebookInstanceLifecycleConfigFuture

	UpdateTrial(ctx workflow.Context, input *sagemaker.UpdateTrialInput) (*sagemaker.UpdateTrialOutput, error)
	UpdateTrialAsync(ctx workflow.Context, input *sagemaker.UpdateTrialInput) *SageMakerUpdateTrialFuture

	UpdateTrialComponent(ctx workflow.Context, input *sagemaker.UpdateTrialComponentInput) (*sagemaker.UpdateTrialComponentOutput, error)
	UpdateTrialComponentAsync(ctx workflow.Context, input *sagemaker.UpdateTrialComponentInput) *SageMakerUpdateTrialComponentFuture

	UpdateUserProfile(ctx workflow.Context, input *sagemaker.UpdateUserProfileInput) (*sagemaker.UpdateUserProfileOutput, error)
	UpdateUserProfileAsync(ctx workflow.Context, input *sagemaker.UpdateUserProfileInput) *SageMakerUpdateUserProfileFuture

	UpdateWorkforce(ctx workflow.Context, input *sagemaker.UpdateWorkforceInput) (*sagemaker.UpdateWorkforceOutput, error)
	UpdateWorkforceAsync(ctx workflow.Context, input *sagemaker.UpdateWorkforceInput) *SageMakerUpdateWorkforceFuture

	UpdateWorkteam(ctx workflow.Context, input *sagemaker.UpdateWorkteamInput) (*sagemaker.UpdateWorkteamOutput, error)
	UpdateWorkteamAsync(ctx workflow.Context, input *sagemaker.UpdateWorkteamInput) *SageMakerUpdateWorkteamFuture

	WaitUntilEndpointDeleted(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error
	WaitUntilEndpointDeletedAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) *awsclients.VoidFuture

	WaitUntilEndpointInService(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) error
	WaitUntilEndpointInServiceAsync(ctx workflow.Context, input *sagemaker.DescribeEndpointInput) *awsclients.VoidFuture

	WaitUntilNotebookInstanceDeleted(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
	WaitUntilNotebookInstanceDeletedAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *awsclients.VoidFuture

	WaitUntilNotebookInstanceInService(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
	WaitUntilNotebookInstanceInServiceAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *awsclients.VoidFuture

	WaitUntilNotebookInstanceStopped(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) error
	WaitUntilNotebookInstanceStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeNotebookInstanceInput) *awsclients.VoidFuture

	WaitUntilProcessingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) error
	WaitUntilProcessingJobCompletedOrStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeProcessingJobInput) *awsclients.VoidFuture

	WaitUntilTrainingJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) error
	WaitUntilTrainingJobCompletedOrStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeTrainingJobInput) *awsclients.VoidFuture

	WaitUntilTransformJobCompletedOrStopped(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) error
	WaitUntilTransformJobCompletedOrStoppedAsync(ctx workflow.Context, input *sagemaker.DescribeTransformJobInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
