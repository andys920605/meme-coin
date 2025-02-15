package meme_coin

import "github.com/andys920605/meme-coin/pkg/dddcore"

const (
	MemeCoinEventName dddcore.DomainEventName = "boost_created"
)

func NewCreatedMemeCoinEvent() dddcore.DomainEvent {
	return dddcore.NewDomainEvent(MemeCoinEventName)
}
