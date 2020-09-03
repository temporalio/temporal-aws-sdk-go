package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/honeycode"
	"github.com/aws/aws-sdk-go/service/honeycode/honeycodeiface"
)

type HoneycodeActivities struct {
	client honeycodeiface.HoneycodeAPI
}

func NewHoneycodeActivities(client honeycodeiface.HoneycodeAPI) *HoneycodeActivities {
    return &HoneycodeActivities{client: client}
}

func (a *HoneycodeActivities) GetScreenData(input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
    return a.client.GetScreenData(input)
}

func (a *HoneycodeActivities) InvokeScreenAutomation(input *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error) {
    return a.client.InvokeScreenAutomation(input)
}
