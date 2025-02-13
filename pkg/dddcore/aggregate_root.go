package dddcore

import "github.com/samber/lo"

type AggregateRoot struct {
	isNew  bool
	events []DomainEvent
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		events: make([]DomainEvent, 0),
	}
}

func (r *AggregateRoot) SetNew() *AggregateRoot {
	r.isNew = true
	return r
}

func (r *AggregateRoot) IsNew() bool {
	return r.isNew
}

func (r *AggregateRoot) DomainEvents() []DomainEvent {
	return lo.UniqBy(r.events, func(event DomainEvent) string {
		return event.Name.String()
	})
}

func (r *AggregateRoot) IsDomainEventsNotEmpty() bool {
	return len(r.events) > 0
}

func (r *AggregateRoot) AppendDomainEvent(event ...DomainEvent) *AggregateRoot {
	if len(event) != 0 {
		r.events = append(r.events, event...)
	}
	return r
}
