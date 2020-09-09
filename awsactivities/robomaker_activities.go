package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/robomaker"
	"github.com/aws/aws-sdk-go/service/robomaker/robomakeriface"
)

type RoboMakerActivities struct {
	client robomakeriface.RoboMakerAPI
}

func NewRoboMakerActivities(session *session.Session, config ...*aws.Config) *RoboMakerActivities {
	client := robomaker.New(session, config...)
	return &RoboMakerActivities{client: client}
}

func (a *RoboMakerActivities) BatchDeleteWorlds(input *robomaker.BatchDeleteWorldsInput) (*robomaker.BatchDeleteWorldsOutput, error) {
	return a.client.BatchDeleteWorlds(input)
}

func (a *RoboMakerActivities) BatchDescribeSimulationJob(input *robomaker.BatchDescribeSimulationJobInput) (*robomaker.BatchDescribeSimulationJobOutput, error) {
	return a.client.BatchDescribeSimulationJob(input)
}

func (a *RoboMakerActivities) CancelDeploymentJob(input *robomaker.CancelDeploymentJobInput) (*robomaker.CancelDeploymentJobOutput, error) {
	return a.client.CancelDeploymentJob(input)
}

func (a *RoboMakerActivities) CancelSimulationJob(input *robomaker.CancelSimulationJobInput) (*robomaker.CancelSimulationJobOutput, error) {
	return a.client.CancelSimulationJob(input)
}

func (a *RoboMakerActivities) CancelSimulationJobBatch(input *robomaker.CancelSimulationJobBatchInput) (*robomaker.CancelSimulationJobBatchOutput, error) {
	return a.client.CancelSimulationJobBatch(input)
}

func (a *RoboMakerActivities) CancelWorldExportJob(input *robomaker.CancelWorldExportJobInput) (*robomaker.CancelWorldExportJobOutput, error) {
	return a.client.CancelWorldExportJob(input)
}

func (a *RoboMakerActivities) CancelWorldGenerationJob(input *robomaker.CancelWorldGenerationJobInput) (*robomaker.CancelWorldGenerationJobOutput, error) {
	return a.client.CancelWorldGenerationJob(input)
}

func (a *RoboMakerActivities) CreateDeploymentJob(input *robomaker.CreateDeploymentJobInput) (*robomaker.CreateDeploymentJobOutput, error) {
	return a.client.CreateDeploymentJob(input)
}

func (a *RoboMakerActivities) CreateFleet(input *robomaker.CreateFleetInput) (*robomaker.CreateFleetOutput, error) {
	return a.client.CreateFleet(input)
}

func (a *RoboMakerActivities) CreateRobot(input *robomaker.CreateRobotInput) (*robomaker.CreateRobotOutput, error) {
	return a.client.CreateRobot(input)
}

func (a *RoboMakerActivities) CreateRobotApplication(input *robomaker.CreateRobotApplicationInput) (*robomaker.CreateRobotApplicationOutput, error) {
	return a.client.CreateRobotApplication(input)
}

func (a *RoboMakerActivities) CreateRobotApplicationVersion(input *robomaker.CreateRobotApplicationVersionInput) (*robomaker.CreateRobotApplicationVersionOutput, error) {
	return a.client.CreateRobotApplicationVersion(input)
}

func (a *RoboMakerActivities) CreateSimulationApplication(input *robomaker.CreateSimulationApplicationInput) (*robomaker.CreateSimulationApplicationOutput, error) {
	return a.client.CreateSimulationApplication(input)
}

func (a *RoboMakerActivities) CreateSimulationApplicationVersion(input *robomaker.CreateSimulationApplicationVersionInput) (*robomaker.CreateSimulationApplicationVersionOutput, error) {
	return a.client.CreateSimulationApplicationVersion(input)
}

func (a *RoboMakerActivities) CreateSimulationJob(input *robomaker.CreateSimulationJobInput) (*robomaker.CreateSimulationJobOutput, error) {
	return a.client.CreateSimulationJob(input)
}

func (a *RoboMakerActivities) CreateWorldExportJob(input *robomaker.CreateWorldExportJobInput) (*robomaker.CreateWorldExportJobOutput, error) {
	return a.client.CreateWorldExportJob(input)
}

func (a *RoboMakerActivities) CreateWorldGenerationJob(input *robomaker.CreateWorldGenerationJobInput) (*robomaker.CreateWorldGenerationJobOutput, error) {
	return a.client.CreateWorldGenerationJob(input)
}

func (a *RoboMakerActivities) CreateWorldTemplate(input *robomaker.CreateWorldTemplateInput) (*robomaker.CreateWorldTemplateOutput, error) {
	return a.client.CreateWorldTemplate(input)
}

func (a *RoboMakerActivities) DeleteFleet(input *robomaker.DeleteFleetInput) (*robomaker.DeleteFleetOutput, error) {
	return a.client.DeleteFleet(input)
}

func (a *RoboMakerActivities) DeleteRobot(input *robomaker.DeleteRobotInput) (*robomaker.DeleteRobotOutput, error) {
	return a.client.DeleteRobot(input)
}

func (a *RoboMakerActivities) DeleteRobotApplication(input *robomaker.DeleteRobotApplicationInput) (*robomaker.DeleteRobotApplicationOutput, error) {
	return a.client.DeleteRobotApplication(input)
}

func (a *RoboMakerActivities) DeleteSimulationApplication(input *robomaker.DeleteSimulationApplicationInput) (*robomaker.DeleteSimulationApplicationOutput, error) {
	return a.client.DeleteSimulationApplication(input)
}

func (a *RoboMakerActivities) DeleteWorldTemplate(input *robomaker.DeleteWorldTemplateInput) (*robomaker.DeleteWorldTemplateOutput, error) {
	return a.client.DeleteWorldTemplate(input)
}

func (a *RoboMakerActivities) DeregisterRobot(input *robomaker.DeregisterRobotInput) (*robomaker.DeregisterRobotOutput, error) {
	return a.client.DeregisterRobot(input)
}

func (a *RoboMakerActivities) DescribeDeploymentJob(input *robomaker.DescribeDeploymentJobInput) (*robomaker.DescribeDeploymentJobOutput, error) {
	return a.client.DescribeDeploymentJob(input)
}

func (a *RoboMakerActivities) DescribeFleet(input *robomaker.DescribeFleetInput) (*robomaker.DescribeFleetOutput, error) {
	return a.client.DescribeFleet(input)
}

func (a *RoboMakerActivities) DescribeRobot(input *robomaker.DescribeRobotInput) (*robomaker.DescribeRobotOutput, error) {
	return a.client.DescribeRobot(input)
}

func (a *RoboMakerActivities) DescribeRobotApplication(input *robomaker.DescribeRobotApplicationInput) (*robomaker.DescribeRobotApplicationOutput, error) {
	return a.client.DescribeRobotApplication(input)
}

func (a *RoboMakerActivities) DescribeSimulationApplication(input *robomaker.DescribeSimulationApplicationInput) (*robomaker.DescribeSimulationApplicationOutput, error) {
	return a.client.DescribeSimulationApplication(input)
}

func (a *RoboMakerActivities) DescribeSimulationJob(input *robomaker.DescribeSimulationJobInput) (*robomaker.DescribeSimulationJobOutput, error) {
	return a.client.DescribeSimulationJob(input)
}

func (a *RoboMakerActivities) DescribeSimulationJobBatch(input *robomaker.DescribeSimulationJobBatchInput) (*robomaker.DescribeSimulationJobBatchOutput, error) {
	return a.client.DescribeSimulationJobBatch(input)
}

func (a *RoboMakerActivities) DescribeWorld(input *robomaker.DescribeWorldInput) (*robomaker.DescribeWorldOutput, error) {
	return a.client.DescribeWorld(input)
}

func (a *RoboMakerActivities) DescribeWorldExportJob(input *robomaker.DescribeWorldExportJobInput) (*robomaker.DescribeWorldExportJobOutput, error) {
	return a.client.DescribeWorldExportJob(input)
}

func (a *RoboMakerActivities) DescribeWorldGenerationJob(input *robomaker.DescribeWorldGenerationJobInput) (*robomaker.DescribeWorldGenerationJobOutput, error) {
	return a.client.DescribeWorldGenerationJob(input)
}

func (a *RoboMakerActivities) DescribeWorldTemplate(input *robomaker.DescribeWorldTemplateInput) (*robomaker.DescribeWorldTemplateOutput, error) {
	return a.client.DescribeWorldTemplate(input)
}

func (a *RoboMakerActivities) GetWorldTemplateBody(input *robomaker.GetWorldTemplateBodyInput) (*robomaker.GetWorldTemplateBodyOutput, error) {
	return a.client.GetWorldTemplateBody(input)
}

func (a *RoboMakerActivities) ListDeploymentJobs(input *robomaker.ListDeploymentJobsInput) (*robomaker.ListDeploymentJobsOutput, error) {
	return a.client.ListDeploymentJobs(input)
}

func (a *RoboMakerActivities) ListFleets(input *robomaker.ListFleetsInput) (*robomaker.ListFleetsOutput, error) {
	return a.client.ListFleets(input)
}

func (a *RoboMakerActivities) ListRobotApplications(input *robomaker.ListRobotApplicationsInput) (*robomaker.ListRobotApplicationsOutput, error) {
	return a.client.ListRobotApplications(input)
}

func (a *RoboMakerActivities) ListRobots(input *robomaker.ListRobotsInput) (*robomaker.ListRobotsOutput, error) {
	return a.client.ListRobots(input)
}

func (a *RoboMakerActivities) ListSimulationApplications(input *robomaker.ListSimulationApplicationsInput) (*robomaker.ListSimulationApplicationsOutput, error) {
	return a.client.ListSimulationApplications(input)
}

func (a *RoboMakerActivities) ListSimulationJobBatches(input *robomaker.ListSimulationJobBatchesInput) (*robomaker.ListSimulationJobBatchesOutput, error) {
	return a.client.ListSimulationJobBatches(input)
}

func (a *RoboMakerActivities) ListSimulationJobs(input *robomaker.ListSimulationJobsInput) (*robomaker.ListSimulationJobsOutput, error) {
	return a.client.ListSimulationJobs(input)
}

func (a *RoboMakerActivities) ListTagsForResource(input *robomaker.ListTagsForResourceInput) (*robomaker.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *RoboMakerActivities) ListWorldExportJobs(input *robomaker.ListWorldExportJobsInput) (*robomaker.ListWorldExportJobsOutput, error) {
	return a.client.ListWorldExportJobs(input)
}

func (a *RoboMakerActivities) ListWorldGenerationJobs(input *robomaker.ListWorldGenerationJobsInput) (*robomaker.ListWorldGenerationJobsOutput, error) {
	return a.client.ListWorldGenerationJobs(input)
}

func (a *RoboMakerActivities) ListWorldTemplates(input *robomaker.ListWorldTemplatesInput) (*robomaker.ListWorldTemplatesOutput, error) {
	return a.client.ListWorldTemplates(input)
}

func (a *RoboMakerActivities) ListWorlds(input *robomaker.ListWorldsInput) (*robomaker.ListWorldsOutput, error) {
	return a.client.ListWorlds(input)
}

func (a *RoboMakerActivities) RegisterRobot(input *robomaker.RegisterRobotInput) (*robomaker.RegisterRobotOutput, error) {
	return a.client.RegisterRobot(input)
}

func (a *RoboMakerActivities) RestartSimulationJob(input *robomaker.RestartSimulationJobInput) (*robomaker.RestartSimulationJobOutput, error) {
	return a.client.RestartSimulationJob(input)
}

func (a *RoboMakerActivities) StartSimulationJobBatch(input *robomaker.StartSimulationJobBatchInput) (*robomaker.StartSimulationJobBatchOutput, error) {
	return a.client.StartSimulationJobBatch(input)
}

func (a *RoboMakerActivities) SyncDeploymentJob(input *robomaker.SyncDeploymentJobInput) (*robomaker.SyncDeploymentJobOutput, error) {
	return a.client.SyncDeploymentJob(input)
}

func (a *RoboMakerActivities) TagResource(input *robomaker.TagResourceInput) (*robomaker.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *RoboMakerActivities) UntagResource(input *robomaker.UntagResourceInput) (*robomaker.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *RoboMakerActivities) UpdateRobotApplication(input *robomaker.UpdateRobotApplicationInput) (*robomaker.UpdateRobotApplicationOutput, error) {
	return a.client.UpdateRobotApplication(input)
}

func (a *RoboMakerActivities) UpdateSimulationApplication(input *robomaker.UpdateSimulationApplicationInput) (*robomaker.UpdateSimulationApplicationOutput, error) {
	return a.client.UpdateSimulationApplication(input)
}

func (a *RoboMakerActivities) UpdateWorldTemplate(input *robomaker.UpdateWorldTemplateInput) (*robomaker.UpdateWorldTemplateOutput, error) {
	return a.client.UpdateWorldTemplate(input)
}
