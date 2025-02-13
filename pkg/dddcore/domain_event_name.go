package dddcore

type DomainEventName string

func (n DomainEventName) String() string {
	return string(n)
}
