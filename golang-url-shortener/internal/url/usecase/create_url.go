package usecase

import (
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
)

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

func (c *CreateURLUsecase) Execute(input port.CreateURLInputDto) (port.CreateURLOutputDto, error) {
	url, _ := c.urlRepository.FindByLongURL(input.LongURL)
	if !url.IsEmpty() {
		err := c.cacheRepository.Set(url)

		if err != nil {
			return port.CreateURLOutputDto{}, err
		}

		return port.CreateURLOutputDto{
			ShortURL: url.ShortURL,
		}, nil
	}

	newAvailableUrlID, err := c.urlRepository.GetNewAvailableID()
	if err != nil {
		return port.CreateURLOutputDto{}, err
	}

	shortURL, err := c.ShortenerService.ShortenURL(newAvailableUrlID)
	if err != nil {
		return port.CreateURLOutputDto{}, err
	}

	url, err = domain.NewURL(newAvailableUrlID, input.LongURL, shortURL)
	if err != nil {
		return port.CreateURLOutputDto{}, err
	}

	err = c.urlRepository.Save(url)
	if err != nil {
		return port.CreateURLOutputDto{}, err
	}

	err = c.cacheRepository.Set(url)
	if err != nil {
		return port.CreateURLOutputDto{}, err
	}

	return port.CreateURLOutputDto{
		ShortURL: (*url).ShortURL,
	}, nil
}
