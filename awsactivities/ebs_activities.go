
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ebs"
	"github.com/aws/aws-sdk-go/service/ebs/ebsiface"
)

type EBSActivities struct {
	client ebsiface.EBSAPI
}

func NewEBSActivities(client ebsiface.EBSAPI) *EBSActivities {
    return &EBSActivities{client: client}
}

func (a *EBSActivities) CompleteSnapshot(input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error) {
    return a.client.CompleteSnapshot(input)
}

func (a *EBSActivities) GetSnapshotBlock(input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error) {
    return a.client.GetSnapshotBlock(input)
}

func (a *EBSActivities) ListChangedBlocks(input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error) {
    return a.client.ListChangedBlocks(input)
}

func (a *EBSActivities) ListSnapshotBlocks(input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error) {
    return a.client.ListSnapshotBlocks(input)
}

func (a *EBSActivities) PutSnapshotBlock(input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error) {
    return a.client.PutSnapshotBlock(input)
}

func (a *EBSActivities) StartSnapshot(input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error) {
    return a.client.StartSnapshot(input)
}
