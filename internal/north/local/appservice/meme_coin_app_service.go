package appservice

import (
	"context"

	"github.com/andys920605/meme-coin/internal/domain/service"
	"github.com/andys920605/meme-coin/internal/north/message"
	"github.com/andys920605/meme-coin/pkg/errors"
	"github.com/andys920605/meme-coin/pkg/logging"
)

type MemeCoinAppService struct {
	logging               *logging.Logging
	memeCoinDomainService *service.MemeCoinDomainService
}

func NewMemeCoinAppService(
	logging *logging.Logging,
	memeCoinDomainService *service.MemeCoinDomainService,
) *MemeCoinAppService {
	return &MemeCoinAppService{
		logging:               logging,
		memeCoinDomainService: memeCoinDomainService,
	}
}

func (s *MemeCoinAppService) Create(ctx context.Context, cmd message.CreateMemeCoinCommand) error {
	if err := s.memeCoinDomainService.Create(ctx, cmd); err != nil {
		return errors.Wrap(err, "create")
	}
	return nil
}
