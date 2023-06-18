package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"

	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger

	Services *Services
}

type Services struct {
	Alert  AlertService
	APM    ApplicationService
	Plugin PluginService
}

func initServices(apiClient *newrelic.NewRelic) *Services {
	return &Services{
		Alert:  &apiClient.Alerts,
		APM:    &apiClient.APM,
		Plugin: &apiClient.Plugins,
	}
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "newrelic"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	token := getApiTokenFromEnv()
	configuration := newrelic.ConfigPersonalAPIKey(token)
	apiClient, err := newrelic.New(configuration)

	if err != nil {
		return nil, fmt.Errorf("failed to create newrelic client: %w", err)
	}

	client := Client{
		logger:   logger,
		Services: initServices(apiClient),
	}

	return &client, nil
}

func getApiTokenFromEnv() string {
	os.Setenv("NEW_RELIC_API_KEY", "NRAK-SMYODD87CECC7JNKN8JVKZRM32P")
	apiToken := os.Getenv("NEW_RELIC_API_KEY")

	if apiToken == "" {
		return ""
	}
	return apiToken
}
