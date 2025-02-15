package meme_coin

import (
	"time"

	"github.com/andys920605/meme-coin/pkg/dddcore"
)

type MemeCoin struct {
	*dddcore.AggregateRoot
	ID              ID
	Name            string
	Description     string
	PopularityScore PopularityScore
	CreatedAt       time.Time
	UpdatedAt       time.Time

	updatedFields UpdatableFields
}

func NewMemeCoin(name, description string) *MemeCoin {
	root := dddcore.NewAggregateRoot().SetNew()
	// Implementation of event-driven design based on CDC (Change Data Capture)
	root.AppendDomainEvent(
		NewCreatedMemeCoinEvent(),
	)
	return &MemeCoin{
		AggregateRoot:   root,
		ID:              NewID(),
		Name:            name,
		Description:     description,
		PopularityScore: NewPopularityScore(),
	}
}

func Rebuild(
	id ID, name, description string, popularityScore PopularityScore, createdAt, updatedAt time.Time) *MemeCoin {
	return &MemeCoin{
		AggregateRoot:   dddcore.NewAggregateRoot(),
		ID:              id,
		Name:            name,
		Description:     description,
		PopularityScore: popularityScore,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func (c *MemeCoin) UpdatedFields() UpdatableFields {
	return c.updatedFields
}

type UpdatableFields struct {
	description     *string
	popularityScore *PopularityScore
}

func (f *UpdatableFields) Description() *string {
	return f.description
}

func (f *UpdatableFields) PopularityScore() *PopularityScore {
	return f.popularityScore
}

func (m *MemeCoin) UpdateDescription(description string) {
	m.Description = description
	m.updatedFields.description = &description
}

func (m *MemeCoin) Poke() {
	score := m.PopularityScore.Increment()
	m.PopularityScore = *score
	m.updatedFields.popularityScore = score
}
