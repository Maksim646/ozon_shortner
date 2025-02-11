package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/Maksim646/ozon_shortner/internal/database/postgresql"
	"github.com/Maksim646/ozon_shortner/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type LinkRepository struct {
	sqalxConn       sqalx.Node
	cache           sync.Map
	cacheTTL        time.Duration
	preparedStmts   map[string]*sql.Stmt
	preparedStmtsMu sync.RWMutex
}

func New(sqalxConn sqalx.Node, cacheTTL time.Duration) model.ILinkRepository {
	return &LinkRepository{
		sqalxConn:     sqalxConn,
		cache:         sync.Map{},
		cacheTTL:      cacheTTL,
		preparedStmts: make(map[string]*sql.Stmt),
	}
}

func (r *LinkRepository) prepareStmt(ctx context.Context, queryName string, query string) (*sql.Stmt, error) {
	r.preparedStmtsMu.RLock()
	stmt, ok := r.preparedStmts[queryName]
	r.preparedStmtsMu.RUnlock()

	if ok {
		return stmt, nil
	}

	stmt, err := r.sqalxConn.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement %s: %w", queryName, err)
	}

	r.preparedStmtsMu.Lock()
	defer r.preparedStmtsMu.Unlock()
	r.preparedStmts[queryName] = stmt

	return stmt, nil
}

func (r *LinkRepository) CreateLink(ctx context.Context, originalLink string, shortLink string) error {
	tx, err := r.sqalxConn.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

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
		return fmt.Errorf("failed to build insert query: %w", err)
	}
	zap.L().Debug(postgresql.BuildQuery(query, params))

	stmt, err := r.prepareStmt(ctx, "createLink", query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, params...)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %w", err)
	}

	r.cache.Store(shortLink, originalLink)

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *LinkRepository) GetOriginalLinkByShortLink(ctx context.Context, shorterLink string) (string, error) {
	if originalLink, ok := r.cache.Load(shorterLink); ok {
		zap.L().Debug("Cache hit for shortLink: " + shorterLink)
		return originalLink.(string), nil
	}

	var originalLink string
	query, params, err := postgresql.Builder.Select(
		"links.original_link",
	).
		From("links").
		Where(sq.Eq{"links.short_link": shorterLink}).ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build select query: %w", err)
	}

	slog.Debug(postgresql.BuildQuery(query, params))

	stmt, err := r.prepareStmt(ctx, "getOriginalLinkByShortLink", query)
	if err != nil {
		return "", err
	}

	err = stmt.QueryRowContext(ctx, params...).Scan(&originalLink)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("link not found for shortLink %s: %w", shorterLink, err)
		}
		return "", fmt.Errorf("failed to execute select query: %w", err)
	}

	r.cache.Store(shorterLink, originalLink)

	return originalLink, nil
}

func (r *LinkRepository) GetShortLinkByOriginalLink(ctx context.Context, originalLink string) (string, error) {
	var shortLink string
	query, params, err := postgresql.Builder.Select(
		"links.short_link",
	).
		From("links").
		Where(sq.Eq{"links.original_link": originalLink}).ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build select query: %w", err)
	}

	slog.Debug(postgresql.BuildQuery(query, params))

	stmt, err := r.prepareStmt(ctx, "getShortLinkByOriginalLink", query)
	if err != nil {
		return "", err
	}
	err = stmt.QueryRowContext(ctx, params...).Scan(&shortLink)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("link not found for originalLink %s: %w", originalLink, err)
		}
		return "", fmt.Errorf("failed to execute select query: %w", err)
	}

	return shortLink, nil
}
