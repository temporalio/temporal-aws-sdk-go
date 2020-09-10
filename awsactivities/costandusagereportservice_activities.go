package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CostandUsageReportServiceActivities struct {
	client costandusagereportserviceiface.CostandUsageReportServiceAPI
}

func NewCostandUsageReportServiceActivities(session *session.Session, config ...*aws.Config) *CostandUsageReportServiceActivities {
	client := costandusagereportservice.New(session, config...)
	return &CostandUsageReportServiceActivities{client: client}
}

func (a *CostandUsageReportServiceActivities) DeleteReportDefinition(ctx context.Context, input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	return a.client.DeleteReportDefinitionWithContext(ctx, input)
}

func (a *CostandUsageReportServiceActivities) DescribeReportDefinitions(ctx context.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	return a.client.DescribeReportDefinitionsWithContext(ctx, input)
}

func (a *CostandUsageReportServiceActivities) ModifyReportDefinition(ctx context.Context, input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
	return a.client.ModifyReportDefinitionWithContext(ctx, input)
}

func (a *CostandUsageReportServiceActivities) PutReportDefinition(ctx context.Context, input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	return a.client.PutReportDefinitionWithContext(ctx, input)
}
