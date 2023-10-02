//go:build !solution

package dist

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"gitlab.com/manytask/itmo-go/private/distbuild/pkg/filecache"
	"gitlab.com/manytask/itmo-go/private/distbuild/pkg/scheduler"
)

type Coordinator struct{}

var defaultConfig = scheduler.Config{
	CacheTimeout: time.Millisecond * 10,
	DepsTimeout:  time.Millisecond * 100,
}

func NewCoordinator(
	log *zap.Logger,
	fileCache *filecache.Cache,
) *Coordinator {
	panic("implement me")
}

func (c *Coordinator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
