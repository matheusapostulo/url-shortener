package usecase

import (
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
)

type CreateURLInputDto struct {
	LongURL string
}

type CreateURLOutputDto struct {
	ShortURL string
}

func NewCreateURLUsecase(urlRp port.URLRepository, cacheRp port.CacheRepository, shortenerSv port.URLShortener) *CreateURLUsecase {
	return &CreateURLUsecase{
		urlRepository:    urlRp,
		cacheRepository:  cacheRp,
		ShortenerService: shortenerSv,
	}
}

type CreateURLUsecase struct {
	urlRepository    port.URLRepository
	cacheRepository  port.CacheRepository
	ShortenerService port.URLShortener
}

func (c *CreateURLUsecase) Execute(input CreateURLInputDto) (CreateURLOutputDto, error) {
	url, _ := c.cacheRepository.Get(input.LongURL)
	if url != nil {
		return CreateURLOutputDto{
			ShortURL: url.ShortURL,
		}, nil
	}

	url, err := c.urlRepository.FindByLongURL(input.LongURL)
	if url != nil {
		return CreateURLOutputDto{
			ShortURL: url.ShortURL,
		}, nil
	}
	if err != nil {
		return CreateURLOutputDto{}, err
	}

	newAvailableUrlID, err := c.urlRepository.GetNewAvailableID()
	if err != nil {
		return CreateURLOutputDto{}, err
	}

	shortURL, err := c.ShortenerService.ShortenURL(newAvailableUrlID)
	if err != nil {
		return CreateURLOutputDto{}, err
	}

	url, err = domain.NewURL(newAvailableUrlID, input.LongURL, shortURL)
	if err != nil {
		return CreateURLOutputDto{}, err
	}

	return CreateURLOutputDto{
		ShortURL: url.ShortURL,
	}, nil
}
