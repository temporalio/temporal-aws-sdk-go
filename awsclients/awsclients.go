package awsclients

import (
	"go.temporal.io/sdk/workflow"
)

type VoidFuture struct {
	Result workflow.Future
}

func (r *VoidFuture) Get(ctx workflow.Context) error {
	return r.Result.Get(ctx, nil)
}
