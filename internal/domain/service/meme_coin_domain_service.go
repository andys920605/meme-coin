package service

import (
	"context"

	"github.com/andys920605/meme-coin/internal/domain/model/meme_coin"
	"github.com/andys920605/meme-coin/internal/north/message"
	"github.com/andys920605/meme-coin/internal/south/port/repository"
	"github.com/andys920605/meme-coin/pkg/database"
	"github.com/andys920605/meme-coin/pkg/errors"
	"github.com/andys920605/meme-coin/pkg/logging"
)

type MemeCoinDomainService struct {
	logging            *logging.Logging
	memeCoinRepository repository.MemeCoinRepository
	transactionManager database.TransactionManager
}

func NewMemeCoinDomainService(
	logging *logging.Logging,
	memeCoinRepository repository.MemeCoinRepository,
	txm database.TransactionManager,
) *MemeCoinDomainService {
	return &MemeCoinDomainService{
		logging:            logging,
		memeCoinRepository: memeCoinRepository,
		transactionManager: txm,
	}
}

func (s *MemeCoinDomainService) CreateMemeCoin(ctx context.Context, cmd message.CreateMemeCoinCommand) (*meme_coin.MemeCoin, error) {
	var memeCoin *meme_coin.MemeCoin
	err := s.transactionManager.Execute(ctx, func(txCtx context.Context) error {
		memeCoin = meme_coin.NewMemeCoin(cmd.Name, cmd.Description)
		if err := s.memeCoinRepository.Save(ctx, memeCoin); err != nil {
			return errors.Wrap(err, "save")
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "transaction manager execute")
	}

	return memeCoin, nil
}

func (s *MemeCoinDomainService) GetMemeCoin(ctx context.Context, query message.GetMemeCoinQuery) (*meme_coin.MemeCoin, error) {
	id, err := meme_coin.ParseID(query.ID)
	if err != nil {
		return nil, errors.Wrap(err, "parse id")
	}
	memeCoin, err := s.memeCoinRepository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get by id")
	}

	return memeCoin, nil
}

func (s *MemeCoinDomainService) UpdateMemeCoin(ctx context.Context, cmd message.UpdateMemeCoinCommand) error {
	id, err := meme_coin.ParseID(cmd.ID)
	if err != nil {
		return errors.Wrap(err, "parse id")
	}
	err = s.transactionManager.Execute(ctx, func(txCtx context.Context) error {
		memeCoin, err := s.memeCoinRepository.GetByID(ctx, id)
		if err != nil {
			return errors.Wrap(err, "get by id")
		}

		memeCoin.UpdateDescription(cmd.Description)

		if err := s.memeCoinRepository.Save(ctx, memeCoin); err != nil {
			return errors.Wrap(err, "save")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction manager execute")
	}

	return nil
}

func (s *MemeCoinDomainService) DeleteMemeCoin(ctx context.Context, cmd message.DeleteMemeCoinCommand) error {
	id, err := meme_coin.ParseID(cmd.ID)
	if err != nil {
		return errors.Wrap(err, "parse id")
	}
	err = s.transactionManager.Execute(ctx, func(txCtx context.Context) error {
		memeCoin, err := s.memeCoinRepository.GetByID(ctx, id)
		if err != nil {
			return errors.Wrap(err, "get by id")
		}

		if err := s.memeCoinRepository.Delete(ctx, memeCoin); err != nil {
			return errors.Wrap(err, "save")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction manager execute")
	}

	return nil
}

func (s *MemeCoinDomainService) PokeMemeCoin(ctx context.Context, cmd message.PokeMemeCoinCommand) error {
	id, err := meme_coin.ParseID(cmd.ID)
	if err != nil {
		return errors.Wrap(err, "parse id")
	}
	err = s.transactionManager.Execute(ctx, func(txCtx context.Context) error {
		memeCoin, err := s.memeCoinRepository.GetByID(ctx, id)
		if err != nil {
			return errors.Wrap(err, "get by id")
		}
		memeCoin.Poke()

		if err := s.memeCoinRepository.Save(ctx, memeCoin); err != nil {
			return errors.Wrap(err, "save")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction manager execute")
	}

	return nil
}
