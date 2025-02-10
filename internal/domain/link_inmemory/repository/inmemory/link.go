package inmemory

import (
	"context"
	"sync"

	"github.com/Maksim646/ozon_shortner/internal/model"
)

// InMemoryLinkRepository implements ILinkRepository in memory.
type InMemoryLinkRepository struct {
	mu      sync.RWMutex
	data    map[string]string // shortLink -> originalLink
	reverse map[string]string // originalLink -> shortLink
}

// NewInMemoryLinkRepository creates a new InMemoryLinkRepository.
func NewInMemoryLinkRepository() *InMemoryLinkRepository {
	return &InMemoryLinkRepository{
		data:    make(map[string]string),
		reverse: make(map[string]string),
	}
}

// CreateLink stores a link in memory.
func (r *InMemoryLinkRepository) CreateLink(ctx context.Context, originalLink string, shortLink string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[shortLink] = originalLink
	r.reverse[originalLink] = shortLink

	return nil
}

// GetOriginalLinkByShortLink retrieves an original link by its short link.
func (r *InMemoryLinkRepository) GetOriginalLinkByShortLink(ctx context.Context, shortLink string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	originalLink, ok := r.data[shortLink]
	if !ok {
		return "", model.ErrLinkNotFound // Assuming this error exists
	}

	return originalLink, nil
}

// GetShortLinkByOriginalLink retrieves a short link by its original link.
func (r *InMemoryLinkRepository) GetShortLinkByOriginalLink(ctx context.Context, originalLink string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	shortLink, ok := r.reverse[originalLink]
	if !ok {
		return "", model.ErrLinkNotFound // Assuming this error exists
	}

	return shortLink, nil
}
