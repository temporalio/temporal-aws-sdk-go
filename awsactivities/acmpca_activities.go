package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/acmpca/acmpcaiface"
)

type ACMPCAActivities struct {
	client acmpcaiface.ACMPCAAPI
}

func NewACMPCAActivities(client acmpcaiface.ACMPCAAPI) *ACMPCAActivities {
    return &ACMPCAActivities{client: client}
}


func (a *ACMPCAActivities) CreateCertificateAuthority(input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
    return a.client.CreateCertificateAuthority(input)
}



func (a *ACMPCAActivities) CreateCertificateAuthorityAuditReport(input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
    return a.client.CreateCertificateAuthorityAuditReport(input)
}



func (a *ACMPCAActivities) CreatePermission(input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error) {
    return a.client.CreatePermission(input)
}



func (a *ACMPCAActivities) DeleteCertificateAuthority(input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error) {
    return a.client.DeleteCertificateAuthority(input)
}



func (a *ACMPCAActivities) DeletePermission(input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error) {
    return a.client.DeletePermission(input)
}



func (a *ACMPCAActivities) DeletePolicy(input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error) {
    return a.client.DeletePolicy(input)
}



func (a *ACMPCAActivities) DescribeCertificateAuthority(input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
    return a.client.DescribeCertificateAuthority(input)
}



func (a *ACMPCAActivities) DescribeCertificateAuthorityAuditReport(input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
    return a.client.DescribeCertificateAuthorityAuditReport(input)
}



func (a *ACMPCAActivities) GetCertificate(input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
    return a.client.GetCertificate(input)
}



func (a *ACMPCAActivities) GetCertificateAuthorityCertificate(input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
    return a.client.GetCertificateAuthorityCertificate(input)
}



func (a *ACMPCAActivities) GetCertificateAuthorityCsr(input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
    return a.client.GetCertificateAuthorityCsr(input)
}



func (a *ACMPCAActivities) GetPolicy(input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
    return a.client.GetPolicy(input)
}



func (a *ACMPCAActivities) ImportCertificateAuthorityCertificate(input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
    return a.client.ImportCertificateAuthorityCertificate(input)
}



func (a *ACMPCAActivities) IssueCertificate(input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error) {
    return a.client.IssueCertificate(input)
}



func (a *ACMPCAActivities) ListCertificateAuthorities(input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
    return a.client.ListCertificateAuthorities(input)
}



func (a *ACMPCAActivities) ListPermissions(input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
    return a.client.ListPermissions(input)
}



func (a *ACMPCAActivities) ListTags(input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
    return a.client.ListTags(input)
}



func (a *ACMPCAActivities) PutPolicy(input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error) {
    return a.client.PutPolicy(input)
}



func (a *ACMPCAActivities) RestoreCertificateAuthority(input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error) {
    return a.client.RestoreCertificateAuthority(input)
}



func (a *ACMPCAActivities) RevokeCertificate(input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error) {
    return a.client.RevokeCertificate(input)
}



func (a *ACMPCAActivities) TagCertificateAuthority(input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error) {
    return a.client.TagCertificateAuthority(input)
}



func (a *ACMPCAActivities) UntagCertificateAuthority(input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error) {
    return a.client.UntagCertificateAuthority(input)
}



func (a *ACMPCAActivities) UpdateCertificateAuthority(input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error) {
    return a.client.UpdateCertificateAuthority(input)
}



func (a *ACMPCAActivities) WaitUntilAuditReportCreated(input *acmpca.DescribeCertificateAuthorityAuditReportInput) error {
    return a.client.WaitUntilAuditReportCreated(input)
}



func (a *ACMPCAActivities) WaitUntilCertificateAuthorityCSRCreated(input *acmpca.GetCertificateAuthorityCsrInput) error {
    return a.client.WaitUntilCertificateAuthorityCSRCreated(input)
}



func (a *ACMPCAActivities) WaitUntilCertificateIssued(input *acmpca.GetCertificateInput) error {
    return a.client.WaitUntilCertificateIssued(input)
}

