
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/textract/textractiface"
)

type TextractActivities struct {
    client textractiface.TextractAPI
}

func NewTextractActivities(session *session.Session, config ...*aws.Config) *TextractActivities {
    client := textract.New(session, config...)
    return &TextractActivities{client: client}
}

func (a *TextractActivities) AnalyzeDocument(input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
    return a.client.AnalyzeDocument(input)
}

func (a *TextractActivities) DetectDocumentText(input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error) {
    return a.client.DetectDocumentText(input)
}

func (a *TextractActivities) GetDocumentAnalysis(input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error) {
    return a.client.GetDocumentAnalysis(input)
}

func (a *TextractActivities) GetDocumentTextDetection(input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error) {
    return a.client.GetDocumentTextDetection(input)
}

func (a *TextractActivities) StartDocumentAnalysis(input *textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error) {
    return a.client.StartDocumentAnalysis(input)
}

func (a *TextractActivities) StartDocumentTextDetection(input *textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error) {
    return a.client.StartDocumentTextDetection(input)
}
