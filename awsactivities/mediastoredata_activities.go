package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"github.com/aws/aws-sdk-go/service/mediastoredata/mediastoredataiface"
)

type MediaStoreDataActivities struct {
	client mediastoredataiface.MediaStoreDataAPI
}

func NewMediaStoreDataActivities(client mediastoredataiface.MediaStoreDataAPI) *MediaStoreDataActivities {
    return &MediaStoreDataActivities{client: client}
}

func (a *MediaStoreDataActivities) DeleteObject(input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error) {
    return a.client.DeleteObject(input)
}

func (a *MediaStoreDataActivities) DescribeObject(input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error) {
    return a.client.DescribeObject(input)
}

func (a *MediaStoreDataActivities) GetObject(input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error) {
    return a.client.GetObject(input)
}

func (a *MediaStoreDataActivities) ListItems(input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error) {
    return a.client.ListItems(input)
}

func (a *MediaStoreDataActivities) PutObject(input *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error) {
    return a.client.PutObject(input)
}
