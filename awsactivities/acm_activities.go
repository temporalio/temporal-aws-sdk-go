
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
)

type ACMActivities struct {
	client acmiface.ACMAPI
}

func NewACMActivities(session *session.Session, config... *aws.Config) *ACMActivities {
    client := acm.New(session, config...)
    return &ACMActivities{client: client}
}

func (a *ACMActivities) AddTagsToCertificate(input *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error) {
    return a.client.AddTagsToCertificate(input)
}

func (a *ACMActivities) DeleteCertificate(input *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error) {
    return a.client.DeleteCertificate(input)
}

func (a *ACMActivities) DescribeCertificate(input *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error) {
    return a.client.DescribeCertificate(input)
}

func (a *ACMActivities) ExportCertificate(input *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error) {
    return a.client.ExportCertificate(input)
}

func (a *ACMActivities) GetCertificate(input *acm.GetCertificateInput) (*acm.GetCertificateOutput, error) {
    return a.client.GetCertificate(input)
}

func (a *ACMActivities) ImportCertificate(input *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error) {
    return a.client.ImportCertificate(input)
}

func (a *ACMActivities) ListCertificates(input *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error) {
    return a.client.ListCertificates(input)
}

func (a *ACMActivities) ListTagsForCertificate(input *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error) {
    return a.client.ListTagsForCertificate(input)
}

func (a *ACMActivities) RemoveTagsFromCertificate(input *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error) {
    return a.client.RemoveTagsFromCertificate(input)
}

func (a *ACMActivities) RenewCertificate(input *acm.RenewCertificateInput) (*acm.RenewCertificateOutput, error) {
    return a.client.RenewCertificate(input)
}

func (a *ACMActivities) RequestCertificate(input *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error) {
    return a.client.RequestCertificate(input)
}

func (a *ACMActivities) ResendValidationEmail(input *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error) {
    return a.client.ResendValidationEmail(input)
}

func (a *ACMActivities) UpdateCertificateOptions(input *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error) {
    return a.client.UpdateCertificateOptions(input)
}

func (a *ACMActivities) WaitUntilCertificateValidated(input *acm.DescribeCertificateInput) error {
    return a.client.WaitUntilCertificateValidated(input)
}
