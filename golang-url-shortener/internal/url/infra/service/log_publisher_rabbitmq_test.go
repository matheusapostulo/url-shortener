package service_test

import (
	"testing"

	"github.com/matheusapostulo/url-shortener/internal/url/infra/service"
	"github.com/matheusapostulo/url-shortener/mocks"
	utils "github.com/matheusapostulo/url-shortener/utils/logger"
	"github.com/stretchr/testify/require"
)

func TestPublishLog(t *testing.T) {
	logData := utils.Log{
		Timestamp: "2021-09-01T00:00:00Z",
		Level:     "INFO",
		Message:   "Log test message",
		Context:   "Log test context",
	}

	byteParam := []byte(`{"timestamp":"2021-09-01T00:00:00Z","level":"INFO","message":"Log test message","context":"Log test context"}`)

	conn := mocks.NewPublisherConnection(t)
	publisherLogRabbitMQ := service.NewLogPublisherService(conn)

	conn.On("PublisherConfig", "logs").Return(nil)
	conn.On("PublishMsg", byteParam).Return(nil)

	err := publisherLogRabbitMQ.PublishLog(logData)

	require.NoError(t, err)

}
