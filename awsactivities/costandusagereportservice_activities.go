package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface"
)

type CostandUsageReportServiceActivities struct {
	client costandusagereportserviceiface.CostandUsageReportServiceAPI
}

func NewCostandUsageReportServiceActivities(client costandusagereportserviceiface.CostandUsageReportServiceAPI) *CostandUsageReportServiceActivities {
    return &CostandUsageReportServiceActivities{client: client}
}


func (a *CostandUsageReportServiceActivities) DeleteReportDefinition(input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
    return a.client.DeleteReportDefinition(input)
}



func (a *CostandUsageReportServiceActivities) DescribeReportDefinitions(input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
    return a.client.DescribeReportDefinitions(input)
}



func (a *CostandUsageReportServiceActivities) ModifyReportDefinition(input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
    return a.client.ModifyReportDefinition(input)
}



func (a *CostandUsageReportServiceActivities) PutReportDefinition(input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error) {
    return a.client.PutReportDefinition(input)
}

