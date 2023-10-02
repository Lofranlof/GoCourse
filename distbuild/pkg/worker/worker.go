//go:build !solution

package worker

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"gitlab.com/manytask/itmo-go/private/distbuild/pkg/api"
	"gitlab.com/manytask/itmo-go/private/distbuild/pkg/artifact"
	"gitlab.com/manytask/itmo-go/private/distbuild/pkg/filecache"
)

type Worker struct{}

func New(
	workerID api.WorkerID,
	coordinatorEndpoint string,
	log *zap.Logger,
	fileCache *filecache.Cache,
	artifacts *artifact.Cache,
) *Worker {
	panic("implement me")
}

func (w *Worker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (w *Worker) Run(ctx context.Context) error {
	panic("implement me")
}
