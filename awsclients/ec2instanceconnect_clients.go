// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"go.temporal.io/sdk/workflow"
)

type EC2InstanceConnectClient interface {
	SendSSHPublicKey(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error)
	SendSSHPublicKeyAsync(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) *Ec2instanceconnectSendSSHPublicKeyResult
}

type EC2InstanceConnectStub struct{}

func NewEC2InstanceConnectStub() EC2InstanceConnectClient {
	return &EC2InstanceConnectStub{}
}

type Ec2instanceconnectSendSSHPublicKeyResult struct {
	Result workflow.Future
}

func (r *Ec2instanceconnectSendSSHPublicKeyResult) Get(ctx workflow.Context) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	var output ec2instanceconnect.SendSSHPublicKeyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *EC2InstanceConnectStub) SendSSHPublicKey(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	var output ec2instanceconnect.SendSSHPublicKeyOutput
	err := workflow.ExecuteActivity(ctx, "EC2InstanceConnect.SendSSHPublicKey", input).Get(ctx, &output)
	return &output, err
}

func (a *EC2InstanceConnectStub) SendSSHPublicKeyAsync(ctx workflow.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) *Ec2instanceconnectSendSSHPublicKeyResult {
	future := workflow.ExecuteActivity(ctx, "EC2InstanceConnect.SendSSHPublicKey", input)
	return &Ec2instanceconnectSendSSHPublicKeyResult{Result: future}
}
