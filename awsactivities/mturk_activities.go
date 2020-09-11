
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/mturk/mturkiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MTurkActivities struct {
    client mturkiface.MTurkAPI
}

func NewMTurkActivities(session *session.Session, config ...*aws.Config) *MTurkActivities {
    client := mturk.New(session, config...)
    return &MTurkActivities{client: client}
}

func (a *MTurkActivities) ApproveAssignment(ctx context.Context, input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error) {
    return a.client.ApproveAssignmentWithContext(ctx, input)
}

func (a *MTurkActivities) AssociateQualificationWithWorker(ctx context.Context, input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error) {
    return a.client.AssociateQualificationWithWorkerWithContext(ctx, input)
}

func (a *MTurkActivities) CreateAdditionalAssignmentsForHIT(ctx context.Context, input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
    return a.client.CreateAdditionalAssignmentsForHITWithContext(ctx, input)
}

func (a *MTurkActivities) CreateHIT(ctx context.Context, input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error) {
    return a.client.CreateHITWithContext(ctx, input)
}

func (a *MTurkActivities) CreateHITType(ctx context.Context, input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error) {
    return a.client.CreateHITTypeWithContext(ctx, input)
}

func (a *MTurkActivities) CreateHITWithHITType(ctx context.Context, input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error) {
    return a.client.CreateHITWithHITTypeWithContext(ctx, input)
}

func (a *MTurkActivities) CreateQualificationType(ctx context.Context, input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error) {
    return a.client.CreateQualificationTypeWithContext(ctx, input)
}

func (a *MTurkActivities) CreateWorkerBlock(ctx context.Context, input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error) {
    return a.client.CreateWorkerBlockWithContext(ctx, input)
}

func (a *MTurkActivities) DeleteHIT(ctx context.Context, input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error) {
    return a.client.DeleteHITWithContext(ctx, input)
}

func (a *MTurkActivities) DeleteQualificationType(ctx context.Context, input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error) {
    return a.client.DeleteQualificationTypeWithContext(ctx, input)
}

func (a *MTurkActivities) DeleteWorkerBlock(ctx context.Context, input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error) {
    return a.client.DeleteWorkerBlockWithContext(ctx, input)
}

func (a *MTurkActivities) DisassociateQualificationFromWorker(ctx context.Context, input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
    return a.client.DisassociateQualificationFromWorkerWithContext(ctx, input)
}

func (a *MTurkActivities) GetAccountBalance(ctx context.Context, input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error) {
    return a.client.GetAccountBalanceWithContext(ctx, input)
}

func (a *MTurkActivities) GetAssignment(ctx context.Context, input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error) {
    return a.client.GetAssignmentWithContext(ctx, input)
}

func (a *MTurkActivities) GetFileUploadURL(ctx context.Context, input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error) {
    return a.client.GetFileUploadURLWithContext(ctx, input)
}

func (a *MTurkActivities) GetHIT(ctx context.Context, input *mturk.GetHITInput) (*mturk.GetHITOutput, error) {
    return a.client.GetHITWithContext(ctx, input)
}

func (a *MTurkActivities) GetQualificationScore(ctx context.Context, input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error) {
    return a.client.GetQualificationScoreWithContext(ctx, input)
}

func (a *MTurkActivities) GetQualificationType(ctx context.Context, input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error) {
    return a.client.GetQualificationTypeWithContext(ctx, input)
}

func (a *MTurkActivities) ListAssignmentsForHIT(ctx context.Context, input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error) {
    return a.client.ListAssignmentsForHITWithContext(ctx, input)
}

func (a *MTurkActivities) ListBonusPayments(ctx context.Context, input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error) {
    return a.client.ListBonusPaymentsWithContext(ctx, input)
}

func (a *MTurkActivities) ListHITs(ctx context.Context, input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error) {
    return a.client.ListHITsWithContext(ctx, input)
}

func (a *MTurkActivities) ListHITsForQualificationType(ctx context.Context, input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error) {
    return a.client.ListHITsForQualificationTypeWithContext(ctx, input)
}

func (a *MTurkActivities) ListQualificationRequests(ctx context.Context, input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error) {
    return a.client.ListQualificationRequestsWithContext(ctx, input)
}

func (a *MTurkActivities) ListQualificationTypes(ctx context.Context, input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error) {
    return a.client.ListQualificationTypesWithContext(ctx, input)
}

func (a *MTurkActivities) ListReviewPolicyResultsForHIT(ctx context.Context, input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
    return a.client.ListReviewPolicyResultsForHITWithContext(ctx, input)
}

func (a *MTurkActivities) ListReviewableHITs(ctx context.Context, input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error) {
    return a.client.ListReviewableHITsWithContext(ctx, input)
}

func (a *MTurkActivities) ListWorkerBlocks(ctx context.Context, input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error) {
    return a.client.ListWorkerBlocksWithContext(ctx, input)
}

func (a *MTurkActivities) ListWorkersWithQualificationType(ctx context.Context, input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
    return a.client.ListWorkersWithQualificationTypeWithContext(ctx, input)
}

func (a *MTurkActivities) NotifyWorkers(ctx context.Context, input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error) {
    return a.client.NotifyWorkersWithContext(ctx, input)
}

func (a *MTurkActivities) RejectAssignment(ctx context.Context, input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error) {
    return a.client.RejectAssignmentWithContext(ctx, input)
}

func (a *MTurkActivities) SendBonus(ctx context.Context, input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error) {
    return a.client.SendBonusWithContext(ctx, input)
}

func (a *MTurkActivities) SendTestEventNotification(ctx context.Context, input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error) {
    return a.client.SendTestEventNotificationWithContext(ctx, input)
}

func (a *MTurkActivities) UpdateExpirationForHIT(ctx context.Context, input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error) {
    return a.client.UpdateExpirationForHITWithContext(ctx, input)
}

func (a *MTurkActivities) UpdateHITReviewStatus(ctx context.Context, input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error) {
    return a.client.UpdateHITReviewStatusWithContext(ctx, input)
}

func (a *MTurkActivities) UpdateHITTypeOfHIT(ctx context.Context, input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error) {
    return a.client.UpdateHITTypeOfHITWithContext(ctx, input)
}

func (a *MTurkActivities) UpdateNotificationSettings(ctx context.Context, input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error) {
    return a.client.UpdateNotificationSettingsWithContext(ctx, input)
}

func (a *MTurkActivities) UpdateQualificationType(ctx context.Context, input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error) {
    return a.client.UpdateQualificationTypeWithContext(ctx, input)
}
