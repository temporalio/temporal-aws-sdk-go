// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"go.temporal.io/sdk/workflow"
)

type CloudHSMClient interface {
	AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *CloudhsmAddTagsToResourceFuture

	CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)
	CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CloudhsmCreateHapgFuture

	CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)
	CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CloudhsmCreateHsmFuture

	CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)
	CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CloudhsmCreateLunaClientFuture

	DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)
	DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *CloudhsmDeleteHapgFuture

	DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)
	DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *CloudhsmDeleteHsmFuture

	DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)
	DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *CloudhsmDeleteLunaClientFuture

	DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)
	DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *CloudhsmDescribeHapgFuture

	DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)
	DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *CloudhsmDescribeHsmFuture

	DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)
	DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *CloudhsmDescribeLunaClientFuture

	GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)
	GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *CloudhsmGetConfigFuture

	ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)
	ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *CloudhsmListAvailableZonesFuture

	ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)
	ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *CloudhsmListHapgsFuture

	ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)
	ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *CloudhsmListHsmsFuture

	ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)
	ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *CloudhsmListLunaClientsFuture

	ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *CloudhsmListTagsForResourceFuture

	ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)
	ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *CloudhsmModifyHapgFuture

	ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)
	ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *CloudhsmModifyHsmFuture

	ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)
	ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *CloudhsmModifyLunaClientFuture

	RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *CloudhsmRemoveTagsFromResourceFuture
}

type CloudHSMStub struct{}

func NewCloudHSMStub() CloudHSMClient {
	return &CloudHSMStub{}
}

type CloudhsmAddTagsToResourceFuture struct {
	Future workflow.Future
}

func (r *CloudhsmAddTagsToResourceFuture) Get(ctx workflow.Context) (*cloudhsm.AddTagsToResourceOutput, error) {
	var output cloudhsm.AddTagsToResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmCreateHapgFuture struct {
	Future workflow.Future
}

func (r *CloudhsmCreateHapgFuture) Get(ctx workflow.Context) (*cloudhsm.CreateHapgOutput, error) {
	var output cloudhsm.CreateHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmCreateHsmFuture struct {
	Future workflow.Future
}

func (r *CloudhsmCreateHsmFuture) Get(ctx workflow.Context) (*cloudhsm.CreateHsmOutput, error) {
	var output cloudhsm.CreateHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmCreateLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudhsmCreateLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.CreateLunaClientOutput, error) {
	var output cloudhsm.CreateLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmDeleteHapgFuture struct {
	Future workflow.Future
}

func (r *CloudhsmDeleteHapgFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteHapgOutput, error) {
	var output cloudhsm.DeleteHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmDeleteHsmFuture struct {
	Future workflow.Future
}

func (r *CloudhsmDeleteHsmFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteHsmOutput, error) {
	var output cloudhsm.DeleteHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmDeleteLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudhsmDeleteLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteLunaClientOutput, error) {
	var output cloudhsm.DeleteLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmDescribeHapgFuture struct {
	Future workflow.Future
}

func (r *CloudhsmDescribeHapgFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeHapgOutput, error) {
	var output cloudhsm.DescribeHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmDescribeHsmFuture struct {
	Future workflow.Future
}

func (r *CloudhsmDescribeHsmFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeHsmOutput, error) {
	var output cloudhsm.DescribeHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmDescribeLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudhsmDescribeLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeLunaClientOutput, error) {
	var output cloudhsm.DescribeLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmGetConfigFuture struct {
	Future workflow.Future
}

func (r *CloudhsmGetConfigFuture) Get(ctx workflow.Context) (*cloudhsm.GetConfigOutput, error) {
	var output cloudhsm.GetConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmListAvailableZonesFuture struct {
	Future workflow.Future
}

func (r *CloudhsmListAvailableZonesFuture) Get(ctx workflow.Context) (*cloudhsm.ListAvailableZonesOutput, error) {
	var output cloudhsm.ListAvailableZonesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmListHapgsFuture struct {
	Future workflow.Future
}

func (r *CloudhsmListHapgsFuture) Get(ctx workflow.Context) (*cloudhsm.ListHapgsOutput, error) {
	var output cloudhsm.ListHapgsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmListHsmsFuture struct {
	Future workflow.Future
}

func (r *CloudhsmListHsmsFuture) Get(ctx workflow.Context) (*cloudhsm.ListHsmsOutput, error) {
	var output cloudhsm.ListHsmsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmListLunaClientsFuture struct {
	Future workflow.Future
}

func (r *CloudhsmListLunaClientsFuture) Get(ctx workflow.Context) (*cloudhsm.ListLunaClientsOutput, error) {
	var output cloudhsm.ListLunaClientsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *CloudhsmListTagsForResourceFuture) Get(ctx workflow.Context) (*cloudhsm.ListTagsForResourceOutput, error) {
	var output cloudhsm.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmModifyHapgFuture struct {
	Future workflow.Future
}

func (r *CloudhsmModifyHapgFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyHapgOutput, error) {
	var output cloudhsm.ModifyHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmModifyHsmFuture struct {
	Future workflow.Future
}

func (r *CloudhsmModifyHsmFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyHsmOutput, error) {
	var output cloudhsm.ModifyHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmModifyLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudhsmModifyLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyLunaClientOutput, error) {
	var output cloudhsm.ModifyLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudhsmRemoveTagsFromResourceFuture struct {
	Future workflow.Future
}

func (r *CloudhsmRemoveTagsFromResourceFuture) Get(ctx workflow.Context) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	var output cloudhsm.RemoveTagsFromResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
	var output cloudhsm.AddTagsToResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.AddTagsToResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *CloudhsmAddTagsToResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.AddTagsToResource", input)
	return &CloudhsmAddTagsToResourceFuture{Future: future}
}

func (a *CloudHSMStub) CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error) {
	var output cloudhsm.CreateHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CloudhsmCreateHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHapg", input)
	return &CloudhsmCreateHapgFuture{Future: future}
}

func (a *CloudHSMStub) CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error) {
	var output cloudhsm.CreateHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CloudhsmCreateHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHsm", input)
	return &CloudhsmCreateHsmFuture{Future: future}
}

func (a *CloudHSMStub) CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error) {
	var output cloudhsm.CreateLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CloudhsmCreateLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateLunaClient", input)
	return &CloudhsmCreateLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error) {
	var output cloudhsm.DeleteHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *CloudhsmDeleteHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHapg", input)
	return &CloudhsmDeleteHapgFuture{Future: future}
}

func (a *CloudHSMStub) DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error) {
	var output cloudhsm.DeleteHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *CloudhsmDeleteHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHsm", input)
	return &CloudhsmDeleteHsmFuture{Future: future}
}

func (a *CloudHSMStub) DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error) {
	var output cloudhsm.DeleteLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *CloudhsmDeleteLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteLunaClient", input)
	return &CloudhsmDeleteLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
	var output cloudhsm.DescribeHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *CloudhsmDescribeHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHapg", input)
	return &CloudhsmDescribeHapgFuture{Future: future}
}

func (a *CloudHSMStub) DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
	var output cloudhsm.DescribeHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *CloudhsmDescribeHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHsm", input)
	return &CloudhsmDescribeHsmFuture{Future: future}
}

func (a *CloudHSMStub) DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
	var output cloudhsm.DescribeLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *CloudhsmDescribeLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeLunaClient", input)
	return &CloudhsmDescribeLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
	var output cloudhsm.GetConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.GetConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *CloudhsmGetConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.GetConfig", input)
	return &CloudhsmGetConfigFuture{Future: future}
}

func (a *CloudHSMStub) ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
	var output cloudhsm.ListAvailableZonesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListAvailableZones", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *CloudhsmListAvailableZonesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListAvailableZones", input)
	return &CloudhsmListAvailableZonesFuture{Future: future}
}

func (a *CloudHSMStub) ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
	var output cloudhsm.ListHapgsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHapgs", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *CloudhsmListHapgsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHapgs", input)
	return &CloudhsmListHapgsFuture{Future: future}
}

func (a *CloudHSMStub) ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
	var output cloudhsm.ListHsmsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHsms", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *CloudhsmListHsmsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHsms", input)
	return &CloudhsmListHsmsFuture{Future: future}
}

func (a *CloudHSMStub) ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
	var output cloudhsm.ListLunaClientsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListLunaClients", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *CloudhsmListLunaClientsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListLunaClients", input)
	return &CloudhsmListLunaClientsFuture{Future: future}
}

func (a *CloudHSMStub) ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
	var output cloudhsm.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *CloudhsmListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListTagsForResource", input)
	return &CloudhsmListTagsForResourceFuture{Future: future}
}

func (a *CloudHSMStub) ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error) {
	var output cloudhsm.ModifyHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *CloudhsmModifyHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHapg", input)
	return &CloudhsmModifyHapgFuture{Future: future}
}

func (a *CloudHSMStub) ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error) {
	var output cloudhsm.ModifyHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *CloudhsmModifyHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHsm", input)
	return &CloudhsmModifyHsmFuture{Future: future}
}

func (a *CloudHSMStub) ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error) {
	var output cloudhsm.ModifyLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *CloudhsmModifyLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyLunaClient", input)
	return &CloudhsmModifyLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	var output cloudhsm.RemoveTagsFromResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.RemoveTagsFromResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *CloudhsmRemoveTagsFromResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.RemoveTagsFromResource", input)
	return &CloudhsmRemoveTagsFromResourceFuture{Future: future}
}
