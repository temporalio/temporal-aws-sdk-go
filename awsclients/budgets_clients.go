package awsclients

import (
	"github.com/aws/aws-sdk-go/service/budgets"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type BudgetsClient interface {
    CreateBudget(ctx workflow.Context, input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error)
    CreateBudgetAsync(ctx workflow.Context, input *budgets.CreateBudgetInput) *BudgetsCreateBudgetResult

    CreateNotification(ctx workflow.Context, input *budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error)
    CreateNotificationAsync(ctx workflow.Context, input *budgets.CreateNotificationInput) *BudgetsCreateNotificationResult

    CreateSubscriber(ctx workflow.Context, input *budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error)
    CreateSubscriberAsync(ctx workflow.Context, input *budgets.CreateSubscriberInput) *BudgetsCreateSubscriberResult

    DeleteBudget(ctx workflow.Context, input *budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error)
    DeleteBudgetAsync(ctx workflow.Context, input *budgets.DeleteBudgetInput) *BudgetsDeleteBudgetResult

    DeleteNotification(ctx workflow.Context, input *budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error)
    DeleteNotificationAsync(ctx workflow.Context, input *budgets.DeleteNotificationInput) *BudgetsDeleteNotificationResult

    DeleteSubscriber(ctx workflow.Context, input *budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error)
    DeleteSubscriberAsync(ctx workflow.Context, input *budgets.DeleteSubscriberInput) *BudgetsDeleteSubscriberResult

    DescribeBudget(ctx workflow.Context, input *budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error)
    DescribeBudgetAsync(ctx workflow.Context, input *budgets.DescribeBudgetInput) *BudgetsDescribeBudgetResult

    DescribeBudgetPerformanceHistory(ctx workflow.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error)
    DescribeBudgetPerformanceHistoryAsync(ctx workflow.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) *BudgetsDescribeBudgetPerformanceHistoryResult

    DescribeBudgets(ctx workflow.Context, input *budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error)
    DescribeBudgetsAsync(ctx workflow.Context, input *budgets.DescribeBudgetsInput) *BudgetsDescribeBudgetsResult

    DescribeNotificationsForBudget(ctx workflow.Context, input *budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error)
    DescribeNotificationsForBudgetAsync(ctx workflow.Context, input *budgets.DescribeNotificationsForBudgetInput) *BudgetsDescribeNotificationsForBudgetResult

    DescribeSubscribersForNotification(ctx workflow.Context, input *budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error)
    DescribeSubscribersForNotificationAsync(ctx workflow.Context, input *budgets.DescribeSubscribersForNotificationInput) *BudgetsDescribeSubscribersForNotificationResult

    UpdateBudget(ctx workflow.Context, input *budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error)
    UpdateBudgetAsync(ctx workflow.Context, input *budgets.UpdateBudgetInput) *BudgetsUpdateBudgetResult

    UpdateNotification(ctx workflow.Context, input *budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error)
    UpdateNotificationAsync(ctx workflow.Context, input *budgets.UpdateNotificationInput) *BudgetsUpdateNotificationResult

    UpdateSubscriber(ctx workflow.Context, input *budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error)
    UpdateSubscriberAsync(ctx workflow.Context, input *budgets.UpdateSubscriberInput) *BudgetsUpdateSubscriberResult
}
type BudgetsCreateBudgetResult struct {
	Result workflow.Future
}

func (r *BudgetsCreateBudgetResult) Get(ctx workflow.Context) (*budgets.CreateBudgetOutput, error) {
    var output budgets.CreateBudgetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsCreateNotificationResult struct {
	Result workflow.Future
}

func (r *BudgetsCreateNotificationResult) Get(ctx workflow.Context) (*budgets.CreateNotificationOutput, error) {
    var output budgets.CreateNotificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsCreateSubscriberResult struct {
	Result workflow.Future
}

func (r *BudgetsCreateSubscriberResult) Get(ctx workflow.Context) (*budgets.CreateSubscriberOutput, error) {
    var output budgets.CreateSubscriberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDeleteBudgetResult struct {
	Result workflow.Future
}

func (r *BudgetsDeleteBudgetResult) Get(ctx workflow.Context) (*budgets.DeleteBudgetOutput, error) {
    var output budgets.DeleteBudgetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDeleteNotificationResult struct {
	Result workflow.Future
}

func (r *BudgetsDeleteNotificationResult) Get(ctx workflow.Context) (*budgets.DeleteNotificationOutput, error) {
    var output budgets.DeleteNotificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDeleteSubscriberResult struct {
	Result workflow.Future
}

func (r *BudgetsDeleteSubscriberResult) Get(ctx workflow.Context) (*budgets.DeleteSubscriberOutput, error) {
    var output budgets.DeleteSubscriberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDescribeBudgetResult struct {
	Result workflow.Future
}

func (r *BudgetsDescribeBudgetResult) Get(ctx workflow.Context) (*budgets.DescribeBudgetOutput, error) {
    var output budgets.DescribeBudgetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDescribeBudgetPerformanceHistoryResult struct {
	Result workflow.Future
}

func (r *BudgetsDescribeBudgetPerformanceHistoryResult) Get(ctx workflow.Context) (*budgets.DescribeBudgetPerformanceHistoryOutput, error) {
    var output budgets.DescribeBudgetPerformanceHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDescribeBudgetsResult struct {
	Result workflow.Future
}

func (r *BudgetsDescribeBudgetsResult) Get(ctx workflow.Context) (*budgets.DescribeBudgetsOutput, error) {
    var output budgets.DescribeBudgetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDescribeNotificationsForBudgetResult struct {
	Result workflow.Future
}

func (r *BudgetsDescribeNotificationsForBudgetResult) Get(ctx workflow.Context) (*budgets.DescribeNotificationsForBudgetOutput, error) {
    var output budgets.DescribeNotificationsForBudgetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsDescribeSubscribersForNotificationResult struct {
	Result workflow.Future
}

func (r *BudgetsDescribeSubscribersForNotificationResult) Get(ctx workflow.Context) (*budgets.DescribeSubscribersForNotificationOutput, error) {
    var output budgets.DescribeSubscribersForNotificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsUpdateBudgetResult struct {
	Result workflow.Future
}

func (r *BudgetsUpdateBudgetResult) Get(ctx workflow.Context) (*budgets.UpdateBudgetOutput, error) {
    var output budgets.UpdateBudgetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsUpdateNotificationResult struct {
	Result workflow.Future
}

func (r *BudgetsUpdateNotificationResult) Get(ctx workflow.Context) (*budgets.UpdateNotificationOutput, error) {
    var output budgets.UpdateNotificationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BudgetsUpdateSubscriberResult struct {
	Result workflow.Future
}

func (r *BudgetsUpdateSubscriberResult) Get(ctx workflow.Context) (*budgets.UpdateSubscriberOutput, error) {
    var output budgets.UpdateSubscriberOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type BudgetsStub struct {
    activities awsactivities.BudgetsActivities
}

func NewBudgetsStub() BudgetsClient {
    return &BudgetsStub{}
}
func (a *BudgetsStub) CreateBudget(ctx workflow.Context, input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error) {
    var output budgets.CreateBudgetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBudget, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) CreateBudgetAsync(ctx workflow.Context, input *budgets.CreateBudgetInput) *BudgetsCreateBudgetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBudget, input)
    return &BudgetsCreateBudgetResult{Result: future}
}
func (a *BudgetsStub) CreateNotification(ctx workflow.Context, input *budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error) {
    var output budgets.CreateNotificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNotification, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) CreateNotificationAsync(ctx workflow.Context, input *budgets.CreateNotificationInput) *BudgetsCreateNotificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateNotification, input)
    return &BudgetsCreateNotificationResult{Result: future}
}
func (a *BudgetsStub) CreateSubscriber(ctx workflow.Context, input *budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error) {
    var output budgets.CreateSubscriberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSubscriber, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) CreateSubscriberAsync(ctx workflow.Context, input *budgets.CreateSubscriberInput) *BudgetsCreateSubscriberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSubscriber, input)
    return &BudgetsCreateSubscriberResult{Result: future}
}
func (a *BudgetsStub) DeleteBudget(ctx workflow.Context, input *budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error) {
    var output budgets.DeleteBudgetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBudget, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DeleteBudgetAsync(ctx workflow.Context, input *budgets.DeleteBudgetInput) *BudgetsDeleteBudgetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBudget, input)
    return &BudgetsDeleteBudgetResult{Result: future}
}
func (a *BudgetsStub) DeleteNotification(ctx workflow.Context, input *budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error) {
    var output budgets.DeleteNotificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNotification, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DeleteNotificationAsync(ctx workflow.Context, input *budgets.DeleteNotificationInput) *BudgetsDeleteNotificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteNotification, input)
    return &BudgetsDeleteNotificationResult{Result: future}
}
func (a *BudgetsStub) DeleteSubscriber(ctx workflow.Context, input *budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error) {
    var output budgets.DeleteSubscriberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSubscriber, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DeleteSubscriberAsync(ctx workflow.Context, input *budgets.DeleteSubscriberInput) *BudgetsDeleteSubscriberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSubscriber, input)
    return &BudgetsDeleteSubscriberResult{Result: future}
}
func (a *BudgetsStub) DescribeBudget(ctx workflow.Context, input *budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error) {
    var output budgets.DescribeBudgetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBudget, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DescribeBudgetAsync(ctx workflow.Context, input *budgets.DescribeBudgetInput) *BudgetsDescribeBudgetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBudget, input)
    return &BudgetsDescribeBudgetResult{Result: future}
}
func (a *BudgetsStub) DescribeBudgetPerformanceHistory(ctx workflow.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error) {
    var output budgets.DescribeBudgetPerformanceHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBudgetPerformanceHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DescribeBudgetPerformanceHistoryAsync(ctx workflow.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) *BudgetsDescribeBudgetPerformanceHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBudgetPerformanceHistory, input)
    return &BudgetsDescribeBudgetPerformanceHistoryResult{Result: future}
}
func (a *BudgetsStub) DescribeBudgets(ctx workflow.Context, input *budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error) {
    var output budgets.DescribeBudgetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBudgets, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DescribeBudgetsAsync(ctx workflow.Context, input *budgets.DescribeBudgetsInput) *BudgetsDescribeBudgetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBudgets, input)
    return &BudgetsDescribeBudgetsResult{Result: future}
}
func (a *BudgetsStub) DescribeNotificationsForBudget(ctx workflow.Context, input *budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error) {
    var output budgets.DescribeNotificationsForBudgetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNotificationsForBudget, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DescribeNotificationsForBudgetAsync(ctx workflow.Context, input *budgets.DescribeNotificationsForBudgetInput) *BudgetsDescribeNotificationsForBudgetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeNotificationsForBudget, input)
    return &BudgetsDescribeNotificationsForBudgetResult{Result: future}
}
func (a *BudgetsStub) DescribeSubscribersForNotification(ctx workflow.Context, input *budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error) {
    var output budgets.DescribeSubscribersForNotificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSubscribersForNotification, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) DescribeSubscribersForNotificationAsync(ctx workflow.Context, input *budgets.DescribeSubscribersForNotificationInput) *BudgetsDescribeSubscribersForNotificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSubscribersForNotification, input)
    return &BudgetsDescribeSubscribersForNotificationResult{Result: future}
}
func (a *BudgetsStub) UpdateBudget(ctx workflow.Context, input *budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error) {
    var output budgets.UpdateBudgetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBudget, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) UpdateBudgetAsync(ctx workflow.Context, input *budgets.UpdateBudgetInput) *BudgetsUpdateBudgetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBudget, input)
    return &BudgetsUpdateBudgetResult{Result: future}
}
func (a *BudgetsStub) UpdateNotification(ctx workflow.Context, input *budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error) {
    var output budgets.UpdateNotificationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateNotification, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) UpdateNotificationAsync(ctx workflow.Context, input *budgets.UpdateNotificationInput) *BudgetsUpdateNotificationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateNotification, input)
    return &BudgetsUpdateNotificationResult{Result: future}
}
func (a *BudgetsStub) UpdateSubscriber(ctx workflow.Context, input *budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error) {
    var output budgets.UpdateSubscriberOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSubscriber, input).Get(ctx, &output)
    return &output, err
}

func (a *BudgetsStub) UpdateSubscriberAsync(ctx workflow.Context, input *budgets.UpdateSubscriberInput) *BudgetsUpdateSubscriberResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSubscriber, input)
    return &BudgetsUpdateSubscriberResult{Result: future}
}
