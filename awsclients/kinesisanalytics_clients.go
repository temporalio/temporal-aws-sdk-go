// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"go.temporal.io/sdk/workflow"
)

type KinesisAnalyticsClient interface {
	AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsAddApplicationCloudWatchLoggingOptionFuture

	AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error)
	AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *KinesisanalyticsAddApplicationInputFuture

	AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *KinesisanalyticsAddApplicationInputProcessingConfigurationFuture

	AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error)
	AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *KinesisanalyticsAddApplicationOutputFuture

	AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *KinesisanalyticsAddApplicationReferenceDataSourceFuture

	CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *KinesisanalyticsCreateApplicationFuture

	DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *KinesisanalyticsDeleteApplicationFuture

	DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionFuture

	DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *KinesisanalyticsDeleteApplicationInputProcessingConfigurationFuture

	DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *KinesisanalyticsDeleteApplicationOutputFuture

	DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *KinesisanalyticsDeleteApplicationReferenceDataSourceFuture

	DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error)
	DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *KinesisanalyticsDescribeApplicationFuture

	DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *KinesisanalyticsDiscoverInputSchemaFuture

	ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *KinesisanalyticsListApplicationsFuture

	ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *KinesisanalyticsListTagsForResourceFuture

	StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error)
	StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *KinesisanalyticsStartApplicationFuture

	StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error)
	StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *KinesisanalyticsStopApplicationFuture

	TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *KinesisanalyticsTagResourceFuture

	UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *KinesisanalyticsUntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *KinesisanalyticsUpdateApplicationFuture
}

type KinesisAnalyticsStub struct{}

func NewKinesisAnalyticsStub() KinesisAnalyticsClient {
	return &KinesisAnalyticsStub{}
}

type KinesisanalyticsAddApplicationCloudWatchLoggingOptionFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsAddApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationInputFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsAddApplicationInputFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationInputProcessingConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsAddApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationOutputFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsAddApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsAddApplicationReferenceDataSourceFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsAddApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsCreateApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsCreateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationInputProcessingConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationOutputFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDeleteApplicationReferenceDataSourceFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsDeleteApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDescribeApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsDescribeApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsDiscoverInputSchemaFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsDiscoverInputSchemaFuture) Get(ctx workflow.Context) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsListApplicationsFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsListApplicationsFuture) Get(ctx workflow.Context) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsListTagsForResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsStartApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsStartApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsStopApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsStopApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsTagResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsTagResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsUntagResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsUntagResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisanalyticsUpdateApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisanalyticsUpdateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsAddApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationCloudWatchLoggingOption", input)
	return &KinesisanalyticsAddApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *KinesisanalyticsAddApplicationInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInput", input)
	return &KinesisanalyticsAddApplicationInputFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *KinesisanalyticsAddApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInputProcessingConfiguration", input)
	return &KinesisanalyticsAddApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *KinesisanalyticsAddApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationOutput", input)
	return &KinesisanalyticsAddApplicationOutputFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *KinesisanalyticsAddApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationReferenceDataSource", input)
	return &KinesisanalyticsAddApplicationReferenceDataSourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *KinesisanalyticsCreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.CreateApplication", input)
	return &KinesisanalyticsCreateApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *KinesisanalyticsDeleteApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplication", input)
	return &KinesisanalyticsDeleteApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationCloudWatchLoggingOption", input)
	return &KinesisanalyticsDeleteApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *KinesisanalyticsDeleteApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationInputProcessingConfiguration", input)
	return &KinesisanalyticsDeleteApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *KinesisanalyticsDeleteApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationOutput", input)
	return &KinesisanalyticsDeleteApplicationOutputFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *KinesisanalyticsDeleteApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationReferenceDataSource", input)
	return &KinesisanalyticsDeleteApplicationReferenceDataSourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DescribeApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *KinesisanalyticsDescribeApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DescribeApplication", input)
	return &KinesisanalyticsDescribeApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DiscoverInputSchema", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *KinesisanalyticsDiscoverInputSchemaFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DiscoverInputSchema", input)
	return &KinesisanalyticsDiscoverInputSchemaFuture{Future: future}
}

func (a *KinesisAnalyticsStub) ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *KinesisanalyticsListApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListApplications", input)
	return &KinesisanalyticsListApplicationsFuture{Future: future}
}

func (a *KinesisAnalyticsStub) ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *KinesisanalyticsListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListTagsForResource", input)
	return &KinesisanalyticsListTagsForResourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StartApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *KinesisanalyticsStartApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StartApplication", input)
	return &KinesisanalyticsStartApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StopApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *KinesisanalyticsStopApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StopApplication", input)
	return &KinesisanalyticsStopApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *KinesisanalyticsTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.TagResource", input)
	return &KinesisanalyticsTagResourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *KinesisanalyticsUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UntagResource", input)
	return &KinesisanalyticsUntagResourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *KinesisanalyticsUpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UpdateApplication", input)
	return &KinesisanalyticsUpdateApplicationFuture{Future: future}
}
