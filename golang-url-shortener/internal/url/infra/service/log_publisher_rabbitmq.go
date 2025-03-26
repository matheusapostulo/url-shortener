package service

import (
	"github.com/matheusapostulo/url-shortener/internal/url/port"
	utils "github.com/matheusapostulo/url-shortener/utils/logger"
)

func NewLogPublisherService(conn port.PublisherConnection) *LogPublisherRabbitMQ {
	return &LogPublisherRabbitMQ{
		conn: conn,
	}
}

type LogPublisherRabbitMQ struct {
	conn port.PublisherConnection
}

func (l *LogPublisherRabbitMQ) PublishLog(log utils.Log) error {
	err := l.conn.PublisherConfig("logs")
	if err != nil {
		return err
	}

	byteLog, err := log.ConvertLogToByte()
	if err != nil {
		return err
	}

	err = l.conn.PublishMsg(byteLog)
	if err != nil {
		return err
	}

	return nil
}
