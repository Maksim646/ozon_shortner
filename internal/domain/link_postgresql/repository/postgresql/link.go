package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/Maksim646/ozon_shortner/internal/database/postgresql"
	"github.com/Maksim646/ozon_shortner/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type LinkRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.ILinkRepository {
	return &LinkRepository{sqalxConn: sqalxConn}
}

func (r *LinkRepository) CreateLink(ctx context.Context, originalLink string, shortLink string) error {
	query, params, err := postgresql.Builder.Insert("links").
		Columns(
			"original_link",
			"short_link",
		).
		Values(
			originalLink,
			shortLink,
		).
		ToSql()
	if err != nil {
		return err
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))

	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}

func (r *LinkRepository) GetOriginalLinkByShortLink(ctx context.Context, shorterLink string) (string, error) {
	var originalLink string
	query, params, err := postgresql.Builder.Select(
		"links.original_link",
	).
		From("links").
		Where(sq.Eq{"links.short_link": shorterLink}).ToSql()
	if err != nil {
		return originalLink, err
	}

	slog.Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &originalLink, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return originalLink, err
		}
	}

	return originalLink, err
}

func (r *LinkRepository) GetShortLinkByOriginalLink(ctx context.Context, originalLink string) (string, error) {
	var shortLink string
	query, params, err := postgresql.Builder.Select(
		"links.short_link",
	).
		From("links").
		Where(sq.Eq{"links.original_link": originalLink}).ToSql()
	if err != nil {
		return shortLink, err
	}

	slog.Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &shortLink, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return shortLink, err
		}
	}

	return shortLink, err
}
