package mysql

import (
	"context"

	"github.com/go-sql-driver/mysql"

	"github.com/andys920605/meme-coin/internal/domain/model/meme_coin"
	"github.com/andys920605/meme-coin/pkg/database"
	"github.com/andys920605/meme-coin/pkg/errors"
	"github.com/andys920605/meme-coin/pkg/mysqlx"
)

type MemeCoinDao struct {
	client *mysqlx.Client
	txm    database.TransactionManager
}

func NewMemeCoinDao(client *mysqlx.Client, txm database.TransactionManager) *MemeCoinDao {
	return &MemeCoinDao{
		client: client,
		txm:    txm,
	}
}

func (d *MemeCoinDao) Create(ctx context.Context, memeCoin *meme_coin.MemeCoin) error {
	dto := &memeCoinDao{
		ID:              memeCoin.ID.Int64(),
		Name:            memeCoin.Name,
		Description:     memeCoin.Description,
		PopularityScore: memeCoin.PopularityScore.Value(),
	}

	tx := d.txm.GetTransaction(ctx)

	err := tx.WithContext(ctx).
		Model(&memeCoinDao{}).
		Create(dto).Error

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errors.NameAlreadyExists.Wrap(err, "name already exists")
		}
		return errors.Wrap(err, "tx create")
	}

	return nil
}

func (d *MemeCoinDao) Update(ctx context.Context, memeCoin *meme_coin.MemeCoin) error {
	tx := d.txm.GetTransaction(ctx)

	updateData := map[string]interface{}{}
	fields := memeCoin.UpdatedFields()
	if fields.Description() != nil {
		updateData["description"] = *fields.Description()
	}
	if fields.PopularityScore() != nil {
		updateData["popularity_score"] = fields.PopularityScore().Value()
	}

	err := tx.WithContext(ctx).
		Model(&memeCoinDao{}).
		Where("id = ?", memeCoin.ID.Int64()).
		Updates(updateData).Error
	if err != nil {
		return errors.Wrap(err, "tx update")
	}

	return nil
}

func (d *MemeCoinDao) Delete(ctx context.Context, memeCoin *meme_coin.MemeCoin) error {
	tx := d.txm.GetTransaction(ctx)
	err := tx.WithContext(ctx).
		Where("id = ?", memeCoin.ID.Int64()).
		Delete(&memeCoinDao{}).Error
	if err != nil {
		return errors.Wrap(err, "tx delete")
	}

	return nil
}

func (d *MemeCoinDao) GetByID(ctx context.Context, id meme_coin.ID) (*meme_coin.MemeCoin, error) {
	tx := d.txm.GetTransaction(ctx)
	var memeCoin *memeCoinDao
	err := tx.WithContext(ctx).
		Model(&memeCoinDao{}).
		Where("id = ?", id.Int64()).
		First(&memeCoin).Error
	if err != nil {
		return nil, errors.Wrap(err, "tx get by id")
	}
	aggregate, err := memeCoin.ToAggregate()
	if err != nil {
		return nil, errors.Wrap(err, "to aggregate")
	}

	return aggregate, nil
}
