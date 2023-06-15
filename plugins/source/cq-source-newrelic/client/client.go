package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"

	"github.com/newrelic/newrelic-client-go/v2/newrelic"
	"github.com/newrelic/newrelic-client-go/v2/pkg/alerts"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger

	NRServices NewRelicServices
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "newrelic"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec
	os.Setenv("NEW_RELIC_API_KEY", "NRAK-SMYODD87CECC7JNKN8JVKZRM32P")

	token := getApiTokenFromEnv()

	configuration := newrelic.ConfigPersonalAPIKey(token)
	apiClient, err := newrelic.New(configuration)

	if err != nil {
		return nil, fmt.Errorf("failed to create newrelic client: %w", err)
	}

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	valor := NewNewRelicServices

	fmt.Println("pluginSpec: %v+\n", &valor)

	policies, err := apiClient.Alerts.ListPolicies(&alerts.ListPoliciesParams{
		Name: "Bewiz",
	})

	if err != nil {
		fmt.Printf("err: %v+\n", policies)
	}

	fmt.Println("The value is: %v\n", policies)

	client := Client{
		logger: logger,
		// NRServices:,
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
