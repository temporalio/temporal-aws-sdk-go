package awsclients

import (
	"go.temporal.io/sdk/workflow"
)

type VoidFuture struct {
	Future workflow.Future
}

func (r *VoidFuture) Get(ctx workflow.Context) error {
	return r.Future.Get(ctx, nil)
}

func NewVoidFuture(future workflow.Future) *VoidFuture {
	return &VoidFuture{Future: future}
}
