package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/acmpca"
	"go.temporal.io/sdk/workflow"
)

type ACMPCAClient interface {
    CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error)
    CreateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) *AcmpcaCreateCertificateAuthorityResult

    CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error)
    CreateCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) *AcmpcaCreateCertificateAuthorityAuditReportResult

    CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error)
    CreatePermissionAsync(ctx workflow.Context, input *acmpca.CreatePermissionInput) *AcmpcaCreatePermissionResult

    DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error)
    DeleteCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) *AcmpcaDeleteCertificateAuthorityResult

    DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error)
    DeletePermissionAsync(ctx workflow.Context, input *acmpca.DeletePermissionInput) *AcmpcaDeletePermissionResult

    DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error)
    DeletePolicyAsync(ctx workflow.Context, input *acmpca.DeletePolicyInput) *AcmpcaDeletePolicyResult

    DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error)
    DescribeCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) *AcmpcaDescribeCertificateAuthorityResult

    DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error)
    DescribeCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *AcmpcaDescribeCertificateAuthorityAuditReportResult

    GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error)
    GetCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *AcmpcaGetCertificateResult

    GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error)
    GetCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) *AcmpcaGetCertificateAuthorityCertificateResult

    GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error)
    GetCertificateAuthorityCsrAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *AcmpcaGetCertificateAuthorityCsrResult

    GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error)
    GetPolicyAsync(ctx workflow.Context, input *acmpca.GetPolicyInput) *AcmpcaGetPolicyResult

    ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error)
    ImportCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) *AcmpcaImportCertificateAuthorityCertificateResult

    IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error)
    IssueCertificateAsync(ctx workflow.Context, input *acmpca.IssueCertificateInput) *AcmpcaIssueCertificateResult

    ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error)
    ListCertificateAuthoritiesAsync(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) *AcmpcaListCertificateAuthoritiesResult

    ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error)
    ListPermissionsAsync(ctx workflow.Context, input *acmpca.ListPermissionsInput) *AcmpcaListPermissionsResult

    ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error)
    ListTagsAsync(ctx workflow.Context, input *acmpca.ListTagsInput) *AcmpcaListTagsResult

    PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error)
    PutPolicyAsync(ctx workflow.Context, input *acmpca.PutPolicyInput) *AcmpcaPutPolicyResult

    RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error)
    RestoreCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) *AcmpcaRestoreCertificateAuthorityResult

    RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error)
    RevokeCertificateAsync(ctx workflow.Context, input *acmpca.RevokeCertificateInput) *AcmpcaRevokeCertificateResult

    TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error)
    TagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) *AcmpcaTagCertificateAuthorityResult

    UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error)
    UntagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) *AcmpcaUntagCertificateAuthorityResult

    UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error)
    UpdateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) *AcmpcaUpdateCertificateAuthorityResult

    WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error
    WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error
    WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error}
type AcmpcaCreateCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaCreateCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityOutput, error) {
    var output acmpca.CreateCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaCreateCertificateAuthorityAuditReportResult struct {
	Result workflow.Future
}

func (r *AcmpcaCreateCertificateAuthorityAuditReportResult) Get(ctx workflow.Context) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
    var output acmpca.CreateCertificateAuthorityAuditReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaCreatePermissionResult struct {
	Result workflow.Future
}

func (r *AcmpcaCreatePermissionResult) Get(ctx workflow.Context) (*acmpca.CreatePermissionOutput, error) {
    var output acmpca.CreatePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaDeleteCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaDeleteCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.DeleteCertificateAuthorityOutput, error) {
    var output acmpca.DeleteCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaDeletePermissionResult struct {
	Result workflow.Future
}

func (r *AcmpcaDeletePermissionResult) Get(ctx workflow.Context) (*acmpca.DeletePermissionOutput, error) {
    var output acmpca.DeletePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaDeletePolicyResult struct {
	Result workflow.Future
}

func (r *AcmpcaDeletePolicyResult) Get(ctx workflow.Context) (*acmpca.DeletePolicyOutput, error) {
    var output acmpca.DeletePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaDescribeCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaDescribeCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityOutput, error) {
    var output acmpca.DescribeCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaDescribeCertificateAuthorityAuditReportResult struct {
	Result workflow.Future
}

func (r *AcmpcaDescribeCertificateAuthorityAuditReportResult) Get(ctx workflow.Context) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
    var output acmpca.DescribeCertificateAuthorityAuditReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaGetCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetCertificateResult) Get(ctx workflow.Context) (*acmpca.GetCertificateOutput, error) {
    var output acmpca.GetCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaGetCertificateAuthorityCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetCertificateAuthorityCertificateResult) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
    var output acmpca.GetCertificateAuthorityCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaGetCertificateAuthorityCsrResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetCertificateAuthorityCsrResult) Get(ctx workflow.Context) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
    var output acmpca.GetCertificateAuthorityCsrOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaGetPolicyResult struct {
	Result workflow.Future
}

func (r *AcmpcaGetPolicyResult) Get(ctx workflow.Context) (*acmpca.GetPolicyOutput, error) {
    var output acmpca.GetPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaImportCertificateAuthorityCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaImportCertificateAuthorityCertificateResult) Get(ctx workflow.Context) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
    var output acmpca.ImportCertificateAuthorityCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaIssueCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaIssueCertificateResult) Get(ctx workflow.Context) (*acmpca.IssueCertificateOutput, error) {
    var output acmpca.IssueCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaListCertificateAuthoritiesResult struct {
	Result workflow.Future
}

func (r *AcmpcaListCertificateAuthoritiesResult) Get(ctx workflow.Context) (*acmpca.ListCertificateAuthoritiesOutput, error) {
    var output acmpca.ListCertificateAuthoritiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaListPermissionsResult struct {
	Result workflow.Future
}

func (r *AcmpcaListPermissionsResult) Get(ctx workflow.Context) (*acmpca.ListPermissionsOutput, error) {
    var output acmpca.ListPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaListTagsResult struct {
	Result workflow.Future
}

func (r *AcmpcaListTagsResult) Get(ctx workflow.Context) (*acmpca.ListTagsOutput, error) {
    var output acmpca.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaPutPolicyResult struct {
	Result workflow.Future
}

func (r *AcmpcaPutPolicyResult) Get(ctx workflow.Context) (*acmpca.PutPolicyOutput, error) {
    var output acmpca.PutPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaRestoreCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaRestoreCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.RestoreCertificateAuthorityOutput, error) {
    var output acmpca.RestoreCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaRevokeCertificateResult struct {
	Result workflow.Future
}

func (r *AcmpcaRevokeCertificateResult) Get(ctx workflow.Context) (*acmpca.RevokeCertificateOutput, error) {
    var output acmpca.RevokeCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaTagCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaTagCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.TagCertificateAuthorityOutput, error) {
    var output acmpca.TagCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaUntagCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaUntagCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.UntagCertificateAuthorityOutput, error) {
    var output acmpca.UntagCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmpcaUpdateCertificateAuthorityResult struct {
	Result workflow.Future
}

func (r *AcmpcaUpdateCertificateAuthorityResult) Get(ctx workflow.Context) (*acmpca.UpdateCertificateAuthorityOutput, error) {
    var output acmpca.UpdateCertificateAuthorityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ACMPCAStub struct {
    activities ACMPCAClient
}

func NewACMPCAStub() ACMPCAClient {
    return &ACMPCAStub{}
}

func (a *ACMPCAStub) CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
    var output acmpca.CreateCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error) {
    var output acmpca.CreateCertificateAuthorityAuditReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCertificateAuthorityAuditReport, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error) {
    var output acmpca.CreatePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error) {
    var output acmpca.DeleteCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error) {
    var output acmpca.DeletePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error) {
    var output acmpca.DeletePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error) {
    var output acmpca.DescribeCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error) {
    var output acmpca.DescribeCertificateAuthorityAuditReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificateAuthorityAuditReport, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error) {
    var output acmpca.GetCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error) {
    var output acmpca.GetCertificateAuthorityCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCertificateAuthorityCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error) {
    var output acmpca.GetCertificateAuthorityCsrOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCertificateAuthorityCsr, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error) {
    var output acmpca.GetPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error) {
    var output acmpca.ImportCertificateAuthorityCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportCertificateAuthorityCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error) {
    var output acmpca.IssueCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.IssueCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error) {
    var output acmpca.ListCertificateAuthoritiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCertificateAuthorities, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error) {
    var output acmpca.ListPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error) {
    var output acmpca.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error) {
    var output acmpca.PutPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error) {
    var output acmpca.RestoreCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error) {
    var output acmpca.RevokeCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error) {
    var output acmpca.TagCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error) {
    var output acmpca.UntagCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMPCAStub) UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error) {
    var output acmpca.UpdateCertificateAuthorityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCertificateAuthority, input).Get(ctx, &output)
    return &output, err
}


func (a *ACMPCAStub) WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error {
    return a.client.WaitUntilAuditReportCreated(input)
}


func (a *ACMPCAStub) WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error {
    return a.client.WaitUntilCertificateAuthorityCSRCreated(input)
}


func (a *ACMPCAStub) WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error {
    return a.client.WaitUntilCertificateIssued(input)
}
