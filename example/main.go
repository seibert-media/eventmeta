package main

import (
	"encoding/json"
	"fmt"

	"github.com/seibert-media/eventmeta"
	metav0 "github.com/seibert-media/eventmeta/v0"
)

func main() {
	mailSent := metav0.NewMailSent()
	mailSent.Data.CustomerID = "test"

	b, err := json.Marshal(&mailSent)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encoded Event to JSON: %s\n", b)

	// b now holds json that can be sent to PubSub
	// and received on a subscription...

	u := &meta.Incomplete{}
	if err := json.Unmarshal(b, &u); err != nil {
		panic(err)
	}

	if equal, err := meta.Matcher(mailSent, u); err != nil || !equal {
		panic(fmt.Errorf("matching objects failed: %v", err))
	}

	receivedEvent := &metav0.MailSent{}
	if err := json.Unmarshal(b, &receivedEvent); err != nil {
		panic(err)
	}

	fmt.Printf("Decoded Event from JSON: %#v\n", receivedEvent)
	fmt.Printf("Decoded EventMeta from JSON: %#v\n", receivedEvent.Metadata)
	fmt.Printf("Decoded EventData from JSON: %#v\n", receivedEvent.Data)
}

