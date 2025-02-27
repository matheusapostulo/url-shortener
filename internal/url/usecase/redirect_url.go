package usecase

import (
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
)

type RedirectURLInputDto struct {
	ShortURL string `json:"short_url"`
}

type RedirectURLOutputDto struct {
	LongURL string `json:"long_url"`
}

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

func (r *RedirectURLUsecase) Execute(input RedirectURLInputDto) (RedirectURLOutputDto, error) {
	url, _ := r.cacheRepository.Get(input.ShortURL)
	if !url.IsEmpty() {
		return RedirectURLOutputDto{
			LongURL: url.LongURL,
		}, nil
	}

	url, err := r.urlRepository.FindByShortURL(input.ShortURL)
	if url.IsEmpty() {
		return RedirectURLOutputDto{}, domain.ErrURLNotFound
	}
	if err != nil {
		return RedirectURLOutputDto{}, domain.ErrInternalServerError
	}

	err = r.cacheRepository.Set(url)
	if err != nil {
		return RedirectURLOutputDto{}, domain.ErrInternalServerError
	}

	return RedirectURLOutputDto{
		LongURL: url.LongURL,
	}, nil
}
