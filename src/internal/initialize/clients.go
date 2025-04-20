package initialize

import (
	"context"
	"github.com/ctfloyd/hazelmere-api/src/pkg/client"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_client"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
)

func InitializeHazelmereClient(config *hz_config.Config, logger hz_logger.Logger) *client.Hazelmere {
	hz, err := client.NewHazelmere(
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
		client.HazelmereConfig{
			Token:              config.ValueOrPanic("clients.hazelmere.token"),
			CallingApplication: "hazelmere-bot",
		},
	)
	if err != nil {
		panic(err)
	}
	return hz
}

func InitializeHazelmereClientResilient(config *hz_config.Config, logger hz_logger.Logger) *client.Hazelmere {
	hz, err := client.NewHazelmere(
		hz_client.NewHttpClient(
			hz_client.HttpClientConfig{
				Host:           config.ValueOrPanic("clients.hazelmere.host"),
				TimeoutMs:      config.IntValueOrPanic("clients.hazelmere.resilient.timeout"),
				Retries:        config.IntValueOrPanic("clients.hazelmere.resilient.retries"),
				RetryWaitMs:    config.IntValueOrPanic("clients.hazelmere.resilient.retryWaitMs"),
				RetryMaxWaitMs: config.IntValueOrPanic("clients.hazelmere.resilient.retryMaxWaitMs"),
			},
			func(msg string) { logger.Error(context.Background(), msg) },
		),
		client.HazelmereConfig{
			Token:              config.ValueOrPanic("clients.hazelmere.token"),
			CallingApplication: "hazelmere-bot",
		},
	)
	if err != nil {
		panic(err)
	}
	return hz
}
