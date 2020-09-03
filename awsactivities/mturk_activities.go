package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/mturk/mturkiface"
)

type MTurkActivities struct {
	client mturkiface.MTurkAPI
}

func NewMTurkActivities(client mturkiface.MTurkAPI) *MTurkActivities {
    return &MTurkActivities{client: client}
}


func (a *MTurkActivities) ApproveAssignment(input *mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error) {
    return a.client.ApproveAssignment(input)
}



func (a *MTurkActivities) AssociateQualificationWithWorker(input *mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error) {
    return a.client.AssociateQualificationWithWorker(input)
}



func (a *MTurkActivities) CreateAdditionalAssignmentsForHIT(input *mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error) {
    return a.client.CreateAdditionalAssignmentsForHIT(input)
}



func (a *MTurkActivities) CreateHIT(input *mturk.CreateHITInput) (*mturk.CreateHITOutput, error) {
    return a.client.CreateHIT(input)
}



func (a *MTurkActivities) CreateHITType(input *mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error) {
    return a.client.CreateHITType(input)
}



func (a *MTurkActivities) CreateHITWithHITType(input *mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error) {
    return a.client.CreateHITWithHITType(input)
}



func (a *MTurkActivities) CreateQualificationType(input *mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error) {
    return a.client.CreateQualificationType(input)
}



func (a *MTurkActivities) CreateWorkerBlock(input *mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error) {
    return a.client.CreateWorkerBlock(input)
}



func (a *MTurkActivities) DeleteHIT(input *mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error) {
    return a.client.DeleteHIT(input)
}



func (a *MTurkActivities) DeleteQualificationType(input *mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error) {
    return a.client.DeleteQualificationType(input)
}



func (a *MTurkActivities) DeleteWorkerBlock(input *mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error) {
    return a.client.DeleteWorkerBlock(input)
}



func (a *MTurkActivities) DisassociateQualificationFromWorker(input *mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error) {
    return a.client.DisassociateQualificationFromWorker(input)
}



func (a *MTurkActivities) GetAccountBalance(input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error) {
    return a.client.GetAccountBalance(input)
}



func (a *MTurkActivities) GetAssignment(input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error) {
    return a.client.GetAssignment(input)
}



func (a *MTurkActivities) GetFileUploadURL(input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error) {
    return a.client.GetFileUploadURL(input)
}



func (a *MTurkActivities) GetHIT(input *mturk.GetHITInput) (*mturk.GetHITOutput, error) {
    return a.client.GetHIT(input)
}



func (a *MTurkActivities) GetQualificationScore(input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error) {
    return a.client.GetQualificationScore(input)
}



func (a *MTurkActivities) GetQualificationType(input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error) {
    return a.client.GetQualificationType(input)
}



func (a *MTurkActivities) ListAssignmentsForHIT(input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error) {
    return a.client.ListAssignmentsForHIT(input)
}



func (a *MTurkActivities) ListBonusPayments(input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error) {
    return a.client.ListBonusPayments(input)
}



func (a *MTurkActivities) ListHITs(input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error) {
    return a.client.ListHITs(input)
}



func (a *MTurkActivities) ListHITsForQualificationType(input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error) {
    return a.client.ListHITsForQualificationType(input)
}



func (a *MTurkActivities) ListQualificationRequests(input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error) {
    return a.client.ListQualificationRequests(input)
}



func (a *MTurkActivities) ListQualificationTypes(input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error) {
    return a.client.ListQualificationTypes(input)
}



func (a *MTurkActivities) ListReviewPolicyResultsForHIT(input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
    return a.client.ListReviewPolicyResultsForHIT(input)
}



func (a *MTurkActivities) ListReviewableHITs(input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error) {
    return a.client.ListReviewableHITs(input)
}



func (a *MTurkActivities) ListWorkerBlocks(input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error) {
    return a.client.ListWorkerBlocks(input)
}



func (a *MTurkActivities) ListWorkersWithQualificationType(input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
    return a.client.ListWorkersWithQualificationType(input)
}



func (a *MTurkActivities) NotifyWorkers(input *mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error) {
    return a.client.NotifyWorkers(input)
}



func (a *MTurkActivities) RejectAssignment(input *mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error) {
    return a.client.RejectAssignment(input)
}



func (a *MTurkActivities) SendBonus(input *mturk.SendBonusInput) (*mturk.SendBonusOutput, error) {
    return a.client.SendBonus(input)
}



func (a *MTurkActivities) SendTestEventNotification(input *mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error) {
    return a.client.SendTestEventNotification(input)
}



func (a *MTurkActivities) UpdateExpirationForHIT(input *mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error) {
    return a.client.UpdateExpirationForHIT(input)
}



func (a *MTurkActivities) UpdateHITReviewStatus(input *mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error) {
    return a.client.UpdateHITReviewStatus(input)
}



func (a *MTurkActivities) UpdateHITTypeOfHIT(input *mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error) {
    return a.client.UpdateHITTypeOfHIT(input)
}



func (a *MTurkActivities) UpdateNotificationSettings(input *mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error) {
    return a.client.UpdateNotificationSettings(input)
}



func (a *MTurkActivities) UpdateQualificationType(input *mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error) {
    return a.client.UpdateQualificationType(input)
}

