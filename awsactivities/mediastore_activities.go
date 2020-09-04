package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastore/mediastoreiface"
)

type MediaStoreActivities struct {
	client mediastoreiface.MediaStoreAPI
}

func NewMediaStoreActivities(client mediastoreiface.MediaStoreAPI) *MediaStoreActivities {
    return &MediaStoreActivities{client: client}
}


func (a *MediaStoreActivities) CreateContainer(input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
    return a.client.CreateContainer(input)
}



func (a *MediaStoreActivities) DeleteContainer(input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error) {
    return a.client.DeleteContainer(input)
}



func (a *MediaStoreActivities) DeleteContainerPolicy(input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error) {
    return a.client.DeleteContainerPolicy(input)
}



func (a *MediaStoreActivities) DeleteCorsPolicy(input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error) {
    return a.client.DeleteCorsPolicy(input)
}



func (a *MediaStoreActivities) DeleteLifecyclePolicy(input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error) {
    return a.client.DeleteLifecyclePolicy(input)
}



func (a *MediaStoreActivities) DeleteMetricPolicy(input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error) {
    return a.client.DeleteMetricPolicy(input)
}



func (a *MediaStoreActivities) DescribeContainer(input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error) {
    return a.client.DescribeContainer(input)
}



func (a *MediaStoreActivities) GetContainerPolicy(input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error) {
    return a.client.GetContainerPolicy(input)
}



func (a *MediaStoreActivities) GetCorsPolicy(input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error) {
    return a.client.GetCorsPolicy(input)
}



func (a *MediaStoreActivities) GetLifecyclePolicy(input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error) {
    return a.client.GetLifecyclePolicy(input)
}



func (a *MediaStoreActivities) GetMetricPolicy(input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error) {
    return a.client.GetMetricPolicy(input)
}



func (a *MediaStoreActivities) ListContainers(input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error) {
    return a.client.ListContainers(input)
}



func (a *MediaStoreActivities) ListTagsForResource(input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *MediaStoreActivities) PutContainerPolicy(input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error) {
    return a.client.PutContainerPolicy(input)
}



func (a *MediaStoreActivities) PutCorsPolicy(input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error) {
    return a.client.PutCorsPolicy(input)
}



func (a *MediaStoreActivities) PutLifecyclePolicy(input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error) {
    return a.client.PutLifecyclePolicy(input)
}



func (a *MediaStoreActivities) PutMetricPolicy(input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error) {
    return a.client.PutMetricPolicy(input)
}



func (a *MediaStoreActivities) StartAccessLogging(input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error) {
    return a.client.StartAccessLogging(input)
}



func (a *MediaStoreActivities) StopAccessLogging(input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error) {
    return a.client.StopAccessLogging(input)
}



func (a *MediaStoreActivities) TagResource(input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *MediaStoreActivities) UntagResource(input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

