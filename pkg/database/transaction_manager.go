package database

import (
	"context"

	"gorm.io/gorm"
)

//go:generate mockgen -destination=../../internal/mock/itransaction_manager_mock_database.go -package=mock github.com/andys920605/meme-coin/pkg/database TransactionManager
type TransactionManager interface {
	Execute(ctx context.Context, fn TxFunc) error
	GetTransaction(ctx context.Context) *gorm.DB
}

type TxFunc func(ctx context.Context) error
type transactionKeyType struct{}

var transactionKey = transactionKeyType{}

type GormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) TransactionManager {
	return &GormTransactionManager{db: db}
}

func (tm *GormTransactionManager) Execute(ctx context.Context, fn TxFunc) error {
	return tm.db.Transaction(func(tx *gorm.DB) error {
		return fn(context.WithValue(ctx, transactionKey, tx))
	})
}

func (tm *GormTransactionManager) GetTransaction(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(transactionKey).(*gorm.DB); ok {
		return tx
	}
	return tm.db
}
