package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/sms/smsiface"
)

type SMSActivities struct {
	client smsiface.SMSAPI
}

func NewSMSActivities(client smsiface.SMSAPI) *SMSActivities {
    return &SMSActivities{client: client}
}


func (a *SMSActivities) CreateApp(input *sms.CreateAppInput) (*sms.CreateAppOutput, error) {
    return a.client.CreateApp(input)
}



func (a *SMSActivities) CreateReplicationJob(input *sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error) {
    return a.client.CreateReplicationJob(input)
}



func (a *SMSActivities) DeleteApp(input *sms.DeleteAppInput) (*sms.DeleteAppOutput, error) {
    return a.client.DeleteApp(input)
}



func (a *SMSActivities) DeleteAppLaunchConfiguration(input *sms.DeleteAppLaunchConfigurationInput) (*sms.DeleteAppLaunchConfigurationOutput, error) {
    return a.client.DeleteAppLaunchConfiguration(input)
}



func (a *SMSActivities) DeleteAppReplicationConfiguration(input *sms.DeleteAppReplicationConfigurationInput) (*sms.DeleteAppReplicationConfigurationOutput, error) {
    return a.client.DeleteAppReplicationConfiguration(input)
}



func (a *SMSActivities) DeleteAppValidationConfiguration(input *sms.DeleteAppValidationConfigurationInput) (*sms.DeleteAppValidationConfigurationOutput, error) {
    return a.client.DeleteAppValidationConfiguration(input)
}



func (a *SMSActivities) DeleteReplicationJob(input *sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error) {
    return a.client.DeleteReplicationJob(input)
}



func (a *SMSActivities) DeleteServerCatalog(input *sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error) {
    return a.client.DeleteServerCatalog(input)
}



func (a *SMSActivities) DisassociateConnector(input *sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error) {
    return a.client.DisassociateConnector(input)
}



func (a *SMSActivities) GenerateChangeSet(input *sms.GenerateChangeSetInput) (*sms.GenerateChangeSetOutput, error) {
    return a.client.GenerateChangeSet(input)
}



func (a *SMSActivities) GenerateTemplate(input *sms.GenerateTemplateInput) (*sms.GenerateTemplateOutput, error) {
    return a.client.GenerateTemplate(input)
}



func (a *SMSActivities) GetApp(input *sms.GetAppInput) (*sms.GetAppOutput, error) {
    return a.client.GetApp(input)
}



func (a *SMSActivities) GetAppLaunchConfiguration(input *sms.GetAppLaunchConfigurationInput) (*sms.GetAppLaunchConfigurationOutput, error) {
    return a.client.GetAppLaunchConfiguration(input)
}



func (a *SMSActivities) GetAppReplicationConfiguration(input *sms.GetAppReplicationConfigurationInput) (*sms.GetAppReplicationConfigurationOutput, error) {
    return a.client.GetAppReplicationConfiguration(input)
}



func (a *SMSActivities) GetAppValidationConfiguration(input *sms.GetAppValidationConfigurationInput) (*sms.GetAppValidationConfigurationOutput, error) {
    return a.client.GetAppValidationConfiguration(input)
}



func (a *SMSActivities) GetAppValidationOutput(input *sms.GetAppValidationOutputInput) (*sms.GetAppValidationOutputOutput, error) {
    return a.client.GetAppValidationOutput(input)
}



func (a *SMSActivities) GetConnectors(input *sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error) {
    return a.client.GetConnectors(input)
}



func (a *SMSActivities) GetReplicationJobs(input *sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error) {
    return a.client.GetReplicationJobs(input)
}



func (a *SMSActivities) GetReplicationRuns(input *sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error) {
    return a.client.GetReplicationRuns(input)
}



func (a *SMSActivities) GetServers(input *sms.GetServersInput) (*sms.GetServersOutput, error) {
    return a.client.GetServers(input)
}



func (a *SMSActivities) ImportAppCatalog(input *sms.ImportAppCatalogInput) (*sms.ImportAppCatalogOutput, error) {
    return a.client.ImportAppCatalog(input)
}



func (a *SMSActivities) ImportServerCatalog(input *sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error) {
    return a.client.ImportServerCatalog(input)
}



func (a *SMSActivities) LaunchApp(input *sms.LaunchAppInput) (*sms.LaunchAppOutput, error) {
    return a.client.LaunchApp(input)
}



func (a *SMSActivities) ListApps(input *sms.ListAppsInput) (*sms.ListAppsOutput, error) {
    return a.client.ListApps(input)
}



func (a *SMSActivities) NotifyAppValidationOutput(input *sms.NotifyAppValidationOutputInput) (*sms.NotifyAppValidationOutputOutput, error) {
    return a.client.NotifyAppValidationOutput(input)
}



func (a *SMSActivities) PutAppLaunchConfiguration(input *sms.PutAppLaunchConfigurationInput) (*sms.PutAppLaunchConfigurationOutput, error) {
    return a.client.PutAppLaunchConfiguration(input)
}



func (a *SMSActivities) PutAppReplicationConfiguration(input *sms.PutAppReplicationConfigurationInput) (*sms.PutAppReplicationConfigurationOutput, error) {
    return a.client.PutAppReplicationConfiguration(input)
}



func (a *SMSActivities) PutAppValidationConfiguration(input *sms.PutAppValidationConfigurationInput) (*sms.PutAppValidationConfigurationOutput, error) {
    return a.client.PutAppValidationConfiguration(input)
}



func (a *SMSActivities) StartAppReplication(input *sms.StartAppReplicationInput) (*sms.StartAppReplicationOutput, error) {
    return a.client.StartAppReplication(input)
}



func (a *SMSActivities) StartOnDemandAppReplication(input *sms.StartOnDemandAppReplicationInput) (*sms.StartOnDemandAppReplicationOutput, error) {
    return a.client.StartOnDemandAppReplication(input)
}



func (a *SMSActivities) StartOnDemandReplicationRun(input *sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error) {
    return a.client.StartOnDemandReplicationRun(input)
}



func (a *SMSActivities) StopAppReplication(input *sms.StopAppReplicationInput) (*sms.StopAppReplicationOutput, error) {
    return a.client.StopAppReplication(input)
}



func (a *SMSActivities) TerminateApp(input *sms.TerminateAppInput) (*sms.TerminateAppOutput, error) {
    return a.client.TerminateApp(input)
}



func (a *SMSActivities) UpdateApp(input *sms.UpdateAppInput) (*sms.UpdateAppOutput, error) {
    return a.client.UpdateApp(input)
}



func (a *SMSActivities) UpdateReplicationJob(input *sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error) {
    return a.client.UpdateReplicationJob(input)
}

