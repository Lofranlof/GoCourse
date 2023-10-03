//go:build !change
// +build !change

package coverme_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"gitlab.com/manytask/itmo-go/public/coverme"
)

// min coverage: .,subpkg 70%

func TestSum(t *testing.T) {
	require.Equal(t, int64(2), coverme.Sum(1, 1))
}
