package handler

import (
	"github.com/Maksim646/ozon_shortner/internal/api/definition"
	"github.com/Maksim646/ozon_shortner/internal/api/server/restapi/api"
	"github.com/Maksim646/ozon_shortner/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) CreateShortLinkHandler(req api.CreateShortLinkParams) middleware.Responder {
	ctx := req.HTTPRequest.Context()

	originalLink := *req.OriginalLink.OriginalLink
	zap.L().Info("CreateShortLinkHandler: handling request", zap.String("original_url", originalLink))

	if originalLink == "" {
		zap.L().Warn("CreateShortLinkHandler: original_url is required")
		err := &definition.Error{Message: useful.StrPtr("original_url is required")}
		return api.NewCreateShortLinkBadRequest().WithPayload(err)
	}

	shortLink, err := h.GenerateShortLink(ctx, originalLink)
	if err != nil {
		zap.L().Error("CreateShortLinkHandler: error generating short link", zap.Error(err))
		return api.NewCreateShortLinkInternalServerError().WithPayload(&definition.Error{
			Message: useful.StrPtr("Failed to generate short link, do not satisfy schema http/https"),
		})
	}

	response := &definition.ShortLink{ShortLink: useful.StrPtr(shortLink)}

	zap.L().Info("CreateShortLinkHandler: short link created successfully",
		zap.String("original_url", originalLink),
		zap.String("short_url", shortLink))

	return api.NewCreateShortLinkOK().WithPayload(response)
}

func (h *Handler) GetOriginalLinkHandler(req api.GetOriginalLinkParams) middleware.Responder {
	ctx := req.HTTPRequest.Context()

	zap.L().Info("GetOriginalLinkHandler: handling request", zap.String("short_link", req.ShortLink))

	originalLink, err := h.linkUsecase.GetOriginalLinkByShortLink(ctx, req.ShortLink)
	if err != nil {
		zap.L().Error("GetOriginalLinkHandler: error fetching original link", zap.Error(err), zap.String("short_link", req.ShortLink))
		apiErr := &definition.Error{Message: useful.StrPtr("original_url not found")}
		return api.NewGetOriginalLinkBadRequest().WithPayload(apiErr)
	}

	response := &definition.OriginalLink{OriginalLink: useful.StrPtr(originalLink)}

	zap.L().Info("GetOriginalLinkHandler: original link retrieved successfully",
		zap.String("short_link", req.ShortLink),
		zap.String("original_url", originalLink))

	return api.NewGetOriginalLinkOK().WithPayload(response)
}
