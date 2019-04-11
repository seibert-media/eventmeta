package meta

import (
	"github.com/seibert-media/eventmeta"
)

// MailSent to notify about a mail sending
type MailSent struct {
	Metadata *meta.EventMeta `json:"metadata"`
	Data     *MailSentData   `json:"data"`
}

// Kind is an access helper to implement the Kind interface
func (c *MailSent) Kind() meta.GroupVersionKind {
	if c.Metadata == nil {
		return meta.GroupVersionKind{}
	}
	return c.Metadata.GetGroupVersionKind()
}

// GetEventMeta is an access helper to implement the EventMetaAccessor interface
func (c *MailSent) GetEventMeta() meta.Event {
	return c.Metadata
}

// MailSentKind defines the event GroupVersionKind
var MailSentKind = Mail.WithKind("sent")

// MailSentData contains the actual payload of a MailSent event
type MailSentData struct {
	CustomerID string `json:"customerID"`
	Address    string `json:"address"`
	Payload    string `json:"payload"`
}

// NewMailSent initializes an empty event with the correct metadata
func NewMailSent() *MailSent {
	return &MailSent{
		Metadata: meta.NewEventMeta(MailSentKind),
		Data:     &MailSentData{},
	}
}
