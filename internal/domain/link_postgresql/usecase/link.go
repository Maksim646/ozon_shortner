package usecase

import (
	"context"

	"github.com/Maksim646/ozon_shortner/internal/model"
)

type Usecase struct {
	linkRepository model.ILinkRepository
}

func New(linkRepository model.ILinkRepository) model.ILinkUsecase {
	return &Usecase{
		linkRepository: linkRepository,
	}
}

func (u *Usecase) CreateLink(ctx context.Context, originalLink string, shortLink string) error {
	return u.linkRepository.CreateLink(ctx, originalLink, shortLink)
}

func (u *Usecase) GetOriginalLinkByShortLink(ctx context.Context, shortLink string) (string, error) {
	return u.linkRepository.GetOriginalLinkByShortLink(ctx, shortLink)
}

func (u *Usecase) GetShortLinkByOriginalLink(ctx context.Context, originalLink string) (string, error) {
	return u.linkRepository.GetShortLinkByOriginalLink(ctx, originalLink)
}
