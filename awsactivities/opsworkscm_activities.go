package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type OpsWorksCMActivities struct {
	client opsworkscmiface.OpsWorksCMAPI
}

func NewOpsWorksCMActivities(session *session.Session, config ...*aws.Config) *OpsWorksCMActivities {
	client := opsworkscm.New(session, config...)
	return &OpsWorksCMActivities{client: client}
}

func (a *OpsWorksCMActivities) AssociateNode(ctx context.Context, input *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error) {
	return a.client.AssociateNodeWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) CreateBackup(ctx context.Context, input *opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error) {
	return a.client.CreateBackupWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) CreateServer(ctx context.Context, input *opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error) {
	return a.client.CreateServerWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DeleteBackup(ctx context.Context, input *opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error) {
	return a.client.DeleteBackupWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DeleteServer(ctx context.Context, input *opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error) {
	return a.client.DeleteServerWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DescribeAccountAttributes(ctx context.Context, input *opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error) {
	return a.client.DescribeAccountAttributesWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DescribeBackups(ctx context.Context, input *opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error) {
	return a.client.DescribeBackupsWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DescribeEvents(ctx context.Context, input *opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error) {
	return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DescribeNodeAssociationStatus(ctx context.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error) {
	return a.client.DescribeNodeAssociationStatusWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DescribeServers(ctx context.Context, input *opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error) {
	return a.client.DescribeServersWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) DisassociateNode(ctx context.Context, input *opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error) {
	return a.client.DisassociateNodeWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) ExportServerEngineAttribute(ctx context.Context, input *opsworkscm.ExportServerEngineAttributeInput) (*opsworkscm.ExportServerEngineAttributeOutput, error) {
	return a.client.ExportServerEngineAttributeWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) ListTagsForResource(ctx context.Context, input *opsworkscm.ListTagsForResourceInput) (*opsworkscm.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) RestoreServer(ctx context.Context, input *opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error) {
	return a.client.RestoreServerWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) StartMaintenance(ctx context.Context, input *opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error) {
	return a.client.StartMaintenanceWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) TagResource(ctx context.Context, input *opsworkscm.TagResourceInput) (*opsworkscm.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) UntagResource(ctx context.Context, input *opsworkscm.UntagResourceInput) (*opsworkscm.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) UpdateServer(ctx context.Context, input *opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error) {
	return a.client.UpdateServerWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) UpdateServerEngineAttributes(ctx context.Context, input *opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error) {
	return a.client.UpdateServerEngineAttributesWithContext(ctx, input)
}

func (a *OpsWorksCMActivities) WaitUntilNodeAssociated(ctx context.Context, input *opsworkscm.DescribeNodeAssociationStatusInput) error {
	return a.client.WaitUntilNodeAssociatedWithContext(ctx, input)

}
