package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ebs"
	"github.com/aws/aws-sdk-go/service/ebs/ebsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type EBSActivities struct {
	client ebsiface.EBSAPI
}

func NewEBSActivities(session *session.Session, config ...*aws.Config) *EBSActivities {
	client := ebs.New(session, config...)
	return &EBSActivities{client: client}
}

func (a *EBSActivities) CompleteSnapshot(ctx context.Context, input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error) {
	return a.client.CompleteSnapshotWithContext(ctx, input)
}

func (a *EBSActivities) GetSnapshotBlock(ctx context.Context, input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error) {
	return a.client.GetSnapshotBlockWithContext(ctx, input)
}

func (a *EBSActivities) ListChangedBlocks(ctx context.Context, input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error) {
	return a.client.ListChangedBlocksWithContext(ctx, input)
}

func (a *EBSActivities) ListSnapshotBlocks(ctx context.Context, input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error) {
	return a.client.ListSnapshotBlocksWithContext(ctx, input)
}

func (a *EBSActivities) PutSnapshotBlock(ctx context.Context, input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error) {
	return a.client.PutSnapshotBlockWithContext(ctx, input)
}

func (a *EBSActivities) StartSnapshot(ctx context.Context, input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.StartSnapshotWithContext(ctx, input)
}
