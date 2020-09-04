
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/aws/aws-sdk-go/service/rdsdataservice/rdsdataserviceiface"
)

type RDSDataServiceActivities struct {
	client rdsdataserviceiface.RDSDataServiceAPI
}

func NewRDSDataServiceActivities(client rdsdataserviceiface.RDSDataServiceAPI) *RDSDataServiceActivities {
    return &RDSDataServiceActivities{client: client}
}

func (a *RDSDataServiceActivities) BatchExecuteStatement(input *rdsdataservice.BatchExecuteStatementInput) (*rdsdataservice.BatchExecuteStatementOutput, error) {
    return a.client.BatchExecuteStatement(input)
}

func (a *RDSDataServiceActivities) BeginTransaction(input *rdsdataservice.BeginTransactionInput) (*rdsdataservice.BeginTransactionOutput, error) {
    return a.client.BeginTransaction(input)
}

func (a *RDSDataServiceActivities) CommitTransaction(input *rdsdataservice.CommitTransactionInput) (*rdsdataservice.CommitTransactionOutput, error) {
    return a.client.CommitTransaction(input)
}

func (a *RDSDataServiceActivities) ExecuteSql(input *rdsdataservice.ExecuteSqlInput) (*rdsdataservice.ExecuteSqlOutput, error) {
    return a.client.ExecuteSql(input)
}

func (a *RDSDataServiceActivities) ExecuteStatement(input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error) {
    return a.client.ExecuteStatement(input)
}

func (a *RDSDataServiceActivities) RollbackTransaction(input *rdsdataservice.RollbackTransactionInput) (*rdsdataservice.RollbackTransactionOutput, error) {
    return a.client.RollbackTransaction(input)
}
