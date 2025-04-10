package initialize

import (
	"context"
	"github.com/ctfloyd/hazelmere-api/src/pkg/client"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_client"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
)

func InitializeHazelmereClient(config *hz_config.Config, logger hz_logger.Logger) *client.Hazelmere {
	return client.NewHazelmere(
		hz_client.NewHttpClient(
			hz_client.HttpClientConfig{
				Host:           config.ValueOrPanic("clients.hazelmere.host"),
				TimeoutMs:      config.IntValueOrPanic("clients.hazelmere.timeout"),
				Retries:        config.IntValueOrPanic("clients.hazelmere.retries"),
				RetryWaitMs:    config.IntValueOrPanic("clients.hazelmere.retryWaitMs"),
				RetryMaxWaitMs: config.IntValueOrPanic("clients.hazelmere.retryMaxWaitMs"),
			},
			func(msg string) { logger.Error(context.Background(), msg) },
		),
	)
}
