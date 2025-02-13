package service

import (
	"context"

	"github.com/andys920605/meme-coin/internal/north/message"
	"github.com/andys920605/meme-coin/internal/south/port/repository"
	"github.com/andys920605/meme-coin/pkg/errors"
	"github.com/andys920605/meme-coin/pkg/logging"
)

type MemeCoinDomainService struct {
	logging            *logging.Logging
	memeCoinRepository repository.MemeCoinRepository
}

func NewMemeCoinDomainService(
	logging *logging.Logging,
	memeCoinRepository repository.MemeCoinRepository,
) *MemeCoinDomainService {
	return &MemeCoinDomainService{
		logging:            logging,
		memeCoinRepository: memeCoinRepository,
	}
}

func (s *MemeCoinDomainService) Create(ctx context.Context, cmd message.CreateMemeCoinCommand) error {
	if err := s.memeCoinRepository.Save(ctx); err != nil {
		return errors.Wrap(err, "save")
	}

	return nil
}
