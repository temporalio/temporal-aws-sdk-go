package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/robomaker"
	"github.com/aws/aws-sdk-go/service/robomaker/robomakeriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type RoboMakerActivities struct {
	client robomakeriface.RoboMakerAPI
}

func NewRoboMakerActivities(session *session.Session, config ...*aws.Config) *RoboMakerActivities {
	client := robomaker.New(session, config...)
	return &RoboMakerActivities{client: client}
}

func (a *RoboMakerActivities) BatchDeleteWorlds(ctx context.Context, input *robomaker.BatchDeleteWorldsInput) (*robomaker.BatchDeleteWorldsOutput, error) {
	return a.client.BatchDeleteWorldsWithContext(ctx, input)
}

func (a *RoboMakerActivities) BatchDescribeSimulationJob(ctx context.Context, input *robomaker.BatchDescribeSimulationJobInput) (*robomaker.BatchDescribeSimulationJobOutput, error) {
	return a.client.BatchDescribeSimulationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CancelDeploymentJob(ctx context.Context, input *robomaker.CancelDeploymentJobInput) (*robomaker.CancelDeploymentJobOutput, error) {
	return a.client.CancelDeploymentJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CancelSimulationJob(ctx context.Context, input *robomaker.CancelSimulationJobInput) (*robomaker.CancelSimulationJobOutput, error) {
	return a.client.CancelSimulationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CancelSimulationJobBatch(ctx context.Context, input *robomaker.CancelSimulationJobBatchInput) (*robomaker.CancelSimulationJobBatchOutput, error) {
	return a.client.CancelSimulationJobBatchWithContext(ctx, input)
}

func (a *RoboMakerActivities) CancelWorldExportJob(ctx context.Context, input *robomaker.CancelWorldExportJobInput) (*robomaker.CancelWorldExportJobOutput, error) {
	return a.client.CancelWorldExportJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CancelWorldGenerationJob(ctx context.Context, input *robomaker.CancelWorldGenerationJobInput) (*robomaker.CancelWorldGenerationJobOutput, error) {
	return a.client.CancelWorldGenerationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateDeploymentJob(ctx context.Context, input *robomaker.CreateDeploymentJobInput) (*robomaker.CreateDeploymentJobOutput, error) {
	return a.client.CreateDeploymentJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateFleet(ctx context.Context, input *robomaker.CreateFleetInput) (*robomaker.CreateFleetOutput, error) {
	return a.client.CreateFleetWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateRobot(ctx context.Context, input *robomaker.CreateRobotInput) (*robomaker.CreateRobotOutput, error) {
	return a.client.CreateRobotWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateRobotApplication(ctx context.Context, input *robomaker.CreateRobotApplicationInput) (*robomaker.CreateRobotApplicationOutput, error) {
	return a.client.CreateRobotApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateRobotApplicationVersion(ctx context.Context, input *robomaker.CreateRobotApplicationVersionInput) (*robomaker.CreateRobotApplicationVersionOutput, error) {
	return a.client.CreateRobotApplicationVersionWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateSimulationApplication(ctx context.Context, input *robomaker.CreateSimulationApplicationInput) (*robomaker.CreateSimulationApplicationOutput, error) {
	return a.client.CreateSimulationApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateSimulationApplicationVersion(ctx context.Context, input *robomaker.CreateSimulationApplicationVersionInput) (*robomaker.CreateSimulationApplicationVersionOutput, error) {
	return a.client.CreateSimulationApplicationVersionWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateSimulationJob(ctx context.Context, input *robomaker.CreateSimulationJobInput) (*robomaker.CreateSimulationJobOutput, error) {
	return a.client.CreateSimulationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateWorldExportJob(ctx context.Context, input *robomaker.CreateWorldExportJobInput) (*robomaker.CreateWorldExportJobOutput, error) {
	return a.client.CreateWorldExportJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateWorldGenerationJob(ctx context.Context, input *robomaker.CreateWorldGenerationJobInput) (*robomaker.CreateWorldGenerationJobOutput, error) {
	return a.client.CreateWorldGenerationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) CreateWorldTemplate(ctx context.Context, input *robomaker.CreateWorldTemplateInput) (*robomaker.CreateWorldTemplateOutput, error) {
	return a.client.CreateWorldTemplateWithContext(ctx, input)
}

func (a *RoboMakerActivities) DeleteFleet(ctx context.Context, input *robomaker.DeleteFleetInput) (*robomaker.DeleteFleetOutput, error) {
	return a.client.DeleteFleetWithContext(ctx, input)
}

func (a *RoboMakerActivities) DeleteRobot(ctx context.Context, input *robomaker.DeleteRobotInput) (*robomaker.DeleteRobotOutput, error) {
	return a.client.DeleteRobotWithContext(ctx, input)
}

func (a *RoboMakerActivities) DeleteRobotApplication(ctx context.Context, input *robomaker.DeleteRobotApplicationInput) (*robomaker.DeleteRobotApplicationOutput, error) {
	return a.client.DeleteRobotApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) DeleteSimulationApplication(ctx context.Context, input *robomaker.DeleteSimulationApplicationInput) (*robomaker.DeleteSimulationApplicationOutput, error) {
	return a.client.DeleteSimulationApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) DeleteWorldTemplate(ctx context.Context, input *robomaker.DeleteWorldTemplateInput) (*robomaker.DeleteWorldTemplateOutput, error) {
	return a.client.DeleteWorldTemplateWithContext(ctx, input)
}

func (a *RoboMakerActivities) DeregisterRobot(ctx context.Context, input *robomaker.DeregisterRobotInput) (*robomaker.DeregisterRobotOutput, error) {
	return a.client.DeregisterRobotWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeDeploymentJob(ctx context.Context, input *robomaker.DescribeDeploymentJobInput) (*robomaker.DescribeDeploymentJobOutput, error) {
	return a.client.DescribeDeploymentJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeFleet(ctx context.Context, input *robomaker.DescribeFleetInput) (*robomaker.DescribeFleetOutput, error) {
	return a.client.DescribeFleetWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeRobot(ctx context.Context, input *robomaker.DescribeRobotInput) (*robomaker.DescribeRobotOutput, error) {
	return a.client.DescribeRobotWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeRobotApplication(ctx context.Context, input *robomaker.DescribeRobotApplicationInput) (*robomaker.DescribeRobotApplicationOutput, error) {
	return a.client.DescribeRobotApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeSimulationApplication(ctx context.Context, input *robomaker.DescribeSimulationApplicationInput) (*robomaker.DescribeSimulationApplicationOutput, error) {
	return a.client.DescribeSimulationApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeSimulationJob(ctx context.Context, input *robomaker.DescribeSimulationJobInput) (*robomaker.DescribeSimulationJobOutput, error) {
	return a.client.DescribeSimulationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeSimulationJobBatch(ctx context.Context, input *robomaker.DescribeSimulationJobBatchInput) (*robomaker.DescribeSimulationJobBatchOutput, error) {
	return a.client.DescribeSimulationJobBatchWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeWorld(ctx context.Context, input *robomaker.DescribeWorldInput) (*robomaker.DescribeWorldOutput, error) {
	return a.client.DescribeWorldWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeWorldExportJob(ctx context.Context, input *robomaker.DescribeWorldExportJobInput) (*robomaker.DescribeWorldExportJobOutput, error) {
	return a.client.DescribeWorldExportJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeWorldGenerationJob(ctx context.Context, input *robomaker.DescribeWorldGenerationJobInput) (*robomaker.DescribeWorldGenerationJobOutput, error) {
	return a.client.DescribeWorldGenerationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) DescribeWorldTemplate(ctx context.Context, input *robomaker.DescribeWorldTemplateInput) (*robomaker.DescribeWorldTemplateOutput, error) {
	return a.client.DescribeWorldTemplateWithContext(ctx, input)
}

func (a *RoboMakerActivities) GetWorldTemplateBody(ctx context.Context, input *robomaker.GetWorldTemplateBodyInput) (*robomaker.GetWorldTemplateBodyOutput, error) {
	return a.client.GetWorldTemplateBodyWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListDeploymentJobs(ctx context.Context, input *robomaker.ListDeploymentJobsInput) (*robomaker.ListDeploymentJobsOutput, error) {
	return a.client.ListDeploymentJobsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListFleets(ctx context.Context, input *robomaker.ListFleetsInput) (*robomaker.ListFleetsOutput, error) {
	return a.client.ListFleetsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListRobotApplications(ctx context.Context, input *robomaker.ListRobotApplicationsInput) (*robomaker.ListRobotApplicationsOutput, error) {
	return a.client.ListRobotApplicationsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListRobots(ctx context.Context, input *robomaker.ListRobotsInput) (*robomaker.ListRobotsOutput, error) {
	return a.client.ListRobotsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListSimulationApplications(ctx context.Context, input *robomaker.ListSimulationApplicationsInput) (*robomaker.ListSimulationApplicationsOutput, error) {
	return a.client.ListSimulationApplicationsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListSimulationJobBatches(ctx context.Context, input *robomaker.ListSimulationJobBatchesInput) (*robomaker.ListSimulationJobBatchesOutput, error) {
	return a.client.ListSimulationJobBatchesWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListSimulationJobs(ctx context.Context, input *robomaker.ListSimulationJobsInput) (*robomaker.ListSimulationJobsOutput, error) {
	return a.client.ListSimulationJobsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListTagsForResource(ctx context.Context, input *robomaker.ListTagsForResourceInput) (*robomaker.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListWorldExportJobs(ctx context.Context, input *robomaker.ListWorldExportJobsInput) (*robomaker.ListWorldExportJobsOutput, error) {
	return a.client.ListWorldExportJobsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListWorldGenerationJobs(ctx context.Context, input *robomaker.ListWorldGenerationJobsInput) (*robomaker.ListWorldGenerationJobsOutput, error) {
	return a.client.ListWorldGenerationJobsWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListWorldTemplates(ctx context.Context, input *robomaker.ListWorldTemplatesInput) (*robomaker.ListWorldTemplatesOutput, error) {
	return a.client.ListWorldTemplatesWithContext(ctx, input)
}

func (a *RoboMakerActivities) ListWorlds(ctx context.Context, input *robomaker.ListWorldsInput) (*robomaker.ListWorldsOutput, error) {
	return a.client.ListWorldsWithContext(ctx, input)
}

func (a *RoboMakerActivities) RegisterRobot(ctx context.Context, input *robomaker.RegisterRobotInput) (*robomaker.RegisterRobotOutput, error) {
	return a.client.RegisterRobotWithContext(ctx, input)
}

func (a *RoboMakerActivities) RestartSimulationJob(ctx context.Context, input *robomaker.RestartSimulationJobInput) (*robomaker.RestartSimulationJobOutput, error) {
	return a.client.RestartSimulationJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) StartSimulationJobBatch(ctx context.Context, input *robomaker.StartSimulationJobBatchInput) (*robomaker.StartSimulationJobBatchOutput, error) {
	return a.client.StartSimulationJobBatchWithContext(ctx, input)
}

func (a *RoboMakerActivities) SyncDeploymentJob(ctx context.Context, input *robomaker.SyncDeploymentJobInput) (*robomaker.SyncDeploymentJobOutput, error) {
	return a.client.SyncDeploymentJobWithContext(ctx, input)
}

func (a *RoboMakerActivities) TagResource(ctx context.Context, input *robomaker.TagResourceInput) (*robomaker.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *RoboMakerActivities) UntagResource(ctx context.Context, input *robomaker.UntagResourceInput) (*robomaker.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *RoboMakerActivities) UpdateRobotApplication(ctx context.Context, input *robomaker.UpdateRobotApplicationInput) (*robomaker.UpdateRobotApplicationOutput, error) {
	return a.client.UpdateRobotApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) UpdateSimulationApplication(ctx context.Context, input *robomaker.UpdateSimulationApplicationInput) (*robomaker.UpdateSimulationApplicationOutput, error) {
	return a.client.UpdateSimulationApplicationWithContext(ctx, input)
}

func (a *RoboMakerActivities) UpdateWorldTemplate(ctx context.Context, input *robomaker.UpdateWorldTemplateInput) (*robomaker.UpdateWorldTemplateOutput, error) {
	return a.client.UpdateWorldTemplateWithContext(ctx, input)
}
