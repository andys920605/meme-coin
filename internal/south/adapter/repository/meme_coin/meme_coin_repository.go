package meme_coin

import (
	"context"

	"golang.org/x/sync/singleflight"

	"github.com/andys920605/meme-coin/internal/south/adapter/repository/dao/meme_coin/mysql"
	"github.com/andys920605/meme-coin/internal/south/port/repository"
)

var _ repository.MemeCoinRepository = (*MemeCoinRepository)(nil)

type MemeCoinRepository struct {
	sg          singleflight.Group
	memeCoinDao *mysql.MemeCoinDao
}

func NewMemeCoinRepository(
	memecoinDao *mysql.MemeCoinDao,
) *MemeCoinRepository {
	return &MemeCoinRepository{
		sg:          singleflight.Group{},
		memeCoinDao: memecoinDao,
	}
}

func (r *MemeCoinRepository) Save(ctx context.Context) error {
	return r.memeCoinDao.Create(ctx)
}
