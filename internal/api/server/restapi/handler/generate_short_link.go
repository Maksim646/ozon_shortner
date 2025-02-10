package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"

	"go.uber.org/zap"
)

const ()

func (h *Handler) GenerateShortLink(ctx context.Context, originalLink string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	shortLink, err := h.linkUsecase.GetShortLinkByOriginalLink(ctx, originalLink)
	if err == nil {
		zap.L().Debug("Found existing short link", zap.String("original", originalLink), zap.String("short", shortLink))
		return shortLink, nil
	}
	zap.L().Debug("No existing short link found, generating new one", zap.String("original", originalLink))

	urlScheme, err := h.getScheme(originalLink)
	if err != nil {
		zap.L().Error("Invalid URL, cannot generate short link", zap.String("url", originalLink), zap.Error(err))
		return "", fmt.Errorf("invalid URL, cannot generate short link: %w", err)
	}

	for i := 0; i < h.maxRetries; i++ {
		shortLinkPart, err := h.generateRandomShortLink(urlScheme)
		if err != nil {
			zap.L().Error("Failed to generate random short link", zap.Error(err))
			return "", fmt.Errorf("failed to generate random short link: %w", err)
		}

		shortURL := urlScheme + "://" + shortLinkPart

		_, err = h.linkUsecase.GetOriginalLinkByShortLink(ctx, shortURL)
		if err != nil {
			err = h.linkUsecase.CreateLink(ctx, originalLink, shortURL)
			if err != nil {
				zap.L().Error("Failed to save new link", zap.Error(err), zap.String("original", originalLink), zap.String("short", shortURL))
				return "", fmt.Errorf("failed to create short link: %w", err)
			}
			zap.L().Info("Generated new link", zap.String("original", originalLink), zap.String("short", shortURL))
			return shortURL, nil
		}
		zap.L().Debug("Collision, generating new short link", zap.String("attempted", shortURL), zap.Int("attempt", i+1))
	}

	zap.L().Error("Failed to generate unique short link after multiple retries", zap.String("original", originalLink))
	return "", fmt.Errorf("failed to generate unique short link after %d retries", h.maxRetries)
}

func (h *Handler) generateRandomShortLink(urlScheme string) (string, error) {
	fmt.Println(h.shortLinkLength, len(urlScheme))
	randomBytes := make([]byte, h.shortLinkLength-len(urlScheme)-4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	shortLink := base64.RawURLEncoding.EncodeToString(randomBytes)
	return shortLink, nil
}

func (h *Handler) getScheme(originalLink string) (string, error) {
	parsedURL, err := url.ParseRequestURI(originalLink)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	scheme := parsedURL.Scheme
	if scheme == "" {
		return "", fmt.Errorf("no scheme provided")
	}
	return scheme, nil
}
