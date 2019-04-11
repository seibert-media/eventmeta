package meta

import (
	"encoding/json"
	"time"
)

// EventMeta is metadata that all internal events must have
type EventMeta struct {
	ID      string           `json:"id,omitempty"`
	Kind    GroupVersionKind `json:"-"`
	Context *Context         `json:"context,omitempty"`
	Created time.Time        `json:"created,omitempty"`
}

// NewEventMeta with kind returns a initialized metadata object with the creation timestamp set
func NewEventMeta(kind GroupVersionKind) *EventMeta {
	return &EventMeta{
		Kind:    kind,
		Created: time.Now().UTC(),
	}
}

func (m *EventMeta) GetEventMeta() Event                     { return m }
func (m *EventMeta) GetID() string                           { return m.ID }
func (m *EventMeta) SetID(to string)                         { m.ID = to }
func (m *EventMeta) GetGroupVersionKind() GroupVersionKind   { return m.Kind }
func (m *EventMeta) SetGroupVersionKind(to GroupVersionKind) { m.Kind = to }
func (m *EventMeta) GetContext() *Context                    { return m.Context }
func (m *EventMeta) SetContext(to *Context)                  { m.Context = to }

// GetCreated returns the creation timestamp of an event in UTC
func (m *EventMeta) GetCreated() time.Time { return m.Created.UTC() }

// SetCreated sets the creation timestamp of an event in UTC
func (m *EventMeta) SetCreated(to time.Time) { m.Created = to.UTC() }

// MarshalJSON implements custom marshaling to encode the objects GVK into two fields, apiVersion and a GroupKind
// It encodes meta timestamps as UTC
func (m *EventMeta) MarshalJSON() ([]byte, error) {
	type Alias EventMeta
	type meta struct {
		APIVersion string `json:"apiVersion"`
		GroupKind  string `json:"kind"`
		*Alias
	}

	obj := meta{
		APIVersion: m.Kind.Version,
		GroupKind:  m.Kind.GroupKind().String(),
		Alias:      (*Alias)(m),
	}
	obj.Alias.Created = obj.Alias.Created.UTC()

	return json.Marshal(obj)
}

// UnmarshalJSON implements custom unmarshalling to decode the objects apiVersion and kind into a full GroupVersionKind
// It decodes meta timestamps as UTC
func (m *EventMeta) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	type Alias EventMeta
	type meta struct {
		APIVersion string `json:"apiVersion"`
		GroupKind  string `json:"kind"`
		*Alias
	}

	obj := &meta{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(b, obj); err != nil {
		return err
	}

	m.Kind = ParseGroupKind(obj.GroupKind).WithVersion(obj.APIVersion)
	m.Created = m.Created.UTC()

	return nil
}
