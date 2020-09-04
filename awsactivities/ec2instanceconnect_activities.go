package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect/ec2instanceconnectiface"
)

type EC2InstanceConnectActivities struct {
	client ec2instanceconnectiface.EC2InstanceConnectAPI
}

func NewEC2InstanceConnectActivities(client ec2instanceconnectiface.EC2InstanceConnectAPI) *EC2InstanceConnectActivities {
    return &EC2InstanceConnectActivities{client: client}
}


func (a *EC2InstanceConnectActivities) SendSSHPublicKey(input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
    return a.client.SendSSHPublicKey(input)
}

