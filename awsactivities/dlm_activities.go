package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/dlm/dlmiface"
)

type DLMActivities struct {
	client dlmiface.DLMAPI
}

func NewDLMActivities(client dlmiface.DLMAPI) *DLMActivities {
    return &DLMActivities{client: client}
}


func (a *DLMActivities) CreateLifecyclePolicy(input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error) {
    return a.client.CreateLifecyclePolicy(input)
}



func (a *DLMActivities) DeleteLifecyclePolicy(input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error) {
    return a.client.DeleteLifecyclePolicy(input)
}



func (a *DLMActivities) GetLifecyclePolicies(input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error) {
    return a.client.GetLifecyclePolicies(input)
}



func (a *DLMActivities) GetLifecyclePolicy(input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error) {
    return a.client.GetLifecyclePolicy(input)
}



func (a *DLMActivities) ListTagsForResource(input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *DLMActivities) TagResource(input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *DLMActivities) UntagResource(input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *DLMActivities) UpdateLifecyclePolicy(input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error) {
    return a.client.UpdateLifecyclePolicy(input)
}

