package meme_coin

import (
	"context"

	"golang.org/x/sync/singleflight"

	"github.com/andys920605/meme-coin/internal/domain/model/meme_coin"
	"github.com/andys920605/meme-coin/internal/south/adapter/repository/dao/meme_coin/mysql"
	"github.com/andys920605/meme-coin/internal/south/port/repository"
	"github.com/andys920605/meme-coin/pkg/errors"
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

func (r *MemeCoinRepository) Save(ctx context.Context, memeCoin *meme_coin.MemeCoin) error {
	if memeCoin.IsNew() {
		if err := r.memeCoinDao.Create(ctx, memeCoin); err != nil {
			return errors.Wrap(err, "create")
		}
	} else {
		if err := r.memeCoinDao.Update(ctx, memeCoin); err != nil {
			return errors.Wrap(err, "create")
		}
	}
	// todo: base on CDC memeCoin.IsDomainEventsNotEmpty() can trigger domain events
	return nil
}

func (r *MemeCoinRepository) Delete(ctx context.Context, memeCoin *meme_coin.MemeCoin) error {
	return r.memeCoinDao.Delete(ctx, memeCoin)
}

func (r *MemeCoinRepository) GetByID(ctx context.Context, id meme_coin.ID) (*meme_coin.MemeCoin, error) {
	return r.memeCoinDao.GetByID(ctx, id)
}
