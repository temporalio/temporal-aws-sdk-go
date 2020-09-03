package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/outposts"
	"github.com/aws/aws-sdk-go/service/outposts/outpostsiface"
)

type OutpostsActivities struct {
	client outpostsiface.OutpostsAPI
}

func NewOutpostsActivities(client outpostsiface.OutpostsAPI) *OutpostsActivities {
    return &OutpostsActivities{client: client}
}

func (a *OutpostsActivities) CreateOutpost(input *outposts.CreateOutpostInput) (*outposts.CreateOutpostOutput, error) {
    return a.client.CreateOutpost(input)
}

func (a *OutpostsActivities) DeleteOutpost(input *outposts.DeleteOutpostInput) (*outposts.DeleteOutpostOutput, error) {
    return a.client.DeleteOutpost(input)
}

func (a *OutpostsActivities) DeleteSite(input *outposts.DeleteSiteInput) (*outposts.DeleteSiteOutput, error) {
    return a.client.DeleteSite(input)
}

func (a *OutpostsActivities) GetOutpost(input *outposts.GetOutpostInput) (*outposts.GetOutpostOutput, error) {
    return a.client.GetOutpost(input)
}

func (a *OutpostsActivities) GetOutpostInstanceTypes(input *outposts.GetOutpostInstanceTypesInput) (*outposts.GetOutpostInstanceTypesOutput, error) {
    return a.client.GetOutpostInstanceTypes(input)
}

func (a *OutpostsActivities) ListOutposts(input *outposts.ListOutpostsInput) (*outposts.ListOutpostsOutput, error) {
    return a.client.ListOutposts(input)
}

func (a *OutpostsActivities) ListSites(input *outposts.ListSitesInput) (*outposts.ListSitesOutput, error) {
    return a.client.ListSites(input)
}
