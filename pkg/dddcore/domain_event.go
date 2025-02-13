package dddcore

import (
	"strings"

	"github.com/google/uuid"
)

type DomainEvent struct {
	ID   string
	Name DomainEventName
}

func NewDomainEvent(name DomainEventName) DomainEvent {
	return DomainEvent{
		ID:   strings.ReplaceAll(uuid.New().String(), "-", ""),
		Name: name,
	}
}
