
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/macie"
	"github.com/aws/aws-sdk-go/service/macie/macieiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MacieActivities struct {
    client macieiface.MacieAPI
}

func NewMacieActivities(session *session.Session, config ...*aws.Config) *MacieActivities {
    client := macie.New(session, config...)
    return &MacieActivities{client: client}
}

func (a *MacieActivities) AssociateMemberAccount(ctx context.Context, input *macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error) {
    return a.client.AssociateMemberAccountWithContext(ctx, input)
}

func (a *MacieActivities) AssociateS3Resources(ctx context.Context, input *macie.AssociateS3ResourcesInput) (*macie.AssociateS3ResourcesOutput, error) {
    return a.client.AssociateS3ResourcesWithContext(ctx, input)
}

func (a *MacieActivities) DisassociateMemberAccount(ctx context.Context, input *macie.DisassociateMemberAccountInput) (*macie.DisassociateMemberAccountOutput, error) {
    return a.client.DisassociateMemberAccountWithContext(ctx, input)
}

func (a *MacieActivities) DisassociateS3Resources(ctx context.Context, input *macie.DisassociateS3ResourcesInput) (*macie.DisassociateS3ResourcesOutput, error) {
    return a.client.DisassociateS3ResourcesWithContext(ctx, input)
}

func (a *MacieActivities) ListMemberAccounts(ctx context.Context, input *macie.ListMemberAccountsInput) (*macie.ListMemberAccountsOutput, error) {
    return a.client.ListMemberAccountsWithContext(ctx, input)
}

func (a *MacieActivities) ListS3Resources(ctx context.Context, input *macie.ListS3ResourcesInput) (*macie.ListS3ResourcesOutput, error) {
    return a.client.ListS3ResourcesWithContext(ctx, input)
}

func (a *MacieActivities) UpdateS3Resources(ctx context.Context, input *macie.UpdateS3ResourcesInput) (*macie.UpdateS3ResourcesOutput, error) {
    return a.client.UpdateS3ResourcesWithContext(ctx, input)
}
