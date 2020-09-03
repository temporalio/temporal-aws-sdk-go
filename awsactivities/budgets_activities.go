package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/budgets/budgetsiface"
)

type BudgetsActivities struct {
	client budgetsiface.BudgetsAPI
}

func NewBudgetsActivities(client budgetsiface.BudgetsAPI) *BudgetsActivities {
    return &BudgetsActivities{client: client}
}

func (a *BudgetsActivities) CreateBudget(input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error) {
    return a.client.CreateBudget(input)
}

func (a *BudgetsActivities) CreateNotification(input *budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error) {
    return a.client.CreateNotification(input)
}

func (a *BudgetsActivities) CreateSubscriber(input *budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error) {
    return a.client.CreateSubscriber(input)
}

func (a *BudgetsActivities) DeleteBudget(input *budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error) {
    return a.client.DeleteBudget(input)
}

func (a *BudgetsActivities) DeleteNotification(input *budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error) {
    return a.client.DeleteNotification(input)
}

func (a *BudgetsActivities) DeleteSubscriber(input *budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error) {
    return a.client.DeleteSubscriber(input)
}

func (a *BudgetsActivities) DescribeBudget(input *budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error) {
    return a.client.DescribeBudget(input)
}

func (a *BudgetsActivities) DescribeBudgetPerformanceHistory(input *budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error) {
    return a.client.DescribeBudgetPerformanceHistory(input)
}

func (a *BudgetsActivities) DescribeBudgets(input *budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error) {
    return a.client.DescribeBudgets(input)
}

func (a *BudgetsActivities) DescribeNotificationsForBudget(input *budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error) {
    return a.client.DescribeNotificationsForBudget(input)
}

func (a *BudgetsActivities) DescribeSubscribersForNotification(input *budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error) {
    return a.client.DescribeSubscribersForNotification(input)
}

func (a *BudgetsActivities) UpdateBudget(input *budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error) {
    return a.client.UpdateBudget(input)
}

func (a *BudgetsActivities) UpdateNotification(input *budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error) {
    return a.client.UpdateNotification(input)
}

func (a *BudgetsActivities) UpdateSubscriber(input *budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error) {
    return a.client.UpdateSubscriber(input)
}
