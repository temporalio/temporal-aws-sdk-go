
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface"
)

type OpsWorksCMActivities struct {
    client opsworkscmiface.OpsWorksCMAPI
}

func NewOpsWorksCMActivities(session *session.Session, config ...*aws.Config) *OpsWorksCMActivities {
    client := opsworkscm.New(session, config...)
    return &OpsWorksCMActivities{client: client}
}

func (a *OpsWorksCMActivities) AssociateNode(input *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error) {
    return a.client.AssociateNode(input)
}

func (a *OpsWorksCMActivities) CreateBackup(input *opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error) {
    return a.client.CreateBackup(input)
}

func (a *OpsWorksCMActivities) CreateServer(input *opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error) {
    return a.client.CreateServer(input)
}

func (a *OpsWorksCMActivities) DeleteBackup(input *opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error) {
    return a.client.DeleteBackup(input)
}

func (a *OpsWorksCMActivities) DeleteServer(input *opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error) {
    return a.client.DeleteServer(input)
}

func (a *OpsWorksCMActivities) DescribeAccountAttributes(input *opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributes(input)
}

func (a *OpsWorksCMActivities) DescribeBackups(input *opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error) {
    return a.client.DescribeBackups(input)
}

func (a *OpsWorksCMActivities) DescribeEvents(input *opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}

func (a *OpsWorksCMActivities) DescribeNodeAssociationStatus(input *opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
    return a.client.DescribeNodeAssociationStatus(input)
}

func (a *OpsWorksCMActivities) DescribeServers(input *opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error) {
    return a.client.DescribeServers(input)
}

func (a *OpsWorksCMActivities) DisassociateNode(input *opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error) {
    return a.client.DisassociateNode(input)
}

func (a *OpsWorksCMActivities) ExportServerEngineAttribute(input *opsworkscm.ExportServerEngineAttributeInput) (*opsworkscm.ExportServerEngineAttributeOutput, error) {
    return a.client.ExportServerEngineAttribute(input)
}

func (a *OpsWorksCMActivities) ListTagsForResource(input *opsworkscm.ListTagsForResourceInput) (*opsworkscm.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *OpsWorksCMActivities) RestoreServer(input *opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error) {
    return a.client.RestoreServer(input)
}

func (a *OpsWorksCMActivities) StartMaintenance(input *opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error) {
    return a.client.StartMaintenance(input)
}

func (a *OpsWorksCMActivities) TagResource(input *opsworkscm.TagResourceInput) (*opsworkscm.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *OpsWorksCMActivities) UntagResource(input *opsworkscm.UntagResourceInput) (*opsworkscm.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *OpsWorksCMActivities) UpdateServer(input *opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error) {
    return a.client.UpdateServer(input)
}

func (a *OpsWorksCMActivities) UpdateServerEngineAttributes(input *opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error) {
    return a.client.UpdateServerEngineAttributes(input)
}

func (a *OpsWorksCMActivities) WaitUntilNodeAssociated(input *opsworkscm.DescribeNodeAssociationStatusInput) error {
    return a.client.WaitUntilNodeAssociated(input)
}
