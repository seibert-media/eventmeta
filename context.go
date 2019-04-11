package meta

// Context of an event to store various information passed on to other services
type Context struct {
	Customer string  `json:"customer,omitempty"`
	Domain   string  `json:"domain,omitempty"`
	Session  string  `json:"session,omitempty"`
	Notify   *Notify `json:"notify,omitempty"`
}

// Notify stores information about which customer to notify about the event
type Notify struct {
	Customer string `json:"customer,omitempty"`
}
