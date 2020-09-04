package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/simpledb"
	"github.com/aws/aws-sdk-go/service/simpledb/simpledbiface"
)

type SimpleDBActivities struct {
	client simpledbiface.SimpleDBAPI
}

func NewSimpleDBActivities(client simpledbiface.SimpleDBAPI) *SimpleDBActivities {
    return &SimpleDBActivities{client: client}
}


func (a *SimpleDBActivities) BatchDeleteAttributes(input *simpledb.BatchDeleteAttributesInput) (*simpledb.BatchDeleteAttributesOutput, error) {
    return a.client.BatchDeleteAttributes(input)
}



func (a *SimpleDBActivities) BatchPutAttributes(input *simpledb.BatchPutAttributesInput) (*simpledb.BatchPutAttributesOutput, error) {
    return a.client.BatchPutAttributes(input)
}



func (a *SimpleDBActivities) CreateDomain(input *simpledb.CreateDomainInput) (*simpledb.CreateDomainOutput, error) {
    return a.client.CreateDomain(input)
}



func (a *SimpleDBActivities) DeleteAttributes(input *simpledb.DeleteAttributesInput) (*simpledb.DeleteAttributesOutput, error) {
    return a.client.DeleteAttributes(input)
}



func (a *SimpleDBActivities) DeleteDomain(input *simpledb.DeleteDomainInput) (*simpledb.DeleteDomainOutput, error) {
    return a.client.DeleteDomain(input)
}



func (a *SimpleDBActivities) DomainMetadata(input *simpledb.DomainMetadataInput) (*simpledb.DomainMetadataOutput, error) {
    return a.client.DomainMetadata(input)
}



func (a *SimpleDBActivities) GetAttributes(input *simpledb.GetAttributesInput) (*simpledb.GetAttributesOutput, error) {
    return a.client.GetAttributes(input)
}



func (a *SimpleDBActivities) ListDomains(input *simpledb.ListDomainsInput) (*simpledb.ListDomainsOutput, error) {
    return a.client.ListDomains(input)
}



func (a *SimpleDBActivities) PutAttributes(input *simpledb.PutAttributesInput) (*simpledb.PutAttributesOutput, error) {
    return a.client.PutAttributes(input)
}



func (a *SimpleDBActivities) Select(input *simpledb.SelectInput) (*simpledb.SelectOutput, error) {
    return a.client.Select(input)
}

