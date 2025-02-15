package mysql

import (
	"time"

	"github.com/andys920605/meme-coin/internal/domain/model/meme_coin"
	"github.com/andys920605/meme-coin/pkg/errors"
)

type memeCoinDao struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:false"`
	Name            string    `gorm:"column:name"`
	Description     string    `gorm:"column:description"`
	PopularityScore int       `gorm:"column:popularity_score"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (m *memeCoinDao) TableName() string {
	return "meme_coin.meme_coin"
}

func (m *memeCoinDao) ToAggregate() (*meme_coin.MemeCoin, error) {
	id := meme_coin.ID(m.ID)
	score, err := meme_coin.ParsePopularityScore(m.PopularityScore)
	if err != nil {
		return nil, errors.New("parse popularity score")
	}

	return meme_coin.Rebuild(id, m.Name, m.Description, score, m.CreatedAt, m.UpdatedAt), nil
}
