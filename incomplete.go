package meta

import "encoding/json"

// Incomplete allows partially parsing unstructured data (like json) to access its metadata, before parsing the payload
type Incomplete struct {
	Metadata *EventMeta      `json:"metadata"`
	Data     json.RawMessage `json:"data,omitempty"`
}

var _ EventMetaAccessor = &Incomplete{}

// GetEventMeta to implement the accessor interface
func (u *Incomplete) GetEventMeta() Event { return u.Metadata }
