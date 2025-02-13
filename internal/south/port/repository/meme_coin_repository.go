package repository

import (
	"context"
)

//go:generate mockgen -destination=../../../mock/imeme_coin_mock_repository.go -package=mock github.com/andys920605/meme-coin/internal/south/port/repository MemeCoinRepository
type MemeCoinRepository interface {
	Save(ctx context.Context) error
}
