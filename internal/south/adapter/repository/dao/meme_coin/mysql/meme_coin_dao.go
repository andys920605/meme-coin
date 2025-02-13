package mysql

import (
	"context"

	"github.com/andys920605/meme-coin/pkg/mysqlx"
)

type MemeCoinDao struct {
	client *mysqlx.Client
}

func NewMemeCoinDao(client *mysqlx.Client) *MemeCoinDao {
	return &MemeCoinDao{
		client: client,
	}
}

func (d *MemeCoinDao) Create(ctx context.Context) error {
	return nil
}
