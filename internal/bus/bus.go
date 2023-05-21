package bus

import partybus "github.com/nextlinux/gopartybus"

var publisher partybus.Publisher
var active bool

func SetPublisher(p partybus.Publisher) {
	publisher = p
	if p != nil {
		active = true
	}
}

func Publish(event partybus.Event) {
	if active {
		publisher.Publish(event)
	}
}
