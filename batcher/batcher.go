//go:build !solution

package batcher

import "gitlab.com/manytask/itmo-go/private/batcher/slow"

type Batcher struct{}

func NewBatcher(v *slow.Value) *Batcher {
	return nil
}
