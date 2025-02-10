package handler

import (
	"context"
	"math/rand"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_httpTransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Maksim646/ozon_shortner/internal/api/client"
	"github.com/Maksim646/ozon_shortner/internal/api/client/operations"
	"github.com/Maksim646/ozon_shortner/internal/database/postgresql/pgtest"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	_linkRepo "github.com/Maksim646/ozon_shortner/internal/domain/link_postgresql/repository/postgresql"
	_linkUsecase "github.com/Maksim646/ozon_shortner/internal/domain/link_postgresql/usecase"
)

const (
	shortLinkLength = 10
	allowedChars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	maxRetries      = 5
)

type Suite struct {
	pgtest.Suite
	ctx     context.Context
	cancel  context.CancelFunc
	api     operations.ClientService
	server  *httptest.Server
	handler *Handler
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	zap.ReplaceGlobals(zaptest.NewLogger(s.T(), zaptest.WrapOptions(zap.AddCaller())))
	s.Suite.SetupSuite("./../../../..")

	linkRepo := _linkRepo.New(s.DB())
	linkUsecase := _linkUsecase.New(linkRepo)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	s.handler = New(
		linkUsecase,
		rng,
		maxRetries,
		shortLinkLength,
		"test",
	)

	s.server = httptest.NewServer(s.handler)
	transport := _httpTransport.New(strings.TrimPrefix(s.server.URL, "http://"), "", nil)
	s.api = client.New(transport, strfmt.Default).Operations
}

func (s *Suite) SetupTest() {
	zap.ReplaceGlobals(zaptest.NewLogger(s.T(), zaptest.WrapOptions(zap.AddCaller())))
	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.Suite.SetupTest()

	linkRepo := _linkRepo.New(s.DB())
	linkUsecase := _linkUsecase.New(linkRepo)

	s.handler.linkUsecase = linkUsecase

}

func (s *Suite) TearDownTest() {
	zap.ReplaceGlobals(zap.NewNop())
	s.cancel()
	s.Suite.TearDownTest()
}

func (s *Suite) TearDownSuite() {
	s.Suite.TearDownSuite()
	s.server.Close()
}
