
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
)

type CloudSearchDomainActivities struct {
	client cloudsearchdomainiface.CloudSearchDomainAPI
}

func NewCloudSearchDomainActivities(session *session.Session, config... *aws.Config) *CloudSearchDomainActivities {
    client := cloudsearchdomain.New(session, config...)
    return &CloudSearchDomainActivities{client: client}
}

func (a *CloudSearchDomainActivities) Search(input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
    return a.client.Search(input)
}

func (a *CloudSearchDomainActivities) Suggest(input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
    return a.client.Suggest(input)
}

func (a *CloudSearchDomainActivities) UploadDocuments(input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
    return a.client.UploadDocuments(input)
}
