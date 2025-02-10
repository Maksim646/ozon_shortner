package usecase

import (
	"context"

	"github.com/Maksim646/ozon_shortner/internal/model"
)

type LinkUsecase struct {
	repo model.ILinkRepository
}

func NewLinkUsecase(repo model.ILinkRepository) *LinkUsecase {
	return &LinkUsecase{
		repo: repo,
	}
}
func (uc *LinkUsecase) CreateLink(ctx context.Context, originalLink string, shortLink string) error {
	return uc.repo.CreateLink(ctx, originalLink, shortLink)
}

func (uc *LinkUsecase) GetOriginalLinkByShortLink(ctx context.Context, shortLink string) (string, error) {
	return uc.repo.GetOriginalLinkByShortLink(ctx, shortLink)
}

func (uc *LinkUsecase) GetShortLinkByOriginalLink(ctx context.Context, originalLink string) (string, error) {
	return uc.repo.GetShortLinkByOriginalLink(ctx, originalLink)
}
