package port

import utils "github.com/matheusapostulo/url-shortener/utils/logger"

type LogPublisherService interface {
	PublishLog(log utils.Log) error
}
