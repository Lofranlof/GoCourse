//go:build !solution
// +build !solution

package sum

import "gitlab.com/manytask/itmo-go/public/sum/pkg"

func Sum(a, b int64) int64 {
	pkg.F()
	return a + b
}
