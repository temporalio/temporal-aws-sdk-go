package internal

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"go.temporal.io/sdk/activity"
	"time"
)

// SetClientToken assigns idempotency token to the same value across retries
func SetClientToken(ctx context.Context, clientToken **string) {
	if clientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		*clientToken = &token
	}
}

// WaitUntilActivity calls function until it returns an error which is not request.WaiterResourceNotReadyErrorCode.
func WaitUntilActivity(ctx context.Context, f func(context.Context, ...request.WaiterOption) error) error {
	// Do not rely on the waiter for retries to support heartbeating
	for {
		err := f(ctx, request.WithWaiterMaxAttempts(1))
		aerr, ok := err.(awserr.Error)
		if !ok || request.WaiterResourceNotReadyErrorCode != aerr.Code() {
			return err
		}
		activity.RecordHeartbeat(ctx)
		//TODO(maxim): Configurable poll interval
		time.Sleep(10 * time.Second)
	}
}
