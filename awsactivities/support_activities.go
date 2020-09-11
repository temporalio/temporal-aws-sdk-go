
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SupportActivities struct {
    client supportiface.SupportAPI
}

func NewSupportActivities(session *session.Session, config ...*aws.Config) *SupportActivities {
    client := support.New(session, config...)
    return &SupportActivities{client: client}
}

func (a *SupportActivities) AddAttachmentsToSet(ctx context.Context, input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
    return a.client.AddAttachmentsToSetWithContext(ctx, input)
}

func (a *SupportActivities) AddCommunicationToCase(ctx context.Context, input *support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error) {
    return a.client.AddCommunicationToCaseWithContext(ctx, input)
}

func (a *SupportActivities) CreateCase(ctx context.Context, input *support.CreateCaseInput) (*support.CreateCaseOutput, error) {
    return a.client.CreateCaseWithContext(ctx, input)
}

func (a *SupportActivities) DescribeAttachment(ctx context.Context, input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error) {
    return a.client.DescribeAttachmentWithContext(ctx, input)
}

func (a *SupportActivities) DescribeCases(ctx context.Context, input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
    return a.client.DescribeCasesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeCommunications(ctx context.Context, input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error) {
    return a.client.DescribeCommunicationsWithContext(ctx, input)
}

func (a *SupportActivities) DescribeServices(ctx context.Context, input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error) {
    return a.client.DescribeServicesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeSeverityLevels(ctx context.Context, input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error) {
    return a.client.DescribeSeverityLevelsWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckRefreshStatuses(ctx context.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
    return a.client.DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckResult(ctx context.Context, input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
    return a.client.DescribeTrustedAdvisorCheckResultWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckSummaries(ctx context.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
    return a.client.DescribeTrustedAdvisorCheckSummariesWithContext(ctx, input)
}

func (a *SupportActivities) DescribeTrustedAdvisorChecks(ctx context.Context, input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error) {
    return a.client.DescribeTrustedAdvisorChecksWithContext(ctx, input)
}

func (a *SupportActivities) RefreshTrustedAdvisorCheck(ctx context.Context, input *support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error) {
    return a.client.RefreshTrustedAdvisorCheckWithContext(ctx, input)
}

func (a *SupportActivities) ResolveCase(ctx context.Context, input *support.ResolveCaseInput) (*support.ResolveCaseOutput, error) {
    return a.client.ResolveCaseWithContext(ctx, input)
}
