//go:build !change
// +build !change

package poorcoverage_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/manytask/itmo-go/public/poorcoverage"
)

// min coverage: . 100%

func TestSum(t *testing.T) {
	require.Equal(t, int64(2), poorcoverage.Sum(1, 1))
}
