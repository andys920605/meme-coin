package meme_coin

import "github.com/andys920605/meme-coin/pkg/errors"

type PopularityScore struct {
	value int
}

func NewPopularityScore() PopularityScore {
	return PopularityScore{value: 0}
}

func ParsePopularityScore(value int) (PopularityScore, error) {
	if value < 0 {
		return PopularityScore{}, errors.New("popularity score cannot be negative")
	}
	return PopularityScore{value: value}, nil
}

func (p PopularityScore) Increment() *PopularityScore {
	return &PopularityScore{value: p.value + 1}
}

func (p PopularityScore) Value() int {
	return p.value
}
