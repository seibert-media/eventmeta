package meta

import (
	"github.com/pkg/errors"
)

// ErrNotEvent is returned when an object does not implement the Event interfaces
var ErrNotEvent = errors.New("object does not implement the Event interfaces")

// Accessor takes an arbitrary event pointer and returns meta.Interface.
// obj must be a pointer to an API type. An error is returned if the minimum
// required fields are missing. Fields that are not required return the default
// value and are a no-op if set.
func Accessor(obj interface{}) (Event, error) {
	switch t := obj.(type) {
	case Event:
		return t, nil
	case EventMetaAccessor:
		if m := t.GetEventMeta(); m != nil {
			return m, nil
		}
		return nil, ErrNotEvent
	default:
		return nil, ErrNotEvent
	}
}

// Matcher takes two arbitrary event pointers, gets their accessors and verifies they match in kind
func Matcher(a, b interface{}) (bool, error) {
	ameta, err := Accessor(a)
	if err != nil {
		return false, errors.Wrap(err, "accessing object a")
	}

	bmeta, err := Accessor(b)
	if err != nil {
		return false, errors.Wrap(err, "accessing object b")
	}

	if ameta.GetGroupVersionKind() != bmeta.GetGroupVersionKind() {
		return false, nil
	}

	return true, nil
}
