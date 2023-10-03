//go:build !solution

package batcher

import "gitlab.com/manytask/itmo-go/public/batcher/slow"

type Batcher struct{}

func NewBatcher(v *slow.Value) *Batcher {
	return nil
}
