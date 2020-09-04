
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/xray"
	"github.com/aws/aws-sdk-go/service/xray/xrayiface"
)

type XRayActivities struct {
	client xrayiface.XRayAPI
}

func NewXRayActivities(client xrayiface.XRayAPI) *XRayActivities {
    return &XRayActivities{client: client}
}

func (a *XRayActivities) BatchGetTraces(input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
    return a.client.BatchGetTraces(input)
}

func (a *XRayActivities) CreateGroup(input *xray.CreateGroupInput) (*xray.CreateGroupOutput, error) {
    return a.client.CreateGroup(input)
}

func (a *XRayActivities) CreateSamplingRule(input *xray.CreateSamplingRuleInput) (*xray.CreateSamplingRuleOutput, error) {
    return a.client.CreateSamplingRule(input)
}

func (a *XRayActivities) DeleteGroup(input *xray.DeleteGroupInput) (*xray.DeleteGroupOutput, error) {
    return a.client.DeleteGroup(input)
}

func (a *XRayActivities) DeleteSamplingRule(input *xray.DeleteSamplingRuleInput) (*xray.DeleteSamplingRuleOutput, error) {
    return a.client.DeleteSamplingRule(input)
}

func (a *XRayActivities) GetEncryptionConfig(input *xray.GetEncryptionConfigInput) (*xray.GetEncryptionConfigOutput, error) {
    return a.client.GetEncryptionConfig(input)
}

func (a *XRayActivities) GetGroup(input *xray.GetGroupInput) (*xray.GetGroupOutput, error) {
    return a.client.GetGroup(input)
}

func (a *XRayActivities) GetGroups(input *xray.GetGroupsInput) (*xray.GetGroupsOutput, error) {
    return a.client.GetGroups(input)
}

func (a *XRayActivities) GetSamplingRules(input *xray.GetSamplingRulesInput) (*xray.GetSamplingRulesOutput, error) {
    return a.client.GetSamplingRules(input)
}

func (a *XRayActivities) GetSamplingStatisticSummaries(input *xray.GetSamplingStatisticSummariesInput) (*xray.GetSamplingStatisticSummariesOutput, error) {
    return a.client.GetSamplingStatisticSummaries(input)
}

func (a *XRayActivities) GetSamplingTargets(input *xray.GetSamplingTargetsInput) (*xray.GetSamplingTargetsOutput, error) {
    return a.client.GetSamplingTargets(input)
}

func (a *XRayActivities) GetServiceGraph(input *xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error) {
    return a.client.GetServiceGraph(input)
}

func (a *XRayActivities) GetTimeSeriesServiceStatistics(input *xray.GetTimeSeriesServiceStatisticsInput) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
    return a.client.GetTimeSeriesServiceStatistics(input)
}

func (a *XRayActivities) GetTraceGraph(input *xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error) {
    return a.client.GetTraceGraph(input)
}

func (a *XRayActivities) GetTraceSummaries(input *xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error) {
    return a.client.GetTraceSummaries(input)
}

func (a *XRayActivities) ListTagsForResource(input *xray.ListTagsForResourceInput) (*xray.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *XRayActivities) PutEncryptionConfig(input *xray.PutEncryptionConfigInput) (*xray.PutEncryptionConfigOutput, error) {
    return a.client.PutEncryptionConfig(input)
}

func (a *XRayActivities) PutTelemetryRecords(input *xray.PutTelemetryRecordsInput) (*xray.PutTelemetryRecordsOutput, error) {
    return a.client.PutTelemetryRecords(input)
}

func (a *XRayActivities) PutTraceSegments(input *xray.PutTraceSegmentsInput) (*xray.PutTraceSegmentsOutput, error) {
    return a.client.PutTraceSegments(input)
}

func (a *XRayActivities) TagResource(input *xray.TagResourceInput) (*xray.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *XRayActivities) UntagResource(input *xray.UntagResourceInput) (*xray.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *XRayActivities) UpdateGroup(input *xray.UpdateGroupInput) (*xray.UpdateGroupOutput, error) {
    return a.client.UpdateGroup(input)
}

func (a *XRayActivities) UpdateSamplingRule(input *xray.UpdateSamplingRuleInput) (*xray.UpdateSamplingRuleOutput, error) {
    return a.client.UpdateSamplingRule(input)
}
