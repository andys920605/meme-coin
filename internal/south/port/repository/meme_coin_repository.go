package repository

import (
	"context"

	"github.com/andys920605/meme-coin/internal/domain/model/meme_coin"
)

//go:generate mockgen -destination=../../../mock/imeme_coin_mock_repository.go -package=mock github.com/andys920605/meme-coin/internal/south/port/repository MemeCoinRepository
type MemeCoinRepository interface {
	Save(ctx context.Context, memeCoin *meme_coin.MemeCoin) error
	Delete(ctx context.Context, memeCoin *meme_coin.MemeCoin) error
	GetByID(ctx context.Context, id meme_coin.ID) (*meme_coin.MemeCoin, error)
}
