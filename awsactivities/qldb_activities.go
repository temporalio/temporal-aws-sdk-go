
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/qldb/qldbiface"
)

type QLDBActivities struct {
	client qldbiface.QLDBAPI
}

func NewQLDBActivities(client qldbiface.QLDBAPI) *QLDBActivities {
    return &QLDBActivities{client: client}
}

func (a *QLDBActivities) CancelJournalKinesisStream(input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error) {
    return a.client.CancelJournalKinesisStream(input)
}

func (a *QLDBActivities) CreateLedger(input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error) {
    return a.client.CreateLedger(input)
}

func (a *QLDBActivities) DeleteLedger(input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error) {
    return a.client.DeleteLedger(input)
}

func (a *QLDBActivities) DescribeJournalKinesisStream(input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error) {
    return a.client.DescribeJournalKinesisStream(input)
}

func (a *QLDBActivities) DescribeJournalS3Export(input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error) {
    return a.client.DescribeJournalS3Export(input)
}

func (a *QLDBActivities) DescribeLedger(input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error) {
    return a.client.DescribeLedger(input)
}

func (a *QLDBActivities) ExportJournalToS3(input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error) {
    return a.client.ExportJournalToS3(input)
}

func (a *QLDBActivities) GetBlock(input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error) {
    return a.client.GetBlock(input)
}

func (a *QLDBActivities) GetDigest(input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error) {
    return a.client.GetDigest(input)
}

func (a *QLDBActivities) GetRevision(input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error) {
    return a.client.GetRevision(input)
}

func (a *QLDBActivities) ListJournalKinesisStreamsForLedger(input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
    return a.client.ListJournalKinesisStreamsForLedger(input)
}

func (a *QLDBActivities) ListJournalS3Exports(input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error) {
    return a.client.ListJournalS3Exports(input)
}

func (a *QLDBActivities) ListJournalS3ExportsForLedger(input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
    return a.client.ListJournalS3ExportsForLedger(input)
}

func (a *QLDBActivities) ListLedgers(input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error) {
    return a.client.ListLedgers(input)
}

func (a *QLDBActivities) ListTagsForResource(input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *QLDBActivities) StreamJournalToKinesis(input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error) {
    return a.client.StreamJournalToKinesis(input)
}

func (a *QLDBActivities) TagResource(input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *QLDBActivities) UntagResource(input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *QLDBActivities) UpdateLedger(input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error) {
    return a.client.UpdateLedger(input)
}
