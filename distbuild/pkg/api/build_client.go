//go:build !solution

package api

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/manytask/itmo-go/private/distbuild/pkg/build"
)

type BuildClient struct{}

func NewBuildClient(l *zap.Logger, endpoint string) *BuildClient {
	panic("implement me")
}

func (c *BuildClient) StartBuild(ctx context.Context, request *BuildRequest) (*BuildStarted, StatusReader, error) {
	panic("implement me")
}

func (c *BuildClient) SignalBuild(ctx context.Context, buildID build.ID, signal *SignalRequest) (*SignalResponse, error) {
	panic("implement me")
}
