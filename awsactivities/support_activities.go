
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
)

type SupportActivities struct {
	client supportiface.SupportAPI
}

func NewSupportActivities(client supportiface.SupportAPI) *SupportActivities {
    return &SupportActivities{client: client}
}

func (a *SupportActivities) AddAttachmentsToSet(input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
    return a.client.AddAttachmentsToSet(input)
}

func (a *SupportActivities) AddCommunicationToCase(input *support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error) {
    return a.client.AddCommunicationToCase(input)
}

func (a *SupportActivities) CreateCase(input *support.CreateCaseInput) (*support.CreateCaseOutput, error) {
    return a.client.CreateCase(input)
}

func (a *SupportActivities) DescribeAttachment(input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error) {
    return a.client.DescribeAttachment(input)
}

func (a *SupportActivities) DescribeCases(input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
    return a.client.DescribeCases(input)
}

func (a *SupportActivities) DescribeCommunications(input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error) {
    return a.client.DescribeCommunications(input)
}

func (a *SupportActivities) DescribeServices(input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error) {
    return a.client.DescribeServices(input)
}

func (a *SupportActivities) DescribeSeverityLevels(input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error) {
    return a.client.DescribeSeverityLevels(input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckRefreshStatuses(input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
    return a.client.DescribeTrustedAdvisorCheckRefreshStatuses(input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckResult(input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
    return a.client.DescribeTrustedAdvisorCheckResult(input)
}

func (a *SupportActivities) DescribeTrustedAdvisorCheckSummaries(input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
    return a.client.DescribeTrustedAdvisorCheckSummaries(input)
}

func (a *SupportActivities) DescribeTrustedAdvisorChecks(input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error) {
    return a.client.DescribeTrustedAdvisorChecks(input)
}

func (a *SupportActivities) RefreshTrustedAdvisorCheck(input *support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error) {
    return a.client.RefreshTrustedAdvisorCheck(input)
}

func (a *SupportActivities) ResolveCase(input *support.ResolveCaseInput) (*support.ResolveCaseOutput, error) {
    return a.client.ResolveCase(input)
}
