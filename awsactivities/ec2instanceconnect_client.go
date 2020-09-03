package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"go.temporal.io/sdk/workflow"
)

type EC2InstanceConnectClient interface {
    SendSSHPublicKey(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error)
    SendSSHPublicKeyAsync(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) *Ec2instanceconnectSendSSHPublicKeyResult
}
type Ec2instanceconnectSendSSHPublicKeyResult struct {
	Result workflow.Future
}

func (r *Ec2instanceconnectSendSSHPublicKeyResult) Get(ctx workflow.Context) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
    var output ec2instanceconnect.SendSSHPublicKeyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type EC2InstanceConnectStub struct {
    activities EC2InstanceConnectClient
}

func NewEC2InstanceConnectStub() EC2InstanceConnectClient {
    return &EC2InstanceConnectStub{}
}

func (a *EC2InstanceConnectStub) SendSSHPublicKey(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
    var output ec2instanceconnect.SendSSHPublicKeyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendSSHPublicKey, input).Get(ctx, &output)
    return &output, err
}
