package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mq/mqiface"
)

type MQActivities struct {
	client mqiface.MQAPI
}

func NewMQActivities(client mqiface.MQAPI) *MQActivities {
    return &MQActivities{client: client}
}

func (a *MQActivities) CreateBroker(input *mq.CreateBrokerInput) (*mq.CreateBrokerOutput, error) {
    return a.client.CreateBroker(input)
}

func (a *MQActivities) CreateConfiguration(input *mq.CreateConfigurationInput) (*mq.CreateConfigurationOutput, error) {
    return a.client.CreateConfiguration(input)
}

func (a *MQActivities) CreateTags(input *mq.CreateTagsInput) (*mq.CreateTagsOutput, error) {
    return a.client.CreateTags(input)
}

func (a *MQActivities) CreateUser(input *mq.CreateUserInput) (*mq.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *MQActivities) DeleteBroker(input *mq.DeleteBrokerInput) (*mq.DeleteBrokerOutput, error) {
    return a.client.DeleteBroker(input)
}

func (a *MQActivities) DeleteTags(input *mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}

func (a *MQActivities) DeleteUser(input *mq.DeleteUserInput) (*mq.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *MQActivities) DescribeBroker(input *mq.DescribeBrokerInput) (*mq.DescribeBrokerOutput, error) {
    return a.client.DescribeBroker(input)
}

func (a *MQActivities) DescribeBrokerEngineTypes(input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error) {
    return a.client.DescribeBrokerEngineTypes(input)
}

func (a *MQActivities) DescribeBrokerInstanceOptions(input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
    return a.client.DescribeBrokerInstanceOptions(input)
}

func (a *MQActivities) DescribeConfiguration(input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error) {
    return a.client.DescribeConfiguration(input)
}

func (a *MQActivities) DescribeConfigurationRevision(input *mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionOutput, error) {
    return a.client.DescribeConfigurationRevision(input)
}

func (a *MQActivities) DescribeUser(input *mq.DescribeUserInput) (*mq.DescribeUserOutput, error) {
    return a.client.DescribeUser(input)
}

func (a *MQActivities) ListBrokers(input *mq.ListBrokersInput) (*mq.ListBrokersOutput, error) {
    return a.client.ListBrokers(input)
}

func (a *MQActivities) ListConfigurationRevisions(input *mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsOutput, error) {
    return a.client.ListConfigurationRevisions(input)
}

func (a *MQActivities) ListConfigurations(input *mq.ListConfigurationsInput) (*mq.ListConfigurationsOutput, error) {
    return a.client.ListConfigurations(input)
}

func (a *MQActivities) ListTags(input *mq.ListTagsInput) (*mq.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *MQActivities) ListUsers(input *mq.ListUsersInput) (*mq.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *MQActivities) RebootBroker(input *mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error) {
    return a.client.RebootBroker(input)
}

func (a *MQActivities) UpdateBroker(input *mq.UpdateBrokerInput) (*mq.UpdateBrokerOutput, error) {
    return a.client.UpdateBroker(input)
}

func (a *MQActivities) UpdateConfiguration(input *mq.UpdateConfigurationInput) (*mq.UpdateConfigurationOutput, error) {
    return a.client.UpdateConfiguration(input)
}

func (a *MQActivities) UpdateUser(input *mq.UpdateUserInput) (*mq.UpdateUserOutput, error) {
    return a.client.UpdateUser(input)
}
