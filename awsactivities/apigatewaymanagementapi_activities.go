
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi/apigatewaymanagementapiiface"
)

type ApiGatewayManagementApiActivities struct {
	client apigatewaymanagementapiiface.ApiGatewayManagementApiAPI
}

func NewApiGatewayManagementApiActivities(session *session.Session, config... *aws.Config) *ApiGatewayManagementApiActivities {
    client := apigatewaymanagementapi.New(session, config...)
    return &ApiGatewayManagementApiActivities{client: client}
}

func (a *ApiGatewayManagementApiActivities) DeleteConnection(input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error) {
    return a.client.DeleteConnection(input)
}

func (a *ApiGatewayManagementApiActivities) GetConnection(input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error) {
    return a.client.GetConnection(input)
}

func (a *ApiGatewayManagementApiActivities) PostToConnection(input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
    return a.client.PostToConnection(input)
}
