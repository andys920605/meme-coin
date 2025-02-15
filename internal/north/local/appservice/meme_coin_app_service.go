package appservice

import (
	"context"
	"time"

	"github.com/andys920605/meme-coin/internal/domain/service"
	"github.com/andys920605/meme-coin/internal/north/message"
	"github.com/andys920605/meme-coin/internal/north/remote/source/handler/response"
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

func (s *MemeCoinAppService) CreateMemeCoin(ctx context.Context, cmd message.CreateMemeCoinCommand) (*response.CreateMemeCoin, error) {
	memeCoin, err := s.memeCoinDomainService.CreateMemeCoin(ctx, cmd)
	if err != nil {
		return nil, errors.Wrap(err, "create meme coin")
	}
	return &response.CreateMemeCoin{
		ID: memeCoin.ID.String(),
	}, nil
}

func (s *MemeCoinAppService) GetMemeCoin(ctx context.Context, query message.GetMemeCoinQuery) (*response.GetMemeCoin, error) {
	memeCoin, err := s.memeCoinDomainService.GetMemeCoin(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get meme coin")
	}
	dto := &response.GetMemeCoin{
		ID:              memeCoin.ID.String(),
		Name:            memeCoin.Name,
		Description:     memeCoin.Description,
		PopularityScore: memeCoin.PopularityScore.Value(),
		CreatedAt:       memeCoin.CreatedAt.Format(time.RFC3339),
	}
	return dto, nil
}

func (s *MemeCoinAppService) UpdateMemeCoin(ctx context.Context, cmd message.UpdateMemeCoinCommand) error {
	if err := s.memeCoinDomainService.UpdateMemeCoin(ctx, cmd); err != nil {
		return errors.Wrap(err, "update meme coin")
	}
	return nil
}

func (s *MemeCoinAppService) DeleteMemeCoin(ctx context.Context, cmd message.DeleteMemeCoinCommand) error {
	if err := s.memeCoinDomainService.DeleteMemeCoin(ctx, cmd); err != nil {
		return errors.Wrap(err, "delete meme coin")
	}
	return nil
}

func (s *MemeCoinAppService) PokeMemeCoin(ctx context.Context, cmd message.PokeMemeCoinCommand) error {
	if err := s.memeCoinDomainService.PokeMemeCoin(ctx, cmd); err != nil {
		return errors.Wrap(err, "poke meme coin")
	}
	return nil
}
