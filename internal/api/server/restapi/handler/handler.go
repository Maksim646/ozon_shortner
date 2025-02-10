package handler

import (
	"math/rand"
	"net/http"
	"strings"
	"sync"

	//"strings"

	"encoding/json"

	"go.uber.org/zap"

	"github.com/Maksim646/ozon_shortner/internal/api/server/restapi"
	"github.com/Maksim646/ozon_shortner/internal/api/server/restapi/api"
	"github.com/Maksim646/ozon_shortner/internal/model"
	"github.com/go-openapi/loads"
)

type Handler struct {
	linkUsecase model.ILinkUsecase
	rng         *rand.Rand
	mu          sync.Mutex

	maxRetries      int
	shortLinkLength int

	router http.Handler
}

func New(
	linkUsecase model.ILinkUsecase,
	rng *rand.Rand,

	maxRetries int,
	shortLinkLength int,

	version string,
) *Handler {

	withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", version)
	swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	if err != nil {
		panic(err)
	}

	h := &Handler{
		linkUsecase:     linkUsecase,
		rng:             rng,
		maxRetries:      maxRetries,
		shortLinkLength: shortLinkLength,
	}

	zap.L().Error("server http handler request")
	router := api.NewOzonShrtnerBackendServiceAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof

	// LINK
	router.CreateShortLinkHandler = api.CreateShortLinkHandlerFunc(h.CreateShortLinkHandler)
	router.GetOriginalLinkHandler = api.GetOriginalLinkHandlerFunc(h.GetOriginalLinkHandler)

	h.router = router.Serve(nil)
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Received HTTP request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	if h.router == nil {
		zap.L().Error("h.router is nil")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	zap.L().Info("h.router is not nil, processing request")
	h.router.ServeHTTP(w, r)
}
