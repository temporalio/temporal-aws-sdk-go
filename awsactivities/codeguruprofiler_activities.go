
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler/codeguruprofileriface"
)

type CodeGuruProfilerActivities struct {
    client codeguruprofileriface.CodeGuruProfilerAPI
}

func NewCodeGuruProfilerActivities(session *session.Session, config ...*aws.Config) *CodeGuruProfilerActivities {
    client := codeguruprofiler.New(session, config...)
    return &CodeGuruProfilerActivities{client: client}
}

func (a *CodeGuruProfilerActivities) AddNotificationChannels(input *codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error) {
    return a.client.AddNotificationChannels(input)
}

func (a *CodeGuruProfilerActivities) BatchGetFrameMetricData(input *codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
    return a.client.BatchGetFrameMetricData(input)
}

func (a *CodeGuruProfilerActivities) ConfigureAgent(input *codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error) {
    return a.client.ConfigureAgent(input)
}

func (a *CodeGuruProfilerActivities) CreateProfilingGroup(input *codeguruprofiler.CreateProfilingGroupInput) (*codeguruprofiler.CreateProfilingGroupOutput, error) {
    return a.client.CreateProfilingGroup(input)
}

func (a *CodeGuruProfilerActivities) DeleteProfilingGroup(input *codeguruprofiler.DeleteProfilingGroupInput) (*codeguruprofiler.DeleteProfilingGroupOutput, error) {
    return a.client.DeleteProfilingGroup(input)
}

func (a *CodeGuruProfilerActivities) DescribeProfilingGroup(input *codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
    return a.client.DescribeProfilingGroup(input)
}

func (a *CodeGuruProfilerActivities) GetFindingsReportAccountSummary(input *codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
    return a.client.GetFindingsReportAccountSummary(input)
}

func (a *CodeGuruProfilerActivities) GetNotificationConfiguration(input *codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
    return a.client.GetNotificationConfiguration(input)
}

func (a *CodeGuruProfilerActivities) GetPolicy(input *codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error) {
    return a.client.GetPolicy(input)
}

func (a *CodeGuruProfilerActivities) GetProfile(input *codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error) {
    return a.client.GetProfile(input)
}

func (a *CodeGuruProfilerActivities) GetRecommendations(input *codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error) {
    return a.client.GetRecommendations(input)
}

func (a *CodeGuruProfilerActivities) ListFindingsReports(input *codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error) {
    return a.client.ListFindingsReports(input)
}

func (a *CodeGuruProfilerActivities) ListProfileTimes(input *codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error) {
    return a.client.ListProfileTimes(input)
}

func (a *CodeGuruProfilerActivities) ListProfilingGroups(input *codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
    return a.client.ListProfilingGroups(input)
}

func (a *CodeGuruProfilerActivities) ListTagsForResource(input *codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *CodeGuruProfilerActivities) PostAgentProfile(input *codeguruprofiler.PostAgentProfileInput) (*codeguruprofiler.PostAgentProfileOutput, error) {
    return a.client.PostAgentProfile(input)
}

func (a *CodeGuruProfilerActivities) PutPermission(input *codeguruprofiler.PutPermissionInput) (*codeguruprofiler.PutPermissionOutput, error) {
    return a.client.PutPermission(input)
}

func (a *CodeGuruProfilerActivities) RemoveNotificationChannel(input *codeguruprofiler.RemoveNotificationChannelInput) (*codeguruprofiler.RemoveNotificationChannelOutput, error) {
    return a.client.RemoveNotificationChannel(input)
}

func (a *CodeGuruProfilerActivities) RemovePermission(input *codeguruprofiler.RemovePermissionInput) (*codeguruprofiler.RemovePermissionOutput, error) {
    return a.client.RemovePermission(input)
}

func (a *CodeGuruProfilerActivities) SubmitFeedback(input *codeguruprofiler.SubmitFeedbackInput) (*codeguruprofiler.SubmitFeedbackOutput, error) {
    return a.client.SubmitFeedback(input)
}

func (a *CodeGuruProfilerActivities) TagResource(input *codeguruprofiler.TagResourceInput) (*codeguruprofiler.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *CodeGuruProfilerActivities) UntagResource(input *codeguruprofiler.UntagResourceInput) (*codeguruprofiler.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *CodeGuruProfilerActivities) UpdateProfilingGroup(input *codeguruprofiler.UpdateProfilingGroupInput) (*codeguruprofiler.UpdateProfilingGroupOutput, error) {
    return a.client.UpdateProfilingGroup(input)
}
