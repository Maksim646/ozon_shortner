package handler

import (
	"strings"

	"github.com/Maksim646/ozon_shortner/internal/api/client/operations"
	"github.com/Maksim646/ozon_shortner/internal/api/definition"
	"github.com/Maksim646/ozon_shortner/pkg/useful"
)

func (s *Suite) TestCreateShortLink_NoValidPass() {
	paramsLink := operations.NewCreateShortLinkParams().WithOriginalLink(&definition.OriginalLink{
		OriginalLink: useful.StrPtr("12345"),
	})

	_, err := s.api.CreateShortLink(paramsLink)
	s.Error(err)
}

func (s *Suite) TestCreateShortLink_Http_ValidPass() {
	paramsLink := operations.NewCreateShortLinkParams().WithOriginalLink(&definition.OriginalLink{
		OriginalLink: useful.StrPtr("http://test/func/service.ru"),
	})
	CreateShortLink, err := s.api.CreateShortLink(paramsLink)
	s.NoError(err)
	s.Equal(10, len(*CreateShortLink.Payload.ShortLink))
}

func (s *Suite) TestCreateShortLink_Https_ValidPass() {
	paramsLink := operations.NewCreateShortLinkParams().WithOriginalLink(&definition.OriginalLink{
		OriginalLink: useful.StrPtr("https://test/func/service.ru"),
	})
	CreateShortLink, err := s.api.CreateShortLink(paramsLink)
	s.NoError(err)

	s.Equal(10, len(*CreateShortLink.Payload.ShortLink))

}

func (s *Suite) TestCreateShortLink_EmptyOriginalLink() {
	paramsLink := operations.NewCreateShortLinkParams().WithOriginalLink(&definition.OriginalLink{
		OriginalLink: useful.StrPtr(""),
	})

	_, err := s.api.CreateShortLink(paramsLink)
	s.Error(err)
}

func (s *Suite) TestCreateShortLink_DuplicateOriginalLink() {
	paramsLink := operations.NewCreateShortLinkParams().WithOriginalLink(&definition.OriginalLink{
		OriginalLink: useful.StrPtr("http://test/func/service.ru"),
	})

	CreateShortLink1, err := s.api.CreateShortLink(paramsLink)
	s.NoError(err)
	CreateShortLink2, err := s.api.CreateShortLink(paramsLink)
	s.NoError(err)
	s.Equal(*CreateShortLink1.Payload.ShortLink, *CreateShortLink2.Payload.ShortLink)
}

func (s *Suite) TestCreateShortLink_LongOriginalLink() {
	longURL := strings.Repeat("a", 2049)
	paramsLink := operations.NewCreateShortLinkParams().WithOriginalLink(&definition.OriginalLink{
		OriginalLink: useful.StrPtr("https://" + longURL),
	})

	CreateShortLink, err := s.api.CreateShortLink(paramsLink)
	s.NoError(err)

	s.NotEmpty(*CreateShortLink.Payload.ShortLink)
	s.Len(*CreateShortLink.Payload.ShortLink, 10)
}

func (s *Suite) TestGetOriginalLink_ValidShortLink() {
	paramsLink := operations.NewCreateShortLinkParams().WithOriginalLink(&definition.OriginalLink{
		OriginalLink: useful.StrPtr("http://test/func/service.ru"),
	})
	shortLink, _ := s.api.CreateShortLink(paramsLink)
	shortLinkRequest := operations.NewGetOriginalLinkParams().WithShortLink(*shortLink.Payload.ShortLink)
	GetOriginalLink, err := s.api.GetOriginalLink(shortLinkRequest)
	s.NoError(err)
	s.Equal(*GetOriginalLink.Payload.OriginalLink, *paramsLink.OriginalLink.OriginalLink)
}

func (s *Suite) TestGetOriginalLink_InvalidShortLink() {
	shortLink := "invalidShortLink"

	params := operations.NewGetOriginalLinkParams().WithShortLink(shortLink)
	_, err := s.api.GetOriginalLink(params)

	s.Error(err)
}

func (s *Suite) TestGetOriginalLink_EmptyShortLink() {
	params := operations.NewGetOriginalLinkParams().WithShortLink("")
	_, err := s.api.GetOriginalLink(params)

	s.Error(err)
}
