package meta

// EventMetaAccessor for accessing event metadata
type EventMetaAccessor interface {
	GetEventMeta() Event
}

// Event lets you work with event metadata
type Event interface {
	GetID() string
	SetID(string)
	GetGroupVersionKind() GroupVersionKind
	SetGroupVersionKind(GroupVersionKind)
	GetContext() *Context
	SetContext(*Context)
}

// Kind provides access to an objects GroupVersionKind
type Kind interface {
	Kind() GroupVersionKind
}
