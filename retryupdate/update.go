//go:build !solution

package retryupdate

import "gitlab.com/manytask/itmo-go/private/retryupdate/kvapi"

func UpdateValue(c kvapi.Client, key string, updateFn func(oldValue *string) (newValue string, err error)) error {
	panic("implement me")
}
