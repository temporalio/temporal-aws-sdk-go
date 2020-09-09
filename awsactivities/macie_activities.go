package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/macie"
	"github.com/aws/aws-sdk-go/service/macie/macieiface"
)

type MacieActivities struct {
	client macieiface.MacieAPI
}

func NewMacieActivities(session *session.Session, config ...*aws.Config) *MacieActivities {
	client := macie.New(session, config...)
	return &MacieActivities{client: client}
}

func (a *MacieActivities) AssociateMemberAccount(input *macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error) {
	return a.client.AssociateMemberAccount(input)
}

func (a *MacieActivities) AssociateS3Resources(input *macie.AssociateS3ResourcesInput) (*macie.AssociateS3ResourcesOutput, error) {
	return a.client.AssociateS3Resources(input)
}

func (a *MacieActivities) DisassociateMemberAccount(input *macie.DisassociateMemberAccountInput) (*macie.DisassociateMemberAccountOutput, error) {
	return a.client.DisassociateMemberAccount(input)
}

func (a *MacieActivities) DisassociateS3Resources(input *macie.DisassociateS3ResourcesInput) (*macie.DisassociateS3ResourcesOutput, error) {
	return a.client.DisassociateS3Resources(input)
}

func (a *MacieActivities) ListMemberAccounts(input *macie.ListMemberAccountsInput) (*macie.ListMemberAccountsOutput, error) {
	return a.client.ListMemberAccounts(input)
}

func (a *MacieActivities) ListS3Resources(input *macie.ListS3ResourcesInput) (*macie.ListS3ResourcesOutput, error) {
	return a.client.ListS3Resources(input)
}

func (a *MacieActivities) UpdateS3Resources(input *macie.UpdateS3ResourcesInput) (*macie.UpdateS3ResourcesOutput, error) {
	return a.client.UpdateS3Resources(input)
}
