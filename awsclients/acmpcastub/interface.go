// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package acmpcastub

import (
	"github.com/aws/aws-sdk-go/service/acmpca"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
	CreateCertificateAuthority(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error)
	CreateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityInput) *ACMPCACreateCertificateAuthorityFuture

	CreateCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) (*acmpca.CreateCertificateAuthorityAuditReportOutput, error)
	CreateCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.CreateCertificateAuthorityAuditReportInput) *ACMPCACreateCertificateAuthorityAuditReportFuture

	CreatePermission(ctx workflow.Context, input *acmpca.CreatePermissionInput) (*acmpca.CreatePermissionOutput, error)
	CreatePermissionAsync(ctx workflow.Context, input *acmpca.CreatePermissionInput) *ACMPCACreatePermissionFuture

	DeleteCertificateAuthority(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) (*acmpca.DeleteCertificateAuthorityOutput, error)
	DeleteCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DeleteCertificateAuthorityInput) *ACMPCADeleteCertificateAuthorityFuture

	DeletePermission(ctx workflow.Context, input *acmpca.DeletePermissionInput) (*acmpca.DeletePermissionOutput, error)
	DeletePermissionAsync(ctx workflow.Context, input *acmpca.DeletePermissionInput) *ACMPCADeletePermissionFuture

	DeletePolicy(ctx workflow.Context, input *acmpca.DeletePolicyInput) (*acmpca.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *acmpca.DeletePolicyInput) *ACMPCADeletePolicyFuture

	DescribeCertificateAuthority(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) (*acmpca.DescribeCertificateAuthorityOutput, error)
	DescribeCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityInput) *ACMPCADescribeCertificateAuthorityFuture

	DescribeCertificateAuthorityAuditReport(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) (*acmpca.DescribeCertificateAuthorityAuditReportOutput, error)
	DescribeCertificateAuthorityAuditReportAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *ACMPCADescribeCertificateAuthorityAuditReportFuture

	GetCertificate(ctx workflow.Context, input *acmpca.GetCertificateInput) (*acmpca.GetCertificateOutput, error)
	GetCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *ACMPCAGetCertificateFuture

	GetCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) (*acmpca.GetCertificateAuthorityCertificateOutput, error)
	GetCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCertificateInput) *ACMPCAGetCertificateAuthorityCertificateFuture

	GetCertificateAuthorityCsr(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) (*acmpca.GetCertificateAuthorityCsrOutput, error)
	GetCertificateAuthorityCsrAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *ACMPCAGetCertificateAuthorityCsrFuture

	GetPolicy(ctx workflow.Context, input *acmpca.GetPolicyInput) (*acmpca.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *acmpca.GetPolicyInput) *ACMPCAGetPolicyFuture

	ImportCertificateAuthorityCertificate(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) (*acmpca.ImportCertificateAuthorityCertificateOutput, error)
	ImportCertificateAuthorityCertificateAsync(ctx workflow.Context, input *acmpca.ImportCertificateAuthorityCertificateInput) *ACMPCAImportCertificateAuthorityCertificateFuture

	IssueCertificate(ctx workflow.Context, input *acmpca.IssueCertificateInput) (*acmpca.IssueCertificateOutput, error)
	IssueCertificateAsync(ctx workflow.Context, input *acmpca.IssueCertificateInput) *ACMPCAIssueCertificateFuture

	ListCertificateAuthorities(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) (*acmpca.ListCertificateAuthoritiesOutput, error)
	ListCertificateAuthoritiesAsync(ctx workflow.Context, input *acmpca.ListCertificateAuthoritiesInput) *ACMPCAListCertificateAuthoritiesFuture

	ListPermissions(ctx workflow.Context, input *acmpca.ListPermissionsInput) (*acmpca.ListPermissionsOutput, error)
	ListPermissionsAsync(ctx workflow.Context, input *acmpca.ListPermissionsInput) *ACMPCAListPermissionsFuture

	ListTags(ctx workflow.Context, input *acmpca.ListTagsInput) (*acmpca.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *acmpca.ListTagsInput) *ACMPCAListTagsFuture

	PutPolicy(ctx workflow.Context, input *acmpca.PutPolicyInput) (*acmpca.PutPolicyOutput, error)
	PutPolicyAsync(ctx workflow.Context, input *acmpca.PutPolicyInput) *ACMPCAPutPolicyFuture

	RestoreCertificateAuthority(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) (*acmpca.RestoreCertificateAuthorityOutput, error)
	RestoreCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.RestoreCertificateAuthorityInput) *ACMPCARestoreCertificateAuthorityFuture

	RevokeCertificate(ctx workflow.Context, input *acmpca.RevokeCertificateInput) (*acmpca.RevokeCertificateOutput, error)
	RevokeCertificateAsync(ctx workflow.Context, input *acmpca.RevokeCertificateInput) *ACMPCARevokeCertificateFuture

	TagCertificateAuthority(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) (*acmpca.TagCertificateAuthorityOutput, error)
	TagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.TagCertificateAuthorityInput) *ACMPCATagCertificateAuthorityFuture

	UntagCertificateAuthority(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) (*acmpca.UntagCertificateAuthorityOutput, error)
	UntagCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UntagCertificateAuthorityInput) *ACMPCAUntagCertificateAuthorityFuture

	UpdateCertificateAuthority(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) (*acmpca.UpdateCertificateAuthorityOutput, error)
	UpdateCertificateAuthorityAsync(ctx workflow.Context, input *acmpca.UpdateCertificateAuthorityInput) *ACMPCAUpdateCertificateAuthorityFuture

	WaitUntilAuditReportCreated(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) error
	WaitUntilAuditReportCreatedAsync(ctx workflow.Context, input *acmpca.DescribeCertificateAuthorityAuditReportInput) *awsclients.VoidFuture

	WaitUntilCertificateAuthorityCSRCreated(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) error
	WaitUntilCertificateAuthorityCSRCreatedAsync(ctx workflow.Context, input *acmpca.GetCertificateAuthorityCsrInput) *awsclients.VoidFuture

	WaitUntilCertificateIssued(ctx workflow.Context, input *acmpca.GetCertificateInput) error
	WaitUntilCertificateIssuedAsync(ctx workflow.Context, input *acmpca.GetCertificateInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}

