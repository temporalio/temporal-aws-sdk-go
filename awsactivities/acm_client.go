package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/acm"
	"go.temporal.io/sdk/workflow"
)

type ACMClient interface {
    AddTagsToCertificate(ctx workflow.Context, input *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error)
    AddTagsToCertificateAsync(ctx workflow.Context, input *acm.AddTagsToCertificateInput) *AcmAddTagsToCertificateResult

    DeleteCertificate(ctx workflow.Context, input *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error)
    DeleteCertificateAsync(ctx workflow.Context, input *acm.DeleteCertificateInput) *AcmDeleteCertificateResult

    DescribeCertificate(ctx workflow.Context, input *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error)
    DescribeCertificateAsync(ctx workflow.Context, input *acm.DescribeCertificateInput) *AcmDescribeCertificateResult

    ExportCertificate(ctx workflow.Context, input *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error)
    ExportCertificateAsync(ctx workflow.Context, input *acm.ExportCertificateInput) *AcmExportCertificateResult

    GetCertificate(ctx workflow.Context, input *acm.GetCertificateInput) (*acm.GetCertificateOutput, error)
    GetCertificateAsync(ctx workflow.Context, input *acm.GetCertificateInput) *AcmGetCertificateResult

    ImportCertificate(ctx workflow.Context, input *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error)
    ImportCertificateAsync(ctx workflow.Context, input *acm.ImportCertificateInput) *AcmImportCertificateResult

    ListCertificates(ctx workflow.Context, input *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error)
    ListCertificatesAsync(ctx workflow.Context, input *acm.ListCertificatesInput) *AcmListCertificatesResult

    ListTagsForCertificate(ctx workflow.Context, input *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error)
    ListTagsForCertificateAsync(ctx workflow.Context, input *acm.ListTagsForCertificateInput) *AcmListTagsForCertificateResult

    RemoveTagsFromCertificate(ctx workflow.Context, input *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error)
    RemoveTagsFromCertificateAsync(ctx workflow.Context, input *acm.RemoveTagsFromCertificateInput) *AcmRemoveTagsFromCertificateResult

    RenewCertificate(ctx workflow.Context, input *acm.RenewCertificateInput) (*acm.RenewCertificateOutput, error)
    RenewCertificateAsync(ctx workflow.Context, input *acm.RenewCertificateInput) *AcmRenewCertificateResult

    RequestCertificate(ctx workflow.Context, input *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error)
    RequestCertificateAsync(ctx workflow.Context, input *acm.RequestCertificateInput) *AcmRequestCertificateResult

    ResendValidationEmail(ctx workflow.Context, input *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error)
    ResendValidationEmailAsync(ctx workflow.Context, input *acm.ResendValidationEmailInput) *AcmResendValidationEmailResult

    UpdateCertificateOptions(ctx workflow.Context, input *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error)
    UpdateCertificateOptionsAsync(ctx workflow.Context, input *acm.UpdateCertificateOptionsInput) *AcmUpdateCertificateOptionsResult

    WaitUntilCertificateValidated(ctx workflow.Context, input *acm.DescribeCertificateInput) error}
type AcmAddTagsToCertificateResult struct {
	Result workflow.Future
}

func (r *AcmAddTagsToCertificateResult) Get(ctx workflow.Context) (*acm.AddTagsToCertificateOutput, error) {
    var output acm.AddTagsToCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmDeleteCertificateResult struct {
	Result workflow.Future
}

func (r *AcmDeleteCertificateResult) Get(ctx workflow.Context) (*acm.DeleteCertificateOutput, error) {
    var output acm.DeleteCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmDescribeCertificateResult struct {
	Result workflow.Future
}

func (r *AcmDescribeCertificateResult) Get(ctx workflow.Context) (*acm.DescribeCertificateOutput, error) {
    var output acm.DescribeCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmExportCertificateResult struct {
	Result workflow.Future
}

func (r *AcmExportCertificateResult) Get(ctx workflow.Context) (*acm.ExportCertificateOutput, error) {
    var output acm.ExportCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmGetCertificateResult struct {
	Result workflow.Future
}

func (r *AcmGetCertificateResult) Get(ctx workflow.Context) (*acm.GetCertificateOutput, error) {
    var output acm.GetCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmImportCertificateResult struct {
	Result workflow.Future
}

func (r *AcmImportCertificateResult) Get(ctx workflow.Context) (*acm.ImportCertificateOutput, error) {
    var output acm.ImportCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmListCertificatesResult struct {
	Result workflow.Future
}

func (r *AcmListCertificatesResult) Get(ctx workflow.Context) (*acm.ListCertificatesOutput, error) {
    var output acm.ListCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmListTagsForCertificateResult struct {
	Result workflow.Future
}

func (r *AcmListTagsForCertificateResult) Get(ctx workflow.Context) (*acm.ListTagsForCertificateOutput, error) {
    var output acm.ListTagsForCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmRemoveTagsFromCertificateResult struct {
	Result workflow.Future
}

func (r *AcmRemoveTagsFromCertificateResult) Get(ctx workflow.Context) (*acm.RemoveTagsFromCertificateOutput, error) {
    var output acm.RemoveTagsFromCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmRenewCertificateResult struct {
	Result workflow.Future
}

func (r *AcmRenewCertificateResult) Get(ctx workflow.Context) (*acm.RenewCertificateOutput, error) {
    var output acm.RenewCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmRequestCertificateResult struct {
	Result workflow.Future
}

func (r *AcmRequestCertificateResult) Get(ctx workflow.Context) (*acm.RequestCertificateOutput, error) {
    var output acm.RequestCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmResendValidationEmailResult struct {
	Result workflow.Future
}

func (r *AcmResendValidationEmailResult) Get(ctx workflow.Context) (*acm.ResendValidationEmailOutput, error) {
    var output acm.ResendValidationEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AcmUpdateCertificateOptionsResult struct {
	Result workflow.Future
}

func (r *AcmUpdateCertificateOptionsResult) Get(ctx workflow.Context) (*acm.UpdateCertificateOptionsOutput, error) {
    var output acm.UpdateCertificateOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ACMStub struct {
    activities ACMClient
}

func NewACMStub() ACMClient {
    return &ACMStub{}
}

func (a *ACMStub) AddTagsToCertificate(ctx workflow.Context, input *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error) {
    var output acm.AddTagsToCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) DeleteCertificate(ctx workflow.Context, input *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error) {
    var output acm.DeleteCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) DescribeCertificate(ctx workflow.Context, input *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error) {
    var output acm.DescribeCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) ExportCertificate(ctx workflow.Context, input *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error) {
    var output acm.ExportCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) GetCertificate(ctx workflow.Context, input *acm.GetCertificateInput) (*acm.GetCertificateOutput, error) {
    var output acm.GetCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) ImportCertificate(ctx workflow.Context, input *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error) {
    var output acm.ImportCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) ListCertificates(ctx workflow.Context, input *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error) {
    var output acm.ListCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) ListTagsForCertificate(ctx workflow.Context, input *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error) {
    var output acm.ListTagsForCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) RemoveTagsFromCertificate(ctx workflow.Context, input *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error) {
    var output acm.RemoveTagsFromCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) RenewCertificate(ctx workflow.Context, input *acm.RenewCertificateInput) (*acm.RenewCertificateOutput, error) {
    var output acm.RenewCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RenewCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) RequestCertificate(ctx workflow.Context, input *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error) {
    var output acm.RequestCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RequestCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) ResendValidationEmail(ctx workflow.Context, input *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error) {
    var output acm.ResendValidationEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResendValidationEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *ACMStub) UpdateCertificateOptions(ctx workflow.Context, input *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error) {
    var output acm.UpdateCertificateOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCertificateOptions, input).Get(ctx, &output)
    return &output, err
}


func (a *ACMStub) WaitUntilCertificateValidated(ctx workflow.Context, input *acm.DescribeCertificateInput) error {
    return a.client.WaitUntilCertificateValidated(input)
}
