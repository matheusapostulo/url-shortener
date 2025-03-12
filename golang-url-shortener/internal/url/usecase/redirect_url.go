package usecase

import (
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
)

func NewRedirectURLUsecase(cacheRp port.CacheRepository, urlRp port.URLRepository) *RedirectURLUsecase {
	return &RedirectURLUsecase{
		cacheRepository: cacheRp,
		urlRepository:   urlRp,
	}
}

type RedirectURLUsecase struct {
	cacheRepository port.CacheRepository
	urlRepository   port.URLRepository
}

func (r *RedirectURLUsecase) Execute(input port.RedirectURLInputDto) (port.RedirectURLOutputDto, error) {
	url, _ := r.cacheRepository.Get(input.ShortURL)
	if !url.IsEmpty() {
		return port.RedirectURLOutputDto{
			LongURL: url.LongURL,
		}, nil
	}

	url, err := r.urlRepository.FindByShortURL(input.ShortURL)
	if url.IsEmpty() {
		return port.RedirectURLOutputDto{}, domain.ErrURLNotFound
	}
	if err != nil {
		return port.RedirectURLOutputDto{}, domain.ErrInternalServerError
	}

	err = r.cacheRepository.Set(url)
	if err != nil {
		return port.RedirectURLOutputDto{}, domain.ErrInternalServerError
	}

	return port.RedirectURLOutputDto{
		LongURL: url.LongURL,
	}, nil
}
