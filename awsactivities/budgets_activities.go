
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/budgets/budgetsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type BudgetsActivities struct {
    client budgetsiface.BudgetsAPI
}

func NewBudgetsActivities(session *session.Session, config ...*aws.Config) *BudgetsActivities {
    client := budgets.New(session, config...)
    return &BudgetsActivities{client: client}
}

func (a *BudgetsActivities) CreateBudget(ctx context.Context, input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error) {
    return a.client.CreateBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) CreateNotification(ctx context.Context, input *budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error) {
    return a.client.CreateNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) CreateSubscriber(ctx context.Context, input *budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error) {
    return a.client.CreateSubscriberWithContext(ctx, input)
}

func (a *BudgetsActivities) DeleteBudget(ctx context.Context, input *budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error) {
    return a.client.DeleteBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) DeleteNotification(ctx context.Context, input *budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error) {
    return a.client.DeleteNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) DeleteSubscriber(ctx context.Context, input *budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error) {
    return a.client.DeleteSubscriberWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeBudget(ctx context.Context, input *budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error) {
    return a.client.DescribeBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeBudgetPerformanceHistory(ctx context.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error) {
    return a.client.DescribeBudgetPerformanceHistoryWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeBudgets(ctx context.Context, input *budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error) {
    return a.client.DescribeBudgetsWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeNotificationsForBudget(ctx context.Context, input *budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error) {
    return a.client.DescribeNotificationsForBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeSubscribersForNotification(ctx context.Context, input *budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error) {
    return a.client.DescribeSubscribersForNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) UpdateBudget(ctx context.Context, input *budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error) {
    return a.client.UpdateBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) UpdateNotification(ctx context.Context, input *budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error) {
    return a.client.UpdateNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) UpdateSubscriber(ctx context.Context, input *budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error) {
    return a.client.UpdateSubscriberWithContext(ctx, input)
}
