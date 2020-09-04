
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/aws/aws-sdk-go/service/translate/translateiface"
)

type TranslateActivities struct {
	client translateiface.TranslateAPI
}

func NewTranslateActivities(client translateiface.TranslateAPI) *TranslateActivities {
    return &TranslateActivities{client: client}
}

func (a *TranslateActivities) DeleteTerminology(input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
    return a.client.DeleteTerminology(input)
}

func (a *TranslateActivities) DescribeTextTranslationJob(input *translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error) {
    return a.client.DescribeTextTranslationJob(input)
}

func (a *TranslateActivities) GetTerminology(input *translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error) {
    return a.client.GetTerminology(input)
}

func (a *TranslateActivities) ImportTerminology(input *translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error) {
    return a.client.ImportTerminology(input)
}

func (a *TranslateActivities) ListTerminologies(input *translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error) {
    return a.client.ListTerminologies(input)
}

func (a *TranslateActivities) ListTextTranslationJobs(input *translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error) {
    return a.client.ListTextTranslationJobs(input)
}

func (a *TranslateActivities) StartTextTranslationJob(input *translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error) {
    return a.client.StartTextTranslationJob(input)
}

func (a *TranslateActivities) StopTextTranslationJob(input *translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error) {
    return a.client.StopTextTranslationJob(input)
}

func (a *TranslateActivities) Text(input *translate.TextInput) (*translate.TextOutput, error) {
    return a.client.Text(input)
}
