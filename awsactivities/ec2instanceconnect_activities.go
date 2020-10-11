// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect/ec2instanceconnectiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type EC2InstanceConnectActivities struct {
	client ec2instanceconnectiface.EC2InstanceConnectAPI

	sessionFactory SessionFactory
}

func NewEC2InstanceConnectActivities(sess *session.Session, config ...*aws.Config) *EC2InstanceConnectActivities {
	client := ec2instanceconnect.New(sess, config...)
	return &EC2InstanceConnectActivities{client: client}
}

func NewEC2InstanceConnectActivitiesWithSessionFactory(sessionFactory SessionFactory) *EC2InstanceConnectActivities {
	return &EC2InstanceConnectActivities{sessionFactory: sessionFactory}
}

func (a *EC2InstanceConnectActivities) getClient(ctx context.Context) (ec2instanceconnectiface.EC2InstanceConnectAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return ec2instanceconnect.New(sess), nil
}

func (a *EC2InstanceConnectActivities) SendSSHPublicKey(ctx context.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendSSHPublicKeyWithContext(ctx, input)
}
