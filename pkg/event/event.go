package event

import (
	"github.com/nextlinux/gopartybus"
)

const (
	PullDockerImage partybus.EventType = "pull-docker-image-event"
	FetchImage      partybus.EventType = "fetch-image-event"
	ReadImage       partybus.EventType = "read-image-event"
	ReadLayer       partybus.EventType = "read-layer-event"
)
