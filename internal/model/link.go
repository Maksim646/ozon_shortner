package model

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

const (
	Https = "https"
	Http  = "https"
)

var (
	ErrLinkNotFound = errors.New("original link not found")
)

type Link struct {
	ID int64 `db:"id"`

	OriginalLink  string        `db:"original_link"`
	ShorterLinkID sql.NullInt64 `db:"short_link"`
	CreatedAt     time.Time     `db:"created_at"`
}

type ILinkRepository interface {
	CreateLink(ctx context.Context, originalLink string, shortLink string) error

	GetOriginalLinkByShortLink(ctx context.Context, shorterLink string) (string, error)
	GetShortLinkByOriginalLink(ctx context.Context, originalLink string) (string, error)
}

type ILinkUsecase interface {
	CreateLink(ctx context.Context, originalLink string, shortLink string) error

	GetOriginalLinkByShortLink(ctx context.Context, shorterLink string) (string, error)
	GetShortLinkByOriginalLink(ctx context.Context, originalLink string) (string, error)
}
