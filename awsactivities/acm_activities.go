package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ACMActivities struct {
	client acmiface.ACMAPI
}

func NewACMActivities(session *session.Session, config ...*aws.Config) *ACMActivities {
	client := acm.New(session, config...)
	return &ACMActivities{client: client}
}

func (a *ACMActivities) AddTagsToCertificate(ctx context.Context, input *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error) {
	return a.client.AddTagsToCertificateWithContext(ctx, input)
}

func (a *ACMActivities) DeleteCertificate(ctx context.Context, input *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error) {
	return a.client.DeleteCertificateWithContext(ctx, input)
}

func (a *ACMActivities) DescribeCertificate(ctx context.Context, input *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error) {
	return a.client.DescribeCertificateWithContext(ctx, input)
}

func (a *ACMActivities) ExportCertificate(ctx context.Context, input *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error) {
	return a.client.ExportCertificateWithContext(ctx, input)
}

func (a *ACMActivities) GetCertificate(ctx context.Context, input *acm.GetCertificateInput) (*acm.GetCertificateOutput, error) {
	return a.client.GetCertificateWithContext(ctx, input)
}

func (a *ACMActivities) ImportCertificate(ctx context.Context, input *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error) {
	return a.client.ImportCertificateWithContext(ctx, input)
}

func (a *ACMActivities) ListCertificates(ctx context.Context, input *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error) {
	return a.client.ListCertificatesWithContext(ctx, input)
}

func (a *ACMActivities) ListTagsForCertificate(ctx context.Context, input *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error) {
	return a.client.ListTagsForCertificateWithContext(ctx, input)
}

func (a *ACMActivities) RemoveTagsFromCertificate(ctx context.Context, input *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error) {
	return a.client.RemoveTagsFromCertificateWithContext(ctx, input)
}

func (a *ACMActivities) RenewCertificate(ctx context.Context, input *acm.RenewCertificateInput) (*acm.RenewCertificateOutput, error) {
	return a.client.RenewCertificateWithContext(ctx, input)
}

func (a *ACMActivities) RequestCertificate(ctx context.Context, input *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error) {
	return a.client.RequestCertificateWithContext(ctx, input)
}

func (a *ACMActivities) ResendValidationEmail(ctx context.Context, input *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error) {
	return a.client.ResendValidationEmailWithContext(ctx, input)
}

func (a *ACMActivities) UpdateCertificateOptions(ctx context.Context, input *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error) {
	return a.client.UpdateCertificateOptionsWithContext(ctx, input)
}

func (a *ACMActivities) WaitUntilCertificateValidated(ctx context.Context, input *acm.DescribeCertificateInput) error {
	return a.client.WaitUntilCertificateValidatedWithContext(ctx, input)

}
