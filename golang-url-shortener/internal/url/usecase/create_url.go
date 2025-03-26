package usecase

import (
	"log"
	"time"

	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
	utils "github.com/matheusapostulo/url-shortener/utils/logger"
)

func NewCreateURLUsecase(
	urlRp port.URLRepository,
	cacheRp port.CacheRepository,
	shortenerSv port.URLShortener,
	logPublisherSv port.LogPublisherService,
) *CreateURLUsecase {
	return &CreateURLUsecase{
		urlRepository:       urlRp,
		cacheRepository:     cacheRp,
		ShortenerService:    shortenerSv,
		logPublisherService: logPublisherSv,
	}
}

type CreateURLUsecase struct {
	urlRepository       port.URLRepository
	cacheRepository     port.CacheRepository
	ShortenerService    port.URLShortener
	logPublisherService port.LogPublisherService
}

func (c *CreateURLUsecase) Execute(input port.CreateURLInputDto) (port.CreateURLOutputDto, error) {
	url, _ := c.urlRepository.FindByLongURL(input.LongURL)
	if !url.IsEmpty() {
		err := c.cacheRepository.Set(url)

		if err != nil {
			return port.CreateURLOutputDto{}, err
		}

		c.logPublisherService.PublishLog(
			utils.Log{
				Timestamp: time.Now().Format(time.RFC3339),
				Level:     "INFO",
				Message:   "URL obtained from cache",
				Context:   "URL cached",
			},
		)

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
		errRepository := c.logPublisherService.PublishLog(
			utils.Log{
				Timestamp: time.Now().Format(time.RFC3339),
				Level:     "ERROR",
				Message:   "Error while creating URL",
				Context:   "Error URL created",
			},
		)
		log.Println(errRepository)
		return port.CreateURLOutputDto{}, err
	}

	err = c.cacheRepository.Set(url)
	if err != nil {
		return port.CreateURLOutputDto{}, err
	}

	c.logPublisherService.PublishLog(
		utils.Log{
			Timestamp: time.Now().Format(time.RFC3339),
			Level:     "INFO",
			Message:   "URL created with success",
			Context:   "URL created",
		},
	)

	return port.CreateURLOutputDto{
		ShortURL: (*url).ShortURL,
	}, nil
}
