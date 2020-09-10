package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/acmpca/acmpcaiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ACMPCAActivities struct {
	client acmpcaiface.ACMPCAAPI
}

func NewACMPCAActivities(session *session.Session, config ...*aws.Config) *ACMPCAActivities {
	client := acmpca.New(session, config...)
	return &ACMPCAActivities{client: client}
}

func (a *ACMPCAActivities) CreateCertificateAuthority(ctx context.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
	return a.client.CreateCertificateAuthorityWithContext(ctx, input)
}

func (a *ACMPCAActivities) CreateCertificateAuthorityAuditReport(ctx context.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
	return a.client.CreateCertificateAuthorityAuditReportWithContext(ctx, input)
}

func (a *ACMPCAActivities) CreatePermission(ctx context.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error) {
	return a.client.CreatePermissionWithContext(ctx, input)
}

func (a *ACMPCAActivities) DeleteCertificateAuthority(ctx context.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error) {
	return a.client.DeleteCertificateAuthorityWithContext(ctx, input)
}

func (a *ACMPCAActivities) DeletePermission(ctx context.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error) {
	return a.client.DeletePermissionWithContext(ctx, input)
}

func (a *ACMPCAActivities) DeletePolicy(ctx context.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error) {
	return a.client.DeletePolicyWithContext(ctx, input)
}

func (a *ACMPCAActivities) DescribeCertificateAuthority(ctx context.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
	return a.client.DescribeCertificateAuthorityWithContext(ctx, input)
}

func (a *ACMPCAActivities) DescribeCertificateAuthorityAuditReport(ctx context.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
	return a.client.DescribeCertificateAuthorityAuditReportWithContext(ctx, input)
}

func (a *ACMPCAActivities) GetCertificate(ctx context.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
	return a.client.GetCertificateWithContext(ctx, input)
}

func (a *ACMPCAActivities) GetCertificateAuthorityCertificate(ctx context.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
	return a.client.GetCertificateAuthorityCertificateWithContext(ctx, input)
}

func (a *ACMPCAActivities) GetCertificateAuthorityCsr(ctx context.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
	return a.client.GetCertificateAuthorityCsrWithContext(ctx, input)
}

func (a *ACMPCAActivities) GetPolicy(ctx context.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
	return a.client.GetPolicyWithContext(ctx, input)
}

func (a *ACMPCAActivities) ImportCertificateAuthorityCertificate(ctx context.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
	return a.client.ImportCertificateAuthorityCertificateWithContext(ctx, input)
}

func (a *ACMPCAActivities) IssueCertificate(ctx context.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error) {
	return a.client.IssueCertificateWithContext(ctx, input)
}

func (a *ACMPCAActivities) ListCertificateAuthorities(ctx context.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
	return a.client.ListCertificateAuthoritiesWithContext(ctx, input)
}

func (a *ACMPCAActivities) ListPermissions(ctx context.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
	return a.client.ListPermissionsWithContext(ctx, input)
}

func (a *ACMPCAActivities) ListTags(ctx context.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *ACMPCAActivities) PutPolicy(ctx context.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error) {
	return a.client.PutPolicyWithContext(ctx, input)
}

func (a *ACMPCAActivities) RestoreCertificateAuthority(ctx context.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error) {
	return a.client.RestoreCertificateAuthorityWithContext(ctx, input)
}

func (a *ACMPCAActivities) RevokeCertificate(ctx context.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error) {
	return a.client.RevokeCertificateWithContext(ctx, input)
}

func (a *ACMPCAActivities) TagCertificateAuthority(ctx context.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error) {
	return a.client.TagCertificateAuthorityWithContext(ctx, input)
}

func (a *ACMPCAActivities) UntagCertificateAuthority(ctx context.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error) {
	return a.client.UntagCertificateAuthorityWithContext(ctx, input)
}

func (a *ACMPCAActivities) UpdateCertificateAuthority(ctx context.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error) {
	return a.client.UpdateCertificateAuthorityWithContext(ctx, input)
}

func (a *ACMPCAActivities) WaitUntilAuditReportCreated(ctx context.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error {
	return a.client.WaitUntilAuditReportCreatedWithContext(ctx, input)

}

func (a *ACMPCAActivities) WaitUntilCertificateAuthorityCSRCreated(ctx context.Context, input *acmpca.GetCertificateAuthorityCsrInput) error {
	return a.client.WaitUntilCertificateAuthorityCSRCreatedWithContext(ctx, input)

}

func (a *ACMPCAActivities) WaitUntilCertificateIssued(ctx context.Context, input *acmpca.GetCertificateInput) error {
	return a.client.WaitUntilCertificateIssuedWithContext(ctx, input)

}
